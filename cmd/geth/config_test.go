package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/netutil"
	"github.com/naoina/toml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/urfave/cli.v1"
)

func TestFlagsConfig(t *testing.T) {
	flags := []interface{}{
		utils.DataDirFlag,
		utils.RaftLogDirFlag,
		utils.AncientFlag,
		utils.MinFreeDiskSpaceFlag,
		utils.KeyStoreDirFlag,
		utils.NoUSBFlag,
		utils.USBFlag,
		utils.SmartCardDaemonPathFlag,
		utils.NetworkIdFlag,
		utils.MainnetFlag,
		utils.GoerliFlag,
		utils.YoloV3Flag,
		utils.RinkebyFlag,
		utils.RopstenFlag,
		utils.DeveloperFlag,
		utils.DeveloperPeriodFlag,
		utils.IdentityFlag,
		utils.DocRootFlag,
		utils.ExitWhenSyncedFlag,
		utils.IterativeOutputFlag,
		utils.ExcludeStorageFlag,
		utils.IncludeIncompletesFlag,
		utils.ExcludeCodeFlag,
		utils.SyncModeFlag,
		utils.GCModeFlag,
		utils.SnapshotFlag,
		utils.TxLookupLimitFlag,
		utils.LightKDFFlag,
		utils.DeprecatedAuthorizationListFlag,
		utils.AuthorizationListFlag,
		utils.BloomFilterSizeFlag,
		utils.OverrideBerlinFlag,
		utils.LightServeFlag,
		utils.LightIngressFlag,
		utils.LightEgressFlag,
		utils.LightMaxPeersFlag,
		utils.UltraLightServersFlag,
		utils.UltraLightFractionFlag,
		utils.UltraLightOnlyAnnounceFlag,
		utils.LightNoPruneFlag,
		utils.LightNoSyncServeFlag,
		utils.EthashCacheDirFlag,
		utils.EthashCachesInMemoryFlag,
		utils.EthashCachesOnDiskFlag,
		utils.EthashCachesLockMmapFlag,
		utils.EthashDatasetDirFlag,
		utils.EthashDatasetsInMemoryFlag,
		utils.EthashDatasetsOnDiskFlag,
		utils.EthashDatasetsLockMmapFlag,
		utils.TxPoolLocalsFlag,
		utils.TxPoolNoLocalsFlag,
		utils.TxPoolJournalFlag,
		utils.TxPoolRejournalFlag,
		utils.TxPoolPriceLimitFlag,
		utils.TxPoolPriceBumpFlag,
		utils.TxPoolAccountSlotsFlag,
		utils.TxPoolGlobalSlotsFlag,
		utils.TxPoolAccountQueueFlag,
		utils.TxPoolGlobalQueueFlag,
		utils.TxPoolLifetimeFlag,
		utils.CacheFlag,
		utils.CacheDatabaseFlag,
		utils.CacheTrieFlag,
		utils.CacheTrieJournalFlag,
		utils.CacheTrieRejournalFlag,
		utils.CacheGCFlag,
		utils.CacheSnapshotFlag,
		utils.CacheNoPrefetchFlag,
		utils.CachePreimagesFlag,
		utils.MiningEnabledFlag,
		utils.MinerThreadsFlag,
		utils.MinerNotifyFlag,
		utils.MinerGasTargetFlag,
		utils.MinerGasLimitFlag,
		utils.MinerGasPriceFlag,
		utils.MinerEtherbaseFlag,
		utils.MinerExtraDataFlag,
		utils.MinerRecommitIntervalFlag,
		utils.MinerNoVerfiyFlag,
		utils.UnlockedAccountFlag,
		utils.PasswordFileFlag,
		utils.ExternalSignerFlag,
		utils.VMEnableDebugFlag,
		utils.InsecureUnlockAllowedFlag,
		utils.RPCGlobalGasCapFlag,
		utils.RPCGlobalTxFeeCapFlag,
		utils.EthStatsURLFlag,
		utils.FakePoWFlag,
		utils.NoCompactionFlag,
		utils.RPCClientToken,
		utils.RPCClientTLSCert,
		utils.RPCClientTLSCaCert,
		utils.RPCClientTLSCipherSuites,
		utils.RPCClientTLSInsecureSkipVerify,
		utils.IPCDisabledFlag,
		utils.IPCPathFlag,
		utils.HTTPEnabledFlag,
		utils.HTTPListenAddrFlag,
		utils.HTTPPortFlag,
		utils.HTTPCORSDomainFlag,
		utils.HTTPVirtualHostsFlag,
		utils.HTTPApiFlag,
		utils.HTTPPathPrefixFlag,
		utils.GraphQLEnabledFlag,
		utils.GraphQLCORSDomainFlag,
		utils.GraphQLVirtualHostsFlag,
		utils.WSEnabledFlag,
		utils.WSListenAddrFlag,
		utils.WSPortFlag,
		utils.WSApiFlag,
		utils.WSAllowedOriginsFlag,
		utils.WSPathPrefixFlag,
		utils.ExecFlag,
		utils.PreloadJSFlag,
		utils.AllowUnprotectedTxs,
		utils.MaxPeersFlag,
		utils.MaxPendingPeersFlag,
		utils.ListenPortFlag,
		utils.BootnodesFlag,
		utils.NodeKeyFileFlag,
		utils.NodeKeyHexFlag,
		utils.NATFlag,
		utils.NoDiscoverFlag,
		utils.DiscoveryV5Flag,
		utils.NetrestrictFlag,
		utils.DNSDiscoveryFlag,
		utils.JSpathFlag,
		utils.GpoBlocksFlag,
		utils.GpoPercentileFlag,
		utils.GpoMaxGasPriceFlag,
		utils.MetricsEnabledFlag,
		utils.MetricsEnabledExpensiveFlag,
		utils.MetricsHTTPFlag,
		utils.MetricsPortFlag,
		utils.MetricsEnableInfluxDBFlag,
		utils.MetricsInfluxDBEndpointFlag,
		utils.MetricsInfluxDBDatabaseFlag,
		utils.MetricsInfluxDBUsernameFlag,
		utils.MetricsInfluxDBPasswordFlag,
		utils.MetricsInfluxDBTagsFlag,
		utils.EWASMInterpreterFlag,
		utils.EVMInterpreterFlag,
		utils.EVMCallTimeOutFlag,
		utils.QuorumImmutabilityThreshold,
		utils.RaftModeFlag,
		utils.RaftBlockTimeFlag,
		utils.RaftJoinExistingFlag,
		utils.EmitCheckpointsFlag,
		utils.RaftPortFlag,
		utils.RaftDNSEnabledFlag,
		utils.EnableNodePermissionFlag,
		utils.AllowedFutureBlockTimeFlag,
		utils.PluginSettingsFlag,
		utils.PluginLocalVerifyFlag,
		utils.PluginPublicKeyFlag,
		utils.PluginSkipVerifyFlag,
		utils.AccountPluginNewAccountConfigFlag,
		utils.IstanbulRequestTimeoutFlag,
		utils.IstanbulBlockPeriodFlag,
		utils.MultitenancyFlag,
		utils.RevertReasonFlag,
		utils.QuorumEnablePrivacyMarker,
		utils.QuorumPTMUnixSocketFlag,
		utils.QuorumPTMUrlFlag,
		utils.QuorumPTMTimeoutFlag,
		utils.QuorumPTMDialTimeoutFlag,
		utils.QuorumPTMHttpIdleTimeoutFlag,
		utils.QuorumPTMHttpWriteBufferSizeFlag,
		utils.QuorumPTMHttpReadBufferSizeFlag,
		utils.QuorumPTMTlsModeFlag,
		utils.QuorumPTMTlsRootCaFlag,
		utils.QuorumPTMTlsClientCertFlag,
		utils.QuorumPTMTlsClientKeyFlag,
		utils.QuorumPTMTlsInsecureSkipVerify,
	}
	nodeKeyFile, err := ioutil.TempFile("/tmp", "nodekey")
	require.NoError(t, err)
	defer os.Remove(nodeKeyFile.Name())

	_, err = nodeKeyFile.WriteString("0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF\n")
	require.NoError(t, err)

	err = nodeKeyFile.Close()
	require.NoError(t, err)

	set := flag.NewFlagSet("dumpconfig", flag.ContinueOnError)
	for _, f := range flags {
		switch f := f.(type) {
		case utils.DirectoryFlag:
			set.String(f.Name, f.Value.String()+"/custom", f.Usage)
		case cli.BoolFlag:
			set.Bool(f.Name, true, f.Usage)
		case cli.BoolTFlag:
			set.Bool(f.Name, false, f.Usage)
		case cli.StringFlag:
			switch f {
			case utils.BootnodesFlag:
				set.String(f.Name, "", f.Usage)
			case utils.GCModeFlag:
				set.String(f.Name, "archive", f.Usage)
			case utils.NodeKeyHexFlag: // either hex or file
			//	set.String(f.Name, "0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF", f.Usage) // either nodeKeyHex or nodeKeyFile
			case utils.NodeKeyFileFlag:
				set.String(f.Name, nodeKeyFile.Name(), f.Usage)
			case utils.NetrestrictFlag:
				set.String(f.Name, "127.0.0.0/16, 23.23.23.23/24,", f.Usage) // TOML problem
			case utils.AuthorizationListFlag:
				set.String(f.Name, "1=0x0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF,2=0x0123456789ABCDEF0123456789ABCDE00123456789ABCDEF0123456789ABCDEF", f.Usage)
			default:
				set.String(f.Name, f.Value+"_custom", f.Usage)
			}
		case cli.Uint64Flag:
			set.Uint64(f.Name, f.Value+10, f.Usage)
		case cli.IntFlag:
			set.Int(f.Name, f.Value+10, f.Usage)
		case utils.TextMarshalerFlag:
			set.String(f.Name, "light", f.Usage)
		case cli.Int64Flag:
			set.Int64(f.Name, f.Value+10, f.Usage)
		case cli.DurationFlag:
			set.Duration(f.Name, f.Value+5*time.Minute, f.Usage)
		case utils.BigFlag:
			set.Uint64(f.Name, f.Value.Uint64()+10, f.Usage)
		case cli.Float64Flag:
			set.Float64(f.Name, f.Value+0.1, f.Name)
		case cli.UintFlag:
			set.Uint(f.Name, f.Value+10, f.Usage)
		default:
			t.Log(fmt.Sprintf("unknown %t", f))
			t.Fail()
		}
	}
	action := utils.MigrateFlags(dumpConfig)
	app := &cli.App{
		Name:   "dumpconfig",
		Usage:  "dump config",
		Action: action,
	}

	ctx := cli.NewContext(app, set, nil)

	out, err := ioutil.TempFile("/tmp", "gethCfg")
	require.NoError(t, err)
	defer out.Close()
	defer os.Remove(out.Name())

	bak := os.Stdout
	defer func() { os.Stdout = bak }()
	os.Stdout = out

	err = action(ctx)
	require.NoError(t, err)

	out2, err := removeComment(out.Name())
	require.NoError(t, err)
	defer os.Remove(out2.Name())

	t.Log(out2.Name())
	val, err := ioutil.ReadFile(out2.Name())
	require.NoError(t, err)
	t.Log(string(val))

	cfg := &gethConfig{}
	err = loadConfig(out2.Name(), cfg)
	require.NoError(t, err)

	// [Eth]
	eth := cfg.Eth
	assert.Equal(t, uint64(1), eth.NetworkId) // mainnet true
	assert.Equal(t, downloader.FastSync, eth.SyncMode)
	assert.Equal(t, []string{"enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@all.mainnet.ethdisco.net"}, eth.EthDiscoveryURLs)
	assert.Equal(t, false, eth.NoPruning)
	assert.Equal(t, false, eth.NoPrefetch)
	assert.Equal(t, 100, eth.LightPeers)
	assert.Equal(t, 75, eth.UltraLightFraction)
	assert.Equal(t, 768, eth.DatabaseCache)
	assert.Equal(t, "", eth.DatabaseFreezer)
	assert.Equal(t, 256, eth.TrieCleanCache)
	assert.Equal(t, "triecache", eth.TrieCleanCacheJournal)
	assert.Equal(t, time.Duration(3600000000000), eth.TrieCleanCacheRejournal)
	assert.Equal(t, 256, eth.TrieDirtyCache)
	assert.Equal(t, time.Duration(3600000000000), eth.TrieTimeout)
	assert.Equal(t, 0, eth.SnapshotCache)
	assert.Equal(t, false, eth.EnablePreimageRecording)
	assert.Equal(t, "", eth.EWASMInterpreter)
	assert.Equal(t, "", eth.EVMInterpreter)
	assert.Equal(t, uint64(25000000), eth.RPCGasCap)
	assert.Equal(t, float64(1), eth.RPCTxFeeCap)
	// Quorum
	assert.Equal(t, time.Duration(15000000000), eth.EVMCallTimeOut)
	// End Quorum

	// [Eth.Miner]
	miner := cfg.Eth.Miner
	assert.Equal(t, uint64(700000000), miner.GasFloor)
	assert.Equal(t, uint64(800000000), miner.GasCeil)
	assert.Equal(t, big.NewInt(1000000000), miner.GasPrice)
	assert.Equal(t, time.Duration(3000000000), miner.Recommit)
	assert.Equal(t, false, miner.Noverify)
	assert.Equal(t, uint64(0), miner.AllowedFutureBlockTime)

	// [Eth.GPO]
	gpo := cfg.Eth.GPO
	assert.Equal(t, 2, gpo.Blocks)
	assert.Equal(t, 60, gpo.Percentile)

	// [Eth.TxPool]
	txPool := cfg.Eth.TxPool
	assert.Equal(t, []common.Address{}, txPool.Locals)
	assert.Equal(t, false, txPool.NoLocals)
	assert.Equal(t, "transactions.rlp", txPool.Journal)
	assert.Equal(t, time.Duration(3600000000000), txPool.Rejournal)
	assert.Equal(t, uint64(1), txPool.PriceLimit)
	assert.Equal(t, uint64(10), txPool.PriceBump)
	assert.Equal(t, uint64(16), txPool.AccountSlots)
	assert.Equal(t, uint64(4096), txPool.GlobalSlots)
	assert.Equal(t, uint64(64), txPool.AccountQueue)
	assert.Equal(t, uint64(1024), txPool.GlobalQueue)
	assert.Equal(t, time.Duration(10800000000000), txPool.Lifetime)
	assert.Equal(t, uint64(64), txPool.TransactionSizeLimit)
	assert.Equal(t, uint64(24), txPool.MaxCodeSize)

	// [Node]
	node := cfg.Node
	assert.Equal(t, "", node.DataDir)
	assert.Equal(t, false, node.InsecureUnlockAllowed)
	assert.Equal(t, false, node.NoUSB)
	assert.Equal(t, "", node.IPCPath)
	assert.Equal(t, "127.0.0.1", node.HTTPHost)
	assert.Equal(t, 8545, node.HTTPPort)
	assert.Equal(t, []string(nil), node.HTTPCors)
	assert.Equal(t, []string{"localhost"}, node.HTTPVirtualHosts)
	assert.Equal(t, []string{"net", "web3", "eth"}, node.HTTPModules)
	assert.Equal(t, "127.0.0.1", node.WSHost)
	assert.Equal(t, 8546, node.WSPort)
	assert.Equal(t, []string(nil), node.WSOrigins)
	assert.Equal(t, []string{"net", "web3", "eth"}, node.WSModules)
	assert.Equal(t, []string(nil), node.GraphQLCors)
	assert.Equal(t, []string{"localhost"}, node.GraphQLVirtualHosts)
	assert.Equal(t, false, node.EnableNodePermission)

	// [Node.P2P]
	p2p := cfg.Node.P2P
	assert.Equal(t, 0, p2p.MaxPeers)
	assert.Equal(t, true, p2p.NoDiscovery)

	assert.Equal(t, bootNodes(t).Nodes, p2p.BootstrapNodes)
	//assert.Equal(t, bootNodesV5(t).Nodes, p2p.BootstrapNodesV5)
	assert.Equal(t, ":0", p2p.ListenAddr)
	assert.Equal(t, false, p2p.EnableMsgEvents)

	type NetRestrictType struct {
		NetRestrict *netutil.Netlist
	}
	var netRestrict NetRestrictType
	err = toml.Unmarshal([]byte(`NetRestrict = ["127.0.0.0/16", "23.23.23.0/24"]`), &netRestrict)
	require.NoError(t, err)
	assert.Equal(t, netRestrict.NetRestrict, p2p.NetRestrict)

	// [Node.HTTPTimeouts]
	httpTimeouts := cfg.Node.HTTPTimeouts
	assert.Equal(t, time.Duration(30000000000), httpTimeouts.ReadTimeout)
	assert.Equal(t, time.Duration(30000000000), httpTimeouts.WriteTimeout)
	assert.Equal(t, time.Duration(120000000000), httpTimeouts.IdleTimeout)

	// QUORUM
	// [Eth.Istanbul]
	quorumIstanbul := eth.Istanbul
	assert.Equal(t, uint64(10000), quorumIstanbul.RequestTimeout)
	assert.Equal(t, uint64(1), quorumIstanbul.BlockPeriod)
	assert.Equal(t, uint64(30000), quorumIstanbul.Epoch)
	assert.Equal(t, big.NewInt(0), quorumIstanbul.Ceil2Nby3Block)
	assert.Equal(t, istanbul.RoundRobin, quorumIstanbul.ProposerPolicy.Id) // conflict with genesis?
	// END QUORUM
}

