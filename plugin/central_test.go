package plugin

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
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
	arbitraryServer := newTestServer("/"+arbitraryDef.RemotePath()+"/"+arbitraryDef.SignatureFileName(), validSignature)
	defer arbitraryServer.Close()
	arbitraryConfig := &PluginCentralConfiguration{
		BaseURL: arbitraryServer.URL,
	}

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
	arbitraryServer := newTestServer("/"+arbitraryDef.RemotePath()+"/"+arbitraryDef.DistFileName(), arbitraryData)
	defer arbitraryServer.Close()
	arbitraryConfig := &PluginCentralConfiguration{
		BaseURL: arbitraryServer.URL,
	}

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
