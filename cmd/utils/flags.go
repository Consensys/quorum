// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// Package utils contains internal helper functions for go-ethereum commands.
package utils

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/big"
	"net/url"
	"os"
	"path/filepath"
	godebug "runtime/debug"
	"strconv"
	"strings"
	"text/tabwriter"
	"text/template"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/fdlimit"
	http2 "github.com/ethereum/go-ethereum/common/http"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	istanbulBackend "github.com/ethereum/go-ethereum/consensus/istanbul/backend"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethstats"
	"github.com/ethereum/go-ethereum/extension"
	"github.com/ethereum/go-ethereum/graphql"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/internal/flags"
	"github.com/ethereum/go-ethereum/les"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/metrics/exp"
	"github.com/ethereum/go-ethereum/metrics/influxdb"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/nat"
	"github.com/ethereum/go-ethereum/p2p/netutil"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/permission"
	"github.com/ethereum/go-ethereum/permission/core/types"
	"github.com/ethereum/go-ethereum/plugin"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/raft"
	pcsclite "github.com/gballet/go-libpcsclite"
	gopsutil "github.com/shirou/gopsutil/mem"
	"gopkg.in/urfave/cli.v1"
)

func init() {
	cli.AppHelpTemplate = `{{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]

VERSION:
   {{.Version}}

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`
	cli.CommandHelpTemplate = flags.CommandHelpTemplate
	cli.HelpPrinter = printHelp
}

