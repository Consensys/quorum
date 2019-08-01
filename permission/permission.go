package permission

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/raft"
	"github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
	pbind "github.com/ethereum/go-ethereum/permission/bind"
)

type NodeOperation uint8

const (
	NodeAdd NodeOperation = iota
	NodeDelete
)

// permission config for bootstrapping
type PermissionLocalConfig struct {
	UpgrdAddress   string `json:"upgrdableAddress"`
	InterfAddress  string `json:"interfaceAddress"`
	ImplAddress    string `json:"implAddress"`
	NodeAddress    string `json:"nodeMgrAddress"`
	AccountAddress string `json:"accountMgrAddress"`
	RoleAddress    string `json:"roleMgrAddress"`
	VoterAddress   string `json:"voterMgrAddress"`
	OrgAddress     string `json:"orgMgrAddress"`
	NwAdminOrg     string `json:"nwAdminOrg"`
	NwAdminRole    string `json:"nwAdminRole"`
	OrgAdminRole   string `json:"orgAdminRole"`

	Accounts      []string `json:"accounts"` //initial list of account that need full access
	SubOrgBreadth string   `json:"subOrgBreadth"`
	SubOrgDepth   string   `json:"subOrgDepth"`
}

type PermissionCtrl struct {
	node       *node.Node
	ethClnt    *ethclient.Client
	eth        *eth.Ethereum
	key        *ecdsa.PrivateKey
	dataDir    string
	permUpgr   *pbind.PermUpgr
	permInterf *pbind.PermInterface
	permNode   *pbind.NodeManager
	permAcct   *pbind.AcctManager
	permRole   *pbind.RoleManager
	permOrg    *pbind.OrgManager
	permConfig *types.PermissionConfig
	orgChan    chan struct{}
	nodeChan   chan struct{}
	roleChan   chan struct{}
	acctChan   chan struct{}
	mux        sync.Mutex
}

// converts local permissions data to global permissions config
func populateConfig(config PermissionLocalConfig) types.PermissionConfig {
	var permConfig types.PermissionConfig
	permConfig.UpgrdAddress = common.HexToAddress(config.UpgrdAddress)
	permConfig.InterfAddress = common.HexToAddress(config.InterfAddress)
	permConfig.ImplAddress = common.HexToAddress(config.ImplAddress)
	permConfig.OrgAddress = common.HexToAddress(config.OrgAddress)
	permConfig.RoleAddress = common.HexToAddress(config.RoleAddress)
	permConfig.NodeAddress = common.HexToAddress(config.NodeAddress)
	permConfig.AccountAddress = common.HexToAddress(config.AccountAddress)
	permConfig.VoterAddress = common.HexToAddress(config.VoterAddress)

	permConfig.NwAdminOrg = config.NwAdminOrg
	permConfig.NwAdminRole = config.NwAdminRole
	permConfig.OrgAdminRole = config.OrgAdminRole

	// populate the account list as passed in config
	for _, val := range config.Accounts {
		permConfig.Accounts = append(permConfig.Accounts, common.HexToAddress(val))
	}
	permConfig.SubOrgBreadth.SetString(config.SubOrgBreadth, 10)
	permConfig.SubOrgDepth.SetString(config.SubOrgDepth, 10)

	return permConfig
}

// function reads the permissions config file passed and populates the
// config structure accordingly
func ParsePermissionConifg(dir string) (types.PermissionConfig, error) {
	fileName := "permission-config.json"
	fullPath := filepath.Join(dir, fileName)
	if _, err := os.Stat(fullPath); err != nil {
		log.Warn("permission-config.json file is missing", "err", err)
		return types.PermissionConfig{}, err
	}

	blob, err := ioutil.ReadFile(fullPath)

	if err != nil {
		log.Error("error reading permission-config.json file", err)
		return types.PermissionConfig{}, err
	}
	var permlocConfig PermissionLocalConfig
	err = json.Unmarshal(blob, &permlocConfig)
	if err != nil {
		log.Error("error unmarshalling permission-config.json file", err)
		return types.PermissionConfig{}, err
	}

	permConfig := populateConfig(permlocConfig)
	if len(permConfig.Accounts) == 0 {
		return types.PermissionConfig{}, errors.New("no accounts given in permission-config.json. Network cannot boot up")
	}

	if permConfig.SubOrgDepth.Cmp(big.NewInt(0)) == 0 || permConfig.SubOrgBreadth.Cmp(big.NewInt(0)) == 0 {
		return types.PermissionConfig{}, errors.New("sub org breadth depth not passed in permission-config.json. Network cannot boot up")
	}

	return permConfig, nil
}

