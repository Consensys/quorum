# Contract Extension APIs
## APIs
### `quorumExtension_extendContract` 
The api for initiating contract extension to a new node.  
#### Parameter
* `toExtend`: address of the private contract which is being extended to the new node
* `newRecipientPtmPublicKey`: Tessera public key of the recipient node
* `voters`: array of ethereum addresses who are required to approve the contract extension
* `txArgs`: standard transaction arguments with `privateFor` info of all nodes involved in contract extension

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"quorumExtension_extendContract", "params":["0x708f521772264f07770c12489af729f26024905c", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8=", ["0xed9d02e382b34818e88b88a309c7fe71e65f419d","0x9186eb3d20cbd1f5f992a950d808c4495153abd5", "0xca843569e3427144cead5e4d5999a3d0ccf92b8e"],{"from": "0xca843569e3427144cead5e4d5999a3d0ccf92b8e", "privateFor":["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="]}], "id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"0x7652916c94961836a6ea57d9af8c655d21f3c206da487122014136213207d2e3"}
```

```javascript tab="geth console"
> quorumExtension.extendContract("0x708f521772264f07770c12489af729f26024905c", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8=", ["0xed9d02e382b34818e88b88a309c7fe71e65f419d","0x9186eb3d20cbd1f5f992a950d808c4495153abd5","0xca843569e3427144cead5e4d5999a3d0ccf92b8e"],{from: "0xca843569e3427144cead5e4d5999a3d0ccf92b8e", privateFor:["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="]})

"0x9e0101dd215281b33989b3ae093147e9009353bb63f531490409a628c6e87310"
```

If the contract is already under the process of extension, api call to extend it again will fail. 

```javascript
> quorumExtension.extendContract("0x708f521772264f07770c12489af729f26024905c", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8=", ["0xed9d02e382b34818e88b88a309c7fe71e65f419d","0x9186eb3d20cbd1f5f992a950d808c4495153abd5", eth.accounts[0]],{from: eth.accounts[0], privateFor:["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="]}
Error: contract extension in progress for the given contract address
    at web3.js:3143:20
    at web3.js:6347:15
    at web3.js:5081:36
    at <anonymous>:1:1
```

### `quorumExtension_approveExtension` 
The api for approving the contract extension to a new node. This can be invoked by the ethereum address which is marked as a voter in the contract extension call
#### Parameter
* `addressVoteOn`: address of the contract extension management contract
* `vote`: bool - `true` indicates approval and `false` disapproval. Contract extension is completed only when all voters vote true. If any participant votes `false` the extension process will be cancelled.

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure
* `txArgs`: standard transaction arguments with `privateFor` info of all nodes involved in contract extension
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumExtension_approveExtension", "params":["0xb1c57951a2f3006910115eadf0f167890e99b9cb", true, {"from": "0xed9d02e382b34818e88b88a309c7fe71e65f419d", "privateFor":["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="]}], "id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"0x8d34a594b286087f45029daad2d5a8fd42f70abb0ae2492429a256a2ba4cb0dd"}
```


```javascript tab="geth console"
> quorumExtension.approveExtension("0xb1c57951a2f3006910115eadf0f167890e99b9cb", true ,{from: "], privateFor:["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="]})

"0x9e0101dd215281b33989b3ae093147e9009353bb63f531490409a628c6e87310"
```


If the contract is already under the process of extension, api call to extend it again will fail.  
```javascript
> quorumExtension.extendContract("0x708f521772264f07770c12489af729f26024905c", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8=", ["0xed9d02e382b34818e88b88a309c7fe71e65f419d","0x9186eb3d20cbd1f5f992a950d808c4495153abd5", eth.accounts[0]],{from: "0x9186eb3d20cbd1f5f992a950d808c4495153abd5", privateFor:["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="]}

"0xcc9b462c5bda5f5579738689304d86dde29f9cd3c0e1053ccc73b7a9fa78efbd"
```

A voter can vote only once. If the account is not a voter the api call will return an error. 
```javascript
> quorumExtension.approveExtension("0xab6669a499938b6fd7a4d9374e7b9a4aee6243b5", true, {from: "0x9186eb3d20cbd1f5f992a950d808c4495153abd5", privateFor: ["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="]})
Error: already voted
    at web3.js:3143:20
    at web3.js:6347:15
    at web3.js:5081:36
    at <anonymous>:1:1
    
> quorumExtension.approveExtension("0xab6669a499938b6fd7a4d9374e7b9a4aee6243b5", true, {from: "0x5907274243eed888265209c97420cd0fdfda4b59", privateFor: ["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="]})
Error: account is not a voter of this extension request
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
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"quorumExtension_cancelExtension", "params":["0xab6669a499938b6fd7a4d9374e7b9a4aee6243b5",{"from": "0xca843569e3427144cead5e4d5999a3d0ccf92b8e", "privateFor":["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="]}], "id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"0xb43da7dbeae5347df86c6933786b8c536b4622463b577a990d4c87214845d16a"}
```


```javascript tab="geth console"
> quorumExtension.cancelExtension("0xa501afd7d6432718daf4458cfae8590d88de818e",  {from: "0xed9d02e382b34818e88b88a309c7fe71e65f419d", privateFor: ["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="]})

"0x9e0101dd215281b33989b3ae093147e9009353bb63f531490409a628c6e87310"
```

If the api is invoked by an ethereum address which is not the creator of the contract extension, an error is thrown.  
```javascript
> quorumExtension.extendContract("0x708f521772264f07770c12489af729f26024905c", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8=", ["0xed9d02e382b34818e88b88a309c7fe71e65f419d","0x9186eb3d20cbd1f5f992a950d808c4495153abd5", eth.accounts[0]],{from: "0x9186eb3d20cbd1f5f992a950d808c4495153abd5", privateFor:["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=","QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="]}

"0xcc9b462c5bda5f5579738689304d86dde29f9cd3c0e1053ccc73b7a9fa78efbd"
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
