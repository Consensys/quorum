package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigBool(t *testing.T) {
	assert.False(t, revertReasonSaved(nil))
	assert.False(t, revertReasonSaved([]BlockchainQuorumConfig{}))
	assert.False(t, revertReasonSaved([]BlockchainQuorumConfig{{SaveRevertReason: false}}))
	assert.True(t, revertReasonSaved([]BlockchainQuorumConfig{{SaveRevertReason: true}}))
	assert.True(t, revertReasonSaved([]BlockchainQuorumConfig{{SaveRevertReason: false}, {SaveRevertReason: true}, {SaveRevertReason: false}}))
}
