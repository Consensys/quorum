package tessera

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/private/engine"

	"github.com/ethereum/go-ethereum/private/cache"

	"github.com/ethereum/go-ethereum/params"

	gocache "github.com/patrickmn/go-cache"

	"github.com/ethereum/go-ethereum/common"
)

type tesseraPrivateTxManager struct {
	client *http.Client
	cache  *gocache.Cache
}

func New(client *http.Client) *tesseraPrivateTxManager {
	return &tesseraPrivateTxManager{
		client: client,
		cache:  gocache.New(cache.DefaultExpiration, cache.CleanupInterval),
	}
}

func (t *tesseraPrivateTxManager) Send(data []byte, from string, to []string, extra *engine.ExtraMetadata) (common.EncryptedPayloadHash, error) {
	req, err := newJSONRequest("POST", "/send", &sendRequest{
		Payload:                      data,
		From:                         from,
		To:                           to,
		AffectedContractTransactions: extra.ACHashes.ToBase64s(),
		ExecHash:                     base64.StdEncoding.EncodeToString(extra.ACMerkleRoot.Bytes()),
		PrivateStateValidation:       extra.PrivateStateValidation,
	})
	if err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	res, err := t.client.Do(req)
	if err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	defer res.Body.Close()

	sendRes := new(sendResponse)
	if err := json.NewDecoder(res.Body).Decode(sendRes); err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	hashBytes, err := base64.StdEncoding.DecodeString(sendRes.Key)
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
	req, err := newJSONRequest("POST", "/sendsignedtx", &sendSignedTxRequest{
		Hash:                         data.Bytes(),
		To:                           to,
		AffectedContractTransactions: extra.ACHashes.ToBase64s(),
		ExecHash:                     base64.StdEncoding.EncodeToString(extra.ACMerkleRoot.Bytes()),
		PrivateStateValidation:       extra.PrivateStateValidation,
	})
	if err != nil {
		return nil, err
	}
	res, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	sendSignedTxRes := new(sendSignedTxResponse)
	if err := json.NewDecoder(res.Body).Decode(sendSignedTxRes); err != nil {
		return nil, err
	}
	hashBytes, err := base64.StdEncoding.DecodeString(sendSignedTxRes.Key)
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
	req, err := newJSONRequest("GET", "/receive", &receiveRequest{
		Key: data.ToBase64(),
	})
	if err != nil {
		return nil, nil, err
	}

	res, err := t.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	receiveRes := new(receiveResponse)
	if err := json.NewDecoder(res.Body).Decode(receiveRes); err != nil {
		return nil, nil, err
	}
	acHashes, err := common.Base64sToEncryptedPayloadHashes(receiveRes.AffectedContractTransactions)
	if err != nil {
		return nil, nil, err
	}
	extra := &engine.ExtraMetadata{
		ACHashes:               acHashes,
		ACMerkleRoot:           common.StringToHash(receiveRes.ExecHash),
		PrivateStateValidation: receiveRes.PrivateStateValidation,
	}

	t.cache.Set(cacheKey, cache.PrivateCacheItem{
		Payload: receiveRes.Payload,
		Extra:   *extra,
	}, gocache.DefaultExpiration)

	return receiveRes.Payload, extra, nil
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
	return request, nil
}
