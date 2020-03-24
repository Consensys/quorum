# Creating a network from scratch

This section details easy to follow step by step instructions of how to setup one or more Quorum nodes from scratch for all new starters.

Let's go through step by step instructions to setup a Quorum node with Raft consensus.

## Quorum with Raft consensus
1. On each machine build Quorum as described in the [Installing](../Installing) section. Ensure that PATH contains geth and bootnode
    ```
    $ git clone https://github.com/jpmorganchase/quorum.git
    $ cd quorum
    $ make all
    $ export PATH=$(pwd)/build/bin:$PATH
    ```

2. Create a working directory which will be the base for the new node(s) and change into it
    ```
    $ mkdir fromscratch
    $ cd fromscratch
    $ mkdir new-node-1
    ```
3. Generate one or more accounts for this node and take down the account address. A funded account may be required depending what you are trying to accomplish
    ```
    $ geth --datadir new-node-1 account new
    
     INFO [06-07|14:52:18.742] Maximum peer count                       ETH=25 LES=0 total=25
     Your new account is locked with a password. Please give a password. Do not forget this password.
     Passphrase: 
     Repeat passphrase: 
     Address: {679fed8f4f3ea421689136b25073c6da7973418f}
     
     Please note the keystore file generated inside new-node-1 includes the address in the last part of its filename.
     
     $ ls new-node-1/keystore
     UTC--2019-06-17T09-29-06.665107000Z--679fed8f4f3ea421689136b25073c6da7973418f
    ```

    !!! note 
        You could generate multiple accounts for a single node, or any number of accounts for additional nodes and pre-allocate them with funds in the genesis.json file (see below)
       
4. Create a `genesis.json` file see example [here](../genesis). The `alloc` field should be pre-populated with the account you generated at previous step
    ```
    $ vim genesis.json
    ... alloc holds 'optional' accounts with a pre-funded amounts. In this example we are funding the accounts 679fed8f4f3ea421689136b25073c6da7973418f (generated from the step above) and c5c7b431e1629fb992eb18a79559f667228cd055.
    {
      "alloc": {
        "0x679fed8f4f3ea421689136b25073c6da7973418f": {
          "balance": "1000000000000000000000000000"
        },
       "0xc5c7b431e1629fb992eb18a79559f667228cd055": {
          "balance": "2000000000000000000000000000"
        }
    },
     "coinbase": "0x0000000000000000000000000000000000000000",
     "config": {
       "homesteadBlock": 0,
       "byzantiumBlock": 0,
       "constantinopleBlock": 0,
       "chainId": 10,
       "eip150Block": 0,
       "eip155Block": 0,
       "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
       "eip158Block": 0,
       "maxCodeSize": 35,
       "maxCodeSizeChangeBlock" : 0,
       "isQuorum": true
     },
     "difficulty": "0x0",
     "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000",
     "gasLimit": "0xE0000000",
     "mixhash": "0x00000000000000000000000000000000000000647572616c65787365646c6578",
     "nonce": "0x0",
     "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
     "timestamp": "0x00"
    }   
    ```
5. Generate node key and copy it into datadir
    ```
    $ bootnode --genkey=nodekey
    $ cp nodekey new-node-1/
    ```
6. Execute below command to display enode id of the new node
    ```
    $ bootnode --nodekey=new-node-1/nodekey --writeaddress > new-node-1/enode
    $ cat new-node-1/enode  
    '70399c3d1654c959a02b73acbdd4770109e39573a27a9b52bd391e5f79b91a42d8f2b9e982959402a97d2cbcb5656d778ba8661ec97909abc72e7bb04392ebd8'
    ```
7. Create a file called `static-nodes.json` and edit it to match this [example](../permissioned-nodes). Your file should contain a single line for your node with your enode's id and the ports you are going to use for devp2p and raft. Ensure that this file is in your nodes data directory
    ```
    $ vim static-nodes.json
    .... paste below lines with enode generated in previous step, port 21000;IP 127.0.0.1 and raft port set as 50000
    [
      "enode://70399c3d1654c959a02b73acbdd4770109e39573a27a9b52bd391e5f79b91a42d8f2b9e982959402a97d2cbcb5656d778ba8661ec97909abc72e7bb04392ebd8@127.0.0.1:21000?discport=0&raftport=50000"
    ] 
    $ cp static-nodes.json new-node-1
    ```
8. Initialize new node with below command.
    ```
    $ geth --datadir new-node-1 init genesis.json
    
    INFO [06-07|15:45:17.508] Maximum peer count                       ETH=25 LES=0 total=25
    INFO [06-07|15:45:17.516] Allocated cache and file handles         database=/Users/krish/new-node-1/geth/chaindata cache=16 handles=16
    INFO [06-07|15:45:17.524] Writing custom genesis block 
    INFO [06-07|15:45:17.524] Persisted trie from memory database      nodes=1 size=152.00B time=75.344µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-07|15:45:17.525] Successfully wrote genesis state         database=chaindata                              hash=ec0542…9665bf
    INFO [06-07|15:45:17.525] Allocated cache and file handles         database=/Users/krish/new-node-1/geth/lightchaindata cache=16 handles=16
    INFO [06-07|15:45:17.527] Writing custom genesis block 
    INFO [06-07|15:45:17.527] Persisted trie from memory database      nodes=1 size=152.00B time=60.76µs  gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-07|15:45:17.527] Successfully wrote genesis state         database=lightchaindata                              hash=ec0542…9665bf
    ```
