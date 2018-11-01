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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/controls"
	pbind "github.com/ethereum/go-ethereum/controls/bind"
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
	node    *node.Node
	ethClnt *ethclient.Client
	eth     *eth.Ethereum
	isRaft  bool
	key     *ecdsa.PrivateKey
}

// Creates the controls structure for permissions
func NewQuorumPermissionCtrl(stack *node.Node, isRaft bool) (*PermissionCtrl, error) {
	// Create a new ethclient to for interfacing with the contract
	stateReader, e, err := controls.CreateEthClient(stack)
	if err != nil {
		log.Error("Unable to create ethereum client for permissions check : ", "err", err)
		return nil, err
	}
	prvKey := stack.GetNodeKey()
	log.Info("mykey value is : ", "prvKey", prvKey)

	return &PermissionCtrl{stack, stateReader, e, isRaft, prvKey}, nil
}

// This function first adds the node list from permissioned-nodes.json to
// the permissiones contract deployed as a precompile via genesis.json
func (p *PermissionCtrl) Start() error {

	// check if permissioning contract is there at address. If not return from here
	if _, err := pbind.NewPermissionsFilterer(params.QuorumPermissionsContract, p.ethClnt); err != nil {
		log.Error("Permissions not enabled for the network : ", "err", err)
		return nil
	}

	// Permissions initialization
	p.init()

	// Monitors node addition and decativation from network
	p.manageNodePermissions()

	// Monitors account level persmissions  update from smart contarct
	p.manageAccountPermissions()
	return nil
}

// This functions updates the initial  values for the network
func (p *PermissionCtrl) init() error {
	// populate the initial list of nodes into the smart contract
	// from permissioned-nodes.json
	p.populateStaticNodesToContract()

	// populate the account access for the genesis.json accounts. these
	// accounts will have full access
	// populateInitAccountAccess()

	// call populates the node details from contract to KnownNodes
	if err := p.populatePermissionedNodes(); err != nil {
		return err
	}
	// call populates the account permissions based on past history
	if err := p.populateAcctPermissions(); err != nil {
		return err
	}

	return nil
}

// Manages node addition and decavtivation from network
func (p *PermissionCtrl) manageNodePermissions() {

	//monitor for new nodes addition via smart contract
	go p.monitorNewNodeAdd()

	//monitor for nodes deletiin via smart contract
	go p.monitorNodeDeactivation()

	//monitor for nodes blacklisting via smart contract
	go p.monitorNodeBlacklisting()
}

