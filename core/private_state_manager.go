package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/psmr"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
)

type PrivateStateManager interface {
	GetPrivateState(psi types.PrivateStateIdentifier) (*state.StateDB, error)
	CommitAndWrite(block *types.Block) error
	Commit(block *types.Block) error
	Reset() error
	GetDefaultState() (*state.StateDB, error)
	GetDefaultStateMetadata() psmr.PrivateStateMetadata
	IsMPS() bool
	MergeReceipts(pub, priv types.Receipts) types.Receipts
} // with two implementations: MultiplePrivateStateManager and LegacyPrivateStateManager

func NewPrivateStateManager(bc *BlockChain, previousBlockHash common.Hash) (PrivateStateManager, error) {
	if bc.chainConfig.IsMPS {
		return NewMultiplePrivateStateManager(bc, previousBlockHash)
	} else {
		return NewLegacyPrivateStateManager(bc, previousBlockHash)
	}
}