// for cases where the node is joining an existing network, permission service
// can be brought up only after block syncing is complete. This function
// waits for block syncing before the starting permissions
func waitForSync(e *eth.Ethereum) {
	for !types.GetSyncStatus() {
		time.Sleep(10 * time.Millisecond)
	}
	for e.Downloader().Synchronising() {
		time.Sleep(10 * time.Millisecond)
	}
}

func StartPermissionService(stack *node.Node) {
	//initialize permission as we can create eth client only after the node and RPC are started
	var permissionService *PermissionCtrl
	if err := stack.Service(&permissionService); err != nil {
		utils.Fatalf("cannot access permissions service %v", err)
	}
	if permissionService == nil {
		utils.Fatalf("permission service unavailable")
	}
	//initialize the service to create eth client and get ethereum service
	if err := permissionService.InitializeService(); err != nil {
		utils.Fatalf("permissions service initialization failed %v", err)
	}
	if err := permissionService.Start(stack.Server()); err != nil {
		utils.Fatalf("permissions service start failed %v", err)
	}
}

// Creates the controls structure for permissions
func NewQuorumPermissionCtrl(stack *node.Node, pconfig *types.PermissionConfig) (*PermissionCtrl, error) {
	// Create a new ethclient to for interfacing with the contract
	return &PermissionCtrl{stack, nil, nil, stack.GetNodeKey(), stack.DataDir(), nil, nil, nil, nil, nil, nil, pconfig, make(chan struct{}), make(chan struct{}), make(chan struct{}), make(chan struct{}), sync.Mutex{}}, nil
}

func (p *PermissionCtrl) InitializeService() error {
	clnt, eth, err := CreateEthClient(p.node)
	if err != nil {
		log.Error("creating eth client failed")
		return err
	}
	waitForSync(eth)
	if err != nil {
		log.Error("Unable to create ethereum client for permissions check", "err", err)
		return err
	}

	if p.permConfig.IsEmpty() {
		log.Error("permission-config.json is missing contract address")
		return errors.New("permission-config.json is missing contract address")
	}
	pu, err := pbind.NewPermUpgr(p.permConfig.UpgrdAddress, clnt)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return err
	}
	// check if permissioning contract is there at address. If not return from here
	pm, err := pbind.NewPermInterface(p.permConfig.InterfAddress, clnt)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return err
	}

	pmAcct, err := pbind.NewAcctManager(p.permConfig.AccountAddress, clnt)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return err
	}

	pmNode, err := pbind.NewNodeManager(p.permConfig.NodeAddress, clnt)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return err
	}

	pmRole, err := pbind.NewRoleManager(p.permConfig.RoleAddress, clnt)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return err
	}

	pmOrg, err := pbind.NewOrgManager(p.permConfig.OrgAddress, clnt)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return err
	}
	p.permUpgr = pu
	p.permInterf = pm
	p.permAcct = pmAcct
	p.permNode = pmNode
	p.permRole = pmRole
	p.permOrg = pmOrg
	p.ethClnt = clnt
	p.eth = eth
	log.Info("permission service initalized")
	return nil
}

// Starts the node permissioning and event monitoring for permissions
// smart contracts
func (p *PermissionCtrl) Start(srvr *p2p.Server) error {

	if p.ethClnt == nil || p.eth == nil {
		log.Info("permission service not initialized")
		return nil
	}
	log.Info("permission service start...")
	// Permissions initialization
	if err := p.init(); err != nil {
		log.Error("Permissions init failed", "err", err)
		return err
	}

	// monitor org management related events
	go p.manageOrgPermissions()

	// monitor org  level node management events
	go p.manageNodePermissions()

	// monitor org level role management events
	go p.manageRolePermissions()

	// monitor org level account management events
	go p.manageAccountPermissions()
	log.Info("permission service started")
	return nil
}

