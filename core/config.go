package core

// QuorumChainConfig is the configuration of Quorum blockchain
type QuorumChainConfig struct {
	revertReasonEnabled bool // if we should save the revert reasons in the Tx Receipts
	multiTenantEnabled  bool // if this blockchain supports multitenancy
}

// NewQuorumChainConfig creates new config for Quorum chain
func NewQuorumChainConfig(multiTenantEnabled, revertReasonEnabled bool) QuorumChainConfig {
	return QuorumChainConfig{
		multiTenantEnabled:  multiTenantEnabled,
		revertReasonEnabled: revertReasonEnabled,
	}
}

// RevertReasonEnabled returns true is revert reason feature is enabled
func (c QuorumChainConfig) RevertReasonEnabled() bool {
	return c.revertReasonEnabled
}

// MultiTenantEnabled returns true is multi-tenancy is enabled
func (c QuorumChainConfig) MultiTenantEnabled() bool {
	return c.multiTenantEnabled
}
