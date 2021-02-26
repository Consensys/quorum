//go:generate mockgen -source interface.go -destination=mock_interface.go -package=mps

package mps

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

type PrivateStateMetadataResolver interface {
	ResolveForManagedParty(managedParty string) (*types.PrivateStateMetadata, error)
	ResolveForUserContext(ctx context.Context) (*types.PrivateStateMetadata, error)
	PSIs() []types.PrivateStateIdentifier
	// NotIncludeAny returns true if NONE of the managedParties is a member
	// of the given psm, otherwise returns false
	NotIncludeAny(psm *types.PrivateStateMetadata, managedParties ...string) bool
}
