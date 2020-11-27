package permission

import (
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/permission/core"
)

func isNodePermissionedBasic(enodeId string, nodename string, currentNode string, direction string) bool {
	permissionedList := core.NodeInfoMap.GetNodeList()

	log.Debug("isNodePermissionedBasic", "permissionedList", permissionedList)
	for _, n := range permissionedList {
		if strings.Contains(n.Url, enodeId) && n.Status == core.NodeApproved {
			log.Debug("isNodePermissionedBasic", "connection", direction, "nodename", nodename[:params.NODE_NAME_LENGTH], "ALLOWED-BY", currentNode[:params.NODE_NAME_LENGTH])
			return true
		}
	}
	log.Debug("isNodePermissionedBasic", "connection", direction, "nodename", nodename[:params.NODE_NAME_LENGTH], "DENIED-BY", currentNode[:params.NODE_NAME_LENGTH])
	return false
}

func isNodePermissionedEEA(node *enode.Node, nodename string, currentNode string, direction string) bool {
	if permissionService == nil {
		log.Debug("isNodePermissionedEEA connection not allowed - permissionService is not set")
		return false
	}
	allowed, err := permissionService.ConnectionAllowed(node.EnodeID(), node.IP().String(), uint16(node.TCP()), uint16(node.RaftPort()))
	log.Debug("isNodePermissionedEEA EEA", "allowed", allowed, "url", node.String())
	if err != nil {
		log.Error("isNodePermissionedEEA connection not allowed", "err", err)
		return false
	}
	if allowed {
		log.Debug("isNodePermissionedEEA", "connection", direction, "nodename", nodename[:params.NODE_NAME_LENGTH], "ALLOWED-BY", currentNode[:params.NODE_NAME_LENGTH])
	} else {
		log.Debug("isNodePermissionedEEA", "connection", direction, "nodename", nodename[:params.NODE_NAME_LENGTH], "DENIED-BY", currentNode[:params.NODE_NAME_LENGTH])

	}
	return allowed

}

func IsNodePermissioned(node *enode.Node, nodename string, currentNode string, datadir string, direction string) bool {

	//if we have not reached QIP714 block return full access
	if !core.PermissionsEnabled() {
		return core.IsNodePermissioned(nodename, currentNode, datadir, direction)
	}

	switch core.PermissionModel {
	case core.Default:
		return core.IsNodePermissioned(nodename, currentNode, datadir, direction)

	case core.Basic:
		return isNodePermissionedBasic(node.EnodeID(), nodename, currentNode, direction)

	case core.EEA:
		return isNodePermissionedEEA(node, nodename, currentNode, direction)
	}
	return false
}
