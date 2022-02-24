package privatecache

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
)

type unifiedCacheProvider struct {
	cache state.Database
}

func (p *unifiedCacheProvider) GetCache() state.Database {
	return p.cache
}

func (p *unifiedCacheProvider) GetCacheWithConfig() state.Database {
	return p.cache
}

func (p *unifiedCacheProvider) Commit(db state.Database, hash common.Hash) error {
	// do nothing since the references will handle the actual commit (when the public root is committed)
	return nil
}

func (p *unifiedCacheProvider) Reference(child, parent common.Hash) {
	p.cache.TrieDB().Reference(child, parent)
}
