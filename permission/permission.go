package permission

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/permission/bind"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/raft"
	"github.com/ethereum/go-ethereum/rpc"
)

// to signal all watches when service is stopped
type StopEvent struct {
}

type NodeOperation uint8

const (
	NodeAdd NodeOperation = iota
	NodeDelete
)

type PermissionCtrl struct {
	Node       *node.Node
	EthClnt    bind.ContractBackend
	Eth        *eth.Ethereum
	Key        *ecdsa.PrivateKey
	DataDir    string
	PermConfig *types.PermissionConfig

	contract *PermissionContractService

	eeaFlag        bool
	StartWaitGroup *sync.WaitGroup // waitgroup to make sure all dependencies are ready before we start the service
	StopFeed       event.Feed      // broadcasting stopEvent when service is being stopped
	ErrorChan      chan error      // channel to capture error when starting aysnc

	mux sync.Mutex
}

// function reads the permissions config file passed and populates the
// config structure accordingly
func ParsePermissionConfig(dir string) (types.PermissionConfig, error) {
	fullPath := filepath.Join(dir, params.PERMISSION_MODEL_CONFIG)
	f, err := os.Open(fullPath)
	if err != nil {
		log.Error("can't open file", "file", fullPath, "error", err)
		return types.PermissionConfig{}, err
	}
	defer func() {
		_ = f.Close()
	}()

	var permConfig types.PermissionConfig
	blob, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Error("error reading file", "err", err, "file", fullPath)
	}

	err = json.Unmarshal(blob, &permConfig)
	if err != nil {
		log.Error("error unmarshalling the file", "err", err, "file", fullPath)
	}

	if len(permConfig.Accounts) == 0 {
		return types.PermissionConfig{}, fmt.Errorf("no accounts given in %s. Network cannot boot up", params.PERMISSION_MODEL_CONFIG)
	}
	if permConfig.SubOrgDepth.Cmp(big.NewInt(0)) == 0 || permConfig.SubOrgBreadth.Cmp(big.NewInt(0)) == 0 {
		return types.PermissionConfig{}, fmt.Errorf("sub org breadth depth not passed in %s. Network cannot boot up", params.PERMISSION_MODEL_CONFIG)
	}
	if permConfig.IsEmpty() {
		return types.PermissionConfig{}, fmt.Errorf("missing contract addresses in %s", params.PERMISSION_MODEL_CONFIG)
	}

	return permConfig, nil
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
		Node:           stack,
		Key:            stack.GetNodeKey(),
		DataDir:        stack.DataDir(),
		PermConfig:     pconfig,
		StartWaitGroup: wg,
		ErrorChan:      make(chan error),
		eeaFlag:        eeaFlag,
	}

	stopChan, stopSubscription := p.SubscribeStopEvent()
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


func (p *PermissionCtrl) BlockChain() *core.BlockChain {
	return p.Eth.BlockChain()
}

func (p *PermissionCtrl) Ethereum() *eth.Ethereum {
	return p.Eth
}

func (p *PermissionCtrl) PermissionConfig() *types.PermissionConfig {
	return p.PermConfig
}

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (p *PermissionCtrl) AfterStart() error {
	log.Debug("permission service: binding contracts")
	err := <-p.ErrorChan // capture any error happened during AsyncStart. Also wait here if AsyncStart is not yet finish
	if err != nil {
		return err
	}
	p.contract.AfterStart()

	// populate the initial list of permissioned nodes and account accesses
	if err := p.PopulateInitPermissions(params.DEFAULT_ORGCACHE_SIZE, params.DEFAULT_ROLECACHE_SIZE,
		params.DEFAULT_NODECACHE_SIZE, params.DEFAULT_ACCOUNTCACHE_SIZE); err != nil {
		return fmt.Errorf("populateInitPermissions failed: %v", err)
	}

	// set the default access to ReadOnly
	types.SetDefaults(p.PermConfig.NwAdminRole, p.PermConfig.OrgAdminRole)

	for _, f := range []func() error{
		p.MonitorQIP714Block,       // monitor block number to activate new permissions controls
		p.ManageOrgPermissions,     // monitor org management related events
		p.ManageNodePermissions,    // monitor org  level Node management events
		p.ManageRolePermissions,    // monitor org level role management events
		p.ManageAccountPermissions, // monitor org level account management events
	} {
		if err := f(); err != nil {
			return err
		}
	}
	log.Info("permission service: is now ready")
	return nil
}

// start service asynchronously due to dependencies
func (p *PermissionCtrl) AsyncStart() {
	var ethereum *eth.Ethereum
	// will be blocked here until Node is up
	if err := p.Node.Service(&ethereum); err != nil {
		p.ErrorChan <- fmt.Errorf("dependent ethereum service not started")
		return
	}
	defer func() {
		p.ErrorChan <- nil
	}()
	// for cases where the Node is joining an existing network, permission service
	// can be brought up only after block syncing is complete. This function
	// waits for block syncing before the starting permissions
	p.StartWaitGroup.Add(1)
	go func(_wg *sync.WaitGroup) {
		log.Debug("permission service: waiting for downloader")
		stopChan, stopSubscription := p.SubscribeStopEvent()
		pollingTicker := time.NewTicker(10 * time.Millisecond)
		defer func(start time.Time) {
			log.Debug("permission service: downloader completed", "took", time.Since(start))
			stopSubscription.Unsubscribe()
			pollingTicker.Stop()
			_wg.Done()
		}(time.Now())
		for {
			select {
			case <-pollingTicker.C:
				if types.GetSyncStatus() && !ethereum.Downloader().Synchronising() {
					return
				}
			case <-stopChan:
				return
			}
		}
	}(p.StartWaitGroup) // wait for downloader to sync if any

	log.Debug("permission service: waiting for all dependencies to be ready")
	p.StartWaitGroup.Wait()
	client, err := p.Node.Attach()
	if err != nil {
		p.ErrorChan <- fmt.Errorf("unable to create rpc client: %v", err)
		return
	}
	p.EthClnt = ethclient.NewClient(client)
	p.Eth = ethereum
	p.contract = &PermissionContractService{EthClnt: p.EthClnt, EeaFlag: p.eeaFlag, Key: p.Key, PermConfig: p.PermConfig}
}

