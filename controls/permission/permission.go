package permission

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"strings"

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
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/raft"
)

const (
	PERMISSIONED_CONFIG = "permissioned-nodes.json"
	BLACKLIST_CONFIG    = "disallowed-nodes.json"
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
	pm               *pbind.Permissions
}

// Creates the controls structure for permissions
func NewQuorumPermissionCtrl(stack *node.Node, permissionedMode, isRaft bool) (*PermissionCtrl, error) {
	// Create a new ethclient to for interfacing with the contract
	stateReader, e, err := controls.CreateEthClient(stack)
	if err != nil {
		log.Error("Unable to create ethereum client for permissions check : ", "err", err)
		return nil, err
	}

	// check if permissioning contract is there at address. If not return from here
	pm, err := pbind.NewPermissions(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error("Permissions not enabled for the network : ", "err", err)
		return nil, err
	}

	return &PermissionCtrl{stack, stateReader, e, isRaft, permissionedMode, stack.GetNodeKey(), stack.DataDir(), pm}, nil
}

// Starts the node permissioning and account access control monitoring
func (p *PermissionCtrl) Start() error {
	// Permissions initialization
	if err := p.init(); err != nil {
		log.Error("Permissions init failed : ", "err", err)
		return err
	}

	// Monitors node addition and decativation from network
	p.manageNodePermissions()

	// Monitors account level persmissions  update from smart contarct
	p.manageAccountPermissions()

	return nil
}

// This functions updates the initial  values for the network
func (p *PermissionCtrl) init() error {
	// populate the initial list of permissioned nodes and account accesses
	if err := p.populateInitPermission(); err != nil {
		return err
	}

	// call populates the account permissions based on past history
	if err := p.populateAcctPermissions(); err != nil {
		return err
	}

	// call populates the node details from contract to KnownNodes
	// this is not required as the permissioned node info is persisted at
	// file level
	// if err := p.populatePermissionedNodes(); err != nil {
	// 	return err
	// }

	return nil
}

// Manages node addition, decavtivation and activation from network
func (p *PermissionCtrl) manageNodePermissions() {

	//monitor for new nodes addition via smart contract
	go p.monitorNewNodeAdd()

	//monitor for nodes deletiin via smart contract
	go p.monitorNodeDeactivation()

	//monitor for nodes activation from deactivation status
	go p.monitorNodeActivation()

	//monitor for nodes blacklisting via smart contract
	go p.monitorNodeBlacklisting()
}

// This functions listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func (p *PermissionCtrl) monitorNewNodeAdd() {
	ch := make(chan *pbind.PermissionsNodeApproved, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var nodeAddEvent *pbind.PermissionsNodeApproved

	_, err := p.pm.PermissionsFilterer.WatchNodeApproved(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeApproved: %v", err)
	}
	for {
		select {
		case nodeAddEvent = <-ch:
			p.updatePermissionedNodes(nodeAddEvent.EnodeId, nodeAddEvent.IpAddrPort, nodeAddEvent.DiscPort, nodeAddEvent.RaftPort, NodeAdd)
		}
	}
}

// This functions listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func (p *PermissionCtrl) monitorNodeDeactivation() {
	ch := make(chan *pbind.PermissionsNodeDeactivated)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newNodeDeleteEvent *pbind.PermissionsNodeDeactivated
	_, err := p.pm.PermissionsFilterer.WatchNodeDeactivated(opts, ch)
	if err != nil {
		log.Info("Failed NodeDeactivated: %v", err)
	}
	for {
		select {
		case newNodeDeleteEvent = <-ch:
			p.updatePermissionedNodes(newNodeDeleteEvent.EnodeId, newNodeDeleteEvent.IpAddrPort, newNodeDeleteEvent.DiscPort, newNodeDeleteEvent.RaftPort, NodeDelete)
		}

	}
}

