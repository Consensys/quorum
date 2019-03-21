package engine

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrPrivateTxManagerNotinUse     = errors.New("private transaction manager is not in use")
	ErrPrivateTxManagerNotReady     = errors.New("private transaction manager is not ready")
	ErrPrivateTxManagerNotSupported = errors.New("private transaction manager does not suppor this operation")
)

type NotInUsePrivateTxManager struct{}

func (dn *NotInUsePrivateTxManager) Send(data []byte, from string, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) (common.EncryptedPayloadHash, error) {
	return common.EncryptedPayloadHash{}, ErrPrivateTxManagerNotinUse
}

func (dn *NotInUsePrivateTxManager) SendSignedTx(data common.EncryptedPayloadHash, to []string, acHashes common.EncryptedPayloadHashes, acMerkleRoot common.Hash) ([]byte, error) {
	return nil, ErrPrivateTxManagerNotinUse
}

func (dn *NotInUsePrivateTxManager) Receive(data common.EncryptedPayloadHash) ([]byte, common.EncryptedPayloadHashes, common.Hash, error) {
	return nil, common.EncryptedPayloadHashes{}, common.Hash{}, ErrPrivateTxManagerNotinUse
}

func (dn *NotInUsePrivateTxManager) Name() string {
	return "NotInUse"
}
