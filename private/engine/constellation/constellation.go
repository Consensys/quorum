package constellation

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/cache"

	gocache "github.com/patrickmn/go-cache"
)

type constellation struct {
	node *Client
	c    *gocache.Cache
}

func New(client *http.Client) *constellation {
	return &constellation{
		node: &Client{
			httpClient: client,
		},
		c: gocache.New(cache.DefaultExpiration, cache.CleanupInterval),
	}
}

func (g *constellation) Send(data []byte, from string, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) (out common.EncryptedPayloadHash, err error) {
	out, err = g.node.SendPayload(data, from, to, acHashes, acMerkleRoot)
	if err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	cacheKey := string(out.Bytes())
	g.c.Set(cacheKey, cache.PrivateCacheItem{
		Payload:      data,
		ACHashes:     acHashes,
		ACMerkleRoot: acMerkleRoot,
	}, cache.DefaultExpiration)
	return out, nil
}

func (g *constellation) SendSignedTx(data common.EncryptedPayloadHash, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) (out []byte, err error) {
	out, err = g.node.SendSignedPayload(data, to, acHashes, acMerkleRoot)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (g *constellation) Receive(data common.EncryptedPayloadHash) ([]byte, common.EncryptedPayloadHashes, common.Hash, error) {
	if common.EmptyEncryptedPayloadHash(data) {
		return data.Bytes(), nil, common.Hash{}, nil
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
			return nil, nil, common.Hash{}, fmt.Errorf("unknown cache item. expected type PrivateCacheItem")
		}
		return cacheItem.Payload, cacheItem.ACHashes, cacheItem.ACMerkleRoot, nil
	}
	privatePayload, acHashes, acMerkleRoot, err := g.node.ReceivePayload(data)
	if nil != err {
		return nil, nil, common.Hash{}, err
	}
	g.c.Set(cacheKey, cache.PrivateCacheItem{
		Payload:      privatePayload,
		ACHashes:     acHashes,
		ACMerkleRoot: acMerkleRoot,
	}, cache.DefaultExpiration)
	return privatePayload, acHashes, acMerkleRoot, nil
}

func (g *constellation) Name() string {
	return "constellation"
}
