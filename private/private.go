package private

import (
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/private/engine/notinuse"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/privatetransactionmanager"
)

type PrivateTransactionManager interface {
	Send(data []byte, from string, to []string) ([]byte, error)
	StoreRaw(data []byte, from string) ([]byte, error)
	SendSignedTx(data []byte, to []string) ([]byte, error)
	Receive(data []byte) ([]byte, error)

	IsSender(txHash common.EncryptedPayloadHash) (bool, error)
	GetParticipants(txHash common.EncryptedPayloadHash) ([]string, error)
}

func FromEnvironmentOrNil(name string) PrivateTransactionManager {
	cfgPath := os.Getenv(name)
	if cfgPath == "" {
		return nil
	}
	if strings.EqualFold(cfgPath, "ignore") {
		return &notinuse.PrivateTransactionManager{}
	}
	return privatetransactionmanager.MustNew(cfgPath)
}

var P = FromEnvironmentOrNil("PRIVATE_CONFIG")
