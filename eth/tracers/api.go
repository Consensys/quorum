// Copyright 2021 The go-ethereum Authors
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

package tracers

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	// defaultTraceTimeout is the amount of time a single transaction can execute
	// by default before being forcefully aborted.
	defaultTraceTimeout = 5 * time.Second

	// defaultTraceReexec is the number of blocks the tracer is willing to go back
	// and reexecute to produce missing historical state necessary to run a specific
	// trace.
	defaultTraceReexec = uint64(128)
)

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
type Backend interface {
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error)
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error)
	GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error)
	RPCGasCap() uint64
	ChainConfig() *params.ChainConfig
	Engine() consensus.Engine
	ChainDb() ethdb.Database
	StateAtBlock(ctx context.Context, block *types.Block, reexec uint64, base *state.StateDB, checkLive bool) (*state.StateDB, mps.PrivateStateRepository, error)
	StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (core.Message, vm.BlockContext, *state.StateDB, *state.StateDB, mps.PrivateStateRepository, error)

	// Quorum
	GetBlockchain() *core.BlockChain
}

// API is the collection of tracing APIs exposed over the private debugging endpoint.
type API struct {
	backend Backend
}

// NewAPI creates a new API definition for the tracing methods of the Ethereum service.
func NewAPI(backend Backend) *API {
	return &API{backend: backend}
}

type chainContext struct {
	api *API
	ctx context.Context
}

var _ core.ChainContext = &chainContext{}

func (context *chainContext) Engine() consensus.Engine {
	return context.api.backend.Engine()
}

func (context *chainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	header, err := context.api.backend.HeaderByNumber(context.ctx, rpc.BlockNumber(number))
	if err != nil {
		return nil
	}
	if header.Hash() == hash {
		return header
	}
	header, err = context.api.backend.HeaderByHash(context.ctx, hash)
	if err != nil {
		return nil
	}
	return header
}

// Quorum

func (context *chainContext) Config() *params.ChainConfig {
	return context.api.backend.ChainConfig()
}

func (context *chainContext) QuorumConfig() *core.QuorumChainConfig {
	return &core.QuorumChainConfig{}
}

func (context *chainContext) PrivateStateManager() mps.PrivateStateManager {
	return context.api.backend.GetBlockchain().PrivateStateManager()
}

func (context *chainContext) CheckAndSetPrivateState(txLogs []*types.Log, privateState *state.StateDB, psi types.PrivateStateIdentifier) {
}

// End Quorum

// chainContext construts the context reader which is used by the evm for reading
// the necessary chain context.
func (api *API) chainContext(ctx context.Context) core.ChainContext {
	return &chainContext{api: api, ctx: ctx}
}

// blockByNumber is the wrapper of the chain access function offered by the backend.
// It will return an error if the block is not found.
func (api *API) blockByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Block, error) {
	block, err := api.backend.BlockByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("block #%d not found", number)
	}
	return block, nil
}

// blockByHash is the wrapper of the chain access function offered by the backend.
// It will return an error if the block is not found.
func (api *API) blockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	block, err := api.backend.BlockByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("block %s not found", hash.Hex())
	}
	return block, nil
}

// blockByNumberAndHash is the wrapper of the chain access function offered by
// the backend. It will return an error if the block is not found.
//
// Note this function is friendly for the light client which can only retrieve the
// historical(before the CHT) header/block by number.
func (api *API) blockByNumberAndHash(ctx context.Context, number rpc.BlockNumber, hash common.Hash) (*types.Block, error) {
	block, err := api.blockByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	if block.Hash() == hash {
		return block, nil
	}
	return api.blockByHash(ctx, hash)
}

// TraceConfig holds extra parameters to trace functions.
type TraceConfig struct {
	*vm.LogConfig
	Tracer  *string
	Timeout *string
	Reexec  *uint64
}

// TraceCallConfig is the config for traceCall API. It holds one more
// field to override the state for tracing.
type TraceCallConfig struct {
	*vm.LogConfig
	Tracer         *string
	Timeout        *string
	Reexec         *uint64
	StateOverrides *ethapi.StateOverride
}

// StdTraceConfig holds extra parameters to standard-json trace functions.
type StdTraceConfig struct {
	vm.LogConfig
	Reexec *uint64
	TxHash common.Hash
}

// txTraceContext is the contextual infos about a transaction before it gets run.
type txTraceContext struct {
	index int         // Index of the transaction within the block
	hash  common.Hash // Hash of the transaction
	block common.Hash // Hash of the block containing the transaction
	// Quorum
	tx *types.Transaction // Transaction is needed
}

