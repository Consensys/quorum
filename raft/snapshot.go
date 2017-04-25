package raft

import (
	"fmt"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
	"github.com/coreos/etcd/wal/walpb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
	"math/big"
	"time"
)

func (pm *ProtocolManager) saveSnapshot(snap raftpb.Snapshot) error {
	if err := pm.snapshotter.SaveSnap(snap); err != nil {
		return err
	}

	walSnap := walpb.Snapshot{
		Index: snap.Metadata.Index,
		Term:  snap.Metadata.Term,
	}

	if err := pm.wal.SaveSnapshot(walSnap); err != nil {
		return err
	}

	return pm.wal.ReleaseLockTo(snap.Metadata.Index)
}

func (pm *ProtocolManager) maybeTriggerSnapshot() {
	if pm.appliedIndex-pm.snapshotIndex < snapshotPeriod {
		return
	}

	pm.triggerSnapshot()
}

func (pm *ProtocolManager) triggerSnapshot() {
	glog.V(logger.Info).Infof("start snapshot [applied index: %d | last snapshot index: %d]", pm.appliedIndex, pm.snapshotIndex)
	snapData := pm.blockchain.CurrentBlock().Hash().Bytes()
	snap, err := pm.raftStorage.CreateSnapshot(pm.appliedIndex, &pm.confState, snapData)
	if err != nil {
		panic(err)
	}
	if err := pm.saveSnapshot(snap); err != nil {
		panic(err)
	}
	// Discard all log entries prior to appliedIndex.
	if err := pm.raftStorage.Compact(pm.appliedIndex); err != nil {
		panic(err)
	}
	glog.V(logger.Info).Infof("compacted log at index %d", pm.appliedIndex)
	pm.snapshotIndex = pm.appliedIndex
}

// For persisting cluster membership changes correctly, we need to trigger a
// snapshot before advancing our persisted appliedIndex in LevelDB.
//
// See handling of EntryConfChange entries in raft/handler.go for details.
func (pm *ProtocolManager) triggerSnapshotWithNextIndex(index uint64) {
	pm.appliedIndex = index
	pm.triggerSnapshot()
}

func (pm *ProtocolManager) loadSnapshot() *raftpb.Snapshot {
	snapshot, err := pm.snapshotter.Load()
	if err != nil && err != snap.ErrNoSnapshot {
		glog.Fatalf("error loading snapshot: %v", err)
	}

	return snapshot
}

func (pm *ProtocolManager) applySnapshot(snap raftpb.Snapshot) {
	glog.V(logger.Info).Infof("applying snapshot to raft storage")
	if err := pm.raftStorage.ApplySnapshot(snap); err != nil {
		glog.Fatalln("failed to apply snapshot: ", err)
	}

	latestBlockHash := common.BytesToHash(snap.Data)

	preSyncHead := pm.blockchain.CurrentBlock()

	glog.V(logger.Info).Infof("before sync, chain head is at block %x", preSyncHead.Hash())

	if latestBlock := pm.blockchain.GetBlockByHash(latestBlockHash); latestBlock == nil {
	syncing:
		for {
			for peerId, peer := range pm.peers {
				glog.V(logger.Info).Infof("synchronizing with peer %v up to block %x", peerId, latestBlockHash)

				peerId := peer.p2pNode.ID.String()
				peerIdPrefix := fmt.Sprintf("%x", peer.p2pNode.ID[:8])

				if err := pm.downloader.Synchronise(peerIdPrefix, latestBlockHash, big.NewInt(0), downloader.BoundedFullSync); err != nil {
					glog.V(logger.Warn).Infof("failed to synchronize with peer %v", peerId)

					time.Sleep(500 * time.Millisecond)
				} else {
					break syncing
				}
			}
		}

		pm.logNewlyAcceptedTransactions(preSyncHead)

		glog.V(logger.Info).Infof("%s: %x\n", chainExtensionMessage, pm.blockchain.CurrentBlock().Hash())
	} else {
		glog.V(logger.Info).Infof("already caught up; no need to synchronize")
	}

	snapMeta := snap.Metadata

	pm.confState = snapMeta.ConfState
	pm.snapshotIndex = snapMeta.Index
	pm.advanceAppliedIndex(snapMeta.Index)
}
func (pm *ProtocolManager) logNewlyAcceptedTransactions(preSyncHead *types.Block) {
	newHead := pm.blockchain.CurrentBlock()
	numBlocks := newHead.NumberU64() - preSyncHead.NumberU64()
	blocks := make([]*types.Block, numBlocks)
	currBlock := newHead
	blocksSeen := 0
	for currBlock.Hash() != preSyncHead.Hash() {
		blocks[int(numBlocks)-(1+blocksSeen)] = currBlock

		blocksSeen += 1
		currBlock = pm.blockchain.GetBlockByHash(currBlock.ParentHash())
	}
	for _, block := range blocks {
		for _, tx := range block.Transactions() {
			logger.LogRaftCheckpoint(logger.TxAccepted, tx.Hash().Hex())
		}
	}
}