9. Start your node by first creating a script as below and then starting it:
    ```
    $ vim startnode1.sh
    ... paste below commands. It will start it in the background.
    #!/bin/bash
    PRIVATE_CONFIG=ignore nohup geth --datadir new-node-1 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50000 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft --emitcheckpoints --port 21000 >> node.log 2>&1 &

    $ chmod +x startnode1.sh 
    $ ./startnode1.sh
    ``` 
    
    !!! note
        This configuration starts Quorum without privacy support as could be evidenced in prefix `PRIVATE_CONFIG=ignore`, please see below sections on [how to enable privacy with privacy transaction managers](#adding-privacy-transaction-manager).

    Your node is now operational and you may attach to it with below commands. 
    ```
    $ geth attach new-node-1/geth.ipc
    Welcome to the Geth JavaScript console!
    
    instance: Geth/v1.8.18-stable-bb88608c(quorum-v2.2.3)/darwin-amd64/go1.10.2
    coinbase: 0xedf53f5bf40c99f48df184441137659aed899c48
    at block: 0 (Thu, 01 Jan 1970 01:00:00 BST)
     datadir: /Users/krish/fromscratch/new-node-1
     modules: admin:1.0 debug:1.0 eth:1.0 ethash:1.0 miner:1.0 net:1.0 personal:1.0 raft:1.0 rpc:1.0 txpool:1.0 web3:1.0
    
    > raft.cluster
    [{
        ip: "127.0.0.1",
        nodeId: "a5596803caebdc9c5326e1a0775563ad8e4aa14aa3530f0ae16d3fd8d7e48bc0b81342064e22094ab5d10303ab5721650af561f2bcdc54d705f8b6a8c07c94c3",
        p2pPort: 21000,
        raftId: 1,
        raftPort: 50000
    }]
    > raft.leader
    "a5596803caebdc9c5326e1a0775563ad8e4aa14aa3530f0ae16d3fd8d7e48bc0b81342064e22094ab5d10303ab5721650af561f2bcdc54d705f8b6a8c07c94c3"
    > raft.role
    "minter"
    > 
    > exit
    ```

### Adding additional node
1. Complete steps 1, 2, 5, and 6 from the previous guide

    ```
    $ mkdir new-node-2
    $ bootnode --genkey=nodekey2
    $ cp nodekey2 new-node-2/nodekey
    $ bootnode --nodekey=new-node-2/nodekey --writeaddress
    56e81550db3ccbfb5eb69c0cfe3f4a7135c931a1bae79ea69a1a1c6092cdcbea4c76a556c3af977756f95d8bf9d7b38ab50ae070da390d3abb3d7e773099c1a9
    ```
2. Retrieve current chains `genesis.json` and `static-nodes.json`. `static-nodes.json` should be placed into new nodes data dir
    ```
    $ cp static-nodes.json new-node-2    
    ```

3. Edit `static-nodes.json` and add new entry for the new node you are configuring (should be last)
   ```
   $ vim new-node-2/static-nodes.json 
   .... append new-node-2's enode generated in step 1, port 21001;IP 127.0.0.1 and raft port set as 50001

   [
     "enode://70399c3d1654c959a02b73acbdd4770109e39573a27a9b52bd391e5f79b91a42d8f2b9e982959402a97d2cbcb5656d778ba8661ec97909abc72e7bb04392ebd8@127.0.0.1:21000?discport=0&raftport=50000",
     "enode://56e81550db3ccbfb5eb69c0cfe3f4a7135c931a1bae79ea69a1a1c6092cdcbea4c76a556c3af977756f95d8bf9d7b38ab50ae070da390d3abb3d7e773099c1a9@127.0.0.1:21001?discport=0&raftport=50001"
   ]
   ```  
    
4. Initialize new node as given below: 

    ```
    $ geth --datadir new-node-2 init genesis.json
    
    INFO [06-07|16:34:39.805] Maximum peer count                       ETH=25 LES=0 total=25
    INFO [06-07|16:34:39.814] Allocated cache and file handles         database=/Users/krish/fromscratch/new-node-2/geth/chaindata cache=16 handles=16
    INFO [06-07|16:34:39.816] Writing custom genesis block 
    INFO [06-07|16:34:39.817] Persisted trie from memory database      nodes=1 size=152.00B time=59.548µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-07|16:34:39.817] Successfully wrote genesis state         database=chaindata                                          hash=f02d0b…ed214a
    INFO [06-07|16:34:39.817] Allocated cache and file handles         database=/Users/krish/fromscratch/new-node-2/geth/lightchaindata cache=16 handles=16
    INFO [06-07|16:34:39.819] Writing custom genesis block 
    INFO [06-07|16:34:39.819] Persisted trie from memory database      nodes=1 size=152.00B time=43.733µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-07|16:34:39.819] Successfully wrote genesis state         database=lightchaindata                                          hash=f02d0b…ed214a
    ```
    
5. Connect to an already running node of the chain and execute raft `addPeer` command.
    ```
    $ geth attach new-node-1/geth.ipc
    Welcome to the Geth JavaScript console!
    
    instance: Geth/v1.8.18-stable-bb88608c(quorum-v2.2.3)/darwin-amd64/go1.10.2
    coinbase: 0xedf53f5bf40c99f48df184441137659aed899c48
    at block: 0 (Thu, 01 Jan 1970 01:00:00 BST)
     datadir: /Users/krish/fromscratch/new-node-1
     modules: admin:1.0 debug:1.0 eth:1.0 ethash:1.0 miner:1.0 net:1.0 personal:1.0 raft:1.0 rpc:1.0 txpool:1.0 web3:1.0
    
    > raft.addPeer('enode://56e81550db3ccbfb5eb69c0cfe3f4a7135c931a1bae79ea69a1a1c6092cdcbea4c76a556c3af977756f95d8bf9d7b38ab50ae070da390d3abb3d7e773099c1a9@127.0.0.1:21001?discport=0&raftport=50001')
    2
    > raft.cluster
    [{
        ip: "127.0.0.1",
        nodeId: "56e81550db3ccbfb5eb69c0cfe3f4a7135c931a1bae79ea69a1a1c6092cdcbea4c76a556c3af977756f95d8bf9d7b38ab50ae070da390d3abb3d7e773099c1a9",
        p2pPort: 21001,
        raftId: 2,
        raftPort: 50001
    }, {
        ip: "127.0.0.1",
        nodeId: "70399c3d1654c959a02b73acbdd4770109e39573a27a9b52bd391e5f79b91a42d8f2b9e982959402a97d2cbcb5656d778ba8661ec97909abc72e7bb04392ebd8",
        p2pPort: 21000,
        raftId: 1,
        raftPort: 50000
    }]
    > exit
    ```
6. Start your node by first creating a script as previous step and changing the ports you are going to use for Devp2p and raft.
    ```
    $ cp startnode1.sh startnode2.sh
    $ vim startnode2.sh
    ..... paste below details
    #!/bin/bash
    PRIVATE_CONFIG=ignore nohup geth --datadir new-node-2 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50001 --raftjoinexisting 2 --rpc --rpcaddr 0.0.0.0 --rpcport 22001 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft --emitcheckpoints --port 21001 2>>node2.log &
    
    $ ./startnode2.sh 
    ```

7. Optional: share new `static-nodes.json` with all other chain participants
    ```
    $ cp new-node-2/static-nodes.json new-node-1
    ```

    Your additional node is now operational and is part of the same chain as the previously set up node.

### Removing node
1. Connect to an already running node of the chain and execute `raft.cluster` and get the `RAFT_ID` corresponding to the node that needs to be removed
2. Run `raft.removePeer(RAFT_ID)` 
3. Stop the `geth` process corresponding to the node that was removed.
    ```
    $ geth attach new-node-1/geth.ipc
    Welcome to the Geth JavaScript console!
    
    instance: Geth/v1.8.18-stable-bb88608c(quorum-v2.2.3)/darwin-amd64/go1.10.2
    coinbase: 0xedf53f5bf40c99f48df184441137659aed899c48
    at block: 0 (Thu, 01 Jan 1970 01:00:00 BST)
     datadir: /Users/krish/fromscratch/new-node-1
     modules: admin:1.0 debug:1.0 eth:1.0 ethash:1.0 miner:1.0 net:1.0 personal:1.0 raft:1.0 rpc:1.0 txpool:1.0 web3:1.0
    
    > raft.cluster
    [{
        ip: "127.0.0.1",
        nodeId: "56e81550db3ccbfb5eb69c0cfe3f4a7135c931a1bae79ea69a1a1c6092cdcbea4c76a556c3af977756f95d8bf9d7b38ab50ae070da390d3abb3d7e773099c1a9",
        p2pPort: 21001,
        raftId: 2,
        raftPort: 50001
    }, {
        ip: "127.0.0.1",
        nodeId: "a5596803caebdc9c5326e1a0775563ad8e4aa14aa3530f0ae16d3fd8d7e48bc0b81342064e22094ab5d10303ab5721650af561f2bcdc54d705f8b6a8c07c94c3",
        p2pPort: 21000,
        raftId: 1,
        raftPort: 50000
    }]
    > 
    > raft.removePeer(2)
    null
    > raft.cluster
    [{
        ip: "127.0.0.1",
        nodeId: "a5596803caebdc9c5326e1a0775563ad8e4aa14aa3530f0ae16d3fd8d7e48bc0b81342064e22094ab5d10303ab5721650af561f2bcdc54d705f8b6a8c07c94c3",
        p2pPort: 21000,
        raftId: 1,
        raftPort: 50000
    }]
    > exit
    $
    $
    $ ps | grep geth
      PID TTY           TIME CMD
    10554 ttys000    0:00.01 -bash
     9125 ttys002    0:00.50 -bash
    10695 ttys002    0:31.42 geth --datadir new-node-1 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50000 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,db,eth,debug,miner,net,shh,txpoo
    10750 ttys002    0:01.94 geth --datadir new-node-2 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50001 --raftjoinexisting 2 --rpc --rpcaddr 0.0.0.0 --rpcport 22001 --rpcapi admin,db,eth,debu
    $ kill 10750
    $
    $
    $ ps
      PID TTY           TIME CMD
    10554 ttys000    0:00.01 -bash
     9125 ttys002    0:00.51 -bash
    10695 ttys002    0:31.76 geth --datadir new-node-1 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50000 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,db,eth,debug,miner,net,shh,txpoo
    ```



## Quorum with Istanbul BFT consensus
1. On each machine build Quorum as described in the [Installing](../Installing) section. Ensure that PATH contains geth and boot node
    ```
    $ git clone https://github.com/jpmorganchase/quorum.git
    $ cd quorum
    $ make all
    $ export PATH=$(pwd)/build/bin:$PATH
    ```
2. Install [istanbul-tools](https://github.com/jpmorganchase/istanbul-tools)
    ```
    $ mkdir fromscratchistanbul
    $ cd fromscratchistanbul
    $ git clone https://github.com/jpmorganchase/istanbul-tools.git
    $ cd istanbul-tools
    $ make
    ```    
3. Create a working directory for each of the X number of initial validator nodes
    ```
    $ mkdir node0 node1 node2 node3 node4
    ```
4. Change into the lead (whichever one you consider first) node's working directory and generate the setup files for X initial validator nodes by executing `istanbul setup --num X --nodes --quorum --save --verbose` **only execute this instruction once, i.e. not X times**. This command will generate several items of interest: `static-nodes.json`, `genesis.json`, and nodekeys for all the initial validator nodes which will sit in numbered directories from 0 to X-1
    ```
    $ cd node0
    $ ../istanbul-tools/build/bin/istanbul setup --num 5 --nodes --quorum --save --verbose
    validators
    {
    	"Address": "0x4c1ccd426833b9782729a212c857f2f03b7b4c0d",
    	"Nodekey": "fe2725c4e8f7617764b845e8d939a65c664e7956eb47ed7d934573f16488efc1",
    	"NodeInfo": "enode://dd333ec28f0a8910c92eb4d336461eea1c20803eed9cf2c056557f986e720f8e693605bba2f4e8f289b1162e5ac7c80c914c7178130711e393ca76abc1d92f57@0.0.0.0:30303?discport=0"
    }
    {
    	"Address": "0x189d23d201b03ae1cf9113672df29a5d672aefa3",
    	"Nodekey": "3434f9efd184f2255f8acc9f4408a5068bd5ae920548044087578ab97ef22f3a",
    	"NodeInfo": "enode://1bb6be462f27e56f901c3fcb2d53a9273565f48e5d354c08f0c044405b29291b405b9f5aa027f3a75f9b058cb43e2f54719f15316979a0e5a2b760fff4631998@0.0.0.0:30303?discport=0"
    }
    {
    	"Address": "0x44b07d2c28b8ed8f02b45bd84ac7d9051b3349e6",
    	"Nodekey": "8183051c9976200d245c59a80ae004f20c3f66e1aa1b8f17458931de91576e05",
    	"NodeInfo": "enode://0df02e94a3befc0683780d898119d3b675e5942c1a2f9ad47d35b4e6ccaf395cd71ec089fcf1d616748bf9871f91e5e3d29c1cf6f8f81de1b279082a104f619d@0.0.0.0:30303?discport=0"
    }
    {
    	"Address": "0xc1056df7c02b6f1a353052eaf0533cc7cb743b52",
    	"Nodekey": "de415c5dbbb9ff0a34dbd3bf871ee41b230f431925e1f4cc1dd225ef47cc066f",
    	"NodeInfo": "enode://3fe0ff0dd2730eaac7b6b379bdb51215b5831f4f48fa54a24a0298ad5ba8c2a332442948d53f4cd4fd28f373089a35e806ef722eb045659910f96a1278120516@0.0.0.0:30303?discport=0"
    }
    {
    	"Address": "0x7ae555d0f6faad7930434abdaac2274fd86ab516",
    	"Nodekey": "768b87473ba96fcfa272f958fc95a3cefdf9aa82110cde6f2f34aa5855eb39db",
    	"NodeInfo": "enode://e53e92e5a51ac2685b0406d0d3c62288b53831c3b0f492b9dc4bc40334783702cfa74c49b836efa2761edde33a3282704273b2453537b855e7a4aeadcccdb43e@0.0.0.0:30303?discport=0"
    }
    
        
    static-nodes.json
    [
    	"enode://dd333ec28f0a8910c92eb4d336461eea1c20803eed9cf2c056557f986e720f8e693605bba2f4e8f289b1162e5ac7c80c914c7178130711e393ca76abc1d92f57@0.0.0.0:30303?discport=0",
    	"enode://1bb6be462f27e56f901c3fcb2d53a9273565f48e5d354c08f0c044405b29291b405b9f5aa027f3a75f9b058cb43e2f54719f15316979a0e5a2b760fff4631998@0.0.0.0:30303?discport=0",
    	"enode://0df02e94a3befc0683780d898119d3b675e5942c1a2f9ad47d35b4e6ccaf395cd71ec089fcf1d616748bf9871f91e5e3d29c1cf6f8f81de1b279082a104f619d@0.0.0.0:30303?discport=0",
    	"enode://3fe0ff0dd2730eaac7b6b379bdb51215b5831f4f48fa54a24a0298ad5ba8c2a332442948d53f4cd4fd28f373089a35e806ef722eb045659910f96a1278120516@0.0.0.0:30303?discport=0",
    	"enode://e53e92e5a51ac2685b0406d0d3c62288b53831c3b0f492b9dc4bc40334783702cfa74c49b836efa2761edde33a3282704273b2453537b855e7a4aeadcccdb43e@0.0.0.0:30303?discport=0"
    ]
    
    
    
    genesis.json
    {
        "config": {
            "chainId": 10,
            "eip150Block": 1,
            "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
            "eip155Block": 1,
            "eip158Block": 1,
            "byzantiumBlock": 1,
            "istanbul": {
                "epoch": 30000,
                "policy": 0
            },
            "isQuorum": true
        },
        "nonce": "0x0",
        "timestamp": "0x5cffc201",
        "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000f8aff869944c1ccd426833b9782729a212c857f2f03b7b4c0d94189d23d201b03ae1cf9113672df29a5d672aefa39444b07d2c28b8ed8f02b45bd84ac7d9051b3349e694c1056df7c02b6f1a353052eaf0533cc7cb743b52947ae555d0f6faad7930434abdaac2274fd86ab516b8410000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0",
        "gasLimit": "0xe0000000",
        "difficulty": "0x1",
        "mixHash": "0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365",
        "coinbase": "0x0000000000000000000000000000000000000000",
        "alloc": {
            "189d23d201b03ae1cf9113672df29a5d672aefa3": {
                "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
            },
            "44b07d2c28b8ed8f02b45bd84ac7d9051b3349e6": {
                "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
            },
            "4c1ccd426833b9782729a212c857f2f03b7b4c0d": {
                "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
            },
            "7ae555d0f6faad7930434abdaac2274fd86ab516": {
                "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
            },
            "c1056df7c02b6f1a353052eaf0533cc7cb743b52": {
                "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
            }
        },
        "number": "0x0",
        "gasUsed": "0x0",
        "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
    }

    $ ls -l
     total 16
     drwxr-xr-x  9 krish  staff   288 11 Jun 16:00 .
     drwxr-xr-x  8 krish  staff   256 11 Jun 15:58 ..
     drwxr-xr-x  3 krish  staff    96 11 Jun 16:00 0
     drwxr-xr-x  3 krish  staff    96 11 Jun 16:00 1
     drwxr-xr-x  3 krish  staff    96 11 Jun 16:00 2
     drwxr-xr-x  3 krish  staff    96 11 Jun 16:00 3
     drwxr-xr-x  3 krish  staff    96 11 Jun 16:00 4
     -rwxr-xr-x  1 krish  staff  1878 11 Jun 16:00 genesis.json
     -rwxr-xr-x  1 krish  staff   832 11 Jun 16:00 static-nodes.json
    ```
    
5. Update `static-nodes.json` to include the intended IP and port numbers of all initial validator nodes. In `static-nodes.json`, you will see a different row for each node. For the rest of the installation guide, row Y refers to node Y and row 1 is assumed to correspond to the lead node
    ```
    $ cat static-nodes.json
    .... update the IP and port numbers as give below...
    [
    	"enode://dd333ec28f0a8910c92eb4d336461eea1c20803eed9cf2c056557f986e720f8e693605bba2f4e8f289b1162e5ac7c80c914c7178130711e393ca76abc1d92f57@127.0.0.1:30300?discport=0",
    	"enode://1bb6be462f27e56f901c3fcb2d53a9273565f48e5d354c08f0c044405b29291b405b9f5aa027f3a75f9b058cb43e2f54719f15316979a0e5a2b760fff4631998@127.0.0.1:30301?discport=0",
    	"enode://0df02e94a3befc0683780d898119d3b675e5942c1a2f9ad47d35b4e6ccaf395cd71ec089fcf1d616748bf9871f91e5e3d29c1cf6f8f81de1b279082a104f619d@127.0.0.1:30302?discport=0",
    	"enode://3fe0ff0dd2730eaac7b6b379bdb51215b5831f4f48fa54a24a0298ad5ba8c2a332442948d53f4cd4fd28f373089a35e806ef722eb045659910f96a1278120516@127.0.0.1:30303?discport=0",
    	"enode://e53e92e5a51ac2685b0406d0d3c62288b53831c3b0f492b9dc4bc40334783702cfa74c49b836efa2761edde33a3282704273b2453537b855e7a4aeadcccdb43e@127.0.0.1:30304?discport=0"
    ]
    ```
6. In each node's working directory, create a data directory called `data`, and inside `data` create the `geth` directory
    ```
    $ cd ..
    $ mkdir -p node0/data/geth
    $ mkdir -p node1/data/geth
    $ mkdir -p node2/data/geth
    $ mkdir -p node3/data/geth
    $ mkdir -p node4/data/geth
    ```
7. Now we will generate initial accounts for any of the nodes in the required node's working directory. The resulting public account address printed in the terminal should be recorded. Repeat as many times as necessary. A set of funded accounts may be required depending what you are trying to accomplish
    ```
    $ geth --datadir node0/data account new
    INFO [06-11|16:05:53.672] Maximum peer count                       ETH=25 LES=0 total=25
    Your new account is locked with a password. Please give a password. Do not forget this password.
    Passphrase: 
    Repeat passphrase: 
    Address: {8fc817d90f179b0b627c2ecbcc1d1b0fcd13ddbd}
    $ geth --datadir node1/data account new
    INFO [06-11|16:06:34.529] Maximum peer count                       ETH=25 LES=0 total=25
    Your new account is locked with a password. Please give a password. Do not forget this password.
    Passphrase: 
    Repeat passphrase: 
    Address: {dce8adeef16a45d94be5e7804df6d35db834d94a}
    $ geth --datadir node2/data account new
    INFO [06-11|16:06:54.365] Maximum peer count                       ETH=25 LES=0 total=25
    Your new account is locked with a password. Please give a password. Do not forget this password.
    Passphrase: 
    Repeat passphrase: 
    Address: {65a3ab6d4cf23395f544833831fc5d42a0b2b43a}
    ```
8. To add accounts to the initial block, edit the `genesis.json` file in the lead node's working directory and update the `alloc` field with the account(s) that were generated at previous step
    ```
    $ vim node0/genesis.json
    .... update the accounts under 'alloc'
    {
        "config": {
            "chainId": 10,
            "eip150Block": 1,
            "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
            "eip155Block": 1,
            "eip158Block": 1,
            "byzantiumBlock": 1,
            "istanbul": {
                "epoch": 30000,
                "policy": 0
            },
            "isQuorum": true
        },
        "nonce": "0x0",
        "timestamp": "0x5cffc201",
        "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000f8aff869944c1ccd426833b9782729a212c857f2f03b7b4c0d94189d23d201b03ae1cf9113672df29a5d672aefa39444b07d2c28b8ed8f02b45bd84ac7d9051b3349e694c1056df7c02b6f1a353052eaf0533cc7cb743b52947ae555d0f6faad7930434abdaac2274fd86ab516b8410000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0",
        "gasLimit": "0xe0000000",
        "difficulty": "0x1",
        "mixHash": "0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365",
        "coinbase": "0x0000000000000000000000000000000000000000",
        "alloc": {
            "8fc817d90f179b0b627c2ecbcc1d1b0fcd13ddbd": {
                "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
            },
            "dce8adeef16a45d94be5e7804df6d35db834d94a": {
                "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
            },
            "65a3ab6d4cf23395f544833831fc5d42a0b2b43a": {
                "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
            }
        },
        "number": "0x0",
        "gasUsed": "0x0",
        "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
    }
    ```
9. Next we need to distribute the files created in part 4, which currently reside in the lead node's working directory, to all other nodes. To do so, place `genesis.json` in the working directory of all nodes, place `static-nodes.json` in the data folder of each node and place `X/nodekey` in node (X-1)'s `data/geth` directory
    ```
    $ cp node0/genesis.json node1
    $ cp node0/genesis.json node2
    $ cp node0/genesis.json node3
    $ cp node0/genesis.json node4
    $ cp node0/static-nodes.json node0/data/
    $ cp node0/static-nodes.json node1/data/
    $ cp node0/static-nodes.json node2/data/
    $ cp node0/static-nodes.json node3/data/
    $ cp node0/static-nodes.json node4/data/
    $ cp node0/0/nodekey node0/data/geth
    $ cp node0/1/nodekey node1/data/geth
    $ cp node0/2/nodekey node2/data/geth
    $ cp node0/3/nodekey node3/data/geth
    $ cp node0/4/nodekey node4/data/geth
    ```
10. Switch into working directory of lead node and initialize it. Repeat for every working directory X created in step 3. *The resulting hash given by executing `geth init` must match for every node*
    ```
    $ cd node0
    $ geth --datadir data init genesis.json
    INFO [06-11|16:14:11.883] Maximum peer count                       ETH=25 LES=0 total=25
    INFO [06-11|16:14:11.894] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node0/data/geth/chaindata cache=16 handles=16
    INFO [06-11|16:14:11.896] Writing custom genesis block 
    INFO [06-11|16:14:11.897] Persisted trie from memory database      nodes=6 size=1.01kB time=76.665µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:14:11.897] Successfully wrote genesis state         database=chaindata                                                  hash=b992be…533db7
    INFO [06-11|16:14:11.897] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node0/data/geth/lightchaindata cache=16 handles=16
    INFO [06-11|16:14:11.898] Writing custom genesis block 
    INFO [06-11|16:14:11.898] Persisted trie from memory database      nodes=6 size=1.01kB time=54.929µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:14:11.898] Successfully wrote genesis state         database=lightchaindata                                                  hash=b992be…533db7
    $
    $ cd ..
    $ cd node1
    $ geth --datadir data init genesis.json
    INFO [06-11|16:14:24.814] Maximum peer count                       ETH=25 LES=0 total=25
    INFO [06-11|16:14:24.824] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node1/data/geth/chaindata cache=16 handles=16
    INFO [06-11|16:14:24.831] Writing custom genesis block 
    INFO [06-11|16:14:24.831] Persisted trie from memory database      nodes=6 size=1.01kB time=82.799µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:14:24.832] Successfully wrote genesis state         database=chaindata                                                  hash=b992be…533db7
    INFO [06-11|16:14:24.832] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node1/data/geth/lightchaindata cache=16 handles=16
    INFO [06-11|16:14:24.833] Writing custom genesis block 
    INFO [06-11|16:14:24.833] Persisted trie from memory database      nodes=6 size=1.01kB time=52.828µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:14:24.834] Successfully wrote genesis state         database=lightchaindata                                                  hash=b992be…533db7
    $
    $ cd ..
    $ cd node2
    $ geth --datadir data init genesis.json
    INFO [06-11|16:14:35.246] Maximum peer count                       ETH=25 LES=0 total=25
    INFO [06-11|16:14:35.257] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node2/data/geth/chaindata cache=16 handles=16
    INFO [06-11|16:14:35.264] Writing custom genesis block 
    INFO [06-11|16:14:35.265] Persisted trie from memory database      nodes=6 size=1.01kB time=124.91µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:14:35.265] Successfully wrote genesis state         database=chaindata                                                  hash=b992be…533db7
    INFO [06-11|16:14:35.265] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node2/data/geth/lightchaindata cache=16 handles=16
    INFO [06-11|16:14:35.267] Writing custom genesis block 
    INFO [06-11|16:14:35.268] Persisted trie from memory database      nodes=6 size=1.01kB time=85.504µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:14:35.268] Successfully wrote genesis state         database=lightchaindata                                                  hash=b992be…533db7
    $ cd ../node3
    $ geth --datadir data init genesis.json
    INFO [06-11|16:14:42.168] Maximum peer count                       ETH=25 LES=0 total=25
    INFO [06-11|16:14:42.178] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node3/data/geth/chaindata cache=16 handles=16
    INFO [06-11|16:14:42.186] Writing custom genesis block 
    INFO [06-11|16:14:42.186] Persisted trie from memory database      nodes=6 size=1.01kB time=124.611µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:14:42.187] Successfully wrote genesis state         database=chaindata                                                  hash=b992be…533db7
    INFO [06-11|16:14:42.187] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node3/data/geth/lightchaindata cache=16 handles=16
    INFO [06-11|16:14:42.189] Writing custom genesis block 
    INFO [06-11|16:14:42.189] Persisted trie from memory database      nodes=6 size=1.01kB time=80.973µs  gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:14:42.189] Successfully wrote genesis state         database=lightchaindata                                                  hash=b992be…533db7
    $ cd ../node4
    $ geth --datadir data init genesis.json
    INFO [06-11|16:14:48.737] Maximum peer count                       ETH=25 LES=0 total=25
    INFO [06-11|16:14:48.747] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node4/data/geth/chaindata cache=16 handles=16
    INFO [06-11|16:14:48.749] Writing custom genesis block 
    INFO [06-11|16:14:48.749] Persisted trie from memory database      nodes=6 size=1.01kB time=71.213µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:14:48.750] Successfully wrote genesis state         database=chaindata                                                  hash=b992be…533db7
    INFO [06-11|16:14:48.750] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node4/data/geth/lightchaindata cache=16 handles=16
    INFO [06-11|16:14:48.751] Writing custom genesis block 
    INFO [06-11|16:14:48.751] Persisted trie from memory database      nodes=6 size=1.01kB time=53.773µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:14:48.751] Successfully wrote genesis state         database=lightchaindata                                                  hash=b992be…533db7
    $ cd..
    ```

11. Start all nodes by first creating a script and running it.
    ```
    $ vim startall.sh
    .... paste below. The port numbers should match the port number for each node decided on in step 5
    #!/bin/bash
    cd node0
    PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30300 2>>node.log &
    
    
    cd ../node1
    PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22001 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30301 2>>node.log &
    
    cd ../node2
    PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22002 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30302 2>>node.log &
     
    cd ../node3
    PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22003 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30303 2>>node.log &
    
    cd ../node4
    PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22004 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30304 2>>node.log &
    $
    See if the any geth nodes are running.
    $ ps | grep geth
    Kill geth processes
    $ killall -INT geth
    $
    $ chmod +x startall.sh
    $ ./startall.sh
    $ ps
      PID TTY           TIME CMD
    10554 ttys000    0:00.11 -bash
    21829 ttys001    0:00.03 -bash
     9125 ttys002    0:00.82 -bash
    36432 ttys002    0:00.19 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,
    36433 ttys002    0:00.18 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22001 --rpcapi admin,
    36434 ttys002    0:00.19 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22002 --rpcapi admin,
    36435 ttys002    0:00.19 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22003 --rpcapi admin,
    36436 ttys002    0:00.19 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22004 --rpcapi admin,
    $ 
    ```
    
    !!! note
        This configuration starts Quorum without privacy support as could be evidenced in prefix `PRIVATE_CONFIG=ignore`, please see below sections on [how to enable privacy with privacy transaction managers](#adding-privacy-transaction-manager).
        Please note that istanbul-tools may be used to generate X number of nodes, more information is available in the [docs](https://github.com/jpmorganchase/istanbul-tools).

    Your node is now operational and you may attach to it with `geth attach node0/data/geth.ipc`. 
    
### Adding additional validator
1. Create a working directory for the new node that needs to be added
    ```
    $ mkdir node5
    ```
2. Change into the working directory for the new node and run `istanbul setup --num 1 --verbose --quorum --save`. This will generate the validator details including Address, NodeInfo and genesis.json
    ```
    $ cd node5
    $ ../istanbul-tools/build/bin/istanbul setup --num 1 --verbose --quorum --save
    validators
    {
    	"Address": "0x2aabbc1bb9bacef60a09764d1a1f4f04a47885c1",
    	"Nodekey": "25b47a49ef08f888c04f30417363e6c6bc33e739147b2f8b5377b3168f9f7435",
    	"NodeInfo": "enode://273eaf48591ce0e77c800b3e6465811d6d2f924c4dcaae016c2c7375256d17876c3e05f91839b741fe12350da0b5a741da4e30f39553fe8790f88503c64f6ef9@0.0.0.0:30303?discport=0"
    }
    
    
    
    genesis.json
    {
        "config": {
            "chainId": 10,
            "eip150Block": 1,
            "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
            "eip155Block": 1,
            "eip158Block": 1,
            "byzantiumBlock": 1,
            "istanbul": {
                "epoch": 30000,
                "policy": 0
            },
            "isQuorum": true
        },
        "nonce": "0x0",
        "timestamp": "0x5cffc942",
        "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000f85ad5942aabbc1bb9bacef60a09764d1a1f4f04a47885c1b8410000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0",
        "gasLimit": "0xe0000000",
        "difficulty": "0x1",
        "mixHash": "0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365",
        "coinbase": "0x0000000000000000000000000000000000000000",
        "alloc": {
            "2aabbc1bb9bacef60a09764d1a1f4f04a47885c1": {
                "balance": "0x446c3b15f9926687d2c40534fdb564000000000000"
            }
        },
        "number": "0x0",
        "gasUsed": "0x0",
        "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
    }
    ```
    
3. Copy the address of the validator and run `istanbul.propose(<address>, true)` from more than half the number of current validators.
    ```
    $ cd ..
    $ geth attach node0/data/geth.ipc
    $ geth attach node0/data/geth.ipc
    Welcome to the Geth JavaScript console!
    
    instance: Geth/v1.8.18-stable-bb88608c(quorum-v2.2.3)/darwin-amd64/go1.10.2
    coinbase: 0x4c1ccd426833b9782729a212c857f2f03b7b4c0d
    at block: 137 (Tue, 11 Jun 2019 16:32:47 BST)
     datadir: /Users/krish/fromscratchistanbul/node0/data
     modules: admin:1.0 debug:1.0 eth:1.0 istanbul:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
    
    > istanbul.getValidators()
    ["0x189d23d201b03ae1cf9113672df29a5d672aefa3", "0x44b07d2c28b8ed8f02b45bd84ac7d9051b3349e6", "0x4c1ccd426833b9782729a212c857f2f03b7b4c0d", "0x7ae555d0f6faad7930434abdaac2274fd86ab516", "0xc1056df7c02b6f1a353052eaf0533cc7cb743b52"]
    > istanbul.propose("0x2aabbc1bb9bacef60a09764d1a1f4f04a47885c1",true)
    null
    $
    $
    $ geth attach node1/data/geth.ipc
    Welcome to the Geth JavaScript console!
    
    instance: Geth/v1.8.18-stable-bb88608c(quorum-v2.2.3)/darwin-amd64/go1.10.2
    coinbase: 0x189d23d201b03ae1cf9113672df29a5d672aefa3
    at block: 176 (Tue, 11 Jun 2019 16:36:02 BST)
     datadir: /Users/krish/fromscratchistanbul/node1/data
     modules: admin:1.0 debug:1.0 eth:1.0 istanbul:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
    
    > istanbul.propose("0x2aabbc1bb9bacef60a09764d1a1f4f04a47885c1",true)
    null
    $
    $
    $ geth attach node2/data/geth.ipc
    Welcome to the Geth JavaScript console!
    
    instance: Geth/v1.8.18-stable-bb88608c(quorum-v2.2.3)/darwin-amd64/go1.10.2
    coinbase: 0x44b07d2c28b8ed8f02b45bd84ac7d9051b3349e6
    at block: 179 (Tue, 11 Jun 2019 16:36:17 BST)
     datadir: /Users/krish/fromscratchistanbul/node2/data
     modules: admin:1.0 debug:1.0 eth:1.0 istanbul:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
    
    > istanbul.propose("0x2aabbc1bb9bacef60a09764d1a1f4f04a47885c1",true)
    null
    $
    $
    $ geth attach node3/data/geth.ipc
    Welcome to the Geth JavaScript console!
    
    instance: Geth/v1.8.18-stable-bb88608c(quorum-v2.2.3)/darwin-amd64/go1.10.2
    coinbase: 0xc1056df7c02b6f1a353052eaf0533cc7cb743b52
    at block: 181 (Tue, 11 Jun 2019 16:36:27 BST)
     datadir: /Users/krish/fromscratchistanbul/node3/data
     modules: admin:1.0 debug:1.0 eth:1.0 istanbul:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
    
    > istanbul.propose("0x2aabbc1bb9bacef60a09764d1a1f4f04a47885c1",true)
    ```
    
4. Verify that the new validator has been added to the list of validators by running `istanbul.getValidators()`
    ```
    ... you can see below command now displays 6 node address as validators.
    > istanbul.getValidators()
    ["0x189d23d201b03ae1cf9113672df29a5d672aefa3", "0x2aabbc1bb9bacef60a09764d1a1f4f04a47885c1", "0x44b07d2c28b8ed8f02b45bd84ac7d9051b3349e6", "0x4c1ccd426833b9782729a212c857f2f03b7b4c0d", "0x7ae555d0f6faad7930434abdaac2274fd86ab516", "0xc1056df7c02b6f1a353052eaf0533cc7cb743b52"]    
    ```
5. Copy `static-nodes.json` and genesis.json from the existing chain. `static-nodes.json` should be placed into new nodes data dir
    ```
    $ cd node5
    $ mkdir -p data/geth
    $ cp ../node0/static-nodes.json data
    $ cp ../node0/genesis.json .
    ```
6. Edit `static-nodes.json` and add the new validators node info to the end of the file. New validators node info can be got from the output of `istanbul setup --num 1 --verbose --quorum --save` command that was run in step 2. Update the IP address and port of the node info to match the IP address of the validator and port you want to use.
    ```
    $ vim data/static-nodes.json
    ...add new validate nodes details with correct IP and port details
    [
    	"enode://dd333ec28f0a8910c92eb4d336461eea1c20803eed9cf2c056557f986e720f8e693605bba2f4e8f289b1162e5ac7c80c914c7178130711e393ca76abc1d92f57@127.0.0.1:30300?discport=0",
    	"enode://1bb6be462f27e56f901c3fcb2d53a9273565f48e5d354c08f0c044405b29291b405b9f5aa027f3a75f9b058cb43e2f54719f15316979a0e5a2b760fff4631998@127.0.0.1:30301?discport=0",
    	"enode://0df02e94a3befc0683780d898119d3b675e5942c1a2f9ad47d35b4e6ccaf395cd71ec089fcf1d616748bf9871f91e5e3d29c1cf6f8f81de1b279082a104f619d@127.0.0.1:30302?discport=0",
    	"enode://3fe0ff0dd2730eaac7b6b379bdb51215b5831f4f48fa54a24a0298ad5ba8c2a332442948d53f4cd4fd28f373089a35e806ef722eb045659910f96a1278120516@127.0.0.1:30303?discport=0",
    	"enode://e53e92e5a51ac2685b0406d0d3c62288b53831c3b0f492b9dc4bc40334783702cfa74c49b836efa2761edde33a3282704273b2453537b855e7a4aeadcccdb43e@127.0.0.1:30304?discport=0",
        "enode://273eaf48591ce0e77c800b3e6465811d6d2f924c4dcaae016c2c7375256d17876c3e05f91839b741fe12350da0b5a741da4e30f39553fe8790f88503c64f6ef9@127.0.0.1:30305?discport=0"
    ]

    ```
7. Copy the nodekey that was generated by `istanbul setup` command to the `geth` directory inside the working directory
    ```
    $ cp 0/nodekey data/geth
    ```
8. Generate one or more accounts for this node and take down the account address.
    ```
    $ geth --datadir data account new
    INFO [06-12|17:45:11.116] Maximum peer count                       ETH=25 LES=0 total=25
    Your new account is locked with a password. Please give a password. Do not forget this password.
    Passphrase: 
    Repeat passphrase: 
    Address: {37922bce824bca2f3206ea53dd50d173b368b572}
    ```
9. Initialize new node with below command
    ```
    $ geth --datadir data init genesis.json
    INFO [06-11|16:42:27.120] Maximum peer count                       ETH=25 LES=0 total=25
    INFO [06-11|16:42:27.130] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node5/data/geth/chaindata cache=16 handles=16
    INFO [06-11|16:42:27.138] Writing custom genesis block 
    INFO [06-11|16:42:27.138] Persisted trie from memory database      nodes=6 size=1.01kB time=163.024µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:42:27.139] Successfully wrote genesis state         database=chaindata                                                  hash=b992be…533db7
    INFO [06-11|16:42:27.139] Allocated cache and file handles         database=/Users/krish/fromscratchistanbul/node5/data/geth/lightchaindata cache=16 handles=16
    INFO [06-11|16:42:27.141] Writing custom genesis block 
    INFO [06-11|16:42:27.142] Persisted trie from memory database      nodes=6 size=1.01kB time=94.57µs   gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-11|16:42:27.142] Successfully wrote genesis state         database=lightchaindata                                                  hash=b992be…533db7
    $
    
    ```
10. Start the node by first creating below script and executing it:
    ```
    $ cd ..
    $ cp startall.sh start6.sh
    $ vim start6.sh
    ... paste below and update IP and port number matching for this node decided on step 6
    #!/bin/bash
    cd node5
    PRIVATE_CONFIG=ignore nohup geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22005 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port 30305 2>>node.log &
    $
    $ ./start6.sh 
    $
    $ ps
      PID TTY           TIME CMD
    10554 ttys000    0:00.11 -bash
    21829 ttys001    0:00.03 -bash
     9125 ttys002    0:00.93 -bash
    36432 ttys002    0:24.48 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,
    36433 ttys002    0:23.36 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22001 --rpcapi admin,
    36434 ttys002    0:24.32 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22002 --rpcapi admin,
    36435 ttys002    0:24.21 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22003 --rpcapi admin,
    36436 ttys002    0:24.17 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22004 --rpcapi admin,
    36485 ttys002    0:00.15 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22005 --rpcapi admin,
    36455 ttys003    0:00.04 -bash
    36467 ttys003    0:00.32 geth attach node3/data/geth.ipc
    ```

### Removing validator
1. Attach to a running validator and run `istanbul.getValidators()` and identify the address of the validator that needs to be removed
    ```
    $ geth attach node0/data/geth.ipc
    Welcome to the Geth JavaScript console!
    
    instance: Geth/v1.8.18-stable-bb88608c(quorum-v2.2.3)/darwin-amd64/go1.10.2
    coinbase: 0xc1056df7c02b6f1a353052eaf0533cc7cb743b52
    at block: 181 (Tue, 11 Jun 2019 16:36:27 BST)
     datadir: /Users/krish/fromscratchistanbul/node0/data
     modules: admin:1.0 debug:1.0 eth:1.0 istanbul:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
    > istanbul.getValidators()
    ["0x189d23d201b03ae1cf9113672df29a5d672aefa3", "0x2aabbc1bb9bacef60a09764d1a1f4f04a47885c1", "0x44b07d2c28b8ed8f02b45bd84ac7d9051b3349e6", "0x4c1ccd426833b9782729a212c857f2f03b7b4c0d", "0x7ae555d0f6faad7930434abdaac2274fd86ab516", "0xc1056df7c02b6f1a353052eaf0533cc7cb743b52"]
    ```
2. Run `istanbul.propose(<address>, false)` by passing the address of the validator that needs to be removed from more than half current validators
    ```
    > istanbul.propose("0x2aabbc1bb9bacef60a09764d1a1f4f04a47885c1",false)
    null
    $
    $ geth attach node1/data/geth.ipc
    Welcome to the Geth JavaScript console!
        
    instance: Geth/v1.8.18-stable-bb88608c(quorum-v2.2.3)/darwin-amd64/go1.10.2
    coinbase: 0xc1056df7c02b6f1a353052eaf0533cc7cb743b52
    at block: 181 (Tue, 11 Jun 2019 16:36:27 BST)
    datadir: /Users/krish/fromscratchistanbul/node1/data
    modules: admin:1.0 debug:1.0 eth:1.0 istanbul:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
    > istanbul.propose("0x2aabbc1bb9bacef60a09764d1a1f4f04a47885c1",false)
    null
    $
    $ geth attach node2/data/geth.ipc
    Welcome to the Geth JavaScript console!
        
    instance: Geth/v1.8.18-stable-bb88608c(quorum-v2.2.3)/darwin-amd64/go1.10.2
    coinbase: 0xc1056df7c02b6f1a353052eaf0533cc7cb743b52
    at block: 181 (Tue, 11 Jun 2019 16:36:27 BST)
    datadir: /Users/krish/fromscratchistanbul/node2/data
    modules: admin:1.0 debug:1.0 eth:1.0 istanbul:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
    > istanbul.propose("0x2aabbc1bb9bacef60a09764d1a1f4f04a47885c1",false)
    null 
    $
    $ geth attach node3/data/geth.ipc
    Welcome to the Geth JavaScript console!
        
    instance: Geth/v1.8.18-stable-bb88608c(quorum-v2.2.3)/darwin-amd64/go1.10.2
    coinbase: 0xc1056df7c02b6f1a353052eaf0533cc7cb743b52
    at block: 181 (Tue, 11 Jun 2019 16:36:27 BST)
    datadir: /Users/krish/fromscratchistanbul/node3/data
    modules: admin:1.0 debug:1.0 eth:1.0 istanbul:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
    > istanbul.propose("0x2aabbc1bb9bacef60a09764d1a1f4f04a47885c1",false)
    null       
    ```
3. Verify that the validator has been removed by running `istanbul.getValidators()`
    ```
    > istanbul.getValidators()
    ["0x189d23d201b03ae1cf9113672df29a5d672aefa3", "0x44b07d2c28b8ed8f02b45bd84ac7d9051b3349e6", "0x4c1ccd426833b9782729a212c857f2f03b7b4c0d", "0x7ae555d0f6faad7930434abdaac2274fd86ab516", "0xc1056df7c02b6f1a353052eaf0533cc7cb743b52"]
    ```
4. Stop the `geth` process corresponding to the validator that was removed.
    ```
    $ ps
      PID TTY           TIME CMD
    10554 ttys000    0:00.11 -bash
    21829 ttys001    0:00.03 -bash
     9125 ttys002    0:00.94 -bash
    36432 ttys002    0:31.93 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,
    36433 ttys002    0:30.75 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22001 --rpcapi admin,
    36434 ttys002    0:31.72 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22002 --rpcapi admin,
    36435 ttys002    0:31.65 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22003 --rpcapi admin,
    36436 ttys002    0:31.63 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22004 --rpcapi admin,
    36485 ttys002    0:06.86 geth --datadir data --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport 22005 --rpcapi admin,
    36455 ttys003    0:00.05 -bash
    36493 ttys003    0:00.22 geth attach node4/data/geth.ipc
    $ kill 36485
    ```

### Adding non-validator node

Same instructions as adding validator node **excluding** step 3 which proposes the node as validator.

### Removing non-validator node

Just execute **step 4** instruction from removing a validator node.


## Adding privacy transaction manager
### Tessera
1. Build Quorum and install [Tessera](https://github.com/jpmorganchase/tessera/releases) as described in the [Installing](../Installing) section. Ensure that PATH contains geth and bootnode. Be aware of the location of the `tessera.jar` release file
    ```
    $ git clone https://github.com/jpmorganchase/quorum.git
    $ cd quorum
    $ make all
    $ export PATH=$(pwd)/build/bin:$PATH
    $ cd ..
    .... copy tessera jar to your desired destination and rename it as tessera
    $ mv tessera-app-0.9.2-app.jar tessera.jar
    
    ```
2. Generate new keys using `java -jar /path-to-tessera/tessera.jar -keygen -filename new-node-1`
    ```
    $ mkdir new-node-1t
    $ cd new-node-1t
    $ java -jar ../tessera.jar -keygen -filename new-node-1
    Enter a password if you want to lock the private key or leave blank
    
    Please re-enter the password (or lack of) to confirm
    
    10:32:51.256 [main] INFO  com.quorum.tessera.nacl.jnacl.Jnacl - Generating new keypair...
    10:32:51.279 [main] INFO  com.quorum.tessera.nacl.jnacl.Jnacl - Generated public key PublicKey[pnesVeDgs805ZPbnulzC5wokDzpdN7CeYKVUBXup/W4=] and private key REDACTED
    10:32:51.624 [main] INFO  c.q.t.k.generation.FileKeyGenerator - Saved public key to /Users/krish/fromscratch/new-node-1t/new-node-1.pub
    10:32:51.624 [main] INFO  c.q.t.k.generation.FileKeyGenerator - Saved private key to /Users/krish/fromscratch/new-node-1t/new-node-1.key
    ```
3. Create new configuration file with newly generated keys referenced. Name it `config.json` as done in this example
    ```
    vim config.json
    {
       "useWhiteList": false,
       "jdbc": {
           "username": "sa",
           "password": "",
           "url": "jdbc:h2:/yourpath/new-node-1t/db1;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0",
           "autoCreateTables": true
       },
       "serverConfigs":[
           {
               "app":"ThirdParty",
               "enabled": true,
               "serverAddress": "http://localhost:9081",
               "communicationType" : "REST"
           },
           {
               "app":"Q2T",
               "enabled": true,
                "serverAddress":"unix:/yourpath/new-node-1t/tm.ipc",
               "communicationType" : "REST"
           },
           {
               "app":"P2P",
               "enabled": true,
               "serverAddress":"http://localhost:9001",
               "sslConfig": {
                   "tls": "OFF"
               },
               "communicationType" : "REST"
           }
       ],
       "peer": [
           {
               "url": "http://localhost:9001"
           },
           {
               "url": "http://localhost:9003"
           }
       ],
       "keys": {
           "passwords": [],
           "keyData": [
               {
                   "privateKeyPath": "/yourpath/new-node-1t/new-node-1.key",
                   "publicKeyPath": "/yourpath/new-node-1t/new-node-1.pub"
               }
           ]
       },
       "alwaysSendTo": []
    }
    ```
    
4. If you want to start another Tessera node, please repeat step 2 & step 3
    ```
    $ cd ..
    $ mkdir new-node-2t
    $ cd new-node-2t
    $ java -jar ../tessera.jar -keygen -filename new-node-2
    Enter a password if you want to lock the private key or leave blank
    
    Please re-enter the password (or lack of) to confirm
    
    10:45:02.567 [main] INFO  com.quorum.tessera.nacl.jnacl.Jnacl - Generating new keypair...
    10:45:02.585 [main] INFO  com.quorum.tessera.nacl.jnacl.Jnacl - Generated public key PublicKey[AeggpVlVsi+rxD6h9tcq/8qL/MsjyipUnkj1nvNPgTU=] and private key REDACTED
    10:45:02.926 [main] INFO  c.q.t.k.generation.FileKeyGenerator - Saved public key to /Users/krish/fromscratch/new-node-2t/new-node-2.pub
    10:45:02.926 [main] INFO  c.q.t.k.generation.FileKeyGenerator - Saved private key to /Users/krish/fromscratch/new-node-2t/new-node-2.key
    $
    $ cp ../new-node-1t/config.json .
    $ vim config.json
    .... paste below
    {
       "useWhiteList": false,
       "jdbc": {
           "username": "sa",
           "password": "",
           "url": "jdbc:h2:yourpath/new-node-2t/db1;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0",
           "autoCreateTables": true
       },
       "serverConfigs":[
           {
               "app":"ThirdParty",
               "enabled": true,
               "serverAddress": "http://localhost:9083",
               "communicationType" : "REST"
           },
           {
               "app":"Q2T",
               "enabled": true,
                "serverAddress":"unix:/yourpath/new-node-2t/tm.ipc",
               "communicationType" : "REST"
           },
           {
               "app":"P2P",
               "enabled": true,
               "serverAddress":"http://localhost:9003",
               "sslConfig": {
                   "tls": "OFF"
               },
               "communicationType" : "REST"
           }
       ],
       "peer": [
           {
               "url": "http://localhost:9001"
           },
           {
               "url": "http://localhost:9003"
           }
       ],
       "keys": {
           "passwords": [],
           "keyData": [
               {
                   "privateKeyPath": "/yourpath/new-node-2t/new-node-2.key",
                   "publicKeyPath": "/yourpath/new-node-2t/new-node-2.pub"
               }
           ]
       },
       "alwaysSendTo": []
    }
    
    ```
5. Start your Tessera nodes and send then into background
    ```
    $ java -jar ../tessera.jar -configfile config.json >> tessera.log 2>&1 &
    [1] 38064
    $ cd ../new-node-1t
    $ java -jar ../tessera.jar -configfile config.json >> tessera.log 2>&1 &
    [2] 38076
    $ ps
      PID TTY           TIME CMD
    10554 ttys000    0:00.12 -bash
    38234 ttys000    1:15.31 /usr/local/Cellar/python/3.7.3/Frameworks/Python.framework/Versions/3.7/Resources/Python.app/Contents/MacOS/Python /usr/local/bin/mkdocs serve
    21829 ttys001    0:00.18 -bash
     9125 ttys002    0:01.52 -bash
    38072 ttys002    1:18.42 /usr/bin/java -jar ../tessera.jar -configfile config.json
    38076 ttys002    1:15.86 /usr/bin/java -jar ../tessera.jar -configfile config.json
    ```

6. Start Quorum nodes attached to running Tessera nodes from above and send it to background
    ```
    ... update the start scripts to include PRIVATE_CONFIG
    $cd ..
    $vim startnode1.sh
    ... paste below
    #!/bin/bash
    PRIVATE_CONFIG=/yourpath/new-node-1t/tm.ipc nohup geth --datadir new-node-1 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50000 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft --emitcheckpoints --port 21000 >> node.log 2>&1 &
    $vim startnode2.sh
    ... paste below
    #!/bin/bash
    PRIVATE_CONFIG=/yourpath/new-node-2t/tm.ipc nohup geth --datadir new-node-2 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50001 --raftjoinexisting 2 --rpc --rpcaddr 0.0.0.0 --rpcport 22001 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft --emitcheckpoints --port 21001 2>>node2.log &
    $ ./startnode1.sh
    $ ./startnode2.sh
    $ ps
      PID TTY           TIME CMD
    10554 ttys000    0:00.12 -bash
    21829 ttys001    0:00.18 -bash
     9125 ttys002    0:01.49 -bash
    38072 ttys002    0:48.92 /usr/bin/java -jar ../tessera.jar -configfile config.json
    38076 ttys002    0:47.60 /usr/bin/java -jar ../tessera.jar -configfile config.json
    38183 ttys002    0:08.15 geth --datadir new-node-1 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50000 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,db,eth,debug,miner,net,shh,txpoo
    38204 ttys002    0:00.19 geth --datadir new-node-2 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50001 --raftjoinexisting 2 --rpc --rpcaddr 0.0.0.0 --rpcport 22001 --rpcapi admin,db,eth,debu
    36455 ttys003    0:00.15 -bash  
    ```

    !!! note
        Tessera IPC bridge will be over a file name defined in your `config.json`, usually named `tm.ipc` as evidenced in prefix `PRIVATE_CONFIG=tm.ipc`. Your node is now able to send and receive private transactions, advertised public node key will be in the `new-node-1.pub` file. Tessera offers a lot of configuration flexibility, please refer [Configuration](../../Privacy/Tessera/Configuration/Configuration%20Overview) section under Tessera for complete and up to date configuration options.
        
    Your node is now operational and you may attach to it with `geth attach new-node-1/geth.ipc` to send private transactions. 
    ```
    $ vim private-contract.js
    ... create simple private contract to send transaction from new-node-1 private for new-node-2's tessera public key created in step 4
    a = eth.accounts[0]
    web3.eth.defaultAccount = a;
    
    // abi and bytecode generated from simplestorage.sol:
    // > solcjs --bin --abi simplestorage.sol
    var abi = [{"constant":true,"inputs":[],"name":"storedData","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"uint256"}],"name":"set","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[{"name":"initVal","type":"uint256"}],"payable":false,"type":"constructor"}];
    
    var bytecode = "0x6060604052341561000f57600080fd5b604051602080610149833981016040528080519060200190919050505b806000819055505b505b610104806100456000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a1afcd914605157806360fe47b11460775780636d4ce63c146097575b600080fd5b3415605b57600080fd5b606160bd565b6040518082815260200191505060405180910390f35b3415608157600080fd5b6095600480803590602001909190505060c3565b005b341560a157600080fd5b60a760ce565b6040518082815260200191505060405180910390f35b60005481565b806000819055505b50565b6000805490505b905600a165627a7a72305820d5851baab720bba574474de3d09dbeaabc674a15f4dd93b974908476542c23f00029";
    
    var simpleContract = web3.eth.contract(abi);
    var simple = simpleContract.new(42, {from:web3.eth.accounts[0], data: bytecode, gas: 0x47b760, privateFor: ["AeggpVlVsi+rxD6h9tcq/8qL/MsjyipUnkj1nvNPgTU="]}, function(e, contract) {
    	if (e) {
    		console.log("err creating contract", e);
    	} else {
    		if (!contract.address) {
    			console.log("Contract transaction send: TransactionHash: " + contract.transactionHash + " waiting to be mined...");
    		} else {
    			console.log("Contract mined! Address: " + contract.address);
    			console.log(contract);
    		}
    	}
    });
    $
    $
    ```
    
    !!! note
        Account opened in geth by default are locked, so please unlock account first before sending transaction
        
    ```
    $
    $ geth attach new-node-1/geth.ipc
    > eth.accounts
    ["0x23214cd88f46865207fa1d2a69971a37cdbf526a"]
    > personal.unlockAccount("0x23214cd88f46865207fa1d2a69971a37cdbf526a");
    Unlock account 0x23214cd88f46865207fa1d2a69971a37cdbf526a
    Passphrase: 
    true
    > loadScript("private-contract.js")
    Contract transaction send: TransactionHash: 0x7d3bf7612ef10c71f752e881648b7c8c4eee3223acab151ee0652447790836a6 waiting to be mined...
    true
    > Contract mined! Address: 0xe975e1e11c5268b1efcbf39b7ee3cf7b8dc85fd7
    ```

    You have **successfully** sent a private transaction from node 1 to node 2 !!
    
    !!! note
        if you do not have an valid public key in the array in private-contract.js, you will see the following error when the script in loaded.
    ```
    > loadScript("private-contract.js")
    err creating contract Error: Non-200 status code: &{Status:400 Bad Request StatusCode:400 Proto:HTTP/1.1 ProtoMajor:1      ProtoMinor:1 Header:map[Date:[Mon, 17 Jun 2019 15:23:53 GMT] Content-Type:[text/plain] Content-Length:[73] Server:[Jetty(9.4.z-SNAPSHOT)]] Body:0xc01997a580 ContentLength:73 TransferEncoding:[] Close:false Uncompressed:false Trailer:map[] Request:0xc019788200 TLS:<nil>}
    ```    

### Constellation
1. Build Quorum and install [Constellation](https://github.com/jpmorganchase/constellation/releases) as described in the [Installing](../Installing) section. Ensure that PATH contains geth, bootnode, and constellation-node binaries
2. Generate new keys with `constellation-node --generatekeys=new-node-1`
3. Start your constellation node and send it into background with `constellation-node --url=https://127.0.0.1:9001/ --port=9001 --workdir=. --socket=tm.ipc --publickeys=new-node-1.pub --privatekeys=new-node-1.key --othernodes=https://127.0.0.1:9001/ >> constellation.log 2>&1 &`
4. Start your node and send it into background with `PRIVATE_CONFIG=tm.ipc nohup geth --datadir new-node-1 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50000 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft --emitcheckpoints --port 21000 2>>node.log &`

    !!! note
        Constellation IPC bridge will be over a file name defined in your configuration: in above step #3 see option `--socket=file-name.ipc`. Your node is now able to send and receive private transactions, advertised public node key will be in the `new-node-1.pub` file.

    Your node is now operational and you may attach to it with `geth attach new-node-1/geth.ipc`. 


## Enabling permissioned configuration
Quorum ships with a permissions system based on a custom whitelist. Detailed documentation is available in [Network Permissioning](../../Security/Framework/Quorum%20Network%20Security/Nodes/Permissioning/Network%20Permissioning).
