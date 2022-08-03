package mps

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/privatecache"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
)

// DefaultPrivateStateRepository acts as the single private state in the original
// Quorum design.
type DefaultPrivateStateRepository struct {
	db ethdb.Database
	// cache of stateDB
	stateCache           state.Database
	privateCacheProvider privatecache.Provider
	// stateDB gives access to the underlying state
	stateDB *state.StateDB
	root    common.Hash
}

func NewDefaultPrivateStateRepository(db ethdb.Database, cache state.Database, privateCacheProvider privatecache.Provider, previousBlockHash common.Hash) (*DefaultPrivateStateRepository, error) {
	root := rawdb.GetPrivateStateRoot(db, previousBlockHash)

	statedb, err := state.New(root, cache, nil)
	if err != nil {
		return nil, err
	}

	return &DefaultPrivateStateRepository{
		db:                   db,
		stateCache:           cache,
		privateCacheProvider: privateCacheProvider,
		stateDB:              statedb,
		root:                 root,
	}, nil
}

func (dpsr *DefaultPrivateStateRepository) DefaultState() (*state.StateDB, error) {
	if dpsr == nil {
		return nil, fmt.Errorf("nil instance")
	}
	return dpsr.stateDB, nil
}

func (dpsr *DefaultPrivateStateRepository) DefaultStateMetadata() *PrivateStateMetadata {
	return DefaultPrivateStateMetadata
}

func (dpsr *DefaultPrivateStateRepository) IsMPS() bool {
	return false
}

func (dpsr *DefaultPrivateStateRepository) PrivateStateRoot(psi types.PrivateStateIdentifier) (common.Hash, error) {
	return dpsr.root, nil
}

func (dpsr *DefaultPrivateStateRepository) StatePSI(psi types.PrivateStateIdentifier) (*state.StateDB, error) {
	if psi != types.DefaultPrivateStateIdentifier {
		return nil, fmt.Errorf("only the 'private' psi is supported by the default private state manager")
	}
	return dpsr.stateDB, nil
}

func (dpsr *DefaultPrivateStateRepository) Reset() error {
	return dpsr.stateDB.Reset(dpsr.root)
}

// CommitAndWrite commits the private state and writes to disk
func (dpsr *DefaultPrivateStateRepository) CommitAndWrite(isEIP158 bool, block *types.Block) error {
	privateRoot, err := dpsr.stateDB.Commit(isEIP158)
	if err != nil {
		return err
	}

	if err := rawdb.WritePrivateStateRoot(dpsr.db, block.Root(), privateRoot); err != nil {
		log.Error("Failed writing private state root", "err", err)
		return err
	}
	dpsr.privateCacheProvider.Commit(dpsr.stateCache, privateRoot)
	dpsr.privateCacheProvider.Reference(privateRoot, block.Root())
	return nil
}

// Commit commits the private state only
func (dpsr *DefaultPrivateStateRepository) Commit(isEIP158 bool, block *types.Block) error {
	var err error
	dpsr.root, err = dpsr.stateDB.Commit(isEIP158)
	return err
}

func (dpsr *DefaultPrivateStateRepository) Copy() PrivateStateRepository {
	return &DefaultPrivateStateRepository{
		db:                   dpsr.db,
		stateCache:           dpsr.stateCache,
		privateCacheProvider: dpsr.privateCacheProvider,
		stateDB:              dpsr.stateDB.Copy(),
		root:                 dpsr.root,
	}
}

// Given a slice of public receipts and an overlapping (smaller) slice of
// private receipts, return a new slice where the default for each location is
// the public receipt but we take the private receipt in each place we have
// one.
func (dpsr *DefaultPrivateStateRepository) MergeReceipts(pub, priv types.Receipts) types.Receipts {
	m := make(map[common.Hash]*types.Receipt)
	for _, receipt := range pub {
		m[receipt.TxHash] = receipt
	}
	for _, receipt := range priv {
		m[receipt.TxHash] = receipt
	}

	ret := make(types.Receipts, 0, len(pub))
	for _, pubReceipt := range pub {
		ret = append(ret, m[pubReceipt.TxHash])
	}

	return ret
}
