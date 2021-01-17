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
socket = "tm.ipc"
workdir = "qdata/c1"
dialTimeout = 8
timeout = 9
`
var socketConfigFileNoTimeouts = `
socket = "tm.ipc"
workdir = "qdata/c1"
`
var httpConfigFileWithTimeouts = `
httpUrl = "http:localhost:9101"
tlsMode = "OFF"
timeout = 101
httpIdleConnTimeout = 102
httpWriteBufferSize = 1001
httpReadBufferSize = 1002
`
var httpConfigFileWithInvalidTls = `
httpUrl = "http:localhost:9101"
tlsMode = "ABC"
`
var httpTlsConfigFileWithTimeouts = `
httpUrl = "https:localhost:9101"
tlsMode = "STRICT"
tlsRootCA = "mydir/rootca.cert.pem"
tlsClientCert = "mydir/client.cert.pem"
tlsClientKey = "mydir/client.key.pem"
timeout = 101
httpIdleConnTimeout = 102
httpWriteBufferSize = 1001
httpReadBufferSize = 1002
`
var httpTlsConfigFileNoTimeouts = `
httpUrl = "https:localhost:9101"
tlsMode = "strict"
tlsRootCA = "mydir/rootca.cert.pem"
tlsClientCert = "mydir/client.cert.pem"
tlsClientKey = "mydir/client.key.pem"
`
var invalidHttpTlsConfigFileNoCerts = `
httpUrl = "https:localhost:9101"
tlsMode = "STRICT"
`
var httpTlsConfigFileWithHTTPOnly = `
httpUrl = "http:localhost:9101"
tlsMode = "strict"
tlsRootCA = "mydir/rootca.cert.pem"
tlsClientCert = "mydir/client.cert.pem"
tlsClientKey = "mydir/client.key.pem"
`
var invalidConfigWithSocketAndHttp = `
socket = "tm.ipc"
workdir = "qdata/c1"
httpUrl = "http:localhost:9101"
`
var invalidConfigWithNoSocketOrHttp = `
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
		assert.Equal(t, socketFile, filepath.Join(cfg.WorkDir, cfg.Socket), "Socket path unexpectedly changed when loading default config")
		assert.Equal(t, DefaultConfig.DialTimeout, cfg.DialTimeout, "Did not get expected socket default DialTimeout")
		assert.Equal(t, DefaultConfig.Timeout, cfg.Timeout, "Did not get expected socket default Timeout")
	}

	err = cfg.Validate()
	assert.NoError(t, err)
}

func TestLoadSocketConfigWithTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "socketConfigFileWithTimeouts.toml")
	if err := ioutil.WriteFile(configFile, []byte(socketConfigFileWithTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.True(t, IsSocketConfigured(cfg), "IsSocketConfigured() returned false, when expecting true")
		assert.Equal(t, "qdata/c1/tm.ipc", filepath.Join(cfg.WorkDir, cfg.Socket), "Did not get expected socket path from config file")
		assert.Equal(t, uint(8), cfg.DialTimeout, "Did not get expected socket DialTimeout from config file")
		assert.Equal(t, uint(9), cfg.Timeout, "Did not get expected socket Timeout from config file")
	}

	err = cfg.Validate()
	assert.NoError(t, err)
}

func TestLoadSocketConfigWithDefaultTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "socketConfigFileNoTimeouts.toml")
	if err := ioutil.WriteFile(configFile, []byte(socketConfigFileNoTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.True(t, IsSocketConfigured(cfg), "IsSocketConfigured() returned false, when expecting true")
		assert.Equal(t, DefaultConfig.DialTimeout, cfg.DialTimeout, "Did not get expected socket default DialTimeout from config file")
		assert.Equal(t, DefaultConfig.Timeout, cfg.Timeout, "Did not get expected socket default Timeout from config file")
	}

	err = cfg.Validate()
	assert.NoError(t, err)
}

func TestLoadHttpConfigWithTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "httpConfigFileWithTimeouts.toml")
	if err := ioutil.WriteFile(configFile, []byte(httpConfigFileWithTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.False(t, IsSocketConfigured(cfg), "IsSocketConfigured() returned true, when expecting false")
		assert.Equal(t, "http:localhost:9101", cfg.HttpUrl, "Did not get expected http url from config file")
		assert.Equal(t, cfg.TlsMode, TlsOff, "Did not get expected IsTlsConfigured() value from config file")
		assert.Equal(t, uint(101), cfg.Timeout, "Did not get expected http Timeout from config file")
		assert.Equal(t, uint(102), cfg.HttpIdleConnTimeout, "Did not get expected http HttpIdleConnTimeout from config file")
		assert.Equal(t, int(1001), cfg.HttpWriteBufferSize, "Did not get expected http HttpWriteBufferSize from config file")
		assert.Equal(t, int(1002), cfg.HttpReadBufferSize, "Did not get expected http HttpReadBufferSize from config file")
	}

	err = cfg.Validate()
	assert.NoError(t, err)
}

