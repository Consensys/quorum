# 7nodes
## Demonstrating Privacy
The [7nodes example](https://github.com/jpmorganchase/quorum-examples/tree/master/examples/7nodes) comes with some simple contracts to demonstrate the privacy features of Quorum. 

In this demo we will:

* Send a private transaction between nodes 1 and 7
* Show that only nodes 1 and 7 are able to view the initial state of the contract
* Have Node 1 update the state of the contract and, once the block containing the updated transaction is validated by the network, again verify that only nodes 1 and 7 are able to see the updated state of the contract 

!!! tip
    [Constellation](../../Privacy/Constellation/Constellation) or [Tessera](../../Privacy/Tessera/Tessera) is used to enable the privacy features of Quorum.  To start a Quorum node without its associated privacy transaction manager, set `PRIVATE_CONFIG=ignore` when starting the node.

### Sending a private transaction

First start running the 7nodes example by following the instructions in the [quorum-examples](../Quorum-Examples#getting-started), then send an example private contract from Node 1 to Node 7 (this is denoted by the public key passed via `privateFor: ["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="]` in `private-contract.js`):
``` bash
./runscript.sh private-contract.js
```
Make note of the `TransactionHash` printed to the terminal.

### Inspecting the Quorum nodes

We can inspect any of the Quorum nodes by using `geth attach` to open the Geth JavaScript console.  For this demo, we will be inspecting Node 1, Node 7 and Node 4.  

It is recommended to use separate terminal windows for each node we are inspecting.  In each terminal, ensure you are in the `path/to/7nodes` directory, then:

- If you aren't already running the 7nodes example, in terminal 1 run `./{consensus}-init.sh` followed by `./{consensus}-start.sh`
- In terminal 1 run `geth attach ipc:qdata/dd1/geth.ipc` to attach to node 1
- In terminal 2 run `geth attach ipc:qdata/dd4/geth.ipc` to attach to node 4
- In terminal 3 run `geth attach ipc:qdata/dd7/geth.ipc` to attach to node 7

To look at the private transaction that was just sent, run the following command in one of the terminals:
``` sh
eth.getTransaction("0xe28912c5694a1b8c4944b2252d5af21724e9f9095daab47bac37b1db0340e0bf")
```
where you should replace this hash with the TransactionHash that was previously printed to the terminal.  This will print something of the form:
``` sh
{
  blockHash: "0x4d6eb0d0f971b5e0394a49e36ba660c69e62a588323a873bb38610f7b9690b34",
  blockNumber: 1,
  from: "0xed9d02e382b34818e88b88a309c7fe71e65f419d",
  gas: 4700000,
  gasPrice: 0,
  hash: "0xe28912c5694a1b8c4944b2252d5af21724e9f9095daab47bac37b1db0340e0bf",
  input: "0x58c0c680ee0b55673e3127eb26e5e537c973cd97c70ec224ccca586cc4d31ae042d2c55704b881d26ca013f15ade30df2dd196da44368b4a7abfec4a2022ec6f",
  nonce: 0,
  r: "0x4952fd6cd1350c283e9abea95a2377ce24a4540abbbf46b2d7a542be6ed7cce5",
  s: "0x4596f7afe2bd23135fa373399790f2d981a9bb8b06144c91f339be1c31ec5aeb",
  to: null,
  transactionIndex: 0,
  v: "0x25",
  value: 0
}
```

Note the `v` field value of `"0x25"` or `"0x26"` (37 or 38 in decimal) which indicates this transaction has a private payload (input). 


#### Checking the state of the contract
For each of the 3 nodes we'll use the Geth JavaScript console to create a variable called `address` which we will assign to the address of the contract created by Node 1.  The contract address can be found in two ways:  

- In Node 1's log file: `7nodes/qdata/logs/1.log`
- By reading the `contractAddress` param after calling `eth.getTransactionReceipt(txHash)` ([Ethereum API documentation](https://github.com/ethereum/wiki/wiki/JavaScript-API#web3ethgettransactionreceipt)) where `txHash` is the hash printed to the terminal after sending the transaction.

Once you've identified the contract address, run the following command in each terminal:
``` javascript
> var address = "0x1932c48b2bf8102ba33b4a6b545c32236e342f34"; //replace with your contract address 
``` 

Next we'll use ```eth.contract``` to define a contract class with the simpleStorage ABI definition in each terminal:
``` javascript
> var abi = [{"constant":true,"inputs":[],"name":"storedData","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"uint256"}],"name":"set","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[{"name":"initVal","type":"uint256"}],"type":"constructor"}];
> var private = eth.contract(abi).at(address)
```

The function calls are now available on the contract instance and you can call those methods on the contract. Let's start by examining the initial value of the contract to make sure that only nodes 1 and 7 can see the initialized value.
- In terminal window 1 (Node 1):
``` javascript
> private.get()
42
```
- In terminal window 2 (Node 4):
``` javascript
> private.get()
0
```
- In terminal window 3 (Node 7):
``` javascript
> private.get()
42
```

So we can see nodes 1 and 7 are able to read the state of the private contract and its initial value is 42.  If you look in `private-contract.js` you will see that this was the value set when the contract was created.  Node 4 is unable to read the state. 

### Updating the state of the contract

Next we'll have Node 1 set the state to the value `4` and verify only nodes 1 and 7 are able to view the new state.

In terminal window 1 (Node 1):
``` javascript
> private.set(4,{from:eth.coinbase,privateFor:["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="]});
"0xacf293b491cccd1b99d0cfb08464a68791cc7b5bc14a9b6e4ff44b46889a8f70"
```
You can check the log files in `7nodes/qdata/logs/` to see each node validating the block with this new private transaction. Once the block containing the transaction has been validated we can once again check the state from each node 1, 4, and 7.
- In terminal window 1 (Node 1):
``` javascript
> private.get()
4
```
- In terminal window 2 (Node 4):
``` javascript
> private.get()
0
```
- In terminal window 3 (Node 7):
``` javascript
> private.get()
4
```
And there you have it. All 7 nodes are validating the same blockchain of transactions, the private transactions carrying only a 512 bit hash, and only the parties to private transactions are able to view and update the state of private contracts.

## Permissions

Node Permissioning is a feature in Quorum that allows only a pre-defined set of nodes (as identified by their remotekey/enodes) to connect to the permissioned network.

In this demo we will:

- Set up a network with a combination of permissioned and non-permissioned nodes in the cluster
- Look at the details of the `permissioned-nodes.json` file
- Demonstrate that only the nodes that are specified in `permissioned-nodes.json` can connect to the network

### Verify only permissioned nodes are connected to the network.

Attach to the individual nodes via `geth attach path/to/geth.ipc` and use `admin.peers` to check the connected nodes:

``` sh
❯ geth attach qdata/dd1/geth.ipc
Welcome to the Geth JavaScript console!

instance: Geth/v1.7.2-stable/darwin-amd64/go1.9.2
coinbase: 0xed9d02e382b34818e88b88a309c7fe71e65f419d
at block: 1 (Mon, 29 Oct 47909665359 22:09:51 EST)
 datadir: /Users/joel/jpm/quorum-examples/examples/7nodes/qdata/dd1
 modules: admin:1.0 debug:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 raft:1.0 rpc:1.0 txpool:1.0 web3:1.0

> admin.peers
[{
    caps: ["eth/63"],
    id: "0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416",
    name: "Geth/v1.7.2-stable/darwin-amd64/go1.9.2",
    network: {
      localAddress: "127.0.0.1:65188",
      remoteAddress: "127.0.0.1:21001"
    },
    protocols: {
      eth: {
        difficulty: 0,
        head: "0xc23b4ebccc79e2636d66939924d46e618269ca1beac5cf1ec83cc862b88b1b71",
        version: 63
      }
    }
},
...
]
```

You can also inspect the log files under `qdata/logs/*.log` for further diagnostics messages around incoming / outgoing connection requests. `grep` for `ALLOWED-BY` or `DENIED-BY`. Be sure to enable verbosity for p2p module.

### Permissioning configuration

Permissioning is granted based on the remote key of the geth node. The remote keys are specified in the `permissioned-nodes.json` and is placed under individual node's `<datadir>`.

The below sample `permissioned-nodes.json` provides a list of nodes permissioned to join the network (node ids truncated for clarity):

``` json
[
   "enode://8475a01f22a1f48116dc1f0d22ecaaaf77e@127.0.0.1:30301",
   "enode://b5660501f496e60e59ded734a889c97b7da@127.0.0.1:30302",
   "enode://54bd7ff4bd971fb80493cf4706455395917@127.0.0.1:30303"
]
```

### Enabling/Disabling permissions

An individual node can enable/disable permissioning by passing the `-permissioned` command line flag. If enabled, then only the nodes that are in the `<datadir>/permissioned-nodes.json` can connect to it. Further, these are the only nodes that this node can make outbound connections to as well.

```
MISCELLANEOUS OPTIONS:
--permissioned          If enabled, the node will allow only a defined list of nodes to connect
```

## Next steps
Additional samples can be found in `quorum-examples/examples/7nodes/samples` for you to use and edit.  You can also create your own contracts to help you understand how the nodes in a Quorum network work together.

## Reducing the number of nodes 
It is easy to reduce the number of nodes used in the example.  You may want to do this for memory usage reasons or just to experiment with a different network configuration.

To run the example with 5 nodes instead of 7, the following changes need to be made:
1. In __`raft-start.sh`__:
 
    Comment out the following lines used to start Quorum nodes 6 & 7
   ```sh
   # PRIVATE_CONFIG=qdata/c6/tm.ipc nohup geth --datadir qdata/dd6 $ARGS --raftport 50406 --rpcport 22005 --port 21005 --unlock 0 --password passwords.txt 2>>qdata/logs/6.log &
   # PRIVATE_CONFIG=qdata/c7/tm.ipc nohup geth --datadir qdata/dd7 $ARGS --raftport 50407 --rpcport 22006 --port 21006 --unlock 0 --password passwords.txt 2>>qdata/logs/7.log &
   ```

1. In __`constellation-start.sh`__ or __`tessera-start.sh`__ (depending on which privacy manager you are using): 

    Change the 2 instances of `for i in {1..7}` to `for i in {1..5}`
    
After making these changes, the `raft-init.sh` and `raft-start.sh` scripts can be run as normal.

`private-contract.js` will also need to be updated as this is set up to send a transaction from node 1 to node 7.  To update the private contract to instead send to node 5, the following steps need to be followed:

1. Copy node 5's public key from `./keys/tm5.pub`

2. Replace the existing `privateFor` in `private-contract.js` with the key copied from `tm5.pub` key, e.g.:
    ``` javascript
    var simple = simpleContract.new(42, {from:web3.eth.accounts[0], data: bytecode, gas: 0x47b760, privateFor: ["R56gy4dn24YOjwyesTczYa8m5xhP6hF2uTMCju/1xkY="]}, function(e, contract) {...}
    ```

After saving this change, the `./runscript.sh private-contract.js` command can be run as usual to submit the private contract.  You can then follow steps described above to verify that node 5 can see the transaction payload and that nodes 2-4 are unable to see the payload.

## Using a Tessera remote enclave
Tessera v0.9 brought with it the option to have an enclave as a separate process from the Transaction
Manager. This is a more secure way of being able to manage and interact with your keys.
To use the remote enclave, call your desired start with using `tessera-remote` as the first
parameter, e.g. `./raft-start.sh tessera-remote`. This will, by default, start 7 Transaction
Managers, the first 4 of which use a remote enclave. If you wish to change this number, you
will need to add the extra parameter `--remoteEnclaves X` in the `--tesseraOptions`, e.g.
`./raft-start.sh tessera-remote --tesseraOptions "--remoteEnclaves 7"`.
