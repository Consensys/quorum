// Overview of the channels used in this module:
//
// Node.
// * quitSync: *Every* channel operation can be unblocked by closing this
//   channel.
//
// ProtocolManager.
// * logCommandC, for committed raft entries and commands to take or load a
//   snapshot, flowing to ethereum.
// * proposeC, for proposals flowing from ethereum to raft
// * confChangeC, for config changes flowing from ethereum to raft
// * peerMsgC, for messages coming from peers, to be dumped into raft
// * roleC, coming from raft notifies us when our role changes
package gethRaft

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/coreos/etcd/pkg/fileutil"
	"github.com/coreos/etcd/snap"
	"github.com/coreos/etcd/wal"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
)

// Messages we send on the logCommandC channel
type LoadSnapshot struct{}

type ProtocolManager struct {
	// peers note -- each node tracks both:
	// * the peers it knows of from discovery
	// * the peers acknowledged by raft
	//
	// only the leader proposes `ConfChangeAddNode` for each peer in the first set
	// but not in the second. this is done:
	// * when a node becomes leader
	// * when the leader learns of new peers

	// This node's rlpx (enode) id
	id string

	// set of rlpx-discovered peers
	rlpxKnownPeers map[string]*peer

	// set of currently active peers known to the raft cluster. this includes self
	raftKnownPeers map[uint64]*raft.Peer

	protocol p2p.Protocol

	blockchain *core.BlockChain

	// to protect the (rlpx and raft) peer maps
	mu sync.RWMutex

	eventMux      *event.TypeMux
	minedBlockSub event.Subscription

	downloader *downloader.Downloader
	peerGetter func() (string, *big.Int)

	rawNode     raft.Node
	raftStorage *raft.MemoryStorage

	// The number of entries applied to the raft log
	appliedIndex uint64

	// The index of the latest snapshot.
	snapshotIndex uint64

	// snapshotting
	snapshotter *snap.Snapshotter
	snapdir     string
	confState   raftpb.ConfState

	// write-ahead log
	waldir string
	wal    *wal.WAL

	peerMsgC    chan p2p.Msg
	proposeC    chan *types.Block
	confChangeC chan raftpb.ConfChange
	// messages committed by raft (right now these are the messages committed
	// right when raft starts)
	logCommandC chan interface{}
	quitSync    chan struct{}

	// Note: we don't actually use this field. We just set it at the same time as
	// starting or stopping the miner in notifyRoleChange. We might want to remove
	// it, but it might also be useful to check.
	role int
}

// Implement the `MsgWriter` interface (necessary for p2p.Send)
func (pm *ProtocolManager) WriteMsg(msg p2p.Msg) error {
	// read *into* buffer
	var buffer = make([]byte, msg.Size)
	msg.Payload.Read(buffer)

	return pm.rawNode.Propose(context.TODO(), buffer)
}

func StartRaftNode(pm *ProtocolManager, storage *raft.MemoryStorage, startPeers []raft.Peer) {
	if !fileutil.Exist(pm.snapdir) {
		if err := os.Mkdir(pm.snapdir, 0750); err != nil {
			log.Fatalf("raftexample: cannot create dir for snapshot (%v)", err)
		}
	}
	pm.snapshotter = snap.New(pm.snapdir)

	oldwal := wal.Exist(pm.waldir)

	// wal deallocated in eventLoop
	pm.wal = pm.replayWAL()

	c := &raft.Config{
		ID: strToIntID(pm.id),
		// TODO(joel): tune these parameters
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}

	if oldwal {
		pm.rawNode = raft.RestartNode(c)
	} else {
		pm.rawNode = raft.StartNode(c, startPeers)
	}

	go pm.serveInternal(pm.proposeC, pm.confChangeC)
	go pm.eventLoop(pm.logCommandC)
	go pm.handlePeerMsgs(pm.peerMsgC)
}

func (pm *ProtocolManager) stop() {
	close(pm.quitSync)
	if pm.rawNode != nil {
		pm.rawNode.Stop()
	}
}

func (pm *ProtocolManager) minedBroadcastLoop(proposeC chan<- *types.Block) {
	for obj := range pm.minedBlockSub.Chan() {
		switch ev := obj.Data.(type) {
		case core.NewMinedBlockEvent:
			select {
			case proposeC <- ev.Block:
			case <-pm.quitSync:
				return
			}
		}
	}
}

