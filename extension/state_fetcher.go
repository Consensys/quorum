package extension

import (
	"encoding/json"
	"fmt"

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
	CurrentBlock() *types.Block
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

// returns the current block hash
func (fetcher *StateFetcher) getCurrentBlockHash() common.Hash {
	return fetcher.chainAccessor.CurrentBlock().Hash()
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

// returns the privacy metadata
func (fetcher *StateFetcher) GetPrivacyMetaData(blockHash common.Hash, address common.Address) (*state.PrivacyMetadata, error) {
	privateState, err := fetcher.privateState(blockHash)
	if err != nil {
		return nil, err
	}

	privacyMetaData, err := privateState.GetStatePrivacyMetadata(address)
	if err != nil {
		return nil, err
	}

	return privacyMetaData, nil
}

// returns the privacy metadata
func (fetcher *StateFetcher) GetStorageRoot(blockHash common.Hash, address common.Address) (common.Hash, error) {
	privateState, err := fetcher.privateState(blockHash)
	if err != nil {
		return common.Hash{}, err
	}

	storageRoot, err := privateState.GetStorageRoot(address)
	if err != nil {
		return common.Hash{}, err
	}

	return storageRoot, nil
}
