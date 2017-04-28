package raft

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

var (
	noFsync = &opt.WriteOptions{
		NoWriteMerge: false,
		Sync:         false,
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
		glog.Fatalln(err)
	} else {
		lastAppliedIndex = binary.LittleEndian.Uint64(dat)
	}

	pm.mu.Lock()
	pm.appliedIndex = lastAppliedIndex
	pm.mu.Unlock()

	glog.V(logger.Info).Infof("loaded the latest applied index: %d", lastAppliedIndex)

	return lastAppliedIndex
}

func (pm *ProtocolManager) writeAppliedIndex(index uint64) {
	glog.V(logger.Info).Infof("persisted the latest applied index: %d", index)
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	pm.quorumRaftDb.Put(appliedDbKey, buf, noFsync)
}
