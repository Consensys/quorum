package ethclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

const (
	arbitraryBase64Data = "YXJiaXRyYXJ5IGRhdGE=" // = "arbitrary data"
)

func TestPrivateTransactionManagerClient_storeRaw(t *testing.T) {
	// mock tessera client
	expectedData := []byte("arbitrary data")
	expectedDataEPH := common.BytesToEncryptedPayloadHash(expectedData)
	arbitraryServer := newStoreRawServer()
	defer arbitraryServer.Close()
	testObject, err := newPrivateTransactionManagerClient(arbitraryServer.URL)
	assert.NoError(t, err)

	key, err := testObject.StoreRaw([]byte("arbitrary payload"), "arbitrary private from")

	assert.NoError(t, err)
	assert.Equal(t, expectedDataEPH, key)
}

func newStoreRawServer() *httptest.Server {
	arbitraryResponse := fmt.Sprintf(`
{
	"key": "%s"
}
`, arbitraryBase64Data)
	mux := http.NewServeMux()
	mux.HandleFunc("/storeraw", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			// parse request
			var storeRawReq storeRawReq
			if err := json.NewDecoder(req.Body).Decode(&storeRawReq); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// send response
			_, _ = fmt.Fprintf(w, "%s", arbitraryResponse)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}

	})
	return httptest.NewServer(mux)
}
