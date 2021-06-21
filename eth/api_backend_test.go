package eth

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/node"
	"github.com/stretchr/testify/require"
)

func TestEthAPIBackend_SubscribePendingLogsEvent(t *testing.T) {
	conf := &Config{
		RaftMode: false,
	}
	stack, err := node.New(&node.Config{})
	if err != nil {
		t.Fatalf("failed to create node, err = %v", err)
	}
	eth, err := New(stack, conf)
	if err != nil {
		t.Fatalf("failed to create eth service, err = %v", err)
	}

	b := &EthAPIBackend{
		eth: eth,
	}

	ch := make(chan []*types.Log, 1)

	_ = b.SubscribePendingLogsEvent(ch)

	recipientCount := eth.ConsensusServicePendingLogsFeed().Send([]*types.Log{})

	require.Zero(t, recipientCount, "not using consensus service so its event feed should not have subscribers")
	require.Zero(t, len(ch), "not using consensus service so subscribed channel should not have received event")
}

func TestEthAPIBackend_SubscribePendingLogsEvent_SubscribesToConsensusServiceFeed(t *testing.T) {
	conf := &Config{
		RaftMode: true,
	}
	stack, err := node.New(&node.Config{})
	if err != nil {
		t.Fatalf("failed to create node, err = %v", err)
	}
	eth, err := New(stack, conf)
	if err != nil {
		t.Fatalf("failed to create eth service, err = %v ", err)
	}

	b := &EthAPIBackend{
		eth: eth,
	}

	ch := make(chan []*types.Log, 1)

	_ = b.SubscribePendingLogsEvent(ch)

	recipientCount := eth.ConsensusServicePendingLogsFeed().Send([]*types.Log{})

	require.NotZero(t, recipientCount, "consensus service in use so its event feed should have subscribers")
	require.Equal(t, 1, len(ch), "consensus service in use so subscribed channel should have received event")
}
