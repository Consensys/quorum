package ethapi

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"

	"github.com/ethereum/go-ethereum/consensus/ethash"

	"github.com/ethereum/go-ethereum/core/state"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	arbitraryCtx  = context.Background()
	privateTxArgs = &PrivateTxArgs{
		PrivateFor: []string{"arbitrary party 1", "arbitrary party 2"},
	}
	arbitraryFrom = common.BytesToAddress([]byte("arbitrary address"))
	arbitraryTo   = common.BytesToAddress([]byte("aribitrary to"))

	arbitrarySimpleStorageContractEncryptedPayloadHash = common.BytesToEncryptedPayloadHash([]byte("arbitrary payload hash"))
	arbitraryChildContractEncryptedPayloadHash         = common.BytesToEncryptedPayloadHash([]byte("arbitrary payload hash 1"))

	simpleStorageContractCreationTx = types.NewContractCreation(
		0,
		big.NewInt(0),
		hexutil.MustDecodeUint64("0x47b760"),
		big.NewInt(0),
		hexutil.MustDecode("0x6060604052341561000f57600080fd5b604051602080610149833981016040528080519060200190919050505b806000819055505b505b610104806100456000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a1afcd914605157806360fe47b11460775780636d4ce63c146097575b600080fd5b3415605b57600080fd5b606160bd565b6040518082815260200191505060405180910390f35b3415608157600080fd5b6095600480803590602001909190505060c3565b005b341560a157600080fd5b60a760ce565b6040518082815260200191505060405180910390f35b60005481565b806000819055505b50565b6000805490505b905600a165627a7a72305820d5851baab720bba574474de3d09dbeaabc674a15f4dd93b974908476542c23f00029"))

	arbitrarySimpleStorageContractAddress = common.HexToAddress("0x9ebd4609f9e416232ebcf0e59b4104faa7504104")
	arbitraryChildContractAddress         = common.HexToAddress("0x8ebd4609F9e416232Ebcf0E59B4104FaA7504104")

	childContractCreationTx = types.NewContractCreation(
		0,
		big.NewInt(0),
		hexutil.MustDecodeUint64("0x47b760"),
		big.NewInt(0),
		hexutil.MustDecode("0x608060405234801561001057600080fd5b506040516020806103028339810180604052602081101561003057600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610271806100916000396000f3fe608060405260043610610046576000357c01000000000000000000000000000000000000000000000000000000009004806360fe47b11461004b5780636d4ce63c14610086575b600080fd5b34801561005757600080fd5b506100846004803603602081101561006e57600080fd5b81019080803590602001909291905050506100b1565b005b34801561009257600080fd5b5061009b610180565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360fe47b1826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561014157600080fd5b505af1158015610155573d6000803e3d6000fd5b505050506040513d602081101561016b57600080fd5b81019080805190602001909291905050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d4ce63c6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561020557600080fd5b505afa158015610219573d6000803e3d6000fd5b505050506040513d602081101561022f57600080fd5b810190808051906020019092919050505090509056fea165627a7a72305820e6629b99a93d3697f4ac2daa2e2a30f3e9126ee6f767a6444fef9d60c77dbd7600290000000000000000000000009ebd4609f9e416232ebcf0e59b4104faa7504104"))

	// transaction to child contract which would invoke simple contract
	childContractTx = types.NewTransaction(
		0,
		arbitraryChildContractAddress,
		big.NewInt(0),
		hexutil.MustDecodeUint64("0x47b760"),
		big.NewInt(0),
		hexutil.MustDecode("0x60fe47b1000000000000000000000000000000000000000000000000000000000000000a"))

	arbitraryCurrentBlockNumber = big.NewInt(1)

	publicStateDB       *state.StateDB
	privateStateDB      *state.StateDB
	arbitraryBlockChain *core.BlockChain

	quorumChainConfig = &params.ChainConfig{big.NewInt(10), big.NewInt(0), nil, false, nil, common.Hash{}, nil, nil, big.NewInt(0), nil, new(params.EthashConfig), nil, nil, true, 64}
)

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func setup() {
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.TerminalFormat(true)))
	testdb := ethdb.NewMemDatabase()
	gspec := &core.Genesis{Config: params.TestChainConfig}
	gspec.MustCommit(testdb)
	var err error
	arbitraryBlockChain, err = core.NewBlockChain(testdb, nil, quorumChainConfig, ethash.NewFaker(), vm.Config{})
	if err != nil {
		panic(err)
	}

	publicStateDB, err = state.New(common.Hash{}, state.NewDatabase(testdb))
	if err != nil {
		panic(err)
	}
	publicStateDB.SetPersistentEthdb(testdb)

	privateStateDB, err = state.New(common.Hash{}, state.NewDatabase(testdb))
	if err != nil {
		panic(err)
	}
	privateStateDB.SetPersistentEthdb(testdb)

	private.P = &StubPrivateTransactionManager{}
}

