package mps

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
)

// manages a number of state DB objects identified by their PSI (private state identifier)
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
	managedStates map[types.PrivateStateIdentifier]*ManagedState
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
		managedStates: make(map[types.PrivateStateIdentifier]*ManagedState)}, nil
}

// A managed state is a pair of stateDb and it's corresponding stateCache objects
// Although right now we may not need a separate stateCache it may be useful if we'll do multiple managed state commits in parallel
type ManagedState struct {
	stateDb    *state.StateDB
	stateCache state.Database
}

// TODO - !!!IMPORTANT!!! review the state delegate logic with the rest of the team
// The empty state is the private state where all private transactions are executed "as if" the private state is not a party.
// ALL private transactions are being applied to this state (with empty private transaction payload). It allows the public
// state to progress as usual (increasing the public state nonce when transactions are sent/contracts are created).
// The empty state also allows us not to execute the non party transactions for each of the managed private states (by
// injecting the empty state into every managed state as a delegate for CERTAIN calls where the state object is nil). When
// invoking an Exist/Empty/GetNonce call on the StateDB if the state object is missing in the current managed private
// the call is delegated to the empty state.
// State delegate consequences:
// * a managed state in a multiple private state environment will NOT have the same state root as the private state in
// a standalone quorum (due to the empty contracts not being part of the actual state in the multiple private states env)
func (psm *MultiplePrivateStateRepository) GetDefaultState() (*state.StateDB, error) {
	return psm.GetPrivateState(types.EmptyPrivateStateMetadata.ID)
}

func (psm *MultiplePrivateStateRepository) GetDefaultStateMetadata() *types.PrivateStateMetadata {
	return types.EmptyPrivateStateMetadata
}

func (psm *MultiplePrivateStateRepository) IsMPS() bool {
	return true
}

func (psm *MultiplePrivateStateRepository) GetPrivateState(psi types.PrivateStateIdentifier) (*state.StateDB, error) {
	managedState, found := psm.managedStates[psi]
	if found {
		return managedState.stateDb, nil
	}
	privateStatesTrieRoot, err := psm.trie.TryGet([]byte(psi))
	if err != nil {
		return nil, err
	}
	stateCache := state.NewDatabase(psm.db)
	statedb, err := state.New(common.BytesToHash(privateStatesTrieRoot), stateCache, nil)
	if err != nil {
		return nil, err
	}
	if psi != types.EmptyPrivateStateMetadata.ID {
		emptyState, err := psm.GetDefaultState()
		if err != nil {
			return nil, err
		}
		statedb.SetEmptyState(emptyState)
	}
	psm.managedStates[psi] = &ManagedState{
		stateCache: stateCache,
		stateDb:    statedb,
	}
	return statedb, nil
}

func (psm *MultiplePrivateStateRepository) Reset() error {
	for psi, managedState := range psm.managedStates {
		root, err := psm.trie.TryGet([]byte(psi))
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
func (psm *MultiplePrivateStateRepository) CommitAndWrite(block *types.Block) error {
	for psi, managedState := range psm.managedStates {
		// commit each managed state
		privateRoot, err := managedState.stateDb.Commit(psm.chainConfig.IsEIP158(block.Number()))
		if err != nil {
			return err
		}
		// update the managed state root in the trie of states
		err = psm.trie.TryUpdate([]byte(psi), privateRoot.Bytes())
		if err != nil {
			return err
		}
		err = managedState.stateCache.TrieDB().Commit(privateRoot, false)
		if err != nil {
			return err
		}
	}
	// commit the trie of states
	mtRoot, err := psm.trie.Commit(nil)
	if err != nil {
		return err
	}
	err = rawdb.WritePrivateStatesTrieRoot(psm.db, block.Root(), mtRoot)
	if err != nil {
		return err
	}
	privateTriedb := psm.repoCache.TrieDB()
	err = privateTriedb.Commit(mtRoot, false)
	return err
}

// commit - commits all private states, updates the trie of private states only
func (psm *MultiplePrivateStateRepository) Commit(block *types.Block) error {
	for psi, managedState := range psm.managedStates {
		// commit each managed state
		privateRoot, err := managedState.stateDb.Commit(psm.chainConfig.IsEIP158(block.Number()))
		if err != nil {
			return err
		}
		// update the managed state root in the trie of states
		err = psm.trie.TryUpdate([]byte(psi), privateRoot.Bytes())
		if err != nil {
			return err
		}
	}
	// commit the trie of states
	_, err := psm.trie.Commit(nil)
	if err != nil {
		return err
	}
	return err
}

func (psm *MultiplePrivateStateRepository) MergeReceipts(pub, priv types.Receipts) types.Receipts {
	m := make(map[common.Hash]*types.Receipt)
	for _, receipt := range pub {
		m[receipt.TxHash] = receipt
	}
	for _, receipt := range priv {
		publicReceipt := m[receipt.TxHash]
		publicReceipt.PSIToReceipt = make(map[types.PrivateStateIdentifier]*types.Receipt)
		publicReceipt.PSIToReceipt[types.EmptyPrivateStateMetadata.ID] = receipt
		for psi, mtReceipt := range receipt.PSIToReceipt {
			publicReceipt.PSIToReceipt[psi] = mtReceipt
		}
	}

	ret := make(types.Receipts, len(pub))
	for idx, pubReceipt := range pub {
		ret[idx] = m[pubReceipt.TxHash]
	}

	return ret
}
