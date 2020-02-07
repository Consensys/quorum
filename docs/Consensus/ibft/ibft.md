# IBFT Consensus Overview

## Introduction

Istanbul Byzantine Fault Tolerant (IBFT) consensus is inspired by Castro-Liskov 99 [paper](http://pmg.csail.mit.edu/papers/osdi99.pdf). IBFT inherits from the original PBFT by using a 3-phase consensus, `PRE-PREPARE`, `PREPARE` and `COMMIT`. The system can tolerate at most `F` faulty nodes in a `N` validator network, where `N = 3F + 1`.  

## Implementation

### Terminology

- `Validator`: Block validation participant.
- `Proposer`: A block validation participant that is chosen to propose block in a consensus round.
- `Round`: Consensus round. A round starts with the proposer creating a block proposal and ends with a block commitment or round change.
- `Proposal`: New block generation proposal which is undergoing consensus processing.
- `Sequence`: Sequence number of a proposal. A sequence number should be greater than all previous sequence numbers. Currently each proposed block height is its associated sequence number.
- `Backlog`: The storage to keep future consensus messages.
- `Round state`: Consensus messages of a specific sequence and round, including pre-prepare message, prepare message, and commit message.
- `Consensus proof`: The commitment signatures of a block that can prove the block has gone through the consensus process.
- `Snapshot`: The validator voting state from last epoch.

### Consensus
Istanbul BFT Consensus protocol begins at Round `0` with the validators picking a proposer from themselves in a round robin fashion. The proposer will then propose a new block proposal and broadcast it along with the `PRE-PREPARE` message. Upon receiving the `PRE-PREPARE` message from the proposer, other validators validate the incoming proposal and enter the state of `PRE-PREPARED` and broadcast `PREPARE` message. This step is to make sure all validators are working on the same sequence and on the same round. When `ceil(2N/3)` of `PREPARE` messages is received by the validator from other validators, the validator switches to the state of `PREPARED` and broadcasts `COMMIT` message. This step is to inform other validators that it accepts the proposed block and is going to insert the block to the chain. Lastly, validators wait for `ceil(2N/3)` of `COMMIT` messages to enter `COMMITTED` state and then append the block to the chain.

Blocks in Istanbul BFT protocol are final, which means that there are no forks and any valid block must be somewhere in the main chain. To prevent a faulty node from generating a totally different chain from the main chain, each validator appends `ceil(2N/3)` of received `COMMIT` signatures to `extraData` field in the header before inserting it into the chain. Thus all blocks are self-verifiable. However, the dynamic `extraData` would cause an issue on block hash calculation. Since the same block from different validators can have different set of `COMMIT` signatures, the same block can have different block hashes as well. To solve this, we calculate the block hash by excluding the `COMMIT` signatures part. Therefore, we can still keep the block/block hash consistency as well as put the consensus proof in the block header.

#### Consensus States
Istanbul BFT is a state machine replication algorithm. Each validator maintains a state machine replica in order to reach block consensus. Various states in IBFT consensus are,

- `NEW ROUND`: Proposer to send new block proposal. Validators wait for `PRE-PREPARE` message.
- `PRE-PREPARED`: A validator has received `PRE-PREPARE` message and broadcasts `PREPARE` message. Then it waits for `ceil(2N/3)` of `PREPARE` or `COMMIT` messages.
- `PREPARED`: A validator has received `ceil(2N/3)` of `PREPARE` messages and broadcasts `COMMIT` messages. Then it waits for `ceil(2N/3)` of `COMMIT` messages.
- `COMMITTED`: A validator has received `ceil(2N/3)` of `COMMIT` messages and is able to insert the proposed block into the blockchain.
- `FINAL COMMITTED`: A new block is successfully inserted into the blockchain and the validator is ready for the next round.
- `ROUND CHANGE`: A validator is waiting for `ceil(2N/3)` of `ROUND CHANGE` messages on the same proposed round number.

**State Transitions**:
![State Transitions](images/IBFTStateTransition.png)

- `NEW ROUND` -> `PRE-PREPARED`:
    - **Proposer** collects transactions from txpool.
    - **Proposer** generates a block proposal and broadcasts it to validators. It then enters the `PRE-PREPARED` state.
    - Each **validator** enters `PRE-PREPARED` upon receiving the `PRE-PREPARE` message with the following conditions:
        - Block proposal is from the valid proposer.
        - Block header is valid.
        - Block proposal's sequence and round match the **validator**'s state.
    - **Validator** broadcasts `PREPARE` message to other validators.
- `PRE-PREPARED` -> `PREPARED`:
    - Validator receives `ceil(2N/3)` of valid `PREPARE` messages to enter `PREPARED` state. Valid messages conform to the following conditions:
        - Matched sequence and round.
        - Matched block hash.
        - Messages are from known validators.
    - Validator broadcasts `COMMIT` message upon entering `PREPARED` state.
- `PREPARED` -> `COMMITTED`:
    - **Validator** receives `ceil(2N/3)` of valid `COMMIT` messages to enter `COMMITTED` state. Valid messages conform to the following conditions:
        - Matched sequence and round.
        - Matched block hash.
        - Messages are from known validators.
- `COMMITTED` -> `FINAL COMMITTED`:
    - **Validator** appends `ceil(2N/3)` commitment signatures to `extraData` and tries to insert the block into the blockchain.
    - **Validator** enters `FINAL COMMITTED` state when insertion succeeds.
- `FINAL COMMITTED` -> `NEW ROUND`:
    - **Validators** pick a new **proposer** and begin a new round timer.

#### Round change flow

- There are three conditions that would trigger `ROUND CHANGE`:
    - Round change timer expires.
    - Invalid `PREPREPARE` message.
    - Block insertion fails.
- When a validator notices that one of the above conditions applies, it broadcasts a `ROUND CHANGE` message along with the proposed round number and waits for `ROUND CHANGE` messages from other validators. The proposed round number is selected based on following condition:
    - If the validator has received `ROUND CHANGE` messages from its peers, it picks the largest round number which has `F + 1` of `ROUND CHANGE` messages.
    - Otherwise, it picks `1 + current round number` as the proposed round number.
- Whenever a validator receives `F + 1` of `ROUND CHANGE` messages on the same proposed round number, it compares the received one with its own. If the received is larger, the validator broadcasts `ROUND CHANGE` message again with the received number.
- Upon receiving `ceil(2N/3)` of `ROUND CHANGE` messages on the same proposed round number, the **validator** exits the round change loop, calculates the new **proposer**, and then enters `NEW ROUND` state.
- Another condition that a validator jumps out of round change loop is when it receives verified block(s) through peer synchronization.

#### Proposer selection

Currently we support two policies: **round robin** and **sticky proposer**.

- Round robin: Round robin is the default proposer selection policy. In this setting proposer will change in every block and round change.
- Sticky proposer: in a sticky proposer setting, proposer will change only when a round change happens.

#### Validator list voting

Istanbul BFT uses a similar validator voting mechanism as Clique and copies most of the content from Clique [EIP](https://github.com/ethereum/EIPs/issues/225). Every epoch transaction resets the validator voting, meaning any pending votes for adding/removing a validator are reset.

For all transactions blocks:

- Proposer can cast one vote to propose a change to the validators list.
- Only the latest proposal per target beneficiary is kept from a single validator.
- Votes are tallied live as the chain progresses (concurrent proposals allowed).
- Proposals reaching majority consensus `VALIDATOR_LIMIT` come into effect immediately.
- Invalid proposals are not to be penalized for client implementation simplicity.
- A proposal coming into effect entails discarding all pending votes for that proposal (both for and against) and starts with a clean slate.

#### Future message and backlog

In an asynchronous network environment, one may receive future messages which cannot be processed in the current state. For example, a validator can receive `COMMIT` messages on `NEW ROUND`. We call this kind of message a "future message." When a validator receives a future message, it will put the message into its **backlog** and try to process later whenever possible.

#### Constants
Istanbul BFT define the following constants

- `EPOCH_LENGTH`: Default: 30000 blocks. Number of blocks after which to checkpoint and reset the pending votes.
- `REQUEST_TIMEOUT`: Timeout for each consensus round before firing a round change in millisecond.
- `BLOCK_PERIOD`: Minimum timestamp difference in seconds between two consecutive blocks.
- `PROPOSER_POLICY`: Proposer selection policy, defaults to round robin.
- `ISTANBUL_DIGEST`: Fixed magic number `0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365` of `mixDigest` in block header for Istanbul block identification.
- `DEFAULT_DIFFICULTY`: Default block difficulty, which is set to `0x0000000000000001`.
- `EXTRA_VANITY`: Fixed number of extra-data prefix bytes reserved for proposer vanity.
    - Suggested `32` bytes to retain the current extra-data allowance and/or use.
- `NONCE_AUTH`: Magic nonce number `0xffffffffffffffff` to vote on adding a validator.
- `NONCE_DROP`: Magic nonce number `0x0000000000000000` to vote on removing a validator.
- `UNCLE_HASH`: Always `Keccak256(RLP([]))` as uncles are meaningless outside of PoW.
- `PREPREPARE_MSG_CODE`: Fixed number `0`. Message code for `PREPREPARE` message.
- `PREPARE_MSG_CODE`: Fixed number `1`. Message code for `PREPARE` message.
- `COMMIT_MSG_CODE`: Fixed number `2`. Message code for `COMMIT` message.
- `ROUND_CHANGE_MSG_CODE`: Fixed number `3`. Message code for `ROUND CHANGE` message 
- `VALIDATOR_LIMIT`: Number of validators to pass an authorization or de-authorization proposal. 
    - Must be `floor(N / 2) + 1` to enforce majority consensus on a chain.

#### Block Header
Istanbul BFT does not add new block header fields. Instead, it follows Clique in repurposing the `ethash` header fields as follows:

- `nonce`: Proposer proposal regarding the account defined by the beneficiary field.
    - Should be `NONCE_DROP` to propose deauthorizing beneficiary as an existing validator.
    - Should be `NONCE_AUTH` to propose authorizing beneficiary as a new validator.
    - **Must** be filled with zeroes, `NONCE_DROP` or `NONCE_AUTH`
- `mixHash`: Fixed magic number `0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365` for Istanbul block identification.
- `ommersHash`: Must be `UNCLE_HASH` as uncles are meaningless outside of PoW.
- `timestamp`: Must be at least the parent timestamp + `BLOCK_PERIOD`
- `difficulty`: Must be filled with `0x0000000000000001`.
- `extraData`: Combined field for signer vanity and RLP encoded Istanbul extra data, where Istanbul extra data contains validator list, proposer seal, and commit seals. Istanbul extra data is defined as follows:
    ```
    type IstanbulExtra struct {
            Validators    []common.Address 	//Validator addresses
            Seal          []byte	        //Proposer seal 65 bytes
            CommittedSeal [][]byte          //Committed seal, 65 * len(Validators) bytes
    }
    ```
    Thus the `extraData` would be in the form of `EXTRA_VANITY | ISTANBUL_EXTRA` where `|` represents a fixed index to separate vanity and Istanbul extra data (not an actual character for separator).
    
    - First `EXTRA_VANITY` bytes (fixed) may contain arbitrary proposer vanity data.
    - `ISTANBUL_EXTRA` bytes are the RLP encoded Istanbul extra data calculated from `RLP(IstanbulExtra)`, where `RLP()` is RLP encoding function, and `IstanbulExtra` is the Istanbul extra data.
        - `Validators`: The list of validators, which **must** be sorted in ascending order.
        - `Seal`: The proposer's signature sealing of the header.
        - `CommittedSeal`: The list of commitment signature seals as consensus proof.

#### Block hash, proposer seal and committed seals
The Istanbul block hash calculation is different from the `ethash` block hash calculation due to the following reasons:

1. The proposer needs to put proposer's seal in `extraData` to prove the block is signed by the chosen proposer.
2. The validators need to put `ceil(2N/3)` of committed seals as consensus proof in `extraData` to prove the block has gone through consensus.

The calculation is still similar to the `ethash` block hash calculation, with the exception that we need to deal with `extraData`. We calculate the fields as follows:

##### Proposer seal calculation
By the time of proposer seal calculation, the committed seals are still unknown, so we calculate the seal with those unknowns empty. The calculation is as follows:

- `Proposer seal`: `SignECDSA(Keccak256(RLP(Header)), PrivateKey)`
- `PrivateKey`: Proposer's private key.
- `Header`: Same as `ethash` header only with a different `extraData`.
- `extraData`: `vanity | RLP(IstanbulExtra)`, where in the `IstanbulExtra`, `CommittedSeal` and `Seal` are empty arrays.

##### Block hash calculation
While calculating block hash, we need to exclude committed seals since that data is dynamic between different validators. Therefore, we make `CommittedSeal` an empty array while calculating the hash. The calculation is:

- `Header`: Same as `ethash` header only with a different `extraData`.
- `extraData`: `vanity | RLP(IstanbulExtra)`, where in the `IstanbulExtra`, `CommittedSeal` is an empty array.

##### Consensus proof
Before inserting a block into the blockchain, each validator needs to collect `ceil(2N/3)` of committed seals from other validators to compose a consensus proof. Once it receives enough committed seals, it will fill the `CommittedSeal` in `IstanbulExtra`, recalculate the `extraData`, and then insert the block into the blockchain. **Note** that since committed seals can differ by different sources, we exclude that part while calculating the block hash as in the previous section.

Committed seal calculation:

Committed seal is calculated by each of the validators signing the hash along with `COMMIT_MSG_CODE` message code of its private key. The calculation is as follows:

- `Committed seal`: `SignECDSA(Keccak256(CONCAT(Hash, COMMIT_MSG_CODE)), PrivateKey)`.
- `CONCAT(Hash, COMMIT_MSG_CODE)`: Concatenate block hash and `COMMIT_MSG_CODE` bytes.
- `PrivateKey`: Signing validator's private key.


## Provenance
Istanbul BFT implementation in Quorum is based on [EIP 650](https://github.com/ethereum/EIPs/issues/650). It has been updated since the EIP was opened to resolve safety issues by introducing locking.
