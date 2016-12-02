# Raft / Ethereum

## Introduction

This directory holds an implementation of [Raft](https://raft.github.io) (using [etcd](https://github.com/coreos/etcd)) as an alternative consensus mechanism for the Ethereum VM instead of the default proof-of-work.

Both Raft and Ethereum have a notion of node:

In Raft a node in normal operation is either a leader or a follower. There is a single leader for the entire cluster, which all log entries must go through. There's also the concept of a "Candidate", but only during leader election. We won't go into more detail about Raft here.

In vanilla Ethereum there is no such thing as a leader or follower. It's possible for any node in the cluster to mine a new block -- which is like being the leader for that round.

We impose a one-to-one correspondence between Raft and Ethereum nodes, so each Ethereum node is also a Raft node, and the leader of the Raft cluster is the only Ethereum node that can mint new blocks.

Ethereum | Raft
-------- | ----
minter   | leader
verifier | follower

Ethereum (and [RLPx](https://github.com/ethereum/devp2p/blob/master/rlpx.md)) [go into details about eth protocol manager] is responsible for sending transactions between nodes. Though it's used for all transactions, this is only important for transactions that originate at a follower. Transactions from the leader can be immediately minted in a block, put into Raft, and then put into the chain.

### Lifecycle of a Transactions

Note: We use "Etcd" the Etcd implementation of Raft, and "Raft" more broadly to also include Raft-Ethereum.

Let's follow the lifecycle of a typical transaction:

[ on a follower ]

1. The transaction is created by a user's rpc call.
2. Using the existing transaction propagation mechanisms in Ethereum, it's announced to all peers and transmitted to a subset of them.

[ on the leader ]

3. It reaches the leader, where it's minted in the next block (see `raftKickOffMining` and `commitNewWork`).
4. That triggers `NewMinedBlockEvent`, which the Raft protocol manager is subscribed to (`minedBlockSub`). The `minedBroadcastLoop` enqueues the new block in `ProtocolManager.proposeC`.
5. `serveInternal` is waiting at the other end of the channel. It proposes the block to Etcd.
6. The block reaches the `eventLoop` (which processes Raft events), as an `AppendEntries` message in `rd.Entries`. It, and all the other messages, are sent through `SendToPeers`. `SendToPeers` is responsible for sending all types of Raft messages (`AppendEntries` and configuration changes) as RLPx messages using the "raft" protocol name. To Ethereum there is only one type of raft message -- the protobuf-marshalled Raft message.

[ on followers ]

7. The message reaches a quorum of peers through the `handleRlpxPeerDiscovery` incoming message loop. It's enqueued in `peerMsgC`, and pulled off by `handlePeerMsgs`, which decodes it and `Step`s raft with it.
8. Etcd creates an `AppendEntries` response, which `eventLoop` handles, calls `SendToPeers`. The message should have only the leader in its `To` field, so `SendRaftPB` sends the protobuffed message back to the leader through RLPx.

[ on the leader ]

9. The leader handles each response through its incoming message loop. When a quorum have been received, Etcd considers the block committed, and it's moved into the `eventLoop` as a `CommittedEntries`.

[ on all ]

10. The block is enqueued in `logCommandC`, where it's handled by `handleLogCommands`. This method checks whether the block extends the chain (it's "valid", see below) and if not just ignores it. If so, the block is validated and set as the new head of the chain by `SetNewHeadBlock`.
11. A `ChainHeadEvent` is posted to notify listeners that a new block has been accepted. This is relevant to us because:
* It removes the newly minted transactions from `raftProposedTxes` (see below).
* It triggers `raftRequestFutureMining` in the (miner's) worker, telling Raft to try minting a new block.

The transaction is now available on all nodes in the cluster.

## Consensus

Raft is responsible for reaching consensus on which transactions should be accepted. More accurately it reaches consensus on blocks, which contain transactions, to be inserted into the blockchain.

We include a "speculative mining" scheme so the minter can create new blocks before the previous has been accepted because we don't want to have to wait for Raft to reach consensus for each block before mining a new one. This introduces some complication in that it's possible we'll need to roll back / ignore blocks that aren't accepted by Raft. Thus, blocks accepted by raft can be either labeled "Valid" or "Invalid".

An example of the current miner (node 1) being partitioned, with node 2 taking over as miner.

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
                                           
The resulting Raft log accepts every block, but due to being serialized after 0xf0ec, block 0x2c52 is marked as invalid and not applied.

[ 0xbeda Parent: 0xacaa Valid ]
[ 0xf0ec Parent: 0xbeda Valid ]
[ 0x839c Parent: 0xf0ec Valid ]
[ 0x2c52 Parent: 0xbeda Invalid ]
[ 0x8b37 Parent: 0x8b37 Valid ]

Note that each block is accepted by Raft and serialized in the log, but the Valid / Invalid marking is at a higher level, in our implementation. From Raft's point of view, each transaction is valid, but at the Raft-Ethereum level, we can see unambiguously which blocks should be applied (and those blocks are exactly the same on each node).

To be clear, this is not the same as the "longest valid chain" mechanism from vanilla Ethereum. LVC is used to resolve forks in a network that doesn't have and can't have assigned leaders. We do have a single leader -- who is the only miner during normal operation -- but forks can be created when leadership changes [fact check]. We mark blocks invalid iff they have the same parent as another block previously in the Raft log.

## Leader Change

[TODO]

## Mining

We mint blocks no more frequently than every 50ms. In other words, when new transactions come in we will mint a new block immediately (so latency is low), but will only mint a block if it's been at least 50ms since the last block (so we don't flood raft with blocks). This scheme is a very simple compromise between raft throughput and transaction latency.

Peers added to ethereum are also added to RLPx. [fact check]. The reverse is not yet implemented, but isn't necessary for correct operation.

The miner currently uses a stripped-down proof of work. One might naively expect we could not use a miner at all, but [fill in].

Raft is defined as a service. What is a service? An interface specified in node/service.go -- "an individual protocol that can be registered into a node". Other services are `Ethereum`, `ReleaseService`, and `Whisper`.

[TODO raftProposedTxes`]

## Managing the chain

Our implementation has a few details which differ from the vanilla Ethereum blockchain.

We write "detached blocks" and set the new block head separately. A detached block is written whenever we receive a new block from the miner, but only set a block as the new head when it's accepted by raft.

We also need to be careful with future blocks [finish]
