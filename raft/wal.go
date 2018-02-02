package raft

import (
	"os"

	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/wal"
	"github.com/coreos/etcd/wal/walpb"
	"github.com/ethereum/go-ethereum/log"
)

func (pm *ProtocolManager) openWAL(maybeRaftSnapshot *raftpb.Snapshot) *wal.WAL {
	if !wal.Exist(pm.waldir) {
		if err := os.Mkdir(pm.waldir, 0750); err != nil {
			fatalf("cannot create waldir: %s", err)
		}

		wal, err := wal.Create(pm.waldir, nil)
		if err != nil {
			fatalf("failed to create waldir: %s", err)
		}
		wal.Close()
	}

	walsnap := walpb.Snapshot{}

	log.Info("loading WAL", "term", walsnap.Term, "index", walsnap.Index)

	if maybeRaftSnapshot != nil {
		walsnap.Index, walsnap.Term = maybeRaftSnapshot.Metadata.Index, maybeRaftSnapshot.Metadata.Term
	}

	wal, err := wal.Open(pm.waldir, walsnap)
	if err != nil {
		fatalf("error loading WAL: %s", err)
	}

	return wal
}

func (pm *ProtocolManager) replayWAL(maybeRaftSnapshot *raftpb.Snapshot) *wal.WAL {
	log.Info("replaying WAL")
	wal := pm.openWAL(maybeRaftSnapshot)

	_, hardState, entries, err := wal.ReadAll()
	if err != nil {
		fatalf("failed to read WAL: %s", err)
	}

	pm.raftStorage.SetHardState(hardState)
	pm.raftStorage.Append(entries)

	return wal
}
