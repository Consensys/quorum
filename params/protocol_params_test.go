package params

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Quorum - test key constant values modified by Quorum
func TestQuorumParams(t *testing.T) {
	type data struct {
		actual   uint64
		expected uint64
	}
	var testData = map[string]data{
		"GasLimitBoundDivisor":       {GasLimitBoundDivisor, 4096},
		"MinGasLimit":                {MinGasLimit, 700000000},
		"GenesisGasLimit":            {GenesisGasLimit, 800000000},
		"QuorumMaximumExtraDataSize": {QuorumMaximumExtraDataSize, 65},
		"QuorumMaxPayloadBufferSize": {QuorumMaxPayloadBufferSize, 128},
	}
	for k, v := range testData {
		assert.Equal(t, v.expected, v.actual, k+" value mismatch")
	}
}
