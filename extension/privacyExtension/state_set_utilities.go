package privacyExtension

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	extension "github.com/ethereum/go-ethereum/extension/extensionContracts"
	"github.com/ethereum/go-ethereum/log"
)

func setState(privateState *state.StateDB, accounts map[string]extension.AccountWithMetadata, privacyMetaData *state.PrivacyMetadata) bool {
	log.Debug("Extension: set private state explicitly from state dump")
	for key, value := range accounts {
		stateDump := value.State

		contractAddress := common.HexToAddress(key)

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
		privateState.SetStatePrivacyMetadata(contractAddress, privacyMetaData)
	}
	return true
}

// updates the privacy metadata
func setPrivacyMetadata(privateState *state.StateDB, address common.Address, hash string) {
	privacyMetaData, err := privateState.GetStatePrivacyMetadata(address)
	if err != nil || privacyMetaData.PrivacyFlag.IsStandardPrivate() {
		return
	}

	ptmHash, err := common.Base64ToEncryptedPayloadHash(hash)
	if err != nil {
		log.Error("setting privacy metadata failed", "err", err)
		return
	}
	pm := state.NewStatePrivacyMetadata(ptmHash, privacyMetaData.PrivacyFlag)
	privateState.SetStatePrivacyMetadata(address, pm)
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
