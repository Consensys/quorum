# Contract Extension APIs
## APIs
### `quorumExtension_extendContract` 
The api for initiating contract extension to a new node.  
#### Parameter
* `toExtend`: address of the private contract which is being extended to the new node
* `newRecipientPtmPublicKey`: Tessera public key of the recipient node
* `recipientAddress`: ethereum addresses which will accept the contract extension in the recipient node
* `txArgs`: standard transaction arguments with `privateFor` info of both nodes involved in contract extension

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"quorumExtension_extendContract","params":["0x9aff347f193ca4560276c3322193224dcdbbe578","BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","0xed9d02e382b34818e88b88a309c7fe71e65f419d",{"from":"0xca843569e3427144cead5e4d5999a3d0ccf92b8e","value":"0x0","privateFor":["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo="],"privacyFlag":1}],"id":15}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"0xceffe8051d098920ac84e33b8a05c48180ed9b26581a6a06ce9874a1bf1502bd"}
```

```javascript tab="geth console"
> quorumExtension.extendContract("0x9aff347f193ca4560276c3322193224dcdbbe578", "BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=", "0xed9d02e382b34818e88b88a309c7fe71e65f419d",{from: "0xca843569e3427144cead5e4d5999a3d0ccf92b8e", privateFor:["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo="]})

"0x9e0101dd215281b33989b3ae093147e9009353bb63f531490409a628c6e87310"
```

If the contract is already under the process of extension, api call to extend it again will fail. 

```javascript
> quorumExtension.extendContract("0x9aff347f193ca4560276c3322193224dcdbbe578", "BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=", "0xed9d02e382b34818e88b88a309c7fe71e65f419d",{from: "0xca843569e3427144cead5e4d5999a3d0ccf92b8e", privateFor:["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo="]})
  Error: contract extension in progress for the given contract address
      at web3.js:3143:20
      at web3.js:6347:15
      at web3.js:5081:36
      at <anonymous>:1:1
```

### `quorumExtension_acceptExtension` 
The api for accepting the contract extension on the recipient node. This can be invoked by the ethereum address which is was given as the `recipientAddress` when creating the contract extension.
#### Parameter
* `addressToVoteOn`: address of the contract extension management contract
* `vote`: bool - `true` indicates approval and `false` disapproval. Contract extension is completed only when all voters vote true. If any participant votes `false` the extension process will be cancelled.

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure
* `txArgs`: standard transaction arguments with `privateFor` info of all nodes involved in contract extension
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumExtension_acceptExtension", "params":["0xb1c57951a2f3006910115eadf0f167890e99b9cb", true, {"from": "0xed9d02e382b34818e88b88a309c7fe71e65f419d", "privateFor":["QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc="]}], "id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"0x8d34a594b286087f45029daad2d5a8fd42f70abb0ae2492429a256a2ba4cb0dd"}
```


```javascript tab="geth console"
> quorumExtension.acceptExtension("0x1349f3e1b8d71effb47b840594ff27da7e603d17", true ,{from: "0x0fbdc686b912d7722dc86510934589e0aaf3b55a", privateFor:["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo="]})

"0x9e0101dd215281b33989b3ae093147e9009353bb63f531490409a628c6e87310"
```


If the contract is already under the process of extension, api call to extend it again will fail.  
```javascript
> quorumExtension.extendContract("0x1932c48b2bf8102ba33b4a6b545c32236e342f34", "1iTZde/ndBHvzhcl7V68x44Vx7pl8nwx9LqnM/AfJUg=", "0x0fbdc686b912d7722dc86510934589e0aaf3b55a", {from: eth.accounts[0], privateFor: ["1iTZde/ndBHvzhcl7V68x44Vx7pl8nwx9LqnM/AfJUg="]})
Error: contract extension in progress for the given contract address
    at web3.js:3143:20
    at web3.js:6347:15
    at web3.js:5081:36
    at <anonymous>:1:1
```

