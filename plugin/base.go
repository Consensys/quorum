package plugin

import (
	"context"
	"fmt"
	"io"
	slog "log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/common"
	iplugin "github.com/ethereum/go-ethereum/internal/plugin"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plugin/initializer"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type managedPlugin interface {
	Start() error
	Stop() error

	Info() (PluginInterfaceName, interface{})
}

// Plugin-meta.json
type MetaData struct {
	Version    string   `json:"version"`
	Os         string   `json:"os"`
	Arch       string   `json:"arch"`
	EntryPoint string   `json:"entrypoint"`
	Parameters []string `json:"parameters,omitempty"`
}

type basePlugin struct {
	pm               *PluginManager
	pluginInterface  PluginInterfaceName // plugin provider name
	pluginDefinition *PluginDefinition
	client           *plugin.Client
	gateways         plugin.PluginSet // gateways to invoke RPC API implementation of interfaces supported by this plugin
	pluginWorkspace  string           // plugin workspace
	commands         []string         // plugin executable commands
	logger           log.Logger
}

var basePluginPointerType = reflect.TypeOf(&basePlugin{})

func newBasePlugin(pm *PluginManager, pluginInterface PluginInterfaceName, pluginDefinition PluginDefinition, gateways plugin.PluginSet) (*basePlugin, error) {
	gateways[initializer.ConnectorName] = &initializer.PluginConnector{}

	// build basePlugin
	return &basePlugin{
		pm:               pm,
		pluginInterface:  pluginInterface,
		logger:           log.New("provider", pluginInterface, "plugin", pluginDefinition.Name, "version", pluginDefinition.Version),
		pluginDefinition: &pluginDefinition,
		gateways:         gateways,
	}, nil

}

// metadata.Command must be populated correctly here
func (bp *basePlugin) load() error {
	// Get plugin distribution path
	pluginDistFilePath, err := bp.pm.downloader.Download(bp.pluginDefinition)
	if err != nil {
		return err
	}
	// get file checksum
	pluginChecksum, err := bp.checksum(pluginDistFilePath)
	if err != nil {
		return err
	}
	bp.logger.Info("verifying plugin integrity", "checksum", pluginChecksum)
	if err := bp.pm.verifier.VerifySignature(bp.pluginDefinition, pluginChecksum); err != nil {
		return fmt.Errorf("unable to verify plugin signature: %v", err)
	}
	bp.logger.Info("unpacking plugin", "checksum", pluginChecksum)
	// Unpack plugin
	unPackDir, pluginMeta, err := unpackPlugin(pluginDistFilePath)
	if err != nil {
		return err
	}
	// Create Execution Command
	var command *exec.Cmd
	executable := path.Join(unPackDir, pluginMeta.EntryPoint)
	if !common.FileExist(executable) {
		return fmt.Errorf("entry point does not exist")
	}
	bp.logger.Debug("Plugin executable", "path", executable)
	if len(pluginMeta.Parameters) == 0 {
		command = exec.Command(executable)
		bp.commands = []string{executable}
	} else {
		command = exec.Command(executable, pluginMeta.Parameters...)
		bp.commands = append([]string{executable}, pluginMeta.Parameters...)
	}
	command.Dir = unPackDir
	bp.client = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  iplugin.DefaultHandshakeConfig,
		Plugins:          bp.gateways,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		Cmd:              command,
		AutoMTLS:         true,
		Logger:           &logDelegate{bp.logger.New("from", "plugin")},
	})

	bp.pluginWorkspace = unPackDir
	return nil
}

