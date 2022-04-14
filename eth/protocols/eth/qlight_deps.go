package eth

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/p2p"
)

func CurrentENREntry(chain *core.BlockChain) *enrEntry {
	return currentENREntry(chain)
}

func NodeInfoFunc(chain *core.BlockChain, network uint64) *NodeInfo {
	return nodeInfo(chain, network)
}

var ETH_65_FULL_SYNC = map[uint64]msgHandler{
	// old 64 messages
	GetBlockHeadersMsg: handleGetBlockHeaders,
	BlockHeadersMsg:    handleBlockHeaders,
	GetBlockBodiesMsg:  handleGetBlockBodies,
	BlockBodiesMsg:     handleBlockBodies,
	NewBlockHashesMsg:  handleNewBlockhashes,
	NewBlockMsg:        handleNewBlock,
	TransactionsMsg:    handleTransactions,
	// New eth65 messages
	NewPooledTransactionHashesMsg: handleNewPooledTransactionHashes,
	GetPooledTransactionsMsg:      handleGetPooledTransactions,
	PooledTransactionsMsg:         handlePooledTransactions,
}

func NewPeerWithTxBroadcast(version uint, p *p2p.Peer, rw p2p.MsgReadWriter, txpool TxPool) *Peer {
	peer := NewPeerNoBroadcast(version, p, rw, txpool)
	// Start up all the broadcasters
	go peer.broadcastTransactions()
	if version >= ETH65 {
		go peer.announceTransactions()
	}
	return peer
}

func NewPeerNoBroadcast(version uint, p *p2p.Peer, rw p2p.MsgReadWriter, txpool TxPool) *Peer {
	peer := &Peer{
		id:              p.ID().String(),
		Peer:            p,
		rw:              rw,
		version:         version,
		knownTxs:        mapset.NewSet(),
		knownBlocks:     mapset.NewSet(),
		queuedBlocks:    make(chan *blockPropagation, maxQueuedBlocks),
		queuedBlockAnns: make(chan *types.Block, maxQueuedBlockAnns),
		txBroadcast:     make(chan []common.Hash),
		txAnnounce:      make(chan []common.Hash),
		txpool:          txpool,
		term:            make(chan struct{}),
	}
	return peer
}

func (p *Peer) MarkBlock(hash common.Hash) {
	p.markBlock(hash)
}

func (p *Peer) MarkTransaction(hash common.Hash) {
	p.markTransaction(hash)
}