func (s *PermissionCtrl) APIs() []rpc.API {
	log.Info("permission rpc API called")
	return []rpc.API{
		{
			Namespace: "quorumPermission",
			Version:   "1.0",
			Service:   NewQuorumControlsAPI(s),
			Public:    true,
		},
	}
}

func (s *PermissionCtrl) Protocols() []p2p.Protocol {
	return []p2p.Protocol{}
}

func (p *PermissionCtrl) Stop() error {
	if p.eth == nil || p.ethClnt == nil {
		return nil
	}
	log.Info("stopping permission service...")
	p.roleChan <- struct{}{}
	p.orgChan <- struct{}{}
	p.acctChan <- struct{}{}
	p.nodeChan <- struct{}{}
	log.Info("stopped permission service")
	return nil
}

// Sets the initial values for the network
func (p *PermissionCtrl) init() error {
	// populate the initial list of permissioned nodes and account accesses
	if err := p.populateInitPermissions(); err != nil {
		return err
	}

	// set the default access to ReadOnly
	types.SetDefaults(p.permConfig.NwAdminRole, p.permConfig.OrgAdminRole)

	return nil
}

// monitors org management related events happening via smart contracts
// and updates cache accordingly
func (p *PermissionCtrl) manageOrgPermissions() {

	chPendingApproval := make(chan *pbind.OrgManagerOrgPendingApproval, 1)
	chOrgApproved := make(chan *pbind.OrgManagerOrgApproved, 1)
	chOrgSuspended := make(chan *pbind.OrgManagerOrgSuspended, 1)
	chOrgReactivated := make(chan *pbind.OrgManagerOrgSuspensionRevoked, 1)

	var evtPendingApproval *pbind.OrgManagerOrgPendingApproval
	var evtOrgApproved *pbind.OrgManagerOrgApproved
	var evtOrgSuspended *pbind.OrgManagerOrgSuspended
	var evtOrgReactivated *pbind.OrgManagerOrgSuspensionRevoked

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.permOrg.OrgManagerFilterer.WatchOrgPendingApproval(opts, chPendingApproval); err != nil {
		log.Info("Failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.permOrg.OrgManagerFilterer.WatchOrgApproved(opts, chOrgApproved); err != nil {
		log.Info("Failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.permOrg.OrgManagerFilterer.WatchOrgSuspended(opts, chOrgSuspended); err != nil {
		log.Info("Failed WatchNodePendingApproval: %v", err)
	}

	if _, err := p.permOrg.OrgManagerFilterer.WatchOrgSuspensionRevoked(opts, chOrgReactivated); err != nil {
		log.Info("Failed WatchNodePendingApproval: %v", err)
	}

	for {
		select {
		case evtPendingApproval = <-chPendingApproval:
			types.OrgInfoMap.UpsertOrg(evtPendingApproval.OrgId, evtPendingApproval.PorgId, evtPendingApproval.UltParent, evtPendingApproval.Level, types.OrgStatus(evtPendingApproval.Status.Uint64()))

		case evtOrgApproved = <-chOrgApproved:
			types.OrgInfoMap.UpsertOrg(evtOrgApproved.OrgId, evtOrgApproved.PorgId, evtOrgApproved.UltParent, evtOrgApproved.Level, types.OrgApproved)

		case evtOrgSuspended = <-chOrgSuspended:
			types.OrgInfoMap.UpsertOrg(evtOrgSuspended.OrgId, evtOrgSuspended.PorgId, evtOrgSuspended.UltParent, evtOrgSuspended.Level, types.OrgSuspended)

		case evtOrgReactivated = <-chOrgReactivated:
			types.OrgInfoMap.UpsertOrg(evtOrgReactivated.OrgId, evtOrgReactivated.PorgId, evtOrgReactivated.UltParent, evtOrgReactivated.Level, types.OrgApproved)
		case <-p.orgChan:
			log.Info("quit org contract watch")
			return
		}
	}
}

// Monitors node management events and updates cache accordingly
func (p *PermissionCtrl) manageNodePermissions() {
	chNodeApproved := make(chan *pbind.NodeManagerNodeApproved, 1)
	chNodeProposed := make(chan *pbind.NodeManagerNodeProposed, 1)
	chNodeDeactivated := make(chan *pbind.NodeManagerNodeDeactivated, 1)
	chNodeActivated := make(chan *pbind.NodeManagerNodeActivated, 1)
	chNodeBlacklisted := make(chan *pbind.NodeManagerNodeBlacklisted)
	chNodeRecovery := make(chan *pbind.NodeManagerNodeRecoveryInitiated, 1)

	var evtNodeApproved *pbind.NodeManagerNodeApproved
	var evtNodeProposed *pbind.NodeManagerNodeProposed
	var evtNodeDeactivated *pbind.NodeManagerNodeDeactivated
	var evtNodeActivated *pbind.NodeManagerNodeActivated
	var evtNodeBlacklisted *pbind.NodeManagerNodeBlacklisted
	var evtNodeRecovery *pbind.NodeManagerNodeRecoveryInitiated

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.permNode.NodeManagerFilterer.WatchNodeApproved(opts, chNodeApproved); err != nil {
		log.Info("Failed WatchNodeApproved", "error", err)
	}

	if _, err := p.permNode.NodeManagerFilterer.WatchNodeProposed(opts, chNodeProposed); err != nil {
		log.Info("Failed WatchNodeProposed", "error", err)
	}

	if _, err := p.permNode.NodeManagerFilterer.WatchNodeDeactivated(opts, chNodeDeactivated); err != nil {
		log.Info("Failed NodeDeactivated", "error", err)
	}
	if _, err := p.permNode.NodeManagerFilterer.WatchNodeActivated(opts, chNodeActivated); err != nil {
		log.Info("Failed WatchNodeActivated", "error", err)
	}

	if _, err := p.permNode.NodeManagerFilterer.WatchNodeBlacklisted(opts, chNodeBlacklisted); err != nil {
		log.Info("Failed NodeBlacklisting", "error", err)
	}

	if _, err := p.permNode.NodeManagerFilterer.WatchNodeRecoveryInitiated(opts, chNodeRecovery); err != nil {
		log.Info("Failed NodeBlacklisting", "error", err)
	}

	for {
		select {
		case evtNodeApproved = <-chNodeApproved:
			p.updatePermissionedNodes(evtNodeApproved.EnodeId, NodeAdd)
			types.NodeInfoMap.UpsertNode(evtNodeApproved.OrgId, evtNodeApproved.EnodeId, types.NodeApproved)

		case evtNodeProposed = <-chNodeProposed:
			types.NodeInfoMap.UpsertNode(evtNodeProposed.OrgId, evtNodeProposed.EnodeId, types.NodePendingApproval)

		case evtNodeDeactivated = <-chNodeDeactivated:
			p.updatePermissionedNodes(evtNodeDeactivated.EnodeId, NodeDelete)
			types.NodeInfoMap.UpsertNode(evtNodeDeactivated.OrgId, evtNodeDeactivated.EnodeId, types.NodeDeactivated)

		case evtNodeActivated = <-chNodeActivated:
			p.updatePermissionedNodes(evtNodeActivated.EnodeId, NodeAdd)
			types.NodeInfoMap.UpsertNode(evtNodeActivated.OrgId, evtNodeActivated.EnodeId, types.NodeApproved)

		case evtNodeBlacklisted = <-chNodeBlacklisted:
			p.updatePermissionedNodes(evtNodeBlacklisted.EnodeId, NodeDelete)
			p.updateDisallowedNodes(evtNodeBlacklisted.EnodeId)
			types.NodeInfoMap.UpsertNode(evtNodeBlacklisted.OrgId, evtNodeBlacklisted.EnodeId, types.NodeBlackListed)

		case evtNodeRecovery = <-chNodeRecovery:
			types.NodeInfoMap.UpsertNode(evtNodeRecovery.OrgId, evtNodeRecovery.EnodeId, types.NodeRecoveryInitiated)

		case <-p.nodeChan:
			log.Info("quit node contract watch")
			return
		}
	}
}

// updates node information in the permissioned-nodes.json file based on node
// management activities in smart contract
func (p *PermissionCtrl) updatePermissionedNodes(enodeId string, operation NodeOperation) {
	log.Debug("updatePermissionedNodes", "DataDir", p.dataDir, "file", params.PERMISSIONED_CONFIG)

	path := filepath.Join(p.dataDir, params.PERMISSIONED_CONFIG)
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for permissioned-nodes.json file. This is because 'permissioned' flag is specified but no permissioned-nodes.json file is present", "err", err)
		return
	}
	// Load the nodes from the config file
	blob, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("updatePermissionedNodes: Failed to access permissioned-nodes.json", "err", err)
		return
	}

	var nodeList []string
	if err := json.Unmarshal(blob, &nodeList); err != nil {
		log.Error("updatePermissionedNodes: Failed to load nodes list", "err", err)
		return
	}

	// logic to update the permissioned-nodes.json file based on action
	index := 0
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
	if operation == NodeAdd {
		nodeList = append(nodeList, enodeId)
	} else {
		nodeList = append(nodeList[:index], nodeList[index+1:]...)
		p.disconnectNode(enodeId)
	}
	blob, _ = json.Marshal(nodeList)

	p.mux.Lock()
	defer p.mux.Unlock()

	if err := ioutil.WriteFile(path, blob, 0644); err != nil {
		log.Error("updatePermissionedNodes: Error writing new node info to file", "err", err)
	}
}

