package qlightptm

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/cache"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/rpc"
	gocache "github.com/patrickmn/go-cache"
)

type CachingProxyTxManager struct {
	features  *engine.FeatureSet
	cache     *gocache.Cache
	rpcClient *rpc.Client
}

type CPItem struct {
	cache.PrivateCacheItem
	IsSender bool
	IsEmpty  bool
}

func Is(ptm interface{}) bool {
	_, ok := ptm.(*CachingProxyTxManager)
	return ok
}

func New() *CachingProxyTxManager {
	return &CachingProxyTxManager{
		features: engine.NewFeatureSet(engine.PrivacyEnhancements),
		cache:    gocache.New(cache.DefaultExpiration, cache.CleanupInterval),
	}
}

func (t *CachingProxyTxManager) SetRPCClient(client *rpc.Client) {
	t.rpcClient = client
}

func (t *CachingProxyTxManager) Send(data []byte, from string, to []string, extra *engine.ExtraMetadata) (string, []string, common.EncryptedPayloadHash, error) {
	panic("implement me")
}

func (t *CachingProxyTxManager) EncryptPayload(data []byte, from string, to []string, extra *engine.ExtraMetadata) ([]byte, error) {
	panic("implement me")
}

func (t *CachingProxyTxManager) StoreRaw(data []byte, from string) (common.EncryptedPayloadHash, error) {
	panic("implement me")
}

func (t *CachingProxyTxManager) SendSignedTx(data common.EncryptedPayloadHash, to []string, extra *engine.ExtraMetadata) (string, []string, []byte, error) {
	panic("implement me")
}

func (t *CachingProxyTxManager) Receive(hash common.EncryptedPayloadHash) (string, []string, []byte, *engine.ExtraMetadata, error) {
	return t.receive(hash, false)
}

// retrieve raw will not return information about medata.
// Related to SendSignedTx
func (t *CachingProxyTxManager) ReceiveRaw(hash common.EncryptedPayloadHash) ([]byte, string, *engine.ExtraMetadata, error) {
	sender, _, data, extra, err := t.receive(hash, true)
	return data, sender, extra, err
}

// retrieve raw will not return information about medata
func (t *CachingProxyTxManager) receive(data common.EncryptedPayloadHash, isRaw bool) (string, []string, []byte, *engine.ExtraMetadata, error) {
	if common.EmptyEncryptedPayloadHash(data) {
		return "", nil, nil, nil, nil
	}
	cacheKey := data.Hex()
	if isRaw {
		// indicate the cache item is incomplete, this will be fulfilled in SendSignedTx
		cacheKey = fmt.Sprintf("%s-incomplete", cacheKey)
	}
	if item, found := t.cache.Get(cacheKey); found {
		cacheItem, ok := item.(CPItem)
		if !ok {
			return "", nil, nil, nil, fmt.Errorf("unknown cache item. expected type PrivateCacheItem")
		}
		if cacheItem.IsEmpty {
			return "", nil, nil, nil, nil
		}
		return cacheItem.Extra.Sender, cacheItem.Extra.ManagedParties, cacheItem.Payload, &cacheItem.Extra, nil
	}

	var result engine.QuorumPayloadExtra
	err := t.rpcClient.Call(&result, "eth_getQuorumPayloadExtra", data.Hex())
	if err != nil {
		return "", nil, nil, nil, err
	}
	if len(result.Payload) > 3 {
		payloadBytes, err := hex.DecodeString(result.Payload[2:])
		if err != nil {
			return "", nil, nil, nil, err
		}
		t.AddToCache(data, payloadBytes, result.ExtraMetaData, result.IsSender)
		return result.ExtraMetaData.Sender, result.ExtraMetaData.ManagedParties, payloadBytes, result.ExtraMetaData, nil
	}
	return "", nil, nil, nil, nil
}

// retrieve raw will not return information about medata
func (t *CachingProxyTxManager) CheckAndAddEmptyToCache(hash common.EncryptedPayloadHash) {
	if common.EmptyEncryptedPayloadHash(hash) {
		return
	}
	cacheKey := hash.Hex()

	if _, found := t.cache.Get(cacheKey); found {
		return
	}

	t.cache.Set(cacheKey, CPItem{
		IsEmpty: true,
	}, gocache.DefaultExpiration)
}

func (t *CachingProxyTxManager) AddToCache(hash common.EncryptedPayloadHash, payload []byte, extra *engine.ExtraMetadata, isSender bool) {
	if common.EmptyEncryptedPayloadHash(hash) {
		return
	}
	cacheKey := hash.Hex()

	t.cache.Set(cacheKey, CPItem{
		PrivateCacheItem: cache.PrivateCacheItem{
			Payload: payload,
			Extra:   *extra,
		},
		IsSender: isSender,
	}, gocache.DefaultExpiration)
}

// retrieve raw will not return information about medata
func (t *CachingProxyTxManager) DecryptPayload(payload common.DecryptRequest) ([]byte, *engine.ExtraMetadata, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}
	payloadHex := fmt.Sprintf("0x%x", payloadBytes)

	var result engine.QuorumPayloadExtra
	err = t.rpcClient.Call(&result, "eth_decryptQuorumPayload", payloadHex)
	if err != nil {
		return nil, nil, err
	}

	responsePayloadHex := result.Payload
	if len(responsePayloadHex) < 3 {
		return nil, nil, fmt.Errorf("Invalid payload hex")
	}
	if responsePayloadHex[:2] == "0x" {
		responsePayloadHex = responsePayloadHex[2:]
	}
	responsePayload, err := hex.DecodeString(responsePayloadHex)
	if err != nil {
		return nil, nil, err
	}
	return responsePayload, result.ExtraMetaData, nil
}

func (t *CachingProxyTxManager) IsSender(data common.EncryptedPayloadHash) (bool, error) {
	_, _, _, _, err := t.receive(data, false)
	if err != nil {
		return false, err
	}
	cacheKey := data.Hex()
	if item, found := t.cache.Get(cacheKey); found {
		cacheItem, ok := item.(CPItem)
		if !ok {
			return false, fmt.Errorf("unknown cache item. expected type PrivateCacheItem")
		}
		return cacheItem.IsSender, nil
	}
	return false, nil
}

func (t *CachingProxyTxManager) GetParticipants(txHash common.EncryptedPayloadHash) ([]string, error) {
	panic("implement me")
}

func (t *CachingProxyTxManager) GetMandatory(txHash common.EncryptedPayloadHash) ([]string, error) {
	panic("implement me")
}

func (t *CachingProxyTxManager) Groups() ([]engine.PrivacyGroup, error) {
	panic("implement me")
}

func (t *CachingProxyTxManager) Name() string {
	return "CachingP2PProxy"
}

func (t *CachingProxyTxManager) HasFeature(f engine.PrivateTransactionManagerFeature) bool {
	return t.features.HasFeature(f)
}

func (t *CachingProxyTxManager) AddPrivateBlockToCache(key string) error {
	var result engine.BlockPrivatePayloads
	err := t.rpcClient.Call(&result, "eth_getQuorumPayloadsForBlock", key)
	if err != nil {
		return err
	}
	for txKey, qpe := range result.Payloads {
		eph, err := common.Base64ToEncryptedPayloadHash(txKey)
		if err != nil {
			return err
		}
		if len(qpe.Payload) > 3 {
			payloadBytes, err := hex.DecodeString(qpe.Payload[2:])
			if err != nil {
				return err
			}
			t.AddToCache(eph, payloadBytes, qpe.ExtraMetaData, qpe.IsSender)
		}
	}
	return nil
}
