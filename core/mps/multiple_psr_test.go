package mps

import (
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/privatecache"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

//TestMultiplePSRCopy tests that copying a the PSR object indeed makes the original and
// the copy and their corresponding managed states independent of each other.
func TestMultiplePSRCopy(t *testing.T) {

	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	privateStateCacheProvider := privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ := NewMultiplePrivateStateRepository(testdb, testCache, common.Hash{}, privateStateCacheProvider)

	testState, _ := psr.StatePSI(types.PrivateStateIdentifier("test"))
	privState, _ := psr.StatePSI(types.DefaultPrivateStateIdentifier)

	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState.AddBalance(addr, big.NewInt(int64(i)))
	}
	testState.Finalise(false)

	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		privState.AddBalance(addr, big.NewInt(int64(i)))
	}
	privState.Finalise(false)

	psrCopy := psr.Copy().(*MultiplePrivateStateRepository)

	testStateCopy, _ := psrCopy.StatePSI(types.PrivateStateIdentifier("test"))
	privStateCopy, _ := psrCopy.StatePSI(types.DefaultPrivateStateIdentifier)
	addedState, _ := psrCopy.StatePSI(types.PrivateStateIdentifier("added"))

	// modify all in memory
	for i := byte(0); i < 255; i++ {
		testState.AddBalance(common.BytesToAddress([]byte{i}), big.NewInt(2*int64(i)))
		privState.AddBalance(common.BytesToAddress([]byte{i}), big.NewInt(2*int64(i)))

		testStateCopy.AddBalance(common.BytesToAddress([]byte{i}), big.NewInt(3*int64(i)))
		privStateCopy.AddBalance(common.BytesToAddress([]byte{i}), big.NewInt(3*int64(i)))
		addedState.AddBalance(common.BytesToAddress([]byte{i}), big.NewInt(3*int64(i)))
	}

	// Finalise the changes on all concurrently
	finalise := func(wg *sync.WaitGroup, db *state.StateDB) {
		defer wg.Done()
		db.Finalise(true)
	}

	var wg sync.WaitGroup
	wg.Add(5)
	go finalise(&wg, testState)
	go finalise(&wg, testStateCopy)
	go finalise(&wg, privState)
	go finalise(&wg, privStateCopy)
	go finalise(&wg, addedState)
	wg.Wait()

	//copies contain correct managed states
	assert.Contains(t, psr.managedStates, types.EmptyPrivateStateIdentifier)
	assert.Contains(t, psr.managedStates, types.DefaultPrivateStateIdentifier)
	assert.Contains(t, psr.managedStates, types.PrivateStateIdentifier("test"))
	assert.NotContains(t, psr.managedStates, types.PrivateStateIdentifier("added"))

	assert.Contains(t, psrCopy.managedStates, types.EmptyPrivateStateIdentifier)
	assert.Contains(t, psrCopy.managedStates, types.DefaultPrivateStateIdentifier)
	assert.Contains(t, psrCopy.managedStates, types.PrivateStateIdentifier("test"))
	assert.Contains(t, psrCopy.managedStates, types.PrivateStateIdentifier("added"))

	assert.Equal(t, psr.db, psrCopy.db)
	assert.Equal(t, psr.repoCache, psrCopy.repoCache)
	assert.NotEqual(t, psr.trie, psrCopy.trie)

	// Verify that the all states have been updated independently
	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		testObj := testState.GetOrNewStateObject(addr)
		testCopyObj := testStateCopy.GetOrNewStateObject(addr)
		privObj := privState.GetOrNewStateObject(addr)
		privCopyObj := privStateCopy.GetOrNewStateObject(addr)
		addedObj := addedState.GetOrNewStateObject(addr)

		if want := big.NewInt(3 * int64(i)); testObj.Balance().Cmp(want) != 0 {
			t.Errorf("empty obj %d: balance mismatch: have %v, want %v", i, testObj.Balance(), want)
		}
		if want := big.NewInt(3 * int64(i)); privObj.Balance().Cmp(want) != 0 {
			t.Errorf("priv obj %d: balance mismatch: have %v, want %v", i, privObj.Balance(), want)
		}
		if want := big.NewInt(4 * int64(i)); testCopyObj.Balance().Cmp(want) != 0 {
			t.Errorf("empty copy obj %d: balance mismatch: have %v, want %v", i, testCopyObj.Balance(), want)
		}
		if want := big.NewInt(4 * int64(i)); privCopyObj.Balance().Cmp(want) != 0 {
			t.Errorf("priv copy obj %d: balance mismatch: have %v, want %v", i, privCopyObj.Balance(), want)
		}
		if want := big.NewInt(3 * int64(i)); addedObj.Balance().Cmp(want) != 0 {
			t.Errorf("added obj %d: balance mismatch: have %v, want %v", i, addedObj.Balance(), want)
		}
	}
}

