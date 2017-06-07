package quorum

import (
	"crypto/ecdsa"
	"math/big"
	"sync"

	"gopkg.in/fatih/set.v0"

	"fmt"

	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	// Create bindings with: go run cmd/abigen/main.go -abi <definition> -pkg quorum -type VotingContract > core/quorum/binding.go
	ABI = `[{"constant":false,"inputs":[{"name":"threshold","type":"uint256"}],"name":"setVoteThreshold","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"removeBlockMaker","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"voterCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"canCreateBlocks","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"voteThreshold","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"height","type":"uint256"}],"name":"getCanonHash","outputs":[{"name":"","type":"bytes32"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"height","type":"uint256"},{"name":"hash","type":"bytes32"}],"name":"vote","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"addBlockMaker","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"removeVoter","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"height","type":"uint256"},{"name":"n","type":"uint256"}],"name":"getEntry","outputs":[{"name":"","type":"bytes32"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"addr","type":"address"}],"name":"isVoter","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"canVote","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"blockMakerCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"getSize","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"addr","type":"address"}],"name":"isBlockMaker","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"addVoter","outputs":[],"payable":false,"type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"sender","type":"address"},{"indexed":false,"name":"blockNumber","type":"uint256"},{"indexed":false,"name":"blockHash","type":"bytes32"}],"name":"Vote","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"","type":"address"}],"name":"AddVoter","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"","type":"address"}],"name":"RemovedVoter","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"","type":"address"}],"name":"AddBlockMaker","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"","type":"address"}],"name":"RemovedBlockMaker","type":"event"}]`

	// browser solidity with optimizations: 0.4.9+commit.364da425.Emscripten.clang (with optimizations enabled)
	RuntimeCode = "606060405236156100ca5763ffffffff60e060020a6000350416631290948581146100cc578063284d163c146100e157806342169e48146100ff578063488099a6146101215780634fe437d514610151578063559c390c1461017357806368bb8bb61461019857806372a571fc146101b057806386c1ff68146101ce57806398ba676d146101ec578063a7771ee314610214578063adfaa72e14610244578063cf52898514610274578063de8fa43114610296578063e814d1c7146102b8578063f4ab9adf146102e8575bfe5b34156100d457fe5b6100df600435610306565b005b34156100e957fe5b6100df600160a060020a036004351661033b565b005b341561010757fe5b61010f6103fc565b60408051918252519081900360200190f35b341561012957fe5b61013d600160a060020a0360043516610402565b604080519115158252519081900360200190f35b341561015957fe5b61010f610417565b60408051918252519081900360200190f35b341561017b57fe5b61010f60043561041d565b60408051918252519081900360200190f35b34156101a057fe5b6100df600435602435610523565b005b34156101b857fe5b6100df600160a060020a0360043516610632565b005b34156101d657fe5b6100df600160a060020a03600435166106e9565b005b34156101f457fe5b61010f6004356024356107aa565b60408051918252519081900360200190f35b341561021c57fe5b61013d600160a060020a03600435166107ff565b604080519115158252519081900360200190f35b341561024c57fe5b61013d600160a060020a0360043516610821565b604080519115158252519081900360200190f35b341561027c57fe5b61010f610836565b60408051918252519081900360200190f35b341561029e57fe5b61010f61083c565b60408051918252519081900360200190f35b34156102c057fe5b61013d600160a060020a0360043516610843565b604080519115158252519081900360200190f35b34156102f057fe5b6100df600160a060020a0360043516610865565b005b600160a060020a03331660009081526003602052604090205460ff16156103325760018190555b610337565b610000565b5b50565b600160a060020a03331660009081526005602052604090205460ff1615610332576004546001141561036c57610000565b600160a060020a03811660009081526005602052604090205460ff161561032d57600160a060020a038116600081815260056020908152604091829020805460ff1916905560048054600019019055815192835290517f8cee3054364d6799f1c8962580ad61273d9d38ca1ff26516bd1ad23c099a60229281900390910190a15b5b610337565b610000565b5b50565b60025481565b60056020526000908152604090205460ff1681565b60015481565b600060006000600060006001860381548110151561043757fe5b906000526020600020906002020160005b509250600090505b60018301548110156105175760018301805484916000918490811061047157fe5b906000526020600020900160005b5054815260208082019290925260409081016000908120548582529286905220541080156104e6575060015483600001600085600101848154811015156104c257fe5b906000526020600020900160005b5054815260208101919091526040016000205410155b1561050e57600183018054829081106104fb57fe5b906000526020600020900160005b505491505b5b600101610450565b8193505b505050919050565b600160a060020a03331660009081526003602052604081205460ff1615610332576000548390101561056457600080548085030190610562908261091c565b505b60008054600019850190811061057657fe5b906000526020600020906002020160005b5060008381526020829052604090205490915015156105ca578060010180548060010182816105b6919061094e565b916000526020600020900160005b50839055505b600082815260208281526040918290208054600101905581518581529081018490528151600160a060020a033316927f3d03ba7f4b5227cdb385f2610906e5bcee147171603ec40005b30915ad20e258928290030190a25b61062c565b610000565b5b505050565b600160a060020a03331660009081526005602052604090205460ff161561033257600160a060020a03811660009081526005602052604090205460ff16151561032d57600160a060020a038116600081815260056020908152604091829020805460ff19166001908117909155600480549091019055815192835290517f1a4ce6942f7aa91856332e618fc90159f13a340611a308f5d7327ba0707e56859281900390910190a15b5b610337565b610000565b5b50565b600160a060020a03331660009081526003602052604090205460ff1615610332576002546001141561071a57610000565b600160a060020a03811660009081526003602052604090205460ff161561032d57600160a060020a038116600081815260036020908152604091829020805460ff1916905560028054600019019055815192835290517f183393fc5cffbfc7d03d623966b85f76b9430f42d3aada2ac3f3deabc78899e89281900390910190a15b5b610337565b610000565b5b50565b600060006000600185038154811015156107c057fe5b906000526020600020906002020160005b50905080600101838154811015156107e557fe5b906000526020600020900160005b505491505b5092915050565b600160a060020a03811660009081526003602052604090205460ff165b919050565b60036020526000908152604090205460ff1681565b60045481565b6000545b90565b600160a060020a03811660009081526005602052604090205460ff165b919050565b600160a060020a03331660009081526003602052604090205460ff161561033257600160a060020a03811660009081526003602052604090205460ff16151561032d57600160a060020a038116600081815260036020908152604091829020805460ff19166001908117909155600280549091019055815192835290517f0ad2eca75347acd5160276fe4b5dad46987e4ff4af9e574195e3e9bc15d7e0ff9281900390910190a15b5b610337565b610000565b5b50565b81548183558181151161062c5760020281600202836000526020600020918201910161062c9190610978565b5b505050565b81548183558181151161062c5760008381526020902061062c9181019083016109a5565b5b505050565b61084091905b8082111561099e57600061099560018301826109c6565b5060020161097e565b5090565b90565b61084091905b8082111561099e57600081556001016109ab565b5090565b90565b508054600082559060005260206000209081019061033791906109a5565b5b505600a165627a7a72305820208b21c6d5a94bcd416ed9e4443edb0a47c45dbc2ddbf68b39009ef8b7ad63ff0029"
)

