package eth

import (
	"errors"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/forkid"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	qlightproto "github.com/ethereum/go-ethereum/eth/protocols/qlight"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/qlight"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

type qlightServerHandler ethHandler

func (h *qlightServerHandler) Chain() *core.BlockChain     { return h.chain }
func (h *qlightServerHandler) StateBloom() *trie.SyncBloom { return h.stateBloom }
func (h *qlightServerHandler) TxPool() eth.TxPool          { return h.txpool }

func (h *qlightServerHandler) RunPeer(peer *eth.Peer, handler eth.Handler) error {
	return nil
}
func (h *qlightServerHandler) Handle(peer *eth.Peer, packet eth.Packet) error {
	return (*ethHandler)(h).Handle(peer, packet)
}

func (h *qlightServerHandler) RunQPeer(peer *qlightproto.Peer, hand qlightproto.Handler) error {
	return (*handler)(h).runQLightServerPeer(peer, hand)
}

// PeerInfo retrieves all known `eth` information about a peer.
func (h *qlightServerHandler) PeerInfo(id enode.ID) interface{} {
	if p := h.peers.peer(id.String()); p != nil {
		return p.info()
	}
	return nil
}

// AcceptTxs retrieves whether transaction processing is enabled on the node
// or if inbound transactions should simply be dropped.
func (h *qlightServerHandler) AcceptTxs() bool {
	return atomic.LoadUint32(&h.acceptTxs) == 1
}

// newHandler returns a handler for all Ethereum chain management protocol.
func newQLightServerHandler(config *handlerConfig) (*handler, error) {
	// Create the protocol manager with the base fields
	h := &handler{
		networkID:                config.Network,
		forkFilter:               forkid.NewFilter(config.Chain),
		eventMux:                 config.EventMux,
		database:                 config.Database,
		txpool:                   config.TxPool,
		chain:                    config.Chain,
		peers:                    newPeerSet(),
		authorizationList:        config.AuthorizationList,
		txsyncCh:                 make(chan *txsync),
		quitSync:                 make(chan struct{}),
		raftMode:                 config.RaftMode,
		engine:                   config.Engine,
		authProvider:             config.authProvider,
		privateBlockDataResolver: config.privateBlockDataResolver,
	}

	return h, nil
}

// runEthPeer registers an eth peer into the joint eth/snap peerset, adds it to
// various subsistems and starts handling messages.
func (h *handler) runQLightServerPeer(peer *qlightproto.Peer, handler qlightproto.Handler) error {
	h.peerWG.Add(1)
	defer h.peerWG.Done()

	// Execute the Ethereum handshake
	var (
		genesis = h.chain.Genesis()
		head    = h.chain.CurrentHeader()
		hash    = head.Hash()
		number  = head.Number.Uint64()
		td      = h.chain.GetTd(hash, number)
	)
	forkID := forkid.NewID(h.chain.Config(), h.chain.Genesis().Hash(), h.chain.CurrentHeader().Number.Uint64())
	if err := peer.EthPeer.Handshake(h.networkID, td, hash, genesis.Hash(), forkID, h.forkFilter); err != nil {
		peer.Log().Debug("Ethereum handshake failed", "err", err)

		// Quorum
		// When the Handshake() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
		peer.EthPeerDisconnected <- struct{}{}
		// End Quorum
		return err
	}

	log.Info("QLight attempting handshake")
	if err := peer.QLightHandshake(true, "", ""); err != nil {
		peer.Log().Debug("QLight handshake failed", "err", err)
		log.Info("QLight handshake failed", "err", err)

		// Quorum
		// When the Handshake() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
		peer.EthPeerDisconnected <- struct{}{}
		// End Quorum
		return err
	}

	peer.Log().Debug("QLight handshake result for peer", "peer", peer.ID(), "server", peer.QLightServer(), "psi", peer.QLightPSI(), "token", peer.QLightToken())
	log.Info("QLight handshake result for peer", "peer", peer.ID(), "server", peer.QLightServer(), "psi", peer.QLightPSI(), "token", peer.QLightToken())
	// if we're not connected to a qlight server - disconnect the peer
	if peer.QLightServer() {
		peer.Log().Debug("QLight server connected to a server peer. Disconnecting.")

		// Quorum
		// When the Handshake() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
		peer.EthPeerDisconnected <- struct{}{}
		// End Quorum
		return fmt.Errorf("connected to a server peer")
	}

	// Ignore maxPeers if this is a trusted peer
	if !peer.Peer.Info().Network.Trusted {
		if h.peers.len() >= h.maxPeers {
			return p2p.DiscTooManyPeers
		}
	}
	peer.Log().Debug("Ethereum peer connected", "name", peer.Name())

	err := h.authProvider.Authorize(peer.QLightToken(), peer.QLightPSI())
	if err != nil {
		peer.Log().Error("Auth error", "err", err)
		return p2p.DiscAuthError
	}

	// Register the peer locally
	if err := h.peers.registerQPeer(peer); err != nil {
		peer.Log().Error("Ethereum peer registration failed", "err", err)

		// Quorum
		// When the Register() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
		peer.EthPeerDisconnected <- struct{}{}
		// End Quorum

		return err
	}
	defer h.removeQLightServerPeer(peer.ID())

	// start periodic auth checks
	peer.QLightPeriodicAuthFunc = func() error { return h.authProvider.Authorize(peer.QLightToken(), peer.QLightPSI()) }
	go peer.PeriodicAuthCheck()

	p := h.peers.peer(peer.ID())
	if p == nil {
		return errors.New("peer dropped during handling")
	}

	// Propagate existing transactions. new transactions appearing
	// after this will be sent via broadcasts.
	h.syncTransactions(peer.EthPeer)

	// Quorum notify other subprotocols that the eth peer is ready, and has been added to the peerset.
	p.EthPeerRegistered <- struct{}{}
	// Quorum

	// Handle incoming messages until the connection is torn down
	return handler(peer)
}

func (h *handler) StartQLightServer(maxPeers int) {
	h.maxPeers = maxPeers
	h.wg.Add(1)
	h.txsCh = make(chan core.NewTxsEvent, txChanSize)
	h.txsSub = h.txpool.SubscribeNewTxsEvent(h.txsCh)
	go h.txBroadcastLoop()

	// broadcast mined blocks
	h.wg.Add(1)
	go h.newBlockBroadcastLoop()

	h.authProvider.Initialize()
}

func (h *handler) StopQLightServer() {
	h.txsSub.Unsubscribe()
	close(h.quitSync)
	h.wg.Wait()

	// Disconnect existing sessions.
	// This also closes the gate for any new registrations on the peer set.
	// sessions which are already established but not added to h.peers yet
	// will exit when they try to register.
	h.peers.close()
	h.peerWG.Wait()
	log.Info("QLight server protocol stopped")
}

func (h *handler) newBlockBroadcastLoop() {
	defer h.wg.Done()

	headCh := make(chan core.ChainHeadEvent, 10)
	headSub := h.chain.SubscribeChainHeadEvent(headCh)
	defer headSub.Unsubscribe()

	for {
		select {
		case ev := <-headCh:
			log.Debug("Announcing block to peers", "number", ev.Block.Number(), "hash", ev.Block.Hash(), "td", ev.Block.Difficulty())
			h.BroadcastBlockQLServer(ev.Block)

		case <-h.quitSync:
			return
		}
	}
}

func (h *handler) BroadcastBlockQLServer(block *types.Block) {
	hash := block.Hash()
	peers := h.peers.qlightPeersWithoutBlock(hash)

	// Calculate the TD of the block (it's not imported yet, so block.Td is not valid)
	var td *big.Int
	if parent := h.chain.GetBlock(block.ParentHash(), block.NumberU64()-1); parent != nil {
		td = new(big.Int).Add(block.Difficulty(), h.chain.GetTd(block.ParentHash(), block.NumberU64()-1))
	} else {
		log.Error("Propagating dangling block", "number", block.Number(), "hash", hash)
		return
	}
	// Send the block to a subset of our peers
	for _, peer := range peers {
		log.Info("Preparing new block private data")
		blockPrivateData, err := h.privateBlockDataResolver.PrepareBlockPrivateData(block, peer.qlight.QLightPSI())
		if err != nil {
			log.Error("Unable to prepare private data for block", "number", block.Number(), "hash", hash, "err", err, "psi", peer.qlight.QLightPSI())
			return
		}
		log.Info("Private transactions data", "is nil", blockPrivateData == nil)
		peer.qlight.AsyncSendNewBlock(block, td, blockPrivateData)
	}
	log.Trace("Propagated block", "hash", hash, "recipients", len(peers), "duration", common.PrettyDuration(time.Since(block.ReceivedAt)))
}

// removePeer unregisters a peer from the downloader and fetchers, removes it from
// the set of tracked peers and closes the network connection to it.
func (h *handler) removeQLightServerPeer(id string) {
	// Create a custom logger to avoid printing the entire id
	var logger log.Logger
	if len(id) < 16 {
		// Tests use short IDs, don't choke on them
		logger = log.New("peer", id)
	} else {
		logger = log.New("peer", id[:8])
	}
	// Abort if the peer does not exist
	peer := h.peers.peer(id)
	if peer == nil {
		logger.Error("Ethereum peer removal failed", "err", errPeerNotRegistered)
		return
	}
	// Remove the `eth` peer if it exists
	logger.Debug("Removing QLight server peer", "snap", peer.snapExt != nil)

	if err := h.peers.unregisterPeer(id); err != nil {
		logger.Error("Ethereum peer removal failed", "err", err)
	}
	// Hard disconnect at the networking layer
	peer.Peer.Disconnect(p2p.DiscUselessPeer)
}

func (ps *peerSet) qlightPeersWithoutBlock(hash common.Hash) []*ethPeer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	list := make([]*ethPeer, 0, len(ps.peers))
	for _, p := range ps.peers {
		if !p.qlight.KnownBlock(hash) {
			list = append(list, p)
		}
	}
	return list
}

