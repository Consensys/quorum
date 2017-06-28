package raft

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"strconv"
)

var (
	noFsync = &opt.WriteOptions{
		NoWriteMerge: false,
		Sync:         false,
	}

	mustFsync = &opt.WriteOptions{
		NoWriteMerge: false,
		Sync:         true,
	}
)

func openQuorumRaftDb(path string) (db *leveldb.DB, err error) {
	// Open the db and recover any potential corruptions
	db, err = leveldb.OpenFile(path, &opt.Options{
		OpenFilesCacheCapacity: -1, // -1 means 0??
		BlockCacheCapacity:     -1,
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(path, nil)
	}
	return
}

func (pm *ProtocolManager) loadAppliedIndex() uint64 {
	dat, err := pm.quorumRaftDb.Get(appliedDbKey, nil)
	var lastAppliedIndex uint64
	if err == errors.ErrNotFound {
		lastAppliedIndex = 0
	} else if err != nil {
		fatalf("loadAppliedIndex error: %s", err)
	} else {
		lastAppliedIndex = binary.LittleEndian.Uint64(dat)
	}

	log.Info("Persistent applied index load", "last applied index", lastAppliedIndex)
	pm.appliedIndex = lastAppliedIndex
	return lastAppliedIndex
}

func (pm *ProtocolManager) writeAppliedIndex(index uint64) {
	log.Info("Persistent applied index write", "index", index)
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	pm.quorumRaftDb.Put(appliedDbKey, buf, noFsync)
}

func (pm *ProtocolManager) loadPeerAddress(raftId uint16) *Address {
	peerUrlKey := []byte(peerUrlKeyPrefix + strconv.Itoa(int(raftId)))
	value, err := pm.quorumRaftDb.Get(peerUrlKey, nil)
	if err != nil {
		fatalf("failed to read address for raft peer %d from leveldb: %v", raftId, err)
	}

	return bytesToAddress(value)
}

func (pm *ProtocolManager) writePeerAddressBytes(raftId uint16, value []byte) {
	key := []byte(peerUrlKeyPrefix + strconv.Itoa(int(raftId)))

	pm.quorumRaftDb.Put(key, value, mustFsync)
}
