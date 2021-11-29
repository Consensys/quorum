// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package eth

import (
	"errors"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/forkid"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/fetcher"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/qlight"
	"github.com/ethereum/go-ethereum/trie"
)

type QLightClientProtocolManager struct {
	*ProtocolManager
	dialCandidates     enode.Iterator
	psi                string
	token              string
	privateClientCache qlight.PrivateClientCache
}

// NewProtocolManager returns a new Ethereum sub protocol manager. The Ethereum sub protocol manages peers capable
// with the Ethereum network.
func NewQLightClientProtocolManager(config *params.ChainConfig, checkpoint *params.TrustedCheckpoint, mode downloader.SyncMode, networkID uint64, mux *event.TypeMux, txpool txPool, engine consensus.Engine, blockchain *core.BlockChain, chaindb ethdb.Database, cacheLimit int, authorizationList map[uint64]common.Hash, raftMode bool, psi string, token string, privateClientCache qlight.PrivateClientCache) (*QLightClientProtocolManager, error) {
	// Create the protocol manager with the base fields
	manager := &QLightClientProtocolManager{
		ProtocolManager: &ProtocolManager{
			networkID:         networkID,
			forkFilter:        forkid.NewFilter(blockchain),
			eventMux:          mux,
			txpool:            txpool,
			blockchain:        blockchain,
			chaindb:           chaindb,
			peers:             newPeerSet(),
			authorizationList: authorizationList,
			txsyncCh:          make(chan *txsync),
			quitSync:          make(chan struct{}),
			raftMode:          raftMode,
			engine:            engine,
		},
		psi:                psi,
		token:              token,
		privateClientCache: privateClientCache,
	}

	// Quorum
	if handler, ok := manager.engine.(consensus.Handler); ok {
		handler.SetBroadcaster(manager)
	}
	// /Quorum

	if mode == downloader.FullSync {
		// The database seems empty as the current block is the genesis. Yet the fast
		// block is ahead, so fast sync was enabled for this node at a certain point.
		// The scenarios where this can happen is
		// * if the user manually (or via a bad block) rolled back a fast sync node
		//   below the sync point.
		// * the last fast sync is not finished while user specifies a full sync this
		//   time. But we don't have any recent state for full sync.
		// In these cases however it's safe to reenable fast sync.
		fullBlock, fastBlock := blockchain.CurrentBlock(), blockchain.CurrentFastBlock()
		if fullBlock.NumberU64() == 0 && fastBlock.NumberU64() > 0 {
			manager.fastSync = uint32(1)
			log.Warn("Switch sync mode from full sync to fast sync")
		}
	} else {
		if blockchain.CurrentBlock().NumberU64() > 0 {
			// Print warning log if database is not empty to run fast sync.
			log.Warn("Switch sync mode from fast sync to full sync")
		} else {
			// If fast sync was requested and our database is empty, grant it
			manager.fastSync = uint32(1)
		}
	}

	// If we have trusted checkpoints, enforce them on the chain
	if checkpoint != nil {
		manager.checkpointNumber = (checkpoint.SectionIndex+1)*params.CHTFrequency - 1
		manager.checkpointHash = checkpoint.SectionHead
	}

	// Construct the downloader (long sync) and its backing state bloom if fast
	// sync is requested. The downloader is responsible for deallocating the state
	// bloom when it's done.
	var stateBloom *trie.SyncBloom
	if atomic.LoadUint32(&manager.fastSync) == 1 {
		stateBloom = trie.NewSyncBloom(uint64(cacheLimit), chaindb)
	}
	manager.downloader = downloader.New(manager.checkpointNumber, chaindb, stateBloom, manager.eventMux, blockchain, nil, manager.removePeer)

	// Construct the fetcher (short sync)
	validator := func(header *types.Header) error {
		return engine.VerifyHeader(blockchain, header, true)
	}
	heighter := func() uint64 {
		return blockchain.CurrentBlock().NumberU64()
	}
	inserter := func(blocks types.Blocks) (int, error) {
		// If sync hasn't reached the checkpoint yet, deny importing weird blocks.
		//
		// Ideally we would also compare the head block's timestamp and similarly reject
		// the propagated block if the head is too old. Unfortunately there is a corner
		// case when starting new networks, where the genesis might be ancient (0 unix)
		// which would prevent full nodes from accepting it.
		if manager.blockchain.CurrentBlock().NumberU64() < manager.checkpointNumber {
			log.Warn("Unsynced yet, discarded propagated block", "number", blocks[0].Number(), "hash", blocks[0].Hash())
			return 0, nil
		}
		// If fast sync is running, deny importing weird blocks. This is a problematic
		// clause when starting up a new network, because fast-syncing miners might not
		// accept each others' blocks until a restart. Unfortunately we haven't figured
		// out a way yet where nodes can decide unilaterally whether the network is new
		// or not. This should be fixed if we figure out a solution.
		if atomic.LoadUint32(&manager.fastSync) == 1 {
			log.Warn("Fast syncing, discarded propagated block", "number", blocks[0].Number(), "hash", blocks[0].Hash())
			return 0, nil
		}
		n, err := manager.blockchain.InsertChain(blocks)
		if err == nil {
			atomic.StoreUint32(&manager.acceptTxs, 1) // Mark initial sync done on any fetcher import
		}
		return n, err
	}
	manager.blockFetcher = fetcher.NewBlockFetcher(false, nil, blockchain.GetBlockByHash, validator, manager.BroadcastBlock, heighter, nil, inserter, manager.removePeer)

	fetchTx := func(peer string, hashes []common.Hash) error {
		p := manager.peers.Peer(peer)
		if p == nil {
			return errors.New("unknown peer")
		}
		return p.RequestTxs(hashes)
	}
	manager.txFetcher = fetcher.NewTxFetcher(txpool.Has, txpool.AddRemotes, fetchTx)

	manager.chainSync = newChainSyncer(manager.ProtocolManager)

	return manager, nil
}

