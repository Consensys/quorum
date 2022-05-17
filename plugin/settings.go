package plugin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/plugin/account"
	"github.com/ethereum/go-ethereum/plugin/helloworld"
	"github.com/ethereum/go-ethereum/plugin/qlight"
	"github.com/ethereum/go-ethereum/plugin/security"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/go-plugin"
	"github.com/naoina/toml"
)

const (
	HelloWorldPluginInterfaceName         = PluginInterfaceName("helloworld") // lower-case always
	SecurityPluginInterfaceName           = PluginInterfaceName("security")
	AccountPluginInterfaceName            = PluginInterfaceName("account")
	QLightTokenManagerPluginInterfaceName = PluginInterfaceName("qlighttokenmanager")
)

var (
	// define additional plugins being supported here
	pluginProviders = map[PluginInterfaceName]pluginProvider{
		HelloWorldPluginInterfaceName: {
			apiProviderFunc: func(ns string, pm *PluginManager) ([]rpc.API, error) {
				template := new(HelloWorldPluginTemplate)
				if err := pm.GetPluginTemplate(HelloWorldPluginInterfaceName, template); err != nil {
					return nil, err
				}
				service, err := template.Get()
				if err != nil {
					return nil, err
				}
				return []rpc.API{{
					Namespace: ns,
					Version:   "1.0.0",
					Service:   service,
					Public:    true,
				}}, nil
			},
			pluginSet: plugin.PluginSet{
				helloworld.ConnectorName: &helloworld.PluginConnector{},
			},
		},
		SecurityPluginInterfaceName: {
			pluginSet: plugin.PluginSet{
				security.TLSConfigurationConnectorName: &security.TLSConfigurationSourcePluginConnector{},
				security.AuthenticationConnectorName:   &security.AuthenticationManagerPluginConnector{},
			},
		},
		AccountPluginInterfaceName: {
			apiProviderFunc: func(ns string, pm *PluginManager) ([]rpc.API, error) {
				f := new(ReloadableAccountServiceFactory)
				if err := pm.GetPluginTemplate(AccountPluginInterfaceName, f); err != nil {
					return nil, err
				}
				service, err := f.Create()
				if err != nil {
					return nil, err
				}
				return []rpc.API{{
					Namespace: ns,
					Version:   "1.0.0",
					Service:   account.NewCreator(service),
					Public:    true,
				}}, nil
			},
			pluginSet: plugin.PluginSet{
				account.ConnectorName: &account.PluginConnector{},
			},
		},
		QLightTokenManagerPluginInterfaceName: {
			pluginSet: plugin.PluginSet{
				qlight.ConnectorName: &qlight.PluginConnector{},
			},
		},
	}

	// this is the place holder for future solution of the plugin central
	quorumPluginCentralConfiguration = &PluginCentralConfiguration{
		CertFingerprint:        "",
		BaseURL:                "https://artifacts.consensys.net/public/quorum-go-plugins/",
		PublicKeyURI:           DefaultPublicKeyFile,
		InsecureSkipTLSVerify:  false,
		PluginDistPathTemplate: "maven/bin/{{.Name}}/{{.Version}}/{{.Name}}-{{.Version}}-{{.OS}}-{{.Arch}}.zip",
		PluginSigPathTemplate:  "maven/bin/{{.Name}}/{{.Version}}/{{.Name}}-{{.Version}}-{{.OS}}-{{.Arch}}-sha256.checksum.asc",
	}
)

type pluginProvider struct {
	// this allows exposing plugin interfaces to geth RPC API automatically.
	// nil value implies that plugin won't expose its methods to geth RPC API
	apiProviderFunc rpcAPIProviderFunc
	// contains connectors being registered to the plugin library
	pluginSet plugin.PluginSet
}

type rpcAPIProviderFunc func(ns string, pm *PluginManager) ([]rpc.API, error)
type Version string

// This is to describe a plugin
//
// Information is used to discover the plugin binary and verify its integrity
// before forking a process running the plugin
type PluginDefinition struct {
	Name string `json:"name" toml:""`
	// the semver version of the plugin
	Version Version `json:"version" toml:""`
	// plugin configuration in a form of map/slice/string
	Config interface{} `json:"config,omitempty" toml:",omitempty"`
}

func ReadMultiFormatConfig(config interface{}) ([]byte, error) {
	if config == nil {
		return []byte{}, nil
	}
	switch k := reflect.TypeOf(config).Kind(); k {
	case reflect.Map, reflect.Slice:
		return json.Marshal(config)
	case reflect.String:
		configStr := config.(string)
		u, err := url.Parse(configStr)
		if err != nil { // just return as is
			return []byte(configStr), nil
		}
		switch s := u.Scheme; s {
		case "file":
			return ioutil.ReadFile(filepath.Join(u.Host, u.Path))
		case "env": // config string in an env variable
			varName := u.Host
			isFile := u.Query().Get("type") == "file"
			if v, ok := os.LookupEnv(varName); ok {
				if isFile {
					return ioutil.ReadFile(v)
				} else {
					return []byte(v), nil
				}
			} else {
				return nil, fmt.Errorf("env variable %s not found", varName)
			}
		default:
			return []byte(configStr), nil
		}
	default:
		return nil, fmt.Errorf("unsupported type of config [%s]", k)
	}
}

// return plugin distribution name. i.e.: <Name>-<Version>-<OS>-<Arch>
func (m *PluginDefinition) FullName() string {
	return fmt.Sprintf("%s-%s-%s-%s", m.Name, m.Version, runtime.GOOS, runtime.GOARCH)
}

