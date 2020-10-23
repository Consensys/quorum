package types

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
)

const (
	NODE_NAME_LENGTH = 32
)

//TODO update this based on permission changes
// check if a given node is permissioned to connect to the change
func IsNodePermissioned(nodename string, currentNode string, datadir string, direction string) bool {
	var permissionedList []string
	nodes := ParsePermissionedNodes(datadir)
	for _, v := range nodes {
		permissionedList = append(permissionedList, v.ID().String())
	}

	log.Debug("isNodePermissioned", "permissionedList", permissionedList)
	for _, v := range permissionedList {
		if v == nodename {
			log.Debug("isNodePermissioned", "connection", direction, "nodename", nodename[:NODE_NAME_LENGTH], "ALLOWED-BY", currentNode[:NODE_NAME_LENGTH])
			// check if the node is blacklisted
			return !isNodeBlackListed(nodename, datadir)
		}
	}
	log.Debug("isNodePermissioned", "connection", direction, "nodename", nodename[:NODE_NAME_LENGTH], "DENIED-BY", currentNode[:NODE_NAME_LENGTH])
	return false
}

//this is a shameless copy from the config.go. It is a duplication of the code
//for the timebeing to allow reload of the permissioned nodes while the server is running

func ParsePermissionedNodes(DataDir string) []*enode.Node {

	log.Debug("parsePermissionedNodes", "DataDir", DataDir, "file", params.PERMISSIONED_CONFIG)

	path := filepath.Join(DataDir, params.PERMISSIONED_CONFIG)
	if _, err := os.Stat(path); err != nil {
		log.Error("Read Error for permissioned-nodes.json file. This is because 'permissioned' flag is specified but no permissioned-nodes.json file is present.", "err", err)
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

// This function checks if the node is black-listed
func isNodeBlackListed(nodeName, dataDir string) bool {
	log.Debug("isNodeBlackListed", "DataDir", dataDir, "file", params.BLACKLIST_CONFIG)

	path := filepath.Join(dataDir, params.BLACKLIST_CONFIG)
	if _, err := os.Stat(path); err != nil {
		log.Debug("Read Error for disallowed-nodes.json file. disallowed-nodes.json file is not present.", "err", err)
		return false
	}
	// Load the nodes from the config file
	blob, err := ioutil.ReadFile(path)
	if err != nil {
		log.Debug("isNodeBlackListed: Failed to access nodes", "err", err)
		return true
	}

	nodelist := []string{}
	if err := json.Unmarshal(blob, &nodelist); err != nil {
		log.Debug("parsePermissionedNodes: Failed to load nodes", "err", err)
		return true
	}

	for _, v := range nodelist {
		if strings.Contains(v, nodeName) {
			return true
		}
	}
	return false
}
