package raft

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/coreos/etcd/pkg/fileutil"
	"github.com/coreos/etcd/snap"
	"github.com/coreos/etcd/wal"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/coreos/etcd/etcdserver/stats"
	raftTypes "github.com/coreos/etcd/pkg/types"
	etcdRaft "github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/rafthttp"
	"github.com/syndtr/goleveldb/leveldb"
	"gopkg.in/fatih/set.v0"
)

type ProtocolManager struct {
	mu       sync.RWMutex // For protecting concurrent JS access to "local peer" and "remote peer" state
	quitSync chan struct{}
	stopped  bool

	// Static configuration
	joinExisting   bool // Whether to join an existing cluster when a WAL doesn't already exist
	bootstrapNodes []*discover.Node
	raftId         uint16
	raftPort       uint16

	// Local peer state (protected by mu vs concurrent access via JS)
	address       *Address
	role          int    // Role: minter or verifier
	appliedIndex  uint64 // The index of the last-applied raft entry
	snapshotIndex uint64 // The index of the latest snapshot.

	// Remote peer state (protected by mu vs concurrent access via JS)
	peers        map[uint16]*Peer
	removedPeers *set.Set // *Permanently removed* peers

	// P2P transport
	p2pServer *p2p.Server // Initialized in start()

	// Blockchain services
	blockchain *core.BlockChain
	downloader *downloader.Downloader
	minter     *minter

	// Blockchain events
	eventMux      *event.TypeMux
	minedBlockSub *event.TypeMuxSubscription

	// Raft proposal events
	blockProposalC      chan *types.Block      // for mined blocks to raft
	confChangeProposalC chan raftpb.ConfChange // for config changes from js console to raft

	// Raft transport
	unsafeRawNode etcdRaft.Node
	transport     *rafthttp.Transport
	httpstopc     chan struct{}
	httpdonec     chan struct{}

	// Raft snapshotting
	snapshotter *snap.Snapshotter
	snapdir     string
	confState   raftpb.ConfState

	// Raft write-ahead log
	waldir string
	wal    *wal.WAL

	// Storage
	quorumRaftDb *leveldb.DB             // Persistent storage for last-applied raft index
	raftStorage  *etcdRaft.MemoryStorage // Volatile raft storage
}

//
// Public interface
//

func NewProtocolManager(raftId uint16, raftPort uint16, blockchain *core.BlockChain, mux *event.TypeMux, bootstrapNodes []*discover.Node, joinExisting bool, datadir string, minter *minter, downloader *downloader.Downloader) (*ProtocolManager, error) {
	waldir := fmt.Sprintf("%s/raft-wal", datadir)
	snapdir := fmt.Sprintf("%s/raft-snap", datadir)
	quorumRaftDbLoc := fmt.Sprintf("%s/quorum-raft-state", datadir)

	manager := &ProtocolManager{
		bootstrapNodes:      bootstrapNodes,
		peers:               make(map[uint16]*Peer),
		removedPeers:        set.New(),
		joinExisting:        joinExisting,
		blockchain:          blockchain,
		eventMux:            mux,
		blockProposalC:      make(chan *types.Block),
		confChangeProposalC: make(chan raftpb.ConfChange),
		httpstopc:           make(chan struct{}),
		httpdonec:           make(chan struct{}),
		waldir:              waldir,
		snapdir:             snapdir,
		snapshotter:         snap.New(snapdir),
		raftId:              raftId,
		raftPort:            raftPort,
		quitSync:            make(chan struct{}),
		raftStorage:         etcdRaft.NewMemoryStorage(),
		minter:              minter,
		downloader:          downloader,
	}

	if db, err := openQuorumRaftDb(quorumRaftDbLoc); err != nil {
		return nil, err
	} else {
		manager.quorumRaftDb = db
	}

	return manager, nil
}

func (pm *ProtocolManager) Start(p2pServer *p2p.Server) {
	log.Info("starting raft protocol handler")

	pm.p2pServer = p2pServer
	pm.minedBlockSub = pm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	pm.startRaft()
	go pm.minedBroadcastLoop()
}

