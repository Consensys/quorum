package privacyExtension

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	extension "github.com/ethereum/go-ethereum/extension/extensionContracts"
	"github.com/ethereum/go-ethereum/log"
)

const stateSharedTopicHash = "0x67a92539f3cbd7c5a9b36c23c0e2beceb27d2e1b3cd8eda02c623689267ae71e"

func setState(privateState *state.StateDB, accounts map[string]extension.AccountWithMetadata) bool {
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
			privateState.SetState(contractAddress, common.HexToHash(keyStore), common.HexToHash(valueStore))
		}
	}
	return true
}

func logContainsExtensionTopic(receivedLog *types.Log) bool {
	if len(receivedLog.Topics) != 1 {
		return false
	}
	return receivedLog.Topics[0].String() == stateSharedTopicHash
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