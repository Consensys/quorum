package mps

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// DefaultPrivateStateRepository acts as the single private state in the original
// Quorum design.
type DefaultPrivateStateRepository struct {
	chainConfig *params.ChainConfig
	db          ethdb.Database
	// cache of stateDB
	stateCache state.Database

	// the trie of private states
	// key - the private state identifier
	// value - the root hash of the private state
	stateDB *state.StateDB
	root    common.Hash
}

func NewDefaultPrivateStateRepository(chainConfig *params.ChainConfig, db ethdb.Database, cache state.Database, previousBlockHash common.Hash) (*DefaultPrivateStateRepository, error) {
	root := rawdb.GetPrivateStateRoot(db, previousBlockHash)

	statedb, err := state.New(root, cache, nil)
	if err != nil {
		return nil, err
	}

	return &DefaultPrivateStateRepository{
		chainConfig: chainConfig,
		db:          db,
		stateCache:  cache,
		stateDB:     statedb,
		root:        root,
	}, nil
}

func (dpsr *DefaultPrivateStateRepository) GetDefaultState() (*state.StateDB, error) {
	return dpsr.stateDB, nil
}

func (dpsr *DefaultPrivateStateRepository) GetDefaultStateMetadata() *types.PrivateStateMetadata {
	return types.DefaultPrivateStateMetadata
}

func (dpsr *DefaultPrivateStateRepository) IsMPS() bool {
	return false
}

func (dpsr *DefaultPrivateStateRepository) GetPrivateState(psi types.PrivateStateIdentifier) (*state.StateDB, error) {
	if psi != types.DefaultPrivateStateIdentifier {
		return nil, fmt.Errorf("only the 'private' psi is supported by the default private state manager")
	}
	return dpsr.stateDB, nil
}

func (dpsr *DefaultPrivateStateRepository) Reset() error {
	// TODO - see if we need to  store the original root
	return dpsr.stateDB.Reset(dpsr.root)
}

// commitAndWrite- commits all private states, updates the state trie and writes to disk
func (dpsr *DefaultPrivateStateRepository) CommitAndWrite(block *types.Block) error {
	privateRoot, err := dpsr.stateDB.Commit(dpsr.chainConfig.IsEIP158(block.Number()))
	if err != nil {
		return err
	}

	if err := rawdb.WritePrivateStateRoot(dpsr.db, block.Root(), privateRoot); err != nil {
		log.Error("Failed writing private state root", "err", err)
		return err
	}
	return dpsr.stateCache.TrieDB().Commit(privateRoot, false)
}

// commit - commits all private states, updates the trie of private states only
func (dpsr *DefaultPrivateStateRepository) Commit(block *types.Block) error {
	var err error
	dpsr.root, err = dpsr.stateDB.Commit(dpsr.chainConfig.IsEIP158(block.Number()))
	return err
}

func (dpsr *DefaultPrivateStateRepository) MergeReceipts(pub, priv types.Receipts) types.Receipts {
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