func (pm *ProtocolManager) Stop() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	defer log.Info("raft protocol handler stopped")

	if pm.stopped {
		return
	}

	log.Info("stopping raft protocol handler...")

	for raftId, peer := range pm.peers {
		pm.disconnectFromPeer(raftId, peer)
	}

	pm.minedBlockSub.Unsubscribe()

	if pm.transport != nil {
		pm.transport.Stop()
	}

	close(pm.httpstopc)
	<-pm.httpdonec
	close(pm.quitSync)

	if pm.unsafeRawNode != nil {
		pm.unsafeRawNode.Stop()
	}

	pm.quorumRaftDb.Close()

	pm.p2pServer = nil

	pm.minter.stop()

	pm.stopped = true
}

func (pm *ProtocolManager) NodeInfo() *RaftNodeInfo {
	pm.mu.RLock() // as we read role and peers
	defer pm.mu.RUnlock()

	var roleDescription string
	if pm.role == minterRole {
		roleDescription = "minter"
	} else {
		roleDescription = "verifier"
	}

	peerAddresses := make([]*Address, len(pm.peers))
	peerIdx := 0
	for _, peer := range pm.peers {
		peerAddresses[peerIdx] = peer.address
		peerIdx += 1
	}

	removedPeerIfaces := pm.removedPeers.List()
	removedPeerIds := make([]uint16, len(removedPeerIfaces))
	for i, removedIface := range removedPeerIfaces {
		removedPeerIds[i] = removedIface.(uint16)
	}

	//
	// NOTE: before exposing any new fields here, make sure that the underlying
	// ProtocolManager members are protected from concurrent access by pm.mu!
	//
	return &RaftNodeInfo{
		ClusterSize:    len(pm.peers) + 1,
		Role:           roleDescription,
		Address:        pm.address,
		PeerAddresses:  peerAddresses,
		RemovedPeerIds: removedPeerIds,
		AppliedIndex:   pm.appliedIndex,
		SnapshotIndex:  pm.snapshotIndex,
	}
}

// There seems to be a very rare race in raft where during `etcdRaft.StartNode`
// it will call back our `Process` method before it's finished returning the
// `raft.Node`, `pm.unsafeRawNode`, to us. This re-entrance through a separate
// thread will cause a nil pointer dereference. To work around this, this
// getter method should be used instead of reading `pm.unsafeRawNode` directly.
func (pm *ProtocolManager) rawNode() etcdRaft.Node {
	for pm.unsafeRawNode == nil {
		time.Sleep(100 * time.Millisecond)
	}

	return pm.unsafeRawNode
}

func (pm *ProtocolManager) nextRaftId() uint16 {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	maxId := pm.raftId

	for peerId := range pm.peers {
		if maxId < peerId {
			maxId = peerId
		}
	}

	removedPeerIfaces := pm.removedPeers.List()
	for _, removedIface := range removedPeerIfaces {
		removedId := removedIface.(uint16)

		if maxId < removedId {
			maxId = removedId
		}
	}

	return maxId + 1
}

func (pm *ProtocolManager) isRaftIdRemoved(id uint16) bool {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.removedPeers.Has(id)
}

func (pm *ProtocolManager) isRaftIdUsed(raftId uint16) bool {
	if pm.raftId == raftId || pm.isRaftIdRemoved(raftId) {
		return true
	}

	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.peers[raftId] != nil
}

func (pm *ProtocolManager) isP2pNodeInCluster(node *discover.Node) bool {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	for _, peer := range pm.peers {
		if peer.p2pNode.ID == node.ID {
			return true
		}
	}

	return false
}

