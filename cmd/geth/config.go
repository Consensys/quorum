// Copyright 2017 The go-ethereum Authors
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

package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/big"
	"os"
	"reflect"
	"unicode"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common/http"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/extension/privacyExtension"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/naoina/toml"
	"gopkg.in/urfave/cli.v1"
)

var (
	dumpConfigCommand = cli.Command{
		Action:      utils.MigrateFlags(dumpConfig),
		Name:        "dumpconfig",
		Usage:       "Show configuration values",
		ArgsUsage:   "",
		Flags:       append(nodeFlags, rpcFlags...),
		Category:    "MISCELLANEOUS COMMANDS",
		Description: `The dumpconfig command shows configuration values.`,
	}

	configFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}
)

// These settings ensure that TOML keys use the same names as Go struct fields.
var tomlSettings = toml.Config{
	NormFieldName: func(rt reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(rt reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		link := ""
		if unicode.IsUpper(rune(rt.Name()[0])) && rt.PkgPath() != "main" {
			link = fmt.Sprintf(", see https://godoc.org/%s#%s for available fields", rt.PkgPath(), rt.Name())
		}
		return fmt.Errorf("field '%s' is not defined in %s%s", field, rt.String(), link)
	},
}

type ethstatsConfig struct {
	URL string `toml:",omitempty"`
}

type gethConfig struct {
	Eth      ethconfig.Config
	Node     node.Config
	Ethstats ethstatsConfig
	Metrics  metrics.Config
}

func loadConfig(file string, cfg *gethConfig) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tomlSettings.NewDecoder(bufio.NewReader(f)).Decode(cfg)
	// Add file name to errors that have a line number.
	if _, ok := err.(*toml.LineError); ok {
		err = errors.New(file + ", " + err.Error())
	}
	return err
}

func defaultNodeConfig() node.Config {
	cfg := node.DefaultConfig
	cfg.Name = clientIdentifier
	cfg.Version = params.VersionWithCommit(gitCommit, gitDate)
	cfg.HTTPModules = append(cfg.HTTPModules, "eth")
	cfg.WSModules = append(cfg.WSModules, "eth")
	cfg.IPCPath = "geth.ipc"
	return cfg
}

// makeConfigNode loads geth configuration and creates a blank node instance.
func makeConfigNode(ctx *cli.Context) (*node.Node, gethConfig) {
	// Quorum: Must occur before setQuorumConfig, as it needs an initialised PTM to be enabled
	// 		   Extension Service and Multitenancy feature validation also depend on PTM availability
	if err := quorumInitialisePrivacy(ctx); err != nil {
		utils.Fatalf("Error initialising Private Transaction Manager: %s", err.Error())
	}

	// Load defaults.
	cfg := gethConfig{
		Eth:     ethconfig.Defaults,
		Node:    defaultNodeConfig(),
		Metrics: metrics.DefaultConfig,
	}

	// Load config file.
	if file := ctx.GlobalString(configFileFlag.Name); file != "" {
		if err := loadConfig(file, &cfg); err != nil {
			utils.Fatalf("%v", err)
		}
	}

	// Apply flags.
	utils.SetNodeConfig(ctx, &cfg.Node)
	stack, err := node.New(&cfg.Node)
	if err != nil {
		utils.Fatalf("Failed to create the protocol stack: %v", err)
	}
	utils.SetEthConfig(ctx, stack, &cfg.Eth)
	if ctx.GlobalIsSet(utils.EthStatsURLFlag.Name) {
		cfg.Ethstats.URL = ctx.GlobalString(utils.EthStatsURLFlag.Name)
	}
	applyMetricConfig(ctx, &cfg)

	return stack, cfg
}

