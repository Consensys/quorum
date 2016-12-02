package gethRaft

import (
	"sort"

	"github.com/coreos/etcd/raft"

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

// ByContext : Machinery to sort raft peers by `Context`
type ByContext []raft.Peer

func (a ByContext) Len() int           { return len(a) }
func (a ByContext) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByContext) Less(i, j int) bool { return string(a[i].Context) < string(a[j].Context) }

// StartNode starts the raft node with its current peers as a cluster.
//
// TODO at some point we need to take care that the node starts only when it
// knows exactly the start peers.
func (s *PublicRaftAPI) StartNode() {
	manager := s.service.raftProtocolManager

	// add our home node to the array of peers
	rpeers := make([]raft.Peer, len(manager.rlpxKnownPeers)+1)
	rpeers[0] = raft.Peer{
		ID:      strToIntID(manager.id),
		Context: []byte(manager.id),
	}

	// and add all the peers we know about
	var i = 1
	for enodeID := range manager.rlpxKnownPeers {
		intID := strToIntID(enodeID)
		peer := raft.Peer{
			ID:      intID,
			Context: []byte(enodeID),
		}
		rpeers[i] = peer
		i++
	}

	// This step is important -- we need the raft log to start out the same across
	// the different nodes -- we already take care that each entry is the same, and
	// this step ensures that they're in the same order.
	sort.Sort(ByContext(rpeers))

	// TODO(joel) how to handle if rpeers differs from the WAL peers?
	StartRaftNode(manager, manager.raftStorage, rpeers)

	go s.service.notifyRoleChange(manager.rawNode.RoleChan().Out())
}

// TODO
// func (s *PublicRaftAPI) StopNode() {
// }

// Version returns the Raft version this node offers.
func (s *PublicRaftAPI) Version() (*rpc.HexNumber, error) {
	return rpc.NewHexNumber(s.version), nil
}
