package constellation

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Client struct {
	httpClient *http.Client
}

func (c *Client) SendPayload(pl []byte, b64From string, b64To []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) (common.EncryptedPayloadHash, error) {
	method := "POST"
	url := "http+unix://c/sendraw"
	buf := bytes.NewBuffer(pl)
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return common.EncryptedPayloadHash{}, fmt.Errorf("unable to build request for (method:%s,path:%s). Cause: %v", method, url, err)
	}
	if b64From != "" {
		req.Header.Set("c11n-from", b64From)
	}
	req.Header.Set("c11n-to", strings.Join(b64To, ","))
	req.Header.Set("Content-Type", "application/octet-stream")
	res, err := c.httpClient.Do(req)

	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return common.EncryptedPayloadHash{}, fmt.Errorf("unable to submit request (method:%s,path:%s). Cause: %v", method, url, err)
	}
	if res.StatusCode != 200 {
		return common.EncryptedPayloadHash{}, fmt.Errorf("Non-200 status code: %+v", res)
	}

	hashBytes, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, res.Body))
	if err != nil {
		return common.EncryptedPayloadHash{}, fmt.Errorf("unable to decode response body for (method:%s,path:%s). Cause: %v", method, url, err)
	}
	return common.BytesToEncryptedPayloadHash(hashBytes), nil
}

func (c *Client) ReceivePayload(key common.EncryptedPayloadHash) ([]byte, common.EncryptedPayloadHashes, common.Hash, error) {
	method := "GET"
	url := "http+unix://c/receiveraw"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("unable to build request for (method:%s,url:%s). Cause: %v", method, url, err)
	}
	req.Header.Set("c11n-key", key.ToBase64())
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("unable to submit request (method:%s,url:%s). Cause: %v", method, url, err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 { // payload not found
		return nil, nil, common.Hash{}, nil // empty payload
	}
	if res.StatusCode != 200 {
		return nil, nil, common.Hash{}, fmt.Errorf("Non-200 status code: %+v", res)
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("unable to read response body for (method:%s,path:%s). Cause: %v", method, url, err)
	}

	return payload, nil, common.Hash{}, nil
}