//this function populates the black listed node information into the disallowed-nodes.json file
func (p *PermissionCtrl) updateDisallowedNodes(url string) {
	log.Debug("updateDisallowedNodes", "DataDir", p.dataDir, "file", params.BLACKLIST_CONFIG)

	fileExisted := true
	path := filepath.Join(p.dataDir, params.BLACKLIST_CONFIG)
	// Check if the file is existing. If the file is not existing create the file
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for disallowed-nodes.json file", "err", err)
		if _, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644); err != nil {
			log.Error("Failed to create disallowed-nodes.json file", "err", err)
			return
		}
		fileExisted = false
	}

	nodelist := []string{}
	// Load the nodes from the config file
	if fileExisted == true {
		blob, err := ioutil.ReadFile(path)
		if err != nil {
			log.Error("updateDisallowedNodes Failed to access disallowed-nodes.json", "err", err)
			return
		}
		if blob != nil {
			if err := json.Unmarshal(blob, &nodelist); err != nil {
				log.Error("updateDisallowedNodes: Failed to load nodes list", "err", err)
				return
			}
		}
	}

	nodelist = append(nodelist, url)
	blob, _ := json.Marshal(nodelist)

	p.mux.Lock()
	defer p.mux.Unlock()

	if err := ioutil.WriteFile(path, blob, 0644); err != nil {
		log.Error("updateDisallowedNodes: Error writing new node info to file", "err", err)
	}
	// Disconnect the peer if it is already connected
	p.disconnectNode(url)
}