func printHelp(out io.Writer, templ string, data interface{}) {
	funcMap := template.FuncMap{"join": strings.Join}
	t := template.Must(template.New("help").Funcs(funcMap).Parse(templ))
	w := tabwriter.NewWriter(out, 38, 8, 2, ' ', 0)
	err := t.Execute(w, data)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// These are all the command line flags we support.
// If you add to this list, please remember to include the
// flag in the appropriate command definition.
//
// The flags are defined here so their names and help texts
// are the same for all commands.

var (
	// General settings
	DataDirFlag = DirectoryFlag{
		Name:  "datadir",
		Usage: "Data directory for the databases and keystore",
		Value: DirectoryString(node.DefaultDataDir()),
	}
	RaftLogDirFlag = DirectoryFlag{
		Name:  "raftlogdir",
		Usage: "Raft log directory for the raft-state, raft-snap and raft-wal folders",
		Value: DirectoryString(node.DefaultDataDir()),
	}
	AncientFlag = DirectoryFlag{
		Name:  "datadir.ancient",
		Usage: "Data directory for ancient chain segments (default = inside chaindata)",
	}
	MinFreeDiskSpaceFlag = DirectoryFlag{
		Name:  "datadir.minfreedisk",
		Usage: "Minimum free disk space in MB, once reached triggers auto shut down (default = --cache.gc converted to MB, 0 = disabled)",
	}
	KeyStoreDirFlag = DirectoryFlag{
		Name:  "keystore",
		Usage: "Directory for the keystore (default = inside the datadir)",
	}
	NoUSBFlag = cli.BoolFlag{
		Name:  "nousb",
		Usage: "Disables monitoring for and managing USB hardware wallets (deprecated)",
	}
	USBFlag = cli.BoolFlag{
		Name:  "usb",
		Usage: "Enable monitoring and management of USB hardware wallets",
	}
	SmartCardDaemonPathFlag = cli.StringFlag{
		Name:  "pcscdpath",
		Usage: "Path to the smartcard daemon (pcscd) socket file",
		Value: pcsclite.PCSCDSockName,
	}
	NetworkIdFlag = cli.Uint64Flag{
		Name:  "networkid",
		Usage: "Explicitly set network id (integer)(For testnets: use --ropsten, --rinkeby, --goerli instead)",
		Value: ethconfig.Defaults.NetworkId,
	}
	MainnetFlag = cli.BoolFlag{
		Name:  "mainnet",
		Usage: "Ethereum mainnet",
	}
	GoerliFlag = cli.BoolFlag{
		Name:  "goerli",
		Usage: "Görli network: pre-configured proof-of-authority test network",
	}
	YoloV3Flag = cli.BoolFlag{
		Name:  "yolov3",
		Usage: "YOLOv3 network: pre-configured proof-of-authority shortlived test network.",
	}
	RinkebyFlag = cli.BoolFlag{
		Name:  "rinkeby",
		Usage: "Rinkeby network: pre-configured proof-of-authority test network",
	}
	RopstenFlag = cli.BoolFlag{
		Name:  "ropsten",
		Usage: "Ropsten network: pre-configured proof-of-work test network",
	}
	DeveloperFlag = cli.BoolFlag{
		Name:  "dev",
		Usage: "Ephemeral proof-of-authority network with a pre-funded developer account, mining enabled",
	}
	DeveloperPeriodFlag = cli.IntFlag{
		Name:  "dev.period",
		Usage: "Block period to use in developer mode (0 = mine only if transaction pending)",
	}
	IdentityFlag = cli.StringFlag{
		Name:  "identity",
		Usage: "Custom node name",
	}
	DocRootFlag = DirectoryFlag{
		Name:  "docroot",
		Usage: "Document Root for HTTPClient file scheme",
		Value: DirectoryString(HomeDir()),
	}
	ExitWhenSyncedFlag = cli.BoolFlag{
		Name:  "exitwhensynced",
		Usage: "Exits after block synchronisation completes",
	}
	IterativeOutputFlag = cli.BoolFlag{
		Name:  "iterative",
		Usage: "Print streaming JSON iteratively, delimited by newlines",
	}
	ExcludeStorageFlag = cli.BoolFlag{
		Name:  "nostorage",
		Usage: "Exclude storage entries (save db lookups)",
	}
	IncludeIncompletesFlag = cli.BoolFlag{
		Name:  "incompletes",
		Usage: "Include accounts for which we don't have the address (missing preimage)",
	}
	ExcludeCodeFlag = cli.BoolFlag{
		Name:  "nocode",
		Usage: "Exclude contract code (save db lookups)",
	}
	defaultSyncMode = ethconfig.Defaults.SyncMode
	SyncModeFlag    = TextMarshalerFlag{
		Name:  "syncmode",
		Usage: `Blockchain sync mode ("fast", "full", "snap" or "light")`,
		Value: &defaultSyncMode,
	}
	GCModeFlag = cli.StringFlag{
		Name:  "gcmode",
		Usage: `Blockchain garbage collection mode ("full", "archive")`,
		Value: "full",
	}
	SnapshotFlag = cli.BoolTFlag{
		Name:  "snapshot",
		Usage: `Enables snapshot-database mode (default = enable)`,
	}
	TxLookupLimitFlag = cli.Uint64Flag{
		Name:  "txlookuplimit",
		Usage: "Number of recent blocks to maintain transactions index for (default = about one year, 0 = entire chain)",
		Value: ethconfig.Defaults.TxLookupLimit,
	}
	LightKDFFlag = cli.BoolFlag{
		Name:  "lightkdf",
		Usage: "Reduce key-derivation RAM & CPU usage at some expense of KDF strength",
	}
	DeprecatedAuthorizationListFlag = cli.StringFlag{
		Name:  "whitelist",
		Usage: "[DEPRECATED: will be replaced by 'authorizationlist'] Comma separated block number-to-hash mappings to authorize (<number>=<hash>)",
	}
	AuthorizationListFlag = cli.StringFlag{
		Name:  "authorizationlist",
		Usage: "Comma separated block number-to-hash mappings to authorize (<number>=<hash>)",
	}
	BloomFilterSizeFlag = cli.Uint64Flag{
		Name:  "bloomfilter.size",
		Usage: "Megabytes of memory allocated to bloom-filter for pruning",
		Value: 2048,
	}
	OverrideBerlinFlag = cli.Uint64Flag{
		Name:  "override.berlin",
		Usage: "Manually specify Berlin fork-block, overriding the bundled setting",
	}
	// Light server and client settings
	LightServeFlag = cli.IntFlag{
		Name:  "light.serve",
		Usage: "Maximum percentage of time allowed for serving LES requests (multi-threaded processing allows values over 100)",
		Value: ethconfig.Defaults.LightServ,
	}
	LightIngressFlag = cli.IntFlag{
		Name:  "light.ingress",
		Usage: "Incoming bandwidth limit for serving light clients (kilobytes/sec, 0 = unlimited)",
		Value: ethconfig.Defaults.LightIngress,
	}
	LightEgressFlag = cli.IntFlag{
		Name:  "light.egress",
		Usage: "Outgoing bandwidth limit for serving light clients (kilobytes/sec, 0 = unlimited)",
		Value: ethconfig.Defaults.LightEgress,
	}
	LightMaxPeersFlag = cli.IntFlag{
		Name:  "light.maxpeers",
		Usage: "Maximum number of light clients to serve, or light servers to attach to",
		Value: ethconfig.Defaults.LightPeers,
	}
	UltraLightServersFlag = cli.StringFlag{
		Name:  "ulc.servers",
		Usage: "List of trusted ultra-light servers",
		Value: strings.Join(ethconfig.Defaults.UltraLightServers, ","),
	}
	UltraLightFractionFlag = cli.IntFlag{
		Name:  "ulc.fraction",
		Usage: "Minimum % of trusted ultra-light servers required to announce a new head",
		Value: ethconfig.Defaults.UltraLightFraction,
	}
	UltraLightOnlyAnnounceFlag = cli.BoolFlag{
		Name:  "ulc.onlyannounce",
		Usage: "Ultra light server sends announcements only",
	}
	LightNoPruneFlag = cli.BoolFlag{
		Name:  "light.nopruning",
		Usage: "Disable ancient light chain data pruning",
	}
	LightNoSyncServeFlag = cli.BoolFlag{
		Name:  "light.nosyncserve",
		Usage: "Enables serving light clients before syncing",
	}
	// Ethash settings
	EthashCacheDirFlag = DirectoryFlag{
		Name:  "ethash.cachedir",
		Usage: "Directory to store the ethash verification caches (default = inside the datadir)",
	}
	EthashCachesInMemoryFlag = cli.IntFlag{
		Name:  "ethash.cachesinmem",
		Usage: "Number of recent ethash caches to keep in memory (16MB each)",
		Value: ethconfig.Defaults.Ethash.CachesInMem,
	}
	EthashCachesOnDiskFlag = cli.IntFlag{
		Name:  "ethash.cachesondisk",
		Usage: "Number of recent ethash caches to keep on disk (16MB each)",
		Value: ethconfig.Defaults.Ethash.CachesOnDisk,
	}
	EthashCachesLockMmapFlag = cli.BoolFlag{
		Name:  "ethash.cacheslockmmap",
		Usage: "Lock memory maps of recent ethash caches",
	}
	EthashDatasetDirFlag = DirectoryFlag{
		Name:  "ethash.dagdir",
		Usage: "Directory to store the ethash mining DAGs",
		Value: DirectoryString(ethconfig.Defaults.Ethash.DatasetDir),
	}
	EthashDatasetsInMemoryFlag = cli.IntFlag{
		Name:  "ethash.dagsinmem",
		Usage: "Number of recent ethash mining DAGs to keep in memory (1+GB each)",
		Value: ethconfig.Defaults.Ethash.DatasetsInMem,
	}
	EthashDatasetsOnDiskFlag = cli.IntFlag{
		Name:  "ethash.dagsondisk",
		Usage: "Number of recent ethash mining DAGs to keep on disk (1+GB each)",
		Value: ethconfig.Defaults.Ethash.DatasetsOnDisk,
	}
	EthashDatasetsLockMmapFlag = cli.BoolFlag{
		Name:  "ethash.dagslockmmap",
		Usage: "Lock memory maps for recent ethash mining DAGs",
	}
	// Transaction pool settings
	TxPoolLocalsFlag = cli.StringFlag{
		Name:  "txpool.locals",
		Usage: "Comma separated accounts to treat as locals (no flush, priority inclusion)",
	}
	TxPoolNoLocalsFlag = cli.BoolFlag{
		Name:  "txpool.nolocals",
		Usage: "Disables price exemptions for locally submitted transactions",
	}
	TxPoolJournalFlag = cli.StringFlag{
		Name:  "txpool.journal",
		Usage: "Disk journal for local transaction to survive node restarts",
		Value: core.DefaultTxPoolConfig.Journal,
	}
	TxPoolRejournalFlag = cli.DurationFlag{
		Name:  "txpool.rejournal",
		Usage: "Time interval to regenerate the local transaction journal",
		Value: core.DefaultTxPoolConfig.Rejournal,
	}
	TxPoolPriceLimitFlag = cli.Uint64Flag{
		Name:  "txpool.pricelimit",
		Usage: "Minimum gas price limit to enforce for acceptance into the pool",
		Value: ethconfig.Defaults.TxPool.PriceLimit,
	}
	TxPoolPriceBumpFlag = cli.Uint64Flag{
		Name:  "txpool.pricebump",
		Usage: "Price bump percentage to replace an already existing transaction",
		Value: ethconfig.Defaults.TxPool.PriceBump,
	}
	TxPoolAccountSlotsFlag = cli.Uint64Flag{
		Name:  "txpool.accountslots",
		Usage: "Minimum number of executable transaction slots guaranteed per account",
		Value: ethconfig.Defaults.TxPool.AccountSlots,
	}
	TxPoolGlobalSlotsFlag = cli.Uint64Flag{
		Name:  "txpool.globalslots",
		Usage: "Maximum number of executable transaction slots for all accounts",
		Value: ethconfig.Defaults.TxPool.GlobalSlots,
	}
	TxPoolAccountQueueFlag = cli.Uint64Flag{
		Name:  "txpool.accountqueue",
		Usage: "Maximum number of non-executable transaction slots permitted per account",
		Value: ethconfig.Defaults.TxPool.AccountQueue,
	}
	TxPoolGlobalQueueFlag = cli.Uint64Flag{
		Name:  "txpool.globalqueue",
		Usage: "Maximum number of non-executable transaction slots for all accounts",
		Value: ethconfig.Defaults.TxPool.GlobalQueue,
	}
	TxPoolLifetimeFlag = cli.DurationFlag{
		Name:  "txpool.lifetime",
		Usage: "Maximum amount of time non-executable transaction are queued",
		Value: ethconfig.Defaults.TxPool.Lifetime,
	}
	// Performance tuning settings
	CacheFlag = cli.IntFlag{
		Name:  "cache",
		Usage: "Megabytes of memory allocated to internal caching (default = 4096 mainnet full node, 128 light mode)",
		Value: 1024,
	}
	CacheDatabaseFlag = cli.IntFlag{
		Name:  "cache.database",
		Usage: "Percentage of cache memory allowance to use for database io",
		Value: 50,
	}
	CacheTrieFlag = cli.IntFlag{
		Name:  "cache.trie",
		Usage: "Percentage of cache memory allowance to use for trie caching (default = 15% full mode, 30% archive mode)",
		Value: 15,
	}
	CacheTrieJournalFlag = cli.StringFlag{
		Name:  "cache.trie.journal",
		Usage: "Disk journal directory for trie cache to survive node restarts",
		Value: ethconfig.Defaults.TrieCleanCacheJournal,
	}
	CacheTrieRejournalFlag = cli.DurationFlag{
		Name:  "cache.trie.rejournal",
		Usage: "Time interval to regenerate the trie cache journal",
		Value: ethconfig.Defaults.TrieCleanCacheRejournal,
	}
	CacheGCFlag = cli.IntFlag{
		Name:  "cache.gc",
		Usage: "Percentage of cache memory allowance to use for trie pruning (default = 25% full mode, 0% archive mode)",
		Value: 25,
	}
	CacheSnapshotFlag = cli.IntFlag{
		Name:  "cache.snapshot",
		Usage: "Percentage of cache memory allowance to use for snapshot caching (default = 10% full mode, 20% archive mode)",
		Value: 10,
	}
	CacheNoPrefetchFlag = cli.BoolFlag{
		Name:  "cache.noprefetch",
		Usage: "Disable heuristic state prefetch during block import (less CPU and disk IO, more time waiting for data)",
	}
	CachePreimagesFlag = cli.BoolFlag{
		Name:  "cache.preimages",
		Usage: "Enable recording the SHA3/keccak preimages of trie keys",
	}
	// Miner settings
	MiningEnabledFlag = cli.BoolFlag{
		Name:  "mine",
		Usage: "Enable mining",
	}
	MinerThreadsFlag = cli.IntFlag{
		Name:  "miner.threads",
		Usage: "Number of CPU threads to use for mining",
		Value: 0,
	}
	MinerNotifyFlag = cli.StringFlag{
		Name:  "miner.notify",
		Usage: "Comma separated HTTP URL list to notify of new work packages",
	}
	MinerNotifyFullFlag = cli.BoolFlag{
		Name:  "miner.notify.full",
		Usage: "Notify with pending block headers instead of work packages",
	}
	MinerGasTargetFlag = cli.Uint64Flag{
		Name:  "miner.gastarget",
		Usage: "Target gas floor for mined blocks",
		Value: ethconfig.Defaults.Miner.GasFloor,
	}
	MinerGasLimitFlag = cli.Uint64Flag{
		Name:  "miner.gaslimit",
		Usage: "Target gas ceiling for mined blocks",
		Value: ethconfig.Defaults.Miner.GasCeil,
	}
	MinerGasPriceFlag = BigFlag{
		Name:  "miner.gasprice",
		Usage: "Minimum gas price for mining a transaction",
		Value: ethconfig.Defaults.Miner.GasPrice,
	}
	MinerEtherbaseFlag = cli.StringFlag{
		Name:  "miner.etherbase",
		Usage: "Public address for block mining rewards (default = first account)",
		Value: "0",
	}
	MinerExtraDataFlag = cli.StringFlag{
		Name:  "miner.extradata",
		Usage: "Block extra data set by the miner (default = client version)",
	}
	MinerRecommitIntervalFlag = cli.DurationFlag{
		Name:  "miner.recommit",
		Usage: "Time interval to recreate the block being mined",
		Value: ethconfig.Defaults.Miner.Recommit,
	}
	MinerNoVerfiyFlag = cli.BoolFlag{
		Name:  "miner.noverify",
		Usage: "Disable remote sealing verification",
	}
	// Account settings
	UnlockedAccountFlag = cli.StringFlag{
		Name:  "unlock",
		Usage: "Comma separated list of accounts to unlock",
		Value: "",
	}
	PasswordFileFlag = cli.StringFlag{
		Name:  "password",
		Usage: "Password file to use for non-interactive password input",
		Value: "",
	}
	ExternalSignerFlag = cli.StringFlag{
		Name:  "signer",
		Usage: "External signer (url or path to ipc file)",
		Value: "",
	}
	VMEnableDebugFlag = cli.BoolFlag{
		Name:  "vmdebug",
		Usage: "Record information useful for VM and contract debugging",
	}
	InsecureUnlockAllowedFlag = cli.BoolFlag{
		Name:  "allow-insecure-unlock",
		Usage: "Allow insecure account unlocking when account-related RPCs are exposed by http",
	}
	RPCGlobalGasCapFlag = cli.Uint64Flag{
		Name:  "rpc.gascap",
		Usage: "Sets a cap on gas that can be used in eth_call/estimateGas (0=infinite)",
		Value: ethconfig.Defaults.RPCGasCap,
	}
	RPCGlobalTxFeeCapFlag = cli.Float64Flag{
		Name:  "rpc.txfeecap",
		Usage: "Sets a cap on transaction fee (in ether) that can be sent via the RPC APIs (0 = no cap)",
		Value: ethconfig.Defaults.RPCTxFeeCap,
	}
	// Logging and debug settings
	EthStatsURLFlag = cli.StringFlag{
		Name:  "ethstats",
		Usage: "Reporting URL of a ethstats service (nodename:secret@host:port)",
	}
	FakePoWFlag = cli.BoolFlag{
		Name:  "fakepow",
		Usage: "Disables proof-of-work verification",
	}
	NoCompactionFlag = cli.BoolFlag{
		Name:  "nocompaction",
		Usage: "Disables db compaction after import",
	}

	// Quorum
	// RPC Client Settings
	RPCClientToken = cli.StringFlag{
		Name:  "rpcclitoken",
		Usage: "RPC Client access token",
	}
	RPCClientTLSCert = cli.StringFlag{
		Name:  "rpcclitls.cert",
		Usage: "Server's TLS certificate PEM file on connection by client",
	}
	RPCClientTLSCaCert = cli.StringFlag{
		Name:  "rpcclitls.cacert",
		Usage: "CA certificate PEM file for provided server's TLS certificate on connection by client",
	}
	RPCClientTLSCipherSuites = cli.StringFlag{
		Name:  "rpcclitls.ciphersuites",
		Usage: "Customize supported cipher suites when using TLS connection. Value is a comma-separated cipher suite string",
	}
	RPCClientTLSInsecureSkipVerify = cli.BoolFlag{
		Name:  "rpcclitls.insecureskipverify",
		Usage: "Disable verification of server's TLS certificate on connection by client",
	}
	// End Quorum

	// RPC settings
	IPCDisabledFlag = cli.BoolFlag{
		Name:  "ipcdisable",
		Usage: "Disable the IPC-RPC server",
	}
	IPCPathFlag = DirectoryFlag{
		Name:  "ipcpath",
		Usage: "Filename for IPC socket/pipe within the datadir (explicit paths escape it)",
	}
	HTTPEnabledFlag = cli.BoolFlag{
		Name:  "http",
		Usage: "Enable the HTTP-RPC server",
	}
	HTTPListenAddrFlag = cli.StringFlag{
		Name:  "http.addr",
		Usage: "HTTP-RPC server listening interface",
		Value: node.DefaultHTTPHost,
	}
	HTTPPortFlag = cli.IntFlag{
		Name:  "http.port",
		Usage: "HTTP-RPC server listening port",
		Value: node.DefaultHTTPPort,
	}
	HTTPCORSDomainFlag = cli.StringFlag{
		Name:  "http.corsdomain",
		Usage: "Comma separated list of domains from which to accept cross origin requests (browser enforced)",
		Value: "",
	}
	HTTPVirtualHostsFlag = cli.StringFlag{
		Name:  "http.vhosts",
		Usage: "Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard.",
		Value: strings.Join(node.DefaultConfig.HTTPVirtualHosts, ","),
	}
	HTTPApiFlag = cli.StringFlag{
		Name:  "http.api",
		Usage: "API's offered over the HTTP-RPC interface",
		Value: "",
	}
	HTTPPathPrefixFlag = cli.StringFlag{
		Name:  "http.rpcprefix",
		Usage: "HTTP path path prefix on which JSON-RPC is served. Use '/' to serve on all paths.",
		Value: "",
	}
	GraphQLEnabledFlag = cli.BoolFlag{
		Name:  "graphql",
		Usage: "Enable GraphQL on the HTTP-RPC server. Note that GraphQL can only be started if an HTTP server is started as well.",
	}
	GraphQLCORSDomainFlag = cli.StringFlag{
		Name:  "graphql.corsdomain",
		Usage: "Comma separated list of domains from which to accept cross origin requests (browser enforced)",
		Value: "",
	}
	GraphQLVirtualHostsFlag = cli.StringFlag{
		Name:  "graphql.vhosts",
		Usage: "Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard.",
		Value: strings.Join(node.DefaultConfig.GraphQLVirtualHosts, ","),
	}
	WSEnabledFlag = cli.BoolFlag{
		Name:  "ws",
		Usage: "Enable the WS-RPC server",
	}
	WSListenAddrFlag = cli.StringFlag{
		Name:  "ws.addr",
		Usage: "WS-RPC server listening interface",
		Value: node.DefaultWSHost,
	}
	WSPortFlag = cli.IntFlag{
		Name:  "ws.port",
		Usage: "WS-RPC server listening port",
		Value: node.DefaultWSPort,
	}
	WSApiFlag = cli.StringFlag{
		Name:  "ws.api",
		Usage: "API's offered over the WS-RPC interface",
		Value: "",
	}
	WSAllowedOriginsFlag = cli.StringFlag{
		Name:  "ws.origins",
		Usage: "Origins from which to accept websockets requests",
		Value: "",
	}
	WSPathPrefixFlag = cli.StringFlag{
		Name:  "ws.rpcprefix",
		Usage: "HTTP path prefix on which JSON-RPC is served. Use '/' to serve on all paths.",
		Value: "",
	}
	ExecFlag = cli.StringFlag{
		Name:  "exec",
		Usage: "Execute JavaScript statement",
	}
	PreloadJSFlag = cli.StringFlag{
		Name:  "preload",
		Usage: "Comma separated list of JavaScript files to preload into the console",
	}
	AllowUnprotectedTxs = cli.BoolFlag{
		Name:  "rpc.allow-unprotected-txs",
		Usage: "Allow for unprotected (non EIP155 signed) transactions to be submitted via RPC",
	}

	// Network Settings
	MaxPeersFlag = cli.IntFlag{
		Name:  "maxpeers",
		Usage: "Maximum number of network peers (network disabled if set to 0)",
		Value: node.DefaultConfig.P2P.MaxPeers,
	}
	MaxPendingPeersFlag = cli.IntFlag{
		Name:  "maxpendpeers",
		Usage: "Maximum number of pending connection attempts (defaults used if set to 0)",
		Value: node.DefaultConfig.P2P.MaxPendingPeers,
	}
	ListenPortFlag = cli.IntFlag{
		Name:  "port",
		Usage: "Network listening port",
		Value: 30303,
	}
	BootnodesFlag = cli.StringFlag{
		Name:  "bootnodes",
		Usage: "Comma separated enode URLs for P2P discovery bootstrap",
		Value: "",
	}
	NodeKeyFileFlag = cli.StringFlag{
		Name:  "nodekey",
		Usage: "P2P node key file",
	}
	NodeKeyHexFlag = cli.StringFlag{
		Name:  "nodekeyhex",
		Usage: "P2P node key as hex (for testing)",
	}
	NATFlag = cli.StringFlag{
		Name:  "nat",
		Usage: "NAT port mapping mechanism (any|none|upnp|pmp|extip:<IP>)",
		Value: "any",
	}
	NoDiscoverFlag = cli.BoolFlag{
		Name:  "nodiscover",
		Usage: "Disables the peer discovery mechanism (manual peer addition)",
	}
	DiscoveryV5Flag = cli.BoolFlag{
		Name:  "v5disc",
		Usage: "Enables the experimental RLPx V5 (Topic Discovery) mechanism",
	}
	NetrestrictFlag = cli.StringFlag{
		Name:  "netrestrict",
		Usage: "Restricts network communication to the given IP networks (CIDR masks)",
	}
	DNSDiscoveryFlag = cli.StringFlag{
		Name:  "discovery.dns",
		Usage: "Sets DNS discovery entry points (use \"\" to disable DNS)",
	}

	// ATM the url is left to the user and deployment to
	JSpathFlag = cli.StringFlag{
		Name:  "jspath",
		Usage: "JavaScript root path for `loadScript`",
		Value: ".",
	}

	// Gas price oracle settings
	GpoBlocksFlag = cli.IntFlag{
		Name:  "gpo.blocks",
		Usage: "Number of recent blocks to check for gas prices",
		Value: ethconfig.Defaults.GPO.Blocks,
	}
	GpoPercentileFlag = cli.IntFlag{
		Name:  "gpo.percentile",
		Usage: "Suggested gas price is the given percentile of a set of recent transaction gas prices",
		Value: ethconfig.Defaults.GPO.Percentile,
	}
	GpoMaxGasPriceFlag = cli.Int64Flag{
		Name:  "gpo.maxprice",
		Usage: "Maximum gas price will be recommended by gpo",
		Value: ethconfig.Defaults.GPO.MaxPrice.Int64(),
	}

	// Metrics flags
	MetricsEnabledFlag = cli.BoolFlag{
		Name:  "metrics",
		Usage: "Enable metrics collection and reporting",
	}
	MetricsEnabledExpensiveFlag = cli.BoolFlag{
		Name:  "metrics.expensive",
		Usage: "Enable expensive metrics collection and reporting",
	}

	// MetricsHTTPFlag defines the endpoint for a stand-alone metrics HTTP endpoint.
	// Since the pprof service enables sensitive/vulnerable behavior, this allows a user
	// to enable a public-OK metrics endpoint without having to worry about ALSO exposing
	// other profiling behavior or information.
	MetricsHTTPFlag = cli.StringFlag{
		Name:  "metrics.addr",
		Usage: "Enable stand-alone metrics HTTP server listening interface",
		Value: metrics.DefaultConfig.HTTP,
	}
	MetricsPortFlag = cli.IntFlag{
		Name:  "metrics.port",
		Usage: "Metrics HTTP server listening port",
		Value: metrics.DefaultConfig.Port,
	}
	MetricsEnableInfluxDBFlag = cli.BoolFlag{
		Name:  "metrics.influxdb",
		Usage: "Enable metrics export/push to an external InfluxDB database",
	}
	MetricsInfluxDBEndpointFlag = cli.StringFlag{
		Name:  "metrics.influxdb.endpoint",
		Usage: "InfluxDB API endpoint to report metrics to",
		Value: metrics.DefaultConfig.InfluxDBEndpoint,
	}
	MetricsInfluxDBDatabaseFlag = cli.StringFlag{
		Name:  "metrics.influxdb.database",
		Usage: "InfluxDB database name to push reported metrics to",
		Value: metrics.DefaultConfig.InfluxDBDatabase,
	}
	MetricsInfluxDBUsernameFlag = cli.StringFlag{
		Name:  "metrics.influxdb.username",
		Usage: "Username to authorize access to the database",
		Value: metrics.DefaultConfig.InfluxDBUsername,
	}
	MetricsInfluxDBPasswordFlag = cli.StringFlag{
		Name:  "metrics.influxdb.password",
		Usage: "Password to authorize access to the database",
		Value: metrics.DefaultConfig.InfluxDBPassword,
	}
	// Tags are part of every measurement sent to InfluxDB. Queries on tags are faster in InfluxDB.
	// For example `host` tag could be used so that we can group all nodes and average a measurement
	// across all of them, but also so that we can select a specific node and inspect its measurements.
	// https://docs.influxdata.com/influxdb/v1.4/concepts/key_concepts/#tag-key
	MetricsInfluxDBTagsFlag = cli.StringFlag{
		Name:  "metrics.influxdb.tags",
		Usage: "Comma-separated InfluxDB tags (key/values) attached to all measurements",
		Value: metrics.DefaultConfig.InfluxDBTags,
	}
	EWASMInterpreterFlag = cli.StringFlag{
		Name:  "vm.ewasm",
		Usage: "External ewasm configuration (default = built-in interpreter)",
		Value: "",
	}
	EVMInterpreterFlag = cli.StringFlag{
		Name:  "vm.evm",
		Usage: "External EVM configuration (default = built-in interpreter)",
		Value: "",
	}

	// Quorum - added configurable call timeout for execution of calls
	EVMCallTimeOutFlag = cli.IntFlag{
		Name:  "vm.calltimeout",
		Usage: "Timeout duration in seconds for message call execution without creating a transaction. Value 0 means no timeout.",
		Value: 5,
	}

	// Quorum
	// immutability threshold which can be passed as a parameter at geth start
	QuorumImmutabilityThreshold = cli.IntFlag{
		Name:  "immutabilitythreshold",
		Usage: "overrides the default immutability threshold for Quorum nodes. Its the threshold beyond which block data will be moved to ancient db",
		Value: 3162240,
	}
	// Raft flags
	RaftModeFlag = cli.BoolFlag{
		Name:  "raft",
		Usage: "If enabled, uses Raft instead of Quorum Chain for consensus",
	}
	RaftBlockTimeFlag = cli.IntFlag{
		Name:  "raftblocktime",
		Usage: "Amount of time between raft block creations in milliseconds",
		Value: 50,
	}
	RaftJoinExistingFlag = cli.IntFlag{
		Name:  "raftjoinexisting",
		Usage: "The raft ID to assume when joining an pre-existing cluster",
		Value: 0,
	}

	EmitCheckpointsFlag = cli.BoolFlag{
		Name:  "emitcheckpoints",
		Usage: "If enabled, emit specially formatted logging checkpoints",
	}
	RaftPortFlag = cli.IntFlag{
		Name:  "raftport",
		Usage: "The port to bind for the raft transport",
		Value: 50400,
	}
	RaftDNSEnabledFlag = cli.BoolFlag{
		Name:  "raftdnsenable",
		Usage: "Enable DNS resolution of peers",
	}

	// Permission
	EnableNodePermissionFlag = cli.BoolFlag{
		Name:  "permissioned",
		Usage: "If enabled, the node will allow only a defined list of nodes to connect",
	}
	AllowedFutureBlockTimeFlag = cli.Uint64Flag{
		Name:  "allowedfutureblocktime",
		Usage: "Max time (in seconds) from current time allowed for blocks, before they're considered future blocks",
		Value: 0,
	}
	// Plugins settings
	PluginSettingsFlag = cli.StringFlag{
		Name:  "plugins",
		Usage: "The URI of configuration which describes plugins being used. E.g.: file:///opt/geth/plugins.json",
	}
	PluginLocalVerifyFlag = cli.BoolFlag{
		Name:  "plugins.localverify",
		Usage: "If enabled, verify plugin integrity from local file system. This requires plugin signature file and PGP public key file to be available",
	}
	PluginPublicKeyFlag = cli.StringFlag{
		Name:  "plugins.publickey",
		Usage: fmt.Sprintf("The URI of PGP public key for local plugin verification. E.g.: file:///opt/geth/pubkey.pgp.asc. This flag is only valid if --%s is set (default = file:///<pluginBaseDir>/%s)", PluginLocalVerifyFlag.Name, plugin.DefaultPublicKeyFile),
	}
	PluginSkipVerifyFlag = cli.BoolFlag{
		Name:  "plugins.skipverify",
		Usage: "If enabled, plugin integrity is NOT verified",
	}
	// account plugin flags
	AccountPluginNewAccountConfigFlag = cli.StringFlag{
		Name:  "plugins.account.config",
		Usage: "Value will be passed to an account plugin if being used.  See the account plugin implementation's documentation for further details",
	}
	// Istanbul settings
	IstanbulRequestTimeoutFlag = cli.Uint64Flag{
		Name:  "istanbul.requesttimeout",
		Usage: "[Deprecated] Timeout for each Istanbul round in milliseconds",
		Value: ethconfig.Defaults.Istanbul.RequestTimeout,
	}
	IstanbulBlockPeriodFlag = cli.Uint64Flag{
		Name:  "istanbul.blockperiod",
		Usage: "[Deprecated] Default minimum difference between two consecutive block's timestamps in seconds",
		Value: ethconfig.Defaults.Istanbul.BlockPeriod,
	}
	// Multitenancy setting
	MultitenancyFlag = cli.BoolFlag{
		Name:  "multitenancy",
		Usage: "Enable multitenancy support for this node. This requires RPC Security Plugin to also be configured.",
	}

	// Revert Reason
	RevertReasonFlag = cli.BoolFlag{
		Name:  "revertreason",
		Usage: "Enable saving revert reason in the transaction receipts for this node.",
	}

	QuorumEnablePrivateTrieCache = cli.BoolFlag{
		Name:  "privatetriecache.enable",
		Usage: "Enable use of private trie cache for this node.",
	}

	QuorumEnablePrivacyMarker = cli.BoolFlag{
		Name:  "privacymarker.enable",
		Usage: "Enable use of privacy marker transactions (PMT) for this node.",
	}

	// Quorum Private Transaction Manager connection options
	QuorumPTMUnixSocketFlag = DirectoryFlag{
		Name:  "ptm.socket",
		Usage: "Path to the ipc file when using unix domain socket for the private transaction manager connection",
	}
	QuorumPTMUrlFlag = cli.StringFlag{
		Name:  "ptm.url",
		Usage: "URL when using http connection to private transaction manager",
	}
	QuorumPTMTimeoutFlag = cli.UintFlag{
		Name:  "ptm.timeout",
		Usage: "Timeout (seconds) for the private transaction manager connection. Zero value means timeout disabled.",
		Value: http2.DefaultConfig.Timeout,
	}
	QuorumPTMDialTimeoutFlag = cli.UintFlag{
		Name:  "ptm.dialtimeout",
		Usage: "Dial timeout (seconds) for the private transaction manager connection. Zero value means timeout disabled.",
		Value: http2.DefaultConfig.DialTimeout,
	}
	QuorumPTMHttpIdleTimeoutFlag = cli.UintFlag{
		Name:  "ptm.http.idletimeout",
		Usage: "Idle timeout (seconds) for the private transaction manager connection. Zero value means timeout disabled.",
		Value: http2.DefaultConfig.HttpIdleConnTimeout,
	}
	QuorumPTMHttpWriteBufferSizeFlag = cli.IntFlag{
		Name:  "ptm.http.writebuffersize",
		Usage: "Size of the write buffer (bytes) for the private transaction manager connection. Zero value uses http.Transport default.",
		Value: 0,
	}
	QuorumPTMHttpReadBufferSizeFlag = cli.IntFlag{
		Name:  "ptm.http.readbuffersize",
		Usage: "Size of the read buffer (bytes) for the private transaction manager connection. Zero value uses http.Transport default.",
		Value: 0,
	}
	QuorumPTMTlsModeFlag = cli.StringFlag{
		Name:  "ptm.tls.mode",
		Usage: `If "off" then TLS disabled (default). If "strict" then will use TLS for http connection to private transaction manager`,
	}
	QuorumPTMTlsRootCaFlag = DirectoryFlag{
		Name:  "ptm.tls.rootca",
		Usage: "Path to file containing root CA certificate for TLS connection to private transaction manager (defaults to host's certificates)",
	}
	QuorumPTMTlsClientCertFlag = DirectoryFlag{
		Name:  "ptm.tls.clientcert",
		Usage: "Path to file containing client certificate (or chain of certs) for TLS connection to private transaction manager",
	}
	QuorumPTMTlsClientKeyFlag = DirectoryFlag{
		Name:  "ptm.tls.clientkey",
		Usage: "Path to file containing client's private key for TLS connection to private transaction manager",
	}
	QuorumPTMTlsInsecureSkipVerify = cli.BoolFlag{
		Name:  "ptm.tls.insecureskipverify",
		Usage: "Disable verification of server's TLS certificate on connection to private transaction manager",
	}
	QuorumLightServerFlag = cli.BoolFlag{
		Name:  "qlight.server",
		Usage: "If enabled, the quorum light P2P protocol is started in addition to the other P2P protocols",
	}
	QuorumLightServerP2PListenPortFlag = cli.IntFlag{
		Name:  "qlight.server.p2p.port",
		Usage: "QLight Network listening port",
		Value: 30305,
	}
	QuorumLightServerP2PMaxPeersFlag = cli.IntFlag{
		Name:  "qlight.server.p2p.maxpeers",
		Usage: "Maximum number of qlight peers",
		Value: 10,
	}
	QuorumLightServerP2PNetrestrictFlag = cli.StringFlag{
		Name:  "qlight.server.p2p.netrestrict",
		Usage: "Restricts network communication to the given IP networks (CIDR masks)",
	}
	QuorumLightServerP2PPermissioningFlag = cli.BoolFlag{
		Name:  "qlight.server.p2p.permissioning",
		Usage: "If enabled, the qlight peers are checked against a permissioned list and a disallowed list.",
	}
	QuorumLightServerP2PPermissioningPrefixFlag = cli.StringFlag{
		Name:  "qlight.server.p2p.permissioning.prefix",
		Usage: "The prefix for the permissioned-nodes.json and disallowed-nodes.json files.",
	}
	QuorumLightClientFlag = cli.BoolFlag{
		Name:  "qlight.client",
		Usage: "If enabled, the quorum light client P2P protocol is started (only)",
	}
	QuorumLightClientPSIFlag = cli.StringFlag{
		Name:  "qlight.client.psi",
		Usage: "The PSI this client will use to connect to a server node.",
	}
	QuorumLightClientTokenEnabledFlag = cli.BoolFlag{
		Name:  "qlight.client.token.enabled",
		Usage: "Whether the client uses a token when connecting to the qlight server.",
	}
	QuorumLightClientTokenValueFlag = cli.StringFlag{
		Name:  "qlight.client.token.value",
		Usage: "The token this client will use to connect to a server node.",
	}
	QuorumLightClientTokenManagementFlag = cli.StringFlag{
		Name:  "qlight.client.token.management",
		Usage: "The mechanism used to refresh the token. Possible values: none (developer mode)/external (new token must be injected via the qlight RPC API)/client-security-plugin (the client security plugin must be deployed/configured).",
	}
	QuorumLightClientRPCTLSFlag = cli.BoolFlag{
		Name:  "qlight.client.rpc.tls",
		Usage: "If enabled, the quorum light client RPC connection will be configured to use TLS",
	}
	QuorumLightClientRPCTLSInsecureSkipVerifyFlag = cli.BoolFlag{
		Name:  "qlight.client.rpc.tls.insecureskipverify",
		Usage: "If enabled, the quorum light client RPC connection skips TLS verification",
	}
	QuorumLightClientRPCTLSCACertFlag = cli.StringFlag{
		Name:  "qlight.client.rpc.tls.cacert",
		Usage: "The quorum light client RPC client certificate authority.",
	}
	QuorumLightClientRPCTLSCertFlag = cli.StringFlag{
		Name:  "qlight.client.rpc.tls.cert",
		Usage: "The quorum light client RPC client certificate.",
	}
	QuorumLightClientRPCTLSKeyFlag = cli.StringFlag{
		Name:  "qlight.client.rpc.tls.key",
		Usage: "The quorum light client RPC client certificate private key.",
	}
	QuorumLightClientServerNodeFlag = cli.StringFlag{
		Name:  "qlight.client.serverNode",
		Usage: "The node ID of the target server node",
	}
	QuorumLightClientServerNodeRPCFlag = cli.StringFlag{
		Name:  "qlight.client.serverNodeRPC",
		Usage: "The RPC URL of the target server node",
	}
	QuorumLightTLSFlag = cli.BoolFlag{
		Name:  "qlight.tls",
		Usage: "If enabled, the quorum light client P2P protocol will use tls",
	}
	QuorumLightTLSCertFlag = cli.StringFlag{
		Name:  "qlight.tls.cert",
		Usage: "The certificate file to use for the qlight P2P connection",
	}
	QuorumLightTLSKeyFlag = cli.StringFlag{
		Name:  "qlight.tls.key",
		Usage: "The key file to use for the qlight P2P connection",
	}
	QuorumLightTLSCACertsFlag = cli.StringFlag{
		Name:  "qlight.tls.cacerts",
		Usage: "The certificate authorities file to use for validating P2P connection",
	}
	QuorumLightTLSClientAuthFlag = cli.IntFlag{
		Name:  "qlight.tls.clientauth",
		Usage: "The way the client is authenticated. Possible values: 0=NoClientCert(default) 1=RequestClientCert 2=RequireAnyClientCert 3=VerifyClientCertIfGiven 4=RequireAndVerifyClientCert",
		Value: 0,
	}
	QuorumLightTLSCipherSuitesFlag = cli.StringFlag{
		Name:  "qlight.tls.ciphersuites",
		Usage: "The cipher suites to use for the qlight P2P connection",
	}
)

// MakeDataDir retrieves the currently requested data directory, terminating
// if none (or the empty string) is specified. If the node is starting a testnet,
// then a subdirectory of the specified datadir will be used.
func MakeDataDir(ctx *cli.Context) string {
	if path := ctx.GlobalString(DataDirFlag.Name); path != "" {
		if ctx.GlobalBool(RopstenFlag.Name) {
			// Maintain compatibility with older Geth configurations storing the
			// Ropsten database in `testnet` instead of `ropsten`.
			return filepath.Join(path, "ropsten")
		}
		if ctx.GlobalBool(RinkebyFlag.Name) {
			return filepath.Join(path, "rinkeby")
		}
		if ctx.GlobalBool(GoerliFlag.Name) {
			return filepath.Join(path, "goerli")
		}
		if ctx.GlobalBool(YoloV3Flag.Name) {
			return filepath.Join(path, "yolo-v3")
		}
		return path
	}
	Fatalf("Cannot determine default data directory, please set manually (--datadir)")
	return ""
}

// setNodeKey creates a node key from set command line flags, either loading it
// from a file or as a specified hex value. If neither flags were provided, this
// method returns nil and an emphemeral key is to be generated.
func setNodeKey(ctx *cli.Context, cfg *p2p.Config) {
	var (
		hex  = ctx.GlobalString(NodeKeyHexFlag.Name)
		file = ctx.GlobalString(NodeKeyFileFlag.Name)
		key  *ecdsa.PrivateKey
		err  error
	)
	switch {
	case file != "" && hex != "":
		Fatalf("Options %q and %q are mutually exclusive", NodeKeyFileFlag.Name, NodeKeyHexFlag.Name)
	case file != "":
		if key, err = crypto.LoadECDSA(file); err != nil {
			Fatalf("Option %q: %v", NodeKeyFileFlag.Name, err)
		}
		cfg.PrivateKey = key
	case hex != "":
		if key, err = crypto.HexToECDSA(hex); err != nil {
			Fatalf("Option %q: %v", NodeKeyHexFlag.Name, err)
		}
		cfg.PrivateKey = key
	}
}

// setNodeUserIdent creates the user identifier from CLI flags.
func setNodeUserIdent(ctx *cli.Context, cfg *node.Config) {
	if identity := ctx.GlobalString(IdentityFlag.Name); len(identity) > 0 {
		cfg.UserIdent = identity
	}
}

// setBootstrapNodes creates a list of bootstrap nodes from the command line
// flags, reverting to pre-configured ones if none have been specified.
func setBootstrapNodes(ctx *cli.Context, cfg *p2p.Config) {
	urls := params.MainnetBootnodes
	switch {
	case ctx.GlobalIsSet(BootnodesFlag.Name):
		urls = SplitAndTrim(ctx.GlobalString(BootnodesFlag.Name))
	case ctx.GlobalBool(RopstenFlag.Name):
		urls = params.RopstenBootnodes
	case ctx.GlobalBool(RinkebyFlag.Name):
		urls = params.RinkebyBootnodes
	case ctx.GlobalBool(GoerliFlag.Name):
		urls = params.GoerliBootnodes
	case ctx.GlobalBool(YoloV3Flag.Name):
		urls = params.YoloV3Bootnodes
	case cfg.BootstrapNodes != nil:
		return // already set, don't apply defaults.
	}

	cfg.BootstrapNodes = make([]*enode.Node, 0, len(urls))
	for _, url := range urls {
		if url != "" {
			node, err := enode.Parse(enode.ValidSchemes, url)
			if err != nil {
				log.Crit("Bootstrap URL invalid", "enode", url, "err", err)
				continue
			}
			cfg.BootstrapNodes = append(cfg.BootstrapNodes, node)
		}
	}
}

// setBootstrapNodesV5 creates a list of bootstrap nodes from the command line
// flags, reverting to pre-configured ones if none have been specified.
func setBootstrapNodesV5(ctx *cli.Context, cfg *p2p.Config) {
	urls := params.V5Bootnodes
	switch {
	case ctx.GlobalIsSet(BootnodesFlag.Name):
		urls = SplitAndTrim(ctx.GlobalString(BootnodesFlag.Name))
	case cfg.BootstrapNodesV5 != nil:
		return // already set, don't apply defaults.
	}

	cfg.BootstrapNodesV5 = make([]*enode.Node, 0, len(urls))
	for _, url := range urls {
		if url != "" {
			node, err := enode.Parse(enode.ValidSchemes, url)
			if err != nil {
				log.Error("Bootstrap URL invalid", "enode", url, "err", err)
				continue
			}
			cfg.BootstrapNodesV5 = append(cfg.BootstrapNodesV5, node)
		}
	}
}

// setListenAddress creates a TCP listening address string from set command
// line flags.
func setListenAddress(ctx *cli.Context, cfg *p2p.Config) {
	if ctx.GlobalIsSet(ListenPortFlag.Name) {
		cfg.ListenAddr = fmt.Sprintf(":%d", ctx.GlobalInt(ListenPortFlag.Name))
	}
}

// setNAT creates a port mapper from command line flags.
func setNAT(ctx *cli.Context, cfg *p2p.Config) {
	if ctx.GlobalIsSet(NATFlag.Name) {
		natif, err := nat.Parse(ctx.GlobalString(NATFlag.Name))
		if err != nil {
			Fatalf("Option %s: %v", NATFlag.Name, err)
		}
		cfg.NAT = natif
	}
}

// SplitAndTrim splits input separated by a comma
// and trims excessive white space from the substrings.
func SplitAndTrim(input string) (ret []string) {
	l := strings.Split(input, ",")
	for _, r := range l {
		if r = strings.TrimSpace(r); r != "" {
			ret = append(ret, r)
		}
	}
	return ret
}

// setHTTP creates the HTTP RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func setHTTP(ctx *cli.Context, cfg *node.Config) {
	if ctx.GlobalBool(LegacyRPCEnabledFlag.Name) && cfg.HTTPHost == "" {
		log.Warn("The flag --rpc is deprecated and will be removed in the future, please use --http")
		cfg.HTTPHost = "127.0.0.1"
		if ctx.GlobalIsSet(LegacyRPCListenAddrFlag.Name) {
			cfg.HTTPHost = ctx.GlobalString(LegacyRPCListenAddrFlag.Name)
			log.Warn("The flag --rpcaddr is deprecated and will be removed in the future, please use --http.addr")
		}
	}
	if ctx.GlobalBool(HTTPEnabledFlag.Name) && cfg.HTTPHost == "" {
		cfg.HTTPHost = "127.0.0.1"
		if ctx.GlobalIsSet(HTTPListenAddrFlag.Name) {
			cfg.HTTPHost = ctx.GlobalString(HTTPListenAddrFlag.Name)
		}
	}

	if ctx.GlobalIsSet(LegacyRPCPortFlag.Name) {
		cfg.HTTPPort = ctx.GlobalInt(LegacyRPCPortFlag.Name)
		log.Warn("The flag --rpcport is deprecated and will be removed in the future, please use --http.port")
	}
	if ctx.GlobalIsSet(HTTPPortFlag.Name) {
		cfg.HTTPPort = ctx.GlobalInt(HTTPPortFlag.Name)
	}

	if ctx.GlobalIsSet(LegacyRPCCORSDomainFlag.Name) {
		cfg.HTTPCors = SplitAndTrim(ctx.GlobalString(LegacyRPCCORSDomainFlag.Name))
		log.Warn("The flag --rpccorsdomain is deprecated and will be removed in the future, please use --http.corsdomain")
	}
	if ctx.GlobalIsSet(HTTPCORSDomainFlag.Name) {
		cfg.HTTPCors = SplitAndTrim(ctx.GlobalString(HTTPCORSDomainFlag.Name))
	}

	if ctx.GlobalIsSet(LegacyRPCApiFlag.Name) {
		cfg.HTTPModules = SplitAndTrim(ctx.GlobalString(LegacyRPCApiFlag.Name))
		log.Warn("The flag --rpcapi is deprecated and will be removed in the future, please use --http.api")
	}
	if ctx.GlobalIsSet(HTTPApiFlag.Name) {
		cfg.HTTPModules = SplitAndTrim(ctx.GlobalString(HTTPApiFlag.Name))
	}

	if ctx.GlobalIsSet(LegacyRPCVirtualHostsFlag.Name) {
		cfg.HTTPVirtualHosts = SplitAndTrim(ctx.GlobalString(LegacyRPCVirtualHostsFlag.Name))
		log.Warn("The flag --rpcvhosts is deprecated and will be removed in the future, please use --http.vhosts")
	}
	if ctx.GlobalIsSet(HTTPVirtualHostsFlag.Name) {
		cfg.HTTPVirtualHosts = SplitAndTrim(ctx.GlobalString(HTTPVirtualHostsFlag.Name))
	}

	if ctx.GlobalIsSet(HTTPPathPrefixFlag.Name) {
		cfg.HTTPPathPrefix = ctx.GlobalString(HTTPPathPrefixFlag.Name)
	}
	if ctx.GlobalIsSet(AllowUnprotectedTxs.Name) {
		cfg.AllowUnprotectedTxs = ctx.GlobalBool(AllowUnprotectedTxs.Name)
	}
}

