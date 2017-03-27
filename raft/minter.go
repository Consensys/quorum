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

package raft

import (
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/eapache/channels"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
)

// Current state information for building the next block
type work struct {
	config       *core.ChainConfig
	publicState  *state.StateDB
	privateState *state.StateDB
	Block        *types.Block
	header       *types.Header
}

type minter struct {
	config           *core.ChainConfig
	mu               sync.Mutex
	mux              *event.TypeMux
	eth              core.Backend
	chain            *core.BlockChain
	chainDb          ethdb.Database
	coinbase         common.Address
	minting          int32 // Atomic status counter
	shouldMine       *channels.RingChannel
	blockTime        time.Duration
	speculativeChain *speculativeChain
}

func newMinter(config *core.ChainConfig, eth core.Backend, blockTime time.Duration) *minter {
	minter := &minter{
		config:           config,
		eth:              eth,
		mux:              eth.EventMux(),
		chainDb:          eth.ChainDb(),
		chain:            eth.BlockChain(),
		shouldMine:       channels.NewRingChannel(1),
		blockTime:        blockTime,
		speculativeChain: newSpeculativeChain(),
	}
	events := minter.mux.Subscribe(
		core.ChainHeadEvent{},
		core.TxPreEvent{},
		InvalidRaftOrdering{},
	)

	minter.speculativeChain.clear(minter.chain.CurrentBlock())

	go minter.eventLoop(events)
	go minter.mintingLoop()

	return minter
}

func (minter *minter) start() {
	atomic.StoreInt32(&minter.minting, 1)
	minter.requestMinting()
}

func (minter *minter) stop() {
	minter.mu.Lock()
	defer minter.mu.Unlock()

	minter.speculativeChain.clear(minter.chain.CurrentBlock())
	atomic.StoreInt32(&minter.minting, 0)
}

// Notify the minting loop that minting should occur, if it's not already been
// requested. Due to the use of a RingChannel, this function is idempotent if
// called multiple times before the minting occurs.
func (minter *minter) requestMinting() {
	minter.shouldMine.In() <- struct{}{}
}

type AddressTxes map[common.Address]types.Transactions

func (minter *minter) updateSpeculativeChainPerNewHead(newHeadBlock *types.Block) {
	minter.mu.Lock()
	defer minter.mu.Unlock()

	minter.speculativeChain.accept(newHeadBlock)
}

func (minter *minter) updateSpeculativeChainPerInvalidOrdering(headBlock *types.Block, invalidBlock *types.Block) {
	invalidHash := invalidBlock.Hash()

	glog.V(logger.Warn).Infof("Handling InvalidRaftOrdering for invalid block %x; current head is %x\n", invalidHash, headBlock.Hash())

	minter.mu.Lock()
	defer minter.mu.Unlock()

	// 1. if the block is not in our db, exit. someone else mined this.
	if !minter.chain.HasBlock(invalidHash) {
		glog.V(logger.Warn).Infof("Someone else mined invalid block %x; ignoring\n", invalidHash)

		return
	}

	minter.speculativeChain.unwindFrom(invalidHash, headBlock)
}

func (minter *minter) eventLoop(events event.Subscription) {
	for event := range events.Chan() {
		switch ev := event.Data.(type) {
		case core.ChainHeadEvent:
			newHeadBlock := ev.Block

			if atomic.LoadInt32(&minter.minting) == 1 {
				minter.updateSpeculativeChainPerNewHead(newHeadBlock)

				//
				// TODO(bts): not sure if this is the place, but we're going to
				// want to put an upper limit on our speculative mining chain
				// length.
				//

				minter.requestMinting()
			} else {
				minter.mu.Lock()
				minter.speculativeChain.setHead(newHeadBlock)
				minter.mu.Unlock()
			}

		case core.TxPreEvent:
			if atomic.LoadInt32(&minter.minting) == 1 {
				minter.requestMinting()
			}

		case InvalidRaftOrdering:
			headBlock := ev.headBlock
			invalidBlock := ev.invalidBlock

			minter.updateSpeculativeChainPerInvalidOrdering(headBlock, invalidBlock)
		}
	}
}

// Returns a wrapper around no-arg func `f` which can be called without limit
// and returns immediately: this will call the underlying func `f` at most once
// every `rate`. If this function is called more than once before the underlying
// `f` is invoked (per this rate limiting), `f` will only be called *once*.
//
// TODO(joel): this has a small bug in that you can't call it *immediately* when
// first allocated.
func throttle(rate time.Duration, f func()) func() {
	request := channels.NewRingChannel(1)

	// every tick, block waiting for another request. then serve it immediately
	go func() {
		ticker := time.NewTicker(rate)
		defer ticker.Stop()

		for range ticker.C {
			<-request.Out()
			go f()
		}
	}()

	return func() {
		request.In() <- struct{}{}
	}
}

// This function spins continuously, blocking until a block should be created
// (via requestMinting()). This is throttled by `minter.blockTime`:
//
//   1. A block is guaranteed to be minted within `blockTime` of being
//      requested.
//   2. We never mint a block more frequently than `blockTime`.
func (minter *minter) mintingLoop() {
	throttledMintNewBlock := throttle(minter.blockTime, func() {
		if atomic.LoadInt32(&minter.minting) == 1 {
			minter.mintNewBlock()
		}
	})

	for range minter.shouldMine.Out() {
		throttledMintNewBlock()
	}
}

func generateNanoTimestamp(parent *types.Block) (tstamp int64) {
	parentTime := parent.Time().Int64()
	tstamp = time.Now().UnixNano()

	if parentTime >= tstamp {
		// Each successive block needs to be after its predecessor.
		tstamp = parentTime + 1
	}

	return
}

