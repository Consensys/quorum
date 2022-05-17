package raft

import (
	mapset "github.com/deckarep/golang-set"
	"gopkg.in/oleiade/lane.v1"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// The speculative chain represents blocks that we have minted which haven't been accepted into the chain yet, building
// on each other in a chain. It has three basic operations:
// * add new block to end
// * accept / remove oldest block
// * unwind / remove invalid blocks to the end
//
// Additionally:
// * clear state when we stop minting
// * set the parent when we're not minting (so it's always current)
type speculativeChain struct {
	head                       *types.Block
	unappliedBlocks            *lane.Deque
	expectedInvalidBlockHashes mapset.Set // This is thread-safe. This set is referred to as our "guard" below.
	proposedTxes               mapset.Set // This is thread-safe.
}

func newSpeculativeChain() *speculativeChain {
	return &speculativeChain{
		head:                       nil,
		unappliedBlocks:            lane.NewDeque(),
		expectedInvalidBlockHashes: mapset.NewSet(),
		proposedTxes:               mapset.NewSet(),
	}
}

func (chain *speculativeChain) clear(block *types.Block) {
	chain.head = block
	chain.unappliedBlocks = lane.NewDeque()
	chain.expectedInvalidBlockHashes.Clear()
	chain.proposedTxes.Clear()
}

// Append a new speculative block
func (chain *speculativeChain) extend(block *types.Block) {
	chain.head = block
	chain.recordProposedTransactions(block.Transactions())
	chain.unappliedBlocks.Append(block)
}

// Set the parent of the speculative chain
//
// Note: This is only called when not minter
func (chain *speculativeChain) setHead(block *types.Block) {
	chain.head = block
}

// Accept this block, removing it from the speculative chain
func (chain *speculativeChain) accept(acceptedBlock *types.Block) {
	earliestProposedI := chain.unappliedBlocks.Shift()
	var earliestProposed *types.Block
	if nil != earliestProposedI {
		earliestProposed = earliestProposedI.(*types.Block)
	}

	// There are three possible scenarios:
	// 1. We don't have a record of this block (or any proposed blocks), meaning someone else minted it and we should
	//    add it as the new head of our speculative chain. New blocks from the old leader are still coming in.
	// 2. This block was the first outstanding one we proposed.
	// 3. This block is different from the block we proposed, (also) meaning new blocks are still coming in from the old
	//    leader, but unlike the first scenario, we need to clear all of the speculative chain state because the
	//    `acceptedBlock` takes precedence over our speculative state.
	if earliestProposed == nil {
		chain.head = acceptedBlock
	} else if expectedBlock := earliestProposed.Hash() == acceptedBlock.Hash(); expectedBlock {
		// Remove the txes in this accepted block from our blacklist.
		chain.removeProposedTxes(acceptedBlock)
	} else {
		log.Info("Another node minted; Clearing speculative state", "block", acceptedBlock.Hash())

		chain.clear(acceptedBlock)
	}
}

// Remove all blocks in the chain from the specified one until the end
func (chain *speculativeChain) unwindFrom(invalidHash common.Hash, headBlock *types.Block) {

	// check our "guard" to see if this is a (descendant) block we're
	// expected to be ruled invalid. if we find it, remove from the guard
	if chain.expectedInvalidBlockHashes.Contains(invalidHash) {
		log.Info("Removing expected-invalid block from guard.", "block", invalidHash)

		chain.expectedInvalidBlockHashes.Remove(invalidHash)

		return
	}

	// pop from the RHS repeatedly, updating minter.parent each time. if not
	// our block, add to guard. in all cases, call removeProposedTxes
	for {
		currBlockI := chain.unappliedBlocks.Pop()

		if nil == currBlockI {
			log.Info("(Popped all blocks from queue.)")

			break
		}

		currBlock := currBlockI.(*types.Block)

		log.Info("Popped block from queue RHS.", "block", currBlock.Hash())

		// Maintain invariant: the parent always points the last speculative block or the head of the blockchain
		// if there are not speculative blocks.
		if speculativeParentI := chain.unappliedBlocks.Last(); nil != speculativeParentI {
			chain.head = speculativeParentI.(*types.Block)
		} else {
			chain.head = headBlock
		}

		chain.removeProposedTxes(currBlock)

		if currBlock.Hash() != invalidHash {
			log.Info("Haven't yet found block; adding descendent to guard.\n", "invalid block", invalidHash, "descendant", currBlock.Hash())

			chain.expectedInvalidBlockHashes.Add(currBlock.Hash())
		} else {
			break
		}
	}
}

// We keep track of txes we've put in all newly-mined blocks since the last
// ChainHeadEvent, and filter them out so that we don't try to create blocks
// with the same transactions. This is necessary because the TX pool will keep
// supplying us these transactions until they are in the chain (after having
// flown through raft).
func (chain *speculativeChain) recordProposedTransactions(txes types.Transactions) {
	txHashIs := make([]interface{}, len(txes))
	for i, tx := range txes {
		txHashIs[i] = tx.Hash()
	}
	for _, i := range txHashIs {
		chain.proposedTxes.Add(i)
	}
}

// Removes txes in block from our "blacklist" of "proposed tx" hashes. When we
// create a new block and use txes from the tx pool, we ignore those that we
// have already used ("proposed"), but that haven't yet officially made it into
// the chain yet.
//
// It's important to remove hashes from this blacklist (once we know we don't
// need them in there anymore) so that it doesn't grow endlessly.
func (chain *speculativeChain) removeProposedTxes(block *types.Block) {
	minedTxes := block.Transactions()
	minedTxInterfaces := make([]interface{}, len(minedTxes))
	for i, tx := range minedTxes {
		minedTxInterfaces[i] = tx.Hash()
	}

	// NOTE: we are using a thread-safe Set here, so it's fine if we access this
	// here and in mintNewBlock concurrently. using a finer-grained set-specific
	// lock here is preferable, because mintNewBlock holds its locks for a
	// nontrivial amount of time.
	for _, i := range minedTxInterfaces {
		chain.proposedTxes.Remove(i)
	}
}

func (chain *speculativeChain) withoutProposedTxes(addrTxes AddressTxes) AddressTxes {
	newMap := make(AddressTxes)

	for addr, txes := range addrTxes {
		filteredTxes := make(types.Transactions, 0)
		for _, tx := range txes {
			if !chain.proposedTxes.Contains(tx.Hash()) {
				filteredTxes = append(filteredTxes, tx)
			}
		}
		if len(filteredTxes) > 0 {
			newMap[addr] = filteredTxes
		}
	}

	return newMap
}
