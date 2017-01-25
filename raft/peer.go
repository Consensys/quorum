package gethRaft

import (
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
	"github.com/ethereum/go-ethereum/p2p"

	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
)

type PeerInfo struct {
	// Role raft.StateType `json:"role"`
	ID uint64
}

var (
	// Connected via ethereum, but not yet connected via raft:
	halfConnectedPeer = &PeerInfo{
		ID: 0,
	}
)

type peer struct {
	// internal (evm) id
	strID string
	// external (etcd) id
	uint64Id uint64

	rawPeer  *p2p.Peer
	raftPeer *raft.Peer

	rw p2p.MsgReadWriter
}

func newPeer(remote *p2p.Peer, rw p2p.MsgReadWriter) *peer {
	strID := remote.ID().String()
	id := strToIntID(strID)
	return &peer{
		strID:    strID,
		uint64Id: id,
		rawPeer:  remote,
		// XXX(joel) fill this in when the raft peer is connected (or do we actually
		// need this? I don't think it's used)
		raftPeer: nil,
		rw:       rw,
	}
}

func (p *peer) Info() *PeerInfo {
	if nil == p.raftPeer {
		return halfConnectedPeer
	}

	return &PeerInfo{
		ID: p.raftPeer.ID,
	}
}

func (p *peer) SendRaftPB(m raftpb.Message) {
	data, err := m.Marshal()

	if err != nil {
		panic(err.Error())
	}

	if err := p2p.Send(p.rw, raftMsg, data); err != nil {
		glog.V(logger.Error).Infof(
			"Failed to send message (%v) to peer (%v): %v", m.Type, m.To, err)
	}
}
