# Raft-based consensus for Ethereum/Quorum

## Introduction

This directory holds an implementation of a [Raft](https://raft.github.io)-based consensus mechanism (using [etcd](https://github.com/coreos/etcd)'s [Raft implementation](https://github.com/coreos/etcd/tree/master/raft)) as an alternative to Ethereum's default proof-of-work. This is useful for closed-membership/consortium settings where byzantine fault tolerance is not a requirement, and there is a desire for faster blocktimes (on the order of milliseconds instead of seconds) and transaction finality (the absence of forking.) Also, compared with QuorumChain, this consensus mechanism does not "unnecessarily" create empty blocks, and effectively creates blocks "on-demand."

When the `geth` binary is passed the `--raft` flag, the node will operate in "raft mode."

## Some implementation basics

Note: Though we use the etcd implementation of the Raft protocol, we speak of "Raft" more broadly to refer to the Raft protocol, and its use to achieve consensus for Quorum/Ethereum.

Both Raft and Ethereum have their own notion of a "node":

In Raft, a node in normal operation is either a "leader" or a "follower." There is a single leader for the entire cluster, which all log entries must flow through. There's also the concept of a "candidate", but only during leader election. We won't go into more detail about Raft here, because by design these details are opaque to applications built on it.

In vanilla Ethereum, there is no such thing as a "leader" or "follower." It's possible for any node in the network to mine a new block -- which is akin to being the leader for that round.

In Raft-based consensus, we impose a one-to-one correspondence between Raft and Ethereum nodes: each Ethereum node is also a Raft node, and by convention, the leader of the Raft cluster is the only Ethereum node that should mine (or "mint") new blocks. A minter is responsible for bundling transactions into a block just like an Ethereum miner, but does not present a proof of work.

Ethereum | Raft
-------- | ----
minter   | leader
verifier | follower

The main reasons we co-locate the leader and minter are (1) convenience, in that Raft ensures there is only one leader at a time, and (2) to avoid a network hop from a node minting blocks to the leader, through which all Raft writes must flow. Our implementation watches Raft leadership changes -- if a node becomes a leader it will start minting, and if a node loses its leadership, it will stop minting.

An observant reader might note that during raft leadership transitions, there could be a small period of time where more than one node might assume that it has minting duties; we detail how correctness is preserved in more detail later in this document.

We use the existing Ethereum p2p transport layer to communicate transactions between nodes, but we communicate blocks only through the Raft transport layer. They are created by the minter and flow from there to the rest of the cluster, always in the same order, via Raft.

When the minter creates a block, unlike in vanilla Ethereum where the block is written to the database and immediately considered the new head of the chain, we only insert the block or set it to be the new head of the chain once the block has flown through Raft. All nodes will extend the chain together in lock-step as they "apply" their Raft log.

From the point of view of Ethereum, Raft is integrated via an implementation of the [`Service`](https://godoc.org/github.com/jpmorganchase/quorum/node#Service) interface in [node/service.go](https://github.com/jpmorganchase/quorum/blob/master/node/service.go): "an individual protocol that can be registered into a node". Other examples of services are [`Ethereum`](https://godoc.org/github.com/jpmorganchase/quorum/eth#Ethereum), [`ReleaseService`](https://godoc.org/github.com/jpmorganchase/quorum/contracts/release#ReleaseService), and [`Whisper`](https://godoc.org/github.com/jpmorganchase/quorum/whisper/whisperv5#Whisper).

## The lifecycle of a transaction

Let's follow the lifecycle of a typical transaction:

#### on any node (whether minter or verifier):

1. The transaction is submitted via an RPC call to geth.
2. Using the existing (p2p) transaction propagation mechanism in Ethereum, the transaction is announced to all peers and, because our cluster is currently configured to use "static nodes," every transaction is sent to all peers in the cluster.

#### on the minter:

3. It reaches the minter, where it's included in the next block (see `mintNewBlock`) via the transaction pool.
4. Block creation triggers a [`NewMinedBlockEvent`](https://godoc.org/github.com/jpmorganchase/quorum/core#NewMinedBlockEvent), which the Raft protocol manager receives via its subscription `minedBlockSub`. The `minedBroadcastLoop` (in raft/handler.go) puts this new block to the `ProtocolManager.proposeC` channel.
5. `serveInternal` is waiting at the other end of the channel. Its job is to RLP-encode blocks and propose them to Raft. Once it flows through Raft, this block will likely become the new head of the blockchain (on all nodes.)

#### on every node:

6. _At this point, Raft comes to consensus and appends the log entry containing our block to the Raft log. (The way this happens at the Raft layer is that the leader sends an `AppendEntries` to all followers, and they acknowledge receipt of the message. Once the leader has received a quorum of such acknowledgements, it notifies each node that this new entry has been committed permanently to the log)._

7. Having crossed the network through Raft, the block reaches the `eventLoop` (which processes new Raft log entries.) It has arrived from the leader through `pm.transport`, an instance of [`rafthttp.Transport`](https://godoc.org/github.com/coreos/etcd/rafthttp#Transport).

8. The block is now handled by `applyNewChainHead`. This method checks whether the block extends the chain (i.e. it's parent is the current head of the chain; see below). If it does not extend the chain, it is simply ignored as a no-op. If it does extend chain, the block is validated and then written as the new head of the chain by [`InsertChain`](https://godoc.org/github.com/jpmorganchase/quorum/core#BlockChain.InsertChain).

9. A [`ChainHeadEvent`](https://godoc.org/github.com/jpmorganchase/quorum/core#ChainHeadEvent) is posted to notify listeners that a new block has been accepted. This is relevant to us because:
* It removes the relevant transaction from the transaction pool.
* It removes the relevant transaction from `speculativeChain`'s `proposedTxes` (see below).
* It triggers `requestMinting` in (minter.go), telling the node to schedule the minting of a new block if any more transactions are pending.

The transaction is now available on all nodes in the cluster with complete finality. Because Raft guarantees a single ordering of entries stored in its log, and because everything that is committed is guaranteed to remain so, there is no forking of the blockchain built upon Raft.

## Chain extension, races, and correctness

Raft is responsible for reaching consensus on which blocks should be accepted into the chain. In the simplest possible scenario, every subsequent block that passes through Raft becomes the new head of the chain.

However, there are rare scenarios in which we can encounter a new block that has passed through Raft that we can not crown as the new head of the chain. In these cases, when applying the raft log in-order, if we come across a block whose parent is not currently the head of the chain, we simply skip the log entry as a no-op.

The most common case where this can occur is during leadership changes. The leader can be thought of as a recommendation or proxy for who should mint -- and it is generally true that there is only a single minter -- but we do not rely on the maximum of one concurrent minter for correctness. During such a transition it's possible that two nodes are both minting for a short period of time. In this scenario there will be a race, the first block that successfully extends the chain will win, and the loser of the race will be ignored.

Consider the following example where this might occur, where Raft entries attempting to extend the chain are denoted like:

`[ 0xbeda Parent: 0xacaa ]`

Where `0xbeda` is the ID of new block, and `0xaa` is the ID of its parent. Here, the initial minter (node 1) is partitioned, and node 2 takes over as the minter.

```
 time                   block submissions
                   node 1                node 2
  |    [ 0xbeda Parent: 0xacaa ]
  |
  |   -- 1 is partitioned; 2 takes over as leader/minter --
  |
  |    [ 0x2c52 Parent: 0xbeda ] [ 0xf0ec Parent: 0xbeda ]
  |                              [ 0x839c Parent: 0xf0ec ]
  |
  |   -- 1 rejoins --
  |
  v                              [ 0x8b37 Parent: 0x8b37 ]
```

Once the partition heals, at the Raft layer node1 will resubmit `0x2c52`, and the resulting serialized log might look as follows:

```
[ 0xbeda Parent: 0xacaa - Extends! ]  (due to node 1)
[ 0xf0ec Parent: 0xbeda - Extends! ]  (due to node 2; let's call this the "winner")
[ 0x839c Parent: 0xf0ec - Extends! ]  (due to node 2)
[ 0x2c52 Parent: 0xbeda - NO-OP.   ]  (due to node 1; let's call this the "loser")
[ 0x8b37 Parent: 0x8b37 - Extends! ]  (due to node 2)
```

Due to being serialized after the "winner," the "loser" entry will not extend the chain, because its parent (`0xbeda`) is no longer at the head of the chain when we apply the entry. The "winner" extended the same parent (`0xbeda`) earlier (and then `0x839c` extended it further.)

Note that each block is accepted by Raft and serialized in the log, and that this "Extends"/"No-op" designation occurs at a higher level in our implementation. From Raft's point of view, each log entry is valid, but at the Quorum-Raft level, we choose which entries will be "used," and will actually extend the chain. This chain extension logic is deterministic: the same exact behavior will occur on every single node in the cluster, keeping the blockchain in sync.

Also note how our approach differs from the "longest valid chain" (LVC) mechanism from vanilla Ethereum. LVC is used to resolve forks in a network that is eventually consistent. Because we use Raft, the state of the blockchain is strongly consistent. There can not be forks in the Raft setting. Once a block has been added as the new head of the chain, it is done so for the entire cluster, and it is permanent.

## Minting frequency

As a default, we mint blocks no more frequently than every 50ms. When new transactions come in we will mint a new block immediately (so latency is low), but we will only mint a block if it's been at least 50ms since the last block (so we don't flood raft with blocks). This rate limiting achieves a balance between transaction throughput and latency.

This default of 50ms is configurable via the `--raftblocktime` flag to geth.

## Speculative minting

One of the ways our approach differs from vanilla Ethereum is that we introduce a new concept of "speculative minting." This is not strictly required for the core functionality of Raft-based Ethereum consensus, but rather it is an optimization that affords lower latency between blocks (or: faster transaction "finality.")

It takes some time for a block to flow through Raft (consensus) to become the head of the chain. If we synchronously waited for a block to become the new head of the chain before creating the new block, any transactions that we receive would take more time to make it into the chain.

In speculative minting we allow the creation of a new block (and its proposal to Raft) before its parent has made it all the way through Raft and into the blockchain.

Since this can happen repeatedly, these blocks (which each have a reference to their parent block) can form a sort of chain. We call this a "speculative chain."

During the course of operation that a speculative chain forms, we keep track of the subset of transactions in the pool that we have already put into blocks (in the speculative chain) that have not yet made it into the blockchain (and whereupon a [`core.ChainHeadEvent`](https://godoc.org/github.com/jpmorganchase/quorum/core#ChainHeadEvent) occurs.) These are called "proposed transactions" (see speculative_chain.go).

Per the presence of "races" (as we detail above), it is possible that a block somewhere in the middle of a speculative chain ends up not making into the chain. In this scenario an [`InvalidRaftOrdering`](https://godoc.org/github.com/jpmorganchase/quorum/raft#InvalidRaftOrdering) event will occur, and we clean up the state of the speculative chain accordingly.

There is currently no limit to the length of these speculative chains, but we plan to add support for this in the future. As a consequence, a minter can currently create arbitrarily many blocks back-to-back in a scenario where Raft stops making progress.

### State in a speculative chain

* `head`: The last-created speculative block. This can be `nil` if the last-created block is already included in the blockchain.
* `proposedTxes`: The set of transactions which have been proposed to Raft in some block, but not yet included in the blockchain.
* `unappliedBlocks`: A queue of blocks which have been proposed to Raft but not yet committed to the blockchain.
  - When minting a new block, we enqueue it at the end of this queue
  - `accept` is called to remove the oldest speculative block when it's accepted into the blockchain.
  - When an [`InvalidRaftOrdering`](https://godoc.org/github.com/jpmorganchase/quorum/raft#InvalidRaftOrdering) occurs, we unwind the queue by popping the most recent blocks from the "new end" of the queue until we find the invalid block. We must repeatedly remove these "newer" speculative blocks because they are all dependent on a block that we know has not been included in the blockchain.
* `expectedInvalidBlockHashes`: The set of blocks which build on an invalid block, but haven't passsed through Raft yet. We remove these as we get them back. When these non-extending blocks come back through Raft we remove them from the speculative chain. We use this set as a "guard" against trying to trim the speculative chain when we shouldn't.

## The Raft transport layer

We communicate blocks over the HTTP transport layer built in to etcd Raft. It's also (at least theoretically) possible to use p2p protocol built-in to Ethereum as a transport for Raft. In our testing we found the default etcd HTTP transport to be more reliable than the p2p (at least as implemented in geth) under high load.

Quorum listens on port 50400 by default for the raft transport, but this is configurable with the `--raftport` flag.

Default number of peers is set to be 25. Max number of peers is configurable with the `--maxpeers N` where N is expected size of the cluster. 

## Initial configuration, and enacting membership changes

Currently Raft-based consensus requires that all _initial_ nodes in the cluster are configured to list the others up-front as [static peers](https://github.com/ethereum/go-ethereum/wiki/Connecting-to-the-network#static-nodes). These enode ID URIs _must_ include a `raftport` querystring parameter specifying the raft port for each peer: e.g. `enode://abcd@127.0.0.1:30400?raftport=50400`. Note that the order of the enodes in the `static-nodes.json` file needs to be the same across all peers.

To remove a node from the cluster, attach to a JS console and issue `raft.removePeer(raftId)`, where `raftId` is the number of the node you wish to remove. For initial nodes in the cluster, this number is the 1-indexed position of the node's enode ID in the static peers list. Once a node has been removed from the cluster, it is permanent; this raft ID can not ever re-connect to the cluster in the future, and the party must re-join the cluster with a new raft ID.

To add a node to the cluster, attach to a JS console and issue `raft.addPeer(enodeId)`. Note that like the enode IDs listed in the static peers JSON file, this enode ID should include a `raftport` querystring parameter. This call will allocate and return a raft ID that was not already in use. After `addPeer`, start the new geth node with the flag `--raftjoinexisting RAFTID` in addition to `--raft`.

## FAQ

### Could you have a single- or two-node cluster? More generally, could you have an even number of nodes?

A cluster can tolerate failures that leave a quorum (majority) available. So a cluster of two nodes can't tolerate any failures, three nodes can tolerate one, and five nodes can tolerate two. Typically Raft clusters have an odd number of nodes, since an even number provides no failure tolerance benefit.

### What happens if you don't assume minter and leader are the same node?

There's no hard reason they couldn't be different. We just co-locate the minter and leader as an optimization.

* It saves one network call communicating the block to the leader.
* It provides a simple way to choose a minter. If we didn't use the Raft leader we'd have to build in "minter election" at a higher level.

Additionally there could even be multiple minters running at the same time, but this would produce contention for which blocks actually extend the chain, reducing the productivity of the cluster (see "races" above).

### I thought there were no forks in a Raft-based blockchain. What's the deal with "speculative minting"?

"Speculative chains" are not forks in the blockchain. They represent a series ("chain") of blocks that have been sent through Raft, after which each of the blocks may or may not actually end up being included in *the blockchain*.

### Can transactions be reversed? Since raft log entries can be disregarded as "no-ops", does this imply transaction reversal?

No. When a Raft log entry containing a new block is disregarded as a "no-op", its transactions will remain in the transaction pool, and so they will be included in a future block in the chain.