func (pm *ProtocolManager) ProposeNewPeer(enodeId string) (uint16, error) {
	node, err := discover.ParseNode(enodeId)
	if err != nil {
		return 0, err
	}

	if pm.isP2pNodeInCluster(node) {
		return 0, fmt.Errorf("node is already in the cluster: %v", enodeId)
	}

	if len(node.IP) != 4 {
		return 0, fmt.Errorf("expected IPv4 address (with length 4), but got IP of length %v", len(node.IP))
	}

	if !node.HasRaftPort() {
		return 0, fmt.Errorf("enodeId is missing raftport querystring parameter: %v", enodeId)
	}

	raftId := pm.nextRaftId()
	address := newAddress(raftId, node.RaftPort, node)

	pm.confChangeProposalC <- raftpb.ConfChange{
		Type:    raftpb.ConfChangeAddNode,
		NodeID:  uint64(raftId),
		Context: address.toBytes(),
	}

	return raftId, nil
}

func (pm *ProtocolManager) ProposePeerRemoval(raftId uint16) {
	pm.confChangeProposalC <- raftpb.ConfChange{
		Type:   raftpb.ConfChangeRemoveNode,
		NodeID: uint64(raftId),
	}
}

//
// MsgWriter interface (necessary for p2p.Send)
//

func (pm *ProtocolManager) WriteMsg(msg p2p.Msg) error {
	// read *into* buffer
	var buffer = make([]byte, msg.Size)
	msg.Payload.Read(buffer)

	return pm.rawNode().Propose(context.TODO(), buffer)
}

//
// Raft interface
//

func (pm *ProtocolManager) Process(ctx context.Context, m raftpb.Message) error {
	return pm.rawNode().Step(ctx, m)
}

func (pm *ProtocolManager) IsIDRemoved(id uint64) bool {
	return pm.isRaftIdRemoved(uint16(id))
}

func (pm *ProtocolManager) ReportUnreachable(id uint64) {
	log.Info("peer is currently unreachable", "peer id", id)

	pm.rawNode().ReportUnreachable(id)
}

func (pm *ProtocolManager) ReportSnapshot(id uint64, status etcdRaft.SnapshotStatus) {
	if status == etcdRaft.SnapshotFailure {
		log.Info("failed to send snapshot", "raft peer", id)
	} else if status == etcdRaft.SnapshotFinish {
		log.Info("finished sending snapshot", "raft peer", id)
	}

	pm.rawNode().ReportSnapshot(id, status)
}

//
// Private methods
//

