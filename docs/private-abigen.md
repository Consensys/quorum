# Abigen with Quorum

### Overview

Abigen is a source code generator that converts smart contract ABI definitions into type-safe Go packages.  In addition to the original capabilities provided by Ethereum described [here](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts). Quorum Abigen also supports private transactions.

### Implementation

`PrivateFrom` and `PrivateFor` fields have been added to the `bind.TransactOpts` which allows users to specify the public keys of the transaction manager (Tessera/Constellation) used to send and receive private transactions. The existing `ethclient` has been extended with a private transaction manager client to support sending `/storeraw` request.
