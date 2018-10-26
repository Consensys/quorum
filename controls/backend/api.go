package backend

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/log"
	)

type PermissionAPI struct {
}

func NewPermissionAPI() *PermissionAPI {
	return &PermissionAPI{}
}

func APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: "permnode",
			Version:   "1.0",
			Service:   NewPermissionAPI(),
			Public:    true,
		},
	}
}

func (s *PermissionAPI) AddVoter(addr string) string {
	log.Info("AJ-called1")
	return "added voter " + addr
}

func (s *PermissionAPI) ProposeNode(enodeId string) string {
	log.Info("AJ-called2")
	return "proposed node " + enodeId
}

func (s *PermissionAPI) BlacklistNode(enodeId string) string {
	log.Info("AJ-called3")
	return "blacklisted node " + enodeId
}

func (s *PermissionAPI) RemoveNode(enodeId string) string {
	log.Info("AJ-called4")
	return "removed node " + enodeId
}

func (s *PermissionAPI) ApproveNode(enodeId string) string {
	log.Info("AJ-called5")
	return "approved node " + enodeId
}

func (s *PermissionAPI) ValidNodes() []string {
	log.Info("AJ-called6")
	return []string{"n1", "n2"}
}
