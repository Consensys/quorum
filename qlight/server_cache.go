package qlight

import (
	"github.com/ethereum/go-ethereum/private/cache"
	gocache "github.com/patrickmn/go-cache"
)

var privateDataCache = gocache.New(cache.DefaultExpiration, cache.CleanupInterval)

func AddDataToServerCache(key *QLightCacheKey, data PrivateTransactionsData) error {
	return privateDataCache.Add(key.String(), data, gocache.DefaultExpiration)
}

func GetDataFromServerCache(key *QLightCacheKey) (PrivateTransactionsData, bool) {
	cacheItem, found := privateDataCache.Get(key.String())
	if !found {
		return nil, false
	}

	privateTxsData, ok := cacheItem.(PrivateTransactionsData)
	if !ok {
		return nil, false
	}
	return privateTxsData, true
}
