package tessera

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/private/engine"

	"github.com/ethereum/go-ethereum/private/cache"

	"github.com/ethereum/go-ethereum/params"

	gocache "github.com/patrickmn/go-cache"

	"github.com/ethereum/go-ethereum/common"
)

type tesseraPrivateTxManager struct {
	client *engine.Client
	cache  *gocache.Cache
}

func New(client *engine.Client) *tesseraPrivateTxManager {
	return &tesseraPrivateTxManager{
		client: client,
		cache:  gocache.New(cache.DefaultExpiration, cache.CleanupInterval),
	}
}

func (t *tesseraPrivateTxManager) submitJSON(method, path string, request interface{}, response interface{}) error {
	req, err := newJSONRequest(method, t.client.FullPath(path), request)
	if err != nil {
		return err
	}
	res, err := t.client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("%d status: %s", res.StatusCode, string(body))
	}
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return err
	}
	return nil
}

func (t *tesseraPrivateTxManager) Send(data []byte, from string, to []string, extra *engine.ExtraMetadata) (common.EncryptedPayloadHash, error) {
	response := new(sendResponse)
	if err := t.submitJSON("POST", "/send", &sendRequest{
		Payload:                      data,
		From:                         from,
		To:                           to,
		AffectedContractTransactions: extra.ACHashes.ToBase64s(),
		ExecHash:                     base64.StdEncoding.EncodeToString(extra.ACMerkleRoot.Bytes()),
		PrivateStateValidation:       extra.PrivateStateValidation,
	}, response); err != nil {
		return common.EncryptedPayloadHash{}, err
	}

	hashBytes, err := base64.StdEncoding.DecodeString(response.Key)
	if err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	eph := common.BytesToEncryptedPayloadHash(hashBytes)

	cacheKey := string(eph.Bytes())
	t.cache.Set(cacheKey, cache.PrivateCacheItem{
		Payload: data,
		Extra:   *extra,
	}, gocache.DefaultExpiration)

	return eph, nil
}

func (t *tesseraPrivateTxManager) SendSignedTx(data common.EncryptedPayloadHash, to []string, extra *engine.ExtraMetadata) ([]byte, error) {
	response := new(sendSignedTxResponse)
	if err := t.submitJSON("POST", "/sendsignedtx", &sendSignedTxRequest{
		Hash:                         data.Bytes(),
		To:                           to,
		AffectedContractTransactions: extra.ACHashes.ToBase64s(),
		ExecHash:                     base64.StdEncoding.EncodeToString(extra.ACMerkleRoot.Bytes()),
		PrivateStateValidation:       extra.PrivateStateValidation,
	}, response); err != nil {
		return nil, err
	}

	hashBytes, err := base64.StdEncoding.DecodeString(response.Key)
	if err != nil {
		return nil, err
	}
	return hashBytes, err
}

func (t *tesseraPrivateTxManager) Receive(data common.EncryptedPayloadHash) ([]byte, *engine.ExtraMetadata, error) {
	if common.EmptyEncryptedPayloadHash(data) {
		return data.Bytes(), nil, nil
	}
	cacheKey := string(data.Bytes())
	if item, found := t.cache.Get(cacheKey); found {
		cacheItem, ok := item.(cache.PrivateCacheItem)
		if !ok {
			return nil, nil, fmt.Errorf("unknown cache item. expected type PrivateCacheItem")
		}
		return cacheItem.Payload, &cacheItem.Extra, nil
	}

	response := new(receiveResponse)
	if err := t.submitJSON("GET", "/receive", &receiveRequest{
		Key: data.ToBase64(),
	}, response); err != nil {
		return nil, nil, err
	}

	acHashes, err := common.Base64sToEncryptedPayloadHashes(response.AffectedContractTransactions)
	if err != nil {
		return nil, nil, err
	}
	acMerkleRootInBytes, err := base64.StdEncoding.DecodeString(response.ExecHash)
	if err != nil {
		return nil, nil, err
	}
	extra := &engine.ExtraMetadata{
		ACHashes:               acHashes,
		ACMerkleRoot:           common.BytesToHash(acMerkleRootInBytes),
		PrivateStateValidation: response.PrivateStateValidation,
	}

	t.cache.Set(cacheKey, cache.PrivateCacheItem{
		Payload: response.Payload,
		Extra:   *extra,
	}, gocache.DefaultExpiration)

	return response.Payload, extra, nil
}

func (t *tesseraPrivateTxManager) Name() string {
	return "Tessera"
}

func newJSONRequest(method string, path string, body interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, path, buf)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", fmt.Sprintf("quorum-v%s", params.QuorumVersion))
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Accept", "application/json")
	return request, nil
}
