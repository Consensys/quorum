# Account/Key Management

Cryptographic keys are an essential component of a Quorum network.  Quorum uses keys to create digital signatures which verify a sender's identity and prevent message tampering.  The Privacy Manager uses keys to encrypt private transaction data.

Both Quorum and the Privacy Manager use user-provided asymmetric key pairs.  Each key pair consists of a public key and a private key.  The public key can be shared freely, but **the private key should never be shared**.

* Quorum derives the account address from the public key by taking the last 20 bytes of its keccak256 hash
* The Privacy Manager uses the public key as an identifier for the target nodes of a private transaction (i.e. the `privateFor` transaction field)

Key management determines how Quorum/the Privacy Manager stores and uses private keys.  The corresponding sections for [Quorum](../Quorum/Overview), [Tessera](../Tessera/Overview), and [Constellation](../Constellation/Overview) detail the methods available to each. 
  