//TestMultiplePSRReset tests that state objects are cleared from all managedState statedbs after reset call
//Any updated stateObjects not committed to statedbs before reset will be cleared
func TestMultiplePSRReset(t *testing.T) {

	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	privateStateCacheProvider := privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ := NewMultiplePrivateStateRepository(testdb, testCache, common.Hash{}, privateStateCacheProvider)

	testState, _ := psr.StatePSI(types.PrivateStateIdentifier("test"))
	emptyState, _ := psr.StatePSI(types.EmptyPrivateStateIdentifier)

	addr := common.BytesToAddress([]byte{254})
	testState.AddBalance(addr, big.NewInt(int64(254)))
	emptyState.AddBalance(addr, big.NewInt(int64(254)))

	// have something to revert to (rather than the empty trie of private states)
	psr.CommitAndWrite(false, types.NewBlockWithHeader(&types.Header{Root: common.Hash{}}))

	// testState2 should branch from the emptyState - so it should contain the contract with address 254...
	testState2, _ := psr.StatePSI(types.PrivateStateIdentifier("test2"))

	for i := byte(0); i < 254; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState.AddBalance(addr, big.NewInt(int64(i)))
		testState2.AddBalance(addr, big.NewInt(int64(i)))
		emptyState.AddBalance(addr, big.NewInt(int64(i)))
	}
	testState.Finalise(false)
	testState2.Finalise(false)
	emptyState.Finalise(false)

	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		assert.True(t, testState.Exist(addr))
		assert.True(t, testState2.Exist(addr))
		assert.True(t, emptyState.Exist(addr))
	}

	psr.Reset()

	// test2 is no longer there in the managed states after reset
	assert.Contains(t, psr.managedStates, types.PrivateStateIdentifier("test"))
	assert.NotContains(t, psr.managedStates, types.PrivateStateIdentifier("test2"))
	assert.Contains(t, psr.managedStates, types.EmptyPrivateStateIdentifier)

	for i := byte(0); i < 254; i++ {
		addr := common.BytesToAddress([]byte{i})
		assert.False(t, testState.Exist(addr))
		assert.False(t, emptyState.Exist(addr))
	}
	addr = common.BytesToAddress([]byte{254})
	assert.True(t, testState.Exist(addr))
	assert.True(t, emptyState.Exist(addr))
}

//TestCreatingManagedStates tests that managed states are created and added to managedState map
func TestCreatingManagedStates(t *testing.T) {
	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	privateStateCacheProvider := privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ := NewMultiplePrivateStateRepository(testdb, testCache, common.Hash{}, privateStateCacheProvider)

	//create some managed states
	psr.DefaultState()
	psr.StatePSI(types.PrivateStateIdentifier("test"))
	psr.StatePSI(types.DefaultPrivateStateIdentifier)

	//check if they exist in managedStates map
	assert.Contains(t, psr.managedStates, types.EmptyPrivateStateIdentifier)
	assert.Contains(t, psr.managedStates, types.DefaultPrivateStateIdentifier)
	assert.Contains(t, psr.managedStates, types.PrivateStateIdentifier("test"))
	assert.NotContains(t, psr.managedStates, types.PrivateStateIdentifier("added"))
}

