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
	"github.com/ethereum/go-ethereum/cmd/utils"
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
	var permConfig types.PermissionConfig
	err = json.Unmarshal(blob, &permConfig)
	if err != nil {
		log.Error("error unmarshalling permission-config.json file", err)
		return types.PermissionConfig{}, err
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
		utils.Fatalf("permission-config.json is missing contract address")
	}

	if !permissionedMode {
		return &PermissionCtrl{stack, stateReader, e, isRaft, permissionedMode, stack.GetNodeKey(), stack.DataDir(), nil, nil, nil, nil, nil, nil, pconfig}, nil
	}
	pu, err := pbind.NewPermUpgr(common.HexToAddress(pconfig.UpgrdAddress), stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}
	// check if permissioning contract is there at address. If not return from here
	pm, err := pbind.NewPermInterface(common.HexToAddress(pconfig.InterfAddress), stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}

	pmAcct, err := pbind.NewAcctManager(common.HexToAddress(pconfig.AccountAddress), stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}

	pmNode, err := pbind.NewNodeManager(common.HexToAddress(pconfig.NodeAddress), stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}

	pmRole, err := pbind.NewRoleManager(common.HexToAddress(pconfig.RoleAddress), stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}

	pmOrg, err := pbind.NewOrgManager(common.HexToAddress(pconfig.OrgAddress), stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network", "err", err)
		return nil, err
	}

	log.Info("AJ-permission contracts initialized")
	return &PermissionCtrl{stack, stateReader, e, isRaft, permissionedMode, stack.GetNodeKey(), stack.DataDir(), pu, pm, pmNode, pmAcct, pmRole, pmOrg, pconfig}, nil
}

// Starts the node permissioning and account access control monitoring
func (p *PermissionCtrl) Start() error {
	// Permissions initialization
	if err := p.init(); err != nil {
		log.Error("Permissions init failed", "err", err)
		return err
	}
	// Monitors node addition and decativation from network
	p.manageNodePermissions()
	// Monitors account level persmissions  update from smart contarct
	p.manageAccountPermissions()

	return nil
}

// Sets the initial values for the network
func (p *PermissionCtrl) init() error {
	if !p.permissionedMode {
		return nil
	}
	// populate the initial list of permissioned nodes and account accesses
	if err := p.populateInitPermissions(); err != nil {
		return err
	}

	// set the default access to ReadOnly
	types.SetDefaultAccess()

	return nil
}

