package privacyExtension

import (
	"bytes"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	extension "github.com/ethereum/go-ethereum/extension/extensionContracts"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private"
)

var DefaultExtensionHandler *ExtensionHandler

type ExtensionHandler struct {
	ptm           private.PrivateTransactionManager
	psis          core.PrivateStateIdentifierService
	isMultitenant bool
}

func Init() {
	DefaultExtensionHandler = NewExtensionHandler(private.P)
}

func NewExtensionHandler(transactionManager private.PrivateTransactionManager) *ExtensionHandler {
	return &ExtensionHandler{ptm: transactionManager}
}

func (handler *ExtensionHandler) SupportMultitenancy(b bool) {
	handler.isMultitenant = b
}

func (handler *ExtensionHandler) SetPSIS(psis core.PrivateStateIdentifierService) {
	handler.psis = psis
}

func (handler *ExtensionHandler) CheckExtensionAndSetPrivateState(txLogs []*types.Log, privateState *state.StateDB, psi string) {
	extraMetaDataUpdated := false
	for _, txLog := range txLogs {
		if !logContainsExtensionTopic(txLog) {
			continue
		}
		//this is a direct state share
		address, hash, uuid, err := extension.UnpackStateSharedLog(txLog.Data)
		if err != nil {
			continue
		}

		// check if state exists for the extension address. If yes then skip
		// processing
		if privateState.GetCode(address) != nil {
			if extraMetaDataUpdated {
				continue
			}
			// check the privacy flag of the contract. if its other than
			// 0 then need to update the privacy metadata for the contract
			//TODO: validate the old and new parties to ensure that all old parties are there
			setPrivacyMetadata(privateState, address, hash)
			if handler.isMultitenant {
				setManagedParties(handler.ptm, privateState, address, hash)
			}
			extraMetaDataUpdated = true
		} else {
			managedParties, accounts, privacyMetaData, found := handler.FetchStateData(txLog.Address, hash, uuid, psi)
			if !found {
				continue
			}
			if !handler.isMultitenant {
				managedParties = nil
			}
			if !validateAccountsExist([]common.Address{address}, accounts) {
				log.Error("Account mismatch", "expected", address, "found", accounts)
				continue
			}
			snapshotId := privateState.Snapshot()

			if success := setState(privateState, accounts, privacyMetaData, managedParties); !success {
				privateState.RevertToSnapshot(snapshotId)
			}
		}
	}
}

func (handler *ExtensionHandler) FetchStateData(address common.Address, hash string, uuid string, psi string) ([]string, map[string]extension.AccountWithMetadata, *state.PrivacyMetadata, bool) {
	if uuidIsSentByUs := handler.UuidIsOwn(address, uuid, psi); !uuidIsSentByUs {
		return nil, nil, nil, false
	}

	managedParties, stateData, privacyMetaData, ok := handler.FetchDataFromPTM(hash)
	if !ok {
		//there is nothing to do here, the state wasn't shared with us
		log.Error("Extension: No state shared with us")
		return nil, nil, nil, false
	}

	var accounts map[string]extension.AccountWithMetadata
	if err := json.Unmarshal(stateData, &accounts); err != nil {
		log.Error("Extension: Could not unmarshal data")
		return nil, nil, nil, false
	}

	return managedParties, accounts, privacyMetaData, true
}

// Checks

func (handler *ExtensionHandler) FetchDataFromPTM(hash string) ([]string, []byte, *state.PrivacyMetadata, bool) {
	ptmHash, _ := common.Base64ToEncryptedPayloadHash(hash)
	_, managedParties, stateData, extraMetaData, err := handler.ptm.Receive(ptmHash)

	if stateData == nil {
		log.Error("No state data found in PTM", "ptm hash", hash)
		return nil, nil, nil, false
	}
	if err != nil {
		log.Error("Error receiving state data from PTM", "ptm hash", hash, "err", err.Error())
		return nil, nil, nil, false
	}

	privacyMetaData := state.NewStatePrivacyMetadata(ptmHash, extraMetaData.PrivacyFlag)
	return managedParties, stateData, privacyMetaData, true
}

func (handler *ExtensionHandler) UuidIsOwn(address common.Address, uuid string, psi string) bool {
	if uuid == "" {
		//we never called accept
		log.Warn("Extension: State shared by accept never called")
		return false
	}
	encryptedTxHash := common.BytesToEncryptedPayloadHash(common.FromHex(uuid))
	isSender, err := handler.ptm.IsSender(encryptedTxHash)
	if err != nil {
		log.Debug("Extension: could not determine if we are sender", "err", err.Error())
		return false
	}

	if !isSender {
		return false
	}

	senderPublicKey, _, encryptedPayload, _, err := handler.ptm.Receive(encryptedTxHash)
	if err != nil {
		log.Debug("Extension: payload not found", "err", err)
		return false
	}

	//check the given PSI is same as PSI of sender key
	senderPsm, err := handler.psis.ResolveForManagedParty(senderPublicKey)
	if err != nil {
		log.Debug("Extension: unable to determine sender public key PSI", "err", err)
		return false
	}

	if psi != "private" && senderPsm.ID != psi {
		// sender was another tenant on this node
		//not an error case, so no need to log an error
		return false
	}

	var payload common.DecryptRequest
	if err := json.Unmarshal(encryptedPayload, &payload); err != nil {
		log.Debug("Extension: payload unmarshal failed", "err", err)
	}

	contractDetails, _, err := handler.ptm.DecryptPayload(payload)
	if err != nil {
		log.Debug("Extension: payload decrypt failed", "err", err)
	}

	if !bytes.Equal(contractDetails, address.Bytes()) {
		log.Error("Extension: wrong address in retrieved UUID")
		return false
	}
	return true
}