func (p *PermissionCtrl) Start(srvr *p2p.Server) error {
	log.Debug("permission service: starting")
	go func() {
		log.Debug("permission service: starting async")
		p.AsyncStart()
	}()
	return nil
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

func (p *PermissionCtrl) Protocols() []p2p.Protocol {
	return []p2p.Protocol{}
}

func (p *PermissionCtrl) Stop() error {
	log.Info("permission service: stopping")
	p.StopFeed.Send(StopEvent{})
	log.Info("permission service: stopped")
	return nil
}

// monitors QIP714Block and set default access
func (p *PermissionCtrl) MonitorQIP714Block() error {
	// if QIP714block is not given, set the default access
	// to readonly
	if p.Eth.BlockChain().Config().QIP714Block == nil {
		types.SetDefaultAccess()
		return nil
	}
	//QIP714block is given, monitor block count
	go func() {
		chainHeadCh := make(chan core.ChainHeadEvent, 1)
		headSub := p.Eth.BlockChain().SubscribeChainHeadEvent(chainHeadCh)
		defer headSub.Unsubscribe()
		stopChan, stopSubscription := p.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case head := <-chainHeadCh:
				if p.Eth.BlockChain().Config().IsQIP714(head.Block.Number()) {
					types.SetDefaultAccess()
					return
				}
			case <-stopChan:
				return
			}
		}
	}()
	return nil
}

// monitors org management related events happening via smart contracts
// and updates cache accordingly
func (p *PermissionCtrl) ManageOrgPermissions() error {
	if p.eeaFlag{
		return p.ManageOrgPermissionsE()
	}
	return p.ManageOrgPermissionsBasic()
}

func (p *PermissionCtrl) ManageOrgPermissionsE() error {
	chPendingApproval := make(chan *permission.EeaOrgManagerOrgPendingApproval, 1)
	chOrgApproved := make(chan *permission.EeaOrgManagerOrgApproved, 1)
	chOrgSuspended := make(chan *permission.EeaOrgManagerOrgSuspended, 1)
	chOrgReactivated := make(chan *permission.EeaOrgManagerOrgSuspensionRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.contract.PermOrgE.EeaOrgManagerFilterer.WatchOrgPendingApproval(opts, chPendingApproval); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.contract.PermOrgE.EeaOrgManagerFilterer.WatchOrgApproved(opts, chOrgApproved); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.contract.PermOrgE.EeaOrgManagerFilterer.WatchOrgSuspended(opts, chOrgSuspended); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.contract.PermOrgE.EeaOrgManagerFilterer.WatchOrgSuspensionRevoked(opts, chOrgReactivated); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtPendingApproval := <-chPendingApproval:
				types.OrgInfoMap.UpsertOrg(evtPendingApproval.OrgId, evtPendingApproval.PorgId, evtPendingApproval.UltParent, evtPendingApproval.Level, types.OrgStatus(evtPendingApproval.Status.Uint64()))

			case evtOrgApproved := <-chOrgApproved:
				types.OrgInfoMap.UpsertOrg(evtOrgApproved.OrgId, evtOrgApproved.PorgId, evtOrgApproved.UltParent, evtOrgApproved.Level, types.OrgApproved)

			case evtOrgSuspended := <-chOrgSuspended:
				types.OrgInfoMap.UpsertOrg(evtOrgSuspended.OrgId, evtOrgSuspended.PorgId, evtOrgSuspended.UltParent, evtOrgSuspended.Level, types.OrgSuspended)

			case evtOrgReactivated := <-chOrgReactivated:
				types.OrgInfoMap.UpsertOrg(evtOrgReactivated.OrgId, evtOrgReactivated.PorgId, evtOrgReactivated.UltParent, evtOrgReactivated.Level, types.OrgApproved)
			case <-stopChan:
				log.Info("quit org contract watch")
				return
			}
		}
	}()
	return nil
}
func (p *PermissionCtrl) ManageOrgPermissionsBasic() error {
	chPendingApproval := make(chan *permission.OrgManagerOrgPendingApproval, 1)
	chOrgApproved := make(chan *permission.OrgManagerOrgApproved, 1)
	chOrgSuspended := make(chan *permission.OrgManagerOrgSuspended, 1)
	chOrgReactivated := make(chan *permission.OrgManagerOrgSuspensionRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.contract.PermOrg.OrgManagerFilterer.WatchOrgPendingApproval(opts, chPendingApproval); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.contract.PermOrg.OrgManagerFilterer.WatchOrgApproved(opts, chOrgApproved); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.contract.PermOrg.OrgManagerFilterer.WatchOrgSuspended(opts, chOrgSuspended); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.contract.PermOrg.OrgManagerFilterer.WatchOrgSuspensionRevoked(opts, chOrgReactivated); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtPendingApproval := <-chPendingApproval:
				types.OrgInfoMap.UpsertOrg(evtPendingApproval.OrgId, evtPendingApproval.PorgId, evtPendingApproval.UltParent, evtPendingApproval.Level, types.OrgStatus(evtPendingApproval.Status.Uint64()))

			case evtOrgApproved := <-chOrgApproved:
				types.OrgInfoMap.UpsertOrg(evtOrgApproved.OrgId, evtOrgApproved.PorgId, evtOrgApproved.UltParent, evtOrgApproved.Level, types.OrgApproved)

			case evtOrgSuspended := <-chOrgSuspended:
				types.OrgInfoMap.UpsertOrg(evtOrgSuspended.OrgId, evtOrgSuspended.PorgId, evtOrgSuspended.UltParent, evtOrgSuspended.Level, types.OrgSuspended)

			case evtOrgReactivated := <-chOrgReactivated:
				types.OrgInfoMap.UpsertOrg(evtOrgReactivated.OrgId, evtOrgReactivated.PorgId, evtOrgReactivated.UltParent, evtOrgReactivated.Level, types.OrgApproved)
			case <-stopChan:
				log.Info("quit org contract watch")
				return
			}
		}
	}()
	return nil
}

