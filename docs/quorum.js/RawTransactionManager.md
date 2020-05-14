The `RawTransactionManager` module of quorum.js provides access to private transaction APIs that require a connection to a [Privacy Manager](../../Privacy/Privacy-Manager).

## Example
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

## Parameters
| param | type | required | description |
| :---: | :---: | :---: | --- |
| <span style="white-space:nowrap">`web3`</span> | `Object` | yes | web3 instance |
| <span style="white-space:nowrap">`enclaveOptions`</span> | `Object` | yes | Privacy Manager connection configuration - see [enclaveOptions](#enclaveoptions) |

### enclaveOptions
| param | type | required | description |
| :---: | :---: | :---: | --- |
| <span style="white-space:nowrap">`privateUrl`</span> | `String` | yes (unless `ipcPath` is provided) | Tessera `ThirdParty` server url (if using the Constellation Privacy Manager use `ipcPath` instead) |
| <span style="white-space:nowrap">`ipcPath`</span> | `String` | no | path to Privacy Manager `.ipc` socket file, `privateUrl` is preferred |
| <span style="white-space:nowrap">`tlsSettings`</span> | `Object` | no | TLS configuration for HTTPS Privacy Manager connections - see [tlsSettings](#tlssettings) |

### tlsSettings
| param | type | required | description |
| :---: | :---: | :---: | --- |
| <span style="white-space:nowrap">`key`</span> | `String` | no  | client private key as byte string |
| <span style="white-space:nowrap">`clcert`</span> | `String` | no | client certificate (signed/unsigned) as byte string |
| <span style="white-space:nowrap">`cacert`</span> | `String` | no | CA certificate as byte string |
| <span style="white-space:nowrap">`allowInsecure`</span> | `boolean` | no | do not verify the Privacy Manager's certificate (can be used to allow self-signed certificates) |

## Methods

### sendRawTransaction

!!! info "If using Constellation"
    Constellation privacy managers do not support this method.  Use [`sendRawTransactionViaSendAPI`](#sendrawtransactionviasendapi) instead.
    
```js
txnMngr.sendRawTransaction(txnParams);
```
Calls Tessera's `ThirdParty` `/storeraw` API, replaces the `data` field in `txnParams` with the response (i.e. encrypted-payload hash), signs the transaction with the `from` account defined in `txnParams`, marks the transaction as private, RLP encodes the transaction in hex format, and submits the signed transaction to the blockchain with `eth_sendRawPrivateTransaction`.

#### Parameters
1. `txnParams` - The transaction to sign and send 
    - `gasPrice`: `Number` - Must always be 0 in Quorum networks
    - `gasLimit`: `Number` - The amount of gas to use for the transaction
    - `to`: `String` - (optional) The destination address of the message, left undefined for a contract-creation transaction 
    - `value`: `Number` - (optional) The value transferred for the transaction
    - `data`: `String` - (optional) Either a byte string containing the associated data of the message, or the initialisation code (bytecode) in the case of a contract-creation transaction
    - `from` - `Object`: [Decrypted account object](https://web3js.readthedocs.io/en/v1.2.7/web3-eth-accounts.html#decrypt) 
    - `nonce`: `Number`  - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce
    - `privateFrom`: `String`  - When sending a private transaction, the sending party's base64-encoded public key to use. If not present *and* passing `privateFor`, the default key as configured in the `TransactionManager` is used
    - `privateFor`: `List<String>`  - When sending a private transaction, an array of the recipients' base64-encoded public keys
    - `isPrivate`: `boolean` - Is the transaction private 
1. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous.

#### Returns
A promise that resolves to the transaction receipt if the transaction was sent successfully, else rejects with an error.

### sendRawTransactionViaSendAPI

!!! info "If using Tessera"
    Tessera privacy managers support [`sendRawTransaction`](#sendrawtransaction) which should be used instead.  `sendRawTransactionViaSendAPI` requires exposing the `Q2T` server to the `js` app.  Ideally only the `ThirdParty` server should be exposed to such applications.

```js
txnMngr.sendRawTransactionViaSendAPI(txnParams);
```

Calls Privacy Manager's `/send` API to encrypt txn data and send to all participant Privacy Manager nodes, replaces `data` field in `txnParams` with response (i.e. encrypted-payload hash), signs the transaction with the `from` account defined in `txnParams`, marks the transaction as private, and submits the signed transaction to the blockchain with `eth_sendRawTransaction`.

#### Parameters
1. `txnParams` - The transaction to sign and send 
    - `gasPrice`: `Number` - Must always be 0 in Quorum networks
    - `gasLimit`: `Number` - The amount of gas to use for the transaction
    - `to`: `String` - (optional) The destination address of the message, left undefined for a contract-creation transaction 
    - `value`: `Number` - (optional) The value transferred for the transaction
    - `data`: `String` - (optional) Either a byte string containing the associated data of the message, or the initialisation code (bytecode) in the case of a contract-creation transaction
    - `from` - `Object`: [Decrypted account object](https://web3js.readthedocs.io/en/v1.2.7/web3-eth-accounts.html#decrypt) 
    - `nonce`: `Number`  - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce
    - `privateFrom`: `String`  - When sending a private transaction, the sending party's base64-encoded public key to use. If not present *and* passing `privateFor`, the default key as configured in the `TransactionManager` is used
    - `privateFor`: `List<String>`  - When sending a private transaction, an array of the recipients' base64-encoded public keys
    - `isPrivate`: `boolean` - Is the transaction private 
1. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous.

#### Returns
A promise that resolves to the transaction receipt if the transaction was sent successfully, else rejects with an error.

### setPrivate
```js
txnMngr.setPrivate(rawTransaction);
```
Marks a signed transaction as private by changing the value of `v` to `37` or `38`.
#### Parameters
1. `rawTransaction`: `String` - RLP-encoded hex-format signed transaction
#### Returns 
Updated RLP-encoded hex-format signed transaction

### storeRawRequest
```js
txnMngr.storeRawRequest(data, privateFrom);
```
Calls Tessera's `ThirdParty` `/storeraw` API to encrypt the provided `data` and store in preparation for a `eth_sendRawPrivateTransaction`.

#### Parameters
1. `data`: `String` - Hex encoded private transaction data (i.e. value of `data`/`input` field in the transaction)
1. `privateFrom`: `String` - Sending party's base64-encoded public key

#### Returns
A promise that resolves to the hex-encoded hash of the encrypted `data` (`key` field) that should be used to replace the `data` field of a transaction if externally signing.  

### sendRawRequest
```js
txnMngr.sendRawRequest(rawTransaction, privateFor);
```
Call `eth_sendRawPrivateTransaction`, sending the signed transaction to the recipients specified in `privateFor`.
#### Parameters
1. `rawTransaction`: `String` - RLP-encoded hex-format signed transaction
1. `privateFor`: `List<String>` - List of the recipients' base64-encoded public keys

#### Returns
A promise that resolves to the transaction receipt if the transaction was sent successfully, else rejects with an error.

## Examples

### Externally signing and sending a private tx

!!!info
    This is not supported by Constellation and requires Quorum v2.2.0+

[Code sample](https://github.com/jpmorganchase/quorum.js/blob/master/7nodes-test/deployContractViaHttp-externalSigningTemplate.js).

1. `storeRawRequest` to encrypt the transaction `data`
    ```js
    txnManager.storeRawRequest(data, from)
    ```
1. Replace `data` field of transaction with `key` field from `storeRawRequest` response
1. Sign the transaction
1. Mark the signed transaction as private with `setPrivate`
    ```js
    txnManager.setPrivate(signedTx)
    ```
1. Send the signed transaction to Quorum with `sendRawRequest`
    ```js
     txnManager.sendRawRequest(serializedTransaction, privateFor)
    ```
 
### Other examples
The [7nodes-test](https://github.com/jpmorganchase/quorum.js/tree/master/7nodes-test) directory in the quorum.js project repo contains examples of quorum.js usage.  These scripts can be tested with a running [7nodes test network](https://github.com/jpmorganchase/quorum-examples/tree/master/examples/7nodes).