// makeFullNode loads geth configuration and creates the Ethereum backend.
func makeFullNode(ctx *cli.Context) (*node.Node, ethapi.Backend) {
	stack, cfg := makeConfigNode(ctx)
	if ctx.GlobalIsSet(utils.OverrideBerlinFlag.Name) {
		cfg.Eth.OverrideBerlin = new(big.Int).SetUint64(ctx.GlobalUint64(utils.OverrideBerlinFlag.Name))
	}

	//Must occur before registering the extension service, as it needs an initialised PTM to be enabled
	if err := quorumInitialisePrivacy(ctx); err != nil {
		utils.Fatalf("Error initialising Private Transaction Manager: %s", err.Error())
	}

	// Quorum - returning `ethService` too for the Raft and extension service
	backend, ethService := utils.RegisterEthService(stack, &cfg.Eth)

	// Quorum
	// plugin service must be after eth service so that eth service will be stopped gradually if any of the plugin
	// fails to start
	if cfg.Node.Plugins != nil {
		utils.RegisterPluginService(stack, &cfg.Node, ctx.Bool(utils.PluginSkipVerifyFlag.Name), ctx.Bool(utils.PluginLocalVerifyFlag.Name), ctx.String(utils.PluginPublicKeyFlag.Name))
	}

	if cfg.Node.IsPermissionEnabled() {
		utils.RegisterPermissionService(stack, ctx.Bool(utils.RaftDNSEnabledFlag.Name), backend.ChainConfig().ChainID)
	}

	if ctx.GlobalBool(utils.RaftModeFlag.Name) {
		utils.RegisterRaftService(stack, ctx, &cfg.Node, ethService)
	}

	if private.IsQuorumPrivacyEnabled() {
		utils.RegisterExtensionService(stack, ethService)
	}
	// End Quorum

	// Configure GraphQL if requested
	if ctx.GlobalIsSet(utils.GraphQLEnabledFlag.Name) {
		utils.RegisterGraphQLService(stack, backend, cfg.Node)
	}
	// Add the Ethereum Stats daemon if requested.
	if cfg.Ethstats.URL != "" {
		utils.RegisterEthStatsService(stack, backend, cfg.Ethstats.URL)
	}
	return stack, backend
}