func (pm *ProtocolManager) startRaft() {
	if !fileutil.Exist(pm.snapdir) {
		if err := os.Mkdir(pm.snapdir, 0750); err != nil {
			fatalf("cannot create dir for snapshot (%v)", err)
		}
	}
	walExisted := wal.Exist(pm.waldir)
	lastAppliedIndex := pm.loadAppliedIndex()

	ss := &stats.ServerStats{}
	ss.Initialize()
	pm.transport = &rafthttp.Transport{
		ID:          raftTypes.ID(pm.raftId),
		ClusterID:   0x1000,
		Raft:        pm,
		ServerStats: ss,
		LeaderStats: stats.NewLeaderStats(strconv.Itoa(int(pm.raftId))),
		ErrorC:      make(chan error),
	}
	pm.transport.Start()

	// We load the snapshot to connect to prev peers before replaying the WAL,
	// which typically goes further into the future than the snapshot.

	var maybeRaftSnapshot *raftpb.Snapshot

	if walExisted {
		maybeRaftSnapshot = pm.loadSnapshot() // re-establishes peer connections
	}

	pm.wal = pm.replayWAL(maybeRaftSnapshot)

	if walExisted {
		if hardState, _, err := pm.raftStorage.InitialState(); err != nil {
			panic(fmt.Sprintf("failed to read initial state from raft while restarting: %v", err))
		} else {
			if lastPersistedCommittedIndex := hardState.Commit; lastPersistedCommittedIndex < lastAppliedIndex {
				log.Info("rolling back applied index to last-durably-committed", "last applied index", lastAppliedIndex, "last persisted index", lastPersistedCommittedIndex)

				// Roll back our applied index. See the logic and explanation around
				// the single call to `pm.applyNewChainHead` for more context.
				lastAppliedIndex = lastPersistedCommittedIndex
			}
		}
	}

	// NOTE: cockroach sets this to false for now until they've "worked out the
	//       bugs"
	enablePreVote := true

	raftConfig := &etcdRaft.Config{
		Applied:       lastAppliedIndex,
		ID:            uint64(pm.raftId),
		ElectionTick:  10, // NOTE: cockroach sets this to 15
		HeartbeatTick: 1,  // NOTE: cockroach sets this to 5
		Storage:       pm.raftStorage,

		// NOTE, from cockroach:
		// "PreVote and CheckQuorum are two ways of achieving the same thing.
		// PreVote is more compatible with quiesced ranges, so we want to switch
		// to it once we've worked out the bugs."
		//
		// TODO: vendor again?
		// PreVote:     enablePreVote,
		CheckQuorum: !enablePreVote,

		// MaxSizePerMsg controls how many Raft log entries the leader will send to
		// followers in a single MsgApp.
		MaxSizePerMsg: 4096, // NOTE: in cockroachdb this is 16*1024

		// MaxInflightMsgs controls how many in-flight messages Raft will send to
		// a follower without hearing a response. The total number of Raft log
		// entries is a combination of this setting and MaxSizePerMsg.
		//
		// NOTE: Cockroach's settings (MaxSizePerMsg of 4k and MaxInflightMsgs
		// of 4) provide for up to 64 KB of raft log to be sent without
		// acknowledgement. With an average entry size of 1 KB that translates
		// to ~64 commands that might be executed in the handling of a single
		// etcdraft.Ready operation.
		MaxInflightMsgs: 256, // NOTE: in cockroachdb this is 4
	}

	log.Info("startRaft", "raft ID", raftConfig.ID)

	if walExisted {
		log.Info("remounting an existing raft log; connecting to peers.")
		pm.unsafeRawNode = etcdRaft.RestartNode(raftConfig)
	} else if pm.joinExisting {
		log.Info("newly joining an existing cluster; waiting for connections.")
		pm.unsafeRawNode = etcdRaft.StartNode(raftConfig, nil)
	} else {
		if numPeers := len(pm.bootstrapNodes); numPeers == 0 {
			panic("exiting due to empty raft peers list")
		} else {
			log.Info("starting a new raft log", "initial cluster size of", numPeers)
		}

		raftPeers, peerAddresses, localAddress := pm.makeInitialRaftPeers()

		pm.setLocalAddress(localAddress)

		// We add all peers up-front even though we will see a ConfChangeAddNode
		// for each shortly. This is because raft's ConfState will contain all of
		// these nodes before we see these log entries, and we always want our
		// snapshots to have all addresses for each of the nodes in the ConfState.
		for _, peerAddress := range peerAddresses {
			pm.addPeer(peerAddress)
		}

		pm.unsafeRawNode = etcdRaft.StartNode(raftConfig, raftPeers)
	}

	go pm.serveRaft()
	go pm.serveLocalProposals()
	go pm.eventLoop()
	go pm.handleRoleChange(pm.rawNode().RoleChan().Out())
}

func (pm *ProtocolManager) setLocalAddress(addr *Address) {
	pm.mu.Lock()
	pm.address = addr
	pm.mu.Unlock()

	// By setting `URLs` on the raft transport, we advertise our URL (in an HTTP
	// header) to any recipient. This is necessary for a newcomer to the cluster
	// to be able to accept a snapshot from us to bootstrap them.
	if urls, err := raftTypes.NewURLs([]string{raftUrl(addr)}); err == nil {
		pm.transport.URLs = urls
	} else {
		panic(fmt.Sprintf("error: could not create URL from local address: %v", addr))
	}
}

func (pm *ProtocolManager) serveRaft() {
	urlString := fmt.Sprintf("http://0.0.0.0:%d", pm.raftPort)
	url, err := url.Parse(urlString)
	if err != nil {
		fatalf("Failed parsing URL (%v)", err)
	}

	listener, err := newStoppableListener(url.Host, pm.httpstopc)
	if err != nil {
		fatalf("Failed to listen rafthttp (%v)", err)
	}
	err = (&http.Server{Handler: pm.transport.Handler()}).Serve(listener)
	select {
	case <-pm.httpstopc:
	default:
		fatalf("Failed to serve rafthttp (%v)", err)
	}
	close(pm.httpdonec)
}

