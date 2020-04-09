## Interacting with the Network
After following the instructions in [Getting Started](./GettingStarted.md), you should have a fully generated local Quorum network. Here are some ways you can interact with the network to try out the features of Quorum.

## Start the Network

If you haven't done so already, go into the network directory and run start.sh (`network/3-nodes-raft-tessera-bash` is the quickstart default, if you changed any settings the folder name will be different). 
```sh
cd network/3-nodes-raft-tessera-bash
./start.sh
```

Note: Run `./stop.sh` if you want to stop all quorum/geth, tessera, and cakeshop instances running on your machine

## Demonstrating Privacy
The network comes with some simple contracts to demonstrate the privacy features of Quorum.  In this demo we will:

- Send a private transaction between nodes 1 and 2
- Show that only nodes 1 and 2 are able to view the initial state of the contract
- Have Node 1 update the state of the contract and, once the block containing the updated transaction is validated by the network, again verify that only nodes 1 and 2 are able to see the updated state of the contract 

## Using Geth from the Command Line

### Sending a private transaction

Send an example private contract from Node 1 to Node 2 (this is denoted by Node 2's public key passed via `privateFor: ["QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc="]` in `private-contract.js`):
```sh
./runscript.sh private-contract.js
```
Make note of the `TransactionHash` printed to the terminal.

### Inspecting the Quorum nodes

We can inspect any of the Quorum nodes by using `geth attach` to open the Geth JavaScript console.  For this demo, we will be inspecting Node 1, Node 2, and Node 3.  

It is recommended to use separate terminal windows for each node we are inspecting.  In each terminal, ensure you are in your network's directory, then:

- In terminal 1 run `geth attach ipc:qdata/dd1/geth.ipc` to attach to node 1
- In terminal 2 run `geth attach ipc:qdata/dd2/geth.ipc` to attach to node 2
- In terminal 3 run `geth attach ipc:qdata/dd3/geth.ipc` to attach to node 3

To look at the private transaction that was just sent, run the following command in one of the terminals:
```sh
eth.getTransaction("0xe28912c5694a1b8c4944b2252d5af21724e9f9095daab47bac37b1db0340e0bf")
```
where you should replace this hash with the TransactionHash that was previously printed to the terminal.  This will print something of the form:
```sh
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

- In Node 1's log file: `qdata/logs/1.log`
- By reading the `contractAddress` param after calling `eth.getTransactionReceipt(txHash)` ([Ethereum API documentation](https://github.com/ethereum/wiki/wiki/JavaScript-API#web3ethgettransactionreceipt)) where `txHash` is the hash printed to the terminal after sending the transaction.

Once you've identified the contract address, run the following command in each terminal:
```
> var address = "0x1932c48b2bf8102ba33b4a6b545c32236e342f34"; //replace with your contract address 
``` 

Next we'll use ```eth.contract``` to define a contract class with the simpleStorage ABI definition in each terminal:
```
> var abi = [{"constant":true,"inputs":[],"name":"storedData","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"uint256"}],"name":"set","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[{"name":"initVal","type":"uint256"}],"type":"constructor"}];
> var private = eth.contract(abi).at(address)
```

The function calls are now available on the contract instance and you can call those methods on the contract. Let's start by examining the initial value of the contract to make sure that only nodes 1 and 2 can see the initialized value.
- In terminal window 1 (Node 1):
```
> private.get()
42
```
- In terminal window 2 (Node 2):
```
> private.get()
42
```
- In terminal window 3 (Node 3):
```
> private.get()
0
```

So we can see nodes 1 and 2 are able to read the state of the private contract and its initial value is 42.  If you look in `private-contract.js` you will see that this was the value set when the contract was created.  Node 3 is unable to read the state. 

### Updating the state of the contract

Next we'll have Node 1 set the state to the value `4` and verify only nodes 1 and 2 are able to view the new state.

In terminal window 1 (Node 1):
```
> private.set(4,{from:eth.accounts[0],privateFor:["QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc="]});
"0xacf293b491cccd1b99d0cfb08464a68791cc7b5bc14a9b6e4ff44b46889a8f70"
```
You can check the log files in `qdata/logs/` to see each node validating the block with this new private transaction. Once the block containing the transaction has been validated we can once again check the state from each node 1, 4, and 2.

- In terminal window 1 (Node 1):
```
> private.get()
4
```

- In terminal window 2 (Node 2):
```
> private.get()
4
```

- In terminal window 3 (Node 3):
```
> private.get()
0
```

And there you have it: All nodes are validating the same blockchain of transactions, with the private transactions containing only a 512 bit hash in place of the transaction data, and only the parties to the private transactions being able to view and update the state of the private contracts.

## Using Remix

You can also try to do all of the above steps using the Storage contract in [Remix](http://remix.ethereum.org) by using our [Quorum Plugin](../RemixPlugin/Overview.md). Follow the instructions for activating the remix plugin in [Getting Started](../RemixPlugin/GettingStarted.md), connect to the nodes using their Quorum and Tessera urls:

Node 1 defaults:

- Quorum RPC: `http://localhost:22000`
- Tessera: `http://localhost:9081`

Node 2 defaults:

- Quorum RPC: `http://localhost:22001`
- Tessera: `http://localhost:9082`

Node 3 defaults:

- Quorum RPC: `http://localhost:22002`
- Tessera: `http://localhost:9083`

## Using Cakeshop

If you chose to include Cakeshop in your network (included in the Quickstart option), you can try to do the above steps in that UI as well.

1. Open http://localhost:8999 in your browser.
2. Go to the Contracts tab and Deploy the contract registry
3. Go to the Sandbox, select the SimpleStorage sample contract from the Contract Library, and deploy with Private For set to the second node's public key (`QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=`)
4. Go back to the main Cakeshop page, go to the Contracts tab again, and you should be able to see the contract you just deployed.
5. Interact with it from there, and switch between nodes using the dropdown in the top right corner of the page.
