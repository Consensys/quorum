package raft

type PublicRaftAPI struct {
	raftService *RaftService
}

func NewPublicRaftAPI(raftService *RaftService) *PublicRaftAPI {
	return &PublicRaftAPI{raftService}
}

func (s *PublicRaftAPI) Role() string {
	role := s.raftService.raftProtocolManager.role
	if role == minterRole {
		return "minter"
	} else {
		return "verifier"
	}
}
