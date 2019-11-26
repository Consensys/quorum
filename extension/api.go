package extension

import (
	"encoding/base64"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/private"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/internal/ethapi"
)

var (
	errNotVoter   = errors.New("account is not a voter of this extension request")
	errNotCreator = errors.New("account is not the creator of this extension request")
)

type PrivateExtensionAPI struct {
	privacyService *PrivacyService
	accountManager IAccountManager
	ptm    		   private.PrivateTransactionManager
}

func NewPrivateExtensionAPI(privacyService *PrivacyService, accountManager IAccountManager, ptm private.PrivateTransactionManager) *PrivateExtensionAPI {
	return &PrivateExtensionAPI{
		privacyService: privacyService,
		accountManager: accountManager,
		ptm: ptm,
	}
}

// ActiveExtensionContracts returns the list of all currently outstanding extension contracts
func (api *PrivateExtensionAPI) ActiveExtensionContracts() []ExtensionContract {
	api.privacyService.mu.Lock()
	defer api.privacyService.mu.Unlock()

	extracted := make([]ExtensionContract, 0, len(api.privacyService.currentContracts))
	for _, contract := range api.privacyService.currentContracts {
		extracted = append(extracted, *contract)
	}
	return extracted
}

// VoteOnContract submits the vote to the specified extension management contract. The vote indicates whether to extend
// a given contract to a new participant or not
func (api *PrivateExtensionAPI) VoteOnContract(addressToVoteOn common.Address, vote bool, txa ethapi.SendTxArgs) (common.Hash, error) {
	txArgs, err := api.accountManager.GenerateTransactOptions(txa)
	if err != nil {
		return common.Hash{}, err
	}

	voterList, err := api.privacyService.managementContractFacade.GetAllVoters(addressToVoteOn)
	if err != nil {
		return common.Hash{}, err
	}
	if isVoter := checkAddressInList(txArgs.From, voterList); !isVoter {
		return common.Hash{}, errNotVoter
	}

	uuid, err := generateUuid(addressToVoteOn, txArgs.PrivateFrom, api.ptm)
	if err != nil {
		return common.Hash{}, err
	}

	//Find the extension contract in order to interact with it
	extender, err := api.privacyService.managementContractFacade.Transactor(addressToVoteOn)
	if err != nil {
		return common.Hash{}, err
	}

	//Perform the vote transaction.
	tx, err := extender.DoVote(txArgs, vote, uuid)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

// ExtendContract deploys a new extension management contract to the blockchain to start the process of extending
// a contract to a new participant
//Create a new extension contract that signifies that we want to add a new participant to an existing contract
//This should contain:
// - arguments for sending a new transaction (the same as sendTransaction)
// - the contract address we want to extend
// - the new PTM public key
// - the Ethereum addresses of who can vote to extend the contract
func (api *PrivateExtensionAPI) ExtendContract(toExtend common.Address, newRecipientPtmPublicKey string, voters []common.Address, txa ethapi.SendTxArgs) (common.Hash, error) {
	// check the new key is valid
	if _, err := base64.StdEncoding.DecodeString(newRecipientPtmPublicKey); err != nil {
		return common.Hash{}, errors.New("invalid new recipient key provided")
	}

	//generate some valid transaction options for sending in the transaction
	txArgs, err := api.accountManager.GenerateTransactOptions(txa)
	if err != nil {
		return common.Hash{}, err
	}

	// check the the intended new recipient will actually receive the extension request
	found := false
	for _, recipient := range txArgs.PrivateFor {
		if recipient == newRecipientPtmPublicKey {
			found = true
			break
		}
	}
	if !found {
		txArgs.PrivateFor = append(txArgs.PrivateFor, newRecipientPtmPublicKey)
	}

	recipientHash, err := api.ptm.Send([]byte(newRecipientPtmPublicKey), txa.PrivateFrom, []string{})
	if err != nil {
		return common.Hash{}, err
	}

	recipientHashBase64 := common.BytesToEncryptedPayloadHash(recipientHash).ToBase64()

	nonce, err := api.privacyService.extClient.NextNonce(txArgs.From)
	if err != nil {
		return common.Hash{}, err
	}
	txArgs.Nonce = new(big.Int).SetUint64(nonce)
	managementAddress := crypto.CreateAddress(txArgs.From, nonce)

	uuid, err := generateUuid(managementAddress, txArgs.PrivateFrom, api.ptm)
	if err != nil {
		return common.Hash{}, err
	}

	//Deploy the contract
	tx, err := api.privacyService.managementContractFacade.Deploy(txArgs, toExtend, voters, recipientHashBase64, uuid)
	if err != nil {
		return common.Hash{}, err
	}

	//Return the transaction hash for later lookup
	return tx.Hash(), nil
}

// Cancel allows the creator to cancel the given extension contract, ensuring
// that no more calls for votes or accepting can be made
func (api *PrivateExtensionAPI) Cancel(extensionContract common.Address, txa ethapi.SendTxArgs) (common.Hash, error) {
	txArgs, err := api.accountManager.GenerateTransactOptions(txa)
	if err != nil {
		return common.Hash{}, err
	}

	caller, err := api.privacyService.managementContractFacade.Caller(extensionContract)
	if err != nil {
		return common.Hash{}, err
	}
	creatorAddress, err := caller.Creator(nil)
	if err != nil {
		return common.Hash{}, err
	}
	if isCreator := checkAddressInList(txArgs.From, []common.Address{creatorAddress}); !isCreator {
		return common.Hash{}, errNotCreator
	}

	extender, err := api.privacyService.managementContractFacade.Transactor(extensionContract)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := extender.Finish(txArgs)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}