// setGraphQL creates the GraphQL listener interface string from the set
// command line flags, returning empty if the GraphQL endpoint is disabled.
func setGraphQL(ctx *cli.Context, cfg *node.Config) {
	if ctx.GlobalIsSet(GraphQLCORSDomainFlag.Name) {
		cfg.GraphQLCors = SplitAndTrim(ctx.GlobalString(GraphQLCORSDomainFlag.Name))
	}
	if ctx.GlobalIsSet(GraphQLVirtualHostsFlag.Name) {
		cfg.GraphQLVirtualHosts = SplitAndTrim(ctx.GlobalString(GraphQLVirtualHostsFlag.Name))
	}
}

// setWS creates the WebSocket RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func setWS(ctx *cli.Context, cfg *node.Config) {
	if ctx.GlobalBool(WSEnabledFlag.Name) && cfg.WSHost == "" {
		cfg.WSHost = "127.0.0.1"
		if ctx.GlobalIsSet(WSListenAddrFlag.Name) {
			cfg.WSHost = ctx.GlobalString(WSListenAddrFlag.Name)
		}
	}
	if ctx.GlobalIsSet(WSPortFlag.Name) {
		cfg.WSPort = ctx.GlobalInt(WSPortFlag.Name)
	}

	if ctx.GlobalIsSet(WSAllowedOriginsFlag.Name) {
		cfg.WSOrigins = SplitAndTrim(ctx.GlobalString(WSAllowedOriginsFlag.Name))
	}

	if ctx.GlobalIsSet(WSApiFlag.Name) {
		cfg.WSModules = SplitAndTrim(ctx.GlobalString(WSApiFlag.Name))
	}

	if ctx.GlobalIsSet(WSPathPrefixFlag.Name) {
		cfg.WSPathPrefix = ctx.GlobalString(WSPathPrefixFlag.Name)
	}
}

