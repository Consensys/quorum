package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/plugin/security"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/qlight"
	"github.com/golang/mock/gomock"
	"github.com/jpmorganchase/quorum-security-plugin-sdk-go/proto"
	"github.com/stretchr/testify/assert"
)

func TestPrivateBlockDataResolverImpl_PrepareBlockPrivateData_EmptyBlock(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockpsm := mps.NewMockPrivateStateManager(ctrl)
	mockptm := private.NewMockPrivateTransactionManager(ctrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockptm.EXPECT().HasFeature(engine.MultiplePrivateStates).Return(true)
	mockptm.EXPECT().Groups().Return(PrivacyGroups, nil).AnyTimes()

	mockpsm.EXPECT().ResolveForUserContext(gomock.Any()).Return(PSI1PSM, nil).AnyTimes()

	pbdr := qlight.NewPrivateBlockDataResolver(mockpsm, mockptm)
	blocks, _, _ := buildTestChainWithZeroTxPerBlock(1, params.QuorumMPSTestChainConfig)

	blockPrivateData, err := pbdr.PrepareBlockPrivateData(blocks[0], PSI1PSM.ID.String())

	assert.Nil(err)
	assert.Nil(blockPrivateData)
}

func TestPrivateBlockDataResolverImpl_PrepareBlockPrivateData_PartyTransaction(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockpsm := mps.NewMockPrivateStateManager(ctrl)
	mockptm := private.NewMockPrivateTransactionManager(ctrl)
	mockstaterepo := mps.NewMockPrivateStateRepository(ctrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockptm.EXPECT().Receive(gomock.Not(common.EncryptedPayloadHash{})).Return("AAA", []string{"AAA", "CCC"}, common.FromHex(testCode), &engine.ExtraMetadata{
		ACHashes:            nil,
		ACMerkleRoot:        common.Hash{},
		PrivacyFlag:         0,
		ManagedParties:      []string{"AAA", "CCC"},
		Sender:              "AAA",
		MandatoryRecipients: nil,
	}, nil).AnyTimes()
	mockptm.EXPECT().HasFeature(engine.MultiplePrivateStates).Return(true)
	mockptm.EXPECT().Groups().Return(PrivacyGroups, nil).AnyTimes()

	mockpsm.EXPECT().ResolveForUserContext(gomock.Any()).Return(PSI1PSM, nil).AnyTimes()
	mockpsm.EXPECT().NotIncludeAny(gomock.Any(), gomock.Any()).Return(false).AnyTimes()
	mockpsm.EXPECT().StateRepository(gomock.Any()).Return(mockstaterepo, nil).AnyTimes()
	mockpsm.EXPECT().PSIs().Return([]types.PrivateStateIdentifier{PSI1PSM.ID, PSI2PSM.ID, types.DefaultPrivateStateIdentifier, types.ToPrivateStateIdentifier("other")}).AnyTimes()

	mockstaterepo.EXPECT().PrivateStateRoot(gomock.Any()).Return(common.StringToHash("PrivateStateRoot"), nil)

	pbdr := qlight.NewPrivateBlockDataResolver(mockpsm, mockptm)
	blocks, _, _ := buildTestChainWithOneTxPerBlock(1, params.QuorumMPSTestChainConfig)

	blockPrivateData, err := pbdr.PrepareBlockPrivateData(blocks[0], PSI1PSM.ID.String())

	assert.Nil(err)
	assert.NotNil(blockPrivateData)
	assert.Equal(common.StringToHash("PrivateStateRoot"), blockPrivateData.PrivateStateRoot)
	assert.Equal(blocks[0].Hash(), blockPrivateData.BlockHash)
	assert.Len(blockPrivateData.PrivateTransactions, 1)
	privateTransactionData := blockPrivateData.PrivateTransactions[0]
	assert.True(privateTransactionData.IsSender)
	assert.Equal(common.FromHex(testCode), privateTransactionData.Payload)
	assert.ElementsMatch(privateTransactionData.Extra.ManagedParties, []string{"AAA"})
}

func TestPrivateBlockDataResolverImpl_PrepareBlockPrivateData_NonPartyTransaction(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockpsm := mps.NewMockPrivateStateManager(ctrl)
	mockptm := private.NewMockPrivateTransactionManager(ctrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockptm.EXPECT().Receive(gomock.Not(common.EncryptedPayloadHash{})).Return("", nil, nil, nil, nil).AnyTimes()
	mockptm.EXPECT().HasFeature(engine.MultiplePrivateStates).Return(true)
	mockptm.EXPECT().Groups().Return(PrivacyGroups, nil).AnyTimes()

	mockpsm.EXPECT().ResolveForUserContext(gomock.Any()).Return(PSI1PSM, nil).AnyTimes()

	pbdr := qlight.NewPrivateBlockDataResolver(mockpsm, mockptm)
	blocks, _, _ := buildTestChainWithOneTxPerBlock(1, params.QuorumMPSTestChainConfig)

	blockPrivateData, err := pbdr.PrepareBlockPrivateData(blocks[0], PSI1PSM.ID.String())

	assert.Nil(err)
	assert.Nil(blockPrivateData)
}

func TestPrivateBlockDataResolverImpl_PrepareBlockPrivateData_PMTTransaction(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockpsm := mps.NewMockPrivateStateManager(ctrl)
	mockptm := private.NewMockPrivateTransactionManager(ctrl)
	mockstaterepo := mps.NewMockPrivateStateRepository(ctrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	tx, err := types.SignTx(types.NewContractCreation(0, big.NewInt(0), testGas, nil, common.BytesToEncryptedPayloadHash([]byte("pmt private tx")).Bytes()), types.QuorumPrivateTxSigner{}, testKey)
	assert.Nil(err)
	txData := new(bytes.Buffer)
	err = json.NewEncoder(txData).Encode(tx)
	assert.Nil(err)

	mockptm.EXPECT().Receive(common.BytesToEncryptedPayloadHash([]byte("pmt inner tx"))).Return("AAA", []string{"AAA", "CCC"}, txData.Bytes(), &engine.ExtraMetadata{
		ACHashes:            nil,
		ACMerkleRoot:        common.Hash{},
		PrivacyFlag:         0,
		ManagedParties:      []string{"AAA", "CCC"},
		Sender:              "AAA",
		MandatoryRecipients: nil,
	}, nil).AnyTimes()
	mockptm.EXPECT().Receive(common.BytesToEncryptedPayloadHash([]byte("pmt private tx"))).Return("AAA", []string{"AAA", "CCC"}, common.FromHex(testCode), &engine.ExtraMetadata{
		ACHashes:            nil,
		ACMerkleRoot:        common.Hash{},
		PrivacyFlag:         0,
		ManagedParties:      []string{"AAA", "CCC"},
		Sender:              "AAA",
		MandatoryRecipients: nil,
	}, nil).AnyTimes()
	mockptm.EXPECT().HasFeature(engine.MultiplePrivateStates).Return(true)
	mockptm.EXPECT().Groups().Return(PrivacyGroups, nil).AnyTimes()

	mockpsm.EXPECT().ResolveForUserContext(gomock.Any()).Return(PSI1PSM, nil).AnyTimes()
	mockpsm.EXPECT().NotIncludeAny(gomock.Any(), gomock.Any()).Return(false).AnyTimes()
	mockpsm.EXPECT().StateRepository(gomock.Any()).Return(mockstaterepo, nil).AnyTimes()
	mockpsm.EXPECT().PSIs().Return([]types.PrivateStateIdentifier{PSI1PSM.ID, PSI2PSM.ID, types.DefaultPrivateStateIdentifier, types.ToPrivateStateIdentifier("other")}).AnyTimes()

	mockstaterepo.EXPECT().PrivateStateRoot(gomock.Any()).Return(common.StringToHash("PrivateStateRoot"), nil)

	pbdr := qlight.NewPrivateBlockDataResolver(mockpsm, mockptm)
	blocks, _, _ := buildTestChainWithOnePMTTxPerBlock(1, params.QuorumMPSTestChainConfig)

	blockPrivateData, err := pbdr.PrepareBlockPrivateData(blocks[0], PSI1PSM.ID.String())

	assert.Nil(err)
	assert.NotNil(blockPrivateData)
	assert.Equal(common.StringToHash("PrivateStateRoot"), blockPrivateData.PrivateStateRoot)
	assert.Equal(blocks[0].Hash(), blockPrivateData.BlockHash)
	assert.Len(blockPrivateData.PrivateTransactions, 2)

	pmtTransactionData := blockPrivateData.PrivateTransactions[0]
	assert.True(pmtTransactionData.IsSender)
	assert.Equal(txData.Bytes(), pmtTransactionData.Payload)
	assert.ElementsMatch(pmtTransactionData.Extra.ManagedParties, []string{"AAA"})

	privateTransactionData := blockPrivateData.PrivateTransactions[1]
	assert.True(privateTransactionData.IsSender)
	assert.Equal(common.FromHex(testCode), privateTransactionData.Payload)
	assert.ElementsMatch(privateTransactionData.Extra.ManagedParties, []string{"AAA"})
}

func TestAuthProviderImpl_Authorize_AuthManagerNil(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockpsm := mps.NewMockPrivateStateManager(ctrl)
	mockpsm.EXPECT().ResolveForUserContext(gomock.Any()).Return(PSI1PSM, nil).AnyTimes()
	authProvider := qlight.NewAuthProvider(mockpsm, func() security.AuthenticationManager { return nil })

	err := authProvider.Initialize()
	assert.Nil(err)

	err = authProvider.Authorize("token", "psi1")
	assert.Nil(err)
}

func TestAuthProviderImpl_Authorize_AuthManagerDisabled(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockpsm := mps.NewMockPrivateStateManager(ctrl)
	mockpsm.EXPECT().ResolveForUserContext(gomock.Any()).Return(PSI1PSM, nil).AnyTimes()
	authProvider := qlight.NewAuthProvider(mockpsm, func() security.AuthenticationManager {
		return &testAuthManager{
			enabled:   false,
			authError: nil,
			authToken: nil,
		}
	})

	err := authProvider.Initialize()
	assert.Nil(err)

	err = authProvider.Authorize("token", "psi1")
	assert.Nil(err)
}

func TestAuthProviderImpl_Authorize_AuthManagerEnabledAuthError(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockpsm := mps.NewMockPrivateStateManager(ctrl)
	mockpsm.EXPECT().ResolveForUserContext(gomock.Any()).Return(PSI1PSM, nil).AnyTimes()
	authProvider := qlight.NewAuthProvider(mockpsm, func() security.AuthenticationManager {
		return &testAuthManager{
			enabled:   true,
			authError: fmt.Errorf("auth error"),
			authToken: nil,
		}
	})

	err := authProvider.Initialize()
	assert.Nil(err)

	err = authProvider.Authorize("token", "psi1")
	assert.EqualError(err, "auth error")
}

func TestAuthProviderImpl_Authorize_AuthManagerEnabledNotEntitledToPSI(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockpsm := mps.NewMockPrivateStateManager(ctrl)
	mockpsm.EXPECT().ResolveForUserContext(gomock.Any()).Return(PSI1PSM, nil).AnyTimes()
	authProvider := qlight.NewAuthProvider(mockpsm, func() security.AuthenticationManager {
		return &testAuthManager{
			enabled:   true,
			authError: nil,
			authToken: &proto.PreAuthenticatedAuthenticationToken{
				RawToken:  nil,
				ExpiredAt: nil,
				Authorities: []*proto.GrantedAuthority{&proto.GrantedAuthority{
					Service:              "psi",
					Method:               "psi2",
					Raw:                  "psi://psi2",
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				}},
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
		}
	})

	err := authProvider.Initialize()
	assert.Nil(err)

	err = authProvider.Authorize("token", "psi1")
	assert.EqualError(err, "PSI not authorized")
}

func TestAuthProviderImpl_Authorize_AuthManagerEnabledMissingEntitlement(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockpsm := mps.NewMockPrivateStateManager(ctrl)
	mockpsm.EXPECT().ResolveForUserContext(gomock.Any()).Return(PSI1PSM, nil).AnyTimes()
	authProvider := qlight.NewAuthProvider(mockpsm, func() security.AuthenticationManager {
		return &testAuthManager{
			enabled:   true,
			authError: nil,
			authToken: &proto.PreAuthenticatedAuthenticationToken{
				RawToken:  nil,
				ExpiredAt: nil,
				Authorities: []*proto.GrantedAuthority{&proto.GrantedAuthority{
					Service:              "psi",
					Method:               "psi1",
					Raw:                  "psi://psi1",
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				}, &proto.GrantedAuthority{
					Service:              "p2p",
					Method:               "qlight",
					Raw:                  "p2p://qlight",
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				},
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
		}
	})

	err := authProvider.Initialize()
	assert.Nil(err)

	err = authProvider.Authorize("token", "psi1")
	assert.EqualError(err, "The P2P token does not have the necessary authorization p2p=true rpcETH=false")
}

func TestAuthProviderImpl_Authorize_AuthManagerEnabledSuccess(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockpsm := mps.NewMockPrivateStateManager(ctrl)
	mockpsm.EXPECT().ResolveForUserContext(gomock.Any()).Return(PSI1PSM, nil).AnyTimes()
	authProvider := qlight.NewAuthProvider(mockpsm, func() security.AuthenticationManager {
		return &testAuthManager{
			enabled:   true,
			authError: nil,
			authToken: &proto.PreAuthenticatedAuthenticationToken{
				RawToken:  nil,
				ExpiredAt: nil,
				Authorities: []*proto.GrantedAuthority{&proto.GrantedAuthority{
					Service:              "psi",
					Method:               "psi1",
					Raw:                  "psi://psi1",
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				}, &proto.GrantedAuthority{
					Service:              "p2p",
					Method:               "qlight",
					Raw:                  "p2p://qlight",
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				}, &proto.GrantedAuthority{
					Service:              "rpc",
					Method:               "eth_*",
					Raw:                  "rpc://eth_*",
					XXX_NoUnkeyedLiteral: struct{}{},
					XXX_unrecognized:     nil,
					XXX_sizecache:        0,
				},
				},
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
		}
	})

	err := authProvider.Initialize()
	assert.Nil(err)

	err = authProvider.Authorize("token", "psi1")
	assert.Nil(err)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////// Helpers /////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

func buildTestChainWithZeroTxPerBlock(n int, config *params.ChainConfig) ([]*types.Block, map[common.Hash]*types.Block, *core.BlockChain) {
	testdb := rawdb.NewMemoryDatabase()
	genesis := core.GenesisBlockForTesting(testdb, testAddress, big.NewInt(1000000000))
	blocks, _ := core.GenerateChain(config, genesis, ethash.NewFaker(), testdb, n, func(i int, block *core.BlockGen) {
		block.SetCoinbase(common.Address{0})
	})

	hashes := make([]common.Hash, n+1)
	hashes[len(hashes)-1] = genesis.Hash()
	blockm := make(map[common.Hash]*types.Block, n+1)
	blockm[genesis.Hash()] = genesis
	for i, b := range blocks {
		hashes[len(hashes)-i-2] = b.Hash()
		blockm[b.Hash()] = b
	}

	blockchain, _ := core.NewBlockChain(testdb, nil, config, ethash.NewFaker(), vm.Config{}, nil, nil, nil)
	return blocks, blockm, blockchain
}

func buildTestChainWithOneTxPerBlock(n int, config *params.ChainConfig) ([]*types.Block, map[common.Hash]*types.Block, *core.BlockChain) {
	testdb := rawdb.NewMemoryDatabase()
	genesis := core.GenesisBlockForTesting(testdb, testAddress, big.NewInt(1000000000))
	blocks, _ := core.GenerateChain(config, genesis, ethash.NewFaker(), testdb, n, func(i int, block *core.BlockGen) {
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

	blockchain, _ := core.NewBlockChain(testdb, nil, config, ethash.NewFaker(), vm.Config{}, nil, nil, nil)
	return blocks, blockm, blockchain
}

func buildTestChainWithOnePMTTxPerBlock(n int, config *params.ChainConfig) ([]*types.Block, map[common.Hash]*types.Block, *core.BlockChain) {
	testdb := rawdb.NewMemoryDatabase()
	genesis := core.GenesisBlockForTesting(testdb, testAddress, big.NewInt(1000000000))
	blocks, _ := core.GenerateChain(config, genesis, ethash.NewFaker(), testdb, n, func(i int, block *core.BlockGen) {
		block.SetCoinbase(common.Address{0})

		signer := types.LatestSigner(config)
		tx, err := types.SignTx(types.NewTransaction(block.TxNonce(testAddress), common.QuorumPrivacyPrecompileContractAddress(), big.NewInt(0), testGas, nil, common.BytesToEncryptedPayloadHash([]byte("pmt inner tx")).Bytes()), signer, testKey)
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

	blockchain, _ := core.NewBlockChain(testdb, nil, config, ethash.NewFaker(), vm.Config{}, nil, nil, nil)
	return blocks, blockm, blockchain
}

var PSI1PSM = mps.NewPrivateStateMetadata("psi1", "psi1", "private state 1", mps.Resident, PG1.Members)

var PSI2PSM = mps.NewPrivateStateMetadata("psi2", "psi2", "private state 2", mps.Resident, PG2.Members)

var PG1 = engine.PrivacyGroup{
	Type:           "RESIDENT",
	Name:           "RG1",
	PrivacyGroupId: "RG1",
	Description:    "Resident Group 1",
	From:           "",
	Members:        []string{"AAA", "BBB"},
}

var PG2 = engine.PrivacyGroup{
	Type:           "RESIDENT",
	Name:           "RG2",
	PrivacyGroupId: "RG2",
	Description:    "Resident Group 2",
	From:           "",
	Members:        []string{"CCC", "DDD"},
}

var PrivacyGroups = []engine.PrivacyGroup{
	PG1,
	PG2,
	{
		Type:           "LEGACY",
		Name:           "LEGACY1",
		PrivacyGroupId: "LEGACY1",
		Description:    "Legacy Group 1",
		From:           "",
		Members:        []string{"LEG1", "LEG2"},
	},
}

type testAuthManager struct {
	enabled   bool
	authError error
	authToken *proto.PreAuthenticatedAuthenticationToken
}

func (am *testAuthManager) Authenticate(ctx context.Context, token string) (*proto.PreAuthenticatedAuthenticationToken, error) {
	return am.authToken, am.authError
}

func (am *testAuthManager) IsEnabled(ctx context.Context) (bool, error) {
	return am.enabled, nil
}
