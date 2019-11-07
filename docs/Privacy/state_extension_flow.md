# Contract state extension flow

---

In this example, we have a smart contract between two existing participants, A & B, who want the contract shared to a 
new recipient, C.

1. Node A proposes the extension of the contract, citing B and C's Private Transaction Manager(PTM) public keys as 
private participants of this extension, node B's public Ethereum key as a voter on this extension, and node C's PTM 
public key as the target receiver.
Before creating the contract, Node A submits Node C's public key to its own PTM as a "private transaction", actually 
using the PTM as a private store.
Node A then creates the contract with this information, substituting the Node C's public key with the returned PTM hash.

2. Nodes A, B and C see an emitted log, and start watching the contract address that emitted the event for subsequent 
events that may happen. It will also write the details of the extension to disk in case of restarts, and make the 
information available over the RPC API.

3a. Node A and B must now submit their vote to the extension, indicating whether they want the extension to take place 
or not. They should include the voters PTM public keys to make sure the transaction is visible to all.

3b. Node C must accept the proposal. Since the state sharing does not execute the transactions that generate the state 
(in order to keep past history private), there is no proof that can be provided by the proposer that the state is 
correct. In order to remedy this, the receiver must accept the proposal for the contract as the proof. It also submits 
a self-transaction to its own PTM, and includes the returned hash in the contract, as a proof that it called the 
accept function when looked at a later date. This prevents someone else calling "accept" impersonating Node C.

4. Once all votes and acceptances have been gathered, Node A fetches the state of the contract and sends it as a 
"private transaction" to Node C. It then submits the PTM hash of that state to the contract, including the recipient's 
PTM public key. This transaction emits a log which will get picked up by the receiver when processing the transaction, 
which signals that the state should be retrieved and set appropriately.

# Q&A

# Why have voters on the contract?
Voting is used as a audit for who agreed to what. Is it possible to not have any voters (except for the proposer and 
receiver), even if more participants do exist for the contract.