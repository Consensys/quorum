package extension

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/extension/extensionContracts"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/rpc"
)

const extensionContractData = "activeExtensions.json"

var (
	//Log queries
	newExtensionQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(extensionContracts.NewContractExtensionContractCreatedTopicHash)}},
		Addresses: []common.Address{},
	}

	finishedExtensionQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(extensionContracts.ExtensionFinishedTopicHash)}},
		Addresses: []common.Address{},
	}

	voteCompletedQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(extensionContracts.AllNodesHaveVotedTopicHash)}},
		Addresses: []common.Address{},
	}

	canPerformStateShareQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(extensionContracts.CanPerformStateShareTopicHash)}},
		Addresses: []common.Address{},
	}
)

type ExtensionContract struct {
	Address                   common.Address `json:"address"`
	AllHaveVoted              bool           `json:"allhavevoted"`
	Initiator                 common.Address `json:"initiator"`
	ManagementContractAddress common.Address `json:"managementcontractaddress"`
	CreationData              []byte         `json:"creationData"`
}

type PrivacyService struct {
	client   *ethclient.Client
	ptm      private.PrivateTransactionManager

	stateFetcher *StateFetcher
	accountManager *AccountManager

	dataDir string

	mu               sync.Mutex
	currentContracts map[common.Address]*ExtensionContract
}

func New(node *node.Node, ptm private.PrivateTransactionManager, thirdpartyunixfile string) (*PrivacyService, error) {
	dataDir := node.InstanceDir()

	service := &PrivacyService{
		currentContracts: make(map[common.Address]*ExtensionContract),
		dataDir:          dataDir,
		ptm:              ptm,
	}

	go service.initialise(node, thirdpartyunixfile)

	return service, nil
}

func (service *PrivacyService) initialise(node *node.Node, thirdpartyunixfile string) {
	service.mu.Lock()
	defer service.mu.Unlock()

	//repopulate existing extensions
	path := filepath.Join(service.dataDir, extensionContractData)

	if _, err := os.Stat(path); err == nil || !os.IsNotExist(err) {
		blob, err := ioutil.ReadFile(path)
		if err != nil {
			panic("could not read existing extension contracts")
		}

		if err = json.Unmarshal(blob, &service.currentContracts); err != nil {
			panic("could not unmarshal existing contract file")
		}
	}

	var ethService *eth.Ethereum
	if err := node.Service(&ethService); err != nil {
		panic("extension: could not connect to ethereum service")
	}

	service.stateFetcher = NewStateFetcher(ethService.ChainDb(), ethService.BlockChain())
	service.accountManager = NewAccountManager(ethService.AccountManager())

	rpcClient, err := node.Attach()
	if err != nil {
		panic("extension: could not connect to ethereum client rpc")
	}

	client := ethclient.NewClient(rpcClient)
	if service.client, err = client.WithIPCPrivateTransactionManager(thirdpartyunixfile); err != nil {
		panic("could not set PTM")
	}

	go service.watchForNewContracts()
	go service.watchForCancelledContracts()
	go service.watchForCompletionEvents()
	go service.watchForVotingCompletedContracts()
}

func (service *PrivacyService) watchForNewContracts() {
	incomingLogs := make(chan types.Log)
	subscription, _ := service.client.SubscribeFilterLogs(context.Background(), newExtensionQuery, incomingLogs)

	for {
		select {
		case err := <-subscription.Err():
			log.Error("Contract extension watcher subscription error", err)
			break
		case foundLog := <-incomingLogs:
			service.mu.Lock()

			tx, _, _ := service.client.TransactionByHash(context.Background(), foundLog.TxHash)
			from, _ := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)

			newExtensionEvent, err := unpackNewExtension(foundLog.Data)
			if err != nil {
				log.Error("Error unpacking extension creation log", err.Error())
				log.Debug("Errored log", foundLog)
				service.mu.Unlock()
				continue
			}

			newContractExtension := ExtensionContract{
				Address:                   newExtensionEvent.ToExtend,
				AllHaveVoted:              false,
				Initiator:                 from,
				ManagementContractAddress: foundLog.Address,
				CreationData:              tx.Data(),
			}

			service.currentContracts[foundLog.Address] = &newContractExtension
			writeContentsToFile(service.currentContracts, service.dataDir)
			service.mu.Unlock()
		}
	}
}

func (service *PrivacyService) watchForCancelledContracts() {
	logsChan := make(chan types.Log)
	subscription, _ := service.client.SubscribeFilterLogs(context.Background(), finishedExtensionQuery, logsChan)

	for {
		select {
		case err := <-subscription.Err():
			log.Error("Contract cancellation extension watcher subscription error", err)
			return
		case l := <-logsChan:
			service.mu.Lock()
			if _, ok := service.currentContracts[l.Address]; ok {
				delete(service.currentContracts, l.Address)
				writeContentsToFile(service.currentContracts, service.dataDir)
			}
			service.mu.Unlock()
		}
	}
}

func (service *PrivacyService) watchForVotingCompletedContracts() {
	logsChan := make(chan types.Log)
	service.client.SubscribeFilterLogs(context.Background(), voteCompletedQuery, logsChan)

	for {
		select {
		case l := <-logsChan:
			service.mu.Lock()
			extensionEntry, ok := service.currentContracts[l.Address]
			if !ok {
				// we didn't have this management contract, so ignore it
				service.mu.Unlock()
				continue
			}
			// we aren't that bothered about the case where someone declines and emits this event
			// because it will immediately be deleted from the API
			extensionEntry.AllHaveVoted = true
			writeContentsToFile(service.currentContracts, service.dataDir)
			service.mu.Unlock()
		}
	}
}


func (service *PrivacyService) watchForCompletionEvents() {
	logsChan := make(chan types.Log)
	service.client.SubscribeFilterLogs(context.Background(), canPerformStateShareQuery, logsChan)

	for {
		select {
		case l := <-logsChan:
			service.mu.Lock()
			extensionEntry, ok := service.currentContracts[l.Address]
			if !ok {
				// we didn't have this management contract, so ignore it
				service.mu.Unlock()
				continue
			}

			//Find the extension contract in order to interact with it
			caller, _ := extensionContracts.NewContractExtenderCaller(l.Address, service.client)
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
				log.Error("Extension", "Unable to fetch parties for PSV extension")
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

			transactor, _ := extensionContracts.NewContractExtenderTransactor(l.Address, service.client)
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
			Service:   NewPrivateExtensionAPI(service),
			Public:    true,
		},
	}
}

func (service *PrivacyService) Start(p2pServer *p2p.Server) error {
	return nil
}

func (service *PrivacyService) Stop() error {
	service.mu.Lock()
	defer service.mu.Unlock()
	return writeContentsToFile(service.currentContracts, service.dataDir)
}
