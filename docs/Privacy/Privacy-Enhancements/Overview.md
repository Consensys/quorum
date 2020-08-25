# Overview

Until Quorum 2.7.0 any node in Quorum network can alter the state of a contract deployed in the network even though it never had the contract byte code. Privacy enhancement feature is to prevent such 'non-party' interaction. This is done by the introduction of `ACOTH` which is short for "Affected Contract's Original Transaction's encrypted payload Hash". The `ACOTH` will be the proof going forward that the node is party to the transaction that created that contract and a non-party node will never have the `ACOTH` preventing it to alter the state. There are two flavours of privacy enhancement implemention ie., a) Counter party protection  and b) Private state validation.

## Counter Party Protection (PP)

This implementation doesn't allow non-party interaction on a private contract but allows state deviation i.e., it will allow nodes to maintain different state through private transaction to 'self' or 'subset of nodes'. 

## Private State Validation (PSV)

This implementation further prevents nodes from state deviation by failing private transactions to 'self' or 'subset of nodes' by sharing the list of recipients among all nodes which is validated against all subsequent transactions (in standard default mode only sending node knows the list of recipients). 

## Key Enhancements

### Privacy Flag

A new additional parameter `PrivacyFlag` in all Quorum [send](http://docs.goquorum.com/en/latest/Getting%20Started/api/) API methods , being passed from the client to enable privacy enhancement feature. This flag is an unsigned integer with the following values: 1 for PP and 3 for PSV transaction. If the flag is missing or zero, the transaction is assumed to be a 'non-privacy enhanced' standard transaction. 

### Privacy Metadata and Privacy Metadata Trie

Privacy Metadata is a new structure introduced in Quorum. It is saved in the quorum DB in the privacy metadata trie (which is linked to the private state - via root hash mappings). Privacy Metadata has the ACOTH and privacyFlag.

Privacy Metadata Trie is a parallel trie that stores the privacy metadata (and whatever extra data we may need) for the private contracts and is linked to the private state by root hash mappings. The records in the trie would be keyed by the contract account address (exactly the same as the contract accounts trie).

Each contract(account) that is created as 'PP' or 'PSV' would have such a structure created and attached to private state trie as it is essential in performing checks on future transactions affection those contracts.

## Configuration Changes

### Quorum

Genesis.json file is modified to include `privacyEnhancementsBlock`. The values for this should be set to an appropriate value in future (and should be initialised with same value across all the nodes in the network) by when the entire network would be ready to transact with privacy enhanced contracts. 

### Tessera

New flag `enableEnhancedPrivacy` has been added to Tessera config defaulting to `FALSE`, and can be enabled by adding the property to the config file the same way as other features. Refer sample configuration for further details.

## Enabling Privacy enhancement in Quorum Network 

For any given node privacy manager(Tessera) is started first and for that reason we allow Tessera node to be upgraded with privacy enhancement support ahead of Quorum upgrade. But when Quorum node is upgraded and geth reinitialised with `privacyEnhancementsBlock`, Quorum node will validate the version of Tessera running and will fail to start if Tessera is not running upgraded version. Quorum node will throw appropriate error message in the console suggesting users to upgrade Tessera first.

If a node wants to upgrade it's Tessera to privacy enhancement release (or further) to avail other features and fixes but not ready to upgrade Quorum, it can do so by not enabling `enableEnhancedPrivacy` in Tessera config. This will allow the node to reject PP and PSV transactions from other nodes until the node is ready to support privacy enhanced contracts.

## Backward compatability

### Quorum

An upgraded Quorum node can coexist on a network where other nodes are running on lower version of Quorum and thus supports node by node upgrade. But it cannot support privacy enhanced contracts until all interested nodes are upgraded and privacy 'enabled'. If a upgraded but privacy not 'enabled' node receives a PSV or PP transaction the node would log a `BAD BLOCK` error with “Privacy enhanced transaction received while privacy enhancements are disabled. Please check your node configuration.” error message. If the consensus algorithm is raft, the node would stop. For Istanbul, the node would keep trying to append the problematic block and reprint the above errors and it wont catch up with rest of nodes until restarted after upgrade.

### Tessera 

On any given node, Tessera can be upgraded to privacy enhanced release anytime but care must be taken as when to enable `enableEnhancedPrivacy` flag in Tessera config as once the flag is enabled, it will accept PSV and PP transactions and can cause the node to crash if Quorum node is not privacy enabled. The upgraded node can continue to communicate on Tessera nodes running on previous versions using `standard` private transactions. API versioning(add hyperlink later) to be introdued along with privacy enhancement enables the upgraded node to determine if the receiving node support privacy enhancement or not and fail the transaction then and there. 

## Tessera P2P communication changes

Refer [here](http://docs.goquorum.com/en/latest/Privacy/Lifecycle-of-a-private-transaction/) to refresh about Tessera P2P communication.

### Party Protection changes

To prevent non-party node from interacting with PP contracts new transactions must be submitted with ACOTH and `PrivacyFlag` from Quorum to Tessera. Tessera node would then generate a secure hashes (hash using new transaction ciphertext, original transaction ciphertext and original transaction master key) for ACOTH and include a) `PrivacyFlag`, b) ACOTH and c) ACOTH Securehash in the transaction payload shared betweek Tessera nodes.

