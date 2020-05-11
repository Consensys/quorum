quorum.js offers a way to add Quorum specific APIs to an instance of `web3`. Current APIs that can be extended are [Raft](../../Consensus/raft/raft-rpc-api), [Istanbul](../../Consensus/ibft/istanbul-rpc-api/), [Privacy](../../Getting%20Started/api/#privacy-apis), and [Permissioning](../../Permissioning/Permissioning%20apis) APIs. 

## Example
```js
const Web3 = require("web3");
const quorumjs = require("quorum-js");

const web3 = new Web3("http://localhost:22000");

quorumjs.extend(web3);

web3.quorum.eth.sendRawPrivateTransaction(signedTx, args);
```
## Parameters
| param | type | required | description |
| :---: | :---: | :---: | --- |
| `web3` | `Object` | yes | web3 instance |
| `apis` | `String` | no | comma-separated list of APIs to extend `web3` with.  Default is to add all APIs, i.e. `quorumjs.extend(web3, 'eth, raft, istanbul, quorumPermission')` | 
