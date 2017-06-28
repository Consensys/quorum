package constellation

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Socket         string   `toml:"socket"`
	PublicKeys     []string `toml:"publickeys"`

	// Deprecated
	SocketPath     string   `toml:"socketPath"`
	PublicKeyPath  string   `toml:"publicKeyPath"`
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
	if len(cfg.PublicKeys) == 0 {
		cfg.PublicKeys = append(cfg.PublicKeys, cfg.PublicKeyPath)
	}
	return cfg, nil
}
