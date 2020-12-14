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

var socketConfigFileWithTimeouts = `
[socketConfig]
socket = "tm.ipc"
workdir = "qdata/c1"
dialTimeout = 8
requestTimeout = 9
responseHeaderTimeout = 10
`
var socketConfigFileNoTimeouts = `
[socketConfig]
socket = "tm.ipc"
workdir = "qdata/c1"
`
var httpConfigFileWithTimeouts = `
[httpConfig]
url = "http:localhost:9101"
clientTimeout = 99
`
var httpConfigFileNoTimeouts = `
[httpConfig]
url = "http:localhost:9101"
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

	cfg, err := FetchConfig(socketFile)
	if assert.NoError(t, err, "Failed to retrieve socket configuration") {
		assert.True(t, IsSocketConfigured(cfg), "IsSocketConfigured() returned false, when expecting true")
		assert.Equal(t, socketFile, filepath.Join(cfg.SocketConfig.WorkDir, cfg.SocketConfig.Socket), "Socket path unexpectedly changed when loading default config")
		assert.Equal(t, DefaultSocketTimeouts.DialTimeout, cfg.SocketConfig.DialTimeout, "Did not get expected socket default DialTimeout")
		assert.Equal(t, DefaultSocketTimeouts.RequestTimeout, cfg.SocketConfig.RequestTimeout, "Did not get expected socket default RequestTimeout")
		assert.Equal(t, DefaultSocketTimeouts.ResponseHeaderTimeout, cfg.SocketConfig.ResponseHeaderTimeout, "Did not get expected socket default ResponseHeaderTimeout")
	}
}

func TestLoadSocketConfigWithTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "config-example1.toml")
	if err := ioutil.WriteFile(configFile, []byte(socketConfigFileWithTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.True(t, IsSocketConfigured(cfg), "IsSocketConfigured() returned false, when expecting true")
		assert.Equal(t, "qdata/c1/tm.ipc", filepath.Join(cfg.SocketConfig.WorkDir, cfg.SocketConfig.Socket), "Did not get expected socket path from config file")
		assert.Equal(t, uint(8), cfg.SocketConfig.DialTimeout, "Did not get expected socket DialTimeout from config file")
		assert.Equal(t, uint(9), cfg.SocketConfig.RequestTimeout, "Did not get expected socket RequestTimeout from config file")
		assert.Equal(t, uint(10), cfg.SocketConfig.ResponseHeaderTimeout, "Did not get expected socket ResponseHeaderTimeout from config file")
	}
}

func TestLoadSocketConfigWithDefaultTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "config-example2.toml")
	if err := ioutil.WriteFile(configFile, []byte(socketConfigFileNoTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.True(t, IsSocketConfigured(cfg), "IsSocketConfigured() returned false, when expecting true")
		assert.Equal(t, DefaultSocketTimeouts.DialTimeout, cfg.SocketConfig.DialTimeout, "Did not get expected socket default DialTimeout from config file")
		assert.Equal(t, DefaultSocketTimeouts.RequestTimeout, cfg.SocketConfig.RequestTimeout, "Did not get expected socket default RequestTimeout from config file")
		assert.Equal(t, DefaultSocketTimeouts.ResponseHeaderTimeout, cfg.SocketConfig.ResponseHeaderTimeout, "Did not get expected socket default ResponseHeaderTimeout from config file")
	}
}

func TestLoadHttpConfigWithTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "config-example3.toml")
	if err := ioutil.WriteFile(configFile, []byte(httpConfigFileWithTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.False(t, IsSocketConfigured(cfg), "IsSocketConfigured() returned true, when expecting false")
		assert.Equal(t, "http:localhost:9101", cfg.HttpConfig.Url, "Did not get expected http url from config file")
		assert.Equal(t, uint(99), cfg.HttpConfig.ClientTimeout, "Did not get expected http clientTimeout from config file")
	}
}

func TestLoadHttpConfigWithDefaultTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "config-example4.toml")
	if err := ioutil.WriteFile(configFile, []byte(httpConfigFileNoTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.False(t, IsSocketConfigured(cfg), "IsSocketConfigured() returned true, when expecting false")
		assert.Equal(t, "http:localhost:9101", cfg.HttpConfig.Url, "Did not get expected http url from config file")
		assert.Equal(t, DefaultHttpTimeouts.ClientTimeout, cfg.HttpConfig.ClientTimeout, "Did not get expected http clientTimeout from config file")
	}
}
