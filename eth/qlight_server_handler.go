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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/forkid"
	"github.com/ethereum/go-ethereum/core/mps"
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
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

type QLightServerProtocolManager struct {
	networkID  uint64
	forkFilter forkid.Filter // Fork ID filter, constant across the lifetime of the node

	fastSync  uint32 // Flag whether fast sync is enabled (gets disabled if we already have blocks)
	acceptTxs uint32 // Flag whether we're considered synchronised (enables transaction processing)

	checkpointNumber uint64      // Block number for the sync progress validator to cross reference
	checkpointHash   common.Hash // Block hash for the sync progress validator to cross reference

	txpool     txPool
	blockchain *core.BlockChain
	chaindb    ethdb.Database
	maxPeers   int

	downloader   *downloader.Downloader
	blockFetcher *fetcher.BlockFetcher
	txFetcher    *fetcher.TxFetcher
	peers        *peerSet

	eventMux      *event.TypeMux
	txsCh         chan core.NewTxsEvent
	txsSub        event.Subscription
	minedBlockSub *event.TypeMuxSubscription

	authorizationList map[uint64]common.Hash

	// channels for fetcher, syncer, txsyncLoop
	txsyncCh chan *txsync
	quitSync chan struct{}

	closeCh chan struct{}

	chainSync *chainSyncer
	wg        sync.WaitGroup
	peerWG    sync.WaitGroup

	// Quorum
	raftMode bool
	engine   consensus.Engine

	// Test fields or hooks
	broadcastTxAnnouncesOnly bool // Testing field, disable transaction propagation
}

// NewProtocolManager returns a new Ethereum sub protocol manager. The Ethereum sub protocol manages peers capable
// with the Ethereum network.
func NewQLightServerProtocolManager(config *params.ChainConfig, checkpoint *params.TrustedCheckpoint, mode downloader.SyncMode, networkID uint64, mux *event.TypeMux, txpool txPool, engine consensus.Engine, blockchain *core.BlockChain, chaindb ethdb.Database, cacheLimit int, authorizationList map[uint64]common.Hash, raftMode bool) (*QLightServerProtocolManager, error) {
	// Create the protocol manager with the base fields
	manager := &QLightServerProtocolManager{
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
	}

	fetchTx := func(peer string, hashes []common.Hash) error {
		p := manager.peers.Peer(peer)
		if p == nil {
			return errors.New("unknown peer")
		}
		return p.RequestTxs(hashes)
	}
	manager.txFetcher = fetcher.NewTxFetcher(txpool.Has, txpool.AddRemotes, fetchTx)

	return manager, nil
}

func (pm *QLightServerProtocolManager) makeProtocol(version uint) p2p.Protocol {
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
	}
}

func (pm *QLightServerProtocolManager) removePeer(id string) {
	// Short circuit if the peer was already removed
	log.Info("QLight removePeer", "id", id)
	peer := pm.peers.Peer(id)
	if peer == nil {
		return
	}
	log.Debug("Removing Ethereum peer", "peer", id)

	log.Info("QLight removePeer - before txFetcher.drop", "id", id)
	pm.txFetcher.Drop(id)
	log.Info("QLight removePeer - after txFetcher.drop", "id", id)

	if err := pm.peers.Unregister(id); err != nil {
		log.Error("Peer removal failed", "peer", id, "err", err)
	}
	log.Info("QLight removePeer - after peers.unregister", "id", id)
	// Hard disconnect at the networking layer
	if peer != nil {
		peer.Peer.Disconnect(p2p.DiscUselessPeer)
	}
	log.Info("QLight removePeer exit", "id", id)
}

func (pm *QLightServerProtocolManager) Start(maxPeers int) {
	pm.maxPeers = maxPeers

	// broadcast transactions
	pm.wg.Add(1)
	pm.txsCh = make(chan core.NewTxsEvent, txChanSize)
	pm.txsSub = pm.txpool.SubscribeNewTxsEvent(pm.txsCh)
	go pm.txBroadcastLoop()
	// qlight - manually start the tx fetcher
	pm.txFetcher.Start()

	// broadcast mined blocks
	pm.wg.Add(1)
	pm.closeCh = make(chan struct{})
	go pm.newBlockBroadcastLoop()

}

