package core

import "time"

type QuorumConfig struct {
	EVMCallTimeOut     time.Duration
	EnableMultitenancy bool
	SaveRevertReason   bool
}

/* TODO uncomment when this will be used
func multitenancyEnabled(c []QuorumConfig) bool {
	return firstBoolean(c, true, func(c QuorumConfig) bool { return c.EnableMultitenancy })
}
*/

func revertReasonSaved(c []QuorumConfig) bool {
	return firstBoolean(c, true, func(c QuorumConfig) bool { return c.SaveRevertReason })
}

func firstBoolean(cfgs []QuorumConfig, v bool, getter func(c QuorumConfig) bool) bool {
	for _, c := range cfgs {
		if getter(c) == v {
			return v
		}
	}
	return !v
}
