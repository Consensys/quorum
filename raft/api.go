package raft

type PublicRaftAPI struct {
	raftService *RaftService
}

func NewPublicRaftAPI(raftService *RaftService) *PublicRaftAPI {
	return &PublicRaftAPI{raftService}
}

func (s *PublicRaftAPI) Role() string {
	return s.raftService.raftProtocolManager.NodeInfo().Role
}

func (s *PublicRaftAPI) AddPeer(raftId uint16, enodeId string) error {
	return s.raftService.raftProtocolManager.ProposeNewPeer(raftId, enodeId)
}

func (s *PublicRaftAPI) RemovePeer(raftId uint16) {
	s.raftService.raftProtocolManager.ProposePeerRemoval(raftId)
}