func (pm *QLightServerProtocolManager) Stop() {
	log.Info("QLight protocol stopping")
	pm.txsSub.Unsubscribe() // quits txBroadcastLoop
	log.Info("QLight protocol after txsSub.Unsubscribe() ")
	// qlight - manually stop the tx fetcher
	pm.txFetcher.Stop()
	log.Info("QLight protocol after txFetcher.Stop() ")

	// Quit chainSync and txsync64.
	// After this is done, no new peers will be accepted.
	close(pm.closeCh)
	pm.wg.Wait()

	log.Info("QLight protocol after WG.Wait()")

	// Disconnect existing sessions.
	// This also closes the gate for any new registrations on the peer set.
	// sessions which are already established but not added to pm.peers yet
	// will exit when they try to register.
	pm.peers.Close()
	pm.peerWG.Wait()

	log.Info("QLight protocol stopped")
}

func (pm *QLightServerProtocolManager) newPeer(pv int, p *p2p.Peer, rw p2p.MsgReadWriter, getPooledTx func(hash common.Hash) *types.Transaction) *peer {
	return newPeer(pv, p, rw, getPooledTx)
}

// Quorum - added protoName argument
func (pm *QLightServerProtocolManager) runPeer(p *peer, protoName string) error {
	//if !pm.chainSync.handlePeerEvent(p) {
	//	return p2p.DiscQuitting
	//}
	pm.peerWG.Add(1)
	defer pm.peerWG.Done()
	return pm.handle(p, protoName)
}

// quorum: protoname is either "eth" or a subprotocol that overrides "eth", e.g. legacy "istanbul/99"
// handle is the callback invoked to manage the life cycle of an eth peer. When
// this function terminates, the peer is disconnected.
func (pm *QLightServerProtocolManager) handle(p *peer, protoName string) error {
	// Ignore maxPeers if this is a trusted peer
	if pm.peers.Len() >= pm.maxPeers && !p.Peer.Info().Network.Trusted {
		return p2p.DiscTooManyPeers
	}
	p.Log().Debug("QLight peer connected", "name", p.Name())

	// Execute the Ethereum handshake
	var (
		genesis = pm.blockchain.Genesis()
		head    = pm.blockchain.CurrentHeader()
		hash    = head.Hash()
		number  = head.Number.Uint64()
		td      = pm.blockchain.GetTd(hash, number)
	)
	forkID := forkid.NewID(pm.blockchain.Config(), pm.blockchain.Genesis().Hash(), pm.blockchain.CurrentHeader().Number.Uint64())
	if err := p.Handshake(pm.networkID, td, hash, genesis.Hash(), forkID, pm.forkFilter, protoName); err != nil {
		p.Log().Debug("Ethereum handshake failed", "protoName", protoName, "err", err)

		// Quorum
		// When the Handshake() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
		p.EthPeerDisconnected <- struct{}{}
		// End Quorum
		return err
	}

	if err := p.QLightHandshake(true, "", ""); err != nil {
		p.Log().Debug("QLight handshake failed", "protoName", protoName, "err", err)

		// Quorum
		// When the Handshake() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
		p.EthPeerDisconnected <- struct{}{}
		// End Quorum
		return err
	}

	p.Log().Debug("QLight handshake result for peer", "peer", p.id, "server", p.qlightServer, "psi", p.qlightPSI, "token", p.qlightToken)

	if p.qlightServer {
		// Register the peer locally
		if err := pm.peers.RegisterIdlePeer(p, pm.removePeer, protoName); err != nil {
			p.Log().Error("Ethereum peer registration failed", "err", err)

			// Quorum
			// When the Register() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
			p.EthPeerDisconnected <- struct{}{}
			// End Quorum

			return err
		}
	} else {
		if err := pm.peers.Register(p, pm.removePeer, protoName); err != nil {
			p.Log().Error("Ethereum peer registration failed", "err", err)

			// Quorum
			// When the Register() returns an error, the Run method corresponding to `eth` protocol returns with the error, causing the peer to drop, signal subprotocol as well to exit the `Run` method
			p.EthPeerDisconnected <- struct{}{}
			// End Quorum

			return err
		}
	}
	defer pm.removePeer(p.id)

	// we're a qlight server connected to another server - do nothing
	if p.qlightServer {
		p.Log().Debug("QLight handshake to another server. Wait for remote disconnect", "protoName", protoName)

		// TODO Qlight - consider a server peers list (with it's own wait group)
		// connected to another server, no messages expected, just wait for disconnection
		msg, err := p.rw.ReadMsg()
		p.Log().Info("QLight - message received on server connection", "msg", msg, "err", err)
		return err
	}

	// Propagate existing transactions. new transactions appearing
	// after this will be sent via broadcasts.

	// TODO Qlight - see what to do about this
	//pm.syncTransactions(p)

	// Quorum notify other subprotocols that the eth peer is ready, and has been added to the peerset.
	p.EthPeerRegistered <- struct{}{}
	// Quorum

	// Handle incoming messages until the connection is torn down
	for {
		if err := pm.handleMsg(p); err != nil {
			log.Info("QLight handleMsg", "err", err)
			p.Log().Debug("QLight message handling failed", "err", err)
			return err
		}
	}
}

