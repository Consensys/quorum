package ethapi

import (
	"context"
	"math/big"
	"os"
	"testing"

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

	arbitraryCurrentBlockNumber = big.NewInt(1)

	publicStateDB       *state.StateDB
	privateStateDB      *state.StateDB
	arbitraryBlockChain *core.BlockChain
)

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func setup() {
	testdb := ethdb.NewMemDatabase()
	gspec := &core.Genesis{Config: params.TestChainConfig}
	gspec.MustCommit(testdb)
	var err error
	arbitraryBlockChain, err = core.NewBlockChain(testdb, nil, params.TestChainConfig, ethash.NewFaker(), vm.Config{})
	if err != nil {
		panic(err)
	}

	publicStateDB, err = state.New(common.Hash{}, state.NewDatabase(ethdb.NewMemDatabase()))
	if err != nil {
		panic(err)
	}

	privateStateDB, err = state.New(common.Hash{}, state.NewDatabase(ethdb.NewMemDatabase()))
	if err != nil {
		panic(err)
	}

	private.P = &StubPrivateTransactionManager{}
}

func teardown() {
	arbitraryBlockChain.Stop()
}

func TestHandlePrivateTransaction_whenCreation(t *testing.T) {
	assert := assert.New(t)

	arbitraryData, err := hexutil.Decode("0x6060604052341561000f57600080fd5b604051602080610149833981016040528080519060200190919050505b806000819055505b505b610104806100456000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a1afcd914605157806360fe47b11460775780636d4ce63c146097575b600080fd5b3415605b57600080fd5b606160bd565b6040518082815260200191505060405180910390f35b3415608157600080fd5b6095600480803590602001909190505060c3565b005b341560a157600080fd5b60a760ce565b6040518082815260200191505060405180910390f35b60005481565b806000819055505b50565b6000805490505b905600a165627a7a72305820d5851baab720bba574474de3d09dbeaabc674a15f4dd93b974908476542c23f00029")
	assert.Nil(err, "%s", err)
	tx := types.NewContractCreation(0, big.NewInt(0), hexutil.MustDecodeUint64("0x47b760"), big.NewInt(0), arbitraryData)
	isPrivate, _, err := handlePrivateTransaction(arbitraryCtx, &StubBackend{}, tx, privateTxArgs, arbitraryFrom, false)

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
	return vm.NewEVM(context, publicStateDB, privateStateDB, params.TestChainConfig, vmCfg), nil, nil
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
	panic("implement me")
}

func (sptm *StubPrivateTransactionManager) Send(data []byte, from string, to []string, extra *engine.ExtraMetadata) (common.EncryptedPayloadHash, error) {
	panic("implement me")
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