// This functions listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func (p *PermissionCtrl) monitorNewNodeAdd() {
	permissions, err := pbind.NewPermissionsFilterer(params.QuorumPermissionsContract, p.ethClnt)
	if err != nil {
		log.Error("failed to monitor new node add : ", "err", err)
	}

	ch := make(chan *pbind.PermissionsNodeApproved, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var nodeAddEvent *pbind.PermissionsNodeApproved

	_, err = permissions.WatchNodeApproved(opts, ch)
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
	permissions, err := pbind.NewPermissionsFilterer(params.QuorumPermissionsContract, p.ethClnt)
	if err != nil {
		log.Error("Failed to monitor node delete: ", "err", err)
	}

	ch := make(chan *pbind.PermissionsNodeDeactivated)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newNodeDeleteEvent *pbind.PermissionsNodeDeactivated

	_, err = permissions.WatchNodeDeactivated(opts, ch)
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
func (p *PermissionCtrl) monitorNodeBlacklisting() {
	permissions, err := pbind.NewPermissionsFilterer(params.QuorumPermissionsContract, p.ethClnt)
	if err != nil {
		log.Error("failed to monitor new node add : ", "err", err)
	}
	ch := make(chan *pbind.PermissionsNodeBlacklisted, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var nodeBlacklistEvent *pbind.PermissionsNodeBlacklisted

	_, err = permissions.WatchNodeBlacklisted(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeBlacklisted: %v", err)
	}

	for {
		select {
		case nodeBlacklistEvent = <-ch:
			p.updateDisallowedNodes(nodeBlacklistEvent)
		}
	}
}

//this function populates the new node information into the permissioned-nodes.json file
func (p *PermissionCtrl) updatePermissionedNodes(enodeId, ipAddrPort, discPort, raftPort string, operation NodeOperation) {
	newEnodeId := formatEnodeId(enodeId, ipAddrPort, discPort, raftPort, p.isRaft)

	//new logic to update the server KnownNodes variable for permissioning
	server := p.node.Server()
	newNode, err := discover.ParseNode(newEnodeId)

	if err != nil {
		log.Error("updatePermissionedNodes: Node URL", "url", newEnodeId, "err", err)
	}

	if operation == NodeAdd {
		// Add the new enode id to server.KnownNodes
		server.KnownNodes = append(server.KnownNodes, newNode)
	} else {
		// delete the new enode id from server.KnownNodes
		index := 0
		for i, node := range server.KnownNodes {
			if node.ID == newNode.ID {
				index = i
			}
		}
		server.KnownNodes = append(server.KnownNodes[:index], server.KnownNodes[index+1:]...)
	}

}

//this function populates the new node information into the permissioned-nodes.json file
func (p *PermissionCtrl) updateDisallowedNodes(nodeBlacklistEvent *pbind.PermissionsNodeBlacklisted) {
	dataDir := p.node.InstanceDir()
	log.Debug("updateDisallowedNodes", "DataDir", dataDir, "file", BLACKLIST_CONFIG)

	fileExisted := true
	path := filepath.Join(dataDir, BLACKLIST_CONFIG)
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

	newEnodeId := formatEnodeId(nodeBlacklistEvent.EnodeId, nodeBlacklistEvent.IpAddrPort, nodeBlacklistEvent.DiscPort, nodeBlacklistEvent.RaftPort, p.isRaft)
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
func (p *PermissionCtrl) manageAccountPermissions() error {
	//monitor for nodes deletiin via smart contract
	go p.monitorAccountPermissions()
	return nil
}

// populates the nodes list from permissioned-nodes.json into the permissions
// smart contract
func (p *PermissionCtrl) populatePermissionedNodes() error {
	permissions, err := pbind.NewPermissionsFilterer(params.QuorumPermissionsContract, p.ethClnt)
	if err != nil {
		log.Error("Failed to monitor node delete: ", "err", err)
		return err
	}

	opts := &bind.FilterOpts{}
	pastAddEvent, err := permissions.FilterNodeApproved(opts)

	recExists := true
	for recExists {
		recExists = pastAddEvent.Next()
		if recExists {
			p.updatePermissionedNodes(pastAddEvent.Event.EnodeId, pastAddEvent.Event.IpAddrPort, pastAddEvent.Event.DiscPort, pastAddEvent.Event.RaftPort, NodeAdd)
		}
	}

	opts = &bind.FilterOpts{}
	pastDelEvent, err := permissions.FilterNodeDeactivated(opts)

	recExists = true
	for recExists {
		recExists = pastDelEvent.Next()
		if recExists {
			p.updatePermissionedNodes(pastDelEvent.Event.EnodeId, pastDelEvent.Event.IpAddrPort, pastDelEvent.Event.DiscPort, pastDelEvent.Event.RaftPort, NodeDelete)
		}
	}
	return nil
}

// populates the nodes list from permissioned-nodes.json into the permissions
// smart contract
func (p *PermissionCtrl) populateAcctPermissions() error {
	permissions, err := pbind.NewPermissionsFilterer(params.QuorumPermissionsContract, p.ethClnt)
	if err != nil {
		log.Error("Failed to monitor node delete: ", "err", err)
		return err
	}

	opts := &bind.FilterOpts{}
	pastEvents, err := permissions.FilterAccountAccessModified(opts)

	recExists := true
	for recExists {
		recExists = pastEvents.Next()
		if recExists {
			types.AddAccountAccess(pastEvents.Event.Address, pastEvents.Event.Access)
		}
	}
	return nil
}

// Monitors permissions changes at acount level and uodate the global permissions
// map with the same
func (p *PermissionCtrl) monitorAccountPermissions() {
	permissions, err := pbind.NewPermissionsFilterer(params.QuorumPermissionsContract, p.ethClnt)
	if err != nil {
		log.Error("Failed to monitor Account permissions : ", "err", err)
	}
	ch := make(chan *pbind.PermissionsAccountAccessModified)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newEvent *pbind.PermissionsAccountAccessModified

	_, err = permissions.WatchAccountAccessModified(opts, ch)
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
// This will format the EnodeId and return
func formatEnodeId(enodeId, ipAddrPort, discPort, raftPort string, isRaft bool) string {
	newEnodeId := "enode://" + enodeId + "@" + ipAddrPort + "?discPort=" + discPort
	if isRaft {
		newEnodeId += "&raftport=" + raftPort
	}
	return newEnodeId
}

//populates the nodes list from permissioned-nodes.json into the permissions
//smart contract
func (p *PermissionCtrl) populateStaticNodesToContract() {

	permissionsContract, err := pbind.NewPermissions(params.QuorumPermissionsContract, p.ethClnt)

	if err != nil {
		utils.Fatalf("Failed to instantiate a Permissions contract: %v", err)
	}
	auth := bind.NewKeyedTransactor(p.key)
	if err != nil {
		utils.Fatalf("Failed to create authorized transactor: %v", err)
	}

	permissionsSession := &pbind.PermissionsSession{
		Contract: permissionsContract,
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
		log.Warn("Failed to udpate network boot status ", "err", err)
	}
	if tx != true {
		datadir := p.node.InstanceDir()

		nodes := p2p.ParsePermissionedNodes(datadir)
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
			}
			log.Debug("Transaction pending", "tx hash", tx.Hash())
		}
		// update the network boot status to true
		// nonce := p.eth.TxPool().Nonce(permissionsSession.TransactOpts.From)
		// permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)

		// _, err := permissionsSession.UpdateNetworkBootStatus()
		// if err != nil {
		// 	log.Warn("Failed to udpate network boot status ", "err", err)
		// }
	}
}
