package raft

import (
	"math/big"
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
