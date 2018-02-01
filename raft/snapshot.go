package raft

import (
	"fmt"
	"io"
	"math/big"
	"sort"
	"time"

	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
	"github.com/coreos/etcd/wal/walpb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"gopkg.in/fatih/set.v0"
)

// Snapshot

type Snapshot struct {
	addresses      []Address
	removedRaftIds []uint16 // Raft IDs for permanently removed peers
	headBlockHash  common.Hash
}

type ByRaftId []Address

func (a ByRaftId) Len() int           { return len(a) }
func (a ByRaftId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRaftId) Less(i, j int) bool { return a[i].raftId < a[j].raftId }

func (pm *ProtocolManager) buildSnapshot() *Snapshot {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	numNodes := len(pm.confState.Nodes)
	numRemovedNodes := pm.removedPeers.Size()

	snapshot := &Snapshot{
		addresses:      make([]Address, numNodes),
		removedRaftIds: make([]uint16, numRemovedNodes),
		headBlockHash:  pm.blockchain.CurrentBlock().Hash(),
	}

	// Populate addresses

	for i, rawRaftId := range pm.confState.Nodes {
		raftId := uint16(rawRaftId)

		if raftId == pm.raftId {
			snapshot.addresses[i] = *pm.address
		} else {
			snapshot.addresses[i] = *pm.peers[raftId].address
		}
	}
	sort.Sort(ByRaftId(snapshot.addresses))

	// Populate removed IDs

	for i, removedIface := range pm.removedPeers.List() {
		snapshot.removedRaftIds[i] = removedIface.(uint16)
	}

	return snapshot
}

// Note that we do *not* read `pm.appliedIndex` here. We only use the `index`
// parameter instead. This is because we need to support a scenario when we
// snapshot for a future index that we have not yet recorded in LevelDB. See
// comments around the use of `forceSnapshot`.
func (pm *ProtocolManager) triggerSnapshot(index uint64) {
	pm.mu.RLock()
	snapshotIndex := pm.snapshotIndex
	pm.mu.RUnlock()

	log.Info("start snapshot", "applied index", pm.appliedIndex, "last snapshot index", snapshotIndex)

	//snapData := pm.blockchain.CurrentBlock().Hash().Bytes()
	//snap, err := pm.raftStorage.CreateSnapshot(pm.appliedIndex, &pm.confState, snapData)
	snapData := pm.buildSnapshot().toBytes()
	snap, err := pm.raftStorage.CreateSnapshot(index, &pm.confState, snapData)
	if err != nil {
		panic(err)
	}
	if err := pm.saveRaftSnapshot(snap); err != nil {
		panic(err)
	}
	// Discard all log entries prior to index.
	if err := pm.raftStorage.Compact(index); err != nil {
		panic(err)
	}
	log.Info("compacted log", "index", pm.appliedIndex)

	pm.mu.Lock()
	pm.snapshotIndex = index
	pm.mu.Unlock()
}

func confStateIdSet(confState raftpb.ConfState) *set.Set {
	set := set.New()
	for _, rawRaftId := range confState.Nodes {
		set.Add(uint16(rawRaftId))
	}
	return set
}

func (pm *ProtocolManager) updateClusterMembership(newConfState raftpb.ConfState, addresses []Address, removedRaftIds []uint16) {
	log.Info("updating cluster membership per raft snapshot")

	prevConfState := pm.confState

	// Update tombstones for permanently removed peers. For simplicity we do not
	// allow the re-use of peer IDs once a peer is removed.

	removedPeers := set.New()
	for _, removedRaftId := range removedRaftIds {
		removedPeers.Add(removedRaftId)
	}
	pm.mu.Lock()
	pm.removedPeers = removedPeers
	pm.mu.Unlock()

	// Remove old peers that we're still connected to

	prevIds := confStateIdSet(prevConfState)
	newIds := confStateIdSet(newConfState)
	idsToRemove := set.Difference(prevIds, newIds)
	for _, idIfaceToRemove := range idsToRemove.List() {
		raftId := idIfaceToRemove.(uint16)
		log.Info("removing old raft peer", "peer id", raftId)

		pm.removePeer(raftId)
	}

	// Update local and remote addresses

	for _, tempAddress := range addresses {
		address := tempAddress // Allocate separately on the heap for each iteration.

		if address.raftId == pm.raftId {
			// If we're a newcomer to an existing cluster, this is where we learn
			// our own Address.
			pm.setLocalAddress(&address)
		} else {
			pm.mu.RLock()
			existingPeer := pm.peers[address.raftId]
			pm.mu.RUnlock()

			if existingPeer == nil {
				log.Info("adding new raft peer", "raft id", address.raftId)
				pm.addPeer(&address)
			}
		}
	}

	pm.mu.Lock()
	pm.confState = newConfState
	pm.mu.Unlock()

	log.Info("updated cluster membership")
}

func (pm *ProtocolManager) maybeTriggerSnapshot() {
	pm.mu.RLock()
	appliedIndex := pm.appliedIndex
	entriesSinceLastSnap := appliedIndex - pm.snapshotIndex
	pm.mu.RUnlock()

	if entriesSinceLastSnap < snapshotPeriod {
		return
	}

	pm.triggerSnapshot(appliedIndex)
}

