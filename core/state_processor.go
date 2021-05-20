// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/permission/core"
	"github.com/ethereum/go-ethereum/private"
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
//
// StateProcessor implements Processor.
type StateProcessor struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
	engine consensus.Engine    // Consensus engine used for block rewards
}

// NewStateProcessor initialises a new StateProcessor.
func NewStateProcessor(config *params.ChainConfig, bc *BlockChain, engine consensus.Engine) *StateProcessor {
	return &StateProcessor{
		config: config,
		bc:     bc,
		engine: engine,
	}
}

// Process processes the state changes according to the Ethereum rules by running
// the transaction messages using the statedb and applying any rewards to both
// the processor (coinbase) and any included uncles.
//
// Process returns the receipts and logs accumulated during the process and
// returns the amount of gas that was used in the process. If any of the
// transactions failed to execute due to insufficient gas it will return an error.
//
// Quorum: Private transactions are handled for the following:
//
// 1. On original single private state (SPS) design
// 2. On multiple private states (MPS) design
// 3. Contract extension callback (p.bc.CheckAndSetPrivateState)
func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, privateStateRepo mps.PrivateStateRepository, cfg vm.Config) (types.Receipts, types.Receipts, []*types.Log, uint64, error) {

	var (
		receipts types.Receipts
		usedGas  = new(uint64)
		header   = block.Header()
		allLogs  []*types.Log
		gp       = new(GasPool).AddGas(block.GasLimit())

		privateReceipts types.Receipts
	)
	// Mutate the block and state according to any hard-fork specs
	if p.config.DAOForkSupport && p.config.DAOForkBlock != nil && p.config.DAOForkBlock.Cmp(block.Number()) == 0 {
		misc.ApplyDAOHardFork(statedb)
	}
	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		mpsReceipt, err := p.handleMPS(i, tx, block, gp, usedGas, cfg, statedb, privateStateRepo)
		if err != nil {
			return nil, nil, nil, 0, err
		}
		// handling transaction in 2 scenarios:
		// 1. For MPS, the target private state being applied would be the EmptyPrivateState.
		//    This must be last to avoid contract address collisions.
		// 2. For orignal SPS design, the target private state is the single private state
		//
		// in both cases, privateStateRepo is responsible to return the appropriate
		// private state for execution and a bool flag to enable the privacy execution
		privateStateDB, err := privateStateRepo.DefaultState()
		if err != nil {
			return nil, nil, nil, 0, err
		}
		privateStateDB.Prepare(tx.Hash(), block.Hash(), i)
		statedb.Prepare(tx.Hash(), block.Hash(), i)

		receipt, privateReceipt, err := ApplyTransaction(p.config, p.bc, nil, gp, statedb, privateStateDB, header, tx, usedGas, cfg, privateStateRepo.IsMPS())
		if err != nil {
			return nil, nil, nil, 0, err
		}

		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)

		// Quorum
		// if the private receipt is nil, or privacy marker receipt is nil,
		// then the tx was public and we do not need to apply the additional logic.
		if privateReceipt != nil {
			privateReceipts = append(privateReceipts, privateReceipt)
			allLogs = append(allLogs, privateReceipt.Logs...)
			p.bc.CheckAndSetPrivateState(privateReceipt.Logs, privateStateDB, privateStateRepo.DefaultStateMetadata().ID)
			// handling the auxiliary receipt from MPS execution
			if mpsReceipt != nil {
				privateReceipt.PSReceipts = mpsReceipt.PSReceipts
				allLogs = append(allLogs, mpsReceipt.Logs...)
			}
		} else {
			if markerReceipt := rawdb.ReadPrivateTransactionReceipt(p.bc.db, tx.Hash()); markerReceipt != nil {
				allLogs = append(allLogs, markerReceipt.Logs...)
				p.bc.CheckAndSetPrivateState(markerReceipt.Logs, privateStateDB, privateStateRepo.DefaultStateMetadata().ID)
			}
			// handling the auxiliary receipt from MPS execution
			if mpsReceipt != nil {
				privateReceipt.PSReceipts = mpsReceipt.PSReceipts
				allLogs = append(allLogs, mpsReceipt.Logs...)
			}
		}
		// End Quorum
	}
	// Finalize the block, applying any consensus engine specific extras (e.g. block rewards)
	p.engine.Finalize(p.bc, header, statedb, block.Transactions(), block.Uncles())

	return receipts, privateReceipts, allLogs, *usedGas, nil
}

