package notinuse

import (
	"github.com/ethereum/go-ethereum/private/engine"
)

// DBUpgradePrivateTransactionManager returns an error for all communication functions,
// while reporting it has the MultiplePrivateStates feature
type DBUpgradePrivateTransactionManager struct {
	PrivateTransactionManager
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

func (ptm *DBUpgradePrivateTransactionManager) Name() string {
	return "dbupgrade"
}

func (ptm *DBUpgradePrivateTransactionManager) HasFeature(f engine.PrivateTransactionManagerFeature) bool {
	return f == engine.MultiplePrivateStates
}
