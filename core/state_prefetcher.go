// Copyright 2019 The go-ethereum Authors
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
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/private"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

// statePrefetcher is a basic Prefetcher, which blindly executes a block on top
// of an arbitrary state with the goal of prefetching potentially useful state
// data from disk before the main block processor start executing.
type statePrefetcher struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
	engine consensus.Engine    // Consensus engine used for block rewards

	pend sync.WaitGroup // Quorum: wait for MPS prefetching
}

// newStatePrefetcher initialises a new statePrefetcher.
func newStatePrefetcher(config *params.ChainConfig, bc *BlockChain, engine consensus.Engine) *statePrefetcher {
	return &statePrefetcher{
		config: config,
		bc:     bc,
		engine: engine,
	}
}

// Prefetch processes the state changes according to the Ethereum rules by running
// the transaction messages using the statedb, but any changes are discarded. The
// only goal is to pre-cache transaction signatures and state trie nodes.
// Quorum: Add privateStateDb argument
func (p *statePrefetcher) Prefetch(block *types.Block, statedb *state.StateDB, privateStateRepo mps.PrivateStateRepository, cfg vm.Config, interrupt *uint32) {
	var (
		header  = block.Header()
		gaspool = new(GasPool).AddGas(block.GasLimit())
	)
	// Iterate over and process the individual transactions
	byzantium := p.config.IsByzantium(block.Number())
	for i, tx := range block.Transactions() {
		// If block precaching was interrupted, abort
		if interrupt != nil && atomic.LoadUint32(interrupt) == 1 {
			return
		}

		// Quorum
		if tx.IsPrivate() && privateStateRepo.IsMPS() {
			p.prefetchMpsTransaction(block, tx, i, statedb, privateStateRepo, cfg, interrupt)
		}
		privateStateDb, _ := privateStateRepo.DefaultState()
		privateStateDb.Prepare(tx.Hash(), block.Hash(), i)
		// End Quorum

		// Block precaching permitted to continue, execute the transaction
		statedb.Prepare(tx.Hash(), block.Hash(), i)

		// Quorum: Add privateStateDb argument
		if err := precacheTransaction(p.config, p.bc, nil, gaspool, statedb, privateStateDb, header, tx, cfg, false); err != nil {
			return // Ugh, something went horribly wrong, bail out
		}
		// If we're pre-byzantium, pre-load trie nodes for the intermediate root
		if !byzantium {
			statedb.IntermediateRoot(true)
		}
	}
	// If were post-byzantium, pre-load trie nodes for the final root hash
	if byzantium {
		statedb.IntermediateRoot(true)
	}
}

// precacheTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. The goal is not to execute
// the transaction successfully, rather to warm up touched data slots.
// Quorum: Add privateStateDb and isMPS arguments
func precacheTransaction(config *params.ChainConfig, bc ChainContext, author *common.Address, gaspool *GasPool, statedb *state.StateDB, privateStateDb *state.StateDB, header *types.Header, tx *types.Transaction, cfg vm.Config, isMPS bool) error {
	// Convert the transaction into an executable message and pre-cache its sender
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return err
	}
	if isMPS {
		msg = msg.WithEmptyPrivateData(tx.IsPrivate())
	}
	// Create the EVM and execute the transaction
	context := NewEVMContext(msg, header, bc, author)
	// Quorum: Add privateStaterDb argument
	var evm *vm.EVM
	// TODO (ricardolyn): this is confusing passing the private state or not
	if tx.IsPrivate() {
		evm = vm.NewEVM(context, statedb, privateStateDb, config, cfg)
	} else {
		evm = vm.NewEVM(context, statedb, statedb, config, cfg)
	}
	evm.SetCurrentTX(tx) // Quorum

	_, err = ApplyMessage(evm, msg, gaspool)
	return err
}

// Quorum

func (p *statePrefetcher) prefetchMpsTransaction(block *types.Block, tx *types.Transaction, txIndex int, statedb *state.StateDB, privateStateRepo mps.PrivateStateRepository, cfg vm.Config, interrupt *uint32) {
	var (
		header  = block.Header()
		gaspool = new(GasPool).AddGas(block.GasLimit())
	)
	byzantium := p.config.IsByzantium(block.Number())
	// Block precaching permitted to continue, execute the transaction
	_, managedParties, _, _, err := private.P.Receive(common.BytesToEncryptedPayloadHash(tx.Data()))
	if err != nil {
		return
	}
	for _, managedParty := range managedParties {
		if interrupt != nil && atomic.LoadUint32(interrupt) == 1 {
			return
		}
		psMetadata, err := p.bc.PrivateStateManager().ResolveForManagedParty(managedParty)
		if err != nil {
			continue
		}

		privateStateDb, err := privateStateRepo.StatePSI(psMetadata.ID)
		if err != nil {
			continue
		}
		p.pend.Add(1)
		go func(start time.Time, followup *types.Block, statedb *state.StateDB, privateStateDb *state.StateDB) {
			privateStateDb.Prepare(tx.Hash(), block.Hash(), txIndex)
			if err := precacheTransaction(p.config, p.bc, nil, gaspool, statedb, privateStateDb, header, tx, cfg, true); err != nil {
				return
			}
			// If we're pre-byzantium, pre-load trie nodes for the intermediate root
			if !byzantium {
				privateStateDb.IntermediateRoot(true)
			}
			p.pend.Done()
		}(time.Now(), block, statedb, privateStateDb)
	}
	p.pend.Wait()
}
