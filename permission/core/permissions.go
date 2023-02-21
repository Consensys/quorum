package core

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
)

type FileBasedPermissioning struct {
	PermissionFile string
	DisallowedFile string
}

var defaultFileBasedPermissioning = FileBasedPermissioning{
	PermissionFile: params.PERMISSIONED_CONFIG,
	DisallowedFile: params.DISALLOWED_CONFIG,
}

func NewFileBasedPermissoningWithPrefix(prefix string) FileBasedPermissioning {
	return FileBasedPermissioning{
		PermissionFile: prefix + "-" + params.PERMISSIONED_CONFIG,
		DisallowedFile: prefix + "-" + params.DISALLOWED_CONFIG,
	}
}

func IsNodePermissioned(nodename string, currentNode string, datadir string, direction string) bool {
	return defaultFileBasedPermissioning.IsNodePermissioned(nodename, currentNode, datadir, direction)
}

func isNodeDisallowed(nodeName, dataDir string) bool {
	return defaultFileBasedPermissioning.isNodeDisallowed(nodeName, dataDir)
}

func (fbp *FileBasedPermissioning) IsNodePermissionedEnode(node *enode.Node, nodename string, currentNode string, datadir string, direction string) bool {
	return fbp.IsNodePermissioned(nodename, currentNode, datadir, direction)
}

// check if a given node is permissioned to connect to the change
func (fbp *FileBasedPermissioning) IsNodePermissioned(nodename string, currentNode string, datadir string, direction string) bool {
	var permissionedList []string
	nodes := fbp.ParsePermissionedNodes(datadir)
	for _, v := range nodes {
		permissionedList = append(permissionedList, v.ID().String())
	}

	log.Debug("IsNodePermissioned", "permissionedList", permissionedList)
	for _, v := range permissionedList {
		if v == nodename {
			log.Debug("IsNodePermissioned", "connection", direction, "nodename", nodename[:params.NODE_NAME_LENGTH], "ALLOWED-BY", currentNode[:params.NODE_NAME_LENGTH])
			// check if the node is disallowed
			return !fbp.isNodeDisallowed(nodename, datadir)
		}
	}
	log.Debug("IsNodePermissioned", "connection", direction, "nodename", nodename[:params.NODE_NAME_LENGTH], "DENIED-BY", currentNode[:params.NODE_NAME_LENGTH])
	return false
}

//this is a shameless copy from the config.go. It is a duplication of the code
//for the timebeing to allow reload of the permissioned nodes while the server is running

func (fbp *FileBasedPermissioning) ParsePermissionedNodes(DataDir string) []*enode.Node {
	log.Debug("parsePermissionedNodes", "DataDir", DataDir, "file", fbp.PermissionFile)

	path := filepath.Join(DataDir, fbp.PermissionFile)
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for permissioned-nodes file. This is because 'permissioned' flag is specified but no permissioned-nodes file is present.", "fileName", fbp.PermissionFile, "err", err)
		return nil
	}
	// Load the nodes from the config file
	blob, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("parsePermissionedNodes: Failed to access nodes", "err", err)
		return nil
	}

	nodelist := []string{}
	if err := json.Unmarshal(blob, &nodelist); err != nil {
		log.Error("parsePermissionedNodes: Failed to load nodes", "err", err)
		return nil
	}
	// Interpret the list as a discovery node array
	var nodes []*enode.Node
	for _, url := range nodelist {
		if url == "" {
			log.Error("parsePermissionedNodes: Node URL blank")
			continue
		}
		node, err := enode.ParseV4(url)
		if err != nil {
			log.Error("parsePermissionedNodes: Node URL", "url", url, "err", err)
			continue
		}
		nodes = append(nodes, node)
	}
	return nodes
}

// This function checks if the node is disallowed
func (fbp *FileBasedPermissioning) isNodeDisallowed(nodeName, dataDir string) bool {
	log.Debug("isNodeDisallowed", "DataDir", dataDir, "file", fbp.DisallowedFile)

	path := filepath.Join(dataDir, fbp.DisallowedFile)
	if _, err := os.Stat(path); err != nil {
		log.Debug("Read Error for disallowed-nodes file. disallowed-nodes file is not present.", "fileName", fbp.DisallowedFile, "err", err)
		return false
	}
	// Load the nodes from the config file
	blob, err := ioutil.ReadFile(path)
	if err != nil {
		log.Debug("isNodeDisallowed: Failed to access nodes", "err", err)
		return true
	}

	nodelist := []string{}
	if err := json.Unmarshal(blob, &nodelist); err != nil {
		log.Debug("parsePermissionedNodes: Failed to load nodes", "err", err)
		return true
	}

	for _, v := range nodelist {
		n, _ := enode.ParseV4(v)
		if nodeName == n.ID().String() {
			return true
		}
	}
	return false
}

// function checks for account access to execute the transaction
func CheckAccountPermission(from common.Address, to *common.Address, value *big.Int, data []byte, gas uint64, gasPrice *big.Int) error {
	transactionType := ValueTransferTxn

	if to == nil {
		transactionType = ContractDeployTxn
	} else if data != nil {
		transactionType = ContractCallTxn
	}

	var toAcct common.Address

	if to == nil {
		toAcct = common.Address{}
	} else {
		toAcct = *to
	}

	return IsTransactionAllowed(from, toAcct, value, gasPrice, big.NewInt(int64(gas)), data, transactionType)
}
