package core

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private/engine"
)

type pmcStateTransitionAPI interface {
	SetTxPrivacyMetadata(pm *types.PrivacyMetadata)
	IsPrivacyEnhancementsEnabled() bool
	RevertToSnapshot(int)
	GetStatePrivacyMetadata(addr common.Address) (*state.PrivacyMetadata, error)
	CalculateMerkleRoot() (common.Hash, error)
	AffectedContracts() []common.Address
}

func newPMH(st pmcStateTransitionAPI) *privateMessageHandler {
	return &privateMessageHandler{stAPI: st}
}

type privateMessageHandler struct {
	stAPI pmcStateTransitionAPI

	hasPrivatePayload bool

	snapshot                int
	receivedPrivacyMetadata *engine.ExtraMetadata
	eph                     common.EncryptedPayloadHash
}

func (pmh *privateMessageHandler) mustVerify() bool {
	return pmh.hasPrivatePayload && pmh.receivedPrivacyMetadata != nil && pmh.stAPI.IsPrivacyEnhancementsEnabled()
}

// checks the privacy metadata in the state transition context
// returns vmError if there is an error in the EVM execution
// returns consensusErr if there is an error in the consensus execution
func (pmh *privateMessageHandler) prepare() (vmError, consensusErr error) {
	if pmh.receivedPrivacyMetadata != nil {
		if !pmh.stAPI.IsPrivacyEnhancementsEnabled() && pmh.receivedPrivacyMetadata.PrivacyFlag.IsNotStandardPrivate() {
			// This situation is only possible if the current node has been upgraded (both quorum and tessera) yet the
			// node did not apply the privacyEnhancementsBlock configuration (with a network agreed block height).
			// Since this would be considered node misconfiguration the behavior should be changed to return an error
			// which would then cause the node not to apply the block (and potentially get stuck and not be able to
			// continue to apply new blocks). The resolution should then be to revert to an appropriate block height and
			// run geth init with the network agreed privacyEnhancementsBlock.
			// The prepare method signature has been changed to allow returning the relevant error.
			return ErrPrivacyEnhancedReceivedWhenDisabled, fmt.Errorf("Privacy enhanced transaction received while privacy enhancements are disabled."+
				" Please check your node configuration. EPH=%s", pmh.eph.ToBase64())
		}

		if pmh.receivedPrivacyMetadata.PrivacyFlag == engine.PrivacyFlagStateValidation && common.EmptyHash(pmh.receivedPrivacyMetadata.ACMerkleRoot) {
			log.Error(ErrPrivacyMetadataInvalidMerkleRoot.Error())
			return ErrPrivacyMetadataInvalidMerkleRoot, nil
		}
		privMetadata := types.NewTxPrivacyMetadata(pmh.receivedPrivacyMetadata.PrivacyFlag)
		pmh.stAPI.SetTxPrivacyMetadata(privMetadata)
	}
	return nil, nil
}

// If the list of affected CA Transactions by the time evm executes is different from the list of affected contract transactions returned from Tessera
// an Error should be thrown and the state should not be updated
// This validation is to prevent cases where the list of affected contract will have changed by the time the evm actually executes transaction
// failed = true will make sure receipt is marked as "failure"
// return error will crash the node and only use when that's the case
func (pmh *privateMessageHandler) verify(vmerr error) (bool, error) {
	// convenient function to return error. It has the same signature as the main function
	returnErrorFunc := func(anError error, logMsg string, ctx ...interface{}) (exitEarly bool, err error) {
		if logMsg != "" {
			log.Debug(logMsg, ctx...)
		}
		pmh.stAPI.RevertToSnapshot(pmh.snapshot)
		exitEarly = true
		if anError != nil {
			err = fmt.Errorf("vmerr=%s, err=%s", vmerr, anError)
		}
		return
	}
	actualACAddresses := pmh.stAPI.AffectedContracts()
	log.Trace("Verify hashes of affected contracts", "expectedHashes", pmh.receivedPrivacyMetadata.ACHashes, "numberOfAffectedAddresses", len(actualACAddresses))
	privacyFlag := pmh.receivedPrivacyMetadata.PrivacyFlag
	for _, addr := range actualACAddresses {
		// GetPrivacyMetadata is invoked on the privateState (as the tx is private) and it returns:
		// 1. public contacts: privacyMetadata = nil, err = nil
		// 2. private contracts of type:
		// 2.1. StandardPrivate:     privacyMetadata = nil, err = "The provided contract does not have privacy metadata"
		// 2.2. PartyProtection/PSV: privacyMetadata = <data>, err = nil
		actualPrivacyMetadata, err := pmh.stAPI.GetStatePrivacyMetadata(addr)
		//when privacyMetadata should have been recovered but wasnt (includes non-party)
		//non party will only be caught here if sender provides privacyFlag
		if err != nil && privacyFlag.IsNotStandardPrivate() {
			return returnErrorFunc(nil, "Unable to find PrivacyMetadata for affected contract", "err", err, "addr", addr.Hex())
		}
		log.Trace("Privacy metadata", "affectedAddress", addr.Hex(), "metadata", actualPrivacyMetadata)
		// both public and standard private contracts will be nil and can be skipped in acoth check
		// public contracts - evm error for write, no error for reads
		// standard private - only error if privacyFlag sent with tx or if no flag sent but other affecteds have privacyFlag
		if actualPrivacyMetadata == nil {
			continue
		}
		// Check that the affected contracts privacy flag matches the transaction privacy flag.
		// I know that this is also checked by tessera, but it only checks for non standard private transactions.
		if actualPrivacyMetadata.PrivacyFlag != pmh.receivedPrivacyMetadata.PrivacyFlag {
			return returnErrorFunc(nil, "Mismatched privacy flags",
				"affectedContract.Address", addr.Hex(),
				"affectedContract.PrivacyFlag", actualPrivacyMetadata.PrivacyFlag,
				"received.PrivacyFlag", pmh.receivedPrivacyMetadata.PrivacyFlag)
		}
		// acoth check - case where node isn't privy to one of actual affecteds
		if pmh.receivedPrivacyMetadata.ACHashes.NotExist(actualPrivacyMetadata.CreationTxHash) {
			return returnErrorFunc(nil, "Participation check failed",
				"affectedContractAddress", addr.Hex(),
				"missingCreationTxHash", actualPrivacyMetadata.CreationTxHash.Hex())
		}
	}

	// check the psv merkle root comparison - for both creation and msg calls
	if !common.EmptyHash(pmh.receivedPrivacyMetadata.ACMerkleRoot) {
		log.Trace("Verify merkle root", "merkleRoot", pmh.receivedPrivacyMetadata.ACMerkleRoot)
		actualACMerkleRoot, err := pmh.stAPI.CalculateMerkleRoot()
		if err != nil {
			return returnErrorFunc(err, "")
		}
		if actualACMerkleRoot != pmh.receivedPrivacyMetadata.ACMerkleRoot {
			return returnErrorFunc(nil, "Merkle Root check failed", "actual", actualACMerkleRoot,
				"expect", pmh.receivedPrivacyMetadata.ACMerkleRoot)
		}
	}
	return false, nil
}