// txTraceResult is the result of a single transaction trace.
type txTraceResult struct {
	Result interface{} `json:"result,omitempty"` // Trace results produced by the tracer
	Error  string      `json:"error,omitempty"`  // Trace failure produced by the tracer
}

// blockTraceTask represents a single block trace task when an entire chain is
// being traced.
type blockTraceTask struct {
	statedb *state.StateDB   // Intermediate state prepped for tracing
	block   *types.Block     // Block to trace the transactions from
	rootref common.Hash      // Trie root reference held for this task
	results []*txTraceResult // Trace results procudes by the task
	// Quorum
	privateStateDb   *state.StateDB
	privateStateRepo mps.PrivateStateRepository
}

// blockTraceResult represets the results of tracing a single block when an entire
// chain is being traced.
type blockTraceResult struct {
	Block  hexutil.Uint64   `json:"block"`  // Block number corresponding to this trace
	Hash   common.Hash      `json:"hash"`   // Block hash corresponding to this trace
	Traces []*txTraceResult `json:"traces"` // Trace results produced by the task
}

// txTraceTask represents a single transaction trace task when an entire block
// is being traced.
type txTraceTask struct {
	statedb *state.StateDB // Intermediate state prepped for tracing
	index   int            // Transaction offset in the block
	// Quorum
	privateStateDb   *state.StateDB
	privateStateRepo mps.PrivateStateRepository
}

// TraceChain returns the structured logs created during the execution of EVM
// between two blocks (excluding start) and returns them as a JSON object.
func (api *API) TraceChain(ctx context.Context, start, end rpc.BlockNumber, config *TraceConfig) (*rpc.Subscription, error) { // Fetch the block interval that we want to trace
	from, err := api.blockByNumber(ctx, start)
	if err != nil {
		return nil, err
	}
	to, err := api.blockByNumber(ctx, end)
	if err != nil {
		return nil, err
	}
	if from.Number().Cmp(to.Number()) >= 0 {
		return nil, fmt.Errorf("end block (#%d) needs to come after start block (#%d)", end, start)
	}
	return api.traceChain(ctx, from, to, config)
}

