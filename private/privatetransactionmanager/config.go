package privatetransactionmanager

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Socket  string `toml:"socket"`  // socket filename
	WorkDir string `toml:"workdir"` // path to socket file

	DialTimeout           uint // timeout for connecting to socket (seconds)
	RequestTimeout        uint // timeout for writing to socket (seconds)
	ResponseHeaderTimeout uint // timeout for reading from socket (seconds)

	// Deprecated
	SocketPath string `toml:"socketPath"`
}

var DefaultConfig = &Config{
	DialTimeout:           1,
	RequestTimeout:        5,
	ResponseHeaderTimeout: 5,
}

func LoadConfig(configPath string) (*Config, error) {
	cfg := *DefaultConfig
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, err
	}
	// Fall back to Constellation 0.0.1 config format if necessary
	if cfg.Socket == "" {
		cfg.Socket = cfg.SocketPath
	}
	return &cfg, nil
}
