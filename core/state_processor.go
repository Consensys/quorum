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
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
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
		mpsReceipt, err := handleMPS(i, tx, gp, usedGas, cfg, statedb, privateStateRepo, p.config, p.bc, header)
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

		receipt, privateReceipt, err := ApplyTransaction(p.config, p.bc, nil, gp, statedb, privateStateDB, header, tx, usedGas, cfg, privateStateRepo.IsMPS(), privateStateRepo)
		if err != nil {
			return nil, nil, nil, 0, err
		}

		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)

		// if the private receipt is nil this means the tx was public
		// and we do not need to apply the additional logic.
		if privateReceipt != nil {
			newPrivateReceipt, privateLogs := HandlePrivateReceipt(receipt, privateReceipt, mpsReceipt, tx, privateStateDB, privateStateRepo, p.bc)
			privateReceipts = append(privateReceipts, newPrivateReceipt)
			allLogs = append(allLogs, privateLogs...)
		}
	}
	// Finalize the block, applying any consensus engine specific extras (e.g. block rewards)
	p.engine.Finalize(p.bc, header, statedb, block.Transactions(), block.Uncles())

	return receipts, privateReceipts, allLogs, *usedGas, nil
}

// Quorum
func HandlePrivateReceipt(receipt *types.Receipt, privateReceipt *types.Receipt, mpsReceipt *types.Receipt, tx *types.Transaction, privateStateDB *state.StateDB, privateStateRepo mps.PrivateStateRepository, bc *BlockChain) (*types.Receipt, []*types.Log) {
	var (
		privateLogs []*types.Log
	)

	if tx.IsPrivacyMarker() {
		// This was a public privacy marker transaction, so we need to handle two scenarios:
		//	1) MPS: privateReceipt is an auxiliary MPS receipt which contains actual private receipts in PSReceipts[]
		//	2) non-MPS: privateReceipt is the actual receipt for the inner private transaction
		// In both cases we return a receipt for the public PMT, which holds the private receipt(s) in PSReceipts[],
		// and we then discard the privateReceipt.
		if privateStateRepo != nil && privateStateRepo.IsMPS() {
			receipt.PSReceipts = privateReceipt.PSReceipts
			privateLogs = append(privateLogs, privateReceipt.Logs...)
		} else {
			receipt.PSReceipts = make(map[types.PrivateStateIdentifier]*types.Receipt)
			receipt.PSReceipts[privateStateRepo.DefaultStateMetadata().ID] = privateReceipt
			privateLogs = append(privateLogs, privateReceipt.Logs...)
			bc.CheckAndSetPrivateState(privateReceipt.Logs, privateStateDB, privateStateRepo.DefaultStateMetadata().ID)
		}

		// There should be no auxiliary receipt from MPS execution, just logging in case this ever occurs
		if mpsReceipt != nil {
			log.Error("Unexpected MPS auxiliary receipt, when processing a privacy marker transaction")
		}
		return privateReceipt, privateLogs
	} else {
		// This was a regular private transaction.
		privateLogs = append(privateLogs, privateReceipt.Logs...)
		bc.CheckAndSetPrivateState(privateReceipt.Logs, privateStateDB, privateStateRepo.DefaultStateMetadata().ID)

		// handling the auxiliary receipt from MPS execution
		if mpsReceipt != nil {
			privateReceipt.PSReceipts = mpsReceipt.PSReceipts
			privateLogs = append(privateLogs, mpsReceipt.Logs...)
		}
		return privateReceipt, privateLogs
	}
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
func handleMPS(ti int, tx *types.Transaction, gp *GasPool, usedGas *uint64, cfg vm.Config, statedb *state.StateDB, privateStateRepo mps.PrivateStateRepository, config *params.ChainConfig, bc *BlockChain, header *types.Header) (mpsReceipt *types.Receipt, err error) {
	if tx.IsPrivate() && privateStateRepo != nil && privateStateRepo.IsMPS() {
		publicStateDBFactory := func() *state.StateDB {
			db := statedb.Copy()
			db.Prepare(tx.Hash(), header.Hash(), ti)
			return db
		}
		privateStateDBFactory := func(psi types.PrivateStateIdentifier) (*state.StateDB, error) {
			db, err := privateStateRepo.StatePSI(psi)
			if err != nil {
				return nil, err
			}
			db.Prepare(tx.Hash(), header.Hash(), ti)
			return db, nil
		}
		mpsReceipt, err = ApplyTransactionOnMPS(config, bc, nil, gp, publicStateDBFactory, privateStateDBFactory, header, tx, usedGas, cfg, privateStateRepo)
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
	header *types.Header, tx *types.Transaction, usedGas *uint64, cfg vm.Config, privateStateRepo mps.PrivateStateRepository) (*types.Receipt, error) {
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
		_, privateReceipt, err := ApplyTransaction(config, bc, author, gp, publicStateDB, privateStateDB, header, tx, usedGas, cfg, !applyAsParty, privateStateRepo)
		if err != nil {
			return nil, err
		}

		// set the PSI for each log (so that the filter system knows for what private state they are)
		// we don't care about the empty privateReceipt (as we'll execute the transaction on the empty state anyway)
		if applyAsParty {
			for _, log := range privateReceipt.Logs {
				log.PSI = psi
				mpsReceipt.Logs = append(mpsReceipt.Logs, log)
			}
			mpsReceipt.PSReceipts[psi] = privateReceipt

			bc.CheckAndSetPrivateState(privateReceipt.Logs, privateStateDB, psi)
		}
	}

	return mpsReceipt, nil
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gas used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(config *params.ChainConfig, bc *BlockChain, author *common.Address, gp *GasPool, statedb, privateStateDB *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *uint64, cfg vm.Config, forceNonParty bool, privateStateRepo mps.PrivateStateRepository) (*types.Receipt, *types.Receipt, error) {
	// Quorum - decide the privateStateDB to use
	privateStateDBToUse := PrivateStateDBForTxn(config.IsQuorum, tx.IsPrivate(), statedb, privateStateDB)
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
	vmenv := vm.NewEVM(context, statedb, privateStateDBToUse, config, cfg)
	// the same transaction object is used for multiple executions (clear the privacy metadata - it should be updated after privacyManager.receive)
	// when running in parallel for multiple private states is implemented - a copy of the tx may be used
	tx.SetTxPrivacyMetadata(nil)
	vmenv.SetCurrentTX(tx)

	if config.IsYoloV2(header.Number) {
		statedb.AddAddressToAccessList(msg.From())
		if dst := msg.To(); dst != nil {
			statedb.AddAddressToAccessList(*dst)
			// If it's a create-tx, the destination will be added inside evm.create
		}
		for _, addr := range vmenv.ActivePrecompiles() {
			statedb.AddAddressToAccessList(addr)
		}
	}

	// Quorum
	txIndex := statedb.TxIndex()
	vmenv.InnerApply = func(innerTx *types.Transaction) error {
		if innerTx.IsPrivate() && privateStateRepo != nil && privateStateRepo.IsMPS() {
			mpsReceipt, err := handleMPS(txIndex, innerTx, gp, usedGas, cfg, statedb, privateStateRepo, config, bc, header)
			if err != nil {
				return err
			}

			// Store the auxiliary MPS receipt for the inner private transaction (this contains private receipts in PSReceipts).
			vmenv.InnerPrivateReceipt = mpsReceipt
			return nil
		}

		defer func() {
			statedb.Prepare(tx.Hash(), statedb.BlockHash(), txIndex)
			privateStateDB.Prepare(tx.Hash(), privateStateDB.BlockHash(), txIndex)
		}()
		statedb.Prepare(innerTx.Hash(), statedb.BlockHash(), txIndex)
		privateStateDB.Prepare(innerTx.Hash(), privateStateDB.BlockHash(), txIndex)

		singleUseGasPool := new(GasPool).AddGas(innerTx.Gas())
		used := uint64(0)
		_, innerPrivateReceipt, err := ApplyTransaction(config, bc, author, singleUseGasPool, statedb, privateStateDB, header, innerTx, &used, cfg, forceNonParty, privateStateRepo)
		if err != nil {
			return err
		}

		if innerPrivateReceipt != nil {
			if innerPrivateReceipt.Logs == nil {
				innerPrivateReceipt.Logs = make([]*types.Log, 0)
			}

			// Store the receipt for the inner private transaction.
			innerPrivateReceipt.TxHash = innerTx.Hash()
			vmenv.InnerPrivateReceipt = innerPrivateReceipt
		}

		return nil
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

	// Quorum
	var privateReceipt *types.Receipt
	if config.IsQuorum {
		if tx.IsPrivate() {
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
		} else {
			// This may have been a privacy marker transaction, in which case need to retrieve the receipt for the
			// inner private transaction (note that this can be an mpsReceipt, containing private receipts in PSReceipts).
			if vmenv.InnerPrivateReceipt != nil {
				privateReceipt = vmenv.InnerPrivateReceipt
			}
		}
	}

	// Save revert reason if feature enabled
	if bc != nil && bc.saveRevertReason {
		revertReason := result.Revert()
		if revertReason != nil {
			if config.IsQuorum && tx.IsPrivate() {
				privateReceipt.RevertReason = revertReason
			} else {
				receipt.RevertReason = revertReason
			}
		}
	}
	// End Quorum

	return receipt, privateReceipt, err
}