// BlockVoting is a type of BlockMaker that uses a smart contract
// to determine the canonical chain. Parties that are allowed to
// vote send vote transactions to the voting contract. Based on
// these transactions the parent block is selected where the next
// block will be build on top of.
type BlockVoting struct {
	bc       *core.BlockChain
	cc       *core.ChainConfig
	txpool   *core.TxPool
	synced   bool
	mux      *event.TypeMux
	db       ethdb.Database
	am       *accounts.Manager
	gasPrice *big.Int

	voteSession  *VotingContractSession
	callContract *VotingContractCaller

	bmk      *ecdsa.PrivateKey
	vk       *ecdsa.PrivateKey
	coinbase common.Address

	pStateMu sync.Mutex
	pState   *pendingState
}

// Vote is posted to the event mux when the BlockVoting instance
// is ordered to send a new vote transaction. Hash is the hash for the
// given number depth.
type Vote struct {
	Hash   common.Hash
	Number *big.Int
	TxHash chan common.Hash
	Err    chan error
}

// CreateBlock is posted to the event mux when the BlockVoting instance
// is ordered to create a new block. Either the hash of the created
// block is returned is hash or an error.
type CreateBlock struct {
	Hash chan common.Hash
	Err  chan error
}

// NewBlockVoting creates a new BlockVoting instance.
// blockMakerKey and/or voteKey can be nil in case this node doesn't create blocks or vote.
// Note, don't forget to call Start.
func NewBlockVoting(bc *core.BlockChain, chainConfig *core.ChainConfig, txpool *core.TxPool, mux *event.TypeMux, db ethdb.Database, accountMgr *accounts.Manager, isSynchronised bool) *BlockVoting {
	bv := &BlockVoting{
		bc:     bc,
		cc:     chainConfig,
		txpool: txpool,
		mux:    mux,
		db:     db,
		am:     accountMgr,
		synced: isSynchronised,
	}

	return bv
}

