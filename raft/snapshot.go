package raft

import (
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
	"github.com/coreos/etcd/wal/walpb"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
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
	if err := pm.raftStorage.ApplySnapshot(snap); err != nil {
		glog.Fatalln("failed to apply snapshot: ", err)
	}

	snapMeta := snap.Metadata

	pm.confState = snapMeta.ConfState
	pm.snapshotIndex = snapMeta.Index
	pm.advanceAppliedIndex(snapMeta.Index)
}