// traceChain configures a new tracer according to the provided configuration, and
// executes all the transactions contained within. The return value will be one item
// per transaction, dependent on the requested tracer.
func (api *API) traceChain(ctx context.Context, start, end *types.Block, config *TraceConfig) (*rpc.Subscription, error) {
	// Tracing a chain is a **long** operation, only do with subscriptions
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}
	sub := notifier.CreateSubscription()

	// Prepare all the states for tracing. Note this procedure can take very
	// long time. Timeout mechanism is necessary.
	reexec := defaultTraceReexec
	if config != nil && config.Reexec != nil {
		reexec = *config.Reexec
	}
	// Quorum
	psm, err := api.chainContext(ctx).PrivateStateManager().ResolveForUserContext(ctx)
	if err != nil {
		return nil, err
	}
	// End Quorum
	blocks := int(end.NumberU64() - start.NumberU64())
	threads := runtime.NumCPU()
	if threads > blocks {
		threads = blocks
	}
	var (
		pend     = new(sync.WaitGroup)
		tasks    = make(chan *blockTraceTask, threads)
		results  = make(chan *blockTraceTask, threads)
		localctx = context.Background()
	)
	for th := 0; th < threads; th++ {
		pend.Add(1)
		go func() {
			defer pend.Done()

			// Fetch and execute the next block trace tasks
			for task := range tasks {
				signer := types.MakeSigner(api.backend.ChainConfig(), task.block.Number())
				blockCtx := core.NewEVMBlockContext(task.block.Header(), api.chainContext(localctx), nil)
				// Trace all the transactions contained within
				for i, tx := range task.block.Transactions() {
					msg, _ := tx.AsMessage(signer)
					msg = api.clearMessageDataIfNonParty(msg, psm) // Quorum
					txctx := &txTraceContext{
						index: i,
						hash:  tx.Hash(),
						block: task.block.Hash(),
						tx:    tx,
					}
					res, err := api.traceTx(localctx, msg, txctx, blockCtx, task.statedb, task.privateStateDb, task.privateStateRepo, config)
					if err != nil {
						task.results[i] = &txTraceResult{Error: err.Error()}
						log.Warn("Tracing failed", "hash", tx.Hash(), "block", task.block.NumberU64(), "err", err)
						break
					}
					// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect
					eip158 := api.backend.ChainConfig().IsEIP158(task.block.Number())
					task.statedb.Finalise(eip158)
					task.privateStateDb.Finalise(eip158)
					task.results[i] = &txTraceResult{Result: res}
				}
				// Stream the result back to the user or abort on teardown
				select {
				case results <- task:
				case <-notifier.Closed():
					return
				}
			}
		}()
	}
	// Start a goroutine to feed all the blocks into the tracers
	begin := time.Now()

	go func() {
		var (
			logged  time.Time
			number  uint64
			traced  uint64
			failed  error
			parent  common.Hash
			statedb *state.StateDB
			// Quorum
			privateStateRepo mps.PrivateStateRepository
			privateState     *state.StateDB
		)
		// Ensure everything is properly cleaned up on any exit path
		defer func() {
			close(tasks)
			pend.Wait()

			switch {
			case failed != nil:
				log.Warn("Chain tracing failed", "start", start.NumberU64(), "end", end.NumberU64(), "transactions", traced, "elapsed", time.Since(begin), "err", failed)
			case number < end.NumberU64():
				log.Warn("Chain tracing aborted", "start", start.NumberU64(), "end", end.NumberU64(), "abort", number, "transactions", traced, "elapsed", time.Since(begin))
			default:
				log.Info("Chain tracing finished", "start", start.NumberU64(), "end", end.NumberU64(), "transactions", traced, "elapsed", time.Since(begin))
			}
			close(results)
		}()
		// Feed all the blocks both into the tracer, as well as fast process concurrently
		for number = start.NumberU64(); number < end.NumberU64(); number++ {
			// Stop tracing if interruption was requested
			select {
			case <-notifier.Closed():
				return
			default:
			}
			// Print progress logs if long enough time elapsed
			if time.Since(logged) > 8*time.Second {
				logged = time.Now()
				log.Info("Tracing chain segment", "start", start.NumberU64(), "end", end.NumberU64(), "current", number, "transactions", traced, "elapsed", time.Since(begin))
			}
			// Retrieve the parent state to trace on top
			block, err := api.blockByNumber(localctx, rpc.BlockNumber(number))
			if err != nil {
				failed = err
				break
			}
			// Prepare the statedb for tracing. Don't use the live database for
			// tracing to avoid persisting state junks into the database.
			statedb, privateStateRepo, err = api.backend.StateAtBlock(localctx, block, reexec, statedb, false)
			if err != nil {
				failed = err
				break
			}
			if statedb.Database().TrieDB() != nil {
				// Hold the reference for tracer, will be released at the final stage
				statedb.Database().TrieDB().Reference(block.Root(), common.Hash{})

				// Release the parent state because it's already held by the tracer
				if parent != (common.Hash{}) {
					statedb.Database().TrieDB().Dereference(parent)
				}
			}
			parent = block.Root()

			privateState, err = privateStateRepo.StatePSI(psm.ID)
			if err != nil {
				failed = err
				break
			}

			next, err := api.blockByNumber(localctx, rpc.BlockNumber(number+1))
			if err != nil {
				failed = err
				break
			}
			// Send the block over to the concurrent tracers (if not in the fast-forward phase)
			txs := next.Transactions()
			select {
			case tasks <- &blockTraceTask{
				statedb: statedb.Copy(),
				block:   next,
				rootref: block.Root(),
				results: make([]*txTraceResult, len(txs)),
				// Quorum
				privateStateDb:   privateState.Copy(),
				privateStateRepo: privateStateRepo,
			}:
			case <-notifier.Closed():
				return
			}
			traced += uint64(len(txs))
		}
	}()

	// Keep reading the trace results and stream the to the user
	go func() {
		var (
			done = make(map[uint64]*blockTraceResult)
			next = start.NumberU64() + 1
		)
		for res := range results {
			// Queue up next received result
			result := &blockTraceResult{
				Block:  hexutil.Uint64(res.block.NumberU64()),
				Hash:   res.block.Hash(),
				Traces: res.results,
			}
			done[uint64(result.Block)] = result

			// Dereference any parent tries held in memory by this task
			if res.statedb.Database().TrieDB() != nil {
				res.statedb.Database().TrieDB().Dereference(res.rootref)
			}
			if res.privateStateDb != nil && res.privateStateDb != res.statedb && res.privateStateDb.Database().TrieDB() != nil {
				res.privateStateDb.Database().TrieDB().Dereference(res.rootref)
			}
			// Stream completed traces to the user, aborting on the first error
			for result, ok := done[next]; ok; result, ok = done[next] {
				if len(result.Traces) > 0 || next == end.NumberU64() {
					notifier.Notify(sub.ID, result)
				}
				delete(done, next)
				next++
			}
		}
	}()
	return sub, nil
}

// TraceBlockByNumber returns the structured logs created during the execution of
// EVM and returns them as a JSON object.
func (api *API) TraceBlockByNumber(ctx context.Context, number rpc.BlockNumber, config *TraceConfig) ([]*txTraceResult, error) {
	block, err := api.blockByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	return api.traceBlock(ctx, block, config)
}