The recipient can accept the extension only once. Executing `quorumExtension.acceptExtension` once extension process is completed will result in the below error
```javascript
> quorumExtension.acceptExtension("0x1349f3e1b8d71effb47b840594ff27da7e603d17", true ,{from: "0x0fbdc686b912d7722dc86510934589e0aaf3b55a", privateFor:["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo="]})
Error: contract extension process complete. nothing to accept
    at web3.js:3143:20
    at web3.js:6347:15
    at web3.js:5081:36
    at <anonymous>:1:1
``` 

Executing `quorumExtension.acceptExtension` from an account which is other than `recipientAddress` will result in the below error:
```javascript
> quorumExtension.acceptExtension("0x4d3bfd7821e237ffe84209d8e638f9f309865b87", true, {from: "0x0bb8aaa95b514d8bff1287c1fb58510479478b4a", privateFor:["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo="]})
Error: account is not acceptor of this extension request
    at web3.js:3143:20
    at web3.js:6347:15
    at web3.js:5081:36
    at <anonymous>:1:1
```

### `quorumExtension_cancelExtension` 
The api for cancelling the contract extension process. This can only be invoked by the ethereum address which initaited the contract extension process.
#### Parameter
* `extensionContract`: address of the contract extension management contract
* `vote`: bool - true indicates approval and false disapproval. Contract extension is completed only when all voters vote true

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure
* `txArgs`: standard transaction arguments with `privateFor` info of all nodes involved in contract extension
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"quorumExtension_cancelExtension","params":["0x622aff909c081783613c9d3f5f4c47be78b310ac",{"from":"0xca843569e3427144cead5e4d5999a3d0ccf92b8e","value":"0x0","privateFor":["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo="],"privacyFlag":1}],"id":63}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"0xb43da7dbeae5347df86c6933786b8c536b4622463b577a990d4c87214845d16a"}
```


```javascript tab="geth console"
> quorumExtension.cancelExtension("0x622aff909c081783613c9d3f5f4c47be78b310ac",{"from":"0xca843569e3427144cead5e4d5999a3d0ccf92b8e","value":"0x0","privateFor": ["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo="]})

"0x9e0101dd215281b33989b3ae093147e9009353bb63f531490409a628c6e87310"
```

If the api is invoked by an ethereum address which is not the creator of the contract extension, an error is thrown.  
```javascript
> quorumExtension.cancelExtension("0x4d3bfd7821e237ffe84209d8e638f9f309865b87", {from: "0xbdafac69ab6c5c2f1c2ba36a462c9d2fb01f877d", privateFor:["1iTZde/ndBHvzhcl7V68x44Vx7pl8nwx9LqnM/AfJUg"]})
Error: account is not the creator of this extension request
    at web3.js:3143:20
    at web3.js:6347:15
    at web3.js:5081:36
    at <anonymous>:1:1
```
### `quorumExtension_activeExtensionContracts` 
Returns the list of all active contract extensions initiated from the node
#### Parameter
None

#### Returns
* `address`: address of the private contract getting extended
* `creationData`: hash of extension creation data
* `initiator`: ethereum address which initiated the contract extension
* `managementcontractaddress`: contract address which manages the extension process
#### Examples
```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"quorumExtension_activeExtensionContracts", "id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":[{"address":"0x708f521772264f07770c12489af729f26024905c","initiator":"0xca843569e3427144cead5e4d5999a3d0ccf92b8e","managementcontractaddress":"0xb1c57951a2f3006910115eadf0f167890e99b9cb","creationData":"rC+qKetN9EbwQNhkuzgvVF7LujUiBekBuKCooJDtZit/+5x0ymXQlj/41iwcoM7SvjEstPg6BKyy1f+NgsMY5g=="}]}
```


```javascript tab="geth console"
> quorumExtension.activeExtensionContracts
[{
    address: "0x708f521772264f07770c12489af729f26024905c",
    creationData: "rC+qKetN9EbwQNhkuzgvVF7LujUiBekBuKCooJDtZit/+5x0ymXQlj/41iwcoM7SvjEstPg6BKyy1f+NgsMY5g==",
    initiator: "0xca843569e3427144cead5e4d5999a3d0ccf92b8e",
    managementcontractaddress: "0xb1c57951a2f3006910115eadf0f167890e99b9cb"
}]
```