func (p *PermissionCtrl) SubscribeStopEvent() (chan StopEvent, event.Subscription) {
	c := make(chan StopEvent)
	s := p.StopFeed.Subscribe(c)
	return c, s
}

// Monitors Node management events and updates cache accordingly
func (p *PermissionCtrl) ManageNodePermissions() error {
	if p.eeaFlag{
		return p.ManageNodePermissionsE()
	}
	return p.ManageNodePermissionsBasic()
}

func (p *PermissionCtrl) ManageNodePermissionsBasic() error {
	chNodeApproved := make(chan *permission.NodeManagerNodeApproved, 1)
	chNodeProposed := make(chan *permission.NodeManagerNodeProposed, 1)
	chNodeDeactivated := make(chan *permission.NodeManagerNodeDeactivated, 1)
	chNodeActivated := make(chan *permission.NodeManagerNodeActivated, 1)
	chNodeBlacklisted := make(chan *permission.NodeManagerNodeBlacklisted)
	chNodeRecoveryInit := make(chan *permission.NodeManagerNodeRecoveryInitiated, 1)
	chNodeRecoveryDone := make(chan *permission.NodeManagerNodeRecoveryCompleted, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.contract.PermNode.NodeManagerFilterer.WatchNodeApproved(opts, chNodeApproved); err != nil {
		return fmt.Errorf("failed WatchNodeApproved: %v", err)
	}

	if _, err := p.contract.PermNode.NodeManagerFilterer.WatchNodeProposed(opts, chNodeProposed); err != nil {
		return fmt.Errorf("failed WatchNodeProposed: %v", err)
	}

	if _, err := p.contract.PermNode.NodeManagerFilterer.WatchNodeDeactivated(opts, chNodeDeactivated); err != nil {
		return fmt.Errorf("failed NodeDeactivated: %v", err)
	}
	if _, err := p.contract.PermNode.NodeManagerFilterer.WatchNodeActivated(opts, chNodeActivated); err != nil {
		return fmt.Errorf("failed WatchNodeActivated: %v", err)
	}

	if _, err := p.contract.PermNode.NodeManagerFilterer.WatchNodeBlacklisted(opts, chNodeBlacklisted); err != nil {
		return fmt.Errorf("failed NodeBlacklisting: %v", err)
	}

	if _, err := p.contract.PermNode.NodeManagerFilterer.WatchNodeRecoveryInitiated(opts, chNodeRecoveryInit); err != nil {
		return fmt.Errorf("failed NodeRecoveryInitiated: %v", err)
	}

	if _, err := p.contract.PermNode.NodeManagerFilterer.WatchNodeRecoveryCompleted(opts, chNodeRecoveryDone); err != nil {
		return fmt.Errorf("failed NodeRecoveryCompleted: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtNodeApproved := <-chNodeApproved:
				p.UpdatePermissionedNodes(evtNodeApproved.EnodeId, NodeAdd)
				types.NodeInfoMap.UpsertNode(evtNodeApproved.OrgId, evtNodeApproved.EnodeId, types.NodeApproved)

			case evtNodeProposed := <-chNodeProposed:
				types.NodeInfoMap.UpsertNode(evtNodeProposed.OrgId, evtNodeProposed.EnodeId, types.NodePendingApproval)

			case evtNodeDeactivated := <-chNodeDeactivated:
				p.UpdatePermissionedNodes(evtNodeDeactivated.EnodeId, NodeDelete)
				types.NodeInfoMap.UpsertNode(evtNodeDeactivated.OrgId, evtNodeDeactivated.EnodeId, types.NodeDeactivated)

			case evtNodeActivated := <-chNodeActivated:
				p.UpdatePermissionedNodes(evtNodeActivated.EnodeId, NodeAdd)
				types.NodeInfoMap.UpsertNode(evtNodeActivated.OrgId, evtNodeActivated.EnodeId, types.NodeApproved)

			case evtNodeBlacklisted := <-chNodeBlacklisted:
				types.NodeInfoMap.UpsertNode(evtNodeBlacklisted.OrgId, evtNodeBlacklisted.EnodeId, types.NodeBlackListed)
				p.UpdateDisallowedNodes(evtNodeBlacklisted.EnodeId, NodeAdd)
				p.UpdatePermissionedNodes(evtNodeBlacklisted.EnodeId, NodeDelete)

			case evtNodeRecoveryInit := <-chNodeRecoveryInit:
				types.NodeInfoMap.UpsertNode(evtNodeRecoveryInit.OrgId, evtNodeRecoveryInit.EnodeId, types.NodeRecoveryInitiated)

			case evtNodeRecoveryDone := <-chNodeRecoveryDone:
				types.NodeInfoMap.UpsertNode(evtNodeRecoveryDone.OrgId, evtNodeRecoveryDone.EnodeId, types.NodeApproved)
				p.UpdateDisallowedNodes(evtNodeRecoveryDone.EnodeId, NodeDelete)
				p.UpdatePermissionedNodes(evtNodeRecoveryDone.EnodeId, NodeAdd)

			case <-stopChan:
				log.Info("quit Node contract watch")
				return
			}
		}
	}()
	return nil
}
func (p *PermissionCtrl) ManageNodePermissionsE() error {
	chNodeApproved := make(chan *permission.EeaNodeManagerNodeApproved, 1)
	chNodeProposed := make(chan *permission.EeaNodeManagerNodeProposed, 1)
	chNodeDeactivated := make(chan *permission.EeaNodeManagerNodeDeactivated, 1)
	chNodeActivated := make(chan *permission.EeaNodeManagerNodeActivated, 1)
	chNodeBlacklisted := make(chan *permission.EeaNodeManagerNodeBlacklisted)
	chNodeRecoveryInit := make(chan *permission.EeaNodeManagerNodeRecoveryInitiated, 1)
	chNodeRecoveryDone := make(chan *permission.EeaNodeManagerNodeRecoveryCompleted, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.contract.PermNodeE.EeaNodeManagerFilterer.WatchNodeApproved(opts, chNodeApproved); err != nil {
		return fmt.Errorf("failed WatchNodeApproved: %v", err)
	}

	if _, err := p.contract.PermNodeE.EeaNodeManagerFilterer.WatchNodeProposed(opts, chNodeProposed); err != nil {
		return fmt.Errorf("failed WatchNodeProposed: %v", err)
	}

	if _, err := p.contract.PermNodeE.EeaNodeManagerFilterer.WatchNodeDeactivated(opts, chNodeDeactivated); err != nil {
		return fmt.Errorf("failed NodeDeactivated: %v", err)
	}
	if _, err := p.contract.PermNodeE.EeaNodeManagerFilterer.WatchNodeActivated(opts, chNodeActivated); err != nil {
		return fmt.Errorf("failed WatchNodeActivated: %v", err)
	}

	if _, err := p.contract.PermNodeE.EeaNodeManagerFilterer.WatchNodeBlacklisted(opts, chNodeBlacklisted); err != nil {
		return fmt.Errorf("failed NodeBlacklisting: %v", err)
	}

	if _, err := p.contract.PermNodeE.EeaNodeManagerFilterer.WatchNodeRecoveryInitiated(opts, chNodeRecoveryInit); err != nil {
		return fmt.Errorf("failed NodeRecoveryInitiated: %v", err)
	}

	if _, err := p.contract.PermNodeE.EeaNodeManagerFilterer.WatchNodeRecoveryCompleted(opts, chNodeRecoveryDone); err != nil {
		return fmt.Errorf("failed NodeRecoveryCompleted: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtNodeApproved := <-chNodeApproved:
				p.UpdatePermissionedNodes(types.GetNodeUrl(evtNodeApproved.EnodeId, string(evtNodeApproved.Ip[:]), evtNodeApproved.Port, evtNodeApproved.Raftport), NodeAdd)
				types.NodeInfoMap.UpsertNode(evtNodeApproved.OrgId, types.GetNodeUrl(evtNodeApproved.EnodeId, string(evtNodeApproved.Ip[:]), evtNodeApproved.Port, evtNodeApproved.Raftport), types.NodeApproved)

			case evtNodeProposed := <-chNodeProposed:
				types.NodeInfoMap.UpsertNode(evtNodeProposed.OrgId, types.GetNodeUrl(evtNodeProposed.EnodeId, string(evtNodeProposed.Ip[:]), evtNodeProposed.Port, evtNodeProposed.Raftport), types.NodePendingApproval)

			case evtNodeDeactivated := <-chNodeDeactivated:
				p.UpdatePermissionedNodes(types.GetNodeUrl(evtNodeDeactivated.EnodeId, string(evtNodeDeactivated.Ip[:]), evtNodeDeactivated.Port, evtNodeDeactivated.Raftport), NodeDelete)
				types.NodeInfoMap.UpsertNode(evtNodeDeactivated.OrgId, types.GetNodeUrl(evtNodeDeactivated.EnodeId, string(evtNodeDeactivated.Ip[:]), evtNodeDeactivated.Port, evtNodeDeactivated.Raftport), types.NodeDeactivated)

			case evtNodeActivated := <-chNodeActivated:
				p.UpdatePermissionedNodes(types.GetNodeUrl(evtNodeActivated.EnodeId, string(evtNodeActivated.Ip[:]), evtNodeActivated.Port, evtNodeActivated.Raftport), NodeAdd)
				types.NodeInfoMap.UpsertNode(evtNodeActivated.OrgId, types.GetNodeUrl(evtNodeActivated.EnodeId, string(evtNodeActivated.Ip[:]), evtNodeActivated.Port, evtNodeActivated.Raftport), types.NodeApproved)

			case evtNodeBlacklisted := <-chNodeBlacklisted:
				types.NodeInfoMap.UpsertNode(evtNodeBlacklisted.OrgId, types.GetNodeUrl(evtNodeBlacklisted.EnodeId, string(evtNodeBlacklisted.Ip[:]), evtNodeBlacklisted.Port, evtNodeBlacklisted.Raftport), types.NodeBlackListed)
				p.UpdateDisallowedNodes(types.GetNodeUrl(evtNodeBlacklisted.EnodeId, string(evtNodeBlacklisted.Ip[:]), evtNodeBlacklisted.Port, evtNodeBlacklisted.Raftport), NodeAdd)
				p.UpdatePermissionedNodes(types.GetNodeUrl(evtNodeBlacklisted.EnodeId, string(evtNodeBlacklisted.Ip[:]), evtNodeBlacklisted.Port, evtNodeBlacklisted.Raftport), NodeDelete)

			case evtNodeRecoveryInit := <-chNodeRecoveryInit:
				types.NodeInfoMap.UpsertNode(evtNodeRecoveryInit.OrgId, types.GetNodeUrl(evtNodeRecoveryInit.EnodeId, string(evtNodeRecoveryInit.Ip[:]), evtNodeRecoveryInit.Port, evtNodeRecoveryInit.Raftport), types.NodeRecoveryInitiated)

			case evtNodeRecoveryDone := <-chNodeRecoveryDone:
				types.NodeInfoMap.UpsertNode(evtNodeRecoveryDone.OrgId, types.GetNodeUrl(evtNodeRecoveryDone.EnodeId, string(evtNodeRecoveryDone.Ip[:]), evtNodeRecoveryDone.Port, evtNodeRecoveryDone.Raftport), types.NodeApproved)
				p.UpdateDisallowedNodes(types.GetNodeUrl(evtNodeRecoveryDone.EnodeId, string(evtNodeRecoveryDone.Ip[:]), evtNodeRecoveryDone.Port, evtNodeRecoveryDone.Raftport), NodeDelete)
				p.UpdatePermissionedNodes(types.GetNodeUrl(evtNodeRecoveryDone.EnodeId, string(evtNodeRecoveryDone.Ip[:]), evtNodeRecoveryDone.Port, evtNodeRecoveryDone.Raftport), NodeAdd)

			case <-stopChan:
				log.Info("quit Node contract watch")
				return
			}
		}
	}()
	return nil
}