// handleMsg is invoked whenever an inbound message is received from a remote
// peer. The remote connection is torn down upon returning any error.
func (pm *QLightServerProtocolManager) handleMsg(p *peer) error {
	// Read the next message from the remote peer, and ensure it's fully consumed
	msg, err := p.rw.ReadMsg()
	log.Info("QLight message received", "err", err)
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

	// Block header query, collect the requested headers and reply
	case msg.Code == GetBlockHeadersMsg:
		// Decode the complex header query
		var query getBlockHeadersData
		if err := msg.Decode(&query); err != nil {
			return errResp(ErrDecode, "%v: %v", msg, err)
		}
		hashMode := query.Origin.Hash != (common.Hash{})
		first := true
		maxNonCanonical := uint64(100)

		// Gather headers until the fetch or network limits is reached
		var (
			bytes   common.StorageSize
			headers []*types.Header
			unknown bool
		)
		for !unknown && len(headers) < int(query.Amount) && bytes < softResponseLimit && len(headers) < downloader.MaxHeaderFetch {
			// Retrieve the next header satisfying the query
			var origin *types.Header
			if hashMode {
				if first {
					first = false
					origin = pm.blockchain.GetHeaderByHash(query.Origin.Hash)
					if origin != nil {
						query.Origin.Number = origin.Number.Uint64()
					}
				} else {
					origin = pm.blockchain.GetHeader(query.Origin.Hash, query.Origin.Number)
				}
			} else {
				origin = pm.blockchain.GetHeaderByNumber(query.Origin.Number)
			}
			if origin == nil {
				break
			}
			headers = append(headers, origin)
			bytes += estHeaderRlpSize

			// Advance to the next header of the query
			switch {
			case hashMode && query.Reverse:
				// Hash based traversal towards the genesis block
				ancestor := query.Skip + 1
				if ancestor == 0 {
					unknown = true
				} else {
					query.Origin.Hash, query.Origin.Number = pm.blockchain.GetAncestor(query.Origin.Hash, query.Origin.Number, ancestor, &maxNonCanonical)
					unknown = (query.Origin.Hash == common.Hash{})
				}
			case hashMode && !query.Reverse:
				// Hash based traversal towards the leaf block
				var (
					current = origin.Number.Uint64()
					next    = current + query.Skip + 1
				)
				if next <= current {
					infos, _ := json.MarshalIndent(p.Peer.Info(), "", "  ")
					p.Log().Warn("GetBlockHeaders skip overflow attack", "current", current, "skip", query.Skip, "next", next, "attacker", infos)
					unknown = true
				} else {
					if header := pm.blockchain.GetHeaderByNumber(next); header != nil {
						nextHash := header.Hash()
						expOldHash, _ := pm.blockchain.GetAncestor(nextHash, next, query.Skip+1, &maxNonCanonical)
						if expOldHash == query.Origin.Hash {
							query.Origin.Hash, query.Origin.Number = nextHash, next
						} else {
							unknown = true
						}
					} else {
						unknown = true
					}
				}
			case query.Reverse:
				// Number based traversal towards the genesis block
				if query.Origin.Number >= query.Skip+1 {
					query.Origin.Number -= query.Skip + 1
				} else {
					unknown = true
				}

			case !query.Reverse:
				// Number based traversal towards the leaf block
				query.Origin.Number += query.Skip + 1
			}
		}
		return p.SendBlockHeaders(headers)

	case msg.Code == BlockHeadersMsg:
		//// A batch of headers arrived to one of our previous requests
		//var headers []*types.Header
		//if err := msg.Decode(&headers); err != nil {
		//	return errResp(ErrDecode, "msg %v: %v", msg, err)
		//}
		//// If no headers were received, but we're expencting a checkpoint header, consider it that
		//if len(headers) == 0 && p.syncDrop != nil {
		//	// Stop the timer either way, decide later to drop or not
		//	p.syncDrop.Stop()
		//	p.syncDrop = nil
		//
		//	// If we're doing a fast sync, we must enforce the checkpoint block to avoid
		//	// eclipse attacks. Unsynced nodes are welcome to connect after we're done
		//	// joining the network
		//	if atomic.LoadUint32(&pm.fastSync) == 1 {
		//		p.Log().Warn("Dropping unsynced node during fast sync", "addr", p.RemoteAddr(), "type", p.Name())
		//		return errors.New("unsynced node cannot serve fast sync")
		//	}
		//}
		//// Filter out any explicitly requested headers, deliver the rest to the downloader
		//filter := len(headers) == 1
		//if filter {
		//	// If it's a potential sync progress check, validate the content and advertised chain weight
		//	if p.syncDrop != nil && headers[0].Number.Uint64() == pm.checkpointNumber {
		//		// Disable the sync drop timer
		//		p.syncDrop.Stop()
		//		p.syncDrop = nil
		//
		//		// Validate the header and either drop the peer or continue
		//		if headers[0].Hash() != pm.checkpointHash {
		//			return errors.New("checkpoint hash mismatch")
		//		}
		//		return nil
		//	}
		//	// Otherwise if it's a whitelisted block, validate against the set
		//	if want, ok := pm.authorizationList[headers[0].Number.Uint64()]; ok {
		//		if hash := headers[0].Hash(); want != hash {
		//			p.Log().Info("Whitelist mismatch, dropping peer", "number", headers[0].Number.Uint64(), "hash", hash, "want", want)
		//			return errors.New("authorizationList block mismatch")
		//		}
		//		p.Log().Debug("Whitelist block verified", "number", headers[0].Number.Uint64(), "hash", want)
		//	}
		//	// Irrelevant of the fork checks, send the header to the fetcher just in case
		//	headers = pm.blockFetcher.FilterHeaders(p.id, headers, time.Now())
		//}
		//if len(headers) > 0 || !filter {
		//	err := pm.downloader.DeliverHeaders(p.id, headers)
		//	if err != nil {
		//		log.Debug("Failed to deliver headers", "err", err)
		//	}
		//}

	case msg.Code == GetBlockBodiesMsg:
		// Decode the retrieval message
		msgStream := rlp.NewStream(msg.Payload, uint64(msg.Size))
		if _, err := msgStream.List(); err != nil {
			return err
		}
		// Gather blocks until the fetch or network limits is reached
		var (
			hash                    common.Hash
			bytes                   int
			bodies                  []rlp.RawValue
			privateTransactionsData PrivateTransactionsData
		)
		for bytes < softResponseLimit && len(bodies) < downloader.MaxBlockFetch {
			// Retrieve the hash of the next block
			if err := msgStream.Decode(&hash); err == rlp.EOL {
				break
			} else if err != nil {
				return errResp(ErrDecode, "msg %v: %v", msg, err)
			}
			// TODO Qlight - loading and parsing the block is a much heavier operation (than just loading the RLP encoded value)
			block := pm.blockchain.GetBlockByHash(hash)
			if block != nil {
				blockPTD, err := pm.preparePrivateTransactionsData(block, p.qlightPSI)
				if err != nil {
					return errResp(ErrDecode, "Unable to produce block private transaction data %v: %v", hash, err)
				}
				if blockPTD != nil {
					privateTransactionsData = append(privateTransactionsData, *blockPTD...)
				}
			}
			// Retrieve the requested block body, stopping if enough was found
			if data := pm.blockchain.GetBodyRLP(hash); len(data) != 0 {
				bodies = append(bodies, data)
				bytes += len(data)
			}
		}
		if len(privateTransactionsData) > 0 {
			err := p2p.Send(p.rw, QLightNewBlockPrivateDataMsg, privateTransactionsData)
			if err != nil {
				log.Info("Error occurred while sending private data msg", "err", err)
				return err
			}
		}
		return p.SendBlockBodiesRLP(bodies)

	case msg.Code == BlockBodiesMsg:
		//// A batch of block bodies arrived to one of our previous requests
		//var request blockBodiesData
		//if err := msg.Decode(&request); err != nil {
		//	return errResp(ErrDecode, "msg %v: %v", msg, err)
		//}
		//// Deliver them all to the downloader for queuing
		//transactions := make([][]*types.Transaction, len(request))
		//uncles := make([][]*types.Header, len(request))
		//
		//for i, body := range request {
		//	transactions[i] = body.Transactions
		//	uncles[i] = body.Uncles
		//}
		//// Filter out any explicitly requested bodies, deliver the rest to the downloader
		//filter := len(transactions) > 0 || len(uncles) > 0
		//if filter {
		//	transactions, uncles = pm.blockFetcher.FilterBodies(p.id, transactions, uncles, time.Now())
		//}
		//if len(transactions) > 0 || len(uncles) > 0 || !filter {
		//	err := pm.downloader.DeliverBodies(p.id, transactions, uncles)
		//	if err != nil {
		//		log.Debug("Failed to deliver bodies", "err", err)
		//	}
		//}

	case p.version >= eth63 && msg.Code == GetNodeDataMsg:
		// Decode the retrieval message
		msgStream := rlp.NewStream(msg.Payload, uint64(msg.Size))
		if _, err := msgStream.List(); err != nil {
			return err
		}
		// Gather state data until the fetch or network limits is reached
		var (
			hash  common.Hash
			bytes int
			data  [][]byte
		)
		for bytes < softResponseLimit && len(data) < downloader.MaxStateFetch {
			// Retrieve the hash of the next state entry
			if err := msgStream.Decode(&hash); err == rlp.EOL {
				break
			} else if err != nil {
				return errResp(ErrDecode, "msg %v: %v", msg, err)
			}
			// Retrieve the requested state entry, stopping if enough was found
			// todo now the code and trienode is mixed in the protocol level,
			// separate these two types.
			if !pm.downloader.SyncBloomContains(hash[:]) {
				// Only lookup the trie node if there's chance that we actually have it
				continue
			}
			entry, err := pm.blockchain.TrieNode(hash)
			if len(entry) == 0 || err != nil {
				// Read the contract code with prefix only to save unnecessary lookups.
				entry, err = pm.blockchain.ContractCodeWithPrefix(hash)
			}
			if err == nil && len(entry) > 0 {
				data = append(data, entry)
				bytes += len(entry)
			}
		}
		return p.SendNodeData(data)

	case p.version >= eth63 && msg.Code == NodeDataMsg:
		// A batch of node state data arrived to one of our previous requests
		var data [][]byte
		if err := msg.Decode(&data); err != nil {
			return errResp(ErrDecode, "msg %v: %v", msg, err)
		}
		// Deliver all to the downloader
		if err := pm.downloader.DeliverNodeData(p.id, data); err != nil {
			log.Debug("Failed to deliver node state data", "err", err)
		}

	case p.version >= eth63 && msg.Code == GetReceiptsMsg:
		// Decode the retrieval message
		msgStream := rlp.NewStream(msg.Payload, uint64(msg.Size))
		if _, err := msgStream.List(); err != nil {
			return err
		}
		// Gather state data until the fetch or network limits is reached
		var (
			hash     common.Hash
			bytes    int
			receipts []rlp.RawValue
		)
		for bytes < softResponseLimit && len(receipts) < downloader.MaxReceiptFetch {
			// Retrieve the hash of the next block
			if err := msgStream.Decode(&hash); err == rlp.EOL {
				break
			} else if err != nil {
				return errResp(ErrDecode, "msg %v: %v", msg, err)
			}
			// Retrieve the requested block's receipts, skipping if unknown to us
			results := pm.blockchain.GetReceiptsByHash(hash)
			if results == nil {
				if header := pm.blockchain.GetHeaderByHash(hash); header == nil || header.ReceiptHash != types.EmptyRootHash {
					continue
				}
			}
			// If known, encode and queue for response packet
			if encoded, err := rlp.EncodeToBytes(results); err != nil {
				log.Error("Failed to encode receipt", "err", err)
			} else {
				receipts = append(receipts, encoded)
				bytes += len(encoded)
			}
		}
		return p.SendReceiptsRLP(receipts)

	case p.version >= eth63 && msg.Code == ReceiptsMsg:
		// A batch of receipts arrived to one of our previous requests
		//var receipts [][]*types.Receipt
		//if err := msg.Decode(&receipts); err != nil {
		//	return errResp(ErrDecode, "msg %v: %v", msg, err)
		//}
		//// Deliver all to the downloader
		//if err := pm.downloader.DeliverReceipts(p.id, receipts); err != nil {
		//	log.Debug("Failed to deliver receipts", "err", err)
		//}

	case msg.Code == NewBlockHashesMsg:
		//var announces newBlockHashesData
		//if err := msg.Decode(&announces); err != nil {
		//	return errResp(ErrDecode, "%v: %v", msg, err)
		//}
		//// Mark the hashes as present at the remote node
		//for _, block := range announces {
		//	p.MarkBlock(block.Hash)
		//}
		//// Schedule all the unknown hashes for retrieval
		//unknown := make(newBlockHashesData, 0, len(announces))
		//for _, block := range announces {
		//	if !pm.blockchain.HasBlock(block.Hash, block.Number) {
		//		unknown = append(unknown, block)
		//	}
		//}
		//for _, block := range unknown {
		//	pm.blockFetcher.Notify(p.id, block.Hash, block.Number, time.Now(), p.RequestOneHeader, p.RequestBodies)
		//}

	case msg.Code == NewBlockMsg:
		// Retrieve and decode the propagated block
		//var request newBlockData
		//if err := msg.Decode(&request); err != nil {
		//	return errResp(ErrDecode, "%v: %v", msg, err)
		//}
		//if hash := types.CalcUncleHash(request.Block.Uncles()); hash != request.Block.UncleHash() {
		//	log.Warn("Propagated block has invalid uncles", "have", hash, "exp", request.Block.UncleHash())
		//	break // TODO(karalabe): return error eventually, but wait a few releases
		//}
		//if hash := types.DeriveSha(request.Block.Transactions(), trie.NewStackTrie(nil)); hash != request.Block.TxHash() {
		//	log.Warn("Propagated block has invalid body", "have", hash, "exp", request.Block.TxHash())
		//	break // TODO(karalabe): return error eventually, but wait a few releases
		//}
		//if err := request.sanityCheck(); err != nil {
		//	return err
		//}
		//request.Block.ReceivedAt = msg.ReceivedAt
		//request.Block.ReceivedFrom = p
		//
		//// Mark the peer as owning the block and schedule it for import
		//p.MarkBlock(request.Block.Hash())
		//pm.blockFetcher.Enqueue(p.id, request.Block)
		//
		//// Assuming the block is importable by the peer, but possibly not yet done so,
		//// calculate the head hash and TD that the peer truly must have.
		//var (
		//	trueHead = request.Block.ParentHash()
		//	trueTD   = new(big.Int).Sub(request.TD, request.Block.Difficulty())
		//)
		//// Update the peer's total difficulty if better than the previous
		//if _, td := p.Head(); trueTD.Cmp(td) > 0 {
		//	p.SetHead(trueHead, trueTD)
		//	pm.chainSync.handlePeerEvent(p)
		//}

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

	case msg.Code == GetPooledTransactionsMsg && p.version >= eth65:
		// Decode the retrieval message
		msgStream := rlp.NewStream(msg.Payload, uint64(msg.Size))
		if _, err := msgStream.List(); err != nil {
			return err
		}
		// Gather transactions until the fetch or network limits is reached
		var (
			hash   common.Hash
			bytes  int
			hashes []common.Hash
			txs    []rlp.RawValue
		)
		for bytes < softResponseLimit {
			// Retrieve the hash of the next block
			if err := msgStream.Decode(&hash); err == rlp.EOL {
				break
			} else if err != nil {
				return errResp(ErrDecode, "msg %v: %v", msg, err)
			}
			// Retrieve the requested transaction, skipping if unknown to us
			tx := pm.txpool.Get(hash)
			if tx == nil {
				continue
			}
			// If known, encode and queue for response packet
			if encoded, err := rlp.EncodeToBytes(tx); err != nil {
				log.Error("Failed to encode transaction", "err", err)
			} else {
				hashes = append(hashes, hash)
				txs = append(txs, encoded)
				bytes += len(encoded)
			}
		}
		return p.SendPooledTransactionsRLP(hashes, txs)

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

// BroadcastBlock will either propagate a block to a subset of its peers, or
// will only announce its availability (depending what's requested).
func (pm *QLightServerProtocolManager) BroadcastBlock(block *types.Block, propagate bool) {
	hash := block.Hash()
	peers := pm.peers.PeersWithoutBlock(hash)

	// If propagation is requested, send to a subset of the peer
	if propagate {
		// Calculate the TD of the block (it's not imported yet, so block.Td is not valid)
		var td *big.Int
		if parent := pm.blockchain.GetBlock(block.ParentHash(), block.NumberU64()-1); parent != nil {
			td = new(big.Int).Add(block.Difficulty(), pm.blockchain.GetTd(block.ParentHash(), block.NumberU64()-1))
		} else {
			log.Error("Propagating dangling block", "number", block.Number(), "hash", hash)
			return
		}
		// Send the block to all the peers
		log.Info("Preparing new block private data")
		for _, peer := range peers {
			if peer.qlightServer {
				continue
			}
			privateTransactionsData, err := pm.preparePrivateTransactionsData(block, peer.qlightPSI)
			if err != nil {
				log.Error("Unable to prepare privateTransactionsData for block", "number", block.Number(), "hash", hash, "err", err, "psi", peer.qlightPSI)
				return
			}
			log.Info("Private transactions data", "is nil", privateTransactionsData == nil)
			peer.AsyncSendNewBlock(block, td, privateTransactionsData)
		}
		log.Trace("Propagated block", "hash", hash, "recipients", len(peers), "duration", common.PrettyDuration(time.Since(block.ReceivedAt)))
		return
	}
	// Otherwise if the block is indeed in out own chain, announce it
	if pm.blockchain.HasBlock(hash, block.NumberU64()) {
		for _, peer := range peers {
			peer.AsyncSendNewBlockHash(block)
		}
		log.Trace("Announced block", "hash", hash, "recipients", len(peers), "duration", common.PrettyDuration(time.Since(block.ReceivedAt)))
	}
}

func (pm *QLightServerProtocolManager) preparePrivateTransactionsData(block *types.Block, psi string) (*PrivateTransactionsData, error) {
	// TODO qlight - this can probably be replaced with loading prepared data (if the block processing is updated to store privateTransactionsData structures / PSI)
	PSI := types.PrivateStateIdentifier(psi)
	result := make(PrivateTransactionsData, 0)
	psm, err := pm.blockchain.PrivateStateManager().ResolveForUserContext(rpc.WithPrivateStateIdentifier(context.Background(), PSI))
	if err != nil {
		return nil, err
	}
	for _, tx := range block.Transactions() {
		if tx.IsPrivacyMarker() {
			_, result, err = pm.fetchPrivateData(tx.Data(), psm, result)
			if err != nil {
				return nil, err
			}

			innerTx, _, _, _ := private.FetchPrivateTransaction(tx.Data())
			if innerTx != nil {
				tx = innerTx
			}
		}

		if tx.IsPrivate() {
			_, result, err = pm.fetchPrivateData(tx.Data(), psm, result)
			if err != nil {
				return nil, err
			}
		}
	}
	if len(result) > 0 {
		return &result, nil
	}
	return nil, nil
}

func (pm *QLightServerProtocolManager) fetchPrivateData(privateData []byte, psm *mps.PrivateStateMetadata, result PrivateTransactionsData) (*PrivateTransactionData, PrivateTransactionsData, error) {
	txHash := common.BytesToEncryptedPayloadHash(privateData)
	_, _, privateTx, extra, err := private.P.Receive(txHash)
	if err != nil {
		return nil, nil, err
	}
	// we're not party to this transaction
	if privateTx == nil {
		return nil, result, nil
	}
	if pm.blockchain.PrivateStateManager().NotIncludeAny(psm, extra.ManagedParties...) {
		return nil, result, nil
	}

	extra.ManagedParties = psm.FilterAddresses(extra.ManagedParties...)

	ptd := PrivateTransactionData{
		Hash:     &txHash,
		Payload:  privateTx,
		Extra:    extra,
		IsSender: false,
	}
	if len(psm.Addresses) == 0 {
		// this is not an MPS node so we have to ask tessera
		ptd.IsSender, _ = private.P.IsSender(txHash)
	} else {
		// this is an MPS node so we can speed up the IsSender logic by checking the addresses in the private state metadata
		ptd.IsSender = !psm.NotIncludeAny(extra.Sender)
	}
	result = append(result, ptd)
	return &ptd, result, nil
}

// BroadcastTransactions will propagate a batch of transactions to all peers which are not known to
// already have the given transaction.
func (pm *QLightServerProtocolManager) BroadcastTransactions(txs types.Transactions, propagate bool) {
	var (
		txset = make(map[*peer][]common.Hash)
		annos = make(map[*peer][]common.Hash)
	)
	// Broadcast transactions to a batch of peers not knowing about it
	// NOTE: Raft-based consensus currently assumes that geth broadcasts
	// transactions to all peers in the network. A previous comment here
	// indicated that this logic might change in the future to only send to a
	// subset of peers. If this change occurs upstream, a merge conflict should
	// arise here, and we should add logic to send to *all* peers in raft mode.

	if propagate {
		for _, tx := range txs {
			peers := pm.peers.PeersWithoutTx(tx.Hash())

			// Send the block to a subset of our peers
			for _, peer := range peers {
				if peer.qlightServer {
					continue
				}
				txset[peer] = append(txset[peer], tx.Hash())
			}
			log.Trace("Broadcast transaction", "hash", tx.Hash(), "recipients", len(peers))
		}
		for peer, hashes := range txset {
			peer.AsyncSendTransactions(hashes)
		}
		return
	}
	// Otherwise only broadcast the announcement to peers
	for _, tx := range txs {
		peers := pm.peers.PeersWithoutTx(tx.Hash())
		for _, peer := range peers {
			if peer.qlightServer {
				continue
			}
			annos[peer] = append(annos[peer], tx.Hash())
		}
	}
	for peer, hashes := range annos {
		if peer.version >= eth65 {
			peer.AsyncSendPooledTransactionHashes(hashes)
		} else {
			peer.AsyncSendTransactions(hashes)
		}
	}
}

// newBlockBroadcastLoop sends mined blocks to connected peers.
func (pm *QLightServerProtocolManager) newBlockBroadcastLoop() {
	defer pm.wg.Done()

	headCh := make(chan core.ChainHeadEvent, 10)
	headSub := pm.blockchain.SubscribeChainHeadEvent(headCh)
	defer headSub.Unsubscribe()

	for {
		select {
		case ev := <-headCh:
			log.Debug("Announcing block to peers", "number", ev.Block.Number(), "hash", ev.Block.Hash(), "td", ev.Block.Difficulty())
			pm.BroadcastBlock(ev.Block, true)

		case <-pm.closeCh:
			return
		}
	}
}

// txBroadcastLoop announces new transactions to connected peers.
func (pm *QLightServerProtocolManager) txBroadcastLoop() {
	defer pm.wg.Done()

	for {
		select {
		case event := <-pm.txsCh:
			// For testing purpose only, disable propagation
			if pm.broadcastTxAnnouncesOnly {
				pm.BroadcastTransactions(event.Txs, false)
				continue
			}
			pm.BroadcastTransactions(event.Txs, true) // First propagate transactions to peers
			//			pm.BroadcastTransactions(event.Txs, false) // Only then announce to the rest

		case <-pm.txsSub.Err():
			return
		}
	}
}

// NodeInfo retrieves some protocol metadata about the running host node.
func (pm *QLightServerProtocolManager) NodeInfo() *NodeInfo {
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

func (self *QLightServerProtocolManager) FindPeers(targets map[common.Address]bool) map[common.Address]consensus.Peer {
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
