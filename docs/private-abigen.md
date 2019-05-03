# Abigen with Quorum

Abigen is a source code generator that converts quorum abi definitions into type-safe Go packages.  In addition to the original capabilities provided by Ethereum described [here](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts) Quorum abigen also supports deploying private transactions.

PrivateFrom and PrivateFor fields have been added to the *bind.TransactOpts type which allows users to specify the public keys of the transaction manager (Tessera/Constellation) used to send and receive transactions.

When using the PrivateFrom and PrivateFor fields, the "PRIVATE_CONFIG" environment variable must be set to point to the running constellation node's .ipc file and this node much match the public key set in the PrivateFrom field.  If not, deploying the private contract will fail.