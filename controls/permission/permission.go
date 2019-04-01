package permission

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/controls"
	pbind "github.com/ethereum/go-ethereum/controls/bind/permission"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/raft"
)

type NodeOperation uint8

const (
	NodeAdd NodeOperation = iota
	NodeDelete
)

// permission config for bootstrapping
type PermissionLocalConfig struct {
	UpgrdAddress   string
	InterfAddress  string
	ImplAddress    string
	NodeAddress    string
	AccountAddress string
	RoleAddress    string
	VoterAddress   string
	OrgAddress     string
	NwAdminOrg     string
	NwAdminRole    string
	OrgAdminRole   string

	Accounts []string //initial list of account that need full access
}

type PermissionCtrl struct {
	node             *node.Node
	ethClnt          *ethclient.Client
	eth              *eth.Ethereum
	isRaft           bool
	permissionedMode bool
	key              *ecdsa.PrivateKey
	dataDir          string
	permUpgr         *pbind.PermUpgr
	permInterf       *pbind.PermInterface
	permNode         *pbind.NodeManager
	permAcct         *pbind.AcctManager
	permRole         *pbind.RoleManager
	permOrg          *pbind.OrgManager
	permConfig       *types.PermissionConfig
}

// This function takes the local config data where all the information is in string
// converts that to address and populates the global permissions config
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
	permConfig.NwAdminRole = config.NwAdminOrg
	permConfig.OrgAdminRole = config.OrgAdminRole

	// populate the account list as passed in config
	for _, val := range config.Accounts {
		permConfig.Accounts = append(permConfig.Accounts, common.HexToAddress(val))
	}

	return permConfig
}

// this function reads the permissions config file passed and populates the
// config structure accrodingly
func ParsePermissionConifg(dir string) (types.PermissionConfig, error) {
	fileName := "permission-config.json"
	fullPath := filepath.Join(dir, fileName)
	if _, err := os.Stat(fullPath); err != nil {
		log.Warn("permission-config.json file is missing", err)
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

	return permConfig, nil
}

// Creates the controls structure for permissions
func NewQuorumPermissionCtrl(stack *node.Node, permissionedMode, isRaft bool, pconfig *types.PermissionConfig) (*PermissionCtrl, error) {
	// Create a new ethclient to for interfacing with the contract
	stateReader, e, err := controls.CreateEthClient(stack)
	if err != nil {
		log.Error("Unable to create ethereum client for permissions check", "err", err)
		return nil, err
	}

	if pconfig.IsEmpty() && permissionedMode {
		log.Error("permission-config.json is missing contract address")
		return nil, errors.New("permission-config.json is missing contract address")
	}
	pu, err := pbind.NewPermUpgr(pconfig.UpgrdAddress, stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}
	// check if permissioning contract is there at address. If not return from here
	pm, err := pbind.NewPermInterface(pconfig.InterfAddress, stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}

	pmAcct, err := pbind.NewAcctManager(pconfig.AccountAddress, stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}

	pmNode, err := pbind.NewNodeManager(pconfig.NodeAddress, stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}

	pmRole, err := pbind.NewRoleManager(pconfig.RoleAddress, stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}

	pmOrg, err := pbind.NewOrgManager(pconfig.OrgAddress, stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}
	return &PermissionCtrl{stack, stateReader, e, isRaft, permissionedMode, stack.GetNodeKey(), stack.DataDir(), pu, pm, pmNode, pmAcct, pmRole, pmOrg, pconfig}, nil
}

// Starts the node permissioning and event monitoring for permissions
// smart contracts
func (p *PermissionCtrl) Start() error {
	// Permissions initialization
	if err := p.init(); err != nil {
		log.Error("Permissions init failed", "err", err)
		return err
	}

	// monitor org management related events
	go p.manageOrgPermissions()

	// monitor org  level node management events
	p.manageNodePermissions()

	// monitor org level role management events
	p.manageRolePermissions()

	// monitor org level account management events
	p.manageAccountPermissions()

	return nil
}

// Sets the initial values for the network
func (p *PermissionCtrl) init() error {
	// populate the initial list of permissioned nodes and account accesses
	if err := p.populateInitPermissions(); err != nil {
		return err
	}

	// set the default access to ReadOnly
	types.SetDefaultAccess()

	return nil
}

// monitors org management related events happening via
// smart contracts
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
		log.Info("AJ-new org pending approval waiting for events...")
		select {
		case evtPendingApproval = <-chPendingApproval:
			log.Info("AJ-newOrgPendingApproval", "node", evtPendingApproval.OrgId)
			types.OrgInfoMap.UpsertOrg(evtPendingApproval.OrgId, types.OrgStatus(evtPendingApproval.Type.Uint64()))
			log.Info("AJ-newOrgPendingApproval cached updated for ", "orgid", evtPendingApproval.OrgId)

		case evtOrgApproved = <-chOrgApproved:
			log.Info("AJ-newOrgPendingApproval", "node", evtOrgApproved.OrgId)
			types.OrgInfoMap.UpsertOrg(evtOrgApproved.OrgId, types.OrgApproved)
			log.Info("AJ-newOrgPendingApproval cached updated for ", "orgid", evtOrgApproved.OrgId)

		case evtOrgSuspended = <-chOrgSuspended:
			log.Info("AJ-newOrgPendingApproval", "node", evtOrgSuspended.OrgId)
			types.OrgInfoMap.UpsertOrg(evtOrgSuspended.OrgId, types.OrgSuspended)
			log.Info("AJ-newOrgPendingApproval cached updated for ", "orgid", evtOrgSuspended.OrgId)

		case evtOrgReactivated = <-chOrgReactivated:
			log.Info("AJ-newOrgPendingApproval", "node", evtOrgReactivated.OrgId)
			types.OrgInfoMap.UpsertOrg(evtOrgReactivated.OrgId, types.OrgApproved)
			log.Info("AJ-newOrgPendingApproval cached updated for ", "orgid", evtOrgReactivated.OrgId)
		}
	}
}

