package extension

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/extension/privacyExtension"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/private"
)

type ServicesFactory interface {
	BackendService() *PrivacyService

	AccountManager() *accounts.Manager
	DataHandler() DataHandler
	StateFetcher() *StateFetcher
}

type DefaultServicesFactory struct {
	backendService *PrivacyService
	accountManager *accounts.Manager
	dataHandler    *JsonFileDataHandler
	stateFetcher   *StateFetcher
}

func NewServicesFactory(node *node.Node, ptm private.PrivateTransactionManager, ethService *eth.Ethereum) (*DefaultServicesFactory, error) {
	factory := &DefaultServicesFactory{}

	factory.accountManager = ethService.AccountManager()
	factory.dataHandler = NewJsonFileDataHandler(node.InstanceDir())
	factory.stateFetcher = NewStateFetcher(ethService.BlockChain())

	backendService, err := New(ptm, factory.AccountManager(), factory.DataHandler(), factory.StateFetcher())
	if err != nil {
		return nil, err
	}
	factory.backendService = backendService

	ethService.BlockChain().PopulateSetPrivateState(privacyExtension.DefaultExtensionHandler.CheckExtensionAndSetPrivateState)

	go backendService.initialise(node)

	return factory, nil
}

func (factory *DefaultServicesFactory) BackendService() *PrivacyService {
	return factory.backendService
}

func (factory *DefaultServicesFactory) AccountManager() *accounts.Manager {
	return factory.accountManager
}

func (factory *DefaultServicesFactory) DataHandler() DataHandler {
	return factory.dataHandler
}

func (factory *DefaultServicesFactory) StateFetcher() *StateFetcher {
	return factory.stateFetcher
}
