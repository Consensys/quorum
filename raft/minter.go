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

package gethRaft

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

	"gopkg.in/fatih/set.v0"
	lane "gopkg.in/oleiade/lane.v1"
)

const throttleMillis = 50

// Work is the minter's current environment and holds
// all of the current state information
type Work struct {
	config *core.ChainConfig

	//
	// TODO(bts): separate private and public state
	//
	state *state.StateDB // apply state changes here

	tcount int // tx count in cycle

	Block *types.Block // the new block

	header   *types.Header
	txs      []*types.Transaction
	receipts []*types.Receipt

	createdAt time.Time
}

type minter struct {
	config *core.ChainConfig

	//
	// TODO(bts): revisit this, vs currentMu.
	//
	mu sync.Mutex

	// event loop
	mux    *event.TypeMux
	events event.Subscription
	wg     sync.WaitGroup

	eth     core.Backend
	chain   *core.BlockChain
	proc    core.Validator
	chainDb ethdb.Database

	coinbase common.Address

	//
	// TODO: do we even need this any more?
	//
	currentMu sync.Mutex
	current   *Work

	// atomic status counter
	minting int32

	proposedTxes               *set.Set
	expectedInvalidBlockHashes *set.Set
	shouldMine                 *channels.RingChannel

	parent *types.Block

	unappliedBlocks *lane.Deque
}

func newMinter(config *core.ChainConfig, eth core.Backend) *minter {
	minter := &minter{
		config:       config,
		eth:          eth,
		mux:          eth.EventMux(),
		chainDb:      eth.ChainDb(),
		chain:        eth.BlockChain(),
		proc:         eth.BlockChain().Validator(),
		proposedTxes: set.New(),
		shouldMine:   channels.NewRingChannel(1),
	}
	minter.events = minter.mux.Subscribe(core.ChainHeadEvent{}, core.TxPreEvent{}, InvalidRaftOrdering{})

	minter.parent = minter.chain.CurrentBlock()
	minter.unappliedBlocks = lane.NewDeque()
	minter.expectedInvalidBlockHashes = set.New()

	go minter.eventLoop()
	go minter.kickOffMinting()

	minter.currentMu.Lock()
	minter.createAndSetCurrentWork()
	minter.currentMu.Unlock()

	return minter
}

func (minter *minter) pending() (*types.Block, *state.StateDB) {
	minter.currentMu.Lock()
	defer minter.currentMu.Unlock()

	if atomic.LoadInt32(&minter.minting) == 0 {
		return types.NewBlock(
			minter.current.header,
			minter.current.txs,
			nil,
			minter.current.receipts,
		), minter.current.state
	}
	return minter.current.Block, minter.current.state
}

func (minter *minter) start() {
	atomic.StoreInt32(&minter.minting, 1)
}

func (minter *minter) stop() {
	minter.wg.Wait()

	minter.mu.Lock()
	defer minter.mu.Unlock()

	//
	// TODO(bts): clear speculative mining state.
	//

	atomic.StoreInt32(&minter.minting, 0)
}

func (minter *minter) requestMinting() {
	minter.shouldMine.In() <- struct{}{}
}

// This is for "speculative" minting. We keep track of txes we've put in all
// newly-mined blocks since the last ChainHeadEvent, and filter them out so that
// we don't try to create blocks with the same transactions. This is necessary
// because the TX pool will keep supplying us these transactions until they are
// in the chain (after having flown through raft).
func (minter *minter) raftFilterProposedTxes(addrTxes map[common.Address]types.Transactions) map[common.Address]types.Transactions {
	newMap := make(map[common.Address]types.Transactions)

	for addr, txes := range addrTxes {
		filteredTxes := make(types.Transactions, 0)
		for _, tx := range txes {
			if !minter.proposedTxes.Has(tx.Hash()) {
				filteredTxes = append(filteredTxes, tx)
			}
		}
		if len(filteredTxes) > 0 {
			newMap[addr] = filteredTxes
		}
	}

	return newMap
}

func (minter *minter) removeProposedTxes(block *types.Block) {
	minedTxes := block.Transactions()
	minedTxInterfaces := make([]interface{}, len(minedTxes))
	for i, tx := range minedTxes {
		minedTxInterfaces[i] = tx.Hash()
	}

	// NOTE: we are using a thread-safe Set here, so it's fine if we access this
	// here and in mintNewBlock concurrently. using a finer-grained set-specific
	// lock here is preferable, because mintNewBlock holds its locks for a
	// nontrivial amount of time.
	minter.proposedTxes.Remove(minedTxInterfaces...)
}

