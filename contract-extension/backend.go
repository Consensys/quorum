package contractExtension

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	extension "github.com/ethereum/go-ethereum/contract-extension/contractExtensionContracts"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"sync"
)

const (
	newExtensionTopic = "0x1bb7909ad96bc757f60de4d9ce11daf7b006e8f398ce028dceb10ce7fdca0f68"
	finishedExtensionTopic = "0x79c47b570b18a8a814b785800e5fcbf104e067663589cef1bba07756e3c6ede9"

	ExtensionContractData = "activeExtensions.json"
)

var (
	//Log queries
	newExtensionQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(newExtensionTopic)}},
		Addresses: []common.Address{},
	}

	finishedExtensionQuery = ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash(finishedExtensionTopic)}},
		Addresses: []common.Address{},
	}

	//default gas limit to use if not passed in sendTxArgs
	defaultGasLimit = uint64(4712384)
	//default gas price to use if not passed in sendTxArgs
	defaultGasPrice = big.NewInt(0)
)

type ExtensionContract struct {
	Address      				common.Address  `json:"address"`
	AllHaveVoted 				bool			`json:"allhavevoted"`
	Initiator					common.Address  `json:"initiator"`
	ManagementContractAddress 	common.Address  `json:"managementcontractaddress"`
	CreationData				[]byte			`json:"creationData"`
	CreatedBlock				uint64			`json:"createdBlock"`

	stopCh						chan struct{}	`json:"-"`
}

type PrivacyService struct {
	ethereum		 *eth.Ethereum
	client			 *ethclient.Client
	ptm				 private.PrivateTransactionManager

	dataDir 		 string

	mu               sync.Mutex
	currentContracts map[common.Address]*ExtensionContract
}

func New(node *node.Node, ptm private.PrivateTransactionManager) (*PrivacyService, error) {
	dataDir := node.InstanceDir()

	service := &PrivacyService{
		currentContracts: make(map[common.Address]*ExtensionContract),
		dataDir:		  dataDir,
		ptm:			  ptm,
	}

	go service.initialise(node)

	return service, nil
}

func (service *PrivacyService) initialise(node *node.Node) {
	service.mu.Lock()
	defer service.mu.Unlock()

	//repopulate existing extensions
	path := filepath.Join(service.dataDir, ExtensionContractData)

	if _, err := os.Stat(path); err == nil || !os.IsNotExist(err) {
		blob, err := ioutil.ReadFile(path)
		if err != nil {
			panic("could not read existing extension contracts")
		}

		if err = json.Unmarshal(blob, &service.currentContracts); err != nil {
			panic("could not unmarshal existing contract file")
		}
	}

	if err := node.Service(&service.ethereum); err != nil {
		panic("extension: could not connect to ethereum service")
	}

	rpcClient, err := node.Attach()
	if err != nil {
		panic("extension: could not connect to ethereum client rpc")
	}

	client := ethclient.NewClient(rpcClient)
	clientAddress := os.Getenv("CONTRACT_EXTENSION_SERVER")
	if service.client, err = client.WithPrivateTransactionManager(clientAddress); err != nil {
		panic("could not set PTM")
	}

	for _, item := range service.currentContracts {
		extensionEntry := item
		extensionEntry.stopCh = make(chan struct{})
		go service.watchForVoteCompleteEvents(extensionEntry.ManagementContractAddress, extensionEntry.CreatedBlock, extensionEntry.stopCh)
	}

	go service.watchForNewContracts()
	go service.watchForCancelledContracts()
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

			a := new(extension.ContractExtenderNewContractExtensionContractCreated)
			if err := extension.ContractExtensionABI.Unpack(a, "NewContractExtensionContractCreated", foundLog.Data); err != nil {
				log.Error("Error unpacking extension creation log", err.Error())
				log.Debug("Errored log", foundLog)
				service.mu.Unlock()
				continue
			}

			newContractExtension := ExtensionContract{
				Address:      				a.ToExtend,
				AllHaveVoted: 				false,
				Initiator:	  				from,
				ManagementContractAddress: 	foundLog.Address,
				CreationData: 				tx.Data(),
				CreatedBlock:				foundLog.BlockNumber,
				stopCh:						make(chan struct{}),
			}

			service.currentContracts[foundLog.Address] = &newContractExtension
			go service.watchForVoteCompleteEvents(foundLog.Address, foundLog.BlockNumber, newContractExtension.stopCh)
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
			if entry, ok := service.currentContracts[l.Address]; ok {
				close(entry.stopCh)
				delete(service.currentContracts, l.Address)
				writeContentsToFile(service.currentContracts, service.dataDir)
			}
			service.mu.Unlock()
		}
	}

}

func (service *PrivacyService) watchForVoteCompleteEvents(address common.Address, startBlock uint64, stopCh <-chan struct{}) {
	logSink := make(chan *extension.ContractExtenderAllNodesHaveVoted)
	filterer, _ := extension.NewContractExtenderFilterer(address, service.client)
	subscription, _ := filterer.WatchAllNodesHaveVoted(&bind.WatchOpts{Start: &startBlock}, logSink)

	select {
	case err := <-subscription.Err():
		log.Error("Contract extension watcher subscription error", err)
	case <-stopCh:
		log.Info("No longer watching extension request", "address", address)
	case event := <-logSink:
		service.mu.Lock()
		defer service.mu.Unlock()

		if !event.Outcome {
			return
		}

		extensionEntry := service.currentContracts[address]
		extensionEntry.AllHaveVoted = true
		writeContentsToFile(service.currentContracts, service.dataDir)

		from := accounts.Account{Address: extensionEntry.Initiator}
		if _, err := service.ethereum.AccountManager().Find(from); err != nil {
			return
		}

		//fetch all the participants and send
		payload := common.BytesToEncryptedPayloadHash(extensionEntry.CreationData)
		fetchedParties, err := service.ptm.GetParticipants(payload)
		if err != nil {
			log.Error("Extension", "Unable to fetch parties for PSV extension")
			return
		}
		txArgs, _ := service.generateTransactOpts(ethapi.SendTxArgs{
			From:          extensionEntry.Initiator,
			PrivateFor:    fetchedParties,
		})

		//Find the extension contract in order to interact with it
		caller, _ := extension.NewContractExtenderCaller(event.Raw.Address, service.client)

		recipientHash, _ := caller.TargetRecipientPublicKeyHash(&bind.CallOpts{Pending: false})
		decoded, _ := base64.StdEncoding.DecodeString(recipientHash)
		recipient, _ := service.ptm.Receive(decoded)

		//we found the account, so we can send
		privateState, _ := service.privateState(event.Raw.BlockHash)
		jsonMap := getAddressState(privateState, extensionEntry.Address)

		//send to PTM
		hash, _ := service.ptm.Send(jsonMap, "", []string{string(recipient)})
		hashB64 := base64.StdEncoding.EncodeToString(hash)

		transactor, _ := extension.NewContractExtenderTransactor(event.Raw.Address, service.client)
		transactor.SetSharedStateHash(txArgs, hashB64)
	}
}

func (service *PrivacyService) generateTransactOpts(txa ethapi.SendTxArgs) (*bind.TransactOpts, error) {
	return generateTransactOpts(service.ethereum.AccountManager(), txa)
}

func (service *PrivacyService) privateState(blockHash common.Hash) (*state.StateDB, error) {
	db := service.ethereum.ChainDb()
	block := service.ethereum.BlockChain().GetBlockByHash(blockHash)

	privateStateRoot := core.GetPrivateStateRoot(db, block.Root())
	return state.New(privateStateRoot, state.NewDatabase(db))
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
