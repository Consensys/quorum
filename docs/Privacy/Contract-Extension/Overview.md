# Overview
Once a private contract is deployed, it is only available on the nodes where it was initially deployed. This means any new node will not have access to the private contract as it does not have the code nor any of the state associated with it. In real life scenarios, we do see a need where a private contract deployed to set of initial participant nodes, need to be extended to a new node which has become part of the business flow. Contract state extension feature addresses this requirement. 

It should be noted that as a part of contract state extension only the state of the contract as of the time of extension is shared. This means that there is no past history of the contract, and attempting to view past history will not yield any result, as the new recipient was not party at that time. This also means that events are not shared either, as the transactions are not shared and no state transitions are calculated. 

## Flow

The overall flow of contract state extension is described below
![contract_state_extension](images/ContractStateExtension.png)  

In this example private contract which has been deployed between Nodes A and B. Now there is a need to extend this private contract to Node C.

1. User in node A proposes the extension of the contract, citing B and C's Private Transaction Manager(PTM) public keys as private participants of this extension, node B and C's public Ethereum key as a voter on this extension, and node C's PTM public key as the target receiver. 
    - **1a** - Node A submits Node C's public key to its own PTM as a "private transaction" and generates hash.
    - **1b** - Node A then creates the extension contract with user given inputs and substituting the Node C's public key with the returned PTM hash in step 1a. This is done to hide the recipient PTM key information.
    - **1c** - the private transaction payload is shared with Tessera nodes of B and C
    - **1d** - The public state is replicated across all nodes. Nodes A, B and C see an emitted log, and start watching the contract address that emitted the event for subsequent events that may happen

1.  Node A automatically approves the to extend by virtue of creating the extension contract. In the approval process:
    - **2a** - Node A generates a random hash H1 by submitting a self transaction to its own Tessera node that would be included in the extension contract as a proof of acceptance
    - **2b** - Node A submits the acceptance to extension contract with hash H1
    - **2c & 2d** - Private transaction payload is shared with Tessera nodes B and C. Public state is replicated across all nodes

1. User in node B owning ethereum public key marked as voter approves the extension to node C using Quorum apis 
    - **3a** - Node B generates a random hash H2 by submitting a self transaction to its own Tessera node that would be included in the extension contract as a proof of acceptance
    - **3b** - Node B submits the acceptance vote to extension contract with hash H2
    - **3c & 3d** - Private transaction payload is shared with Tessera nodes A and C. Public state is replicated across all nodes

1. Since the state sharing does not execute the transactions that generate the state 
   (in order to keep past history private), there is no proof that can be provided by the proposer that the state is correct. In order to remedy this, the receiver must accept the proposal for the contract as the proof. In this step, the user owning the ethereum public key of node C which was marked as voter approves the contract extension using Quorum apis
    - **4a** - Node C generates a random hash H2 by submitting a self transaction to its own Tessera node that would be included in the extension contract as a proof of acceptance
    - **4b** - Node C submits the acceptance vote to extension contract with hash H3
    - **4c & 4d** - Private transaction payload is shared with Tessera nodes A and B. Public state is replicated across all nodes

1. Node A monitors for all approvals to be completed for contract extension. Once all approvals are completed
    - **5a & 5b** - Node A fetches the state of the contract and sends it as a "private transaction" to Node C. It then submits the PTM hash of that state to the contract, including the recipient's PTM public key.
    - **5c** - Node A submits a transactions to mark completion of state share. This transaction emits a log which will get picked up by the receiver when processing the transaction
    - **5d & 5e** - Private transaction payload is shared with Tessera nodes C. Public state is replicated across all nodes

1. Node C monitors for state share event
    - **6a** - Upon noticing the state share event as a part of block processing, node C fetches the contract private state data from Tessera of node C
    - **6b** - Node C applies the fetched state to the contract address and becomes party of the private contract
  
## FAQs
  
### Why have voters on the contract?

Voting is used as a audit for who agreed to what. It is possible to not have any voters (except for the proposer and receiver), even if more participants do exist for the contract.