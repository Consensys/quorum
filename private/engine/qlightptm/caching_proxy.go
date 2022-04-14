package qlightptm

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private/cache"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/rpc"
	gocache "github.com/patrickmn/go-cache"
)

type RPCClientCaller interface {
	Call(result interface{}, method string, args ...interface{}) error
}

type CachingProxyTxManager struct {
	features  *engine.FeatureSet
	cache     *gocache.Cache
	rpcClient RPCClientCaller
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

func (t *CachingProxyTxManager) SetRPCClientCaller(client RPCClientCaller) {
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
func (t *CachingProxyTxManager) receive(hash common.EncryptedPayloadHash, isRaw bool) (string, []string, []byte, *engine.ExtraMetadata, error) {
	if common.EmptyEncryptedPayloadHash(hash) {
		return "", nil, nil, nil, nil
	}
	cacheKey := hash.Hex()
	if isRaw {
		// indicate the cache item is incomplete, this will be fulfilled in SendSignedTx
		cacheKey = fmt.Sprintf("%s-incomplete", cacheKey)
	}
	log.Info("qlight: retrieving private data from ptm cache")
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

	log.Info("qlight: no private data in ptm cache, retrieving from qlight server node")
	var result engine.QuorumPayloadExtra
	err := t.rpcClient.Call(&result, "eth_getQuorumPayloadExtra", hash.Hex())
	if err != nil {
		return "", nil, nil, nil, err
	}
	if len(result.Payload) > 3 {
		payloadBytes, err := hex.DecodeString(result.Payload[2:])
		if err != nil {
			return "", nil, nil, nil, err
		}

		toCache := &CachablePrivateTransactionData{
			Hash:                hash,
			QuorumPrivateTxData: result,
		}
		if err := t.Cache(toCache); err != nil {
			log.Warn("unable to cache ptm data", "err", err)
		}

		return result.ExtraMetaData.Sender, result.ExtraMetaData.ManagedParties, payloadBytes, result.ExtraMetaData, nil
	}
	return "", nil, nil, nil, nil
}

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

type CachablePrivateTransactionData struct {
	Hash                common.EncryptedPayloadHash
	QuorumPrivateTxData engine.QuorumPayloadExtra
}

func (t *CachingProxyTxManager) Cache(privateTxData *CachablePrivateTransactionData) error {
	if common.EmptyEncryptedPayloadHash(privateTxData.Hash) {
		return nil
	}
	cacheKey := privateTxData.Hash.Hex()

	payload, err := hexutil.Decode(privateTxData.QuorumPrivateTxData.Payload)
	if err != nil {
		return err
	}

	t.cache.Set(cacheKey, CPItem{
		PrivateCacheItem: cache.PrivateCacheItem{
			Payload: payload,
			Extra:   *privateTxData.QuorumPrivateTxData.ExtraMetaData,
		},
		IsSender: privateTxData.QuorumPrivateTxData.IsSender,
	}, gocache.DefaultExpiration)

	return nil
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