type BootNodesV5Type struct {
	Nodes []*enode.Node
}

func bootNodesV5(t *testing.T) BootNodesV5Type {
	var bootNodesV5 BootNodesV5Type
	err := toml.Unmarshal([]byte(`Nodes = ["enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303", "enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303", "enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", "enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303"]`), &bootNodesV5)
	require.NoError(t, err)
	return bootNodesV5
}

type BootNodesType struct {
	Nodes []*enode.Node
}

func bootNodes(t *testing.T) BootNodesType {
	var bootNodes BootNodesType
	err := toml.Unmarshal([]byte(`Nodes = ["enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303", "enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303", "enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", "enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303"]`), &bootNodes)
	require.NoError(t, err)
	return bootNodes
}

func removeComment(name string) (*os.File, error) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	out, err := ioutil.TempFile("/tmp", "gethCfg")
	if err != nil {
		return nil, fmt.Errorf("create temp file: %w", err)
	}
	defer out.Close()
	text := string(file)
	lines := strings.Split(text, "\n")
	first := true
	for _, line := range lines {
		if strings.Index(line, "#") != 0 && !(first && line == "") {
			line = strings.Replace(line, "e+00", ".0", 1)
			line = strings.Replace(line, "[[", "[", 1)
			line = strings.Replace(line, "]]", "]", 1)
			_, err = out.WriteString(line + "\n")
			if err != nil {
				return nil, fmt.Errorf("write line: %w", err)
			}
			first = false
		}
	}
	return out, nil
}

