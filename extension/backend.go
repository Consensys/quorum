package extension

import (
	"encoding/base64"
	"errors"
	"sync"

	"github.com/ethereum/go-ethereum/extension/extensionContracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/rpc"
)

type PrivacyService struct {
	ptm private.PrivateTransactionManager

	stateFetcher             *StateFetcher
	accountManager           IAccountManager
	dataHandler              DataHandler
	managementContractFacade ManagementContractFacade
	extClient                Client

	mu               sync.Mutex
	currentContracts map[common.Address]*ExtensionContract
}

func New(ptm private.PrivateTransactionManager, manager IAccountManager, handler DataHandler, fetcher *StateFetcher) (*PrivacyService, error) {
	service := &PrivacyService{
		currentContracts: make(map[common.Address]*ExtensionContract),
		ptm:              ptm,
		dataHandler:      handler,
		stateFetcher:     fetcher,
		accountManager:   manager,
	}

	var err error
	service.currentContracts, err = service.dataHandler.Load()
	if err != nil {
		return nil, errors.New("could not load existing extension contracts: " + err.Error())
	}

	return service, nil
}

func (service *PrivacyService) initialise(node *node.Node, thirdpartyunixfile string) {
	service.mu.Lock()
	defer service.mu.Unlock()

	rpcClient, err := node.Attach()
	if err != nil {
		panic("extension: could not connect to ethereum client rpc")
	}

	client, _ := ethclient.NewClient(rpcClient).WithIPCPrivateTransactionManager(thirdpartyunixfile)
	service.managementContractFacade = NewManagementContractFacade(client)
	service.extClient = NewInProcessClient(client)

	go service.watchForNewContracts()
	go service.watchForCancelledContracts()
	go service.watchForCompletionEvents()
	go service.watchForNewVotes()
}

func (service *PrivacyService) watchForNewContracts() {
	incomingLogs, subscription, _ := service.extClient.SubscribeToLogs(newExtensionQuery)

	for {
		select {
		case err := <-subscription.Err():
			log.Error("Contract extension watcher subscription error", err)
			break
		case foundLog := <-incomingLogs:
			service.mu.Lock()

			tx, _ := service.extClient.TransactionByHash(foundLog.TxHash)
			from, _ := types.QuorumPrivateTxSigner{}.Sender(tx)

			newExtensionEvent, err := extensionContracts.UnpackNewExtensionCreatedLog(foundLog.Data)
			if err != nil {
				log.Error("Error unpacking extension creation log", err.Error())
				log.Debug("Errored log", foundLog)
				service.mu.Unlock()
				continue
			}

			newContractExtension := ExtensionContract{
				Address:                   newExtensionEvent.ToExtend,
				HasVoted:                  false,
				Initiator:                 from,
				ManagementContractAddress: foundLog.Address,
				CreationData:              tx.Data(),
			}

			service.currentContracts[foundLog.Address] = &newContractExtension
			err = service.dataHandler.Save(service.currentContracts)
			if err != nil {
				log.Error("Error writing extension data to file", err.Error())
				service.mu.Unlock()
				continue
			}
			service.mu.Unlock()

			// if party is sender then complete self voting
			data := common.BytesToEncryptedPayloadHash(newContractExtension.CreationData)
			isSender, _ := service.ptm.IsSender(data)

			if isSender {
				fetchedParties, err := service.ptm.GetParticipants(data)
				if err != nil {
					log.Error("Extension", "Unable to fetch all parties for extension management contract")
					continue
				}
				//Find the extension contract in order to interact with it
				caller, _ := service.managementContractFacade.Caller(newContractExtension.ManagementContractAddress)
				contractCreator, _ := caller.Creator(nil)

				txArgs := ethapi.SendTxArgs{From: contractCreator, PrivateFor: fetchedParties}

				extensionAPI := NewPrivateExtensionAPI(service, service.accountManager, service.ptm)
				_, err = extensionAPI.VoteOnContract(newContractExtension.ManagementContractAddress, true, txArgs)

				if err != nil {
					log.Error("Extension","Unable initiator vote on management contract failed" )
				}

			}
		}
	}
}