// TraceBlockByHash returns the structured logs created during the execution of
// EVM and returns them as a JSON object.
func (api *API) TraceBlockByHash(ctx context.Context, hash common.Hash, config *TraceConfig) ([]*txTraceResult, error) {
	block, err := api.blockByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	return api.traceBlock(ctx, block, config)
}

// TraceBlock returns the structured logs created during the execution of EVM
// and returns them as a JSON object.
func (api *API) TraceBlock(ctx context.Context, blob []byte, config *TraceConfig) ([]*txTraceResult, error) {
	block := new(types.Block)
	if err := rlp.Decode(bytes.NewReader(blob), block); err != nil {
		return nil, fmt.Errorf("could not decode block: %v", err)
	}
	return api.traceBlock(ctx, block, config)
}

// TraceBlockFromFile returns the structured logs created during the execution of
// EVM and returns them as a JSON object.
func (api *API) TraceBlockFromFile(ctx context.Context, file string, config *TraceConfig) ([]*txTraceResult, error) {
	blob, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}
	return api.TraceBlock(ctx, blob, config)
}

// TraceBadBlock returns the structured logs created during the execution of
// EVM against a block pulled from the pool of bad ones and returns them as a JSON
// object.
func (api *API) TraceBadBlock(ctx context.Context, hash common.Hash, config *TraceConfig) ([]*txTraceResult, error) {
	for _, block := range rawdb.ReadAllBadBlocks(api.backend.ChainDb()) {
		if block.Hash() == hash {
			return api.traceBlock(ctx, block, config)
		}
	}
	return nil, fmt.Errorf("bad block %#x not found", hash)
}

// StandardTraceBlockToFile dumps the structured logs created during the
// execution of EVM to the local file system and returns a list of files
// to the caller.
func (api *API) StandardTraceBlockToFile(ctx context.Context, hash common.Hash, config *StdTraceConfig) ([]string, error) {
	block, err := api.blockByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	return api.standardTraceBlockToFile(ctx, block, config)
}

// StandardTraceBadBlockToFile dumps the structured logs created during the
// execution of EVM against a block pulled from the pool of bad ones to the
// local file system and returns a list of files to the caller.
func (api *API) StandardTraceBadBlockToFile(ctx context.Context, hash common.Hash, config *StdTraceConfig) ([]string, error) {
	for _, block := range rawdb.ReadAllBadBlocks(api.backend.ChainDb()) {
		if block.Hash() == hash {
			return api.standardTraceBlockToFile(ctx, block, config)
		}
	}
	return nil, fmt.Errorf("bad block %#x not found", hash)
}

