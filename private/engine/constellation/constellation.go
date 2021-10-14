package constellation

import (
	"fmt"

	"github.com/ethereum/go-ethereum/private/engine"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/cache"

	gocache "github.com/patrickmn/go-cache"
)

type constellation struct {
	node *Client
	c    *gocache.Cache
}

func Is(ptm interface{}) bool {
	_, ok := ptm.(*constellation)
	return ok
}

func New(client *engine.Client) *constellation {
	return &constellation{
		node: &Client{
			httpClient: client.HttpClient,
		},
		c: gocache.New(cache.DefaultExpiration, cache.CleanupInterval),
	}
}

func (g *constellation) Send(data []byte, from string, to []string, extra *engine.ExtraMetadata) (string, []string, common.EncryptedPayloadHash, error) {
	if extra.PrivacyFlag.IsNotStandardPrivate() {
		return "", nil, common.EncryptedPayloadHash{}, engine.ErrPrivateTxManagerDoesNotSupportPrivacyEnhancements
	}
	out, err := g.node.SendPayload(data, from, to, extra.ACHashes, extra.ACMerkleRoot)
	if err != nil {
		return "", nil, common.EncryptedPayloadHash{}, err
	}
	cacheKey := string(out.Bytes())
	g.c.Set(cacheKey, cache.PrivateCacheItem{
		Payload: data,
		Extra:   *extra,
	}, cache.DefaultExpiration)
	return "", nil, out, nil
}

func (g *constellation) EncryptPayload(data []byte, from string, to []string, extra *engine.ExtraMetadata) ([]byte, error) {
	return nil, engine.ErrPrivateTxManagerNotSupported
}

func (g *constellation) DecryptPayload(payload common.DecryptRequest) ([]byte, *engine.ExtraMetadata, error) {
	return nil, nil, engine.ErrPrivateTxManagerNotSupported
}

func (g *constellation) StoreRaw(data []byte, from string) (common.EncryptedPayloadHash, error) {
	return common.EncryptedPayloadHash{}, engine.ErrPrivateTxManagerNotSupported
}

func (g *constellation) SendSignedTx(data common.EncryptedPayloadHash, to []string, extra *engine.ExtraMetadata) (string, []string, []byte, error) {
	return "", nil, nil, engine.ErrPrivateTxManagerNotSupported
}

func (g *constellation) ReceiveRaw(data common.EncryptedPayloadHash) ([]byte, string, *engine.ExtraMetadata, error) {
	return nil, "", nil, engine.ErrPrivateTxManagerNotSupported
}

func (g *constellation) IsSender(txHash common.EncryptedPayloadHash) (bool, error) {
	return false, engine.ErrPrivateTxManagerNotSupported
}

func (ptm *constellation) Groups() ([]engine.PrivacyGroup, error) {
	return nil, engine.ErrPrivateTxManagerNotSupported
}

func (g *constellation) GetParticipants(txHash common.EncryptedPayloadHash) ([]string, error) {
	return nil, engine.ErrPrivateTxManagerNotSupported
}

func (g *constellation) GetMandatory(txHash common.EncryptedPayloadHash) ([]string, error) {
	return nil, engine.ErrPrivateTxManagerNotSupported
}

func (g *constellation) Receive(data common.EncryptedPayloadHash) (string, []string, []byte, *engine.ExtraMetadata, error) {
	if common.EmptyEncryptedPayloadHash(data) {
		return "", nil, nil, nil, nil
	}
	// Ignore this error since not being a recipient of
	// a payload isn't an error.
	// TODO: Return an error if it's anything OTHER than
	// 'you are not a recipient.'
	cacheKey := string(data.Bytes())
	x, found := g.c.Get(cacheKey)
	if found {
		cacheItem, ok := x.(cache.PrivateCacheItem)
		if !ok {
			return "", nil, nil, nil, fmt.Errorf("unknown cache item. expected type PrivateCacheItem")
		}
		return "", nil, cacheItem.Payload, &cacheItem.Extra, nil
	}
	privatePayload, acHashes, acMerkleRoot, err := g.node.ReceivePayload(data)
	if nil != err {
		return "", nil, nil, nil, err
	}
	extra := engine.ExtraMetadata{
		ACHashes:     acHashes,
		ACMerkleRoot: acMerkleRoot,
	}
	g.c.Set(cacheKey, cache.PrivateCacheItem{
		Payload: privatePayload,
		Extra:   extra,
	}, cache.DefaultExpiration)
	return "", nil, privatePayload, &extra, nil
}

func (g *constellation) Name() string {
	return "Constellation"
}

func (g *constellation) HasFeature(f engine.PrivateTransactionManagerFeature) bool {
	return false
}
