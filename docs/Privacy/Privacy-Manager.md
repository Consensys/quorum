# Privacy Manager

* Private data shared only with a subset of nodes
* PGP https://protonmail.com/blog/what-is-pgp-encryption/
* NaCl
* Alternate encryption mechanisms

## Implementations
* [Tessera](../../Privacy/Tessera/Tessera) is a production-ready implementation of Quorum's privacy manager.  It is undergoing active development with new features being added regularly.

* [Constellation](../../Privacy/Constellation/Constellation) is the reference implementation of Quorum's privacy manager.  It is still supported but no longer undergoing active development of new features.  

## Private transaction flow
See [How Tessera Works](../../Privacy/Tessera/How%20Tessera%20Works) for a description of the lifecycle of a private transaction.

## Encryption process flow

The end-to-end flow for the submission of a private transaction to a Quorum network, including the process of data encryption, is described in [How Tessera Works](Tessera/How%20Tessera%20Works).  This section provides a more detailed explanation of the Privacy Manager's encryption process.

Encryption of private transaction data is performed by the Privacy Manager's enclave.    

The enclave performs the following steps to encrypt data:

1. generating a random master key (RMK) and a random Nonce 
1. encrypting the Transaction payload with the nonce and RMK from step a.
1. iterating through the list of transaction recipients, in this case parties A and B, and encrypting the RMK from a. with the shared key derived from Party A's private key and the recipient's public key, along with another randomly generated nonce. Each encrypted RMK is unique for each recipient and will only be shared with the respective recipient along with encrypted payload.
1. returning the encrypted payload from step b. and all encrypted RMKs from step c. to the Transaction Manager   
