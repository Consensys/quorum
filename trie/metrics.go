package trie

import "github.com/ethereum/go-ethereum/metrics"

type trieDbMetrics struct {
	memcacheCleanHitMeter   metrics.Meter
	memcacheCleanMissMeter  metrics.Meter
	memcacheCleanReadMeter  metrics.Meter
	memcacheCleanWriteMeter metrics.Meter

	memcacheDirtyHitMeter   metrics.Meter
	memcacheDirtyMissMeter  metrics.Meter
	memcacheDirtyReadMeter  metrics.Meter
	memcacheDirtyWriteMeter metrics.Meter

	memcacheFlushTimeTimer  metrics.ResettingTimer
	memcacheFlushNodesMeter metrics.Meter
	memcacheFlushSizeMeter  metrics.Meter

	memcacheGCTimeTimer  metrics.ResettingTimer
	memcacheGCNodesMeter metrics.Meter
	memcacheGCSizeMeter  metrics.Meter

	memcacheCommitTimeTimer  metrics.ResettingTimer
	memcacheCommitNodesMeter metrics.Meter
	memcacheCommitSizeMeter  metrics.Meter
}

func newMetrics(prefix string) *trieDbMetrics {
	return &trieDbMetrics{
		memcacheCleanHitMeter:   metrics.NewRegisteredMeter(prefix+"trie/memcache/clean/hit", nil),
		memcacheCleanMissMeter:  metrics.NewRegisteredMeter(prefix+"trie/memcache/clean/miss", nil),
		memcacheCleanReadMeter:  metrics.NewRegisteredMeter(prefix+"trie/memcache/clean/read", nil),
		memcacheCleanWriteMeter: metrics.NewRegisteredMeter(prefix+"trie/memcache/clean/write", nil),

		memcacheDirtyHitMeter:   metrics.NewRegisteredMeter(prefix+"trie/memcache/dirty/hit", nil),
		memcacheDirtyMissMeter:  metrics.NewRegisteredMeter(prefix+"trie/memcache/dirty/miss", nil),
		memcacheDirtyReadMeter:  metrics.NewRegisteredMeter(prefix+"trie/memcache/dirty/read", nil),
		memcacheDirtyWriteMeter: metrics.NewRegisteredMeter(prefix+"trie/memcache/dirty/write", nil),

		memcacheFlushTimeTimer:  metrics.NewRegisteredResettingTimer(prefix+"trie/memcache/flush/time", nil),
		memcacheFlushNodesMeter: metrics.NewRegisteredMeter(prefix+"trie/memcache/flush/nodes", nil),
		memcacheFlushSizeMeter:  metrics.NewRegisteredMeter(prefix+"trie/memcache/flush/size", nil),

		memcacheGCTimeTimer:  metrics.NewRegisteredResettingTimer(prefix+"trie/memcache/gc/time", nil),
		memcacheGCNodesMeter: metrics.NewRegisteredMeter(prefix+"trie/memcache/gc/nodes", nil),
		memcacheGCSizeMeter:  metrics.NewRegisteredMeter(prefix+"trie/memcache/gc/size", nil),

		memcacheCommitTimeTimer:  metrics.NewRegisteredResettingTimer(prefix+"trie/memcache/commit/time", nil),
		memcacheCommitNodesMeter: metrics.NewRegisteredMeter(prefix+"trie/memcache/commit/nodes", nil),
		memcacheCommitSizeMeter:  metrics.NewRegisteredMeter(prefix+"trie/memcache/commit/size", nil),
	}
}
