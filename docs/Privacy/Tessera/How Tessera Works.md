### Private Transaction Process Flow

Below is a description of how Private Transactions are processed in Quorum:

![Quorum Tessera Privacy Flow](https://github.com/jpmorganchase/tessera/raw/master/Tessera%20Privacy%20flow.jpeg)

In this example, Party A and Party B are party to Transaction AB, whilst Party C is not.

1. Party A sends a Transaction to their Quorum Node, specifying the Transaction payload and setting `privateFor` to be the public keys for Parties A and B (Party A is optional)
1. Party A's Quorum Node passes the Transaction on to its paired Transaction Manager, requesting for it to store the Transaction payload
1. Party A's Transaction Manager makes a call to its associated Enclave to validate the sender and encrypt the payload
1. Party A's Enclave checks the private key for Party A and, once validated, performs the Transaction conversion. This entails: 
      
    1. generating a random master key (RMK) and a random Nonce 
    1. encrypting the Transaction payload with the nonce and RMK from step a.
    1. iterating through the list of transaction recipients, in this case parties A and B, and encrypting the RMK from a. with the shared key derived from Party A's private key and the recipient's public key, along with another randomly generated nonce. Each encrypted RMK is unique for each recipient and will only be shared with the respective recipient along with encrypted payload.
    1. returning the encrypted payload from step b. and all encrypted RMKs from step c. to the Transaction Manager   

1. Party A's Transaction Manager calculates the SHA3-512 hash of the encrypted payload then stores the encrypted payload and encrypted RMKs against the hash in the database
1. Party A's Transaction Manager then securely transfers (via HTTPS) the encrypted payload, and RMK that has been encrypted with shared key from previous step 4.c, the nonce's to Party B's Transaction Manager.  Party B's Transaction Manager responds with an Ack/Nack response. Note that if Party A does not receive a response/receives a Nack from Party B then the Transaction will not be propagated to the network.  It is a prerequisite for the recipients to store the communicated payload.
1. Once the data transmission to Party B's Transaction Manager has been successful, Party A's Transaction Manager returns the hash to the Quorum Node which then replaces the Transaction's original payload with that hash, and changes the transaction's `V` value to 37 or 38, which will indicate to other nodes that this hash represents a private transaction with an associated encrypted payload as opposed to a public transaction with nonsensical bytecode.
1. The Transaction is then propagated to the rest of the network using the standard Ethereum P2P Protocol.
1. A block containing Transaction AB is created and distributed to each Party on the network.
1. In processing the block, all Parties will attempt to process the Transaction.  Each Quorum node will recognise a `V` value of 37 or 38, identifying the Transaction as one whose payload requires decrypting, and make a call to their local Transaction Manager to determine if they hold the Transaction (using the hash as the index to look up).
1. Since Party C does not hold the Transaction, it will receive a `NotARecipient` message and will skip the Transaction - it will not update its Private StateDB.  Party A and B will look up the hash in their local Transaction Managers and identify that they do hold the Transaction. Each will then make a call to its Enclave, passing in the Encrypted Payload, Encrypted symmetric key and Signature.
1. The Enclave validates the signature and then decrypts the symmetric key using the Party's private key that is held in The Enclave, decrypts the Transaction Payload using the now-revealed symmetric key and returns the decrypted payload to the Transaction Manager.
1. The Transaction Managers for Parties A and B then send the decrypted payload to the EVM for contract code execution.  This execution will update the state in the Quorum Node's Private StateDB only. NOTE: once the code has been executed it is discarded so is never available for reading without going through the above process.


