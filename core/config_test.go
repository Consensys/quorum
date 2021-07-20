package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigRevertReasonEnabled(t *testing.T) {
	assert.False(t, revertReasonEnabled(nil))
	assert.False(t, revertReasonEnabled([]QuorumChainConfig{}))
	assert.False(t, revertReasonEnabled([]QuorumChainConfig{{RevertReasonEnabled: false}}))
	assert.True(t, revertReasonEnabled([]QuorumChainConfig{{RevertReasonEnabled: true}}))
	assert.True(t, revertReasonEnabled([]QuorumChainConfig{{RevertReasonEnabled: false}, {RevertReasonEnabled: true}, {RevertReasonEnabled: false}}))
}

func TestConfigMultiTenantEnabled(t *testing.T) {
	assert.False(t, multiTenantEnabled(nil))
	assert.False(t, multiTenantEnabled([]QuorumChainConfig{}))
	assert.False(t, multiTenantEnabled([]QuorumChainConfig{{MultiTenantEnabled: false}}))
	assert.True(t, multiTenantEnabled([]QuorumChainConfig{{MultiTenantEnabled: true}}))
	assert.True(t, multiTenantEnabled([]QuorumChainConfig{{MultiTenantEnabled: false}, {MultiTenantEnabled: true}, {MultiTenantEnabled: false}}))
}
