package core

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"
)

// TestPrefetch_PublicTransactions_Disabled tests that prefetching is not happening
// it confirms the expected read count based on the number of transactions
func TestPrefetch_PublicTransactions_Disabled(t *testing.T) {
	var (
		privateTx         = false
		txCount           = 100
		expectedReadCount = calculateExpectedReadCountWithoutPrefetch(txCount)
		// prefetch enabled?
		prefetchEnabled = false
	)

	mockTxDataArr := createMockTxData(txCount, privateTx)
	chain, blocks, diskDb := createBlocks(mockTxDataArr, prefetchEnabled)
	actualReadCount := insertBlocks(chain, blocks, diskDb)
	publicState, _, _ := chain.State()

	assert.Equal(t, expectedReadCount, actualReadCount)
	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(2), publicState.GetNonce(data.fromAddress))
		assert.Equal(t, common.BigToHash(big.NewInt(15)), publicState.GetState(data.toAddress, common.HexToHash("00")))
	}
}

// TestPrefetch_PublicTransactions_Enabled tests that prefetching is being done
// it confirms the expected read count on DB is higher that what will be on a normal case scenario
func TestPrefetch_PublicTransactions_Enabled(t *testing.T) {
	var (
		privateTx = false
		// number of transactions being sent (1 per block)
		txCount = 100
		// TODO ricardolyn: verify this function
		expectedReadCount = int(float64(calculateExpectedReadCountWithoutPrefetch(txCount)) * 1.10)
		// prefetch enabled?
		prefetchEnabled = true
	)

	mockTxDataArr := createMockTxData(txCount, privateTx)
	chain, blocks, diskDb := createBlocks(mockTxDataArr, prefetchEnabled)
	actualReadCount := insertBlocks(chain, blocks, diskDb)
	publicState, _, _ := chain.State()

	assert.Less(t, expectedReadCount, actualReadCount)
	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(2), publicState.GetNonce(data.fromAddress))
		assert.Equal(t, common.BigToHash(big.NewInt(15)), publicState.GetState(data.toAddress, common.HexToHash("00")))
	}
}

func TestPrefetch_PrivateTransactions_Enabled(t *testing.T) {
	var (
		privateTx = true
		// number of transactions being sent (1 per block)
		txCount = 100
		// TODO ricardolyn: verify this function
		expectedReadCount = int(float64(calculateExpectedReadCountWithoutPrefetch(txCount)) * 1.10)
		// prefetch enabled?
		prefetchEnabled = true
	)

	mockTxDataArr := createMockTxData(txCount, privateTx)
	chain, blocks, diskDb := createBlocks(mockTxDataArr, prefetchEnabled)
	readCount := insertBlocks(chain, blocks, diskDb)
	publicState, privateStateRepo, _ := chain.State()
	privateState, _ := privateStateRepo.DefaultState()

	assert.Equal(t, expectedReadCount, readCount)
	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(1), publicState.GetNonce(data.fromAddress))
		assert.Equal(t, common.Hash{}, publicState.GetState(data.toAddress, common.HexToHash("00")))
		assert.Equal(t, common.BigToHash(big.NewInt(15)), privateState.GetState(data.toAddress, common.HexToHash("00")))
	}
}

// Utility types

type mockTxData struct {
	fromAddress    common.Address
	fromPrivateKey *ecdsa.PrivateKey

	toAddress common.Address

	funds *big.Int

	isPrivate bool
}

// Utility functions

func calculateExpectedReadCountWithoutPrefetch(txCount int) int {
	return txCount * 21
}

func createMockTxData(n int, private bool) []*mockTxData {
	result := make([]*mockTxData, n)
	for i := 0; i < n; i++ {
		fromKey, _ := crypto.GenerateKey()
		fromAddress := crypto.PubkeyToAddress(fromKey.PublicKey)
		result[i] = &mockTxData{
			fromPrivateKey: fromKey,
			fromAddress:    fromAddress,
			funds:          big.NewInt(1000000000),
			isPrivate:      private}
	}
	return result
}

func createBlocks(mockTxDataArr []*mockTxData, usePrefetch bool) (*BlockChain, []*types.Block, ethdb.Database) {
	var (
		// Generate a canonical chain to act as the main dataset
		engine      = ethash.NewFaker()
		db          = rawdb.NewMemoryDatabase()
		cacheConfig = *defaultCacheConfig
	)
	// Generate different accounts to create transactions from randomly. with prefetching, the number of writes should be higher than 11+n where n is the number of tx
	cacheConfig.TrieCleanNoPrefetch = !usePrefetch

	allocation := GenesisAlloc{}
	for _, data := range mockTxDataArr {
		allocation[data.fromAddress] = GenesisAccount{
			Balance: data.funds,
			Nonce:   0,
		}
	}
	gspec := &Genesis{
		Config: params.QuorumTestChainConfig,
		Alloc:  allocation,
	}
	genesis := gspec.MustCommit(db)

	blocks, _ := GenerateChain(params.QuorumTestChainConfig, genesis, engine, db, len(mockTxDataArr), func(i int, b *BlockGen) {
		b.SetCoinbase(common.Address{1})
		mockTxData := mockTxDataArr[i]
		var signer types.Signer = types.HomesteadSigner{}

		createTransaction := types.NewContractCreation(0, common.Big0, uint64(3000000), common.Big0, c1.create(big.NewInt(10)))
		if mockTxData.isPrivate {
			createTransaction.SetPrivate()
			signer = types.QuorumPrivateTxSigner{}
		}
		tx1, _ := types.SignTx(createTransaction, signer, mockTxData.fromPrivateKey)
		b.AddTx(tx1)
		mockTxData.toAddress = b.receipts[0].ContractAddress
	})
	blocks2, _ := GenerateChain(params.QuorumTestChainConfig, blocks[len(blocks)-1], engine, db, len(mockTxDataArr), func(i int, b *BlockGen) {
		b.SetCoinbase(common.Address{1})
		mockTxData := mockTxDataArr[i]
		var signer types.Signer = types.HomesteadSigner{}

		setTransaction := types.NewTransaction(1, mockTxData.toAddress, common.Big0, uint64(3000000), common.Big0, c1.set(15))
		if mockTxData.isPrivate {
			setTransaction.SetPrivate()
			signer = types.QuorumPrivateTxSigner{}
		}
		tx2, _ := types.SignTx(setTransaction, signer, mockTxData.fromPrivateKey)
		b.AddTx(tx2)

	})
	blocks = append(blocks, blocks2...)

	// Import the canonical chain
	diskdb := rawdb.NewDatabase(memorydb.NewMetered())
	gspec.MustCommit(diskdb)
	chain, _ := NewBlockChain(diskdb, &cacheConfig, params.TestChainConfig, engine, vm.Config{
		Debug:  true,
		Tracer: vm.NewJSONLogger(nil, os.Stdout),
	}, nil, nil)
	//if err != nil {
	//	t.Fatalf("failed to create tester chain: %v", err)
	//}
	return chain, blocks, diskdb
}

func insertBlocks(chain *BlockChain, blocks []*types.Block, diskDb ethdb.Database) (readCount int) {
	const propertyName = "readCount"

	startReadCountStr, _ := diskDb.Stat(propertyName)
	chain.InsertChain(blocks)
	time.Sleep(1 * time.Second)

	endReadCountStr, _ := diskDb.Stat(propertyName)
	startReadCount, _ := strconv.Atoi(startReadCountStr)
	endReadCount, _ := strconv.Atoi(endReadCountStr)

	return endReadCount - startReadCount
}
