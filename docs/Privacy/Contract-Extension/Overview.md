# Overview
Once a private contract is deployed, it is only available on the nodes where it was initially deployed. This means any new node will not have access to the private contract as it does not have the code nor any of the state associated with it. In real life scenarios, we do see a need where a private contract deployed to set of initial participant nodes, need to be extended to a new node which has become part of the business flow. Contract state extension feature addresses this requirement. 

It should be noted that as a part of contract state extension only the state of the contract as of the time of extension is shared. This means that there is no past history of the contract, and attempting to view past history will not yield any result, as the new recipient was not party at that time. This also means that events are not shared either, as the transactions are not shared and no state transitions are calculated. 

## Flow

The overall flow of contract state extension is described below
![contract_state_extension](images/ContractStateExtension.png)  

In this example private contract is being extended from Nodes A to Node B. 

1. User in node A proposes the extension of the contract, citing B's Private Transaction Manager(PTM) public keys as private participants of this extension, node B's public Ethereum key as a receiving address of this extension, and node B's PTM public key as the target receiver. 
    - **1a** - Node A creates the extension contract with user given inputs and Node B's PTM public key
    - **1b** - the private transaction payload is shared with Tessera of node B.
    - **1c** - The public state is replicated across both nodes. Nodes A and B  see an emitted log, and start watching the contract address that emitted the event for subsequent events that may happen

1.  Node A automatically approves the contract extension by virtue of creating the extension contract. In the approval process:
    - **2a** - Node A submits the approval to extension contract 
    - **2c & 2d** - Private transaction payload is shared with Tessera node B. Public state is replicated across all nodes

1. Since the state sharing does not execute the transactions that generate the state 
   (in order to keep past history private), there is no proof that can be provided by the proposer that the state is correct. In order to remedy this, the receiver must accept the proposal for the contract as the proof. In this step, the user owning the ethereum public key of node B which was marked as receiving address, approves the contract extension using Quorum apis
    - **3a** - Node B submits the acceptance vote to extension contract
    - **3c & 3d** - Private transaction payload is shared with Tessera nodes A. Public state is replicated across all nodes

1. Node A monitors for acceptance of contract extension by Node B. Once accepted
    - **4a & 4b** - Node A fetches the state of the contract and sends it as a "private transaction" to Node B. It then submits the PTM hash of that state to the contract, including the recipient's PTM public key.
    - **4c** - Node A submits a transactions to mark completion of state share. This transaction emits a log which will get picked up by the receiver when processing the transaction
    - **4d & 4e** - Private transaction payload is shared with Tessera nodes B. Public state is replicated across all nodes

1. Node B monitors for state share event
    - **5a** - Upon noticing the state share event as a part of block processing, node B fetches the contract private state data from Tessera of node B
    - **5b** - Node B applies the fetched state to the contract address and becomes party of the private contract
    

## Note
If the network is running with [Enhanced network permissioning](http://docs.goquorum.com/en/latest/Permissioning/Enhanced%20Permissions%20Model/Overview/), then:
* Initiation of contract extension can only be done by a network admin or org admin account

* Similarly on the receiving node, contract extension can be accepted only by an network admin or org admin account
 