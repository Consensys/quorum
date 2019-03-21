package tessera

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

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

func (t *tesseraPrivateTxManager) Send(data []byte, from string, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) (common.EncryptedPayloadHash, error) {
	req, err := newJSONRequest("POST", "/send", &sendRequest{
		Payload:                      data,
		From:                         from,
		To:                           to,
		AffectedContractTransactions: acHashes.ToBase64s(),
		ExecHash:                     base64.StdEncoding.EncodeToString(acMerkleRoot.Bytes()),
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
		Payload:      data,
		ACHashes:     acHashes,
		ACMerkleRoot: acMerkleRoot,
	}, gocache.DefaultExpiration)

	return eph, nil
}

func (t *tesseraPrivateTxManager) SendSignedTx(data common.EncryptedPayloadHash, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) ([]byte, error) {
	req, err := newJSONRequest("POST", "/sendsignedtx", &sendSignedTxRequest{
		Hash:                         data.Bytes(),
		To:                           to,
		AffectedContractTransactions: acHashes.ToBase64s(),
		ExecHash:                     base64.StdEncoding.EncodeToString(acMerkleRoot.Bytes()),
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

func (t *tesseraPrivateTxManager) Receive(data common.EncryptedPayloadHash) ([]byte, common.EncryptedPayloadHashes, common.Hash, error) {
	if common.EmptyEncryptedPayloadHash(data) {
		return data.Bytes(), nil, common.Hash{}, nil
	}
	cacheKey := string(data.Bytes())
	if item, found := t.cache.Get(cacheKey); found {
		cacheItem, ok := item.(cache.PrivateCacheItem)
		if !ok {
			return nil, nil, common.Hash{}, fmt.Errorf("unknown cache item. expected type PrivateCacheItem")
		}
		return cacheItem.Payload, cacheItem.ACHashes, cacheItem.ACMerkleRoot, nil
	}
	req, err := newJSONRequest("GET", "/receive", &receiveRequest{
		Key: data.ToBase64(),
	})
	if err != nil {
		return nil, nil, common.Hash{}, err
	}

	res, err := t.client.Do(req)
	if err != nil {
		return nil, nil, common.Hash{}, err
	}
	defer res.Body.Close()
	receiveRes := new(receiveResponse)
	if err := json.NewDecoder(res.Body).Decode(receiveRes); err != nil {
		return nil, nil, common.Hash{}, err
	}
	acHashes, err := common.Base64sToEncryptedPayloadHashes(receiveRes.AffectedContractTransactions)
	if err != nil {
		return nil, nil, common.Hash{}, err
	}
	acMerkleRoot := common.StringToHash(receiveRes.ExecHash)

	t.cache.Set(cacheKey, cache.PrivateCacheItem{
		Payload:      receiveRes.Payload,
		ACHashes:     acHashes,
		ACMerkleRoot: acMerkleRoot,
	}, gocache.DefaultExpiration)

	return receiveRes.Payload, acHashes, acMerkleRoot, nil
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
