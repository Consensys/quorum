package permissions

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"os"
	// "sync"

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
	go monitorNewNodeAdd(stack)

	//monitor for nodes deletiin via smart contract
	go monitorNodeDelete(stack, stateReader)
}

// This functions listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func monitorNewNodeAdd(stack *node.Node) {

	stateReader, err := createEthClient(stack)

	permissions, err := NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("failed to monitor new node add : ", "err" , err)
	}
	datadir := stack.DataDir()

	ch := make(chan *PermissionsNodeApproved)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	for {
		_, err = permissions.WatchNodeApproved(opts, ch)
		if err != nil {
			log.Info("Failed NewNodeProposed: %v", err)
		}
		var nodeAddEvent *PermissionsNodeApproved = <-ch

		log.Info("calling update permissions, length is", "event", nodeAddEvent)

		updatePermissionedNodes(nodeAddEvent.EnodeId, nodeAddEvent.IpAddrPort, nodeAddEvent.DiscPort, nodeAddEvent.RaftPort, datadir, NodeAdd)
    }
}

// This functions listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func monitorNodeDelete(stack *node.Node, stateReader *ethclient.Client) {

	permissions, err := NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("Failed to monitor node delete: ", "err" , err)
	}
	datadir := stack.DataDir()

	ch := make(chan *PermissionsNodeDeactivated)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	for {
		_, err = permissions.WatchNodeDeactivated(opts, ch)
		if err != nil {
			log.Info("Failed NodeDeactivated: %v", err)
		}
		var newEvent *PermissionsNodeDeactivated = <-ch

		updatePermissionedNodes(newEvent.EnodeId, newEvent.IpAddrPort, newEvent.DiscPort, newEvent.RaftPort, datadir, NodeDelete)
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
		log.Error("parsePermissionedNodes: Failed to load nodes list", "err", err)
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

	// mu := sync.RWMutex{}
	blob, _ = json.Marshal(nodelist)

	// mu.Lock()
	if err:= ioutil.WriteFile(path, blob, 0644); err!= nil{
		log.Error("updatePermissionedNodes: Error writing new node info to file", "err", err)
	}
	// mu.Unlock()

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

	pastEvents, err := permissions.FilterAcctAccessModified(opts)

	recExists := true
	for recExists {
		recExists = pastEvents.Next()
		if recExists {
			types.AddAccountAccess(pastEvents.Event.AcctId, pastEvents.Event.Access)
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
	ch := make(chan *PermissionsAcctAccessModified)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	// const addr1 = "ed9d02e382b34818e88b88a309c7fe71e65f419d"
	// var acctAddr1 = common.HexToAddress(addr1)
	// types.AddAccountAccess(acctAddr1, 0)

	// const addr2 = "ca843569e3427144cead5e4d5999a3d0ccf92b8e"
	// var acctAddr2 = common.HexToAddress(addr2)
	// types.AddAccountAccess(acctAddr2, 1)

	// const addr3 = "0fbdc686b912d7722dc86510934589e0aaf3b55a"
	// var acctAddr3= common.HexToAddress(addr3)
	// types.AddAccountAccess(acctAddr3, 2)

	// const addr4 = "9186eb3d20cbd1f5f992a950d808c4495153abd5"
	// var acctAddr4= common.HexToAddress(addr4)
	// types.AddAccountAccess(acctAddr4, 3)

	for {
		_, err = permissions.WatchAcctAccessModified(opts, ch)
		if err != nil {
			log.Info("Failed NewNodeProposed: %v", err)
		}
		var newEvent *PermissionsAcctAccessModified = <-ch
		log.Info("caught the event and calling PutAcctMap")
		types.AddAccountAccess(newEvent.AcctId, newEvent.Access)
    }
}
