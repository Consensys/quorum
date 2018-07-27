package permissions

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"gopkg.in/urfave/cli.v1"
)
const (
	PERMISSIONED_CONFIG = "permissioned-nodes.json"
)

//This function first adds the node list from permissioned-nodes.json to
//the permissiones contract deployed as a precompile via genesis.json
func QuorumPermissioning(ctx *cli.Context, stack *node.Node ){
	//Create a new ethclient to for interfacing with the contract
	e, stateReader := createEthClient(stack)

	//call populate nodes to populate the nodes into contract
	populateNodesToContract (ctx, stack, e, stateReader)

	//monitor for new nodes addition via smart contract
	go monitorNewNodeAdd(stack, stateReader)

	//monitor for nodes deletiin via smart contract
	go monitorNodeDelete(stack, stateReader)
}

//populates the nodes list from permissioned-nodes.json into the permissions
//smart contract
func populateNodesToContract(ctx *cli.Context, stack *node.Node, e *eth.Ethereum, stateReader *ethclient.Client){

	//Read the key file from key store. SHOULD WE MAKE IT CONFIG value
	key := getKeyFromKeyStore(ctx)

	permissionsContract, err := NewPermissions(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		utils.Fatalf("Failed to instantiate a Permissions contract: %v", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	if err != nil {
		utils.Fatalf("Failed to create authorized transactor: %v", err)
	}

	permissionsSession := &PermissionsSession{
		Contract: permissionsContract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 3558096384,
			GasPrice: big.NewInt(0),
		},
	}

	// datadir := ctx.GlobalString(utils.DataDirFlag.Name)
	datadir := stack.DataDir()

	nodes := p2p.ParsePermissionedNodes(datadir)
	for _, node := range nodes {
		enodeID := fmt.Sprintf("%x", node.ID[:])
		log.Trace("Adding node to permissions contract", "enodeID", enodeID)

		nonce := e.TxPool().Nonce(permissionsSession.TransactOpts.From)
		permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)

		tx, err := permissionsSession.ProposeNode(enodeID, true, true)
		if err != nil {
			log.Warn("Failed to propose node", "err", err)
		}
		log.Debug("Transaction pending", "tx hash", tx.Hash())
	}
}

//This functions listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func monitorNewNodeAdd(stack *node.Node, stateReader *ethclient.Client){

	permissions, err := NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		utils.Fatalf("Failed to instantiate a Permissions Filterer: %v", err)
	}
	datadir := stack.DataDir()

	ch := make(chan *PermissionsNewNodeProposed)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	var operation string = "ADD"

	for {
		_, err = permissions.WatchNewNodeProposed(opts, ch)
		if err != nil {
			log.Info("Failed NewNodeProposed: %v", err)
		}
		var newEvent *PermissionsNewNodeProposed = <-ch
		updatePermissionedNodes(newEvent.EnodeId, datadir, operation)
    }
}

//This functions listens on the channel for new node approval via smart contract and
// adds the same into permissioned-nodes.json
func monitorNodeDelete(stack *node.Node, stateReader *ethclient.Client){

	permissions, err := NewPermissionsFilterer(params.QuorumPermissionsContract, stateReader)
	if err != nil {
		utils.Fatalf("Failed to instantiate a Permissions Filterer: %v", err)
	}
	datadir := stack.DataDir()

	ch := make(chan *PermissionsNodeDeactivated)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	var operation string = "DEL"

	for {
		_, err = permissions.WatchNodeDeactivated(opts, ch)
		if err != nil {
			log.Info("Failed NodeDeactivated: %v", err)
		}
		var newEvent *PermissionsNodeDeactivated = <-ch

		updatePermissionedNodes(newEvent.EnodeId, datadir, operation)
    }
}
//Create an RPC client for the contract interface
func createEthClient(stack *node.Node ) (*eth.Ethereum, *ethclient.Client){
	var e *eth.Ethereum
	if err := stack.Service(&e); err != nil {
		utils.Fatalf("Ethereum service not running: %v", err)
	}

	rpcClient, err := stack.Attach()
	if err != nil {
		utils.Fatalf("Failed to attach to self: %v", err)
	}

	return e, ethclient.NewClient(rpcClient)

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
	n := len(keyBlob)

	return string(keyBlob[:n])

}

//this function populates the new node information into the permissioned-nodes.json file
func updatePermissionedNodes(enodeId string, dataDir string, operation string){
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
	newEnodeId := "enode://" + enodeId + "@127.0.0.1:21005?discport=0&raftport=50406"

	if (operation == "ADD"){
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

	blob, _ = json.Marshal(nodelist)
	if err:= ioutil.WriteFile(path, blob, 0644); err!= nil{
		log.Error("updatePermissionedNodes: Error writing new node info to file", "err", err)
	}

}