// traceBlock configures a new tracer according to the provided configuration, and
// executes all the transactions contained within. The return value will be one item
// per transaction, dependent on the requestd tracer.
func (api *API) traceBlock(ctx context.Context, block *types.Block, config *TraceConfig) ([]*txTraceResult, error) {
	if block.NumberU64() == 0 {
		return nil, errors.New("genesis is not traceable")
	}
	parent, err := api.blockByNumberAndHash(ctx, rpc.BlockNumber(block.NumberU64()-1), block.ParentHash())
	if err != nil {
		return nil, err
	}
	reexec := defaultTraceReexec
	if config != nil && config.Reexec != nil {
		reexec = *config.Reexec
	}
	statedb, privateStateRepo, err := api.backend.StateAtBlock(ctx, parent, reexec, nil, true)
	if err != nil {
		return nil, err
	}

	// Quorum
	psm, err := api.chainContext(ctx).PrivateStateManager().ResolveForUserContext(ctx)
	if err != nil {
		return nil, err
	}
	privateStateDb, err := privateStateRepo.StatePSI(psm.ID)
	if err != nil {
		return nil, err
	}
	// End Quorum

	// Execute all the transaction contained within the block concurrently
	var (
		signer  = types.MakeSigner(api.backend.ChainConfig(), block.Number())
		txs     = block.Transactions()
		results = make([]*txTraceResult, len(txs))

		pend = new(sync.WaitGroup)
		jobs = make(chan *txTraceTask, len(txs))
	)
	threads := runtime.NumCPU()
	if threads > len(txs) {
		threads = len(txs)
	}
	blockCtx := core.NewEVMBlockContext(block.Header(), api.chainContext(ctx), nil)
	blockHash := block.Hash()
	for th := 0; th < threads; th++ {
		pend.Add(1)
		go func() {
			defer pend.Done()
			// Fetch and execute the next transaction trace tasks
			for task := range jobs {
				tx := txs[task.index]
				msg, _ := tx.AsMessage(signer)
				msg = api.clearMessageDataIfNonParty(msg, psm)
				txctx := &txTraceContext{
					index: task.index,
					hash:  tx.Hash(),
					block: blockHash,
					tx:    tx,
				}
				res, err := api.traceTx(ctx, msg, txctx, blockCtx, task.statedb, task.privateStateDb, task.privateStateRepo, config)
				if err != nil {
					results[task.index] = &txTraceResult{Error: err.Error()}
					continue
				}
				results[task.index] = &txTraceResult{Result: res}
			}
		}()
	}
	// Feed the transactions into the tracers and return
	var failed error
	for i, tx := range txs {
		// Send the trace task over for execution
		jobs <- &txTraceTask{
			statedb: statedb.Copy(),
			index:   i,
			// Quorum
			privateStateDb:   privateStateDb.Copy(),
			privateStateRepo: privateStateRepo,
		}

		// Generate the next state snapshot fast without tracing
		msg, _ := tx.AsMessage(signer)
		privateStateDbToUse := core.PrivateStateDBForTxn(api.chainContext(ctx).Config().IsQuorum, tx, statedb, privateStateDb)
		statedb.Prepare(tx.Hash(), block.Hash(), i)
		privateStateDbToUse.Prepare(tx.Hash(), block.Hash(), i)
		vmenv := vm.NewEVM(blockCtx, core.NewEVMTxContext(msg), statedb, privateStateDbToUse, api.backend.ChainConfig(), vm.Config{})
		vmenv.SetCurrentTX(tx)
		vmenv.InnerApply = func(innerTx *types.Transaction) error {
			return applyInnerTransaction(api.backend.GetBlockchain(), statedb, privateStateDbToUse, block.Header(), tx, vm.Config{}, privateStateRepo.IsMPS(), privateStateRepo, vmenv, innerTx, i)
		}
		if _, err := core.ApplyMessage(vmenv, msg, new(core.GasPool).AddGas(msg.Gas())); err != nil {
			failed = err
			break
		}
		// Finalize the state so any modifications are written to the trie
		// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect
		eip158 := vmenv.ChainConfig().IsEIP158(block.Number())
		statedb.Finalise(eip158)
		privateStateDb.Finalise(eip158)
	}
	close(jobs)
	pend.Wait()

	// If execution failed in between, abort
	if failed != nil {
		return nil, failed
	}
	return results, nil
}