func teardown() {
	arbitraryBlockChain.Stop()
	log.Root().SetHandler(log.DiscardHandler())
}

func TestSimulateExecution_whenTypicalCreation(t *testing.T) {
	assert := assert.New(t)

	affectedCACreationTxHashes, merkleRoot, err := simulateExecution(arbitraryCtx, &StubBackend{}, arbitraryFrom, simpleStorageContractCreationTx, privateTxArgs)

	assert.NoError(err, "simulation execution")
	assert.Empty(affectedCACreationTxHashes, "creation tx should not have any affected contract creation tx hashes")
	assert.Equal(common.Hash{}, merkleRoot, "no private state validation")
}

func TestSimulateExecution_whenCreationWithStateValidation(t *testing.T) {
	assert := assert.New(t)
	privateTxArgs.PrivateStateValidation = true

	affectedCACreationTxHashes, merkleRoot, err := simulateExecution(arbitraryCtx, &StubBackend{}, arbitraryFrom, simpleStorageContractCreationTx, privateTxArgs)

	assert.NoError(err, "simulation execution")
	assert.Empty(affectedCACreationTxHashes, "creation tx should not have any affected contract creation tx hashes")
	assert.NotEqual(common.Hash{}, merkleRoot, "no private state validation")
}

func TestSimulateExecution_whenNestedContractInteraction(t *testing.T) {
	assert := assert.New(t)

	backend := &StubBackend{}
	privateStateDB.SetCode(arbitrarySimpleStorageContractAddress, hexutil.MustDecode("0x6080604052600436106043576000357c01000000000000000000000000000000000000000000000000000000009004806360fe47b11460485780636d4ce63c146093575b600080fd5b348015605357600080fd5b50607d60048036036020811015606857600080fd5b810190808035906020019092919050505060bb565b6040518082815260200191505060405180910390f35b348015609e57600080fd5b5060a560ce565b6040518082815260200191505060405180910390f35b6000816000819055506000549050919050565b6000805490509056fea165627a7a7230582000f86094047c0a1b33312cb6f75b5980eac40efc58a6589752689031beb32cf50029"))
	_ = privateStateDB.SetStatePrivacyMetadata(arbitrarySimpleStorageContractAddress, &state.PrivacyMetadata{
		PrivateStateValidation: privateTxArgs.PrivateStateValidation,
		CreationTxHash:         arbitrarySimpleStorageContractEncryptedPayloadHash,
	})
	privateStateDB.SetState(arbitrarySimpleStorageContractAddress, common.Hash{0}, common.Hash{100})
	privateStateDB.Commit(true)
	// privateStateDB.SetCode(arbitraryChildContractAddress, hexutil.MustDecode("0x60806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806360fe47b1146100515780636d4ce63c1461008c575b600080fd5b34801561005d57600080fd5b5061008a6004803603602081101561007457600080fd5b81019080803590602001909291905050506100b7565b005b34801561009857600080fd5b506100a1610186565b6040518082815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360fe47b1826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561014757600080fd5b505af115801561015b573d6000803e3d6000fd5b505050506040513d602081101561017157600080fd5b81019080805190602001909291905050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d4ce63c6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561020b57600080fd5b505afa15801561021f573d6000803e3d6000fd5b505050506040513d602081101561023557600080fd5b810190808051906020019092919050505090509056fea165627a7a72305820b2e3e502953d6a6abeed571a333c8f35135b47e2cced8e4a23eab6a16dfa624c0029"))
	_, _, err := simulateExecution(arbitraryCtx, backend, arbitraryFrom, childContractCreationTx, privateTxArgs)
	_ = privateStateDB.SetStatePrivacyMetadata(arbitraryChildContractAddress, &state.PrivacyMetadata{
		PrivateStateValidation: privateTxArgs.PrivateStateValidation,
		CreationTxHash:         arbitraryChildContractEncryptedPayloadHash,
	})
	log.Debug("state", "state", privateStateDB.GetState(arbitraryChildContractAddress, common.Hash{0}))

	log.Debug("execute child contract function")
	affectedCACreationTxHashes, _, err := simulateExecution(arbitraryCtx, backend, arbitraryFrom, childContractTx, privateTxArgs)

	assert.NoError(err, "simulation execution")
	assert.NotEmpty(affectedCACreationTxHashes, "affected contract accounts' creation transacton hashes")
	assert.True(!affectedCACreationTxHashes.NotExist(arbitrarySimpleStorageContractEncryptedPayloadHash), "%s is an affected contract account", arbitrarySimpleStorageContractAddress.Hex())

}

