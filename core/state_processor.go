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
	"math/big"

	"github.com/ethereum/go-ethereum/common"
<<<<<<< HEAD
=======
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
>>>>>>> 7cc6abeef6ec0b6c5fd5a94920fa79157cdfcd37
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
<<<<<<< HEAD
)

var (
	big8  = big.NewInt(8)
	big32 = big.NewInt(32)
=======
	"github.com/ethereum/go-ethereum/params"
>>>>>>> 7cc6abeef6ec0b6c5fd5a94920fa79157cdfcd37
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
<<<<<<< HEAD
func (p *StateProcessor) Process(block *types.Block, publicState, privateState *state.StateDB, cfg vm.Config) (types.Receipts, types.Receipts, vm.Logs, *big.Int, error) {
	var (
		publicReceipts  types.Receipts
		privateReceipts types.Receipts
		totalUsedGas    = big.NewInt(0)
		err             error
		header          = block.Header()
		allLogs         vm.Logs
		gp              = new(GasPool).AddGas(block.GasLimit())
	)

	for i, tx := range block.Transactions() {
		publicState.StartRecord(tx.Hash(), block.Hash(), i)
		privateState.StartRecord(tx.Hash(), block.Hash(), i)

		publicReceipt, privateReceipt, _, err := ApplyTransaction(p.config, p.bc, gp, publicState, privateState, header, tx, totalUsedGas, cfg)
		if err != nil {
			return nil, nil, nil, totalUsedGas, err
		}
		publicReceipts = append(publicReceipts, publicReceipt)
		allLogs = append(allLogs, publicReceipt.Logs...)

		// if the private receipt is nil this means the tx was public
		// and we do not need to apply the additional logic.
		if privateReceipt != nil {
			privateReceipts = append(privateReceipts, privateReceipt)
			allLogs = append(allLogs, privateReceipt.Logs...)
		}
	}
	AccumulateRewards(publicState, header, block.Uncles())

	return publicReceipts, privateReceipts, allLogs, totalUsedGas, err
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment.
//
// ApplyTransactions returns the generated receipts and vm logs during the
// execution of the state transition phase.
func ApplyTransaction(config *ChainConfig, bc *BlockChain, gp *GasPool, publicState, privateState *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *big.Int, cfg vm.Config) (*types.Receipt, *types.Receipt, *big.Int, error) {
	if !tx.IsPrivate() {
		privateState = publicState
	}

	if tx.GasPrice() != nil && tx.GasPrice().Cmp(common.Big0) > 0 {
		return nil, nil, nil, ErrInvalidGasPrice
	}

	_, gas, err := ApplyMessage(NewEnv(publicState, privateState, config, bc, tx, header, cfg), tx, gp)
=======
func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, cfg vm.Config) (types.Receipts, []*types.Log, *big.Int, error) {
	var (
		receipts     types.Receipts
		totalUsedGas = big.NewInt(0)
		header       = block.Header()
		allLogs      []*types.Log
		gp           = new(GasPool).AddGas(block.GasLimit())
	)
	// Mutate the the block and state according to any hard-fork specs
	if p.config.DAOForkSupport && p.config.DAOForkBlock != nil && p.config.DAOForkBlock.Cmp(block.Number()) == 0 {
		misc.ApplyDAOHardFork(statedb)
	}
	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		statedb.StartRecord(tx.Hash(), block.Hash(), i)
		receipt, _, err := ApplyTransaction(p.config, p.bc, nil, gp, statedb, header, tx, totalUsedGas, cfg)
		if err != nil {
			return nil, nil, nil, err
		}
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)
	}
	// Finalize the block, applying any consensus engine specific extras (e.g. block rewards)
	p.engine.Finalize(p.bc, header, statedb, block.Transactions(), block.Uncles(), receipts)

	return receipts, allLogs, totalUsedGas, nil
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gas used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(config *params.ChainConfig, bc *BlockChain, author *common.Address, gp *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *big.Int, cfg vm.Config) (*types.Receipt, *big.Int, error) {
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return nil, nil, err
	}
	// Create a new context to be used in the EVM environment
	context := NewEVMContext(msg, header, bc, author)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, statedb, config, cfg)
	// Apply the transaction to the current state (included in the env)
	_, gas, err := ApplyMessage(vmenv, msg, gp)
>>>>>>> 7cc6abeef6ec0b6c5fd5a94920fa79157cdfcd37
	if err != nil {
		return nil, nil, err
	}

	// Update the state with pending changes
	usedGas.Add(usedGas, gas)
<<<<<<< HEAD
	publicReceipt := types.NewReceipt(publicState.IntermediateRoot().Bytes(), usedGas)
	publicReceipt.TxHash = tx.Hash()
	publicReceipt.GasUsed = new(big.Int).Set(gas)
	if MessageCreatesContract(tx) {
		from, _ := tx.From()
		publicReceipt.ContractAddress = crypto.CreateAddress(from, tx.Nonce())
	}

	logs := publicState.GetLogs(tx.Hash())
	publicReceipt.Logs = logs
	publicReceipt.Bloom = types.CreateBloom(types.Receipts{publicReceipt})

	var privateReceipt *types.Receipt
	if tx.IsPrivate() {
		privateReceipt = types.NewReceipt(privateState.IntermediateRoot().Bytes(), usedGas)
		privateReceipt.TxHash = tx.Hash()
		privateReceipt.GasUsed = new(big.Int).Set(gas)
		if MessageCreatesContract(tx) {
			from, _ := tx.From()
			privateReceipt.ContractAddress = crypto.CreateAddress(from, tx.Nonce())
		}

		logs := privateState.GetLogs(tx.Hash())
		privateReceipt.Logs = logs
		privateReceipt.Bloom = types.CreateBloom(types.Receipts{privateReceipt})
	}

	return publicReceipt, privateReceipt, gas, err
}

// AccumulateRewards credits the coinbase of the given block with the
// mining reward. The total reward consists of the static block reward
// and rewards for included uncles. The coinbase of each uncle block is
// also rewarded.
func AccumulateRewards(statedb *state.StateDB, header *types.Header, uncles []*types.Header) {
	reward := new(big.Int).Set(BlockReward)
	r := new(big.Int)
	for _, uncle := range uncles {
		r.Add(uncle.Number, big8)
		r.Sub(r, header.Number)
		r.Mul(r, BlockReward)
		r.Div(r, big8)
		statedb.AddBalance(uncle.Coinbase, r)

		r.Div(BlockReward, big32)
		reward.Add(reward, r)
	}
	statedb.AddBalance(header.Coinbase, reward)
=======
	// Create a new receipt for the transaction, storing the intermediate root and gas used by the tx
	// based on the eip phase, we're passing wether the root touch-delete accounts.
	receipt := types.NewReceipt(statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes(), usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)
	// if the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
	}

	// Set the receipt logs and create a bloom for filtering
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	return receipt, gas, err
>>>>>>> 7cc6abeef6ec0b6c5fd5a94920fa79157cdfcd37
}
