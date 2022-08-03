package mps

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/privatecache"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	// emptyRoot is the known root hash of an empty trie.
	emptyRoot = common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")
)

type StateRootProviderFunc func(isEIP158 bool) (common.Hash, error)

// MultiplePrivateStateRepository manages a number of state DB objects
// identified by their types.PrivateStateIdentifier. It also maintains a trie
// of private states whose root hash is mapped with a block hash.
type MultiplePrivateStateRepository struct {
	db ethdb.Database
	// trie of private states cache
	repoCache            state.Database
	privateCacheProvider privatecache.Provider

	// the trie of private states
	// key - the private state identifier
	// value - the root hash of the private state
	trie state.Trie

	// mux protects concurrent access to managedStates map
	mux sync.Mutex
	// managed states map
	managedStates map[types.PrivateStateIdentifier]*managedState
}

func NewMultiplePrivateStateRepository(db ethdb.Database, cache state.Database, privateStatesTrieRoot common.Hash, privateCacheProvider privatecache.Provider) (*MultiplePrivateStateRepository, error) {
	tr, err := cache.OpenTrie(privateStatesTrieRoot)
	if err != nil {
		return nil, err
	}
	repo := &MultiplePrivateStateRepository{
		db:                   db,
		repoCache:            cache,
		privateCacheProvider: privateCacheProvider,
		trie:                 tr,
		managedStates:        make(map[types.PrivateStateIdentifier]*managedState),
	}
	return repo, nil
}

// A managed state is a pair of stateDb and it's corresponding stateCache objects
// Although right now we may not need a separate stateCache it may be useful if we'll do multiple managed state commits in parallel
type managedState struct {
	stateDb               *state.StateDB
	stateCache            state.Database
	privateCacheProvider  privatecache.Provider
	stateRootProviderFunc StateRootProviderFunc
}

func (ms *managedState) Copy() *managedState {
	copy := &managedState{
		stateDb:              ms.stateDb.Copy(),
		stateCache:           ms.stateCache,
		privateCacheProvider: ms.privateCacheProvider,
	}
	copy.stateRootProviderFunc = copy.calPrivateStateRoot
	return copy
}

// calPrivateStateRoot is to return state root hash from the commit of a managedState identified by psi
func (ms *managedState) calPrivateStateRoot(isEIP158 bool) (common.Hash, error) {
	privateRoot, err := ms.stateDb.Commit(isEIP158)
	if err != nil {
		return common.Hash{}, err
	}
	err = ms.privateCacheProvider.Commit(ms.stateCache, privateRoot)
	if err != nil {
		return common.Hash{}, err
	}
	return privateRoot, nil
}

func (mpsr *MultiplePrivateStateRepository) DefaultState() (*state.StateDB, error) {
	return mpsr.StatePSI(EmptyPrivateStateMetadata.ID)
}

func (mpsr *MultiplePrivateStateRepository) DefaultStateMetadata() *PrivateStateMetadata {
	return EmptyPrivateStateMetadata
}

func (mpsr *MultiplePrivateStateRepository) IsMPS() bool {
	return true
}

func (mpsr *MultiplePrivateStateRepository) PrivateStateRoot(psi types.PrivateStateIdentifier) (common.Hash, error) {
	privateStateRoot, err := mpsr.trie.TryGet([]byte(psi))
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(privateStateRoot), nil
}

func (mpsr *MultiplePrivateStateRepository) StatePSI(psi types.PrivateStateIdentifier) (*state.StateDB, error) {
	mpsr.mux.Lock()
	ms, found := mpsr.managedStates[psi]
	mpsr.mux.Unlock()
	if found {
		return ms.stateDb, nil
	}
	privateStateRoot, err := mpsr.trie.TryGet([]byte(psi))
	if err != nil {
		return nil, err
	}
	var stateCache state.Database
	var stateDB *state.StateDB
	if privateStateRoot == nil && psi != EmptyPrivateStateMetadata.ID {
		// this is the first time we are trying to use this private state so branch from the empty state
		emptyState, err := mpsr.DefaultState()
		if err != nil {
			return nil, err
		}
		mpsr.mux.Lock()
		ms := mpsr.managedStates[EmptyPrivateStateMetadata.ID]
		mpsr.mux.Unlock()

		stateDB = emptyState.Copy()
		stateCache = ms.stateCache
	} else {
		stateCache = mpsr.privateCacheProvider.GetCache()
		stateDB, err = state.New(common.BytesToHash(privateStateRoot), stateCache, nil)
		if err != nil {
			return nil, err
		}
	}
	mpsr.mux.Lock()
	defer mpsr.mux.Unlock()
	managedState := &managedState{
		stateCache:           stateCache,
		privateCacheProvider: mpsr.privateCacheProvider,
		stateDb:              stateDB,
	}
	managedState.stateRootProviderFunc = managedState.calPrivateStateRoot
	mpsr.managedStates[psi] = managedState
	return stateDB, nil
}

