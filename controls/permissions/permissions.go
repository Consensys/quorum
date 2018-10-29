package permissions

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"math/big"
	"os"
	"sync"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/controls"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/raft"
	"gopkg.in/urfave/cli.v1"
	"github.com/ethereum/go-ethereum/controls/permbind"
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

func SayHello(n string) string{
	return "Hello " + n + "!"
}

// This function first adds the node list from permissioned-nodes.json to
// the permissiones contract deployed as a precompile via genesis.json
func QuorumPermissioning(ctx *cli.Context, stack *node.Node ) error {
	// Create a new ethclient to for interfacing with the contract
	stateReader, e, err := controls.CreateEthClient(stack)
	if err != nil {
		log.Error ("Unable to create ethereum client for permissions check : ", "err" , err)
		return err
	}

	// check if permissioning contract is there at address. If not return from here
	if _ , err = permbind.NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader); err != nil {
		log.Error ("Permissions not enabled for the network : ", "err" , err)
		return nil
	}
	isRaft := false
	if ctx.GlobalBool(utils.RaftModeFlag.Name) {
		isRaft = true
	}

	// Permissions initialization
	err = permissionsInit(ctx, stack, e, stateReader, isRaft)

	// Monitors node addition and decativation from network
	manageNodePermissions(ctx, stack, e, stateReader, isRaft);

	// Monitors account level persmissions  update from smart contarct 
	manageAccountPermissions(stack, stateReader);

	return nil
}

// This functions updates the initial  values for the network
func permissionsInit(ctx *cli.Context, stack *node.Node, e *eth.Ethereum, stateReader *ethclient.Client, isRaft bool) error {
	// populate the initial list of nodes into the smart contract
	// from permissioned-nodes.json
	populateStaticNodesToContract(ctx, stack, e, stateReader)

	// populate the account access for the genesis.json accounts. these
	// accounts will have full access
	// populateInitAccountAccess()


	// call populates the node details from contract to KnownNodes
	if err := populatePermissionedNodes (stack, stateReader, isRaft); err != nil {
		return err
	}
	// call populates the account permissions based on past history
	if err := populateAcctPermissions (stack, stateReader); err != nil {
		return err
	}

	return nil
}

// Manages node addition and decavtivation from network
func manageNodePermissions(ctx *cli.Context, stack *node.Node, e *eth.Ethereum, stateReader *ethclient.Client, isRaft bool) {

	//monitor for new nodes addition via smart contract
	go monitorNewNodeAdd(stack, stateReader, isRaft)

	//monitor for nodes deletiin via smart contract
	go monitorNodeDeactivation(stack, stateReader, isRaft)

	//monitor for nodes blacklisting via smart contract
	go monitorNodeBlacklisting(stack, stateReader, isRaft)
}

// This functions listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func monitorNewNodeAdd(stack *node.Node, stateReader *ethclient.Client, isRaft bool) {
	permissions, err := permbind.NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("failed to monitor new node add : ", "err" , err)
	}

	ch := make(chan *permbind.PermissionsNodeApproved, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var nodeAddEvent *permbind.PermissionsNodeApproved

	_, err = permissions.WatchNodeApproved(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeApproved: %v", err)
	}

	for {
		select {
		case nodeAddEvent = <-ch:
			updatePermissionedNodes(stack, nodeAddEvent.EnodeId, nodeAddEvent.IpAddrPort, nodeAddEvent.DiscPort, nodeAddEvent.RaftPort, isRaft, NodeAdd)
		}
    }
}

// This functions listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func monitorNodeDeactivation(stack *node.Node, stateReader *ethclient.Client, isRaft bool) {
	permissions, err := permbind.NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("Failed to monitor node delete: ", "err" , err)
	}

	ch := make(chan *permbind.PermissionsNodeDeactivated)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newNodeDeleteEvent *permbind.PermissionsNodeDeactivated

	_, err = permissions.WatchNodeDeactivated(opts, ch)
	if err != nil {
		log.Info("Failed NodeDeactivated: %v", err)
	}

	for {
		select {
		case newNodeDeleteEvent = <-ch:
			updatePermissionedNodes(stack, newNodeDeleteEvent.EnodeId, newNodeDeleteEvent.IpAddrPort, newNodeDeleteEvent.DiscPort, newNodeDeleteEvent.RaftPort, isRaft, NodeDelete)
	  }

	}
}