// Monitors account access related events and updates the cache accordingly
func (p *PermissionCtrl) manageAccountPermissions() {
	chAccessModified := make(chan *pbind.AcctManagerAccountAccessModified)
	chAccessRevoked := make(chan *pbind.AcctManagerAccountAccessRevoked)
	chStatusChanged := make(chan *pbind.AcctManagerAccountStatusChanged)

	var evtAccessModified *pbind.AcctManagerAccountAccessModified
	var evtAccessRevoked *pbind.AcctManagerAccountAccessRevoked
	var evtStatusChanged *pbind.AcctManagerAccountStatusChanged

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.permAcct.AcctManagerFilterer.WatchAccountAccessModified(opts, chAccessModified); err != nil {
		log.Info("Failed NewNodeProposed", "error", err)
	}

	if _, err := p.permAcct.AcctManagerFilterer.WatchAccountAccessRevoked(opts, chAccessRevoked); err != nil {
		log.Info("Failed NewNodeProposed", "error", err)
	}

	if _, err := p.permAcct.AcctManagerFilterer.WatchAccountStatusChanged(opts, chStatusChanged); err != nil {
		log.Info("Failed NewNodeProposed", "error", err)
	}

	for {
		select {
		case evtAccessModified = <-chAccessModified:
			types.AcctInfoMap.UpsertAccount(evtAccessModified.OrgId, evtAccessModified.RoleId, evtAccessModified.Account, evtAccessModified.OrgAdmin, types.AcctStatus(int(evtAccessModified.Status.Uint64())))

		case evtAccessRevoked = <-chAccessRevoked:
			types.AcctInfoMap.UpsertAccount(evtAccessRevoked.OrgId, evtAccessRevoked.RoleId, evtAccessRevoked.Account, evtAccessRevoked.OrgAdmin, types.AcctActive)

		case evtStatusChanged = <-chStatusChanged:
			ac := types.AcctInfoMap.GetAccount(evtStatusChanged.Account)
			types.AcctInfoMap.UpsertAccount(evtStatusChanged.OrgId, ac.RoleId, evtStatusChanged.Account, ac.IsOrgAdmin, types.AcctStatus(int(evtStatusChanged.Status.Uint64())))
		case <-p.acctChan:
			log.Info("quit account contract watch")
			return
		}
	}
}

