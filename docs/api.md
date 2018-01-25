
# API

## Privacy APIs

### `web3.eth.sendTransaction(object)` was modified to support private transactions

```js
web3.eth.sendTransaction(transactionObject [, callback])
```

Sends a transaction to the network.

##### Parameters

1. `Object` - The transaction object to send:
  - `from`: `String` - The address for the sending account. Uses the [web3.eth.defaultAccount](#web3ethdefaultaccount) property, if not specified.
  - `to`: `String` - (optional) The destination address of the message, left undefined for a contract-creation transaction.
  - `value`: `Number|String|BigNumber` - (optional) The value transferred for the transaction in Wei, also the endowment if it's a contract-creation transaction.
  - `gas`: `Number|String|BigNumber` - (optional, default: To-Be-Determined) The amount of gas to use for the transaction (unused gas is refunded).
  - <strike>`gasPrice`: `Number|String|BigNumber` - (optional, default: To-Be-Determined) The price of gas for this transaction in wei, defaults to the mean network gas price.</strike>
  - `data`: `String` - (optional) Either a [byte string](https://github.com/ethereum/wiki/wiki/Solidity,-Docs-and-ABI) containing the associated data of the message, or in the case of a contract-creation transaction, the initialisation code.
  - `nonce`: `Number`  - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.
  - `privateFrom`: `String`  - (optional) When sending a private transaction, the sending party's base64-encoded public key to use. If not present *and* passing `privateFor`, use the default key as configured in the `TransactionManager`.
  - `privateFor`: `List<String>`  - (optional) When sending a private transaction, an array of the recipients' base64-encoded public keys.
2. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous. See [this note](#using-callbacks) for details.

##### Returns

`String` - The 32 Bytes transaction hash as HEX string.

If the transaction was a contract creation use [web3.eth.getTransactionReceipt()](#web3gettransactionreceipt) to get the contract address, after the transaction was mined.

##### Example

```js
// compiled solidity source code using https://chriseth.github.io/cpp-ethereum/
var code = "603d80600c6000396000f3007c01000000000000000000000000000000000000000000000000000000006000350463c6888fa18114602d57005b6007600435028060005260206000f3";

web3.eth.sendTransaction({
    data: code,
    privateFor: ["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="]
  },
  function(err, address) {
    if (!err) {
      console.log(address); // "0x7f9fade1c0d57a7af66ab4ead7c2eb7b11a91385"
    }
  }
});
```
