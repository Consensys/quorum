package permission

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
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
	pbind "github.com/ethereum/go-ethereum/permission/bind"
	"github.com/ethereum/go-ethereum/raft"
	"github.com/ethereum/go-ethereum/rpc"
)


type BasicPermissionCtrl struct {
	Node       *node.Node
	EthClnt    bind.ContractBackend
	Eth        *eth.Ethereum
	Key        *ecdsa.PrivateKey
	DataDir    string
	PermUpgr   *pbind.PermUpgr
	PermInterf *pbind.PermInterface
	PermNode   *pbind.NodeManager
	PermAcct   *pbind.AcctManager
	PermRole   *pbind.RoleManager
	PermOrg    *pbind.OrgManager
	PermConfig *types.PermissionConfig

	StartWaitGroup *sync.WaitGroup // waitgroup to make sure all dependencies are ready before we start the service
	StopFeed       event.Feed      // broadcasting stopEvent when service is being stopped
	ErrorChan      chan error      // channel to capture error when starting aysnc

	mux sync.Mutex
}

/*func (p *BasicPermissionCtrl) MonitorQIP714Block() error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) ManageOrgPermissions() error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) SubscribeStopEvent() (chan types.StopEvent, event.Subscription) {
	panic("implement me")
}

func (p *BasicPermissionCtrl) ManageNodePermissions() error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) UpdateFile(fileName, enodeId string, operation types.types.NodeOperation, createFile bool) {
	panic("implement me")
}

func (p *BasicPermissionCtrl) UpdatePermissionedNodes(enodeId string, operation types.types.NodeOperation) {
	panic("implement me")
}

func (p *BasicPermissionCtrl) UpdateDisallowedNodes(url string, operation types.types.NodeOperation) {
	panic("implement me")
}

func (p *BasicPermissionCtrl) ManageAccountPermissions() error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) DisconnectNode(enodeId string) {
	panic("implement me")
}

func (p *BasicPermissionCtrl) InstantiateCache(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int) {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateInitPermissions(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int) error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) BootupNetwork(permInterfSession *pbind.EeaPermInterfaceSession) error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateAccountsFromContract(auth *bind.TransactOpts) error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateRolesFromContract(auth *bind.TransactOpts) error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateNodesFromContract(auth *bind.TransactOpts) error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateOrgsFromContract(auth *bind.TransactOpts) error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateStaticNodesToContract(permissionsSession *pbind.EeaPermInterfaceSession) error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateInitAccountAccess(permissionsSession *pbind.EeaPermInterfaceSession) error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) UpdateNetworkStatus(permissionsSession *pbind.EeaPermInterfaceSession) error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) ManageRolePermissions() error {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateAccountToCache(acctId common.Address) (*types.AccountInfo, error) {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateOrgToCache(orgId string) (*types.OrgInfo, error) {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateRoleToCache(roleKey *types.RoleKey) (*types.RoleInfo, error) {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateNodeCache(url string) (*types.NodeInfo, error) {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PopulateNodeCacheAndValidate(hexNodeId, ultimateParentId string) bool {
	panic("implement me")
}

func (p *BasicPermissionCtrl) BlockChain() *core.BlockChain {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PermissionConfig() *types.PermissionConfig {
	panic("implement me")
}

func (p *BasicPermissionCtrl) Ethereum() *eth.Ethereum {
	panic("implement me")
}

func (p *BasicPermissionCtrl) PermissionInterface() interface{} {
	panic("implement me")
}
*/

func (p *BasicPermissionCtrl) PermissionInterface() interface{} {
	return p.PermInterf
}

func (p *BasicPermissionCtrl) BlockChain() *core.BlockChain {
	return p.Eth.BlockChain()
}

func (p *BasicPermissionCtrl) Ethereum() *eth.Ethereum {
	return p.Eth
}

