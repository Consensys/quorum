package qlight

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/cache"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/private/engine/qlightptm"
	gocache "github.com/patrickmn/go-cache"
)

type clientCache struct {
	cachingTXManager  *qlightptm.CachingProxyTxManager
	privateBlockCache *gocache.Cache
	db                ethdb.Database
}

func NewClientCache(db ethdb.Database) (PrivateClientCache, error) {
	cachingTXManager, ok := private.P.(*qlightptm.CachingProxyTxManager)
	if !ok {
		return nil, fmt.Errorf("unable to initialize cachingTXManager")
	}
	return &clientCache{
		cachingTXManager:  cachingTXManager,
		privateBlockCache: gocache.New(cache.DefaultExpiration, cache.CleanupInterval),
		db:                db,
	}, nil
}

func (c *clientCache) AddPrivateBlock(key QLightCacheKey) error {
	var result engine.BlockPrivatePayloads
	err := c.cachingTXManager.GetRPCClient().Call(&result, "eth_getQuorumPayloadsForBlock", key.String())
	if err != nil {
		return err
	}
	for txKey, qpe := range result.Payloads {
		eph, err := common.Base64ToEncryptedPayloadHash(txKey)
		if err != nil {
			return err
		}
		if len(qpe.Payload) > 3 {
			payloadBytes, err := hex.DecodeString(qpe.Payload[2:])
			if err != nil {
				return err
			}
			c.cachingTXManager.AddToCache(eph, payloadBytes, qpe.ExtraMetaData, qpe.IsSender)
		}
	}
	return c.privateBlockCache.Add(result.BlockHash, result.PrivateStateRoot, gocache.DefaultExpiration)
}

func (c *clientCache) CheckAndAddEmptyEntry(hash common.EncryptedPayloadHash) {
	c.cachingTXManager.CheckAndAddEmptyToCache(hash)
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
