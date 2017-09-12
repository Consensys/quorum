package private

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/constellation"
	"os"
)

type PrivateTransactionManager interface {
	Send(realTo *common.Address, data []byte, from string, to []string) (*common.Address, []byte, error)
	Receive(data []byte) (*common.Address, []byte, error)
	NullAddressProxy() common.Address
	ParseConstellationPayload(data []byte) (realTo *common.Address, realData []byte, err error)
}

func FromEnvironmentOrNil(name string) PrivateTransactionManager {
	cfgPath := os.Getenv(name)
	if cfgPath == "" {
		return nil
	}
	return constellation.MustNew(cfgPath)
}

var P = FromEnvironmentOrNil("PRIVATE_CONFIG")
