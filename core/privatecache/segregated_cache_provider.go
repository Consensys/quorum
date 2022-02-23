package privatecache

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
)

type segregatedCacheProvider struct {
	db ethdb.Database
}

func (p *segregatedCacheProvider) GetCache() state.Database {
	return state.NewDatabase(p.db)
}

func (p *segregatedCacheProvider) Commit(db state.Database, hash common.Hash) error {
	return db.TrieDB().Commit(hash, false, nil)
}
func (p *segregatedCacheProvider) Reference(child, parent common.Hash) {
	// do nothing
}
