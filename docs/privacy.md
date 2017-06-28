
# Privacy

## Sending Private Transactions

To send a private transaction, a `PrivateTransactionManager` must be configured. This is the
service which transfers private payloads to their intended recipients, performing
encryption and related operations in the process.

Currently, `constellation` is supported out of the box via the `PRIVATE_CONFIG` environment
variable (please note that this integration method will change in the near future.) See the
`7nodes` folder in the `quorum-examples` repository for a complete example of how to use it.
The transaction sent in `script1.js` is private for node 7's `PrivateTransactionManager`
public key.

Once `constellation` is launched and `PRIVATE_CONFIG` points to a valid configuration file,
a `SendTransaction` call can be made private by specifying the `privateFor` argument.
`privateFor` is a list of public keys of the intended recipients. (Note that in the case of
`constellation`, this public key is distinct from Ethereum account keys.) When a transaction
is private, the transaction contents will be sent to the `PrivateTransactionManager` and the
identifier returned will be placed in the transaction instead. When other Quorum nodes
receive a private transaction, they will query their `PrivateTransactionManager` for the
identifier and replace the transaction contents with the result (if any; nodes which are
not party to a transaction will not be able to retrieve the original contents.)
