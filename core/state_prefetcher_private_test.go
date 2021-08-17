package core

import (
	"crypto/ecdsa"
	"encoding/base64"
	"math/big"
	"os"
	"testing"

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
)

var (
	contractDeployed = &contract{
		name:     "contractDeployed",
		abi:      mustParse(contractDeployedDefinition),
		bytecode: common.Hex2Bytes("608060405234801561001057600080fd5b506040516020806105a88339810180604052602081101561003057600080fd5b81019080805190602001909291905050508060008190555050610550806100586000396000f3fe608060405260043610610051576000357c01000000000000000000000000000000000000000000000000000000009004806360fe47b1146100565780636d4ce63c146100a5578063d7139463146100d0575b600080fd5b34801561006257600080fd5b5061008f6004803603602081101561007957600080fd5b810190808035906020019092919050505061010b565b6040518082815260200191505060405180910390f35b3480156100b157600080fd5b506100ba61011e565b6040518082815260200191505060405180910390f35b3480156100dc57600080fd5b50610109600480360360208110156100f357600080fd5b8101908080359060200190929190505050610127565b005b6000816000819055506000549050919050565b60008054905090565b600030610132610212565b808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051809103906000f080158015610184573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff166360fe47b1836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b1580156101f657600080fd5b505af115801561020a573d6000803e3d6000fd5b505050505050565b604051610302806102238339019056fe608060405234801561001057600080fd5b506040516020806103028339810180604052602081101561003057600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610271806100916000396000f3fe608060405260043610610046576000357c01000000000000000000000000000000000000000000000000000000009004806360fe47b11461004b5780636d4ce63c14610086575b600080fd5b34801561005757600080fd5b506100846004803603602081101561006e57600080fd5b81019080803590602001909291905050506100b1565b005b34801561009257600080fd5b5061009b610180565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360fe47b1826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561014157600080fd5b505af1158015610155573d6000803e3d6000fd5b505050506040513d602081101561016b57600080fd5b81019080805190602001909291905050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d4ce63c6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561020557600080fd5b505afa158015610219573d6000803e3d6000fd5b505050506040513d602081101561022f57600080fd5b810190808051906020019092919050505090509056fea165627a7a72305820a537f4c360ce5c6f55523298e314e6456e5c3e02c170563751dfda37d3aeddb30029a165627a7a7230582060396bfff29d2dfc5a9f4216bfba5e24d031d54fd4b26ebebde1a26c59df0c1e0029"),
	}

	contractDeploymentCount = 1

	contractArgumentInitValue = int64(10)
	contractArgumentSetValue  = int64(15)

	contractCreateABIPayloadBytes = contractDeployed.create(big.NewInt(contractArgumentInitValue))
	contractSetABIPayloadBytes    = contractDeployed.set(contractArgumentSetValue)

	encryptedPayloadHashForContractDeployment = common.BytesToEncryptedPayloadHash(common.Hex2Bytes("41a982be5d1f3d92d57487d7d9a905c1d92d3353570730464639affc964bcc83ea24e5b449140a2216ecc3f1d11d3dfd3663c6a9a4f18a7c837a9e4d8bfc81ce"))
	encryptedPayloadHashForSetFunction        = common.BytesToEncryptedPayloadHash(common.Hex2Bytes("93f769208aa744b6d65310ab191f1fe22f8508ad069810f06889381b89d8c03ade785c7b14230439673f76e08ec84bad611d95d1cbb66dbcf548acbf93db0296"))

	slot0OnAccountStorage = common.HexToHash("00")
)

func TestPrefetch_PublicTransaction(t *testing.T) {
	var (
		engine = ethash.NewFaker()
	)
	mockTxDataArr := createMockTxData(contractDeploymentCount, Public)
	chain, gspec := createBlockchain(params.QuorumTestChainConfig, mockTxDataArr)
	minedBlock, futureBlock := createBlocks(chain, gspec, mockTxDataArr, nil)

	// Import the canonical chain
	chain.InsertChain(types.Blocks{minedBlock})

	prefetcher := newStatePrefetcher(gspec.Config, chain, engine)

	throwaway, _ := state.New(minedBlock.Root(), chain.stateCache, chain.snaps)
	privateRepo, _ := chain.PrivateStateManager().StateRepository(minedBlock.Root())
	throwawayRepo := privateRepo.Copy()

	// When
	prefetcher.Prefetch(futureBlock, throwaway, throwawayRepo, vm.Config{}, nil)

	// Then
	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(2), throwaway.GetNonce(data.fromAddress))
		assert.Equal(t, common.BigToHash(big.NewInt(contractArgumentSetValue)), throwaway.GetState(data.toAddress, slot0OnAccountStorage))
	}
}

