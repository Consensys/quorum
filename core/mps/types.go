package mps

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

var (
	// DefaultPrivateStateMetadata is the metadata for the single private state being used
	// when MPS is disabled
	DefaultPrivateStateMetadata = NewPrivateStateMetadata(
		types.DefaultPrivateStateIdentifier,
		"private",
		"legacy private state",
		Resident,
		nil,
	)
	// EmptyPrivateStateMetadata is the metadata for the empty private state being used
	// when MPS is enabled
	EmptyPrivateStateMetadata = NewPrivateStateMetadata(
		types.EmptyPrivateStateIdentifier,
		"empty",
		"empty state metadata",
		Resident,
		nil,
	)
)

type PrivateStateType uint64

const (
	Resident PrivateStateType = iota                          // 0
	Legacy   PrivateStateType = 1 << PrivateStateType(iota-1) // 1
	Pantheon PrivateStateType = 1 << PrivateStateType(iota-1) // 2
)

// PrivateStateMetadata is the domain model in Quorum which maps with
// PrivacyGroup domain in Tessera
type PrivateStateMetadata struct {
	ID          types.PrivateStateIdentifier
	Name        string
	Description string
	Type        PrivateStateType
	// Addresses stores the public keys in the SAME ORDER being configured
	// in Tessera
	Addresses []string
	// addressIndex is to facilitate fast searching
	addressIndex map[string]struct{}
}

func (psm *PrivateStateMetadata) NotIncludeAny(addresses ...string) bool {
	for _, addr := range addresses {
		if _, found := psm.addressIndex[addr]; found {
			return false
		}
	}
	return true
}

func (psm *PrivateStateMetadata) FilterAddresses(addresses ...string) []string {
	result := make([]string, 0)
	for _, addr := range addresses {
		if _, found := psm.addressIndex[addr]; found {
			result = append(result, addr)
		}
	}
	return result
}

func (psm *PrivateStateMetadata) String() string {
	return fmt.Sprintf("ID=%s,Name=%s,Desc=%s,Type=%d,Addresses=%v", psm.ID, psm.Name, psm.Description, psm.Type, psm.Addresses)
}

func NewPrivateStateMetadata(id types.PrivateStateIdentifier, name, desc string, t PrivateStateType, addresses []string) *PrivateStateMetadata {
	index := make(map[string]struct{}, len(addresses))
	for _, a := range addresses {
		index[a] = struct{}{}
	}
	return &PrivateStateMetadata{
		ID:           id,
		Name:         name,
		Description:  desc,
		Type:         t,
		Addresses:    addresses[:],
		addressIndex: index,
	}
}
