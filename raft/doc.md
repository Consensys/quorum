# Raft / Ethereum

## Introduction

This directory holds an implementation of [Raft](https://raft.github.io) (using [etcd](https://github.com/coreos/etcd)) as an alternative consensus mechanism for the Ethereum VM instead of the default proof-of-work.

Both Raft and Ethereum have their own notion of node:

In Raft a node in normal operation is either a leader or a follower. There is a single leader for the entire cluster, which all log entries must go through. There's also the concept of a "Candidate", but only during leader election. We won't go into more detail about Raft here, because by design these details should be opaque to applications built on it.

In vanilla Ethereum there is no such thing as a leader or follower. It's possible for any node in the network to mine a new block -- which is like being the leader for that round.

In Quorum-Raft We impose a one-to-one correspondence between Raft and Ethereum nodes, so each Ethereum node is also a Raft node, and the leader of the Raft cluster is the only Ethereum node that can mint new blocks. The reason we make sure to colocate the leader and minter is to avoid the latency of transferring the block from minter to raft leader.

Ethereum | Raft
-------- | ----
minter   | leader
verifier | follower

Note that Raft is responsible for creating blocks (bundles of transactions), but at a lower level we rely on Ethereum's built-in p2p network ([RLPx](https://github.com/ethereum/devp2p/blob/master/rlpx.md)) to communicate transactions to the leader. Though it's used for all transactions, this is only important for transactions that originate at a follower. Transactions from the leader can be immediately minted in a block, put into Raft, and then put into the chain.

### Lifecycle of a Transactions

Note: We use "Etcd" for the Etcd implementation of Raft, and "Raft" more broadly to also include Raft-Ethereum.

Let's follow the lifecycle of a typical transaction:

#### on a follower:

1. The transaction is created by a user's rpc call.
2. Using the existing transaction propagation mechanisms in Ethereum, it's announced to all peers and transmitted to a subset of them.

#### on the leader:

3. It reaches the leader, where it's minted in the next block (see `mintNewBlock`).
4. That triggers `NewMinedBlockEvent`, which the Raft protocol manager is subscribed to (`minedBlockSub`). The `minedBroadcastLoop` enqueues the new block in `ProtocolManager.proposeC`.
5. `serveInternal` is waiting at the other end of the channel. Its job is to rlp-encode blocks and propose them to Raft.
6. The block reaches the `eventLoop` (which processes Raft events), as an `AppendEntries` message in `rd.Entries`. It, and all the other messages, are sent through `pm.transport` -- an instance of [`rafthttp.Transport`](https://godoc.org/github.com/coreos/etcd/rafthttp#Transport) -- responsible for communicating raft messages to all peers.

#### on followers:

7. The message reaches peers through `ProtocolManager.Process`, which `Step`s the raft state machine.
8. Etcd creates an `AppendEntries` response (acknowledging its acceptance of the block), which `eventLoop` handles, calls `transport.Send` (the message should have only the leader in its `To` field).

#### on the leader:

9. The leader handles each response through its incoming message loop. When a quorum have been received, Etcd considers the block committed, and it's moved into the `eventLoop` as a `CommittedEntries`.

#### on all:

10. The block is delivered by raft, then handled by `applyNewChainHead`. This method checks whether the block extends the chain (it's "valid", see below) and if not just ignores it. If so, the block is validated and set as the new head of the chain by `SetNewHeadBlock`.
11. A `ChainHeadEvent` is posted to notify listeners that a new block has been accepted. This is relevant to us because:
* It removes the newly minted transactions from `proposedTxes` (see below).
* It triggers `requestMinting` in the (minter's) worker, telling Raft to try minting a new block.

The transaction is now available on all nodes in the cluster.

## Consensus

Raft is responsible for reaching consensus on which transactions should be accepted. More accurately it reaches consensus on blocks, which contain transactions, to be inserted into the blockchain.

We include a "speculative mining" scheme so the minter can create new blocks before the previous has been accepted because we don't want to have to wait for Raft to reach consensus for each block before minting a new one (we currently provide a block latency (the max time before minting a new block) of 50ms). This introduces some complication in that it's possible we'll need to roll back / ignore blocks that aren't accepted by Raft. Thus, blocks accepted by raft can be either labeled "Valid" or "Invalid".

An example of the current minter (node 1) being partitioned, with node 2 taking over as minter.

```
 time                   block submissions
                   node 1                node 2
  |    [ 0xbeda Parent: 0xacaa ]
  |
  |    -- 1 is partitioned, 2 takes over --
  |
  |    [ 0x2c52 Parent: 0xbeda ] [ 0xf0ec Parent: 0xbeda ]
  |                              [ 0x839c Parent: 0xf0ec ]
  |
  |    -- 1 rejoins --
  |
  v                              [ 0x8b37 Parent: 0x8b37 ]
```

The resulting Raft log accepts every block, but due to being serialized after 0xf0ec, block 0x2c52 is marked as invalid and not applied.

```
[ 0xbeda Parent: 0xacaa Valid ]
[ 0xf0ec Parent: 0xbeda Valid ]
[ 0x839c Parent: 0xf0ec Valid ]
[ 0x2c52 Parent: 0xbeda Invalid ]
[ 0x8b37 Parent: 0x8b37 Valid ]
```

Note that each block is accepted by Raft and serialized in the log, but the Valid / Invalid marking is at a higher level, in our implementation. From Raft's point of view, each transaction is valid, but at the Quorum-Raft level, we can see unambiguously which blocks should be applied (and those blocks are exactly the same on each node).

To be clear, this is not the same as the "longest valid chain" mechanism from vanilla Ethereum. LVC is used to resolve forks in a network that doesn't have and can't have assigned leaders. We do have a single leader -- who is the only minter during normal operation -- but forks can be created when leadership changes. We mark blocks invalid iff they have the same parent as another block previously in the Raft log.

## Leadership Change

The leader can be thought of as a recommendation for who should mint. It's important that we never rely on leadership for correctness. During a transition it's possible that two nodes are both minting at the same time, and that's fine. Raft will serialize their blocks, so one of these competing minters will have the first block accepted by raft, and will take over as minter.

## Minting

We mint blocks no more frequently than every 50ms. In other words, when new transactions come in we will mint a new block immediately (so latency is low), but will only mint a block if it's been at least 50ms since the last block (so we don't flood raft with blocks). This scheme is a very simple compromise between raft throughput and transaction latency.

The miner currently uses a stripped-down proof of work, which we call "minting". The minter is responsible for bundling transactions into a block just like an Ethereum miner, but doesn't have to present a proof of work.

Raft is defined as a service. What is a service? An interface specified in node/service.go -- "an individual protocol that can be registered into a node". Other examples of services are `Ethereum`, `ReleaseService`, and `Whisper`.

## Speculative Minting

### Managing the chain

Our implementation has a few details which differ from the vanilla Ethereum blockchain.

We write "detached blocks" and set the new block head separately. A detached block is written whenever we receive a new block from the minter, but only set a block as the new head when it's accepted by raft.

We also need to be careful with future blocks. Since they might not be applied, we need to track the transactions which are bundled in future blocks to make sure they're submitted at some point even if their original carrying block is not applied.

### Terminology

#### "speculative block" / "speculative chain":

There is some delay between a block being minted by the leader and accepted into the blockchain (due to passing through Raft). In this time it's called a *speculative block* and the chain of speculative blocks is called the *speculative chain*. There is currently no limit on the number of speculative blocks in the speculative chain, ie a minter can create arbitrarily many blocks during a delay in Raft's acceptance.

#### "chain head event" / `core.ChainHeadEvent`:

An event posted when the blockchain head is updated.

#### "invalid ordering" / `InvalidRaftOrdering`:

Raft is a mechanism for ordering log messages, not determining their validity. It's possible for two nodes in the same cluster to create two "competing" blocks, only one of which will be applied, though both are valid.

For example, node *A* starts as minter, but experiences network failure as it mints a new block. Node *B* then takes over as minter and creates a new block which is accepted into the raft log and the blockchain. When node *A* rejoins, its block must make it into the raft log, but will be ignored by all parties (including *A*).

When a node receives a non-extending block from raft, it creates an `InvalidRaftOrdering` event. The event is ignored if another node minted the non-extending block, but it causes us to update our speculative state if we minted it. This means that we were previously the minter, but someone else took over.

### Pieces of State

* `parent`: The last speculative block.
* `proposedTxes`: The set of transactions which have been proposed to Raft in some block, but not yet committed.
* `unappliedBlocks`: A queue of blocks which have been proposed to raft but not yet committed.
  - When minting a new block we enqueue it at the end
  - `updateSpeculativeChainPerNewHead` is called on the node which minted a block to remove the oldest speculative block when it's accepted into the chain. This is more complicated in practice since other nodes could take over as minter and update the chain.
  - When an invalid ordering is found we unwind the queue by popping the most recent blocks from the right until we find the invalid block (this can remove blocks that depend on it).
* `expectedInvalidBlockHashes`: The set of blocks which build on an invalid block, but haven't passsed through Raft yet. We remove these as we get them back. When invalid blocks come back from raft we remove them from the speculative chain, so we use this set as a "guard" against trying to trim the speculative chain when we shouldn't.