// Manages node addition, decavtivation and activation from network
func (p *PermissionCtrl) manageNodePermissions() {

	if p.permissionedMode {
		log.Info("AJ-permission start")
		//monitor for new nodes addition via smart contract
		go p.monitorNewNodeAdd()

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
	var nodeAddEvent *pbind.NodeManagerNodeApproved

	_, err := p.permNode.NodeManagerFilterer.WatchNodeApproved(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeApproved: %v", err)
	}
	for {
		log.Info("AJ-new node approved waiting for events...")
		select {
		case nodeAddEvent = <-ch:
			log.Info("AJ-newNodeApproved", "node", nodeAddEvent.EnodeId)
			p.updatePermissionedNodes(nodeAddEvent.EnodeId, NodeAdd)

			auth := bind.NewKeyedTransactor(p.key)
			permAcctSession := &pbind.NodeManagerSession{
				Contract: p.permNode,
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
			if rs, err := permAcctSession.GetNodeDetails(nodeAddEvent.EnodeId); err != nil {
				log.Error("AJ-failed to read node info ", "err", err)
			} else {
				types.NodeInfoMap.UpsertNode(rs.OrgId, rs.EnodeId, int(rs.NodeStatus.Uint64()))
				log.Info("AJ-newNodeApproved cached updated for ", "enode", rs.EnodeId)
			}
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
	var newNodeDeleteEvent *pbind.NodeManagerNodeDeactivated
	_, err := p.permNode.NodeManagerFilterer.WatchNodeDeactivated(opts, ch)
	if err != nil {
		log.Info("Failed NodeDeactivated: %v", err)
	}
	for {
		select {
		case newNodeDeleteEvent = <-ch:
			p.updatePermissionedNodes(newNodeDeleteEvent.EnodeId, NodeDelete)
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
	var nodeActivatedEvent *pbind.NodeManagerNodeActivated

	_, err := p.permNode.NodeManagerFilterer.WatchNodeActivated(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeActivated: %v", err)
	}
	for {
		select {
		case nodeActivatedEvent = <-ch:
			p.updatePermissionedNodes(nodeActivatedEvent.EnodeId, NodeAdd)
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
	var newNodeBlacklistEvent *pbind.NodeManagerNodeBlacklisted

	_, err := p.permNode.NodeManagerFilterer.WatchNodeBlacklisted(opts, ch)
	if err != nil {
		log.Info("Failed NodeBlacklisting: %v", err)
	}
	for {
		select {
		case newNodeBlacklistEvent = <-ch:
			log.Info("AJ-nodeBlackListed", "event", newNodeBlacklistEvent)
			//p.updatePermissionedNodes(newNodeBlacklistEvent., newNodeBlacklistEvent.IpAddrPort, newNodeBlacklistEvent.DiscPort, newNodeBlacklistEvent.RaftPort, NodeDelete)
			//p.updateDisallowedNodes(newNodeBlacklistEvent)
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
func (p *PermissionCtrl) updateDisallowedNodes(nodeBlacklistEvent *pbind.PermissionsNodeBlacklisted) {
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

	newEnodeId := p.formatEnodeId(nodeBlacklistEvent.EnodeId, nodeBlacklistEvent.IpAddrPort, nodeBlacklistEvent.DiscPort, nodeBlacklistEvent.RaftPort)
	nodelist = append(nodelist, newEnodeId)
	mu := sync.RWMutex{}
	blob, _ := json.Marshal(nodelist)
	mu.Lock()
	if err := ioutil.WriteFile(path, blob, 0644); err != nil {
		log.Error("updateDisallowedNodes: Error writing new node info to file", "err", err)
	}
	mu.Unlock()

	// Disconnect the peer if it is already connected
	p.disconnectNode(newEnodeId)
}

// Manages account level permissions update
func (p *PermissionCtrl) manageAccountPermissions() {
	if !p.permissionedMode {
		return
	}
	//monitor for nodes deletiin via smart contract
	go p.monitorAccountPermissions()

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
func (p *PermissionCtrl) monitorAccountPermissions() {
	ch := make(chan *pbind.AcctManagerAccountAccessModified)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newEvent *pbind.AcctManagerAccountAccessModified

	_, err := p.permAcct.AcctManagerFilterer.WatchAccountAccessModified(opts, ch)
	if err != nil {
		log.Info("AJ-Failed NewNodeProposed: %v", err)
	}

	for {
		select {
		case newEvent = <-ch:
			log.Info("AJ-AccountAccessModified", "address", newEvent.Address, "role", newEvent.RoleId)
			types.AddAccountAccess(newEvent.Address, newEvent.RoleId)
			auth := bind.NewKeyedTransactor(p.key)
			permAcctSession := &pbind.AcctManagerSession{
				Contract: p.permAcct,
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
			if addr, org, role, status, orgAdmin, err := permAcctSession.GetAccountDetails(newEvent.Address); err != nil {
				log.Error("AJ-failed to read account info ", "err", err)
			} else {
				types.AcctInfoMap.UpsertAccount(org, role, addr, orgAdmin, int(status.Uint64()))
				log.Info("AJ-AccountAccessModified cached updated for ", "acct", addr)
			}

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

// Thus function checks if the its the initial network boot up and if yes
// populates the initial network enode details from static-nodes.json into
// smart contracts. Sets the accounts access to full access for the initial
// initial list of accounts as given in genesis.json file
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
		if err.Error() == "no contract code at given address" {
			return err
		}
		log.Warn("Failed to retrieve network boot status ", "err", err)
		return err
	}

	if networkInitialized && !p.permissionedMode {
		// Network is initialized with permissions and node is joining in a non-permissioned
		// option. stop the node from coming up
		utils.Fatalf("Joining a permissioned network in non-permissioned mode is not permitted. Bring up geth with --permissioned.")
	}

	if !p.permissionedMode {
		log.Info("Node started in non-permissioned mode")
		return errors.New("Node started in non-permissioned mode")
	}

	if !networkInitialized {
		if err := p.bootupNetwork(permInterfSession); err != nil {
			return err
		}
		log.Info("AJ-network boot completed")
	} else {
		log.Info("AJ-network already booted")
		//populate orgs, nodes, roles and accounts from contract

		p.populateOrgsFromContract(auth)

		p.populateNodesFromContract(auth)

		p.populateRolesFromContract(auth)

		p.populateAccountsFromContract(auth)
		log.Info("AJ-all data loaded from contract")
	}

	ShowCacheData()

	return nil
}

func (p *PermissionCtrl) bootupNetwork(permInterfSession *pbind.PermInterfaceSession) error {
	// Ensure that there is at least one account given as a part of genesis.json
	// which will have full access. If not throw a fatal error
	// Do not want a network with no access
	log.Info("AJ-network not initialized")

	permInterfSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permInterfSession.TransactOpts.From))
	if _, err := permInterfSession.SetPolicy(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, p.permConfig.OrgAdminRole); err != nil {
		log.Error("AJ-permIntr.setPolicy failed", "err", err)
		return err
	}
	log.Info("AJ-permInter setPolicy done")
	permInterfSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permInterfSession.TransactOpts.From))
	if _, err := permInterfSession.Init(common.HexToAddress(p.permConfig.OrgAddress), common.HexToAddress(p.permConfig.RoleAddress), common.HexToAddress(p.permConfig.AccountAddress), common.HexToAddress(p.permConfig.VoterAddress), common.HexToAddress(p.permConfig.NodeAddress)); err != nil {
		log.Error("AJ-permIntr.init failed", "err", err)
		return err
	}
	log.Info("AJ-permInter init done")

	types.OrgInfoMap.UpsertOrg(p.permConfig.NwAdminOrg, 2)
	types.RoleInfoMap.UpsertRole(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, true, 3, true)
	// populate the initial node list from static-nodes.json
	err := p.populateStaticNodesToContract(permInterfSession)
	if err != nil {
		return err
	}
	log.Info("AJ-permInter init node population done")
	// populate initial account access to full access
	err = p.populateInitAccountAccess(permInterfSession)
	if err != nil {
		return err
	}
	log.Info("AJ-permInter init account population done")
	if err == nil && len(p.permConfig.Accounts) == 0 {

		//utils.Fatalf("Permissioned network being brought up with zero accounts having full access. Add permissioned full access accounts in genesis.json and bring up the network")
	}

	// update network status to boot completed
	err = p.updateNetworkStatus(permInterfSession)
	if err != nil {
		log.Info("AJ-failed to updated network boot status")
		return err
	}
	return nil
}

func ShowCacheData() {
	log.Info("AJ-======org info ======")
	types.OrgInfoMap.Show()
	log.Info("AJ-======Node info======")
	types.NodeInfoMap.Show()
	log.Info("AJ-======role info======")
	types.RoleInfoMap.Show()
	log.Info("AJ-======act info======")
	types.AcctInfoMap.Show()
	log.Info("AJ-============")
}

func (p *PermissionCtrl) populateAccountsFromContract(auth *bind.TransactOpts) {
	//populate accounts
	permAcctSession := &pbind.AcctManagerSession{
		Contract: p.permAcct,
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
	permAcctSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permAcctSession.TransactOpts.From))
	if numberOfRoles, err := permAcctSession.GetNumberOfAccounts(); err != nil {
		log.Error("AJ-reading acct num failed")
	} else {
		log.Info("AJ-acct num ", "num", numberOfRoles.Int64())
		iOrgNum := numberOfRoles.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			permAcctSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permAcctSession.TransactOpts.From))
			if addr, org, role, status, orgAdmin, err := permAcctSession.GetAccountDetailsFromIndex(big.NewInt(int64(k))); err != nil {
				log.Error("AJ-Org reading org info failed")
			} else {
				types.AcctInfoMap.UpsertAccount(org, role, addr, orgAdmin, int(status.Int64()))
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
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 47000000,
			GasPrice: big.NewInt(0),
		},
	}
	permRoleSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permRoleSession.TransactOpts.From))
	if numberOfRoles, err := permRoleSession.GetNumberOfRoles(); err != nil {
		log.Error("AJ-reading role num failed")
	} else {
		log.Info("AJ-role num ", "num", numberOfRoles.Int64())
		iOrgNum := numberOfRoles.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			permRoleSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permRoleSession.TransactOpts.From))
			if roleStruct, err := permRoleSession.GetRoleDetailsFromIndex(big.NewInt(int64(k))); err != nil {
				log.Error("AJ-role reading org info failed")
			} else {
				types.RoleInfoMap.UpsertRole(roleStruct.OrgId, roleStruct.RoleId, roleStruct.Voter, int(roleStruct.AccessType.Int64()), roleStruct.Active)
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
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 47000000,
			GasPrice: big.NewInt(0),
		},
	}
	permNodeSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permNodeSession.TransactOpts.From))
	if numberOfNodes, err := permNodeSession.GetNumberOfNodes(); err != nil {
		log.Error("AJ-reading node num failed")
	} else {
		log.Info("AJ-node num ", "num", numberOfNodes.Int64())
		iOrgNum := numberOfNodes.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			permNodeSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permNodeSession.TransactOpts.From))
			if nodeStruct, err := permNodeSession.GetNodeDetailsFromIndex(big.NewInt(int64(k))); err != nil {
				log.Error("AJ-node reading org info failed")
			} else {
				types.NodeInfoMap.UpsertNode(nodeStruct.OrgId, nodeStruct.EnodeId, int(nodeStruct.NodeStatus.Int64()))
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
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 47000000,
			GasPrice: big.NewInt(0),
		},
	}
	permOrgSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permOrgSession.TransactOpts.From))
	if numberOfOrgs, err := permOrgSession.GetNumberOfOrgs(); err != nil {
		log.Error("AJ-reading org num failed")
	} else {
		log.Info("AJ-org num ", "num", numberOfOrgs.Int64())
		iOrgNum := numberOfOrgs.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			permOrgSession.TransactOpts.Nonce = new(big.Int).SetUint64(p.eth.TxPool().Nonce(permOrgSession.TransactOpts.From))
			if o, s, err := permOrgSession.GetOrgInfo(big.NewInt(int64(k))); err != nil {
				log.Error("AJ-Org reading org info failed")
			} else {
				types.OrgInfoMap.UpsertOrg(o, int(s.Int64()))
			}
		}

	}
}

