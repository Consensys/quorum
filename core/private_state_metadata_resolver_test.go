package core

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestNewPrivateStateMetadata_whenTypical(t *testing.T) {
	psi := types.ToPrivateStateIdentifier("arbitrary")
	testObject := NewPrivateStateMetadata(psi, "name", "desc", Resident, []string{"a"})

	assert.Equal(t, psi, testObject.ID)
	assert.Equal(t, "name", testObject.Name)
	assert.Equal(t, "desc", testObject.Description)
	assert.Equal(t, Resident, testObject.Type)
	assert.Contains(t, testObject.Addresses, "a")
	assert.Contains(t, testObject.addressIndex, "a")
}

func TestNewPrivateStateMetadata_whenNoIndex(t *testing.T) {
	psi := types.ToPrivateStateIdentifier("arbitrary")
	testObject := NewPrivateStateMetadata(psi, "name", "desc", Resident, nil)

	assert.Equal(t, psi, testObject.ID)
	assert.Equal(t, "name", testObject.Name)
	assert.Equal(t, "desc", testObject.Description)
	assert.Equal(t, Resident, testObject.Type)
	assert.Empty(t, testObject.Addresses)
	assert.Empty(t, testObject.addressIndex)
}

func TestPrivateStateMetadata_NotIncludeAny_whenTypical(t *testing.T) {
	pk1 := "arbitrary pk1"
	pk2 := "arbitrary pk2"
	testObject := NewPrivateStateMetadata(types.ToPrivateStateIdentifier("arbitrary"), "name", "desc", Resident, []string{pk1, pk2})

	assert.True(t, testObject.NotIncludeAny("arbitrary pk"))
}

func TestPrivateStateMetadata_NotIncludeAny_whenMatch(t *testing.T) {
	pk1 := "arbitrary pk1"
	pk2 := "arbitrary pk2"
	testObject := NewPrivateStateMetadata(types.ToPrivateStateIdentifier("arbitrary"), "name", "desc", Resident, []string{pk1, pk2})

	assert.False(t, testObject.NotIncludeAny(pk1))
}
