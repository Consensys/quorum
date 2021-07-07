package core

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
)

type DefaultPrivateStateManager struct {
	// Low level persistent database to store final content in
	db        ethdb.Database
	repoCache state.Database
}

func newDefaultPrivateStateManager(db ethdb.Database, cacheConfig *CacheConfig) *DefaultPrivateStateManager {
	return &DefaultPrivateStateManager{
		db:        db,
		repoCache: state.NewDatabaseWithCache(db, cacheConfig.TrieCleanLimit, cacheConfig.PrivateTrieCleanJournal),
	}
}

func (d *DefaultPrivateStateManager) StateRepository(blockHash common.Hash) (mps.PrivateStateRepository, error) {
	return mps.NewDefaultPrivateStateRepository(d.db, d.repoCache, blockHash)
}

func (d *DefaultPrivateStateManager) ResolveForManagedParty(_ string) (*mps.PrivateStateMetadata, error) {
	return mps.DefaultPrivateStateMetadata, nil
}

func (d *DefaultPrivateStateManager) ResolveForUserContext(ctx context.Context) (*mps.PrivateStateMetadata, error) {
	psi, ok := rpc.PrivateStateIdentifierFromContext(ctx)
	if !ok {
		psi = types.DefaultPrivateStateIdentifier
	}
	return &mps.PrivateStateMetadata{ID: psi, Type: mps.Resident}, nil
}

func (d *DefaultPrivateStateManager) PSIs() []types.PrivateStateIdentifier {
	return []types.PrivateStateIdentifier{
		types.DefaultPrivateStateIdentifier,
	}
}

func (d *DefaultPrivateStateManager) NotIncludeAny(_ *mps.PrivateStateMetadata, _ ...string) bool {
	// with default implementation, all managedParties are members of the psm
	return false
}

func (d *DefaultPrivateStateManager) CheckAt(root common.Hash) error {
	_, err := state.New(rawdb.GetPrivateStateRoot(d.db, root), d.repoCache, nil)
	return err
}

func (d *DefaultPrivateStateManager) TrieDB() *trie.Database {
	return d.repoCache.TrieDB()
}