func (bv *BlockVoting) resetPendingState(parent *types.Block) {
	publicState, privateState, err := bv.bc.State()
	if err != nil {
		panic(fmt.Sprintf("State error: %v", err))
	}

	ps := &pendingState{
		parent:        parent,
		publicState:   publicState,
		privateState:  privateState,
		header:        bv.makeHeader(parent),
		gp:            new(core.GasPool),
		ownedAccounts: accountAddressesSet(bv.am.Accounts()),
	}

	ps.gp.AddGas(ps.header.GasLimit)

	txs := types.NewTransactionsByPriorityAndNonce(bv.txpool.Pending())

	lowGasTxs, failedTxs := ps.applyTransactions(txs, bv.mux, bv.bc, bv.cc)
	bv.txpool.RemoveBatch(lowGasTxs)
	bv.txpool.RemoveBatch(failedTxs)

	bv.pStateMu.Lock()
	bv.pState = ps
	bv.pStateMu.Unlock()
}

func (bv *BlockVoting) makeHeader(parent *types.Block) *types.Header {
	tstart := time.Now()
	tstamp := tstart.Unix()
	if parent.Time().Cmp(new(big.Int).SetInt64(tstamp)) >= 0 {
		tstamp = parent.Time().Int64() + 1
	}
	// this will ensure we're not going off too far in the future
	if now := time.Now().Unix(); tstamp > now+4 {
		wait := time.Duration(tstamp-now) * time.Second
		glog.V(logger.Info).Infoln("We are too far in the future. Waiting for", wait)
		time.Sleep(wait)
	}

	num := parent.Number()
	header := &types.Header{
		Number:     num.Add(num, common.Big1),
		ParentHash: parent.Hash(),
		Difficulty: core.CalcDifficulty(bv.cc, uint64(tstamp), parent.Time().Uint64(), parent.Number(), parent.Difficulty()),
		GasLimit:   core.CalcGasLimit(parent),
		GasUsed:    new(big.Int),
		Time:       big.NewInt(tstamp),
	}

	if bv.bmk != nil {
		header.Coinbase = crypto.PubkeyToAddress(bv.bmk.PublicKey)
	}

	return header
}

