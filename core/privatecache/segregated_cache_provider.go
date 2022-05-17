package privatecache

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
)

type segregatedCacheProvider struct {
	db     ethdb.Database
	config *trie.Config
}

func (p *segregatedCacheProvider) GetCache() state.Database {
	return state.NewDatabase(p.db)
}

func (p *segregatedCacheProvider) GetCacheWithConfig() state.Database {
	return state.NewDatabaseWithConfig(p.db, p.config)
}

func (p *segregatedCacheProvider) Commit(db state.Database, hash common.Hash) error {
	return db.TrieDB().Commit(hash, false, nil)
}
func (p *segregatedCacheProvider) Reference(child, parent common.Hash) {
	// do nothing
}
