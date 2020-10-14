package permission

import (
	"crypto/ecdsa"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/permission/basic"
	bb "github.com/ethereum/go-ethereum/permission/basic/bind"
	"github.com/ethereum/go-ethereum/permission/eea"
	eb "github.com/ethereum/go-ethereum/permission/eea/bind"
	ptype "github.com/ethereum/go-ethereum/permission/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type PermissionCtrl struct {
	node               *node.Node
	ethClnt            bind.ContractBackend
	eth                *eth.Ethereum
	key                *ecdsa.PrivateKey
	dataDir            string
	permConfig         *types.PermissionConfig
	contract           ptype.InitService
	backend            ptype.Backend
	eeaFlag            bool
	useDns             bool
	isRaft             bool
	startWaitGroup     *sync.WaitGroup // waitgroup to make sure all dependencies are ready before we start the service
	errorChan          chan error      // channel to capture error when starting aysnc
	networkInitialized bool
	controlService     ptype.ControlService
}

// Create a service instance for permissioning
//
// Permission Service depends on the following:
// 1. EthService to be ready
// 2. Downloader to sync up blocks
// 3. InProc RPC server to be ready
func NewQuorumPermissionCtrl(stack *node.Node, pconfig *types.PermissionConfig, eeaFlag, useDns bool) (*PermissionCtrl, error) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	p := &PermissionCtrl{
		node:           stack,
		key:            stack.GetNodeKey(),
		dataDir:        stack.DataDir(),
		permConfig:     pconfig,
		startWaitGroup: wg,
		errorChan:      make(chan error),
		eeaFlag:        eeaFlag,
		useDns:         useDns,
		isRaft:         false,
	}

	p.populateBackEnd()
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
	log.Debug("permission service: starting")
	go func() {
		log.Debug("permission service: starting async")
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

func NewPermissionContractService(ethClnt bind.ContractBackend, eeaFlag bool, key *ecdsa.PrivateKey,
	permConfig *types.PermissionConfig, isRaft, useDns bool) ptype.InitService {

	contractBackEnd := ptype.ContractBackend{
		EthClnt:    ethClnt,
		Key:        key,
		PermConfig: permConfig,
	}

	if eeaFlag {
		return &eea.Init{
			Backend: contractBackEnd,
			IsRaft:  isRaft,
			UseDns:  useDns,
		}
	}
	return &basic.Init{
		Backend: contractBackEnd,
	}
}

func (p *PermissionCtrl) NewPermissionRoleService(transactOpts *bind.TransactOpts) (ptype.RoleService, error) {
	roleBackend := ptype.ContractBackend{EthClnt: p.ethClnt, Key: p.key, PermConfig: p.permConfig}

	switch p.eeaFlag {
	case true:
		backEnd, err := getEeaBackEndWithTransactOpts(roleBackend, p.isRaft, p.useDns, transactOpts)
		if err != nil {
			return nil, err
		}
		return &eea.Role{Backend: backEnd}, nil
	default:
		backEnd, err := getBasicBackEndWithTransactOpts(roleBackend, transactOpts)
		if err != nil {
			return nil, err
		}
		return &basic.Role{Backend: backEnd}, nil
	}
}

func (p *PermissionCtrl) NewPermissionAuditService() (ptype.AuditService, error) {
	auditBackend := ptype.ContractBackend{EthClnt: p.ethClnt, Key: p.key, PermConfig: p.permConfig}
	switch p.eeaFlag {
	case true:
		backEnd, err := getEeaBackEnd(auditBackend, p.isRaft, p.useDns)
		if err != nil {
			return nil, err
		}
		return &eea.Audit{Backend: backEnd}, nil
	default:
		backEnd, err := getBasicBackEnd(auditBackend)
		if err != nil {
			return nil, err
		}
		return &basic.Audit{Backend: backEnd}, nil
	}
}

func (p *PermissionCtrl) NewPermissionControlService() (ptype.ControlService, error) {
	controlBackend := ptype.ContractBackend{EthClnt: p.ethClnt, Key: p.key, PermConfig: p.permConfig}
	switch p.eeaFlag {
	case true:
		backEnd, err := getEeaBackEnd(controlBackend, p.isRaft, p.useDns)
		if err != nil {
			return nil, err
		}
		return &eea.Control{Backend: backEnd}, nil
	default:
		return &basic.Control{}, nil
	}
}

func (p *PermissionCtrl) GetPermissionInitialized() bool {
	return p.networkInitialized
}

func (p *PermissionCtrl) ConnectionAllowed(_enodeId, _ip string, _port, _raftPort uint16) (bool, error) {
	if p.controlService == nil {
		controlBackend := ptype.ContractBackend{EthClnt: p.ethClnt, Key: p.key, PermConfig: p.permConfig}
		switch p.eeaFlag {
		case true:
			backEnd, err := getEeaBackEnd(controlBackend, p.isRaft, p.useDns)
			if err != nil {
				return false, err
			}
			p.controlService = &eea.Control{Backend: backEnd}
		default:
			p.controlService = &basic.Control{}
		}
	}
	return p.controlService.ConnectionAllowed(_enodeId, _ip, _port, _raftPort)
}

func (p *PermissionCtrl) TransactionAllowed(_sender common.Address, _target common.Address, _value *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _payload []byte) (bool, error) {
	if p.controlService == nil {
		controlBackend := ptype.ContractBackend{EthClnt: p.ethClnt, Key: p.key, PermConfig: p.permConfig}
		switch p.eeaFlag {
		case true:
			backEnd, err := getEeaBackEnd(controlBackend, p.isRaft, p.useDns)
			if err != nil {
				return false, err
			}
			p.controlService = &eea.Control{Backend: backEnd}
		default:
			p.controlService = &basic.Control{}
		}
	}

	return p.controlService.TransactionAllowed(_sender, _target, _value, _gasPrice, _gasLimit, _payload)
}

func (p *PermissionCtrl) NewPermissionOrgService(transactOpts *bind.TransactOpts) (ptype.OrgService, error) {
	orgBackend := ptype.ContractBackend{EthClnt: p.ethClnt, Key: p.key, PermConfig: p.permConfig}
	switch p.eeaFlag {
	case true:
		backEnd, err := getEeaBackEndWithTransactOpts(orgBackend, p.isRaft, p.useDns, transactOpts)
		if err != nil {
			return nil, err
		}
		return &eea.Org{Backend: backEnd}, nil
	default:
		backEnd, err := getBasicBackEndWithTransactOpts(orgBackend, transactOpts)
		if err != nil {
			return nil, err
		}
		return &basic.Org{Backend: backEnd}, nil
	}
}

func (p *PermissionCtrl) NewPermissionNodeService(transactOpts *bind.TransactOpts) (ptype.NodeService, error) {
	nodeBackend := ptype.ContractBackend{EthClnt: p.ethClnt, Key: p.key, PermConfig: p.permConfig}
	switch p.eeaFlag {
	case true:
		backEnd, err := getEeaBackEndWithTransactOpts(nodeBackend, p.isRaft, p.useDns, transactOpts)
		if err != nil {
			return nil, err
		}
		return &eea.Node{Backend: backEnd}, nil
	default:
		backEnd, err := getBasicBackEndWithTransactOpts(nodeBackend, transactOpts)
		if err != nil {
			return nil, err
		}
		return &basic.Node{Backend: backEnd}, nil
	}
}

func (p *PermissionCtrl) NewPermissionAccountService(transactOpts *bind.TransactOpts) (ptype.AccountService, error) {
	accountBackend := ptype.ContractBackend{EthClnt: p.ethClnt, Key: p.key, PermConfig: p.permConfig}
	switch p.eeaFlag {
	case true:
		backEnd, err := getEeaBackEndWithTransactOpts(accountBackend, p.isRaft, p.useDns, transactOpts)
		if err != nil {
			return nil, err
		}
		return &eea.Account{Backend: backEnd}, nil
	default:
		backEnd, err := getBasicBackEndWithTransactOpts(accountBackend, transactOpts)
		if err != nil {
			return nil, err
		}
		return &basic.Account{Backend: backEnd}, nil
	}
}

func getBasicInterfaceContractSession(permInterfaceInstance *bb.PermInterface, contractAddress common.Address, backend bind.ContractBackend) (*bb.PermInterfaceSession, error) {
	if err := ptype.BindContract(&permInterfaceInstance, func() (interface{}, error) { return bb.NewPermInterface(contractAddress, backend) }); err != nil {
		return nil, err
	}
	ps := &bb.PermInterfaceSession{
		Contract: permInterfaceInstance,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	return ps, nil
}

func getBasicBackEndWithTransactOpts(contractBackend ptype.ContractBackend, transactOpts *bind.TransactOpts) (*basic.Basic, error) {
	basicBackend, err := getBasicBackEnd(contractBackend)
	if err != nil {
		return nil, err
	}
	basicBackend.PermInterfSession.TransactOpts = *transactOpts

	return basicBackend, nil
}

func getBasicBackEnd(contractBackend ptype.ContractBackend) (*basic.Basic, error) {
	basicBackend := basic.Basic{ContractBackend: contractBackend}
	ps, err := getBasicInterfaceContractSession(basicBackend.PermInterf, contractBackend.PermConfig.InterfAddress, contractBackend.EthClnt)
	if err != nil {
		return nil, err
	}
	basicBackend.PermInterfSession = ps
	return &basicBackend, nil
}

func getEeaBackEndWithTransactOpts(contractBackend ptype.ContractBackend, isRaft, useDns bool, transactOpts *bind.TransactOpts) (*eea.Eea, error) {
	eeaBackend, err := getEeaBackEnd(contractBackend, isRaft, useDns)
	if err != nil {
		return nil, err
	}
	eeaBackend.PermInterfSession.TransactOpts = *transactOpts
	return eeaBackend, nil
}

func getEeaBackEnd(contractBackend ptype.ContractBackend, isRaft, useDns bool) (*eea.Eea, error) {
	eeaBackend := eea.Eea{ContractBackend: contractBackend, IsRaft: isRaft, UseDns: useDns}
	ps, err := getEeaInterfaceContractSession(eeaBackend.PermInterf, contractBackend.PermConfig.InterfAddress, contractBackend.EthClnt)
	if err != nil {
		return nil, err
	}
	eeaBackend.PermInterfSession = ps
	return &eeaBackend, nil
}

func getEeaInterfaceContractSession(permInterfaceInstance *eb.PermInterface, contractAddress common.Address, backend bind.ContractBackend) (*eb.PermInterfaceSession, error) {
	if err := ptype.BindContract(&permInterfaceInstance, func() (interface{}, error) { return eb.NewPermInterface(contractAddress, backend) }); err != nil {
		return nil, err
	}
	ps := &eb.PermInterfaceSession{
		Contract: permInterfaceInstance,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	return ps, nil
}

func (p *PermissionCtrl) populateBackEnd() {
	backend := ptype.NewInterfaceBackend(p.node, false, p.dataDir)

	switch p.eeaFlag {
	case true:
		p.backend = &eea.Backend{
			Ib: *backend,
		}

	default:
		p.backend = &basic.Backend{
			Ib: *backend,
		}
	}
}

func (p *PermissionCtrl) updateBackEnd() {
	p.contract = NewPermissionContractService(p.ethClnt, p.eeaFlag, p.key, p.permConfig, p.isRaft, p.useDns)
	switch p.eeaFlag {
	case true:
		p.backend.(*eea.Backend).Contr = p.contract.(*eea.Init)
		p.backend.(*eea.Backend).Ib.SetIsRaft(p.isRaft)

	default:
		p.backend.(*basic.Backend).Contr = p.contract.(*basic.Init)
		p.backend.(*basic.Backend).Ib.SetIsRaft(p.isRaft)
	}
}