// serve two channels (proposeC, confChangeC) to handle changes originating
// internally
func (pm *ProtocolManager) serveInternal(proposeC <-chan *types.Block, confChangeC <-chan raftpb.ConfChange) {
	var confChangeCount uint64

	for {
		select {
		case block, ok := <-proposeC:
			if !ok {
				log.Println("error: read from proposeC failed")
				return
			}

			size, r, err := rlp.EncodeToReader(block)
			if err != nil {
				panic(fmt.Sprintf("error: failed to send RLP-encoded block: %s", err.Error()))
			}
			var buffer = make([]byte, uint32(size))
			r.Read(buffer)

			// blocks until accepted by the raft state machine
			pm.rawNode.Propose(context.TODO(), buffer)
		case cc, ok := <-confChangeC:
			if !ok {
				log.Println("error: read from confChangeC failed")
				return
			}

			confChangeCount++
			cc.ID = confChangeCount
			pm.rawNode.ProposeConfChange(context.TODO(), cc)
		case <-pm.quitSync:
			return
		}
	}
}

func strToIntID(strID string) uint64 {
	// take 64 bits
	intID, err := strconv.ParseUint(strID[:16], 16, 64)
	if err != nil {
		log.Fatalf("Failed to parse string id: %v", err)
	}
	return intID
}

func (pm *ProtocolManager) eventLoop(logCommandC chan<- interface{}) {
	snap, err := pm.raftStorage.Snapshot()
	if err != nil {
		panic(err)
	}
	pm.confState = snap.Metadata.ConfState
	pm.snapshotIndex = snap.Metadata.Index
	pm.appliedIndex = snap.Metadata.Index

	ticker := time.NewTicker(tickerMS * time.Millisecond)
	defer ticker.Stop()
	defer pm.wal.Close()

	for {
		select {
		case <-ticker.C:
			pm.rawNode.Tick()

		// when the node is first ready it gives us entries to commit and messages
		// to immediately publish
		case rd := <-pm.rawNode.Ready():
			pm.wal.Save(rd.HardState, rd.Entries)
			if !raft.IsEmptySnap(rd.Snapshot) {
				pm.saveSnap(rd.Snapshot)
				pm.raftStorage.ApplySnapshot(rd.Snapshot)
				pm.publishSnapshot(rd.Snapshot)
			}

			// 1: Write HardState, Entries, and Snapshot to persistent storage if they
			// are not empty.
			pm.raftStorage.Append(rd.Entries)

			// 2: Send all Messages to the nodes named in the To field.
			pm.SendToPeers(rd.Messages)

			// 3: Apply Snapshot (if any) and CommittedEntries to the state machine.
			for _, entry := range rd.CommittedEntries {
				switch entry.Type {
				case raftpb.EntryNormal:
					if len(entry.Data) == 0 {
						break
					}
					var block types.Block
					err := rlp.DecodeBytes(entry.Data, &block)
					if err != nil {
						log.Println("error decoding block: ", err)
					}
					select {
					case logCommandC <- &block:
					case <-pm.quitSync:
						return
					}
				case raftpb.EntryConfChange:
					var cc raftpb.ConfChange
					cc.Unmarshal(entry.Data)
					pm.rawNode.ApplyConfChange(cc)
					pm.mu.Lock()
					switch cc.Type {
					case raftpb.ConfChangeAddNode:
						p := &raft.Peer{
							ID: cc.NodeID,
							// We use the context to hold the enode id
							Context: cc.Context,
						}
						pm.raftKnownPeers[cc.NodeID] = p
						// if _, ok := pm.protocolManager.rlpxKnownPeers[string(p.Context)]; !ok {
						// 	// TODO would be cool if we could hint to rlpx to look
						// 	// for a new peer
						// }

					case raftpb.ConfChangeRemoveNode:
						if cc.NodeID == strToIntID(pm.id) {
							glog.V(logger.Info).Infoln("I've been removed from the cluster -- shutting down!")
							return
						}
						delete(pm.raftKnownPeers, cc.NodeID)
					}

					pm.appliedIndex = entry.Index

					pm.mu.Unlock()
				}
			}

			// 4: Call Node.Advance() to signal readiness for the next batch of
			// updates.
			pm.maybeTriggerSnapshot()
			pm.rawNode.Advance()

		case <-pm.quitSync:
			return
		}
	}
}

// Protocol Manager