func (minter *minter) eventLoop() {
	for event := range minter.events.Chan() {
		switch ev := event.Data.(type) {
		case core.ChainHeadEvent:
			newHeadBlock := ev.Block

			minter.mu.Lock()
			if atomic.LoadInt32(&minter.minting) == 1 {
				earliestProposedI := minter.unappliedBlocks.Shift()

				var earliestProposed *types.Block
				if nil != earliestProposedI {
					earliestProposed = earliestProposedI.(*types.Block)
				}

				if earliestProposed != nil && earliestProposed.Hash() != newHeadBlock.Hash() {
					// Another node has minted and had its block accepted. We clear all
					// speculative mining state.

					glog.V(logger.Warn).Infof("Another node mined %x; Clearing speculative minting state\n", newHeadBlock.Hash())

					minter.parent = newHeadBlock
					minter.proposedTxes.Clear()
					minter.unappliedBlocks = lane.NewDeque()
				} else {
					// We're receiving the acceptance for an expected block. Remove its
					// txes from our blacklist.
					minter.removeProposedTxes(newHeadBlock)
				}

				//
				// TODO(bts): not sure if this is the place, but we're going to need to put
				// an upper limit on our speculative mining chain length.
				//

				minter.requestMinting()
			} else {
				minter.parent = newHeadBlock
			}
			minter.mu.Unlock()

		case core.TxPreEvent:
			// Apply transaction to the pending state if we're not minting
			if atomic.LoadInt32(&minter.minting) == 0 {
				minter.currentMu.Lock()
				if from, err := ev.Tx.From(); err != nil {
					txMap := map[common.Address]types.Transactions{from: types.Transactions{ev.Tx}}
					txes := types.NewTransactionsByPriceAndNonce(txMap)
					// TODO(bts): check whether this is needed to broadcast txes while not minting
					minter.current.commitTransactions(minter.mux, txes, minter.chain)
				}
				minter.currentMu.Unlock()
			} else {
				minter.requestMinting()
			}

		case InvalidRaftOrdering:
			headBlock := ev.headBlock
			invalidBlock := ev.invalidBlock
			invalidHash := invalidBlock.Hash()

			glog.V(logger.Warn).Infof("Handling InvalidRaftOrdering for invalid block %x; current head is %x\n", invalidHash, headBlock.Hash())

			minter.mu.Lock()

			// 1. if the block is not in our db, exit. someone else mined this.
			if !minter.chain.HasBlock(invalidBlock.Hash()) {
				glog.V(logger.Warn).Infof("Someone else mined invalid block %x; ignoring\n", invalidHash)

				minter.mu.Unlock()
				continue
			}

			// 2. check our guard to see if this is a (descendent) block we're
			// expected to be ruled invalid. if we find it, remove from the guard
			if minter.expectedInvalidBlockHashes.Has(invalidHash) {
				glog.V(logger.Warn).Infof("Removing expected-invalid block %x from guard.\n", invalidHash)

				minter.expectedInvalidBlockHashes.Remove(invalidHash)

				minter.mu.Unlock()
				continue
			}

			// 3. pop from the RHS repeatedly, updating minter.parent each time. if not
			// our block, add to guard. in all cases, call removeProposedTxes
			for {
				currBlockI := minter.unappliedBlocks.Pop()

				if nil == currBlockI {
					glog.V(logger.Warn).Infof("(Popped all blocks from queue.)\n")

					break
				}

				currBlock := currBlockI.(*types.Block)

				glog.V(logger.Info).Infof("Popped block %x from queue RHS.\n", currBlock.Hash())

				if speculativeParentI := minter.unappliedBlocks.Last(); nil != speculativeParentI {
					minter.parent = speculativeParentI.(*types.Block)
				} else {
					minter.parent = headBlock
				}

				minter.removeProposedTxes(currBlock)

				if currBlock.Hash() != invalidHash {
					glog.V(logger.Warn).Infof("Haven't yet found block %x; adding descendent %x to guard.\n", invalidHash, currBlock.Hash())

					minter.expectedInvalidBlockHashes.Add(currBlock.Hash())
				} else {
					break
				}
			}
			minter.mu.Unlock()
		}
	}
}

