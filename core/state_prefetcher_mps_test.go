package core

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/private"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"math/big"
	"os"
	"testing"
)

var (
	contractInitValue = int64(10)
	contractSetValue  = int64(15)

	contractCreateBytes = c1.create(big.NewInt(contractInitValue))
	contractSetBytes    = c1.set(contractSetValue)

	privateCreatePayload = common.BytesToEncryptedPayloadHash(common.Hex2Bytes("c3bb5ad6a32f5fa3d3e9fd5d3f85b9a7e51ec126c86d984cc3aecb5618d7fe72"))
	privateSetPayload    = common.BytesToEncryptedPayloadHash(common.Hex2Bytes("9975a239f0e897656145022ed3714ecc9218b380e95a946eea8a47c9c9058e9b"))
)

func TestPrefetch_PublicTransaction(t *testing.T) {
	var (
		engine    = ethash.NewFaker()
		interrupt = uint32(0)
		privateTx = false
		txCount   = 1
	)

	mockTxDataArr := createMockTxData(txCount, privateTx)
	chain, gspec := createBlockchain(mockTxDataArr)
	_, minedBlock, futureBlock := createBlocks(gspec, mockTxDataArr)

	// Import the canonical chain
	chain.InsertChain(types.Blocks{minedBlock, futureBlock})

	prefetcher := newStatePrefetcher(gspec.Config, chain, engine)

	throwaway, _ := state.New(minedBlock.Root(), chain.stateCache, chain.snaps)
	privateRepo, _ := chain.PrivateStateManager().StateRepository(minedBlock.Root())
	throwawayRepo := privateRepo.Copy()

	prefetcher.Prefetch(futureBlock, throwaway, throwawayRepo, vm.Config{}, &interrupt)

	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(2), throwaway.GetNonce(data.fromAddress))
		assert.Equal(t, common.BigToHash(big.NewInt(15)), throwaway.GetState(data.toAddress, common.HexToHash("00")))
	}
}

func TestPrefetch_PrivateDualStateTransaction(t *testing.T) {
	var (
		engine    = ethash.NewFaker()
		interrupt = uint32(0)
		privateTx = true
		txCount   = 1
	)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)
	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockptm.EXPECT().Receive(privateCreatePayload).Return("", []string{}, contractCreateBytes, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(privateSetPayload).Return("", []string{}, contractSetBytes, nil, nil).AnyTimes()

	mockTxDataArr := createMockTxData(txCount, privateTx)
	chain, gspec := createBlockchain(mockTxDataArr)
	genesisBlock, minedBlock, futureBlock := createBlocks(gspec, mockTxDataArr)

	assert.Equal(t, genesisBlock.ParentHash(), common.Hash{})
	assert.Equal(t, minedBlock.ParentHash(), genesisBlock.Hash())
	assert.Equal(t, futureBlock.ParentHash(), minedBlock.Hash())
	// Import the canonical chain

	if n, err := chain.InsertChain(types.Blocks{minedBlock, futureBlock}); n == 0 || err != nil {
		t.Fatal("Failure when inserting blocks", "n", n, "err", err)
	}

	prefetcher := newStatePrefetcher(gspec.Config, chain, engine)

	throwaway, _ := state.New(minedBlock.Root(), chain.stateCache, chain.snaps)
	privateRepo, _ := chain.PrivateStateManager().StateRepository(minedBlock.Root())
	throwawayRepo := privateRepo.Copy()

	prefetcher.Prefetch(futureBlock, throwaway, throwawayRepo, vm.Config{}, &interrupt)

	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(2), throwaway.GetNonce(data.fromAddress))
		assert.Equal(t, common.BigToHash(big.NewInt(15)), throwaway.GetState(data.toAddress, common.HexToHash("00")))
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

func createBlockchain(mockTxDataArr []*mockTxData) (*BlockChain, *Genesis) {
	var (
		// Generate a canonical chain to act as the main dataset
		engine      = ethash.NewFaker()
		cacheConfig = *defaultCacheConfig
	)
	// We are going to manually run prefetch
	cacheConfig.TrieCleanNoPrefetch = true

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
	diskdb := rawdb.NewDatabase(memorydb.NewMetered())
	gspec.MustCommit(diskdb)
	vmConfig := vm.Config{
		Debug:  true,
		Tracer: vm.NewJSONLogger(nil, os.Stdout),
	}
	chain, _ := NewBlockChain(diskdb, &cacheConfig, gspec.Config, engine, vmConfig, nil, nil)

	return chain, gspec
}

func createBlocks(gspec *Genesis, mockTxDataArr []*mockTxData) (*types.Block, *types.Block, *types.Block) {
	var (
		// Generate a canonical chain to act as the main dataset
		engine = ethash.NewFaker()
		db     = rawdb.NewMemoryDatabase()
	)
	genesisBlock := gspec.MustCommit(db)
	minedBlocks, _ := GenerateChain(gspec.Config, genesisBlock, engine, db, 1, func(i int, b *BlockGen) {
		b.SetCoinbase(common.Address{1})
		var signer types.Signer = types.HomesteadSigner{}
		for _, mockTxData := range mockTxDataArr {
			data := contractCreateBytes
			if mockTxData.isPrivate {
				data = privateCreatePayload.Bytes()
			}
			createTransaction := types.NewContractCreation(0, common.Big0, uint64(3000000), common.Big0, data)
			if mockTxData.isPrivate {
				createTransaction.SetPrivate()
				signer = types.QuorumPrivateTxSigner{}
			}
			tx1, _ := types.SignTx(createTransaction, signer, mockTxData.fromPrivateKey)
			b.AddTx(tx1)
			mockTxData.toAddress = b.receipts[0].ContractAddress
		}
	})
	futureBlocks, _ := GenerateChain(gspec.Config, minedBlocks[0], engine, db, 1, func(i int, b *BlockGen) {
		b.SetCoinbase(common.Address{1})
		var signer types.Signer = types.HomesteadSigner{}
		for _, mockTxData := range mockTxDataArr {
			data := contractSetBytes
			if mockTxData.isPrivate {
				data = privateSetPayload.Bytes()
			}
			setTransaction := types.NewTransaction(1, mockTxData.toAddress, common.Big0, uint64(3000000), common.Big0, data)
			if mockTxData.isPrivate {
				setTransaction.SetPrivate()
				signer = types.QuorumPrivateTxSigner{}
			}
			tx2, _ := types.SignTx(setTransaction, signer, mockTxData.fromPrivateKey)
			b.AddTx(tx2)
		}
	})

	return genesisBlock, minedBlocks[0], futureBlocks[0]
}

//
//func insertBlocks(chain *BlockChain, blocks []*types.Block, diskDb ethdb.Database) (readCount int) {
//	const propertyName = "readCount"
//
//	startReadCountStr, _ := diskDb.Stat(propertyName)
//	chain.InsertChain(blocks)
//	time.Sleep(1 * time.Second)
//
//	endReadCountStr, _ := diskDb.Stat(propertyName)
//	startReadCount, _ := strconv.Atoi(startReadCountStr)
//	endReadCount, _ := strconv.Atoi(endReadCountStr)
//
//	return endReadCount - startReadCount
//}