func NewProtocolManager(strID string, blockchain *core.BlockChain, mux *event.TypeMux, downloader *downloader.Downloader, peerGetter func() (string, *big.Int)) (*ProtocolManager, error) {
	waldir := fmt.Sprintf("raft-%s-wal", strID)
	snapdir := fmt.Sprintf("raft-%s-snap", strID)

	manager := &ProtocolManager{
		rlpxKnownPeers: make(map[string]*peer),
		raftKnownPeers: make(map[uint64]*raft.Peer),
		blockchain:     blockchain,
		eventMux:       mux,
		downloader:     downloader,
		peerGetter:     peerGetter,
		peerMsgC:       make(chan p2p.Msg, msgChanSize),
		logCommandC:    make(chan interface{}),
		proposeC:       make(chan *types.Block),
		confChangeC:    make(chan raftpb.ConfChange),

		waldir:      waldir,
		snapdir:     snapdir,
		snapshotter: nil, // snap.New(snapdir),
		id:          strID,
		quitSync:    make(chan struct{}),
		raftStorage: raft.NewMemoryStorage(),
	}

	manager.protocol = p2p.Protocol{
		Name:    protocolName,
		Version: uint(protocolVersion),

		// number of message codes used
		Length: 1,

		NodeInfo: func() interface{} {
			return manager.NodeInfo()
		},

		PeerInfo: func(id discover.NodeID) interface{} {
			manager.mu.RLock()
			defer manager.mu.RUnlock()

			if p, ok := manager.rlpxKnownPeers[id.String()]; ok {
				return p.Info()
			}

			return nil
		},

		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			peer := newPeer(p, rw)
			return manager.handleRlpxPeerDiscovery(peer, manager.confChangeC, manager.peerMsgC)
		},
	}

	go manager.handleLogCommands(manager.logCommandC)

	return manager, nil
}

func (pm *ProtocolManager) SendToPeers(messages []raftpb.Message) {
	for _, m := range messages {
		if m.To == 0 {
			// ignore intentionally dropped message
			continue
		}

		pm.mu.RLock()
		raftPeer, ok := pm.raftKnownPeers[m.To]
		if ok {
			rlpxName := string(raftPeer.Context)
			ethPeer, ok := pm.rlpxKnownPeers[rlpxName]

			if ok {
				defer ethPeer.SendRaftPB(m)
			} else {
				glog.V(logger.Error).Infof(
					"Ignoring %v sent to unknown p2p peer: %v\n", m.Type, rlpxName)
			}
		} else {
			glog.V(logger.Error).Infof(
				"Ignoring %v sent to unknown raft peer: %v\n", m.Type, m.To)
		}
		pm.mu.RUnlock()

	}
}

func (pm *ProtocolManager) NodeInfo() *RaftNodeInfo {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	var roleDescription string
	if pm.role == minterRole {
		roleDescription = "minter"
	} else {
		roleDescription = "verifier"
	}

	return &RaftNodeInfo{
		ClusterSize: len(pm.raftKnownPeers),
		Genesis:     pm.blockchain.Genesis().Hash(),
		Head:        pm.blockchain.CurrentBlock().Hash(),
		Role:        roleDescription,
	}
}

func (pm *ProtocolManager) Start() {
	pm.minedBlockSub = pm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	go pm.minedBroadcastLoop(pm.proposeC)
}

func (pm *ProtocolManager) Stop() {
	glog.V(logger.Info).Infoln("Stopping ethereum protocol handler...")

	pm.stop()

	glog.V(logger.Info).Infoln("Ethereum protocol handler stopped")
}

func logCheckpoint(checkpointName string, iface interface{}) {
	log.Printf("RAFT-CHECKPOINT %s %v\n", checkpointName, iface)
}

// manage the lifecycle of a peer. the peer is disconnected when this function
// terminates.
func (pm *ProtocolManager) handleRlpxPeerDiscovery(p *peer, confChangeC chan<- raftpb.ConfChange, peerMsgC chan<- p2p.Msg) error {
	glog.V(logger.Debug).Infof("%v: peer connected [%s]", p, p.strID)
	logCheckpoint("PEER-CONNECTED", p.uint64Id)

	pm.mu.Lock()
	pm.rlpxKnownPeers[p.strID] = p
	pm.mu.Unlock()

	if raftRunning := pm.rawNode != nil; raftRunning {
		cc := raftpb.ConfChange{
			Type:    raftpb.ConfChangeAddNode,
			NodeID:  p.uint64Id,
			Context: []byte(p.strID),
		}

		// TODO(bts): only propose conf change if our raft conf doesn't include this node
		select {
		case <-pm.quitSync:
			return nil
		case confChangeC <- cc:
		}
	}

	defer pm.removeRlpxPeer(p.strID)

	// incoming message loop
	for {
		msg, err := p.rw.ReadMsg()
		if err != nil {
			return err
		}
		select {
		case peerMsgC <- msg:
		case <-pm.quitSync:
			return nil
		}
	}
}