// This function listnes on the channel for any node blacklisting event via smart contract
// and adds the same disallowed-nodes.json
func monitorNodeBlacklisting(stack *node.Node, stateReader *ethclient.Client, isRaft bool) {
	permissions, err := permbind.NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("failed to monitor new node add : ", "err" , err)
	}
	ch := make(chan *permbind.PermissionsNodeBlacklisted, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var nodeBlacklistEvent *permbind.PermissionsNodeBlacklisted

	_, err = permissions.WatchNodeBlacklisted(opts, ch)
	if err != nil {
		log.Info("Failed WatchNodeBlacklisted: %v", err)
	}

	for {
		select {
		case nodeBlacklistEvent = <-ch:
			updateDisallowedNodes(nodeBlacklistEvent, stack, isRaft)
		}
    }
}

//this function populates the new node information into the permissioned-nodes.json file
func updatePermissionedNodes(stack *node.Node, enodeId , ipAddrPort, discPort, raftPort string, isRaft bool, operation NodeOperation){
	newEnodeId := formatEnodeId(enodeId, ipAddrPort, discPort, raftPort, isRaft)

	//new logic to update the server KnownNodes variable for permissioning
	server := stack.Server();
	newNode, err := discover.ParseNode(newEnodeId)

	if err != nil {
		log.Error("updatePermissionedNodes: Node URL", "url", newEnodeId, "err", err)
	}

	if (operation == NodeAdd){
		// Add the new enode id to server.KnownNodes
		server.KnownNodes = append(server.KnownNodes, newNode)
	} else {
		// delete the new enode id from server.KnownNodes
		index := 0
		for i, node := range server.KnownNodes {
			if (node.ID == newNode.ID){
				index = i
			}
		}
		server.KnownNodes = append (server.KnownNodes[:index], server.KnownNodes[index+1:]...)
	}

}

//this function populates the new node information into the permissioned-nodes.json file
func updateDisallowedNodes(nodeBlacklistEvent *permbind.PermissionsNodeBlacklisted, stack *node.Node, isRaft bool){
	dataDir := stack.DataDir()
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

	newEnodeId := formatEnodeId (nodeBlacklistEvent.EnodeId, nodeBlacklistEvent.IpAddrPort, nodeBlacklistEvent.DiscPort, nodeBlacklistEvent.RaftPort, isRaft )
	nodelist = append(nodelist, newEnodeId)
	mu := sync.RWMutex{}
	blob, _ := json.Marshal(nodelist)
	mu.Lock()
	if err:= ioutil.WriteFile(path, blob, 0644); err!= nil{
		log.Error("updateDisallowedNodes: Error writing new node info to file", "err", err)
	}
	mu.Unlock()

	// Disconnect the peer if it is already connected
	disconnectNode(stack, newEnodeId, isRaft)
}

// Manages account level permissions update
func manageAccountPermissions(stack *node.Node, stateReader *ethclient.Client) error {
	//monitor for nodes deletiin via smart contract
	go monitorAccountPermissions(stack, stateReader)
	return nil
}
// populates the nodes list from permissioned-nodes.json into the permissions
// smart contract
func populatePermissionedNodes (stack *node.Node, stateReader *ethclient.Client, isRaft bool) error{
	permissions, err := permbind.NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("Failed to monitor node delete: ", "err" , err)
		return err
	}

	opts := &bind.FilterOpts{}
	pastAddEvent, err := permissions.FilterNodeApproved(opts)

	recExists := true
	for recExists {
		recExists = pastAddEvent.Next()
		if recExists {
			updatePermissionedNodes(stack, pastAddEvent.Event.EnodeId, pastAddEvent.Event.IpAddrPort, pastAddEvent.Event.DiscPort, pastAddEvent.Event.RaftPort, isRaft, NodeAdd)
		}
	}

	opts = &bind.FilterOpts{}
	pastDelEvent, err := permissions.FilterNodeDeactivated(opts)

	recExists = true
	for recExists {
		recExists = pastDelEvent.Next()
		if recExists {
			updatePermissionedNodes(stack, pastDelEvent.Event.EnodeId, pastDelEvent.Event.IpAddrPort, pastDelEvent.Event.DiscPort, pastDelEvent.Event.RaftPort, isRaft, NodeDelete)
		}
	}
	return nil
}

