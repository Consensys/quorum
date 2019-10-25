package privacy_extension

import (
	"github.com/ethereum/go-ethereum/common"
	extension "github.com/ethereum/go-ethereum/contract-extension/contractExtensionContracts"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

const stateSharedTopicHash = "0x40b79448ff8678eac1487385427aa682ee6ee831ce0702c09f95255645428531"

func setState(privateState *state.StateDB, accounts map[string]extension.AccountWithMetadata) bool {
	for key, value := range accounts {
		stateDump := value.State

		contractAddress := common.HexToAddress(key)

		newBalance, errBalanceSet := new(big.Int).SetString(stateDump.Balance, 10)
		if !errBalanceSet {
			log.Warn("could not set address balance", "address", key, "balance", stateDump.Balance)
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