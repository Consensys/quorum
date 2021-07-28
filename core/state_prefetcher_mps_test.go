package core

import (
	"crypto/ecdsa"
	"encoding/base64"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/private"
	privateEngine "github.com/ethereum/go-ethereum/private/engine"
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

	privateCreatePayload = common.BytesToEncryptedPayloadHash(common.Hex2Bytes("41a982be5d1f3d92d57487d7d9a905c1d92d3353570730464639affc964bcc83ea24e5b449140a2216ecc3f1d11d3dfd3663c6a9a4f18a7c837a9e4d8bfc81ce"))
	privateSetPayload    = common.BytesToEncryptedPayloadHash(common.Hex2Bytes("93f769208aa744b6d65310ab191f1fe22f8508ad069810f06889381b89d8c03ade785c7b14230439673f76e08ec84bad611d95d1cbb66dbcf548acbf93db0296"))
)

func TestPrefetch_PublicTransaction(t *testing.T) {
	var (
		engine    = ethash.NewFaker()
		interrupt = uint32(0)
		privateTx = false
		txCount   = 1
	)

	mockTxDataArr := createMockTxData(txCount, privateTx)
	chain, gspec := createBlockchain(params.QuorumTestChainConfig, mockTxDataArr)
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
		isPrivate = true
		txCount   = 1
		mockCtrl  = gomock.NewController(t)
	)
	defer mockCtrl.Finish()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)
	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockptm.EXPECT().Receive(privateCreatePayload).Return("", []string{}, contractCreateBytes, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(privateSetPayload).Return("", []string{}, contractSetBytes, nil, nil).AnyTimes()

	mockTxDataArr := createMockTxData(txCount, isPrivate)
	chain, gspec := createBlockchain(params.QuorumTestChainConfig, mockTxDataArr)
	_, minedBlock, futureBlock := createBlocks(gspec, mockTxDataArr)

	// Import the canonical chain
	if n, err := chain.InsertChain(types.Blocks{minedBlock, futureBlock}); n == 0 || err != nil {
		t.Fatal("Failure when inserting blocks", "n", n, "err", err)
	}
	prefetcher := newStatePrefetcher(gspec.Config, chain, engine)

	throwaway, _ := state.New(minedBlock.Root(), chain.stateCache, chain.snaps)
	privateRepo, _ := chain.PrivateStateManager().StateRepository(minedBlock.Root())
	throwawayRepo := privateRepo.Copy()

	// When
	prefetcher.Prefetch(futureBlock, throwaway, throwawayRepo, vm.Config{}, &interrupt)

	// Then
	throwawayPrivateState, _ := throwawayRepo.DefaultState()
	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(2), throwaway.GetNonce(data.fromAddress))
		assert.Equal(t, common.Hash{}, throwaway.GetState(data.toAddress, common.HexToHash("00")))
		assert.Equal(t, common.BigToHash(big.NewInt(15)), throwawayPrivateState.GetState(data.toAddress, common.HexToHash("00")))
	}
}

func TestPrefetch_PrivateMPSTransaction(t *testing.T) {
	var (
		engine    = ethash.NewFaker()
		interrupt = uint32(0)
		isPrivate = true
		txCount   = 1
		mockCtrl  = gomock.NewController(t)
	)
	defer mockCtrl.Finish()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)
	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockptm.EXPECT().Receive(common.EncryptedPayloadHash{}).Return("", []string{}, nil, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(privateCreatePayload).Return("", []string{"psi2"}, contractCreateBytes, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(privateSetPayload).Return("", []string{"psi2"}, contractSetBytes, nil, nil).AnyTimes()
	mockptm.EXPECT().HasFeature(privateEngine.MultiplePrivateStates).Return(true).AnyTimes()
	mockptm.EXPECT().Groups().Return([]privateEngine.PrivacyGroup{
		{
			Type:           privateEngine.PrivacyGroupResident,
			Name:           PSI1PSM.Name,
			PrivacyGroupId: base64.StdEncoding.EncodeToString([]byte(PSI1PSM.ID)),
			Description:    "Resident Group 1",
			From:           "",
			Members:        []string{"psi1"},
		},
		{
			Type:           privateEngine.PrivacyGroupResident,
			Name:           PSI2PSM.Name,
			PrivacyGroupId: base64.StdEncoding.EncodeToString([]byte(PSI2PSM.ID)),
			Description:    "Resident Group 2",
			From:           "",
			Members:        []string{"psi2"},
		},
	}, nil)

	mockTxDataArr := createMockTxData(txCount, isPrivate)
	chain, gspec := createBlockchain(params.QuorumMPSTestChainConfig, mockTxDataArr)
	_, minedBlock, futureBlock := createBlocks(gspec, mockTxDataArr)

	// Import the canonical chain
	if n, err := chain.InsertChain(types.Blocks{minedBlock, futureBlock}); n == 0 || err != nil {
		t.Fatal("Failure when inserting blocks", "n", n, "err", err)
	}
	prefetcher := newStatePrefetcher(gspec.Config, chain, engine)

	throwaway, _ := state.New(minedBlock.Root(), chain.stateCache, chain.snaps)
	privateRepo, _ := chain.PrivateStateManager().StateRepository(minedBlock.Root())
	throwawayRepo := privateRepo.Copy()

	// When
	prefetcher.Prefetch(futureBlock, throwaway, throwawayRepo, vm.Config{}, &interrupt)

	// Then
	throwawayDefaultPrivateState, _ := throwawayRepo.DefaultState()
	throwawayPS1PrivateState, _ := throwawayRepo.StatePSI(PSI1PSM.ID)
	throwawayPS2PrivateState, _ := throwawayRepo.StatePSI(PSI2PSM.ID)
	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(2), throwaway.GetNonce(data.fromAddress))
		assert.Equal(t, common.Hash{}, throwaway.GetState(data.toAddress, common.HexToHash("00")))
		assert.Equal(t, common.Hash{}, throwawayDefaultPrivateState.GetState(data.toAddress, common.HexToHash("00")))
		assert.Equal(t, common.Hash{}, throwawayPS1PrivateState.GetState(data.toAddress, common.HexToHash("00")))
		assert.Equal(t, common.BigToHash(big.NewInt(15)), throwawayPS2PrivateState.GetState(data.toAddress, common.HexToHash("00")))
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

func createBlockchain(chainConfig *params.ChainConfig, mockTxDataArr []*mockTxData) (*BlockChain, *Genesis) {
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
		Config: chainConfig,
		Alloc:  allocation,
	}
	diskdb := rawdb.NewMemoryDatabase()
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
		engine      = ethash.NewFaker()
		temporaryDb = rawdb.NewMemoryDatabase()
	)
	genesisBlock := gspec.MustCommit(temporaryDb)
	minedBlocks, _ := GenerateChain(gspec.Config, genesisBlock, engine, temporaryDb, 1, func(i int, b *BlockGen) {
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
			signedTx, _ := types.SignTx(createTransaction, signer, mockTxData.fromPrivateKey)
			b.AddTx(signedTx)

			// save the contract address to use when calling `set()`
			mockTxData.toAddress = b.receipts[0].ContractAddress
		}
	})
	futureBlocks, _ := GenerateChain(gspec.Config, minedBlocks[0], engine, temporaryDb, 1, func(i int, b *BlockGen) {
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
			signedTx, _ := types.SignTx(setTransaction, signer, mockTxData.fromPrivateKey)
			b.AddTx(signedTx)
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
