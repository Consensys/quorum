package engine

import (
	"fmt"
	"os"
	"path/filepath"

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
	Url           string // transaction manager URL for HTTP connection
	ClientTimeout uint   // timeout for overall client call (seconds), not valid for IPC socket
}

var DefaultSocketTimeouts = socketConfig{
	DialTimeout:           1,
	RequestTimeout:        5,
	ResponseHeaderTimeout: 5,
}
var DefaultHttpTimeouts = httpConfig{
	ClientTimeout: 10,
}

func IsSocketConfigured(cfg Config) bool {
	return cfg.ConnectionType == socketConnection
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

	if cfg.SocketConfig.Socket != "" && cfg.HttpConfig.Url != "" {
		return fmt.Errorf("cannot specify both Socket and HTTP connections in config file")
	}

	if cfg.SocketConfig.Socket != "" {
		cfg.ConnectionType = socketConnection
	} else if cfg.HttpConfig.Url != "" {
		cfg.ConnectionType = httpConnection
	} else {
		return fmt.Errorf("either Socket or HTTP connection must be specified in config file")
	}

	return nil
}
