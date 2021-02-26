package mps

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type DefaultPrivateStateMetadataResolver struct {
}

func (dpsmr *DefaultPrivateStateMetadataResolver) ResolveForManagedParty(_ string) (*types.PrivateStateMetadata, error) {
	return types.DefaultPrivateStateMetadata, nil
}

func (dpsmr *DefaultPrivateStateMetadataResolver) ResolveForUserContext(ctx context.Context) (*types.PrivateStateMetadata, error) {
	psi, ok := ctx.Value(rpc.CtxPrivateStateIdentifier).(types.PrivateStateIdentifier)
	if !ok {
		psi = types.DefaultPrivateStateIdentifier
	}
	return &types.PrivateStateMetadata{ID: psi, Type: types.Resident}, nil
}

func (dpsmr *DefaultPrivateStateMetadataResolver) PSIs() []types.PrivateStateIdentifier {
	return []types.PrivateStateIdentifier{
		types.DefaultPrivateStateIdentifier,
	}
}

func (dpsmr *DefaultPrivateStateMetadataResolver) NotIncludeAny(_ *types.PrivateStateMetadata, _ ...string) bool {
	// with default implementation, all managedParties are members of the psm
	return false
}
