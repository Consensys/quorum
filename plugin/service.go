package plugin

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

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
}

func (s *PluginManager) Protocols() []p2p.Protocol { return nil }

func (s *PluginManager) APIs() []rpc.API {
	// the below code show how to expose APIs of a plugin via JSON RPC
	// this is only for demonstration purposes
	helloWorldAPI := make([]rpc.API, 0)
	helloWorldPluginTemplate := new(HelloWorldPluginTemplate)
	if err := s.GetPluginTemplate(HelloWorldPluginInterfaceName, helloWorldPluginTemplate); err != nil {
		log.Info("plugin: not configured", "name", HelloWorldPluginInterfaceName, "err", err)
	} else {
		pluginInstance, err := helloWorldPluginTemplate.Get()
		if err != nil {
			log.Info("plugin: instance not ready", "name", HelloWorldPluginInterfaceName, "err", err)
		} else {
			helloWorldAPI = append(helloWorldAPI, rpc.API{
				Namespace: fmt.Sprintf("plugin@%s", HelloWorldPluginInterfaceName),
				Service:   pluginInstance,
				Version:   "1.0",
				Public:    true,
			})
		}
	}
	return append([]rpc.API{
		{
			Namespace: "admin",
			Service:   NewPluginManagerAPI(s),
			Version:   "1.0",
			Public:    false,
		},
	}, helloWorldAPI...)
}

func (s *PluginManager) Start(_ *p2p.Server) (err error) {
	log.Info("Starting all plugins", "count", len(s.initializedPlugins))
	startedPlugins := make([]managedPlugin, 0, len(s.initializedPlugins))
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
	}
	return
}

func (s *PluginManager) getPlugin(name PluginInterfaceName) (managedPlugin, bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	p, ok := s.plugins[name]
	return p, ok
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
			panic(fmt.Sprintf("plugin template must extend *basePlugin"))
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
	log.Info("Stopping all plugins", "count", len(s.initializedPlugins))
	allErrors := make([]error, 0)
	for _, p := range s.initializedPlugins {
		if err := p.Stop(); err != nil {
			allErrors = append(allErrors, err)
		}
	}
	log.Info("All plugins stopped", "errors", allErrors)
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

func NewPluginManager(nodeName string, settings *Settings, skipVerify bool, localVerify bool, publicKey string) (*PluginManager, error) {
	pm := &PluginManager{
		nodeName:           nodeName,
		pluginBaseDir:      settings.BaseDir.String(),
		centralClient:      NewPluginCentralClient(settings.CentralConfig),
		plugins:            make(map[PluginInterfaceName]managedPlugin),
		initializedPlugins: make(map[PluginInterfaceName]managedPlugin),
		settings:           settings,
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
		base, err := newBasePlugin(pm, pluginName, pluginDefinition, pluginProvider)
		if err != nil {
			return nil, fmt.Errorf("plugin [%s] %s", pluginName, err.Error())
		}
		pm.initializedPlugins[pluginName] = base
	}
	return pm, nil
}

func NewEmptyPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make(map[PluginInterfaceName]managedPlugin),
	}
}
