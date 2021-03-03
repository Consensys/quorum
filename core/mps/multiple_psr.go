package mps

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
)

// MultiplePrivateStateRepository manages a number of state DB objects
// identified by their types.PrivateStateIdentifier. It also maintains a trie
// of private states whose root hash is mapped with a block hash.
type MultiplePrivateStateRepository struct {
	chainConfig *params.ChainConfig
	db          ethdb.Database
	// trie of private states cache
	repoCache state.Database

	// the trie of private states
	// key - the private state identifier
	// value - the root hash of the private state
	trie state.Trie

	// managed states map
	managedStates map[types.PrivateStateIdentifier]*managedState
}

func NewMultiplePrivateStateRepository(chainConfig *params.ChainConfig, db ethdb.Database, cache state.Database, previousBlockHash common.Hash) (*MultiplePrivateStateRepository, error) {
	privateStatesTrieRoot := rawdb.GetPrivateStatesTrieRoot(db, previousBlockHash)
	tr, err := cache.OpenTrie(privateStatesTrieRoot)
	if err != nil {
		return nil, err
	}
	return &MultiplePrivateStateRepository{
		chainConfig:   chainConfig,
		db:            db,
		repoCache:     cache,
		trie:          tr,
		managedStates: make(map[types.PrivateStateIdentifier]*managedState)}, nil
}

// A managed state is a pair of stateDb and it's corresponding stateCache objects
// Although right now we may not need a separate stateCache it may be useful if we'll do multiple managed state commits in parallel
type managedState struct {
	stateDb    *state.StateDB
	stateCache state.Database
}

func (ms *managedState) Copy() *managedState {
	return &managedState{
		stateDb:    ms.stateDb.Copy(),
		stateCache: ms.stateCache,
	}
}

func (mpsr *MultiplePrivateStateRepository) GetDefaultState() (*state.StateDB, error) {
	return mpsr.GetPrivateState(types.EmptyPrivateStateMetadata.ID)
}

func (mpsr *MultiplePrivateStateRepository) GetDefaultStateMetadata() *types.PrivateStateMetadata {
	return types.EmptyPrivateStateMetadata
}

func (mpsr *MultiplePrivateStateRepository) IsMPS() bool {
	return true
}

func (mpsr *MultiplePrivateStateRepository) GetPrivateState(psi types.PrivateStateIdentifier) (*state.StateDB, error) {
	ms, found := mpsr.managedStates[psi]
	if found {
		return ms.stateDb, nil
	}
	privateStatesTrieRoot, err := mpsr.trie.TryGet([]byte(psi))
	if err != nil {
		return nil, err
	}
	stateCache := state.NewDatabase(mpsr.db)
	statedb, err := state.New(common.BytesToHash(privateStatesTrieRoot), stateCache, nil)
	if err != nil {
		return nil, err
	}
	if psi != types.EmptyPrivateStateMetadata.ID {
		emptyState, err := mpsr.GetDefaultState()
		if err != nil {
			return nil, err
		}
		statedb.SetEmptyState(emptyState)
	}
	mpsr.managedStates[psi] = &managedState{
		stateCache: stateCache,
		stateDb:    statedb,
	}
	return statedb, nil
}

func (mpsr *MultiplePrivateStateRepository) Reset() error {
	for psi, managedState := range mpsr.managedStates {
		root, err := mpsr.trie.TryGet([]byte(psi))
		if err != nil {
			return err
		}
		err = managedState.stateDb.Reset(common.BytesToHash(root))
		if err != nil {
			return err
		}
	}
	return nil
}

// commitAndWrite- commits all private states, updates the trie of private states, writes to disk
func (mpsr *MultiplePrivateStateRepository) CommitAndWrite(block *types.Block) error {
	for psi, managedState := range mpsr.managedStates {
		// commit each managed state
		privateRoot, err := managedState.stateDb.Commit(mpsr.chainConfig.IsEIP158(block.Number()))
		if err != nil {
			return err
		}
		// update the managed state root in the trie of states
		err = mpsr.trie.TryUpdate([]byte(psi), privateRoot.Bytes())
		if err != nil {
			return err
		}
		err = managedState.stateCache.TrieDB().Commit(privateRoot, false)
		if err != nil {
			return err
		}
	}
	// commit the trie of states
	mtRoot, err := mpsr.trie.Commit(nil)
	if err != nil {
		return err
	}
	err = rawdb.WritePrivateStatesTrieRoot(mpsr.db, block.Root(), mtRoot)
	if err != nil {
		return err
	}
	privateTriedb := mpsr.repoCache.TrieDB()
	err = privateTriedb.Commit(mtRoot, false)
	return err
}

// commit - commits all private states, updates the trie of private states only
func (mpsr *MultiplePrivateStateRepository) Commit(block *types.Block) error {
	for psi, managedState := range mpsr.managedStates {
		// commit each managed state
		privateRoot, err := managedState.stateDb.Commit(mpsr.chainConfig.IsEIP158(block.Number()))
		if err != nil {
			return err
		}
		// update the managed state root in the trie of states
		err = mpsr.trie.TryUpdate([]byte(psi), privateRoot.Bytes())
		if err != nil {
			return err
		}
	}
	// commit the trie of states
	_, err := mpsr.trie.Commit(nil)
	if err != nil {
		return err
	}
	return err
}

func (mpsr *MultiplePrivateStateRepository) Copy() PrivateStateRepository {
	managedStatesCopy := make(map[types.PrivateStateIdentifier]*managedState)
	for key, value := range mpsr.managedStates {
		managedStatesCopy[key] = value.Copy()
	}
	return &MultiplePrivateStateRepository{
		chainConfig:   mpsr.chainConfig,
		db:            mpsr.db,
		repoCache:     mpsr.repoCache,
		trie:          mpsr.repoCache.CopyTrie(mpsr.trie),
		managedStates: managedStatesCopy,
	}
}

func (mpsr *MultiplePrivateStateRepository) MergeReceipts(pub, priv types.Receipts) types.Receipts {
	m := make(map[common.Hash]*types.Receipt)
	for _, receipt := range pub {
		m[receipt.TxHash] = receipt
	}
	for _, receipt := range priv {
		publicReceipt := m[receipt.TxHash]
		publicReceipt.PSReceipts = make(map[types.PrivateStateIdentifier]*types.Receipt)
		publicReceipt.PSReceipts[types.EmptyPrivateStateMetadata.ID] = receipt
		for psi, receipt := range receipt.PSReceipts {
			publicReceipt.PSReceipts[psi] = receipt
		}
	}

	ret := make(types.Receipts, len(pub))
	for idx, pubReceipt := range pub {
		ret[idx] = m[pubReceipt.TxHash]
	}

	return ret
}
