package raft

import (
	"os"

	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/wal"
	"github.com/coreos/etcd/wal/walpb"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
)

func (pm *ProtocolManager) openWAL(maybeRaftSnapshot *raftpb.Snapshot) *wal.WAL {
	if !wal.Exist(pm.waldir) {
		if err := os.Mkdir(pm.waldir, 0750); err != nil {
			glog.Fatalf("cannot create waldir: %v", err)
		}

		wal, err := wal.Create(pm.waldir, nil)
		if err != nil {
			glog.Fatalf("failed to create waldir: %v", err)
		}
		wal.Close()
	}

	glog.V(logger.Info).Infof("loading WAL")

	walsnap := walpb.Snapshot{}

	if maybeRaftSnapshot != nil {
		walsnap.Index, walsnap.Term = maybeRaftSnapshot.Metadata.Index, maybeRaftSnapshot.Metadata.Term
	}

	wal, err := wal.Open(pm.waldir, walsnap)
	if err != nil {
		glog.Fatalf("error loading WAL: %v", err)
	}

	return wal
}

func (pm *ProtocolManager) replayWAL(maybeRaftSnapshot *raftpb.Snapshot) *wal.WAL {
	glog.V(logger.Info).Infoln("replaying WAL")
	wal := pm.openWAL(maybeRaftSnapshot)

	_, hardState, entries, err := wal.ReadAll()
	if err != nil {
		glog.Fatalf("failed to read WAL: %v", err)
	}

	pm.raftStorage.SetHardState(hardState)
	pm.raftStorage.Append(entries)

	return wal
}
