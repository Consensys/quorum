package extension

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
)

var (
	errNotAcceptor = errors.New("account is not acceptor of this extension request")
	errNotCreator  = errors.New("account is not the creator of this extension request")
)

const extensionCompleted = "DONE"
const extensionInProgress = "ACTIVE"

type PrivateExtensionAPI struct {
	privacyService *PrivacyService
}

func NewPrivateExtensionAPI(privacyService *PrivacyService) *PrivateExtensionAPI {
	return &PrivateExtensionAPI{
		privacyService: privacyService,
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

// checks of the passed contract address is under extension process
func (api *PrivateExtensionAPI) checkIfContractUnderExtension(toExtend common.Address) bool {
	for _, v := range api.ActiveExtensionContracts() {
		if v.ContractExtended == toExtend {
			return true
		}
	}
	return false
}

// checks if the voter has already voted on the contract.
func (api *PrivateExtensionAPI) checkAlreadyVoted(addressToVoteOn, from common.Address) bool {
	caller, _ := api.privacyService.managementContractFacade.Caller(addressToVoteOn)
	opts := bind.CallOpts{Pending: true, From: from}

	voted, _ := caller.CheckIfVoted(&opts)
	return voted
}

// checks if the voter has already voted on the contract.
func (api *PrivateExtensionAPI) checkIfExtensionComplete(addressToVoteOn, from common.Address) (bool, error) {
	caller, _ := api.privacyService.managementContractFacade.Caller(addressToVoteOn)
	opts := bind.CallOpts{Pending: true, From: from}

	status, err := caller.CheckIfExtensionFinished(&opts)
	if err != nil {
		return true, err
	}
	return status, nil
}

// checks if the contract being extended is a public contract
func (api *PrivateExtensionAPI) checkIfPublicContract(toExtend common.Address) bool {
	// check if the passed contract is public contract
	publicStateDb, _, _ := api.privacyService.stateFetcher.chainAccessor.State()
	if publicStateDb != nil && publicStateDb.Exist(toExtend) {
		return true
	}
	return false
}

// ApproveContractExtension submits the vote to the specified extension management contract. The vote indicates whether to extend
// a given contract to a new participant or not
func (api *PrivateExtensionAPI) ApproveExtension(addressToVoteOn common.Address, vote bool, txa ethapi.SendTxArgs) (string, error) {
	// check if the extension has been completed. if yes
	// no acceptance required
	status, err := api.checkIfExtensionComplete(addressToVoteOn, txa.From)
	if err != nil {
		return "", err
	}

	if status {
		return "", errors.New("contract extension process complete. nothing to accept")
	}

	if !types.CheckIfAdminAccount(txa.From) {
		return "", errors.New("account cannot accept extension")
	}

	txArgs, err := api.privacyService.GenerateTransactOptions(txa)
	if err != nil {
		return "", err
	}

	voterList, err := api.privacyService.managementContractFacade.GetAllVoters(addressToVoteOn)
	if err != nil {
		return "", err
	}
	if isVoter := checkAddressInList(txArgs.From, voterList); !isVoter {
		return "", errNotAcceptor
	}

	if api.checkAlreadyVoted(addressToVoteOn, txArgs.From) {
		return "", errors.New("already voted")
	}
	uuid, err := generateUuid(addressToVoteOn, txArgs.PrivateFrom, api.privacyService.ptm)
	if err != nil {
		return "", err
	}

	//Find the extension contract in order to interact with it
	extender, err := api.privacyService.managementContractFacade.Transactor(addressToVoteOn)
	if err != nil {
		return "", err
	}

	//Perform the vote transaction.
	tx, err := extender.DoVote(txArgs, vote, uuid)
	if err != nil {
		return "", err
	}
	msg := fmt.Sprintf("0x%x", tx.Hash())
	return msg, nil
}

// ExtendContract deploys a new extension management contract to the blockchain to start the process of extending
// a contract to a new participant
//Create a new extension contract that signifies that we want to add a new participant to an existing contract
//This should contain:
// - arguments for sending a new transaction (the same as sendTransaction)
// - the contract address we want to extend
// - the new PTM public key
// - the Ethereum addresses of who can vote to extend the contract
func (api *PrivateExtensionAPI) ExtendContract(toExtend common.Address, newRecipientPtmPublicKey string, recipientAddr common.Address, txa ethapi.SendTxArgs) (string, error) {

	// check if the contract to be extended is already under extension
	// if yes throw an error
	if api.checkIfContractUnderExtension(toExtend) {
		return "", errors.New("contract extension in progress for the given contract address")
	}

	// check if a public contract is being extended
	if api.checkIfPublicContract(toExtend) {
		return "", errors.New("extending a public contract!!! not allowed")
	}

	// check if recipient address is 0x0
	if recipientAddr == (common.Address{0}) {
		return "", errors.New("invalid recipient address")

	}

	// if running in permissioned mode with new permissions model
	// ensure that the account extending the contract is an admin
	// account and recipient account is an admin account as well
	if txa.From == recipientAddr {
		return "", errors.New("account accepting the extension cannot be the account initiating extension")
	}
	if !types.CheckIfAdminAccount(txa.From) {
		return "", errors.New("account not an org admin account, cannot initiate extension")
	}
	if !types.CheckIfAdminAccount(recipientAddr) {
		return "", errors.New("recipient account address is not an org admin account. cannot accept extension")
	}

	// check the new key is valid
	if _, err := base64.StdEncoding.DecodeString(newRecipientPtmPublicKey); err != nil {
		return "", errors.New("invalid new recipient transaction manager key provided")
	}

	//generate some valid transaction options for sending in the transaction
	txArgs, err := api.privacyService.GenerateTransactOptions(txa)
	if err != nil {
		return "", err
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

	//Deploy the contract
	tx, err := api.privacyService.managementContractFacade.Deploy(txArgs, toExtend, recipientAddr, newRecipientPtmPublicKey)
	if err != nil {
		return "", err
	}

	//Return the transaction hash for later lookup
	msg := fmt.Sprintf("0x%x", tx.Hash())
	return msg, nil
}

// CancelExtension allows the creator to cancel the given extension contract, ensuring
// that no more calls for votes or accepting can be made
func (api *PrivateExtensionAPI) CancelExtension(extensionContract common.Address, txa ethapi.SendTxArgs) (string, error) {
	status, err := api.checkIfExtensionComplete(extensionContract, txa.From)
	if err != nil {
		return "", err
	}
	if status {
		return "", errors.New("contract extension process complete. nothing to cancel")
	}

	txArgs, err := api.privacyService.GenerateTransactOptions(txa)
	if err != nil {
		return "", err
	}

	caller, err := api.privacyService.managementContractFacade.Caller(extensionContract)
	if err != nil {
		return "", err
	}
	creatorAddress, err := caller.Creator(nil)
	if err != nil {
		return "", err
	}
	if isCreator := checkAddressInList(txArgs.From, []common.Address{creatorAddress}); !isCreator {
		return "", errNotCreator
	}

	extender, err := api.privacyService.managementContractFacade.Transactor(extensionContract)
	if err != nil {
		return "", err
	}

	tx, err := extender.Finish(txArgs)
	if err != nil {
		return "", err
	}
	msg := fmt.Sprintf("0x%x", tx.Hash())
	return msg, nil
}

// Returns the extension status from management contract
func (api *PrivateExtensionAPI) GetExtensionStatus(extensionContract common.Address) (string, error) {

	status, err := api.checkIfExtensionComplete(extensionContract, common.Address{})
	if err != nil {
		return "", err
	}

	if status {
		return extensionCompleted, nil
	}

	return extensionInProgress, nil
}
