package core

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

// manages a number of state DB objects identified by their PSI (private state identifier)
type LegacyPrivateStateManager struct {
	// the trie of private states
	// key - the private state identifier
	// value - the root hash of the private state
	stateDB *state.StateDB
	root    common.Hash

	bc *BlockChain // Canonical block chain
}

func NewLegacyPrivateStateManager(bc *BlockChain, previousBlockHash common.Hash) (PrivateStateManager, error) {
	root := rawdb.GetPrivateStateRoot(bc.db, previousBlockHash)

	statedb, err := state.New(root, bc.psManagerCache, nil)
	if err != nil {
		return nil, err
	}

	return &LegacyPrivateStateManager{
		bc:      bc,
		stateDB: statedb,
		root:    root,
	}, nil
}

func (psm *LegacyPrivateStateManager) GetDefaultState() (*state.StateDB, error) {
	return psm.stateDB, nil
}

func (psm *LegacyPrivateStateManager) GetDefaultStateMetadata() PrivateStateMetadata {
	return DefaultPrivateStateMetadata
}

func (psm *LegacyPrivateStateManager) IsMPS() bool {
	return false
}

func (psm *LegacyPrivateStateManager) GetPrivateState(psi string) (*state.StateDB, error) {
	if psi != "private" {
		return nil, fmt.Errorf("Only the 'private' psi is supported by the legacy private state manager")
	}
	return psm.stateDB, nil
}

func (psm *LegacyPrivateStateManager) Reset() error {
	// TODO - see if we need to  store the original root
	return psm.stateDB.Reset(psm.root)
}

// commitAndWrite- commits all private states, updates the trie of private states, writes to disk
func (psm *LegacyPrivateStateManager) CommitAndWrite(block *types.Block) error {
	privateRoot, err := psm.stateDB.Commit(psm.bc.chainConfig.IsEIP158(block.Number()))
	if err != nil {
		return err
	}

	if err := rawdb.WritePrivateStateRoot(psm.bc.db, block.Root(), privateRoot); err != nil {
		log.Error("Failed writing private state root", "err", err)
		return err
	}
	return psm.bc.psManagerCache.TrieDB().Commit(privateRoot, false)
}

// commit - commits all private states, updates the trie of private states only
func (psm *LegacyPrivateStateManager) Commit(block *types.Block) error {
	var err error
	psm.root, err = psm.stateDB.Commit(psm.bc.chainConfig.IsEIP158(block.Number()))
	return err
}

func (psm *LegacyPrivateStateManager) MergeReceipts(pub, priv types.Receipts) types.Receipts {
	m := make(map[common.Hash]*types.Receipt)
	for _, receipt := range pub {
		m[receipt.TxHash] = receipt
	}
	for _, receipt := range priv {
		m[receipt.TxHash] = receipt
	}

	ret := make(types.Receipts, 0, len(pub))
	for _, pubReceipt := range pub {
		ret = append(ret, m[pubReceipt.TxHash])
	}

	return ret
}
