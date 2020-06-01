package cache

import (
	"time"

	"github.com/ethereum/go-ethereum/private/engine"
	gocache "github.com/patrickmn/go-cache"
)

const (
	DefaultExpiration = 5 * time.Minute
	CleanupInterval   = 5 * time.Minute
)

func NewDefaultCache() *gocache.Cache {
	return gocache.New(DefaultExpiration, CleanupInterval)
}

type PrivateCacheItem struct {
	Payload []byte
	Extra   engine.ExtraMetadata
}
