package raft

type RaftNodeInfo struct {
	ClusterSize    int        `json:"clusterSize"`
	Role           string     `json:"role"`
	Address        *Address   `json:"address"`
	PeerAddresses  []*Address `json:"peerAddresses"`
	RemovedPeerIds []uint16   `json:"removedPeerIds"`
	AppliedIndex   uint64     `json:"appliedIndex"`
	SnapshotIndex  uint64     `json:"snapshotIndex"`
}

type PublicRaftAPI struct {
	raftService *RaftService
}

func NewPublicRaftAPI(raftService *RaftService) *PublicRaftAPI {
	return &PublicRaftAPI{raftService}
}

func (s *PublicRaftAPI) Role() string {
	return s.raftService.raftProtocolManager.NodeInfo().Role
}

func (s *PublicRaftAPI) AddPeer(enodeId string) (uint16, error) {
	return s.raftService.raftProtocolManager.ProposeNewPeer(enodeId, false)
}

func (s *PublicRaftAPI) AddLearner(enodeId string) (uint16, error) {
	return s.raftService.raftProtocolManager.ProposeNewPeer(enodeId, true)
}

func (s *PublicRaftAPI) PromoteToPeer(raftId uint16) (bool, error) {
	return s.raftService.raftProtocolManager.PromoteToPeer(raftId)
}

func (s *PublicRaftAPI) RemovePeer(raftId uint16) error {
	return s.raftService.raftProtocolManager.ProposePeerRemoval(raftId)
}

func (s *PublicRaftAPI) Leader() (string, error) {

	addr, err := s.raftService.raftProtocolManager.LeaderAddress()
	if err != nil {
		return "", err
	}
	return addr.NodeId.String(), nil
}

func (s *PublicRaftAPI) Cluster() ([]ClusterInfo, error) {
	nodeInfo := s.raftService.raftProtocolManager.NodeInfo()
	if nodeInfo.Role == "" {
		return []ClusterInfo{}, nil
	}
	leaderAddr, err := s.raftService.raftProtocolManager.LeaderAddress()
	if err != nil {
		if err == errNoLeaderElected && s.Role() == "" {
			return []ClusterInfo{}, nil
		}
		return []ClusterInfo{}, err
	}
	peerAddresses := append(nodeInfo.PeerAddresses, nodeInfo.Address)
	clustInfo := make([]ClusterInfo, len(peerAddresses))
	for i, a := range peerAddresses {
		role := ""
		if a.RaftId == leaderAddr.RaftId {
			role = "minter"
		} else if s.raftService.raftProtocolManager.isLearner(a.RaftId) {
			role = "learner"
		} else if s.raftService.raftProtocolManager.isVerifier(a.RaftId) {
			role = "verifier"
		}
		clustInfo[i] = ClusterInfo{*a, role}
	}
	return clustInfo, nil
}

func (s *PublicRaftAPI) GetRaftId(enodeId string) (uint16, error) {
	return s.raftService.raftProtocolManager.FetchRaftId(enodeId)
}
