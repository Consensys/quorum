package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigRevertReasonEnabled(t *testing.T) {
	assert.False(t, revertReasonEnabled(nil))
	assert.False(t, revertReasonEnabled([]BlockchainQuorumConfig{}))
	assert.False(t, revertReasonEnabled([]BlockchainQuorumConfig{{RevertReasonEnabled: false}}))
	assert.True(t, revertReasonEnabled([]BlockchainQuorumConfig{{RevertReasonEnabled: true}}))
	assert.True(t, revertReasonEnabled([]BlockchainQuorumConfig{{RevertReasonEnabled: false}, {RevertReasonEnabled: true}, {RevertReasonEnabled: false}}))
}

func TestConfigMultiTenantEnabled(t *testing.T) {
	assert.False(t, multiTenantEnabled(nil))
	assert.False(t, multiTenantEnabled([]BlockchainQuorumConfig{}))
	assert.False(t, multiTenantEnabled([]BlockchainQuorumConfig{{MultiTenantEnabled: false}}))
	assert.True(t, multiTenantEnabled([]BlockchainQuorumConfig{{MultiTenantEnabled: true}}))
	assert.True(t, multiTenantEnabled([]BlockchainQuorumConfig{{MultiTenantEnabled: false}, {MultiTenantEnabled: true}, {MultiTenantEnabled: false}}))
}
