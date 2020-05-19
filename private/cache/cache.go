package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

const (
	DefaultExpiration = 5 * time.Minute
	CleanupInterval   = 5 * time.Minute
)

func NewDefaultCache() *gocache.Cache {
	return gocache.New(DefaultExpiration, CleanupInterval)
}
