package extension

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/extension/extensionContracts"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/rpc"
)

type PrivacyService struct {
	ptm                      private.PrivateTransactionManager
	stateFetcher             *StateFetcher
	accountManager           *accounts.Manager
	dataHandler              DataHandler
	managementContractFacade ManagementContractFacade
	extClient                Client
	stopFeed                 event.Feed

	mu               sync.Mutex
	currentContracts map[common.Address]*ExtensionContract
}

var (
	//default gas limit to use if not passed in sendTxArgs
	defaultGasLimit = uint64(4712384)
	//default gas price to use if not passed in sendTxArgs
	defaultGasPrice = big.NewInt(0)

	//Private participants must be specified for contract extension related transactions
	errNotPrivate = errors.New("must specify private participants")
)

// to signal all watches when service is stopped
type stopEvent struct {
}

func (service *PrivacyService) subscribeStopEvent() (chan stopEvent, event.Subscription) {
	c := make(chan stopEvent)
	s := service.stopFeed.Subscribe(c)
	return c, s
}

func New(ptm private.PrivateTransactionManager, manager *accounts.Manager, handler DataHandler, fetcher *StateFetcher) (*PrivacyService, error) {
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

func (service *PrivacyService) initialise(node *node.Node) {
	service.mu.Lock()
	defer service.mu.Unlock()

	rpcClient, err := node.Attach()
	if err != nil {
		panic("extension: could not connect to ethereum client rpc")
	}

	client := ethclient.NewClientWithPTM(rpcClient, service.ptm)
	service.managementContractFacade = NewManagementContractFacade(client)
	service.extClient = NewInProcessClient(client)

	for _, f := range []func() error{
		service.watchForNewContracts,       // watch for new extension contract creation event
		service.watchForCancelledContracts, // watch for extension contract cancellation event
		service.watchForCompletionEvents,   // watch for extension contract voting complete event
	} {
		if err := f(); err != nil {
			log.Error("")
		}
	}

}

func (service *PrivacyService) watchForNewContracts() error {
	incomingLogs, subscription, err := service.extClient.SubscribeToLogs(newExtensionQuery)

	if err != nil {
		return err
	}

	go func() {
		stopChan, stopSubscription := service.subscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case err := <-subscription.Err():
				log.Error("Contract extension watcher subscription error", "error", err)
				break

			case foundLog := <-incomingLogs:
				service.mu.Lock()

				tx, _ := service.extClient.TransactionByHash(foundLog.TxHash)
				from, _ := types.QuorumPrivateTxSigner{}.Sender(tx)

				newExtensionEvent, err := extensionContracts.UnpackNewExtensionCreatedLog(foundLog.Data)
				if err != nil {
					log.Error("Error unpacking extension creation log", "error", err)
					log.Debug("Errored log", foundLog)
					service.mu.Unlock()
					continue
				}

				newContractExtension := ExtensionContract{
					ContractExtended:          newExtensionEvent.ToExtend,
					Initiator:                 from,
					Recipient:                 newExtensionEvent.RecipientAddress,
					RecipientPtmKey:           newExtensionEvent.RecipientPTMKey,
					ManagementContractAddress: foundLog.Address,
					CreationData:              tx.Data(),
				}

				service.currentContracts[foundLog.Address] = &newContractExtension
				err = service.dataHandler.Save(service.currentContracts)
				if err != nil {
					log.Error("Error writing extension data to file", "error", err)
					service.mu.Unlock()
					continue
				}
				service.mu.Unlock()

				// if party is sender then complete self voting
				data := common.BytesToEncryptedPayloadHash(newContractExtension.CreationData)
				isSender, _ := service.ptm.IsSender(data)

				if isSender {
					fetchedParties, err := service.ptm.GetParticipants(data)
					if err != nil || len(fetchedParties) == 0 {
						log.Error("Extension: unable to fetch all parties for extension management contract", "error", err)
						continue
					}
					//Find the extension contract in order to interact with it
					caller, _ := service.managementContractFacade.Caller(newContractExtension.ManagementContractAddress)
					contractCreator, _ := caller.Creator(nil)

					txArgs := ethapi.SendTxArgs{From: contractCreator, PrivateTxArgs: ethapi.PrivateTxArgs{PrivateFor: fetchedParties}}

					extensionAPI := NewPrivateExtensionAPI(service)
					_, err = extensionAPI.ApproveExtension(newContractExtension.ManagementContractAddress, true, txArgs)

					if err != nil {
						log.Error("Extension: initiator vote on management contract failed", "error", err)
					}
				}

			case <-stopChan:
				return
			}
		}
	}()

	return nil
}

