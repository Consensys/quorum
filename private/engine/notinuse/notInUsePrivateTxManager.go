package notinuse

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/engine"
)

var ErrPrivateTxManagerNotInUse = errors.New("private transaction manager is not in use")

// NotInUsePrivateTxManager returns an error for all communication functions,
// stating that no private transaction manager is being used by the node
type PrivateTransactionManager struct{}

func (ptm *PrivateTransactionManager) IsSender(txHash common.EncryptedPayloadHash) (bool, error) {
	panic("implement me")
}

func (ptm *PrivateTransactionManager) Groups() ([]engine.PrivacyGroup, error) {
	panic("implement me")
}

func (ptm *PrivateTransactionManager) GetParticipants(txHash common.EncryptedPayloadHash) ([]string, error) {
	panic("implement me")
}

func (ptm *PrivateTransactionManager) GetMandatory(txHash common.EncryptedPayloadHash) ([]string, error) {
	panic("implement me")
}

func (ptm *PrivateTransactionManager) Send(data []byte, from string, to []string, extra *engine.ExtraMetadata) (string, []string, common.EncryptedPayloadHash, error) {
	return "", nil, common.EncryptedPayloadHash{}, engine.ErrPrivateTxManagerNotinUse
}

func (ptm *PrivateTransactionManager) EncryptPayload(data []byte, from string, to []string, extra *engine.ExtraMetadata) ([]byte, error) {
	return nil, engine.ErrPrivateTxManagerNotinUse
}

func (ptm *PrivateTransactionManager) DecryptPayload(payload common.DecryptRequest) ([]byte, *engine.ExtraMetadata, error) {
	return nil, nil, engine.ErrPrivateTxManagerNotSupported
}

func (ptm *PrivateTransactionManager) StoreRaw(data []byte, from string) (common.EncryptedPayloadHash, error) {
	return common.EncryptedPayloadHash{}, engine.ErrPrivateTxManagerNotinUse
}

func (ptm *PrivateTransactionManager) SendSignedTx(data common.EncryptedPayloadHash, to []string, extra *engine.ExtraMetadata) (string, []string, []byte, error) {
	return "", nil, nil, engine.ErrPrivateTxManagerNotinUse
}

func (ptm *PrivateTransactionManager) Receive(data common.EncryptedPayloadHash) (string, []string, []byte, *engine.ExtraMetadata, error) {
	//error not thrown here, acts as though no private data to fetch
	return "", nil, nil, nil, nil
}

func (ptm *PrivateTransactionManager) ReceiveRaw(data common.EncryptedPayloadHash) ([]byte, string, *engine.ExtraMetadata, error) {
	return nil, "", nil, engine.ErrPrivateTxManagerNotinUse
}

func (ptm *PrivateTransactionManager) Name() string {
	return "NotInUse"
}

func (ptm *PrivateTransactionManager) HasFeature(f engine.PrivateTransactionManagerFeature) bool {
	return false
}
