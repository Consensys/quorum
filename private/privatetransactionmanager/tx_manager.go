package privatetransactionmanager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/private/cache"
	gocache "github.com/patrickmn/go-cache"
)

type PrivateTransactionManager struct {
	node *Client
	c    *gocache.Cache
}

func (g *PrivateTransactionManager) Send(data []byte, from string, to []string) (out []byte, err error) {
	out, err = g.node.SendPayload(data, from, to)
	if err != nil {
		return nil, err
	}
	g.c.Set(string(out), data, cache.DefaultExpiration)
	return out, nil
}

func (g *PrivateTransactionManager) StoreRaw(data []byte, from string) (out []byte, err error) {
	out, err = g.node.StorePayload(data, from)
	if err != nil {
		return nil, err
	}
	g.c.Set(string(out), data, cache.DefaultExpiration)
	return out, nil
}

func (g *PrivateTransactionManager) SendSignedTx(data []byte, to []string) (out []byte, err error) {
	out, err = g.node.SendSignedPayload(data, to)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (g *PrivateTransactionManager) Receive(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return data, nil
	}
	// Ignore this error since not being a recipient of
	// a payload isn't an error.
	// TODO: Return an error if it's anything OTHER than
	// 'you are not a recipient.'
	dataStr := string(data)
	x, found := g.c.Get(dataStr)
	if found {
		return x.([]byte), nil
	}
	pl, _ := g.node.ReceivePayload(data)
	g.c.Set(dataStr, pl, cache.DefaultExpiration)
	return pl, nil
}

func New(path string) (*PrivateTransactionManager, error) {
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