func TestLoadAndDumpGethConfig(t *testing.T) {
	out, err := ioutil.TempFile("/tmp", "gethCfg")
	require.NoError(t, err)
	defer out.Close()
	_, err = out.WriteString(`[Eth]
NetworkId = 1337
SyncMode = "full"
EthDiscoveryURLs = ["enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@all.mainnet.ethdisco.net"]
SnapDiscoveryURLs = []
NoPruning = false
NoPrefetch = false
LightPeers = 100
UltraLightFraction = 75
DatabaseCache = 768
DatabaseFreezer = ""
TrieCleanCache = 256
TrieCleanCacheJournal = "triecache-test"
TrieCleanCacheRejournal = 3600000000000
TrieDirtyCache = 256
TrieTimeout = 3600000000000
SnapshotCache = 0
Preimages = true
EnablePreimageRecording = false
EWASMInterpreter = ""
EVMInterpreter = ""
RPCGasCap = 25000000
RPCTxFeeCap = 1e+00
RaftMode = true
EnableNodePermission = true
EVMCallTimeOut = 3600000000000

[Eth.Miner]
GasFloor = 700000000
GasCeil = 800000000
GasPrice = 0
Recommit = 3000000000
Noverify = false
AllowedFutureBlockTime = 0

[Eth.GPO]
Blocks = 20
Percentile = 60
MaxPrice = 500000000000

[Eth.TxPool]
Locals = []
NoLocals = false
Journal = "transactions.rlp"
Rejournal = 3600000000000
PriceLimit = 1
PriceBump = 10
AccountSlots = 16
GlobalSlots = 4096
AccountQueue = 64
GlobalQueue = 1024
Lifetime = 10800000000000
TransactionSizeLimit = 64
MaxCodeSize = 24

[Eth.Istanbul]
RequestTimeout = 10000
BlockPeriod = 5
ProposerPolicy = "id = 0\n"
Epoch = 30000
Ceil2Nby3Block = 0
TestQBFTBlock = 0

[Node]
UserIdent = "_custom"
DataDir = "/data"
RaftLogDir = ""
InsecureUnlockAllowed = true
NoUSB = true
IPCPath = "geth.ipc"
HTTPHost = "0.0.0.0"
HTTPPort = 8545
HTTPCors = ["'*'"]
HTTPVirtualHosts = ["'*'"]
HTTPModules = ["admin", "db", "eth", "debug", "miner", "net", "txpool", "personal", "web3", "quorum", "istanbul"]
WSHost = "0.0.0.0"
WSPort = 8546
WSOrigins = ["'*'"]
WSModules = ["admin", "db", "eth", "debug", "miner", "net", "txpool", "personal", "web3", "quorum", "istanbul"]
GraphQLCors = ["'*'"]
GraphQLVirtualHosts = ["'*'"]
EnableNodePermission = true

[Node.P2P]
MaxPeers = 50
NoDiscovery = true
BootstrapNodes = ["enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303", "enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303", "enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", "enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303"]
BootstrapNodesV5 = ["enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303", "enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303", "enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", "enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303"]
StaticNodes = []
TrustedNodes = []
NetRestrict = ["127.0.0.0/16", "23.23.23.0/24"]
ListenAddr = ":30303"
EnableMsgEvents = false

[Node.HTTPTimeouts]
ReadTimeout = 30000000000
WriteTimeout = 30000000000
IdleTimeout = 120000000000
        
[Metrics]
HTTP = "127.0.0.1"
Port = 6060
InfluxDBEndpoint = "http://localhost:8086"
InfluxDBDatabase = "geth"
InfluxDBUsername = "test"
InfluxDBPassword = "test"
InfluxDBTags = "host=localhost"
`)
	require.NoError(t, err)
	err = out.Close()
	require.NoError(t, err)
	cfg := &gethConfig{}

	err = loadConfig(out.Name(), cfg)
	require.NoError(t, err)

	testConfig(t, cfg)

	out, err = ioutil.TempFile("/tmp", "gethCfg")
	require.NoError(t, err)

	err = tomlSettings.NewEncoder(out).Encode(cfg)
	require.NoError(t, err)

	cfg = &gethConfig{}
	err = loadConfig(out.Name(), cfg)
	require.NoError(t, err)

	testConfig(t, cfg)
}

