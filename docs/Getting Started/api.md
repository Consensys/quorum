
# API

## Privacy APIs

#### eth.sendTransaction

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

#### eth.sendRawPrivateTransaction

__To support sending raw transactions in Quorum, the `web3.eth.sendRawPrivateTransaction(string, object)` API method has been created.__

```js
web3.eth.sendRawPrivateTransaction(signedTransactionData [, privateData] [, callback])
```

Sends a pre-signed transaction. For example can be signed using: https://github.com/SilentCicero/ethereumjs-accounts

__Important:__ Please note that before calling this API, a `storeraw` api need to be called first to Quorum's private transaction manager. Instructions on how to do this can be found [here](https://github.com/jpmorganchase/tessera/wiki/Interface-&-API).

##### Parameters
 1. `String` - Signed transaction data in HEX format
 2. `Object` - Private data to send
    - `privateFor`: `List<String>`  - When sending a private transaction, an array of the recipients' base64-encoded public keys.
3. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous. See [this note](#using-callbacks) for details.

 ##### Returns
 `String` - The 32 Bytes transaction hash as HEX string.
 If the transaction was a contract creation use [web3.eth.getTransactionReceipt()](#web3ethgettransactionreceipt) to get the contract address, after the transaction was mined.
 
 
 ##### Example
  ```js
 var Tx = require('ethereumjs-tx');
 var privateKey = new Buffer('e331b6d69882b4cb4ea581d88e0b604039a3de5967688d3dcffdd2270c0fd109', 'hex')
  var rawTx = {
   nonce: '0x00',
   gasPrice: '0x09184e72a000', 
   gasLimit: '0x2710',
   to: '0x0000000000000000000000000000000000000000', 
   value: '0x00', 
   // This data should be the hex value of the hash returned by Quorum's privacy transaction manager after invoking storeraw api
   data: '0x7f7465737432000000000000000000000000000000000000000000000000000000600057'
 }
  var tx = new Tx(rawTx);
 tx.sign(privateKey);
  var serializedTx = tx.serialize();
  //console.log(serializedTx.toString('hex'));
 //f889808609184e72a00082271094000000000000000000000000000000000000000080a47f74657374320000000000000000000000000000000000000000000000000000006000571ca08a8bbf888cfa37bbf0bb965423625641fc956967b81d12e23709cead01446075a01ce999b56a8a88504be365442ea61239198e23d1fce7d00fcfc5cd3b44b7215f
  web3.eth.sendRawPrivateTransaction('0x' + serializedTx.toString('hex'), {privateFor: ["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="]}, function(err, hash) {
   if (!err)
     console.log(hash); // "0x7f9fade1c0d57a7af66ab4ead79fade1c0d57a7af66ab4ead7c2c2eb7b11a91385"
 });
 ```
 
 

## JSON RPC Privacy API Reference

__In addition to the JSON-RPC provided by Ethereum, Quorum exposes below two API calls.__


#### eth_storageRoot

Returns the storage root of given address (Contract/Account etc)

##### Parameters

1. `address`: `String` - The address to fetch the storage root for in hex
2. `block`: `String` - (optional) The block number to look at in hex (e.g. `0x15` for block 21). Uses the latest block if not specified. 

##### Returns

`String` - 32 Bytes storageroot hash as HEX string at latest block height. When blocknumber is given, it provides the storageroot hash at that block height. 

##### Example

```js
// Request

curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc": "2.0", "method": "eth_storageRoot", "params":["0x1349f3e1b8d71effb47b840594ff27da7e603d17"], "id": 67}'

// Response
{
  "id":67,
  "jsonrpc": "2.0",
  "result": "0x81d1fa699f807735499cf6f7df860797cf66f6a66b565cfcda3fae3521eb6861"
}

```

```js
// When block number is provided...
// Request

curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc": "2.0", "method": "eth_storageRoot", "params":["0x1349f3e1b8d71effb47b840594ff27da7e603d17","0x1"], "id": 67}'

// Response
{
  "id":67,
  "jsonrpc": "2.0",
  "result": "0x81d1fa699f807735499cf6f7df860797cf66f6a66b565cfcda3fae3521eb6861"
}

// After private state of the contract is changed from '42' to '99'
// Request

curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc": "2.0", "method": "eth_storageRoot", "params":["0x1349f3e1b8d71effb47b840594ff27da7e603d17", "0x2"], "id": 67}'

// Response
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

1. `id`: `String` - the HEX formatted generated Sha3-512 hash of the encrypted payload from the Private Transaction Manager. This is seen in the transaction as the `input` field

##### Returns

`String` - unencrypted transaction payload in HEX format.

##### Example

```js
// Request

curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0", "method":"eth_getQuorumPayload", "params":["0x5e902fa2af51b186468df6ffc21fd2c26235f4959bf900fc48c17dc1774d86d046c0e466230225845ddf2cf98f23ede5221c935aac27476e77b16604024bade0"], "id":67}'

// Response
{
  "id":67,
  "jsonrpc": "2.0",
  "result": "0x6060604052341561000f57600080fd5b604051602080610149833981016040528080519060200190919050505b806000819055505b505b610104806100456000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a1afcd914605157806360fe47b11460775780636d4ce63c146097575b600080fd5b3415605b57600080fd5b606160bd565b6040518082815260200191505060405180910390f35b3415608157600080fd5b6095600480803590602001909190505060c3565b005b341560a157600080fd5b60a760ce565b6040518082815260200191505060405180910390f35b60005481565b806000819055505b50565b6000805490505b905600a165627a7a72305820d5851baab720bba574474de3d09dbeaabc674a15f4dd93b974908476542c23f00029000000000000000000000000000000000000000000000000000000000000002a"
}

// On a node which is not party to the transaction
// Response
{
  "id":67,
  "jsonrpc": "2.0",
  "result": "0x"
}
```

***

#### eth_sendTransactionAsync
 
 Sends a transaction to the network asynchronously. This will return 
 immediately, potentially before the transaction has been submitted to the 
 transaction pool. A callback can be provided to receive the result of 
 submitting the transaction; a server must be set up to receive POST requests
 at the given URL.
 
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
     - `callbackUrl`: `String` - (optional) the URL to perform a POST request to to post the result of submitted the transaction
 
##### Returns
 
 1. `String` - The empty hash, defined as `0x0000000000000000000000000000000000000000000000000000000000000000`
 
 The callback URL receives the following object:
 
 2. `Object` - The result object:
    - `id`: `String` - the identifier in the original RPC call, used to match this result to the request
    - `txHash`: `String` - the transaction hash that was generated, if successful
    - `error`: `String` - the error that occurred whilst submitting the transaction.
 
 If the transaction was a contract creation use `web3.eth.getTransactionReceipt()` to get the contract address, after the transaction was mined.
 
##### Example

For the RPC call and the immediate response:

```
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0", "method":"eth_sendTransactionAsync", "params":[{"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d", "data": "0x6060604052341561000f57600080fd5b604051602080610149833981016040528080519060200190919050505b806000819055505b505b610104806100456000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a1afcd914605157806360fe47b11460775780636d4ce63c146097575b600080fd5b3415605b57600080fd5b606160bd565b6040518082815260200191505060405180910390f35b3415608157600080fd5b6095600480803590602001909190505060c3565b005b341560a157600080fd5b60a760ce565b6040518082815260200191505060405180910390f35b60005481565b806000819055505b50565b6000805490505b905600a165627a7a72305820d5851baab720bba574474de3d09dbeaabc674a15f4dd93b974908476542c23f00029000000000000000000000000000000000000000000000000000000000000002a", "gas": "0x47b760", "privateFor": ["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="]}], "id":67}'

// Response
{
  "id": 67,
  "jsonrpc": "2.0",
  "result": "0x0000000000000000000000000000000000000000000000000000000000000000"
}


// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0", "method":"eth_sendTransactionAsync", "params":[{"from":"0xe2e382b3b8871e65f419d", "data": "0x6060604052341561000f57600080fd5b604051602080610149833981016040528080519060200190919050505b806000819055505b505b610104806100456000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a1afcd914605157806360fe47b11460775780636d4ce63c146097575b600080fd5b3415605b57600080fd5b606160bd565b6040518082815260200191505060405180910390f35b3415608157600080fd5b6095600480803590602001909190505060c3565b005b341560a157600080fd5b60a760ce565b6040518082815260200191505060405180910390f35b60005481565b806000819055505b50565b6000805490505b905600a165627a7a72305820d5851baab720bba574474de3d09dbeaabc674a15f4dd93b974908476542c23f00029000000000000000000000000000000000000000000000000000000000000002a", "gas": "0x47b760", "privateFor": ["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="]}], "id":67}'

//If a syntactic error occured with the RPC call. 
//In this example the wallet address is the wrong length
//so the error is it cannot convert the parameter to the correct type
//it is NOT an error relating the the address not being managed by this node.

//Response
{
    "id": 67,
    "jsonrpc": "2.0",
    "error": {
        "code": -32602,
        "message": "invalid argument 0: json: cannot unmarshal hex string of odd length into Go struct field AsyncSendTxArgs.from of type common.Address"
    }
}
```

If the callback URL is provided, the following response will be received after
the transaction has been submitted; this example assumes a webserver that can
be accessed by calling http://localhost:8080 has been set up to accept POST 
requests:

```


// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0", "method":"eth_sendTransactionAsync", "params":[{"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d", "data": "0x6060604052341561000f57600080fd5b604051602080610149833981016040528080519060200190919050505b806000819055505b505b610104806100456000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a1afcd914605157806360fe47b11460775780636d4ce63c146097575b600080fd5b3415605b57600080fd5b606160bd565b6040518082815260200191505060405180910390f35b3415608157600080fd5b6095600480803590602001909190505060c3565b005b341560a157600080fd5b60a760ce565b6040518082815260200191505060405180910390f35b60005481565b806000819055505b50565b6000805490505b905600a165627a7a72305820d5851baab720bba574474de3d09dbeaabc674a15f4dd93b974908476542c23f00029000000000000000000000000000000000000000000000000000000000000002a", "gas": "0x47b760", "privateFor": ["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="], "callbackUrl": "http://localhost:8080"}], "id":67}'

// Response
//Note that the ID is the same in the callback as the request - this can be used to match the request to the response.
{
  "id": 67,
  "txHash": "0x75ebbf4fbe29355fc8a4b8d1e14ecddf0228b64ef41e6d2fce56047650e2bf17"
}




// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0", "method":"eth_sendTransactionAsync", "params":[{"from":"0xae9bc6cd5145e67fbd1887a5145271fd182f0ee7", "callbackUrl": "http://localhost:8080", "data": "0x6060604052341561000f57600080fd5b604051602080610149833981016040528080519060200190919050505b806000819055505b505b610104806100456000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a1afcd914605157806360fe47b11460775780636d4ce63c146097575b600080fd5b3415605b57600080fd5b606160bd565b6040518082815260200191505060405180910390f35b3415608157600080fd5b6095600480803590602001909190505060c3565b005b341560a157600080fd5b60a760ce565b6040518082815260200191505060405180910390f35b60005481565b806000819055505b50565b6000805490505b905600a165627a7a72305820d5851baab720bba574474de3d09dbeaabc674a15f4dd93b974908476542c23f00029000000000000000000000000000000000000000000000000000000000000002a", "gas": "0x47b760", "privateFor": ["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="]}], "id":67}'

//If a semantic error occured with the RPC call. 
//In this example the wallet address is not managed by the node
//So the RPC call will succeed (giving the empty hash), but the callback will show a failure

// In the callback
{
    "id": 67,
    "error":"unknown account"
}
```
