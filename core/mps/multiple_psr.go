package mps

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
)

// MultiplePrivateStateRepository manages a number of state DB objects
// identified by their types.PrivateStateIdentifier. It also maintains a trie
// of private states whose root hash is mapped with a block hash.
type MultiplePrivateStateRepository struct {
	db ethdb.Database
	// trie of private states cache
	repoCache state.Database

	// the trie of private states
	// key - the private state identifier
	// value - the root hash of the private state
	trie state.Trie

	// mux protects concurrent access to managedStates map
	mux sync.Mutex
	// managed states map
	managedStates map[types.PrivateStateIdentifier]*managedState
}

func NewMultiplePrivateStateRepository(db ethdb.Database, cache state.Database, previousBlockHash common.Hash) (*MultiplePrivateStateRepository, error) {
	privateStatesTrieRoot := rawdb.GetPrivateStatesTrieRoot(db, previousBlockHash)
	tr, err := cache.OpenTrie(privateStatesTrieRoot)
	if err != nil {
		return nil, err
	}
	return &MultiplePrivateStateRepository{
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

func (mpsr *MultiplePrivateStateRepository) DefaultState() (*state.StateDB, error) {
	return mpsr.StatePSI(types.EmptyPrivateStateMetadata.ID)
}

func (mpsr *MultiplePrivateStateRepository) DefaultStateMetadata() *types.PrivateStateMetadata {
	return types.EmptyPrivateStateMetadata
}

func (mpsr *MultiplePrivateStateRepository) IsMPS() bool {
	return true
}

func (mpsr *MultiplePrivateStateRepository) StatePSI(psi types.PrivateStateIdentifier) (*state.StateDB, error) {
	mpsr.mux.Lock()
	ms, found := mpsr.managedStates[psi]
	mpsr.mux.Unlock()
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
		emptyState, err := mpsr.DefaultState()
		if err != nil {
			return nil, err
		}
		statedb.SetEmptyState(emptyState)
	}
	mpsr.mux.Lock()
	defer mpsr.mux.Unlock()
	mpsr.managedStates[psi] = &managedState{
		stateCache: stateCache,
		stateDb:    statedb,
	}
	return statedb, nil
}

func (mpsr *MultiplePrivateStateRepository) Reset() error {
	mpsr.mux.Lock()
	defer mpsr.mux.Unlock()
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
func (mpsr *MultiplePrivateStateRepository) CommitAndWrite(isEIP158 bool, block *types.Block) error {
	mpsr.mux.Lock()
	defer mpsr.mux.Unlock()
	for psi, managedState := range mpsr.managedStates {
		// commit each managed state
		privateRoot, err := managedState.stateDb.Commit(isEIP158)
		if err != nil {
			return err
		}
		// update the managed state root in the trie of states
		err = mpsr.trie.TryUpdate([]byte(psi), privateRoot.Bytes())
		if err != nil {
			return err
		}
		err = managedState.stateCache.TrieDB().Commit(privateRoot, false, nil)
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
	err = privateTriedb.Commit(mtRoot, false, nil)
	return err
}

// commit - commits all private states, updates the trie of private states only
func (mpsr *MultiplePrivateStateRepository) Commit(isEIP158 bool, block *types.Block) error {
	mpsr.mux.Lock()
	defer mpsr.mux.Unlock()
	for psi, managedState := range mpsr.managedStates {
		// commit each managed state
		privateRoot, err := managedState.stateDb.Commit(isEIP158)
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
	mpsr.mux.Lock()
	defer mpsr.mux.Unlock()
	managedStatesCopy := make(map[types.PrivateStateIdentifier]*managedState)
	for key, value := range mpsr.managedStates {
		managedStatesCopy[key] = value.Copy()
	}
	return &MultiplePrivateStateRepository{
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
