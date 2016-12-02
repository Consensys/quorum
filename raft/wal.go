package gethRaft

import (
	"log"
	"os"

	"github.com/coreos/etcd/wal"
	"github.com/coreos/etcd/wal/walpb"
)

func (pm *ProtocolManager) openWAL() *wal.WAL {
	if !wal.Exist(pm.waldir) {
		if err := os.Mkdir(pm.waldir, 0750); err != nil {
			log.Fatalf("cannot create waldir (%v)", err)
		}

		wal, err := wal.Create(pm.waldir, nil)
		if err != nil {
			log.Fatalf("create wal error (%v)", err)
		}
		wal.Close()
	}

	wal, err := wal.Open(pm.waldir, walpb.Snapshot{})
	if err != nil {
		log.Fatalf("error loading WAL (%v)", err)
	}

	return wal
}

func (pm *ProtocolManager) replayWAL() *wal.WAL {
	wal := pm.openWAL()
	_, st, ents, err := wal.ReadAll()
	if err != nil {
		log.Fatalf("failed to read WAL (%v)", err)
	}

	// append to storage so raft starts at the right place in log
	pm.raftStorage.Append(ents)
	pm.raftStorage.SetHardState(st)
	return wal
}
