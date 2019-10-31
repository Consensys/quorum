package private

import (
	"os"

	"github.com/ethereum/go-ethereum/private/constellation"
)

type PrivateTransactionManager interface {
	Send(data []byte, from string, to []string) ([]byte, error)
	SendSignedTx(data []byte, to []string) ([]byte, error)
	Receive(data []byte) ([]byte, error)
}

func FromEnvironment(name string) PrivateTransactionManager {
	return constellation.MustNew(os.Getenv(name))
}

var P = FromEnvironment("PRIVATE_CONFIG")
