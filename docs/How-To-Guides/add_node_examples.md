# Node addition examples

Below are some scenarios for adding a new node into a network, with a mix of different options such as 
consensus algorithm, permissioning and discovery.

You can find the resources required to run the examples in the 
[quorum-examples](https://github.com/jpmorganchase/quorum-examples/tree/master/examples/adding_nodes) repository. 
Checkout the repository through `git` or otherwise download all the resources your local machine to follow along.

The examples use `docker-compose` for the container definitions. If you are following along by copying the commands 
described, then it is important to set the project name for Docker Compose, or to remember to change the prefix for 
your directory. See [Docker documentation](https://docs.docker.com/compose/reference/envvars/#compose_project_name) 
for more details.

To set the project name, run the following:
```bash
$ export COMPOSE_PROJECT_NAME=addnode
```

## Non-permimssioned IBFT with discovery

An example using IBFT, no permissioning and discover enabled via a bootnode.
There are no static peers in this network; instead, every node is set to talk to node 1 via the CLI flag 
`--bootnodes enode://ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef@172.16.239.11:21000`.
Node 1 will forward the details of all the nodes it knows about (in this case, everyone) and they will then initiate their
own connections.

1. Bring up an initial network of 6 nodes.

    ```bash
    # Ensure any old network is removed
    $ docker-compose -f ibft-non-perm-bootnode.yml down
    
    # Bring up 6 nodes
    $ docker-compose -f ibft-non-perm-bootnode.yml up node1 node2 node3 node4 node5 node6
    ```
   
2. Send in a public transaction and check it is minted.

    !!! note
        * The block creation period is set to 2 seconds, so you may have to wait upto that amount of time for the transaction to be minted.
        * The transaction hashes will likely be different, but the contract addresses will be the same for your network.
    
    ```bash
    # Send in the transaction
    $ docker exec -it addnode_node1_1 geth --exec 'loadScript("/examples/public-contract.js")' attach /qdata/dd/geth.ipc
    Contract transaction send: TransactionHash: 0xd1bf0c15546802e5a121f79d0d8e6f0fa45d4961ef8ab9598885d28084cfa909 waiting to be mined...
    true
        
    # Retrieve the value of the contract
    $ docker exec -it addnode_node1_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1932c48b2bf8102ba33b4a6b545c32236e342f34"); private.get();' attach /qdata/dd/geth.ipc
    42
    ```
    
    We created a transaction, in this case with hash `0xd1bf0c15546802e5a121f79d0d8e6f0fa45d4961ef8ab9598885d28084cfa909`,
    and then retrieved its value, which was set to be `42`.

3. Bring up the last node. This node also has its bootnodes set to be node 1, so at startup will try to establish a 
connection to node 1 only. After this, node 1 will share which nodes it knows about, and node 7 can then initiate 
connections with those peers.
    
    ```bash
    # Bring up node 7
    $ docker-compose -f ibft-non-perm-bootnode.yml up node7
    ```
   
4. Let's check to see if the nodes are in sync. If they are, they will have similar block numbers, which is enough for 
this example; there are other ways to tell if nodes are on the same chain, e.g. matching block hashes.

    !!! note
        Depending on timing, the second may have an extra block or two.

    ```bash
    # Fetch the latest block number for node 1
    $ docker exec -it addnode_node1_1 geth --exec 'eth.blockNumber' attach /qdata/dd/geth.ipc
    45
    
    # Fetch the latest block number for node 7
    $ docker exec -it addnode_node7_1 geth --exec 'eth.blockNumber' attach /qdata/dd/geth.ipc
    45
    ```
    
5. We can check that the transaction and contract we sent earlier now exist on node 7.

    ```bash
    $ docker exec -it addnode_node7_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1932c48b2bf8102ba33b4a6b545c32236e342f34"); private.get();' attach /qdata/dd/geth.ipc
    42
    ```

6. To be sure we have two way communication, let's send a transaction from node 7 to the network.
    
    ```bash
    $ docker exec -it addnode_node7_1 geth --exec 'loadScript("/examples/public-contract.js")' attach /qdata/dd/geth.ipc
    Contract transaction send: TransactionHash: 0x84cefc3aab8ce5797dc73c70db604e5c8830fc7c2cf215876eb34fff533e2725 waiting to be mined...
    true
    ```
   
7. Finally, we can check if the transaction was minted and the contract executed on each node.

    ```bash
    # Check on node 1
    $ docker exec -it addnode_node1_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1349f3e1b8d71effb47b840594ff27da7e603d17"); private.get();' attach /qdata/dd/geth.ipc
    42
   
    # Check on node 7
    $ docker exec -it addnode_node7_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1349f3e1b8d71effb47b840594ff27da7e603d17"); private.get();' attach /qdata/dd/geth.ipc
    42
    ```
   
And that's it. We deployed a working 6 node network, and then added a 7th node afterwards; this 7th node was able to 
read existing public data, as well as deploy its own transactions and contracts for others to see!

## Non-permissioned RAFT with discovery disabled

This example walks through adding a new node to a RAFT network. This network does not have permissioning for the 
Ethereum peer-to-peer layer, and makes it connections solely based on who is listed in the nodes `static-nodes.json` 
file.

1. Bring up an initial network of 6 nodes.

    ```bash
    # Ensure any old network is removed
    $ docker-compose -f raft-non-perm-nodiscover.yml down
    
    # Bring up 6 nodes
    $ docker-compose -f raft-non-perm-nodiscover.yml up node1 node2 node3 node4 node5 node6
    ```
   
2. Send in a public transaction and check it is minted.

    !!! note
        * The transaction hashes will likely be different, but the contract addresses will be the same for your network.
    
    ```bash
    # Send in the transaction
    $ docker exec -it addnode_node1_1 geth --exec 'loadScript("/examples/public-contract.js")' attach /qdata/dd/geth.ipc
    Contract transaction send: TransactionHash: 0xd1bf0c15546802e5a121f79d0d8e6f0fa45d4961ef8ab9598885d28084cfa909 waiting to be mined...
    true
        
    # Retrieve the value of the contract
    $ docker exec -it addnode_node1_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1932c48b2bf8102ba33b4a6b545c32236e342f34"); private.get();' attach /qdata/dd/geth.ipc
    42
    ```
    
    We created a transaction, in this case with hash `0xd1bf0c15546802e5a121f79d0d8e6f0fa45d4961ef8ab9598885d28084cfa909`,
    and then retrieved its value, which was set to be `42`.

3. We need to add the new peer to the RAFT network before it joins, otherwise the existing nodes will reject it from 
the RAFT communication layer; we also need to know what ID the new node should join with.

    ```bash
    # Add the new node
    $ docker exec -it addnode_node1_1 geth --exec 'raft.addPeer("enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@172.16.239.17:21000?discport=0&raftport=50400")' attach /qdata/dd/geth.ipc
    7
    ```
   
   The return value is the RAFT ID of the new node. When the node joins the network for the first time, it will need 
   this ID number handy. If it was lost, you can always view the full network, including IDs, by running the 
   `raft.cluster` command on an existing node.

4. Bring up the last node. Here, we pass the newly created ID number as a flag into the startup of node 7. This lets 
the node know to not bootstrap a new network from the contents of `static-nodes.json`, but to connect to an existing 
node there are fetch any bootstrap information.
    
    ```bash
    # Bring up node 7
    $ QUORUM_GETH_ARGS="--raftjoinexisting 7" docker-compose -f raft-non-perm-nodiscover.yml up node7
    ```
   
5. Let's check to see if the nodes are in sync. We can do by seeing if we have the contract that we viewer earlier on 
node 7.

    ```bash
    # Fetch the contracts value on node 7
    $ docker exec -it addnode_node7_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1932c48b2bf8102ba33b4a6b545c32236e342f34"); private.get();' attach /qdata/dd/geth.ipc
    42
    ```

6. To be sure we have two way communication, let's send a transaction from node 7 to the network.
    
    ```bash
    $ docker exec -it addnode_node7_1 geth --exec 'loadScript("/examples/public-contract.js")' attach /qdata/dd/geth.ipc
    Contract transaction send: TransactionHash: 0x84cefc3aab8ce5797dc73c70db604e5c8830fc7c2cf215876eb34fff533e2725 waiting to be mined...
    true
    ```
   
7. Finally, we can check if the transaction was minted and the contract executed on each node.

    ```bash
    # Check on node 1
    $ docker exec -it addnode_node1_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1349f3e1b8d71effb47b840594ff27da7e603d17"); private.get();' attach /qdata/dd/geth.ipc
    42
   
    # Check on node 7
    $ docker exec -it addnode_node7_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1349f3e1b8d71effb47b840594ff27da7e603d17"); private.get();' attach /qdata/dd/geth.ipc
    42
    ```
   
And that's it. We deployed a working 6 node network, and then added a 7th node afterwards; this 7th node was able to 
read existing public data, as well as deploy its own transactions and contracts for others to see!

## Permissioned RAFT with discovery disabled

This example walks through adding a new node to a RAFT network. This network does have permissioning enabled for the 
Ethereum peer-to-peer layer; this means that for any Ethereum tasks, such as syncing the initial blockchain or 
propagating transactions, the node must appear is others nodes' `permissioned-nodes.json` file.

1. Bring up an initial network of 6 nodes.

    ```bash
    # Ensure any old network is removed
    $ docker-compose -f raft-perm-nodiscover.yml down
    
    # Bring up 6 nodes
    $ docker-compose -f raft-perm-nodiscover.yml up node1 node2 node3 node4 node5 node6
    ```
   
2. Send in a public transaction and check it is minted.

    !!! note
        * The transaction hashes will likely be different, but the contract addresses will be the same for your network.
    
    ```bash
    # Send in the transaction
    $ docker exec -it addnode_node1_1 geth --exec 'loadScript("/examples/public-contract.js")' attach /qdata/dd/geth.ipc
    Contract transaction send: TransactionHash: 0xd1bf0c15546802e5a121f79d0d8e6f0fa45d4961ef8ab9598885d28084cfa909 waiting to be mined...
    true
        
    # Retrieve the value of the contract
    $ docker exec -it addnode_node1_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1932c48b2bf8102ba33b4a6b545c32236e342f34"); private.get();' attach /qdata/dd/geth.ipc
    42
    ```
    
    We created a transaction, in this case with hash `0xd1bf0c15546802e5a121f79d0d8e6f0fa45d4961ef8ab9598885d28084cfa909`,
    and then retrieved its value, which was set to be `42`.

3. We need to add the new peer to the RAFT network before it joins, otherwise the existing nodes will reject it from 
the RAFT communication layer; we also need to know what ID the new node should join with.

    ```bash
    # Add the new node
    $ docker exec -it addnode_node1_1 geth --exec 'raft.addPeer("enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@172.16.239.17:21000?discport=0&raftport=50400")' attach /qdata/dd/geth.ipc
    7
    ```
   
   The return value is the RAFT ID of the new node. When the node joins the network for the first time, it will need 
   this ID number handy. If it was lost, you can always view the full network, including IDs, by running the 
   `raft.cluster` command on an existing node.

4. Bring up the last node. Here, we pass the newly created ID number as a flag into the startup of node 7. This lets 
the node know to not bootstrap a new network from the contents of `static-nodes.json`, but to connect to an existing 
node there are fetch any bootstrap information.
    
    ```bash
    # Bring up node 7
    $ QUORUM_GETH_ARGS="--raftjoinexisting 7" docker-compose -f raft-non-perm-nodiscover.yml up node7
    ```
   
5. Let's check to see if the nodes are in sync. We can do by seeing if we have the contract that we viewer earlier on 
node 7.

    ```bash
    # Fetch the contracts value on node 7
    $ docker exec -it addnode_node7_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1932c48b2bf8102ba33b4a6b545c32236e342f34"); private.get();' attach /qdata/dd/geth.ipc
    0
    ```
   
   The value here is `0`, not the expected `42`! Node 7 is unable to sync the blockchain because the other peers in the 
   network are refusing to allow connections from node 7, due to it being missing in the `permissioned-nodes.json` file.
   
   This does not affect the RAFT layer, so if node 7 was already is sync, it could still receive new blocks; this is 
   okay though, since it would be permissioned on the RAFT side by virtue of being part of the RAFT cluster.
   
6. Let's update the permissioned nodes list on node 1, which will allow node 7 to connect to it.

    ```bash
    $ docker exec -it addnode_node1_1 cp /extradata/static-nodes-7.json /qdata/dd/permissioned-nodes.json
    $
    ```

7. Node 7 should now be synced up through node 1. Let's see if we can see the contract we made earlier.

    !!! note
        Quorum attempts to re-establish nodes every 30 seconds, so you may have to wait for the sync to happen.

    ```bash
    $ docker exec -it addnode_node7_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1932c48b2bf8102ba33b4a6b545c32236e342f34"); private.get();' attach /qdata/dd/geth.ipc
    42
    ```

8. To be sure we have two way communication, let's send a transaction from node 7 to the network.
    
    ```bash
    $ docker exec -it addnode_node7_1 geth --exec 'loadScript("/examples/public-contract.js")' attach /qdata/dd/geth.ipc
    Contract transaction send: TransactionHash: 0x84cefc3aab8ce5797dc73c70db604e5c8830fc7c2cf215876eb34fff533e2725 waiting to be mined...
    true
    ```
   
9. Finally, we can check if the transaction was minted and the contract executed on each node.

    ```bash
    # Check on node 1
    $ docker exec -it addnode_node1_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1349f3e1b8d71effb47b840594ff27da7e603d17"); private.get();' attach /qdata/dd/geth.ipc
    42
   
    # Check on node 7
    $ docker exec -it addnode_node7_1 geth --exec 'var private = eth.contract([{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"}]).at("0x1349f3e1b8d71effb47b840594ff27da7e603d17"); private.get();' attach /qdata/dd/geth.ipc
    42
    ```
   
And that's it. We deployed a working 6 node network, and then added a 7th node afterwards; this 7th node was able to 
   read existing public data, as well as deploy its own transactions and contracts for others to see!
   
## Adding a Private Transaction Manager

This is a simple example of adding a new Tessera instance to an existing network. For simplicity, 
the steps to add the Quorum node are omitted, but are those followed in the IBFT example.
Here, a Tessera node is added without any of the discovery options specified, meaning that the 
IP Whitelist isn't used, nor is key discovery disabled.

1. Start up the initial 6 node network.

    ```bash
    # Ensure any old network is removed
    $ docker-compose -f tessera-add.yml down
    
    # Bring up 6 nodes
    $ docker-compose -f tessera-add.yml up node1 node2 node3 node4 node5 node6
    ```
   
2. We can verify that private transactions can be sent by sending one from node 1 to node 6.
We can also see that since node 7 doesn't exist yet, we can't send private transactions to it.

    ```bash
    # Send a private transaction from node 1 to node 6
    $ docker exec -it addnode_node1_1 geth --exec 'loadScript("/examples/private-contract-6.js")' attach /qdata/dd/geth.ipc
    Contract transaction send: TransactionHash: 0xc8a5de4bb79d4a8c3c1156917968ca9b2965f2514732fc1cff357ec999b9aba4 waiting to be mined...
    true
    # Success! 
   
    $ docker exec -it addnode_node1_1 geth --exec 'loadScript("/examples/private-contract-7.js")' attach /qdata/dd/geth.ipc
    err creating contract Error: Non-200 status code: &{Status:404 Not Found StatusCode:404 Proto:HTTP/1.1 ProtoMajor:1 ProtoMinor:1 Header:map[Server:[Jetty(9.4.z-SNAPSHOT)] Date:[Thu, 16 Jan 2020 12:44:19 GMT] Content-Type:[text/plain] Content-Length:[73]] Body:0xc028e87d40 ContentLength:73 TransferEncoding:[] Close:false Uncompressed:false Trailer:map[] Request:0xc000287200 TLS:<nil>}
    true
    # An expected failure. The script content didn't succeed, but the script itself was run okay, so true was still returned
    ```
   
3. Let's first bring up node 7, then we can inspect what is happening and the configuration used.
    
    ```bash
    # Bring up node 7
    $ docker-compose -f tessera-add.yml up node7 
   
    $ docker exec -it addnode_node7_1 cat /qdata/tm/tessera-config.json
    # ...some output...
    ```
   
    The last command will output Tessera 7's configuration.
    The pieces we are interested in here are the following:
   
    ```json
    {
      "useWhiteList": false,
       "peer": [
          {
              "url": "http://txmanager1:9000"
          }
       ],
       ...
    }
    ```
   
   We can see that the whitelist is not enabled, discovery is not specified so defaults to enabled,
   and we have a single peer to start off with, which is node 1.
   This is all that is needed to connect to an existing network. Shortly after starting up, Tessera 
   will ask node 1 about all it's peers, and then will keep a record of them for it's own use. From 
   then on, all the nodes will know about node 7 and can send private transactions to it.
   
4. Let's try it! Let's send a private transaction from node 1 to the newly added node 7.

    ```bash
    # Sending a transaction from node 1 to node 7
    $ docker exec -it addnode_node1_1 geth --exec 'loadScript("/examples/private-contract-7.js")' attach /qdata/dd/geth.ipc
    Contract transaction send: TransactionHash: 0x3e3b50768ffdb51979677ddb58f48abdabb82a3fd4f0bac5b3d1ad8014e954e9 waiting to be mined...
    true
    ```
   
   We got a success this time! Tessera 7 has been accepted into the network and can interact with the 
   other existing nodes.