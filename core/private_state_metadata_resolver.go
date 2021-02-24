package core

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
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
	Addresses   []string
}

func (psm *PrivateStateMetadata) HasAddress(address string) bool {
	for _, addr := range psm.Addresses {
		if addr == address {
			return true
		}
	}
	return false
}

func (psm *PrivateStateMetadata) HasAnyAddress(addresses []string) bool {
	for _, addr := range addresses {
		if psm.HasAddress(addr) {
			return true
		}
	}
	return false
}

var EmptyPrivateStateMetadata = PrivateStateMetadata{
	ID:          types.ToPrivateStateIdentifier("empty"),
	Name:        "empty",
	Description: "empty state",
	Type:        Resident,
	Addresses:   nil,
}

var DefaultPrivateStateMetadata = PrivateStateMetadata{
	ID:          types.DefaultPrivateStateIdentifier,
	Name:        "private",
	Description: "legacy private state",
	Type:        Resident,
	Addresses:   nil,
}

type PrivateStateMetadataResolver interface {
	ResolveForManagedParty(managedParty string) (*PrivateStateMetadata, error)
	ResolveForUserContext(ctx context.Context) (*PrivateStateMetadata, error)
	PSIs() []types.PrivateStateIdentifier
}

type DefaultPrivateStateMetadataResolver struct {
}

func (t *DefaultPrivateStateMetadataResolver) ResolveForManagedParty(managedParty string) (*PrivateStateMetadata, error) {
	return &PrivateStateMetadata{ID: types.DefaultPrivateStateIdentifier, Type: Resident}, nil
}

func (t *DefaultPrivateStateMetadataResolver) ResolveForUserContext(ctx context.Context) (*PrivateStateMetadata, error) {
	psi, ok := ctx.Value(rpc.CtxPrivateStateIdentifier).(types.PrivateStateIdentifier)
	if !ok {
		psi = types.DefaultPrivateStateIdentifier
	}
	return &PrivateStateMetadata{ID: psi, Type: Resident}, nil
}

func (t *DefaultPrivateStateMetadataResolver) PSIs() []types.PrivateStateIdentifier {
	return []types.PrivateStateIdentifier{
		types.DefaultPrivateStateIdentifier,
	}
}
