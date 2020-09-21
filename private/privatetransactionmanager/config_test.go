package privatetransactionmanager

import (
	"path/filepath"
	"testing"
)

func TestLoadConfigWithTimeouts(t *testing.T) {

	expectedPath := "qdata/c1/tm.ipc"
	var expectedDialTimeout uint = 8
	var expectedRequestTimeout uint = 9
	var expectedResponseHeaderTimeout uint = 10

	cfg, err := LoadConfig("config-example1.toml")
	if err != nil {
		t.Errorf("Failed to open config file: %v", err)
	}
	path := filepath.Join(cfg.WorkDir, cfg.Socket)
	if path != expectedPath {
		t.Errorf("Incorrect socket path from config file: got '%v', expected '%v'", path, expectedPath)
	}
	if cfg.DialTimeout != expectedDialTimeout {
		t.Errorf("Incorrect DialTimeout from config file: got '%v', expected '%v'", cfg.DialTimeout, expectedDialTimeout)
	}
	if cfg.RequestTimeout != expectedRequestTimeout {
		t.Errorf("Incorrect RequestTimeout from config file: got '%v', expected '%v'", cfg.RequestTimeout, expectedRequestTimeout)
	}
	if cfg.ResponseHeaderTimeout != expectedResponseHeaderTimeout {
		t.Errorf("Incorrect ResponseHeaderTimeout from config file: got '%v', expected '%v'", cfg.ResponseHeaderTimeout, expectedResponseHeaderTimeout)
	}
}

func TestLoadConfigWithDefaultTimeouts(t *testing.T) {

	expectedDialTimeout := DefaultConfig.DialTimeout
	expectedRequestTimeout := DefaultConfig.RequestTimeout
	expectedResponseHeaderTimeout := DefaultConfig.ResponseHeaderTimeout

	cfg, err := LoadConfig("config-example2.toml")
	if err != nil {
		t.Errorf("Failed to open config file: %v", err)
	}
	dialTimeout := cfg.DialTimeout
	if dialTimeout != expectedDialTimeout {
		t.Errorf("Unexpected default DialTimeout: got '%v', expected '%v'", dialTimeout, expectedDialTimeout)
	}
	requestTimeout := cfg.RequestTimeout
	if requestTimeout != expectedRequestTimeout {
		t.Errorf("Unexpected default RequestTimeout: got '%v', expected '%v'", requestTimeout, expectedRequestTimeout)
	}
	responseHeaderTimeout := cfg.ResponseHeaderTimeout
	if responseHeaderTimeout != expectedResponseHeaderTimeout {
		t.Errorf("Unexpected default ResponseHeaderTimeout: got '%v', expected '%v'", responseHeaderTimeout, expectedResponseHeaderTimeout)
	}
}