// Reads the node list from static-nodes.json and populates into the contract
func (p *PermissionCtrl) populateStaticNodesToContract(permissionsSession *pbind.PermInterfaceSession) error {
	nodes := p2p.ParsePermissionedNodes(p.dataDir)
	for _, node := range nodes {

		enodeID := node.EnodeID()

		log.Info("AJ-Adding node to permissions contract", "enodeID", enodeID)

		nonce := p.eth.TxPool().Nonce(permissionsSession.TransactOpts.From)
		permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)

		tx, err := permissionsSession.AddAdminNodes(node.String())
		if err != nil {
			log.Warn("Failed to propose node", "err", err)
			return err
		}
		types.NodeInfoMap.UpsertNode(p.permConfig.NwAdminOrg, node.String(), 2)
		log.Debug("Transaction pending", "tx hash", tx.Hash())
	}
	return nil
}

// Invokes the initAccounts function of smart contract to set the initial
// set of accounts access to full access
func (p *PermissionCtrl) populateInitAccountAccess(permissionsSession *pbind.PermInterfaceSession) error {

	if !p.permConfig.IsEmpty() {
		log.Info("AJ-add initial account list ...")
		for _, a := range p.permConfig.Accounts {
			log.Info("AJ-adding account ", "A", a)
			nonce := p.eth.TxPool().Nonce(permissionsSession.TransactOpts.From)
			permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)
			_, er := permissionsSession.AddAdminAccounts(common.HexToAddress(a))
			if er != nil {
				utils.Fatalf("error adding permission initial account list account: %s, error:%v", a, er)
			}
			types.AcctInfoMap.UpsertAccount(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, common.HexToAddress(a), true, 2)
		}
		log.Info("AJ-add initial account list ...done")
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
