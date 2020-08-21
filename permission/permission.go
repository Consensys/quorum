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

	"github.com/ethereum/go-ethereum/accounts"
	basicbind "github.com/ethereum/go-ethereum/permission/basic/bind"
	eeabind "github.com/ethereum/go-ethereum/permission/eea/bind"

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
	"github.com/ethereum/go-ethereum/permission/basic"
	"github.com/ethereum/go-ethereum/permission/eea"
	ptype "github.com/ethereum/go-ethereum/permission/types"
	"github.com/ethereum/go-ethereum/raft"
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
	stopFeed       event.Feed      // broadcasting stopEvent when service is being stopped
	errorChan      chan error      // channel to capture error when starting aysnc
	mux            sync.Mutex
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
		node:           stack,
		key:            stack.GetNodeKey(),
		dataDir:        stack.DataDir(),
		permConfig:     pconfig,
		startWaitGroup: wg,
		errorChan:      make(chan error),
		eeaFlag:        eeaFlag,
	}
	if eeaFlag {
		p.backend = &eea.Backend{SubscribeStopEvent: p.subscribeStopEvent, UpdatePermissionedNodes: p.UpdatePermissionedNodes, UpdateDisallowedNodes: p.UpdateDisallowedNodes}
	} else {
		p.backend = &basic.Backend{SubscribeStopEvent: p.subscribeStopEvent, UpdatePermissionedNodes: p.UpdatePermissionedNodes, UpdateDisallowedNodes: p.UpdateDisallowedNodes}
	}

	stopChan, stopSubscription := p.subscribeStopEvent()
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

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (p *PermissionCtrl) AfterStart() error {
	log.Debug("permission service: binding contracts")
	err := <-p.errorChan // capture any error happened during asyncStart. Also wait here if asyncStart is not yet finish
	if err != nil {
		return err
	}
	p.contract.AfterStart()

	// populate the initial list of permissioned nodes and account accesses
	if err := p.populateInitPermissions(params.DEFAULT_ORGCACHE_SIZE, params.DEFAULT_ROLECACHE_SIZE,
		params.DEFAULT_NODECACHE_SIZE, params.DEFAULT_ACCOUNTCACHE_SIZE); err != nil {
		return fmt.Errorf("populateInitPermissions failed: %v", err)
	}

	// set the default access to ReadOnly
	types.SetDefaults(p.permConfig.NwAdminRole, p.permConfig.OrgAdminRole)
	for _, f := range []func() error{
		p.monitorQIP714Block,               // monitor block number to activate new permissions controls
		p.backend.ManageOrgPermissions,     // monitor org management related events
		p.backend.ManageNodePermissions,    // monitor org  level Node management events
		p.backend.ManageRolePermissions,    // monitor org level role management events
		p.backend.ManageAccountPermissions, // monitor org level account management events
	} {
		if err := f(); err != nil {
			return err
		}
	}
	log.Info("permission service: is now ready")
	return nil
}

// start service asynchronously due to dependencies
func (p *PermissionCtrl) asyncStart() {
	var ethereum *eth.Ethereum
	// will be blocked here until Node is up
	if err := p.node.Service(&ethereum); err != nil {
		p.errorChan <- fmt.Errorf("dependent ethereum service not started")
		return
	}
	defer func() {
		p.errorChan <- nil
	}()
	// for cases where the Node is joining an existing network, permission service
	// can be brought up only after block syncing is complete. This function
	// waits for block syncing before the starting permissions
	p.startWaitGroup.Add(1)
	go func(_wg *sync.WaitGroup) {
		log.Debug("permission service: waiting for downloader")
		stopChan, stopSubscription := p.subscribeStopEvent()
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
	}(p.startWaitGroup) // wait for downloader to sync if any

	log.Debug("permission service: waiting for all dependencies to be ready")
	p.startWaitGroup.Wait()
	client, err := p.node.Attach()
	if err != nil {
		p.errorChan <- fmt.Errorf("unable to create rpc client: %v", err)
		return
	}
	p.ethClnt = ethclient.NewClient(client)
	p.eth = ethereum
	p.contract = NewPermissionContractService(p.ethClnt, p.eeaFlag, p.key, p.permConfig)
	if p.eeaFlag {
		b := p.backend.(*eea.Backend)
		b.Contr = p.contract.(*eea.Contract)
	} else {
		b := p.backend.(*basic.Backend)
		b.Contr = p.contract.(*basic.Contract)
	}
}

