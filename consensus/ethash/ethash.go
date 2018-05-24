// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package ethash implements the ethash proof-of-work consensus engine.
package ethash

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
<<<<<<< HEAD
=======
	"runtime"
>>>>>>> core/release/1.8
	"strconv"
	"sync"
	"time"
	"unsafe"

	mmap "github.com/edsrzf/mmap-go"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/log"
<<<<<<< HEAD
	"github.com/ethereum/go-ethereum/rpc"
	metrics "github.com/rcrowley/go-metrics"
=======
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/golang-lru/simplelru"
>>>>>>> core/release/1.8
)

var ErrInvalidDumpMagic = errors.New("invalid dump magic")

var (
	// maxUint256 is a big integer representing 2^256-1
	maxUint256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))

	// sharedEthash is a full instance that can be shared between multiple users.
<<<<<<< HEAD
	sharedEthash = New("", 3, 0, "", 1, 0)
=======
	sharedEthash = New(Config{"", 3, 0, "", 1, 0, ModeNormal})
>>>>>>> core/release/1.8

	// algorithmRevision is the data structure version used for file naming.
	algorithmRevision = 23

	// dumpMagic is a dataset dump header to sanity check a data dump.
	dumpMagic = []uint32{0xbaddcafe, 0xfee1dead}
)

// isLittleEndian returns whether the local system is running in little or big
// endian byte order.
func isLittleEndian() bool {
	n := uint32(0x01020304)
	return *(*byte)(unsafe.Pointer(&n)) == 0x04
}

// memoryMap tries to memory map a file of uint32s for read only access.
func memoryMap(path string) (*os.File, mmap.MMap, []uint32, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, nil, nil, err
	}
	mem, buffer, err := memoryMapFile(file, false)
	if err != nil {
		file.Close()
		return nil, nil, nil, err
	}
	for i, magic := range dumpMagic {
		if buffer[i] != magic {
			mem.Unmap()
			file.Close()
			return nil, nil, nil, ErrInvalidDumpMagic
		}
	}
	return file, mem, buffer[len(dumpMagic):], err
}

// memoryMapFile tries to memory map an already opened file descriptor.
func memoryMapFile(file *os.File, write bool) (mmap.MMap, []uint32, error) {
	// Try to memory map the file
	flag := mmap.RDONLY
	if write {
		flag = mmap.RDWR
	}
	mem, err := mmap.Map(file, flag, 0)
	if err != nil {
		return nil, nil, err
	}
	// Yay, we managed to memory map the file, here be dragons
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&mem))
	header.Len /= 4
	header.Cap /= 4

	return mem, *(*[]uint32)(unsafe.Pointer(&header)), nil
}

// memoryMapAndGenerate tries to memory map a temporary file of uint32s for write
// access, fill it with the data from a generator and then move it into the final
// path requested.
func memoryMapAndGenerate(path string, size uint64, generator func(buffer []uint32)) (*os.File, mmap.MMap, []uint32, error) {
	// Ensure the data folder exists
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, nil, nil, err
	}
	// Create a huge temporary empty file to fill with data
	temp := path + "." + strconv.Itoa(rand.Int())

	dump, err := os.Create(temp)
	if err != nil {
		return nil, nil, nil, err
	}
	if err = dump.Truncate(int64(len(dumpMagic))*4 + int64(size)); err != nil {
		return nil, nil, nil, err
	}
	// Memory map the file for writing and fill it with the generator
	mem, buffer, err := memoryMapFile(dump, true)
	if err != nil {
		dump.Close()
		return nil, nil, nil, err
	}
	copy(buffer, dumpMagic)

	data := buffer[len(dumpMagic):]
	generator(data)

	if err := mem.Unmap(); err != nil {
		return nil, nil, nil, err
	}
	if err := dump.Close(); err != nil {
		return nil, nil, nil, err
	}
	if err := os.Rename(temp, path); err != nil {
		return nil, nil, nil, err
	}
	return memoryMap(path)
}

