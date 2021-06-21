package ethclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

type privateTransactionManagerClient interface {
	StoreRaw(data []byte, privateFrom string) (common.EncryptedPayloadHash, error)
}

type privateTransactionManagerDefaultClient struct {
	rawurl     string
	httpClient *http.Client
}

// Create a new client to interact with private transaction manager via a HTTP endpoint
func newPrivateTransactionManagerClient(endpoint string) (privateTransactionManagerClient, error) {
	_, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &privateTransactionManagerDefaultClient{
		rawurl:     endpoint,
		httpClient: &http.Client{},
	}, nil
}

type storeRawReq struct {
	Payload string `json:"payload"`
	From    string `json:"from,omitempty"`
}

type storeRawResp struct {
	Key string `json:"key"`
}

func (pc *privateTransactionManagerDefaultClient) StoreRaw(data []byte, privateFrom string) (common.EncryptedPayloadHash, error) {
	storeRawReq := &storeRawReq{
		Payload: base64.StdEncoding.EncodeToString(data),
		From:    privateFrom,
	}
	reqBodyBuf := new(bytes.Buffer)
	if err := json.NewEncoder(reqBodyBuf).Encode(storeRawReq); err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	resp, err := pc.httpClient.Post(pc.rawurl+"/storeraw", "application/json", reqBodyBuf)
	if err != nil {
		return common.EncryptedPayloadHash{}, fmt.Errorf("unable to invoke /storeraw due to %s", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return common.EncryptedPayloadHash{}, fmt.Errorf("server returns %s", resp.Status)
	}
	// parse response
	var storeRawResp storeRawResp
	if err := json.NewDecoder(resp.Body).Decode(&storeRawResp); err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	encryptedPayloadHash, err := common.Base64ToEncryptedPayloadHash(storeRawResp.Key)
	if err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	return encryptedPayloadHash, nil
}