// Manages node addition, decavtivation and activation from network
func (p *PermissionCtrl) manageNodePermissions() {

	if p.permissionedMode {
		log.Info("AJ-manage node start")
		//monitor for new nodes addition via smart contract
		go p.monitorNewNodeAdd()

		go p.monitorNewNodePendingApproval()

		//monitor for nodes deletion via smart contract
		go p.monitorNodeDeactivation()

		//monitor for nodes activation from deactivation status
		go p.monitorNodeActivation()

		//monitor for nodes blacklisting via smart contract
		go p.monitorNodeBlacklisting()
	}
}

// Listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func (p *PermissionCtrl) monitorNewNodeAdd() {
	log.Info("AJ-new node approved event monitor started...")
	ch := make(chan *pbind.NodeManagerNodeApproved, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var evt *pbind.NodeManagerNodeApproved

	_, err := p.permNode.NodeManagerFilterer.WatchNodeApproved(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeApproved: %v", err)
	}
	for {
		log.Info("AJ-new node approved waiting for events...")
		select {
		case evt = <-ch:
			log.Info("AJ-newNodeApproved", "node", evt.EnodeId)
			p.updatePermissionedNodes(evt.EnodeId, NodeAdd)
			types.NodeInfoMap.UpsertNode(evt.OrgId, evt.EnodeId, types.NodeApproved)
			log.Info("AJ-newNodeApproved cached updated for ", "enode", evt.EnodeId)
		}
	}
}

