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

//TestDefaultPSRCopy tests that copying a the PSR object indeed makes the original and
// the copy and their states independent of each other.
func TestDefaultPSRCopy(t *testing.T) {

	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	psr, _ := NewDefaultPrivateStateRepository(params.QuorumTestChainConfig, testdb, testCache, common.Hash{})

	testState, _ := psr.GetDefaultState()

	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState.AddBalance(addr, big.NewInt(int64(i)))
	}
	testState.Finalise(false)

	psrCopy := psr.Copy().(*DefaultPrivateStateRepository)

	testStateCopy, _ := psrCopy.GetDefaultState()

	// modify all in memory
	for i := byte(0); i < 255; i++ {
		testState.AddBalance(common.BytesToAddress([]byte{i}), big.NewInt(2*int64(i)))
		testStateCopy.AddBalance(common.BytesToAddress([]byte{i}), big.NewInt(3*int64(i)))
	}

	// Finalise the changes on all concurrently
	finalise := func(wg *sync.WaitGroup, db *state.StateDB) {
		defer wg.Done()
		db.Finalise(true)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go finalise(&wg, testState)
	go finalise(&wg, testStateCopy)
	wg.Wait()

	assert.Equal(t, psr.chainConfig, psrCopy.chainConfig)
	assert.Equal(t, psr.db, psrCopy.db)
	assert.Equal(t, psr.stateCache, psrCopy.stateCache)

	// Verify that the all states have been updated independently
	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		testObj := testState.GetOrNewStateObject(addr)
		testCopyObj := testStateCopy.GetOrNewStateObject(addr)

		if want := big.NewInt(3 * int64(i)); testObj.Balance().Cmp(want) != 0 {
			t.Errorf("empty obj %d: balance mismatch: have %v, want %v", i, testObj.Balance(), want)
		}
		if want := big.NewInt(4 * int64(i)); testCopyObj.Balance().Cmp(want) != 0 {
			t.Errorf("empty copy obj %d: balance mismatch: have %v, want %v", i, testCopyObj.Balance(), want)
		}
	}
}

//TestDefaultPSRReset tests that state objects are cleared from statedb after reset call
//Any updated stateObjects not committed before reset will be cleared
func TestDefaultPSRReset(t *testing.T) {

	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	psr, _ := NewDefaultPrivateStateRepository(params.QuorumTestChainConfig, testdb, testCache, common.Hash{})

	testState, _ := psr.GetDefaultState()

	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState.AddBalance(addr, big.NewInt(int64(i)))
	}
	testState.Finalise(false)

	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		assert.True(t, testState.Exist(addr))
	}

	psr.Reset()

	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		assert.False(t, testState.Exist(addr))
	}
}

func TestOnlyPrivateStateAccessible(t *testing.T) {
	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	psr, _ := NewDefaultPrivateStateRepository(params.QuorumTestChainConfig, testdb, testCache, common.Hash{})

	privateState, _ := psr.GetDefaultState()
	assert.NotEqual(t, privateState, nil)
	privateState, _ = psr.GetPrivateState(types.DefaultPrivateStateIdentifier)
	assert.NotEqual(t, privateState, nil)
	_, err := psr.GetPrivateState(types.PrivateStateIdentifier("test"))
	assert.Error(t, err, "only the 'private' psi is supported by the default private state manager")
}

//TestDefaultPSRCommitAndWrite tests that statedb is updated but not written to db
func TestDefaultPSRCommit(t *testing.T) {
	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	psr, _ := NewDefaultPrivateStateRepository(params.QuorumTestChainConfig, testdb, testCache, common.Hash{})
	header := &types.Header{Number: big.NewInt(int64(1)), Root: common.Hash{123}}
	block := types.NewBlockWithHeader(header)

	testState, _ := psr.GetDefaultState()

	testRoot := testState.IntermediateRoot(false)
	assert.Equal(t, testRoot, common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"))

	//make updates to states
	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState.AddBalance(addr, big.NewInt(int64(i)))
	}
	assert.Equal(t, rawdb.GetPrivateStateRoot(testdb, block.Root()), common.Hash{})

	psr.Commit(block)

	//private root updated but not committed
	assert.NotEqual(t, psr.root, common.Hash{})
	assert.Equal(t, rawdb.GetPrivateStateRoot(testdb, block.Root()), common.Hash{})

	testRoot = testState.IntermediateRoot(false)
	assert.NotEqual(t, testRoot, common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"))
}

//TestDefaultPSRCommitAndWrite tests that statedb is updated and written to db
func TestDefaultPSRCommitAndWrite(t *testing.T) {
	testdb := rawdb.NewMemoryDatabase()
	testCache := state.NewDatabase(testdb)
	psr, _ := NewDefaultPrivateStateRepository(params.QuorumTestChainConfig, testdb, testCache, common.Hash{})
	header := &types.Header{Number: big.NewInt(int64(1)), Root: common.Hash{123}}
	block := types.NewBlockWithHeader(header)

	testState, _ := psr.GetDefaultState()

	testRoot := testState.IntermediateRoot(false)
	assert.Equal(t, testRoot, common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"))

	//make updates to states
	for i := byte(0); i < 255; i++ {
		addr := common.BytesToAddress([]byte{i})
		testState.AddBalance(addr, big.NewInt(int64(i)))
	}
	assert.Equal(t, rawdb.GetPrivateStateRoot(testdb, block.Root()), common.Hash{})

	psr.CommitAndWrite(block)

	//private root gets committed to db, but isn't updated on psr (only needed for commit)
	assert.Equal(t, psr.root, common.Hash{})
	assert.NotEqual(t, rawdb.GetPrivateStateRoot(testdb, block.Root()), common.Hash{})

	testRoot = testState.IntermediateRoot(false)
	assert.NotEqual(t, testRoot, common.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"))
}
