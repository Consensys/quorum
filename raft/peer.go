package raft

import (
	"io"
	"net"

	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/rlp"
)

// Serializable information about a Peer. Sufficient to build `etcdRaft.Peer`
// or `discover.Node`.
type Address struct {
	raftId   uint16          `json:"raftId"`
	nodeId   discover.NodeID `json:"nodeId"`
	ip       net.IP          `json:"ip"`
	p2pPort  uint16          `json:"p2pPort"`
	raftPort uint16          `json:"raftPort"`
}

func newAddress(raftId uint16, raftPort uint16, node *discover.Node) *Address {
	return &Address{
		raftId:   raftId,
		nodeId:   node.ID,
		ip:       node.IP,
		p2pPort:  node.TCP,
		raftPort: raftPort,
	}
}

type RaftPeerInfo struct {
	raftId   string `json:"raftId"`
	nodeId   string `json:"nodeId"`
	ip       string `json:"ip"`
	p2pPort  string `json:"p2pPort"`
	raftPort string `json:"raftPort"`
}

// A peer that we're connected to via both raft's http transport, and ethereum p2p
type Peer struct {
	address *Address       // For raft transport
	p2pNode *discover.Node // For ethereum transport
}

func (addr *Address) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{addr.raftId, addr.nodeId, addr.ip, addr.p2pPort, addr.raftPort})
}

func (addr *Address) DecodeRLP(s *rlp.Stream) error {
	// These fields need to be public:
	var temp struct {
		RaftId   uint16
		NodeId   discover.NodeID
		Ip       net.IP
		P2pPort  uint16
		RaftPort uint16
	}

	if err := s.Decode(&temp); err != nil {
		return err
	} else {
		addr.raftId, addr.nodeId, addr.ip, addr.p2pPort, addr.raftPort = temp.RaftId, temp.NodeId, temp.Ip, temp.P2pPort, temp.RaftPort
		return nil
	}
}

// RLP Address encoding, for transport over raft and storage in LevelDB.

func (addr *Address) toBytes() []byte {
	size, r, err := rlp.EncodeToReader(addr)
	if err != nil {
		panic(fmt.Sprintf("error: failed to RLP-encode Address: %s", err.Error()))
	}
	var buffer = make([]byte, uint32(size))
	r.Read(buffer)

	return buffer
}

func bytesToAddress(bytes []byte) *Address {
	var addr Address
	if err := rlp.DecodeBytes(bytes, &addr); err != nil {
		log.Fatalf("failed to RLP-decode Address: %v", err)
	}
	return &addr
}

func (address *Address) toDisplay() *RaftPeerInfo {
	info := &RaftPeerInfo{
		raftId:   fmt.Sprintf("%v", address.raftId),
		nodeId:   address.nodeId.String(),
		ip:       address.ip.String(),
		p2pPort:  fmt.Sprintf("%v", address.p2pPort),
		raftPort: fmt.Sprintf("%v", address.raftPort),
	}
	return info
}
