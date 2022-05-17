package raft

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"
)

func Test_New_RegistersEthServicePendingLogsFeed(t *testing.T) {
	conf := &eth.Config{
		RaftMode: true,
	}
	stack, err := node.New(&node.Config{})
	if err != nil {
		t.Fatalf("failed to create node, err = %v", err)
	}
	ethService, err := eth.New(stack, conf)
	if err != nil {
		t.Fatalf("failed to create eth service, err = %v", err)
	}

	tmpWorkingDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpWorkingDir)
	}()

	raftService, err := New(stack, &params.ChainConfig{}, 0, 0, false, time.Second, ethService, nil, tmpWorkingDir, false)
	if err != nil {
		t.Fatalf("failed to create raft service, err = %v", err)
	}

	require.Equal(t, ethService.ConsensusServicePendingLogsFeed(), raftService.pendingLogsFeed, "raft service has not been set up with Ethereum service's consensusServicePendingLogsFeed")
}