// adds or deletes and entry from a given file
func (p *PermissionCtrl) UpdateFile(fileName, enodeId string, operation NodeOperation, createFile bool) {
	// Load the nodes from the config file
	var nodeList []string
	index := 0
	// if createFile is false means the file is already existing. read the file
	if !createFile {
		blob, err := ioutil.ReadFile(fileName)
		if err != nil && !createFile {
			log.Error("Failed to access the file", "fileName", fileName, "err", err)
			return
		}

		if err := json.Unmarshal(blob, &nodeList); err != nil {
			log.Error("Failed to load nodes list from file", "fileName", fileName, "err", err)
			return
		}

		// logic to update the permissioned-nodes.json file based on action

		recExists := false
		for i, eid := range nodeList {
			if eid == enodeId {
				index = i
				recExists = true
				break
			}
		}
		if (operation == NodeAdd && recExists) || (operation == NodeDelete && !recExists) {
			return
		}
	}
	if operation == NodeAdd {
		nodeList = append(nodeList, enodeId)
	} else {
		nodeList = append(nodeList[:index], nodeList[index+1:]...)
	}
	blob, _ := json.Marshal(nodeList)

	p.mux.Lock()
	defer p.mux.Unlock()

	if err := ioutil.WriteFile(fileName, blob, 0644); err != nil {
		log.Error("Error writing new Node info to file", "fileName", fileName, "err", err)
	}
}

