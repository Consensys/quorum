package core

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type DefaultPrivateStateManager struct {
	bc        *BlockChain
	repoCache state.Database
}

func NewDefaultPrivateStateManager(bc *BlockChain) *DefaultPrivateStateManager {
	return &DefaultPrivateStateManager{
		bc:        bc,
		repoCache: state.NewDatabase(bc.db),
	}
}

func (d *DefaultPrivateStateManager) GetPrivateStateRepository(blockHash common.Hash) (mps.PrivateStateRepository, error) {
	return mps.NewDefaultPrivateStateRepository(d.bc.chainConfig, d.bc.db, d.repoCache, blockHash)
}

func (d *DefaultPrivateStateManager) ResolveForManagedParty(_ string) (*types.PrivateStateMetadata, error) {
	return types.DefaultPrivateStateMetadata, nil
}

func (d *DefaultPrivateStateManager) ResolveForUserContext(ctx context.Context) (*types.PrivateStateMetadata, error) {
	psi, ok := ctx.Value(rpc.CtxPrivateStateIdentifier).(types.PrivateStateIdentifier)
	if !ok {
		psi = types.DefaultPrivateStateIdentifier
	}
	return &types.PrivateStateMetadata{ID: psi, Type: types.Resident}, nil
}

func (d *DefaultPrivateStateManager) PSIs() []types.PrivateStateIdentifier {
	return []types.PrivateStateIdentifier{
		types.DefaultPrivateStateIdentifier,
	}
}

func (d *DefaultPrivateStateManager) NotIncludeAny(_ *types.PrivateStateMetadata, _ ...string) bool {
	// with default implementation, all managedParties are members of the psm
	return false
}

func (d *DefaultPrivateStateManager) GetCache() state.Database {
	return d.repoCache
}
