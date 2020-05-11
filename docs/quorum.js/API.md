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

const enclaveOptions = {
    privateUrl: "http://localhost:9081"
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
| `key` | `String` | no  | path to client key file |
| `clcert` | `String` | no | path to client cert file |
| `cacert` | `String` | no | path to ca cert file |
| `allowInsecure` | `boolean` | no | do not verify the Privacy Manager's certificate |

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

## Enclaves

The library supports connection to Quorum private transaction manager and execution of a raw transaction. Example **pseudo** code:

```js

const web3 = new Web3(new Web3.providers.HttpProvider(address));
const quorumjs = require("quorum-js");

const enclaveOptions = {
  /* at least one enclave option must be provided     */
  /* ipcPath is preferred for utilizing older API     */
  /* Constellation only supports ipcPath              */
  /* For Tessera: privateUrl is ThirdParty server url */
  ipcPath: "/quorum-examples/examples/7nodes/qdata/c1/tm.ipc",
  publicUrl: "http://localhost:8080",
  privateUrl: "http://localhost:8090"
};

const rawTransactionManager = quorumjs.RawTransactionManager(web3, enclaveOptions);

const txnParams = {
  gasPrice: 0,
  gasLimit: 4300000,
  to: null,
  value: 0,
  data: deploy,
  from: decryptedAccount,
  isPrivate: true,
  privateFrom: TM1_PUBLIC_KEY,
  privateFor: [TM2_PUBLIC_KEY],
  nonce
};

// Older API: txn manager and Quorum version agnostic
// requires the IPC path to be set in enclaveOptions
rawTransactionManager.sendRawTransactionViaSendAPI(txnParams);

// Newer API: Quorum v2.2.1+ and Tessera
// requires the private URL to be set in enclaveOptions
rawTransactionManager.sendRawTransaction(txnParams);
```

It sends a private transaction to the network [ this transaction can be either a contract deployment or a contract call ].


##### Parameters

1. `Object` - The transaction object to send:
    - <strike>`gasPrice`: `Number` - The price of gas for this transaction, defaults to the mean 
    network gas price [ because we work in a private network the gasPrice is 0 ].</strike>
    - `gasLimit`: `Number` - The amount of gas to use for the transaction.
    - `to`: `String` - (optional) The destination address of the message, left undefined for a contract-creation 
    transaction [in case of a contract creation the to field must be `null`].
    - `value`: `Number` - (optional) The value transferred for the transaction, also the 
    endowment if it's a contract-creation transaction.
    - `data`: `String` - (optional) Either a [byte string](https://github.com/ethereum/wiki/wiki/Solidity,-Docs-and-ABI) 
    containing the associated data of the message, or in the case of a contract-creation transaction, the initialisation code (bytecode).
    - `decryptedAccount` : `String` - the public key of the sender's account;
    - `nonce`: `Number`  - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.
    - `privateFrom`: `String`  - When sending a private transaction, the sending party's base64-encoded public key to use. If not present *and* passing `privateFor`, use the default key as configured in the `TransactionManager`.
    - `privateFor`: `List<String>`  - When sending a private transaction, an array of the recipients' base64-encoded public keys.
2. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous.

##### Returns

`String` - The 32 Bytes transaction hash as HEX string.


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

## Example for connecting to Tessera node with TLS enabled

To send a request to Tessera node with TLS enabled for raw transactions, add the cert information as specified below. 

```js

let tlsOptions = {
        key: fs.readFileSync('./cert.key'),
        clcert: fs.readFileSync('./cert.pem'),
        ca: fs.readFileSync('./cert.pem')
        allowInsecure: false
    }

const rawTransactionManager = quorumjs.RawTransactionManager(web31, {
        privateUrl:toPrivateURL,
        tlsSettings: tlsOptions
    });

```

##### Parameters

  - `key` : `String` - a byte string with private key of the client
  - `clcert` : `String` - a byte string with client certificate (signed / unsigned)
  - `ca` : `String` - (Optional) a byte string with CA certificate
  - `allowInsecure` : `Boolean` - to accept self signed certificates