func (service *PrivacyService) watchForCancelledContracts() error {
	incomingLogs, subscription, err := service.extClient.SubscribeToLogs(finishedExtensionQuery)

	if err != nil {
		return err
	}

	go func() {
		stopChan, stopSubscription := service.subscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case err := <-subscription.Err():
				log.Error("Contract cancellation extension watcher subscription error", "error", err)
				return
			case l := <-incomingLogs:
				service.mu.Lock()
				if _, ok := service.currentContracts[l.Address]; ok {
					delete(service.currentContracts, l.Address)
					if err := service.dataHandler.Save(service.currentContracts); err != nil {
						log.Error("Faile to store list of contracts being extended", "error", err)
					}
				}
				service.mu.Unlock()
			case <-stopChan:
				return
			}
		}

	}()

	return nil
}

func (service *PrivacyService) watchForCompletionEvents() error {
	incomingLogs, _, err := service.extClient.SubscribeToLogs(canPerformStateShareQuery)

	if err != nil {
		return err
	}

	go func() {
		stopChan, stopSubscription := service.subscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case l := <-incomingLogs:
				log.Debug("Extension: Received a completion event", "address", l.Address.Hex(), "blockNumber", l.BlockNumber)
				service.mu.Lock()
				func() {
					defer func() {
						service.mu.Unlock()
					}()
					extensionEntry, ok := service.currentContracts[l.Address]
					if !ok {
						// we didn't have this management contract, so ignore it
						log.Debug("Extension: this node doesn't participate in the contract extender", "address", l.Address.Hex())
						return
					}

					//Find the extension contract in order to interact with it
					caller, err := service.managementContractFacade.Caller(l.Address)
					if err != nil {
						log.Error("service.managementContractFacade.Caller", "address", l.Address.Hex(), "error", err)
						return
					}
					contractCreator, err := caller.Creator(nil)
					if err != nil {
						log.Error("[contract] caller.Creator", "error", err)
						return
					}
					log.Debug("Extension: check if this node has the account that created the contract extender", "account", contractCreator)
					if _, err := service.accountManager.Find(accounts.Account{Address: contractCreator}); err != nil {
						log.Warn("Account used to sign extension contract no longer available", "account", contractCreator.Hex())
						return
					}

					// fetch all the participants and send
					payload := common.BytesToEncryptedPayloadHash(extensionEntry.CreationData)
					fetchedParties, err := service.ptm.GetParticipants(payload)
					if err != nil || len(fetchedParties) == 0 {
						log.Error("Extension: Unable to fetch all parties for extension management contract", "error", err)
						return
					}
					log.Debug("Extension: able to fetch all parties", "parties", fetchedParties)

					txArgs, err := service.GenerateTransactOptions(ethapi.SendTxArgs{From: contractCreator, PrivateTxArgs: ethapi.PrivateTxArgs{PrivateFor: fetchedParties}})
					if err != nil {
						log.Error("service.accountManager.GenerateTransactOptions", "error", err, "contractCreator", contractCreator.Hex(), "privateFor", fetchedParties)
						return
					}

					//we found the account, so we can send
					contractToExtend, err := caller.ContractToExtend(nil)
					if err != nil {
						log.Error("[contract] caller.ContractToExtend", "error", err)
						return
					}
					log.Debug("Extension: dump current state", "block", l.BlockHash, "contract", contractToExtend.Hex())
					entireStateData, err := service.stateFetcher.GetAddressStateFromBlock(l.BlockHash, contractToExtend)
					if err != nil {
						log.Error("[state] service.stateFetcher.GetAddressStateFromBlock", "block", l.BlockHash.Hex(), "contract", contractToExtend.Hex(), "error", err)
						return
					}

					log.Debug("Extension: send the state dump to the new recipient", "recipients", fetchedParties)

					// PSV & PP changes
					// send the new transaction with state dump to all participants
					extraMetaData := engine.ExtraMetadata{PrivacyFlag: engine.PrivacyFlagStandardPrivate}
					privacyMetaData, err := service.stateFetcher.GetPrivacyMetaData(l.BlockHash, contractToExtend)
					if err != nil {
						log.Error("[privacyMetaData] fetch err", "err", err)
					} else {
						extraMetaData.PrivacyFlag = privacyMetaData.PrivacyFlag
						if privacyMetaData.PrivacyFlag == engine.PrivacyFlagStateValidation {
							storageRoot, err := service.stateFetcher.GetStorageRoot(l.BlockHash, contractToExtend)
							if err != nil {
								log.Error("[storageRoot] fetch err", "err", err)
							}
							extraMetaData.ACMerkleRoot = storageRoot
						}
					}
					hashOfStateData, err := service.ptm.Send(entireStateData, "", fetchedParties, &extraMetaData)

					if err != nil {
						log.Error("[ptm] service.ptm.Send", "stateDataInHex", hex.EncodeToString(entireStateData[:]), "recipients", fetchedParties, "error", err)
						return
					}
					hashofStateDataBase64 := hashOfStateData.ToBase64()

					transactor, err := service.managementContractFacade.Transactor(l.Address)
					if err != nil {
						log.Error("service.managementContractFacade.Transactor", "address", l.Address.Hex(), "error", err)
						return
					}
					log.Debug("Extension: store the encrypted payload hash of dump state", "contract", l.Address.Hex())
					if tx, err := transactor.SetSharedStateHash(txArgs, hashofStateDataBase64); err != nil {
						log.Error("[contract] transactor.SetSharedStateHash", "error", err, "hashOfStateInBase64", hashofStateDataBase64)
					} else {
						log.Debug("Extension: transaction carrying shared state", "txhash", tx.Hash(), "private", tx.IsPrivate())
					}
				}()
			case <-stopChan:
				return
			}
		}

	}()
	return nil
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
	log.Debug("extension service: starting")
	return nil
}