func (pm *QLightClientProtocolManager) makeProtocol(version uint) p2p.Protocol {
	return p2p.Protocol{
		Name:    "qlight",
		Version: version,
		Length:  qlightProtocolLength,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			return pm.runPeer(pm.newPeer(int(version), p, rw, pm.txpool.Get), "qlight")
		},
		NodeInfo: func() interface{} {
			return pm.NodeInfo()
		},
		PeerInfo: func(id enode.ID) interface{} {
			if p := pm.peers.Peer(fmt.Sprintf("%x", id[:8])); p != nil {
				return p.Info()
			}
			return nil
		},
		DialCandidates: pm.dialCandidates,
	}
}

func (pm *QLightClientProtocolManager) removePeer(id string) {
	// Short circuit if the peer was already removed
	peer := pm.peers.Peer(id)
	if peer == nil {
		return
	}
	log.Debug("Removing Ethereum peer", "peer", id)

	// Unregister the peer from the downloader and Ethereum peer set
	pm.downloader.UnregisterPeer(id)
	pm.txFetcher.Drop(id)

	if err := pm.peers.Unregister(id); err != nil {
		log.Error("Peer removal failed", "peer", id, "err", err)
	}
	// Hard disconnect at the networking layer
	if peer != nil {
		peer.Peer.Disconnect(p2p.DiscUselessPeer)
	}
}

func (pm *QLightClientProtocolManager) Start(maxPeers int) {
	pm.maxPeers = maxPeers

	// Quorum
	// We set this immediately to make sure the qlight client never drops
	// incoming txes. Qlight client doesn't use the fetcher or downloader, and so
	// this would never be set otherwise.
	atomic.StoreUint32(&pm.acceptTxs, 1)
	// /Quorum

	// start sync handlers
	pm.wg.Add(2)
	go pm.chainSync.loop()
	go pm.txsyncLoop64() // TODO(karalabe): Legacy initial tx echange, drop with eth/64.
}

func (pm *QLightClientProtocolManager) Stop() {
	// Quit chainSync and txsync64.
	// After this is done, no new peers will be accepted.
	close(pm.quitSync)
	pm.wg.Wait()

	// Disconnect existing sessions.
	// This also closes the gate for any new registrations on the peer set.
	// sessions which are already established but not added to pm.peers yet
	// will exit when they try to register.
	pm.peers.Close()
	pm.peerWG.Wait()

	log.Info("Ethereum protocol stopped")
}