// Disconnect the node from the network
func (p *PermissionCtrl) disconnectNode(enodeId string) {
	if p.eth.ChainConfig().Istanbul == nil && p.eth.ChainConfig().Clique == nil {
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
				log.Error("failed parse node id", "err", err, "enodeId", enodeId)
			}
		}
	}

}

// Thus function checks if the initial network boot up status and if no
// populates permissions model with details from permission-config.json
func (p *PermissionCtrl) populateInitPermissions() error {
	auth := bind.NewKeyedTransactor(p.key)
	permInterfSession := &pbind.PermInterfaceSession{
		Contract: p.permInterf,
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

	networkInitialized, err := permInterfSession.GetNetworkBootStatus()
	if err != nil {
		// handle the scenario of no contract code.
		log.Warn("Failed to retrieve network boot status ", "err", err)
		return err
	}

	if !networkInitialized {
		if err := p.bootupNetwork(permInterfSession); err != nil {
			return err
		}
	} else {
		//populate orgs, nodes, roles and accounts from contract
		if err := p.populateOrgsFromContract(auth); err != nil {
			return err
		}
		if err := p.populateNodesFromContract(auth); err != nil {
			return err
		}
		if err := p.populateRolesFromContract(auth); err != nil {
			return err
		}
		if err := p.populateAccountsFromContract(auth); err != nil {
			return err
		}
	}

	return nil
}

// initialize the permissions model and populate initial values
func (p *PermissionCtrl) bootupNetwork(permInterfSession *pbind.PermInterfaceSession) error {
	permInterfSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permInterfSession.TransactOpts.From))
	if _, err := permInterfSession.SetPolicy(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, p.permConfig.OrgAdminRole); err != nil {
		log.Error("bootupNetwork SetPolicy failed", "err", err)
		return err
	}
	permInterfSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permInterfSession.TransactOpts.From))
	if _, err := permInterfSession.Init(&p.permConfig.SubOrgBreadth, &p.permConfig.SubOrgDepth); err != nil {
		log.Error("bootupNetwork init failed", "err", err)
		return err
	}

	types.OrgInfoMap.UpsertOrg(p.permConfig.NwAdminOrg, "", p.permConfig.NwAdminOrg, big.NewInt(1), types.OrgApproved)
	types.RoleInfoMap.UpsertRole(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, true, true, types.FullAccess, true)
	// populate the initial node list from static-nodes.json
	if err := p.populateStaticNodesToContract(permInterfSession); err != nil {
		return err
	}
	// populate initial account access to full access
	if err := p.populateInitAccountAccess(permInterfSession); err != nil {
		return err
	}

	// update network status to boot completed
	if err := p.updateNetworkStatus(permInterfSession); err != nil {
		log.Error("failed to updated network boot status", "error", err)
		return err
	}
	return nil
}

