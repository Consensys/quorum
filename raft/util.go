package raft

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"io"
	"os"
	"runtime"
	"strconv"
)

// TODO: this is just copied over from cmd/utils/cmd.go. dedupe

// Fatalf formats a message to standard error and exits the program.
// The message is also printed to standard output if standard error
// is redirected to a different file.
func fatalf(format string, args ...interface{}) {
	w := io.MultiWriter(os.Stdout, os.Stderr)
	if runtime.GOOS == "windows" {
		// The SameFile check below doesn't work on Windows.
		// stdout is unlikely to get redirected though, so just print there.
		w = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			w = os.Stderr
		}
	}
	fmt.Fprintf(w, "Fatal: "+format+"\n", args...)
	os.Exit(1)
}

// maps a node id (512 bit public key of the node) to a 64 bit raft id required by the etcd raft core.
// note: enode ids in the network / static-nodes.json need to have unique 15 char prefixes.
func nodeIdToRaftId(nodeId string) (uint64, error) {
	log.Info("Converting node id to raft id", "node id", nodeId)
	nodeIdChars := []rune(nodeId)
	// a rune is an alias for int32
	// raft id is an uint64, uint64 is the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	// 0x7FFFFFFFFFFFFFFF == 9223372036854775807 (dec)
	// len("7FFFFFFFFFFFFFFF")  == 16, so take the first 15 chars (16-1) of the node id to make sure it fits into an uint64
	idShort := string(nodeIdChars[0:15])
	log.Info("raft idshort ", "idShort", idShort)
	raftId, err := strconv.ParseUint(idShort, 16, 64)
	log.Info("raft id as uint64", "raftId uint64", raftId)
	if err != nil {
		log.Error("Error converting node id to uint64 raft id", "err", err)
		return 0, err
	}
	return raftId, nil
}

// used to convert a uint64 to a string, needed for displaying the raftId properly in Javascript (raft/api.go), e.g.
// raft.addPeer(enodeUrl) as JS Number and uint64 are not compatible.
func RaftIdToString(raftId uint64) string {
	return fmt.Sprintf("%d", raftId)
}
