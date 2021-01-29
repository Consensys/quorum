package private

import (
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"syscall"
	"testing"

	http2 "github.com/ethereum/go-ethereum/common/http"
	"github.com/ethereum/go-ethereum/private/engine/constellation"
	"github.com/ethereum/go-ethereum/private/engine/tessera"
	"github.com/stretchr/testify/assert"
)

func TestFromEnvironmentOrNil_whenNoConfig(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic received: %s", r)
		}
	}()
	os.Unsetenv("ARBITRARY_CONFIG_ENV")
	cfg, err := FromEnvironmentOrNil("ARBITRARY_CONFIG_ENV")

	assert.NoError(t, err, "unexpected error")
	assert.Equal(t, cfg.ConnectionType, http2.NoConnection, "expected no instance to be set")
}

func TestFromEnvironmentOrNil_whenUsingUnixSocketWithConstellation(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("this test case is not supported for windows")
	}
	testServer, socketFile := startUnixSocketHTTPServer(t, map[string]http.HandlerFunc{
		"/upcheck": MockEmptySuccessHandler,
	})
	defer testServer.Close()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic received: %s", r)
		}
	}()
	os.Setenv("ARBITRARY_CONFIG_ENV", socketFile)
	cfg, err := FromEnvironmentOrNil("ARBITRARY_CONFIG_ENV")
	p, err := NewPrivateTxManager(cfg)

	assert.NoError(t, err, "unexpected error")
	if !constellation.Is(p) {
		t.Errorf("expected Constellation to be used but found %v", reflect.TypeOf(p))
	}
}

func TestFromEnvironmentOrNil_whenUsingUnixSocketWithTessera(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("this test case is not supported for windows")
	}
	testServer, socketFile := startUnixSocketHTTPServer(t, map[string]http.HandlerFunc{
		"/upcheck": MockEmptySuccessHandler,
		"/version": MockEmptySuccessHandler,
	})
	defer testServer.Close()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic received: %s", r)
		}
	}()
	os.Setenv("ARBITRARY_CONFIG_ENV", socketFile)
	cfg, err := FromEnvironmentOrNil("ARBITRARY_CONFIG_ENV")
	p, err := NewPrivateTxManager(cfg)

	assert.NoError(t, err, "unexpected error")
	if !tessera.Is(p) {
		t.Errorf("expected Tessera to be used but found %v", reflect.TypeOf(p))
	}
}

func MockEmptySuccessHandler(_ http.ResponseWriter, _ *http.Request) {

}

func startUnixSocketHTTPServer(t *testing.T, handlers map[string]http.HandlerFunc) (*httptest.Server, string) {
	tmpFile := filepath.Join(os.TempDir(), "temp.sock")
	syscall.Unlink(tmpFile)
	l, err := net.Listen("unix", tmpFile)
	if err != nil {
		t.Fatalf("can't start a unix socket server due to %s", err)
	}
	os.Chmod(tmpFile, 0600)
	mux := http.NewServeMux()
	for k, v := range handlers {
		mux.HandleFunc(k, v)
	}

	testServer := httptest.Server{
		Listener: l,
		Config:   &http.Server{Handler: mux},
	}
	testServer.Start()
	t.Log("Unix Socket HTTP server started")
	return &testServer, tmpFile
}
