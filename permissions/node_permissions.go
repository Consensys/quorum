package permissions

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"os"
	"sync"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"gopkg.in/urfave/cli.v1"
)
const (
	PERMISSIONED_CONFIG = "permissioned-nodes.json"
	BLACKLIST_CONFIG = "disallowed-nodes.json"
)

type NodeOperation uint8

const (
	NodeAdd NodeOperation = iota
	NodeDelete
)

// This function first adds the node list from permissioned-nodes.json to
// the permissiones contract deployed as a precompile via genesis.json
func QuorumPermissioning(ctx *cli.Context, stack *node.Node ) error {

	// Create a new ethclient to for interfacing with the contract
	stateReader, err := createEthClient(stack)
	if err != nil {
		log.Error ("Unable to create ethereum client for permissions check : ", "err" , err)
		return err
	}

	// check if permissioning contract is there at address. If not return from here
	if _ , err = NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader); err != nil {
		log.Error ("Permissions not enabled for the network : ", "err" , err)
		return nil
	}

	// Monitors node addition and decativation from network
	manageNodePermissions(stack, stateReader);

	// Monitors account level persmissions  update from smart contarct 
	manageAccountPermissions(stack, stateReader);

	return nil
}

// Create an RPC client for the contract interface
func createEthClient(stack *node.Node ) (*ethclient.Client, error){
	var e *eth.Ethereum
	if err := stack.Service(&e); err != nil {
		return nil, err
	}

	rpcClient, err := stack.Attach()
	if err != nil {
		return nil, err
	}

	return ethclient.NewClient(rpcClient), nil
}

// Manages node addition and decavtivation from network
func manageNodePermissions(stack *node.Node, stateReader *ethclient.Client) {
	//monitor for new nodes addition via smart contract
	go monitorNewNodeAdd(stack, stateReader)

	//monitor for nodes deletiin via smart contract
	go monitorNodeDeactivation(stack, stateReader)

	//monitor for nodes blacklisting via smart contract
	go monitorNodeBlacklisting(stack, stateReader)
}

// This functions listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func monitorNewNodeAdd(stack *node.Node, stateReader *ethclient.Client) {

	permissions, err := NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("failed to monitor new node add : ", "err" , err)
	}
	datadir := stack.DataDir()

	ch := make(chan *PermissionsNodeApproved, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var nodeAddEvent *PermissionsNodeApproved 

	_, err = permissions.WatchNodeApproved(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeApproved: %v", err)
	}

	for {
		select {
		case nodeAddEvent = <-ch:
			updatePermissionedNodes(nodeAddEvent.EnodeId, nodeAddEvent.IpAddrPort, nodeAddEvent.DiscPort, nodeAddEvent.RaftPort, datadir, NodeAdd)
		}
    }
}

// This functions listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func monitorNodeDeactivation(stack *node.Node, stateReader *ethclient.Client) {

	permissions, err := NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("Failed to monitor node delete: ", "err" , err)
	}
	datadir := stack.DataDir()

	ch := make(chan *PermissionsNodeDeactivated)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newNodeDeleteEvent *PermissionsNodeDeactivated 

	_, err = permissions.WatchNodeDeactivated(opts, ch)
	if err != nil {
		log.Info("Failed NodeDeactivated: %v", err)
	}

	for {
		select {
		case newNodeDeleteEvent = <-ch:
			updatePermissionedNodes(newNodeDeleteEvent.EnodeId, newNodeDeleteEvent.IpAddrPort, newNodeDeleteEvent.DiscPort, newNodeDeleteEvent.RaftPort, datadir, NodeDelete)
	  }

	}
}

// This function listnes on the channel for any node blacklisting event via smart contract
// and adds the same disallowed-nodes.json
func monitorNodeBlacklisting(stack *node.Node, stateReader *ethclient.Client) {

	permissions, err := NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("failed to monitor new node add : ", "err" , err)
	}
	datadir := stack.DataDir()

	ch := make(chan *PermissionsNodeBlacklisted, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var nodeBlacklistEvent *PermissionsNodeBlacklisted 

	_, err = permissions.WatchNodeBlacklisted(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeBlacklisted: %v", err)
	}

	for {
		select {
		case nodeBlacklistEvent = <-ch:
			updateDisallowedNodes(nodeBlacklistEvent.EnodeId, nodeBlacklistEvent.IpAddrPort, nodeBlacklistEvent.DiscPort, nodeBlacklistEvent.RaftPort, datadir)
		}
    }
}