func (pm *ProtocolManager) handleRoleChange(roleC <-chan interface{}) {
	for {
		select {
		case role := <-roleC:
			intRole, ok := role.(int)

			if !ok {
				panic("Couldn't cast role to int")
			}

			if intRole == minterRole {
				log.EmitCheckpoint(log.BecameMinter)
				pm.minter.start()
			} else { // verifier
				log.EmitCheckpoint(log.BecameVerifier)
				pm.minter.stop()
			}

			pm.mu.Lock()
			pm.role = intRole
			pm.mu.Unlock()

		case <-pm.quitSync:
			return
		}
	}
}

func (pm *ProtocolManager) minedBroadcastLoop() {
	for obj := range pm.minedBlockSub.Chan() {
		switch ev := obj.Data.(type) {
		case core.NewMinedBlockEvent:
			select {
			case pm.blockProposalC <- ev.Block:
			case <-pm.quitSync:
				return
			}
		}
	}
}

// Serve two channels to handle new blocks and raft configuration changes originating locally.
func (pm *ProtocolManager) serveLocalProposals() {
	//
	// TODO: does it matter that this will restart from 0 whenever we restart a cluster?
	//
	var confChangeCount uint64

	for {
		select {
		case block, ok := <-pm.blockProposalC:
			if !ok {
				log.Info("error: read from proposeC failed")
				return
			}

			size, r, err := rlp.EncodeToReader(block)
			if err != nil {
				panic(fmt.Sprintf("error: failed to send RLP-encoded block: %s", err.Error()))
			}
			var buffer = make([]byte, uint32(size))
			r.Read(buffer)

			// blocks until accepted by the raft state machine
			pm.rawNode().Propose(context.TODO(), buffer)
		case cc, ok := <-pm.confChangeProposalC:
			if !ok {
				log.Info("error: read from confChangeC failed")
				return
			}

			confChangeCount++
			cc.ID = confChangeCount
			pm.rawNode().ProposeConfChange(context.TODO(), cc)
		case <-pm.quitSync:
			return
		}
	}
}

func (pm *ProtocolManager) entriesToApply(allEntries []raftpb.Entry) (entriesToApply []raftpb.Entry) {
	if len(allEntries) == 0 {
		return
	}

	first := allEntries[0].Index
	pm.mu.RLock()
	lastApplied := pm.appliedIndex
	pm.mu.RUnlock()

	if first > lastApplied+1 {
		fatalf("first index of committed entry[%d] should <= appliedIndex[%d] + 1", first, lastApplied)
	}

	firstToApply := lastApplied - first + 1

	if firstToApply < uint64(len(allEntries)) {
		entriesToApply = allEntries[firstToApply:]
	}
	return
}

func raftUrl(address *Address) string {
	return fmt.Sprintf("http://%s:%d", address.ip, address.raftPort)
}

func (pm *ProtocolManager) addPeer(address *Address) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	raftId := address.raftId

	// Add P2P connection:
	p2pNode := discover.NewNode(address.nodeId, address.ip, 0, uint16(address.p2pPort))
	pm.p2pServer.AddPeer(p2pNode)

	// Add raft transport connection:
	pm.transport.AddPeer(raftTypes.ID(raftId), []string{raftUrl(address)})
	pm.peers[raftId] = &Peer{address, p2pNode}
}

func (pm *ProtocolManager) disconnectFromPeer(raftId uint16, peer *Peer) {
	pm.p2pServer.RemovePeer(peer.p2pNode)
	pm.transport.RemovePeer(raftTypes.ID(raftId))
}

