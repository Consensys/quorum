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

// Package ethconfig contains the configuration of the ETH and LES protocols.
package ethconfig

import (
	"math/big"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	istanbulBackend "github.com/ethereum/go-ethereum/consensus/istanbul/backend"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
)

// FullNodeGPO contains default gasprice oracle settings for full node.
var FullNodeGPO = gasprice.Config{
	Blocks:     20,
	Percentile: 60,
	MaxPrice:   gasprice.DefaultMaxPrice,
}

// LightClientGPO contains default gasprice oracle settings for light client.
var LightClientGPO = gasprice.Config{
	Blocks:     2,
	Percentile: 60,
	MaxPrice:   gasprice.DefaultMaxPrice,
}

// Defaults contains default settings for use on the Ethereum main net.
var Defaults = Config{
	// Quorum - make full sync the default sync mode in quorum (as opposed to upstream geth)
	SyncMode: downloader.FullSync,
	// End Quorum
	Ethash: ethash.Config{
		CacheDir:         "ethash",
		CachesInMem:      2,
		CachesOnDisk:     3,
		CachesLockMmap:   false,
		DatasetsInMem:    1,
		DatasetsOnDisk:   2,
		DatasetsLockMmap: false,
	},
	NetworkId:               1337,
	TxLookupLimit:           2350000,
	LightPeers:              100,
	UltraLightFraction:      75,
	DatabaseCache:           768,
	TrieCleanCache:          154,
	TrieCleanCacheJournal:   "triecache",
	TrieCleanCacheRejournal: 60 * time.Minute,
	TrieDirtyCache:          256,
	TrieTimeout:             60 * time.Minute,
	SnapshotCache:           102,
	Miner: miner.Config{
		GasFloor: params.DefaultMinGasLimit,
		GasCeil:  params.GenesisGasLimit,
		GasPrice: big.NewInt(params.GWei),
		Recommit: 3 * time.Second,
	},
	TxPool:      core.DefaultTxPoolConfig,
	RPCGasCap:   25000000,
	GPO:         FullNodeGPO,
	RPCTxFeeCap: 1, // 1 ether

	// Quorum
	Istanbul: *istanbul.DefaultConfig, // Quorum
}

func init() {
	home := os.Getenv("HOME")
	if home == "" {
		if user, err := user.Current(); err == nil {
			home = user.HomeDir
		}
	}
	if runtime.GOOS == "darwin" {
		Defaults.Ethash.DatasetDir = filepath.Join(home, "Library", "Ethash")
	} else if runtime.GOOS == "windows" {
		localappdata := os.Getenv("LOCALAPPDATA")
		if localappdata != "" {
			Defaults.Ethash.DatasetDir = filepath.Join(localappdata, "Ethash")
		} else {
			Defaults.Ethash.DatasetDir = filepath.Join(home, "AppData", "Local", "Ethash")
		}
	} else {
		Defaults.Ethash.DatasetDir = filepath.Join(home, ".ethash")
	}
}

//go:generate gencodec -type Config -formats toml -out gen_config.go

