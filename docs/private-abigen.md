# Abigen with Quorum

Abigen is a source code generator that converts quorum abi definitions into type-safe Go packages.  In addition to the original capabilities provided by Ethereum described [here](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts) Quorum abigen also supports private transactions.

`PrivateFrom` and `PrivateFor` fields have been added to the `bind.TransactOpts` type which allows users to specify the public keys of the transaction manager (Tessera/Constellation) used to send and receive private transactions.