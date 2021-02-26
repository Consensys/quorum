package core

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/private"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	// testCode is the testing contract binary code which will initialises some
	// variables in constructor
	testCode = "0x60806040527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060005534801561003457600080fd5b5060fc806100436000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80630c4dae8814603757806398a213cf146053575b600080fd5b603d607e565b6040518082815260200191505060405180910390f35b607c60048036036020811015606757600080fd5b81019080803590602001909291905050506084565b005b60005481565b806000819055507fe9e44f9f7da8c559de847a3232b57364adc0354f15a2cd8dc636d54396f9587a6000546040518082815260200191505060405180910390a15056fea265627a7a723058208ae31d9424f2d0bc2a3da1a5dd659db2d71ec322a17db8f87e19e209e3a1ff4a64736f6c634300050a0032"

	// testGas is the gas required for contract deployment.
	testGas = 144109
)

var (
	testKey, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testAddress = crypto.PubkeyToAddress(testKey.PublicKey)
)

func buildTestChain(n int, config *params.ChainConfig) ([]*types.Block, map[common.Hash]*types.Block, *BlockChain) {
	testdb := rawdb.NewMemoryDatabase()
	genesis := GenesisBlockForTesting(testdb, testAddress, big.NewInt(1000000000))
	blocks, _ := GenerateChain(params.QuorumMPSTestChainConfig, genesis, ethash.NewFaker(), testdb, n, func(i int, block *BlockGen) {
		block.SetCoinbase(common.Address{0})

		signer := types.QuorumPrivateTxSigner{}
		tx, err := types.SignTx(types.NewContractCreation(block.TxNonce(testAddress), big.NewInt(0), testGas, nil, common.FromHex(testCode)), signer, testKey)
		if err != nil {
			panic(err)
		}
		block.AddTx(tx)
	})

	hashes := make([]common.Hash, n+1)
	hashes[len(hashes)-1] = genesis.Hash()
	blockm := make(map[common.Hash]*types.Block, n+1)
	blockm[genesis.Hash()] = genesis
	for i, b := range blocks {
		hashes[len(hashes)-i-2] = b.Hash()
		blockm[b.Hash()] = b
	}

	blockchain, _ := NewBlockChain(testdb, nil, config, ethash.NewFaker(), vm.Config{}, nil, nil)
	return blocks, blockm, blockchain
}

