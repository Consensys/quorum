package extension

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/multitenancy"
	"github.com/ethereum/go-ethereum/permission/core"
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
func (api *PrivateExtensionAPI) ActiveExtensionContracts(ctx context.Context) []ExtensionContract {
	api.privacyService.mu.Lock()
	defer api.privacyService.mu.Unlock()

	psi, err := api.privacyService.apiBackendHelper.PSMR().ResolveForUserContext(ctx)
	if err != nil {
		return nil
	}

	extracted := make([]ExtensionContract, 0)
	for _, contract := range api.privacyService.psiContracts[psi.ID] {
		extracted = append(extracted, *contract)
	}

	return extracted
}

// checks of the passed contract address is under extension process
func (api *PrivateExtensionAPI) checkIfContractUnderExtension(ctx context.Context, toExtend common.Address) bool {
	for _, v := range api.ActiveExtensionContracts(ctx) {
		if v.ContractExtended == toExtend {
			return true
		}
	}
	return false
}

// checks if the voter has already voted on the contract.
func (api *PrivateExtensionAPI) checkAlreadyVoted(addressToVoteOn, from common.Address, psi types.PrivateStateIdentifier) bool {
	psiManagementContractClient := api.privacyService.managementContract(psi)
	defer psiManagementContractClient.Close()
	caller, _ := psiManagementContractClient.Caller(addressToVoteOn)
	opts := bind.CallOpts{Pending: true, From: from}

	voted, _ := caller.CheckIfVoted(&opts)
	return voted
}

// checks if the contract extension is completed
func (api *PrivateExtensionAPI) checkIfExtensionComplete(addressToVoteOn, from common.Address, psi types.PrivateStateIdentifier) (bool, error) {
	psiManagementContractClient := api.privacyService.managementContract(psi)
	defer psiManagementContractClient.Close()
	caller, _ := psiManagementContractClient.Caller(addressToVoteOn)
	opts := bind.CallOpts{Pending: true, From: from}

	status, err := caller.CheckIfExtensionFinished(&opts)
	if err != nil {
		return true, err
	}
	return status, nil
}

// checks if the contract being extended is a public contract
func (api *PrivateExtensionAPI) checkIfPublicContract(toExtend common.Address) (bool, error) {
	// check if the passed contract is public contract
	chain := api.privacyService.stateFetcher.chainAccessor
	publicStateDb, _, err := chain.StateAtPSI(chain.CurrentBlock().Root(), types.DefaultPrivateStateIdentifier)
	if err != nil {
		return false, err
	}
	return publicStateDb != nil && publicStateDb.Exist(toExtend), nil
}

// checks if the contract being extended is available on the node
func (api *PrivateExtensionAPI) checkIfPrivateStateExists(psi types.PrivateStateIdentifier, toExtend common.Address) (bool, error) {
	// check if the private contract exists on the node extending the contract
	chain := api.privacyService.stateFetcher.chainAccessor
	_, privateStateDb, err := chain.StateAtPSI(chain.CurrentBlock().Root(), psi)
	if err != nil {
		return false, err
	}
	return privateStateDb != nil && privateStateDb.GetCode(toExtend) != nil, nil
}

func (api *PrivateExtensionAPI) doMultiTenantChecks(ctx context.Context, address common.Address, txa ethapi.SendTxArgs) error {
	backendHelper := api.privacyService.apiBackendHelper
	if token, ok := backendHelper.SupportsMultitenancy(ctx); ok {
		psm, err := backendHelper.PSMR().ResolveForUserContext(ctx)
		if err != nil {
			return err
		}
		eoaSecAttr := (&multitenancy.PrivateStateSecurityAttribute{}).WithPSI(psm.ID).WithNodeEOA(address)
		psm, err = backendHelper.PSMR().ResolveForManagedParty(txa.PrivateFrom)
		if err != nil {
			return err
		}
		privateFromSecAttr := (&multitenancy.PrivateStateSecurityAttribute{}).WithPSI(psm.ID).WithNodeEOA(address)
		if isAuthorized, _ := multitenancy.IsAuthorized(token, eoaSecAttr, privateFromSecAttr); !isAuthorized {
			return multitenancy.ErrNotAuthorized
		}
	}
	return nil
}

