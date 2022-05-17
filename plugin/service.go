package plugin

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/pluggable"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

type PluginManagerInterface interface {
	APIs() []rpc.API
	Start() (err error)
	Stop() error
	IsEnabled(name PluginInterfaceName) bool
	PluginsInfo() interface{}
	Reload(name PluginInterfaceName) (bool, error)
	GetPluginTemplate(name PluginInterfaceName, v managedPlugin) error
}

//go:generate mockgen -source=service.go -destination plugin_manager_mockery.go -package plugin
var _ PluginManagerInterface = &PluginManager{}
var _ PluginManagerInterface = &MockPluginManagerInterface{}

// this implements geth service
type PluginManager struct {
	nodeName           string // geth node name
	pluginBaseDir      string // base directory for all the plugins
	verifier           Verifier
	centralClient      *CentralClient
	downloader         *Downloader
	settings           *Settings
	mux                sync.Mutex                            // control concurrent access to plugins cache
	plugins            map[PluginInterfaceName]managedPlugin // lazy load the actual plugin templates
	initializedPlugins map[PluginInterfaceName]managedPlugin // prepopulate during initialization of plugin manager, needed for starting/stopping/getting info
	pluginsStarted     *int32
}

// this is called after PluginManager service has been successfully started
// See node/node.go#Start()
func (s *PluginManager) APIs() []rpc.API {
	return append([]rpc.API{
		{
			Namespace: "admin",
			Service:   NewPluginManagerAPI(s),
			Version:   "1.0",
			Public:    false,
		},
	}, s.delegateAPIs()...)
}

func (s *PluginManager) Start() (err error) {
	initializedPluginsCount := len(s.initializedPlugins)
	if initializedPluginsCount == 0 {
		log.Info("No plugins to initialise")
		return
	}
	if atomic.LoadInt32(s.pluginsStarted) != 0 {
		log.Info("Plugins already started")
		return
	}
	log.Info("Starting all plugins", "count", initializedPluginsCount)
	startedPlugins := make([]managedPlugin, 0, initializedPluginsCount)
	for _, p := range s.initializedPlugins {
		if err = p.Start(); err != nil {
			break
		} else {
			startedPlugins = append(startedPlugins, p)
		}
	}
	if err != nil {
		for _, p := range startedPlugins {
			_ = p.Stop()
		}
	} else {
		atomic.StoreInt32(s.pluginsStarted, 1)
	}
	return
}

func (s *PluginManager) getPlugin(name PluginInterfaceName) (managedPlugin, bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	p, ok := s.plugins[name] // check if it's been used before
	if !ok {
		p, ok = s.initializedPlugins[name] // check if it's been initialized before
	}
	return p, ok
}

// Check if a plugin is enabled/setup
func (s *PluginManager) IsEnabled(name PluginInterfaceName) bool {
	if s == nil {
		return false
	}
	_, ok := s.initializedPlugins[name]
	return ok
}

// store the plugin instance to the value of the pointer v and cache it
// this function makes sure v value will never be nil
func (s *PluginManager) GetPluginTemplate(name PluginInterfaceName, v managedPlugin) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("invalid argument value, expected a pointer but got %s", reflect.TypeOf(v))
	}
	recoverToErrorFunc := func(f func()) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%s", r)
			}
		}()
		f()
		return
	}
	if p, ok := s.plugins[name]; ok {
		return recoverToErrorFunc(func() {
			cachedValue := reflect.ValueOf(p)
			rv.Elem().Set(cachedValue.Elem())
		})
	}
	base, ok := s.initializedPlugins[name]
	if !ok {
		return fmt.Errorf("plugin: [%s] is not found", name)
	}
	if err := recoverToErrorFunc(func() {
		basePluginValue := reflect.ValueOf(base)
		// the first field in the plugin template object is the basePlugin
		// it indicates that the plugin template "extends" basePlugin
		basePluginField := rv.Elem().FieldByName("basePlugin")
		if !basePluginField.IsValid() || basePluginField.Type() != basePluginPointerType {
			panic("plugin template must extend *basePlugin")
		}
		// need to have write access to the unexported field in the target object
		basePluginField = reflect.NewAt(basePluginField.Type(), unsafe.Pointer(basePluginField.UnsafeAddr())).Elem()
		basePluginField.Set(basePluginValue)
	}); err != nil {
		return err
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	s.plugins[name] = v
	return nil
}

