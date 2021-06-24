package eth

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestBestPeer(t *testing.T) {
	ps := newPeerSet()
	err := ps.Register(nil, nil, "")
	assert.Error(t, err, "expect error telling that peer can't be nil")

	p1 := &peer{
		id:   "1",
		head: common.Hash{},
	}
	err = ps.Register(p1, nil, "")
	assert.NoError(t, err)

	p2 := &peer{
		id:   "2",
		head: common.Hash{},
	}
	err = ps.Register(p2, nil, "")
	assert.NoError(t, err)

	bp := ps.BestPeer()
	assert.NotNil(t, bp)

	p1.td = big.NewInt(12)
	bp = ps.BestPeer()
	assert.Equal(t, p1, bp)

	p2.td = big.NewInt(13)
	bp = ps.BestPeer()
	assert.Equal(t, p2, bp)

	p1.td = big.NewInt(14)
	bp = ps.BestPeer()
	assert.Equal(t, p1, bp)

	p1.td = nil
	bp = ps.BestPeer()
	assert.Equal(t, p2, bp)
}