//this function populates the new node information into the permissioned-nodes.json file
func updatePermissionedNodes(enodeId , ipAddrPort, discPort, raftPort, dataDir string, operation NodeOperation){
	log.Debug("updatePermissionedNodes", "DataDir", dataDir, "file", PERMISSIONED_CONFIG)

	path := filepath.Join(dataDir, PERMISSIONED_CONFIG)
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

	// HACK: currently the ip, discpot and raft port are hard coded. Need to enhance the
	//contract to pass these variables as part of the event and change this
	// newEnodeId := "enode://" + enodeId + "@127.0.0.1:21005?discport=0&raftport=50406"
	newEnodeId := "enode://" + enodeId + "@" + ipAddrPort + "?discPort=" + discPort + "&raftport=" + raftPort
	log.Info("Enode id is : " , "newEnodeId", newEnodeId)


	if (operation == NodeAdd){
		nodelist = append(nodelist, newEnodeId)
	} else {
		index := 0
		for i, enodeId := range nodelist {
			if (enodeId == newEnodeId){
				index = i
				break
			}
		}
		nodelist = append(nodelist[:index], nodelist[index+1:]...)
	}

	mu := sync.RWMutex{}
	blob, _ = json.Marshal(nodelist)

	mu.Lock()
	if err:= ioutil.WriteFile(path, blob, 0644); err!= nil{
		log.Error("updatePermissionedNodes: Error writing new node info to file", "err", err)
	}
	mu.Unlock()

}

//this function populates the new node information into the permissioned-nodes.json file
func updateDisallowedNodes(enodeId , ipAddrPort, discPort, raftPort, dataDir string){
	log.Debug("updateDisallowedNodes", "DataDir", dataDir, "file", BLACKLIST_CONFIG)

	fileExisted := true
	path := filepath.Join(dataDir, BLACKLIST_CONFIG)
	// Check if the file is existing. If the file is not existing create the file
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for disallowed-nodes.json file." , "err", err)
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
		if (blob != nil) {
			if err := json.Unmarshal(blob, &nodelist); err != nil {
				log.Error("updateDisallowedNodes: Failed to load nodes list", "err", err)
				return 
			}
		}
	}

	newEnodeId := "enode://" + enodeId + "@" + ipAddrPort + "?discPort=" + discPort + "&raftport=" + raftPort
	log.Info("Enode id is : " , "newEnodeId", newEnodeId)

	nodelist = append(nodelist, newEnodeId)

	mu := sync.RWMutex{}
	blob, _ := json.Marshal(nodelist)
	mu.Lock()
	if err:= ioutil.WriteFile(path, blob, 0644); err!= nil{
		log.Error("updateDisallowedNodes: Error writing new node info to file", "err", err)
	}
	mu.Unlock()
}

// Manages account level permissions update
func manageAccountPermissions(stack *node.Node, stateReader *ethclient.Client) error {

	//call populate nodes to populate the nodes into contract
	if err := populateAcctPermissions (stack, stateReader); err != nil {
		return err;
	}

	//monitor for nodes deletiin via smart contract
	go monitorAccountPermissions(stack, stateReader)
	return nil
}

// populates the nodes list from permissioned-nodes.json into the permissions
// smart contract
func populateAcctPermissions(stack *node.Node, stateReader *ethclient.Client) error{

	permissions, err := NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("Failed to monitor node delete: ", "err" , err)
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
func monitorAccountPermissions(stack *node.Node, stateReader *ethclient.Client) {

	log.Info("Inside monotorAccountPermissions")

	permissions, err := NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("Failed to monitor Account permissions : ", "err" , err)
	}
	ch := make(chan *PermissionsAccountAccessModified)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newEvent *PermissionsAccountAccessModified

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
