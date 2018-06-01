package raft

import (
	"testing"
	"strconv"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rlp"
	"fmt"
)

func TestSignHeader(t *testing.T){
	//create only what we need to test the seal
	var testRaftId uint16 = 5
	config := &node.Config{Name: "unit-test", DataDir: ""}

	nodeKey := config.NodeKey()

	raftProtocolManager := &ProtocolManager{raftId:testRaftId}
	raftService := &RaftService{nodeKey: nodeKey, raftProtocolManager: raftProtocolManager}
	minter := minter{eth: raftService,}

	//create some fake header to sign
	fakeParentHash := common.HexToHash("0xc2c1dc1be8054808c69e06137429899d")

	header := &types.Header{
		ParentHash: fakeParentHash,
		Number:     big.NewInt(1),
		Difficulty: big.NewInt(1),
		GasLimit:   new(big.Int),
		GasUsed:    new(big.Int),
		Coinbase:   minter.coinbase,
		Time:       big.NewInt(time.Now().UnixNano()),
	}

	headerHash := header.Hash()
	extraDataBytes := minter.buildExtraSeal(headerHash)
	print(extraDataBytes)
	seal := new(extraSeal)
	err := rlp.DecodeBytes(extraDataBytes, seal)
	if err != nil {
		t.Fatalf("Unable to decode seal: %s", err.Error())
	}
	print("\n")
	fmt.Printf("%v\n", seal)
	print("\n")
	print(seal.raftId)
	print("\n")
	print(seal.signature)
	print("\n")

	// Check raftId
	sealRaftId, err := strconv.ParseInt(string(seal.raftId), 16, 64)
	if err != nil {
		t.Errorf("Unable to get RaftId")
	}
	if  sealRaftId != int64	(testRaftId) {
		t.Errorf("RaftID does not match. Expected: %d, Actual: %d", testRaftId, sealRaftId)
	}

	//Identify who signed it
	sig:= seal.signature
	pubKey, err := crypto.SigToPub(headerHash.Bytes(), sig)
	if err != nil {
		t.Fatalf("Unable to get public key from signature: %s", err.Error())
	}

	//Compare derived public key to original public key
	if pubKey.X.Cmp(nodeKey.X) != 0 {
		t.Errorf("Signature incorrect!")
	}

}