// updates Node information in the permissioned-nodes.json file based on Node
// management activities in smart contract
func (p *PermissionCtrl) UpdatePermissionedNodes(enodeId string, operation NodeOperation) {
	log.Debug("updatePermissionedNodes", "DataDir", p.DataDir, "file", params.PERMISSIONED_CONFIG)

	path := filepath.Join(p.DataDir, params.PERMISSIONED_CONFIG)
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for permissioned-nodes.json file. This is because 'permissioned' flag is specified but no permissioned-nodes.json file is present", "err", err)
		return
	}

	p.UpdateFile(path, enodeId, operation, false)
	if operation == NodeDelete {
		p.DisconnectNode(enodeId)
	}
}

//this function populates the black listed Node information into the disallowed-nodes.json file
func (p *PermissionCtrl) UpdateDisallowedNodes(url string, operation NodeOperation) {
	log.Debug("updateDisallowedNodes", "DataDir", p.DataDir, "file", params.BLACKLIST_CONFIG)

	fileExists := true
	path := filepath.Join(p.DataDir, params.BLACKLIST_CONFIG)
	// Check if the file is existing. If the file is not existing create the file
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for disallowed-nodes.json file", "err", err)
		if _, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644); err != nil {
			log.Error("Failed to create disallowed-nodes.json file", "err", err)
			return
		}
		fileExists = false
	}

	if fileExists {
		p.UpdateFile(path, url, operation, false)
	} else {
		p.UpdateFile(path, url, operation, true)
	}
}

// Monitors account access related events and updates the cache accordingly
func (p *PermissionCtrl) ManageAccountPermissions() error {
	if p.eeaFlag{
		return p.ManageAccountPermissionsE()
	}
	return p.ManageAccountPermissionsBasic()
}

func (p *PermissionCtrl) ManageAccountPermissionsE() error {
	chAccessModified := make(chan *permission.EeaAcctManagerAccountAccessModified)
	chAccessRevoked := make(chan *permission.EeaAcctManagerAccountAccessRevoked)
	chStatusChanged := make(chan *permission.EeaAcctManagerAccountStatusChanged)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.contract.PermAcctE.EeaAcctManagerFilterer.WatchAccountAccessModified(opts, chAccessModified); err != nil {
		return fmt.Errorf("failed AccountAccessModified: %v", err)
	}

	if _, err := p.contract.PermAcctE.EeaAcctManagerFilterer.WatchAccountAccessRevoked(opts, chAccessRevoked); err != nil {
		return fmt.Errorf("failed AccountAccessRevoked: %v", err)
	}

	if _, err := p.contract.PermAcctE.EeaAcctManagerFilterer.WatchAccountStatusChanged(opts, chStatusChanged); err != nil {
		return fmt.Errorf("failed AccountStatusChanged: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtAccessModified := <-chAccessModified:
				types.AcctInfoMap.UpsertAccount(evtAccessModified.OrgId, evtAccessModified.RoleId, evtAccessModified.Account, evtAccessModified.OrgAdmin, types.AcctStatus(int(evtAccessModified.Status.Uint64())))

			case evtAccessRevoked := <-chAccessRevoked:
				types.AcctInfoMap.UpsertAccount(evtAccessRevoked.OrgId, evtAccessRevoked.RoleId, evtAccessRevoked.Account, evtAccessRevoked.OrgAdmin, types.AcctActive)

			case evtStatusChanged := <-chStatusChanged:
				if ac, err := types.AcctInfoMap.GetAccount(evtStatusChanged.Account); ac != nil {
					types.AcctInfoMap.UpsertAccount(evtStatusChanged.OrgId, ac.RoleId, evtStatusChanged.Account, ac.IsOrgAdmin, types.AcctStatus(int(evtStatusChanged.Status.Uint64())))
				} else {
					log.Info("error fetching account information", "err", err)
				}
			case <-stopChan:
				log.Info("quit account contract watch")
				return
			}
		}
	}()
	return nil
}
func (p *PermissionCtrl) ManageAccountPermissionsBasic() error {
	chAccessModified := make(chan *permission.AcctManagerAccountAccessModified)
	chAccessRevoked := make(chan *permission.AcctManagerAccountAccessRevoked)
	chStatusChanged := make(chan *permission.AcctManagerAccountStatusChanged)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.contract.PermAcct.AcctManagerFilterer.WatchAccountAccessModified(opts, chAccessModified); err != nil {
		return fmt.Errorf("failed AccountAccessModified: %v", err)
	}

	if _, err := p.contract.PermAcct.AcctManagerFilterer.WatchAccountAccessRevoked(opts, chAccessRevoked); err != nil {
		return fmt.Errorf("failed AccountAccessRevoked: %v", err)
	}

	if _, err := p.contract.PermAcct.AcctManagerFilterer.WatchAccountStatusChanged(opts, chStatusChanged); err != nil {
		return fmt.Errorf("failed AccountStatusChanged: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtAccessModified := <-chAccessModified:
				types.AcctInfoMap.UpsertAccount(evtAccessModified.OrgId, evtAccessModified.RoleId, evtAccessModified.Account, evtAccessModified.OrgAdmin, types.AcctStatus(int(evtAccessModified.Status.Uint64())))

			case evtAccessRevoked := <-chAccessRevoked:
				types.AcctInfoMap.UpsertAccount(evtAccessRevoked.OrgId, evtAccessRevoked.RoleId, evtAccessRevoked.Account, evtAccessRevoked.OrgAdmin, types.AcctActive)

			case evtStatusChanged := <-chStatusChanged:
				if ac, err := types.AcctInfoMap.GetAccount(evtStatusChanged.Account); ac != nil {
					types.AcctInfoMap.UpsertAccount(evtStatusChanged.OrgId, ac.RoleId, evtStatusChanged.Account, ac.IsOrgAdmin, types.AcctStatus(int(evtStatusChanged.Status.Uint64())))
				} else {
					log.Info("error fetching account information", "err", err)
				}
			case <-stopChan:
				log.Info("quit account contract watch")
				return
			}
		}
	}()
	return nil
}