//TestMultiplePSRCommit tests that managedStates are updated, trie of states is updated but not written to db
func TestMultiplePSRCommit(t *testing.T) {
	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	privateStateCacheProvider := privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ := NewMultiplePrivateStateRepository(testdb, testCache, common.Hash{}, privateStateCacheProvider)
	header := &types.Header{Number: big.NewInt(int64(1)), Root: common.Hash{123}}
	block := types.NewBlockWithHeader(header)

	testState, _ := psr.StatePSI(types.PrivateStateIdentifier("test"))
	privState, _ := psr.StatePSI(types.DefaultPrivateStateIdentifier)

	//states have empty tries first
	testRoot := testState.IntermediateRoot(false)
	privRoot := privState.IntermediateRoot(false)
	assert.Equal(t, testRoot, emptyRoot)
	assert.Equal(t, privRoot, emptyRoot)

	//make updates to states
	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState.AddBalance(addr, big.NewInt(int64(i)))
		privState.AddBalance(addr, big.NewInt(int64(i)))
	}

	assert.Equal(t, rawdb.GetPrivateStatesTrieRoot(testdb, block.Root()), common.Hash{})

	psr.Commit(false, block)

	//trie root updated but not committed to db
	assert.NotEqual(t, psr.trie.Hash(), common.Hash{})
	assert.Equal(t, rawdb.GetPrivateStatesTrieRoot(testdb, block.Root()), common.Hash{})

	privateKey, _ := psr.trie.TryGet([]byte(types.DefaultPrivateStateIdentifier))
	assert.NotEqual(t, len(privateKey), 0)
	testKey, _ := psr.trie.TryGet([]byte(types.PrivateStateIdentifier("test")))
	assert.NotEqual(t, len(testKey), 0)
	emptyKey, _ := psr.trie.TryGet([]byte(types.EmptyPrivateStateIdentifier))
	assert.NotEqual(t, len(emptyKey), 0)
	notKey, _ := psr.trie.TryGet([]byte(types.PrivateStateIdentifier("notKey")))
	assert.Equal(t, len(notKey), 0)

	//managed state tries updated
	testRoot = testState.IntermediateRoot(false)
	privRoot = privState.IntermediateRoot(false)
	assert.NotEqual(t, testRoot, emptyRoot)
	assert.NotEqual(t, privRoot, emptyRoot)
}

//TestMultiplePSRCommitAndWrite tests that managedStates are updated, trie of states is updated and written to db
func TestMultiplePSRCommitAndWrite(t *testing.T) {
	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	privateStateCacheProvider := privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ := NewMultiplePrivateStateRepository(testdb, testCache, common.Hash{}, privateStateCacheProvider)
	header := &types.Header{Number: big.NewInt(int64(1)), Root: common.Hash{123}}
	block := types.NewBlockWithHeader(header)

	testState, _ := psr.StatePSI(types.PrivateStateIdentifier("test"))
	privState, _ := psr.StatePSI(types.DefaultPrivateStateIdentifier)

	//states have empty tries first
	testRoot := testState.IntermediateRoot(false)
	privRoot := privState.IntermediateRoot(false)
	assert.Equal(t, testRoot, emptyRoot)
	assert.Equal(t, privRoot, emptyRoot)

	//make updates to states
	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState.AddBalance(addr, big.NewInt(int64(i)))
		privState.AddBalance(addr, big.NewInt(int64(i)))
	}

	assert.Equal(t, rawdb.GetPrivateStatesTrieRoot(testdb, block.Root()), common.Hash{})

	psr.CommitAndWrite(false, block)

	//trie root updated and committed to db
	assert.NotEqual(t, psr.trie.Hash(), common.Hash{})
	assert.NotEqual(t, rawdb.GetPrivateStatesTrieRoot(testdb, block.Root()), common.Hash{})

	privateKey, _ := psr.trie.TryGet([]byte(types.DefaultPrivateStateIdentifier))
	assert.NotEqual(t, len(privateKey), 0)
	testKey, _ := psr.trie.TryGet([]byte(types.PrivateStateIdentifier("test")))
	assert.NotEqual(t, len(testKey), 0)
	emptyKey, _ := psr.trie.TryGet([]byte(types.EmptyPrivateStateIdentifier))
	assert.NotEqual(t, len(emptyKey), 0)
	notKey, _ := psr.trie.TryGet([]byte(types.PrivateStateIdentifier("notKey")))
	assert.Equal(t, len(notKey), 0)

	//managed state tries updated
	testRoot = testState.IntermediateRoot(false)
	privRoot = privState.IntermediateRoot(false)
	assert.NotEqual(t, testRoot, emptyRoot)
	assert.NotEqual(t, privRoot, emptyRoot)
}

