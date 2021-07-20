package core

type QuorumChainConfig struct {
	RevertReasonEnabled bool // if we should save the revert reasons in the Tx Receipts
	MultiTenantEnabled  bool // if this blockchain supports multitenancy
}
