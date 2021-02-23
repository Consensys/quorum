package types

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodePSI_whenTypical(t *testing.T) {
	actual := EncodePSI(strconv.AppendUint(nil, 32, 10), "ARBITRARY")

	assert.Equal(t, "\"ARBITRARY/32\"", string(actual))
}

func TestEncodePSI_whenNoPSI(t *testing.T) {
	actual := EncodePSI(strconv.AppendUint(nil, 32, 10), "")

	assert.Equal(t, "32", string(actual))
}

func TestDecodePSI_whenTypical(t *testing.T) {
	input := "\"ARBITRARY/1\""

	psi := DecodePSI([]byte(input))

	assert.Equal(t, PrivateStateIdentifier("ARBITRARY"), psi)
}

func TestDecodePSI_whenNoPSI(t *testing.T) {
	inputs := []string{
		"1",
		"\"1",
		"1\"",
		"\"xyz\"",
	}
	for _, input := range inputs {
		psi := DecodePSI([]byte(input))

		assert.Equal(t, DefaultPrivateStateIdentifier, psi, "input: %s", input)
	}
}
