package extension

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/extension/extensionContracts"
)

// ChainAccessor provides methods to fetch state and blocks from the local blockchain
type ChainAccessor interface {
	// GetBlockByHash retrieves a block from the local chain.
	GetBlockByHash(common.Hash) *types.Block
	StateAt(root common.Hash) (*state.StateDB, *state.StateDB, error)
	State() (*state.StateDB, *state.StateDB, error)
	GetReceiptsByHash(hash common.Hash) types.Receipts
}

// StateFetcher manages retrieving state from the database and returning it in
// a usable form by the extension API.
type StateFetcher struct {
	chainAccessor ChainAccessor
	ethDb         ethdb.Database
}

// Creates a new StateFetcher from the ethereum service
func NewStateFetcher(chainAccessor ChainAccessor, chainDb ethdb.Database) *StateFetcher {
	return &StateFetcher{
		chainAccessor: chainAccessor,
		ethDb:         chainDb,
	}
}

// GetAddressStateFromBlock is a public method that combines the other
// functions of a StateFetcher, retrieving the state of an address at a given
// block, represented in JSON.
func (fetcher *StateFetcher) GetAddressStateFromBlock(blockHash common.Hash, addressToFetch common.Address) ([]byte, error) {
	privateState, err := fetcher.privateState(blockHash)
	if err != nil {
		return nil, err
	}
	stateData, err := fetcher.addressStateAsJson(privateState, addressToFetch)
	if err != nil {
		return nil, err
	}
	return stateData, nil
}

// privateState returns the private state database for a given block hash.
func (fetcher *StateFetcher) privateState(blockHash common.Hash) (*state.StateDB, error) {
	block := fetcher.chainAccessor.GetBlockByHash(blockHash)
	_, privateState, err := fetcher.chainAccessor.StateAt(block.Root())
	return privateState, err
}

// addressStateAsJson returns the state of an address, including the balance,
// nonce, code and state data as a JSON map.
func (fetcher *StateFetcher) addressStateAsJson(privateState *state.StateDB, addressToShare common.Address) ([]byte, error) {
	keepAddresses := make(map[string]extensionContracts.AccountWithMetadata)

	if account, found := privateState.DumpAddress(addressToShare); found {
		keepAddresses[addressToShare.Hex()] = extensionContracts.AccountWithMetadata{
			State: account,
		}
	} else {
		return nil, fmt.Errorf("error in contract state fetch")
	}
	//types can be marshalled, so errors can't occur
	out, _ := json.Marshal(&keepAddresses)
	return out, nil
}

// fethches the transaction object from transaction hash given
func (fetcher *StateFetcher) GetTransaction(txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64) {
	tx, blockHash, blockNumber, index := rawdb.ReadTransaction(fetcher.ethDb, txHash)
	return tx, blockHash, blockNumber, index
}

// fetches the transaction receipts object for the given blockhash and transaction index
func (fetcher *StateFetcher) GetTransactionReceipt(blockHash common.Hash, index uint64) *types.Receipt {
	receipts := fetcher.chainAccessor.GetReceiptsByHash(blockHash)

	if len(receipts) <= int(index) {
		return nil
	}
	receipt := receipts[index]
	return receipt
}
