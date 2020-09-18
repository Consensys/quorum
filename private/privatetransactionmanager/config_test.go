package privatetransactionmanager

import (
	"path/filepath"
	"testing"
)

func TestLoadConfigWithRequestTimeout(t *testing.T) {

	expectedPath := "qdata/c1/tm.ipc"
	var expectedRequestTimeout uint = 10

	cfg, err := LoadConfig("config-example1.toml")
	if err != nil {
		t.Errorf("Failed to open config file: %v", err)
	}
	path := filepath.Join(cfg.WorkDir, cfg.Socket)
	if path != expectedPath {
		t.Errorf("Incorrect socket path read from config file: got '%v', expected '%v'", path, expectedPath)
	}
	requestTimeout := cfg.RequestTimeout
	if requestTimeout != expectedRequestTimeout {
		t.Errorf("Incorrect readTimeout read from config file: got '%v', expected '%v'", requestTimeout, expectedRequestTimeout)
	}
}

func TestLoadConfigWithoutRequestTimeout(t *testing.T) {

	var expectedRequestTimeout uint = 0

	cfg, err := LoadConfig("config-example2.toml")
	if err != nil {
		t.Errorf("Failed to open config file: %v", err)
	}
	requestTimeout := cfg.RequestTimeout
	if requestTimeout != expectedRequestTimeout {
		t.Errorf("Incorrect readTimeout read from config file: got '%v', expected '%v'", requestTimeout, expectedRequestTimeout)
	}
}
