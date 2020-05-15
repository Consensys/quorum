package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
	"github.com/ethereum/go-ethereum/private/engine"
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