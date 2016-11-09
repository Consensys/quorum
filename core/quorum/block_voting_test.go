package quorum_test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/core/quorum"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

var (
	voteKey1, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addrVoteKey1 = crypto.PubkeyToAddress(voteKey1.PublicKey)

	voteKey2, _  = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	addrVoteKey2 = crypto.PubkeyToAddress(voteKey2.PublicKey)

	blockMakerKey1, _ = crypto.HexToECDSA("49a7b37aa6f6645917e7b807e9d1c00d4fa71f18343b0d4122a4d2df64dd6fee")
	addrBlockMaker1   = crypto.PubkeyToAddress(blockMakerKey1.PublicKey)
)

func genesisBlock(voteThreshold int) string {
	return fmt.Sprintf(`{
		   "coinbase": "0x0000000000000000000000000000000000000000",
		   "config": {
		      "homesteadBlock": 0
		   },
		   "difficulty": "0x0",
		   "extraData": "0x",
		   "gasLimit": "0x2FEFD800",
		   "mixhash": "0x00000000000000000000000000000000000000647572616c65787365646c6578",
		   "nonce": "0x0",
		   "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
		   "timestamp": "0x00",
		   "alloc": {
			"%s": {
			"balance":"100000000000000000"
			},
			"%s": {
			"balance":"100000000000000000"
			},
			"%s": {
			"balance":"100000000000000000"
			},
		      "%s": {
			 "code": "%s",
			 "storage": {
			    "0x0000000000000000000000000000000000000000000000000000000000000001": "%#x",

			    "0x0000000000000000000000000000000000000000000000000000000000000002": "0x02",
			    "0x9ba0793ab2c61c9fc0f4204891503530bba7c28e9bf7671b90e9f5010eb132c5": "0x01",
			    "0x706439b5bc31f518c47b3ca07b45a1b4b7ff29deae44d5f3450e5c76b207c890": "0x01",

			    "0x0000000000000000000000000000000000000000000000000000000000000004": "0x01",
			    "0xc02baebfa4f2f2bec2e73904266d093f2a0b2c1b54617f55347d7c4e6ef48047": "0x01"
			 }
		      }
		   }
		}`,
		addrVoteKey1.Hex(),
		addrVoteKey2.Hex(),
		addrBlockMaker1.Hex(),
		params.QuorumVotingContractAddr.Hex(),
		quorum.RuntimeCode,
		voteThreshold,
	)
}

func TestBlockVotingBlockMaker(t *testing.T) {
	//var (
	//	db, _       = ethdb.NewMemDatabase()
	//	mux         event.TypeMux
	//	chainConfig = core.MakeChainConfig()
	//	bv          *quorum.BlockVoting
	//
	//	err error
	//)
	//
	//genesis, err := core.WriteGenesisBlock(db, strings.NewReader(genesisBlock(2)))
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//blockchain, _ := core.NewBlockChain(db, chainConfig, core.FakePow{}, &mux, false)
	//
	//txpool := core.NewTxPool(chainConfig, &mux, func() (*state.StateDB, error) {
	//	_, state := bv.Pending()
	//	return state, nil
	//}, func() *big.Int { return big.NewInt(1000000) })
	//
	//// func NewBlockVoting(client *rpc.Client, bc *core.BlockChain, chainConfig *core.ChainConfig, txpool *core.TxPool, mux *event.TypeMux, db ethdb.Database, blockMakerKey, voteKey *ecdsa.PrivateKey, isSynchronised bool) *BlockVoting {
	//bv, err = quorum.NewBlockVoting(nil, blockchain, chainConfig, txpool, &mux, db, blockMakerKey1, voteKey1, true)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//pBlock, _ := bv.Pending()
	//if pBlock.NumberU64() != 1 {
	//	t.Fatalf("Expected pending block to have block number 1, got %d", pBlock.NumberU64())
	//}
	//
	//createBlockReq := quorum.CreateBlock{
	//	Hash: make(chan common.Hash),
	//	Err:  make(chan error),
	//}
	//
	//mux.Post(createBlockReq)
	//
	//select {
	//case <-createBlockReq.Hash:
	//	t.Fatalf("Block was create while vote threshold was not met")
	//case <-createBlockReq.Err:
	//	break
	//}
	//
	//// send in 2 votes and creating a block should work
	//vote := quorum.Vote{
	//	Hash:   genesis.Hash(),
	//	Number: genesis.Number(),
	//	TxHash: make(chan common.Hash),
	//	Err:    make(chan error),
	//}
	//mux.Post(vote)
	//
	//select {
	//case err := <-vote.Err:
	//	t.Fatalf("Unable to vote: %s", err)
	//case <-vote.TxHash:
	//	break
	//}
	//
	//mux.Post(vote)
	//
	//select {
	//case err := <-vote.Err:
	//	t.Fatalf("Unable to vote: %s", err)
	//case <-vote.TxHash:
	//	break
	//}
	//
	//mux.Post(createBlockReq)
	//
	//select {
	//case hash := <-createBlockReq.Hash:
	//	t.Logf("Created block %s", hash.Hex())
	//case err := <-createBlockReq.Err:
	//	t.Fatalf("Block should be created but got error: %s", err)
	//}
}