func TestLoadHttpConfigWithInvalidTls(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "httpConfigFileWithInvalidTls.toml")
	if err := ioutil.WriteFile(configFile, []byte(httpConfigFileWithInvalidTls), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	assert.NoError(t, err)

	err = cfg.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "invalid value for TLS mode in config file, must be either OFF or STRICT")
	}
}

func TestLoadHttpTlsConfigWithTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "httpTlsConfigFileWithTimeouts.toml")
	if err := ioutil.WriteFile(configFile, []byte(httpTlsConfigFileWithTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.False(t, IsSocketConfigured(cfg), "IsSocketConfigured() returned true, when expecting false")
		assert.Equal(t, "https:localhost:9101", cfg.HttpUrl, "Did not get expected http url from config file")
		assert.Equal(t, cfg.TlsMode, TlsStrict, "Did not get expected IsTlsConfigured() value from config file")
		assert.Equal(t, "mydir/rootca.cert.pem", cfg.TlsRootCA, "Did not get expected TlsRootCA from config file")
		assert.Equal(t, "mydir/client.cert.pem", cfg.TlsClientCert, "Did not get expected TlsClientCert from config file")
		assert.Equal(t, "mydir/client.key.pem", cfg.TlsClientKey, "Did not get expected TlsClientKey from config file")
		assert.Equal(t, uint(101), cfg.Timeout, "Did not get expected http Timeout from config file")
		assert.Equal(t, uint(102), cfg.HttpIdleConnTimeout, "Did not get expected http HttpIdleConnTimeout from config file")
		assert.Equal(t, int(1001), cfg.HttpWriteBufferSize, "Did not get expected http HttpWriteBufferSize from config file")
		assert.Equal(t, int(1002), cfg.HttpReadBufferSize, "Did not get expected http HttpReadBufferSize from config file")
	}

	err = cfg.Validate()
	assert.NoError(t, err)
}

func TestLoadHttpConfigWithDefaultTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "httpTlsConfigFileNoTimeouts.toml")
	if err := ioutil.WriteFile(configFile, []byte(httpTlsConfigFileNoTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.False(t, IsSocketConfigured(cfg), "IsSocketConfigured() returned true, when expecting false")
		assert.Equal(t, "https:localhost:9101", cfg.HttpUrl, "Did not get expected http url from config file")
		assert.Equal(t, DefaultConfig.Timeout, cfg.Timeout, "Did not get expected http Timeout from config file")
	}

	err = cfg.Validate()
	assert.NoError(t, err)
}

func TestHTTPMissingCerts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "invalidHttpTlsConfigFileNoCerts.toml")
	if err := ioutil.WriteFile(configFile, []byte(invalidHttpTlsConfigFileNoCerts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	assert.NoError(t, err)

	err = cfg.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "missing details for HTTP connection with TLS, configuration must specify: rootCA, clientCert, clientKey")
	}
}

func TestTlsWithHTTPOnly(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "httpTlsConfigFileWithHTTPOnly.toml")
	if err := ioutil.WriteFile(configFile, []byte(httpTlsConfigFileWithHTTPOnly), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	assert.NoError(t, err)

	err = cfg.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "connection is configured with TLS but HTTPS url is not specified")
	}
}

func TestSocketWithHTTPNotAllowed(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "invalidConfigWithSocketAndHttp.toml")
	if err := ioutil.WriteFile(configFile, []byte(invalidConfigWithSocketAndHttp), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	assert.NoError(t, err)

	err = cfg.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "HTTP URL and unix ipc file cannot both be specified for private transaction manager connection")
	}
}

func TestEitherSocketOrHTTPMustBeSpecified(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "invalidConfigWithNoSocketOrHttp.toml")
	if err := ioutil.WriteFile(configFile, []byte(invalidConfigWithNoSocketOrHttp), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	_, err := FetchConfig(configFile)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "either Socket or HTTP connection must be specified in config file")
	}
}