// standardTraceBlockToFile configures a new tracer which uses standard JSON output,
// and traces either a full block or an individual transaction. The return value will
// be one filename per transaction traced.
func (api *API) standardTraceBlockToFile(ctx context.Context, block *types.Block, config *StdTraceConfig) ([]string, error) {
	// If we're tracing a single transaction, make sure it's present
	if config != nil && config.TxHash != (common.Hash{}) {
		if !containsTx(block, config.TxHash) {
			return nil, fmt.Errorf("transaction %#x not found in block", config.TxHash)
		}
	}
	if block.NumberU64() == 0 {
		return nil, errors.New("genesis is not traceable")
	}
	parent, err := api.blockByNumberAndHash(ctx, rpc.BlockNumber(block.NumberU64()-1), block.ParentHash())
	if err != nil {
		return nil, err
	}
	reexec := defaultTraceReexec
	if config != nil && config.Reexec != nil {
		reexec = *config.Reexec
	}
	statedb, privateStateRepo, err := api.backend.StateAtBlock(ctx, parent, reexec, nil, true)
	if err != nil {
		return nil, err
	}
	psm, err := api.chainContext(ctx).PrivateStateManager().ResolveForUserContext(ctx)
	if err != nil {
		return nil, err
	}
	privateStateDb, err := privateStateRepo.StatePSI(psm.ID)
	if err != nil {
		return nil, err
	}
	// Retrieve the tracing configurations, or use default values
	var (
		logConfig vm.LogConfig
		txHash    common.Hash
	)
	if config != nil {
		logConfig = config.LogConfig
		txHash = config.TxHash
	}
	logConfig.Debug = true

	// Execute transaction, either tracing all or just the requested one
	var (
		dumps       []string
		signer      = types.MakeSigner(api.backend.ChainConfig(), block.Number())
		chainConfig = api.backend.ChainConfig()
		vmctx       = core.NewEVMBlockContext(block.Header(), api.chainContext(ctx), nil)
		canon       = true
	)
	// Check if there are any overrides: the caller may wish to enable a future
	// fork when executing this block. Note, such overrides are only applicable to the
	// actual specified block, not any preceding blocks that we have to go through
	// in order to obtain the state.
	// Therefore, it's perfectly valid to specify `"futureForkBlock": 0`, to enable `futureFork`

	if config != nil && config.Overrides != nil {
		// Copy the config, to not screw up the main config
		// Note: the Clique-part is _not_ deep copied
		chainConfigCopy := new(params.ChainConfig)
		*chainConfigCopy = *chainConfig
		chainConfig = chainConfigCopy
		if berlin := config.LogConfig.Overrides.BerlinBlock; berlin != nil {
			chainConfig.BerlinBlock = berlin
			canon = false
		}
	}
	for i, tx := range block.Transactions() {
		// Prepare the trasaction for un-traced execution
		var (
			msg, _    = tx.AsMessage(signer)
			txContext = core.NewEVMTxContext(msg)
			vmConf    vm.Config
			dump      *os.File
			writer    *bufio.Writer
			err       error
		)
		// If the transaction needs tracing, swap out the configs
		if tx.Hash() == txHash || txHash == (common.Hash{}) {
			// Generate a unique temporary file to dump it into
			prefix := fmt.Sprintf("block_%#x-%d-%#x-", block.Hash().Bytes()[:4], i, tx.Hash().Bytes()[:4])
			if !canon {
				prefix = fmt.Sprintf("%valt-", prefix)
			}
			dump, err = ioutil.TempFile(os.TempDir(), prefix)
			if err != nil {
				return nil, err
			}
			dumps = append(dumps, dump.Name())

			// Swap out the noop logger to the standard tracer
			writer = bufio.NewWriter(dump)
			vmConf = vm.Config{
				Debug:                   true,
				Tracer:                  vm.NewJSONLogger(&logConfig, writer),
				EnablePreimageRecording: true,
			}
		}
		// Execute the transaction and flush any traces to disk

		// Quorum
		privateStateDbToUse := core.PrivateStateDBForTxn(chainConfig.IsQuorum, tx, statedb, privateStateDb)
		vmConf.ApplyOnPartyOverride = &psm.ID

		vmenv := vm.NewEVM(vmctx, txContext, statedb, privateStateDbToUse, chainConfig, vmConf)

		// Quorum
		vmenv.SetCurrentTX(tx)
		vmenv.InnerApply = func(innerTx *types.Transaction) error {
			return applyInnerTransaction(api.backend.GetBlockchain(), statedb, privateStateDbToUse, block.Header(), tx, vmConf, privateStateRepo.IsMPS(), privateStateRepo, vmenv, innerTx, i)
		}
		msg = api.clearMessageDataIfNonParty(msg, psm)

		statedb.Prepare(tx.Hash(), block.Hash(), i)
		privateStateDbToUse.Prepare(tx.Hash(), block.Hash(), i)

		_, err = core.ApplyMessage(vmenv, msg, new(core.GasPool).AddGas(msg.Gas()))
		if writer != nil {
			writer.Flush()
		}
		if dump != nil {
			dump.Close()
			log.Info("Wrote standard trace", "file", dump.Name())
		}
		if err != nil {
			return dumps, err
		}
		// Finalize the state so any modifications are written to the trie
		// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect
		eip158 := vmenv.ChainConfig().IsEIP158(block.Number())
		statedb.Finalise(eip158)
		privateStateDbToUse.Finalise(eip158)

		// If we've traced the transaction we were looking for, abort
		if tx.Hash() == txHash {
			break
		}
	}
	return dumps, nil
}

// containsTx reports whether the transaction with a certain hash
// is contained within the specified block.
func containsTx(block *types.Block, hash common.Hash) bool {
	for _, tx := range block.Transactions() {
		if tx.Hash() == hash {
			return true
		}
	}
	return false
}

// TraceTransaction returns the structured logs created during the execution of EVM
// and returns them as a JSON object.
func (api *API) TraceTransaction(ctx context.Context, hash common.Hash, config *TraceConfig) (interface{}, error) {
	tx, blockHash, blockNumber, index, err := api.backend.GetTransaction(ctx, hash)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, fmt.Errorf("transaction %#x not found", hash)
	}
	// It shouldn't happen in practice.
	if blockNumber == 0 {
		return nil, errors.New("genesis is not traceable")
	}
	reexec := defaultTraceReexec
	if config != nil && config.Reexec != nil {
		reexec = *config.Reexec
	}
	block, err := api.blockByNumberAndHash(ctx, rpc.BlockNumber(blockNumber), blockHash)
	if err != nil {
		return nil, err
	}
	msg, vmctx, statedb, privateState, privateStateRepo, err := api.backend.StateAtTransaction(ctx, block, int(index), reexec)
	if err != nil {
		return nil, err
	}
	txctx := &txTraceContext{
		index: int(index),
		hash:  hash,
		block: blockHash,
		tx:    tx,
	}
	return api.traceTx(ctx, msg, txctx, vmctx, statedb, privateState, privateStateRepo, config)
}