// setIPC creates an IPC path configuration from the set command line flags,
// returning an empty string if IPC was explicitly disabled, or the set path.
func setIPC(ctx *cli.Context, cfg *node.Config) {
	CheckExclusive(ctx, IPCDisabledFlag, IPCPathFlag)
	switch {
	case ctx.GlobalBool(IPCDisabledFlag.Name):
		cfg.IPCPath = ""
	case ctx.GlobalIsSet(IPCPathFlag.Name):
		cfg.IPCPath = ctx.GlobalString(IPCPathFlag.Name)
	}
}

// setLes configures the les server and ultra light client settings from the command line flags.
func setLes(ctx *cli.Context, cfg *ethconfig.Config) {
	if ctx.GlobalIsSet(LightServeFlag.Name) {
		cfg.LightServ = ctx.GlobalInt(LightServeFlag.Name)
	}
	if ctx.GlobalIsSet(LightIngressFlag.Name) {
		cfg.LightIngress = ctx.GlobalInt(LightIngressFlag.Name)
	}
	if ctx.GlobalIsSet(LightEgressFlag.Name) {
		cfg.LightEgress = ctx.GlobalInt(LightEgressFlag.Name)
	}
	if ctx.GlobalIsSet(LightMaxPeersFlag.Name) {
		cfg.LightPeers = ctx.GlobalInt(LightMaxPeersFlag.Name)
	}
	if ctx.GlobalIsSet(UltraLightServersFlag.Name) {
		cfg.UltraLightServers = strings.Split(ctx.GlobalString(UltraLightServersFlag.Name), ",")
	}
	if ctx.GlobalIsSet(UltraLightFractionFlag.Name) {
		cfg.UltraLightFraction = ctx.GlobalInt(UltraLightFractionFlag.Name)
	}
	if cfg.UltraLightFraction <= 0 && cfg.UltraLightFraction > 100 {
		log.Error("Ultra light fraction is invalid", "had", cfg.UltraLightFraction, "updated", ethconfig.Defaults.UltraLightFraction)
		cfg.UltraLightFraction = ethconfig.Defaults.UltraLightFraction
	}
	if ctx.GlobalIsSet(UltraLightOnlyAnnounceFlag.Name) {
		cfg.UltraLightOnlyAnnounce = ctx.GlobalBool(UltraLightOnlyAnnounceFlag.Name)
	}
	if ctx.GlobalIsSet(LightNoPruneFlag.Name) {
		cfg.LightNoPrune = ctx.GlobalBool(LightNoPruneFlag.Name)
	}
	if ctx.GlobalIsSet(LightNoSyncServeFlag.Name) {
		cfg.LightNoSyncServe = ctx.GlobalBool(LightNoSyncServeFlag.Name)
	}
}

// MakeDatabaseHandles raises out the number of allowed file handles per process
// for Geth and returns half of the allowance to assign to the database.
func MakeDatabaseHandles() int {
	limit, err := fdlimit.Maximum()
	if err != nil {
		Fatalf("Failed to retrieve file descriptor allowance: %v", err)
	}
	raised, err := fdlimit.Raise(uint64(limit))
	if err != nil {
		Fatalf("Failed to raise file descriptor allowance: %v", err)
	}
	return int(raised / 2) // Leave half for networking and other stuff
}

// MakeAddress converts an account specified directly as a hex encoded string or
// a key index in the key store to an internal account representation.
func MakeAddress(ks *keystore.KeyStore, account string) (accounts.Account, error) {
	// If the specified account is a valid address, return it
	if common.IsHexAddress(account) {
		return accounts.Account{Address: common.HexToAddress(account)}, nil
	}
	// Otherwise try to interpret the account as a keystore index
	index, err := strconv.Atoi(account)
	if err != nil || index < 0 {
		return accounts.Account{}, fmt.Errorf("invalid account address or index %q", account)
	}
	log.Warn("-------------------------------------------------------------------")
	log.Warn("Referring to accounts by order in the keystore folder is dangerous!")
	log.Warn("This functionality is deprecated and will be removed in the future!")
	log.Warn("Please use explicit addresses! (can search via `geth account list`)")
	log.Warn("-------------------------------------------------------------------")

	accs := ks.Accounts()
	if len(accs) <= index {
		return accounts.Account{}, fmt.Errorf("index %d higher than number of accounts %d", index, len(accs))
	}
	return accs[index], nil
}