// Disconnect the Node from the network
func (p *PermissionCtrl) DisconnectNode(enodeId string) {
	if p.Eth.BlockChain().Config().Istanbul == nil && p.Eth.BlockChain().Config().Clique == nil {
		var raftService *raft.RaftService
		if err := p.Node.Service(&raftService); err == nil {
			raftApi := raft.NewPublicRaftAPI(raftService)

			//get the raftId for the given enodeId
			raftId, err := raftApi.GetRaftId(enodeId)
			if err == nil {
				raftApi.RemovePeer(raftId)
			} else {
				log.Error("failed to get raft id", "err", err, "enodeId", enodeId)
			}
		}
	} else {
		// Istanbul  or clique - disconnect the peer
		server := p.Node.Server()
		if server != nil {
			node, err := enode.ParseV4(enodeId)
			if err == nil {
				server.RemovePeer(node)
			} else {
				log.Error("failed parse Node id", "err", err, "enodeId", enodeId)
			}
		}
	}

}

func (p *PermissionCtrl) InstantiateCache(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int) {
	// instantiate the cache objects for permissions
	types.OrgInfoMap = types.NewOrgCache(orgCacheSize)
	types.OrgInfoMap.PopulateCacheFunc(p.PopulateOrgToCache)

	types.RoleInfoMap = types.NewRoleCache(roleCacheSize)
	types.RoleInfoMap.PopulateCacheFunc(p.PopulateRoleToCache)

	types.NodeInfoMap = types.NewNodeCache(nodeCacheSize)
	types.NodeInfoMap.PopulateCacheFunc(p.PopulateNodeCache)
	types.NodeInfoMap.PopulateValidateFunc(p.PopulateNodeCacheAndValidate)

	types.AcctInfoMap = types.NewAcctCache(accountCacheSize)
	types.AcctInfoMap.PopulateCacheFunc(p.PopulateAccountToCache)
}

// Thus function checks if the initial network boot up status and if no
// populates permissions model with details from permission-config.json
func (p *PermissionCtrl) PopulateInitPermissions(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int) error {
	p.InstantiateCache(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize)
	networkInitialized, err := p.contract.GetNetworkBootStatus()
	if err != nil {
		// handle the scenario of no contract code.
		log.Warn("Failed to retrieve network boot status ", "err", err)
		return err
	}

	if !networkInitialized {
		if err := p.BootupNetwork(); err != nil {
			return err
		}
	} else {
		//populate orgs, nodes, roles and accounts from contract
		for _, f := range []func() error{
			p.PopulateOrgsFromContract,
			p.PopulateNodesFromContract,
			p.PopulateRolesFromContract,
			p.PopulateAccountsFromContract,
		} {
			if err := f(); err != nil {
				return err
			}
		}
	}
	return nil
}

// initialize the permissions model and populate initial values
func (p *PermissionCtrl) BootupNetwork() error {
	if _, err := p.contract.SetPolicy(p.PermConfig.NwAdminOrg, p.PermConfig.NwAdminRole, p.PermConfig.OrgAdminRole); err != nil {
		log.Error("bootupNetwork SetPolicy failed", "err", err)
		return err
	}
	if _, err := p.contract.Init(p.PermConfig.SubOrgBreadth, p.PermConfig.SubOrgDepth); err != nil {
		log.Error("bootupNetwork init failed", "err", err)
		return err
	}

	types.OrgInfoMap.UpsertOrg(p.PermConfig.NwAdminOrg, "", p.PermConfig.NwAdminOrg, big.NewInt(1), types.OrgApproved)
	types.RoleInfoMap.UpsertRole(p.PermConfig.NwAdminOrg, p.PermConfig.NwAdminRole, true, true, types.FullAccess, true)
	// populate the initial Node list from static-nodes.json
	if err := p.PopulateStaticNodesToContract(); err != nil {
		return err
	}
	// populate initial account access to full access
	if err := p.PopulateInitAccountAccess(); err != nil {
		return err
	}

	// update network status to boot completed
	if err := p.UpdateNetworkStatus(); err != nil {
		log.Error("failed to updated network boot status", "error", err)
		return err
	}
	return nil
}

