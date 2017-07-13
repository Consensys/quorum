# How a block gets added to the chain in quorum

A code walkthrough showing the making of a new block and how nodes receive and vote for it.

## Pre-requisites for all nodes:

The [BlockVoting service has been started](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/cmd/geth/main.go#L332) with a [block maker strategy in place](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/eth/block_voting.go#L11-L13)

## On the block maker

1. BlockMaker reaches its deadline waiting for a new block from others and [fires an event to trigger the creation of a new one itself](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/quorum/vote_strategy.go#L91)
2. The event is [picked up by the BlockVoting loop](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/quorum/block_voting.go#L261)
3. A new block is [constructed, inserted into the chain](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/quorum/block_voting.go#L348-L349)
4. A [ChainHeadEvent is published on BlockChain.chainEvents](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/blockchain.go#L1239)
5. A [NewMinedBlockEvent is published](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/quorum/block_voting.go#L353)
6. The NewMinedBlockEvent will eventually cause the ProtocolManager to [broadcast the new block](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/eth/handler.go#L759-L761)
7. Finally, the ChainHeadEvent will cause the node to [reset its pending state and vote for its own block](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/quorum/block_voting.go#L232-L236)


## On other nodes

1. Node [receives a message notifying about new block](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/eth/handler.go#L653)
2. [Enqueues it for import](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/eth/handler.go#L664)
3. The fetcher [detects a new import has been queued](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/eth/fetcher/fetcher.go#L293)
4. And [imports it](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/eth/fetcher/fetcher.go#L313) by [calling ProtocolManager.insertChain()](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/eth/handler.go#L178)
5. Which in turn calls BlockChain.InsertChain(), which checks that [the block was created by a block maker](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/block_validator.go#L334) (via [BlockValidator.ValidateBlock()](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/blockchain.go#L879) -> [ValidateHeader()](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/block_validator.go#L97) -> [ValidateExtraData()](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/block_validator.go#L329)),
and [call BlockValidator.ValidateState()](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/blockchain.go#L945)
to check that [its parent is the current head](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/block_validator.go#L172)
6. It will then send a [ChainHeadEvent on BlockChain.chainEvents](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/blockchain.go#L1239)
6. Which is processed in BlockVoting.run(), causing the Node's [pending state to be reset and the node to vote for the new block](https://github.com/jpmorganchase/quorum/blob/856c9fe43e3eb675678c55ceda7eb4d298a7ca7d/core/quorum/block_voting.go#L232-L236)

