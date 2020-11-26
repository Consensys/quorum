package permission

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/permission/basic"
	"github.com/ethereum/go-ethereum/permission/cache"
	"github.com/ethereum/go-ethereum/permission/eea"
	ptype "github.com/ethereum/go-ethereum/permission/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type PermissionCtrl struct {
	node               *node.Node
	ethClnt            bind.ContractBackend
	eth                *eth.Ethereum
	key                *ecdsa.PrivateKey
	dataDir            string
	permConfig         *ptype.PermissionConfig
	contract           ptype.InitService
	backend            ptype.Backend
	useDns             bool
	isRaft             bool
	startWaitGroup     *sync.WaitGroup // waitgroup to make sure all dependencies are ready before we start the service
	errorChan          chan error      // channel to capture error when starting aysnc
	networkInitialized bool
	controlService     ptype.ControlService
}

var permissionService *PermissionCtrl

func setPermissionService(ps *PermissionCtrl) {
	if permissionService == nil {
		permissionService = ps
	}
}

// Create a service instance for permissioning
//
// Permission Service depends on the following:
// 1. EthService to be ready
// 2. Downloader to sync up blocks
// 3. InProc RPC server to be ready
func NewQuorumPermissionCtrl(stack *node.Node, pconfig *ptype.PermissionConfig, useDns bool) (*PermissionCtrl, error) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	p := &PermissionCtrl{
		node:           stack,
		key:            stack.GetNodeKey(),
		dataDir:        stack.DataDir(),
		permConfig:     pconfig,
		startWaitGroup: wg,
		errorChan:      make(chan error),
		useDns:         useDns,
		isRaft:         false,
	}

	err := p.populateBackEnd()
	if err != nil {
		return nil, err
	}
	stopChan, stopSubscription := ptype.SubscribeStopEvent()
	inProcRPCServerSub := stack.EventMux().Subscribe(rpc.InProcServerReadyEvent{})
	log.Debug("permission service: waiting for InProcRPC Server")

	go func(_wg *sync.WaitGroup) {
		defer func(start time.Time) {
			log.Debug("permission service: InProcRPC server is ready", "took", time.Since(start))
			stopSubscription.Unsubscribe()
			inProcRPCServerSub.Unsubscribe()
			_wg.Done()
		}(time.Now())
		select {
		case <-inProcRPCServerSub.Chan():
		case <-stopChan:
		}
	}(wg) // wait for inproc RPC to be ready
	return p, nil
}

func (p *PermissionCtrl) Start(srvr *p2p.Server) error {
	log.Info("permission service: starting")
	go func() {
		log.Info("permission service: starting async")
		p.asyncStart()
	}()
	return nil
}

func (p *PermissionCtrl) Protocols() []p2p.Protocol {
	return []p2p.Protocol{}
}

func (p *PermissionCtrl) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "quorumPermission",
			Version:   "1.0",
			Service:   NewQuorumControlsAPI(p),
			Public:    true,
		},
	}
}

func (p *PermissionCtrl) Stop() error {
	log.Info("permission service: stopping")
	ptype.StopFeed.Send(ptype.StopEvent{})
	log.Info("permission service: stopped")
	return nil
}

func (p *PermissionCtrl) IsEEAPermission() bool {
	return p.permConfig.PermissionsModel == ptype.PERMISSION_EEA
}

func NewPermissionContractService(ethClnt bind.ContractBackend, eeaFlag bool, key *ecdsa.PrivateKey,
	permConfig *ptype.PermissionConfig, isRaft, useDns bool) ptype.InitService {

	contractBackEnd := ptype.ContractBackend{
		EthClnt:    ethClnt,
		Key:        key,
		PermConfig: permConfig,
		IsRaft:     isRaft,
		UseDns:     useDns,
	}

	if eeaFlag {
		return &eea.Init{
			Backend: contractBackEnd,
		}
	}
	return &basic.Init{
		Backend: contractBackEnd,
	}
}

func (p *PermissionCtrl) NewPermissionRoleService(txa ethapi.SendTxArgs) (ptype.RoleService, error) {
	transactOpts, err := p.getTxParams(txa)
	if err != nil {
		return nil, err
	}
	return p.backend.GetRoleService(transactOpts, p.getContractBackend())
}

func (p *PermissionCtrl) NewPermissionOrgService(txa ethapi.SendTxArgs) (ptype.OrgService, error) {
	transactOpts, err := p.getTxParams(txa)
	if err != nil {
		return nil, err
	}
	return p.backend.GetOrgService(transactOpts, p.getContractBackend())
}