// populates the account access details from contract into cache
func (p *PermissionCtrl) populateAccountsFromContract(auth *bind.TransactOpts) error {
	//populate accounts
	permAcctSession := &pbind.AcctManagerSession{
		Contract: p.permAcct,
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
func (p *PermissionCtrl) populateRolesFromContract(auth *bind.TransactOpts) error {
	//populate roles
	permRoleSession := &pbind.RoleManagerSession{
		Contract: p.permRole,
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

// populates the node details from contract into cache
func (p *PermissionCtrl) populateNodesFromContract(auth *bind.TransactOpts) error {
	//populate nodes
	permNodeSession := &pbind.NodeManagerSession{
		Contract: p.permNode,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	if numberOfNodes, err := permNodeSession.GetNumberOfNodes(); err == nil {
		iOrgNum := numberOfNodes.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			permNodeSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permNodeSession.TransactOpts.From))
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
func (p *PermissionCtrl) populateOrgsFromContract(auth *bind.TransactOpts) error {
	//populate orgs
	permOrgSession := &pbind.OrgManagerSession{
		Contract: p.permOrg,
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

// Reads the node list from static-nodes.json and populates into the contract
func (p *PermissionCtrl) populateStaticNodesToContract(permissionsSession *pbind.PermInterfaceSession) error {
	nodes := p2p.ParsePermissionedNodes(p.dataDir)
	for _, node := range nodes {

		enodeID := node.EnodeID()
		nonce := p.eth.TxPool().Nonce(permissionsSession.TransactOpts.From)
		permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)

		_, err := permissionsSession.AddAdminNode(node.String())
		if err != nil {
			log.Warn("Failed to propose node", "err", err, "enode", enodeID)
			return err
		}
		types.NodeInfoMap.UpsertNode(p.permConfig.NwAdminOrg, node.String(), 2)
	}
	return nil
}

// Invokes the initAccounts function of smart contract to set the initial
// set of accounts access to full access
func (p *PermissionCtrl) populateInitAccountAccess(permissionsSession *pbind.PermInterfaceSession) error {
	for _, a := range p.permConfig.Accounts {
		nonce := p.eth.TxPool().Nonce(permissionsSession.TransactOpts.From)
		permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)
		_, er := permissionsSession.AddAdminAccount(a)
		if er != nil {
			log.Warn("Error adding permission initial account list", "err", er, "account", a)
			return er
		}
		types.AcctInfoMap.UpsertAccount(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, a, true, 2)
	}
	return nil
}

// updates network boot status to true
func (p *PermissionCtrl) updateNetworkStatus(permissionsSession *pbind.PermInterfaceSession) error {
	nonce := p.eth.TxPool().Nonce(permissionsSession.TransactOpts.From)
	permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)
	_, err := permissionsSession.UpdateNetworkBootStatus()
	if err != nil {
		log.Warn("Failed to udpate network boot status ", "err", err)
		return err
	}
	return nil
}

// monitors role management related events and updated cache
func (p *PermissionCtrl) manageRolePermissions() {
	chRoleCreated := make(chan *pbind.RoleManagerRoleCreated, 1)
	chRoleRevoked := make(chan *pbind.RoleManagerRoleRevoked, 1)

	var evtRoleCreated *pbind.RoleManagerRoleCreated
	var evtRoleRevoked *pbind.RoleManagerRoleRevoked

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := p.permRole.RoleManagerFilterer.WatchRoleCreated(opts, chRoleCreated); err != nil {
		log.Info("Failed WatchRoleCreated: %v", err)
	}

	if _, err := p.permRole.RoleManagerFilterer.WatchRoleRevoked(opts, chRoleRevoked); err != nil {
		log.Info("Failed WatchRoleRemoved: %v", err)
	}

	for {
		select {
		case evtRoleCreated = <-chRoleCreated:
			types.RoleInfoMap.UpsertRole(evtRoleCreated.OrgId, evtRoleCreated.RoleId, evtRoleCreated.IsVoter, evtRoleCreated.IsAdmin, types.AccessType(int(evtRoleCreated.BaseAccess.Uint64())), true)

		case evtRoleRevoked = <-chRoleRevoked:
			if r := types.RoleInfoMap.GetRole(evtRoleRevoked.OrgId, evtRoleRevoked.RoleId); r != nil {
				types.RoleInfoMap.UpsertRole(evtRoleRevoked.OrgId, evtRoleRevoked.RoleId, r.IsVoter, r.IsAdmin, r.Access, false)
			} else {
				log.Error("Revoke role - cache is missing role", "org", evtRoleRevoked.OrgId, "role", evtRoleRevoked.RoleId)
			}
		case <-p.roleChan:
			log.Info("quit role contract watch")
			return
		}
	}
}
