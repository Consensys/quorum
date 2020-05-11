## extend
quorum.js offers a way to add Quorum specific APIs to an instance of `web3`. Current APIs that can be extended are [Raft](../../Consensus/raft/raft-rpc-api), [Istanbul](../../Consensus/ibft/istanbul-rpc-api/), [Privacy](../../Getting%20Started/api/#privacy-apis), and [Permissioning](../../Permissioning/Permissioning%20apis) APIs. 

### Example
```js
const Web3 = require("web3");
const quorumjs = require("quorum-js");

const web3 = new Web3("http://localhost:22000");

quorumjs.extend(web3);

web3.quorum.eth.sendRawPrivateTransaction(signedTx, args);
```

### Parameters
| param | type | required | description |
| :---: | :---: | :---: | --- |
| `web3` | `Object` | yes | web3 instance |
| `apis` | `String` | no | comma-separated list of APIs to extend `web3` with.  Default is to add all APIs, i.e. `quorumjs.extend(web3, 'eth, raft, istanbul, quorumPermission')` | 

## RawTransactionManager
**TODO - add an overview**

### Example
```js
const Web3 = require("web3");
const quorumjs = require("quorum-js");

const web3 = new Web3("http://localhost:22000");

const tlsOptions = {
    key: fs.readFileSync("./cert.key"),
    clcert: fs.readFileSync("./cert.pem"),
    cacert: fs.readFileSync("./cacert.pem"),
    allowInsecure: false
};
const enclaveOptions = {
    privateUrl: "http://localhost:9081",
    tlsSettings: tlsOptions
};

const txnMngr = quorumjs.RawTransactionManager(web3, enclaveOptions);

txnMngr.sendRawTransaction(args);
``` 

### Parameters
| param | type | required | description |
| :---: | :---: | :---: | --- |
| `web3` | `Object` | yes | web3 instance |
| `enclaveOptions` | `Object` | yes | Privacy Manager connection configuration - see enclaveOptions below |

#### enclaveOptions
| param | type | required | description |
| :---: | :---: | :---: | --- |
| `privateUrl` | `String` | yes (unless `ipcPath` is provided) | Tessera `ThirdParty` server url (if using the Constellation Privacy Manager use `ipcPath` instead) |
| `ipcPath` | `String` | no | path to Privacy Manager `.ipc` socket file, `privateUrl` is preferred |
| `tlsSettings` | `Object` | no | TLS configuration for HTTPS Privacy Manager connections - see tlsSettings below |

#### tlsSettings
| param | type | required | description |
| :---: | :---: | :---: | --- |
| `key` | `String` | no  | client private key as byte string |
| `clcert` | `String` | no | client certificate (signed/unsigned) as byte string |
| `cacert` | `String` | no | CA certificate as byte string |
| `allowInsecure` | `boolean` | no | do not verify the Privacy Manager's certificate (can be used to allow self-signed certificates) |

### Methods

#### sendRawTransaction
```js
txnMngr.sendRawTransaction(txnParams);
```
Calls Tessera's `ThirdParty` `/storeraw`, replaces `data` field in `txnParams` with response (i.e. encrypted-payload hash), signs the transaction with the `from` account defined in `txnParams`, marks the transaction as private, RLP encodes the transaction in hex format, submits the signed transaction to the blockchain with `eth_sendRawPrivateTransaction`.

##### Parameters
1. `txnParams` - The transaction to sign and send 
    - `gasPrice`: `Number` - Must always be 0 in Quorum networks
    - `gasLimit`: `Number` - The amount of gas to use for the transaction
    - `to`: `String` - (optional) The destination address of the message, left undefined for a contract-creation transaction 
    - `value`: `Number` - (optional) The value transferred for the transaction, also the 
    endowment if it's a contract-creation transaction
    - `data`: `String` - (optional) Either a [byte string](https://github.com/ethereum/wiki/wiki/Solidity,-Docs-and-ABI) containing the associated data of the message, or in the case of a contract-creation transaction, the initialisation code (bytecode)
    - `decryptedAccount` : `String` - The public key of the sender's account
    - `nonce`: `Number`  - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce
    - `privateFrom`: `String`  - When sending a private transaction, the sending party's base64-encoded public key to use. If not present *and* passing `privateFor`, the default key as configured in the `TransactionManager` is used
    - `privateFor`: `List<String>`  - When sending a private transaction, an array of the recipients' base64-encoded public keys
    - `isPrivate`: `boolean` - Is the transaction private 
1. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous.

##### Returns
A promise that resolves to the transaction receipt if the transaction was sent successfully, else rejects with an error.

#### sendRawTransactionViaSendAPI

```js
txnMngr.sendRawTransactionViaSendAPI(txnParams);
```

!!! info
    `sendRawTransaction` should be used where possible.  `sendRawTransactionViaSendAPI` is necessary when using Constellation but requires providing the `ipcPath` in `enclaveOptions`.

Calls `Q2T` `/send` to encrypt txn data and send to all participant Privacy Manager nodes, replaces `data` field in `txnParams` with response (i.e. encrypted-payload hash), signs the transaction with the `from` account defined in `txnParams`, marks the transaction as private, submits the signed transaction to the blockchain with `eth_sendRawTransaction`.

##### Parameters
1. `txnParams` - The transaction to sign and send 
    - `gasPrice`: `Number` - Must always be 0 in Quorum networks
    - `gasLimit`: `Number` - The amount of gas to use for the transaction
    - `to`: `String` - (optional) The destination address of the message, left undefined for a contract-creation transaction 
    - `value`: `Number` - (optional) The value transferred for the transaction, also the 
    endowment if it's a contract-creation transaction
    - `data`: `String` - (optional) Either a [byte string](https://github.com/ethereum/wiki/wiki/Solidity,-Docs-and-ABI) containing the associated data of the message, or in the case of a contract-creation transaction, the initialisation code (bytecode)
    - `decryptedAccount` : `String` - The public key of the sender's account
    - `nonce`: `Number`  - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce
    - `privateFrom`: `String`  - When sending a private transaction, the sending party's base64-encoded public key to use. If not present *and* passing `privateFor`, the default key as configured in the `TransactionManager` is used
    - `privateFor`: `List<String>`  - When sending a private transaction, an array of the recipients' base64-encoded public keys
    - `isPrivate`: `boolean` - Is the transaction private 
1. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous.

##### Returns
A promise that resolves to the transaction receipt if the transaction was sent successfully, else rejects with an error.

#### setPrivate
```js
txnMngr.setPrivate(rawTransaction);
```
Marks a signed transaction as private by changing the value of `v` to `37` or `38`.
##### Parameters
1. `rawTransaction`: `String` - RLP-encoded signed transaction
##### Returns 
Updated RLP-encoded signed transaction

#### sendRawRequest
```js
txnMngr.sendRawRequest(rawTransaction, privateFor);
```
Call `eth_sendRawPrivateTransaction`, sending the signed transaction to the recipients specified in `privateFor`.
##### Parameters
1. `rawTransaction`: `String` - RLP-encoded signed transaction
1. `privateFor`: `List<String>` - List of the recipients' base64-encoded public keys

##### Returns
A promise that resolves to the transaction receipt if the transaction was sent successfully, else rejects with an error.

#### storeRawRequest
```js
txnMngr.storeRawRequest(data, privateFrom);
```
Calls Tessera's `ThirdParty` `/storeraw`.
##### Parameters
1. `data`: `String` - Hex encoded private transaction data (i.e. value of `data` field in the transaction)
1. `privateFrom`: `String` - When sending a private transaction, the sending party's base64-encoded public key to use. If not present *and* passing `privateFor`, the default key as configured in the `TransactionManager` is used

##### Returns
A promise that resolves to the hex-encoded hash of the encrypted `data`.  

## Start sending requests

## Starting Web3 on HTTP

To send asynchronous requests we need to instantiate `web3` with a `HTTP` address that points to the `Quorum` node.

```js
      const Web3 = require("web3");
      const web3 = new Web3(
        new Web3.providers.HttpProvider("http://localhost:22001")
      );
      const account = web3.eth.accounts[0];
```

### Send raw transactions using external signer. [Only available in Tessera with Quorum v2.2.0+]

If you want to use a different transaction signing mechanism, here are the steps to invoke the relevant APIs separately.

Firstly, a `storeRawRequest` function would need to be called by the enclave:

```js

const web3 = new Web3(new Web3.providers.HttpProvider(address));
const quorumjs = require("quorum-js");

const txnManager = quorumjs.RawTransactionManager(web3, {
  publicUrl: "http://localhost:8080",
  privateUrl: "http://localhost:8090"
});

txnManager.storeRawRequest(data, from)

```

##### Parameters

  - `data`: `String` - Either a [byte string](https://github.com/ethereum/wiki/wiki/Solidity,-Docs-and-ABI) 
    containing the associated data of the message, or in the case of a contract-creation transaction, the initialisation code (bytecode).
  - `from`: `String` (Optional) - Sender public key

A raw transaction will then need to be formed and signed, please note the data field will need to be replaced with the transaction hash which was returned from the privacy manager (the `key` field of the response data from `storeRawRequest` api call).


Secondly, the raw transaction can then be sent to Quorum by `sendRawRequest` function:

```js

var privateFor = ["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="]

txnManager.sendRawRequest(serializedTransaction, privateFor)

```

##### Parameters

  - `serializedTransaction`: `String` - Signed transaction data in HEX format.
  - `privateFor`: `List<String>` - When sending a private transaction, an array of the recipients' base64-encoded public keys.

## Examples for using quorum.js with [quorum-examples/7nodes](https://github.com/jpmorganchase/quorum-examples/tree/master/examples/7nodes)

Please see using Constellation and Quorum implementation private txn [example](https://github.com/jpmorganchase/quorum.js/blob/master/7nodes-test/deployContractViaIpc.js) and Tessera implementation [example](https://github.com/jpmorganchase/quorum.js/blob/master/7nodes-test/deployContractViaHttp.js). An extension sample is also provided.