func TestPrefetch_PrivateDualStateTransaction(t *testing.T) {
	var (
		engine   = ethash.NewFaker()
		mockCtrl = gomock.NewController(t)
	)
	defer mockCtrl.Finish()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)
	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockptm.EXPECT().Receive(encryptedPayloadHashForContractDeployment).Return("", []string{}, contractCreateABIPayloadBytes, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(encryptedPayloadHashForSetFunction).Return("", []string{}, contractSetABIPayloadBytes, nil, nil).AnyTimes()

	mockTxDataArr := createMockTxData(contractDeploymentCount, Private)
	chain, gspec := createBlockchain(params.QuorumTestChainConfig, mockTxDataArr)
	minedBlock, futureBlock := createBlocks(chain, gspec, mockTxDataArr, nil)

	// Import the canonical chain
	if n, err := chain.InsertChain(types.Blocks{minedBlock}); n == 0 || err != nil {
		t.Fatal("Failure when inserting blocks", "n", n, "err", err)
	}
	prefetcher := newStatePrefetcher(gspec.Config, chain, engine)

	throwaway, _ := state.New(minedBlock.Root(), chain.stateCache, chain.snaps)
	privateRepo, _ := chain.PrivateStateManager().StateRepository(minedBlock.Root())
	throwawayRepo := privateRepo.Copy()

	// When
	prefetcher.Prefetch(futureBlock, throwaway, throwawayRepo, vm.Config{}, nil)

	// Then
	throwawayPrivateState, _ := throwawayRepo.DefaultState()
	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(2), throwaway.GetNonce(data.fromAddress))
		assert.Equal(t, common.Hash{}, throwaway.GetState(data.toAddress, slot0OnAccountStorage))
		assert.Equal(t, common.BigToHash(big.NewInt(contractArgumentSetValue)), throwawayPrivateState.GetState(data.toAddress, slot0OnAccountStorage))
	}
}

func TestPrefetch_PrivateMPSTransaction(t *testing.T) {
	var (
		engine   = ethash.NewFaker()
		mockCtrl = gomock.NewController(t)
	)
	defer mockCtrl.Finish()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)
	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockptm.EXPECT().Receive(common.EncryptedPayloadHash{}).Return("", []string{}, nil, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(encryptedPayloadHashForContractDeployment).Return("", []string{"BBB"}, contractCreateABIPayloadBytes, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(encryptedPayloadHashForSetFunction).Return("", []string{"BBB"}, contractSetABIPayloadBytes, nil, nil).AnyTimes()
	mockptm.EXPECT().HasFeature(privateEngine.MultiplePrivateStates).Return(true).AnyTimes()
	mockptm.EXPECT().Groups().Return([]privateEngine.PrivacyGroup{
		{
			Type:           privateEngine.PrivacyGroupResident,
			Name:           PSI1PSM.Name,
			PrivacyGroupId: base64.StdEncoding.EncodeToString([]byte(PSI1PSM.ID)),
			Description:    "Resident Group 1",
			From:           "",
			Members:        []string{"AAA"},
		},
		{
			Type:           privateEngine.PrivacyGroupResident,
			Name:           PSI2PSM.Name,
			PrivacyGroupId: base64.StdEncoding.EncodeToString([]byte(PSI2PSM.ID)),
			Description:    "Resident Group 2",
			From:           "",
			Members:        []string{"BBB"},
		},
	}, nil)

	mockTxDataArr := createMockTxData(contractDeploymentCount, Private)
	chain, gspec := createBlockchain(params.QuorumMPSTestChainConfig, mockTxDataArr)
	minedBlock, futureBlock := createBlocks(chain, gspec, mockTxDataArr, nil)

	// Import the canonical chain
	if n, err := chain.InsertChain(types.Blocks{minedBlock}); n == 0 || err != nil {
		t.Fatal("Failure when inserting blocks", "n", n, "err", err)
	}
	prefetcher := newStatePrefetcher(gspec.Config, chain, engine)

	throwaway, _ := state.New(minedBlock.Root(), chain.stateCache, chain.snaps)
	privateRepo, _ := chain.PrivateStateManager().StateRepository(minedBlock.Root())
	throwawayRepo := privateRepo.Copy()

	// When
	prefetcher.Prefetch(futureBlock, throwaway, throwawayRepo, vm.Config{}, nil)

	// Then
	throwawayDefaultPrivateState, _ := throwawayRepo.DefaultState()
	throwawayPS1PrivateState, _ := throwawayRepo.StatePSI(PSI1PSM.ID)
	throwawayPS2PrivateState, _ := throwawayRepo.StatePSI(PSI2PSM.ID)
	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(2), throwaway.GetNonce(data.fromAddress))
		assert.Equal(t, common.Hash{}, throwaway.GetState(data.toAddress, slot0OnAccountStorage))
		assert.Equal(t, common.Hash{}, throwawayDefaultPrivateState.GetState(data.toAddress, slot0OnAccountStorage))
		assert.Equal(t, common.Hash{}, throwawayPS1PrivateState.GetState(data.toAddress, slot0OnAccountStorage))
		assert.Equal(t, common.BigToHash(big.NewInt(contractArgumentSetValue)), throwawayPS2PrivateState.GetState(data.toAddress, slot0OnAccountStorage))
	}
}

