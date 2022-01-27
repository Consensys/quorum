package extension

import (
	"context"
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
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/rpc"
)

type PrivacyService struct {
	ptm              private.PrivateTransactionManager
	stateFetcher     *StateFetcher
	accountManager   *accounts.Manager
	dataHandler      DataHandler
	stopFeed         event.Feed
	apiBackendHelper APIBackendHelper

	mu           sync.Mutex
	psiContracts map[types.PrivateStateIdentifier]map[common.Address]*ExtensionContract

	node   *node.Node
	config *params.ChainConfig

	isQlightClient bool
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

func (service *PrivacyService) newEthClient(psi types.PrivateStateIdentifier) *ethclient.Client {
	rpcClient, err := service.node.AttachWithPSI(psi)
	if err != nil {
		// AttachWithPSI does not return non-nil error. This is just a defensive check
		panic("this should not happen: " + err.Error())
	}
	return ethclient.NewClientWithPTM(rpcClient, service.ptm)
}

func (service *PrivacyService) client(psi types.PrivateStateIdentifier) Client {
	return NewInProcessClient(service.newEthClient(psi))
}

func (service *PrivacyService) managementContract(psi types.PrivateStateIdentifier) ManagementContractFacade {
	return NewManagementContractFacade(service.newEthClient(psi))
}

func (service *PrivacyService) subscribeStopEvent() (chan stopEvent, event.Subscription) {
	c := make(chan stopEvent)
	s := service.stopFeed.Subscribe(c)
	return c, s
}

func New(stack *node.Node, ptm private.PrivateTransactionManager, manager *accounts.Manager, handler DataHandler, fetcher *StateFetcher, apiBackendHelper APIBackendHelper, config *params.ChainConfig) (*PrivacyService, error) {
	service := &PrivacyService{
		psiContracts:     make(map[types.PrivateStateIdentifier]map[common.Address]*ExtensionContract),
		ptm:              ptm,
		dataHandler:      handler,
		stateFetcher:     fetcher,
		accountManager:   manager,
		apiBackendHelper: apiBackendHelper,
		node:             stack,
		config:           config,
		isQlightClient:   false,
	}

	apiSupport, ok := service.apiBackendHelper.(ethapi.ProxyAPISupport)
	if ok {
		if apiSupport.ProxyEnabled() {
			service.isQlightClient = true
		}
	}

	var err error
	service.psiContracts, err = service.dataHandler.Load()
	if err != nil {
		return nil, errors.New("could not load existing extension contracts: " + err.Error())
	}

	// Register service to node
	stack.RegisterAPIs(service.apis())
	stack.RegisterLifecycle(service)

	return service, nil
}

func (service *PrivacyService) watchForNewContracts(psi types.PrivateStateIdentifier) error {
	handler := NewSubscriptionHandler(service.node, psi, service.ptm, service)

	cb := func(foundLog types.Log) {
		service.mu.Lock()
		psiClient := service.client(psi)
		defer psiClient.Close()
		tx, _ := service.client(psi).TransactionInBlock(foundLog.BlockHash, foundLog.TxIndex)
		from, _ := types.QuorumPrivateTxSigner{}.Sender(tx)

		newExtensionEvent, err := extensionContracts.UnpackNewExtensionCreatedLog(foundLog.Data)
		if err != nil {
			log.Error("Error unpacking extension creation log", "error", err)
			log.Debug("Errored log", foundLog)
			service.mu.Unlock()
			return
		}

		newContractExtension := ExtensionContract{
			ContractExtended:          newExtensionEvent.ToExtend,
			Initiator:                 from,
			Recipient:                 newExtensionEvent.RecipientAddress,
			RecipientPtmKey:           newExtensionEvent.RecipientPTMKey,
			ManagementContractAddress: foundLog.Address,
			CreationData:              tx.Data(),
		}

		enclaveKey := common.BytesToEncryptedPayloadHash(tx.Data())
		privateFrom, _, _, _, err := service.ptm.Receive(enclaveKey)
		if err != nil {
			log.Error("Error receiving private payload", "error", err)
			service.mu.Unlock()
			return
		}

		if service.psiContracts[psi] == nil {
			service.psiContracts[psi] = make(map[common.Address]*ExtensionContract)
		}
		service.psiContracts[psi][foundLog.Address] = &newContractExtension

		if err := service.dataHandler.Save(service.psiContracts); err != nil {
			log.Error("Error writing extension data to file", "error", err)
			service.mu.Unlock()
			return
		}
		service.mu.Unlock()

		// if party is sender then complete self voting

		isSender, _ := service.ptm.IsSender(enclaveKey)

		if service.isQlightClient {
			log.Debug("Extension: this is a light node and it does not handle self vote events", "address", newContractExtension.ContractExtended.Hex())
			return
		}

		if isSender {
			fetchedParties, err := service.ptm.GetParticipants(enclaveKey)
			if err != nil || len(fetchedParties) == 0 {
				log.Error("Extension: unable to fetch all parties for extension management contract", "error", err)
				return
			}

			psm, _ := service.apiBackendHelper.PSMR().ResolveForManagedParty(privateFrom)
			if psm.ID != psi {
				return
			}

			psiManagementContractClient := service.managementContract(psi)
			defer psiManagementContractClient.Close()
			//Find the extension contract in order to interact with it
			caller, _ := psiManagementContractClient.Caller(newContractExtension.ManagementContractAddress)
			contractCreator, _ := caller.Creator(nil)

			txArgs := ethapi.SendTxArgs{From: contractCreator, PrivateTxArgs: ethapi.PrivateTxArgs{PrivateFor: fetchedParties, PrivateFrom: privateFrom}}

			extensionAPI := NewPrivateExtensionAPI(service)
			ctx := rpc.WithPrivateStateIdentifier(context.Background(), psm.ID)
			_, err = extensionAPI.ApproveExtension(ctx, newContractExtension.ManagementContractAddress, true, txArgs)

			if err != nil {
				log.Error("Extension: initiator vote on management contract failed", "error", err)
			}
		}
	}

	return handler.createSub(newExtensionQuery, cb)
}

func (service *PrivacyService) watchForCancelledContracts(psi types.PrivateStateIdentifier) error {
	handler := NewSubscriptionHandler(service.node, psi, service.ptm, service)

	cb := func(l types.Log) {
		service.mu.Lock()
		if _, ok := service.psiContracts[psi][l.Address]; ok {
			delete(service.psiContracts[psi], l.Address)
			if err := service.dataHandler.Save(service.psiContracts); err != nil {
				log.Error("Failed to store list of contracts being extended", "error", err)
			}
		}
		service.mu.Unlock()
	}

	return handler.createSub(finishedExtensionQuery, cb)
}

func (service *PrivacyService) watchForCompletionEvents(psi types.PrivateStateIdentifier) error {
	handler := NewSubscriptionHandler(service.node, psi, service.ptm, service)

	cb := func(l types.Log) {
		log.Debug("Extension: Received a completion event", "address", l.Address.Hex(), "blockNumber", l.BlockNumber)

		if service.isQlightClient {
			log.Debug("Extension: this is a light node and it does not handle completion events", "address", l.Address.Hex())
			return
		}

		service.mu.Lock()
		defer func() {
			service.mu.Unlock()
		}()
		extensionEntry, ok := service.psiContracts[psi][l.Address]
		if !ok {
			// we didn't have this management contract, so ignore it
			log.Debug("Extension: this node doesn't participate in the contract extender", "address", l.Address.Hex())
			return
		}

		psiManagementContractClient := service.managementContract(psi)
		defer psiManagementContractClient.Close()
		//Find the extension contract in order to interact with it
		caller, err := psiManagementContractClient.Caller(l.Address)
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

		privateFrom, _, _, _, err := service.ptm.Receive(payload)
		if err != nil || len(privateFrom) == 0 {
			log.Error("Extension: unable to fetch privateFrom(sender) for extension management contract", "error", err)
			return
		}
		log.Debug("Extension: able to fetch privateFrom(sender)", "privateFrom", privateFrom)

		txPsi, err := service.apiBackendHelper.PSMR().ResolveForManagedParty(privateFrom)
		if err != nil {
			log.Error("Extension: unable to resolve private state metadata for sender", "error", err)
			return
		}
		if txPsi.ID != psi {
			return
		}
		txArgs, err := service.GenerateTransactOptions(ethapi.SendTxArgs{From: contractCreator, PrivateTxArgs: ethapi.PrivateTxArgs{PrivateFor: fetchedParties, PrivateFrom: privateFrom}})
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
		log.Debug("Extension: dump current state", "block", l.BlockHash, "contract", contractToExtend.Hex(), "psi", txPsi.ID)
		entireStateData, err := service.stateFetcher.GetAddressStateFromBlock(l.BlockHash, contractToExtend, txPsi.ID)
		if err != nil {
			log.Error("[state] service.stateFetcher.GetAddressStateFromBlock", "block", l.BlockHash.Hex(), "contract", contractToExtend.Hex(), "error", err)
			return
		}

		log.Debug("Extension: send the state dump to the new recipient", "recipients", fetchedParties)

		// PSV & PP changes
		// send the new transaction with state dump to all participants
		extraMetaData := engine.ExtraMetadata{PrivacyFlag: engine.PrivacyFlagStandardPrivate}
		privacyMetaData, err := service.stateFetcher.GetPrivacyMetaData(l.BlockHash, contractToExtend, txPsi.ID)
		if err != nil {
			log.Error("[privacyMetaData] fetch err", "err", err)
		} else {
			extraMetaData.PrivacyFlag = privacyMetaData.PrivacyFlag
			if privacyMetaData.PrivacyFlag == engine.PrivacyFlagStateValidation {
				storageRoot, err := service.stateFetcher.GetStorageRoot(l.BlockHash, contractToExtend, txPsi.ID)
				if err != nil {
					log.Error("[storageRoot] fetch err", "err", err)
				}
				extraMetaData.ACMerkleRoot = storageRoot
			}
			// Fetch mandatory recipients data from Tessera - only when privacy flag is 2
			if privacyMetaData.PrivacyFlag == engine.PrivacyFlagMandatoryRecipients {
				fetchedMandatoryRecipients, err := service.ptm.GetMandatory(privacyMetaData.CreationTxHash)
				if err != nil || len(fetchedMandatoryRecipients) == 0 {
					log.Error("Extension: Unable to fetch mandatory parties for extension management contract", "error", err)
					return
				}
				log.Debug("Extension: able to fetch mandatory recipients", "mandatory", fetchedMandatoryRecipients)
				extraMetaData.MandatoryRecipients = fetchedMandatoryRecipients
			}
		}

		_, _, hashOfStateData, err := service.ptm.Send(entireStateData, privateFrom, fetchedParties, &extraMetaData)

		if err != nil {
			log.Error("[ptm] service.ptm.Send", "stateDataInHex", hex.EncodeToString(entireStateData[:]), "recipients", fetchedParties, "error", err)
			return
		}
		hashofStateDataBase64 := hashOfStateData.ToBase64()

		transactor, err := psiManagementContractClient.Transactor(l.Address)
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
	}

	return handler.createSub(canPerformStateShareQuery, cb)
}

// utility methods
func (service *PrivacyService) apis() []rpc.API {
	return []rpc.API{
		{
			Namespace: "quorumExtension",
			Version:   "1.0",
			Service:   NewPrivateExtensionProxyAPI(service),
			Public:    true,
		},
	}
}

// node.Lifecycle interface methods:

func (service *PrivacyService) Start() error {
	log.Debug("extension service: starting")
	service.mu.Lock()
	defer service.mu.Unlock()

	for _, psi := range service.apiBackendHelper.PSMR().PSIs() {
		for _, f := range []func(identifier types.PrivateStateIdentifier) error{
			service.watchForNewContracts,       // watch for new extension contract creation event
			service.watchForCancelledContracts, // watch for extension contract cancellation event
			service.watchForCompletionEvents,   // watch for extension contract voting complete event
		} {
			if err := f(psi); err != nil {
				return err
			}
		}
	}

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

	txArgs := bind.NewWalletTransactor(wallet, from, service.config.ChainID)
	txArgs.PrivateFrom = txa.PrivateFrom
	txArgs.PrivateFor = txa.PrivateFor
	txArgs.GasLimit = defaultGasLimit
	txArgs.GasPrice = defaultGasPrice
	txArgs.IsUsingPrivacyPrecompile = service.apiBackendHelper.IsPrivacyMarkerTransactionCreationEnabled()

	if txa.GasPrice != nil {
		txArgs.GasPrice = txa.GasPrice.ToInt()
	}
	if txa.Gas != nil {
		txArgs.GasLimit = uint64(*txa.Gas)
	}
	return txArgs, nil
}

// returns the participant list for a given private contract
func (service *PrivacyService) GetAllParticipants(blockHash common.Hash, address common.Address, psi types.PrivateStateIdentifier) ([]string, error) {
	privacyMetaData, err := service.stateFetcher.GetPrivacyMetaData(blockHash, address, psi)
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
func (service *PrivacyService) CheckIfContractCreator(blockHash common.Hash, address common.Address, psi types.PrivateStateIdentifier) bool {
	privacyMetaData, err := service.stateFetcher.GetPrivacyMetaData(blockHash, address, psi)
	if err != nil {
		return true
	}

	isCreator, err := service.ptm.IsSender(privacyMetaData.CreationTxHash)
	if err != nil {
		return false
	}

	if !isCreator {
		return false
	}

	privateFrom, _, _, _, err := service.ptm.Receive(privacyMetaData.CreationTxHash)
	if err != nil || len(privateFrom) == 0 {
		return false
	}
	psm, _ := service.apiBackendHelper.PSMR().ResolveForManagedParty(privateFrom)
	return psm.ID == psi
}
