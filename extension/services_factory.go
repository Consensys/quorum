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

	// TODO - I have no idea what other bits (other than ethService.APIBackend) I may need from ethService so for now just keep a handle to it
	// must revisit later and reduce access to only what is needed
	backendService, err := New(ptm, factory.AccountManager(), factory.DataHandler(), factory.StateFetcher(), ethService)
	if err != nil {
		return nil, err
	}
	factory.backendService = backendService

	privacyExtension.DefaultExtensionHandler.MultitenancyEnabled = ethService.BlockChain().SupportsMultitenancy()
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