func (p *PermissionCtrl) Start(srvr *p2p.Server) error {
	log.Debug("permission service: starting")
	go func() {
		log.Debug("permission service: starting async")
		p.asyncStart()
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
	p.stopFeed.Send(ptype.StopEvent{})
	log.Info("permission service: stopped")
	return nil
}

// monitors QIP714Block and set default access
func (p *PermissionCtrl) monitorQIP714Block() error {
	// if QIP714block is not given, set the default access
	// to readonly
	if p.eth.BlockChain().Config().QIP714Block == nil {
		types.SetDefaultAccess()
		return nil
	}
	//QIP714block is given, monitor block count
	go func() {
		chainHeadCh := make(chan core.ChainHeadEvent, 1)
		headSub := p.eth.BlockChain().SubscribeChainHeadEvent(chainHeadCh)
		defer headSub.Unsubscribe()
		stopChan, stopSubscription := p.subscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case head := <-chainHeadCh:
				if p.eth.BlockChain().Config().IsQIP714(head.Block.Number()) {
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

func (p *PermissionCtrl) subscribeStopEvent() (chan ptype.StopEvent, event.Subscription) {
	c := make(chan ptype.StopEvent)
	s := p.stopFeed.Subscribe(c)
	return c, s
}

// adds or deletes and entry from a given file
func (p *PermissionCtrl) updateFile(fileName, enodeId string, operation ptype.NodeOperation, createFile bool) {
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
		if (operation == ptype.NodeAdd && recExists) || (operation == ptype.NodeDelete && !recExists) {
			return
		}
	}
	if operation == ptype.NodeAdd {
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
func (p *PermissionCtrl) UpdatePermissionedNodes(enodeId string, operation ptype.NodeOperation) {
	log.Debug("UpdatePermissionedNodes", "DataDir", p.dataDir, "file", params.PERMISSIONED_CONFIG)

	path := filepath.Join(p.dataDir, params.PERMISSIONED_CONFIG)
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for permissioned-nodes.json file. This is because 'permissioned' flag is specified but no permissioned-nodes.json file is present", "err", err)
		return
	}

	p.updateFile(path, enodeId, operation, false)
	if operation == ptype.NodeDelete {
		p.disconnectNode(enodeId)
	}
}

//this function populates the black listed Node information into the disallowed-nodes.json file
func (p *PermissionCtrl) UpdateDisallowedNodes(url string, operation ptype.NodeOperation) {
	log.Debug("UpdateDisallowedNodes", "DataDir", p.dataDir, "file", params.BLACKLIST_CONFIG)

	fileExists := true
	path := filepath.Join(p.dataDir, params.BLACKLIST_CONFIG)
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
		p.updateFile(path, url, operation, false)
	} else {
		p.updateFile(path, url, operation, true)
	}
}

// Disconnect the Node from the network
func (p *PermissionCtrl) disconnectNode(enodeId string) {
	if p.eth.BlockChain().Config().Istanbul == nil && p.eth.BlockChain().Config().Clique == nil {
		var raftService *raft.RaftService
		if err := p.node.Service(&raftService); err == nil {
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
		server := p.node.Server()
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

func (p *PermissionCtrl) instantiateCache(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int) {
	// instantiate the cache objects for permissions
	types.OrgInfoMap = types.NewOrgCache(orgCacheSize)
	types.OrgInfoMap.PopulateCacheFunc(p.populateOrgToCache)

	types.RoleInfoMap = types.NewRoleCache(roleCacheSize)
	types.RoleInfoMap.PopulateCacheFunc(p.populateRoleToCache)

	types.NodeInfoMap = types.NewNodeCache(nodeCacheSize)
	types.NodeInfoMap.PopulateCacheFunc(p.populateNodeCache)
	types.NodeInfoMap.PopulateValidateFunc(p.populateNodeCacheAndValidate)

	types.AcctInfoMap = types.NewAcctCache(accountCacheSize)
	types.AcctInfoMap.PopulateCacheFunc(p.populateAccountToCache)
}

// Thus function checks if the initial network boot up status and if no
// populates permissions model with details from permission-config.json
func (p *PermissionCtrl) populateInitPermissions(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int) error {
	p.instantiateCache(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize)
	networkInitialized, err := p.contract.GetNetworkBootStatus()
	if err != nil {
		// handle the scenario of no contract code.
		log.Warn("Failed to retrieve network boot status ", "err", err)
		return err
	}

	if !networkInitialized {
		if err := p.bootupNetwork(); err != nil {
			return err
		}
	} else {
		//populate orgs, nodes, roles and accounts from contract
		for _, f := range []func() error{
			p.populateOrgsFromContract,
			p.populateNodesFromContract,
			p.populateRolesFromContract,
			p.populateAccountsFromContract,
		} {
			if err := f(); err != nil {
				return err
			}
		}
	}
	return nil
}

// initialize the permissions model and populate initial values
func (p *PermissionCtrl) bootupNetwork() error {
	if _, err := p.contract.SetPolicy(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, p.permConfig.OrgAdminRole); err != nil {
		log.Error("bootupNetwork SetPolicy failed", "err", err)
		return err
	}
	if _, err := p.contract.Init(p.permConfig.SubOrgBreadth, p.permConfig.SubOrgDepth); err != nil {
		log.Error("bootupNetwork init failed", "err", err)
		return err
	}

	types.OrgInfoMap.UpsertOrg(p.permConfig.NwAdminOrg, "", p.permConfig.NwAdminOrg, big.NewInt(1), types.OrgApproved)
	types.RoleInfoMap.UpsertRole(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, true, true, types.FullAccess, true)
	// populate the initial Node list from static-nodes.json
	if err := p.populateStaticNodesToContract(); err != nil {
		return err
	}
	// populate initial account access to full access
	if err := p.populateInitAccountAccess(); err != nil {
		return err
	}

	// update network status to boot completed
	if err := p.updateNetworkStatus(); err != nil {
		log.Error("failed to updated network boot status", "error", err)
		return err
	}
	return nil
}

// populates the account access details from contract into cache
func (p *PermissionCtrl) populateAccountsFromContract() error {
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
func (p *PermissionCtrl) populateRolesFromContract() error {
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
func (p *PermissionCtrl) populateNodesFromContract() error {
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
func (p *PermissionCtrl) populateOrgsFromContract() error {

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
func (p *PermissionCtrl) populateStaticNodesToContract() error {
	nodes := p.node.Server().Config.StaticNodes
	for _, node := range nodes {
		enodeId, ip, port, raftPort := node.NodeDetails()
		_, err := p.contract.AddAdminNode(enodeId, ip, port, raftPort)
		if err != nil {
			log.Warn("Failed to propose Node", "err", err, "enode", node.EnodeID())
			return err
		}
		types.NodeInfoMap.UpsertNode(p.permConfig.NwAdminOrg, types.GetNodeUrl(enodeId, ip[:], port, raftPort), 2)
	}
	return nil
}

// Invokes the initAccounts function of smart contract to set the initial
// set of accounts access to full access
func (p *PermissionCtrl) populateInitAccountAccess() error {
	for _, a := range p.permConfig.Accounts {
		_, er := p.contract.AddAdminAccount(a)
		if er != nil {
			log.Warn("Error adding permission initial account list", "err", er, "account", a)
			return er
		}
		types.AcctInfoMap.UpsertAccount(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, a, true, 2)
	}
	return nil
}

// updates network boot status to true
func (p *PermissionCtrl) updateNetworkStatus() error {
	_, err := p.contract.UpdateNetworkBootStatus()
	if err != nil {
		log.Warn("Failed to udpate network boot status ", "err", err)
		return err
	}
	return nil
}

// getter to get an account record from the contract
func (p *PermissionCtrl) populateAccountToCache(acctId common.Address) (*types.AccountInfo, error) {
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
func (p *PermissionCtrl) populateOrgToCache(orgId string) (*types.OrgInfo, error) {
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
func (p *PermissionCtrl) populateRoleToCache(roleKey *types.RoleKey) (*types.RoleInfo, error) {
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
func (p *PermissionCtrl) populateNodeCache(url string) (*types.NodeInfo, error) {
	orgId, url, status, err := p.contract.GetNodeDetails(url)
	if err != nil {
		return nil, err
	}

	if status.Int64() == 0 {
		return nil, types.ErrNodeDoesNotExists
	}
	return &types.NodeInfo{OrgId: orgId, Url: url, Status: types.NodeStatus(status.Int64())}, nil
}

// getter to get a Node record from the contract
func (p *PermissionCtrl) populateNodeCacheAndValidate(hexNodeId, ultimateParentId string) bool {
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

func NewPermissionContractService(ethClnt bind.ContractBackend, eeaFlag bool, key *ecdsa.PrivateKey,
	permConfig *types.PermissionConfig) ptype.ContractService {
	if eeaFlag {
		return &eea.Contract{EthClnt: ethClnt, Key: key, PermConfig: permConfig}
	}
	return &basic.Contract{EthClnt: ethClnt, Key: key, PermConfig: permConfig}
}

func NewPermissionContractServiceForApi(p *PermissionCtrl, frmAcct accounts.Account, transactOpts *bind.TransactOpts, gasLimit uint64, gasPrice *big.Int) ptype.ContractService {
	if p.eeaFlag {
		pc := p.contract.(*eea.Contract)
		ps := &eeabind.EeaPermInterfaceSession{
			Contract: pc.PermInterf,
			CallOpts: bind.CallOpts{
				Pending: true,
			},
			TransactOpts: bind.TransactOpts{
				From:     frmAcct.Address,
				GasLimit: gasLimit,
				GasPrice: gasPrice,
				Signer:   transactOpts.Signer,
			},
		}
		return &eea.Contract{PermInterfSession: ps, PermConfig: p.permConfig}
	}
	pc := p.contract.(*basic.Contract)
	ps := &basicbind.PermInterfaceSession{
		Contract: pc.PermInterf,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     frmAcct.Address,
			GasLimit: gasLimit,
			GasPrice: gasPrice,
			Signer:   transactOpts.Signer,
		},
	}
	return &basic.Contract{PermInterfSession: ps, PermConfig: p.permConfig}

}
