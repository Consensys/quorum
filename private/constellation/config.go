package constellation

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Url            string   `toml:"url"`
	Port           int      `toml:"port"`
	SocketPath     string   `toml:"socketPath"`
	OtherNodeUrls  []string `toml:"otherNodeUrls"`
	PublicKeyPath  string   `toml:"publicKeyPath"`
	PrivateKeyPath string   `toml:"privateKeyPath"`
	StoragePath    string   `toml:"storagePath"`
}

func LoadConfig(configPath string) (*Config, error) {
	cfg := new(Config)
	if _, err := toml.DecodeFile(configPath, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
