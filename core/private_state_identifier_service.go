package core

import (
	"context"
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

var PSIS PrivateStateIdentifierService
