package private

import (
	"os"

	"github.com/ethereum/go-ethereum/private/privatetransactionmanager"
)

type PrivateTransactionManager interface {
	Send(data []byte, from string, to []string) ([]byte, error)
	SendSignedTx(data []byte, to []string) ([]byte, error)
	Receive(data []byte) ([]byte, error)
}

func FromEnvironmentOrNil(name string) PrivateTransactionManager {
	cfgPath := os.Getenv(name)
	if cfgPath == "" {
		return nil
	}
	return privatetransactionmanager.MustNew(cfgPath)
}

var P = FromEnvironmentOrNil("PRIVATE_CONFIG")
