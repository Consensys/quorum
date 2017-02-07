package gethRaft

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"

	"github.com/syndtr/goleveldb/leveldb/errors"
)

func (pm *ProtocolManager) loadAppliedIndex() uint64 {
	dat, err := pm.appliedDb.Get(appliedDbKey, nil)
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
	pm.appliedDb.Put(appliedDbKey, buf, nil)
}

func (pm *ProtocolManager) loadPeerUrl(nodeId uint64) string {
	peerUrlKey := []byte(peerUrlKeyPrefix + string(nodeId))
	value, err := pm.appliedDb.Get(peerUrlKey, nil)
	if err != nil {
		glog.Fatalf("failed to read peer url for peer %d from leveldb: %v", nodeId, err)
	}
	return string(value)
}

func (pm *ProtocolManager) writePeerUrl(nodeId uint64, url string) {
	key := []byte(peerUrlKeyPrefix + string(nodeId))
	value := []byte(url)

	pm.appliedDb.Put(key, value, nil)
}
