package constellation

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/patrickmn/go-cache"
)

type Constellation struct {
	node   *Client
	c      *cache.Cache
	ignore bool
}

var (
	ErrConstellationIsntInit = errors.New("ignoreConstellation")
)

func (g *Constellation) Send(data []byte, from string, to []string) (out []byte, err error) {
	if g.ignore {
		return nil, ErrConstellationIsntInit
	}
	out, err = g.node.SendPayload(data, from, to)
	if err != nil {
		return nil, err
	}
	g.c.Set(string(out), data, cache.DefaultExpiration)
	return out, nil
}

func (g *Constellation) Receive(data []byte) ([]byte, error) {
	if g.ignore {
		log.Trace("should ignore", "data", fmt.Sprintf("%v", data))
		return nil, nil
	}
	if len(data) == 0 {
		log.Trace("data has no length", "data", fmt.Sprintf("%v", data))
		return data, nil
	}
	// Ignore this error since not being a recipient of
	// a payload isn't an error.
	// TODO: Return an error if it's anything OTHER than
	// 'you are not a recipient.'
	dataStr := string(data)
	log.Trace("data being processed", "datastring", fmt.Sprintf("%v", dataStr))
	log.Trace("data being processed", "datastring", fmt.Sprintf("%v", data))
	x, found := g.c.Get(dataStr)
	if found {
		return x.([]byte), nil
	}
	pl, _ := g.node.ReceivePayload(data)
	g.c.Set(dataStr, pl, cache.DefaultExpiration)
	log.Trace("data being returned", "data", fmt.Sprintf("%v", pl))
	return pl, nil
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
		node:   n,
		c:      cache.New(5*time.Minute, 5*time.Minute),
		ignore: false,
	}, nil
}

func MustNew(path string) *Constellation {
	if strings.EqualFold(path, "ignore") {
		return &Constellation{
			node:   nil,
			c:      nil,
			ignore: true,
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
