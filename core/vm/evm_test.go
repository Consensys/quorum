package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAffectedMode_Update_whenTypical(t *testing.T) {
	testObject := ModeUnknown
	authorizedReads := []bool{true, false}
	authorizedWrites := []bool{true, false}
	for _, authorizedRead := range authorizedReads {
		for _, authorizedWrite := range authorizedWrites {
			actual := testObject.Update(authorizedRead, authorizedWrite)

			assert.True(t, actual.Has(ModeUpdated))
			assert.Equal(t, authorizedRead, actual.Has(ModeRead))
			assert.Equal(t, authorizedWrite, actual.Has(ModeWrite))
			assert.False(t, testObject.Has(ModeUpdated))
		}
	}
}
