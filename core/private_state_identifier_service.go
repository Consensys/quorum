package core

import (
	"context"

	"github.com/ethereum/go-ethereum/private/engine"
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
	ID          string
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
	ID:          "empty",
	Name:        "empty",
	Description: "empty state",
	Type:        Resident,
	Addresses:   nil,
}

var DefaultPrivateStateMetadata = PrivateStateMetadata{
	ID:          "private",
	Name:        "private",
	Description: "legacy private state",
	Type:        Resident,
	Addresses:   nil,
}

type PrivateStateIdentifierService interface {
	ResolveForManagedParty(managedParty string) (*PrivateStateMetadata, error)
	ResolveForUserContext(ctx context.Context) (*PrivateStateMetadata, error)
	Groups() []engine.PrivacyGroup
}

type PrivatePSISImpl struct {
}

func (t *PrivatePSISImpl) ResolveForManagedParty(managedParty string) (*PrivateStateMetadata, error) {
	return &PrivateStateMetadata{ID: "private", Type: Resident}, nil
}

func (t *PrivatePSISImpl) ResolveForUserContext(ctx context.Context) (*PrivateStateMetadata, error) {
	psi, ok := ctx.Value(rpc.CtxPrivateStateIdentifier).(string)
	if !ok {
		psi = "private"
	}
	return &PrivateStateMetadata{ID: psi, Type: Resident}, nil
}

func (t *PrivatePSISImpl) Groups() []engine.PrivacyGroup {
	return []engine.PrivacyGroup{
		{
			Type:           "Resident",
			Name:           "private",
			PrivacyGroupId: "private",
			Description:    "private",
			From:           "",
			Members:        []string{},
		},
	}
}