// After `f` been called, wait before calling it again.
// Note that we will allow `f` to be called immediately, but not again before
// a timeout.
//
// TODO(joel) this has a small bug in that you can't call it *immediately* when
// first allocated
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

// NOTE: We kick off a goroutine to the side of the minter so the `eventLoop`
// can run continuously (it adds an anonymous struct to `shouldMine` to
// request a block creation, when it sees a transaction created; this is
// non-blocking and can be done >= 1 times with the same effect). This goroutine
// spins continuously, blocking until a block should be created.
func (minter *minter) kickOffMinting() {
	// Throttling is simple but has two nice properties for this use case:
	// 1. A block is guaranteed to be created within throttleMillis of being requested
	// 2. We never create a block more frequently than throttleMillis
	rate := throttleMillis * time.Millisecond
	throttledCommitNewWork := throttle(rate, func() {
		if atomic.LoadInt32(&minter.minting) == 1 {
			minter.mintNewBlock()
		}
	})

	for range minter.shouldMine.Out() {
		throttledCommitNewWork()
	}
}

func (minter *minter) broadcastWork(work *Work) {
	block := work.Block

	go func(block *types.Block, logs vm.Logs, receipts []*types.Receipt) {
		minter.mux.Post(core.NewMinedBlockEvent{Block: block})

		minter.mux.Post(core.ChainEvent{Block: block, Hash: block.Hash(), Logs: logs})

		// NOTE: we're currently not doing this anymore because the block is not
		// in the chain yet (in previous wait(), we only do this when stat is
		// CanonStatTy):
		//
		// minter.mux.Post(logs)

		if err := core.WriteBlockReceipts(minter.chainDb, block.Hash(), block.Number().Uint64(), receipts); err != nil {
			glog.V(logger.Warn).Infoln("error writing block receipts:", err)
		}
	}(block, work.state.Logs(), work.receipts)

	elapsed := time.Since(time.Unix(0, work.header.Time.Int64()))
	glog.V(logger.Info).Infof("🔨  Mined block (#%v / %x) in %v", block.Number(), block.Hash().Bytes()[:4], elapsed)
}

// makeCurrent creates a new environment for the current cycle.
// This method assumes currentMu is being held
func (minter *minter) createWork(tstamp int64) (*Work, error) {
	parent := minter.parent
	parentNumber := parent.Number()

	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     parentNumber.Add(parentNumber, common.Big1),
		Difficulty: core.CalcDifficulty(minter.config, uint64(tstamp), parent.Time().Uint64(), parent.Number(), parent.Difficulty()),
		GasLimit:   core.CalcGasLimit(parent),
		GasUsed:    new(big.Int),
		Coinbase:   minter.coinbase,
		Time:       big.NewInt(tstamp),
	}

	state, err := state.New(minter.parent.Root(), minter.eth.ChainDb())
	if err != nil {
		return nil, err
	}
	work := &Work{
		config:    minter.config,
		state:     state,
		header:    header,
		createdAt: time.Now(),
	}

	work.tcount = 0

	return work, nil
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

// We assume that currentMu is being held
func (minter *minter) createAndSetCurrentWork() {
	tstamp := generateNanoTimestamp(minter.parent)

	// An error could potentially happen if starting to mine in an odd state.
	work, workErr := minter.createWork(tstamp)
	if workErr != nil {
		glog.V(logger.Info).Infoln("Could not create new env for minting, retrying on next block.")
		return
	}
	minter.current = work
}