### Private State Validation changes

Besides the ACOTH, a PSV transaction has an execution hash (merkle root) calculated from all the affected contract(s) resulting from the transaction simulation(at the time of submission) included from Quorum to Tessera. The d) execution hash and e) list of participants are also shared between Tessera nodes.

## Privacy Enhanced Transaction End to End Flow

Refer [here](http://docs.goquorum.com/en/latest/Privacy/Lifecycle-of-a-private-transaction/) for end to end flow of 'standard' private transaction.

![](Privacy_Enhancement.png)

In this example we walk through the flow of a private transaction on a 'privacy enhanced contract' between Nodes A & B.

1. User pushing a private transaction from Node A private for Node B

    - The transaction payload will include the `PrivacyFlag` with value `1` for PP and `3` for PSV contract

2. Node A reading the `PrivacyFlag` will run EVM simulation to get all affected contracts and the ACOTH(s) associated to contract account. For PSV transactions, it calculates an execution hash (merkle root) from all the affected contracts resulting from the transaction simulation.

3. Node A pushes the transaction payload, `PrivacyFlag`, ACOTH (& the merkle root for `PSV`) to Node A Tessera.

4. Node A Tessera would then generate secure hashes for ACOTH and use them to validate that the originating party has access to all relevant transactions. In addition for `PSV` it would also verify participants list against the list in the ACOTH transaction (as in `PSV` contract the recipient list is shared across all nodes party to the contract). If the list doesn't match it will return failure on `/send` to Node A Quorum.

5. Node A Tessera pushes to Node B Tessera encrypted payload, ACOTH <-> Securehash mapping (for `PSV` transaction it will further push `privateFor` list and merkle root).

6. Node B Tessera will compute and compare secure hash from Node A Tessera (for `PSV` it will also verify paricipant list of ACOTH against `privateFor` list). 

7. Node B Tessera will return SUCCESS to Node A Tessera always even if the compute and compare mismatched (to prevent Node A snipping out recipient of a contract) but it will not store the payload/ACOTH<->Securehash mapping based on the outcome.

8. Node A Tessera returns hash for the encrytped transaction payload to Node A Quorum

9. Node A mines the transaction across the network.

10. Node A & Node B being party to the contract will `/receive` decrypted payload, ACOTH (for `PSV` also merkle root) from respective Tessera Nodes.

11. Both Nodes execute the transaction and compare the ACOTH (and execution has for `PSV`) and update the transaction receipt accordingly to mark transaction execution completion..

**Note : If the EVM simulation impact more than one contract, all contracts should have the same `PrivacyFlag`, else transaction is rejected. All contracts ACOTH is included in the transmission and the Tessera node will create individual secure hash for each ACOTH.** 
