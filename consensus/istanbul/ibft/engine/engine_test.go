package ibftengine

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestEngine(t *testing.T) {
	_ = NewEngine(nil, common.Address{}, nil)
}
