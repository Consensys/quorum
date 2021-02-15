package core

import (
	"context"
	"github.com/ethereum/go-ethereum/private/engine"
)

type PrivateStateType uint64

const (
	Resident PrivateStateType = iota                          // 0
	Legacy   PrivateStateType = 1 << PrivateStateType(iota-1) // 1
	Pantheon PrivateStateType = 1 << PrivateStateType(iota-1) // 2
)

type PrivateStateMetadata struct {
	ID          string
	Name        string
	Description string
	Type        PrivateStateType
	Addresses   []string
}

func (self *PrivateStateMetadata) HasAddress(address string) bool {
	for _, addr := range self.Addresses {
		if addr == address {
			return true
		}
	}
	return false
}

func (self *PrivateStateMetadata) HasAnyAddress(addresses []string) bool {
	for _, addr := range addresses {
		if self.HasAddress(addr) {
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
	psi, ok := ctx.Value("PSI").(string)
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
