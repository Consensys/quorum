# Raft RPC API
# APIs
### raft_cluster
Returns the details of all nodes part of the raft cluster
#### Parameters
None
#### Returns
* `hostName`: DNS name or the host IP address 
* `nodeActive`: true if the node is active in raft cluster else false
* `nodeId`: enode id of the node
* `p2pPort`: p2p port 
* `raftId`: raft id of the node
* `raftPort`: raft port 
* `role`: role of the node in raft quorum. Can be minter/ verifier/ learner. In case there is no leader at network level it will be returned as `""`
#### Examples
```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"raft_cluster", "id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":[{"raftId":1,"nodeId":"ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef","p2pPort":21000,"raftPort":50401,"hostname":"127.0.0.1","role":"minter","nodeActive":true},{"raftId":3,"nodeId":"579f786d4e2830bbcc02815a27e8a9bacccc9605df4dc6f20bcc1a6eb391e7225fff7cb83e5b4ecd1f3a94d8b733803f2f66b7e871961e7b029e22c155c3a778","p2pPort":21002,"raftPort":50403,"hostname":"127.0.0.1","role":"verifier","nodeActive":true},{"raftId":2,"nodeId":"0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416","p2pPort":21001,"raftPort":50402,"hostname":"127.0.0.1","role":"verifier","nodeActive":true}]}
```

```javascript tab="geth console"
> raft.cluster
[{
    hostname: "127.0.0.1",
    nodeActive: true,
    nodeId: "0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416",
    p2pPort: 21001,
    raftId: 2,
    raftPort: 50402,
    role: "verifier"
}, {
    hostname: "127.0.0.1",
    nodeActive: true,
    nodeId: "579f786d4e2830bbcc02815a27e8a9bacccc9605df4dc6f20bcc1a6eb391e7225fff7cb83e5b4ecd1f3a94d8b733803f2f66b7e871961e7b029e22c155c3a778",
    p2pPort: 21002,
    raftId: 3,
    raftPort: 50403,
    role: "verifier"
}, {
    hostname: "127.0.0.1",
    nodeActive: true,
    nodeId: "ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef",
    p2pPort: 21000,
    raftId: 1,
    raftPort: 50401,
    role: "minter"
}]
```
### raft_role 
Returns the role of the current node in raft cluster
#### Parameters
None
#### Returns
* `result`: role of the node in raft cluster. Can be minter/ verifier/ learner. In case there is no leader at network level it will be returned as `""`
#### Examples
```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"raft_role", "id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"verifier"}
```

```javascript tab="geth console"
> raft.role
"minter"
```
### raft_leader
Returns enode id of the leader node
#### Parameters
None
#### Returns
* `result`: enode id of the leader
#### Examples
```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"raft_leader", "id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef"}
```


```javascript tab="geth console"
> raft.leader
"ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef"
```

If there is no leader at the network level, the call to the api will result in the following error:
```javascript
> raft.leader
Error: no leader is currently elected
    at web3.js:3143:20
    at web3.js:6347:15
    at get (web3.js:6247:38)
    at <unknown>
```

### raft_addPeer
API for adding a new peer to the network. 
#### Parameters
* `enodeId`: enode id of the node to be added to the network
#### Returns
* `result`: raft id for the node being added
#### Examples
```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"raft_addPeer","params": ["enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405"], "id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":5}
```

```javascript tab="geth console"
> raft.addPeer("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405")
5
```

The new node can join the network with `geth` option of `--raftjoinexisting <<raftId>>`

If the node being added is already part of the network the of the network, the following error is thrown:
```javascript
> raft.addPeer("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405")
Error: node with this enode has already been added to the cluster: f06c06f1e958cb2edf90d8bfb912de287f9b047b4228436e94b5b78e3ee16171
    at web3.js:3143:20
    at web3.js:6347:15
    at web3.js:5081:36
    at <anonymous>:1:1
```
### raft_removePeer
API to remove a node from raft cluster
#### Parameters
* `raftId` : raft id of  the node to be removed from the cluster
#### Returns
* `result`: null
#### Examples
```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"raft_removePeer","params": [4], "id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":null}
```

```javascript tab="geth console"
> raft.removePeer(4)
null
```

### raft_addLearner
API to add a new node to the network as a learner node. The learner node syncs with network and can transact but will not be part of raft quorum and hence will not provide block confirmation to minter node. 
#### Parameters
* `enodeId`
#### Returns
* `result`: raft id for the node being added
#### Examples
```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"raft_addLearner","params": ["enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405"], "id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":5}

```

```javascript tab="geth console"
> raft.addLearner("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405")
5
```
### raft_promoteToPeer
API for promoting a learner node to peer and thus be part of the raft quorum. 
#### Parameters
* `raftId`: raft id of the node to be promoted
#### Returns
* `result`: true or false
#### Examples
```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22001 --data '{"jsonrpc":"2.0","method":"raft_promoteToPeer","params": [4], "id":10}' --header "Content-Type: application/json"

// Response
{// Response
 {"jsonrpc":"2.0","id":10,"result":true}
```

```javascript tab="geth console"
> raft.promoteToPeer(4)
true
```
