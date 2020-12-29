package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type connectionType uint

const (
	socketConnection connectionType = iota
	httpConnection
)

type Config struct {
	ConnectionType connectionType
	SocketConfig   socketConfig `toml:",omitempty"`
	HttpConfig     httpConfig   `toml:",omitempty"`
}

type socketConfig struct {
	Socket                string // socket filename
	WorkDir               string // path to socket file
	DialTimeout           uint   // timeout for connecting to socket (seconds)
	RequestTimeout        uint   // timeout for writing to socket (seconds)
	ResponseHeaderTimeout uint   // timeout for reading from socket (seconds)
}

type httpConfig struct {
	Url             string // transaction manager URL for HTTP connection
	Tls             string // if "OFF" then TLS disabled, if "STRICT" then TLS enabled
	RootCA          string // path to file containing certificate for root CA
	ClientCert      string // path to file containing client certificate (or chain of certs)
	ClientKey       string // path to file containing client's private key
	ClientTimeout   uint   // timeout for overall client call (seconds), zero means timeout disabled
	IdleConnTimeout uint   // timeout for idle connection (seconds), zero means no limit
	WriteBufferSize int    // size of the write buffer (bytes), if zero then uses http.Transport default
	ReadBufferSize  int    // size of the read buffer (bytes), if zero then uses http.Transport default
}

var DefaultSocketTimeouts = socketConfig{
	DialTimeout:           1,
	RequestTimeout:        5,
	ResponseHeaderTimeout: 5,
}
var DefaultHttpTimeouts = httpConfig{
	ClientTimeout:   10,
	IdleConnTimeout: 10,
}

func IsSocketConfigured(cfg Config) bool {
	return cfg.ConnectionType == socketConnection
}

func IsTlsConfigured(cfg Config) bool {
	return cfg.HttpConfig.Tls == "STRICT"
}

// FetchConfig sets up the configuration for the connection to a txn manager.
// It will accept a path to a socket file or a path to a config file,
// and returns the full configuration info for either a socket file or HTTP.
func FetchConfig(path string) (Config, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	isSocket := info.Mode()&os.ModeSocket != 0
	if isSocket {
		cfg.ConnectionType = socketConnection
		cfg.SocketConfig = DefaultSocketTimeouts
		cfg.SocketConfig.WorkDir, cfg.SocketConfig.Socket = filepath.Split(path)

	} else {
		err = LoadConfigFile(path, &cfg)
		if err != nil {
			return Config{}, err
		}
	}

	return cfg, nil
}

func LoadConfigFile(path string, cfg *Config) error {
	cfg.SocketConfig = DefaultSocketTimeouts
	cfg.HttpConfig = DefaultHttpTimeouts

	if _, err := toml.DecodeFile(path, cfg); err != nil {
		return err
	}
	cfg.HttpConfig.Tls = strings.ToUpper(cfg.HttpConfig.Tls)

	if cfg.SocketConfig.Socket != "" && cfg.HttpConfig.Url != "" {
		return fmt.Errorf("cannot specify both Socket and HTTP connections in config file")
	}

	if cfg.SocketConfig.Socket != "" {
		cfg.ConnectionType = socketConnection
	} else if cfg.HttpConfig.Url != "" {
		cfg.ConnectionType = httpConnection
		switch cfg.HttpConfig.Tls {
		case "OFF":
			//no action needed
		case "STRICT":
			if cfg.HttpConfig.RootCA == "" || cfg.HttpConfig.ClientCert == "" || cfg.HttpConfig.ClientKey == "" {
				return fmt.Errorf("missing details for HTTP connection with TLS, config file must specify: rootCA, clientCert, clientKey")
			}
		default:
			return fmt.Errorf("invalid value for 'Tls' in config file, must be either OFF or STRICT")
		}
	} else {
		return fmt.Errorf("either Socket or HTTP connection must be specified in config file")
	}

	return nil
}
