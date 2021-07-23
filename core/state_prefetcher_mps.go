package core

import (
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/log"
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

// MPSPrefetcher is an interface for pre-caching transaction signatures and state for MPS
type MPSPrefetcher interface {
	
	// Prefetch analogous to Prefetcher.Prefetch but applied to MPS
	Prefetch(block *types.Block, statedb *state.StateDB, privateStateRepo mps.PrivateStateRepository, cfg vm.Config, interrupt *uint32)
}

func runMPSPrefetch(mpsPrefetcher MPSPrefetcher, followup *types.Block, throwaway *state.StateDB, throwawayPrivateStateRepo mps.PrivateStateRepository, vmConfig vm.Config, followupInterrupt *uint32) {
	log.Debug("#### blockchain Private Prefetch 1", "number", followup.Number(), "hash", followup.Hash())
	go func(start time.Time, followup *types.Block, throwaway *state.StateDB, throwawayPrivateStateRepo mps.PrivateStateRepository, interrupt *uint32) {
		mpsPrefetcher.Prefetch(followup, throwaway, throwawayPrivateStateRepo, vmConfig, interrupt)

		// TODO ricardolyn
		//blockPrefetchExecuteTimer.Update(time.Since(start))
		//if atomic.LoadUint32(interrupt) == 1 {
		//	blockPrefetchInterruptMeter.Mark(1)
		//}
	}(time.Now(), followup, throwaway, throwawayPrivateStateRepo, followupInterrupt)
}

// stateMPSPrefetcher is a basic Prefetcher, which blindly executes a block on top
// of an arbitrary state with the goal of prefetching potentially useful state
// data from disk before the main block processor start executing.
type stateMPSPrefetcher struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
	engine consensus.Engine    // Consensus engine used for block rewards

	pend sync.WaitGroup
}

// newStateMPSPrefetcher initialises a new stateMPSPrefetcher.
func newStateMPSPrefetcher(config *params.ChainConfig, bc *BlockChain, engine consensus.Engine) *stateMPSPrefetcher {
	return &stateMPSPrefetcher{
		config: config,
		bc:     bc,
		engine: engine,
	}
}

// Prefetch processes the state changes according to the Ethereum rules by running
// the transaction messages using the statedb, but any changes are discarded. The
// only goal is to pre-cache transaction signatures and state trie nodes.
func (p *stateMPSPrefetcher) Prefetch(block *types.Block, statedb *state.StateDB, privateStateRepo mps.PrivateStateRepository, cfg vm.Config, interrupt *uint32) {
	var (
		header  = block.Header()
		gaspool = new(GasPool).AddGas(block.GasLimit())
	)
	log.Debug("#### Private Prefetch 1", "number", block.Number(), "hash", block.Hash())
	// Iterate over and process the individual transactions
	byzantium := p.config.IsByzantium(block.Number())
	for i, tx := range block.Transactions() {
		// If block precaching was interrupted, abort
		if interrupt != nil && atomic.LoadUint32(interrupt) == 1 {
			return
		}
		if !tx.IsPrivate() {
			// to ignore
			log.Trace("#### Private Prefetch 2.1", "tx.hash", tx.Hash())
			return
		}

		// Block precaching permitted to continue, execute the transaction
		_, managedParties, _, _, err := private.P.Receive(common.BytesToEncryptedPayloadHash(tx.Data()))
		if err != nil {
			// TODO ricardolyn: log error?
			log.Trace("#### Private Prefetch 2.2", "err", err)
			return
		}
		for _, managedParty := range managedParties {
			if interrupt != nil && atomic.LoadUint32(interrupt) == 1 {
				return
			}
			psMetadata, err := p.bc.PrivateStateManager().ResolveForManagedParty(managedParty)
			if err != nil {
				// TODO ricardolyn: skip this PSI as it doesn't exist locally?
				continue
			}

			privateStateDb, err := privateStateRepo.StatePSI(psMetadata.ID)
			p.pend.Add(1)
			go func(start time.Time, followup *types.Block, statedb *state.StateDB, privateStateDb *state.StateDB) {
				privateStateDb.Prepare(tx.Hash(), block.Hash(), i)
				if err := precachePrivateTransaction(p.config, p.bc, nil, gaspool, statedb, privateStateDb, header, tx, cfg); err != nil {
					return
				}
				// If we're pre-byzantium, pre-load trie nodes for the intermediate root
				if !byzantium {
					privateStateDb.IntermediateRoot(true)
				}
				p.pend.Done()
			}(time.Now(), block, statedb, privateStateDb)
			p.pend.Wait()
		}
	}
	// TODO ricardolyn: should we do it?
	//// If were post-byzantium, pre-load trie nodes for the final root hash
	//if byzantium {
	//	statedb.IntermediateRoot(true)
	//}
}

// precacheTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. The goal is not to execute
// the transaction successfully, rather to warm up touched data slots.
func precachePrivateTransaction(config *params.ChainConfig, bc ChainContext, author *common.Address, gaspool *GasPool, statedb *state.StateDB, privateStatedb *state.StateDB, header *types.Header, tx *types.Transaction, cfg vm.Config) error {
	// Convert the transaction into an executable message and pre-cache its sender
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return err
	}
	// Create the EVM and execute the transaction
	context := NewEVMContext(msg, header, bc, author)

	//only precaching public txs
	vm := vm.NewEVM(context, statedb, privateStatedb, config, cfg)
	vm.SetCurrentTX(tx)
	// /Quorum

	_, err = ApplyMessage(vm, msg, gaspool)
	return err
}
