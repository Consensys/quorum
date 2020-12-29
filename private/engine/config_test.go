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
tls = "OFF"
clientTimeout = 101
idleConnTimeout = 102
writeBufferSize = 1001
readBufferSize = 1002
`
var httpConfigFileWithInvalidTls = `
[httpConfig]
url = "http:localhost:9101"
tls = "ABC"
`
var httpTlsConfigFileWithTimeouts = `
[httpConfig]
url = "http:localhost:9101"
tls = "STRICT"
rootCA = "mydir/rootca.cert.pem"
clientCert = "mydir/client.cert.pem"
clientKey = "mydir/client.key.pem"
clientTimeout = 101
idleConnTimeout = 102
writeBufferSize = 1001
readBufferSize = 1002
`
var httpConfigFileNoTimeouts = `
[httpConfig]
url = "http:localhost:9101"
Tls = "strict"
rootCA = "mydir/rootca.cert.pem"
clientCert = "mydir/client.cert.pem"
clientKey = "mydir/client.key.pem"
`
var invalidHttpTlsConfigFileNoCerts = `
[httpConfig]
url = "http:localhost:9101"
Tls = "STRICT"
`
var invalidConfigWithSocketAndHttp = `
[socketConfig]
socket = "tm.ipc"
workdir = "qdata/c1"
[httpConfig]
url = "http:localhost:9101"
`
var invalidConfigWithNoSocketOrHttp = `
socket = "tm.ipc"
workdir = "qdata/c1"
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
	configFile := filepath.Join(os.TempDir(), "socketConfigFileWithTimeouts.toml")
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
	configFile := filepath.Join(os.TempDir(), "socketConfigFileNoTimeouts.toml")
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
	configFile := filepath.Join(os.TempDir(), "httpConfigFileWithTimeouts.toml")
	if err := ioutil.WriteFile(configFile, []byte(httpConfigFileWithTimeouts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	cfg, err := FetchConfig(configFile)
	if assert.NoError(t, err, "Failed to load config file") {
		assert.False(t, IsSocketConfigured(cfg), "IsSocketConfigured() returned true, when expecting false")
		assert.Equal(t, "http:localhost:9101", cfg.HttpConfig.Url, "Did not get expected http url from config file")
		assert.False(t, IsTlsConfigured(cfg), "Did not get expected IsTlsConfigured() value from config file")
		assert.Equal(t, uint(101), cfg.HttpConfig.ClientTimeout, "Did not get expected http ClientTimeout from config file")
		assert.Equal(t, uint(102), cfg.HttpConfig.IdleConnTimeout, "Did not get expected http IdleConnTimeout from config file")
		assert.Equal(t, int(1001), cfg.HttpConfig.WriteBufferSize, "Did not get expected http WriteBufferSize from config file")
		assert.Equal(t, int(1002), cfg.HttpConfig.ReadBufferSize, "Did not get expected http ReadBufferSize from config file")
	}
}

func TestLoadHttpConfigWithInvalidTls(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "httpConfigFileWithInvalidTls.toml")
	if err := ioutil.WriteFile(configFile, []byte(httpConfigFileWithInvalidTls), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	_, err := FetchConfig(configFile)
	assert.EqualError(t, err, "invalid value for 'Tls' in config file, must be either OFF or STRICT")
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
		assert.Equal(t, "http:localhost:9101", cfg.HttpConfig.Url, "Did not get expected http url from config file")
		assert.True(t, IsTlsConfigured(cfg), "Did not get expected IsTlsConfigured() value from config file")
		assert.Equal(t, "mydir/rootca.cert.pem", cfg.HttpConfig.RootCA, "Did not get expected RootCA from config file")
		assert.Equal(t, "mydir/client.cert.pem", cfg.HttpConfig.ClientCert, "Did not get expected ClientCert from config file")
		assert.Equal(t, "mydir/client.key.pem", cfg.HttpConfig.ClientKey, "Did not get expected ClientKey from config file")
		assert.Equal(t, uint(101), cfg.HttpConfig.ClientTimeout, "Did not get expected http ClientTimeout from config file")
		assert.Equal(t, uint(102), cfg.HttpConfig.IdleConnTimeout, "Did not get expected http IdleConnTimeout from config file")
		assert.Equal(t, int(1001), cfg.HttpConfig.WriteBufferSize, "Did not get expected http WriteBufferSize from config file")
		assert.Equal(t, int(1002), cfg.HttpConfig.ReadBufferSize, "Did not get expected http ReadBufferSize from config file")
	}
}

func TestLoadHttpConfigWithDefaultTimeouts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "httpConfigFileNoTimeouts.toml")
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

func TestHTTPMissingCerts(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "invalidHttpTlsConfigFileNoCerts.toml")
	if err := ioutil.WriteFile(configFile, []byte(invalidHttpTlsConfigFileNoCerts), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	_, err := FetchConfig(configFile)
	assert.EqualError(t, err, "missing details for HTTP connection with TLS, config file must specify: rootCA, clientCert, clientKey")
}

func TestSocketWithHTTPNotAllowed(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "invalidConfigWithSocketAndHttp.toml")
	if err := ioutil.WriteFile(configFile, []byte(invalidConfigWithSocketAndHttp), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	_, err := FetchConfig(configFile)
	assert.EqualError(t, err, "cannot specify both Socket and HTTP connections in config file")
}

func TestEitherSocketOrHTTPMustBeSpecified(t *testing.T) {
	configFile := filepath.Join(os.TempDir(), "invalidConfigWithNoSocketOrHttp.toml")
	if err := ioutil.WriteFile(configFile, []byte(invalidConfigWithNoSocketOrHttp), 0600); err != nil {
		t.Fatalf("Failed to create config file for unit test, error: %v", err)
	}
	defer os.Remove(configFile)

	_, err := FetchConfig(configFile)
	assert.EqualError(t, err, "either Socket or HTTP connection must be specified in config file")
}
