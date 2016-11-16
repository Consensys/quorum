// Copyright 2014 The go-ethereum Authors
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
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
)

// GetHashFn returns a function for which the VM env can query block hashes through
// up to the limit defined by the Yellow Paper and uses the given block chain
// to query for information.
func GetHashFn(ref common.Hash, chain *BlockChain) func(n uint64) common.Hash {
	return func(n uint64) common.Hash {
		for block := chain.GetBlockByHash(ref); block != nil; block = chain.GetBlock(block.ParentHash(), block.NumberU64()-1) {
			if block.NumberU64() == n {
				return block.Hash()
			}
		}

		return common.Hash{}
	}
}

type DualStateEnv interface {
	vm.Environment

	PublicState() *state.StateDB
	PrivateState() *state.StateDB

	Push(*state.StateDB)
	Pop()
}

type VMEnv struct {
	publicState, privateState *state.StateDB // State to use for executing
	states                    [1027]*state.StateDB
	currentStateDepth         uint
	readOnly                  bool
	readOnlyDepth             uint

	chainConfig *ChainConfig // Chain configuration
	evm         *vm.EVM      // The Ethereum Virtual Machine
	depth       int          // Current execution depth
	msg         Message      // Message appliod

	header    *types.Header            // Header information
	chain     *BlockChain              // Blockchain handle
	getHashFn func(uint64) common.Hash // getHashFn callback is used to retrieve block hashes
}

// NewEnv creates a new environment for executing a transaction.
// In case the transaction is public its the responsibility from
// the caller to supply publicState for the privateState argument.
func NewEnv(publicState, privateState *state.StateDB, chainConfig *ChainConfig, chain *BlockChain, msg Message, header *types.Header, cfg vm.Config) *VMEnv {
	env := &VMEnv{
		chainConfig:  chainConfig,
		chain:        chain,
		publicState:  publicState,
		privateState: privateState,
		header:       header,
		msg:          msg,
		getHashFn:    GetHashFn(header.ParentHash, chain),
	}

	env.Push(privateState)

	env.evm = vm.New(env, cfg)
	return env
}

func (env *VMEnv) ReadOnly() bool               { return env.readOnly }
func (env *VMEnv) PublicState() *state.StateDB  { return env.publicState }
func (env *VMEnv) PrivateState() *state.StateDB { return env.privateState }
func (env *VMEnv) Push(state *state.StateDB) {
	if env.privateState != state {
		env.readOnly = true
		env.readOnlyDepth = env.currentStateDepth
	}

	env.states[env.currentStateDepth] = state
	env.currentStateDepth++
}
func (env *VMEnv) Pop() {
	env.currentStateDepth--
	if env.readOnly && env.currentStateDepth == env.readOnlyDepth {
		env.readOnly = false
	}
}
func (env *VMEnv) currentState() *state.StateDB { return env.states[env.currentStateDepth-1] }

func (self *VMEnv) RuleSet() vm.RuleSet      { return self.chainConfig }
func (self *VMEnv) Vm() vm.Vm                { return self.evm }
func (self *VMEnv) Origin() common.Address   { f, _ := self.msg.From(); return f }
func (self *VMEnv) BlockNumber() *big.Int    { return self.header.Number }
func (self *VMEnv) Coinbase() common.Address { return self.header.Coinbase }
func (self *VMEnv) Time() *big.Int           { return self.header.Time }
func (self *VMEnv) Difficulty() *big.Int     { return self.header.Difficulty }
func (self *VMEnv) GasLimit() *big.Int       { return self.header.GasLimit }
func (self *VMEnv) Value() *big.Int          { return self.msg.Value() }
func (self *VMEnv) Db() vm.Database {
	return self.currentState()
}
func (self *VMEnv) Depth() int     { return self.depth }
func (self *VMEnv) SetDepth(i int) { self.depth = i }
func (self *VMEnv) GetHash(n uint64) common.Hash {
	return self.getHashFn(n)
}

func (self *VMEnv) AddLog(log *vm.Log) {
	self.currentState().AddLog(log)
}
func (self *VMEnv) CanTransfer(from common.Address, balance *big.Int) bool {
	return self.currentState().GetBalance(from).Cmp(balance) >= 0
}

func (self *VMEnv) SnapshotDatabase() int {
	return self.currentState().Snapshot()
}

// We only need to revert the current state because when we call from private
// public state it's read only, there wouldn't be anything to reset.
// (A)->(B)->C->(B): A failure in (B) wouldn't need to reset C, as C was flagged
// read only.
func (self *VMEnv) RevertToSnapshot(snapshot int) {
	self.currentState().RevertToSnapshot(snapshot)
}

func (self *VMEnv) Transfer(from, to vm.Account, amount *big.Int) {
	Transfer(from, to, amount)
}

func (self *VMEnv) Call(me vm.ContractRef, addr common.Address, data []byte, gas, price, value *big.Int) ([]byte, error) {
	return Call(self, me, addr, data, gas, price, value)
}
func (self *VMEnv) CallCode(me vm.ContractRef, addr common.Address, data []byte, gas, price, value *big.Int) ([]byte, error) {
	return CallCode(self, me, addr, data, gas, price, value)
}

func (self *VMEnv) DelegateCall(me vm.ContractRef, addr common.Address, data []byte, gas, price *big.Int) ([]byte, error) {
	return DelegateCall(self, me, addr, data, gas, price)
}

func (self *VMEnv) Create(me vm.ContractRef, data []byte, gas, price, value *big.Int) ([]byte, common.Address, error) {
	return Create(self, me, data, gas, price, value)
}
