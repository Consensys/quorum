package constellation

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/patrickmn/go-cache"
	"time"
)

func copyBytes(b []byte) []byte {
	ob := make([]byte, len(b))
	copy(ob, b)
	return ob
}

type Constellation struct {
	node             *Client
	c                *cache.Cache
	maskAddress      common.Address
	nullAddressProxy common.Address
}

func (g *Constellation) Send(toField *common.Address, data []byte, from string, to []string) (maskTo *common.Address, out []byte, err error) {
	var realToField common.Address
	if toField == nil {
		realToField = g.nullAddressProxy
	} else {
		realToField = *toField
	}
	payload := append(realToField.Bytes(), data...)
	log.Info("Sending payload to constellation", "payload", common.ToHex(payload))
	if len(data) > 0 {
		if len(to) == 0 {
			out = copyBytes(payload)
		} else {
			var err error
			out, err = g.node.SendPayload(payload, from, to)
			if err != nil {
				return nil, nil, err
			}
		}
	}
	g.c.Set(string(out), payload, cache.DefaultExpiration)
	return &g.maskAddress, out, nil
}

func (g *Constellation) ParseConstellationPayload(dataWithTo []byte) (*common.Address, []byte, error) {
	if len(dataWithTo) < 20 {
		log.Info("Didn't find a valid payload in constellation, indicating this transaction is not for us.", "payload", common.ToHex(dataWithTo))
		return nil, nil, fmt.Errorf("Malformed constellation payload")

	}
	realTo := common.BytesToAddress(dataWithTo[:20])
	realPayload := dataWithTo[20:]
	if realTo != g.nullAddressProxy {
		return &realTo, realPayload, nil
	} else {
		return nil, realPayload, nil
	}

}

func (g *Constellation) NullAddressProxy() common.Address {
	return g.nullAddressProxy
}

func (g *Constellation) Receive(data []byte) (*common.Address, []byte, error) {
	dataStr := string(data)
	x, found := g.c.Get(dataStr)
	if found {
		realTo, realData, err := g.ParseConstellationPayload(x.([]byte))
		if err != nil {
			return nil, nil, err
		}
		if realTo == nil {
			log.Info("Received private contract creation payload from cache", "data", common.ToHex(realData))
		} else {
			log.Info("Received private payload from cache to address", "to", realTo.Hex(), "data", common.ToHex(realData))
		}
		return realTo, realData, nil
	}
	// Ignore this error since not being a recipient of
	// a payload isn't an error.
	// TODO: Return an error if it's anything OTHER than
	// 'you are not a recipient.'
	dataWithTo, _ := g.node.ReceivePayload(data)
	realTo, realData, err := g.ParseConstellationPayload(dataWithTo)
	if err != nil {
		return nil, nil, err
	}
	g.c.Set(dataStr, dataWithTo, cache.DefaultExpiration)
	if realTo == nil {
		log.Info("Received contract creation payload from constellation", "data", common.ToHex(realData))
	} else {
		log.Info("Received payload from constellation", "to", realTo.Hex(), "data", common.ToHex(realData))
	}
	return realTo, realData, nil
}

func New(configPath string) (*Constellation, error) {
	cfg, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}
	err = RunNode(configPath, cfg.Socket)
	if err != nil {
		return nil, err
	}
	n, err := NewClient(cfg.PublicKeys[0], cfg.Socket)
	if err != nil {
		return nil, err
	}
	maskAddr := common.BytesToAddress(common.FromHex(cfg.ToMask))
	nullProxy := common.BytesToAddress(common.FromHex(cfg.NullProxy))
	return &Constellation{
		node:             n,
		c:                cache.New(5*time.Minute, 5*time.Minute),
		maskAddress:      maskAddr,
		nullAddressProxy: nullProxy,
	}, nil
}

func MustNew(configPath string) *Constellation {
	g, err := New(configPath)
	if err != nil {
		panic(fmt.Sprintf("MustNew error: %v", err))
	}
	return g
}

func MaybeNew(configPath string) *Constellation {
	if configPath == "" {
		return nil
	}
	return MustNew(configPath)
}
