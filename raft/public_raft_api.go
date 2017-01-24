package gethRaft

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type PublicRaftAPI struct {
	version uint64
	service *RaftService
}

func NewPublicRaftAPI(service *RaftService) *PublicRaftAPI {
	return &PublicRaftAPI{
		version: protocolVersion,
		service: service,
	}
}

// The Raft version this node offers.
func (s *PublicRaftAPI) Version() (*rpc.HexNumber, error) {
	return rpc.NewHexNumber(s.version), nil
}
