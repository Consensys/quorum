
# Design

## Public/Private State

Quorum supports dual state:

- public state, accessible by all nodes within the network
- private state, only accessible by nodes with the correct permissions

The difference is made through the use of transactions with encrypted (private) and non-encrypted payloads (public).
Nodes can determine if a transaction is private by looking at the `v` value of the signature.
Public transactions have a `v` value of 27 or 28, private transactions have a value of 37 or 38.

If the transaction is private and the node has the ability to decrypt the payload it can execute the transaction.
Nodes who are not involved in the transaction cannot decrypt the payload and process the transaction.
As a result all nodes share a common public state which is created through public transactions and have a local unique private state.

This model imposes a restriction in the ability to modify state in private transactions.
Since it's a common use case for a (private) contract to read data from a public contract the virtual machine has the ability to jump into read only mode.
For each call from a private contract to a public contract the virtual machine will change to read only mode.
If the virtual machine is in read only mode and the code tries to make a state change the virtual machine stops execution and throws an exception.

The following transactions are allowed:

S: sender, (X): private, X: public, ->: direction, []: read only mode
```
1. S -> A -> B
2. S -> (A) -> (B)
3. S -> (A) -> [B -> C]
```
The following transaction are unsupported:

```
1. (S) -> A
2. (S) -> (A)
```

### State verification

To determine if nodes are in sync the public state root hash is included in the block.
Since private transactions can only be processed by nodes that are involved its impossible to get global consensus on the private state.
To overcome this issue the RPC method `eth_storageRoot(address[, blockNumber]) -> hash` can be used.
It returns the storage root for the given address at an (optional) block number.
If the optional block number is not given the latest block number is used.
The storage root hash can be on or off chain compared by the parties involved.
