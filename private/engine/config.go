package engine

import (
	"os"
	"path/filepath"

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

var DefaultConfig = Config{
	DialTimeout:           1,
	RequestTimeout:        5,
	ResponseHeaderTimeout: 5,
}

// LoadConfig sets up the configuration for the connection to a txn manager.
// It will accept a path to a socket file or a path to a config file,
// and returns the full configuration info for the socket file.
func LoadConfig(path string) (Config, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return Config{}, err
	}

	cfg := DefaultConfig
	isSocket := info.Mode()&os.ModeSocket != 0
	if !isSocket {
		if _, err := toml.DecodeFile(path, &cfg); err != nil {
			return Config{}, err
		}
	} else {
		cfg.WorkDir, cfg.Socket = filepath.Split(path)
	}

	// Fall back to Constellation 0.0.1 config format if necessary
	if cfg.Socket == "" {
		cfg.Socket = cfg.SocketPath
	}
	return cfg, nil
}
