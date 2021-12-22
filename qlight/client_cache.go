package qlight

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/cache"
	"github.com/ethereum/go-ethereum/private/engine/qlightptm"
	gocache "github.com/patrickmn/go-cache"
)

type clientCache struct {
	cacheWithEmpty    CacheWithEmpty
	privateBlockCache *gocache.Cache
	db                ethdb.Database
}

func NewClientCache(db ethdb.Database) (PrivateClientCache, error) {
	cachingTXManager, ok := private.P.(*qlightptm.CachingProxyTxManager)
	if !ok {
		return nil, fmt.Errorf("unable to initialize cacheWithEmpty")
	}
	return NewClientCacheWithEmpty(db, cachingTXManager, gocache.New(cache.DefaultExpiration, cache.CleanupInterval))
}

func NewClientCacheWithEmpty(db ethdb.Database, cacheWithEmpty CacheWithEmpty, gocache *gocache.Cache) (PrivateClientCache, error) {
	return &clientCache{
		cacheWithEmpty:    cacheWithEmpty,
		privateBlockCache: gocache,
		db:                db,
	}, nil
}

func (c *clientCache) AddPrivateBlock(blockPrivateData BlockPrivateData) error {
	for _, pvtTx := range blockPrivateData.PrivateTransactions {
		if err := c.cacheWithEmpty.Cache(pvtTx.ToCachable()); err != nil {
			return err
		}
	}
	return c.privateBlockCache.Add(blockPrivateData.BlockHash.ToBase64(), blockPrivateData.PrivateStateRoot.ToBase64(), gocache.DefaultExpiration)
}

func (c *clientCache) CheckAndAddEmptyEntry(hash common.EncryptedPayloadHash) {
	c.cacheWithEmpty.CheckAndAddEmptyToCache(hash)
}

func (c *clientCache) ValidatePrivateStateRoot(blockHash common.Hash, publicStateRoot common.Hash) error {
	dbPrivateStateRoot := rawdb.GetPrivateStateRoot(c.db, publicStateRoot)

	cachePrivateStateRootStr, found := c.privateBlockCache.Get(blockHash.ToBase64())
	if !found {
		// this means that we don't have any private data for this block thus the private state should not have changed
		return nil
	}
	cachePrivateStateRootB64, ok := cachePrivateStateRootStr.(string)
	if !ok {
		return fmt.Errorf("Invalid private block cache item")
	}
	cachePrivateStateRoot, err := common.Base64ToHash(cachePrivateStateRootB64)
	if err != nil {
		return fmt.Errorf("Invalid encoding for private state root: %s", cachePrivateStateRootB64)
	}
	if !bytes.Equal(cachePrivateStateRoot.Bytes(), dbPrivateStateRoot.Bytes()) {
		log.Error("QLight - Private state root hash check failure for block", "hash", blockHash)
		return fmt.Errorf("Private root hash missmatch for block %s", blockHash)
	}
	log.Info("QLight - Private state root hash check successful for block", "hash", blockHash)
	return nil
}