// Quorum
// returns the privateStateDB to be used for a transaction
func PrivateStateDBForTxn(isQuorum, isPrivate bool, stateDb, privateStateDB *state.StateDB) *state.StateDB {
	if !isQuorum || !isPrivate {
		return stateDb
	}
	return privateStateDB
}

// Quorum
// handling MPS scenario for a private transaction
//
// handleMPS returns the auxiliary receipt and not the standard receipt
func (p *StateProcessor) handleMPS(ti int, tx *types.Transaction, block *types.Block, gp *GasPool, usedGas *uint64, cfg vm.Config, statedb *state.StateDB, privateStateRepo mps.PrivateStateRepository) (mpsReceipt *types.Receipt, err error) {
	if tx.IsPrivate() && privateStateRepo.IsMPS() {
		publicStateDBFactory := func() *state.StateDB {
			db := statedb.Copy()
			db.Prepare(tx.Hash(), block.Hash(), ti)
			return db
		}
		privateStateDBFactory := func(psi types.PrivateStateIdentifier) (*state.StateDB, error) {
			db, err := privateStateRepo.StatePSI(psi)
			if err != nil {
				return nil, err
			}
			db.Prepare(tx.Hash(), block.Hash(), ti)
			return db, nil
		}
		mpsReceipt, err = ApplyTransactionOnMPS(p.config, p.bc, nil, gp, publicStateDBFactory, privateStateDBFactory, block.Header(), tx, usedGas, cfg)
	}
	return
}