func (p *BasicPermissionCtrl) PermissionConfig() *types.PermissionConfig {
	return p.PermConfig
}

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (p *BasicPermissionCtrl) AfterStart() error {
	log.Debug("permission service: binding contracts")
	err := <-p.ErrorChan // capture any error happened during AsyncStart. Also wait here if AsyncStart is not yet finish
	if err != nil {
		return err
	}
	if err := types.BindContract(&p.PermUpgr, func() (interface{}, error) { return pbind.NewPermUpgr(p.PermConfig.UpgrdAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermInterf, func() (interface{}, error) { return pbind.NewPermInterface(p.PermConfig.InterfAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermAcct, func() (interface{}, error) { return pbind.NewAcctManager(p.PermConfig.AccountAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermNode, func() (interface{}, error) { return pbind.NewNodeManager(p.PermConfig.NodeAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermRole, func() (interface{}, error) { return pbind.NewRoleManager(p.PermConfig.RoleAddress, p.EthClnt) }); err != nil {
		return err
	}
	if err := types.BindContract(&p.PermOrg, func() (interface{}, error) { return pbind.NewOrgManager(p.PermConfig.OrgAddress, p.EthClnt) }); err != nil {
		return err
	}

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
func (p *BasicPermissionCtrl) AsyncStart() {
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
}

func (p *BasicPermissionCtrl) Start(srvr *p2p.Server) error {
	log.Debug("permission service: starting")
	go func() {
		log.Debug("permission service: starting async")
		p.AsyncStart()
	}()
	return nil
}

func (p *BasicPermissionCtrl) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "quorumPermission",
			Version:   "1.0",
			Service:   NewQuorumControlsAPI(p),
			Public:    true,
		},
	}
}

func (p *BasicPermissionCtrl) Protocols() []p2p.Protocol {
	return []p2p.Protocol{}
}

func (p *BasicPermissionCtrl) Stop() error {
	log.Info("permission service: stopping")
	p.StopFeed.Send(types.StopEvent{})
	log.Info("permission service: stopped")
	return nil
}

// monitors QIP714Block and set default access
func (p *BasicPermissionCtrl) MonitorQIP714Block() error {
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
func (p *BasicPermissionCtrl) ManageOrgPermissions() error {
	chPendingApproval := make(chan *pbind.OrgManagerOrgPendingApproval, 1)
	chOrgApproved := make(chan *pbind.OrgManagerOrgApproved, 1)
	chOrgSuspended := make(chan *pbind.OrgManagerOrgSuspended, 1)
	chOrgReactivated := make(chan *pbind.OrgManagerOrgSuspensionRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.PermOrg.OrgManagerFilterer.WatchOrgPendingApproval(opts, chPendingApproval); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.PermOrg.OrgManagerFilterer.WatchOrgApproved(opts, chOrgApproved); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.PermOrg.OrgManagerFilterer.WatchOrgSuspended(opts, chOrgSuspended); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.PermOrg.OrgManagerFilterer.WatchOrgSuspensionRevoked(opts, chOrgReactivated); err != nil {
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

func (p *BasicPermissionCtrl) SubscribeStopEvent() (chan types.StopEvent, event.Subscription) {
	c := make(chan types.StopEvent)
	s := p.StopFeed.Subscribe(c)
	return c, s
}

// Monitors Node management events and updates cache accordingly
func (p *BasicPermissionCtrl) ManageNodePermissions() error {
	chNodeApproved := make(chan *pbind.NodeManagerNodeApproved, 1)
	chNodeProposed := make(chan *pbind.NodeManagerNodeProposed, 1)
	chNodeDeactivated := make(chan *pbind.NodeManagerNodeDeactivated, 1)
	chNodeActivated := make(chan *pbind.NodeManagerNodeActivated, 1)
	chNodeBlacklisted := make(chan *pbind.NodeManagerNodeBlacklisted)
	chNodeRecoveryInit := make(chan *pbind.NodeManagerNodeRecoveryInitiated, 1)
	chNodeRecoveryDone := make(chan *pbind.NodeManagerNodeRecoveryCompleted, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.PermNode.NodeManagerFilterer.WatchNodeApproved(opts, chNodeApproved); err != nil {
		return fmt.Errorf("failed WatchNodeApproved: %v", err)
	}

	if _, err := p.PermNode.NodeManagerFilterer.WatchNodeProposed(opts, chNodeProposed); err != nil {
		return fmt.Errorf("failed WatchNodeProposed: %v", err)
	}

	if _, err := p.PermNode.NodeManagerFilterer.WatchNodeDeactivated(opts, chNodeDeactivated); err != nil {
		return fmt.Errorf("failed NodeDeactivated: %v", err)
	}
	if _, err := p.PermNode.NodeManagerFilterer.WatchNodeActivated(opts, chNodeActivated); err != nil {
		return fmt.Errorf("failed WatchNodeActivated: %v", err)
	}

	if _, err := p.PermNode.NodeManagerFilterer.WatchNodeBlacklisted(opts, chNodeBlacklisted); err != nil {
		return fmt.Errorf("failed NodeBlacklisting: %v", err)
	}

	if _, err := p.PermNode.NodeManagerFilterer.WatchNodeRecoveryInitiated(opts, chNodeRecoveryInit); err != nil {
		return fmt.Errorf("failed NodeRecoveryInitiated: %v", err)
	}

	if _, err := p.PermNode.NodeManagerFilterer.WatchNodeRecoveryCompleted(opts, chNodeRecoveryDone); err != nil {
		return fmt.Errorf("failed NodeRecoveryCompleted: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtNodeApproved := <-chNodeApproved:
				p.UpdatePermissionedNodes(evtNodeApproved.EnodeId, types.NodeAdd)
				types.NodeInfoMap.UpsertNode(evtNodeApproved.OrgId, evtNodeApproved.EnodeId, types.NodeApproved)

			case evtNodeProposed := <-chNodeProposed:
				types.NodeInfoMap.UpsertNode(evtNodeProposed.OrgId, evtNodeProposed.EnodeId, types.NodePendingApproval)

			case evtNodeDeactivated := <-chNodeDeactivated:
				p.UpdatePermissionedNodes(evtNodeDeactivated.EnodeId, types.NodeDelete)
				types.NodeInfoMap.UpsertNode(evtNodeDeactivated.OrgId, evtNodeDeactivated.EnodeId, types.NodeDeactivated)

			case evtNodeActivated := <-chNodeActivated:
				p.UpdatePermissionedNodes(evtNodeActivated.EnodeId, types.NodeAdd)
				types.NodeInfoMap.UpsertNode(evtNodeActivated.OrgId, evtNodeActivated.EnodeId, types.NodeApproved)

			case evtNodeBlacklisted := <-chNodeBlacklisted:
				types.NodeInfoMap.UpsertNode(evtNodeBlacklisted.OrgId, evtNodeBlacklisted.EnodeId, types.NodeBlackListed)
				p.UpdateDisallowedNodes(evtNodeBlacklisted.EnodeId, types.NodeAdd)
				p.UpdatePermissionedNodes(evtNodeBlacklisted.EnodeId, types.NodeDelete)

			case evtNodeRecoveryInit := <-chNodeRecoveryInit:
				types.NodeInfoMap.UpsertNode(evtNodeRecoveryInit.OrgId, evtNodeRecoveryInit.EnodeId, types.NodeRecoveryInitiated)

			case evtNodeRecoveryDone := <-chNodeRecoveryDone:
				types.NodeInfoMap.UpsertNode(evtNodeRecoveryDone.OrgId, evtNodeRecoveryDone.EnodeId, types.NodeApproved)
				p.UpdateDisallowedNodes(evtNodeRecoveryDone.EnodeId, types.NodeDelete)
				p.UpdatePermissionedNodes(evtNodeRecoveryDone.EnodeId, types.NodeAdd)

			case <-stopChan:
				log.Info("quit Node contract watch")
				return
			}
		}
	}()
	return nil
}

// adds or deletes and entry from a given file
func (p *BasicPermissionCtrl) UpdateFile(fileName, enodeId string, operation types.types.NodeOperation, createFile bool) {
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
		if (operation == types.NodeAdd && recExists) || (operation == types.NodeDelete && !recExists) {
			return
		}
	}
	if operation == types.NodeAdd {
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
func (p *BasicPermissionCtrl) UpdatePermissionedNodes(enodeId string, operation types.NodeOperation) {
	log.Debug("updatePermissionedNodes", "DataDir", p.DataDir, "file", params.PERMISSIONED_CONFIG)

	path := filepath.Join(p.DataDir, params.PERMISSIONED_CONFIG)
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for permissioned-nodes.json file. This is because 'permissioned' flag is specified but no permissioned-nodes.json file is present", "err", err)
		return
	}

	p.UpdateFile(path, enodeId, operation, false)
	if operation == types.NodeDelete {
		p.DisconnectNode(enodeId)
	}
}

//this function populates the black listed Node information into the disallowed-nodes.json file
func (p *BasicPermissionCtrl) UpdateDisallowedNodes(url string, operation types.NodeOperation) {
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
func (p *BasicPermissionCtrl) ManageAccountPermissions() error {
	chAccessModified := make(chan *pbind.AcctManagerAccountAccessModified)
	chAccessRevoked := make(chan *pbind.AcctManagerAccountAccessRevoked)
	chStatusChanged := make(chan *pbind.AcctManagerAccountStatusChanged)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.PermAcct.AcctManagerFilterer.WatchAccountAccessModified(opts, chAccessModified); err != nil {
		return fmt.Errorf("failed AccountAccessModified: %v", err)
	}

	if _, err := p.PermAcct.AcctManagerFilterer.WatchAccountAccessRevoked(opts, chAccessRevoked); err != nil {
		return fmt.Errorf("failed AccountAccessRevoked: %v", err)
	}

	if _, err := p.PermAcct.AcctManagerFilterer.WatchAccountStatusChanged(opts, chStatusChanged); err != nil {
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
func (p *BasicPermissionCtrl) DisconnectNode(enodeId string) {
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

func (p *BasicPermissionCtrl) InstantiateCache(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int) {
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
func (p *BasicPermissionCtrl) PopulateInitPermissions(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int) error {
	auth := bind.NewKeyedTransactor(p.Key)
	permInterfSession := &pbind.PermInterfaceSession{
		Contract: p.PermInterf,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 47000000,
			GasPrice: big.NewInt(0),
		},
	}

	p.InstantiateCache(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize)

	networkInitialized, err := permInterfSession.GetNetworkBootStatus()
	if err != nil {
		// handle the scenario of no contract code.
		log.Warn("Failed to retrieve network boot status ", "err", err)
		return err
	}

	if !networkInitialized {
		if err := p.BootupNetwork(permInterfSession); err != nil {
			return err
		}
	} else {
		//populate orgs, nodes, roles and accounts from contract
		for _, f := range []func(auth *bind.TransactOpts) error{
			p.PopulateOrgsFromContract,
			p.PopulateNodesFromContract,
			p.PopulateRolesFromContract,
			p.PopulateAccountsFromContract,
		} {
			if err := f(auth); err != nil {
				return err
			}
		}
	}
	return nil
}

// initialize the permissions model and populate initial values
func (p *BasicPermissionCtrl) BootupNetwork(permInterfSession *pbind.PermInterfaceSession) error {
	if _, err := permInterfSession.SetPolicy(p.PermConfig.NwAdminOrg, p.PermConfig.NwAdminRole, p.PermConfig.OrgAdminRole); err != nil {
		log.Error("bootupNetwork SetPolicy failed", "err", err)
		return err
	}
	if _, err := permInterfSession.Init(p.PermConfig.SubOrgBreadth, p.PermConfig.SubOrgDepth); err != nil {
		log.Error("bootupNetwork init failed", "err", err)
		return err
	}

	types.OrgInfoMap.UpsertOrg(p.PermConfig.NwAdminOrg, "", p.PermConfig.NwAdminOrg, big.NewInt(1), types.OrgApproved)
	types.RoleInfoMap.UpsertRole(p.PermConfig.NwAdminOrg, p.PermConfig.NwAdminRole, true, true, types.FullAccess, true)
	// populate the initial Node list from static-nodes.json
	if err := p.PopulateStaticNodesToContract(permInterfSession); err != nil {
		return err
	}
	// populate initial account access to full access
	if err := p.PopulateInitAccountAccess(permInterfSession); err != nil {
		return err
	}

	// update network status to boot completed
	if err := p.UpdateNetworkStatus(permInterfSession); err != nil {
		log.Error("failed to updated network boot status", "error", err)
		return err
	}
	return nil
}

// populates the account access details from contract into cache
func (p *BasicPermissionCtrl) PopulateAccountsFromContract(auth *bind.TransactOpts) error {
	//populate accounts
	permAcctSession := &pbind.AcctManagerSession{
		Contract: p.PermAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}

	if numberOfRoles, err := permAcctSession.GetNumberOfAccounts(); err == nil {
		iOrgNum := numberOfRoles.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if addr, org, role, status, orgAdmin, err := permAcctSession.GetAccountDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				types.AcctInfoMap.UpsertAccount(org, role, addr, orgAdmin, types.AcctStatus(int(status.Int64())))
			}
		}
	} else {
		return err
	}
	return nil
}

// populates the role details from contract into cache
func (p *BasicPermissionCtrl) PopulateRolesFromContract(auth *bind.TransactOpts) error {
	//populate roles
	permRoleSession := &pbind.RoleManagerSession{
		Contract: p.PermRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	if numberOfRoles, err := permRoleSession.GetNumberOfRoles(); err == nil {
		iOrgNum := numberOfRoles.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if roleStruct, err := permRoleSession.GetRoleDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				types.RoleInfoMap.UpsertRole(roleStruct.OrgId, roleStruct.RoleId, roleStruct.Voter, roleStruct.Admin, types.AccessType(int(roleStruct.AccessType.Int64())), roleStruct.Active)
			}
		}

	} else {
		return err
	}
	return nil
}

// populates the Node details from contract into cache
func (p *BasicPermissionCtrl) PopulateNodesFromContract(auth *bind.TransactOpts) error {
	//populate nodes
	permNodeSession := &pbind.NodeManagerSession{
		Contract: p.PermNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	if numberOfNodes, err := permNodeSession.GetNumberOfNodes(); err == nil {
		iOrgNum := numberOfNodes.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if nodeStruct, err := permNodeSession.GetNodeDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				types.NodeInfoMap.UpsertNode(nodeStruct.OrgId, nodeStruct.EnodeId, types.NodeStatus(int(nodeStruct.NodeStatus.Int64())))
			}
		}
	} else {
		return err
	}
	return nil
}

// populates the org details from contract into cache
func (p *BasicPermissionCtrl) PopulateOrgsFromContract(auth *bind.TransactOpts) error {
	//populate orgs
	permOrgSession := &pbind.OrgManagerSession{
		Contract: p.PermOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	if numberOfOrgs, err := permOrgSession.GetNumberOfOrgs(); err == nil {
		iOrgNum := numberOfOrgs.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if orgId, porgId, ultParent, level, status, err := permOrgSession.GetOrgInfo(big.NewInt(int64(k))); err == nil {
				types.OrgInfoMap.UpsertOrg(orgId, porgId, ultParent, level, types.OrgStatus(int(status.Int64())))
			}
		}
	} else {
		return err
	}
	return nil
}

// Reads the Node list from static-nodes.json and populates into the contract
func (p *BasicPermissionCtrl) PopulateStaticNodesToContract(permissionsSession *pbind.PermInterfaceSession) error {
	nodes := p.Node.Server().Config.StaticNodes
	for _, node := range nodes {
		_, err := permissionsSession.AddAdminNode(node.String())
		if err != nil {
			log.Warn("Failed to propose Node", "err", err, "enode", node.EnodeID())
			return err
		}
		types.NodeInfoMap.UpsertNode(p.PermConfig.NwAdminOrg, node.String(), 2)
	}
	return nil
}

// Invokes the initAccounts function of smart contract to set the initial
// set of accounts access to full access
func (p *BasicPermissionCtrl) PopulateInitAccountAccess(permissionsSession *pbind.PermInterfaceSession) error {
	for _, a := range p.PermConfig.Accounts {
		_, er := permissionsSession.AddAdminAccount(a)
		if er != nil {
			log.Warn("Error adding permission initial account list", "err", er, "account", a)
			return er
		}
		types.AcctInfoMap.UpsertAccount(p.PermConfig.NwAdminOrg, p.PermConfig.NwAdminRole, a, true, 2)
	}
	return nil
}

// updates network boot status to true
func (p *BasicPermissionCtrl) UpdateNetworkStatus(permissionsSession *pbind.PermInterfaceSession) error {
	_, err := permissionsSession.UpdateNetworkBootStatus()
	if err != nil {
		log.Warn("Failed to udpate network boot status ", "err", err)
		return err
	}
	return nil
}

// monitors role management related events and updated cache
func (p *BasicPermissionCtrl) ManageRolePermissions() error {
	chRoleCreated := make(chan *pbind.RoleManagerRoleCreated, 1)
	chRoleRevoked := make(chan *pbind.RoleManagerRoleRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.PermRole.RoleManagerFilterer.WatchRoleCreated(opts, chRoleCreated); err != nil {
		return fmt.Errorf("failed WatchRoleCreated: %v", err)
	}

	if _, err := p.PermRole.RoleManagerFilterer.WatchRoleRevoked(opts, chRoleRevoked); err != nil {
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
func (p *BasicPermissionCtrl) PopulateAccountToCache(acctId common.Address) (*types.AccountInfo, error) {
	permAcctInterface := &pbind.AcctManagerSession{
		Contract: p.PermAcct,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	account, orgId, roleId, status, isAdmin, err := permAcctInterface.GetAccountDetails(acctId)
	if err != nil {
		return nil, err
	}

	if status.Int64() == 0 {
		return nil, types.ErrAccountNotThere
	}
	return &types.AccountInfo{AcctId: account, OrgId: orgId, RoleId: roleId, Status: types.AcctStatus(status.Int64()), IsOrgAdmin: isAdmin}, nil
}

// getter to get a org record from the contract
func (p *BasicPermissionCtrl) PopulateOrgToCache(orgId string) (*types.OrgInfo, error) {
	permOrgInterface := &pbind.OrgManagerSession{
		Contract: p.PermOrg,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	org, parentOrgId, ultimateParentId, orgLevel, orgStatus, err := permOrgInterface.GetOrgDetails(orgId)
	if err != nil {
		return nil, err
	}
	if orgStatus.Int64() == 0 {
		return nil, types.ErrOrgDoesNotExists
	}
	orgInfo := types.OrgInfo{OrgId: org, ParentOrgId: parentOrgId, UltimateParent: ultimateParentId, Status: types.OrgStatus(orgStatus.Int64()), Level: orgLevel}
	// now need to build the list of sub orgs for this org
	subOrgIndexes, err := permOrgInterface.GetSubOrgIndexes(orgId)
	if err != nil {
		return nil, err
	}

	if len(subOrgIndexes) == 0 {
		return &orgInfo, nil
	}

	// range through the sub org indexes and get the org ids to populate the suborg list
	for _, s := range subOrgIndexes {
		subOrgId, _, _, _, _, err := permOrgInterface.GetOrgInfo(s)

		if err != nil {
			return nil, err
		}
		orgInfo.SubOrgList = append(orgInfo.SubOrgList, orgId+"."+subOrgId)

	}
	return &orgInfo, nil
}

// getter to get a role record from the contract
func (p *BasicPermissionCtrl) PopulateRoleToCache(roleKey *types.RoleKey) (*types.RoleInfo, error) {
	permRoleInterface := &pbind.RoleManagerSession{
		Contract: p.PermRole,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	roleDetails, err := permRoleInterface.GetRoleDetails(roleKey.RoleId, roleKey.OrgId)

	if err != nil {
		return nil, err
	}

	if roleDetails.OrgId == "" {
		return nil, types.ErrInvalidRole
	}
	return &types.RoleInfo{OrgId: roleDetails.OrgId, RoleId: roleDetails.RoleId, IsVoter: roleDetails.Voter, IsAdmin: roleDetails.Admin, Access: types.AccessType(roleDetails.AccessType.Int64()), Active: roleDetails.Active}, nil
}

// getter to get a role record from the contract
func (p *BasicPermissionCtrl) PopulateNodeCache(url string) (*types.NodeInfo, error) {
	permNodeInterface := &pbind.NodeManagerSession{
		Contract: p.PermNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	nodeDetails, err := permNodeInterface.GetNodeDetails(url)
	if err != nil {
		return nil, err
	}

	if nodeDetails.NodeStatus.Int64() == 0 {
		return nil, types.ErrNodeDoesNotExists
	}
	return &types.NodeInfo{OrgId: nodeDetails.OrgId, Url: nodeDetails.EnodeId, Status: types.NodeStatus(nodeDetails.NodeStatus.Int64())}, nil
}

// getter to get a Node record from the contract
func (p *BasicPermissionCtrl) PopulateNodeCacheAndValidate(hexNodeId, ultimateParentId string) bool {
	permNodeInterface := &pbind.NodeManagerSession{
		Contract: p.PermNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	txnAllowed := false
	passedEnode, _ := enode.ParseV4(hexNodeId)
	if numberOfNodes, err := permNodeInterface.GetNumberOfNodes(); err == nil {
		numNodes := numberOfNodes.Uint64()
		for k := uint64(0); k < numNodes; k++ {
			if nodeStruct, err := permNodeInterface.GetNodeDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				if orgRec, err := types.OrgInfoMap.GetOrg(nodeStruct.OrgId); err != nil {
					if orgRec.UltimateParent == ultimateParentId {
						recEnode, _ := enode.ParseV4(nodeStruct.EnodeId)
						if recEnode.ID() == passedEnode.ID() {
							txnAllowed = true
							types.NodeInfoMap.UpsertNode(nodeStruct.OrgId, nodeStruct.EnodeId, types.NodeStatus(int(nodeStruct.NodeStatus.Int64())))
						}
					}
				}
			}
		}
	}
	return txnAllowed
}