func (p *PermissionCtrl) monitorNewNodePendingApproval() {
	log.Info("AJ-new node proposed event monitor started...")
	ch := make(chan *pbind.NodeManagerNodeProposed, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var evt *pbind.NodeManagerNodeProposed

	_, err := p.permNode.NodeManagerFilterer.WatchNodeProposed(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeProposed: %v", err)
	}
	for {
		log.Info("AJ-new node proposed waiting for events...")
		select {
		case evt = <-ch:
			log.Info("AJ-newNodeProposed", "node", evt.EnodeId)
			p.updatePermissionedNodes(evt.EnodeId, NodeAdd)
			types.NodeInfoMap.UpsertNode(evt.OrgId, evt.EnodeId, types.NodePendingApproval)
			log.Info("AJ-newNodeProposed cached updated for ", "enode", evt.EnodeId)
		}
	}
}

// Listens on the channel for new node deactivation via smart contract
// and removes the same from permissioned-nodes.json
func (p *PermissionCtrl) monitorNodeDeactivation() {
	ch := make(chan *pbind.NodeManagerNodeDeactivated)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var evt *pbind.NodeManagerNodeDeactivated
	_, err := p.permNode.NodeManagerFilterer.WatchNodeDeactivated(opts, ch)
	if err != nil {
		log.Info("Failed NodeDeactivated: %v", err)
	}
	for {
		select {
		case evt = <-ch:
			p.updatePermissionedNodes(evt.EnodeId, NodeDelete)
			types.NodeInfoMap.UpsertNode(evt.OrgId, evt.EnodeId, types.NodeDeactivated)
			log.Info("AJ-NodeDeactivated cached updated for ", "enode", evt.EnodeId)
		}

	}
}

// Listnes on the channel for any node activation via smart contract
// and adds the same permissioned-nodes.json
func (p *PermissionCtrl) monitorNodeActivation() {
	ch := make(chan *pbind.NodeManagerNodeActivated, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var evt *pbind.NodeManagerNodeActivated

	_, err := p.permNode.NodeManagerFilterer.WatchNodeActivated(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeActivated: %v", err)
	}
	for {
		select {
		case evt = <-ch:
			p.updatePermissionedNodes(evt.EnodeId, NodeAdd)
			types.NodeInfoMap.UpsertNode(evt.OrgId, evt.EnodeId, types.NodeActivated)
			log.Info("AJ-newNodeActivated cached updated for ", "enode", evt.EnodeId)
		}
	}
}

// Listens on the channel for node blacklisting via smart contract and
// adds the same into disallowed-nodes.json
func (p *PermissionCtrl) monitorNodeBlacklisting() {
	ch := make(chan *pbind.NodeManagerNodeBlacklisted)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var evt *pbind.NodeManagerNodeBlacklisted

	_, err := p.permNode.NodeManagerFilterer.WatchNodeBlacklisted(opts, ch)
	if err != nil {
		log.Info("Failed NodeBlacklisting: %v", err)
	}
	for {
		select {
		case evt = <-ch:
			log.Info("AJ-nodeBlackListed", "event", evt)
			p.updatePermissionedNodes(evt.EnodeId, NodeDelete)
			p.updateDisallowedNodes(evt.EnodeId)
			types.NodeInfoMap.UpsertNode(evt.OrgId, evt.EnodeId, types.NodeBlackListed)
			log.Info("AJ-newNodeABlacklisted cached updated for ", "enode", evt.EnodeId)
		}

	}
}

// Populates the new node information into the permissioned-nodes.json file
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

	nodelist := []string{}
	if err := json.Unmarshal(blob, &nodelist); err != nil {
		log.Error("updatePermissionedNodes: Failed to load nodes list", "err", err)
		return
	}

	// logic to update the permissioned-nodes.json file based on action
	index := 0
	recExists := false
	for i, eid := range nodelist {
		if eid == enodeId {
			index = i
			recExists = true
			break
		}
	}
	if operation == NodeAdd {
		if !recExists {
			nodelist = append(nodelist, enodeId)
		}
	} else {
		if recExists {
			nodelist = append(nodelist[:index], nodelist[index+1:]...)
		}
		p.disconnectNode(enodeId)
	}
	mu := sync.RWMutex{}
	blob, _ = json.Marshal(nodelist)

	mu.Lock()
	if err := ioutil.WriteFile(path, blob, 0644); err != nil {
		log.Error("updatePermissionedNodes: Error writing new node info to file", "err", err)
	}
	mu.Unlock()
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
	mu := sync.RWMutex{}
	blob, _ := json.Marshal(nodelist)
	mu.Lock()
	if err := ioutil.WriteFile(path, blob, 0644); err != nil {
		log.Error("updateDisallowedNodes: Error writing new node info to file", "err", err)
	}
	mu.Unlock()

	// Disconnect the peer if it is already connected
	p.disconnectNode(url)
}

// Manages account level permissions update
func (p *PermissionCtrl) manageAccountPermissions() {
	if !p.permissionedMode {
		return
	}
	go p.monitorAccountPermissionsAccessModified()
	go p.monitorAccountPermissionsAccessRevoked()
	return
}

// populates the nodes list from permissioned-nodes.json into the permissions smart contract
func (p *PermissionCtrl) populatePermissionedNodes() error {
	opts := &bind.FilterOpts{}
	pastAddEvent, err := p.permNode.NodeManagerFilterer.FilterNodeApproved(opts)

	if err == nil {
		recExists := true
		for recExists {
			recExists = pastAddEvent.Next()
			if recExists {
				p.updatePermissionedNodes(pastAddEvent.Event.EnodeId, NodeAdd)
			}
		}
	}

	opts = &bind.FilterOpts{}
	pastDelEvent, err := p.permNode.NodeManagerFilterer.FilterNodeDeactivated(opts)
	if err == nil {
		recExists := true
		for recExists {
			recExists = pastDelEvent.Next()
			if recExists {
				p.updatePermissionedNodes(pastDelEvent.Event.EnodeId, NodeDelete)
			}
		}
	}
	return nil
}

// Monitors permissions changes at acount level and uodate the account permissions cache
func (p *PermissionCtrl) monitorAccountPermissionsAccessModified() {
	ch := make(chan *pbind.AcctManagerAccountAccessModified)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var evt *pbind.AcctManagerAccountAccessModified

	_, err := p.permAcct.AcctManagerFilterer.WatchAccountAccessModified(opts, ch)
	if err != nil {
		log.Info("AJ-Failed NewNodeProposed: %v", err)
	}

	for {
		select {
		case evt = <-ch:
			log.Info("AJ-AccountAccessModified", "address", evt.Address, "role", evt.RoleId)
			types.AcctInfoMap.UpsertAccount(evt.OrgId, evt.RoleId, evt.Address, evt.OrgAdmin, types.AcctStatus(int(evt.Status.Uint64())))
			log.Info("AJ-AccountAccessModified cached updated for ", "acct", evt.Address)
		}

	}
}

func (p *PermissionCtrl) monitorAccountPermissionsAccessRevoked() {
	ch := make(chan *pbind.AcctManagerAccountAccessRevoked)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var evt *pbind.AcctManagerAccountAccessRevoked

	_, err := p.permAcct.AcctManagerFilterer.WatchAccountAccessRevoked(opts, ch)
	if err != nil {
		log.Info("AJ-Failed NewNodeProposed: %v", err)
	}

	for {
		select {
		case evt = <-ch:
			log.Info("AJ-AccountAccessModified", "address", evt.Address, "role", evt.RoleId)
			types.AcctInfoMap.UpsertAccount(evt.OrgId, evt.RoleId, evt.Address, evt.OrgAdmin, types.AcctActive)
			log.Info("AJ-AccountAccessModified cached updated for ", "acct", evt.Address)
		}

	}
}

// Disconnect the node from the network
func (p *PermissionCtrl) disconnectNode(enodeId string) {
	if p.isRaft {
		var raftService *raft.RaftService
		if err := p.node.Service(&raftService); err == nil {
			raftApi := raft.NewPublicRaftAPI(raftService)

			//get the raftId for the given enodeId
			raftId, err := raftApi.GetRaftId(enodeId)
			if err == nil {
				raftApi.RemovePeer(raftId)
			}
		}
	} else {
		// Istanbul - disconnect the peer
		server := p.node.Server()
		if server != nil {
			node, err := enode.ParseV4(enodeId)
			if err == nil {
				server.RemovePeer(node)
			}
		}
	}
}

// helper function to format EnodeId
func (p *PermissionCtrl) formatEnodeId(enodeId, ipAddrPort, discPort, raftPort string) string {
	newEnodeId := "enode://" + enodeId + "@" + ipAddrPort + "?discport=" + discPort
	if p.isRaft {
		newEnodeId += "&raftport=" + raftPort
	}
	return newEnodeId
}

// Thus function checks if the its the initial network boot up status and if no
// populates permissioning model with details from permission-config.json
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
		p.populateOrgsFromContract(auth)

		p.populateNodesFromContract(auth)

		p.populateRolesFromContract(auth)

		p.populateAccountsFromContract(auth)
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
	if _, err := permInterfSession.Init(p.permConfig.OrgAddress, p.permConfig.RoleAddress, p.permConfig.AccountAddress, p.permConfig.VoterAddress, p.permConfig.NodeAddress); err != nil {
		log.Error("bootupNetwork init failed", "err", err)
		return err
	}

	types.OrgInfoMap.UpsertOrg(p.permConfig.NwAdminOrg, 2)
	types.RoleInfoMap.UpsertRole(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, true, types.FullAccess, true)
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

func (p *PermissionCtrl) populateAccountsFromContract(auth *bind.TransactOpts) {
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

	}
}