func TestPrefetch_PrivateDualState_PMTTransaction(t *testing.T) {
	var (
		engine   = ethash.NewFaker()
		mockCtrl = gomock.NewController(t)
	)
	defer mockCtrl.Finish()

	// Activate PMT
	params.QuorumTestChainConfig.PrivacyPrecompileBlock = big.NewInt(0)
	defer func() { params.QuorumTestChainConfig.PrivacyPrecompileBlock = nil }()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)
	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockTxDataArr := createMockTxData(contractDeploymentCount, PMT)

	mockptm.EXPECT().Receive(encryptedPayloadHashForContractDeployment).Return("", []string{}, contractCreateABIPayloadBytes, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(encryptedPayloadHashForSetFunction).Return("", []string{}, contractSetABIPayloadBytes, nil, nil).AnyTimes()

	chain, gspec := createBlockchain(params.QuorumTestChainConfig, mockTxDataArr)
	minedBlock, futureBlock := createBlocks(chain, gspec, mockTxDataArr, func(outerTx *types.Transaction, mockTxData *mockTxData) {
		enclaveHash := common.BytesToEncryptedPayloadHash(outerTx.Data())
		mockptm.EXPECT().Receive(enclaveHash).DoAndReturn(func(hash common.EncryptedPayloadHash) (string, []string, []byte, *privateEngine.ExtraMetadata, error) {
			innerTx := types.NewTransaction(1, mockTxData.toAddress, common.Big0, uint64(3000000), common.Big0, encryptedPayloadHashForSetFunction.Bytes())
			innerTx.SetPrivate()
			signedTx, _ := types.SignTx(innerTx, types.QuorumPrivateTxSigner{}, mockTxData.fromPrivateKey)
			jsonSignedTx, _ := signedTx.MarshalJSON()
			return "", []string{}, jsonSignedTx, nil, nil
		}).AnyTimes()
	})

	// Import the canonical chain
	if n, err := chain.InsertChain(types.Blocks{minedBlock}); n == 0 || err != nil {
		t.Fatal("Failure when inserting blocks", "n", n, "err", err)
	}
	prefetcher := newStatePrefetcher(gspec.Config, chain, engine)

	throwaway, _ := state.New(minedBlock.Root(), chain.stateCache, chain.snaps)
	privateRepo, _ := chain.PrivateStateManager().StateRepository(minedBlock.Root())
	throwawayRepo := privateRepo.Copy()

	// When
	prefetcher.Prefetch(futureBlock, throwaway, throwawayRepo, vm.Config{}, nil)

	// Then
	throwawayPrivateState, _ := throwawayRepo.DefaultState()
	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(2), throwaway.GetNonce(data.fromAddress))
		assert.Equal(t, common.Hash{}, throwaway.GetState(data.toAddress, slot0OnAccountStorage))
		assert.Equal(t, common.BigToHash(big.NewInt(contractArgumentSetValue)), throwawayPrivateState.GetState(data.toAddress, slot0OnAccountStorage))
	}
}

