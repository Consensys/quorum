package core

type BlockchainQuorumConfig struct {
	RevertReasonEnabled bool // if we should save the revert reasons in the Tx Receipts
	MultiTenantEnabled  bool // if this blockchain supports multitenancy
}

func blockchainQuorumConfig(cfg ...BlockchainQuorumConfig) BlockchainQuorumConfig {
	return BlockchainQuorumConfig{
		RevertReasonEnabled: revertReasonEnabled(cfg),
		MultiTenantEnabled:  multiTenantEnabled(cfg),
	}
}

func revertReasonEnabled(c []BlockchainQuorumConfig) bool {
	return firstBoolean(c, true, func(c BlockchainQuorumConfig) bool { return c.RevertReasonEnabled })
}

func multiTenantEnabled(c []BlockchainQuorumConfig) bool {
	return firstBoolean(c, true, func(c BlockchainQuorumConfig) bool { return c.MultiTenantEnabled })
}

func firstBoolean(cfgs []BlockchainQuorumConfig, v bool, getter func(c BlockchainQuorumConfig) bool) bool {
	for _, c := range cfgs {
		if getter(c) == v {
			return v
		}
	}
	return !v
}
