
# API

## Privacy APIs

__To support private transactions in Quorum, the `web3.eth.sendTransaction(object)` API method has been modified.__

```js
web3.eth.sendTransaction(transactionObject [, callback])
```

Sends a transaction to the network.

##### Parameters

1. `Object` - The transaction object to send:
    - `from`: `String` - The address for the sending account. Uses the `web3.eth.defaultAccount` property, if not specified.
    - `to`: `String` - (optional) The destination address of the message, left undefined for a contract-creation transaction.
    - `value`: `Number|String|BigNumber` - (optional) The value transferred for the transaction in Wei, also the endowment if it's a contract-creation transaction.
    - `gas`: `Number|String|BigNumber` - (optional, default: To-Be-Determined) The amount of gas to use for the transaction (unused gas is refunded).
    - <strike>`gasPrice`: `Number|String|BigNumber` - (optional, default: To-Be-Determined) The price of gas for this transaction in wei, defaults to the mean network gas price.</strike>
    - `data`: `String` - (optional) Either a [byte string](https://github.com/ethereum/wiki/wiki/Solidity,-Docs-and-ABI) containing the associated data of the message, or in the case of a contract-creation transaction, the initialisation code.
    - `nonce`: `Number`  - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.
    - `privateFrom`: `String`  - (optional) When sending a private transaction, the sending party's base64-encoded public key to use. If not present *and* passing `privateFor`, use the default key as configured in the `TransactionManager`.
    - `privateFor`: `List<String>`  - (optional) When sending a private transaction, an array of the recipients' base64-encoded public keys.
2. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous.

##### Returns

`String` - The 32 Bytes transaction hash as HEX string.

If the transaction was a contract creation use `web3.eth.getTransactionReceipt()` to get the contract address, after the transaction was mined.

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
***

## JSON RPC Privacy API Reference

#### eth_storageRoot

Returns the storage root of given address (Contract/Account etc)

##### Parameters

address, block number (hex)

##### Returns

`String` - 32 Bytes storageroot hash as HEX string at latest block height. It provides history of storage root hash when block number is provided.

##### Example

```js
//Request

curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc": "2.0", "method": "eth_storageRoot", "params":["0x1349f3e1b8d71effb47b840594ff27da7e603d17"], "id": 67}'

//Response

{

  "id":67,

  "jsonrpc": "2.0",

  "result": "0x81d1fa699f807735499cf6f7df860797cf66f6a66b565cfcda3fae3521eb6861"

}

```

```js
//When block number is provided...
//Request

curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc": "2.0", "method": "eth_storageRoot", "params":["0x1349f3e1b8d71effb47b840594ff27da7e603d17","0x1"], "id": 67}'

//Response

{

  "id":67,

  "jsonrpc": "2.0",

  "result": "0x81d1fa699f807735499cf6f7df860797cf66f6a66b565cfcda3fae3521eb6861"

}
//After private state of the contract is changed from '42' to '99'
//Request

curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc": "2.0", "method": "eth_storageRoot", "params":["0x1349f3e1b8d71effb47b840594ff27da7e603d17","0x2"], "id": 67}'

//Response

{

  "id":67,

  "jsonrpc": "2.0",

  "result": "0x0edb0e520c35df37a0d080d5245c9b8f9e1f9d7efab77c067d1e12c0a71299da"

}

```

***

#### eth_getQuorumPayload

Returns the unencrypted payload from Tessera/constellation

##### Parameters

Transaction payload hash in Hex format

##### Returns

`String` - unencrypted transaction payload in HEX format.

##### Example

```js
//Request

curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0", "method":"eth_getQuorumPayload", "params":["0x5e902fa2af51b186468df6ffc21fd2c26235f4959bf900fc48c17dc1774d86d046c0e466230225845ddf2cf98f23ede5221c935aac27476e77b16604024bade0"], "id":67}'

//Response

{

  "id":67,

  "jsonrpc": "2.0",

  "result": "0x6060604052341561000f57600080fd5b604051602080610149833981016040528080519060200190919050505b806000819055505b505b610104806100456000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a1afcd914605157806360fe47b11460775780636d4ce63c146097575b600080fd5b3415605b57600080fd5b606160bd565b6040518082815260200191505060405180910390f35b3415608157600080fd5b6095600480803590602001909190505060c3565b005b341560a157600080fd5b60a760ce565b6040518082815260200191505060405180910390f35b60005481565b806000819055505b50565b6000805490505b905600a165627a7a72305820d5851baab720bba574474de3d09dbeaabc674a15f4dd93b974908476542c23f00029000000000000000000000000000000000000000000000000000000000000002a"

}
// On a node which is not party to the transaction
//Response

{

  "id":67,

  "jsonrpc": "2.0",

  "result": "0x"

}
```


