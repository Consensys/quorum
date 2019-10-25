package contractExtension

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	extension "github.com/ethereum/go-ethereum/contract-extension/contractExtensionContracts"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/private"
)

//this exists to send to the PTM to be encrypted, returning a non-deterministic hash
var ptmMessage = []byte("extension-data")

type PrivateExtensionAPI struct {
	privacyService 	*PrivacyService
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

	var extracted []ExtensionContract
	for _, contract := range api.privacyService.currentContracts {
		extracted = append(extracted, *contract)
	}
	return extracted
}

// VoteOnContract submits the vote to the specified extension management contract. The vote indicates whether to extend
// a given contract to a new participant or not
func (api *PrivateExtensionAPI) VoteOnContract(addressToVoteOn common.Address, vote bool, txa ethapi.SendTxArgs) (common.Hash, error) {
	txArgs, err := api.privacyService.generateTransactOpts(txa)
	if err != nil {
		return common.Hash{}, err
	}

	//Find the extension contract in order to interact with it
	extender, err := extension.NewContractExtenderTransactor(addressToVoteOn, api.privacyService.client)
	if err != nil {
		return common.Hash{}, err
	}

	//Perform the vote transaction.
	tx, err := extender.DoVote(txArgs, vote)
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
func (api *PrivateExtensionAPI) ExtendContract(ctx context.Context, toExtend common.Address, newRecipient string, voters []common.Address, txa ethapi.SendTxArgs) (common.Hash, error) {
	txArgs, err := api.privacyService.generateTransactOpts(txa)
	if err != nil {
		return common.Hash{}, err
	}

	recipientHash, err := api.privacyService.ptm.Send([]byte(newRecipient), txa.PrivateFrom, []string{})
	if err != nil {
		return common.Hash{}, err
	}

	recipientHashBase64 := common.BytesToEncryptedPayloadHash(recipientHash).ToBase64()

	//Deploy the contract
	_, tx, _, err := extension.DeployContractExtender(txArgs, api.privacyService.client, toExtend, voters, recipientHashBase64)
	if err != nil {
		return common.Hash{}, err
	}

	//Return the address of the deployed contract
	return tx.Hash(), nil
}

// Accept
func (api *PrivateExtensionAPI) Accept(ctx context.Context, addressToVoteOn common.Address, txa ethapi.SendTxArgs) (common.Hash, error){
	txArgs, err := api.privacyService.generateTransactOpts(txa)
	if err != nil {
		return common.Hash{}, err
	}

	uuid, err := generateUuid(txArgs.PrivateFrom, api.privacyService.ptm)
	if err != nil {
		return common.Hash{}, err
	}

	//Find the extension contract in order to interact with it
	extender, err := extension.NewContractExtenderTransactor(addressToVoteOn, api.privacyService.client)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := extender.ShareAcceptStatus(txArgs, uuid)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (api *PrivateExtensionAPI) setUuid(addressToVoteOn common.Address, txArgs *bind.TransactOpts) (common.Hash, error) {
	//Find the extension contract in order to interact with it
	extender, err := extension.NewContractExtenderTransactor(addressToVoteOn, api.privacyService.client)
	if err != nil {
		return common.Hash{}, err
	}

	uuid, err := generateUuid(txArgs.PrivateFrom, api.privacyService.ptm)
	if err != nil {
		return common.Hash{}, err
	}

	//Perform the vote transaction.
	tx, err := extender.SetUuid(txArgs, uuid)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (api *PrivateExtensionAPI) Cancel(extensionContract common.Address, txa ethapi.SendTxArgs) (common.Hash, error) {
	txArgs, err := api.privacyService.generateTransactOpts(txa)
	if err != nil {
		return common.Hash{}, err
	}

	extender, err := extension.NewContractExtenderTransactor(extensionContract, api.privacyService.client)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := extender.Finish(txArgs)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func generateUuid(privateFrom string, ptm private.PrivateTransactionManager) (string, error) {
	hash, err := ptm.Send(ptmMessage, privateFrom, []string{})
	if err != nil {
		return "", err
	}
	return common.BytesToEncryptedPayloadHash(hash).String(), nil
}