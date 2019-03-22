package constellation

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func launchNode(cfgPath string) (*exec.Cmd, error) {
	cmd := exec.Command("constellation-node", cfgPath)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	go io.Copy(os.Stderr, stderr)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	time.Sleep(100 * time.Millisecond)
	return cmd, nil
}

type Client struct {
	httpClient *http.Client
}

func (c *Client) doJson(path string, apiReq interface{}) (*http.Response, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(apiReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "http+unix://c/"+path, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.httpClient.Do(req)
	if err == nil && res.StatusCode != 200 {
		return nil, fmt.Errorf("Non-200 status code: %+v", res)
	}
	return res, err
}

func (c *Client) SendPayload(pl []byte, b64From string, b64To []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) (common.EncryptedPayloadHash, error) {
	buf := bytes.NewBuffer(pl)
	req, err := http.NewRequest("POST", "http+unix://c/sendraw", buf)
	if err != nil {
		return common.EncryptedPayloadHash{}, err
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
		return common.EncryptedPayloadHash{}, err
	}
	if res.StatusCode != 200 {
		return common.EncryptedPayloadHash{}, fmt.Errorf("Non-200 status code: %+v", res)
	}

	hashBytes, err := ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, res.Body))
	if err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	return common.BytesToEncryptedPayloadHash(hashBytes), nil
}

func (c *Client) ReceivePayload(key common.EncryptedPayloadHash) ([]byte, common.EncryptedPayloadHashes, common.Hash, error) {
	req, err := http.NewRequest("GET", "http+unix://c/receiveraw", nil)
	if err != nil {
		return nil, nil, common.Hash{}, err
	}
	req.Header.Set("c11n-key", key.ToBase64())
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, common.Hash{}, err
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
		return nil, nil, common.Hash{}, err
	}

	return payload, nil, common.Hash{}, nil
}
