package notinuse

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

var ErrPrivateTxManagerNotInUse = errors.New("private transaction manager is not in use")

// NotInUsePrivateTxManager returns an error for all communication functions,
// stating that no private transaction manager is being used by the node
type PrivateTransactionManager struct{}

func (ptm *PrivateTransactionManager) IsSender(txHash common.EncryptedPayloadHash) (bool, error) {
	panic("implement me")
}

func (ptm *PrivateTransactionManager) GetParticipants(txHash common.EncryptedPayloadHash) ([]string, error) {
	panic("implement me")
}

func (ptm *PrivateTransactionManager) Send(data []byte, from string, to []string) (common.EncryptedPayloadHash, error) {
	return common.EncryptedPayloadHash{}, ErrPrivateTxManagerNotInUse
}

func (ptm *PrivateTransactionManager) StoreRaw(data []byte, from string) (common.EncryptedPayloadHash, error) {
	return common.EncryptedPayloadHash{}, ErrPrivateTxManagerNotInUse
}

func (ptm *PrivateTransactionManager) SendSignedTx(txHash common.EncryptedPayloadHash, to []string) ([]byte, error) {
	return nil, ErrPrivateTxManagerNotInUse
}

func (ptm *PrivateTransactionManager) Receive(txHash common.EncryptedPayloadHash) ([]byte, error) {
	//error not thrown here, acts as though no private data to fetch
	return nil, nil
}

func (ptm *PrivateTransactionManager) Name() string {
	return "NotInUse"
}
