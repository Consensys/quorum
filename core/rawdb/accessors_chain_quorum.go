package rawdb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
)

// HasBadBlock returns whether the block with the hash is a bad block. dep: Istanbul
func HasBadBlock(db ethdb.Reader, hash common.Hash) bool {
	return ReadBadBlock(db, hash) != nil
}
