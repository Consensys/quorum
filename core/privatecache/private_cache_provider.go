package privatecache

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
)

type PrivateCacheProvider interface {
	GetCache() state.Database
	Commit(db state.Database, hash common.Hash) error
	Reference(child, parent common.Hash)
}

func NewPrivateCacheProvider(db ethdb.Database, cache state.Database, privateCacheEnabled bool) PrivateCacheProvider {
	if privateCacheEnabled {
		log.Info("Using UnifiedCacheProvider.")
		return &unifiedCacheProvider{
			cache: cache,
		}
	}
	log.Info("Using SegregatedCacheProvider.")
	return &segregatedCacheProvider{db: db}
}
