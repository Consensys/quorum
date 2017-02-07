// Overview of the channels used in this module:
//
// Node.
// * quitSync: *Every* channel operation can be unblocked by closing this
//   channel.
//
// ProtocolManager.
// * proposeC, for proposals flowing from ethereum to raft
// * confChangeC, currently unused; in the future for adding new, non-initial, raft peers
// * roleC, coming from raft notifies us when our role changes
package gethRaft

import (
	"fmt"
	"math/big"
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
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/coreos/etcd/etcdserver/stats"
	raftTypes "github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/rafthttp"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type ProtocolManager struct {
	// peers note -- each node tracks the peers acknowledged by raft
	//
	// only the leader proposes `ConfChangeAddNode` for each peer in the first set
	// but not in the second. this is done:
	// * when a node becomes leader
	// * when the leader learns of new peers

	// This node's raft id
	id int

	// set of currently active peers known to the raft cluster. this includes self
	raftPeers []raft.Peer
	peerUrls  []string
	p2pNodes  []*discover.Node

	blockchain *core.BlockChain

	// to protect the raft peers and addresses
	mu sync.RWMutex

	eventMux      *event.TypeMux
	minedBlockSub event.Subscription

	downloader *downloader.Downloader
	peerGetter func() (string, *big.Int)

	rawNode     raft.Node
	raftStorage *raft.MemoryStorage

	transport *rafthttp.Transport
	httpstopc chan struct{}
	httpdonec chan struct{}

	// The number of entries applied to the raft log
	appliedIndex uint64

	// The index of the latest snapshot.
	snapshotIndex uint64

	// Snapshotting
	snapshotter                      *snap.Snapshotter
	snapdir                          string
	confState                        raftpb.ConfState
	// when we restart a cluster, the first snapshot we receive is important: we
	// read its raftpb.ConfState to reconnect to the last-known members in the
	// cluster. once we perform this reconnection on startup, we will not do it
	// again for the life of this geth process.
	awaitingFirstSnapshotPostRestart bool

	// write-ahead log
	waldir string
	wal    *wal.WAL

	// We create a database with a single value -- the last applied index, used as an optimization for faster
	// startup
	appliedDb *leveldb.DB

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

//
// Raft interface methods
//
func (pm *ProtocolManager) Process(ctx context.Context, m raftpb.Message) error {
	return pm.rawNode.Step(ctx, m)
}
func (pm *ProtocolManager) IsIDRemoved(id uint64) bool {
	//
	// TODO: IMPLEMENT in the future once we support dynamic membership
	//

	glog.V(logger.Info).Infof("Reporting that raft ID %d is not removed", id)
	return false
}
func (pm *ProtocolManager) ReportUnreachable(id uint64) {
	//
	// TODO: Is there anything we need to do here?
	//

	glog.V(logger.Error).Infof("UNIMPLEMENTED Raft: ReportUnreachable. delegating to RawNode's impl")

	pm.rawNode.ReportUnreachable(id)
}
func (pm *ProtocolManager) ReportSnapshot(id uint64, status raft.SnapshotStatus) {
	//
	// TODO: Is there anything we need to do here?
	//

	glog.V(logger.Error).Infof("UNIMPLEMENTED Raft: ReportSnapshot. delegating to RawNode's impl")

	pm.rawNode.ReportSnapshot(id, status)
}

func (pm *ProtocolManager) startRaftNode(minter *minter) {
	if !fileutil.Exist(pm.snapdir) {
		if err := os.Mkdir(pm.snapdir, 0750); err != nil {
			glog.Fatalf("cannot create dir for snapshot (%v)", err)
		}
	}

	walExisted := wal.Exist(pm.waldir)

	pm.wal = pm.replayWAL()

	// NOTE: cockroach sets this to false for now until they've "worked out the
	//       bugs"
	enablePreVote := true

	lastAppliedIndex := pm.loadAppliedIndex()

	c := &raft.Config{
		Applied: lastAppliedIndex,
		ID:      uint64(pm.id),
		// TODO(joel): tune these parameters
		ElectionTick:  10, // NOTE: cockroach sets this to 15
		HeartbeatTick: 1,  // NOTE: cockroach sets this to 5
		Storage:       pm.raftStorage,

		// NOTE, from cockroach:
		// "PreVote and CheckQuorum are two ways of achieving the same thing.
		// PreVote is more compatible with quiesced ranges, so we want to switch
		// to it once we've worked out the bugs."
		PreVote:     enablePreVote,
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
		// raft.Ready operation.
		MaxInflightMsgs: 256, // NOTE: in cockroachdb this is 4
	}

	glog.V(logger.Info).Infof("local raft ID is %v", c.ID)

	if walExisted {
		if lastAppliedIndex > 0 {
			pm.awaitingFirstSnapshotPostRestart = true
		}

		pm.rawNode = raft.RestartNode(c)
	} else {
		if numPeers := len(pm.raftPeers); numPeers == 0 {
			panic("exiting due to empty raft peers list")
		} else {
			glog.V(logger.Info).Infof("starting raft with %v total peers.", numPeers)
		}

		pm.rawNode = raft.StartNode(c, pm.raftPeers)
	}

	ss := &stats.ServerStats{}
	ss.Initialize()

	pm.transport = &rafthttp.Transport{
		ID:          raftTypes.ID(pm.id),
		ClusterID:   0x1000,
		Raft:        pm,
		ServerStats: ss,
		LeaderStats: stats.NewLeaderStats(strconv.Itoa(pm.id)),
		ErrorC:      make(chan error),
	}

	pm.transport.Start()

	go pm.serveRaft()
	go pm.serveInternal(pm.proposeC, pm.confChangeC)
	go pm.eventLoop()
	go pm.handleRoleChange(pm.rawNode.RoleChan().Out(), minter)
}

func (pm *ProtocolManager) serveRaft() {
	urlString := fmt.Sprintf("http://0.0.0.0:%d", nodeHttpPort(pm.p2pNodes[pm.id-1]))
	url, err := url.Parse(urlString)
	if err != nil {
		glog.Fatalf("Failed parsing URL (%v)", err)
	}

	listener, err := newStoppableListener(url.Host, pm.httpstopc)
	if err != nil {
		glog.Fatalf("Failed to listen rafthttp (%v)", err)
	}
	err = (&http.Server{Handler: pm.transport.Handler()}).Serve(listener)
	select {
	case <-pm.httpstopc:
	default:
		glog.Fatalf("Failed to serve rafthttp (%v)", err)
	}
	close(pm.httpdonec)
}

func (pm *ProtocolManager) handleRoleChange(roleC <-chan interface{}, minter *minter) {
	for {
		select {
		case role := <-roleC:
			intRole, ok := role.(int)

			if !ok {
				panic("Couldn't cast role to int")
			}

			if intRole == minterRole {
				logCheckpoint(BECAME_MINTER, "")
				minter.start()
			} else { // verifier
				logCheckpoint(BECAME_VERIFIER, "")
				minter.stop()
			}

			pm.mu.Lock()
			pm.role = intRole
			pm.mu.Unlock()

		case <-pm.quitSync:
			return
		}
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
				glog.V(logger.Info).Infoln("error: read from proposeC failed")
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
				glog.V(logger.Info).Infoln("error: read from confChangeC failed")
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

func (pm *ProtocolManager) entriesToApply(ents []raftpb.Entry) (nents []raftpb.Entry) {
	if len(ents) == 0 {
		return
	}

	first := ents[0].Index
	lastApplied := pm.appliedIndex

	if first > lastApplied+1 {
		glog.Fatalf("first index of committed entry[%d] should <= appliedIndex[%d] + 1", first, lastApplied)
	}

	firstToApply := lastApplied - first + 1

	if firstToApply < uint64(len(ents)) {
		nents = ents[firstToApply:]
	}
	return
}

func (pm *ProtocolManager) addPeer(nodeId uint64, peerUrl string) {
	pm.transport.AddPeer(raftTypes.ID(nodeId), []string{peerUrl})
}

func (pm *ProtocolManager) reconnectToPreviousPeers(nodeIds []uint64) {
	for _, nodeId := range nodeIds {
		peerUrl := pm.loadPeerUrl(nodeId)

		pm.addPeer(nodeId, peerUrl)
	}
}

func (pm *ProtocolManager) eventLoop() {
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

			if snap := rd.Snapshot; !raft.IsEmptySnap(snap) {
				pm.saveSnapshot(snap)
				pm.applySnapshot(snap)

				if pm.awaitingFirstSnapshotPostRestart {
					pm.reconnectToPreviousPeers(snap.Metadata.ConfState.Nodes)

					pm.awaitingFirstSnapshotPostRestart = false
				}
			} else {
				if pm.awaitingFirstSnapshotPostRestart {
					panic("expected an initial snapshot post restart but did not receive one: unable to connect to peers.")
				}
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
						glog.V(logger.Error).Infoln("error decoding block: ", err)
					}
					pm.applyNewChainHead(&block)

				case raftpb.EntryConfChange:
					var cc raftpb.ConfChange
					cc.Unmarshal(entry.Data)

					// We lock access to this, in case we want to read the list of
					// cluster members concurrently via RPC (e.g. from NodeInfo()):
					pm.mu.Lock()
					pm.confState = *pm.rawNode.ApplyConfChange(cc)
					pm.mu.Unlock()

					switch cc.Type {
					case raftpb.ConfChangeAddNode:
						glog.V(logger.Info).Infof("adding peer %v due to ConfChangeAddNode", cc.NodeID)

						nodeId := cc.NodeID
						peerUrl := string(cc.Context)
						pm.addPeer(nodeId, peerUrl)
						pm.writePeerUrl(nodeId, peerUrl)

					case raftpb.ConfChangeRemoveNode:
						glog.V(logger.Info).Infof("removing peer %v due to ConfChangeRemoveNode", cc.NodeID)

						if cc.NodeID == uint64(pm.id) {
							glog.V(logger.Warn).Infoln("removing self from the cluster due to ConfChangeRemoveNode")

							pm.advanceAppliedIndex(entry.Index)

							// TODO: we might want to completely exit(0) geth here
							return
						}

						pm.transport.RemovePeer(raftTypes.ID(cc.NodeID))

					case raftpb.ConfChangeUpdateNode:
						glog.Fatalln("not yet handled: ConfChangeUpdateNode")
					}

					// We force a snapshot here to persist our updated confState, so we
					// know our fellow cluster members when we come back online.
					//
					// It is critical here to snapshot *before* writing our applied
					// index in LevelDB, otherwise a crash while/before snapshotting
					// (after advancing our applied index) would result in the loss of a
					// cluster member upon restart: we would re-mount with an old
					// ConfState.
					pm.triggerSnapshotWithNextIndex(entry.Index)
				}

				pm.advanceAppliedIndex(entry.Index)
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

func makeRaftPeers(urls []string) []raft.Peer {
	peers := make([]raft.Peer, len(urls))
	for i, url := range urls {
		peerId := i + 1

		peers[i] = raft.Peer{
			ID:      uint64(peerId),
			Context: []byte(url),
		}
	}
	return peers
}

func nodeHttpPort(node *discover.Node) uint16 {
	//
	// TODO: we should probably read this from the commandline, but it's a little trickier because we wouldn't be
	// accepting a single port like with --port or --rpcport; we'd have to ask for a base HTTP port (e.g. 50400)
	// with the convention/understanding that the port used by each node would be base + raft ID, which quorum is
	// otherwise not aware of.
	//
	return 20000 + node.TCP
}

func makePeerUrls(nodes []*discover.Node) []string {
	urls := make([]string, len(nodes))
	for i, node := range nodes {
		ip := node.IP.String()
		port := nodeHttpPort(node)
		urls[i] = fmt.Sprintf("http://%s:%d", ip, port)
	}

	return urls
}

func NewProtocolManager(id int, blockchain *core.BlockChain, mux *event.TypeMux, downloader *downloader.Downloader, peers []*discover.Node, peerGetter func() (string, *big.Int), datadir string) (*ProtocolManager, error) {
	waldir := fmt.Sprintf("%s/raft-wal", datadir)
	snapdir := fmt.Sprintf("%s/raft-snap", datadir)
	appliedDbLoc := fmt.Sprintf("%s/raft-applied", datadir)

	peerUrls := makePeerUrls(peers)
	manager := &ProtocolManager{
		raftPeers:   makeRaftPeers(peerUrls),
		peerUrls:    peerUrls,
		p2pNodes:    peers,
		blockchain:  blockchain,
		eventMux:    mux,
		downloader:  downloader,
		peerGetter:  peerGetter,
		proposeC:    make(chan *types.Block),
		confChangeC: make(chan raftpb.ConfChange),
		httpstopc:   make(chan struct{}),
		httpdonec:   make(chan struct{}),
		waldir:      waldir,
		snapdir:     snapdir,
		snapshotter: snap.New(snapdir),
		id:          id,
		quitSync:    make(chan struct{}),
		raftStorage: raft.NewMemoryStorage(),
	}

	// Open the db and recover any potential corruptions
	db, err := leveldb.OpenFile(appliedDbLoc, &opt.Options{
		OpenFilesCacheCapacity: -1, // -1 means 0??
		BlockCacheCapacity:     -1,
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(appliedDbLoc, nil)
	}
	// (Re)check for errors and abort if opening of the db failed
	if err != nil {
		return nil, err
	}
	manager.appliedDb = db

	return manager, nil
}

func (pm *ProtocolManager) NodeInfo() *RaftNodeInfo {
	pm.mu.RLock() // as we read pm.role
	defer pm.mu.RUnlock()

	var roleDescription string
	if pm.role == minterRole {
		roleDescription = "minter"
	} else {
		roleDescription = "verifier"
	}

	return &RaftNodeInfo{
		ClusterSize: len(pm.raftPeers),
		Genesis:     pm.blockchain.Genesis().Hash(),
		Head:        pm.blockchain.CurrentBlock().Hash(),
		Role:        roleDescription,
	}
}

func sleep(duration time.Duration) {
	<-time.NewTimer(duration).C
}

func (pm *ProtocolManager) Start(minter *minter) {
	pm.minedBlockSub = pm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	go pm.minedBroadcastLoop(pm.proposeC)
	pm.startRaftNode(minter)
}

func (pm *ProtocolManager) Stop() {
	glog.V(logger.Info).Infoln("Stopping raft protocol handler...")

	pm.transport.Stop()
	close(pm.httpstopc)
	<-pm.httpdonec
	close(pm.quitSync)
	if pm.rawNode != nil {
		pm.rawNode.Stop()
	}

	pm.appliedDb.Close()

	//
	// TODO: stop minting here
	//

	glog.V(logger.Info).Infoln("Raft protocol handler stopped")
}

func logCheckpoint(checkpointName string, iface interface{}) {
	glog.V(logger.Info).Infof("RAFT-CHECKPOINT %s %v\n", checkpointName, iface)
}

func blockExtendsChain(block *types.Block, chain *core.BlockChain) bool {
	return block.ParentHash() == chain.CurrentBlock().Hash()
}

func (pm *ProtocolManager) applyNewChainHead(block *types.Block) {
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

		for _, tx := range block.Transactions() {
			logCheckpoint(TX_ACCEPTED, tx.Hash().Hex())
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
}

// Sets new appliedIndex in-memory, *and* writes this appliedIndex to LevelDB.
func (pm *ProtocolManager) advanceAppliedIndex(index uint64) {
	pm.appliedIndex = index

	pm.writeAppliedIndex(index)
}