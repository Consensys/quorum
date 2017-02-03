package gethRaft

import (
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/wal/walpb"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
)

func (pm *ProtocolManager) saveSnap(snap raftpb.Snapshot) error {
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

func (pm *ProtocolManager) publishSnapshot(snap raftpb.Snapshot) {
	if raft.IsEmptySnap(snap) {
		return
	}

	glog.V(logger.Info).Infof("publishing snapshot at index %d", pm.snapshotIndex)
	defer glog.V(logger.Info).Infof("finished publishing snapshot at index %d", pm.snapshotIndex)

	if snap.Metadata.Index <= pm.appliedIndex {
		glog.Fatalf("snapshot index [%d] should > progress.appliedIndex [%d] + 1", snap.Metadata.Index, pm.appliedIndex)
	}

	pm.logCommandC <- LoadSnapshot{}

	pm.confState = snap.Metadata.ConfState
	pm.snapshotIndex = snap.Metadata.Index
	pm.appliedIndex = snap.Metadata.Index
}

func (pm *ProtocolManager) maybeTriggerSnapshot() {
	if pm.appliedIndex-pm.snapshotIndex <= defaultSnapCount {
		return
	}

	glog.V(logger.Info).Infof("start snapshot [applied index: %d | last snapshot index: %d]", pm.appliedIndex, pm.snapshotIndex)
	snapData := pm.blockchain.CurrentBlock().Hash().Bytes()
	snap, err := pm.raftStorage.CreateSnapshot(pm.appliedIndex, &pm.confState, snapData)
	if err != nil {
		panic(err)
	}
	if err := pm.saveSnap(snap); err != nil {
		panic(err)
	}

	// Discard all log entries prior to appliedIndex.
	if err := pm.raftStorage.Compact(pm.appliedIndex); err != nil {
		panic(err)
	}

	glog.V(logger.Info).Infof("compacted log at index %d", pm.appliedIndex)
	pm.snapshotIndex = pm.appliedIndex
}
