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

func NewServicesFactory(stack *node.Node, ptm private.PrivateTransactionManager, ethService *eth.Ethereum) (*DefaultServicesFactory, error) {
	factory := &DefaultServicesFactory{}

	factory.accountManager = ethService.AccountManager()
	factory.dataHandler = NewJsonFileDataHandler(stack.InstanceDir())
	factory.stateFetcher = NewStateFetcher(ethService.BlockChain())

	backendService, err := New(stack, ptm, factory.AccountManager(), factory.DataHandler(), factory.StateFetcher(), ethService.APIBackend)
	if err != nil {
		return nil, err
	}
	factory.backendService = backendService

	isMultitenant := ethService.BlockChain().SupportsMultitenancy(context.Background())
	privacyExtension.DefaultExtensionHandler.SupportMultitenancy(isMultitenant)
	privacyExtension.DefaultExtensionHandler.SetPSMR(ethService.BlockChain().PrivateStateManager())

	ethService.BlockChain().PopulateSetPrivateState(privacyExtension.DefaultExtensionHandler.CheckExtensionAndSetPrivateState)

	return factory, nil
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
