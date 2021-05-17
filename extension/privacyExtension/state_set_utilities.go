package privacyExtension

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	extension "github.com/ethereum/go-ethereum/extension/extensionContracts"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"
)

func setState(privateState *state.StateDB, accounts map[string]extension.AccountWithMetadata, privacyMetaData *state.PrivacyMetadata, managedParties []string) bool {
	log.Debug("Extension: set private state explicitly from state dump")
	for key, value := range accounts {
		stateDump := value.State

		contractAddress := common.HexToAddress(key)

		privateState.CreateAccount(contractAddress)
		newBalance, errBalanceSet := new(big.Int).SetString(stateDump.Balance, 10)
		if !errBalanceSet {
			log.Error("could not set address balance", "address", key, "balance", stateDump.Balance)
			return false
		}
		privateState.SetBalance(contractAddress, newBalance)
		privateState.SetNonce(contractAddress, stateDump.Nonce)
		privateState.SetCode(contractAddress, common.Hex2Bytes(stateDump.Code))
		for keyStore, valueStore := range stateDump.Storage {
			privateState.SetState(contractAddress, keyStore, common.HexToHash(valueStore))
		}
		if privacyMetaData.PrivacyFlag != engine.PrivacyFlagStandardPrivate {
			privateState.SetPrivacyMetadata(contractAddress, privacyMetaData)
		}
		if managedParties != nil {
			privateState.SetManagedParties(contractAddress, managedParties)
		}
	}
	return true
}

// updates the privacy metadata
func setPrivacyMetadata(privateState *state.StateDB, address common.Address, hash string) {
	privacyMetaData, err := privateState.GetPrivacyMetadata(address)
	if err != nil || privacyMetaData.PrivacyFlag.IsStandardPrivate() {
		return
	}

	ptmHash, err := common.Base64ToEncryptedPayloadHash(hash)
	if err != nil {
		log.Error("setting privacy metadata failed", "err", err)
		return
	}
	pm := state.NewStatePrivacyMetadata(ptmHash, privacyMetaData.PrivacyFlag)
	privateState.SetPrivacyMetadata(address, pm)
}

func setManagedParties(ptm private.PrivateTransactionManager, privateState *state.StateDB, address common.Address, hash string) {
	existingManagedParties, err := privateState.GetManagedParties(address)
	if err != nil {
		return
	}

	ptmHash, err := common.Base64ToEncryptedPayloadHash(hash)
	if err != nil {
		log.Error("setting privacy metadata failed", "err", err)
		return
	}

	_, managedParties, _, _, _ := ptm.Receive(ptmHash)
	newManagedParties := common.AppendSkipDuplicates(existingManagedParties, managedParties...)
	privateState.SetManagedParties(address, newManagedParties)
}

func logContainsExtensionTopic(receivedLog *types.Log) bool {
	if len(receivedLog.Topics) != 1 {
		return false
	}
	return receivedLog.Topics[0].String() == extension.StateSharedTopicHash
}

// validateAccountsExist checks that all the accounts in the expected list are
// present in the state map, and that no  other accounts exist in the state map
// that are unexpected
func validateAccountsExist(expectedAccounts []common.Address, actualAccounts map[string]extension.AccountWithMetadata) bool {
	if len(expectedAccounts) != len(actualAccounts) {
		return false
	}
	for _, account := range expectedAccounts {
		_, exists := actualAccounts[account.String()]
		if !exists {
			return false
		}
	}
	return true
}