// GenerateExtensionApprovalUuid generates a uuid to be used for contract state extension approval when calling doVote within the management contract,
// allowing the approval method to be called with an external signer
func (api *PrivateExtensionAPI) GenerateExtensionApprovalUuid(ctx context.Context, addressToVoteOn common.Address, txa ethapi.SendTxArgs) (string, error) {
	err := api.doMultiTenantChecks(ctx, txa.From, txa)
	if err != nil {
		return "", err
	}

	psm, err := api.privacyService.apiBackendHelper.PSMR().ResolveForUserContext(ctx)
	if err != nil {
		return "", err
	}
	psi := psm.ID

	// check if the extension has been completed. if yes
	// no acceptance required
	status, err := api.checkIfExtensionComplete(addressToVoteOn, txa.From, psi)
	if err != nil {
		return "", err
	}

	if status {
		return "", errors.New("contract extension process complete. nothing to accept")
	}

	if !core.CheckIfAdminAccount(txa.From) {
		return "", errors.New("account cannot accept extension")
	}

	// get all participants for the contract being extended
	participants, err := api.privacyService.GetAllParticipants(api.privacyService.stateFetcher.getCurrentBlockHash(), addressToVoteOn, psi)
	if err == nil {
		txa.PrivateFor = append(txa.PrivateFor, participants...)
	}

	txArgs, err := api.privacyService.GenerateTransactOptions(txa)
	if err != nil {
		return "", err
	}

	psiManagementContractClient := api.privacyService.managementContract(psi)
	defer psiManagementContractClient.Close()
	voterList, err := psiManagementContractClient.GetAllVoters(addressToVoteOn)
	if err != nil {
		return "", err
	}
	if isVoter := checkAddressInList(txArgs.From, voterList); !isVoter {
		return "", errNotAcceptor
	}

	if api.checkAlreadyVoted(addressToVoteOn, txArgs.From, psi) {
		return "", errors.New("already voted")
	}
	uuid, err := generateUuid(addressToVoteOn, txArgs.PrivateFrom, txArgs.PrivateFor, api.privacyService.ptm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

// ApproveContractExtension submits the vote to the specified extension management contract. The vote indicates whether to extend
// a given contract to a new participant or not
func (api *PrivateExtensionAPI) ApproveExtension(ctx context.Context, addressToVoteOn common.Address, vote bool, txa ethapi.SendTxArgs) (string, error) {
	err := api.doMultiTenantChecks(ctx, txa.From, txa)
	if err != nil {
		return "", err
	}

	psm, err := api.privacyService.apiBackendHelper.PSMR().ResolveForUserContext(ctx)
	if err != nil {
		return "", err
	}
	psi := psm.ID

	// check if the extension has been completed. if yes
	// no acceptance required
	status, err := api.checkIfExtensionComplete(addressToVoteOn, txa.From, psi)
	if err != nil {
		return "", err
	}

	if status {
		return "", errors.New("contract extension process complete. nothing to accept")
	}

	if !core.CheckIfAdminAccount(txa.From) {
		return "", errors.New("account cannot accept extension")
	}

	// get all participants for the contract being extended
	participants, err := api.privacyService.GetAllParticipants(api.privacyService.stateFetcher.getCurrentBlockHash(), addressToVoteOn, psi)
	if err == nil {
		txa.PrivateFor = append(txa.PrivateFor, participants...)
	}

	txArgs, err := api.privacyService.GenerateTransactOptions(txa)
	if err != nil {
		return "", err
	}

	psiManagementContractClient := api.privacyService.managementContract(psi)
	defer psiManagementContractClient.Close()
	voterList, err := psiManagementContractClient.GetAllVoters(addressToVoteOn)
	if err != nil {
		return "", err
	}
	if isVoter := checkAddressInList(txArgs.From, voterList); !isVoter {
		return "", errNotAcceptor
	}

	if api.checkAlreadyVoted(addressToVoteOn, txArgs.From, psi) {
		return "", errors.New("already voted")
	}
	uuid, err := generateUuid(addressToVoteOn, txArgs.PrivateFrom, txArgs.PrivateFor, api.privacyService.ptm)
	if err != nil {
		return "", err
	}

	//Find the extension contract in order to interact with it
	extender, err := psiManagementContractClient.Transactor(addressToVoteOn)
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
// Create a new extension contract that signifies that we want to add a new participant to an existing contract
// This should contain:
// - arguments for sending a new transaction (the same as sendTransaction)
// - the contract address we want to extend
// - the new PTM public key
// - the Ethereum addresses of who can vote to extend the contract
func (api *PrivateExtensionAPI) ExtendContract(ctx context.Context, toExtend common.Address, newRecipientPtmPublicKey string, recipientAddr common.Address, txa ethapi.SendTxArgs) (string, error) {
	// check if the contract to be extended is already under extension
	// if yes throw an error
	if api.checkIfContractUnderExtension(ctx, toExtend) {
		return "", errors.New("contract extension in progress for the given contract address")
	}

	// check if a public contract is being extended
	isPublic, err := api.checkIfPublicContract(toExtend)
	if err != nil {
		return "", err
	}
	if isPublic {
		return "", errors.New("extending a public contract!!! not allowed")
	}

	err = api.doMultiTenantChecks(ctx, txa.From, txa)
	if err != nil {
		return "", err
	}

	// check if recipient address is 0x0
	if recipientAddr == (common.Address{0}) {
		return "", errors.New("invalid recipient address")
	}

	psm, err := api.privacyService.apiBackendHelper.PSMR().ResolveForUserContext(ctx)
	if err != nil {
		return "", err
	}

	// check if a private contract exists
	privateContractExists, err := api.checkIfPrivateStateExists(psm.ID, toExtend)
	if err != nil {
		return "", err
	}
	if !privateContractExists {
		return "", errors.New("extending a non-existent private contract!!! not allowed")
	}

	// check if contract creator
	if !api.privacyService.CheckIfContractCreator(api.privacyService.stateFetcher.getCurrentBlockHash(), toExtend, psm.ID) {
		return "", errors.New("operation not allowed")
	}

	// if running in permissioned mode with new permissions model
	// ensure that the account extending the contract is an admin
	// account and recipient account is an admin account as well
	if txa.From == recipientAddr {
		return "", errors.New("account accepting the extension cannot be the account initiating extension")
	}
	if !core.CheckIfAdminAccount(txa.From) {
		return "", errors.New("account not an org admin account, cannot initiate extension")
	}
	if !core.CheckIfAdminAccount(recipientAddr) {
		return "", errors.New("recipient account address is not an org admin account. cannot accept extension")
	}

	// check the new key is valid
	if _, err := base64.StdEncoding.DecodeString(newRecipientPtmPublicKey); err != nil {
		return "", errors.New("invalid new recipient transaction manager key provided")
	}

	// check the the intended new recipient will actually receive the extension request
	switch len(txa.PrivateFor) {
	case 0:
		txa.PrivateFor = append(txa.PrivateFor, newRecipientPtmPublicKey)
	case 1:
		if txa.PrivateFor[0] != newRecipientPtmPublicKey {
			return "", errors.New("mismatch between recipient transaction manager key and privateFor argument")
		}
	default:
		return "", errors.New("invalid transaction manager keys given in privateFor argument")
	}

	// get all participants for the contract being extended
	participants, err := api.privacyService.GetAllParticipants(api.privacyService.stateFetcher.getCurrentBlockHash(), toExtend, psm.ID)
	if err == nil {
		txa.PrivateFor = append(txa.PrivateFor, participants...)
	}

	//generate some valid transaction options for sending in the transaction
	txArgs, err := api.privacyService.GenerateTransactOptions(txa)
	if err != nil {
		return "", err
	}

	psiManagementContractClient := api.privacyService.managementContract(psm.ID)
	defer psiManagementContractClient.Close()
	//Deploy the contract
	tx, err := psiManagementContractClient.Deploy(txArgs, toExtend, recipientAddr, newRecipientPtmPublicKey)
	if err != nil {
		return "", err
	}

	//Return the transaction hash for later lookup
	msg := fmt.Sprintf("0x%x", tx.Hash())
	return msg, nil
}

// CancelExtension allows the creator to cancel the given extension contract, ensuring
// that no more calls for votes or accepting can be made
func (api *PrivateExtensionAPI) CancelExtension(ctx context.Context, extensionContract common.Address, txa ethapi.SendTxArgs) (string, error) {
	err := api.doMultiTenantChecks(ctx, txa.From, txa)
	if err != nil {
		return "", err
	}

	psm, err := api.privacyService.apiBackendHelper.PSMR().ResolveForUserContext(ctx)
	if err != nil {
		return "", err
	}
	// get all participants for the contract being extended
	status, err := api.checkIfExtensionComplete(extensionContract, txa.From, psm.ID)
	if err != nil {
		return "", err
	}
	if status {
		return "", errors.New("contract extension process complete. nothing to cancel")
	}

	participants, err := api.privacyService.GetAllParticipants(api.privacyService.stateFetcher.getCurrentBlockHash(), extensionContract, psm.ID)
	if err == nil {
		txa.PrivateFor = append(txa.PrivateFor, participants...)
	}

	txArgs, err := api.privacyService.GenerateTransactOptions(txa)
	if err != nil {
		return "", err
	}

	psiManagementContractClient := api.privacyService.managementContract(psm.ID)
	defer psiManagementContractClient.Close()
	caller, err := psiManagementContractClient.Caller(extensionContract)
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

	extender, err := psiManagementContractClient.Transactor(extensionContract)
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
func (api *PrivateExtensionAPI) GetExtensionStatus(ctx context.Context, extensionContract common.Address) (string, error) {
	psm, err := api.privacyService.apiBackendHelper.PSMR().ResolveForUserContext(ctx)
	if err != nil {
		return "", err
	}
	status, err := api.checkIfExtensionComplete(extensionContract, common.Address{}, psm.ID)
	if err != nil {
		return "", err
	}

	if status {
		return extensionCompleted, nil
	}

	return extensionInProgress, nil
}
