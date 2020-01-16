package privatetransactionmanager

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

type PrivateTransactionManager struct {
	node                                *Client
	c                                   *cache.Cache
	isPrivateTransactionManagerNotInUse bool
}

var (
	errPrivateTransactionManagerNotUsed = errors.New("private transaction manager not in use")
)

func (g *PrivateTransactionManager) Send(data []byte, from string, to []string) (out []byte, err error) {
	if g.isPrivateTransactionManagerNotInUse {
		return nil, errPrivateTransactionManagerNotUsed
	}
	out, err = g.node.SendPayload(data, from, to)
	if err != nil {
		return nil, err
	}
	g.c.Set(string(out), data, cache.DefaultExpiration)
	return out, nil
}

func (g *PrivateTransactionManager) SendSignedTx(data []byte, to []string) (out []byte, err error) {
	if g.isPrivateTransactionManagerNotInUse {
		return nil, errPrivateTransactionManagerNotUsed
	}
	out, err = g.node.SendSignedPayload(data, to)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (g *PrivateTransactionManager) Receive(data []byte) ([]byte, error) {
	if g.isPrivateTransactionManagerNotInUse {
		return nil, nil
	}
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

func (g *PrivateTransactionManager) IsSender(txHash common.EncryptedPayloadHash) (bool, error) {
	if g.isPrivateTransactionManagerNotInUse {
		return false, errPrivateTransactionManagerNotUsed
	}
	return g.node.IsSender(txHash)
}

func (g *PrivateTransactionManager) GetParticipants(txHash common.EncryptedPayloadHash) ([]string, error) {
	if g.isPrivateTransactionManagerNotInUse {
		return nil, errPrivateTransactionManagerNotUsed
	}
	return g.node.GetParticipants(txHash)
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
		node:                                n,
		c:                                   cache.New(5*time.Minute, 5*time.Minute),
		isPrivateTransactionManagerNotInUse: false,
	}, nil
}

func MustNew(path string) *PrivateTransactionManager {
	if strings.EqualFold(path, "ignore") {
		return &PrivateTransactionManager{
			node:                                nil,
			c:                                   nil,
			isPrivateTransactionManagerNotInUse: true,
		}
	}
	g, err := New(path)
	if err != nil {
		panic(fmt.Sprintf("MustNew: Failed to connect to private transaction manager (%s): %v", path, err))
	}
	return g
}