// Handle is invoked from a peer's message handler when it receives a new remote
// message that the handler couldn't consume and serve itself.
func (h *qlightServerHandler) QHandle(peer *qlightproto.Peer, packet eth.Packet) error {
	// Consume any broadcasts and announces, forwarding the rest to the downloader
	switch packet := packet.(type) {
	case *eth.NewPooledTransactionHashesPacket:
		return (*ethHandler)(h).Handle(peer.EthPeer, packet)
	case *eth.TransactionsPacket:
		return (*ethHandler)(h).Handle(peer.EthPeer, packet)
	case *eth.PooledTransactionsPacket:
		return (*ethHandler)(h).Handle(peer.EthPeer, packet)
	case *eth.GetBlockBodiesPacket:
		return h.handleGetBlockBodies(packet, peer)
	default:
		return fmt.Errorf("unexpected eth packet type: %T", packet)
	}
}

func (h *qlightServerHandler) handleGetBlockBodies(query *eth.GetBlockBodiesPacket, peer *qlightproto.Peer) error {
	blockPublicData, blockPrivateData, err := h.answerGetBlockBodiesQuery(query, peer)
	if err != nil {
		return err
	}
	if len(blockPrivateData) > 0 {
		err := peer.SendBlockPrivateData(blockPrivateData)
		if err != nil {
			log.Info("Error occurred while sending private data msg", "err", err)
			return err
		}
	}
	return peer.EthPeer.SendBlockBodiesRLP(blockPublicData)
}

