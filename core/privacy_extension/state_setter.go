package privacy_extension

import (
	"encoding/base64"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	extension "github.com/ethereum/go-ethereum/contract-extension/contractExtensionContracts"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private"
	"math/big"
)

const stateSharedTopicHash = "0x40b79448ff8678eac1487385427aa682ee6ee831ce0702c09f95255645428531"

func CheckIfExtensionHappened(txLogs []*types.Log, privateState *state.StateDB) {
	// there should be two logs,
	// the first being the state extension log, the second being the event finished log
	if len(txLogs) != 2 {
		//not an extension transaction, so don't check
		return
	}

	if txLog := txLogs[0]; LogContainsExtensionTopic(txLog) {
		//this is a direct state share
		decodedLog := new(extension.ContractExtenderStateShared)
		if err := extension.ContractExtensionABI.Unpack(decodedLog, "StateShared", txLog.Data); err != nil {
			return
		}
		snapshotId := privateState.Snapshot()
		if success := HandleExtensionRequest(decodedLog.Hash, decodedLog.Uuid, privateState); !success {
			privateState.RevertToSnapshot(snapshotId)
		}
	}
}

func LogContainsExtensionTopic(receivedLog *types.Log) bool {
	if len(receivedLog.Topics) != 1 {
		return false
	}
	return receivedLog.Topics[0].String() == stateSharedTopicHash
}

func HandleExtensionRequest(hash string, uuid string, privateState *state.StateDB) bool {
	if uuidIsSentByUs := UuidIsOwn(uuid); !uuidIsSentByUs {
		return false
	}

	stateData, ok := FetchData(hash)
	if !ok {
		//there is nothing to do here, the state wasn't shared with us
		log.Info("Extension", "No state shared with us")
		return false
	}

	var accounts map[string]extension.AccountWithMetadata
	if err := json.Unmarshal(stateData, &accounts); err != nil {
		log.Info("Extension", "Could not unmarshal data")
		return false
	}

	SetState(privateState, accounts)

	return true
}

// Checks

func FetchData(hash string) ([]byte, bool){
	ptmHash, _ := base64.StdEncoding.DecodeString(hash)
	stateData, err := private.P.Receive(ptmHash)

	if stateData == nil || err != nil {
		return nil, false
	}
	return stateData, true
}

func UuidIsOwn(uuid string) bool {
	if uuid == "" {
		//we never called accept
		log.Info("Extension", "State shared by accept never called")
		return false
	}
	encryptedTxHash := common.BytesToEncryptedPayloadHash(common.FromHex(uuid))

	isSender, err := private.P.IsSender(encryptedTxHash)
	if err != nil || !isSender {
		log.Info("Extension", "We are not the sender")
		return false
	}
	return true
}

func SetState(privateState *state.StateDB, accounts map[string]extension.AccountWithMetadata) bool {
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