func TestHandlePrivateTransaction_whenTypicalCreation(t *testing.T) {
	assert := assert.New(t)

	isPrivate, _, err := handlePrivateTransaction(arbitraryCtx, &StubBackend{}, simpleStorageContractCreationTx, privateTxArgs, arbitraryFrom, false)

	if err != nil {
		t.Fatalf("%s", err)
	}

	assert.True(isPrivate, "must be a private transaction")
}

type StubBackend struct {
}

func (sb *StubBackend) GetEVM(ctx context.Context, msg core.Message, state vm.MinimalApiState, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	context := core.NewEVMContext(msg, &types.Header{
		Coinbase:   arbitraryFrom,
		Number:     arbitraryCurrentBlockNumber,
		Time:       big.NewInt(0),
		Difficulty: big.NewInt(0),
		GasLimit:   0,
	}, arbitraryBlockChain, nil)
	return vm.NewEVM(context, publicStateDB, privateStateDB, quorumChainConfig, vmCfg), nil, nil
}

func (sb *StubBackend) CurrentBlock() *types.Block {
	return types.NewBlock(&types.Header{
		Number: arbitraryCurrentBlockNumber,
	}, nil, nil, nil)
}

func (sb *StubBackend) Downloader() *downloader.Downloader {
	panic("implement me")
}

func (sb *StubBackend) ProtocolVersion() int {
	panic("implement me")
}

func (sb *StubBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	panic("implement me")
}

func (sb *StubBackend) ChainDb() ethdb.Database {
	panic("implement me")
}

func (sb *StubBackend) EventMux() *event.TypeMux {
	panic("implement me")
}

func (sb *StubBackend) AccountManager() *accounts.Manager {
	panic("implement me")
}

func (sb *StubBackend) SetHead(number uint64) {
	panic("implement me")
}

func (sb *StubBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	panic("implement me")
}

func (sb *StubBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	panic("implement me")
}

func (sb *StubBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (vm.MinimalApiState, *types.Header, error) {
	return &StubMinimalApiState{}, nil, nil
}

func (sb *StubBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	panic("implement me")
}

func (sb *StubBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	panic("implement me")
}

func (sb *StubBackend) GetTd(blockHash common.Hash) *big.Int {
	panic("implement me")
}

func (sb *StubBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	panic("implement me")
}

func (sb *StubBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	panic("implement me")
}

func (sb *StubBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	panic("implement me")
}

func (sb *StubBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	panic("implement me")
}

func (sb *StubBackend) GetPoolTransactions() (types.Transactions, error) {
	panic("implement me")
}

func (sb *StubBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	panic("implement me")
}

func (sb *StubBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	panic("implement me")
}

func (sb *StubBackend) Stats() (pending int, queued int) {
	panic("implement me")
}

func (sb *StubBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	panic("implement me")
}

func (sb *StubBackend) SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription {
	panic("implement me")
}

func (sb *StubBackend) ChainConfig() *params.ChainConfig {
	panic("implement me")
}

type StubMinimalApiState struct {
}

func (smps *StubMinimalApiState) GetBalance(addr common.Address) *big.Int {
	panic("implement me")
}

func (smps *StubMinimalApiState) GetCode(addr common.Address) []byte {
	panic("implement me")
}

func (smps *StubMinimalApiState) GetState(a common.Address, b common.Hash) common.Hash {
	panic("implement me")
}

func (smps *StubMinimalApiState) GetNonce(addr common.Address) uint64 {
	panic("implement me")
}

func (smps *StubMinimalApiState) GetStatePrivacyMetadata(addr common.Address) (*state.PrivacyMetadata, error) {
	panic("implement me")
}

func (smps *StubMinimalApiState) GetRLPEncodedStateObject(addr common.Address) ([]byte, error) {
	panic("implement me")
}

type StubPrivateTransactionManager struct {
}

func (sptm *StubPrivateTransactionManager) Name() string {
	return "StubPrivateTransactionManager"
}

func (sptm *StubPrivateTransactionManager) Send(data []byte, from string, to []string, extra *engine.ExtraMetadata) (common.EncryptedPayloadHash, error) {
	return arbitrarySimpleStorageContractEncryptedPayloadHash, nil
}

func (sptm *StubPrivateTransactionManager) SendSignedTx(data common.EncryptedPayloadHash, to []string, extra *engine.ExtraMetadata) ([]byte, error) {
	panic("implement me")
}

func (sptm *StubPrivateTransactionManager) Receive(data common.EncryptedPayloadHash) ([]byte, *engine.ExtraMetadata, error) {
	panic("implement me")
}

func (sptm *StubPrivateTransactionManager) ReceiveRaw(data common.EncryptedPayloadHash) ([]byte, *engine.ExtraMetadata, error) {
	panic("implement me")
}
