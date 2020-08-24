package types

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/raft"
)

// to signal all watches when service is stopped
type StopEvent struct {
}

// broadcasting stopEvent when service is being stopped
var StopFeed event.Feed

type NodeOperation uint8

const (
	NodeAdd NodeOperation = iota
	NodeDelete
)

type Backend interface {
	// Monitors account access related events and updates the cache accordingly
	ManageAccountPermissions() error
	// Monitors Node management events and updates cache accordingly
	ManageNodePermissions() error
	// monitors org management related events happening via smart contracts
	// and updates cache accordingly
	ManageOrgPermissions() error
	// monitors role management related events and updated cache
	ManageRolePermissions() error
}

// adds or deletes and entry from a given file
func UpdateFile(fileName, enodeId string, operation NodeOperation, createFile bool) error {
	// Load the nodes from the config file
	var nodeList []string
	index := 0
	// if createFile is false means the file is already existing. read the file
	if !createFile {
		blob, err := ioutil.ReadFile(fileName)
		if err != nil && !createFile {
			return err
		}

		if err := json.Unmarshal(blob, &nodeList); err != nil {
			return err
		}

		// logic to update the permissioned-nodes.json file based on action

		recExists := false
		for i, eid := range nodeList {
			if eid == enodeId {
				index = i
				recExists = true
				break
			}
		}
		if (operation == NodeAdd && recExists) || (operation == NodeDelete && !recExists) {
			return nil
		}
	}
	if operation == NodeAdd {
		nodeList = append(nodeList, enodeId)
	} else {
		nodeList = append(nodeList[:index], nodeList[index+1:]...)
	}
	blob, _ := json.Marshal(nodeList)

	var mux sync.Mutex
	mux.Lock()
	defer mux.Unlock()

	err := ioutil.WriteFile(fileName, blob, 0644)
	return err
}

//this function populates the black listed Node information into the disallowed-nodes.json file
func UpdateDisallowedNodes(dataDir, url string, operation NodeOperation) error {

	fileExists := true
	path := filepath.Join(dataDir, params.BLACKLIST_CONFIG)
	// Check if the file is existing. If the file is not existing create the file
	if _, err := os.Stat(path); err != nil {
		if _, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644); err != nil {
			return err
		}
		fileExists = false
	}

	if fileExists {
		err := UpdateFile(path, url, operation, false)
		return err
	} else {
		err := UpdateFile(path, url, operation, true)
		return err
	}
}

// Disconnect the Node from the network
func DisconnectNode(node *node.Node, enodeId string, isRaft bool) error {
	if isRaft {
		var raftService *raft.RaftService
		if err := node.Service(&raftService); err == nil {
			raftApi := raft.NewPublicRaftAPI(raftService)

			//get the raftId for the given enodeId
			raftId, err := raftApi.GetRaftId(enodeId)
			if err == nil {
				raftApi.RemovePeer(raftId)
			} else {
				return err
			}
		}
	} else {
		// Istanbul  or clique - disconnect the peer
		server := node.Server()
		if server != nil {
			node, err := enode.ParseV4(enodeId)
			if err == nil {
				server.RemovePeer(node)
			} else {
				return err
			}
		}
	}
	return nil
}

// updates Node information in the permissioned-nodes.json file based on Node
// management activities in smart contract
func UpdatePermissionedNodes(node *node.Node, dataDir, enodeId string, operation NodeOperation, isRaft bool) error {
	path := filepath.Join(dataDir, params.PERMISSIONED_CONFIG)
	if _, err := os.Stat(path); err != nil {
		return err
	}

	err := UpdateFile(path, enodeId, operation, false)
	if err != nil {
		return err
	}
	if operation == NodeDelete {
		err := DisconnectNode(node, enodeId, isRaft)
		if err != nil {
			return err
		}
	}
	return nil
}

// function to subscribe to the stop event
func SubscribeStopEvent() (chan StopEvent, event.Subscription) {
	c := make(chan StopEvent)
	s := StopFeed.Subscribe(c)
	return c, s
}
