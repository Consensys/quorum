// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package eth

import (
	"math/big"

	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	rpc "github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/net/context"
)

var raftHasNoPending = fmt.Errorf("Raft mode has no Pending block. Use latest instead.")

// EthApiBackend implements ethapi.Backend for full nodes
type EthApiBackend struct {
	eth *Ethereum
}

func (b *EthApiBackend) SetHead(number uint64) {
	b.eth.blockchain.SetHead(number)
}

func (b *EthApiBackend) HeaderByNumber(blockNr rpc.BlockNumber) *types.Header {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		block, _, _ := b.eth.blockVoting.Pending()
		return block.Header()
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.eth.blockchain.CurrentBlock().Header()
	}
	return b.eth.blockchain.GetHeaderByNumber(uint64(blockNr))
}

func (b *EthApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	// Pending block is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		if b.eth.protocolManager.raftMode {
			return nil, raftHasNoPending
		}
		block, _, _ := b.eth.blockVoting.Pending()
		return block, nil
	}
	// Otherwise resolve and return the block
	if blockNr == rpc.LatestBlockNumber {
		return b.eth.blockchain.CurrentBlock(), nil
	}
	return b.eth.blockchain.GetBlockByNumber(uint64(blockNr)), nil
}

func (b *EthApiBackend) StateAndHeaderByNumber(blockNr rpc.BlockNumber) (ethapi.State, *types.Header, error) {
	// Pending state is only known by the miner
	if blockNr == rpc.PendingBlockNumber {
		if b.eth.protocolManager.raftMode {
			return nil, nil, raftHasNoPending
		}
		block, publicState, privateState := b.eth.blockVoting.Pending()
		return EthApiState{publicState, privateState}, block.Header(), nil
	}
	// Otherwise resolve the block number and return its state
	header := b.HeaderByNumber(blockNr)
	if header == nil {
		return nil, nil, nil
	}
	publicState, privateState, err := b.eth.BlockChain().StateAt(header.Root)
	return EthApiState{publicState, privateState}, header, err
}

func (b *EthApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	return b.eth.blockchain.GetBlockByHash(blockHash), nil
}

func (b *EthApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return core.GetBlockReceipts(b.eth.chainDb, blockHash, core.GetBlockNumber(b.eth.chainDb, blockHash)), nil
}

func (b *EthApiBackend) GetTd(blockHash common.Hash) *big.Int {
	return b.eth.blockchain.GetTdByHash(blockHash)
}

func (b *EthApiBackend) GetVMEnv(ctx context.Context, msg core.Message, state ethapi.State, header *types.Header) (vm.Environment, func() error, error) {
	var (
		statedb      = state.(EthApiState)
		publicState  = statedb.publicState
		privateState = statedb.privateState
	)

	if msg.To() != nil && !privateState.Exist(*msg.To()) {
		privateState = publicState
	}

	addr, _ := msg.From()
	from := privateState.GetOrNewStateObject(addr)
	from.SetBalance(common.MaxBig)
	vmError := func() error { return nil }
	return core.NewEnv(publicState, privateState, b.eth.chainConfig, b.eth.blockchain, msg, header, b.eth.chainConfig.VmConfig), vmError, nil
}

func (b *EthApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	b.eth.txMu.Lock()
	defer b.eth.txMu.Unlock()

	b.eth.txPool.SetLocal(signedTx)
	return b.eth.txPool.Add(signedTx)
}

func (b *EthApiBackend) RemoveTx(txHash common.Hash) {
	b.eth.txMu.Lock()
	defer b.eth.txMu.Unlock()

	b.eth.txPool.Remove(txHash)
}

func (b *EthApiBackend) GetPoolTransactions() types.Transactions {
	b.eth.txMu.Lock()
	defer b.eth.txMu.Unlock()

	var txs types.Transactions
	for _, batch := range b.eth.txPool.Pending() {
		txs = append(txs, batch...)
	}
	return txs
}

func (b *EthApiBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	b.eth.txMu.Lock()
	defer b.eth.txMu.Unlock()

	return b.eth.txPool.Get(hash)
}

func (b *EthApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	b.eth.txMu.Lock()
	defer b.eth.txMu.Unlock()

	return b.eth.txPool.State().GetNonce(addr), nil
}

func (b *EthApiBackend) Stats() (pending int, queued int) {
	b.eth.txMu.Lock()
	defer b.eth.txMu.Unlock()

	return b.eth.txPool.Stats()
}

func (b *EthApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	b.eth.txMu.Lock()
	defer b.eth.txMu.Unlock()

	return b.eth.TxPool().Content()
}

func (b *EthApiBackend) Downloader() *downloader.Downloader {
	return b.eth.Downloader()
}

func (b *EthApiBackend) ProtocolVersion() int {
	return b.eth.EthVersion()
}

func (b *EthApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (b *EthApiBackend) ChainDb() ethdb.Database {
	return b.eth.ChainDb()
}

func (b *EthApiBackend) EventMux() *event.TypeMux {
	return b.eth.EventMux()
}

func (b *EthApiBackend) AccountManager() *accounts.Manager {
	return b.eth.AccountManager()
}

type EthApiState struct {
	publicState, privateState *state.StateDB
}

func (s EthApiState) GetBalance(ctx context.Context, addr common.Address) (*big.Int, error) {
	if s.publicState.Exist(addr) {
		return s.publicState.GetBalance(addr), nil
	}
	return s.privateState.GetBalance(addr), nil
}

func (s EthApiState) GetCode(ctx context.Context, addr common.Address) ([]byte, error) {
	if s.publicState.Exist(addr) {
		return s.publicState.GetCode(addr), nil
	}
	return s.privateState.GetCode(addr), nil
}

func (s EthApiState) GetState(ctx context.Context, a common.Address, b common.Hash) (common.Hash, error) {
	if s.publicState.Exist(a) {
		return s.publicState.GetState(a, b), nil
	}
	return s.privateState.GetState(a, b), nil
}

func (s EthApiState) GetNonce(ctx context.Context, addr common.Address) (uint64, error) {
	if s.publicState.Exist(addr) {
		return s.publicState.GetNonce(addr), nil
	}
	return s.privateState.GetNonce(addr), nil
}
