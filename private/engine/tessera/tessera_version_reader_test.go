package tessera

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/private/engine"
	testifyassert "github.com/stretchr/testify/assert"
)

func TestVersionApi_404NotFound(t *testing.T) {
	assert := testifyassert.New(t)

	mux := http.NewServeMux()

	testServer = httptest.NewServer(mux)
	defer testServer.Close()

	version := RetrieveTesseraAPIVersion(&engine.Client{
		HttpClient: &http.Client{},
		BaseURL:    testServer.URL,
	})

	assert.Equal(apiVersion1, version)
}

func TestVersionApi_GarbageData(t *testing.T) {
	assert := testifyassert.New(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/version/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("GARBAGE"))
	})

	testServer = httptest.NewServer(mux)
	defer testServer.Close()

	version := RetrieveTesseraAPIVersion(&engine.Client{
		HttpClient: &http.Client{},
		BaseURL:    testServer.URL,
	})

	assert.Equal(apiVersion1, version)
}

func TestVersionApi_emptyVersionsArray(t *testing.T) {
	assert := testifyassert.New(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/version/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("[]"))
	})

	testServer = httptest.NewServer(mux)
	defer testServer.Close()

	version := RetrieveTesseraAPIVersion(&engine.Client{
		HttpClient: &http.Client{},
		BaseURL:    testServer.URL,
	})

	assert.Equal(apiVersion1, version)
}

func TestVersionApi_invalidVersionItem(t *testing.T) {
	assert := testifyassert.New(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/version/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("{\"versions\":[{\"version\":\"1.0\"},{}]}"))
	})

	testServer = httptest.NewServer(mux)
	defer testServer.Close()

	version := RetrieveTesseraAPIVersion(&engine.Client{
		HttpClient: &http.Client{},
		BaseURL:    testServer.URL,
	})

	assert.Equal(apiVersion1, version)
}

func TestVersionApi_validVersionInWrongOrder(t *testing.T) {
	assert := testifyassert.New(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/version/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("[\"2.0\",\"1.0\"]"))
	})

	testServer = httptest.NewServer(mux)
	defer testServer.Close()

	version := RetrieveTesseraAPIVersion(&engine.Client{
		HttpClient: &http.Client{},
		BaseURL:    testServer.URL,
	})

	assert.Equal("2.0", version)
}

func TestVersionApi_validVersion(t *testing.T) {
	assert := testifyassert.New(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/version/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("[\"1.0\",\"2.0\"]"))
	})

	testServer = httptest.NewServer(mux)
	defer testServer.Close()

	version := RetrieveTesseraAPIVersion(&engine.Client{
		HttpClient: &http.Client{},
		BaseURL:    testServer.URL,
	})

	assert.Equal("2.0", version)
}