// TraceCall lets you trace a given eth_call. It collects the structured logs
// created during the execution of EVM if the given transaction was added on
// top of the provided block and returns them as a JSON object.
// You can provide -2 as a block number to trace on top of the pending block.
func (api *API) TraceCall(ctx context.Context, args ethapi.CallArgs, blockNrOrHash rpc.BlockNumberOrHash, config *TraceCallConfig) (interface{}, error) {
	// Try to retrieve the specified block
	var (
		err   error
		block *types.Block
	)
	if hash, ok := blockNrOrHash.Hash(); ok {
		block, err = api.blockByHash(ctx, hash)
	} else if number, ok := blockNrOrHash.Number(); ok {
		block, err = api.blockByNumber(ctx, number)
	} else {
		return nil, errors.New("invalid arguments; neither block nor hash specified")
	}
	if err != nil {
		return nil, err
	}
	// try to recompute the state
	reexec := defaultTraceReexec
	if config != nil && config.Reexec != nil {
		reexec = *config.Reexec
	}
	statedb, privateStateRepo, err := api.backend.StateAtBlock(ctx, block, reexec, nil, true)
	if err != nil {
		return nil, err
	}
	privateStateDb, err := privateStateRepo.DefaultState()
	if err != nil {
		return nil, err
	}

	// Apply the customized state rules if required.
	if config != nil {
		if err := config.StateOverrides.Apply(statedb); err != nil {
			return nil, err
		}
	}

	// Execute the trace
	msg := args.ToMessage(api.backend.RPCGasCap())
	vmctx := core.NewEVMBlockContext(block.Header(), api.chainContext(ctx), nil)

	// Quorum: we run the call with privateState as publicState to check if we have a result, if it is not empty, then it is a private call
	var noTracerConfig *TraceConfig
	if config != nil {
		// create a new config without the tracer so that we have a ExecutionResult returned by api.traceTx
		noTracerConfig = &TraceConfig{
			LogConfig: config.LogConfig,
			Reexec:    config.Reexec,
			Timeout:   config.Timeout,
		}
	}
	var traceConfig *TraceConfig
	if config != nil {
		// create a new config without the tracer so that we have a ExecutionResult returned by api.traceTx
		traceConfig = &TraceConfig{
			LogConfig: config.LogConfig,
			Tracer:    config.Tracer,
			Timeout:   config.Timeout,
			Reexec:    config.Reexec,
		}
	}
	res, err := api.traceTx(ctx, msg, new(txTraceContext), vmctx, statedb, privateStateDb, privateStateRepo, noTracerConfig) // test private with no config
	if exeRes, ok := res.(*ethapi.ExecutionResult); ok && err == nil && len(exeRes.StructLogs) > 0 {                         // check there is a result
		if config != nil && config.Tracer != nil { // re-run the private call with the custom JS tracer
			return api.traceTx(ctx, msg, new(txTraceContext), vmctx, statedb, privateStateDb, privateStateRepo, traceConfig) // re-run with trace
		}
		return res, err // return private result with no tracer
	} else if err == nil && !ok {
		return nil, fmt.Errorf("can not cast traceTx result to *ethapi.ExecutionResult: %#v, %#v", res, err) // better error formatting than "method handler failed"
	}
	// End Quorum

	return api.traceTx(ctx, msg, new(txTraceContext), vmctx, statedb, privateStateDb, privateStateRepo, traceConfig)
}