func (pm *ProtocolManager) removePeer(raftId uint16) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if peer := pm.peers[raftId]; peer != nil {
		pm.disconnectFromPeer(raftId, peer)

		delete(pm.peers, raftId)
	}

	// This is only necessary sometimes, but it's idempotent. Also, we *always*
	// do this, and not just when there's still a peer in the map, because we
	// need to do it for our *own* raft ID before we get booted from the cluster
	// so that snapshots are identical on all nodes. It's important for a booted
	// node to have a snapshot identical to every other node because that node
	// can potentially re-enter the cluster with a new raft ID.
	pm.removedPeers.Add(raftId)
}

func (pm *ProtocolManager) eventLoop() {
	ticker := time.NewTicker(tickerMS * time.Millisecond)
	defer ticker.Stop()
	defer pm.wal.Close()

	exitAfterApplying := false

	for {
		select {
		case <-ticker.C:
			pm.rawNode().Tick()

		// when the node is first ready it gives us entries to commit and messages
		// to immediately publish
		case rd := <-pm.rawNode().Ready():
			pm.wal.Save(rd.HardState, rd.Entries)

			if snap := rd.Snapshot; !etcdRaft.IsEmptySnap(snap) {
				pm.saveRaftSnapshot(snap)
				pm.applyRaftSnapshot(snap)
				pm.advanceAppliedIndex(snap.Metadata.Index)
			}

			// 1: Write HardState, Entries, and Snapshot to persistent storage if they
			// are not empty.
			pm.raftStorage.Append(rd.Entries)

			// 2: Send all Messages to the nodes named in the To field.
			pm.transport.Send(rd.Messages)

			// 3: Apply Snapshot (if any) and CommittedEntries to the state machine.
			for _, entry := range pm.entriesToApply(rd.CommittedEntries) {
				switch entry.Type {
				case raftpb.EntryNormal:
					if len(entry.Data) == 0 {
						break
					}
					var block types.Block
					err := rlp.DecodeBytes(entry.Data, &block)
					if err != nil {
						log.Error("error decoding block: ", err)
					}

					if pm.blockchain.HasBlock(block.Hash(), block.NumberU64()) {
						// This can happen:
						//
						// if (1) we crashed after applying this block to the chain, but
						//        before writing appliedIndex to LDB.
						// or (2) we crashed in a scenario where we applied further than
						//        raft *durably persisted* its committed index (see
						//        https://github.com/coreos/etcd/pull/7899). In this
						//        scenario, when the node comes back up, we will re-apply
						//        a few entries.

						headBlockHash := pm.blockchain.CurrentBlock().Hash()
						log.Warn("not applying already-applied block", "block hash", block.Hash(), "parent", block.ParentHash(), "head", headBlockHash)
					} else {
						pm.applyNewChainHead(&block)
					}

				case raftpb.EntryConfChange:
					var cc raftpb.ConfChange
					cc.Unmarshal(entry.Data)
					raftId := uint16(cc.NodeID)

					pm.confState = *pm.rawNode().ApplyConfChange(cc)

					forceSnapshot := false

					switch cc.Type {
					case raftpb.ConfChangeAddNode:
						if pm.isRaftIdRemoved(raftId) {
							log.Info("ignoring ConfChangeAddNode for permanently-removed peer", "raft id", raftId)
						} else if raftId <= uint16(len(pm.bootstrapNodes)) {
							// See initial cluster logic in startRaft() for more information.
							log.Info("ignoring expected ConfChangeAddNode for initial peer", "raft id", raftId)

							// We need a snapshot to exist to reconnect to peers on start-up after a crash.
							forceSnapshot = true
						} else if pm.isRaftIdUsed(raftId) {
							log.Info("ignoring ConfChangeAddNode for already-used raft ID", "raft id", raftId)
						} else {
							log.Info("adding peer due to ConfChangeAddNode", "raft id", raftId)

							forceSnapshot = true
							pm.addPeer(bytesToAddress(cc.Context))
						}

					case raftpb.ConfChangeRemoveNode:
						if pm.isRaftIdRemoved(raftId) {
							log.Info("ignoring ConfChangeRemoveNode for already-removed peer", "raft id", raftId)
						} else {
							log.Info("removing peer due to ConfChangeRemoveNode", "raft id", raftId)

							forceSnapshot = true

							if raftId == pm.raftId {
								exitAfterApplying = true
							}

							pm.removePeer(raftId)
						}

					case raftpb.ConfChangeUpdateNode:
						// NOTE: remember to forceSnapshot in this case, if we add support
						// for this.
						fatalf("not yet handled: ConfChangeUpdateNode")
					}

					if forceSnapshot {
						// We force a snapshot here to persist our updated confState, so we
						// know our fellow cluster members when we come back online.
						//
						// It is critical here to snapshot *before* writing our applied
						// index in LevelDB, otherwise a crash while/before snapshotting
						// (after advancing our applied index) would result in the loss of a
						// cluster member upon restart: we would re-mount with an old
						// ConfState.
						pm.triggerSnapshot(entry.Index)
					}
				}

				pm.advanceAppliedIndex(entry.Index)
			}

			pm.maybeTriggerSnapshot()

			if exitAfterApplying {
				log.Warn("permanently removing self from the cluster")
				pm.Stop()
				log.Warn("permanently exited the cluster")

				return
			}

			// 4: Call Node.Advance() to signal readiness for the next batch of
			// updates.
			pm.rawNode().Advance()

		case <-pm.quitSync:
			return
		}
	}
}