func (minter *minter) mintNewBlock() {
	minter.mu.Lock()
	defer minter.mu.Unlock()
	minter.currentMu.Lock()
	defer minter.currentMu.Unlock()

	minter.createAndSetCurrentWork()
	work := minter.current

	allAddrTxes := minter.eth.TxPool().Pending()
	addrTxes := minter.raftFilterProposedTxes(allAddrTxes)
	transactions := types.NewTransactionsByPriceAndNonce(addrTxes)

	committedTxes := work.commitTransactions(minter.mux, transactions, minter.chain)

	committedTxIs := make([]interface{}, len(committedTxes))
	for i, tx := range committedTxes {
		committedTxIs[i] = tx.Hash()
	}
	minter.proposedTxes.Add(committedTxIs...)

	var emptyUncles []*types.Header
	header := work.header

	if atomic.LoadInt32(&minter.minting) == 1 {
		// commit state root after all state transitions.
		core.AccumulateRewards(work.state, header, emptyUncles)
		header.Root = work.state.IntermediateRoot()
	}

	// create the new block whose nonce will be mined.
	work.Block = types.NewBlock(header, work.txs, emptyUncles, work.receipts)

	if work.tcount == 0 {
		// Don't mine if there is nothing to do..
		glog.V(logger.Info).Infoln("Not generating new block since there are no pending transactions")
		return
	}

	diffchange := new(big.Int).Sub(header.Difficulty, minter.parent.Difficulty())
	glog.V(logger.Info).Infof("Generated next block #%v with [%d txns] new difficulty: %d (%+v)", work.Block.Number(), work.tcount, header.Difficulty, diffchange)

	work.state.Commit()

	block := work.Block

	//
	// TODO(bts): what to do for validation here now that AuxValidator has been
	//            removed in Quorum?
	//
	// auxValidator := minter.eth.BlockChain().AuxValidator()
	// if err := core.ValidateHeader(minter.config, auxValidator, block.Header(), minter.parent.Header(), true, false); err != nil && err != core.BlockFutureErr {
	// 	panic(fmt.Sprint("Invalid header on mined block: ", err))
	// }

	if err := minter.chain.WriteDetachedBlock(block); err != nil {
		panic(fmt.Sprint("error writing block to chain", err))
	}

	// update block hash since it is now available and not when the receipt/log of individual transactions were created
	for _, r := range work.receipts {
		for _, l := range r.Logs {
			l.BlockHash = block.Hash()
		}
	}
	for _, log := range work.state.Logs() {
		log.BlockHash = block.Hash()
	}

	// NOTE: We are currently writing txes, receipts, and the bloom filters
	// even though this block might not end up becoming the next head of the
	// chain.
	//
	// This puts transactions in a extra db for rpc
	core.WriteTransactions(minter.chainDb, block)
	// store the receipts
	core.WriteReceipts(minter.chainDb, work.receipts)
	// Write map map bloom filters
	core.WriteMipmapBloom(minter.chainDb, block.NumberU64(), work.receipts)

	minter.parent = block
	minter.unappliedBlocks.Append(block)

	minter.broadcastWork(work)
}

func (env *Work) commitTransactions(mux *event.TypeMux, txes *types.TransactionsByPriceAndNonce, bc *core.BlockChain) types.Transactions {
	committedTxes := make(types.Transactions, 0)
	gp := new(core.GasPool).AddGas(env.header.GasLimit)

	var coalescedLogs vm.Logs
	for {
		tx := txes.Peek()
		if tx == nil {
			break
		}

		env.state.StartRecord(tx.Hash(), common.Hash{}, 0)

		logs, err := env.commitTransaction(tx, bc, gp)
		switch {
		case err != nil:
			if glog.V(logger.Detail) {
				glog.Infof("TX (%x) failed, will be removed: %v\n", tx.Hash().Bytes()[:4], err)
			}
			txes.Pop() // skip rest of txes from this account
		default:
			env.tcount++
			coalescedLogs = append(coalescedLogs, logs...)
			committedTxes = append(committedTxes, tx)
			txes.Shift()
		}
	}
	if len(coalescedLogs) > 0 || env.tcount > 0 {
		// make a copy, the state caches the logs and these logs get "upgraded" from pending to mined
		// logs by filling in the block hash when the block was mined by the local miner. This can
		// cause a race condition if a log was "upgraded" before the PendingLogsEvent is processed.
		cpy := make(vm.Logs, len(coalescedLogs))
		for i, l := range coalescedLogs {
			cpy[i] = new(vm.Log)
			*cpy[i] = *l
		}
		go func(logs vm.Logs, tcount int) {
			if len(logs) > 0 {
				mux.Post(core.PendingLogsEvent{Logs: logs})
			}
			if tcount > 0 {
				mux.Post(core.PendingStateEvent{})
			}
		}(cpy, env.tcount)
	}

	return committedTxes
}

func (env *Work) commitTransaction(tx *types.Transaction, bc *core.BlockChain, gp *core.GasPool) (vm.Logs, error) {
	snap := env.state.Snapshot()

	//
	// TODO: separate private and public state here
	//
	receipt, logs, _, err := core.ApplyTransaction(env.config, bc, gp, env.state, env.state, env.header, tx, env.header.GasUsed, env.config.VmConfig)
	if err != nil {
		env.state.RevertToSnapshot(snap)
		return nil, err
	}
	env.txs = append(env.txs, tx)
	env.receipts = append(env.receipts, receipt)

	return logs, nil
}