// Start runs the event loop.
func (bv *BlockVoting) Start(client *rpc.Client, strat BlockMakerStrategy, voteKey, blockMakerKey *ecdsa.PrivateKey) error {
	bv.bmk = blockMakerKey
	bv.vk = voteKey

	ethClient := ethclient.NewClient(client)
	callContract, err := NewVotingContractCaller(params.QuorumVotingContractAddr, ethClient)
	if err != nil {
		return err
	}
	bv.callContract = callContract

	if voteKey != nil {
		contract, err := NewVotingContract(params.QuorumVotingContractAddr, ethClient)
		if err != nil {
			return err
		}

		auth := bind.NewKeyedTransactor(voteKey)
		bv.voteSession = &VotingContractSession{
			Contract: contract,
			CallOpts: bind.CallOpts{
				Pending: true,
			},
			TransactOpts: bind.TransactOpts{
				From:   auth.From,
				Signer: auth.Signer,
			},
		}
	}

	bv.run(strat)

	return nil
}

func (bv *BlockVoting) run(strat BlockMakerStrategy) {
	if bv.bmk != nil {
		glog.Infof("Node configured for block creation: %s", crypto.PubkeyToAddress(bv.bmk.PublicKey).Hex())
	}
	if bv.vk != nil {
		glog.Infof("Node configured for block voting: %s", crypto.PubkeyToAddress(bv.vk.PublicKey).Hex())
	}

	sub := bv.mux.Subscribe(downloader.StartEvent{},
		downloader.DoneEvent{},
		downloader.FailedEvent{},
		core.ChainHeadEvent{},
		core.TxPreEvent{},
		Vote{},
		CreateBlock{})

	bv.resetPendingState(bv.bc.CurrentBlock())

	go func() {
		defer sub.Unsubscribe()

		strat.Start()

		for {
			select {
			case event, ok := <-sub.Chan():
				if !ok {
					return
				}

				switch e := event.Data.(type) {
				case downloader.StartEvent: // begin synchronising, stop block creation and/or voting
					bv.synced = false
					strat.Pause()
				case downloader.DoneEvent, downloader.FailedEvent: // caught up, or got an error, start block createion and/or voting
					bv.synced = true
					strat.Resume()
				case core.ChainHeadEvent: // got a new header, reset pending state
					bv.resetPendingState(e.Block)
					if bv.synced {
						number := new(big.Int)
						number.Add(e.Block.Number(), common.Big1)
						if tx, err := bv.vote(number, e.Block.Hash()); err == nil {
							if glog.V(logger.Debug) {
								glog.Infof("Voted for %s on height %d in tx %s", e.Block.Hash().Hex(), number, tx.Hex())
							}
						} else if glog.V(logger.Debug) {
							glog.Errorf("Unable to vote: %v", err)
						}
					}
				case core.TxPreEvent: // tx entered pool, apply to pending state
					bv.applyTransaction(e.Tx)
				case Vote:
					if bv.synced {
						txHash, err := bv.vote(e.Number, e.Hash)
						if err == nil && e.TxHash != nil {
							e.TxHash <- txHash
						} else if err != nil && e.Err != nil {
							e.Err <- err
						} else if err != nil {
							if glog.V(logger.Debug) {
								glog.Errorf("Unable to vote: %v", err)
							}
						}
					} else {
						e.Err <- fmt.Errorf("Node not synced")
					}
				case CreateBlock:
					block, err := bv.createBlock()
					if err == nil && e.Hash != nil {
						e.Hash <- block.Hash()
					} else if err != nil && e.Err != nil {
						e.Err <- err
					} else if err != nil {
						if glog.V(logger.Debug) {
							glog.Errorf("Unable to create block: %v", err)
						}
					}

					if err != nil {
						bv.pStateMu.Lock()
						cBlock := bv.pState.parent
						bv.pStateMu.Unlock()
						num := new(big.Int).Add(cBlock.Number(), common.Big1)
						_, err := bv.vote(num, cBlock.Hash())
						if err != nil {
							glog.Errorf("Unable to vote: %v", err)
							bv.resetPendingState(bv.bc.CurrentBlock())
						}
					}
				}
			}
		}
	}()
}

func (bv *BlockVoting) applyTransaction(tx *types.Transaction) {
	acc, _ := tx.From()
	txs := map[common.Address]types.Transactions{acc: types.Transactions{tx}}
	txset := types.NewTransactionsByPriorityAndNonce(txs)

	bv.pStateMu.Lock()
	bv.pState.applyTransactions(txset, bv.mux, bv.bc, bv.cc)
	bv.pStateMu.Unlock()
}