<<<<<<< HEAD
// cache wraps an ethash cache with some metadata to allow easier concurrent use.
type cache struct {
	epoch uint64 // Epoch for which this cache is relevant

	dump *os.File  // File descriptor of the memory mapped cache
	mmap mmap.MMap // Memory map itself to unmap before releasing

	cache []uint32   // The actual cache data content (may be memory mapped)
	used  time.Time  // Timestamp of the last use for smarter eviction
	once  sync.Once  // Ensures the cache is generated only once
	lock  sync.Mutex // Ensures thread safety for updating the usage time
=======
// lru tracks caches or datasets by their last use time, keeping at most N of them.
type lru struct {
	what string
	new  func(epoch uint64) interface{}
	mu   sync.Mutex
	// Items are kept in a LRU cache, but there is a special case:
	// We always keep an item for (highest seen epoch) + 1 as the 'future item'.
	cache      *simplelru.LRU
	future     uint64
	futureItem interface{}
}

// newlru create a new least-recently-used cache for either the verification caches
// or the mining datasets.
func newlru(what string, maxItems int, new func(epoch uint64) interface{}) *lru {
	if maxItems <= 0 {
		maxItems = 1
	}
	cache, _ := simplelru.NewLRU(maxItems, func(key, value interface{}) {
		log.Trace("Evicted ethash "+what, "epoch", key)
	})
	return &lru{what: what, new: new, cache: cache}
}

// get retrieves or creates an item for the given epoch. The first return value is always
// non-nil. The second return value is non-nil if lru thinks that an item will be useful in
// the near future.
func (lru *lru) get(epoch uint64) (item, future interface{}) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	// Get or create the item for the requested epoch.
	item, ok := lru.cache.Get(epoch)
	if !ok {
		if lru.future > 0 && lru.future == epoch {
			item = lru.futureItem
		} else {
			log.Trace("Requiring new ethash "+lru.what, "epoch", epoch)
			item = lru.new(epoch)
		}
		lru.cache.Add(epoch, item)
	}
	// Update the 'future item' if epoch is larger than previously seen.
	if epoch < maxEpoch-1 && lru.future < epoch+1 {
		log.Trace("Requiring new future ethash "+lru.what, "epoch", epoch+1)
		future = lru.new(epoch + 1)
		lru.future = epoch + 1
		lru.futureItem = future
	}
	return item, future
}

// cache wraps an ethash cache with some metadata to allow easier concurrent use.
type cache struct {
	epoch uint64    // Epoch for which this cache is relevant
	dump  *os.File  // File descriptor of the memory mapped cache
	mmap  mmap.MMap // Memory map itself to unmap before releasing
	cache []uint32  // The actual cache data content (may be memory mapped)
	once  sync.Once // Ensures the cache is generated only once
}

// newCache creates a new ethash verification cache and returns it as a plain Go
// interface to be usable in an LRU cache.
func newCache(epoch uint64) interface{} {
	return &cache{epoch: epoch}
>>>>>>> core/release/1.8
}

// generate ensures that the cache content is generated before use.
func (c *cache) generate(dir string, limit int, test bool) {
	c.once.Do(func() {
<<<<<<< HEAD
		// If we have a testing cache, generate and return
		if test {
			c.cache = make([]uint32, 1024/4)
			generateCache(c.cache, c.epoch, seedHash(c.epoch*epochLength+1))
			return
		}
		// If we don't store anything on disk, generate and return
		size := cacheSize(c.epoch*epochLength + 1)
		seed := seedHash(c.epoch*epochLength + 1)

=======
		size := cacheSize(c.epoch*epochLength + 1)
		seed := seedHash(c.epoch*epochLength + 1)
		if test {
			size = 1024
		}
		// If we don't store anything on disk, generate and return.
>>>>>>> core/release/1.8
		if dir == "" {
			c.cache = make([]uint32, size/4)
			generateCache(c.cache, c.epoch, seed)
			return
		}
		// Disk storage is needed, this will get fancy
		var endian string
		if !isLittleEndian() {
			endian = ".be"
		}
		path := filepath.Join(dir, fmt.Sprintf("cache-R%d-%x%s", algorithmRevision, seed[:8], endian))
		logger := log.New("epoch", c.epoch)

<<<<<<< HEAD
=======
		// We're about to mmap the file, ensure that the mapping is cleaned up when the
		// cache becomes unused.
		runtime.SetFinalizer(c, (*cache).finalizer)

>>>>>>> core/release/1.8
		// Try to load the file from disk and memory map it
		var err error
		c.dump, c.mmap, c.cache, err = memoryMap(path)
		if err == nil {
			logger.Debug("Loaded old ethash cache from disk")
			return
		}
		logger.Debug("Failed to load old ethash cache", "err", err)

		// No previous cache available, create a new cache file to fill
		c.dump, c.mmap, c.cache, err = memoryMapAndGenerate(path, size, func(buffer []uint32) { generateCache(buffer, c.epoch, seed) })
		if err != nil {
			logger.Error("Failed to generate mapped ethash cache", "err", err)

			c.cache = make([]uint32, size/4)
			generateCache(c.cache, c.epoch, seed)
		}
		// Iterate over all previous instances and delete old ones
		for ep := int(c.epoch) - limit; ep >= 0; ep-- {
			seed := seedHash(uint64(ep)*epochLength + 1)
			path := filepath.Join(dir, fmt.Sprintf("cache-R%d-%x%s", algorithmRevision, seed[:8], endian))
			os.Remove(path)
		}
	})
}

