package permission

import (
	"crypto/ecdsa"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	node           *node.Node
	ethClnt        bind.ContractBackend
	eth            *eth.Ethereum
	key            *ecdsa.PrivateKey
	dataDir        string
	permConfig     *types.PermissionConfig
	contract       ptype.ContractService
	backend        ptype.Backend
	eeaFlag        bool
	startWaitGroup *sync.WaitGroup // waitgroup to make sure all dependencies are ready before we start the service
	errorChan      chan error      // channel to capture error when starting aysnc
}

// Create a service instance for permissioning
//
// Permission Service depends on the following:
// 1. EthService to be ready
// 2. Downloader to sync up blocks
// 3. InProc RPC server to be ready
func NewQuorumPermissionCtrl(stack *node.Node, pconfig *types.PermissionConfig, eeaFlag bool) (*PermissionCtrl, error) {
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
	permConfig *types.PermissionConfig) ptype.ContractService {
	if eeaFlag {
		return &eea.Contract{EthClnt: ethClnt, Key: key, PermConfig: permConfig}
	}
	return &basic.Contract{EthClnt: ethClnt, Key: key, PermConfig: permConfig}
}

func NewPermissionContractServiceForApi(p *PermissionCtrl, transactOpts *bind.TransactOpts) ptype.ContractService {
	switch p.eeaFlag {
	case true:
		pc := p.contract.(*eea.Contract)
		ps := &eb.PermInterfaceSession{
			Contract: pc.PermInterf,
			CallOpts: bind.CallOpts{
				Pending: true,
			},
			TransactOpts: *transactOpts,
		}
		return &eea.Contract{PermInterfSession: ps, PermConfig: p.permConfig}

	default:
		pc := p.contract.(*basic.Contract)
		ps := &bb.PermInterfaceSession{
			Contract: pc.PermInterf,
			CallOpts: bind.CallOpts{
				Pending: true,
			},
			TransactOpts: *transactOpts,
		}
		return &basic.Contract{PermInterfSession: ps, PermConfig: p.permConfig}
	}
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
	isRaft := p.eth.BlockChain().Config().Istanbul == nil && p.eth.BlockChain().Config().Clique == nil
	p.contract = NewPermissionContractService(p.ethClnt, p.eeaFlag, p.key, p.permConfig)
	switch p.eeaFlag {
	case true:
		p.backend.(*eea.Backend).Contr = p.contract.(*eea.Contract)
		p.backend.(*eea.Backend).Ib.SetIsRaft(isRaft)
	default:
		p.backend.(*basic.Backend).Contr = p.contract.(*basic.Contract)
		p.backend.(*basic.Backend).Ib.SetIsRaft(isRaft)
	}
}
