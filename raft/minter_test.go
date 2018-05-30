package raft

import (
	"testing"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/node"
	"math/big"
	"time"
)

func TestSignHeader(t *testing.T){
	//create only what we need to test the signature
	config := &node.Config{Name: "unit-test", DataDir: ""}

	nodeKey := config.NodeKey()
	raftService := &RaftService{nodeKey: nodeKey,}
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
	sig := minter.signHeader(headerHash)

	//Identify who signed it
	pubKey, err := crypto.SigToPub(headerHash.Bytes(), sig)
	if err != nil {
		t.Errorf("Unable to get public key from signature: %s", err.Error())
	}

	//Compare derived public key to original public key
	if pubKey.X.Cmp(nodeKey.X) != 0 {
		t.Errorf("Signature incorrect!")
	}

}
