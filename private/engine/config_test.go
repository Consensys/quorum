package engine

import (
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

var configFileWithTimeouts = `
    socket = "tm.ipc"
    workdir = "qdata/c1"
    dialTimeout = 8
    requestTimeout = 9
    responseHeaderTimeout = 10
`
var configFileNoTimeouts = `
    socket = "tm.ipc"
    workdir = "qdata/c1"
`

func TestDefaultTimeoutsUsedWhenNoConfigFileSpecified(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("this test case is not supported for windows")
	}

	socketFile := filepath.Join(os.TempDir(), "socket-file.ipc")
	syscall.Unlink(socketFile)
	l, err := net.Listen("unix", socketFile)
	if err != nil {
		t.Fatalf("Could not create socket file '%v' for unit test, error: %v", socketFile, err)
	}
	defer l.Close()

	cfg, err := LoadConfig(socketFile)
	if assert.NoError(t, err, "Failed to retrieve socket configuration") {
		assert.Equal(t, socketFile, filepath.Join(cfg.WorkDir, cfg.Socket), "Socket path unexpectedly changed when loading default config")
		assert.Equal(t, DefaultConfig.DialTimeout, cfg.DialTimeout, "Did not get expected default DialTimeout")
		assert.Equal(t, DefaultConfig.RequestTimeout, cfg.RequestTimeout, "Did not get expected default RequestTimeout")
		assert.Equal(t, DefaultConfig.ResponseHeaderTimeout, cfg.ResponseHeaderTimeout, "Did not get expected default ResponseHeaderTimeout")
	}
}

func TestLoadConfigWithTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "config-example1.toml")
	if err := ioutil.WriteFile(configFile, []byte(configFileWithTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := LoadConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.Equal(t, "qdata/c1/tm.ipc", filepath.Join(cfg.WorkDir, cfg.Socket), "Did not get expected socket path from config file")
		assert.Equal(t, uint(8), cfg.DialTimeout, "Did not get expected DialTimeout from config file")
		assert.Equal(t, uint(9), cfg.RequestTimeout, "Did not get expected RequestTimeout from config file")
		assert.Equal(t, uint(10), cfg.ResponseHeaderTimeout, "Did not get expected ResponseHeaderTimeout from config file")
	}
}

func TestLoadConfigWithDefaultTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "config-example2.toml")
	if err := ioutil.WriteFile(configFile, []byte(configFileNoTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := LoadConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.Equal(t, DefaultConfig.DialTimeout, cfg.DialTimeout, "Did not get expected default DialTimeout from config file")
		assert.Equal(t, DefaultConfig.RequestTimeout, cfg.RequestTimeout, "Did not get expected default RequestTimeout from config file")
		assert.Equal(t, DefaultConfig.ResponseHeaderTimeout, cfg.ResponseHeaderTimeout, "Did not get expected default ResponseHeaderTimeout from config file")
	}
}
