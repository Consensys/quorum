package core

import "time"

type BlockchainQuorumConfig struct {
	EVMCallTimeOut     time.Duration
	EnableMultitenancy bool
	SaveRevertReason   bool
}

/* TODO uncomment when this will be used
func multitenancyEnabled(c []QuorumConfig) bool {
	return firstBoolean(c, true, func(c QuorumConfig) bool { return c.EnableMultitenancy })
}
*/

func revertReasonSaved(c []BlockchainQuorumConfig) bool {
	return firstBoolean(c, true, func(c BlockchainQuorumConfig) bool { return c.SaveRevertReason })
}

func firstBoolean(cfgs []BlockchainQuorumConfig, v bool, getter func(c BlockchainQuorumConfig) bool) bool {
	for _, c := range cfgs {
		if getter(c) == v {
			return v
		}
	}
	return !v
}
