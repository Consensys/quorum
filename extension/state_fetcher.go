package extension

import (
	"encoding/json"

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
}

// StateFetcher manages retrieving state from the database and returning it in
// a usable form by the extension API.
type StateFetcher struct {
	chainAccessor ChainAccessor
}

// Creates a new StateFetcher from the ethereum service
func NewStateFetcher(chainAccessor ChainAccessor) *StateFetcher {
	return &StateFetcher{
		chainAccessor: chainAccessor,
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
	block := fetcher.chainAccessor.GetBlockByHash(blockHash)
	_, privateState, err := fetcher.chainAccessor.StateAt(block.Root())

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
