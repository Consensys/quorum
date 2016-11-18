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

	// browser solidity with optimizations: 0.4.2+commit.af6afb04.mod.Emscripten.clang
	RuntimeCode = "606060405236156100c45760e060020a60003504631290948581146100c9578063284d163c146100fe57806342169e481461013a578063488099a6146101485780634fe437d514610168578063559c390c1461017657806368bb8bb61461027b57806372a571fc146102eb57806386c1ff681461039157806398ba676d146103cd578063a7771ee31461043d578063adfaa72e1461046a578063cf5289851461048a578063de8fa43114610498578063e814d1c7146104b3578063f4ab9adf146104df575b610002565b3461000257610598600435600160a060020a03331660009081526003602052604090205460ff16156100c45760018190555b50565b3461000257610598600435600160a060020a03331660009081526005602052604090205460ff16156100c457600454600114156105ae57610002565b34610002576104a160025481565b346100025761059a60043560056020526000908152604090205460ff1681565b34610002576104a160015481565b34610002576104a160043560006000600060006000600050600186038154811015610002579080526002027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630192505b60018301548110156106275760018301805484916000918490811015610002576000918252602080832090910154835282810193909352604091820181205485825292869052205410801561025057506001805490840180548591600091859081101561000257906000526020600020900160005054815260208101919091526040016000205410155b156102735760018301805482908110156100025760009182526020909120015491505b6001016101c6565b3461000257610598600435602435600160a060020a03331660009081526003602052604081205460ff16156100c45780548390101561063457805480840381018083559082908290801582901161062f5760020281600202836000526020600020918201910161062f91906106bb565b3461000257610598600435600160a060020a03331660009081526005602052604090205460ff16156100c457600160a060020a0381166000908152604090205460ff1615156100fb5760406000819020805460ff191660019081179091556004805490910190558051600160a060020a038316815290517f1a4ce6942f7aa91856332e618fc90159f13a340611a308f5d7327ba0707e56859181900360200190a16100fb565b3461000257610598600435600160a060020a03331660009081526003602052604090205460ff16156100c4576002546001141561076457610002565b34610002576104a1600435602435600060006000600050600185038154811015610002579080526002027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630181509050806001016000508381548110156100025750825250602090200154919050565b346100025761059a600435600160a060020a03811660009081526003602052604090205460ff165b919050565b346100025761059a60043560036020526000908152604090205460ff1681565b34610002576104a160045481565b34610002576000545b60408051918252519081900360200190f35b346100025761059a600435600160a060020a03811660009081526005602052604090205460ff16610465565b3461000257610598600435600160a060020a03331660009081526003602052604090205460ff16156100c457600160a060020a03811660009081526003602052604090205460ff1615156100fb5760406000818120600160a060020a0384169182905260036020908152815460ff1916600190811790925560028054909201909155825191825291517f0ad2eca75347acd5160276fe4b5dad46987e4ff4af9e574195e3e9bc15d7e0ff929181900390910190a16100fb565b005b604080519115158252519081900360200190f35b600160a060020a03811660009081526005602052604090205460ff16156100fb5760406000819020805460ff19169055600480546000190190558051600160a060020a038316815290517f8cee3054364d6799f1c8962580ad61273d9d38ca1ff26516bd1ad23c099a60229181900360200190a16100fb565b509392505050565b505050505b60008054600019850190811015610002578382526002027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301602081905260408220549092501415610708578060010160005080548060010182818154818355818115116106f5578183600052602060002091820191016106f591906106dd565b50506002015b808211156106f1576001810180546000808355918252602082206106b5918101905b808211156106f157600081556001016106dd565b5090565b5050506000928352506020909120018290555b600082815260208281526040918290208054600101905581514381529081018490528151600160a060020a033316927f3d03ba7f4b5227cdb385f2610906e5bcee147171603ec40005b30915ad20e258928290030190a2505050565b600160a060020a03811660009081526003602052604090205460ff16156100fb5760406000819020805460ff19169055600280546000190190558051600160a060020a038316815290517f183393fc5cffbfc7d03d623966b85f76b9430f42d3aada2ac3f3deabc78899e89181900360200190a16100fb56"
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
