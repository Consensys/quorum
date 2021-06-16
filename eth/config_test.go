package eth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuorumDefautConfig(t *testing.T) {
	type data struct {
		actual   uint64
		expected uint64
	}
	var testData = map[string]data{
		"eth.DefaultConfig.Miner.GasFloor": {DefaultConfig.Miner.GasFloor, 700000000},
		"eth.DefaultConfig.Miner.GasCeil":  {DefaultConfig.Miner.GasCeil, 800000000},
	}
	for k, v := range testData {
		assert.Equal(t, v.expected, v.actual, k+" value mismatch")
	}
}
