package ethclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type privateTransactionManagerClient interface {
	StoreRaw(data []byte, privateFrom string) ([]byte, error)
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
	return newPrivateTransactionManagerClientNoValidation(endpoint, &http.Client{}), nil
}

func newPrivateTransactionManagerClientNoValidation(endpoint string, httpclient *http.Client) privateTransactionManagerClient {
	return &privateTransactionManagerDefaultClient{
		rawurl:     endpoint,
		httpClient: httpclient,
	}
}

type storeRawReq struct {
	Payload string `json:"payload"`
	From    string `json:"from,omitempty"`
}

type storeRawResp struct {
	Key string `json:"key"`
}

func (pc *privateTransactionManagerDefaultClient) StoreRaw(data []byte, privateFrom string) ([]byte, error) {
	storeRawReq := &storeRawReq{
		Payload: base64.StdEncoding.EncodeToString(data),
		From:    privateFrom,
	}
	reqBodyBuf := new(bytes.Buffer)
	if err := json.NewEncoder(reqBodyBuf).Encode(storeRawReq); err != nil {
		return nil, err
	}
	resp, err := pc.httpClient.Post(pc.rawurl+"/storeraw", "application/json", reqBodyBuf)
	if err != nil {
		return nil, fmt.Errorf("unable to invoke /storeraw due to %s", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returns %s", resp.Status)
	}
	// parse response
	var storeRawResp storeRawResp
	if err := json.NewDecoder(resp.Body).Decode(&storeRawResp); err != nil {
		return nil, err
	}
	encryptedPayloadHash, err := base64.StdEncoding.DecodeString(storeRawResp.Key)
	if err != nil {
		return nil, err
	}
	return encryptedPayloadHash, nil
}
