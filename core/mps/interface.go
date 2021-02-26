//go:generate mockgen -source interface.go -destination=mock_interface.go -package=mps

package mps

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
)

type PrivateStateManager interface {
	PrivateStateMetadataResolver
	GetPrivateStateRepository(blockHash common.Hash) (PrivateStateRepository, error)
}

type PrivateStateMetadataResolver interface {
	ResolveForManagedParty(managedParty string) (*types.PrivateStateMetadata, error)
	ResolveForUserContext(ctx context.Context) (*types.PrivateStateMetadata, error)
	PSIs() []types.PrivateStateIdentifier
	// NotIncludeAny returns true if NONE of the managedParties is a member
	// of the given psm, otherwise returns false
	NotIncludeAny(psm *types.PrivateStateMetadata, managedParties ...string) bool
}

type PrivateStateRepository interface {
	GetPrivateState(psi types.PrivateStateIdentifier) (*state.StateDB, error)
	CommitAndWrite(block *types.Block) error
	Commit(block *types.Block) error
	Reset() error
	GetDefaultState() (*state.StateDB, error)
	GetDefaultStateMetadata() *types.PrivateStateMetadata
	IsMPS() bool
	MergeReceipts(pub, priv types.Receipts) types.Receipts
} // with two implementations: MultiplePrivateStateRepository and DefaultPrivateStateRepository
