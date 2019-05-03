# Quorum -  Enterprise Ethereum Client

## What is Quorum?
Quorum is an Ethereum-based distributed ledger protocol that has been developed to provide industries such as finance, supply chain, retail, real estate, etc. with a permissioned implementation of Ethereum that supports transaction and contract privacy.  

Quorum includes a minimalistic fork of the [Go Ethereum client](https://github.com/ethereum/go-ethereum) (a.k.a geth), and as such, leverages the work that the Ethereum developer community has undertaken.  
 
The primary features of Quorum, and therefore extensions over public Ethereum, are:

* Transaction and contract privacy
* Multiple voting-based consensus mechanisms
* Network/Peer permissions management
* Higher performance

Quorum currently includes the following components:

* Quorum Node (modified Geth Client)
* Privacy Manager (Constellation/Tessera)
    * Transaction Manager
    * Enclave

!!! info "Background Reading"
    For more information on the design rationale and background to Quorum, please read the [**Quorum Whitepaper**](https://github.com/jpmorganchase/quorum/blob/master/docs/Quorum%20Whitepaper%20v0.2.pdf), view the [Hyperledger deck](https://drive.google.com/open?id=0B8rVouOzG7cOeHo0M2ZBejZTdGs) or watch the [presentation](https://drive.google.com/open?id=0B8rVouOzG7cOcDg4UkxqdTBacm8) given to the Hyperledger Project Technical Steering Committee meeting on 22-Sept-16. Also see quick overview of sending private transactions [here](https://vimeo.com/user5833792/review/210456729/8f70cfaaa5)


## Logical Architecture Diagram
![](Quorum%20Design.png)

## Design
### Public/Private State

Quorum supports dual state:

- Public state: accessible by all nodes within the network
- Private state: only accessible by nodes with the correct permissions

The difference is made through the use of transactions with encrypted (private) and non-encrypted payloads (public).
Nodes can determine if a transaction is private by looking at the `v` value of the signature.
Public transactions have a `v` value of `27` or `28`, private transactions have a value of `37` or `38`.

If the transaction is private, the node can only execute the transaction if it has the ability to access and decrypt the payload. Nodes who are not involved in the transaction do not have the private payload at all. As a result all nodes share a common public state which is created through public transactions and have a local unique private state.

This model imposes a restriction in the ability to modify state in private transactions.
Since it's a common use case for a (private) contract to read data from a public contract the virtual machine has the ability to jump into read only mode.
For each call from a private contract to a public contract the virtual machine will change to read only mode.
If the virtual machine is in read only mode and the code tries to make a state change the virtual machine stops execution and throws an exception.

The following transactions are allowed:

```
1. S -> A -> B
2. S -> (A) -> (B)
3. S -> (A) -> [B -> C]
```

and the following transaction are unsupported:

```
1. (S) -> A
2. (S) -> (A)
```

where:
- `S` = sender
- `(X)` = private
- `X` = public
- `->` = direction
- `[]` = read only mode

### State verification

To determine if nodes are in sync the public state root hash is included in the block.
Since private transactions can only be processed by nodes that are involved its impossible to get global consensus on the private state.

To overcome this issue the RPC method `eth_storageRoot(address[, blockNumber]) -> hash` can be used.
It returns the storage root for the given address at an (optional) block number.
If the optional block number is not given the latest block number is used.
The storage root hash can be on or off chain compared by the parties involved.

## Component Overview
### Quorum Node
The Quorum Node is intentionally designed to be a lightweight fork of geth in order that it can continue to take advantage of the R&D that is taking place within the ever growing Ethereum community.  To that end, Quorum will be updated in-line with future geth releases.

The Quorum Node includes the following modifications to geth:

 * Consensus is achieved with the Raft or Istanbul BFT consensus algorithms instead of using Proof-of-Work.
 * The P2P layer has been modified to only allow connections to/from permissioned nodes.
 * The block generation logic has been modified to replace the ‘global state root’ check with a new ‘global public state root’.
 * The block validation logic has been modified to replace the ‘global state root’ in the block header with the ‘global public state root’
 * The State Patricia trie has been split into two: a public state trie and a private state trie.
 * Block validation logic has been modified to handle ‘Private Transactions’
 * Transaction creation has been modified to allow for Transaction data to be replaced by encrypted hashes in order to preserve private data where required
 * The pricing of Gas has been removed, although Gas itself remains

### Constellation & Tessera
[Constellation](Privacy/Constellation/Constellation) and [Tessera](Privacy/Tessera/Tessera) are Haskell and Java implementations of a general-purpose system for submitting information in a secure way. They are comparable to a network of MTA (Message Transfer Agents) where messages are encrypted with PGP. It is not blockchain-specific, and are potentially applicable in many other types of applications where you want individually-sealed message exchange within a network of counterparties. The Constellation and Tessera modules consist of two sub-modules: 

* The Node (which is used for Quorum's default implementation of a `PrivateTransactionManager`) 
* The Enclave


#### Transaction Manager
Quorum’s Transaction Manager is responsible for Transaction privacy.  It stores and allows access to encrypted transaction data, exchanges encrypted payloads with other participant's Transaction Managers but does not have access to any sensitive private keys. It utilizes the Enclave for cryptographic functionality (although the Enclave can optionally be hosted by the Transaction Manager itself.)

The Transaction Manager is restful/stateless and can be load balanced easily.

For further details on how the Transaction Manager interacts with the Enclave, please refer [here](Privacy/Tessera/Tessera%20Services/Transaction%20Manager)

#### The Enclave

Distributed Ledger protocols typically leverage cryptographic techniques for transaction authenticity, participant authentication, and historical data preservation (i.e. through a chain of cryptographically hashed data.)  In order to achieve a separation of concerns, as well as to provide performance improvements through parallelization of certain crypto-operations, much of the cryptographic work including symmetric key generation and data encryption/decryption is delegated to the Enclave.  

The Enclave works hand in hand with the Transaction Manager to strengthen privacy by managing the encryption/decryption in an isolated way.  It holds private keys and is essentially a “virtual HSM” isolated from other components.

For further details on the Enclave, please refer [here](Privacy/Tessera/Tessera%20Services/Enclave).