const (
	// softResponseLimit is the target maximum size of replies to data retrievals.
	softResponseLimit = 2 * 1024 * 1024

	maxBodiesServe = 1024
)

func (h *qlightServerHandler) answerGetBlockBodiesQuery(query *eth.GetBlockBodiesPacket, peer *qlightproto.Peer) ([]rlp.RawValue, []qlight.BlockPrivateData, error) {
	// Gather blocks until the fetch or network limits is reached
	var (
		bytes             int
		bodies            []rlp.RawValue
		blockPrivateDatas []qlight.BlockPrivateData
	)
	for lookups, hash := range *query {
		if bytes >= softResponseLimit || len(bodies) >= maxBodiesServe ||
			lookups >= 2*maxBodiesServe {
			break
		}
		block := h.chain.GetBlockByHash(hash)
		if block != nil {
			if bpd, err := h.privateBlockDataResolver.PrepareBlockPrivateData(block, peer.QLightPSI()); err != nil {
				return nil, nil, fmt.Errorf("Unable to produce block private transaction data %v: %v", hash, err)
			} else if bpd != nil {
				blockPrivateDatas = append(blockPrivateDatas, *bpd)
			}
			// TODO qlight - add soft limits for block private data as well
		}
		if data := h.chain.GetBodyRLP(hash); len(data) != 0 {
			bodies = append(bodies, data)
			bytes += len(data)
		}
	}
	return bodies, blockPrivateDatas, nil
}
