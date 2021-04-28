package plugin

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCentralClient_PublicKey(t *testing.T) {
	arbitraryServer := newTestServer("/"+DefaultPublicKeyFile, arbitraryPubKey)
	defer arbitraryServer.Close()
	arbitraryConfig := &PluginCentralConfiguration{
		BaseURL:      arbitraryServer.URL,
		PublicKeyURI: DefaultPublicKeyFile,
	}

	testObject := NewPluginCentralClient(arbitraryConfig)

	actualValue, err := testObject.PublicKey()

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, arbitraryPubKey, actualValue)
}

func TestCentralClient_PublicKey_RelativePath(t *testing.T) {
	arbitraryServer := newTestServer("/aa/"+DefaultPublicKeyFile, arbitraryPubKey)
	defer arbitraryServer.Close()
	arbitraryConfig := &PluginCentralConfiguration{
		BaseURL:      arbitraryServer.URL + "/aa/bb/cc/", // without postfix /, the relative path would be diff
		PublicKeyURI: "../../" + DefaultPublicKeyFile,
	}

	testObject := NewPluginCentralClient(arbitraryConfig)

	actualValue, err := testObject.PublicKey()

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, arbitraryPubKey, actualValue)
}

func TestCentralClient_PublicKey_withSSL(t *testing.T) {
	arbitraryServer := httptest.NewTLSServer(newMux("/"+DefaultPublicKeyFile, arbitraryPubKey))
	defer arbitraryServer.Close()
	arbitraryConfig := &PluginCentralConfiguration{
		CertFingerprint:       string(arbitraryServer.Certificate().Signature),
		BaseURL:               arbitraryServer.URL,
		PublicKeyURI:          DefaultPublicKeyFile,
		InsecureSkipTLSVerify: true,
	}

	testObject := NewPluginCentralClient(arbitraryConfig)

	actualValue, err := testObject.PublicKey()

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, arbitraryPubKey, actualValue)
}

func TestCentralClient_PluginSignature(t *testing.T) {
	arbitraryDef := &PluginDefinition{
		Name:    "arbitrary-plugin",
		Version: "1.0.0",
	}
	expectedPath := fmt.Sprintf("/maven/bin/%s/%s/%s-%s-%s-%s-sha256.checksum.asc", arbitraryDef.Name, arbitraryDef.Version, arbitraryDef.Name, arbitraryDef.Version, runtime.GOOS, runtime.GOARCH)
	arbitraryServer := newTestServer(expectedPath, validSignature)
	defer arbitraryServer.Close()
	arbitraryConfig := &PluginCentralConfiguration{
		BaseURL: arbitraryServer.URL,
	}
	arbitraryConfig.SetDefaults()

	testObject := NewPluginCentralClient(arbitraryConfig)

	actualValue, err := testObject.PluginSignature(arbitraryDef)

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, validSignature, actualValue)
}

func TestCentralClient_PluginDistribution(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "q-")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()
	arbitraryDef := &PluginDefinition{
		Name:    "arbitrary-plugin",
		Version: "1.0.0",
	}
	arbitraryData := []byte("arbitrary data")
	expectedPath := fmt.Sprintf("/maven/bin/%s/%s/%s-%s-%s-%s.zip", arbitraryDef.Name, arbitraryDef.Version, arbitraryDef.Name, arbitraryDef.Version, runtime.GOOS, runtime.GOARCH)
	arbitraryServer := newTestServer(expectedPath, arbitraryData)
	defer arbitraryServer.Close()
	arbitraryConfig := &PluginCentralConfiguration{
		BaseURL: arbitraryServer.URL,
	}
	arbitraryConfig.SetDefaults()

	testObject := NewPluginCentralClient(arbitraryConfig)

	err = testObject.PluginDistribution(arbitraryDef, path.Join(tmpDir, "download.zip"))

	assert.NoError(t, err)
}

func newTestServer(pattern string, returnedData []byte) *httptest.Server {
	return httptest.NewServer(newMux(pattern, returnedData))
}

func newMux(pattern string, returnedData []byte) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(pattern, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_, err := io.Copy(w, bytes.NewReader(returnedData))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))
	return mux
}