// This function listnes on the channel for any node blacklisting event via smart contract
// and adds the same disallowed-nodes.json
func (p *PermissionCtrl) monitorNodeActivation() {
	ch := make(chan *pbind.PermissionsNodeActivated, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var nodeActivatedEvent *pbind.PermissionsNodeActivated

	_, err := p.pm.PermissionsFilterer.WatchNodeActivated(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeActivated: %v", err)
	}
	for {
		select {
		case nodeActivatedEvent = <-ch:
			p.updatePermissionedNodes(nodeActivatedEvent.EnodeId, nodeActivatedEvent.IpAddrPort, nodeActivatedEvent.DiscPort, nodeActivatedEvent.RaftPort, NodeAdd)
		}
	}
}

// This functions listens on the channel for node blacklisting via smart contract and
// adds the same into disallowed-nodes.json
func (p *PermissionCtrl) monitorNodeBlacklisting() {
	ch := make(chan *pbind.PermissionsNodeBlacklisted)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newNodeBlacklistEvent *pbind.PermissionsNodeBlacklisted

	_, err := p.pm.PermissionsFilterer.WatchNodeBlacklisted(opts, ch)
	if err != nil {
		log.Info("Failed NodeBlacklisting: %v", err)
	}
	for {
		select {
		case newNodeBlacklistEvent = <-ch:
			p.updatePermissionedNodes(newNodeBlacklistEvent.EnodeId, newNodeBlacklistEvent.IpAddrPort, newNodeBlacklistEvent.DiscPort, newNodeBlacklistEvent.RaftPort, NodeDelete)
			p.updateDisallowedNodes(newNodeBlacklistEvent)
		}

	}
}

//this function populates the new node information into the permissioned-nodes.json file
func (p *PermissionCtrl) updatePermissionedNodes(enodeId, ipAddrPort, discPort, raftPort string, operation NodeOperation) {
	log.Debug("updatePermissionedNodes", "DataDir", p.dataDir, "file", PERMISSIONED_CONFIG)

	path := filepath.Join(p.dataDir, PERMISSIONED_CONFIG)
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for permissioned-nodes.json file. This is because 'permissioned' flag is specified but no permissioned-nodes.json file is present.", "err", err)
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

	newEnodeId := p.formatEnodeId(enodeId, ipAddrPort, discPort, raftPort)

	// logic to update the permissioned-nodes.json file based on action
	index := 0
	recExists := false
	for i, enodeId := range nodelist {
		if strings.EqualFold(enodeId, newEnodeId){
			index = i
			recExists = true
			break
		}
	}
	if operation == NodeAdd {
		if !recExists {
			nodelist = append(nodelist, newEnodeId)
		}
	} else {
		if recExists {
			nodelist = append(nodelist[:index], nodelist[index+1:]...)
		}
		p.disconnectNode(newEnodeId)
	}
	mu := sync.RWMutex{}
	blob, _ = json.Marshal(nodelist)

	mu.Lock()
	if err:= ioutil.WriteFile(path, blob, 0644); err!= nil{
		log.Error("updatePermissionedNodes: Error writing new node info to file", "err", err)
	}
	mu.Unlock()
}

//this function populates the black listed node information into the permissioned-nodes.json file
func (p *PermissionCtrl) updateDisallowedNodes(nodeBlacklistEvent *pbind.PermissionsNodeBlacklisted) {
	log.Debug("updateDisallowedNodes", "DataDir", p.dataDir, "file", BLACKLIST_CONFIG)

	fileExisted := true
	path := filepath.Join(p.dataDir, BLACKLIST_CONFIG)
	// Check if the file is existing. If the file is not existing create the file
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for disallowed-nodes.json file.", "err", err)
		if _, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644); err != nil {
			log.Error("Failed to create disallowed-nodes.json file ", "err", err)
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
	//monitor for nodes deletiin via smart contract
	go p.monitorAccountPermissions()

	return
}

// populates the nodes list from permissioned-nodes.json into the permissions
// smart contract
func (p *PermissionCtrl) populatePermissionedNodes() error {
	opts := &bind.FilterOpts{}
	pastAddEvent, err := p.pm.PermissionsFilterer.FilterNodeApproved(opts)

	if err == nil {
		recExists := true
		for recExists {
			recExists = pastAddEvent.Next()
			if recExists {
				p.updatePermissionedNodes(pastAddEvent.Event.EnodeId, pastAddEvent.Event.IpAddrPort, pastAddEvent.Event.DiscPort, pastAddEvent.Event.RaftPort, NodeAdd)
			}
		}
	}

	opts = &bind.FilterOpts{}
	pastDelEvent, err := p.pm.PermissionsFilterer.FilterNodeDeactivated(opts)
	if err == nil {
		recExists := true
		for recExists {
			recExists = pastDelEvent.Next()
			if recExists {
				p.updatePermissionedNodes(pastDelEvent.Event.EnodeId, pastDelEvent.Event.IpAddrPort, pastDelEvent.Event.DiscPort, pastDelEvent.Event.RaftPort, NodeDelete)
			}
		}
	}
	return nil
}

// populates the nodes list from permissioned-nodes.json into the permissions
// smart contract
func (p *PermissionCtrl) populateAcctPermissions() error {
	opts := &bind.FilterOpts{}
	pastEvents, err := p.pm.PermissionsFilterer.FilterAccountAccessModified(opts)

	if err == nil {
		recExists := true
		for recExists {
			recExists = pastEvents.Next()
			if recExists {
				types.AddAccountAccess(pastEvents.Event.Address, pastEvents.Event.Access)
			}
		}
	}

	return nil
}

// Monitors permissions changes at acount level and uodate the global permissions
// map with the same
func (p *PermissionCtrl) monitorAccountPermissions() {
	ch := make(chan *pbind.PermissionsAccountAccessModified)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newEvent *pbind.PermissionsAccountAccessModified

	_, err := p.pm.PermissionsFilterer.WatchAccountAccessModified(opts, ch)
	if err != nil {
		log.Info("Failed NewNodeProposed: %v", err)
	}

	for {
		select {
		case newEvent = <-ch:
			types.AddAccountAccess(newEvent.Address, newEvent.Access)
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
			node, err := discover.ParseNode(enodeId)
			if err == nil {
				server.RemovePeer(node)
			}
		}
	}
}

// helper function to format EnodeId
func (p *PermissionCtrl) formatEnodeId(enodeId, ipAddrPort, discPort, raftPort string) string {
	newEnodeId := "enode://" + enodeId + "@" + ipAddrPort + "?discPort=" + discPort
	if p.isRaft {
		newEnodeId += "&raftport=" + raftPort
	}
	return newEnodeId
}

//populates the nodes list from permissioned-nodes.json into the permissions
//smart contract
func (p *PermissionCtrl) populateInitPermission() error {
	auth := bind.NewKeyedTransactor(p.key)
	permissionsSession := &pbind.PermissionsSession{
		Contract: p.pm,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 4700000,
			GasPrice: big.NewInt(0),
		},
	}

	tx, err := permissionsSession.GetNetworkBootStatus()
	if err != nil {
		// handle the scenario of no contract code.
		if err.Error() == "no contract code at given address"{
			return err
		}
		log.Warn("Failed to retrieve network boot status ", "err", err)
	}

	if tx && !p.permissionedMode {
		// Network is initialized with permissions and node is joining in a non-permissioned
		// option. stop the node from coming up
		utils.Fatalf("Joining a permissioned network in non-permissioned mode. Bring up geth with --permissioned.")
	}

	if !p.permissionedMode {
		return errors.New("Node started in non-permissioned mode")
	}
	if tx != true {
		// populate initial account access to full access
		err = p.populateInitAccountAccess(permissionsSession)
		if err != nil {
			return err
		}

		// populate the initial node list from static-nodes.json
		err := p.populateStaticNodesToContract(permissionsSession)
		if err != nil {
			return err
		}

		// update network status to boot completed
		err = p.updateNetworkStatus(permissionsSession)
		if err != nil {
			return err
		}
	}
	return nil
}

