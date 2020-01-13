package extension

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/extension/extensionContracts"
)

// BlockRetriever provides methods to fetch blocks from the local blockchain
type BlockRetriever interface {
	// GetBlockByHash retrieves a block from the local chain.
	GetBlockByHash(common.Hash) *types.Block
	StateAt(root common.Hash) (*state.StateDB, *state.StateDB, error)
}

// StateFetcher manages retrieving state from the database and returning it in
// a usable form by the extension API.
type StateFetcher struct {
	db             ethdb.Database
	blockRetriever BlockRetriever
}

// Creates a new StateFetcher from the ethereum service
func NewStateFetcher(db ethdb.Database, blockRetriever BlockRetriever) *StateFetcher {
	return &StateFetcher{
		db:             db,
		blockRetriever: blockRetriever,
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
	return fetcher.addressStateAsJson(privateState, addressToFetch), nil
}

// privateState returns the private state database for a given block hash.
func (fetcher *StateFetcher) privateState(blockHash common.Hash) (*state.StateDB, error) {
	block := fetcher.blockRetriever.GetBlockByHash(blockHash)
	_, privateState, err := fetcher.blockRetriever.StateAt(block.Root())

	return privateState, err
}

// addressStateAsJson returns the state of an address, including the balance,
// nonce, code and state data as a JSON map.
func (fetcher *StateFetcher) addressStateAsJson(privateState *state.StateDB, addressToShare common.Address) []byte {
	keepAddresses := make(map[string]extensionContracts.AccountWithMetadata)

	if account, found := privateState.DumpAddress(addressToShare); found {
		keepAddresses[addressToShare.Hex()] = extensionContracts.AccountWithMetadata{
			State: account,
		}
	}

	//types can be marshalled, so errors can't occur
	out, _ := json.Marshal(&keepAddresses)
	return out
}
