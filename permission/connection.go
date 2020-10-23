package permission

import (
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

func isNodePermissionedBasic(enodeId string) bool {
	for _, n := range types.NodeInfoMap.GetNodeList() {
		if n.Status == types.NodeApproved && strings.Contains(n.Url, enodeId) {
			log.Info("isNodePermissionedBasic check passed", "lk_url", enodeId, "c_url", n.Url)
			return true
		}
	}
	return false
}

func IsNodePermissioned(node *enode.Node, nodename string, currentNode string, datadir string, direction string) bool {
	var permissionType types.PermissionModelType

	if permissionService == nil {
		permissionType = types.Default
	} else if permissionService.eeaFlag {
		permissionType = types.EEA
	} else {
		permissionType = types.Basic
	}
	log.Info("AJ-permission IsNodePermissioned", "permType", permissionType, "url", node.String())
	switch permissionType {
	case types.Default:
		allowed := types.IsNodePermissioned(nodename, currentNode, datadir, direction)
		log.Info("AJ-isNodePermissioned Default", "allowed", allowed, "url", node.String())
		return allowed

	case types.Basic:
		allowed := isNodePermissionedBasic(node.EnodeID())
		log.Info("AJ-isNodePermissioned Basic", "allowed", allowed, "url", node.String())
		return allowed
	case types.EEA:
		allowed, err := permissionService.ConnectionAllowed(node.EnodeID(), node.IP().String(), uint16(node.TCP()), uint16(node.RaftPort()))
		log.Info("AJ-isNodePermissioned EEA", "allowed", allowed, "url", node.String())
		if err != nil {
			log.Error("AJ-isNodePermissioned EEA ERRORED", "err", err, "allowed", allowed, "url", node.String())
			return false
		}
		return allowed
	}
	return false
}
