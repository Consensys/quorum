package log

const (
	TxCreated          = "TX-CREATED"
	TxAccepted         = "TX-ACCEPTED"
	TxCompleted        = "TX-COMPLETED"
	BecameMinter       = "BECAME-MINTER"
	BecameVerifier     = "BECAME-VERIFIER"
	BecameLearner      = "BECAME-LEARNER"
	BlockCreated       = "BLOCK-CREATED"
	BlockVotingStarted = "BLOCK-VOTING-STARTED"
)

var DoEmitCheckpoints = false

func EmitCheckpoint(checkpointName string, logValues ...interface{}) {
	args := []interface{}{"name", checkpointName}
	args = append(args, logValues...)
	if DoEmitCheckpoints {
		Info("QUORUM-CHECKPOINT", args...)
	}
}