// populates the nodes list from permissioned-nodes.json into the permissions
// smart contract
func populateAcctPermissions(stack *node.Node, stateReader *ethclient.Client) error{
	permissions, err := permbind.NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
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
	permissions, err := permbind.NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		log.Error ("Failed to monitor Account permissions : ", "err" , err)
	}
	ch := make(chan *permbind.PermissionsAccountAccessModified)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var newEvent *permbind.PermissionsAccountAccessModified

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
func disconnectNode (stack *node.Node, enodeId string, isRaft bool){
	if isRaft {
		var raftService *raft.RaftService
		if err := stack.Service(&raftService); err == nil {
			raftApi := raft.NewPublicRaftAPI(raftService)

			//get the raftId for the given enodeId
			raftId, err := raftApi.GetRaftId(enodeId)
			if err == nil {
				raftApi.RemovePeer(raftId)
			}
		}
	} else {
		// Istanbul - disconnect the peer
		server := stack.Server()
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
func formatEnodeId( enodeId , ipAddrPort, discPort, raftPort string, isRaft bool) string {
	newEnodeId := "enode://" + enodeId + "@" + ipAddrPort + "?discPort=" + discPort
	if isRaft {
		newEnodeId +=  "&raftport=" + raftPort
	}
	return newEnodeId
}
//populates the nodes list from permissioned-nodes.json into the permissions
//smart contract
func populateStaticNodesToContract(ctx *cli.Context, stack *node.Node, e *eth.Ethereum, stateReader *ethclient.Client){
	//Read the key file from key store. SHOULD WE MAKE IT CONFIG value
	key := getKeyFromKeyStore(ctx)

	permissionsContract, err := permbind.NewPermissions(params.QuorumPermissionsContract, stateReader)

	if err != nil {
		utils.Fatalf("Failed to instantiate a Permissions contract: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	if err != nil {
		utils.Fatalf("Failed to create authorized transactor: %v", err)
	}

	permissionsSession := &permbind.PermissionsSession{
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
		datadir := ctx.GlobalString(utils.DataDirFlag.Name)

		nodes := p2p.ParsePermissionedNodes(datadir)
		for _, node := range nodes {

			enodeID := node.ID.String()
			ipAddr := node.IP.String()
			port := fmt.Sprintf("%v", node.TCP)
			discPort := fmt.Sprintf("%v", node.UDP)
			raftPort := fmt.Sprintf("%v", node.RaftPort)

			ipAddrPort := ipAddr + ":" + port

			log.Trace("Adding node to permissions contract", "enodeID", enodeID)

			nonce := e.TxPool().Nonce(permissionsSession.TransactOpts.From)
			permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)

			tx, err := permissionsSession.ProposeNode(enodeID, ipAddrPort, discPort, raftPort)
			if err != nil {
				log.Warn("Failed to propose node", "err", err)
			}
			log.Debug("Transaction pending", "tx hash", tx.Hash())
		}
		// update the network boot status to true
		nonce := e.TxPool().Nonce(permissionsSession.TransactOpts.From)
		permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)

		_, err := permissionsSession.UpdateNetworkBootStatus()
		if err != nil {
			log.Warn("Failed to udpate network boot status ", "err", err)
		}
	}
}

//This functions reads the first file in key store directory, reads the key
//value and returns the same
func getKeyFromKeyStore(ctx *cli.Context) string {
	datadir := ctx.GlobalString(utils.DataDirFlag.Name)

	files, err := ioutil.ReadDir(filepath.Join(datadir, "keystore"))
	if err != nil {
		utils.Fatalf("Failed to read keystore directory: %v", err)
	}

	// HACK: here we always use the first key as transactor
	var keyPath string
	for _, f := range files {
		keyPath = filepath.Join(datadir, "keystore", f.Name())
		break
	}
	keyBlob, err := ioutil.ReadFile(keyPath)
	if err != nil {
		utils.Fatalf("Failed to read key file: %v", err)
	}
	// n := bytes.IndexByte(keyBlob, 0)
	n := len(keyBlob)

	return string(keyBlob[:n])
}