func (pm *QLightClientProtocolManager) newPeer(pv int, p *p2p.Peer, rw p2p.MsgReadWriter, getPooledTx func(hash common.Hash) *types.Transaction) *peer {
	return newPeer(pv, p, rw, getPooledTx)
}

// Quorum - added protoName argument
func (pm *QLightClientProtocolManager) runPeer(p *peer, protoName string) error {
	if !pm.chainSync.handlePeerEvent(p) {
		return p2p.DiscQuitting
	}
	pm.peerWG.Add(1)
	defer pm.peerWG.Done()
	return pm.handle(p, protoName)
}

// quorum: protoname is either "eth" or a subprotocol that overrides "eth", e.g. legacy "istanbul/99"
// handle is the callback invoked to manage the life cycle of an eth peer. When
// this function terminates, the peer is disconnected.
func (pm *QLightClientProtocolManager) handle(p *peer, protoName string) error {
	// Ignore maxPeers if this is a trusted peer
	if pm.peers.Len() >= pm.maxPeers && !p.Peer.Info().Network.Trusted {
		return p2p.DiscTooManyPeers
	}
	p.Log().Debug("Ethereum peer connected", "name", p.Name())

	// Execute the Ethereum handshake
	var (
		genesis = pm.blockchain.Genesis()
		head    = pm.blockchain.CurrentHeader()
		hash    = head.Hash()
		number  = head.Number.Uint64()
		td      = pm.blockchain.GetTd(hash, number)
	)
	forkID := forkid.NewID(pm.blockchain.Config(), pm.blockchain.Genesis().Hash(), pm.blockchain.CurrentHeader().Number.Uint64())
	log.Info("QLight attempting eth handshake")
	if err := p.Handshake(pm.networkID, td, hash, genesis.Hash(), forkID, pm.forkFilter, protoName); err != nil {
		p.Log().Debug("Ethereum handshake failed", "protoName", protoName, "err", err)

		// Quorum
		// When the Handshake() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
		p.EthPeerDisconnected <- struct{}{}
		// End Quorum
		return err
	}

	log.Info("QLight attempting handshake")
	if err := p.QLightHandshake(false, pm.psi, pm.token); err != nil {
		p.Log().Debug("QLight handshake failed", "protoName", protoName, "err", err)
		log.Info("QLight handshake failed", "protoName", protoName, "err", err)

		// Quorum
		// When the Handshake() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
		p.EthPeerDisconnected <- struct{}{}
		// End Quorum
		return err
	}

	p.Log().Debug("QLight handshake result for peer", "peer", p.id, "server", p.qlightServer, "psi", p.qlightPSI, "token", p.qlightToken)
	log.Info("QLight handshake result for peer", "peer", p.id, "server", p.qlightServer, "psi", p.qlightPSI, "token", p.qlightToken)
	// if we're not connected to a qlight server - disconnect the peer
	if !p.qlightServer {
		p.Log().Debug("QLight connected to a non server peer. Disconnecting.", "protoName", protoName)

		// Quorum
		// When the Handshake() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
		p.EthPeerDisconnected <- struct{}{}
		// End Quorum
		return fmt.Errorf("connected to a non server peer")
	}

	// Register the peer locally
	if err := pm.peers.Register(p, pm.removePeer, protoName); err != nil {
		p.Log().Error("Ethereum peer registration failed", "err", err)

		// Quorum
		// When the Register() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
		p.EthPeerDisconnected <- struct{}{}
		// End Quorum

		return err
	}
	defer pm.removePeer(p.id)

	// Register the peer in the downloader. If the downloader considers it banned, we disconnect
	if err := pm.downloader.RegisterPeer(p.id, p.version, p); err != nil {
		return err
	}
	pm.chainSync.handlePeerEvent(p)

	// Propagate existing transactions. new transactions appearing
	// after this will be sent via broadcasts.
	pm.syncTransactions(p)

	// If we have a trusted CHT, reject all peers below that (avoid fast sync eclipse)
	if pm.checkpointHash != (common.Hash{}) {
		// Request the peer's checkpoint header for chain height/weight validation
		if err := p.RequestHeadersByNumber(pm.checkpointNumber, 1, 0, false); err != nil {
			return err
		}
		// Start a timer to disconnect if the peer doesn't reply in time
		p.syncDrop = time.AfterFunc(syncChallengeTimeout, func() {
			p.Log().Warn("Checkpoint challenge timed out, dropping", "addr", p.RemoteAddr(), "type", p.Name())
			pm.removePeer(p.id)
		})
		// Make sure it's cleaned up if the peer dies off
		defer func() {
			if p.syncDrop != nil {
				p.syncDrop.Stop()
				p.syncDrop = nil
			}
		}()
	}
	// If we have any explicit authorized block hashes, request them
	for number := range pm.authorizationList {
		if err := p.RequestHeadersByNumber(number, 1, 0, false); err != nil {
			return err
		}
	}

	// Quorum notify other subprotocols that the eth peer is ready, and has been added to the peerset.
	p.EthPeerRegistered <- struct{}{}
	// Quorum

	// Handle incoming messages until the connection is torn down
	for {
		if err := pm.handleMsg(p); err != nil {
			p.Log().Debug("Ethereum message handling failed", "err", err)
			return err
		}
	}
}

