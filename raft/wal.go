package raft

import (
	"os"

	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/wal"
	"github.com/coreos/etcd/wal/walpb"
	"github.com/ethereum/go-ethereum/log"
)

func (pm *ProtocolManager) openWAL(maybeSnapshot *raftpb.Snapshot) *wal.WAL {
	if !wal.Exist(pm.waldir) {
		if err := os.Mkdir(pm.waldir, 0750); err != nil {
			fatalf("cannot create waldir (%v)", err)
		}

		wal, err := wal.Create(pm.waldir, nil)
		if err != nil {
			fatalf("failed to create waldir (%v)", err)
		}
		wal.Close()
	}

	walsnap := walpb.Snapshot{}
	if maybeSnapshot != nil {
		walsnap.Index = maybeSnapshot.Metadata.Index
		walsnap.Term = maybeSnapshot.Metadata.Term
	}

	log.Info("loading WAL at term %d and index %d", walsnap.Term, walsnap.Index)

	wal, err := wal.Open(pm.waldir, walsnap)
	if err != nil {
		fatalf("error loading WAL (%v)", err)
	}

	return wal
}

func (pm *ProtocolManager) replayWAL() *wal.WAL {
	log.Info("replaying WAL")
	maybeSnapshot := pm.loadSnapshot()
	wal := pm.openWAL(maybeSnapshot)

	_, hardState, entries, err := wal.ReadAll()
	if err != nil {
		fatalf("failed to read WAL (%v)", err)
	}

	if maybeSnapshot != nil {
		pm.applySnapshot(*maybeSnapshot)
	}

	pm.raftStorage.SetHardState(hardState)
	pm.raftStorage.Append(entries)

	return wal
}
