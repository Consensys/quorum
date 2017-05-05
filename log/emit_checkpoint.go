package log

const (
	TxCreated          = "TX-CREATED"
	TxAccepted         = "TX-ACCEPTED"
	BecameMinter       = "BECAME-MINTER"
	BecameVerifier     = "BECAME-VERIFIER"
	BlockCreated       = "BLOCK-CREATED"
	BlockVotingStarted = "BLOCK-VOTING-STARTED"
)

var DoEmitCheckpoints = false

func EmitCheckpoint(checkpointName string, logValues ...interface{}) {
	if DoEmitCheckpoints {
		Info("QUORUM-CHECKPOINT", "name", checkpointName, "data", logValues)
	}
}
