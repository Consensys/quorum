package mpsdbupgrade

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/engine"
)

var ErrOperationNotSupportedByDBUpgradeTXManager = errors.New("opperation not supported by db upgrade tx manager")

// DBUpgradePrivateTransactionManager returns an error for all communication functions,
// while reporting it has the MultiplePrivateStates feature
type DBUpgradePrivateTransactionManager struct{}

func (ptm *DBUpgradePrivateTransactionManager) IsSender(txHash common.EncryptedPayloadHash) (bool, error) {
	return false, ErrOperationNotSupportedByDBUpgradeTXManager
}

func (ptm *DBUpgradePrivateTransactionManager) Groups() ([]engine.PrivacyGroup, error) {
	return []engine.PrivacyGroup{
		{
			Type:           "resident",
			Name:           "private",
			PrivacyGroupId: "private",
			Description:    "default resident group",
			From:           "",
			Members:        nil,
		},
	}, nil
}

func (ptm *DBUpgradePrivateTransactionManager) GetParticipants(txHash common.EncryptedPayloadHash) ([]string, error) {
	return nil, ErrOperationNotSupportedByDBUpgradeTXManager
}

func (ptm *DBUpgradePrivateTransactionManager) Send(data []byte, from string, to []string, extra *engine.ExtraMetadata) (string, []string, common.EncryptedPayloadHash, error) {
	return "", nil, common.EncryptedPayloadHash{}, ErrOperationNotSupportedByDBUpgradeTXManager
}

func (ptm *DBUpgradePrivateTransactionManager) EncryptPayload(data []byte, from string, to []string, extra *engine.ExtraMetadata) ([]byte, error) {
	return nil, ErrOperationNotSupportedByDBUpgradeTXManager
}

func (ptm *DBUpgradePrivateTransactionManager) DecryptPayload(payload common.DecryptRequest) ([]byte, *engine.ExtraMetadata, error) {
	return nil, nil, ErrOperationNotSupportedByDBUpgradeTXManager
}

func (ptm *DBUpgradePrivateTransactionManager) StoreRaw(data []byte, from string) (common.EncryptedPayloadHash, error) {
	return common.EncryptedPayloadHash{}, ErrOperationNotSupportedByDBUpgradeTXManager
}

func (ptm *DBUpgradePrivateTransactionManager) SendSignedTx(data common.EncryptedPayloadHash, to []string, extra *engine.ExtraMetadata) (string, []string, []byte, error) {
	return "", nil, nil, ErrOperationNotSupportedByDBUpgradeTXManager
}

func (ptm *DBUpgradePrivateTransactionManager) Receive(data common.EncryptedPayloadHash) (string, []string, []byte, *engine.ExtraMetadata, error) {
	//error not thrown here, acts as though no private data to fetch
	return "", nil, nil, nil, nil
}

func (ptm *DBUpgradePrivateTransactionManager) ReceiveRaw(data common.EncryptedPayloadHash) ([]byte, string, *engine.ExtraMetadata, error) {
	return nil, "", nil, nil
}

func (ptm *DBUpgradePrivateTransactionManager) Name() string {
	return "dbupgrade"
}

func (ptm *DBUpgradePrivateTransactionManager) HasFeature(f engine.PrivateTransactionManagerFeature) bool {
	if f == engine.MultiplePrivateStates {
		return true
	}
	return false
}