func (s *PluginManager) Stop() error {
	initializedPluginsCount := len(s.initializedPlugins)
	log.Info("Stopping all plugins", "count", initializedPluginsCount)
	allErrors := make([]error, 0)
	for _, p := range s.initializedPlugins {
		if err := p.Stop(); err != nil {
			allErrors = append(allErrors, err)
		}
	}
	log.Info("All plugins stopped", "errors", allErrors)
	if initializedPluginsCount > 0 {
		atomic.StoreInt32(s.pluginsStarted, 0)
	}
	if len(allErrors) == 0 {
		return nil
	}
	return fmt.Errorf("%s", allErrors)
}

// Provide details of current plugins being used
func (s *PluginManager) PluginsInfo() interface{} {
	info := make(map[PluginInterfaceName]interface{})
	if len(s.initializedPlugins) == 0 {
		return info
	}
	info["baseDir"] = s.pluginBaseDir
	for _, p := range s.initializedPlugins {
		k, v := p.Info()
		info[k] = v
	}
	return info
}

// AddAccountPluginToBackend adds the account plugin to the provided account backend
func (s *PluginManager) AddAccountPluginToBackend(b *pluggable.Backend) error {
	v := new(ReloadableAccountServiceFactory)
	if err := s.GetPluginTemplate(AccountPluginInterfaceName, v); err != nil {
		return err
	}
	service, err := v.Create()
	if err != nil {
		return err
	}
	if err := b.SetPluginService(service); err != nil {
		return err
	}
	return nil
}

func (s *PluginManager) Reload(name PluginInterfaceName) (bool, error) {
	p, ok := s.getPlugin(name)
	if !ok {
		return false, fmt.Errorf("no such plugin provider: %s", name)
	}
	_ = p.Stop()
	if err := p.Start(); err != nil {
		return false, err
	}
	return true, nil
}

// this is to configure delegate APIs call to the plugins
func (s *PluginManager) delegateAPIs() []rpc.API {
	apis := make([]rpc.API, 0)
	for _, p := range s.initializedPlugins {
		interfaceName, _ := p.Info()
		if pluginProvider, ok := pluginProviders[interfaceName]; ok {
			if pluginProvider.apiProviderFunc != nil {
				namespace := fmt.Sprintf("plugin@%s", interfaceName)
				log.Debug("adding RPC API delegate for plugin", "provider", interfaceName, "namespace", namespace)
				if delegates, err := pluginProvider.apiProviderFunc(namespace, s); err != nil {
					log.Error("unable to delegate RPC API calls to plugin", "provider", interfaceName, "error", err)
				} else {
					apis = append(apis, delegates...)
				}
			}
		}
	}
	return apis
}

func NewPluginManager(nodeName string, settings *Settings, skipVerify bool, localVerify bool, publicKey string) (*PluginManager, error) {
	pm := &PluginManager{
		nodeName:           nodeName,
		pluginBaseDir:      settings.BaseDir.String(),
		centralClient:      NewPluginCentralClient(settings.CentralConfig),
		plugins:            make(map[PluginInterfaceName]managedPlugin),
		initializedPlugins: make(map[PluginInterfaceName]managedPlugin),
		settings:           settings,
		pluginsStarted:     new(int32),
	}
	pm.downloader = NewDownloader(pm)
	if skipVerify {
		log.Warn("plugin: ignore integrity verification")
		pm.verifier = NewNonVerifier()
	} else {
		var err error
		if pm.verifier, err = NewVerifier(pm, localVerify, publicKey); err != nil {
			return nil, err
		}
	}
	for pluginName, pluginDefinition := range settings.Providers {
		log.Debug("Preparing plugin", "provider", pluginName, "name", pluginDefinition.Name, "version", pluginDefinition.Version)
		pluginProvider, ok := pluginProviders[pluginName]
		if !ok {
			return nil, fmt.Errorf("plugin: [%s] is not supported", pluginName)
		}
		base, err := newBasePlugin(pm, pluginName, pluginDefinition, pluginProvider.pluginSet)
		if err != nil {
			return nil, fmt.Errorf("plugin [%s] %s", pluginName, err.Error())
		}
		pm.initializedPlugins[pluginName] = base
	}
	log.Debug("Created plugin manager", "PluginsInfo()", pm.PluginsInfo())
	return pm, nil
}

func NewEmptyPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make(map[PluginInterfaceName]managedPlugin),
	}
}
