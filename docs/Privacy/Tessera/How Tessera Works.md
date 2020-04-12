### Private Transaction Process Flow

Below is a description of how Private Transactions are processed in Quorum:

![Quorum Tessera Privacy Flow](https://github.com/jpmorganchase/tessera/raw/master/Tessera%20Privacy%20flow.jpeg)

In this example, Party A and Party B are party to Transaction AB, whilst Party C is not.

1. Party A sends a Transaction to their Quorum Node, specifying the Transaction payload and setting `privateFor` to be the public keys for Parties A and B (Party A is optional)
1. Party A's Quorum Node passes the Transaction on to its paired Transaction Manager, requesting that it encrypt and store the Transaction payload before forwarding it on to the recipients of the transaction (i.e. Party B)
1. Party A's Transaction Manager makes a call to its associated Enclave to encrypt the payload for the given recipients
1. Party A's Enclave encrypts the private transaction payload by: 
      
    1. generating a symmetric key (which will be referred to as *tx-key* from here on) and two random nonces 
    1. encrypting the Transaction payload with this tx-key and one of the nonces
    1. encrypting the tx-key separately for each recipient by:
    
        1. deriving the ECDH (elliptic-curve Diffie-Hellman) shared symmetric key (*shared-key*) from the sender's (A's) private key and the recipient's (B's) public key
        1. encrypting the tx-key with the shared-key and the other nonce
        1. repeating this for all recipients (i.e. for *n* recipients there will be *n* unique encrypted tx-keys) (the nonce is unchanged for each recipient as the shared-key being used changes)
        1. returning to the Transaction Manager: the encrypted transaction payload, all encrypted tx-keys, both nonces, and the public keys of the sender and all recipients

1. Party A's Transaction Manager stores the response from the Enclave and forwards to the private transaction recipients by:
    1. calculating the SHA3-512 hash of the encrypted payload (this acts as the unique identifier/primary key in the database)
    1. storing the hash, encrypted payload, all encrypted tx-keys, both nonces, and the public keys of the sender and all recipients in the database
    1. sending the response from the Enclave to each recipient by:
        1. sanitising the response for each recipient (i.e. removing all encrypted tx-keys except for the one intended for that recipient)
        1. serialising the data
        1. pushing the serialised data to the recipient (in this case Party B's Transaction Manager)
        1. ensuring the push was successful (if a single recipient fails to respond or returns an error then the process will stop here - i.e. it is a requirement that all recipients have stored the encrypted payload before the transaction can be propagated at the Quorum level and written to the blockchain)
        1. repeating this for all recipients 
1. Party A's Transaction Manager returns the hash of the encrypted payload to the Quorum Node.  Quorum replaces the `data` field of the Transaction with that hash, and changes the transaction's `v` value to `37` or `38`, thus marking the transaction as private and indicating that the transaction's `data` field represents the hash of an encrypted payload as opposed to executable EVM bytecode
1. The Transaction is then propagated to the rest of the network using the standard geth P2P Protocol
1. A block containing Transaction AB is created and distributed to each Quorum node in the network
1. In processing the block, all Quorum nodes attempt to process the Transaction.  Recognising that the transaction `data` is a hash due to the `v` value of `37` or `38`, each node will make a call to its Transaction Manager to determine if it is party to the transaction (i.e. there is an entry for the given hash in its database).  In this example, Party A & B's Transaction Managers will determine that they are party to the transaction whereas Party C's Transaction Manager will determine that it is not party
1. Party A & B's Transaction Managers make a call to their associated Enclaves to decrypt the payload
1. Party A & B's Enclaves decrypt the private transaction by:
    1. deriving the shared-key used in the encryption:
        1. as Party A was the sender of this transaction, it will derive the shared-key using its private key and the receiver's (B's) public key
        1. as Party B was a recipient of this transaction it will derive the shared-key using its private key and the sender's (A's) public key
    1. decrypting the tx-key with the shared-key and the encrypted data and nonce retrieved from the database
    1. decrypting the private transaction payload with the tx-key and the encrytped data and nonce retrieved from the database
    1. returning to the Transaction Manager: the decrypted private transaction data
1. The Transaction Manager's return their results to their Quorum nodes:
    1. Party A & B's Transaction Managers return the decrypted private transaction data to their Quorum nodes which can now execute the transaction as normal, thus updating their respective Private StateDB.  Quorum discards the decrypted private transaction data once used 
    1. Party C's Transaction Manager returns a 404 NOT FOUND to its Quorum node as it is not a recipient of the transaction.  Recognising that it is not party to this private transaction, the Quorum node will skip the execution of the transaction, so that no changes to its Private StateDB are made    


