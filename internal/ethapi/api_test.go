package ethapi

import (
	"context"
	"math/big"
	"testing"

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
	arbitraryCtx = context.Background()
	args         = &PrivateTxArgs{
		PrivateFor: []string{"arbitrary party 1", "arbitrary party 2"},
	}
	arbitraryFrom = common.BytesToAddress([]byte("arbitrary address"))
	arbitraryTo   = common.BytesToAddress([]byte("aribitrary to"))

	arbitraryCurrentBlockNumber = big.NewInt(1)
)

func TestHandlePrivateTransaction_whenCreation(t *testing.T) {
	assert := assert.New(t)

	arbitraryData := []byte("data")
	tx := types.NewTransaction(0, arbitraryTo, big.NewInt(1), 0, big.NewInt(1), arbitraryData)
	isPrivate, _, err := handlePrivateTransaction(arbitraryCtx, &StubBackend{}, tx, args, arbitraryFrom, false)

	if err != nil {
		t.Fatalf("%s", err)
	}

	assert.True(isPrivate, "must be a private transaction")
}

type StubBackend struct {
}

func (mb *StubBackend) Downloader() *downloader.Downloader {
	panic("implement me")
}

func (mb *StubBackend) ProtocolVersion() int {
	panic("implement me")
}

func (mb *StubBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	panic("implement me")
}

func (mb *StubBackend) ChainDb() ethdb.Database {
	panic("implement me")
}

func (mb *StubBackend) EventMux() *event.TypeMux {
	panic("implement me")
}

func (mb *StubBackend) AccountManager() *accounts.Manager {
	panic("implement me")
}

func (mb *StubBackend) SetHead(number uint64) {
	panic("implement me")
}

func (mb *StubBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	panic("implement me")
}

func (mb *StubBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	panic("implement me")
}

func (mb *StubBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (vm.MinimalApiState, *types.Header, error) {
	panic("implement me")
}

func (mb *StubBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	panic("implement me")
}

func (mb *StubBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	panic("implement me")
}

func (mb *StubBackend) GetTd(blockHash common.Hash) *big.Int {
	panic("implement me")
}

func (mb *StubBackend) GetEVM(ctx context.Context, msg core.Message, state vm.MinimalApiState, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	panic("implement me")
}

func (mb *StubBackend) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	panic("implement me")
}

func (mb *StubBackend) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	panic("implement me")
}

func (mb *StubBackend) SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription {
	panic("implement me")
}

func (mb *StubBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	panic("implement me")
}

func (mb *StubBackend) GetPoolTransactions() (types.Transactions, error) {
	panic("implement me")
}

func (mb *StubBackend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	panic("implement me")
}

func (mb *StubBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	panic("implement me")
}

func (mb *StubBackend) Stats() (pending int, queued int) {
	panic("implement me")
}

func (mb *StubBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	panic("implement me")
}

func (mb *StubBackend) SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription {
	panic("implement me")
}

func (mb *StubBackend) ChainConfig() *params.ChainConfig {
	panic("implement me")
}

func (mb *StubBackend) CurrentBlock() *types.Block {
	return types.NewBlock(&types.Header{
		Number: arbitraryCurrentBlockNumber,
	}, nil, nil, nil)
}
