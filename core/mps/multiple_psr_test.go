package mps

import (
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
)

//TestMultiplePSRCopy tests that copying a the PSR object indeed makes the original and
// the copy and their corresponding managed states independent of each other.
func TestMultiplePSRCopy(t *testing.T) {

	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	psr, _ := NewMultiplePrivateStateRepository(params.QuorumMPSTestChainConfig, testdb, testCache, common.Hash{})

	testState, _ := psr.GetPrivateState(types.PrivateStateIdentifier("test"))
	privState, _ := psr.GetPrivateState(types.DefaultPrivateStateIdentifier)

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

	testStateCopy, _ := psrCopy.GetPrivateState(types.PrivateStateIdentifier("test"))
	privStateCopy, _ := psrCopy.GetPrivateState(types.DefaultPrivateStateIdentifier)
	addedState, _ := psrCopy.GetPrivateState(types.PrivateStateIdentifier("added"))

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

	assert.Equal(t, psr.chainConfig, psrCopy.chainConfig)
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
	psr, _ := NewMultiplePrivateStateRepository(params.QuorumMPSTestChainConfig, testdb, testCache, common.Hash{})

	testState, _ := psr.GetPrivateState(types.PrivateStateIdentifier("test"))
	privState, _ := psr.GetPrivateState(types.DefaultPrivateStateIdentifier)

	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState.AddBalance(addr, big.NewInt(int64(i)))
		privState.AddBalance(addr, big.NewInt(int64(i)))
	}
	testState.Finalise(false)
	privState.Finalise(false)

	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		assert.True(t, testState.Exist(addr))
		assert.True(t, privState.Exist(addr))
	}

	psr.Reset()

	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		assert.False(t, testState.Exist(addr))
		assert.False(t, privState.Exist(addr))
	}
}

//TestCreatingManagedStates tests that managed states are created and added to managedState map
func TestCreatingManagedStates(t *testing.T) {
	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	psr, _ := NewMultiplePrivateStateRepository(params.QuorumMPSTestChainConfig, testdb, testCache, common.Hash{})

	//create some managed states
	psr.GetDefaultState()
	psr.GetPrivateState(types.PrivateStateIdentifier("test"))
	psr.GetPrivateState(types.DefaultPrivateStateIdentifier)

	//check if they exist in managedStates map
	assert.Contains(t, psr.managedStates, types.EmptyPrivateStateIdentifier)
	assert.Contains(t, psr.managedStates, types.DefaultPrivateStateIdentifier)
	assert.Contains(t, psr.managedStates, types.PrivateStateIdentifier("test"))
	assert.NotContains(t, psr.managedStates, types.PrivateStateIdentifier("added"))
}
