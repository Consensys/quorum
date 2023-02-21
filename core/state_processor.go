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
	"errors"
	"fmt"

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
	blockContext := NewEVMBlockContext(header, p.bc, nil)
	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		mpsReceipt, err := handleMPS(i, tx, gp, usedGas, cfg, statedb, privateStateRepo, p.config, p.bc, header, false, false)
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

		privateStateDBToUse := PrivateStateDBForTxn(p.config.IsQuorum, tx, statedb, privateStateDB)

		// Quorum - check for account permissions to execute the transaction
		if core.IsV2Permission() {
			if err := core.CheckAccountPermission(tx.From(), tx.To(), tx.Value(), tx.Data(), tx.Gas(), tx.GasPrice()); err != nil {
				return nil, nil, nil, 0, fmt.Errorf("could not apply tx %d [%v]: %w", i, tx.Hash().Hex(), err)
			}
		}

		if p.config.IsQuorum && !p.config.IsGasPriceEnabled(header.Number) && tx.GasPrice() != nil && tx.GasPrice().Cmp(common.Big0) > 0 {
			return nil, nil, nil, 0, fmt.Errorf("could not apply tx %d [%v]: %w", i, tx.Hash().Hex(), ErrInvalidGasPrice)
		}

		msg, err := tx.AsMessage(types.MakeSigner(p.config, header.Number))
		if err != nil {
			return nil, nil, nil, 0, fmt.Errorf("could not apply tx %d [%v]: %w", i, tx.Hash().Hex(), err)
		}

		// Quorum: this tx needs to be applied as if we were not a party
		msg = msg.WithEmptyPrivateData(privateStateRepo.IsMPS() && tx.IsPrivate())

		// the same transaction object is used for multiple executions (clear the privacy metadata - it should be updated after privacyManager.receive)
		// when running in parallel for multiple private states is implemented - a copy of the tx may be used
		tx.SetTxPrivacyMetadata(nil)

		txContext := NewEVMTxContext(msg)
		vmenv := vm.NewEVM(blockContext, txContext, statedb, privateStateDBToUse, p.config, cfg)
		vmenv.SetCurrentTX(tx)
		receipt, privateReceipt, err := applyTransaction(msg, p.config, p.bc, nil, gp, statedb, privateStateDB, header, tx, usedGas, vmenv, cfg, privateStateRepo.IsMPS(), privateStateRepo)
		if err != nil {
			return nil, nil, nil, 0, fmt.Errorf("could not apply tx %d [%v]: %w", i, tx.Hash().Hex(), err)
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
func PrivateStateDBForTxn(isQuorum bool, tx *types.Transaction, stateDb, privateStateDB *state.StateDB) *state.StateDB {
	if isQuorum && (tx.IsPrivate() || tx.IsPrivacyMarker()) {
		return privateStateDB
	}
	return stateDb
}

// Quorum
// handling MPS scenario for a private transaction
//
// handleMPS returns the auxiliary receipt and not the standard receipt
func handleMPS(ti int, tx *types.Transaction, gp *GasPool, usedGas *uint64, cfg vm.Config, statedb *state.StateDB, privateStateRepo mps.PrivateStateRepository, config *params.ChainConfig, bc ChainContext, header *types.Header, applyOnPartiesOnly bool, isInnerPrivateTxn bool) (mpsReceipt *types.Receipt, err error) {
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
		mpsReceipt, err = ApplyTransactionOnMPS(config, bc, nil, gp, publicStateDBFactory, privateStateDBFactory, header, tx, usedGas, cfg, privateStateRepo, applyOnPartiesOnly, isInnerPrivateTxn)
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
func ApplyTransactionOnMPS(config *params.ChainConfig, bc ChainContext, author *common.Address, originalGP *GasPool,
	publicStateDBFactory func() *state.StateDB, privateStateDBFactory func(psi types.PrivateStateIdentifier) (*state.StateDB, error),
	header *types.Header, tx *types.Transaction, usedGas *uint64, cfg vm.Config, privateStateRepo mps.PrivateStateRepository,
	applyOnPartiesOnly bool, isInnerPrivateTxn bool) (*types.Receipt, error) {
	mpsReceipt := &types.Receipt{
		QuorumReceiptExtraData: types.QuorumReceiptExtraData{
			PSReceipts: make(map[types.PrivateStateIdentifier]*types.Receipt),
		},
		Logs: make([]*types.Log, 0),
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
		if cfg.ApplyOnPartyOverride != nil && *cfg.ApplyOnPartyOverride != psi {
			continue
		}
		_, applyAsParty := targetPsi[psi]
		if !applyAsParty && applyOnPartiesOnly {
			continue
		}
		privateStateDB, err := privateStateDBFactory(psi)
		if err != nil {
			return nil, err
		}
		publicStateDB := publicStateDBFactory()

		// use a clone of the gas pool (as we don't want to consume gas multiple times for each MPS execution, which might blow the block gasLimit on MPS node)
		gp := new(GasPool).AddGas(originalGP.Gas())

		_, privateReceipt, err := ApplyTransaction(config, bc, author, gp, publicStateDB, privateStateDB, header, tx, usedGas, cfg, !applyAsParty, privateStateRepo, isInnerPrivateTxn)
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

// /Quorum

func applyTransaction(msg types.Message, config *params.ChainConfig, bc ChainContext, author *common.Address, gp *GasPool, statedb, privateStateDB *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *uint64, evm *vm.EVM, cfg vm.Config, forceNonParty bool, privateStateRepo mps.PrivateStateRepository) (*types.Receipt, *types.Receipt, error) {
	// Create a new context to be used in the EVM environment.

	// Quorum
	txIndex := statedb.TxIndex()
	evm.InnerApply = func(innerTx *types.Transaction) error {
		return ApplyInnerTransaction(bc, author, gp, statedb, privateStateDB, header, tx, usedGas, cfg, forceNonParty, privateStateRepo, evm, innerTx, txIndex)
	}
	// End Quorum

	// Apply the transaction to the current state (included in the env)
	result, err := ApplyMessage(evm, msg, gp)
	if err != nil {
		return nil, nil, err
	}

	// Update the state with pending changes.
	var root []byte
	if config.IsByzantium(header.Number) {
		statedb.Finalise(true)
	} else {
		root = statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
	}
	*usedGas += result.UsedGas

	// Create a new receipt for the transaction, storing the intermediate root and gas used
	// by the tx.
	receipt := &types.Receipt{Type: tx.Type(), PostState: common.CopyBytes(root), CumulativeGasUsed: *usedGas}

	// If this is a private transaction, the public receipt should always
	// indicate success.
	if !(config.IsQuorum && tx.IsPrivate()) && result.Failed() {
		receipt.Status = types.ReceiptStatusFailed
	} else {
		receipt.Status = types.ReceiptStatusSuccessful
	}
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = result.UsedGas

	// If the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(evm.TxContext.Origin, tx.Nonce())
	}

	// Set the receipt logs and create the bloom filter.
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
				privateReceipt.ContractAddress = crypto.CreateAddress(evm.TxContext.Origin, tx.Nonce())
			}

			privateReceipt.Logs = privateStateDB.GetLogs(tx.Hash())
			privateReceipt.Bloom = types.CreateBloom(types.Receipts{privateReceipt})
		} else {
			// This may have been a privacy marker transaction, in which case need to retrieve the receipt for the
			// inner private transaction (note that this can be an mpsReceipt, containing private receipts in PSReceipts).
			if evm.InnerPrivateReceipt != nil {
				privateReceipt = evm.InnerPrivateReceipt
			}
		}
	}

	// Save revert reason if feature enabled
	if bc != nil && bc.QuorumConfig().RevertReasonEnabled() {
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

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gas used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(config *params.ChainConfig, bc ChainContext, author *common.Address, gp *GasPool, statedb, privateStateDB *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *uint64, cfg vm.Config, forceNonParty bool, privateStateRepo mps.PrivateStateRepository, isInnerPrivateTxn bool) (*types.Receipt, *types.Receipt, error) {
	// Quorum - decide the privateStateDB to use
	privateStateDbToUse := PrivateStateDBForTxn(config.IsQuorum, tx, statedb, privateStateDB)
	// End Quorum

	// Quorum - check for account permissions to execute the transaction
	if core.IsV2Permission() {
		if err := core.CheckAccountPermission(tx.From(), tx.To(), tx.Value(), tx.Data(), tx.Gas(), tx.GasPrice()); err != nil {
			return nil, nil, err
		}
	}

	if config.IsQuorum && !config.IsGasPriceEnabled(header.Number) && tx.GasPrice() != nil && tx.GasPrice().Cmp(common.Big0) > 0 {
		return nil, nil, ErrInvalidGasPrice
	}

	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return nil, nil, err
	}
	// Quorum: this tx needs to be applied as if we were not a party
	msg = msg.WithEmptyPrivateData(forceNonParty && tx.IsPrivate())
	// Quorum: if this is the inner private txn of a PMT then need to indicate this
	msg = msg.WithInnerPrivateFlag(isInnerPrivateTxn)

	// Create a new context to be used in the EVM environment
	blockContext := NewEVMBlockContext(header, bc, author)
	txContext := NewEVMTxContext(msg)
	vmenv := vm.NewEVM(blockContext, txContext, statedb, privateStateDbToUse, config, cfg)

	// the same transaction object is used for multiple executions (clear the privacy metadata - it should be updated after privacyManager.receive)
	// when running in parallel for multiple private states is implemented - a copy of the tx may be used
	tx.SetTxPrivacyMetadata(nil)
	vmenv.SetCurrentTX(tx)

	return applyTransaction(msg, config, bc, author, gp, statedb, privateStateDB, header, tx, usedGas, vmenv, cfg, forceNonParty, privateStateRepo)
}

// Quorum

// ApplyInnerTransaction is called from within the Quorum precompile for privacy marker transactions.
// It's a call back which essentially duplicates the logic in Process(),
// in this case to process the actual private transaction.
func ApplyInnerTransaction(bc ChainContext, author *common.Address, gp *GasPool, stateDB *state.StateDB, privateStateDB *state.StateDB, header *types.Header, outerTx *types.Transaction, usedGas *uint64, evmConf vm.Config, forceNonParty bool, privateStateRepo mps.PrivateStateRepository, vmenv *vm.EVM, innerTx *types.Transaction, txIndex int) error {
	// this should never happen, but added as sanity check
	if !innerTx.IsPrivate() {
		return errors.New("attempt to process non-private transaction from within ApplyInnerTransaction()")
	}

	// create a single use gas pool (as we don't want the gas consumed by the inner tx to blow the block gasLimit on a participant node)
	singleUseGasPool := new(GasPool).AddGas(innerTx.Gas())

	if privateStateRepo != nil && privateStateRepo.IsMPS() {
		mpsReceipt, err := handleMPS(txIndex, innerTx, singleUseGasPool, usedGas, evmConf, stateDB, privateStateRepo, bc.Config(), bc, header, true, true)
		if err != nil {
			return err
		}

		// Store the auxiliary MPS receipt for the inner private transaction (this contains private receipts in PSReceipts).
		vmenv.InnerPrivateReceipt = mpsReceipt
		return nil
	}

	defer prepareStates(outerTx, stateDB, privateStateDB, txIndex)
	prepareStates(innerTx, stateDB, privateStateDB, txIndex)

	used := uint64(0)
	_, innerPrivateReceipt, err := ApplyTransaction(bc.Config(), bc, author, singleUseGasPool, stateDB, privateStateDB, header, innerTx, &used, evmConf, forceNonParty, privateStateRepo, true)
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

// Quorum
func prepareStates(tx *types.Transaction, stateDB *state.StateDB, privateStateDB *state.StateDB, txIndex int) {
	stateDB.Prepare(tx.Hash(), stateDB.BlockHash(), txIndex)
	privateStateDB.Prepare(tx.Hash(), privateStateDB.BlockHash(), txIndex)
}