// populates the account access details from contract into cache
func (p *PermissionCtrl) PopulateAccountsFromContract() error {
	if numberOfRoles, err := p.contract.GetNumberOfAccounts(); err == nil {
		iOrgNum := numberOfRoles.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if addr, org, role, status, orgAdmin, err := p.contract.GetAccountDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				types.AcctInfoMap.UpsertAccount(org, role, addr, orgAdmin, types.AcctStatus(int(status.Int64())))
			}
		}
	} else {
		return err
	}
	return nil
}

// populates the role details from contract into cache
func (p *PermissionCtrl) PopulateRolesFromContract() error {
	if numberOfRoles, err := p.contract.GetNumberOfRoles(); err == nil {
		iOrgNum := numberOfRoles.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if roleStruct, err := p.contract.GetRoleDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				types.RoleInfoMap.UpsertRole(roleStruct.OrgId, roleStruct.RoleId, roleStruct.Voter, roleStruct.Admin, types.AccessType(int(roleStruct.AccessType.Int64())), roleStruct.Active)
			}
		}

	} else {
		return err
	}
	return nil
}

// populates the Node details from contract into cache
func (p *PermissionCtrl) PopulateNodesFromContract() error {
	if numberOfNodes, err := p.contract.GetNumberOfNodes(); err == nil {
		iOrgNum := numberOfNodes.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if orgId, url, status, err := p.contract.GetNodeDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				types.NodeInfoMap.UpsertNode(orgId, url, types.NodeStatus(int(status.Int64())))
			}
		}
	} else {
		return err
	}
	return nil
}

// populates the org details from contract into cache
func (p *PermissionCtrl) PopulateOrgsFromContract() error {

	if numberOfOrgs, err := p.contract.GetNumberOfOrgs(); err == nil {
		iOrgNum := numberOfOrgs.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if orgId, porgId, ultParent, level, status, err := p.contract.GetOrgInfo(big.NewInt(int64(k))); err == nil {
				types.OrgInfoMap.UpsertOrg(orgId, porgId, ultParent, level, types.OrgStatus(int(status.Int64())))
			}
		}
	} else {
		return err
	}
	return nil
}

// Reads the Node list from static-nodes.json and populates into the contract
func (p *PermissionCtrl) PopulateStaticNodesToContract() error {
	nodes := p.Node.Server().Config.StaticNodes
	for _, node := range nodes {
		enodeId, ip, port, raftPort := node.NodeDetails()
		_, err := p.contract.AddAdminNode(enodeId, ip, port, raftPort)
		if err != nil {
			log.Warn("Failed to propose Node", "err", err, "enode", node.EnodeID())
			return err
		}
		types.NodeInfoMap.UpsertNode(p.PermConfig.NwAdminOrg, types.GetNodeUrl(enodeId, string(ip[:]), port, raftPort), 2)
	}
	return nil
}

// Invokes the initAccounts function of smart contract to set the initial
// set of accounts access to full access
func (p *PermissionCtrl) PopulateInitAccountAccess() error {
	for _, a := range p.PermConfig.Accounts {
		_, er := p.contract.AddAdminAccount(a)
		if er != nil {
			log.Warn("Error adding permission initial account list", "err", er, "account", a)
			return er
		}
		types.AcctInfoMap.UpsertAccount(p.PermConfig.NwAdminOrg, p.PermConfig.NwAdminRole, a, true, 2)
	}
	return nil
}

// updates network boot status to true
func (p *PermissionCtrl) UpdateNetworkStatus() error {
	_, err := p.contract.UpdateNetworkBootStatus()
	if err != nil {
		log.Warn("Failed to udpate network boot status ", "err", err)
		return err
	}
	return nil
}

// monitors role management related events and updated cache
func (p *PermissionCtrl) ManageRolePermissions() error {
	if p.eeaFlag {
		return p.ManageRolePermissionsE()
	}
	return p.ManageRolePermissionsBasic()
}

func (p *PermissionCtrl) ManageRolePermissionsE() error {
	chRoleCreated := make(chan *permission.EeaRoleManagerRoleCreated, 1)
	chRoleRevoked := make(chan *permission.EeaRoleManagerRoleRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.contract.PermRoleE.EeaRoleManagerFilterer.WatchRoleCreated(opts, chRoleCreated); err != nil {
		return fmt.Errorf("failed WatchRoleCreated: %v", err)
	}

	if _, err := p.contract.PermRoleE.EeaRoleManagerFilterer.WatchRoleRevoked(opts, chRoleRevoked); err != nil {
		return fmt.Errorf("failed WatchRoleRemoved: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtRoleCreated := <-chRoleCreated:
				types.RoleInfoMap.UpsertRole(evtRoleCreated.OrgId, evtRoleCreated.RoleId, evtRoleCreated.IsVoter, evtRoleCreated.IsAdmin, types.AccessType(int(evtRoleCreated.BaseAccess.Uint64())), true)

			case evtRoleRevoked := <-chRoleRevoked:
				if r, _ := types.RoleInfoMap.GetRole(evtRoleRevoked.OrgId, evtRoleRevoked.RoleId); r != nil {
					types.RoleInfoMap.UpsertRole(evtRoleRevoked.OrgId, evtRoleRevoked.RoleId, r.IsVoter, r.IsAdmin, r.Access, false)
				} else {
					log.Error("Revoke role - cache is missing role", "org", evtRoleRevoked.OrgId, "role", evtRoleRevoked.RoleId)
				}
			case <-stopChan:
				log.Info("quit role contract watch")
				return
			}
		}
	}()
	return nil
}
func (p *PermissionCtrl) ManageRolePermissionsBasic() error {
	chRoleCreated := make(chan *permission.RoleManagerRoleCreated, 1)
	chRoleRevoked := make(chan *permission.RoleManagerRoleRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.contract.PermRole.RoleManagerFilterer.WatchRoleCreated(opts, chRoleCreated); err != nil {
		return fmt.Errorf("failed WatchRoleCreated: %v", err)
	}

	if _, err := p.contract.PermRole.RoleManagerFilterer.WatchRoleRevoked(opts, chRoleRevoked); err != nil {
		return fmt.Errorf("failed WatchRoleRemoved: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtRoleCreated := <-chRoleCreated:
				types.RoleInfoMap.UpsertRole(evtRoleCreated.OrgId, evtRoleCreated.RoleId, evtRoleCreated.IsVoter, evtRoleCreated.IsAdmin, types.AccessType(int(evtRoleCreated.BaseAccess.Uint64())), true)

			case evtRoleRevoked := <-chRoleRevoked:
				if r, _ := types.RoleInfoMap.GetRole(evtRoleRevoked.OrgId, evtRoleRevoked.RoleId); r != nil {
					types.RoleInfoMap.UpsertRole(evtRoleRevoked.OrgId, evtRoleRevoked.RoleId, r.IsVoter, r.IsAdmin, r.Access, false)
				} else {
					log.Error("Revoke role - cache is missing role", "org", evtRoleRevoked.OrgId, "role", evtRoleRevoked.RoleId)
				}
			case <-stopChan:
				log.Info("quit role contract watch")
				return
			}
		}
	}()
	return nil
}