<<<<<<< HEAD
// release closes any file handlers and memory maps open.
func (c *cache) release() {
	if c.mmap != nil {
		c.mmap.Unmap()
		c.mmap = nil
	}
	if c.dump != nil {
		c.dump.Close()
		c.dump = nil
=======
// finalizer unmaps the memory and closes the file.
func (c *cache) finalizer() {
	if c.mmap != nil {
		c.mmap.Unmap()
		c.dump.Close()
		c.mmap, c.dump = nil, nil
>>>>>>> core/release/1.8
	}
}

// dataset wraps an ethash dataset with some metadata to allow easier concurrent use.
type dataset struct {
<<<<<<< HEAD
	epoch uint64 // Epoch for which this cache is relevant

	dump *os.File  // File descriptor of the memory mapped cache
	mmap mmap.MMap // Memory map itself to unmap before releasing

	dataset []uint32   // The actual cache data content
	used    time.Time  // Timestamp of the last use for smarter eviction
	once    sync.Once  // Ensures the cache is generated only once
	lock    sync.Mutex // Ensures thread safety for updating the usage time
=======
	epoch   uint64    // Epoch for which this cache is relevant
	dump    *os.File  // File descriptor of the memory mapped cache
	mmap    mmap.MMap // Memory map itself to unmap before releasing
	dataset []uint32  // The actual cache data content
	once    sync.Once // Ensures the cache is generated only once
}

// newDataset creates a new ethash mining dataset and returns it as a plain Go
// interface to be usable in an LRU cache.
func newDataset(epoch uint64) interface{} {
	return &dataset{epoch: epoch}
>>>>>>> core/release/1.8
}

// generate ensures that the dataset content is generated before use.
func (d *dataset) generate(dir string, limit int, test bool) {
	d.once.Do(func() {
<<<<<<< HEAD
		// If we have a testing dataset, generate and return
		if test {
			cache := make([]uint32, 1024/4)
			generateCache(cache, d.epoch, seedHash(d.epoch*epochLength+1))

			d.dataset = make([]uint32, 32*1024/4)
			generateDataset(d.dataset, d.epoch, cache)

			return
		}
		// If we don't store anything on disk, generate and return
		csize := cacheSize(d.epoch*epochLength + 1)
		dsize := datasetSize(d.epoch*epochLength + 1)
		seed := seedHash(d.epoch*epochLength + 1)

=======
		csize := cacheSize(d.epoch*epochLength + 1)
		dsize := datasetSize(d.epoch*epochLength + 1)
		seed := seedHash(d.epoch*epochLength + 1)
		if test {
			csize = 1024
			dsize = 32 * 1024
		}
		// If we don't store anything on disk, generate and return
>>>>>>> core/release/1.8
		if dir == "" {
			cache := make([]uint32, csize/4)
			generateCache(cache, d.epoch, seed)

			d.dataset = make([]uint32, dsize/4)
			generateDataset(d.dataset, d.epoch, cache)
		}
		// Disk storage is needed, this will get fancy
		var endian string
		if !isLittleEndian() {
			endian = ".be"
		}
		path := filepath.Join(dir, fmt.Sprintf("full-R%d-%x%s", algorithmRevision, seed[:8], endian))
		logger := log.New("epoch", d.epoch)

<<<<<<< HEAD
=======
		// We're about to mmap the file, ensure that the mapping is cleaned up when the
		// cache becomes unused.
		runtime.SetFinalizer(d, (*dataset).finalizer)

>>>>>>> core/release/1.8
		// Try to load the file from disk and memory map it
		var err error
		d.dump, d.mmap, d.dataset, err = memoryMap(path)
		if err == nil {
			logger.Debug("Loaded old ethash dataset from disk")
			return
		}
		logger.Debug("Failed to load old ethash dataset", "err", err)

		// No previous dataset available, create a new dataset file to fill
		cache := make([]uint32, csize/4)
		generateCache(cache, d.epoch, seed)

		d.dump, d.mmap, d.dataset, err = memoryMapAndGenerate(path, dsize, func(buffer []uint32) { generateDataset(buffer, d.epoch, cache) })
		if err != nil {
			logger.Error("Failed to generate mapped ethash dataset", "err", err)

			d.dataset = make([]uint32, dsize/2)
			generateDataset(d.dataset, d.epoch, cache)
		}
		// Iterate over all previous instances and delete old ones
		for ep := int(d.epoch) - limit; ep >= 0; ep-- {
			seed := seedHash(uint64(ep)*epochLength + 1)
			path := filepath.Join(dir, fmt.Sprintf("full-R%d-%x%s", algorithmRevision, seed[:8], endian))
			os.Remove(path)
		}
	})
}

<<<<<<< HEAD
// release closes any file handlers and memory maps open.
func (d *dataset) release() {
	if d.mmap != nil {
		d.mmap.Unmap()
		d.mmap = nil
	}
	if d.dump != nil {
		d.dump.Close()
		d.dump = nil
=======
// finalizer closes any file handlers and memory maps open.
func (d *dataset) finalizer() {
	if d.mmap != nil {
		d.mmap.Unmap()
		d.dump.Close()
		d.mmap, d.dump = nil, nil
>>>>>>> core/release/1.8
	}
}

// MakeCache generates a new ethash cache and optionally stores it to disk.
func MakeCache(block uint64, dir string) {
	c := cache{epoch: block / epochLength}
	c.generate(dir, math.MaxInt32, false)
<<<<<<< HEAD
	c.release()
=======
>>>>>>> core/release/1.8
}

// MakeDataset generates a new ethash dataset and optionally stores it to disk.
func MakeDataset(block uint64, dir string) {
	d := dataset{epoch: block / epochLength}
	d.generate(dir, math.MaxInt32, false)
<<<<<<< HEAD
	d.release()
=======
}

// Mode defines the type and amount of PoW verification an ethash engine makes.
type Mode uint

const (
	ModeNormal Mode = iota
	ModeShared
	ModeTest
	ModeFake
	ModeFullFake
)

// Config are the configuration parameters of the ethash.
type Config struct {
	CacheDir       string
	CachesInMem    int
	CachesOnDisk   int
	DatasetDir     string
	DatasetsInMem  int
	DatasetsOnDisk int
	PowMode        Mode
>>>>>>> core/release/1.8
}

// Ethash is a consensus engine based on proot-of-work implementing the ethash
// algorithm.
type Ethash struct {
<<<<<<< HEAD
	cachedir     string // Data directory to store the verification caches
	cachesinmem  int    // Number of caches to keep in memory
	cachesondisk int    // Number of caches to keep on disk
	dagdir       string // Data directory to store full mining datasets
	dagsinmem    int    // Number of mining datasets to keep in memory
	dagsondisk   int    // Number of mining datasets to keep on disk

	caches   map[uint64]*cache   // In memory caches to avoid regenerating too often
	fcache   *cache              // Pre-generated cache for the estimated future epoch
	datasets map[uint64]*dataset // In memory datasets to avoid regenerating too often
	fdataset *dataset            // Pre-generated dataset for the estimated future epoch
=======
	config Config

	caches   *lru // In memory caches to avoid regenerating too often
	datasets *lru // In memory datasets to avoid regenerating too often
>>>>>>> core/release/1.8

	// Mining related fields
	rand     *rand.Rand    // Properly seeded random source for nonces
	threads  int           // Number of threads to mine on if mining
	update   chan struct{} // Notification channel to update mining parameters
	hashrate metrics.Meter // Meter tracking the average hashrate

	// The fields below are hooks for testing
<<<<<<< HEAD
	tester    bool          // Flag whether to use a smaller test dataset
	shared    *Ethash       // Shared PoW verifier to avoid cache regeneration
	fakeMode  bool          // Flag whether to disable PoW checking
	fakeFull  bool          // Flag whether to disable all consensus rules
=======
	shared    *Ethash       // Shared PoW verifier to avoid cache regeneration
>>>>>>> core/release/1.8
	fakeFail  uint64        // Block number which fails PoW check even in fake mode
	fakeDelay time.Duration // Time delay to sleep for before returning from verify

	lock sync.Mutex // Ensures thread safety for the in-memory caches and mining fields
}

// New creates a full sized ethash PoW scheme.
<<<<<<< HEAD
func New(cachedir string, cachesinmem, cachesondisk int, dagdir string, dagsinmem, dagsondisk int) *Ethash {
	if cachesinmem <= 0 {
		log.Warn("One ethash cache must always be in memory", "requested", cachesinmem)
		cachesinmem = 1
	}
	if cachedir != "" && cachesondisk > 0 {
		log.Info("Disk storage enabled for ethash caches", "dir", cachedir, "count", cachesondisk)
	}
	if dagdir != "" && dagsondisk > 0 {
		log.Info("Disk storage enabled for ethash DAGs", "dir", dagdir, "count", dagsondisk)
	}
	return &Ethash{
		cachedir:     cachedir,
		cachesinmem:  cachesinmem,
		cachesondisk: cachesondisk,
		dagdir:       dagdir,
		dagsinmem:    dagsinmem,
		dagsondisk:   dagsondisk,
		caches:       make(map[uint64]*cache),
		datasets:     make(map[uint64]*dataset),
		update:       make(chan struct{}),
		hashrate:     metrics.NewMeter(),
=======
func New(config Config) *Ethash {
	if config.CachesInMem <= 0 {
		log.Warn("One ethash cache must always be in memory", "requested", config.CachesInMem)
		config.CachesInMem = 1
	}
	if config.CacheDir != "" && config.CachesOnDisk > 0 {
		log.Info("Disk storage enabled for ethash caches", "dir", config.CacheDir, "count", config.CachesOnDisk)
	}
	if config.DatasetDir != "" && config.DatasetsOnDisk > 0 {
		log.Info("Disk storage enabled for ethash DAGs", "dir", config.DatasetDir, "count", config.DatasetsOnDisk)
	}
	return &Ethash{
		config:   config,
		caches:   newlru("cache", config.CachesInMem, newCache),
		datasets: newlru("dataset", config.DatasetsInMem, newDataset),
		update:   make(chan struct{}),
		hashrate: metrics.NewMeter(),
>>>>>>> core/release/1.8
	}
}

// NewTester creates a small sized ethash PoW scheme useful only for testing
// purposes.
func NewTester() *Ethash {
<<<<<<< HEAD
	return &Ethash{
		cachesinmem: 1,
		caches:      make(map[uint64]*cache),
		datasets:    make(map[uint64]*dataset),
		tester:      true,
		update:      make(chan struct{}),
		hashrate:    metrics.NewMeter(),
	}
=======
	return New(Config{CachesInMem: 1, PowMode: ModeTest})
>>>>>>> core/release/1.8
}

// NewFaker creates a ethash consensus engine with a fake PoW scheme that accepts
// all blocks' seal as valid, though they still have to conform to the Ethereum
// consensus rules.
func NewFaker() *Ethash {
<<<<<<< HEAD
	return &Ethash{fakeMode: true}
=======
	return &Ethash{
		config: Config{
			PowMode: ModeFake,
		},
	}
>>>>>>> core/release/1.8
}

// NewFakeFailer creates a ethash consensus engine with a fake PoW scheme that
// accepts all blocks as valid apart from the single one specified, though they
// still have to conform to the Ethereum consensus rules.
func NewFakeFailer(fail uint64) *Ethash {
<<<<<<< HEAD
	return &Ethash{fakeMode: true, fakeFail: fail}
=======
	return &Ethash{
		config: Config{
			PowMode: ModeFake,
		},
		fakeFail: fail,
	}
>>>>>>> core/release/1.8
}

// NewFakeDelayer creates a ethash consensus engine with a fake PoW scheme that
// accepts all blocks as valid, but delays verifications by some time, though
// they still have to conform to the Ethereum consensus rules.
func NewFakeDelayer(delay time.Duration) *Ethash {
<<<<<<< HEAD
	return &Ethash{fakeMode: true, fakeDelay: delay}
=======
	return &Ethash{
		config: Config{
			PowMode: ModeFake,
		},
		fakeDelay: delay,
	}
>>>>>>> core/release/1.8
}

// NewFullFaker creates an ethash consensus engine with a full fake scheme that
// accepts all blocks as valid, without checking any consensus rules whatsoever.
func NewFullFaker() *Ethash {
<<<<<<< HEAD
	return &Ethash{fakeMode: true, fakeFull: true}
=======
	return &Ethash{
		config: Config{
			PowMode: ModeFullFake,
		},
	}
>>>>>>> core/release/1.8
}

// NewShared creates a full sized ethash PoW shared between all requesters running
// in the same process.
func NewShared() *Ethash {
	return &Ethash{shared: sharedEthash}
}

// cache tries to retrieve a verification cache for the specified block number
// by first checking against a list of in-memory caches, then against caches
// stored on disk, and finally generating one if none can be found.
<<<<<<< HEAD
func (ethash *Ethash) cache(block uint64) []uint32 {
	epoch := block / epochLength

	// If we have a PoW for that epoch, use that
	ethash.lock.Lock()

	current, future := ethash.caches[epoch], (*cache)(nil)
	if current == nil {
		// No in-memory cache, evict the oldest if the cache limit was reached
		for len(ethash.caches) > 0 && len(ethash.caches) >= ethash.cachesinmem {
			var evict *cache
			for _, cache := range ethash.caches {
				if evict == nil || evict.used.After(cache.used) {
					evict = cache
				}
			}
			delete(ethash.caches, evict.epoch)
			evict.release()

			log.Trace("Evicted ethash cache", "epoch", evict.epoch, "used", evict.used)
		}
		// If we have the new cache pre-generated, use that, otherwise create a new one
		if ethash.fcache != nil && ethash.fcache.epoch == epoch {
			log.Trace("Using pre-generated cache", "epoch", epoch)
			current, ethash.fcache = ethash.fcache, nil
		} else {
			log.Trace("Requiring new ethash cache", "epoch", epoch)
			current = &cache{epoch: epoch}
		}
		ethash.caches[epoch] = current

		// If we just used up the future cache, or need a refresh, regenerate
		if ethash.fcache == nil || ethash.fcache.epoch <= epoch {
			if ethash.fcache != nil {
				ethash.fcache.release()
			}
			log.Trace("Requiring new future ethash cache", "epoch", epoch+1)
			future = &cache{epoch: epoch + 1}
			ethash.fcache = future
		}
		// New current cache, set its initial timestamp
		current.used = time.Now()
	}
	ethash.lock.Unlock()

	// Wait for generation finish, bump the timestamp and finalize the cache
	current.generate(ethash.cachedir, ethash.cachesondisk, ethash.tester)

	current.lock.Lock()
	current.used = time.Now()
	current.lock.Unlock()

	// If we exhausted the future cache, now's a good time to regenerate it
	if future != nil {
		go future.generate(ethash.cachedir, ethash.cachesondisk, ethash.tester)
	}
	return current.cache
=======
func (ethash *Ethash) cache(block uint64) *cache {
	epoch := block / epochLength
	currentI, futureI := ethash.caches.get(epoch)
	current := currentI.(*cache)

	// Wait for generation finish.
	current.generate(ethash.config.CacheDir, ethash.config.CachesOnDisk, ethash.config.PowMode == ModeTest)

	// If we need a new future cache, now's a good time to regenerate it.
	if futureI != nil {
		future := futureI.(*cache)
		go future.generate(ethash.config.CacheDir, ethash.config.CachesOnDisk, ethash.config.PowMode == ModeTest)
	}
	return current
>>>>>>> core/release/1.8
}

// dataset tries to retrieve a mining dataset for the specified block number
// by first checking against a list of in-memory datasets, then against DAGs
// stored on disk, and finally generating one if none can be found.
<<<<<<< HEAD
func (ethash *Ethash) dataset(block uint64) []uint32 {
	epoch := block / epochLength

	// If we have a PoW for that epoch, use that
	ethash.lock.Lock()

	current, future := ethash.datasets[epoch], (*dataset)(nil)
	if current == nil {
		// No in-memory dataset, evict the oldest if the dataset limit was reached
		for len(ethash.datasets) > 0 && len(ethash.datasets) >= ethash.dagsinmem {
			var evict *dataset
			for _, dataset := range ethash.datasets {
				if evict == nil || evict.used.After(dataset.used) {
					evict = dataset
				}
			}
			delete(ethash.datasets, evict.epoch)
			evict.release()

			log.Trace("Evicted ethash dataset", "epoch", evict.epoch, "used", evict.used)
		}
		// If we have the new cache pre-generated, use that, otherwise create a new one
		if ethash.fdataset != nil && ethash.fdataset.epoch == epoch {
			log.Trace("Using pre-generated dataset", "epoch", epoch)
			current = &dataset{epoch: ethash.fdataset.epoch} // Reload from disk
			ethash.fdataset = nil
		} else {
			log.Trace("Requiring new ethash dataset", "epoch", epoch)
			current = &dataset{epoch: epoch}
		}
		ethash.datasets[epoch] = current

		// If we just used up the future dataset, or need a refresh, regenerate
		if ethash.fdataset == nil || ethash.fdataset.epoch <= epoch {
			if ethash.fdataset != nil {
				ethash.fdataset.release()
			}
			log.Trace("Requiring new future ethash dataset", "epoch", epoch+1)
			future = &dataset{epoch: epoch + 1}
			ethash.fdataset = future
		}
		// New current dataset, set its initial timestamp
		current.used = time.Now()
	}
	ethash.lock.Unlock()

	// Wait for generation finish, bump the timestamp and finalize the cache
	current.generate(ethash.dagdir, ethash.dagsondisk, ethash.tester)

	current.lock.Lock()
	current.used = time.Now()
	current.lock.Unlock()

	// If we exhausted the future dataset, now's a good time to regenerate it
	if future != nil {
		go future.generate(ethash.dagdir, ethash.dagsondisk, ethash.tester)
	}
	return current.dataset
=======
func (ethash *Ethash) dataset(block uint64) *dataset {
	epoch := block / epochLength
	currentI, futureI := ethash.datasets.get(epoch)
	current := currentI.(*dataset)

	// Wait for generation finish.
	current.generate(ethash.config.DatasetDir, ethash.config.DatasetsOnDisk, ethash.config.PowMode == ModeTest)

	// If we need a new future dataset, now's a good time to regenerate it.
	if futureI != nil {
		future := futureI.(*dataset)
		go future.generate(ethash.config.DatasetDir, ethash.config.DatasetsOnDisk, ethash.config.PowMode == ModeTest)
	}

	return current
>>>>>>> core/release/1.8
}

// Threads returns the number of mining threads currently enabled. This doesn't
// necessarily mean that mining is running!
func (ethash *Ethash) Threads() int {
	ethash.lock.Lock()
	defer ethash.lock.Unlock()

	return ethash.threads
}

// SetThreads updates the number of mining threads currently enabled. Calling
// this method does not start mining, only sets the thread count. If zero is
// specified, the miner will use all cores of the machine. Setting a thread
// count below zero is allowed and will cause the miner to idle, without any
// work being done.
func (ethash *Ethash) SetThreads(threads int) {
	ethash.lock.Lock()
	defer ethash.lock.Unlock()

	// If we're running a shared PoW, set the thread count on that instead
	if ethash.shared != nil {
		ethash.shared.SetThreads(threads)
		return
	}
	// Update the threads and ping any running seal to pull in any changes
	ethash.threads = threads
	select {
	case ethash.update <- struct{}{}:
	default:
	}
}

// Hashrate implements PoW, returning the measured rate of the search invocations
// per second over the last minute.
func (ethash *Ethash) Hashrate() float64 {
	return ethash.hashrate.Rate1()
}

// APIs implements consensus.Engine, returning the user facing RPC APIs. Currently
// that is empty.
func (ethash *Ethash) APIs(chain consensus.ChainReader) []rpc.API {
	return nil
}

// SeedHash is the seed to use for generating a verification cache and the mining
// dataset.
func SeedHash(block uint64) []byte {
	return seedHash(block)
}
<<<<<<< HEAD

// Protocol implements consensus.Engine.Protocol
func (ethash *Ethash) Protocol() consensus.Protocol {
	return consensus.EthProtocol
}
=======
>>>>>>> core/release/1.8
