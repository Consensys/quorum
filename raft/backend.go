package gethRaft

import (
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

type RaftService struct {
	blockchain     *core.BlockChain
	chainDb        ethdb.Database // Block chain database
	txMu           sync.Mutex
	txPool         *core.TxPool
	accountManager *accounts.Manager

	raftProtocolManager *ProtocolManager

	// we need an event mux to instantiate the blockchain
	eventMux *event.TypeMux
	minter   *minter
}

type RaftNodeInfo struct {
	ClusterSize int         `json:"clusterSize"`
	Genesis     common.Hash `json:"genesis"` // SHA3 hash of the host's genesis block
	Head        common.Hash `json:"head"`    // SHA3 hash of the host's best owned block
	Role        string      `json:"role"`
}

func New(ctx *node.ServiceContext, chainConfig *core.ChainConfig, strID string, e *eth.Ethereum) (*RaftService, error) {
	service := &RaftService{
		eventMux:       ctx.EventMux,
		chainDb:        e.ChainDb(),
		blockchain:     e.BlockChain(),
		txPool:         e.TxPool(),
		accountManager: e.AccountManager(),
	}

	ethProxy := e.GetProxy()

	var err error
	if service.raftProtocolManager, err = NewProtocolManager(strID,
		service.blockchain, service.eventMux, ethProxy.Downloader,
		ethProxy.GetBestRaftPeer); err != nil {
		return nil, err
	}

	service.minter = newMinter(chainConfig, service)

	return service, nil
}

func (service *RaftService) APIs() []rpc.API {
	return []rpc.API{
		{
			// startNode, (TODO) stopNode, version
			Namespace: "raft",
			Version:   "1.0",
			Service:   NewPublicRaftAPI(service),
			Public:    true,
		}, {
			// sendTransaction
			Namespace: "raft",
			Version:   "1.0",
			Service:   NewPublicTransactionPoolAPI(service),
			Public:    true,
		},
	}
}

func (service *RaftService) startMinting() {
	go service.minter.start()
}

func (service *RaftService) stopMinting() {
	service.minter.stop()
}

// Backend interface methods
func (service *RaftService) AccountManager() *accounts.Manager { return service.accountManager }
func (service *RaftService) BlockChain() *core.BlockChain      { return service.blockchain }
func (service *RaftService) ChainDb() ethdb.Database           { return service.chainDb }
func (service *RaftService) DappDb() ethdb.Database            { return nil }
func (service *RaftService) EventMux() *event.TypeMux          { return service.eventMux }
func (service *RaftService) TxPool() *core.TxPool              { return service.txPool }

// node.Service interface methods
func (service *RaftService) Protocols() []p2p.Protocol {
	return []p2p.Protocol{service.raftProtocolManager.protocol}
}

// Start implements node.Service, starting the background data propagation thread
// of the protocol.
func (service *RaftService) Start(*p2p.Server) error {
	service.raftProtocolManager.Start()
	return nil
}

func (service *RaftService) notifyRoleChange(roleC <-chan interface{}) {
	for {
		select {
		case role := <-roleC:
			intRole, ok := role.(int)

			if !ok {
				panic("Couldn't cast role to int")
			}

			if intRole == minterRole {
				service.startMinting()
			} else { // verifier
				service.stopMinting()
			}

			service.raftProtocolManager.role = intRole
		case <-service.raftProtocolManager.quitSync:
			return
		}
	}
}

// Stop implements node.Service, stopping the background data propagation thread
// of the protocol.
func (service *RaftService) Stop() error {
	service.blockchain.Stop()
	service.raftProtocolManager.Stop()
	service.stopMinting()
	service.eventMux.Stop()

	service.chainDb.Close()

	glog.V(logger.Info).Infoln("Raft stopped")
	return nil
}
