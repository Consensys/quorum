package constellation

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/patrickmn/go-cache"
)

type Constellation struct {
	node                    *Client
	c                       *cache.Cache
	isConstellationNotInUse bool
}

var (
	ErrConstellationIsntInit = errors.New("Constellation not in use")
)

type PrivateCacheItem struct {
	payload      []byte
	acHashes     common.EncryptedPayloadHashes // hashes of affected contracts
	acMerkleRoot common.Hash                   // merkle root of all affected contracts
}

func (g *Constellation) Send(data []byte, from string, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) (out common.EncryptedPayloadHash, err error) {
	if g.isConstellationNotInUse {
		return common.EncryptedPayloadHash{}, ErrConstellationIsntInit
	}
	out, err = g.node.SendPayload(data, from, to, acHashes, acMerkleRoot)
	if err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	cacheKey := string(out.Bytes())
	g.c.Set(cacheKey, PrivateCacheItem{
		payload:      data,
		acHashes:     acHashes,
		acMerkleRoot: acMerkleRoot,
	}, cache.DefaultExpiration)
	return out, nil
}

func (g *Constellation) SendSignedTx(data common.EncryptedPayloadHash, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) (out []byte, err error) {
	if g.isConstellationNotInUse {
		return nil, ErrConstellationIsntInit
	}
	out, err = g.node.SendSignedPayload(data, to, acHashes, acMerkleRoot)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (g *Constellation) Receive(data common.EncryptedPayloadHash) ([]byte, common.EncryptedPayloadHashes, common.Hash, error) {
	if g.isConstellationNotInUse {
		return nil, nil, common.Hash{}, nil
	}
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
		cacheItem, ok := x.(PrivateCacheItem)
		if !ok {
			return nil, nil, common.Hash{}, fmt.Errorf("unknown cache item. expected type PrivateCacheItem")
		}
		return cacheItem.payload, cacheItem.acHashes, cacheItem.acMerkleRoot, nil
	}
	privatePayload, acHashes, acMerkleRoot, err := g.node.ReceivePayload(data)
	if nil != err {
		return nil, nil, common.Hash{}, err
	}
	g.c.Set(cacheKey, PrivateCacheItem{
		payload:      privatePayload,
		acHashes:     acHashes,
		acMerkleRoot: acMerkleRoot,
	}, cache.DefaultExpiration)
	return privatePayload, acHashes, acMerkleRoot, nil
}

func New(path string) (*Constellation, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	// We accept either the socket or a configuration file that points to
	// a socket.
	isSocket := info.Mode()&os.ModeSocket != 0
	if !isSocket {
		cfg, err := LoadConfig(path)
		if err != nil {
			return nil, err
		}
		path = filepath.Join(cfg.WorkDir, cfg.Socket)
	}
	err = RunNode(path)
	if err != nil {
		return nil, err
	}
	n, err := NewClient(path)
	if err != nil {
		return nil, err
	}
	return &Constellation{
		node:                    n,
		c:                       cache.New(5*time.Minute, 5*time.Minute),
		isConstellationNotInUse: false,
	}, nil
}

func MustNew(path string) *Constellation {
	if strings.EqualFold(path, "ignore") {
		return &Constellation{
			node:                    nil,
			c:                       nil,
			isConstellationNotInUse: true,
		}
	}
	g, err := New(path)
	if err != nil {
		panic(fmt.Sprintf("MustNew: Failed to connect to Constellation (%s): %v", path, err))
	}
	return g
}

func MaybeNew(path string) *Constellation {
	if path == "" {
		return nil
	}
	return MustNew(path)
}
