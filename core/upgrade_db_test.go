package core

import (
	"encoding/base64"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	PrivatePG = engine.PrivacyGroup{
		Type:           "RESIDENT",
		Name:           "private",
		PrivacyGroupId: base64.StdEncoding.EncodeToString([]byte("private")),
		Description:    "private",
		From:           "",
		Members:        []string{"CCC", "DDD"},
	}
	DBUpgradeQuorumTestChainConfig = &params.ChainConfig{
		ChainID:                  params.QuorumTestChainConfig.ChainID,
		HomesteadBlock:           params.QuorumTestChainConfig.HomesteadBlock,
		DAOForkBlock:             params.QuorumTestChainConfig.DAOForkBlock,
		DAOForkSupport:           params.QuorumTestChainConfig.DAOForkSupport,
		EIP150Block:              params.QuorumTestChainConfig.EIP150Block,
		EIP150Hash:               params.QuorumTestChainConfig.EIP150Hash,
		EIP155Block:              params.QuorumTestChainConfig.EIP155Block,
		EIP158Block:              params.QuorumTestChainConfig.EIP158Block,
		ByzantiumBlock:           params.QuorumTestChainConfig.ByzantiumBlock,
		ConstantinopleBlock:      params.QuorumTestChainConfig.ConstantinopleBlock,
		PetersburgBlock:          params.QuorumTestChainConfig.PetersburgBlock,
		IstanbulBlock:            params.QuorumTestChainConfig.IstanbulBlock,
		MuirGlacierBlock:         params.QuorumTestChainConfig.MuirGlacierBlock,
		YoloV2Block:              params.QuorumTestChainConfig.YoloV2Block,
		EWASMBlock:               params.QuorumTestChainConfig.EWASMBlock,
		Ethash:                   params.QuorumTestChainConfig.Ethash,
		Clique:                   params.QuorumTestChainConfig.Clique,
		Istanbul:                 params.QuorumTestChainConfig.Istanbul,
		IsQuorum:                 params.QuorumTestChainConfig.IsQuorum,
		TransactionSizeLimit:     params.QuorumTestChainConfig.TransactionSizeLimit,
		MaxCodeSize:              params.QuorumTestChainConfig.MaxCodeSize,
		QIP714Block:              params.QuorumTestChainConfig.QIP714Block,
		MaxCodeSizeChangeBlock:   params.QuorumTestChainConfig.MaxCodeSizeChangeBlock,
		MaxCodeSizeConfig:        params.QuorumTestChainConfig.MaxCodeSizeConfig,
		PrivacyEnhancementsBlock: params.QuorumTestChainConfig.PrivacyEnhancementsBlock,
		IsMPS:                    params.QuorumTestChainConfig.IsMPS,
	}
)

