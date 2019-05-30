# Transaction Processing

One of the key features of Quorum is that of Transaction Privacy.  To that end, we introduce the notion of 'Public Transactions' and 'Private Transactions'.  Note that this is a notional concept only and Quorum does not introduce new Transaction Types, but rather, the Ethereum Transaction Model has been extended to include an optional `privateFor` parameter (the population of which results in a Transaction being treated as private by Quorum) and the Transaction Type has a new `IsPrivate` method to identify such Transactions.

[Constellation](../../Privacy/Constellation/Constellation) / [Tessera](../../Privacy/Tessera/Tessera) are used by Quorum to transfer private payloads to their intended recipients, performing encryption and related operations in the process.

## Public Transactions
So called 'Public Transactions' are those Transactions whose payload is visible to all participants of the same Quorum network. These are [created as standard Ethereum Transactions in the usual way](https://github.com/ethereum/wiki/wiki/JavaScript-API#web3ethsendtransaction).

Examples of Public Transactions may include Market Data updates from some service provider, or some reference data update such as a correction to a Bond Security definition.

!!! Note
    'Public' Transactions are not Transactions from the public Ethereum network.  Perhaps a more appropriate term would be 'common' or 'global' Transactions, but 'Public' is used to contrast with 'Private' Transactions.

## Private Transactions
So called 'Private Transactions' are those Transactions whose payload is only visible to the network participants whose public keys are specified in the `privateFor` parameter of the Transaction .  `privateFor` can take multiple addresses in a comma separated list. (See Creating Private Transactions under the [Running Quorum](../../Getting Started/running) section).  

When the Quorum Node encounters a Transaction with a non-null `privateFor` value, it sets the `V` value of the Transaction Signature to be either `37` or `38` (as opposed to `27` or `28` which are the values used to indicate a Transaction is 'public' as per standard Ethereum as specified in the Ethereum yellow paper).

## Processing Transactions

### Public vs Private Transaction handling
Public Transactions are executed in the standard Ethereum way, and so if a Public Transaction is sent to an Account that holds Contract code, each participant will execute the same code and their underlying StateDBs will be updated accordingly.

Private Transactions, however, are not executed per standard Ethereum: prior to the sender's Quorum Node propagating the Transaction to the rest of the network, it replaces the original Transaction Payload with a hash of the encrypted Payload that it receives from Constellation/Tessera. Participants that are party to the Transaction will be able to replace the hash with the actual payload via their Constellation/Tessera instance, whilst those Participants that are not party will only see the hash. 

The result is that if a Private Transaction is sent to an Account that holds Contract code, those participants who are not party to the Transaction will simply end up skipping the Transaction, and therefore not execute the Contract code.  However those participants that are party to the Transaction will replace the hash with the original Payload before calling the EVM for execution, and their StateDB will be updated accordingly.  In absence of making corresponding changes to the geth client, these two sets of participants would therefore end up with different StateDBs and not be able to reach consensus. So in order to support this bifurcation of contract state, Quorum stores the state of Public contracts in a Public State Trie that is globally synchronised, and it stores the state of Private contracts in a Private State Trie that is not synchronised globally.  For details on how Consensus is achieved in light of this, please refer to [Quorum Consensus](../../Consensus/Consensus).

### Private Transaction Process Flow

Please refer [Private Transaction Flow](../../Privacy/Tessera/How%20Tessera%20Works) section under Tessera
