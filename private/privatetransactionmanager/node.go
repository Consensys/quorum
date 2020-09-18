package privatetransactionmanager

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/tv42/httpunix"
)

type storeRawReq struct {
	Payload string `json:"payload"`
	From    string `json:"from,omitempty"`
}

type storeRawResp struct {
	Key string `json:"key"`
}

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

func unixTransport(socketPath string, readTimeout uint) *httpunix.Transport {
	if readTimeout == 0 {
		readTimeout = 5
	}

	t := &httpunix.Transport{
		DialTimeout:           1 * time.Second,
		RequestTimeout:        5 * time.Second,
		ResponseHeaderTimeout: time.Duration(readTimeout) * time.Second,
	}
	t.RegisterLocation("c", socketPath)
	return t
}

func unixClient(socketPath string, readTimeout uint) *http.Client {
	return &http.Client{
		Transport: unixTransport(socketPath, readTimeout),
	}
}

func RunNode(socketPath string, readTimeout uint) error {
	c := unixClient(socketPath, readTimeout)
	res, err := c.Get("http+unix://c/upcheck")
	if err != nil {
		return err
	}
	if res.StatusCode == 200 {
		return nil
	}
	return errors.New("private transaction manager did not respond to upcheck request")
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

func (c *Client) SendPayload(pl []byte, b64From string, b64To []string) ([]byte, error) {
	buf := bytes.NewBuffer(pl)
	req, err := http.NewRequest("POST", "http+unix://c/sendraw", buf)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Non-200 status code: %+v", res)
	}

	return ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, res.Body))
}

func (c *Client) StorePayload(pl []byte, b64From string) ([]byte, error) {
	storeRawReq := &storeRawReq{
		Payload: base64.StdEncoding.EncodeToString(pl),
		From:    b64From,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(storeRawReq); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", "http+unix://c/storeraw", buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.httpClient.Do(req)

	if res.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("Non-200 status code, verify that tessera is running and version is 0.10.5+: %v", res)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Non-200 status code: %+v", res)
	}
	// parse response
	var storeRawResp storeRawResp
	if err := json.NewDecoder(res.Body).Decode(&storeRawResp); err != nil {
		return nil, err
	}
	encryptedPayloadHash, err := base64.StdEncoding.DecodeString(storeRawResp.Key)
	if err != nil {
		return nil, err
	}
	return encryptedPayloadHash, nil
}

func (c *Client) SendSignedPayload(signedPayload []byte, b64To []string) ([]byte, error) {
	buf := bytes.NewBuffer(signedPayload)
	req, err := http.NewRequest("POST", "http+unix://c/sendsignedtx", buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("c11n-to", strings.Join(b64To, ","))
	req.Header.Set("Content-Type", "application/octet-stream")
	res, err := c.httpClient.Do(req)

	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Non-200 status code: %+v", res)
	}

	return ioutil.ReadAll(base64.NewDecoder(base64.StdEncoding, res.Body))
}

func (c *Client) ReceivePayload(key []byte) ([]byte, error) {
	req, err := http.NewRequest("GET", "http+unix://c/receiveraw", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("c11n-key", base64.StdEncoding.EncodeToString(key))
	res, err := c.httpClient.Do(req)

	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Non-200 status code: %+v", res)
	}

	return ioutil.ReadAll(res.Body)
}

func (c *Client) IsSender(txHash common.EncryptedPayloadHash) (bool, error) {
	req, err := http.NewRequest("GET", "http+unix://c/transaction/"+url.PathEscape(txHash.ToBase64())+"/isSender", nil)
	if err != nil {
		return false, err
	}

	res, err := c.httpClient.Do(req)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return false, err
	}

	if res.StatusCode != 200 {
		return false, fmt.Errorf("non-200 status code: %+v", res)
	}

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(out))
}

func (c *Client) GetParticipants(txHash common.EncryptedPayloadHash) ([]string, error) {
	requestUrl := "http+unix://c/transaction/" + url.PathEscape(txHash.ToBase64()) + "/participants"
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Non-200 status code: %+v", res)
	}

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	split := strings.Split(string(out), ",")

	return split, nil
}

func NewClient(socketPath string, readTimeout uint) (*Client, error) {
	return &Client{
		httpClient: unixClient(socketPath, readTimeout),
	}, nil
}