// return plugin distribution file name stored locally
func (m *PluginDefinition) DistFileName() string {
	return fmt.Sprintf("%s.zip", m.FullName())
}

// return plugin distribution signature file name stored locally
func (m *PluginDefinition) SignatureFileName() string {
	return fmt.Sprintf("%s.sha256sum.asc", m.DistFileName())
}

// must be always be lowercase when define constants
// as when unmarshaling from config, value will be case-lowered
type PluginInterfaceName string

// When this is used as a key in map. This function is not invoked.
func (p *PluginInterfaceName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*p = PluginInterfaceName(strings.ToLower(v))
	return nil
}

func (p *PluginInterfaceName) UnmarshalTOML(data []byte) error {
	var v string
	if err := toml.Unmarshal(data, &v); err != nil {
		return err
	}
	*p = PluginInterfaceName(strings.ToLower(v))
	return nil
}

func (p *PluginInterfaceName) UnmarshalText(data []byte) error {
	*p = PluginInterfaceName(strings.ToLower(string(data)))
	return nil
}

func (p PluginInterfaceName) String() string {
	return string(p)
}

// this defines plugins used in the geth node
type Settings struct {
	BaseDir       EnvironmentAwaredValue                   `json:"baseDir" toml:""`
	CentralConfig *PluginCentralConfiguration              `json:"central" toml:"Central"`
	Providers     map[PluginInterfaceName]PluginDefinition `json:"providers" toml:""`
}

func (s *Settings) GetPluginDefinition(name PluginInterfaceName) (*PluginDefinition, bool) {
	m, ok := s.Providers[name]
	return &m, ok
}

func (s *Settings) SetDefaults() {
	if s.CentralConfig == nil {
		s.CentralConfig = quorumPluginCentralConfiguration
	} else {
		s.CentralConfig.SetDefaults()
	}
}

// CheckSettingsAreSupported validates Settings by ensuring that only supportedPlugins are defined.
// It is not required for all supportedPlugins to be defined.
// An error containing plugin details is returned if one or more unsupported plugins are defined.
func (s *Settings) CheckSettingsAreSupported(supportedPlugins []PluginInterfaceName) error {
	errList := []PluginInterfaceName{}
	for name := range s.Providers {
		isValid := false
		for _, supportedPlugin := range supportedPlugins {
			if supportedPlugin == name {
				isValid = true
				break
			}
		}
		if !isValid {
			errList = append(errList, name)
		}
	}
	if len(errList) != 0 {
		return fmt.Errorf("unsupported plugins configured: %v", errList)
	}
	return nil
}

type PluginCentralConfiguration struct {
	// To implement certificate pinning while communicating with PluginCentral
	// if it's empty, we skip cert pinning logic
	CertFingerprint       string `json:"certFingerprint" toml:""`
	BaseURL               string `json:"baseURL" toml:""`
	PublicKeyURI          string `json:"publicKeyURI" toml:""`
	InsecureSkipTLSVerify bool   `json:"insecureSkipTLSVerify" toml:""`

	// URL path template to the plugin distribution file.
	// It uses Golang text template.
	PluginDistPathTemplate string `json:"pluginDistPathTemplate" toml:""`
	// URL path template to the plugin sha256 checksum signature file.
	// It uses Golang text template.
	PluginSigPathTemplate string `json:"pluginSigPathTemplate" toml:""`
}

// populate default values from quorumPluginCentralConfiguration
func (c *PluginCentralConfiguration) SetDefaults() {
	if len(c.BaseURL) == 0 {
		c.BaseURL = quorumPluginCentralConfiguration.BaseURL
	}
	if len(c.PublicKeyURI) == 0 {
		c.PublicKeyURI = quorumPluginCentralConfiguration.PublicKeyURI
	}
	if len(c.PluginDistPathTemplate) == 0 {
		c.PluginDistPathTemplate = quorumPluginCentralConfiguration.PluginDistPathTemplate
	}
	if len(c.PluginSigPathTemplate) == 0 {
		c.PluginSigPathTemplate = quorumPluginCentralConfiguration.PluginSigPathTemplate
	}
}

// support URI format with 'env' scheme during JSON/TOML/TEXT unmarshalling
// e.g.: env://FOO_VAR means read a string value from FOO_VAR environment variable
type EnvironmentAwaredValue string

func (d *EnvironmentAwaredValue) UnmarshalJSON(data []byte) error {
	return d.unmarshal(data)
}

func (d *EnvironmentAwaredValue) UnmarshalTOML(data []byte) error {
	return d.unmarshal(data)
}

func (d *EnvironmentAwaredValue) UnmarshalText(data []byte) error {
	return d.unmarshal(data)
}

func (d *EnvironmentAwaredValue) unmarshal(data []byte) error {
	v := string(data)
	isString := strings.HasPrefix(v, "\"") && strings.HasSuffix(v, "\"")
	if !isString {
		return fmt.Errorf("not a string")
	}
	v = strings.TrimFunc(v, func(r rune) bool {
		return r == '"'
	})
	if u, err := url.Parse(v); err == nil {
		switch u.Scheme {
		case "env":
			v = os.Getenv(u.Host)
		}
	}
	*d = EnvironmentAwaredValue(v)
	return nil
}

func (d EnvironmentAwaredValue) String() string {
	return string(d)
}