// setEtherbase retrieves the etherbase either from the directly specified
// command line flags or from the keystore if CLI indexed.
func setEtherbase(ctx *cli.Context, ks *keystore.KeyStore, cfg *ethconfig.Config) {
	// Extract the current etherbase
	var etherbase string
	if ctx.GlobalIsSet(MinerEtherbaseFlag.Name) {
		etherbase = ctx.GlobalString(MinerEtherbaseFlag.Name)
	}
	// Convert the etherbase into an address and configure it
	if etherbase != "" {
		if ks != nil {
			account, err := MakeAddress(ks, etherbase)
			if err != nil {
				Fatalf("Invalid miner etherbase: %v", err)
			}
			cfg.Miner.Etherbase = account.Address
		} else {
			Fatalf("No etherbase configured")
		}
	}
}

// MakePasswordList reads password lines from the file specified by the global --password flag.
func MakePasswordList(ctx *cli.Context) []string {
	path := ctx.GlobalString(PasswordFileFlag.Name)
	if path == "" {
		return nil
	}
	text, err := ioutil.ReadFile(path)
	if err != nil {
		Fatalf("Failed to read password file: %v", err)
	}
	lines := strings.Split(string(text), "\n")
	// Sanitise DOS line endings.
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}
	return lines
}

func SetP2PConfig(ctx *cli.Context, cfg *p2p.Config) {
	setNodeKey(ctx, cfg)
	setNAT(ctx, cfg)
	setListenAddress(ctx, cfg)
	setBootstrapNodes(ctx, cfg)
	setBootstrapNodesV5(ctx, cfg)

	lightClient := ctx.GlobalString(SyncModeFlag.Name) == "light"
	lightServer := (ctx.GlobalInt(LightServeFlag.Name) != 0)

	lightPeers := ctx.GlobalInt(LightMaxPeersFlag.Name)
	if lightClient && !ctx.GlobalIsSet(LightMaxPeersFlag.Name) {
		// dynamic default - for clients we use 1/10th of the default for servers
		lightPeers /= 10
	}

	if ctx.GlobalIsSet(MaxPeersFlag.Name) {
		cfg.MaxPeers = ctx.GlobalInt(MaxPeersFlag.Name)
		if lightServer && !ctx.GlobalIsSet(LightMaxPeersFlag.Name) {
			cfg.MaxPeers += lightPeers
		}
	} else {
		if lightServer {
			cfg.MaxPeers += lightPeers
		}
		if lightClient && ctx.GlobalIsSet(LightMaxPeersFlag.Name) && cfg.MaxPeers < lightPeers {
			cfg.MaxPeers = lightPeers
		}
	}
	if !(lightClient || lightServer) {
		lightPeers = 0
	}
	ethPeers := cfg.MaxPeers - lightPeers
	if lightClient {
		ethPeers = 0
	}
	log.Info("Maximum peer count", "ETH", ethPeers, "LES", lightPeers, "total", cfg.MaxPeers)

	if ctx.GlobalIsSet(MaxPendingPeersFlag.Name) {
		cfg.MaxPendingPeers = ctx.GlobalInt(MaxPendingPeersFlag.Name)
	}
	if ctx.GlobalIsSet(NoDiscoverFlag.Name) || lightClient {
		cfg.NoDiscovery = true
	}

	// if we're running a light client or server, force enable the v5 peer discovery
	// unless it is explicitly disabled with --nodiscover note that explicitly specifying
	// --v5disc overrides --nodiscover, in which case the later only disables v4 discovery
	forceV5Discovery := (lightClient || lightServer) && !ctx.GlobalBool(NoDiscoverFlag.Name)
	if ctx.GlobalIsSet(DiscoveryV5Flag.Name) {
		cfg.DiscoveryV5 = ctx.GlobalBool(DiscoveryV5Flag.Name)
	} else if forceV5Discovery {
		cfg.DiscoveryV5 = true
	}

	if netrestrict := ctx.GlobalString(NetrestrictFlag.Name); netrestrict != "" {
		list, err := netutil.ParseNetlist(netrestrict)
		if err != nil {
			Fatalf("Option %q: %v", NetrestrictFlag.Name, err)
		}
		cfg.NetRestrict = list
	}

	if ctx.GlobalBool(DeveloperFlag.Name) {
		// --dev mode can't use p2p networking.
		cfg.MaxPeers = 0
		cfg.ListenAddr = ":0"
		cfg.NoDiscovery = true
		cfg.DiscoveryV5 = false
	}
}

func SetQP2PConfig(ctx *cli.Context, cfg *p2p.Config) {
	setNodeKey(ctx, cfg)
	//setNAT(ctx, cfg)
	cfg.NAT = nil
	if ctx.GlobalIsSet(QuorumLightServerP2PListenPortFlag.Name) {
		cfg.ListenAddr = fmt.Sprintf(":%d", ctx.GlobalInt(QuorumLightServerP2PListenPortFlag.Name))
	}

	cfg.EnableNodePermission = ctx.GlobalIsSet(QuorumLightServerP2PPermissioningFlag.Name)

	cfg.MaxPeers = 10
	if ctx.GlobalIsSet(QuorumLightServerP2PMaxPeersFlag.Name) {
		cfg.MaxPeers = ctx.GlobalInt(QuorumLightServerP2PMaxPeersFlag.Name)
	}

	if netrestrict := ctx.GlobalString(QuorumLightServerP2PNetrestrictFlag.Name); netrestrict != "" {
		list, err := netutil.ParseNetlist(netrestrict)
		if err != nil {
			Fatalf("Option %q: %v", QuorumLightServerP2PNetrestrictFlag.Name, err)
		}
		cfg.NetRestrict = list
	}

	cfg.MaxPendingPeers = 0
	cfg.NoDiscovery = true
	cfg.DiscoveryV5 = false
	cfg.NoDial = true
}

// SetNodeConfig applies node-related command line flags to the config.
func SetNodeConfig(ctx *cli.Context, cfg *node.Config) {
	SetP2PConfig(ctx, &cfg.P2P)
	if cfg.QP2P != nil {
		SetQP2PConfig(ctx, cfg.QP2P)
	}
	setIPC(ctx, cfg)
	setHTTP(ctx, cfg)
	setGraphQL(ctx, cfg)
	setWS(ctx, cfg)
	setNodeUserIdent(ctx, cfg)
	setDataDir(ctx, cfg)
	setRaftLogDir(ctx, cfg)
	setSmartCard(ctx, cfg)

	if ctx.GlobalIsSet(ExternalSignerFlag.Name) {
		cfg.ExternalSigner = ctx.GlobalString(ExternalSignerFlag.Name)
	}

	if ctx.GlobalIsSet(KeyStoreDirFlag.Name) {
		cfg.KeyStoreDir = ctx.GlobalString(KeyStoreDirFlag.Name)
	}
	if ctx.GlobalIsSet(LightKDFFlag.Name) {
		cfg.UseLightweightKDF = ctx.GlobalBool(LightKDFFlag.Name)
	}
	if ctx.GlobalIsSet(NoUSBFlag.Name) || cfg.NoUSB {
		log.Warn("Option nousb is deprecated and USB is deactivated by default. Use --usb to enable")
	}
	if ctx.GlobalIsSet(USBFlag.Name) {
		cfg.USB = ctx.GlobalBool(USBFlag.Name)
	}
	if ctx.GlobalIsSet(InsecureUnlockAllowedFlag.Name) {
		cfg.InsecureUnlockAllowed = ctx.GlobalBool(InsecureUnlockAllowedFlag.Name)
	}

	// Quorum
	if ctx.GlobalIsSet(EnableNodePermissionFlag.Name) {
		cfg.EnableNodePermission = ctx.GlobalBool(EnableNodePermissionFlag.Name)
	}
	if ctx.GlobalIsSet(MultitenancyFlag.Name) {
		cfg.EnableMultitenancy = ctx.GlobalBool(MultitenancyFlag.Name)
	}
}

func setSmartCard(ctx *cli.Context, cfg *node.Config) {
	// Skip enabling smartcards if no path is set
	path := ctx.GlobalString(SmartCardDaemonPathFlag.Name)
	if path == "" {
		return
	}
	// Sanity check that the smartcard path is valid
	fi, err := os.Stat(path)
	if err != nil {
		log.Info("Smartcard socket not found, disabling", "err", err)
		return
	}
	if fi.Mode()&os.ModeType != os.ModeSocket {
		log.Error("Invalid smartcard daemon path", "path", path, "type", fi.Mode().String())
		return
	}
	// Smartcard daemon path exists and is a socket, enable it
	cfg.SmartCardDaemonPath = path
}

func setDataDir(ctx *cli.Context, cfg *node.Config) {
	switch {
	case ctx.GlobalIsSet(DataDirFlag.Name):
		cfg.DataDir = ctx.GlobalString(DataDirFlag.Name)
	case ctx.GlobalBool(DeveloperFlag.Name):
		cfg.DataDir = "" // unless explicitly requested, use memory databases
	case ctx.GlobalBool(RopstenFlag.Name) && cfg.DataDir == node.DefaultDataDir():
		// Maintain compatibility with older Geth configurations storing the
		// Ropsten database in `testnet` instead of `ropsten`.
		legacyPath := filepath.Join(node.DefaultDataDir(), "testnet")
		if _, err := os.Stat(legacyPath); !os.IsNotExist(err) {
			log.Warn("Using the deprecated `testnet` datadir. Future versions will store the Ropsten chain in `ropsten`.")
			cfg.DataDir = legacyPath
		} else {
			cfg.DataDir = filepath.Join(node.DefaultDataDir(), "ropsten")
		}

		cfg.DataDir = filepath.Join(node.DefaultDataDir(), "ropsten")
	case ctx.GlobalBool(RinkebyFlag.Name) && cfg.DataDir == node.DefaultDataDir():
		cfg.DataDir = filepath.Join(node.DefaultDataDir(), "rinkeby")
	case ctx.GlobalBool(GoerliFlag.Name) && cfg.DataDir == node.DefaultDataDir():
		cfg.DataDir = filepath.Join(node.DefaultDataDir(), "goerli")
	case ctx.GlobalBool(YoloV3Flag.Name) && cfg.DataDir == node.DefaultDataDir():
		cfg.DataDir = filepath.Join(node.DefaultDataDir(), "yolo-v3")
	}
	if err := SetPlugins(ctx, cfg); err != nil {
		Fatalf(err.Error())
	}
}

func setRaftLogDir(ctx *cli.Context, cfg *node.Config) {
	if ctx.GlobalIsSet(RaftLogDirFlag.Name) {
		cfg.RaftLogDir = ctx.GlobalString(RaftLogDirFlag.Name)
	} else {
		cfg.RaftLogDir = cfg.DataDir
	}
}

// Quorum
//
// Read plugin settings from --plugins flag. Overwrite settings defined in --config if any
func SetPlugins(ctx *cli.Context, cfg *node.Config) error {
	if ctx.GlobalIsSet(PluginSettingsFlag.Name) {
		// validate flag combination
		if ctx.GlobalBool(PluginSkipVerifyFlag.Name) && ctx.GlobalBool(PluginLocalVerifyFlag.Name) {
			return fmt.Errorf("only --%s or --%s must be set", PluginSkipVerifyFlag.Name, PluginLocalVerifyFlag.Name)
		}
		if !ctx.GlobalBool(PluginLocalVerifyFlag.Name) && ctx.GlobalIsSet(PluginPublicKeyFlag.Name) {
			return fmt.Errorf("--%s is required for setting --%s", PluginLocalVerifyFlag.Name, PluginPublicKeyFlag.Name)
		}
		pluginSettingsURL, err := url.Parse(ctx.GlobalString(PluginSettingsFlag.Name))
		if err != nil {
			return fmt.Errorf("plugins: Invalid URL for --%s due to %s", PluginSettingsFlag.Name, err)
		}
		var pluginSettings plugin.Settings
		r, err := urlReader(pluginSettingsURL)
		if err != nil {
			return fmt.Errorf("plugins: unable to create reader due to %s", err)
		}
		defer func() {
			_ = r.Close()
		}()
		if err := json.NewDecoder(r).Decode(&pluginSettings); err != nil {
			return fmt.Errorf("plugins: unable to parse settings due to %s", err)
		}
		pluginSettings.SetDefaults()
		cfg.Plugins = &pluginSettings
	}
	return nil
}

func urlReader(u *url.URL) (io.ReadCloser, error) {
	s := u.Scheme
	switch s {
	case "file":
		return os.Open(filepath.Join(u.Host, u.Path))
	}
	return nil, fmt.Errorf("unsupported scheme %s", s)
}

func setGPO(ctx *cli.Context, cfg *gasprice.Config, light bool) {
	// If we are running the light client, apply another group
	// settings for gas oracle.
	if light {
		cfg.Blocks = ethconfig.LightClientGPO.Blocks
		cfg.Percentile = ethconfig.LightClientGPO.Percentile
	}
	if ctx.GlobalIsSet(GpoBlocksFlag.Name) {
		cfg.Blocks = ctx.GlobalInt(GpoBlocksFlag.Name)
	}
	if ctx.GlobalIsSet(GpoPercentileFlag.Name) {
		cfg.Percentile = ctx.GlobalInt(GpoPercentileFlag.Name)
	}
	if ctx.GlobalIsSet(GpoMaxGasPriceFlag.Name) {
		cfg.MaxPrice = big.NewInt(ctx.GlobalInt64(GpoMaxGasPriceFlag.Name))
	}
}

