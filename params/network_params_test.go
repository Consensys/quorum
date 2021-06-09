package params

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuorumImmutabilityThresholdParams(t *testing.T) {
	immutabilityThreshold := GetImmutabilityThreshold()
	assert.Equal(t, 90000, immutabilityThreshold)

	// call Set to set the values
	SetQuorumImmutabilityThreshold(20000)
	immutabilityThreshold = GetImmutabilityThreshold()
	assert.Equal(t, 20000, immutabilityThreshold)
}