// handleMsg is invoked whenever an inbound message is received from a remote
// peer. The remote connection is torn down upon returning any error.
func (pm *QLightClientProtocolManager) handleMsg(p *peer) error {
	// Read the next message from the remote peer, and ensure it's fully consumed
	msg, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Size > protocolMaxMsgSize {
		return errResp(ErrMsgTooLarge, "%v > %v", msg.Size, protocolMaxMsgSize)
	}
	defer msg.Discard()

	// Handle the message depending on its contents
	switch {
	case msg.Code == StatusMsg:
		// Status messages should never arrive after the handshake
		return errResp(ErrExtraStatusMsg, "uncontrolled status message")

	case msg.Code == BlockHeadersMsg:
		// A batch of headers arrived to one of our previous requests
		var headers []*types.Header
		if err := msg.Decode(&headers); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		// If no headers were received, but we're expencting a checkpoint header, consider it that
		if len(headers) == 0 && p.syncDrop != nil {
			// Stop the timer either way, decide later to drop or not
			p.syncDrop.Stop()
			p.syncDrop = nil

			// If we're doing a fast sync, we must enforce the checkpoint block to avoid
			// eclipse attacks. Unsynced nodes are welcome to connect after we're done
			// joining the network
			if atomic.LoadUint32(&pm.fastSync) == 1 {
				p.Log().Warn("Dropping unsynced node during fast sync", "addr", p.RemoteAddr(), "type", p.Name())
				return errors.New("unsynced node cannot serve fast sync")
			}
		}
		// Filter out any explicitly requested headers, deliver the rest to the downloader
		filter := len(headers) == 1
		if filter {
			// If it's a potential sync progress check, validate the content and advertised chain weight
			if p.syncDrop != nil && headers[0].Number.Uint64() == pm.checkpointNumber {
				// Disable the sync drop timer
				p.syncDrop.Stop()
				p.syncDrop = nil

				// Validate the header and either drop the peer or continue
				if headers[0].Hash() != pm.checkpointHash {
					return errors.New("checkpoint hash mismatch")
				}
				return nil
			}
			// Otherwise if it's a authorized block, validate against the set
			if want, ok := pm.authorizationList[headers[0].Number.Uint64()]; ok {
				if hash := headers[0].Hash(); want != hash {
					p.Log().Info("AuthorizationList mismatch, dropping peer", "number", headers[0].Number.Uint64(), "hash", hash, "want", want)
					return errors.New("authorizationList block mismatch")
				}
				p.Log().Debug("AuthorizationList block verified", "number", headers[0].Number.Uint64(), "hash", want)
			}
			// Irrelevant of the fork checks, send the header to the fetcher just in case
			headers = pm.blockFetcher.FilterHeaders(p.id, headers, time.Now())
		}
		if len(headers) > 0 || !filter {
			err := pm.downloader.DeliverHeaders(p.id, headers)
			if err != nil {
				log.Debug("Failed to deliver headers", "err", err)
			}
		}

	case msg.Code == BlockBodiesMsg:
		// A batch of block bodies arrived to one of our previous requests
		var request blockBodiesData
		if err := msg.Decode(&request); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		// Deliver them all to the downloader for queuing
		transactions := make([][]*types.Transaction, len(request))
		uncles := make([][]*types.Header, len(request))

		for i, body := range request {
			transactions[i] = body.Transactions
			pm.updateCacheWithNonPartyTxData(body.Transactions)
			uncles[i] = body.Uncles
		}
		// Filter out any explicitly requested bodies, deliver the rest to the downloader
		filter := len(transactions) > 0 || len(uncles) > 0
		if filter {
			transactions, uncles = pm.blockFetcher.FilterBodies(p.id, transactions, uncles, time.Now())
		}
		if len(transactions) > 0 || len(uncles) > 0 || !filter {
			err := pm.downloader.DeliverBodies(p.id, transactions, uncles)
			if err != nil {
				log.Debug("Failed to deliver bodies", "err", err)
			}
		}

	case p.version >= eth63 && msg.Code == ReceiptsMsg:
		// A batch of receipts arrived to one of our previous requests
		var receipts [][]*types.Receipt
		if err := msg.Decode(&receipts); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		// Deliver all to the downloader
		if err := pm.downloader.DeliverReceipts(p.id, receipts); err != nil {
			log.Debug("Failed to deliver receipts", "err", err)
		}

	case msg.Code == NewBlockHashesMsg:
		var announces newBlockHashesData
		if err := msg.Decode(&announces); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		// Mark the hashes as present at the remote node
		for _, block := range announces {
			p.MarkBlock(block.Hash)
		}
		// Schedule all the unknown hashes for retrieval
		unknown := make(newBlockHashesData, 0, len(announces))
		for _, block := range announces {
			if !pm.blockchain.HasBlock(block.Hash, block.Number) {
				unknown = append(unknown, block)
			}
		}
		for _, block := range unknown {
			pm.blockFetcher.Notify(p.id, block.Hash, block.Number, time.Now(), p.RequestOneHeader, p.RequestBodies)
		}

	case msg.Code == QLightNewBlockPrivateDataMsg:
		log.Info("Received new block private data")
		// Retrieve and decode the propagated block
		var request []engine.BlockPrivatePayloads
		if err := msg.Decode(&request); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		for _, blockPrivateData := range request {
			err := pm.privateClientCache.AddPrivateBlock(blockPrivateData)
			if err != nil {
				return errResp(ErrDecode, "%v: %v", msg, err)
			}
		}

	case msg.Code == NewBlockMsg:
		// Retrieve and decode the propagated block
		var request newBlockData
		if err := msg.Decode(&request); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		if hash := types.CalcUncleHash(request.Block.Uncles()); hash != request.Block.UncleHash() {
			log.Warn("Propagated block has invalid uncles", "have", hash, "exp", request.Block.UncleHash())
			break // TODO(karalabe): return error eventually, but wait a few releases
		}
		if hash := types.DeriveSha(request.Block.Transactions(), trie.NewStackTrie(nil)); hash != request.Block.TxHash() {
			log.Warn("Propagated block has invalid body", "have", hash, "exp", request.Block.TxHash())
			break // TODO(karalabe): return error eventually, but wait a few releases
		}
		if err := request.sanityCheck(); err != nil {
			return err
		}
		request.Block.ReceivedAt = msg.ReceivedAt
		request.Block.ReceivedFrom = p

		pm.updateCacheWithNonPartyTxData(request.Block.Transactions())

		// Mark the peer as owning the block and schedule it for import
		p.MarkBlock(request.Block.Hash())
		pm.blockFetcher.Enqueue(p.id, request.Block)

		// Assuming the block is importable by the peer, but possibly not yet done so,
		// calculate the head hash and TD that the peer truly must have.
		var (
			trueHead = request.Block.ParentHash()
			trueTD   = new(big.Int).Sub(request.TD, request.Block.Difficulty())
		)
		// Update the peer's total difficulty if better than the previous
		if _, td := p.Head(); trueTD.Cmp(td) > 0 {
			p.SetHead(trueHead, trueTD)
			pm.chainSync.handlePeerEvent(p)
		}

	case msg.Code == NewPooledTransactionHashesMsg && p.version >= eth65:
		// New transaction announcement arrived, make sure we have
		// a valid and fresh chain to handle them
		if atomic.LoadUint32(&pm.acceptTxs) == 0 {
			break
		}
		var hashes []common.Hash
		if err := msg.Decode(&hashes); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		// Schedule all the unknown hashes for retrieval
		for _, hash := range hashes {
			p.MarkTransaction(hash)
		}
		pm.txFetcher.Notify(p.id, hashes)

	case msg.Code == TransactionMsg || (msg.Code == PooledTransactionsMsg && p.version >= eth65):
		// Transactions arrived, make sure we have a valid and fresh chain to handle them
		if atomic.LoadUint32(&pm.acceptTxs) == 0 {
			break
		}
		// Transactions can be processed, parse all of them and deliver to the pool
		var txs []*types.Transaction
		if err := msg.Decode(&txs); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		for i, tx := range txs {
			// Validate and mark the remote transaction
			if tx == nil {
				return errResp(ErrDecode, "transaction %d is nil", i)
			}
			p.MarkTransaction(tx.Hash())
		}
		pm.txFetcher.Enqueue(p.id, txs, msg.Code == PooledTransactionsMsg)

	default:
		return errResp(ErrInvalidMsgCode, "%v", msg.Code)
	}
	return nil
}

