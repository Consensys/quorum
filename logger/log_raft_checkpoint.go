package logger

import (
	"github.com/ethereum/go-ethereum/logger/glog"
)

const (
	TxCreated      = "TX-CREATED"
	TxAccepted     = "TX-ACCEPTED"
	BecameMinter   = "BECAME-MINTER"
	BecameVerifier = "BECAME-VERIFIER"
)

var DoLogRaft = false

func LogRaftCheckpoint(checkpointName string, logValues ...interface{}) {
	if DoLogRaft {
		glog.V(Info).Infof("RAFT-CHECKPOINT %s %v\n", checkpointName, logValues)
	}
}
