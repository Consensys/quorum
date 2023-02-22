package core

import (
	"encoding/base64"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

/*
pragma solidity ^0.5.0;

contract Accumulator {
  uint public storedData;

  event IncEvent(uint value);

  constructor(uint initVal) public{
    storedData = initVal;
  }

  function inc(uint x) public {
    storedData = storedData + x;
    emit IncEvent(storedData);
  }

  function get() view public returns (uint retVal) {
    return storedData;
  }
}
*/

const AccumulatorABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initVal\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"IncEvent\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"retVal\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"inc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"storedData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

var (
	deployContract       = common.BytesToEncryptedPayloadHash([]byte("deployContract"))
	incrementByOnePS1    = common.BytesToEncryptedPayloadHash([]byte("incContractPS1"))
	incrementByOnePS1PS2 = common.BytesToEncryptedPayloadHash([]byte("incContractPS1PS2"))

	AccumulatorParsedABI, _         = abi.JSON(strings.NewReader(AccumulatorABI))
	AccumulatorBin                  = "0x608060405234801561001057600080fd5b5060405161018a38038061018a8339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000819055505061012f8061005b6000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80632a1afcd91460415780636d4ce63c14605d578063812600df146079575b600080fd5b604760a4565b6040518082815260200191505060405180910390f35b606360aa565b6040518082815260200191505060405180910390f35b60a260048036036020811015608d57600080fd5b810190808035906020019092919050505060b3565b005b60005481565b60008054905090565b80600054016000819055507fc13aa85405f3616d514cfd2316b12181b047ed7f229bce08ce53c671f6f94f986000546040518082815260200191505060405180910390a15056fea265627a7a723158208fb1390ecdc6d669bf1855aed67a225931aee2c14ac2b8f5cd2ac5a8fb3a21af64736f6c63430005110032"
	Contract1AddressAfterDeployment = crypto.CreateAddress(testAddress, 0)
	Contract2AddressAfterDeployment = crypto.CreateAddress(testAddress, 1)
	PS1PG                           = engine.PrivacyGroup{
		Type:           "RESIDENT",
		Name:           "PS1",
		PrivacyGroupId: base64.StdEncoding.EncodeToString([]byte("PS1")),
		Description:    "Resident Group 1",
		From:           "",
		Members:        []string{"AAA", "BBB"},
	}

	PS2PG = engine.PrivacyGroup{
		Type:           "RESIDENT",
		Name:           "PS2",
		PrivacyGroupId: base64.StdEncoding.EncodeToString([]byte("PS2")),
		Description:    "Resident Group 2",
		From:           "",
		Members:        []string{"CCC", "DDD"},
	}
)

func buildCacheProviderMPSTestChain(n int, config *params.ChainConfig, quorumChainConfig *QuorumChainConfig) ([]*types.Block, map[common.Hash]*types.Block, *BlockChain) {
	testdb := rawdb.NewMemoryDatabase()
	genesis := GenesisBlockForTesting(testdb, testAddress, big.NewInt(1000000000))

	// The generated chain deploys two Accumulator contracts
	// - Accumulator contract 1 is incremented every 1 and 2 blocks for PS1 and PS1&PS2 respectively
	// - Accumulator contract 2 is incremented every block for both PS1 and PS2
	blocks, _ := GenerateChain(config, genesis, ethash.NewFaker(), testdb, n, func(i int, block *BlockGen) {
		block.SetCoinbase(common.Address{0})

		signer := types.QuorumPrivateTxSigner{}
		var tx *types.Transaction
		var err error
		if i == 0 {
			tx, err = types.SignTx(types.NewContractCreation(block.TxNonce(testAddress), big.NewInt(0), testGas, nil, deployContract.Bytes()), signer, testKey)
			if err != nil {
				panic(err)
			}
			block.AddTx(tx)
			tx, err = types.SignTx(types.NewContractCreation(block.TxNonce(testAddress), big.NewInt(0), testGas, nil, deployContract.Bytes()), signer, testKey)
			if err != nil {
				panic(err)
			}
			block.AddTx(tx)
		} else {
			if i%2 == 1 {
				tx, err = types.SignTx(types.NewTransaction(block.TxNonce(testAddress), Contract1AddressAfterDeployment, big.NewInt(0), testGas, nil, incrementByOnePS1.Bytes()), signer, testKey)
				if err != nil {
					panic(err)
				}
				block.AddTx(tx)
				tx, err = types.SignTx(types.NewTransaction(block.TxNonce(testAddress), Contract2AddressAfterDeployment, big.NewInt(0), testGas, nil, incrementByOnePS1PS2.Bytes()), signer, testKey)
				if err != nil {
					panic(err)
				}
				block.AddTx(tx)
			} else {
				tx, err = types.SignTx(types.NewTransaction(block.TxNonce(testAddress), Contract1AddressAfterDeployment, big.NewInt(0), testGas, nil, incrementByOnePS1PS2.Bytes()), signer, testKey)
				if err != nil {
					panic(err)
				}
				block.AddTx(tx)
				tx, err = types.SignTx(types.NewTransaction(block.TxNonce(testAddress), Contract2AddressAfterDeployment, big.NewInt(0), testGas, nil, incrementByOnePS1PS2.Bytes()), signer, testKey)
				if err != nil {
					panic(err)
				}
				block.AddTx(tx)
			}
		}
	})

	hashes := make([]common.Hash, n+1)
	hashes[len(hashes)-1] = genesis.Hash()
	blockm := make(map[common.Hash]*types.Block, n+1)
	blockm[genesis.Hash()] = genesis
	for i, b := range blocks {
		hashes[len(hashes)-i-2] = b.Hash()
		blockm[b.Hash()] = b
	}

	// recreate the DB so that we don't have the public state written already by the block generation logic
	testdb = rawdb.NewMemoryDatabase()
	_ = GenesisBlockForTesting(testdb, testAddress, big.NewInt(1000000000))

	// disable snapshots
	testingCacheConfig := &CacheConfig{
		TrieCleanLimit: 256,
		TrieDirtyLimit: 256,
		TrieTimeLimit:  5 * time.Minute,
		// TODO - figure out why is the snapshot causing panics when enabled during the test
		SnapshotLimit: 0,
		SnapshotWait:  true,
	}

	blockchain, err := NewBlockChain(testdb, testingCacheConfig, config, ethash.NewFaker(), vm.Config{}, nil, nil, quorumChainConfig)
	if err != nil {
		return nil, nil, nil
	}
	return blocks, blockm, blockchain
}

func buildMockMPSPTM(mockCtrl *gomock.Controller) private.PrivateTransactionManager {
	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)
	deployAccumulatorContractConstructor, _ := AccumulatorParsedABI.Pack("", big.NewInt(1))
	deployAccumulatorContract := append(common.FromHex(AccumulatorBin), deployAccumulatorContractConstructor...)
	incrementAccumulatorContract, _ := AccumulatorParsedABI.Pack("inc", big.NewInt(1))

	mockptm.EXPECT().Receive(deployContract).Return("", []string{"AAA", "CCC"}, deployAccumulatorContract, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(incrementByOnePS1).Return("", []string{"AAA"}, incrementAccumulatorContract, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(incrementByOnePS1PS2).Return("", []string{"AAA", "CCC"}, incrementAccumulatorContract, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(common.EncryptedPayloadHash{}).Return("", []string{}, common.EncryptedPayloadHash{}.Bytes(), nil, nil).AnyTimes()
	mockptm.EXPECT().HasFeature(engine.MultiplePrivateStates).Return(true)
	mockptm.EXPECT().Groups().Return([]engine.PrivacyGroup{PS1PG, PS2PG}, nil).AnyTimes()

	return mockptm
}

func TestSegregatedCacheProviderMPS(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockptm := buildMockMPSPTM(mockCtrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	blocks, _, blockchain := buildCacheProviderMPSTestChain(11, params.QuorumMPSTestChainConfig, nil)

	count, err := blockchain.InsertChain(blocks)

	assert.Nil(t, err)
	assert.Equal(t, len(blocks), count)

	lastBlock := blocks[len(blocks)-1]

	statedbLast, privateStateRepoLast, _ := blockchain.StateAt(lastBlock.Root())

	assert.Equal(t, uint64(2*len(blocks)), statedbLast.GetNonce(testAddress))
	PS1, _ := privateStateRepoLast.StatePSI(types.PrivateStateIdentifier("PS1"))
	accPS1StateLast := PS1.GetState(Contract1AddressAfterDeployment, common.Hash{})
	assert.Equal(t, common.BytesToHash(big.NewInt(int64(len(blocks))).Bytes()), accPS1StateLast)

	PS2, _ := privateStateRepoLast.StatePSI(types.PrivateStateIdentifier("PS2"))
	accPS2StateLast := PS2.GetState(Contract1AddressAfterDeployment, common.Hash{})
	// PS2 is incremented every other block - thus the (len(blocks)+1)/2 formula
	assert.Equal(t, common.BytesToHash(big.NewInt(int64((len(blocks)+1)/2)).Bytes()), accPS2StateLast)

	// retrieve the state at block height 1
	block1 := blocks[1]

	statedbB1, privateStateRepoB1, _ := blockchain.StateAt(block1.Root())

	assert.Equal(t, uint64(4), statedbB1.GetNonce(testAddress))
	PS1, _ = privateStateRepoB1.StatePSI(types.PrivateStateIdentifier("PS1"))
	PS1Root := PS1.IntermediateRoot(false)
	accPS1StateB1 := PS1.GetState(Contract1AddressAfterDeployment, common.Hash{})
	assert.Equal(t, common.BytesToHash([]byte{2}), accPS1StateB1)

	PS2, _ = privateStateRepoB1.StatePSI(types.PrivateStateIdentifier("PS2"))
	PS2Root := PS2.IntermediateRoot(false)
	accPS2StateB1 := PS2.GetState(Contract1AddressAfterDeployment, common.Hash{})
	// PS2 is incremented every other block - thus the (len(blocks)+1)/2 formula
	assert.Equal(t, common.BytesToHash([]byte{1}), accPS2StateB1)

	// check that both roots have already been written to the underlying DB

	contains, err := blockchain.db.Has(PS1Root.Bytes())
	assert.Nil(t, err)
	assert.True(t, contains)
	contains, err = blockchain.db.Has(PS2Root.Bytes())
	assert.Nil(t, err)
	assert.True(t, contains)
}

func TestUnifiedCacheProviderMPS(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockptm := buildMockMPSPTM(mockCtrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	blocks, _, blockchain := buildCacheProviderMPSTestChain(130, params.QuorumMPSTestChainConfig, &QuorumChainConfig{multiTenantEnabled: true, privateTrieCacheEnabled: true})

	count, err := blockchain.InsertChain(blocks[:129])

	assert.Nil(t, err)
	assert.Equal(t, 129, count)

	lastBlock := blocks[128]

	statedb, privateStateRepo, _ := blockchain.StateAt(lastBlock.Root())

	assert.Equal(t, uint64(258), statedb.GetNonce(testAddress))
	PS1, _ := privateStateRepo.StatePSI(types.PrivateStateIdentifier("PS1"))
	accPS1State := PS1.GetState(Contract1AddressAfterDeployment, common.Hash{})
	assert.Equal(t, common.BytesToHash(big.NewInt(129).Bytes()), accPS1State)

	PS2, _ := privateStateRepo.StatePSI(types.PrivateStateIdentifier("PS2"))
	accPS2State := PS2.GetState(Contract1AddressAfterDeployment, common.Hash{})
	// PS2 is incremented every other block - thus the (len(blocks)+1)/2 formula
	assert.Equal(t, common.BytesToHash(big.NewInt(65).Bytes()), accPS2State)

	// The following is an attempt to explain the process by which the trie nodes corresponding to PS1 and PS2 are being
	// garbage collected due to the TriesInMemory limit implemented in the blockchain

	// Expected state structure (block 2 - index 1 in the blocks array)
	// Public state just contains the testAddress(testKey) with nonce 4 and has PUB(BL2) root hash
	// PS1 has C1(2) and C2(2) with PS1(BL2) as root hash
	// PS2 has C1(1) and C2(2) with PS2(BL2) as root hash
	// the Trie of private states contains PS1(BL2) and PS2(BL2) and has TPS(BL2) root hash
	// the public state root references the trie of private states PUB(BL2) -> TPS(BL2)
	// the trie of private states leaves reference PS1(BL2) and PS2(BL2)

	// Considering the above we can establish that
	// PS1(BL2) is referenced once by the PS1 leaf in the trie of private states at block height 2
	// PS2(BL2) is referenced once by the PS2 leaf in the trie of private states at block height 2

	// retrieve the state at block height 2
	block1 := blocks[1]
	statedbB1, privateStateRepoB1, _ := blockchain.StateAt(block1.Root())

	assert.Equal(t, uint64(4), statedbB1.GetNonce(testAddress))
	PS1, _ = privateStateRepoB1.StatePSI(types.PrivateStateIdentifier("PS1"))
	PS1Root := PS1.IntermediateRoot(false)
	accPS1StateB1 := PS1.GetState(Contract1AddressAfterDeployment, common.Hash{})
	assert.Equal(t, common.BytesToHash([]byte{2}), accPS1StateB1)

	PS2, _ = privateStateRepoB1.StatePSI(types.PrivateStateIdentifier("PS2"))
	PS2Root := PS2.IntermediateRoot(false)
	accPS2StateB1 := PS2.GetState(Contract1AddressAfterDeployment, common.Hash{})
	// PS2 is incremented every other block - thus the (len(blocks)+1)/2 formula
	assert.Equal(t, common.BytesToHash([]byte{1}), accPS2StateB1)

	// check that the roots have NOT been written to the underlying DB
	contains, err := blockchain.db.Has(block1.Root().Bytes())
	assert.Nil(t, err)
	assert.False(t, contains)
	contains, err = blockchain.db.Has(PS1Root.Bytes())
	assert.Nil(t, err)
	assert.False(t, contains)
	contains, err = blockchain.db.Has(PS2Root.Bytes())
	assert.Nil(t, err)
	assert.False(t, contains)

	// check that the roots are available in the cache
	data, err := blockchain.stateCache.TrieDB().Node(block1.Root())
	assert.Nil(t, err)
	assert.True(t, len(data) > 0)
	data, err = blockchain.stateCache.TrieDB().Node(PS1Root)
	assert.Nil(t, err)
	assert.True(t, len(data) > 0)
	data, err = blockchain.stateCache.TrieDB().Node(PS2Root)
	assert.Nil(t, err)
	assert.True(t, len(data) > 0)

	// Process block 130 and reassess the underlying DB and the cache
	// When block 130 is processed the "chosen" in blockchain becomes 2 (130 - TriesInMemory) and the public root hash
	// PUB(BL2) is being de-referenced (reference count is reduced by 1) and the reference count (parents) becomes 0
	// As a result the trie of private states TPS(BL2) is also being de-referenced and becomes 0
	// Each leaf in the trie of private states is being de-referenced which then causes the private state roots PS1(BL2)
	// and PS2(BL2) to at block height 2 be de-referenced as well

	// All nodes with reference counts (parents) equal to 0 are being garbage collected (removed from the cache)
	count, err = blockchain.InsertChain(blocks[129:])

	assert.Nil(t, err)
	assert.Equal(t, 1, count)

	// check that the roots have NOT been written to the underlying DB
	contains, err = blockchain.db.Has(block1.Root().Bytes())
	assert.Nil(t, err)
	assert.False(t, contains)
	contains, err = blockchain.db.Has(PS1Root.Bytes())
	assert.Nil(t, err)
	assert.False(t, contains)
	contains, err = blockchain.db.Has(PS2Root.Bytes())
	assert.Nil(t, err)
	assert.False(t, contains)

	// check that the roots have been garbage collected (removed) from the cache (other intermediate trie nodes may have
	// been eliminated from the cache as well)
	data, err = blockchain.stateCache.TrieDB().Node(block1.Root())
	assert.Error(t, err, "not found")
	assert.Nil(t, data)
	data, err = blockchain.stateCache.TrieDB().Node(PS1Root)
	assert.Error(t, err, "not found")
	assert.Nil(t, data)
	data, err = blockchain.stateCache.TrieDB().Node(PS2Root)
	assert.Error(t, err, "not found")
	assert.Nil(t, data)
}