func TestMultiplePSMRStateCreated(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockpsm := mps.NewMockPrivateStateManager(mockCtrl)

	mockptm.EXPECT().Receive(gomock.Not(common.EncryptedPayloadHash{})).Return("", []string{"psi1", "psi2"}, common.FromHex(testCode), nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(common.EncryptedPayloadHash{}).Return("", []string{}, common.EncryptedPayloadHash{}.Bytes(), nil, nil).AnyTimes()

	mockpsm.EXPECT().ResolveForManagedParty("psi1").Return(&PSI1PSM, nil).AnyTimes()
	mockpsm.EXPECT().ResolveForManagedParty("psi2").Return(&PSI2PSM, nil).AnyTimes()

	blocks, blockmap, blockchain := buildTestChain(2, params.QuorumMPSTestChainConfig)
	cache := state.NewDatabase(blockchain.db)
	blockchain.SetPrivateStateManager(mockpsm)

	for _, block := range blocks {
		parent := blockmap[block.ParentHash()]
		statedb, _ := state.New(parent.Root(), blockchain.StateCache(), nil)
		mockpsm.EXPECT().GetPrivateStateRepository(gomock.Any()).Return(mps.NewMultiplePrivateStateRepository(blockchain.chainConfig, blockchain.db, cache, parent.Root())).AnyTimes()

		privateStateRepo, err := blockchain.PrivateStateManager().GetPrivateStateRepository(parent.Root())
		assert.NoError(t, err)

		_, privateReceipts, _, _, _ := blockchain.Processor().Process(block, statedb, privateStateRepo, vm.Config{})

		for _, privateReceipt := range privateReceipts {
			expectedContractAddress := privateReceipt.ContractAddress

			emptyState, _ := privateStateRepo.GetDefaultState()
			assert.True(t, emptyState.Exist(expectedContractAddress))
			assert.Equal(t, emptyState.GetCodeSize(expectedContractAddress), 0)
			ps1, _ := privateStateRepo.GetPrivateState(types.PrivateStateIdentifier("psi1"))
			assert.True(t, ps1.Exist(expectedContractAddress))
			assert.NotEqual(t, ps1.GetCodeSize(expectedContractAddress), 0)
			ps2, _ := privateStateRepo.GetPrivateState(types.PrivateStateIdentifier("psi2"))
			assert.True(t, ps2.Exist(expectedContractAddress))
			assert.NotEqual(t, ps2.GetCodeSize(expectedContractAddress), 0)

		}
		//CommitAndWrite to db
		privateStateRepo.CommitAndWrite(block)

		for _, privateReceipt := range privateReceipts {
			expectedContractAddress := privateReceipt.ContractAddress
			latestBlockRoot := block.Root()
			_, privDb, _ := blockchain.StateAtPSI(latestBlockRoot, types.ToPrivateStateIdentifier("empty"))
			assert.True(t, privDb.Exist(expectedContractAddress))
			assert.Equal(t, privDb.GetCodeSize(expectedContractAddress), 0)
			//contract exists on both psi states
			_, privDb, _ = blockchain.StateAtPSI(latestBlockRoot, types.PrivateStateIdentifier("psi1"))
			assert.True(t, privDb.Exist(expectedContractAddress))
			assert.NotEqual(t, privDb.GetCodeSize(expectedContractAddress), 0)
			_, privDb, _ = blockchain.StateAtPSI(latestBlockRoot, types.PrivateStateIdentifier("psi2"))
			assert.True(t, privDb.Exist(expectedContractAddress))
			assert.NotEqual(t, privDb.GetCodeSize(expectedContractAddress), 0)
			//contract should exist on default private state (delegated to emptystate) but no contract code
			_, privDb, _ = blockchain.StateAtPSI(latestBlockRoot, types.DefaultPrivateStateIdentifier)
			assert.True(t, privDb.Exist(expectedContractAddress))
			assert.Equal(t, privDb.GetCodeSize(expectedContractAddress), 0)
			//contract should exist on random state (delegated to emptystate) but no contract code
			_, privDb, _ = blockchain.StateAtPSI(latestBlockRoot, types.ToPrivateStateIdentifier("other"))
			assert.True(t, privDb.Exist(expectedContractAddress))
			assert.Equal(t, privDb.GetCodeSize(expectedContractAddress), 0)
		}
	}
}

func TestMPSReset(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockpsm := mps.NewMockPrivateStateManager(mockCtrl)

	mockptm.EXPECT().Receive(gomock.Not(common.EncryptedPayloadHash{})).Return("", []string{"psi1", "psi2"}, common.FromHex(testCode), nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(common.EncryptedPayloadHash{}).Return("", []string{}, common.EncryptedPayloadHash{}.Bytes(), nil, nil).AnyTimes()

	mockpsm.EXPECT().ResolveForManagedParty("psi1").Return(&PSI1PSM, nil).AnyTimes()
	mockpsm.EXPECT().ResolveForManagedParty("psi2").Return(&PSI2PSM, nil).AnyTimes()

	blocks, blockmap, blockchain := buildTestChain(2, params.QuorumMPSTestChainConfig)
	cache := state.NewDatabase(blockchain.db)
	blockchain.SetPrivateStateManager(mockpsm)

	for _, block := range blocks {
		parent := blockmap[block.ParentHash()]
		statedb, _ := state.New(parent.Root(), blockchain.StateCache(), nil)
		mockpsm.EXPECT().GetPrivateStateRepository(gomock.Any()).Return(mps.NewMultiplePrivateStateRepository(blockchain.chainConfig, blockchain.db, cache, parent.Root())).AnyTimes()

		privateStateRepo, err := blockchain.PrivateStateManager().GetPrivateStateRepository(parent.Root())
		assert.NoError(t, err)

		_, privateReceipts, _, _, _ := blockchain.Processor().Process(block, statedb, privateStateRepo, vm.Config{})

		for _, privateReceipt := range privateReceipts {
			expectedContractAddress := privateReceipt.ContractAddress

			emptyState, _ := privateStateRepo.GetDefaultState()
			assert.True(t, emptyState.Exist(expectedContractAddress))
			assert.Equal(t, emptyState.GetCodeSize(expectedContractAddress), 0)
			ps1, _ := privateStateRepo.GetPrivateState(types.PrivateStateIdentifier("psi1"))
			assert.True(t, ps1.Exist(expectedContractAddress))
			assert.NotEqual(t, ps1.GetCodeSize(expectedContractAddress), 0)
			ps2, _ := privateStateRepo.GetPrivateState(types.PrivateStateIdentifier("psi2"))
			assert.True(t, ps2.Exist(expectedContractAddress))
			assert.NotEqual(t, ps2.GetCodeSize(expectedContractAddress), 0)

			privateStateRepo.Reset()

			emptyState, _ = privateStateRepo.GetDefaultState()
			assert.False(t, emptyState.Exist(expectedContractAddress))
			assert.Equal(t, emptyState.GetCodeSize(expectedContractAddress), 0)
			ps1, _ = privateStateRepo.GetPrivateState(types.PrivateStateIdentifier("psi1"))
			assert.False(t, ps1.Exist(expectedContractAddress))
			assert.Equal(t, ps1.GetCodeSize(expectedContractAddress), 0)
			ps2, _ = privateStateRepo.GetPrivateState(types.PrivateStateIdentifier("psi2"))
			assert.False(t, ps2.Exist(expectedContractAddress))
			assert.Equal(t, ps2.GetCodeSize(expectedContractAddress), 0)
		}
	}
}

var PSI1PSM = types.PrivateStateMetadata{
	ID:          "psi1",
	Name:        "psi1",
	Description: "private state 1",
	Type:        types.Resident,
	Addresses:   nil,
}

var PSI2PSM = types.PrivateStateMetadata{
	ID:          "psi2",
	Name:        "psi2",
	Description: "private state 2",
	Type:        types.Resident,
	Addresses:   nil,
}
