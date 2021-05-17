package extension

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/private"
)

type subscriptionHandler struct {
	facade  ManagementContractFacade
	client  Client
	service *PrivacyService
}

func NewSubscriptionHandler(node *node.Node, psi types.PrivateStateIdentifier, ptm private.PrivateTransactionManager, service *PrivacyService) *subscriptionHandler {
	rpcClient, err := node.AttachWithPSI(psi)
	if err != nil {
		panic("extension: could not connect to ethereum client rpc")
	}

	client := ethclient.NewClientWithPTM(rpcClient, ptm)

	return &subscriptionHandler{
		facade:  NewManagementContractFacade(client),
		client:  NewInProcessClient(client),
		service: service,
	}
}

func (handler *subscriptionHandler) createSub(query ethereum.FilterQuery, logHandlerCb func(types.Log)) error {
	incomingLogs, subscription, err := handler.client.SubscribeToLogs(query)

	if err != nil {
		return err
	}

	go func() {
		stopChan, stopSubscription := handler.service.subscribeStopEvent()
		defer stopSubscription.Unsubscribe()

		for {
			select {
			case err := <-subscription.Err():
				log.Error("Contract extension watcher subscription error", "error", err)
				break
			case foundLog := <-incomingLogs:
				logHandlerCb(foundLog)
			case <-stopChan:
				return
			}
		}
	}()

	return nil
}
