package extension

import (
	"context"

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

	rpcClient, err := node.Attach()
	if err != nil {
		panic("extension: could not connect to ethereum client rpc")
	}

	backendService, err := New(ptm, factory.AccountManager(), factory.DataHandler(), factory.StateFetcher(), ethService.APIBackend, rpcClient)
	if err != nil {
		return nil, err
	}
	factory.backendService = backendService

	_, isMultitenant := ethService.BlockChain().SupportsMultitenancy(context.Background())
	privacyExtension.DefaultExtensionHandler.SupportMultitenancy(isMultitenant)

	ethService.BlockChain().PopulateSetPrivateState(privacyExtension.DefaultExtensionHandler.CheckExtensionAndSetPrivateState)

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