// traceTx configures a new tracer according to the provided configuration, and
// executes the given message in the provided environment. The return value will
// be tracer dependent.
func (api *API) traceTx(ctx context.Context, message core.Message, txctx *txTraceContext, vmctx vm.BlockContext, statedb, privateStateDb *state.StateDB, privateStateRepo mps.PrivateStateRepository, config *TraceConfig) (interface{}, error) {
	// Assemble the structured logger or the JavaScript tracer
	var (
		tracer    vm.Tracer
		err       error
		txContext = core.NewEVMTxContext(message)
	)
	switch {
	case config != nil && config.Tracer != nil:
		// Define a meaningful timeout of a single transaction trace
		timeout := defaultTraceTimeout
		if config.Timeout != nil {
			if timeout, err = time.ParseDuration(*config.Timeout); err != nil {
				return nil, err
			}
		}
		// Constuct the JavaScript tracer to execute with
		if tracer, err = New(*config.Tracer, txContext); err != nil {
			return nil, err
		}
		// Handle timeouts and RPC cancellations
		deadlineCtx, cancel := context.WithTimeout(ctx, timeout)
		go func() {
			<-deadlineCtx.Done()
			if deadlineCtx.Err() == context.DeadlineExceeded {
				tracer.(*Tracer).Stop(errors.New("execution timeout"))
			}
		}()
		defer cancel()

	case config == nil:
		tracer = vm.NewStructLogger(nil)

	default:
		tracer = vm.NewStructLogger(config.LogConfig)
	}

	// Quorum
	// Set the private state to public state if it is not a private tx
	privateStateDbToUse := core.PrivateStateDBForTxn(api.backend.ChainConfig().IsQuorum, txctx.tx, statedb, privateStateDb)

	psm, err := api.chainContext(ctx).PrivateStateManager().ResolveForUserContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("tracing failed: %v", err)
	}

	// Run the transaction with tracing enabled.
	vmconf := &vm.Config{Debug: true, Tracer: tracer, ApplyOnPartyOverride: &psm.ID}
	vmenv := vm.NewEVM(vmctx, txContext, statedb, privateStateDbToUse, api.backend.ChainConfig(), *vmconf)
	vmenv.SetCurrentTX(txctx.tx)
	vmenv.InnerApply = func(innerTx *types.Transaction) error {
		header, err := api.backend.HeaderByHash(ctx, txctx.block)
		if err != nil {
			return err
		}
		return applyInnerTransaction(api.backend.GetBlockchain(), statedb, privateStateDbToUse, header, txctx.tx, *vmconf, privateStateRepo.IsMPS(), privateStateRepo, vmenv, innerTx, txctx.index)
	}
	// End Quorum

	// Call Prepare to clear out the statedb access list
	statedb.Prepare(txctx.hash, txctx.block, txctx.index)
	privateStateDbToUse.Prepare(txctx.hash, txctx.block, txctx.index)

	result, err := core.ApplyMessage(vmenv, message, new(core.GasPool).AddGas(message.Gas()))
	if err != nil {
		return nil, fmt.Errorf("tracing failed: %w", err)
	}

	// Depending on the tracer type, format and return the output.
	switch tracer := tracer.(type) {
	case *vm.StructLogger:
		// If the result contains a revert reason, return it.
		returnVal := fmt.Sprintf("%x", result.Return())
		if len(result.Revert()) > 0 {
			returnVal = fmt.Sprintf("%x", result.Revert())
		}
		return &ethapi.ExecutionResult{
			Gas:         result.UsedGas,
			Failed:      result.Failed(),
			ReturnValue: returnVal,
			StructLogs:  ethapi.FormatLogs(tracer.StructLogs()),
		}, nil

	case *Tracer:
		return tracer.GetResult()

	default:
		panic(fmt.Sprintf("bad tracer type %T", tracer))
	}
}

// APIs return the collection of RPC services the tracer package offers.
func APIs(backend Backend) []rpc.API {
	// Append all the local APIs and return
	return []rpc.API{
		{
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewAPI(backend),
			Public:    false,
		},
	}
}

// Quorum

// clearMessageDataIfNonParty sets the message data to empty hash in case the private state is not party to the
// transaction. The effect is that when the private tx payload is resolved using the privacy manager the private part of
// the transaction is not retrieved and the transaction is being executed as if the node/private state is not party to
// the transaction.
func (api *API) clearMessageDataIfNonParty(msg types.Message, psm *mps.PrivateStateMetadata) types.Message {
	if msg.IsPrivate() {
		_, managedParties, _, _, _ := private.P.Receive(common.BytesToEncryptedPayloadHash(msg.Data()))

		if api.chainContext(nil).PrivateStateManager().NotIncludeAny(psm, managedParties...) {
			return msg.WithEmptyPrivateData(true)
		}
	}
	return msg
}

func applyInnerTransaction(bc *core.BlockChain, stateDB *state.StateDB, privateStateDB *state.StateDB, header *types.Header, outerTx *types.Transaction, evmConf vm.Config, forceNonParty bool, privateStateRepo mps.PrivateStateRepository, vmenv *vm.EVM, innerTx *types.Transaction, txIndex int) error {
	var (
		author  *common.Address = nil // ApplyTransaction will determine the author from the header so we won't do it here
		gp      *core.GasPool   = new(core.GasPool).AddGas(outerTx.Gas())
		usedGas uint64          = 0
	)
	return core.ApplyInnerTransaction(bc, author, gp, stateDB, privateStateDB, header, outerTx, &usedGas, evmConf, forceNonParty, privateStateRepo, vmenv, innerTx, txIndex)
}