func (pm *QLightClientProtocolManager) updateCacheWithNonPartyTxData(transactions types.Transactions) {
	for _, tx := range transactions {
		if tx.IsPrivate() || tx.IsPrivacyMarker() {
			txHash := common.BytesToEncryptedPayloadHash(tx.Data())
			pm.privateClientCache.CheckAndAddEmptyEntry(txHash)
		}
	}
}

// Quorum
func (pm *QLightClientProtocolManager) Enqueue(id string, block *types.Block) {
	pm.blockFetcher.Enqueue(id, block)
}

// BroadcastBlock is a no-op for QLightClientProtocolManager
func (pm *QLightClientProtocolManager) BroadcastBlock(_ *types.Block, _ bool) {
	// do nothing
}

// BroadcastTransactions is a no-op for QLightClientProtocolManager
func (pm *QLightClientProtocolManager) BroadcastTransactions(_ types.Transactions, _ bool) {
	// do nothing
}

// minedBroadcastLoop is a no-op for QLightClientProtocolManager
func (pm *QLightClientProtocolManager) minedBroadcastLoop() {
	// do nothing
}

// txBroadcastLoop is a no-op for QLightClientProtocolManager
func (pm *QLightClientProtocolManager) txBroadcastLoop() {
	// do nothing
}

