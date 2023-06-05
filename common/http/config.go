package http

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

const (
	NoConnection               string = "none"
	UnixDomainSocketConnection string = "unix"
	HttpConnection             string = "http"
)

const (
	TlsOff    string = "off"
	TlsStrict string = "strict"
)

type Config struct {
	ConnectionType        string `toml:"-"` // connection type is not loaded from toml
	Socket                string // filename for unix domain socket
	WorkDir               string // directory for unix domain socket
	HttpUrl               string // transaction manager URL for HTTP connection
	Timeout               uint   // timeout for overall client call (seconds), zero means timeout disabled
	DialTimeout           uint   // timeout for connecting to unix socket (seconds)
	HttpIdleConnTimeout   uint   // timeout for idle http connection (seconds), zero means timeout disabled
	HttpWriteBufferSize   int    // size of http connection write buffer (bytes), if zero then uses http.Transport default
	HttpReadBufferSize    int    // size of http connection read buffer (bytes), if zero then uses http.Transport default
	TlsMode               string // whether TLS is enabled on HTTP connection (can be "off" or "strict")
	TlsRootCA             string // path to file containing certificate for root CA (defaults to host's certificates)
	TlsClientCert         string // path to file containing client certificate (or chain of certs)
	TlsClientKey          string // path to file containing client's private key
	TlsInsecureSkipVerify bool   // if true then does not verify that server certificate is CA signed
}

var NoConnectionConfig = Config{
	ConnectionType: NoConnection,
	TlsMode:        TlsOff,
}

var DefaultConfig = Config{
	Timeout:             5,
	DialTimeout:         1,
	HttpIdleConnTimeout: 10,
	TlsMode:             TlsOff,
}

func IsSocketConfigured(cfg Config) bool {
	return cfg.ConnectionType == UnixDomainSocketConnection
}

// This will accept path as any of the following and return relevant configuration:
//   - path set to "ignore"
//   - path to an ipc file
//   - path to a config file
func FetchConfigOrIgnore(path string) (Config, error) {
	if path == "" || strings.EqualFold(path, "ignore") {
		return NoConnectionConfig, nil
	}

	return FetchConfig(path)
}

// FetchConfig sets up the configuration for the connection to a txn manager.
// It will accept a path to an ipc file or a path to a config file,
// and returns the full configuration info for the specified type of connection.
func FetchConfig(path string) (Config, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return Config{}, fmt.Errorf("unable to check whether connection details are specified as a config file or ipc file '%s', due to: %s", path, err)
	}

	var cfg Config
	isSocket := info.Mode()&os.ModeSocket != 0
	if isSocket {
		cfg = DefaultConfig
		cfg.ConnectionType = UnixDomainSocketConnection
		cfg.WorkDir, cfg.Socket = filepath.Split(path)
	} else {
		cfg, err = LoadConfigFile(path)
		if err != nil {
			return Config{}, fmt.Errorf("error reading config from '%s' due to: %s", path, err)
		}
	}

	return cfg, nil
}

func LoadConfigFile(path string) (Config, error) {
	cfg := DefaultConfig
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return Config{}, err
	}
	cfg.TlsMode = strings.ToLower(cfg.TlsMode)

	if cfg.Socket != "" {
		cfg.ConnectionType = UnixDomainSocketConnection
	} else if cfg.HttpUrl != "" {
		cfg.ConnectionType = HttpConnection
	} else {
		return Config{}, fmt.Errorf("either Socket or HTTP connection must be specified in config file")
	}

	return cfg, nil
}

func (cfg *Config) Validate() error {
	switch cfg.ConnectionType {
	case "": // no connection type defined
	case NoConnection:
	case UnixDomainSocketConnection:
		if len(cfg.Socket) == 0 { //sanity check - should never occur
			return fmt.Errorf("ipc file configuration is missing for private transaction manager connection")
		}
		if len(cfg.HttpUrl) != 0 {
			return fmt.Errorf("HTTP URL and unix ipc file cannot both be specified for private transaction manager connection")
		}
		if cfg.TlsMode != TlsOff {
			return fmt.Errorf("TLS is not supported over unix domain socket for private transaction manager connection")
		}
	case HttpConnection:
		if len(cfg.Socket) != 0 {
			return fmt.Errorf("HTTP URL and unix ipc file cannot both be specified for private transaction manager connection")
		}
		if len(cfg.HttpUrl) == 0 { //sanity check - should never occur
			return fmt.Errorf("URL configuration is missing for private transaction manager HTTP connection")
		}
		switch cfg.TlsMode {
		case TlsOff:
			//no action needed
		case TlsStrict:
			if !strings.Contains(strings.ToLower(cfg.HttpUrl), "https") {
				return fmt.Errorf("connection is configured with TLS but HTTPS url is not specified")
			}
			if (len(cfg.TlsClientCert) == 0 && len(cfg.TlsClientKey) != 0) || (len(cfg.TlsClientCert) != 0 && len(cfg.TlsClientKey) == 0) {
				return fmt.Errorf("invalid details for HTTP connection with TLS, configuration must specify both clientCert and clientKey, or neither one")
			}
		default:
			return fmt.Errorf("invalid value for TLS mode in config file, must be either OFF or STRICT")
		}
	}

	return nil
}

//
// Setters for the various config fields
//

func (cfg *Config) SetSocket(socketPath string) {
	cfg.ConnectionType = UnixDomainSocketConnection
	workDir, socketFilename := filepath.Split(socketPath)
	if workDir != "" {
		cfg.WorkDir = workDir
	}
	cfg.Socket = socketFilename
}

func (cfg *Config) SetHttpUrl(httpUrl string) {
	cfg.ConnectionType = HttpConnection
	cfg.HttpUrl = httpUrl
}

func (cfg *Config) SetTimeout(timeout uint) {
	cfg.Timeout = timeout
}

func (cfg *Config) SetDialTimeout(dialTimeout uint) {
	cfg.DialTimeout = dialTimeout
}

func (cfg *Config) SetHttpIdleConnTimeout(httpIdleConnTimeout uint) {
	cfg.HttpIdleConnTimeout = httpIdleConnTimeout
}

func (cfg *Config) SetHttpWriteBufferSize(httpWriteBufferSize int) {
	cfg.HttpWriteBufferSize = httpWriteBufferSize
}

func (cfg *Config) SetHttpReadBufferSize(httpReadBufferSize int) {
	cfg.HttpReadBufferSize = httpReadBufferSize
}

func (cfg *Config) SetTlsMode(tlsMode string) {
	cfg.TlsMode = tlsMode
}

func (cfg *Config) SetTlsRootCA(tlsRootCA string) {
	cfg.TlsRootCA = tlsRootCA
}

func (cfg *Config) SetTlsClientCert(tlsClientCert string) {
	cfg.TlsClientCert = tlsClientCert
}

func (cfg *Config) SetTlsClientKey(tlsClientKey string) {
	cfg.TlsClientKey = tlsClientKey
}

func (cfg *Config) SetTlsInsecureSkipVerify(tlsInsecureSkipVerify bool) {
	cfg.TlsInsecureSkipVerify = tlsInsecureSkipVerify
}
