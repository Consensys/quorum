package privatetransactionmanager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/cache"
	gocache "github.com/patrickmn/go-cache"
)

type PrivateTransactionManager struct {
	node *Client
	c    *gocache.Cache
}

func (g *PrivateTransactionManager) Send(data []byte, from string, to []string) (out common.EncryptedPayloadHash, err error) {
	var b []byte
	b, err = g.node.SendPayload(data, from, to)
	if err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	g.c.Set(string(b), data, cache.DefaultExpiration)
	out = common.BytesToEncryptedPayloadHash(b)
	return
}

func (g *PrivateTransactionManager) StoreRaw(data []byte, from string) (out common.EncryptedPayloadHash, err error) {
	var b []byte
	b, err = g.node.StorePayload(data, from)
	if err != nil {
		return common.EncryptedPayloadHash{}, err
	}
	g.c.Set(string(b), data, cache.DefaultExpiration)
	out = common.BytesToEncryptedPayloadHash(b)
	return out, nil
}

func (g *PrivateTransactionManager) SendSignedTx(txHash common.EncryptedPayloadHash, to []string) (out []byte, err error) {
	out, err = g.node.SendSignedPayload(txHash.Bytes(), to)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (g *PrivateTransactionManager) Receive(txHash common.EncryptedPayloadHash) ([]byte, error) {
	if common.EmptyEncryptedPayloadHash(txHash) {
		return []byte{}, nil
	}
	// Ignore this error since not being a recipient of
	// a payload isn't an error.
	// TODO: Return an error if it's anything OTHER than
	// 'you are not a recipient.'
	dataStr := string(txHash.Bytes())
	x, found := g.c.Get(dataStr)
	if found {
		return x.([]byte), nil
	}
	pl, _ := g.node.ReceivePayload(txHash.Bytes())
	g.c.Set(dataStr, pl, cache.DefaultExpiration)
	return pl, nil
}

func (g *PrivateTransactionManager) IsSender(txHash common.EncryptedPayloadHash) (bool, error) {
	return g.node.IsSender(txHash)
}

func (g *PrivateTransactionManager) GetParticipants(txHash common.EncryptedPayloadHash) ([]string, error) {

	return g.node.GetParticipants(txHash)
}

func New(path string) (*PrivateTransactionManager, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	// We accept either the socket or a configuration file that points to
	// a socket.
	cfg := new(Config)
	isSocket := info.Mode()&os.ModeSocket != 0
	if !isSocket {
		cfg, err = LoadConfig(path)
		if err != nil {
			return nil, err
		}
	} else {
		cfg.WorkDir, cfg.Socket = filepath.Split(path)
	}
	err = RunNode(*cfg)
	if err != nil {
		return nil, err
	}
	n, err := NewClient(*cfg)
	if err != nil {
		return nil, err
	}
	return &PrivateTransactionManager{
		node: n,
		c:    cache.NewDefaultCache(),
	}, nil
}

func MustNew(path string) *PrivateTransactionManager {
	g, err := New(path)
	if err != nil {
		panic(fmt.Sprintf("MustNew: Failed to connect to private transaction manager (%s): %v", path, err))
	}
	return g
}
