package raft

import (
	"io"
	"net"

	"fmt"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
)

// Serializable information about a Peer. Sufficient to build `etcdRaft.Peer`
// or `enode.Node`.
// As NodeId is mainly used to derive the `ecdsa.pubkey` to build `enode.Node` it is kept as [64]byte instead of ID [32]byte used by `enode.Node`.
type Address struct {
	RaftId   uint16        `json:"raftId"`
	NodeId   enode.EnodeID `json:"nodeId"`
	Ip       net.IP        `json:"ip"`
	P2pPort  enr.TCP       `json:"p2pPort"`
	RaftPort enr.RaftPort  `json:"raftPort"`
	IsLearner bool			`json:"isLearner"`
}

func newAddress(raftId uint16, raftPort int, node *enode.Node, isLearner bool) *Address {
	// derive 64 byte nodeID from 128 byte enodeID
	id, err := enode.RaftHexID(node.EnodeID())
	if err != nil {
		panic(err)
	}
	return &Address{
		RaftId:   raftId,
		NodeId:   id,
		Ip:       node.IP(),
		P2pPort:  enr.TCP(node.TCP()),
		RaftPort: enr.RaftPort(raftPort),
		IsLearner: isLearner,
	}
}

// A peer that we're connected to via both raft's http transport, and ethereum p2p
type Peer struct {
	address *Address    // For raft transport
	p2pNode *enode.Node // For ethereum transport
}

func (addr *Address) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{addr.RaftId, addr.NodeId, addr.Ip, addr.P2pPort, addr.RaftPort, addr.IsLearner})
}

func (addr *Address) DecodeRLP(s *rlp.Stream) error {
	// These fields need to be public:
	var temp struct {
		RaftId   uint16
		NodeId   enode.EnodeID
		Ip       net.IP
		P2pPort  enr.TCP
		RaftPort enr.RaftPort
		IsLearner bool
	}

	var tempOld struct {
		RaftId   uint16
		NodeId   enode.EnodeID
		Ip       net.IP
		P2pPort  enr.TCP
		RaftPort enr.RaftPort
	}

	//TODO (Amal): review
	if err := s.Decode(&temp); err != nil {
		log.Printf("AJ - error decoding %v", err)
		if err1 := s.Decode(&tempOld); err1 != nil {
			return err1
		}
		addr.RaftId, addr.NodeId, addr.Ip, addr.P2pPort, addr.RaftPort, addr.IsLearner = tempOld.RaftId, tempOld.NodeId, tempOld.Ip, tempOld.P2pPort, tempOld.RaftPort, false
		return nil
		//return err
	} else {
		addr.RaftId, addr.NodeId, addr.Ip, addr.P2pPort, addr.RaftPort, addr.IsLearner = temp.RaftId, temp.NodeId, temp.Ip, temp.P2pPort, temp.RaftPort, temp.IsLearner
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