func (mpsr *MultiplePrivateStateRepository) Reset() error {
	mpsr.mux.Lock()
	defer mpsr.mux.Unlock()
	for psi, managedState := range mpsr.managedStates {
		root, err := mpsr.trie.TryGet([]byte(psi))
		if err != nil {
			return err
		}
		// if this was a newly created private state (branched from the empty state) - remove it from the managedStates map
		if root == nil {
			delete(mpsr.managedStates, psi)
			continue
		}
		err = managedState.stateDb.Reset(common.BytesToHash(root))
		if err != nil {
			return err
		}
	}
	return nil
}

// CommitAndWrite commits all private states, updates the trie of private states, writes to disk
func (mpsr *MultiplePrivateStateRepository) CommitAndWrite(isEIP158 bool, block *types.Block) error {
	mpsr.mux.Lock()
	defer mpsr.mux.Unlock()
	// commit each managed state
	for psi, managedState := range mpsr.managedStates {
		// calculate and commit state root if required
		privateRoot, err := managedState.stateRootProviderFunc(isEIP158)
		if err != nil {
			return err
		}

		// update the managed state root in the trie of state roots
		if err := mpsr.trie.TryUpdate([]byte(psi), privateRoot.Bytes()); err != nil {
			return err
		}
	}
	// commit the trie of states
	mtRoot, err := mpsr.trie.Commit(func(paths [][]byte, hexpath []byte, leaf []byte, parent common.Hash) error {
		privateRoot := common.BytesToHash(leaf)
		if privateRoot != emptyRoot {
			mpsr.privateCacheProvider.Reference(privateRoot, parent)
		}
		return nil
	})
	if err != nil {
		return err
	}
	err = rawdb.WritePrivateStatesTrieRoot(mpsr.db, block.Root(), mtRoot)
	if err != nil {
		return err
	}
	mpsr.privateCacheProvider.Commit(mpsr.repoCache, mtRoot)
	mpsr.privateCacheProvider.Reference(mtRoot, block.Root())
	return nil
}

// Commit commits all private states, updates the trie of private states only
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
	_, err := mpsr.trie.Commit(func(paths [][]byte, hexpath []byte, leaf []byte, parent common.Hash) error {
		privateRoot := common.BytesToHash(leaf)
		if privateRoot != emptyRoot {
			mpsr.privateCacheProvider.Reference(privateRoot, parent)
		}
		return nil
	})
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
		db:                   mpsr.db,
		repoCache:            mpsr.repoCache,
		privateCacheProvider: mpsr.privateCacheProvider,
		trie:                 mpsr.repoCache.CopyTrie(mpsr.trie),
		managedStates:        managedStatesCopy,
	}
}

// Given a slice of public receipts and an overlapping (smaller) slice of
// private receipts, return a new slice where the default for each location is
// the public receipt but we take the private receipt in each place we have
// one.
// Each entry for a private receipt will actually consist of a copy of a dummy auxiliary receipt,
// which holds the real private receipts for each PSI under PSReceipts[].
// Note that we also add a private receipt for the "empty" PSI.
func (mpsr *MultiplePrivateStateRepository) MergeReceipts(pub, priv types.Receipts) types.Receipts {
	m := make(map[common.Hash]*types.Receipt)
	for _, receipt := range pub {
		m[receipt.TxHash] = receipt
	}
	for _, receipt := range priv {
		publicReceipt, found := m[receipt.TxHash]
		if !found {
			// this is a PMT receipt - no merging required as it already has the relevant PSReceipts set
			continue
		}
		publicReceipt.PSReceipts = make(map[types.PrivateStateIdentifier]*types.Receipt)
		publicReceipt.PSReceipts[EmptyPrivateStateMetadata.ID] = receipt
		for psi, psReceipt := range receipt.PSReceipts {
			publicReceipt.PSReceipts[psi] = psReceipt
		}
	}

	ret := make(types.Receipts, len(pub))
	for idx, pubReceipt := range pub {
		ret[idx] = m[pubReceipt.TxHash]
	}

	return ret
}