// 1. Start the chain with isMPS=false and insert 3 blocks that each create a private contract.
// 2. Iterate over the private receipts and extract each contract address. Verify that the contracts have non empty bytecode.
// 3. Run mpsdbupbrade and start the chain with isMPS=true.
// 4. Insert an extra block which adds a new private contract. Verify that all the contracts identified at step 2 are
// still available in the "private" state and that all of the contracts are available as empty contracts in the empty
// state.
func TestMultiplePSMRDBUpgrade(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockptm.EXPECT().Receive(gomock.Not(common.EncryptedPayloadHash{})).Return("", []string{"CCC"}, common.FromHex(testCode), nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(common.EncryptedPayloadHash{}).Return("", []string{}, common.EncryptedPayloadHash{}.Bytes(), nil, nil).AnyTimes()
	mockptm.EXPECT().HasFeature(engine.MultiplePrivateStates).Return(true)
	mockptm.EXPECT().Groups().Return([]engine.PrivacyGroup{PrivatePG}, nil).AnyTimes()

	blocks, _, blockchain := buildTestChain(4, DBUpgradeQuorumTestChainConfig)
	db := blockchain.db

	count, err := blockchain.InsertChain(blocks[0:3])
	assert.NoError(t, err)
	assert.Equal(t, 3, count)

	c1Address := blockchain.GetReceiptsByHash(blocks[0].Hash())[0].ContractAddress
	assert.NotNil(t, c1Address)
	c2Address := blockchain.GetReceiptsByHash(blocks[1].Hash())[0].ContractAddress
	assert.NotNil(t, c2Address)
	c3Address := blockchain.GetReceiptsByHash(blocks[2].Hash())[0].ContractAddress
	assert.NotNil(t, c3Address)
	// check that the C3 receipt is a flat receipt (PSReceipts field is nil)
	c3Receipt := blockchain.GetReceiptsByHash(blocks[2].Hash())[0]
	assert.Empty(t, c3Receipt.PSReceipts)

	standaloneStateRepo, err := blockchain.PrivateStateManager().StateRepository(blocks[2].Root())
	assert.NoError(t, err)

	standaloneStateDB, err := standaloneStateRepo.DefaultState()
	assert.NoError(t, err)

	assert.True(t, standaloneStateDB.Exist(c1Address))
	assert.NotEqual(t, standaloneStateDB.GetCodeSize(c1Address), 0)
	assert.True(t, standaloneStateDB.Exist(c2Address))
	assert.NotEqual(t, standaloneStateDB.GetCodeSize(c2Address), 0)
	assert.True(t, standaloneStateDB.Exist(c3Address))
	assert.NotEqual(t, standaloneStateDB.GetCodeSize(c3Address), 0)

	// execute mpsdbupgrade
	assert.Nil(t, mps.UpgradeDB(db, blockchain))
	// UpgradeDB updates the chainconfig isMPS to true so set it back to false at the end of the test
	defer func() { DBUpgradeQuorumTestChainConfig.IsMPS = false }()
	assert.True(t, DBUpgradeQuorumTestChainConfig.IsMPS)

	blockchain.Stop()

	// reinstantiate the blockchain with isMPS enabled
	blockchain, err = NewBlockChain(db, nil, DBUpgradeQuorumTestChainConfig, ethash.NewFaker(), vm.Config{}, nil, nil, nil)
	assert.Nil(t, err)

	count, err = blockchain.InsertChain(blocks[3:])
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	c4Address := blockchain.GetReceiptsByHash(blocks[3].Hash())[0].ContractAddress
	assert.NotNil(t, c4Address)

	mpsStateRepo, err := blockchain.PrivateStateManager().StateRepository(blocks[3].Root())
	assert.NoError(t, err)

	emptyStateDB, err := mpsStateRepo.DefaultState()
	assert.NoError(t, err)
	privateStateDB, err := mpsStateRepo.StatePSI(types.DefaultPrivateStateIdentifier)
	assert.NoError(t, err)

	assert.True(t, privateStateDB.Exist(c1Address))
	assert.NotEqual(t, privateStateDB.GetCodeSize(c1Address), 0)
	assert.True(t, privateStateDB.Exist(c2Address))
	assert.NotEqual(t, privateStateDB.GetCodeSize(c2Address), 0)
	assert.True(t, privateStateDB.Exist(c3Address))
	assert.NotEqual(t, privateStateDB.GetCodeSize(c3Address), 0)
	assert.True(t, privateStateDB.Exist(c4Address))
	assert.NotEqual(t, privateStateDB.GetCodeSize(c4Address), 0)

	assert.True(t, emptyStateDB.Exist(c1Address))
	assert.Equal(t, emptyStateDB.GetCodeSize(c1Address), 0)
	assert.True(t, emptyStateDB.Exist(c2Address))
	assert.Equal(t, emptyStateDB.GetCodeSize(c2Address), 0)
	assert.True(t, emptyStateDB.Exist(c3Address))
	assert.Equal(t, emptyStateDB.GetCodeSize(c3Address), 0)
	assert.True(t, emptyStateDB.Exist(c4Address))
	assert.Equal(t, emptyStateDB.GetCodeSize(c4Address), 0)

	// check the receipts has the PSReceipts field populated (due to the newly applied block)
	c4Receipt := blockchain.GetReceiptsByHash(blocks[3].Hash())[0]
	assert.NotNil(t, c4Receipt.PSReceipts)
	assert.Contains(t, c4Receipt.PSReceipts, types.DefaultPrivateStateIdentifier)
	assert.Contains(t, c4Receipt.PSReceipts, types.EmptyPrivateStateIdentifier)

	// check the block 3 receipts has been upgraded and it has PSReceipts field populated (by the upgrade process)
	c3Receipt = blockchain.GetReceiptsByHash(blocks[2].Hash())[0]
	assert.NotNil(t, c3Receipt.PSReceipts)
	assert.Contains(t, c3Receipt.PSReceipts, types.DefaultPrivateStateIdentifier)
	assert.Contains(t, c3Receipt.PSReceipts, types.EmptyPrivateStateIdentifier)
}
