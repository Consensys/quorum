package cache

import (
	"github.com/ethereum/go-ethereum/private/engine"
	"time"
)

const (
	DefaultExpiration = 5 * time.Minute
	CleanupInterval   = 5 * time.Minute
)

type PrivateCacheItem struct {
	Payload []byte
	Extra   engine.ExtraMetadata
}
