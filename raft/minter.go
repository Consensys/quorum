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

// Current state information for building the next block
type work struct {
	config       *core.ChainConfig
	publicState  *state.StateDB
	privateState *state.StateDB
	Block        *types.Block
	header       *types.Header
}

type minter struct {
	config                     *core.ChainConfig
	mu                         sync.Mutex
	mux                        *event.TypeMux
	eth                        core.Backend
	chain                      *core.BlockChain
	chainDb                    ethdb.Database
	coinbase                   common.Address
	minting                    int32 // atomic status counter
	proposedTxes               *set.Set
	expectedInvalidBlockHashes *set.Set
	shouldMine                 *channels.RingChannel
	blockTime                  time.Duration
	parent                     *types.Block
	unappliedBlocks            *lane.Deque
}

// Assumes mu is held.
// TODO(bts): extract all speculative fields into a new MintingState datatype.
func (minter *minter) clearSpeculativeState(parent *types.Block) {
	minter.parent = parent
	minter.proposedTxes.Clear()
	minter.unappliedBlocks = lane.NewDeque()
	minter.expectedInvalidBlockHashes = set.New()
}

func newMinter(config *core.ChainConfig, eth core.Backend, blockTime time.Duration) *minter {
	minter := &minter{
		config:       config,
		eth:          eth,
		mux:          eth.EventMux(),
		chainDb:      eth.ChainDb(),
		chain:        eth.BlockChain(),
		proposedTxes: set.New(),
		shouldMine:   channels.NewRingChannel(1),
		blockTime:    blockTime,
	}
	events := minter.mux.Subscribe(
		core.ChainHeadEvent{},
		core.TxPreEvent{},
		InvalidRaftOrdering{},
	)

	minter.clearSpeculativeState(minter.chain.CurrentBlock())

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

	minter.clearSpeculativeState(minter.chain.CurrentBlock())
	atomic.StoreInt32(&minter.minting, 0)
}

// Notify the minting loop that minting should occur, if it's not already been
// requested. Due to the use of a RingChannel, this function is idempotent if
// called multiple times before the minting occurs.
func (minter *minter) requestMinting() {
	minter.shouldMine.In() <- struct{}{}
}

type AddressTxes map[common.Address]types.Transactions