//TestMultiplePSRIntroduceNewPrivateState tests that a newly introduced private state is branched from the empty state and maintained accordingly
func TestMultiplePSRIntroduceNewPrivateState(t *testing.T) {

	testPS1 := types.PrivateStateIdentifier("PS1")
	testPS2 := types.PrivateStateIdentifier("PS2")
	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	privateStateCacheProvider := privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ := NewMultiplePrivateStateRepository(testdb, testCache, common.Hash{}, privateStateCacheProvider)
	header1 := &types.Header{Number: big.NewInt(int64(1)), Root: common.Hash{123}}
	block1 := types.NewBlockWithHeader(header1)

	header2 := &types.Header{Number: big.NewInt(int64(2)), Root: common.Hash{124}}
	block2 := types.NewBlockWithHeader(header2)

	testState1, _ := psr.StatePSI(testPS1)
	emptyState, _ := psr.StatePSI(types.EmptyPrivateStateIdentifier)

	//states have empty tries first
	testState1Root := testState1.IntermediateRoot(false)
	emptyStateRoot := emptyState.IntermediateRoot(false)
	assert.Equal(t, testState1Root, emptyRoot)
	assert.Equal(t, emptyStateRoot, emptyRoot)

	//make updates to states
	for i := byte(0); i < 10; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState1.AddBalance(addr, big.NewInt(int64(i)))
		emptyState.AddBalance(addr, big.NewInt(int64(i)))
	}

	assert.Equal(t, rawdb.GetPrivateStatesTrieRoot(testdb, block1.Root()), common.Hash{})

	psr.CommitAndWrite(false, block1)

	//trie root updated and committed to db
	psrRootHash1 := psr.trie.Hash()
	assert.NotEqual(t, psrRootHash1, emptyRoot)
	assert.Equal(t, rawdb.GetPrivateStatesTrieRoot(testdb, block1.Root()), psrRootHash1)

	emptyStateRootHash, _ := psr.trie.TryGet([]byte(types.EmptyPrivateStateIdentifier))
	assert.NotEqual(t, len(emptyStateRootHash), 0)
	ps1RootHash, _ := psr.trie.TryGet([]byte(testPS1))
	assert.NotEqual(t, len(ps1RootHash), 0)
	notKeyRootHash, _ := psr.trie.TryGet([]byte(types.PrivateStateIdentifier("notKey")))
	assert.Equal(t, len(notKeyRootHash), 0)

	//managed state tries updated
	testState1Root = testState1.IntermediateRoot(false)
	emptyStateRoot = emptyState.IntermediateRoot(false)
	assert.NotEqual(t, testState1Root, emptyRoot)
	assert.NotEqual(t, emptyStateRoot, emptyRoot)

	// begin adding state at block2
	privateStateCacheProvider = privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ = NewMultiplePrivateStateRepository(testdb, testCache, rawdb.GetPrivateStatesTrieRoot(testdb, block1.Root()), privateStateCacheProvider)

	testState1, _ = psr.StatePSI(testPS1)
	testState2, _ := psr.StatePSI(testPS2)
	emptyState, _ = psr.StatePSI(types.EmptyPrivateStateIdentifier)

	//make updates to states
	for i := byte(10); i < 20; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState1.AddBalance(addr, big.NewInt(int64(i)))
		testState2.AddBalance(addr, big.NewInt(int64(i)))
		emptyState.AddBalance(addr, big.NewInt(int64(i)))
	}

	psr.CommitAndWrite(false, block2)

	privateStateCacheProvider = privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ = NewMultiplePrivateStateRepository(testdb, testCache, rawdb.GetPrivateStatesTrieRoot(testdb, block2.Root()), privateStateCacheProvider)

	testState1, _ = psr.StatePSI(testPS1)
	testState2, _ = psr.StatePSI(testPS2)
	emptyState, _ = psr.StatePSI(types.EmptyPrivateStateIdentifier)

	// we've only added addresses from 10 to 20 to testState2 but since it branched from emptyState it should also contain addresses from 0 to 10
	for i := byte(0); i < 20; i++ {
		addr := common.BytesToAddress([]byte{i})
		assert.True(t, testState1.Exist(addr))
		assert.True(t, testState2.Exist(addr))
		assert.True(t, emptyState.Exist(addr))
	}

	// check that PS2 does not exist in the PSR at block1 height
	privateStateCacheProvider = privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ = NewMultiplePrivateStateRepository(testdb, testCache, rawdb.GetPrivateStatesTrieRoot(testdb, block1.Root()), privateStateCacheProvider)

	emptyStateRootHash, _ = psr.trie.TryGet([]byte(types.EmptyPrivateStateIdentifier))
	assert.NotEqual(t, len(emptyStateRootHash), 0)
	ps1RootHash, _ = psr.trie.TryGet([]byte(testPS1))
	assert.NotEqual(t, len(ps1RootHash), 0)
	ps2RootHash, _ := psr.trie.TryGet([]byte(testPS2))
	assert.Equal(t, len(ps2RootHash), 0)

	// check that PS2 does exist in the PSR at block2 height
	privateStateCacheProvider = privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ = NewMultiplePrivateStateRepository(testdb, testCache, rawdb.GetPrivateStatesTrieRoot(testdb, block2.Root()), privateStateCacheProvider)

	emptyStateRootHash, _ = psr.trie.TryGet([]byte(types.EmptyPrivateStateIdentifier))
	assert.NotEqual(t, len(emptyStateRootHash), 0)
	ps1RootHash, _ = psr.trie.TryGet([]byte(testPS1))
	assert.NotEqual(t, len(ps1RootHash), 0)
	ps2RootHash, _ = psr.trie.TryGet([]byte(testPS2))
	assert.NotEqual(t, len(ps2RootHash), 0)
}

