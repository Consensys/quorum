package privatecache

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/trie"
)

type Provider interface {
	GetCache() state.Database
	GetCacheWithConfig() state.Database
	Commit(db state.Database, hash common.Hash) error
	Reference(child, parent common.Hash)
}

func NewPrivateCacheProvider(db ethdb.Database, config *trie.Config, cache state.Database, privateCacheEnabled bool) Provider {
	if privateCacheEnabled {
		log.Info("Using UnifiedCacheProvider.")
		return &unifiedCacheProvider{
			cache: cache,
		}
	}
	log.Info("Using SegregatedCacheProvider.")
	return &segregatedCacheProvider{db: db, config: config}
}