func (pm *ProtocolManager) makeInitialRaftPeers() (raftPeers []etcdRaft.Peer, peerAddresses []*Address, localAddress *Address) {
	initialNodes := pm.bootstrapNodes
	raftPeers = make([]etcdRaft.Peer, len(initialNodes))  // Entire cluster
	peerAddresses = make([]*Address, len(initialNodes)-1) // Cluster without *this* node

	peersSeen := 0
	for i, node := range initialNodes {
		raftId := uint16(i + 1)
		// We initially get the raftPort from the enode ID's query string. As an alternative, we can move away from
		// requiring the use of static peers for the initial set, and load them from e.g. another JSON file which
		// contains pairs of enodes and raft ports, or we can get this initial peer list from commandline flags.
		address := newAddress(raftId, node.RaftPort, node)

		raftPeers[i] = etcdRaft.Peer{
			ID:      uint64(raftId),
			Context: address.toBytes(),
		}

		if raftId == pm.raftId {
			localAddress = address
		} else {
			peerAddresses[peersSeen] = address
			peersSeen += 1
		}
	}

	return
}

func sleep(duration time.Duration) {
	<-time.NewTimer(duration).C
}

func blockExtendsChain(block *types.Block, chain *core.BlockChain) bool {
	return block.ParentHash() == chain.CurrentBlock().Hash()
}

func (pm *ProtocolManager) applyNewChainHead(block *types.Block) {
	if !blockExtendsChain(block, pm.blockchain) {
		headBlock := pm.blockchain.CurrentBlock()

		log.Info("Non-extending block", "block", block.Hash(), "parent", block.ParentHash(), "head", headBlock.Hash())

		pm.minter.invalidRaftOrderingChan <- InvalidRaftOrdering{headBlock: headBlock, invalidBlock: block}
	} else {
		if existingBlock := pm.blockchain.GetBlockByHash(block.Hash()); nil == existingBlock {
			if err := pm.blockchain.Validator().ValidateBody(block); err != nil {
				panic(fmt.Sprintf("failed to validate block %x (%v)", block.Hash(), err))
			}
		}

		for _, tx := range block.Transactions() {
			log.EmitCheckpoint(log.TxAccepted, "tx", tx.Hash().Hex())
		}

		_, err := pm.blockchain.InsertChain([]*types.Block{block})

		if err != nil {
			panic(fmt.Sprintf("failed to extend chain: %s", err.Error()))
		}

		log.EmitCheckpoint(log.BlockCreated, "block", fmt.Sprintf("%x", block.Hash()))
	}
}

// Sets new appliedIndex in-memory, *and* writes this appliedIndex to LevelDB.
func (pm *ProtocolManager) advanceAppliedIndex(index uint64) {
	pm.writeAppliedIndex(index)

	pm.mu.Lock()
	pm.appliedIndex = index
	pm.mu.Unlock()
}
