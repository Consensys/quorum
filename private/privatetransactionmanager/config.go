package privatetransactionmanager

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Socket         string `toml:"socket"`  // socket filename
	WorkDir        string `toml:"workdir"` // path to socket file
	RequestTimeout uint   // optional override for request timeout (seconds)

	// Deprecated
	SocketPath string `toml:"socketPath"`
}

func LoadConfig(configPath string) (*Config, error) {
	cfg := new(Config)
	if _, err := toml.DecodeFile(configPath, cfg); err != nil {
		return nil, err
	}
	// Fall back to Constellation 0.0.1 config format if necessary
	if cfg.Socket == "" {
		cfg.Socket = cfg.SocketPath
	}
	return cfg, nil
}
