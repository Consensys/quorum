package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
)

// manages a number of state DB objects identified by their PSI (private state identifier)
type PrivateStateService struct {
	// the trie of private states
	// key - the private state identifier
	// value - the root hash of the private state
	trie state.Trie

	bc *BlockChain // Canonical block chain

	// not sure if relevant - maybe remove
	previousBlockHash common.Hash

	// managed states map
	managedStates map[string]*ManagedState
}

// A managed state is a pair of stateDb and it's corresponding stateCache objects
// Although right now we may not need a separate stateCache it may be useful if we'll do multiple managed state commits in parallel
type ManagedState struct {
	stateDb    *state.StateDB
	stateCache state.Database
}

func NewPrivateStateService(bc *BlockChain, previousBlockHash common.Hash) (*PrivateStateService, error) {
	mtPrivateStateTrieRoot := rawdb.GetMTPrivateStateRoot(bc.db, previousBlockHash)
	tr, err := bc.psServiceCache.OpenTrie(mtPrivateStateTrieRoot)
	if err != nil {
		return nil, err
	}
	return &PrivateStateService{
		trie:          tr,
		bc:            bc,
		managedStates: make(map[string]*ManagedState)}, nil
}

//utility function for debugging
func (mt *PrivateStateService) GetManagedStateRoots() []string {
	myMap := mt.managedStates
	keys := make([]string, 0, len(myMap))

	for k := range myMap {
		keys = append(keys, k)
	}
	return keys
}

//utility function for debugging
func (mt *PrivateStateService) GetManagedStates() []*ManagedState {
	myMap := mt.managedStates
	keys := make([]*ManagedState, 0, len(myMap))

	for _, v := range myMap {
		keys = append(keys, v)
	}
	return keys
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
func (mt *PrivateStateService) GetEmptyState() (*state.StateDB, error) {
	return mt.GetPrivateState(EmptyPrivateStateMetadata.ID)
}

func (mt *PrivateStateService) GetPrivateState(psi string) (*state.StateDB, error) {
	managedState, found := mt.managedStates[psi]
	if found {
		return managedState.stateDb, nil
	}
	mtPrivateStateRoot, err := mt.trie.TryGet([]byte(psi))
	if err != nil {
		return nil, err
	}
	stateCache := state.NewDatabase(mt.bc.db)
	statedb, err := state.New(common.BytesToHash(mtPrivateStateRoot), stateCache, nil)
	if err != nil {
		return nil, err
	}
	if psi != EmptyPrivateStateMetadata.ID {
		emptyState, err := mt.GetEmptyState()
		if err != nil {
			return nil, err
		}
		statedb.SetEmptyState(emptyState)
	}
	mt.managedStates[psi] = &ManagedState{
		stateCache: stateCache,
		stateDb:    statedb,
	}
	return statedb, nil
}

func (mt *PrivateStateService) Reset() error {
	for psi, managedState := range mt.managedStates {
		root, err := mt.trie.TryGet([]byte(psi))
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
func (mt *PrivateStateService) CommitAndWrite(block *types.Block) error {
	for psi, managedState := range mt.managedStates {
		// commit each managed state
		privateRoot, err := managedState.stateDb.Commit(mt.bc.chainConfig.IsEIP158(block.Number()))
		if err != nil {
			return err
		}
		// update the managed state root in the trie of states
		err = mt.trie.TryUpdate([]byte(psi), privateRoot.Bytes())
		if err != nil {
			return err
		}
		err = managedState.stateCache.TrieDB().Commit(privateRoot, false)
		if err != nil {
			return err
		}
	}
	// commit the trie of states
	mtRoot, err := mt.trie.Commit(nil)
	if err != nil {
		return err
	}
	err = rawdb.WriteMTPrivateStateRoot(mt.bc.db, block.Root(), mtRoot)
	if err != nil {
		return err
	}
	privateTriedb := mt.bc.psServiceCache.TrieDB()
	err = privateTriedb.Commit(mtRoot, false)
	return err
}

// commit - commits all private states, updates the trie of private states only
func (mt *PrivateStateService) Commit(block *types.Block) error {
	for psi, managedState := range mt.managedStates {
		// commit each managed state
		privateRoot, err := managedState.stateDb.Commit(mt.bc.chainConfig.IsEIP158(block.Number()))
		if err != nil {
			return err
		}
		// update the managed state root in the trie of states
		err = mt.trie.TryUpdate([]byte(psi), privateRoot.Bytes())
		if err != nil {
			return err
		}
	}
	// commit the trie of states
	_, err := mt.trie.Commit(nil)
	if err != nil {
		return err
	}
	return err
}
