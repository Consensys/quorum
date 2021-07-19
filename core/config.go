package core

type BlockchainQuorumConfig struct {
	SaveRevertReason bool
}

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