func (pm *ProtocolManager) loadSnapshot() *raftpb.Snapshot {
	if raftSnapshot := pm.readRaftSnapshot(); raftSnapshot != nil {
		log.Info("loading snapshot")

		pm.applyRaftSnapshot(*raftSnapshot)

		return raftSnapshot
	} else {
		log.Info("no snapshot to load")

		return nil
	}
}

func (snapshot *Snapshot) toBytes() []byte {
	size, r, err := rlp.EncodeToReader(snapshot)
	if err != nil {
		panic(fmt.Sprintf("error: failed to RLP-encode Snapshot: %s", err.Error()))
	}
	var buffer = make([]byte, uint32(size))
	r.Read(buffer)

	return buffer
}

func bytesToSnapshot(bytes []byte) *Snapshot {
	var snapshot Snapshot
	if err := rlp.DecodeBytes(bytes, &snapshot); err != nil {
		fatalf("failed to RLP-decode Snapshot: %v", err)
	}
	return &snapshot
}

func (snapshot *Snapshot) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{snapshot.addresses, snapshot.removedRaftIds, snapshot.headBlockHash})
}

func (snapshot *Snapshot) DecodeRLP(s *rlp.Stream) error {
	// These fields need to be public:
	var temp struct {
		Addresses      []Address
		RemovedRaftIds []uint16
		HeadBlockHash  common.Hash
	}

	if err := s.Decode(&temp); err != nil {
		return err
	} else {
		snapshot.addresses, snapshot.removedRaftIds, snapshot.headBlockHash = temp.Addresses, temp.RemovedRaftIds, temp.HeadBlockHash
		return nil
	}
}

// Raft snapshot

func (pm *ProtocolManager) saveRaftSnapshot(snap raftpb.Snapshot) error {
	if err := pm.snapshotter.SaveSnap(snap); err != nil {
		return err
	}

	walSnap := walpb.Snapshot{
		Index: snap.Metadata.Index,
		Term:  snap.Metadata.Term,
	}

	if err := pm.wal.SaveSnapshot(walSnap); err != nil {
		return err
	}

	return pm.wal.ReleaseLockTo(snap.Metadata.Index)
}

func (pm *ProtocolManager) readRaftSnapshot() *raftpb.Snapshot {
	snapshot, err := pm.snapshotter.Load()
	if err != nil && err != snap.ErrNoSnapshot {
		fatalf("error loading snapshot: %v", err)
	}

	return snapshot
}

func (pm *ProtocolManager) applyRaftSnapshot(raftSnapshot raftpb.Snapshot) {
	log.Info("applying snapshot to raft storage")
	if err := pm.raftStorage.ApplySnapshot(raftSnapshot); err != nil {
		fatalf("failed to apply snapshot: %s", err)
	}
	snapshot := bytesToSnapshot(raftSnapshot.Data)

	latestBlockHash := snapshot.headBlockHash

	pm.updateClusterMembership(raftSnapshot.Metadata.ConfState, snapshot.addresses, snapshot.removedRaftIds)

	preSyncHead := pm.blockchain.CurrentBlock()

	if latestBlock := pm.blockchain.GetBlockByHash(latestBlockHash); latestBlock == nil {
		pm.syncBlockchainUntil(latestBlockHash)
		pm.logNewlyAcceptedTransactions(preSyncHead)

		log.Info(chainExtensionMessage, "hash", pm.blockchain.CurrentBlock().Hash())
	} else {
		log.Info("blockchain is caught up; no need to synchronize")
	}

	snapMeta := raftSnapshot.Metadata
	pm.confState = snapMeta.ConfState
	pm.mu.Lock()
	pm.snapshotIndex = snapMeta.Index
	pm.mu.Unlock()
}

func (pm *ProtocolManager) syncBlockchainUntil(hash common.Hash) {
	pm.mu.RLock()
	peerMap := make(map[uint16]*Peer, len(pm.peers))
	for raftId, peer := range pm.peers {
		peerMap[raftId] = peer
	}
	pm.mu.RUnlock()

	for {
		for peerId, peer := range peerMap {
			log.Info("synchronizing with peer", "peer id", peerId, "hash", hash)

			peerId := peer.p2pNode.ID.String()
			peerIdPrefix := fmt.Sprintf("%x", peer.p2pNode.ID[:8])

			if err := pm.downloader.Synchronise(peerIdPrefix, hash, big.NewInt(0), downloader.BoundedFullSync); err != nil {
				log.Info("failed to synchronize with peer", "peer id", peerId)

				time.Sleep(500 * time.Millisecond)
			} else {
				return
			}
		}
	}
}

func (pm *ProtocolManager) logNewlyAcceptedTransactions(preSyncHead *types.Block) {
	newHead := pm.blockchain.CurrentBlock()
	numBlocks := newHead.NumberU64() - preSyncHead.NumberU64()
	blocks := make([]*types.Block, numBlocks)
	currBlock := newHead
	blocksSeen := 0
	for currBlock.Hash() != preSyncHead.Hash() {
		blocks[int(numBlocks)-(1+blocksSeen)] = currBlock

		blocksSeen += 1
		currBlock = pm.blockchain.GetBlockByHash(currBlock.ParentHash())
	}
	for _, block := range blocks {
		for _, tx := range block.Transactions() {
			log.EmitCheckpoint(log.TxAccepted, "tx", tx.Hash().Hex())
		}
	}
}
