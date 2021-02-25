package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/psmr"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
)

// manages a number of state DB objects identified by their PSI (private state identifier)
type MultiplePrivateStateManager struct {
	// the trie of private states
	// key - the private state identifier
	// value - the root hash of the private state
	trie state.Trie

	bc *BlockChain // Canonical block chain

	// not sure if relevant - maybe remove
	previousBlockHash common.Hash

	// managed states map
	managedStates map[types.PrivateStateIdentifier]*ManagedState
}

func NewMultiplePrivateStateManager(bc *BlockChain, previousBlockHash common.Hash) (PrivateStateManager, error) {
	privateStatesTrieRoot := rawdb.GetPrivateStatesTrieRoot(bc.db, previousBlockHash)
	tr, err := bc.psManagerCache.OpenTrie(privateStatesTrieRoot)
	if err != nil {
		return nil, err
	}
	return &MultiplePrivateStateManager{
		trie:          tr,
		bc:            bc,
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
func (psm *MultiplePrivateStateManager) GetDefaultState() (*state.StateDB, error) {
	return psm.GetPrivateState(psmr.EmptyPrivateStateMetadata.ID)
}

func (psm *MultiplePrivateStateManager) GetDefaultStateMetadata() psmr.PrivateStateMetadata {
	return psmr.EmptyPrivateStateMetadata
}

func (psm *MultiplePrivateStateManager) IsMPS() bool {
	return true
}

func (psm *MultiplePrivateStateManager) GetPrivateState(psi types.PrivateStateIdentifier) (*state.StateDB, error) {
	managedState, found := psm.managedStates[psi]
	if found {
		return managedState.stateDb, nil
	}
	privateStatesTrieRoot, err := psm.trie.TryGet([]byte(psi))
	if err != nil {
		return nil, err
	}
	stateCache := state.NewDatabase(psm.bc.db)
	statedb, err := state.New(common.BytesToHash(privateStatesTrieRoot), stateCache, nil)
	if err != nil {
		return nil, err
	}
	if psi != psmr.EmptyPrivateStateMetadata.ID {
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

func (psm *MultiplePrivateStateManager) Reset() error {
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
func (psm *MultiplePrivateStateManager) CommitAndWrite(block *types.Block) error {
	for psi, managedState := range psm.managedStates {
		// commit each managed state
		privateRoot, err := managedState.stateDb.Commit(psm.bc.chainConfig.IsEIP158(block.Number()))
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
	err = rawdb.WritePrivateStatesTrieRoot(psm.bc.db, block.Root(), mtRoot)
	if err != nil {
		return err
	}
	privateTriedb := psm.bc.psManagerCache.TrieDB()
	err = privateTriedb.Commit(mtRoot, false)
	return err
}

// commit - commits all private states, updates the trie of private states only
func (psm *MultiplePrivateStateManager) Commit(block *types.Block) error {
	for psi, managedState := range psm.managedStates {
		// commit each managed state
		privateRoot, err := managedState.stateDb.Commit(psm.bc.chainConfig.IsEIP158(block.Number()))
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

func (psm *MultiplePrivateStateManager) MergeReceipts(pub, priv types.Receipts) types.Receipts {
	m := make(map[common.Hash]*types.Receipt)
	for _, receipt := range pub {
		m[receipt.TxHash] = receipt
	}
	for _, receipt := range priv {
		publicReceipt := m[receipt.TxHash]
		publicReceipt.MTVersions = make(map[types.PrivateStateIdentifier]*types.Receipt)
		publicReceipt.MTVersions[psmr.EmptyPrivateStateMetadata.ID] = receipt
		for psi, mtReceipt := range receipt.MTVersions {
			publicReceipt.MTVersions[psi] = mtReceipt
		}
	}

	ret := make(types.Receipts, len(pub))
	for idx, pubReceipt := range pub {
		ret[idx] = m[pubReceipt.TxHash]
	}

	return ret
}
