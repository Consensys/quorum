package core

// QuorumChainConfig is the configuration of Quorum blockchain
type QuorumChainConfig struct {
	revertReasonEnabled     bool // if we should save the revert reasons in the Tx Receipts
	multiTenantEnabled      bool // if this blockchain supports multitenancy
	privacyMarkerEnabled    bool // if the privacy marker is activated
	privateTrieCacheEnabled bool // if the private trie cache is enabled
}

// NewQuorumChainConfig creates new config for Quorum chain
func NewQuorumChainConfig(multiTenantEnabled, revertReasonEnabled, privacyMarkerEnabled bool, privateTrieCacheEnabled bool) QuorumChainConfig {
	return QuorumChainConfig{
		multiTenantEnabled:      multiTenantEnabled,
		revertReasonEnabled:     revertReasonEnabled,
		privacyMarkerEnabled:    privacyMarkerEnabled,
		privateTrieCacheEnabled: privateTrieCacheEnabled,
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

// PrivacyMarkerEnabled returns true is privacy marker is enabled
func (c QuorumChainConfig) PrivacyMarkerEnabled() bool {
	return c.privacyMarkerEnabled
}

// PrivateTrieCacheEnabled returns true if the private cache is enabled
func (c QuorumChainConfig) PrivateTrieCacheEnabled() bool {
	return c.privateTrieCacheEnabled
}