// Reads the node list from static-nodes.json and populates into the contract
func (p *PermissionCtrl) populateStaticNodesToContract(permissionsSession *pbind.PermissionsSession) error {
	nodes := p2p.ParsePermissionedNodes(p.dataDir)
	for _, node := range nodes {

		enodeID := node.ID.String()
		ipAddr := node.IP.String()
		port := fmt.Sprintf("%v", node.TCP)
		discPort := fmt.Sprintf("%v", node.UDP)
		raftPort := fmt.Sprintf("%v", node.RaftPort)

		ipAddrPort := ipAddr + ":" + port

		log.Trace("Adding node to permissions contract", "enodeID", enodeID)

		nonce := p.eth.TxPool().Nonce(permissionsSession.TransactOpts.From)
		permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)

		tx, err := permissionsSession.ProposeNode(enodeID, ipAddrPort, discPort, raftPort)
		if err != nil {
			log.Warn("Failed to propose node", "err", err)
			return err
		}
		log.Debug("Transaction pending", "tx hash", tx.Hash())
	}
	return nil
}

// Reads the acount from geth keystore and grants full access to these accounts
func (p *PermissionCtrl) populateInitAccountAccess(permissionsSession *pbind.PermissionsSession) error {
	_, err := permissionsSession.InitAccounts()
	if err != nil {
		log.Error("calling init accounts failed", "err", err)
		return err
	}
	return nil
}

// update network boot status to true
func (p *PermissionCtrl) updateNetworkStatus(permissionsSession *pbind.PermissionsSession) error {
	nonce := p.eth.TxPool().Nonce(permissionsSession.TransactOpts.From)
	permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)
	_, err := permissionsSession.UpdateNetworkBootStatus()
	if err != nil {
		log.Warn("Failed to udpate network boot status ", "err", err)
		return err
	}
	return nil
}