// Config contains configuration options for of the ETH and LES protocols.
type Config struct {
	// The genesis block, which is inserted if the database is empty.
	// If nil, the Ethereum main net block is used.
	Genesis *core.Genesis `toml:",omitempty"`

	// Protocol options
	NetworkId uint64 // Network ID to use for selecting peers to connect to
	SyncMode  downloader.SyncMode

	// This can be set to list of enrtree:// URLs which will be queried for
	// for nodes to connect to.
	EthDiscoveryURLs  []string
	SnapDiscoveryURLs []string

	NoPruning  bool // Whether to disable pruning and flush everything to disk
	NoPrefetch bool // Whether to disable prefetching and only load state on demand

	TxLookupLimit uint64 `toml:",omitempty"` // The maximum number of blocks from head whose tx indices are reserved.

	// AuthorizationList of required block number -> hash values to accept
	AuthorizationList map[uint64]common.Hash `toml:"-"` // not in the TOML configuration

	// Light client options
	LightServ          int  `toml:",omitempty"` // Maximum percentage of time allowed for serving LES requests
	LightIngress       int  `toml:",omitempty"` // Incoming bandwidth limit for light servers
	LightEgress        int  `toml:",omitempty"` // Outgoing bandwidth limit for light servers
	LightPeers         int  `toml:",omitempty"` // Maximum number of LES client peers
	LightNoPrune       bool `toml:",omitempty"` // Whether to disable light chain pruning
	LightNoSyncServe   bool `toml:",omitempty"` // Whether to serve light clients before syncing
	SyncFromCheckpoint bool `toml:",omitempty"` // Whether to sync the header chain from the configured checkpoint

	// Ultra Light client options
	UltraLightServers      []string `toml:",omitempty"` // List of trusted ultra light servers
	UltraLightFraction     int      `toml:",omitempty"` // Percentage of trusted servers to accept an announcement
	UltraLightOnlyAnnounce bool     `toml:",omitempty"` // Whether to only announce headers, or also serve them

	// Database options
	SkipBcVersionCheck bool `toml:"-"`
	DatabaseHandles    int  `toml:"-"`
	DatabaseCache      int
	DatabaseFreezer    string

	TrieCleanCache          int
	TrieCleanCacheJournal   string        `toml:",omitempty"` // Disk journal directory for trie cache to survive node restarts
	TrieCleanCacheRejournal time.Duration `toml:",omitempty"` // Time interval to regenerate the journal for clean cache
	TrieDirtyCache          int
	TrieTimeout             time.Duration `toml:",omitempty"` // Cumulative Time interval spent on gc, after which to flush trie cache to disk
	SnapshotCache           int
	Preimages               bool

	// Mining options
	Miner miner.Config

	// Ethash options
	Ethash ethash.Config

	// Transaction pool options
	TxPool core.TxPoolConfig

	// Gas Price Oracle options
	GPO gasprice.Config

	// Enables tracking of SHA3 preimages in the VM
	EnablePreimageRecording bool

	// Miscellaneous options
	DocRoot string `toml:"-"`

	// Type of the EWASM interpreter ("" for default)
	EWASMInterpreter string

	// Type of the EVM interpreter ("" for default)
	EVMInterpreter string

	// RPCGasCap is the global gas cap for eth-call variants.
	RPCGasCap uint64

	// RPCTxFeeCap is the global transaction fee(price * gaslimit) cap for
	// send-transction variants. The unit is ether.
	RPCTxFeeCap float64

	// Checkpoint is a hardcoded checkpoint which can be nil.
	Checkpoint *params.TrustedCheckpoint `toml:",omitempty"`

	// CheckpointOracle is the configuration for checkpoint oracle.
	CheckpointOracle *params.CheckpointOracleConfig `toml:",omitempty"`

	// Berlin block override (TODO: remove after the fork)
	OverrideBerlin *big.Int `toml:",omitempty"`

	// Quorum

	RaftMode             bool
	EnableNodePermission bool
	// Istanbul options
	Istanbul istanbul.Config

	// timeout value for call
	EVMCallTimeOut time.Duration

	// Quorum
	core.QuorumChainConfig `toml:"-"`

	// QuorumLight
	QuorumLightServer bool               `toml:",omitempty"`
	QuorumLightClient *QuorumLightClient `toml:",omitempty"`
}

