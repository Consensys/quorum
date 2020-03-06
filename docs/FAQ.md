### Quorum FAQ

??? question "I've run into an issue with Quorum, where do I get support?"
    The [Quorum Slack channels](https://clh7rniov2.execute-api.us-east-1.amazonaws.com/Express/) are the best place to query the community and get immediate help.
 
    The Quorum engineering team monitors Slack as well as any issues raised on the Quorum GitHub repositories (e.g. [Quorum](https://github.com/jpmorganchase/quorum/), [Tessera](https://github.com/jpmorganchase/tessera), [Quorum-Examples](https://github.com/jpmorganchase/quorum-examples), etc.).  
    
??? question "How does Quorum achieve Transaction Privacy?"
    Quorum achieves Transaction Privacy by:
    
     1. Enabling transaction Senders to create a private transaction by marking who is privy to that transaction via the `privateFor` parameter
     2. Replacing the payload of a private transaction with a hash of the encrypted payload, such that the original payload is not visible to participants who are not privy to the transaction
     3. Storing encrypted private data off-chain in a separate component called the Transaction Manager (provided by [Constellation](https://github.com/jpmorganchase/constellation) or [Tessera](https://github.com/jpmorganchase/tessera)).  The Transaction Manager distributes the encrypted data to other parties that are privy to the transaction and returns the decrypted payload to those parties 
    
    Please see the [Transaction Processing](../Transaction%20Processing/Transaction%20Processing) page for more info.
    
??? question "How does Quorum achieve consensus on Private Transactions?"
    In standard Ethereum, all nodes process all transactions and so each node has the same state root.  In Quorum, nodes process all 'public' transactions (which might include reference data or market data contracts for example) but only process the private transactions that they are party to.  
    
    Quorum nodes maintain two Patricia Merkle Tries: one for private state and one for public state. As a result, block validation includes a **state** check on the new-to-Quorum `public state root`. Block validation also includes a check of the `global Transaction hash`, which is a hash of **all** Transactions in a block - private and public. This means that each node is able to validate that it has the same set of Transactions as other nodes.  Since the EVM is provably deterministic through the synchronized public state root, and that the Private Transaction inputs are known to be in sync across nodes (global Transaction Hash), private state synchronization across nodes can be implied.  In addition, Quorum provides an API call, `eth_storageRoot`, that returns the private state hash for a given transaction at a given block height, that can optionally be called at the application layer to specifically perform an off-chain state validation with a counterparty.
    
    Please see the [Quorum Consensus](../Consensus/Consensus) and [Transaction Processing](../Transaction%20Processing/Transaction%20Processing) pages for more info.

??? question "Are there any restrictions on the transaction size for private transactions (since they are encrypted)?"
    The only restriction is the gas limit on the transaction. Constellation/Tessera does not have a size limit (although maybe it should be possible to set one). If anything, performing large transactions as private transactions will improve performance because most of the network only sees hash digests. In terms of performance of transferring large data blobs between geographically distributed nodes, it would be equivalent performance to PGP encrypting the file and transferring it over http/https..so very fast. If you are doing sequential transactions then of course you will have to wait for those transfers, but there is no special overhead by the payload being large if you are doing separate/concurrent transactions, subject to network bandwidth limits. Constellation/Tessera does everything in parallel.

??? question "Should I include originating node in private transaction?"
    No, you should not. In Quorum, including originating node's `privateFor` will result in an error. If you would like to create a private contract that is visible to the originating node only please use this format: `privateFor: []` per https://github.com/jpmorganchase/quorum/pull/165

??? question "Is it possible to run a Quorum node without a Transaction Manager?"
    Starting a Quorum node with `PRIVATE_CONFIG=ignore` (instead of `PRIVATE_CONFIG=path/to/tm.ipc`) will start the node without a Transaction Manager. The node will not broadcast matching private keys (please ensure that there is no transaction manager running for it) and will be unable to participate in any private transactions.
    
??? question "Is there an official docker image for Quorum/Constellation/Tessera?"
    Yes! The [official docker containers](https://hub.docker.com/u/quorumengineering/):
    
    `quorumengineering/quorum:latest`
    `quorumengineering/constellation:latest`
    `quorumengineering/tessera:latest`
    
??? question "Can I create a network of Quorum nodes using different consensus mechanisms?"
    Unfortunately, that is not possible. Quorum nodes configured with raft will only be able to work correctly with other nodes running raft consensus. This applies to all other supported consensus algorithms.

??? info "Quorum version compatibility table"
    |                                     | Adding new node v2.0.x | Adding new node v2.1.x | Adding new node v2.2.x |
    | ----------------------------------- | ---------------------- | ---------------------- | ---------------------- |
    | Existing chain consisting of v2.0.x | <span style="color:green;">block sync<br /> public txn<br /> private txn</span>  | <span style="color:red;">block sync</span>  | <span style="color:red;">block sync</span> |
    | Existing chain consisting of v2.1.x | <span style="color:red;">block sync</span>  | <span style="color:green;">block sync<br /> public txn<br /> private txn</span> | <span style="color:green;">block sync<br /> public txn<br /> private txn</span> |
    | Existing chain consisting of v2.2.x | <span style="color:red;">block sync</span>  | <span style="color:green;">block sync<br /> public txn<br /> private txn</span> | <span style="color:green;">block sync<br /> public txn<br /> private txn</span> |

    **Note:** While every Quorum v2 client will be able to connect to any other v2 client, the usefullness will be severely degraded. <span style="color:red;">Red color</span> signifies that while connectivity is possible, <span style="color:red;">red colored</span> versions will be unable to send public or private txns to the rest of the net due to the EIP155 changes in the signer implemented in newer versions.

### Tessera FAQ

??? question "What does enabling 'disablePeerDiscovery' mean?"
    It means the node will only communicate with the nodes defined in the configuration file. Upto version 0.10.2, the nodes still accepts transactions from undiscovered nodes. From version 0.10.3 the node blocks all communication with undiscovered nodes.

??? info "Upgrading to Tessera version 0.10.+ from verion 0.9.+ and below"
    Due to 'database file unable to open' issue with H2 DB upgrade from version 1.4.196 direct to version 1.4.200 as explained  [here](https://github.com/h2database/h2database/issues/2263), our recommended mitigation strategy is to upgrade to version 1.4.199 first before upgrading to version 1.4.200 i.e., first upgrade to Tessera 0.10.0 before upgrading to higher versions. 

### Raft FAQ

??? question "Could you have a single- or two-node cluster? More generally, could you have an even number of nodes?"
    A cluster can tolerate failures that leave a quorum (majority) available. So a cluster of two nodes can't tolerate any failures, three nodes can tolerate one, and five nodes can tolerate two. Typically Raft clusters have an odd number of nodes, since an even number provides no failure tolerance benefit.

??? question "What happens if you don't assume minter and leader are the same node?"
    There's no hard reason they couldn't be different. We just co-locate the minter and leader as an optimization.
    
    * It saves one network call communicating the block to the leader.
    * It provides a simple way to choose a minter. If we didn't use the Raft leader we'd have to build in "minter election" at a higher level.

    Additionally there could even be multiple minters running at the same time, but this would produce contention for which blocks actually extend the chain, reducing the productivity of the cluster (see [Raft: Chain extension, races, and correctness](../Consensus/raft/#chain-extension-races-and-correctness) above).

??? question "I thought there were no forks in a Raft-based blockchain. What's the deal with "speculative minting"?"
    "Speculative chains" are not forks in the blockchain. They represent a series ("chain") of blocks that have been sent through Raft, after which each of the blocks may or may not actually end up being included in *the blockchain*.

??? question "Can transactions be reversed? Since raft log entries can be disregarded as "no-ops", does this imply transaction reversal?"
    No. When a Raft log entry containing a new block is disregarded as a "no-op", its transactions will remain in the transaction pool, and so they will be included in a future block in the chain.

??? question "What's the deal with the block timestamp being stored in nanoseconds (instead of seconds, like other consensus mechanisms)?"
    With raft-based consensus we can produce far more than one block per second, which vanilla Ethereum implicitly disallows (as the default timestamp resolution is in seconds and every block must have a timestamp greater than its parent). For Raft, we store the timestamp in nanoseconds and ensure it is incremented by at least 1 nanosecond per block.

??? question "Why do I see "Error: Number can only safely store up to 53 bits" when using web3js with Raft?"
    As mentioned above, Raft stores the timestamp in nanoseconds, so it is too large to be held as a number in javascript.
    You need to modify your code to take account of this. An example can be seen [here](https://github.com/jpmorganchase/quorum.js/blob/master/lib/index.js#L35).
    A future quorum release will address this issue.

??? info "Known Raft consensus node misconfiguration"
    Please see https://github.com/jpmorganchase/quorum/issues/410
