package private

import (
	"os"

	"github.com/jpmorganchase/quorum/private/constellation"
)

type PrivateTransactionManager interface {
	Send(data []byte, from string, to []string) ([]byte, error)
	Receive(data []byte) ([]byte, error)
}

func FromEnvironmentOrNil(name string) PrivateTransactionManager {
	cfgPath := os.Getenv(name)
	if cfgPath == "" {
		return nil
	}
	return constellation.MustNew(cfgPath)
}

var P = FromEnvironmentOrNil("PRIVATE_CONFIG")
