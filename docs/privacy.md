
# Privacy

## Sending Private Transactions

To send a private transaction, a private Transaction Manager must be configured. This is the
service which transfers private payloads to their intended recipients, performing
encryption and related operations in the process.

[Constellation](https://github.com/jpmorganchase/constellation) / [Tessera](https://github.com/jpmorganchase/tessera) is used to provide the private Transaction Manager for a Quorum node.  Once a Constellation/Tessera node is running, the `PRIVATE_CONFIG` environment variable is used to point the Quorum node to the transaction manager instance.  Examples of this can be seen in the [quorum-examples 7nodes](https://github.com/jpmorganchase/quorum-examples) source files.

Once Constellation/Tessera is launched and `PRIVATE_CONFIG` points to a valid configuration file,
a `SendTransaction` call can be made private by specifying the `privateFor` argument.
`privateFor` is a list of public keys of the intended recipients (these public keys are distinct from Ethereum account keys). When a transaction
is private, the transaction contents will be sent to the `PrivateTransactionManager` and the
identifier returned will be placed in the transaction instead. When other Quorum nodes
receive a private transaction, they will query their `PrivateTransactionManager` for the
identifier and replace the transaction contents with the result.  Nodes which are
not party to a transaction will not be able to retrieve the original contents.