func (pm *ProtocolManager) handlePeerMsgs(peerMsgC <-chan p2p.Msg) error {
	for {
		select {
		case msg := <-peerMsgC:
			defer msg.Discard()

			decoded := make([]byte, msg.Size)
			if err := msg.Decode(&decoded); err != nil {
				return err
			}

			var pbDecoded raftpb.Message
			if err := pbDecoded.Unmarshal(decoded); err != nil {
				return err
			}

			pm.rawNode.Step(context.TODO(), pbDecoded)
		case <-pm.quitSync:
			return nil
		}
	}
}

func blockExtendsChain(block *types.Block, chain *core.BlockChain) bool {
	return block.ParentHash() == chain.CurrentBlock().Hash()
}

func (pm *ProtocolManager) handleLogCommands(logCommandC <-chan interface{}) {
	for {
		select {
		case iface := <-logCommandC:

			//
			// TODO(bts): we need to keep track of what we've applied in case we crash. don't replay everything
			//

			//
			// TODO(joel): make sure snapshotting/compaction is consistent. i.e. no extra blocks
			//

			switch cmd := iface.(type) {
			case *types.Block:
				block := cmd
				if !blockExtendsChain(block, pm.blockchain) {
					headBlock := pm.blockchain.CurrentBlock()

					glog.V(logger.Warn).Infof("Non-extending block: %x (parent is %x; current head is %x)\n", block.Hash(), block.ParentHash(), headBlock.Hash())

					pm.eventMux.Post(InvalidRaftOrdering{headBlock: headBlock, invalidBlock: block})
				} else {
					if existingBlock := pm.blockchain.GetBlockByHash(block.Hash()); nil == existingBlock {
						if err := pm.blockchain.Validator().ValidateBlock(block); err != nil {
							panic(fmt.Sprintf("failed to validate block %x (%v)", block.Hash(), err))
						}
					}

					if pm.blockchain.HasBlock(block.Hash()) {
						// This node mined the block, so it was already in the
						// DB. We simply extend the chain:
						pm.blockchain.SetNewHeadBlock(block)
					} else {
						//
						// This will broadcast a CHE *almost always*. It does its
						// broadcasting at the end in a goroutine, but only conditionally if
						// the chain head is in a certain state. For now, we will broadcast
						// a CHE ourselves below to guarantee correctness.
						//
						_, err := pm.blockchain.InsertChain([]*types.Block{block})

						if err != nil {
							panic(fmt.Sprintf("failed to extend chain: %s", err.Error()))
						}
					}

					pm.eventMux.Post(core.ChainHeadEvent{Block: block})
					glog.V(logger.Info).Infof("Successfully extended chain: %x\n", block.Hash())
				}
			case LoadSnapshot:
				snapshot, err := pm.snapshotter.Load()
				if err == snap.ErrNoSnapshot {
					panic(err)
				}
				hash := common.BytesToHash(snapshot.Data)
				block := pm.blockchain.GetBlockByHash(hash)

				// it's possible for the block to not yet be in the chain if we just
				// joined and have to load a snapshot. Block here and wait for the
				// downloader to catch us up.
				if block == nil {
					peerID, peerTd := pm.peerGetter()
					pm.downloader.Synchronise(peerID, hash, peerTd, downloader.FullSync)

					if pm.blockchain.GetBlockByHash(hash) == nil {
						panic(fmt.Sprintf("downloader failed to synchronize block %x\n", hash))
					}
				}

				if err = pm.blockchain.FastSyncCommitHead(hash); err != nil {
					panic(err)
				}
			default:
				panic("Unhandled handleLogCommands type")
			}

		case <-pm.quitSync:
			return
		}
	}
}

func (pm *ProtocolManager) removeRlpxPeer(id string) error {
	pm.mu.Lock()
	if peer, ok := pm.rlpxKnownPeers[id]; ok {
		glog.V(logger.Debug).Infoln("Removing peer", id)
		logCheckpoint("PEER-DISCONNECTED", peer.uint64Id)
		delete(pm.rlpxKnownPeers, id)
		pm.mu.Unlock()
		peer.rawPeer.Disconnect(p2p.DiscUselessPeer)
	} else {
		pm.mu.Unlock()
	}

	return nil
}
