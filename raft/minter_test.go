package raft

import (
	"testing"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jpmorganchase/quorum/raft/backend"
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
		Difficulty: big.NewInt(12),
		GasLimit:   0,
		GasUsed:    0,
		Coinbase:   minter.coinbase,
		Time:       big.NewInt(time.Now().UnixNano()),
	}

	sig := minter.signHeader(header.Hash())

	//Identify who signed it
	pubKey, err := crypto.SigToPub(header.Hash().Bytes(), sig)
	if err != nil {
		t.Errorf("Unable to get public key from signature: %s", err.Error())
	}

	//Compare derived publick key to original public key
	if pubKey != nodeKey.Public() {
		t.Errorf("Signature incorrect!")
	}

}

