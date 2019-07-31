package raft

import (
	"errors"
	"fmt"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rlp"
)

func TestSignHeader(t *testing.T) {
	//create only what we need to test the seal
	var testRaftId uint16 = 5
	config := &node.Config{Name: "unit-test", DataDir: ""}

	nodeKey := config.NodeKey()

	raftProtocolManager := &ProtocolManager{raftId: testRaftId}
	raftService := &RaftService{nodeKey: nodeKey, raftProtocolManager: raftProtocolManager}
	minter := minter{eth: raftService}

	//create some fake header to sign
	fakeParentHash := common.HexToHash("0xc2c1dc1be8054808c69e06137429899d")

	header := &types.Header{
		ParentHash: fakeParentHash,
		Number:     big.NewInt(1),
		Difficulty: big.NewInt(1),
		GasLimit:   uint64(0),
		GasUsed:    uint64(0),
		Coinbase:   minter.coinbase,
		Time:       big.NewInt(time.Now().UnixNano()),
	}

	headerHash := header.Hash()
	extraDataBytes := minter.buildExtraSeal(headerHash)
	var seal *extraSeal
	err := rlp.DecodeBytes(extraDataBytes[:], &seal)
	if err != nil {
		t.Fatalf("Unable to decode seal: %s", err.Error())
	}

	// Check raftId
	sealRaftId, err := hexutil.DecodeUint64("0x" + string(seal.RaftId)) //add the 0x prefix
	if err != nil {
		t.Errorf("Unable to get RaftId: %s", err.Error())
	}
	if sealRaftId != uint64(testRaftId) {
		t.Errorf("RaftID does not match. Expected: %d, Actual: %d", testRaftId, sealRaftId)
	}

	//Identify who signed it
	sig := seal.Signature
	pubKey, err := crypto.SigToPub(headerHash.Bytes(), sig)
	if err != nil {
		t.Fatalf("Unable to get public key from signature: %s", err.Error())
	}

	//Compare derived public key to original public key
	if pubKey.X.Cmp(nodeKey.X) != 0 {
		t.Errorf("Signature incorrect!")
	}

}

func TestAddLearner_whenTypical(t *testing.T) {
	//create only what we need to test add learner node
	var testRaftId uint16 = 1
	config := &node.Config{Name: "unit-test", DataDir: ""}

	nodeKey := config.NodeKey()
	enodeIdStr := fmt.Sprintf("%x", crypto.FromECDSAPub(&nodeKey.PublicKey)[1:])
	url := enodeId(enodeIdStr, "127.0.0.1:21001", 50401)
	err, peers := peerList(url)
	if err != nil {
		t.Errorf("getting peers failed %v", err)
	}

	raftProtocolManager := &ProtocolManager{
		raftId:              testRaftId,
		bootstrapNodes:      peers,
		confChangeProposalC: make(chan raftpb.ConfChange),
		removedPeers:        mapset.NewSet(),
		confState:           raftpb.ConfState{Nodes: []uint64{uint64(testRaftId)}, Learners: []uint64{}},
	}

	raftService := &RaftService{nodeKey: nodeKey, raftProtocolManager: raftProtocolManager}

	propPeer := func() {
		raftid, err := raftService.raftProtocolManager.ProposeNewPeer("enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", true)
		if err != nil {
			t.Errorf("propose new peer failed %v\n", err)
		}
		if raftid != testRaftId+1 {
			t.Errorf("1. wrong raft id. expected %d got %d\n", testRaftId+1, raftid)
		}
	}
	go propPeer()
	select {
	case confChange := <-raftService.raftProtocolManager.confChangeProposalC:
		if confChange.Type != raftpb.ConfChangeAddLearnerNode {
			t.Errorf("expected ConfChangeAddLearnerNode but got %s", confChange.Type.String())
		}
		if uint16(confChange.NodeID) != testRaftId+1 {
			t.Errorf("2. wrong raft id. expected %d got %d\n", testRaftId+1, uint16(confChange.NodeID))
		}
	case <-time.After(time.Millisecond * 200):
		t.Errorf("add learner conf change not recieved")
	}
}