// Assumes mu is held.
func (minter *minter) createWork() *work {
	parent := minter.speculativeChain.head
	parentNumber := parent.Number()
	tstamp := generateNanoTimestamp(parent)

	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     parentNumber.Add(parentNumber, common.Big1),
		Difficulty: core.CalcDifficulty(minter.config, uint64(tstamp), parent.Time().Uint64(), parent.Number(), parent.Difficulty()),
		GasLimit:   core.CalcGasLimit(parent),
		GasUsed:    new(big.Int),
		Coinbase:   minter.coinbase,
		Time:       big.NewInt(tstamp),
	}

	publicState, privateState, err := minter.chain.StateAt(parent.Root())
	if err != nil {
		panic(fmt.Sprint("failed to get parent state: ", err))
	}

	return &work{
		config:       minter.config,
		publicState:  publicState,
		privateState: privateState,
		header:       header,
	}
}

func (minter *minter) getTransactions() *types.TransactionsByPriceAndNonce {
	allAddrTxes := minter.eth.TxPool().Pending()
	addrTxes := minter.speculativeChain.withoutProposedTxes(allAddrTxes)
	return types.NewTransactionsByPriceAndNonce(addrTxes)
}

// Sends-off events asynchronously.
func (minter *minter) firePendingBlockEvents(logs vm.Logs) {
	// Copy logs before we mutate them, adding a block hash.
	copiedLogs := make(vm.Logs, len(logs))
	for i, l := range logs {
		copiedLogs[i] = new(vm.Log)
		*copiedLogs[i] = *l
	}

	go func() {
		minter.mux.Post(core.PendingLogsEvent{Logs: copiedLogs})
		minter.mux.Post(core.PendingStateEvent{})
	}()
}

func (minter *minter) mintNewBlock() {
	minter.mu.Lock()
	defer minter.mu.Unlock()

	work := minter.createWork()
	transactions := minter.getTransactions()

	committedTxes, publicReceipts, privateReceipts, logs := work.commitTransactions(transactions, minter.chain)
	txCount := len(committedTxes)

	if txCount == 0 {
		glog.V(logger.Info).Infoln("Not minting a new block since there are no pending transactions")
		return
	}

	minter.firePendingBlockEvents(logs)

	header := work.header

	// commit state root after all state transitions.
	core.AccumulateRewards(work.publicState, header, nil)
	header.Root = work.publicState.IntermediateRoot()

	// NOTE: < QuorumChain creates a signature here and puts it in header.Extra. >

	allReceipts := append(publicReceipts, privateReceipts...)
	header.Bloom = types.CreateBloom(allReceipts)

	// update block hash since it is now available, but was not when the
	// receipt/log of individual transactions were created:
	headerHash := header.Hash()
	for _, l := range logs {
		l.BlockHash = headerHash
	}

	block := types.NewBlock(header, committedTxes, nil, publicReceipts)

	glog.V(logger.Info).Infof("Generated next block #%v with [%d txns]", block.Number(), txCount)

	if _, err := work.publicState.Commit(); err != nil {
		panic(fmt.Sprint("error committing public state: ", err))
	}
	if _, privStateErr := work.privateState.Commit(); privStateErr != nil {
		panic(fmt.Sprint("error committing private state: ", privStateErr))
	}

	minter.speculativeChain.extend(block)

	minter.mux.Post(core.NewMinedBlockEvent{Block: block})

	elapsed := time.Since(time.Unix(0, header.Time.Int64()))
	glog.V(logger.Info).Infof("🔨  Mined block (#%v / %x) in %v", block.Number(), block.Hash().Bytes()[:4], elapsed)
}

func (env *work) commitTransactions(txes *types.TransactionsByPriceAndNonce, bc *core.BlockChain) (types.Transactions, types.Receipts, types.Receipts, vm.Logs) {
	var logs vm.Logs
	var committedTxes types.Transactions
	var publicReceipts types.Receipts
	var privateReceipts types.Receipts

	gp := new(core.GasPool).AddGas(env.header.GasLimit)
	txCount := 0

	for {
		tx := txes.Peek()
		if tx == nil {
			break
		}

		env.publicState.StartRecord(tx.Hash(), common.Hash{}, 0)

		publicReceipt, privateReceipt, err := env.commitTransaction(tx, bc, gp)
		switch {
		case err != nil:
			if glog.V(logger.Detail) {
				glog.Infof("TX (%x) failed, will be removed: %v\n", tx.Hash().Bytes()[:4], err)
			}
			txes.Pop() // skip rest of txes from this account
		default:
			txCount++
			committedTxes = append(committedTxes, tx)

			logs = append(logs, publicReceipt.Logs...)
			publicReceipts = append(publicReceipts, publicReceipt)

			if privateReceipt != nil {
				logs = append(logs, privateReceipt.Logs...)
				privateReceipts = append(privateReceipts, privateReceipt)
			}

			txes.Shift()
		}
	}

	return committedTxes, publicReceipts, privateReceipts, logs
}

func (env *work) commitTransaction(tx *types.Transaction, bc *core.BlockChain, gp *core.GasPool) (*types.Receipt, *types.Receipt, error) {
	publicSnapshot := env.publicState.Snapshot()
	privateSnapshot := env.privateState.Snapshot()

	publicReceipt, privateReceipt, _, err := core.ApplyTransaction(env.config, bc, gp, env.publicState, env.privateState, env.header, tx, env.header.GasUsed, env.config.VmConfig)
	if err != nil {
		env.publicState.RevertToSnapshot(publicSnapshot)
		env.privateState.RevertToSnapshot(privateSnapshot)

		return nil, nil, err
	}

	return publicReceipt, privateReceipt, nil
}