// This is for "speculative" minting. We keep track of txes we've put in all
// newly-mined blocks since the last ChainHeadEvent, and filter them out so that
// we don't try to create blocks with the same transactions. This is necessary
// because the TX pool will keep supplying us these transactions until they are
// in the chain (after having flown through raft).
func (minter *minter) withoutProposedTxes(addrTxes AddressTxes) AddressTxes {
	newMap := make(AddressTxes)

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

// Removes txes in block from our "blacklist" of "proposed tx" hashes. When we
// create a new block and use txes from the tx pool, we ignore those that we
// have already used ("proposed"), but that haven't yet officially made it into
// the chain yet.
//
// It's important to remove hashes from this blacklist (once we know we don't
// need them in there anymore) so that it doesn't grow endlessly.
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

func (minter *minter) eventLoop(events event.Subscription) {
	for event := range events.Chan() {
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
					glog.V(logger.Warn).Infof("Another node minted %x; Clearing speculative state\n", newHeadBlock.Hash())

					minter.clearSpeculativeState(newHeadBlock)
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
			if atomic.LoadInt32(&minter.minting) == 1 {
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

// This function spins continuously, blocking until a block should be created
// (via requestMinting()). This is throttled by `minter.blockTime`:
//
//   1. A block is guaranteed to be created within `blockTime` of being
//      requested.
//   2. We never create a block more frequently than `blockTime`.
func (minter *minter) mintingLoop() {
	throttledCommitNewWork := throttle(minter.blockTime, func() {
		if atomic.LoadInt32(&minter.minting) == 1 {
			minter.mintNewBlock()
		}
	})

	for range minter.shouldMine.Out() {
		throttledCommitNewWork()
	}
}

func (minter *minter) broadcastWork(work *work, receipts types.Receipts) {
	block := work.Block

	go func(block *types.Block, logs vm.Logs, receipts types.Receipts) {
		minter.mux.Post(core.NewMinedBlockEvent{Block: block})

		minter.mux.Post(core.ChainEvent{Block: block, Hash: block.Hash(), Logs: logs})

		// NOTE: we're currently not doing this because the block is not in the
		// chain yet:
		//
		// minter.mux.Post(logs)

		if err := core.WriteBlockReceipts(minter.chainDb, block.Hash(), block.Number().Uint64(), receipts); err != nil {
			glog.V(logger.Warn).Infoln("error writing block receipts:", err)
		}
	}(block, work.publicState.Logs(), receipts)

	elapsed := time.Since(time.Unix(0, work.header.Time.Int64()))
	glog.V(logger.Info).Infof("🔨  Mined block (#%v / %x) in %v", block.Number(), block.Hash().Bytes()[:4], elapsed)
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

func (minter *minter) createWork() *work {
	parent := minter.parent
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

	publicState, privateState, err := minter.chain.StateAt(minter.parent.Root())
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
	addrTxes := minter.withoutProposedTxes(allAddrTxes)
	return types.NewTransactionsByPriceAndNonce(addrTxes)
}

func (minter *minter) mintNewBlock() {
	minter.mu.Lock()
	defer minter.mu.Unlock()

	work := minter.createWork()
	transactions := minter.getTransactions()

	committedTxes, receipts, logs := work.commitTransactions(transactions, minter.chain)
	txCount := len(committedTxes)

	if txCount == 0 {
		glog.V(logger.Info).Infoln("Not generating new block since there are no pending transactions")
		return
	}

	if len(logs) > 0 {
		// make a copy, the state caches the logs and these logs get "upgraded" from pending to mined
		// logs by filling in the block hash when the block was mined by the local miner. This can
		// cause a race condition if a log was "upgraded" before the PendingLogsEvent is processed.
		cpy := make(vm.Logs, len(logs))
		for i, l := range logs {
			cpy[i] = new(vm.Log)
			*cpy[i] = *l
		}
		go func(logs vm.Logs, tcount int) {
			if len(logs) > 0 {
				minter.mux.Post(core.PendingLogsEvent{Logs: logs})
			}
			if tcount > 0 {
				minter.mux.Post(core.PendingStateEvent{})
			}
		}(cpy, txCount)
	}

	committedTxIs := make([]interface{}, len(committedTxes))
	for i, tx := range committedTxes {
		committedTxIs[i] = tx.Hash()
	}
	minter.proposedTxes.Add(committedTxIs...)

	header := work.header

	if atomic.LoadInt32(&minter.minting) == 1 {
		// commit state root after all state transitions.
		core.AccumulateRewards(work.publicState, header, nil)
		header.Root = work.publicState.IntermediateRoot()
	}

	// NOTE: < QuorumChain creates a signature here and puts it in header.Extra. >

	// update block hash since it is now available, but was not when the
	// receipt/log of individual transactions were created:
	headerHash := header.Hash()
	for _, r := range receipts {
		for _, l := range r.Logs {
			l.BlockHash = headerHash
		}
	}
	// NOTE: QuorumChain does not do the following:
	for _, log := range work.publicState.Logs() {
		log.BlockHash = headerHash
	}

	header.Bloom = types.CreateBloom(receipts)

	block := types.NewBlock(header, committedTxes, nil, receipts)
	work.Block = block

	glog.V(logger.Info).Infof("Generated next block #%v with [%d txns]", block.Number(), txCount)

	work.publicState.Commit()

	_, pubStateErr := work.publicState.Commit()
	if pubStateErr != nil {
		panic(fmt.Sprint("error committing public state: ", pubStateErr))
	}

	privateStateRoot, privStateErr := work.privateState.Commit()
	if privStateErr != nil {
		panic(fmt.Sprint("error committing private state: ", privStateErr))
	}

	if err := core.WritePrivateStateRoot(minter.chainDb, block.Root(), privateStateRoot); err != nil {
		panic(fmt.Sprint("error writing private state root: ", err))
	}

	//
	// TODO(bts): what to do for validation here now that AuxValidator has been
	//            removed in Quorum? looking at InsertChain might help here.
	//
	// We might not want/need any validation here. it will occur anyway once the
	// message has flown through raft.
	//
	// auxValidator := minter.eth.BlockChain().AuxValidator()
	// if err := core.ValidateHeader(minter.config, auxValidator, block.Header(), minter.parent.Header(), true, false); err != nil && err != core.BlockFutureErr {
	// 	panic(fmt.Sprint("Invalid header on mined block: ", err))
	// }

	if err := minter.chain.WriteDetachedBlock(block); err != nil {
		panic(fmt.Sprint("error writing block to chain: ", err))
	}

	// NOTE: We are currently writing txes, receipts, and the bloom filters
	// even though this block might not end up becoming the next head of the
	// chain.
	//
	// This puts transactions in a extra db for rpc
	core.WriteTransactions(minter.chainDb, block)
	// store the receipts
	core.WriteReceipts(minter.chainDb, receipts)
	// Write map map bloom filters
	core.WriteMipmapBloom(minter.chainDb, block.NumberU64(), receipts)

	minter.parent = block
	minter.unappliedBlocks.Append(block)

	minter.broadcastWork(work, receipts)
}

func (env *work) commitTransactions(txes *types.TransactionsByPriceAndNonce, bc *core.BlockChain) (types.Transactions, types.Receipts, vm.Logs) {
	var logs vm.Logs
	var committedTxes types.Transactions
	var receipts types.Receipts

	gp := new(core.GasPool).AddGas(env.header.GasLimit)
	txCount := 0

	for {
		tx := txes.Peek()
		if tx == nil {
			break
		}

		env.publicState.StartRecord(tx.Hash(), common.Hash{}, 0)

		receipt, txLogs, err := env.commitTransaction(tx, bc, gp)
		switch {
		case err != nil:
			if glog.V(logger.Detail) {
				glog.Infof("TX (%x) failed, will be removed: %v\n", tx.Hash().Bytes()[:4], err)
			}
			txes.Pop() // skip rest of txes from this account
		default:
			txCount++
			logs = append(logs, txLogs...)
			committedTxes = append(committedTxes, tx)
			receipts = append(receipts, receipt)
			txes.Shift()
		}
	}

	return committedTxes, receipts, logs
}

func (env *work) commitTransaction(tx *types.Transaction, bc *core.BlockChain, gp *core.GasPool) (*types.Receipt, vm.Logs, error) {
	publicSnapshot := env.publicState.Snapshot()
	privateSnapshot := env.privateState.Snapshot()

	// NOTE: QuorumChain disables forcing JIT here.

	receipt, logs, _, err := core.ApplyTransaction(env.config, bc, gp, env.publicState, env.privateState, env.header, tx, env.header.GasUsed, env.config.VmConfig)
	if err != nil {
		env.publicState.RevertToSnapshot(publicSnapshot)
		env.privateState.RevertToSnapshot(privateSnapshot)

		return nil, nil, err
	}

	return receipt, logs, nil
}