func TestPrefetch_PrivateMPS_PMTTransaction(t *testing.T) {
	var (
		engine   = ethash.NewFaker()
		mockCtrl = gomock.NewController(t)
	)
	defer mockCtrl.Finish()

	// Activate PMT
	params.QuorumMPSTestChainConfig.PrivacyPrecompileBlock = big.NewInt(0)
	defer func() { params.QuorumMPSTestChainConfig.PrivacyPrecompileBlock = nil }()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)
	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockTxDataArr := createMockTxData(contractDeploymentCount, PMT)

	mockptm.EXPECT().Receive(common.EncryptedPayloadHash{}).Return("", []string{}, nil, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(encryptedPayloadHashForContractDeployment).Return("", []string{"BBB"}, contractCreateABIPayloadBytes, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(encryptedPayloadHashForSetFunction).Return("", []string{"BBB"}, contractSetABIPayloadBytes, nil, nil).AnyTimes()
	mockptm.EXPECT().HasFeature(privateEngine.MultiplePrivateStates).Return(true).AnyTimes()
	mockptm.EXPECT().Groups().Return([]privateEngine.PrivacyGroup{
		{
			Type:           privateEngine.PrivacyGroupResident,
			Name:           PSI1PSM.Name,
			PrivacyGroupId: base64.StdEncoding.EncodeToString([]byte(PSI1PSM.ID)),
			Description:    "Resident Group 1",
			From:           "",
			Members:        []string{"AAA"},
		},
		{
			Type:           privateEngine.PrivacyGroupResident,
			Name:           PSI2PSM.Name,
			PrivacyGroupId: base64.StdEncoding.EncodeToString([]byte(PSI2PSM.ID)),
			Description:    "Resident Group 2",
			From:           "",
			Members:        []string{"BBB"},
		},
	}, nil)

	chain, gspec := createBlockchain(params.QuorumMPSTestChainConfig, mockTxDataArr)
	minedBlock, futureBlock := createBlocks(chain, gspec, mockTxDataArr, func(outerTx *types.Transaction, mockTxData *mockTxData) {
		enclaveHash := common.BytesToEncryptedPayloadHash(outerTx.Data())
		mockptm.EXPECT().Receive(enclaveHash).DoAndReturn(func(hash common.EncryptedPayloadHash) (string, []string, []byte, *privateEngine.ExtraMetadata, error) {
			innerTx := types.NewTransaction(1, mockTxData.toAddress, common.Big0, uint64(3000000), common.Big0, encryptedPayloadHashForSetFunction.Bytes())
			innerTx.SetPrivate()
			signedTx, _ := types.SignTx(innerTx, types.QuorumPrivateTxSigner{}, mockTxData.fromPrivateKey)
			jsonSignedTx, _ := signedTx.MarshalJSON()
			return "", []string{}, jsonSignedTx, nil, nil
		}).AnyTimes()
	})

	// Import the canonical chain
	if n, err := chain.InsertChain(types.Blocks{minedBlock}); n == 0 || err != nil {
		t.Fatal("Failure when inserting blocks", "n", n, "err", err)
	}
	prefetcher := newStatePrefetcher(gspec.Config, chain, engine)

	throwaway, _ := state.New(minedBlock.Root(), chain.stateCache, chain.snaps)
	privateRepo, _ := chain.PrivateStateManager().StateRepository(minedBlock.Root())
	throwawayRepo := privateRepo.Copy()

	// When
	prefetcher.Prefetch(futureBlock, throwaway, throwawayRepo, vm.Config{}, nil)

	// Then
	throwawayDefaultPrivateState, _ := throwawayRepo.DefaultState()
	throwawayPS1PrivateState, _ := throwawayRepo.StatePSI(PSI1PSM.ID)
	throwawayPS2PrivateState, _ := throwawayRepo.StatePSI(PSI2PSM.ID)
	for _, data := range mockTxDataArr {
		assert.Equal(t, uint64(2), throwaway.GetNonce(data.fromAddress))
		assert.Equal(t, common.Hash{}, throwaway.GetState(data.toAddress, slot0OnAccountStorage))
		assert.Equal(t, common.Hash{}, throwawayDefaultPrivateState.GetState(data.toAddress, slot0OnAccountStorage))
		assert.Equal(t, common.Hash{}, throwawayPS1PrivateState.GetState(data.toAddress, slot0OnAccountStorage))
		assert.Equal(t, common.BigToHash(big.NewInt(contractArgumentSetValue)), throwawayPS2PrivateState.GetState(data.toAddress, slot0OnAccountStorage))
	}
}

// Utility types

type txType int

const (
	Public txType = iota
	Private
	PMT
)

type mockTxData struct {
	fromAddress    common.Address
	fromPrivateKey *ecdsa.PrivateKey

	toAddress common.Address

	funds *big.Int

	txType txType
}

// Utility functions

func createMockTxData(n int, txType txType) []*mockTxData {
	result := make([]*mockTxData, n)
	for i := 0; i < n; i++ {
		fromKey, _ := crypto.GenerateKey()
		fromAddress := crypto.PubkeyToAddress(fromKey.PublicKey)
		result[i] = &mockTxData{
			fromPrivateKey: fromKey,
			fromAddress:    fromAddress,
			funds:          big.NewInt(1000000000),
			txType:         txType}
	}
	return result
}

func createBlockchain(chainConfig *params.ChainConfig, mockTxDataArr []*mockTxData) (*BlockChain, *Genesis) {
	var (
		engine      = ethash.NewFaker()
		cacheConfig = *defaultCacheConfig
	)
	// Disable prefetching. We are going to manually run prefetch
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

func createBlocks(chain *BlockChain, gspec *Genesis, mockTxDataArr []*mockTxData, decorateSetTransaction func(*types.Transaction, *mockTxData)) (*types.Block, *types.Block) {
	var (
		engine      = ethash.NewFaker()
		temporaryDb = rawdb.NewMemoryDatabase()
	)
	genesisBlock := gspec.MustCommit(temporaryDb)
	minedBlocks, _ := GenerateChain(gspec.Config, genesisBlock, engine, temporaryDb, 1, func(i int, b *BlockGen) {
		b.SetCoinbase(common.Address{1})
		var signer types.Signer = types.HomesteadSigner{}
		for _, mockTxData := range mockTxDataArr {
			var data []byte
			switch mockTxData.txType {
			case Public:
				data = contractCreateABIPayloadBytes
			case Private, PMT:
				data = encryptedPayloadHashForContractDeployment.Bytes()
			}

			createTransaction := types.NewContractCreation(0, common.Big0, uint64(3000000), common.Big0, data)

			switch mockTxData.txType {
			case Private, PMT:
				createTransaction.SetPrivate()
				signer = types.QuorumPrivateTxSigner{}
			}
			signedTx, _ := types.SignTx(createTransaction, signer, mockTxData.fromPrivateKey)
			b.AddTxWithChain(chain, signedTx)

			// save the contract address to use when calling `set()`
			mockTxData.toAddress = b.receipts[0].ContractAddress
		}
	})
	futureBlocks, _ := GenerateChain(gspec.Config, minedBlocks[0], engine, temporaryDb, 1, func(i int, b *BlockGen) {
		b.SetCoinbase(common.Address{1})
		var signer types.Signer = types.HomesteadSigner{}
		for _, mockTxData := range mockTxDataArr {
			var data []byte
			var setTransaction *types.Transaction

			switch mockTxData.txType {
			case Public:
				data = contractSetABIPayloadBytes
				setTransaction = types.NewTransaction(1, mockTxData.toAddress, common.Big0, uint64(3000000), common.Big0, data)
			case Private:
				data = encryptedPayloadHashForSetFunction.Bytes()
				setTransaction = types.NewTransaction(1, mockTxData.toAddress, common.Big0, uint64(3000000), common.Big0, data)
				setTransaction.SetPrivate()
				signer = types.QuorumPrivateTxSigner{}
			case PMT:
				data = common.LeftPadBytes(mockTxData.toAddress.Bytes(), 64)
				setTransaction = types.NewTransaction(1, common.QuorumPrivacyPrecompileContractAddress(), common.Big0, uint64(3000000), common.Big0, data)
			}

			if decorateSetTransaction != nil {
				decorateSetTransaction(setTransaction, mockTxData)
			}

			signedTx, _ := types.SignTx(setTransaction, signer, mockTxData.fromPrivateKey)
			b.AddTxWithChain(chain, signedTx)
		}
	})

	return minedBlocks[0], futureBlocks[0]
}

const (
	contractDeployedDefinition = `
[
	{
		"constant": false,
		"inputs": [
			{
				"name": "newValue",
				"type": "uint256"
			}
		],
		"name": "set",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "get",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "newValue",
				"type": "uint256"
			}
		],
		"name": "newContractC2",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"name": "initVal",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	}
]
`
)