func (service *PrivacyService) Stop() error {
	log.Info("extension service: stopping")
	service.stopFeed.Send(stopEvent{})
	log.Info("extension service: stopped")
	return nil
}

func (service *PrivacyService) GenerateTransactOptions(txa ethapi.SendTxArgs) (*bind.TransactOpts, error) {
	if txa.PrivateFor == nil {
		return nil, errNotPrivate
	}
	from := accounts.Account{Address: txa.From}
	wallet, err := service.accountManager.Find(from)

	if err != nil {
		return nil, fmt.Errorf("no wallet found for account %s", txa.From.String())
	}

	//Find the account we plan to send the transaction from

	txArgs := bind.NewWalletTransactor(wallet, from)
	txArgs.PrivateFrom = txa.PrivateFrom
	txArgs.PrivateFor = txa.PrivateFor
	txArgs.GasLimit = defaultGasLimit
	txArgs.GasPrice = defaultGasPrice
	if txa.GasPrice != nil {
		txArgs.GasPrice = txa.GasPrice.ToInt()
	}
	if txa.Gas != nil {
		txArgs.GasLimit = uint64(*txa.Gas)
	}
	return txArgs, nil
}

// returns the participant list for a given private contract
func (service *PrivacyService) GetAllParticipants(blockHash common.Hash, address common.Address) ([]string, error) {
	privacyMetaData, err := service.stateFetcher.GetPrivacyMetaData(blockHash, address)
	if err != nil {
		return nil, err
	}
	if privacyMetaData.PrivacyFlag.IsStandardPrivate() {
		return nil, nil
	}

	participants, err := service.ptm.GetParticipants(privacyMetaData.CreationTxHash)
	if err != nil {
		return nil, err
	}
	return participants, nil
}

// check if the node had created the contract
func (service *PrivacyService) CheckIfContractCreator(blockHash common.Hash, address common.Address) bool {
	privacyMetaData, err := service.stateFetcher.GetPrivacyMetaData(blockHash, address)
	if err != nil {
		return true
	}

	isCreator, err := service.ptm.IsSender(privacyMetaData.CreationTxHash)
	if err != nil {
		return false
	}

	return isCreator
}