func (bp *basePlugin) Start() (err error) {
	startTime := time.Now()
	defer func(startTime time.Time) {
		if err == nil {
			bp.logger.Info("Plugin started", "took", time.Since(startTime))
		} else {
			bp.logger.Error("Plugin failed to start", "error", err, "took", time.Since(startTime))
			_ = bp.Stop()
		}
	}(startTime)
	bp.logger.Info("Starting plugin")
	bp.logger.Debug("Starting plugin: Loading")
	err = bp.load()
	if err != nil {
		return
	}
	bp.logger.Debug("Starting plugin: Creating client")
	_, err = bp.client.Client()
	if err != nil {
		return
	}
	bp.logger.Debug("Starting plugin: Initializing")
	err = bp.init()
	return
}

func (bp *basePlugin) Stop() error {
	if bp.client != nil {
		bp.client.Kill()
	}
	if bp.pluginWorkspace == "" {
		return nil
	}
	return bp.cleanPluginWorkspace()
}

func (bp *basePlugin) cleanPluginWorkspace() error {
	workspace, err := os.Open(bp.pluginWorkspace)
	if err != nil {
		return err
	}
	defer workspace.Close()
	names, err := workspace.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(bp.pluginWorkspace, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func (bp *basePlugin) init() error {
	bp.logger.Info("Initializing plugin")
	raw, err := bp.dispense(initializer.ConnectorName)
	if err != nil {
		return err
	}
	c, ok := raw.(initializer.PluginInitializer)
	if !ok {
		return fmt.Errorf("missing plugin initializer. Make sure it is in the plugin set")
	}
	rawConfig, err := ReadMultiFormatConfig(bp.pluginDefinition.Config)
	if err != nil {
		return err
	}
	return c.Init(context.Background(), bp.pm.nodeName, rawConfig)
}

func (bp *basePlugin) dispense(name string) (interface{}, error) {
	rpcClient, err := bp.client.Client()
	if err != nil {
		return nil, err
	}
	return rpcClient.Dispense(name)
}

func (bp *basePlugin) Config() *PluginDefinition {
	return bp.pluginDefinition
}

func (bp *basePlugin) checksum(pluginFile string) (string, error) {
	return getSha256Checksum(pluginFile)
}

func (bp *basePlugin) Info() (PluginInterfaceName, interface{}) {
	info := make(map[string]interface{})
	info["name"] = bp.pluginDefinition.Name
	info["version"] = bp.pluginDefinition.Version
	info["config"] = bp.pluginDefinition.Config
	info["executable"] = bp.commands
	return bp.pluginInterface, info
}

type logDelegate struct {
	eLogger log.Logger
}

func (ld *logDelegate) Trace(msg string, args ...interface{}) {
	ld.eLogger.Trace(msg, args...)
}

func (ld *logDelegate) Debug(msg string, args ...interface{}) {
	ld.eLogger.Debug(msg, args...)
}

func (ld *logDelegate) Info(msg string, args ...interface{}) {
	ld.eLogger.Info(msg, args...)
}

func (ld *logDelegate) Warn(msg string, args ...interface{}) {
	ld.eLogger.Warn(msg, args...)
}

func (ld *logDelegate) Error(msg string, args ...interface{}) {
	ld.eLogger.Error(msg, args...)
}

func (ld *logDelegate) IsTrace() bool {
	return true
}

func (*logDelegate) IsDebug() bool {
	return true
}

func (*logDelegate) IsInfo() bool {
	return true
}

func (*logDelegate) IsWarn() bool {
	return true
}

func (*logDelegate) IsError() bool {
	return true
}

func (ld *logDelegate) With(args ...interface{}) hclog.Logger {
	return &logDelegate{ld.eLogger.New(args...)}
}

func (ld *logDelegate) Named(name string) hclog.Logger {
	return ld
}

func (ld *logDelegate) ResetNamed(name string) hclog.Logger {
	return ld
}

func (ld *logDelegate) SetLevel(level hclog.Level) {
}

func (*logDelegate) StandardLogger(opts *hclog.StandardLoggerOptions) *slog.Logger {
	return nil
}

func (*logDelegate) StandardWriter(opts *hclog.StandardLoggerOptions) io.Writer {
	return nil
}
