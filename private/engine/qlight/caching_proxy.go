package qlight

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/cache"
	"github.com/ethereum/go-ethereum/private/engine"
	gocache "github.com/patrickmn/go-cache"
)

type CachingProxyTxManager struct {
	features *engine.FeatureSet
	cache    *gocache.Cache
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
		cacheItem, ok := item.(cache.PrivateCacheItem)
		if !ok {
			return "", nil, nil, nil, fmt.Errorf("unknown cache item. expected type PrivateCacheItem")
		}
		return cacheItem.Extra.Sender, cacheItem.Extra.ManagedParties, cacheItem.Payload, &cacheItem.Extra, nil
	}
	return "", nil, nil, nil, nil
}

// retrieve raw will not return information about medata
func (t *CachingProxyTxManager) AddToCache(hash common.EncryptedPayloadHash, payload []byte, extra *engine.ExtraMetadata) {
	if common.EmptyEncryptedPayloadHash(hash) {
		return
	}
	cacheKey := hash.Hex()

	t.cache.Set(cacheKey, cache.PrivateCacheItem{
		Payload: payload,
		Extra:   *extra,
	}, gocache.DefaultExpiration)
}

// retrieve raw will not return information about medata
func (t *CachingProxyTxManager) DecryptPayload(payload common.DecryptRequest) ([]byte, *engine.ExtraMetadata, error) {
	panic("implement me")
}

func (t *CachingProxyTxManager) IsSender(txHash common.EncryptedPayloadHash) (bool, error) {
	panic("implement me")
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
