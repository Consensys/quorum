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
		glog.Fatalln(err)
	} else {
		lastAppliedIndex = binary.LittleEndian.Uint64(dat)
	}

	glog.V(logger.Info).Infof("Persistent applied index load: %d", lastAppliedIndex)
	pm.appliedIndex = lastAppliedIndex
	return lastAppliedIndex
}

func (pm *ProtocolManager) writeAppliedIndex(index uint64) {
	glog.V(logger.Info).Infof("Persistent applied index write: %d", index)
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	pm.quorumRaftDb.Put(appliedDbKey, buf, noFsync)
}

func (pm *ProtocolManager) loadPeerUrl(nodeId uint64) string {
	peerUrlKey := []byte(peerUrlKeyPrefix + string(nodeId))
	value, err := pm.quorumRaftDb.Get(peerUrlKey, nil)
	if err != nil {
		glog.Fatalf("failed to read peer url for peer %d from leveldb: %v", nodeId, err)
	}
	return string(value)
}

func (pm *ProtocolManager) writePeerUrl(nodeId uint64, url string) {
	key := []byte(peerUrlKeyPrefix + string(nodeId))
	value := []byte(url)

	pm.quorumRaftDb.Put(key, value, mustFsync)
}