// getter to get an account record from the contract
func (p *PermissionCtrl) PopulateAccountToCache(acctId common.Address) (*types.AccountInfo, error) {
	account, orgId, roleId, status, isAdmin, err := p.contract.GetAccountDetails(acctId)
	if err != nil {
		return nil, err
	}

	if status.Int64() == 0 {
		return nil, types.ErrAccountNotThere
	}
	return &types.AccountInfo{AcctId: account, OrgId: orgId, RoleId: roleId, Status: types.AcctStatus(status.Int64()), IsOrgAdmin: isAdmin}, nil
}

// getter to get a org record from the contract
func (p *PermissionCtrl) PopulateOrgToCache(orgId string) (*types.OrgInfo, error) {
	org, parentOrgId, ultimateParentId, orgLevel, orgStatus, err := p.contract.GetOrgDetails(orgId)
	if err != nil {
		return nil, err
	}
	if orgStatus.Int64() == 0 {
		return nil, types.ErrOrgDoesNotExists
	}
	orgInfo := types.OrgInfo{OrgId: org, ParentOrgId: parentOrgId, UltimateParent: ultimateParentId, Status: types.OrgStatus(orgStatus.Int64()), Level: orgLevel}
	// now need to build the list of sub orgs for this org
	subOrgIndexes, err := p.contract.GetSubOrgIndexes(orgId)
	if err != nil {
		return nil, err
	}

	if len(subOrgIndexes) == 0 {
		return &orgInfo, nil
	}

	// range through the sub org indexes and get the org ids to populate the suborg list
	for _, s := range subOrgIndexes {
		subOrgId, _, _, _, _, err := p.contract.GetOrgInfo(s)

		if err != nil {
			return nil, err
		}
		orgInfo.SubOrgList = append(orgInfo.SubOrgList, orgId+"."+subOrgId)

	}
	return &orgInfo, nil
}

// getter to get a role record from the contract
func (p *PermissionCtrl) PopulateRoleToCache(roleKey *types.RoleKey) (*types.RoleInfo, error) {
	roleDetails, err := p.contract.GetRoleDetails(roleKey.RoleId, roleKey.OrgId)

	if err != nil {
		return nil, err
	}

	if roleDetails.OrgId == "" {
		return nil, types.ErrInvalidRole
	}
	return &types.RoleInfo{OrgId: roleDetails.OrgId, RoleId: roleDetails.RoleId, IsVoter: roleDetails.Voter, IsAdmin: roleDetails.Admin, Access: types.AccessType(roleDetails.AccessType.Int64()), Active: roleDetails.Active}, nil
}

// getter to get a role record from the contract
func (p *PermissionCtrl) PopulateNodeCache(url string) (*types.NodeInfo, error) {
	orgId, url, status, err := p.contract.GetNodeDetails(url)
	if err != nil {
		return nil, err
	}

	if status.Int64() == 0 {
		return nil, types.ErrNodeDoesNotExists
	}
	// TODO Amal
	return &types.NodeInfo{OrgId: orgId, Url: url, Status: types.NodeStatus(status.Int64())}, nil
}

// getter to get a Node record from the contract
func (p *PermissionCtrl) PopulateNodeCacheAndValidate(hexNodeId, ultimateParentId string) bool {
	txnAllowed := false
	passedEnode, _ := enode.ParseV4(hexNodeId)
	if numberOfNodes, err := p.contract.GetNumberOfNodes(); err == nil {
		numNodes := numberOfNodes.Uint64()
		for k := uint64(0); k < numNodes; k++ {
			if orgId, url, status, err := p.contract.GetNodeDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				if orgRec, err := types.OrgInfoMap.GetOrg(orgId); err != nil {
					if orgRec.UltimateParent == ultimateParentId {
						recEnode, _ := enode.ParseV4(url)
						if recEnode.ID() == passedEnode.ID() {
							txnAllowed = true
							types.NodeInfoMap.UpsertNode(orgId, url, types.NodeStatus(int(status.Int64())))
						}
					}
				}
			}
		}
	}
	return txnAllowed
}

func BindContract(contractInstance interface{}, bindFunc func() (interface{}, error)) error {
	element := reflect.ValueOf(contractInstance).Elem()
	instance, err := bindFunc()
	if err != nil {
		return err
	}
	element.Set(reflect.ValueOf(instance))
	return nil
}