//TestMultiplePSRRemovalFromPrivateState tests that exist no longer picks suicided accounts
func TestMultiplePSRRemovalFromPrivateState(t *testing.T) {

	testPS1 := types.PrivateStateIdentifier("PS1")
	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	privateStateCacheProvider := privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ := NewMultiplePrivateStateRepository(testdb, testCache, common.Hash{}, privateStateCacheProvider)
	header1 := &types.Header{Number: big.NewInt(int64(1)), Root: common.Hash{123}}
	block1 := types.NewBlockWithHeader(header1)

	header2 := &types.Header{Number: big.NewInt(int64(2)), Root: common.Hash{124}}
	block2 := types.NewBlockWithHeader(header2)

	testState1, _ := psr.StatePSI(testPS1)
	emptyState, _ := psr.StatePSI(types.EmptyPrivateStateIdentifier)

	//make updates to states
	for i := byte(0); i < 10; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState1.AddBalance(addr, big.NewInt(int64(i)))
		emptyState.AddBalance(addr, big.NewInt(int64(i)))
	}

	assert.Equal(t, rawdb.GetPrivateStatesTrieRoot(testdb, block1.Root()), common.Hash{})

	psr.CommitAndWrite(false, block1)

	privateStateCacheProvider = privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ = NewMultiplePrivateStateRepository(testdb, testCache, rawdb.GetPrivateStatesTrieRoot(testdb, block1.Root()), privateStateCacheProvider)

	testState1, _ = psr.StatePSI(testPS1)
	emptyState, _ = psr.StatePSI(types.EmptyPrivateStateIdentifier)

	removedAddress := common.BytesToAddress([]byte{1})
	testState1.Suicide(removedAddress)

	psr.CommitAndWrite(false, block2)
	privateStateCacheProvider = privatecache.NewPrivateCacheProvider(testdb, testCache, false)
	psr, _ = NewMultiplePrivateStateRepository(testdb, testCache, rawdb.GetPrivateStatesTrieRoot(testdb, block2.Root()), privateStateCacheProvider)
	testState1, _ = psr.StatePSI(testPS1)
	emptyState, _ = psr.StatePSI(types.EmptyPrivateStateIdentifier)

	assert.False(t, testState1.Exist(removedAddress))
	assert.True(t, emptyState.Exist(removedAddress))
}