func (service *PrivacyService) watchForNewVotes() {
	incomingLogs, _, _ := service.extClient.SubscribeToLogs(newVoteQuery)

	for {
		select {
		case l := <-incomingLogs:
			service.mu.Lock()

			managementContract, ok := service.currentContracts[l.Address]
			if !ok {
				// we didn't have this management contract, so ignore it
				service.mu.Unlock()
				continue
			}

			tx, err := service.extClient.TransactionInBlock(l.BlockHash, l.TxIndex)
			if err != nil {
				service.mu.Unlock()
				continue
			}

			data := common.BytesToEncryptedPayloadHash(tx.Data())
			isSender, err := service.ptm.IsSender(data)
			if err != nil {
				log.Warn("couldn't determine if sender of private transaction", "tx", tx.Hash().String(), "err", err.Error())
			}
			managementContract.HasVoted = isSender
			service.mu.Unlock()
		}
	}
}

func (service *PrivacyService) watchForCancelledContracts() {
	incomingLogs, subscription, _ := service.extClient.SubscribeToLogs(finishedExtensionQuery)

	for {
		select {
		case err := <-subscription.Err():
			log.Error("Contract cancellation extension watcher subscription error", err)
			return
		case l := <-incomingLogs:
			service.mu.Lock()
			if _, ok := service.currentContracts[l.Address]; ok {
				delete(service.currentContracts, l.Address)
				service.dataHandler.Save(service.currentContracts)
			}
			service.mu.Unlock()
		}
	}
}

func (service *PrivacyService) watchForCompletionEvents() {
	incomingLogs, _, _ := service.extClient.SubscribeToLogs(canPerformStateShareQuery)

	for {
		select {
		case l := <-incomingLogs:
			service.mu.Lock()
			extensionEntry, ok := service.currentContracts[l.Address]
			if !ok {
				// we didn't have this management contract, so ignore it
				service.mu.Unlock()
				continue
			}

			//Find the extension contract in order to interact with it
			caller, _ := service.managementContractFacade.Caller(l.Address)
			contractCreator, _ := caller.Creator(nil)

			if !service.accountManager.Exists(contractCreator) {
				log.Warn("Account used to sign extension contract no longer available", "account", contractCreator.Hex())
				service.mu.Unlock()
				continue
			}

			//fetch all the participants and send
			payload := common.BytesToEncryptedPayloadHash(extensionEntry.CreationData)
			fetchedParties, err := service.ptm.GetParticipants(payload)
			if err != nil {
				log.Error("Extension", "Unable to fetch all parties for extension management contract")
				service.mu.Unlock()
				continue
			}

			txArgs, _ := service.accountManager.GenerateTransactOptions(ethapi.SendTxArgs{From: contractCreator, PrivateFor: fetchedParties})

			recipientHash, _ := caller.TargetRecipientPublicKeyHash(&bind.CallOpts{Pending: false})
			decoded, _ := base64.StdEncoding.DecodeString(recipientHash)
			recipient, _ := service.ptm.Receive(decoded)

			//we found the account, so we can send
			contractToExtend, _ := caller.ContractToExtend(nil)
			entireStateData, _ := service.stateFetcher.GetAddressStateFromBlock(l.BlockHash, contractToExtend)

			//send to PTM
			hashOfStateData, _ := service.ptm.Send(entireStateData, "", []string{string(recipient)})
			hashofStateDataBase64 := base64.StdEncoding.EncodeToString(hashOfStateData)

			transactor, _ := service.managementContractFacade.Transactor(l.Address)
			transactor.SetSharedStateHash(txArgs, hashofStateDataBase64)
			service.mu.Unlock()
		}
	}
}

// node.Service interface methods:
func (service *PrivacyService) Protocols() []p2p.Protocol {
	return []p2p.Protocol{}
}

func (service *PrivacyService) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "quorumExtension",
			Version:   "1.0",
			Service:   NewPrivateExtensionAPI(service, service.accountManager, service.ptm),
			Public:    true,
		},
	}
}

func (service *PrivacyService) Start(p2pServer *p2p.Server) error {
	return nil
}

func (service *PrivacyService) Stop() error {
	return nil
}