func TestPromoteLearnerToPeer_whenTypical(t *testing.T) {
	//create only what we need to test add learner node
	var testRaftId uint16 = 2
	config := &node.Config{Name: "unit-test", DataDir: ""}

	nodeKey := config.NodeKey()
	enodeIdStr := fmt.Sprintf("%x", crypto.FromECDSAPub(&nodeKey.PublicKey)[1:])
	enodeStr := enodeId(enodeIdStr, "127.0.0.1:21001", 50401)
	err, peers := peerList(enodeStr)
	if err != nil {
		t.Errorf("getting peers failed %v", err)
	}
	learnerRaftId := uint16(1)

	raftProtocolManager := &ProtocolManager{
		raftId:              testRaftId,
		bootstrapNodes:      peers,
		confChangeProposalC: make(chan raftpb.ConfChange),
		removedPeers:        mapset.NewSet(),
		confState:           raftpb.ConfState{Nodes: []uint64{uint64(testRaftId)}, Learners: []uint64{uint64(learnerRaftId)}},
	}
	raftService := &RaftService{nodeKey: nodeKey, raftProtocolManager: raftProtocolManager}
	promoteToPeer := func() {
		ok, err := raftService.raftProtocolManager.PromoteToPeer(learnerRaftId)
		if err != nil || !ok {
			t.Errorf("promote learner to peer failed %v\n", err)
		}
	}
	go promoteToPeer()
	select {
	case confChange := <-raftService.raftProtocolManager.confChangeProposalC:
		if confChange.Type != raftpb.ConfChangeAddNode {
			t.Errorf("expected ConfChangeAddNode but got %s", confChange.Type.String())
		}
		if uint16(confChange.NodeID) != learnerRaftId {
			t.Errorf("2. wrong raft id. expected %d got %d\n", learnerRaftId, uint16(confChange.NodeID))
		}
	case <-time.After(time.Millisecond * 200):
		t.Errorf("add learner conf change not recieved")
	}
}

func TestAddLearnerOrPeer_fromLearner(t *testing.T) {
	//create only what we need to test add learner node
	var testRaftId uint16 = 1
	config := &node.Config{Name: "unit-test", DataDir: ""}

	nodeKey := config.NodeKey()
	enodeIdStr := fmt.Sprintf("%x", crypto.FromECDSAPub(&nodeKey.PublicKey)[1:])
	enodeStr := enodeId(enodeIdStr, "127.0.0.1:21001", 50401)
	err, peers := peerList(enodeStr)
	if err != nil {
		t.Errorf("getting peers failed %v", err)
	}

	raftProtocolManager := &ProtocolManager{
		raftId:              testRaftId,
		bootstrapNodes:      peers,
		confChangeProposalC: make(chan raftpb.ConfChange),
		removedPeers:        mapset.NewSet(),
		confState:           raftpb.ConfState{Nodes: []uint64{}, Learners: []uint64{uint64(testRaftId)}},
	}

	raftService := &RaftService{nodeKey: nodeKey, raftProtocolManager: raftProtocolManager}

	_, err = raftService.raftProtocolManager.ProposeNewPeer("enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", true)

	if err == nil {
		t.Errorf("learner should not be allowed to add learner or peer")
	}

	if err != nil && strings.Index(err.Error(), "learner node can't add peer or learner") == -1 {
		t.Errorf("expect error message: propose new peer failed, got: %v\n", err)
	}

	_, err = raftService.raftProtocolManager.ProposeNewPeer("enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", false)

	if err == nil {
		t.Errorf("learner should not be allowed to add learner or peer")
	}

	if err != nil && strings.Index(err.Error(), "learner node can't add peer or learner") == -1 {
		t.Errorf("expect error message: propose new peer failed, got: %v\n", err)
	}

}

func TestPromoteLearnerToPeer_fromLearner(t *testing.T) {
	//create only what we need to test add learner node
	var testRaftId uint16 = 1
	config := &node.Config{Name: "unit-test", DataDir: ""}

	nodeKey := config.NodeKey()
	enodeIdStr := fmt.Sprintf("%x", crypto.FromECDSAPub(&nodeKey.PublicKey)[1:])
	enodeStr := enodeId(enodeIdStr, "127.0.0.1:21001", 50401)
	err, peers := peerList(enodeStr)
	if err != nil {
		t.Errorf("getting peers failed %v", err)
	}
	learnerRaftId := uint16(2)
	raftProtocolManager := &ProtocolManager{
		raftId:              testRaftId,
		bootstrapNodes:      peers,
		confChangeProposalC: make(chan raftpb.ConfChange),
		removedPeers:        mapset.NewSet(),
		confState:           raftpb.ConfState{Nodes: []uint64{}, Learners: []uint64{uint64(testRaftId), uint64(learnerRaftId)}},
	}

	raftService := &RaftService{nodeKey: nodeKey, raftProtocolManager: raftProtocolManager}

	_, err = raftService.raftProtocolManager.PromoteToPeer(learnerRaftId)

	if err == nil {
		t.Errorf("learner should not be allowed to promote to peer")
	}

	if err != nil && strings.Index(err.Error(), "learner node can't promote to peer") == -1 {
		t.Errorf("expect error message: propose new peer failed, got: %v\n", err)
	}

}

func enodeId(id string, ip string, raftPort int) string {
	return fmt.Sprintf("enode://%s@%s?discport=0&raftport=%d", id, ip, raftPort)
}

func peerList(url string) (error, []*enode.Node) {
	var nodes []*enode.Node
	node, err := enode.ParseV4(url)
	if err != nil {
		return errors.New(fmt.Sprintf("Node URL %s: %v\n", url, err)), nil
	}
	nodes = append(nodes, node)
	return nil, nodes
}
