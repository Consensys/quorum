package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
)

// manages a number of state DB objects identified by their PSI (private state identifier)
type MTStateService struct {
	// the trie of private states
	// key - the private state identifier
	// value - the root hash of the private state
	trie state.Trie

	bc *BlockChain // Canonical block chain

	// not sure if relevant - maybe remove
	previousBlockHash common.Hash

	// managed states map
	managedStates map[string]*MTManagedState
}

// A managed state is a pair of stateDb and it's corresponding stateCache objects
// Although right now we may not need a separate stateCache it may be useful if we'll do multiple managed state commits in parallel
type MTManagedState struct {
	stateDb    *state.StateDB
	stateCache state.Database
}

func NewMTStateService(bc *BlockChain, previousBlockHash common.Hash) (*MTStateService, error) {
	mtPrivateStateTrieRoot := rawdb.GetMTPrivateStateRoot(bc.db, previousBlockHash)
	tr, err := bc.mtServiceCache.OpenTrie(mtPrivateStateTrieRoot)
	if err != nil {
		return nil, err
	}
	return &MTStateService{
		trie:          tr,
		bc:            bc,
		managedStates: make(map[string]*MTManagedState)}, nil
}

//utility function for debugging
func (mt *MTStateService) GetManagedStateRoots() []string {
	myMap := mt.managedStates
	keys := make([]string, 0, len(myMap))

	for k, _ := range myMap {
		keys = append(keys, k)
	}
	return keys
}

//utility function for debugging
func (mt *MTStateService) GetManagedStates() []*MTManagedState {
	myMap := mt.managedStates
	keys := make([]*MTManagedState, 0, len(myMap))

	for _, v := range myMap {
		keys = append(keys, v)
	}
	return keys
}

func (mt *MTStateService) GetOverallPrivateState() (*state.StateDB, error) {
	return mt.GetPrivateState("private")
}

func (mt *MTStateService) GetPrivateState(psi string) (*state.StateDB, error) {
	managedState, found := mt.managedStates[psi]
	if found {
		return managedState.stateDb, nil
	}
	mtPrivateStateRoot, err := mt.trie.TryGet([]byte(psi))
	if err != nil {
		return nil, err
	}
	stateCache := state.NewDatabase(mt.bc.db)
	statedb, err := state.New(common.BytesToHash(mtPrivateStateRoot), stateCache)
	if err != nil {
		return nil, err
	}
	mt.managedStates[psi] = &MTManagedState{
		stateCache: stateCache,
		stateDb:    statedb,
	}
	return statedb, nil
}

func (mt *MTStateService) Reset() error {
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
func (mt *MTStateService) CommitAndWrite(block *types.Block) error {
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
	privateTriedb := mt.bc.mtServiceCache.TrieDB()
	err = privateTriedb.Commit(mtRoot, false)
	return err
}

// commit - commits all private states, updates the trie of private states only
func (mt *MTStateService) Commit(block *types.Block) error {
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