// dumpConfig is the dumpconfig command.
func dumpConfig(ctx *cli.Context) error {
	_, cfg := makeConfigNode(ctx)
	comment := ""

	if cfg.Eth.Genesis != nil {
		cfg.Eth.Genesis = nil
		comment += "# Note: this config doesn't contain the genesis block.\n\n"
	}

	out, err := tomlSettings.Marshal(&cfg)
	if err != nil {
		return err
	}

	dump := os.Stdout
	if ctx.NArg() > 0 {
		dump, err = os.OpenFile(ctx.Args().Get(0), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer dump.Close()
	}
	dump.WriteString(comment)
	dump.Write(out)

	return nil
}

func applyMetricConfig(ctx *cli.Context, cfg *gethConfig) {
	if ctx.GlobalIsSet(utils.MetricsEnabledFlag.Name) {
		cfg.Metrics.Enabled = ctx.GlobalBool(utils.MetricsEnabledFlag.Name)
	}
	if ctx.GlobalIsSet(utils.MetricsEnabledExpensiveFlag.Name) {
		cfg.Metrics.EnabledExpensive = ctx.GlobalBool(utils.MetricsEnabledExpensiveFlag.Name)
	}
	if ctx.GlobalIsSet(utils.MetricsHTTPFlag.Name) {
		cfg.Metrics.HTTP = ctx.GlobalString(utils.MetricsHTTPFlag.Name)
	}
	if ctx.GlobalIsSet(utils.MetricsPortFlag.Name) {
		cfg.Metrics.Port = ctx.GlobalInt(utils.MetricsPortFlag.Name)
	}
	if ctx.GlobalIsSet(utils.MetricsEnableInfluxDBFlag.Name) {
		cfg.Metrics.EnableInfluxDB = ctx.GlobalBool(utils.MetricsEnableInfluxDBFlag.Name)
	}
	if ctx.GlobalIsSet(utils.MetricsInfluxDBEndpointFlag.Name) {
		cfg.Metrics.InfluxDBEndpoint = ctx.GlobalString(utils.MetricsInfluxDBEndpointFlag.Name)
	}
	if ctx.GlobalIsSet(utils.MetricsInfluxDBDatabaseFlag.Name) {
		cfg.Metrics.InfluxDBDatabase = ctx.GlobalString(utils.MetricsInfluxDBDatabaseFlag.Name)
	}
	if ctx.GlobalIsSet(utils.MetricsInfluxDBUsernameFlag.Name) {
		cfg.Metrics.InfluxDBUsername = ctx.GlobalString(utils.MetricsInfluxDBUsernameFlag.Name)
	}
	if ctx.GlobalIsSet(utils.MetricsInfluxDBPasswordFlag.Name) {
		cfg.Metrics.InfluxDBPassword = ctx.GlobalString(utils.MetricsInfluxDBPasswordFlag.Name)
	}
	if ctx.GlobalIsSet(utils.MetricsInfluxDBTagsFlag.Name) {
		cfg.Metrics.InfluxDBTags = ctx.GlobalString(utils.MetricsInfluxDBTagsFlag.Name)
	}
}

// quorumValidateEthService checks quorum features that depend on the ethereum service
func quorumValidateEthService(stack *node.Node, isRaft bool) {
	var ethereum *eth.Ethereum

	err := stack.Lifecycle(&ethereum)
	if err != nil {
		utils.Fatalf("Error retrieving Ethereum service: %v", err)
	}

	quorumValidateConsensus(ethereum, isRaft)

	quorumValidatePrivacyEnhancements(ethereum)
}

// quorumValidateConsensus checks if a consensus was used. The node is killed if consensus was not used
func quorumValidateConsensus(ethereum *eth.Ethereum, isRaft bool) {
	if !isRaft && ethereum.BlockChain().Config().Istanbul == nil && ethereum.BlockChain().Config().Clique == nil {
		utils.Fatalf("Consensus not specified. Exiting!!")
	}
}

// quorumValidatePrivacyEnhancements checks if privacy enhancements are configured the transaction manager supports
// the PrivacyEnhancements feature
func quorumValidatePrivacyEnhancements(ethereum *eth.Ethereum) {
	privacyEnhancementsBlock := ethereum.BlockChain().Config().PrivacyEnhancementsBlock
	if privacyEnhancementsBlock != nil {
		log.Info("Privacy enhancements is configured to be enabled from block ", "height", privacyEnhancementsBlock)
		if !private.P.HasFeature(engine.PrivacyEnhancements) {
			utils.Fatalf("Cannot start quorum with privacy enhancements enabled while the transaction manager does not support it")
		}
	}
}

// configure and set up quorum transaction privacy
func quorumInitialisePrivacy(ctx *cli.Context) error {
	cfg, err := QuorumSetupPrivacyConfiguration(ctx)
	if err != nil {
		return err
	}

	err = private.InitialiseConnection(cfg)
	if err != nil {
		return err
	}
	privacyExtension.Init()

	return nil
}

// Get private transaction manager configuration
func QuorumSetupPrivacyConfiguration(ctx *cli.Context) (http.Config, error) {
	// get default configuration
	cfg, err := private.GetLegacyEnvironmentConfig()
	if err != nil {
		return http.Config{}, err
	}

	// override the config with command line parameters
	if ctx.GlobalIsSet(utils.QuorumPTMUnixSocketFlag.Name) {
		cfg.SetSocket(ctx.GlobalString(utils.QuorumPTMUnixSocketFlag.Name))
	}
	if ctx.GlobalIsSet(utils.QuorumPTMUrlFlag.Name) {
		cfg.SetHttpUrl(ctx.GlobalString(utils.QuorumPTMUrlFlag.Name))
	}
	if ctx.GlobalIsSet(utils.QuorumPTMTimeoutFlag.Name) {
		cfg.SetTimeout(ctx.GlobalUint(utils.QuorumPTMTimeoutFlag.Name))
	}
	if ctx.GlobalIsSet(utils.QuorumPTMDialTimeoutFlag.Name) {
		cfg.SetDialTimeout(ctx.GlobalUint(utils.QuorumPTMDialTimeoutFlag.Name))
	}
	if ctx.GlobalIsSet(utils.QuorumPTMHttpIdleTimeoutFlag.Name) {
		cfg.SetHttpIdleConnTimeout(ctx.GlobalUint(utils.QuorumPTMHttpIdleTimeoutFlag.Name))
	}
	if ctx.GlobalIsSet(utils.QuorumPTMHttpWriteBufferSizeFlag.Name) {
		cfg.SetHttpWriteBufferSize(ctx.GlobalInt(utils.QuorumPTMHttpWriteBufferSizeFlag.Name))
	}
	if ctx.GlobalIsSet(utils.QuorumPTMHttpReadBufferSizeFlag.Name) {
		cfg.SetHttpReadBufferSize(ctx.GlobalInt(utils.QuorumPTMHttpReadBufferSizeFlag.Name))
	}
	if ctx.GlobalIsSet(utils.QuorumPTMTlsModeFlag.Name) {
		cfg.SetTlsMode(ctx.GlobalString(utils.QuorumPTMTlsModeFlag.Name))
	}
	if ctx.GlobalIsSet(utils.QuorumPTMTlsRootCaFlag.Name) {
		cfg.SetTlsRootCA(ctx.GlobalString(utils.QuorumPTMTlsRootCaFlag.Name))
	}
	if ctx.GlobalIsSet(utils.QuorumPTMTlsClientCertFlag.Name) {
		cfg.SetTlsClientCert(ctx.GlobalString(utils.QuorumPTMTlsClientCertFlag.Name))
	}
	if ctx.GlobalIsSet(utils.QuorumPTMTlsClientKeyFlag.Name) {
		cfg.SetTlsClientKey(ctx.GlobalString(utils.QuorumPTMTlsClientKeyFlag.Name))
	}
	if ctx.GlobalIsSet(utils.QuorumPTMTlsInsecureSkipVerify.Name) {
		cfg.SetTlsInsecureSkipVerify(ctx.Bool(utils.QuorumPTMTlsInsecureSkipVerify.Name))
	}

	if err = cfg.Validate(); err != nil {
		return cfg, err
	}
	return cfg, nil
}