func testConfig(t *testing.T, cfg *gethConfig) {
	// [Eth]
	eth := cfg.Eth
	assert.Equal(t, uint64(1337), eth.NetworkId)
	assert.Equal(t, downloader.FullSync, eth.SyncMode)
	assert.Equal(t, []string{"enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@all.mainnet.ethdisco.net"}, eth.EthDiscoveryURLs)
	assert.Equal(t, false, eth.NoPruning)
	assert.Equal(t, false, eth.NoPrefetch)
	assert.Equal(t, 100, eth.LightPeers)
	assert.Equal(t, 75, eth.UltraLightFraction)
	assert.Equal(t, 768, eth.DatabaseCache)
	assert.Equal(t, "", eth.DatabaseFreezer)
	assert.Equal(t, 256, eth.TrieCleanCache)
	assert.Equal(t, "triecache-test", eth.TrieCleanCacheJournal)
	assert.Equal(t, time.Duration(3600000000000), eth.TrieCleanCacheRejournal)
	assert.Equal(t, 256, eth.TrieDirtyCache)
	assert.Equal(t, time.Duration(3600000000000), eth.TrieTimeout)
	assert.Equal(t, 0, eth.SnapshotCache)
	assert.Equal(t, false, eth.EnablePreimageRecording)
	assert.Equal(t, "", eth.EWASMInterpreter)
	assert.Equal(t, "", eth.EVMInterpreter)
	assert.Equal(t, uint64(25000000), eth.RPCGasCap)
	assert.Equal(t, float64(1), eth.RPCTxFeeCap)
	// Quorum
	assert.Equal(t, time.Duration(3600000000000), eth.EVMCallTimeOut)
	assert.Equal(t, true, eth.EnableNodePermission)
	// End Quorum

	// [Eth.Miner]
	miner := eth.Miner
	assert.Equal(t, uint64(700000000), miner.GasFloor)
	assert.Equal(t, uint64(800000000), miner.GasCeil)
	assert.Equal(t, big.NewInt(0), miner.GasPrice)
	assert.Equal(t, time.Duration(3000000000), miner.Recommit)
	assert.Equal(t, false, miner.Noverify)
	assert.Equal(t, uint64(0), miner.AllowedFutureBlockTime)

	// [Eth.GPO]
	gpo := eth.GPO
	assert.Equal(t, 20, gpo.Blocks)
	assert.Equal(t, 60, gpo.Percentile)

	// [Eth.TxPool]
	txPool := eth.TxPool
	assert.Equal(t, []common.Address{}, txPool.Locals)
	assert.Equal(t, false, txPool.NoLocals)
	assert.Equal(t, "transactions.rlp", txPool.Journal)
	assert.Equal(t, time.Duration(3600000000000), txPool.Rejournal)
	assert.Equal(t, uint64(1), txPool.PriceLimit)
	assert.Equal(t, uint64(10), txPool.PriceBump)
	assert.Equal(t, uint64(16), txPool.AccountSlots)
	assert.Equal(t, uint64(4096), txPool.GlobalSlots)
	assert.Equal(t, uint64(64), txPool.AccountQueue)
	assert.Equal(t, uint64(1024), txPool.GlobalQueue)
	assert.Equal(t, time.Duration(10800000000000), txPool.Lifetime)
	assert.Equal(t, uint64(64), txPool.TransactionSizeLimit)
	assert.Equal(t, uint64(24), txPool.MaxCodeSize)

	// [Node]
	node := cfg.Node
	assert.Equal(t, "/data", node.DataDir)
	assert.Equal(t, true, node.InsecureUnlockAllowed)
	assert.Equal(t, true, node.NoUSB)
	assert.Equal(t, "geth.ipc", node.IPCPath)
	assert.Equal(t, "0.0.0.0", node.HTTPHost)
	assert.Equal(t, 8545, node.HTTPPort)
	assert.Equal(t, []string{"'*'"}, node.HTTPCors)
	assert.Equal(t, []string{"'*'"}, node.HTTPVirtualHosts)
	assert.Equal(t, []string{"admin", "db", "eth", "debug", "miner", "net", "txpool", "personal", "web3", "quorum", "istanbul"}, node.HTTPModules)
	assert.Equal(t, "0.0.0.0", node.WSHost)
	assert.Equal(t, 8546, node.WSPort)
	assert.Equal(t, []string{"'*'"}, node.WSOrigins)
	assert.Equal(t, []string{"admin", "db", "eth", "debug", "miner", "net", "txpool", "personal", "web3", "quorum", "istanbul"}, node.WSModules)
	assert.Equal(t, []string{"'*'"}, node.GraphQLCors)
	assert.Equal(t, []string{"'*'"}, node.GraphQLVirtualHosts)
	assert.Equal(t, true, node.EnableNodePermission)

	// [Node.P2P]
	p2p := cfg.Node.P2P
	assert.Equal(t, 50, p2p.MaxPeers)
	assert.Equal(t, true, p2p.NoDiscovery)
	assert.Equal(t, bootNodes(t).Nodes, p2p.BootstrapNodes)
	assert.Equal(t, bootNodesV5(t).Nodes, p2p.BootstrapNodesV5)

	/*assert.Equal(t, []*enode.Node{}, p2p.BootstrapNodes)
	if p2p.BootstrapNodesV5 != nil {
		assert.Equal(t, []*enode.Node{}, p2p.BootstrapNodesV5)
	}*/
	assert.Equal(t, ":30303", p2p.ListenAddr)
	assert.Equal(t, false, p2p.EnableMsgEvents)

	// [Node.HTTPTimeouts]
	httpTimeouts := cfg.Node.HTTPTimeouts
	assert.Equal(t, time.Duration(30000000000), httpTimeouts.ReadTimeout)
	assert.Equal(t, time.Duration(30000000000), httpTimeouts.WriteTimeout)
	assert.Equal(t, time.Duration(120000000000), httpTimeouts.IdleTimeout)

	// QUORUM
	// [Eth.Quorum.Istanbul]
	istanbul := cfg.Eth.Istanbul
	assert.Equal(t, uint64(10000), istanbul.RequestTimeout)
	assert.Equal(t, uint64(5), istanbul.BlockPeriod)
	assert.Equal(t, uint64(30000), istanbul.Epoch)
	assert.Equal(t, big.NewInt(0), istanbul.Ceil2Nby3Block)
	// END QUORUM
}