// Quorum
// ApplyTransactionOnMPS runs the transaction on multiple private states which
// the transaction is designated to.
//
// For each designated private state, the transaction is ran only ONCE.
//
// ApplyTransactionOnMPS returns the auxiliary receipt which is mainly used to capture
// multiple private receipts and logs array. Logs are decorated with types.PrivateStateIdentifier
//
// The originalGP gas pool will not be modified
func ApplyTransactionOnMPS(config *params.ChainConfig, bc *BlockChain, author *common.Address, originalGP *GasPool,
	publicStateDBFactory func() *state.StateDB, privateStateDBFactory func(psi types.PrivateStateIdentifier) (*state.StateDB, error),
	header *types.Header, tx *types.Transaction, usedGas *uint64, cfg vm.Config) (*types.Receipt, error) {
	// clone the gas pool (as we don't want to keep consuming intrinsic gas multiple times for each MPS execution)
	gp := new(GasPool).AddGas(originalGP.Gas())
	mpsReceipt := &types.Receipt{
		PSReceipts: make(map[types.PrivateStateIdentifier]*types.Receipt),
		Logs:       make([]*types.Log, 0),
	}
	_, managedParties, _, _, err := private.P.Receive(common.BytesToEncryptedPayloadHash(tx.Data()))
	if err != nil {
		return nil, err
	}
	targetPsi := make(map[types.PrivateStateIdentifier]struct{})
	for _, managedParty := range managedParties {
		psMetadata, err := bc.PrivateStateManager().ResolveForManagedParty(managedParty)
		if err != nil {
			return nil, err
		}
		targetPsi[psMetadata.ID] = struct{}{}
	}
	// execute in all the managed private states
	// TODO this could be enhanced to run in parallel
	for _, psi := range bc.PrivateStateManager().PSIs() {
		_, applyAsParty := targetPsi[psi]
		privateStateDB, err := privateStateDBFactory(psi)
		if err != nil {
			return nil, err
		}
		publicStateDB := publicStateDBFactory()
		_, receipt, err := ApplyTransaction(config, bc, author, gp, publicStateDB, privateStateDB, header, tx, usedGas, cfg, !applyAsParty)
		if err != nil {
			return nil, err
		}
		// set the PSI for each log (so that the filter system knows for what private state they are)
		// we don't care about the empty receipt (as we'll execute the transaction on the empty state anyway)
		if applyAsParty {
			for _, log := range receipt.Logs {
				log.PSI = psi
				mpsReceipt.Logs = append(mpsReceipt.Logs, log)
			}
			mpsReceipt.PSReceipts[psi] = receipt

			bc.CheckAndSetPrivateState(receipt.Logs, privateStateDB, psi)
		}
	}
	return mpsReceipt, nil
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gas used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(config *params.ChainConfig, bc *BlockChain, author *common.Address, gp *GasPool, statedb, privateStateDB *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *uint64, cfg vm.Config, forceNonParty bool) (*types.Receipt, *types.Receipt, error) {
	// Quorum - decide the privateStateDB to use
	privateStateDbToUse := PrivateStateDBForTxn(config.IsQuorum, tx.IsPrivate(), statedb, privateStateDB)
	// /Quorum

	// Quorum - check for account permissions to execute the transaction
	if core.IsV2Permission() {
		if err := core.CheckAccountPermission(tx.From(), tx.To(), tx.Value(), tx.Data(), tx.Gas(), tx.GasPrice()); err != nil {
			return nil, nil, err
		}
	}

	if config.IsQuorum && tx.GasPrice() != nil && tx.GasPrice().Cmp(common.Big0) > 0 {
		return nil, nil, ErrInvalidGasPrice
	}

	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return nil, nil, err
	}

	// Quorum: this tx needs to be applied as if we were not a party
	msg = msg.WithEmptyPrivateData(forceNonParty && tx.IsPrivate())
	// Create a new context to be used in the EVM environment
	context := NewEVMContext(msg, header, bc, author)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, statedb, privateStateDbToUse, config, cfg)
	// the same transaction object is used for multiple executions (clear the privacy metadata - it should be updated after privacyManager.receive)
	// when running in parallel for multiple private states is implemented - a copy of the tx may be used
	tx.SetTxPrivacyMetadata(nil)
	vmenv.SetCurrentTX(tx)

	// Quorum
	txIndex := statedb.TxIndex()
	vmenv.InnerApply = func(innerTx *types.Transaction) (*types.Receipt, error) {
		defer func() {
			statedb.Prepare(tx.Hash(), header.Hash(), txIndex)
			privateStateDB.Prepare(tx.Hash(), header.Hash(), txIndex)
		}()
		statedb.Prepare(innerTx.Hash(), header.Hash(), txIndex)
		privateStateDB.Prepare(innerTx.Hash(), header.Hash(), txIndex)

		singleUseGasPool := new(GasPool).AddGas(innerTx.Gas())
		used := uint64(0)
		_, privReceipt, err := ApplyTransaction(config, bc, author, singleUseGasPool, statedb, privateStateDB, header, innerTx, &used, cfg, forceNonParty)

		if privReceipt != nil {
			if privReceipt.Logs == nil {
				privReceipt.Logs = make([]*types.Log, 0)
			}
			privateStateDB.MarkerTransactionReceipts = append(privateStateDB.MarkerTransactionReceipts, privReceipt)
			rawdb.WritePrivateTransactionReceipt(bc.db, tx.Hash(), privReceipt)
		}

		return privReceipt, err
	}
	// End Quorum

	// Apply the transaction to the current state (included in the env)
	result, err := ApplyMessage(vmenv, msg, gp)
	if err != nil {
		return nil, nil, err
	}
	// Update the state with pending changes
	var root []byte
	if config.IsByzantium(header.Number) {
		statedb.Finalise(true)
	} else {
		root = statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
	}
	*usedGas += result.UsedGas

	// If this is a private transaction, the public receipt should always
	// indicate success.
	publicFailed := !(config.IsQuorum && tx.IsPrivate()) && result.Failed()

	// Create a new receipt for the transaction, storing the intermediate root and gas used by the tx
	// based on the eip phase, we're passing wether the root touch-delete accounts.
	receipt := types.NewReceipt(root, publicFailed, *usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = result.UsedGas
	// if the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
	}
	// Set the receipt logs and create a bloom for filtering
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
	receipt.BlockHash = statedb.BlockHash()
	receipt.BlockNumber = header.Number
	receipt.TransactionIndex = uint(statedb.TxIndex())

	var privateReceipt *types.Receipt
	if config.IsQuorum && tx.IsPrivate() {
		var privateRoot []byte
		if config.IsByzantium(header.Number) {
			privateStateDB.Finalise(true)
		} else {
			privateRoot = privateStateDB.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
		}
		privateReceipt = types.NewReceipt(privateRoot, result.Failed(), *usedGas)
		privateReceipt.TxHash = tx.Hash()
		privateReceipt.GasUsed = result.UsedGas
		if msg.To() == nil {
			privateReceipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
		}

		privateReceipt.Logs = privateStateDB.GetLogs(tx.Hash())
		privateReceipt.Bloom = types.CreateBloom(types.Receipts{privateReceipt})
	}

	return receipt, privateReceipt, err
}
