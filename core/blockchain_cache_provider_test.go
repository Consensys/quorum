package core

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/private"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	incrementByOne = common.BytesToEncryptedPayloadHash([]byte("incContract"))
)

func buildCacheProviderTestChain(n int, config *params.ChainConfig, quorumChainConfig *QuorumChainConfig) ([]*types.Block, map[common.Hash]*types.Block, *BlockChain) {
	testdb := rawdb.NewMemoryDatabase()
	genesis := GenesisBlockForTesting(testdb, testAddress, big.NewInt(1000000000))

	// The generated chain deploys one Accumulator contracts which is incremented every block
	blocks, _ := GenerateChain(config, genesis, ethash.NewFaker(), testdb, n, func(i int, block *BlockGen) {
		block.SetCoinbase(common.Address{0})

		signer := types.QuorumPrivateTxSigner{}
		var tx *types.Transaction
		var err error
		if i == 0 {
			tx, err = types.SignTx(types.NewContractCreation(block.TxNonce(testAddress), big.NewInt(0), testGas, nil, deployContract.Bytes()), signer, testKey)
		} else {
			tx, err = types.SignTx(types.NewTransaction(block.TxNonce(testAddress), Contract1AddressAfterDeployment, big.NewInt(0), testGas, nil, incrementByOne.Bytes()), signer, testKey)
		}
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

	// recreate the DB so that we don't have the public state written already by the block generation logic
	testdb = rawdb.NewMemoryDatabase()
	genesis = GenesisBlockForTesting(testdb, testAddress, big.NewInt(1000000000))

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

func buildMockPTM(mockCtrl *gomock.Controller) private.PrivateTransactionManager {
	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)
	deployAccumulatorContractConstructor, _ := AccumulatorParsedABI.Pack("", big.NewInt(1))
	deployAccumulatorContract := append(common.FromHex(AccumulatorBin), deployAccumulatorContractConstructor...)
	incrementAccumulatorContract, _ := AccumulatorParsedABI.Pack("inc", big.NewInt(1))

	mockptm.EXPECT().Receive(deployContract).Return("", []string{"AAA"}, deployAccumulatorContract, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(incrementByOne).Return("", []string{"AAA"}, incrementAccumulatorContract, nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(common.EncryptedPayloadHash{}).Return("", []string{}, common.EncryptedPayloadHash{}.Bytes(), nil, nil).AnyTimes()

	return mockptm
}

func TestSegregatedCacheProvider(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockptm := buildMockPTM(mockCtrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	blocks, _, blockchain := buildCacheProviderTestChain(11, params.QuorumTestChainConfig, nil)

	count, err := blockchain.InsertChain(blocks)

	assert.Nil(t, err)
	assert.Equal(t, len(blocks), count)

	lastBlock := blocks[len(blocks)-1]

	statedbLast, privateStateRepoLast, _ := blockchain.StateAt(lastBlock.Root())

	assert.Equal(t, uint64(len(blocks)), statedbLast.GetNonce(testAddress))
	privateState, _ := privateStateRepoLast.DefaultState()
	accPrivateStateStateLast := privateState.GetState(Contract1AddressAfterDeployment, common.Hash{})
	assert.Equal(t, common.BytesToHash(big.NewInt(int64(len(blocks))).Bytes()), accPrivateStateStateLast)

	// retrieve the state at block height 1
	block1 := blocks[1]

	statedbB1, privateStateRepoB1, _ := blockchain.StateAt(block1.Root())

	assert.Equal(t, uint64(2), statedbB1.GetNonce(testAddress))
	privateState, _ = privateStateRepoB1.DefaultState()
	privateStateRoot := privateState.IntermediateRoot(false)
	accPrivateStateStateB1 := privateState.GetState(Contract1AddressAfterDeployment, common.Hash{})
	assert.Equal(t, common.BytesToHash([]byte{2}), accPrivateStateStateB1)

	// check that the private state root has already been written to the DB
	contains, err := blockchain.db.Has(privateStateRoot.Bytes())
	assert.Nil(t, err)
	assert.True(t, contains)
}

func TestUnifiedCacheProvider(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockptm := buildMockPTM(mockCtrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	blocks, _, blockchain := buildCacheProviderTestChain(130, params.QuorumTestChainConfig, &QuorumChainConfig{privateTrieCacheEnabled: true})

	count, err := blockchain.InsertChain(blocks[:129])

	assert.Nil(t, err)
	assert.Equal(t, 129, count)

	lastBlock := blocks[128]

	statedb, privateStateRepo, _ := blockchain.StateAt(lastBlock.Root())

	assert.Equal(t, uint64(129), statedb.GetNonce(testAddress))
	privateState, _ := privateStateRepo.DefaultState()
	accPrivateState := privateState.GetState(Contract1AddressAfterDeployment, common.Hash{})
	assert.Equal(t, common.BytesToHash(big.NewInt(129).Bytes()), accPrivateState)

	// The following is an attempt to explain the process by which the trie nodes corresponding to privateState are being
	// garbage collected due to the TriesInMemory limit implemented in the blockchain

	// Expected state structure (block 2 - index 1 in the blocks array)
	// Public state just contains the testAddress(testKey) with nonce 2 and has PUB(BL2) root hash
	// privateState has C1(2) privateState(BL2) as root hash
	// the public state root PUB(BL2) references the private state root privateState(BL2)

	// Considering the above we can establish that
	// privateState(BL2) is referenced once by the public state root PUB(BL2) at block height 2

	// retrieve the state at block height 2
	block1 := blocks[1]
	statedbB1, privateStateRepoB1, _ := blockchain.StateAt(block1.Root())

	assert.Equal(t, uint64(2), statedbB1.GetNonce(testAddress))
	privateState, _ = privateStateRepoB1.DefaultState()
	privateStateRoot := privateState.IntermediateRoot(false)
	accPrivateStateB1 := privateState.GetState(Contract1AddressAfterDeployment, common.Hash{})
	assert.Equal(t, common.BytesToHash([]byte{2}), accPrivateStateB1)

	// check that the roots have NOT been written to the underlying DB
	contains, err := blockchain.db.Has(block1.Root().Bytes())
	assert.Nil(t, err)
	assert.False(t, contains)
	contains, err = blockchain.db.Has(privateStateRoot.Bytes())
	assert.Nil(t, err)
	assert.False(t, contains)

	// check that the roots are available in the cache
	data, err := blockchain.stateCache.TrieDB().Node(block1.Root())
	assert.Nil(t, err)
	assert.True(t, len(data) > 0)
	data, err = blockchain.stateCache.TrieDB().Node(privateStateRoot)
	assert.Nil(t, err)
	assert.True(t, len(data) > 0)

	// Process block 130 and reassess the underlying DB and the cache
	// When block 130 is processed the "chosen" in blockchain becomes 2 (130 - TriesInMemory) and the public root hash
	// PUB(BL2) is being de-referenced (reference count is reduced by 1) and the reference count (parents) becomes 0
	// As a result private state root privateState(BL2) is de-referenced and the reference count becomes 0

	// All nodes with reference counts (parents) equal to 0 are being garbage collected (removed from the cache)
	count, err = blockchain.InsertChain(blocks[129:])

	assert.Nil(t, err)
	assert.Equal(t, 1, count)

	// check that the roots have NOT been written to the underlying DB
	contains, err = blockchain.db.Has(block1.Root().Bytes())
	assert.Nil(t, err)
	assert.False(t, contains)
	contains, err = blockchain.db.Has(privateStateRoot.Bytes())
	assert.Nil(t, err)
	assert.False(t, contains)

	// check that the roots have been garbage collected (removed) from the cache (other intermediate trie nodes may have
	// been eliminated from the cache as well)
	data, err = blockchain.stateCache.TrieDB().Node(block1.Root())
	assert.Error(t, err, "not found")
	assert.Nil(t, data)
	data, err = blockchain.stateCache.TrieDB().Node(privateStateRoot)
	assert.Error(t, err, "not found")
	assert.Nil(t, data)
}
