package core

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/privatecache"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
)

type MultiplePrivateStateManager struct {
	// Low level persistent database to store final content in
	db                     ethdb.Database
	privateStatesTrieCache state.Database
	privateCacheProvider   privatecache.Provider

	residentGroupByKey map[string]*mps.PrivateStateMetadata
	privacyGroupById   map[types.PrivateStateIdentifier]*mps.PrivateStateMetadata
}

func newMultiplePrivateStateManager(db ethdb.Database, privateCacheProvider privatecache.Provider, residentGroupByKey map[string]*mps.PrivateStateMetadata, privacyGroupById map[types.PrivateStateIdentifier]*mps.PrivateStateMetadata) (*MultiplePrivateStateManager, error) {
	return &MultiplePrivateStateManager{
		db:                     db,
		privateStatesTrieCache: privateCacheProvider.GetCacheWithConfig(),
		privateCacheProvider:   privateCacheProvider,
		residentGroupByKey:     residentGroupByKey,
		privacyGroupById:       privacyGroupById,
	}, nil
}

func (m *MultiplePrivateStateManager) StateRepository(blockHash common.Hash) (mps.PrivateStateRepository, error) {
	privateStatesTrieRoot := rawdb.GetPrivateStatesTrieRoot(m.db, blockHash)
	return mps.NewMultiplePrivateStateRepository(m.db, m.privateStatesTrieCache, privateStatesTrieRoot, m.privateCacheProvider)
}

func (m *MultiplePrivateStateManager) ResolveForManagedParty(managedParty string) (*mps.PrivateStateMetadata, error) {
	psm, found := m.residentGroupByKey[managedParty]
	if !found {
		return nil, fmt.Errorf("unable to find private state metadata for managed party %s", managedParty)
	}
	return psm, nil
}

func (m *MultiplePrivateStateManager) ResolveForUserContext(ctx context.Context) (*mps.PrivateStateMetadata, error) {
	psi, ok := rpc.PrivateStateIdentifierFromContext(ctx)
	if !ok {
		psi = types.DefaultPrivateStateIdentifier
	}
	psm, found := m.privacyGroupById[psi]
	if !found {
		return nil, fmt.Errorf("unable to find private state for context psi %s", psi)
	}
	return psm, nil
}

func (m *MultiplePrivateStateManager) PSIs() []types.PrivateStateIdentifier {
	psis := make([]types.PrivateStateIdentifier, 0, len(m.privacyGroupById))
	for psi := range m.privacyGroupById {
		psis = append(psis, psi)
	}
	return psis
}

func (m *MultiplePrivateStateManager) NotIncludeAny(psm *mps.PrivateStateMetadata, managedParties ...string) bool {
	return psm.NotIncludeAny(managedParties...)
}

func (m *MultiplePrivateStateManager) CheckAt(root common.Hash) error {
	_, err := state.New(rawdb.GetPrivateStatesTrieRoot(m.db, root), m.privateStatesTrieCache, nil)
	return err
}

func (m *MultiplePrivateStateManager) TrieDB() *trie.Database {
	return m.privateStatesTrieCache.TrieDB()
}