func setTxPool(ctx *cli.Context, cfg *core.TxPoolConfig) {
	if ctx.GlobalIsSet(TxPoolLocalsFlag.Name) {
		locals := strings.Split(ctx.GlobalString(TxPoolLocalsFlag.Name), ",")
		for _, account := range locals {
			if trimmed := strings.TrimSpace(account); !common.IsHexAddress(trimmed) {
				Fatalf("Invalid account in --txpool.locals: %s", trimmed)
			} else {
				cfg.Locals = append(cfg.Locals, common.HexToAddress(account))
			}
		}
	}
	if ctx.GlobalIsSet(TxPoolNoLocalsFlag.Name) {
		cfg.NoLocals = ctx.GlobalBool(TxPoolNoLocalsFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolJournalFlag.Name) {
		cfg.Journal = ctx.GlobalString(TxPoolJournalFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolRejournalFlag.Name) {
		cfg.Rejournal = ctx.GlobalDuration(TxPoolRejournalFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolPriceLimitFlag.Name) {
		cfg.PriceLimit = ctx.GlobalUint64(TxPoolPriceLimitFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolPriceBumpFlag.Name) {
		cfg.PriceBump = ctx.GlobalUint64(TxPoolPriceBumpFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolAccountSlotsFlag.Name) {
		cfg.AccountSlots = ctx.GlobalUint64(TxPoolAccountSlotsFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolGlobalSlotsFlag.Name) {
		cfg.GlobalSlots = ctx.GlobalUint64(TxPoolGlobalSlotsFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolAccountQueueFlag.Name) {
		cfg.AccountQueue = ctx.GlobalUint64(TxPoolAccountQueueFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolGlobalQueueFlag.Name) {
		cfg.GlobalQueue = ctx.GlobalUint64(TxPoolGlobalQueueFlag.Name)
	}
	if ctx.GlobalIsSet(TxPoolLifetimeFlag.Name) {
		cfg.Lifetime = ctx.GlobalDuration(TxPoolLifetimeFlag.Name)
	}
}

func setEthash(ctx *cli.Context, cfg *ethconfig.Config) {
	if ctx.GlobalIsSet(EthashCacheDirFlag.Name) {
		cfg.Ethash.CacheDir = ctx.GlobalString(EthashCacheDirFlag.Name)
	}
	if ctx.GlobalIsSet(EthashDatasetDirFlag.Name) {
		cfg.Ethash.DatasetDir = ctx.GlobalString(EthashDatasetDirFlag.Name)
	}
	if ctx.GlobalIsSet(EthashCachesInMemoryFlag.Name) {
		cfg.Ethash.CachesInMem = ctx.GlobalInt(EthashCachesInMemoryFlag.Name)
	}
	if ctx.GlobalIsSet(EthashCachesOnDiskFlag.Name) {
		cfg.Ethash.CachesOnDisk = ctx.GlobalInt(EthashCachesOnDiskFlag.Name)
	}
	if ctx.GlobalIsSet(EthashCachesLockMmapFlag.Name) {
		cfg.Ethash.CachesLockMmap = ctx.GlobalBool(EthashCachesLockMmapFlag.Name)
	}
	if ctx.GlobalIsSet(EthashDatasetsInMemoryFlag.Name) {
		cfg.Ethash.DatasetsInMem = ctx.GlobalInt(EthashDatasetsInMemoryFlag.Name)
	}
	if ctx.GlobalIsSet(EthashDatasetsOnDiskFlag.Name) {
		cfg.Ethash.DatasetsOnDisk = ctx.GlobalInt(EthashDatasetsOnDiskFlag.Name)
	}
	if ctx.GlobalIsSet(EthashDatasetsLockMmapFlag.Name) {
		cfg.Ethash.DatasetsLockMmap = ctx.GlobalBool(EthashDatasetsLockMmapFlag.Name)
	}
}

func setMiner(ctx *cli.Context, cfg *miner.Config) {
	if ctx.GlobalIsSet(MinerNotifyFlag.Name) {
		cfg.Notify = strings.Split(ctx.GlobalString(MinerNotifyFlag.Name), ",")
	}
	cfg.NotifyFull = ctx.GlobalBool(MinerNotifyFullFlag.Name)
	if ctx.GlobalIsSet(MinerExtraDataFlag.Name) {
		cfg.ExtraData = []byte(ctx.GlobalString(MinerExtraDataFlag.Name))
	}
	if ctx.GlobalIsSet(MinerGasTargetFlag.Name) {
		cfg.GasFloor = ctx.GlobalUint64(MinerGasTargetFlag.Name)
	}
	if ctx.GlobalIsSet(MinerGasLimitFlag.Name) {
		cfg.GasCeil = ctx.GlobalUint64(MinerGasLimitFlag.Name)
	}
	if ctx.GlobalIsSet(MinerGasPriceFlag.Name) {
		cfg.GasPrice = GlobalBig(ctx, MinerGasPriceFlag.Name)
	}
	if ctx.GlobalIsSet(MinerRecommitIntervalFlag.Name) {
		cfg.Recommit = ctx.GlobalDuration(MinerRecommitIntervalFlag.Name)
	}
	if ctx.GlobalIsSet(MinerNoVerfiyFlag.Name) {
		cfg.Noverify = ctx.GlobalBool(MinerNoVerfiyFlag.Name)
	}
	if ctx.GlobalIsSet(AllowedFutureBlockTimeFlag.Name) {
		cfg.AllowedFutureBlockTime = ctx.GlobalUint64(AllowedFutureBlockTimeFlag.Name) //Quorum
	}
}

func setAuthorizationList(ctx *cli.Context, cfg *ethconfig.Config) {
	authorizationList := ctx.GlobalString(AuthorizationListFlag.Name)
	if authorizationList == "" {
		authorizationList = ctx.GlobalString(DeprecatedAuthorizationListFlag.Name)
		if authorizationList != "" {
			log.Warn("The flag --whitelist is deprecated and will be removed in the future, please use --authorizationlist")
		}
	}
	if authorizationList == "" {
		return
	}
	cfg.AuthorizationList = make(map[uint64]common.Hash)
	for _, entry := range strings.Split(authorizationList, ",") {
		parts := strings.Split(entry, "=")
		if len(parts) != 2 {
			Fatalf("Invalid authorized entry: %s", entry)
		}
		number, err := strconv.ParseUint(parts[0], 0, 64)
		if err != nil {
			Fatalf("Invalid authorized block number %s: %v", parts[0], err)
		}
		var hash common.Hash
		if err = hash.UnmarshalText([]byte(parts[1])); err != nil {
			Fatalf("Invalid authorized hash %s: %v", parts[1], err)
		}
		cfg.AuthorizationList[number] = hash
	}
}

// Quorum
func setIstanbul(ctx *cli.Context, cfg *eth.Config) {
	if ctx.GlobalIsSet(IstanbulRequestTimeoutFlag.Name) {
		log.Warn("WARNING: The flag --istanbul.requesttimeout is deprecated and will be removed in the future, please use ibft.requesttimeoutseconds on genesis file")
		cfg.Istanbul.RequestTimeout = ctx.GlobalUint64(IstanbulRequestTimeoutFlag.Name)
	}
	if ctx.GlobalIsSet(IstanbulBlockPeriodFlag.Name) {
		log.Warn("WARNING: The flag --istanbul.blockperiod is deprecated and will be removed in the future, please use ibft.blockperiodseconds on genesis file")
		cfg.Istanbul.BlockPeriod = ctx.GlobalUint64(IstanbulBlockPeriodFlag.Name)
	}
}

func setRaft(ctx *cli.Context, cfg *eth.Config) {
	cfg.RaftMode = ctx.GlobalBool(RaftModeFlag.Name)
}

func setQuorumConfig(ctx *cli.Context, cfg *eth.Config) error {
	cfg.EVMCallTimeOut = time.Duration(ctx.GlobalInt(EVMCallTimeOutFlag.Name)) * time.Second
	cfg.QuorumChainConfig = core.NewQuorumChainConfig(ctx.GlobalBool(MultitenancyFlag.Name),
		ctx.GlobalBool(RevertReasonFlag.Name), ctx.GlobalBool(QuorumEnablePrivacyMarker.Name),
		ctx.GlobalBool(QuorumEnablePrivateTrieCache.Name))
	setIstanbul(ctx, cfg)
	setRaft(ctx, cfg)
	return nil
}

// CheckExclusive verifies that only a single instance of the provided flags was
// set by the user. Each flag might optionally be followed by a string type to
// specialize it further.
func CheckExclusive(ctx *cli.Context, args ...interface{}) {
	set := make([]string, 0, 1)
	for i := 0; i < len(args); i++ {
		// Make sure the next argument is a flag and skip if not set
		flag, ok := args[i].(cli.Flag)
		if !ok {
			panic(fmt.Sprintf("invalid argument, not cli.Flag type: %T", args[i]))
		}
		// Check if next arg extends current and expand its name if so
		name := flag.GetName()

		if i+1 < len(args) {
			switch option := args[i+1].(type) {
			case string:
				// Extended flag check, make sure value set doesn't conflict with passed in option
				if ctx.GlobalString(flag.GetName()) == option {
					name += "=" + option
					set = append(set, "--"+name)
				}
				// shift arguments and continue
				i++
				continue

			case cli.Flag:
			default:
				panic(fmt.Sprintf("invalid argument, not cli.Flag or string extension: %T", args[i+1]))
			}
		}
		// Mark the flag if it's set
		if ctx.GlobalIsSet(flag.GetName()) {
			set = append(set, "--"+name)
		}
	}
	if len(set) > 1 {
		Fatalf("Flags %v can't be used at the same time", strings.Join(set, ", "))
	}
}

func SetQLightConfig(ctx *cli.Context, nodeCfg *node.Config, ethCfg *ethconfig.Config) {
	if ctx.GlobalIsSet(QuorumLightServerFlag.Name) {
		ethCfg.QuorumLightServer = ctx.GlobalBool(QuorumLightServerFlag.Name)
	}

	if ethCfg.QuorumLightServer {
		if nodeCfg.QP2P == nil {
			nodeCfg.QP2P = &p2p.Config{
				ListenAddr:  ":30305",
				MaxPeers:    10,
				NAT:         nil,
				NoDial:      true,
				NoDiscovery: true,
			}
			SetQP2PConfig(ctx, nodeCfg.QP2P)
		}
	} else {
		nodeCfg.QP2P = nil
	}

	ethCfg.QuorumLightClient = &ethconfig.QuorumLightClient{}
	if ctx.GlobalIsSet(QuorumLightClientFlag.Name) {
		ethCfg.QuorumLightClient.Use = ctx.GlobalBool(QuorumLightClientFlag.Name)
	}

	if len(ethCfg.QuorumLightClient.PSI) == 0 {
		ethCfg.QuorumLightClient.PSI = "private"
	}
	if ctx.GlobalIsSet(QuorumLightClientPSIFlag.Name) {
		ethCfg.QuorumLightClient.PSI = ctx.GlobalString(QuorumLightClientPSIFlag.Name)
	}

	if ctx.GlobalIsSet(QuorumLightClientTokenEnabledFlag.Name) {
		ethCfg.QuorumLightClient.TokenEnabled = ctx.GlobalBool(QuorumLightClientTokenEnabledFlag.Name)
	}

	if ctx.GlobalIsSet(QuorumLightClientTokenValueFlag.Name) {
		ethCfg.QuorumLightClient.TokenValue = ctx.GlobalString(QuorumLightClientTokenValueFlag.Name)
	}

	if len(ethCfg.QuorumLightClient.TokenManagement) == 0 {
		ethCfg.QuorumLightClient.TokenManagement = "client-security-plugin"
	}
	if ctx.GlobalIsSet(QuorumLightClientTokenManagementFlag.Name) {
		ethCfg.QuorumLightClient.TokenManagement = ctx.GlobalString(QuorumLightClientTokenManagementFlag.Name)
	}
	if !isValidTokenManagement(ethCfg.QuorumLightClient.TokenManagement) {
		Fatalf("Invalid value specified '%s' for flag '%s'.", ethCfg.QuorumLightClient.TokenManagement, QuorumLightClientTokenManagementFlag.Name)
	}

	if ctx.GlobalIsSet(QuorumLightClientRPCTLSFlag.Name) {
		ethCfg.QuorumLightClient.RPCTLS = ctx.GlobalBool(QuorumLightClientRPCTLSFlag.Name)
	}

	if ctx.GlobalIsSet(QuorumLightClientRPCTLSCACertFlag.Name) {
		ethCfg.QuorumLightClient.RPCTLSCACert = ctx.GlobalString(QuorumLightClientRPCTLSCACertFlag.Name)
	}

	if ctx.GlobalIsSet(QuorumLightClientRPCTLSInsecureSkipVerifyFlag.Name) {
		ethCfg.QuorumLightClient.RPCTLSInsecureSkipVerify = ctx.GlobalBool(QuorumLightClientRPCTLSInsecureSkipVerifyFlag.Name)
	}

	if ctx.GlobalIsSet(QuorumLightClientRPCTLSCertFlag.Name) && ctx.GlobalIsSet(QuorumLightClientRPCTLSKeyFlag.Name) {
		ethCfg.QuorumLightClient.RPCTLSCert = ctx.GlobalString(QuorumLightClientRPCTLSCertFlag.Name)
		ethCfg.QuorumLightClient.RPCTLSKey = ctx.GlobalString(QuorumLightClientRPCTLSKeyFlag.Name)
	} else if ctx.GlobalIsSet(QuorumLightClientRPCTLSCertFlag.Name) {
		Fatalf("'%s' specified without specifying '%s'", QuorumLightClientRPCTLSCertFlag.Name, QuorumLightClientRPCTLSKeyFlag.Name)
	} else if ctx.GlobalIsSet(QuorumLightClientRPCTLSKeyFlag.Name) {
		Fatalf("'%s' specified without specifying '%s'", QuorumLightClientRPCTLSKeyFlag.Name, QuorumLightClientRPCTLSCertFlag.Name)
	}

	if ctx.GlobalIsSet(QuorumLightClientServerNodeRPCFlag.Name) {
		ethCfg.QuorumLightClient.ServerNodeRPC = ctx.GlobalString(QuorumLightClientServerNodeRPCFlag.Name)
	}

	if ctx.GlobalIsSet(QuorumLightClientServerNodeFlag.Name) {
		ethCfg.QuorumLightClient.ServerNode = ctx.GlobalString(QuorumLightClientServerNodeFlag.Name)
		// This is already done in geth/config - before the node.New invocation (at which point the StaticNodes is already copied)
		//stack.Config().P2P.StaticNodes = []*enode.Node{enode.MustParse(ethCfg.QuorumLightClientServerNode)}
	}

	if ethCfg.QuorumLightClient.Enabled() {
		if ctx.GlobalBool(MiningEnabledFlag.Name) {
			Fatalf("QLight clients do not support mining")
		}
		if len(ethCfg.QuorumLightClient.ServerNode) == 0 {
			Fatalf("Please specify the '%s' when running a qlight client.", QuorumLightClientServerNodeFlag.Name)
		}
		if len(ethCfg.QuorumLightClient.ServerNodeRPC) == 0 {
			Fatalf("Please specify the '%s' when running a qlight client.", QuorumLightClientServerNodeRPCFlag.Name)
		}

		nodeCfg.P2P.StaticNodes = []*enode.Node{enode.MustParse(ethCfg.QuorumLightClient.ServerNode)}
		log.Info("The node is configured to run as a qlight client. 'maxpeers' is overridden to `1` and the P2P listener is disabled.")
		nodeCfg.P2P.MaxPeers = 1
		// force the qlight client node to disable the local P2P listener
		nodeCfg.P2P.ListenAddr = ""
	}
}

// SetEthConfig applies eth-related command line flags to the config.
func SetEthConfig(ctx *cli.Context, stack *node.Node, cfg *ethconfig.Config) {
	// Avoid conflicting network flags
	CheckExclusive(ctx, MainnetFlag, DeveloperFlag, RopstenFlag, RinkebyFlag, GoerliFlag, YoloV3Flag)
	CheckExclusive(ctx, LightServeFlag, SyncModeFlag, "light")
	CheckExclusive(ctx, DeveloperFlag, ExternalSignerFlag) // Can't use both ephemeral unlocked and external signer
	if ctx.GlobalString(GCModeFlag.Name) == "archive" && ctx.GlobalUint64(TxLookupLimitFlag.Name) != 0 {
		ctx.GlobalSet(TxLookupLimitFlag.Name, "0")
		log.Warn("Disable transaction unindexing for archive node")
	}
	if ctx.GlobalIsSet(LightServeFlag.Name) && ctx.GlobalUint64(TxLookupLimitFlag.Name) != 0 {
		log.Warn("LES server cannot serve old transaction status and cannot connect below les/4 protocol version if transaction lookup index is limited")
	}
	var ks *keystore.KeyStore
	if keystores := stack.AccountManager().Backends(keystore.KeyStoreType); len(keystores) > 0 {
		ks = keystores[0].(*keystore.KeyStore)
	}
	setEtherbase(ctx, ks, cfg)
	setGPO(ctx, &cfg.GPO, ctx.GlobalString(SyncModeFlag.Name) == "light")
	setTxPool(ctx, &cfg.TxPool)
	setEthash(ctx, cfg)
	setMiner(ctx, &cfg.Miner)
	setAuthorizationList(ctx, cfg)
	setLes(ctx, cfg)

	// Cap the cache allowance and tune the garbage collector
	mem, err := gopsutil.VirtualMemory()
	if err == nil {
		if 32<<(^uintptr(0)>>63) == 32 && mem.Total > 2*1024*1024*1024 {
			log.Warn("Lowering memory allowance on 32bit arch", "available", mem.Total/1024/1024, "addressable", 2*1024)
			mem.Total = 2 * 1024 * 1024 * 1024
		}
		allowance := int(mem.Total / 1024 / 1024 / 3)
		if cache := ctx.GlobalInt(CacheFlag.Name); cache > allowance {
			log.Warn("Sanitizing cache to Go's GC limits", "provided", cache, "updated", allowance)
			ctx.GlobalSet(CacheFlag.Name, strconv.Itoa(allowance))
		}
	}
	// Ensure Go's GC ignores the database cache for trigger percentage
	cache := ctx.GlobalInt(CacheFlag.Name)
	gogc := math.Max(20, math.Min(100, 100/(float64(cache)/1024)))

	log.Debug("Sanitizing Go's GC trigger", "percent", int(gogc))
	godebug.SetGCPercent(int(gogc))

	// Quorum
	err = setQuorumConfig(ctx, cfg)
	if err != nil {
		Fatalf("Quorum configuration has an error: %v", err)
	}

	if ctx.GlobalIsSet(SyncModeFlag.Name) {
		cfg.SyncMode = *GlobalTextMarshaler(ctx, SyncModeFlag.Name).(*downloader.SyncMode)
	}

	// Quorum
	if cfg.QuorumLightClient.Enabled() && cfg.SyncMode != downloader.FullSync {
		Fatalf("Only the 'full' syncmode is supported for the qlight client.")
	}
	if private.IsQuorumPrivacyEnabled() && cfg.SyncMode != downloader.FullSync {
		Fatalf("Only the 'full' syncmode is supported when quorum privacy is enabled.")
	}
	// End Quorum

	if ctx.GlobalIsSet(NetworkIdFlag.Name) {
		cfg.NetworkId = ctx.GlobalUint64(NetworkIdFlag.Name)
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheDatabaseFlag.Name) {
		cfg.DatabaseCache = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheDatabaseFlag.Name) / 100
	}
	cfg.DatabaseHandles = MakeDatabaseHandles()
	if ctx.GlobalIsSet(AncientFlag.Name) {
		cfg.DatabaseFreezer = ctx.GlobalString(AncientFlag.Name)
	}

	if gcmode := ctx.GlobalString(GCModeFlag.Name); gcmode != "full" && gcmode != "archive" {
		Fatalf("--%s must be either 'full' or 'archive'", GCModeFlag.Name)
	}
	if ctx.GlobalIsSet(GCModeFlag.Name) {
		cfg.NoPruning = ctx.GlobalString(GCModeFlag.Name) == "archive"
	}
	if ctx.GlobalIsSet(CacheNoPrefetchFlag.Name) {
		cfg.NoPrefetch = ctx.GlobalBool(CacheNoPrefetchFlag.Name)
	}
	// Read the value from the flag no matter if it's set or not.
	cfg.Preimages = ctx.GlobalBool(CachePreimagesFlag.Name)
	if true || cfg.NoPruning && !cfg.Preimages { // TODO: Quorum; force preimages for contract extension and dump of states compatibility, until a fix is found
		cfg.Preimages = true
		log.Info("Enabling recording of key preimages since archive mode is used")
	}
	if ctx.GlobalIsSet(TxLookupLimitFlag.Name) {
		cfg.TxLookupLimit = ctx.GlobalUint64(TxLookupLimitFlag.Name)
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheTrieFlag.Name) {
		cfg.TrieCleanCache = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheTrieFlag.Name) / 100
	}
	if ctx.GlobalIsSet(CacheTrieJournalFlag.Name) {
		cfg.TrieCleanCacheJournal = ctx.GlobalString(CacheTrieJournalFlag.Name)
	}
	if ctx.GlobalIsSet(CacheTrieRejournalFlag.Name) {
		cfg.TrieCleanCacheRejournal = ctx.GlobalDuration(CacheTrieRejournalFlag.Name)
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheGCFlag.Name) {
		cfg.TrieDirtyCache = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheGCFlag.Name) / 100
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheSnapshotFlag.Name) {
		cfg.SnapshotCache = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheSnapshotFlag.Name) / 100
	}
	if !ctx.GlobalBool(SnapshotFlag.Name) {
		// If snap-sync is requested, this flag is also required
		if cfg.SyncMode == downloader.SnapSync {
			log.Info("Snap sync requested, enabling --snapshot")
		} else {
			cfg.TrieCleanCache += cfg.SnapshotCache
			cfg.SnapshotCache = 0 // Disabled
		}
	}
	if ctx.GlobalIsSet(DocRootFlag.Name) {
		cfg.DocRoot = ctx.GlobalString(DocRootFlag.Name)
	}
	if ctx.GlobalIsSet(VMEnableDebugFlag.Name) {
		// TODO(fjl): force-enable this in --dev mode
		cfg.EnablePreimageRecording = ctx.GlobalBool(VMEnableDebugFlag.Name)
	}

	if ctx.GlobalIsSet(EWASMInterpreterFlag.Name) {
		cfg.EWASMInterpreter = ctx.GlobalString(EWASMInterpreterFlag.Name)
	}

	if ctx.GlobalIsSet(EVMInterpreterFlag.Name) {
		cfg.EVMInterpreter = ctx.GlobalString(EVMInterpreterFlag.Name)
	}
	if ctx.GlobalIsSet(RPCGlobalGasCapFlag.Name) {
		cfg.RPCGasCap = ctx.GlobalUint64(RPCGlobalGasCapFlag.Name)
	}
	if cfg.RPCGasCap != 0 {
		log.Info("Set global gas cap", "cap", cfg.RPCGasCap)
	} else {
		log.Info("Global gas cap disabled")
	}
	if ctx.GlobalIsSet(RPCGlobalTxFeeCapFlag.Name) {
		cfg.RPCTxFeeCap = ctx.GlobalFloat64(RPCGlobalTxFeeCapFlag.Name)
	}
	if ctx.GlobalIsSet(NoDiscoverFlag.Name) {
		cfg.EthDiscoveryURLs, cfg.SnapDiscoveryURLs = []string{}, []string{}
	} else if ctx.GlobalIsSet(DNSDiscoveryFlag.Name) {
		urls := ctx.GlobalString(DNSDiscoveryFlag.Name)
		if urls == "" {
			cfg.EthDiscoveryURLs = []string{}
		} else {
			cfg.EthDiscoveryURLs = SplitAndTrim(urls)
		}
	}

	// set immutability threshold in config
	params.SetQuorumImmutabilityThreshold(ctx.GlobalInt(QuorumImmutabilityThreshold.Name))

	// Override any default configs for hard coded networks.
	switch {
	case ctx.GlobalBool(MainnetFlag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 1
		}
		cfg.Genesis = core.DefaultGenesisBlock()
		SetDNSDiscoveryDefaults(cfg, params.MainnetGenesisHash)
	case ctx.GlobalBool(RopstenFlag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 3
		}
		cfg.Genesis = core.DefaultRopstenGenesisBlock()
		SetDNSDiscoveryDefaults(cfg, params.RopstenGenesisHash)
	case ctx.GlobalBool(RinkebyFlag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 4
		}
		cfg.Genesis = core.DefaultRinkebyGenesisBlock()
		SetDNSDiscoveryDefaults(cfg, params.RinkebyGenesisHash)
	case ctx.GlobalBool(GoerliFlag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 5
		}
		cfg.Genesis = core.DefaultGoerliGenesisBlock()
		SetDNSDiscoveryDefaults(cfg, params.GoerliGenesisHash)
	case ctx.GlobalBool(YoloV3Flag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = new(big.Int).SetBytes([]byte("yolov3x")).Uint64() // "yolov3x"
		}
		cfg.Genesis = core.DefaultYoloV3GenesisBlock()
	case ctx.GlobalBool(DeveloperFlag.Name):
		if !ctx.GlobalIsSet(NetworkIdFlag.Name) {
			cfg.NetworkId = 1337
		}
		// Create new developer account or reuse existing one
		var (
			developer  accounts.Account
			passphrase string
			err        error
		)
		if list := MakePasswordList(ctx); len(list) > 0 {
			// Just take the first value. Although the function returns a possible multiple values and
			// some usages iterate through them as attempts, that doesn't make sense in this setting,
			// when we're definitely concerned with only one account.
			passphrase = list[0]
		}
		// setEtherbase has been called above, configuring the miner address from command line flags.
		if cfg.Miner.Etherbase != (common.Address{}) {
			developer = accounts.Account{Address: cfg.Miner.Etherbase}
		} else if accs := ks.Accounts(); len(accs) > 0 {
			developer = ks.Accounts()[0]
		} else {
			developer, err = ks.NewAccount(passphrase)
			if err != nil {
				Fatalf("Failed to create developer account: %v", err)
			}
		}
		if err := ks.Unlock(developer, passphrase); err != nil {
			Fatalf("Failed to unlock developer account: %v", err)
		}
		log.Info("Using developer account", "address", developer.Address)

		// Create a new developer genesis block or reuse existing one
		cfg.Genesis = core.DeveloperGenesisBlock(uint64(ctx.GlobalInt(DeveloperPeriodFlag.Name)), developer.Address)
		if ctx.GlobalIsSet(DataDirFlag.Name) {
			// Check if we have an already initialized chain and fall back to
			// that if so. Otherwise we need to generate a new genesis spec.
			chaindb := MakeChainDatabase(ctx, stack, true)
			if rawdb.ReadCanonicalHash(chaindb, 0) != (common.Hash{}) {
				cfg.Genesis = nil // fallback to db content
			}
			chaindb.Close()
		}
		if !ctx.GlobalIsSet(MinerGasPriceFlag.Name) {
			cfg.Miner.GasPrice = big.NewInt(1)
		}
	default:
		if cfg.NetworkId == 1 {
			SetDNSDiscoveryDefaults(cfg, params.MainnetGenesisHash)
		}
	}
}

// SetDNSDiscoveryDefaults configures DNS discovery with the given URL if
// no URLs are set.
func SetDNSDiscoveryDefaults(cfg *ethconfig.Config, genesis common.Hash) {
	if cfg.EthDiscoveryURLs != nil {
		return // already set through flags/config
	}
	protocol := "all"
	if cfg.SyncMode == downloader.LightSync {
		protocol = "les"
	}
	if url := params.KnownDNSNetwork(genesis, protocol); url != "" {
		cfg.EthDiscoveryURLs = []string{url}
	}
	if cfg.SyncMode == downloader.SnapSync {
		if url := params.KnownDNSNetwork(genesis, "snap"); url != "" {
			cfg.SnapDiscoveryURLs = []string{url}
		}
	}
}

// RegisterEthService adds an Ethereum client to the stack.
// Quorum => returns also the ethereum service which is used by the raft service
func RegisterEthService(stack *node.Node, cfg *ethconfig.Config) (ethapi.Backend, *eth.Ethereum) {
	if cfg.SyncMode == downloader.LightSync {
		backend, err := les.New(stack, cfg)
		if err != nil {
			Fatalf("Failed to register the Ethereum service: %v", err)
		}
		stack.RegisterAPIs(tracers.APIs(backend.ApiBackend))
		return backend.ApiBackend, nil
	}

	// Quorum
	client, err := stack.Attach()
	if err != nil {
		Fatalf("Failed to attach to self: %v", err)
	}
	cfg.Istanbul.Client = ethclient.NewClient(client)
	// End Quorum

	backend, err := eth.New(stack, cfg)
	if err != nil {
		Fatalf("Failed to register the Ethereum service: %v", err)
	}
	if cfg.LightServ > 0 {
		_, err := les.NewLesServer(stack, backend, cfg)
		if err != nil {
			Fatalf("Failed to create the LES server: %v", err)
		}
	}
	stack.RegisterAPIs(tracers.APIs(backend.APIBackend))
	return backend.APIBackend, backend
}

// RegisterEthStatsService configures the Ethereum Stats daemon and adds it to
// the given node.
func RegisterEthStatsService(stack *node.Node, backend ethapi.Backend, url string) {
	if err := ethstats.New(stack, backend, backend.Engine(), url); err != nil {
		Fatalf("Failed to register the Ethereum Stats service: %v", err)
	}
}

// RegisterGraphQLService is a utility function to construct a new service and register it against a node.
func RegisterGraphQLService(stack *node.Node, backend ethapi.Backend, cfg node.Config) {
	if err := graphql.New(stack, backend, cfg.GraphQLCors, cfg.GraphQLVirtualHosts); err != nil {
		Fatalf("Failed to register the GraphQL service: %v", err)
	}
}

// Quorum
//
// Register plugin manager as a service in geth
func RegisterPluginService(stack *node.Node, cfg *node.Config, skipVerify bool, localVerify bool, publicKey string) {
	// ricardolyn: I can't adapt this Plugin Service construction to the new approach as there are circular dependencies between Node and Plugin
	if err := cfg.ResolvePluginBaseDir(); err != nil {
		Fatalf("plugins: unable to resolve plugin base dir due to %s", err)
	}
	pluginManager, err := plugin.NewPluginManager(cfg.UserIdent, cfg.Plugins, skipVerify, localVerify, publicKey)
	if err != nil {
		Fatalf("plugins: Failed to register the Plugins service: %v", err)
	}
	stack.SetPluginManager(pluginManager)
	stack.RegisterAPIs(pluginManager.APIs())
	stack.RegisterLifecycle(pluginManager)
	log.Info("plugin service registered")
}

// Configure smart-contract-based permissioning service
func RegisterPermissionService(stack *node.Node, useDns bool, chainID *big.Int) {
	permissionConfig, err := types.ParsePermissionConfig(stack.DataDir())
	if err != nil {
		Fatalf("loading of %s failed due to %v", params.PERMISSION_MODEL_CONFIG, err)
	}
	// start the permissions management service
	_, err = permission.NewQuorumPermissionCtrl(stack, &permissionConfig, useDns, chainID)
	if err != nil {
		Fatalf("failed to load the permission contracts as given in %s due to %v", params.PERMISSION_MODEL_CONFIG, err)
	}
	log.Info("permission service registered")
}

func RegisterRaftService(stack *node.Node, ctx *cli.Context, nodeCfg *node.Config, ethService *eth.Ethereum) {
	blockTimeMillis := ctx.GlobalInt(RaftBlockTimeFlag.Name)
	raftLogDir := nodeCfg.RaftLogDir // default value is set either 'datadir' or 'raftlogdir'
	joinExistingId := ctx.GlobalInt(RaftJoinExistingFlag.Name)
	useDns := ctx.GlobalBool(RaftDNSEnabledFlag.Name)
	raftPort := uint16(ctx.GlobalInt(RaftPortFlag.Name))

	privkey := nodeCfg.NodeKey()
	strId := enode.PubkeyToIDV4(&privkey.PublicKey).String()
	blockTimeNanos := time.Duration(blockTimeMillis) * time.Millisecond
	peers := nodeCfg.StaticNodes()

	var myId uint16
	var joinExisting bool

	if joinExistingId > 0 {
		myId = uint16(joinExistingId)
		joinExisting = true
	} else if len(peers) == 0 {
		Fatalf("Raft-based consensus requires either (1) an initial peers list (in static-nodes.json) including this enode hash (%v), or (2) the flag --raftjoinexisting RAFT_ID, where RAFT_ID has been issued by an existing cluster member calling `raft.addPeer(ENODE_ID)` with an enode ID containing this node's enode hash.", strId)
	} else {
		peerIds := make([]string, len(peers))

		for peerIdx, peer := range peers {
			if !peer.HasRaftPort() {
				Fatalf("raftport querystring parameter not specified in static-node enode ID: %v. please check your static-nodes.json file.", peer.String())
			}

			peerId := peer.ID().String()
			peerIds[peerIdx] = peerId
			if peerId == strId {
				myId = uint16(peerIdx) + 1
			}
		}

		if myId == 0 {
			Fatalf("failed to find local enode ID (%v) amongst peer IDs: %v", strId, peerIds)
		}
	}

	_, err := raft.New(stack, ethService.BlockChain().Config(), myId, raftPort, joinExisting, blockTimeNanos, ethService, peers, raftLogDir, useDns)
	if err != nil {
		Fatalf("raft: Failed to register the Raft service: %v", err)
	}

	log.Info("raft service registered")
}

func RegisterExtensionService(stack *node.Node, ethService *eth.Ethereum) {
	_, err := extension.NewServicesFactory(stack, private.P, ethService)
	if err != nil {
		Fatalf("Failed to register the Extension service: %v", err)
	}

	log.Info("extension service registered")
}

func SetupMetrics(ctx *cli.Context) {
	if metrics.Enabled {
		log.Info("Enabling metrics collection")

		var (
			enableExport = ctx.GlobalBool(MetricsEnableInfluxDBFlag.Name)
			endpoint     = ctx.GlobalString(MetricsInfluxDBEndpointFlag.Name)
			database     = ctx.GlobalString(MetricsInfluxDBDatabaseFlag.Name)
			username     = ctx.GlobalString(MetricsInfluxDBUsernameFlag.Name)
			password     = ctx.GlobalString(MetricsInfluxDBPasswordFlag.Name)
		)

		if enableExport {
			tagsMap := SplitTagsFlag(ctx.GlobalString(MetricsInfluxDBTagsFlag.Name))

			log.Info("Enabling metrics export to InfluxDB")

			go influxdb.InfluxDBWithTags(metrics.DefaultRegistry, 10*time.Second, endpoint, database, username, password, "geth.", tagsMap)
		}

		if ctx.GlobalIsSet(MetricsHTTPFlag.Name) {
			address := fmt.Sprintf("%s:%d", ctx.GlobalString(MetricsHTTPFlag.Name), ctx.GlobalInt(MetricsPortFlag.Name))
			log.Info("Enabling stand-alone metrics HTTP endpoint", "address", address)
			exp.Setup(address)
		}
	}
}

func SplitTagsFlag(tagsFlag string) map[string]string {
	tags := strings.Split(tagsFlag, ",")
	tagsMap := map[string]string{}

	for _, t := range tags {
		if t != "" {
			kv := strings.Split(t, "=")

			if len(kv) == 2 {
				tagsMap[kv[0]] = kv[1]
			}
		}
	}

	return tagsMap
}

// MakeChainDatabase open an LevelDB using the flags passed to the client and will hard crash if it fails.
func MakeChainDatabase(ctx *cli.Context, stack *node.Node, readonly bool) ethdb.Database {
	var (
		cache   = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheDatabaseFlag.Name) / 100
		handles = MakeDatabaseHandles()

		err     error
		chainDb ethdb.Database
	)
	if ctx.GlobalString(SyncModeFlag.Name) == "light" {
		name := "lightchaindata"
		chainDb, err = stack.OpenDatabase(name, cache, handles, "", readonly)
	} else {
		name := "chaindata"
		chainDb, err = stack.OpenDatabaseWithFreezer(name, cache, handles, ctx.GlobalString(AncientFlag.Name), "", readonly)
	}
	if err != nil {
		Fatalf("Could not open database: %v", err)
	}
	return chainDb
}

func MakeGenesis(ctx *cli.Context) *core.Genesis {
	var genesis *core.Genesis
	switch {
	case ctx.GlobalBool(MainnetFlag.Name):
		genesis = core.DefaultGenesisBlock()
	case ctx.GlobalBool(RopstenFlag.Name):
		genesis = core.DefaultRopstenGenesisBlock()
	case ctx.GlobalBool(RinkebyFlag.Name):
		genesis = core.DefaultRinkebyGenesisBlock()
	case ctx.GlobalBool(GoerliFlag.Name):
		genesis = core.DefaultGoerliGenesisBlock()
	case ctx.GlobalBool(YoloV3Flag.Name):
		genesis = core.DefaultYoloV3GenesisBlock()
	case ctx.GlobalBool(DeveloperFlag.Name):
		Fatalf("Developer chains are ephemeral")
	}
	return genesis
}

// MakeChain creates a chain manager from set command line flags.
func MakeChain(ctx *cli.Context, stack *node.Node, useExist bool) (chain *core.BlockChain, chainDb ethdb.Database) {
	var err error
	var config *params.ChainConfig
	chainDb = MakeChainDatabase(ctx, stack, false) // TODO(rjl493456442) support read-only database
	if useExist {
		stored := rawdb.ReadCanonicalHash(chainDb, 0)
		if (stored == common.Hash{}) {
			Fatalf("No existing genesis")
		}
		config = rawdb.ReadChainConfig(chainDb, stored)
	} else {
		config, _, err = core.SetupGenesisBlock(chainDb, MakeGenesis(ctx))
		if err != nil {
			Fatalf("%v", err)
		}
	}

	var engine consensus.Engine

	client, err := stack.Attach()
	if err != nil {
		Fatalf("Failed to attach to self: %v", err)
	}

	if config.Clique != nil {
		engine = clique.New(config.Clique, chainDb)
	} else if config.Istanbul != nil {
		log.Warn("WARNING: The attribute config.istanbul is deprecated and will be removed in the future, please use config.ibft on genesis file")
		// for IBFT
		istanbulConfig := istanbul.DefaultConfig
		if config.Istanbul.Epoch != 0 {
			istanbulConfig.Epoch = config.Istanbul.Epoch
		}
		istanbulConfig.ProposerPolicy = istanbul.NewProposerPolicy(istanbul.ProposerPolicyId(config.Istanbul.ProposerPolicy))
		istanbulConfig.Ceil2Nby3Block = config.Istanbul.Ceil2Nby3Block
		istanbulConfig.TestQBFTBlock = config.Istanbul.TestQBFTBlock
		if config.Transitions != nil && len(config.Transitions) != 0 {
			istanbulConfig.Transitions = config.Transitions
		}
		istanbulConfig.Client = ethclient.NewClient(client)
		engine = istanbulBackend.New(istanbulConfig, stack.GetNodeKey(), chainDb)
	} else if config.IBFT != nil {
		ibftConfig := setBFTConfig(config.IBFT.BFTConfig)
		ibftConfig.TestQBFTBlock = nil
		if config.Transitions != nil && len(config.Transitions) != 0 {
			ibftConfig.Transitions = config.Transitions
		}
		ibftConfig.Client = ethclient.NewClient(client)
		engine = istanbulBackend.New(ibftConfig, stack.GetNodeKey(), chainDb)
	} else if config.QBFT != nil {
		qbftConfig := setBFTConfig(config.QBFT.BFTConfig)
		qbftConfig.TestQBFTBlock = big.NewInt(0)
		if config.Transitions != nil && len(config.Transitions) != 0 {
			qbftConfig.Transitions = config.Transitions
		}
		if config.QBFT.ValidatorContractAddress != (common.Address{}) {
			qbftConfig.ValidatorContract = config.QBFT.ValidatorContractAddress
		}
		qbftConfig.Client = ethclient.NewClient(client)
		engine = istanbulBackend.New(qbftConfig, stack.GetNodeKey(), chainDb)
	} else if config.IsQuorum {
		// for Raft
		engine = ethash.NewFullFaker()
	} else {
		engine = ethash.NewFaker()
		if !ctx.GlobalBool(FakePoWFlag.Name) {
			engine = ethash.New(ethash.Config{
				CacheDir:         stack.ResolvePath(ethconfig.Defaults.Ethash.CacheDir),
				CachesInMem:      ethconfig.Defaults.Ethash.CachesInMem,
				CachesOnDisk:     ethconfig.Defaults.Ethash.CachesOnDisk,
				CachesLockMmap:   ethconfig.Defaults.Ethash.CachesLockMmap,
				DatasetDir:       stack.ResolvePath(ethconfig.Defaults.Ethash.DatasetDir),
				DatasetsInMem:    ethconfig.Defaults.Ethash.DatasetsInMem,
				DatasetsOnDisk:   ethconfig.Defaults.Ethash.DatasetsOnDisk,
				DatasetsLockMmap: ethconfig.Defaults.Ethash.DatasetsLockMmap,
			}, nil, false)
		}
	}
	if gcmode := ctx.GlobalString(GCModeFlag.Name); gcmode != "full" && gcmode != "archive" {
		Fatalf("--%s must be either 'full' or 'archive'", GCModeFlag.Name)
	}
	cache := &core.CacheConfig{
		TrieCleanLimit:      ethconfig.Defaults.TrieCleanCache,
		TrieCleanNoPrefetch: ctx.GlobalBool(CacheNoPrefetchFlag.Name),
		TrieDirtyLimit:      ethconfig.Defaults.TrieDirtyCache,
		TrieDirtyDisabled:   ctx.GlobalString(GCModeFlag.Name) == "archive",
		TrieTimeLimit:       ethconfig.Defaults.TrieTimeout,
		SnapshotLimit:       ethconfig.Defaults.SnapshotCache,
		Preimages:           ctx.GlobalBool(CachePreimagesFlag.Name),
	}
	if true || cache.TrieDirtyDisabled && !cache.Preimages { // TODO: Quorum; force preimages for contract extension and dump of states compatibility, until a fix is found
		cache.Preimages = true
		log.Info("Enabling recording of key preimages since archive mode is used")
	}
	if !ctx.GlobalBool(SnapshotFlag.Name) {
		cache.SnapshotLimit = 0 // Disabled
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheTrieFlag.Name) {
		cache.TrieCleanLimit = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheTrieFlag.Name) / 100
	}
	if ctx.GlobalIsSet(CacheFlag.Name) || ctx.GlobalIsSet(CacheGCFlag.Name) {
		cache.TrieDirtyLimit = ctx.GlobalInt(CacheFlag.Name) * ctx.GlobalInt(CacheGCFlag.Name) / 100
	}
	vmcfg := vm.Config{EnablePreimageRecording: ctx.GlobalBool(VMEnableDebugFlag.Name)}

	// Quorum
	var limit *uint64
	if ctx.GlobalIsSet(TxLookupLimitFlag.Name) {
		l := ctx.GlobalUint64(TxLookupLimitFlag.Name)
		limit = &l
	}
	// End Quorum

	// TODO(rjl493456442) disable snapshot generation/wiping if the chain is read only.
	// Disable transaction indexing/unindexing by default.
	// TODO should multiple private states work with import/export/inspect commands
	chain, err = core.NewBlockChain(chainDb, cache, config, engine, vmcfg, nil, limit, nil)
	if err != nil {
		Fatalf("Can't create BlockChain: %v", err)
	}
	return chain, chainDb
}

func setBFTConfig(bftConfig *params.BFTConfig) *istanbul.Config {
	istanbulConfig := istanbul.DefaultConfig
	if bftConfig.BlockPeriodSeconds != 0 {
		istanbulConfig.BlockPeriod = bftConfig.BlockPeriodSeconds
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
	return istanbulConfig
}

// MakeConsolePreloads retrieves the absolute paths for the console JavaScript
// scripts to preload before starting.
func MakeConsolePreloads(ctx *cli.Context) []string {
	// Skip preloading if there's nothing to preload
	if ctx.GlobalString(PreloadJSFlag.Name) == "" {
		return nil
	}
	// Otherwise resolve absolute paths and return them
	var preloads []string

	for _, file := range strings.Split(ctx.GlobalString(PreloadJSFlag.Name), ",") {
		preloads = append(preloads, strings.TrimSpace(file))
	}
	return preloads
}

// MigrateFlags sets the global flag from a local flag when it's set.
// This is a temporary function used for migrating old command/flags to the
// new format.
//
// e.g. geth account new --keystore /tmp/mykeystore --lightkdf
//
// is equivalent after calling this method with:
//
// geth --keystore /tmp/mykeystore --lightkdf account new
//
// This allows the use of the existing configuration functionality.
// When all flags are migrated this function can be removed and the existing
// configuration functionality must be changed that is uses local flags
func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
				ctx.GlobalSet(name, ctx.String(name))
			}
		}
		return action(ctx)
	}
}

func isValidTokenManagement(value string) bool {
	switch value {
	case
		"none",
		"external",
		"client-security-plugin":
		return true
	}
	return false
}