// CreateConsensusEngine creates a consensus engine for the given chain configuration.
func CreateConsensusEngine(stack *node.Node, chainConfig *params.ChainConfig, config *Config, notify []string, noverify bool, db ethdb.Database) consensus.Engine {
	// If proof-of-authority is requested, set it up
	if chainConfig.Clique != nil {
		chainConfig.Clique.AllowedFutureBlockTime = config.Miner.AllowedFutureBlockTime //Quorum
		return clique.New(chainConfig.Clique, db)
	}
	if chainConfig.Transitions != nil && len(chainConfig.Transitions) != 0 {
		config.Istanbul.Transitions = chainConfig.Transitions
	}
	// If Istanbul is requested, set it up
	if chainConfig.Istanbul != nil {
		log.Warn("WARNING: The attribute config.istanbul is deprecated and will be removed in the future, please use config.ibft on genesis file")
		if chainConfig.Istanbul.Epoch != 0 {
			config.Istanbul.Epoch = chainConfig.Istanbul.Epoch
		}
		config.Istanbul.ProposerPolicy = istanbul.NewProposerPolicy(istanbul.ProposerPolicyId(chainConfig.Istanbul.ProposerPolicy))
		config.Istanbul.Ceil2Nby3Block = chainConfig.Istanbul.Ceil2Nby3Block
		config.Istanbul.AllowedFutureBlockTime = config.Miner.AllowedFutureBlockTime //Quorum
		config.Istanbul.TestQBFTBlock = chainConfig.Istanbul.TestQBFTBlock

		return istanbulBackend.New(&config.Istanbul, stack.GetNodeKey(), db)
	}
	if chainConfig.IBFT != nil {
		setBFTConfig(&config.Istanbul, chainConfig.IBFT.BFTConfig)
		config.Istanbul.TestQBFTBlock = nil
		if chainConfig.IBFT.ValidatorContractAddress != (common.Address{}) {
			config.Istanbul.ValidatorContract = chainConfig.IBFT.ValidatorContractAddress
		}
		return istanbulBackend.New(&config.Istanbul, stack.GetNodeKey(), db)
	}
	if chainConfig.QBFT != nil {
		setBFTConfig(&config.Istanbul, chainConfig.QBFT.BFTConfig)
		config.Istanbul.TestQBFTBlock = big.NewInt(0)
		if chainConfig.QBFT.ValidatorContractAddress != (common.Address{}) {
			config.Istanbul.ValidatorContract = chainConfig.QBFT.ValidatorContractAddress
		}
		return istanbulBackend.New(&config.Istanbul, stack.GetNodeKey(), db)
	}
	// For Quorum, Raft run as a separate service, so
	// the Ethereum service still needs a consensus engine,
	// use the consensus with the lightest overhead
	engine := ethash.NewFullFaker()
	engine.SetThreads(-1) // Disable CPU Mining
	return engine
}

// Quorum

type QuorumLightClient struct {
	Use                      bool   `toml:",omitempty"`
	PSI                      string `toml:",omitempty"`
	TokenEnabled             bool   `toml:",omitempty"`
	TokenValue               string `toml:",omitempty"`
	TokenManagement          string `toml:",omitempty"`
	RPCTLS                   bool   `toml:",omitempty"`
	RPCTLSInsecureSkipVerify bool   `toml:",omitempty"`
	RPCTLSCACert             string `toml:",omitempty"`
	RPCTLSCert               string `toml:",omitempty"`
	RPCTLSKey                string `toml:",omitempty"`
	ServerNode               string `toml:",omitempty"`
	ServerNodeRPC            string `toml:",omitempty"`
}

func (q *QuorumLightClient) Enabled() bool {
	return q != nil && q.Use
}

func setBFTConfig(istanbulConfig *istanbul.Config, bftConfig *params.BFTConfig) {
	if bftConfig.BlockPeriodSeconds != 0 {
		istanbulConfig.BlockPeriod = bftConfig.BlockPeriodSeconds
	}
	if bftConfig.EmptyBlockPeriodSeconds != 0 {
		istanbulConfig.EmptyBlockPeriod = bftConfig.EmptyBlockPeriodSeconds
	}
	if bftConfig.EmptyBlockPeriodSeconds < bftConfig.BlockPeriodSeconds {
		istanbulConfig.EmptyBlockPeriod = bftConfig.BlockPeriodSeconds
	}
	if bftConfig.RequestTimeoutSeconds != 0 {
		istanbulConfig.RequestTimeout = bftConfig.RequestTimeoutSeconds * 1000
	}
	if bftConfig.EpochLength != 0 {
		istanbulConfig.Epoch = bftConfig.EpochLength
	}
	if bftConfig.ProposerPolicy != 0 {
		istanbulConfig.ProposerPolicy = istanbul.NewProposerPolicy(istanbul.ProposerPolicyId(bftConfig.ProposerPolicy))
	}
	if bftConfig.Ceil2Nby3Block != nil {
		istanbulConfig.Ceil2Nby3Block = bftConfig.Ceil2Nby3Block
	}
}
