package p2p

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
	"github.com/ethereum/go-ethereum/p2p/discover"
)

const (
	NODE_NAME_LENGTH    = 32
	PERMISSIONED_CONFIG = "permissioned-nodes.json"
)

// check if a given node is permissioned to connect to the change
func isNodePermissioned(nodename string, currentNode string, datadir string, direction string) bool {

	var permissonedList []string
	nodes := parsePermissionedNodes(datadir)
	for _, v := range nodes {
		permissonedList = append(permissonedList, v.ID.String())
	}

	glog.V(logger.Debug).Infof("Permisssioned_list %v", permissonedList)
	for _, v := range permissonedList {
		if v == nodename {
			glog.V(logger.Debug).Infof("isNodePermissioned <%v> connection:: nodename <%v> ALLOWED-BY <%v>", direction, nodename[:NODE_NAME_LENGTH], currentNode[:NODE_NAME_LENGTH])
			return true
		}
		glog.V(logger.Debug).Infof("isNodePermissioned <%v> connection:: nodename <%v> DENIED-BY <%v>", direction, nodename[:NODE_NAME_LENGTH], currentNode[:NODE_NAME_LENGTH])
	}
	glog.V(logger.Debug).Infof("isNodePermissioned <%v> connection:: nodename <%v> DENIED-BY <%v>", direction, nodename[:NODE_NAME_LENGTH], currentNode[:NODE_NAME_LENGTH])
	return false
}

//this is a shameless copy from the config.go. It is a duplication of the code
//for the timebeing to allow reload of the permissioned nodes while the server is running

func parsePermissionedNodes(DataDir string) []*discover.Node {

	glog.V(logger.Debug).Infof("parsePermissionedNodes DataDir %v, file %v", DataDir, PERMISSIONED_CONFIG)

	path := filepath.Join(DataDir, PERMISSIONED_CONFIG)
	if _, err := os.Stat(path); err != nil {
		glog.V(logger.Error).Infof("Read Error for permissioned-nodes.json file %v. This is because 'permissioned' flag is specified but no permissioned-nodes.json file is present.", err)
		return nil
	}
	// Load the nodes from the config file
	blob, err := ioutil.ReadFile(path)
	if err != nil {
		glog.V(logger.Error).Infof("parsePermissionedNodes: Failed to access nodes: %v", err)
		return nil
	}

	nodelist := []string{}
	if err := json.Unmarshal(blob, &nodelist); err != nil {
		glog.V(logger.Error).Infof("parsePermissionedNodes: Failed to load nodes: %v", err)
		return nil
	}
	// Interpret the list as a discovery node array
	var nodes []*discover.Node
	for _, url := range nodelist {
		if url == "" {
			glog.V(logger.Error).Infof("parsePermissionedNodes: Node URL blank")
			continue
		}
		node, err := discover.ParseNode(url)
		if err != nil {
			glog.V(logger.Error).Infof("parsePermissionedNodes: Node URL %s: %v\n", url, err)
			continue
		}
		nodes = append(nodes, node)
	}
	return nodes
}