func (bv *BlockVoting) Pending() (*types.Block, *state.StateDB, *state.StateDB) {
	bv.pStateMu.Lock()
	defer bv.pStateMu.Unlock()
	return types.NewBlock(bv.pState.header, bv.pState.txs, nil, bv.pState.receipts), bv.pState.publicState.Copy(), bv.pState.privateState.Copy()
}

func (bv *BlockVoting) createBlock() (*types.Block, error) {
	if bv.bmk == nil {
		return nil, fmt.Errorf("Node not configured for block creation")
	}

	ch, err := bv.canonHash(bv.pState.header.Number.Uint64())
	if err != nil {
		return nil, err
	}
	if ch != bv.pState.parent.Hash() {
		return nil, fmt.Errorf("invalid canonical hash, expected %s got %s", ch.Hex(), bv.pState.header.Hash().Hex())
	}

	bv.pStateMu.Lock()
	defer bv.pStateMu.Unlock()

	state := bv.pState.publicState // shortcut
	header := bv.pState.header
	receipts := bv.pState.receipts

	core.AccumulateRewards(state, header, nil)

	header.Root = state.IntermediateRoot()

	// Quorum blocks contain a signature of the header in the Extra field.
	// This signature is verified during block import and ensures that the
	// block is created by a party that is allowed to create blocks.
	signature, err := crypto.Sign(header.QuorumHash().Bytes(), bv.bmk)
	if err != nil {
		return nil, err
	}
	header.Extra = signature

	// update block hash in receipts and logs now it is available
	for _, r := range receipts {
		for _, l := range r.Logs {
			l.BlockHash = header.Hash()
		}
	}

	header.Bloom = types.CreateBloom(receipts)

	block := types.NewBlock(header, bv.pState.txs, nil, receipts)
	if _, err := bv.bc.InsertChain(types.Blocks{block}); err != nil {
		return nil, err
	}

	bv.mux.Post(core.NewMinedBlockEvent{Block: block})

	return block, nil
}

func (bv *BlockVoting) vote(height *big.Int, hash common.Hash) (common.Hash, error) {
	if bv.voteSession == nil {
		return common.Hash{}, fmt.Errorf("Node is not configured for voting")
	}
	cv, err := bv.callContract.CanVote(nil, bv.voteSession.TransactOpts.From)
	if err != nil {
		return common.Hash{}, err
	}
	if !cv {
		return common.Hash{}, fmt.Errorf("%s is not allowed to vote", bv.voteSession.TransactOpts.From.Hex())
	}

	if glog.V(logger.Detail) {
		glog.Infof("vote for %s on height %d", hash.Hex(), height)
	}

	nonce := bv.txpool.Nonce(bv.voteSession.TransactOpts.From)
	bv.voteSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)
	defer func() { bv.voteSession.TransactOpts.Nonce = nil }()

	tx, err := bv.voteSession.Vote(height, hash)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

// CanonHash returns the canonical block hash on the given height.
func (bv *BlockVoting) canonHash(height uint64) (common.Hash, error) {
	opts := &bind.CallOpts{Pending: true}
	return bv.callContract.GetCanonHash(opts, new(big.Int).SetUint64(height))
}

// isVoter returns an indication if the given address is allowed
// to vote.
func (bv *BlockVoting) isVoter(addr common.Address) (bool, error) {
	return bv.callContract.IsVoter(nil, addr)
}

// isBlockMaker returns an indication if the given address is allowed
// to make blocks
func (bv *BlockVoting) isBlockMaker(addr common.Address) (bool, error) {
	return bv.callContract.IsBlockMaker(nil, addr)
}

func accountAddressesSet(accounts []accounts.Account) *set.Set {
	accountSet := set.New()
	for _, account := range accounts {
		accountSet.Add(account.Address)
	}
	return accountSet
}