// NodeInfo retrieves some protocol metadata about the running host node.
func (pm *QLightClientProtocolManager) NodeInfo() *NodeInfo {
	currentBlock := pm.blockchain.CurrentBlock()
	// //Quorum
	//
	// changes done to fetch maxCodeSize dynamically based on the
	// maxCodeSizeConfig changes
	// /Quorum
	chainConfig := pm.blockchain.Config()
	chainConfig.MaxCodeSize = uint64(chainConfig.GetMaxCodeSize(pm.blockchain.CurrentBlock().Number()) / 1024)

	return &NodeInfo{
		Network:    pm.networkID,
		Difficulty: pm.blockchain.GetTd(currentBlock.Hash(), currentBlock.NumberU64()),
		Genesis:    pm.blockchain.Genesis().Hash(),
		Config:     chainConfig,
		Head:       currentBlock.Hash(),
		Consensus:  "qlight",
	}
}

// Quorum
func (pm *QLightClientProtocolManager) getConsensusAlgorithm() string {
	var consensusAlgo string
	if pm.raftMode { // raft does not use consensus interface
		consensusAlgo = "raft"
	} else {
		switch pm.engine.(type) {
		case consensus.Istanbul:
			consensusAlgo = "istanbul"
		case *clique.Clique:
			consensusAlgo = "clique"
		case *ethash.Ethash:
			consensusAlgo = "ethash"
		default:
			consensusAlgo = "unknown"
		}
	}
	return consensusAlgo
}

func (self *QLightClientProtocolManager) FindPeers(targets map[common.Address]bool) map[common.Address]consensus.Peer {
	m := make(map[common.Address]consensus.Peer)
	for _, p := range self.peers.Peers() {
		pubKey := p.Node().Pubkey()
		addr := crypto.PubkeyToAddress(*pubKey)
		if targets[addr] {
			m[addr] = p
		}
	}
	return m
}

// End Quorum
