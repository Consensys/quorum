package core

type QuorumChainConfig struct {
	RevertReasonEnabled bool // if we should save the revert reasons in the Tx Receipts
	MultiTenantEnabled  bool // if this blockchain supports multitenancy
}

func quorumChainConfig(cfg ...QuorumChainConfig) QuorumChainConfig {
	return QuorumChainConfig{
		RevertReasonEnabled: revertReasonEnabled(cfg),
		MultiTenantEnabled:  multiTenantEnabled(cfg),
	}
}

func revertReasonEnabled(c []QuorumChainConfig) bool {
	return firstBoolean(c, true, func(c QuorumChainConfig) bool { return c.RevertReasonEnabled })
}

func multiTenantEnabled(c []QuorumChainConfig) bool {
	return firstBoolean(c, true, func(c QuorumChainConfig) bool { return c.MultiTenantEnabled })
}

func firstBoolean(cfgs []QuorumChainConfig, v bool, getter func(c QuorumChainConfig) bool) bool {
	for _, c := range cfgs {
		if getter(c) == v {
			return v
		}
	}
	return !v
}
