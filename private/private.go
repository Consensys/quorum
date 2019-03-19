package private

import (
	"os"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/private/constellation"
)

type PrivateTransactionManager interface {
	Send(data []byte, from string, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) (common.EncryptedPayloadHash, error)
	SendSignedTx(data common.EncryptedPayloadHash, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) ([]byte, error)
	Receive(data common.EncryptedPayloadHash) ([]byte, common.EncryptedPayloadHashes, common.Hash, error)
}

func FromEnvironmentOrNil(name string) PrivateTransactionManager {
	cfgPath := os.Getenv(name)
	if cfgPath == "" {
		return nil
	}
	return constellation.MustNew(cfgPath)
}

var P = FromEnvironmentOrNil("PRIVATE_CONFIG")
