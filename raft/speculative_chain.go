package raft

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"

	"gopkg.in/fatih/set.v0"
	lane "gopkg.in/oleiade/lane.v1"
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
	parent                     *types.Block
	unappliedBlocks            *lane.Deque
	expectedInvalidBlockHashes *set.Set // This is thread-safe.
}

func (chain *speculativeChain) clear(parent *types.Block) {
	chain.parent = parent
	chain.unappliedBlocks = lane.NewDeque()
	chain.expectedInvalidBlockHashes = set.New()
}

// Append a new speculative block
func (chain *speculativeChain) extend(block *types.Block) {
	chain.parent = block
	chain.unappliedBlocks.Append(block)
}

// Set the parent of the speculative chain
//
// Note: This is only called when not minter
func (chain *speculativeChain) setParent(block *types.Block) {
	chain.parent = block
}

// Accept this block, removing it from the head of the speculative chain
func (chain *speculativeChain) accept(acceptedBlock *types.Block) (foundExpected bool) {
	earliestProposedI := chain.unappliedBlocks.Shift()
	var earliestProposed *types.Block
	if nil != earliestProposedI {
		earliestProposed = earliestProposedI.(*types.Block)
	}

	expectedBlock := earliestProposed == nil || earliestProposed.Hash() == acceptedBlock.Hash()
	if !expectedBlock {
		glog.V(logger.Warn).Infof("Another node minted %x; Clearing speculative state\n", acceptedBlock.Hash())

		chain.clear(acceptedBlock)
	}
	return expectedBlock
}

// Remove all blocks in the chain from the specified one until the end
func (chain *speculativeChain) unwindFrom(invalidHash common.Hash, headBlock *types.Block, removeProposedTxes func(*types.Block)) {

	// check our guard to see if this is a (descendant) block we're
	// expected to be ruled invalid. if we find it, remove from the guard
	if chain.expectedInvalidBlockHashes.Has(invalidHash) {
		glog.V(logger.Warn).Infof("Removing expected-invalid block %x from guard.\n", invalidHash)

		chain.expectedInvalidBlockHashes.Remove(invalidHash)

		return
	}

	// pop from the RHS repeatedly, updating minter.parent each time. if not
	// our block, add to guard. in all cases, call removeProposedTxes
	for {
		currBlockI := chain.unappliedBlocks.Pop()

		if nil == currBlockI {
			glog.V(logger.Warn).Infof("(Popped all blocks from queue.)\n")

			break
		}

		currBlock := currBlockI.(*types.Block)

		glog.V(logger.Info).Infof("Popped block %x from queue RHS.\n", currBlock.Hash())

		// Maintain invariant: the parent always points the last speculative block or the head of the blockchain
		// if there are not speculative blocks.
		if speculativeParentI := chain.unappliedBlocks.Last(); nil != speculativeParentI {
			chain.parent = speculativeParentI.(*types.Block)
		} else {
			chain.parent = headBlock
		}

		// Each time we remove a block call back to the minter to remove proposed txes from that block
		removeProposedTxes(currBlock)

		if currBlock.Hash() != invalidHash {
			glog.V(logger.Warn).Infof("Haven't yet found block %x; adding descendent %x to guard.\n", invalidHash, currBlock.Hash())

			chain.expectedInvalidBlockHashes.Add(currBlock.Hash())
		} else {
			break
		}
	}
}
