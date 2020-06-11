package notinuse

import (
	"errors"

	"github.com/ethereum/go-ethereum/private/engine"
)

var ErrPrivateTxManagerNotInUse = errors.New("private transaction manager is not in use")

// NotInUsePrivateTxManager returns an error for all communication functions,
// stating that no private transaction manager is being used by the node
type PrivateTransactionManager struct{}

func (ptm *PrivateTransactionManager) Send(data []byte, from string, to []string) ([]byte, error) {
	return nil, ErrPrivateTxManagerNotInUse
}

func (ptm *PrivateTransactionManager) StoreRaw(data []byte, from string) ([]byte, error) {
	return nil, ErrPrivateTxManagerNotInUse
}

func (ptm *PrivateTransactionManager) SendSignedTx(data []byte, to []string) ([]byte, error) {
	return nil, ErrPrivateTxManagerNotInUse
}

func (ptm *PrivateTransactionManager) Receive(data []byte) ([]byte, error) {
	//error not thrown here, acts as though no private data to fetch
	return nil, nil
}

func (ptm *PrivateTransactionManager) Name() string {
	return "NotInUse"
}

func (ptm *PrivateTransactionManager) HasFeature(f engine.PrivateTransactionManagerFeature) bool {
	return false
}
