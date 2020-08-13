package private

import (
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/private/engine/notinuse"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/privatetransactionmanager"
)

// Interacting with Private Transaction Manager APIs
type PrivateTransactionManager interface {
	Send(data []byte, from string, to []string) (common.EncryptedPayloadHash, error)
	StoreRaw(data []byte, from string) (common.EncryptedPayloadHash, error)
	SendSignedTx(txHash common.EncryptedPayloadHash, to []string) ([]byte, error)
	Receive(txHash common.EncryptedPayloadHash) ([]byte, error)

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