func (p *PermissionCtrl) populateRolesFromContract(auth *bind.TransactOpts) {
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
				types.RoleInfoMap.UpsertRole(roleStruct.OrgId, roleStruct.RoleId, roleStruct.Voter, types.AccessType(int(roleStruct.AccessType.Int64())), roleStruct.Active)
			}
		}

	}
}

func (p *PermissionCtrl) populateNodesFromContract(auth *bind.TransactOpts) {
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

	}
}

func (p *PermissionCtrl) populateOrgsFromContract(auth *bind.TransactOpts) {
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
			if o, s, err := permOrgSession.GetOrgInfo(big.NewInt(int64(k))); err == nil {
				types.OrgInfoMap.UpsertOrg(o, types.OrgStatus(int(s.Int64())))
			}
		}

	}
}

// Reads the node list from static-nodes.json and populates into the contract
func (p *PermissionCtrl) populateStaticNodesToContract(permissionsSession *pbind.PermInterfaceSession) error {
	nodes := p2p.ParsePermissionedNodes(p.dataDir)
	for _, node := range nodes {

		enodeID := node.EnodeID()
		nonce := p.eth.TxPool().Nonce(permissionsSession.TransactOpts.From)
		permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)

		_, err := permissionsSession.AddAdminNodes(node.String())
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
		_, er := permissionsSession.AddAdminAccounts(a)
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

func (p *PermissionCtrl) manageRolePermissions() {
	if p.permissionedMode {
		log.Info("AJ-manage role start")
		//monitor for new nodes addition via smart contract
		go p.monitorNewRoleAdd()
		go p.monitorNewRoleRemove()
	}
}

func (p *PermissionCtrl) monitorNewRoleAdd() {
	log.Info("AJ-new role added event monitor started...")
	ch := make(chan *pbind.RoleManagerRoleCreated, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var evt *pbind.RoleManagerRoleCreated

	_, err := p.permRole.RoleManagerFilterer.WatchRoleCreated(opts, ch)
	if err != nil {
		log.Info("Failed WatchRoleCreated: %v", err)
	}
	for {
		log.Info("AJ-new role created waiting for events...")
		select {
		case evt = <-ch:
			log.Info("AJ-newRoleCreated", "org", evt.OrgId, "role", evt.RoleId)
			types.RoleInfoMap.UpsertRole(evt.OrgId, evt.RoleId, evt.IsVoter, types.AccessType(int(evt.BaseAccess.Uint64())), true)
			log.Info("AJ-newRoleCreated cached updated for ", "orgid", evt.OrgId, "role", evt.RoleId)
		}
	}
}

func (p *PermissionCtrl) monitorNewRoleRemove() {
	log.Info("AJ-new role remove event monitor started...")
	ch := make(chan *pbind.RoleManagerRoleRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var evt *pbind.RoleManagerRoleRevoked

	_, err := p.permRole.RoleManagerFilterer.WatchRoleRevoked(opts, ch)
	if err != nil {
		log.Info("Failed WatchRoleRemoved: %v", err)
	}
	for {
		log.Info("AJ-new role removed waiting for events...")
		select {
		case evt = <-ch:
			log.Info("AJ-newRoleRemoved", "org", evt.OrgId, "role", evt.RoleId)
			if r := types.RoleInfoMap.GetRole(evt.OrgId, evt.RoleId); r != nil {
				types.RoleInfoMap.UpsertRole(evt.OrgId, evt.RoleId, r.IsVoter, r.Access, false)
				log.Info("AJ-newRoleRemoved cached updated for ", "orgid", evt.OrgId, "role", evt.RoleId)
			} else {
				log.Error("AJ-revoke role - cache is missing role", "org", evt.OrgId, "role", evt.RoleId)
			}

		}
	}
}