func (p *PermissionCtrl) NewPermissionNodeService(txa ethapi.SendTxArgs) (ptype.NodeService, error) {
	transactOpts, err := p.getTxParams(txa)
	if err != nil {
		return nil, err
	}
	return p.backend.GetNodeService(transactOpts, p.getContractBackend())
}

func (p *PermissionCtrl) NewPermissionAccountService(txa ethapi.SendTxArgs) (ptype.AccountService, error) {
	transactOpts, err := p.getTxParams(txa)
	if err != nil {
		return nil, err
	}
	return p.backend.GetAccountService(transactOpts, p.getContractBackend())
}

func (p *PermissionCtrl) NewPermissionAuditService() (ptype.AuditService, error) {
	return p.backend.GetAuditService(p.getContractBackend())
}

func (p *PermissionCtrl) NewPermissionControlService() (ptype.ControlService, error) {
	return p.backend.GetControlService(p.getContractBackend())
}

func (p *PermissionCtrl) getContractBackend() ptype.ContractBackend {
	return ptype.ContractBackend{EthClnt: p.ethClnt, Key: p.key, PermConfig: p.permConfig, IsRaft: p.isRaft, UseDns: p.isRaft}
}

func (p *PermissionCtrl) ConnectionAllowed(_enodeId, _ip string, _port, _raftPort uint16) (bool, error) {
	cs, err := p.backend.GetControlService(p.getContractBackend())
	if err != nil {
		return false, err
	}
	return cs.ConnectionAllowed(_enodeId, _ip, _port, _raftPort)
}

func (p *PermissionCtrl) IsTransactionAllowed(_sender common.Address, _target common.Address, _value *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _payload []byte, transactionType cache.TransactionType) error {
	// If permissions model is not in use return nil
	if cache.PermissionModel == cache.Default {
		return nil
	}

	cs, err := p.backend.GetControlService(p.getContractBackend())
	if err != nil {
		return err
	}

	return cs.TransactionAllowed(_sender, _target, _value, _gasPrice, _gasLimit, _payload, transactionType)
}

func (p *PermissionCtrl) populateBackEnd() error {
	backend := ptype.NewInterfaceBackend(p.node, false, p.dataDir)

	switch p.permConfig.PermissionsModel {
	case ptype.PERMISSION_EEA:
		p.backend = &eea.Backend{
			Ib: *backend,
		}
		log.Debug("permission service: using eea permissions model")
		return nil

	case ptype.PERMISSION_BASIC:
		p.backend = &basic.Backend{
			Ib: *backend,
		}
		log.Debug("permission service: using basic permissions model")
		return nil

	default:
		return errors.New("permission: invalid permissions model passed")
	}

}

func (p *PermissionCtrl) updateBackEnd() {
	p.contract = NewPermissionContractService(p.ethClnt, p.IsEEAPermission(), p.key, p.permConfig, p.isRaft, p.useDns)
	switch p.IsEEAPermission() {
	case true:
		p.backend.(*eea.Backend).Contr = p.contract.(*eea.Init)
		p.backend.(*eea.Backend).Ib.SetIsRaft(p.isRaft)

	default:
		p.backend.(*basic.Backend).Contr = p.contract.(*basic.Init)
		p.backend.(*basic.Backend).Ib.SetIsRaft(p.isRaft)
	}
}

// validateAccount validates the account and returns the wallet associated with that for signing the transaction
func (p *PermissionCtrl) validateAccount(from common.Address) (accounts.Wallet, error) {
	acct := accounts.Account{Address: from}
	w, err := p.eth.AccountManager().Find(acct)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// getTxParams extracts the transaction related parameters
func (p *PermissionCtrl) getTxParams(txa ethapi.SendTxArgs) (*bind.TransactOpts, error) {
	w, err := p.validateAccount(txa.From)
	if err != nil {
		return nil, ptype.ErrInvalidAccount
	}
	fromAcct := accounts.Account{Address: txa.From}
	transactOpts := bind.NewWalletTransactor(w, fromAcct)

	transactOpts.GasPrice = defaultGasPrice
	if txa.GasPrice != nil {
		transactOpts.GasPrice = txa.GasPrice.ToInt()
	}

	transactOpts.GasLimit = defaultGasLimit
	if txa.Gas != nil {
		transactOpts.GasLimit = uint64(*txa.Gas)
	}
	transactOpts.From = fromAcct.Address

	return transactOpts, nil
}
