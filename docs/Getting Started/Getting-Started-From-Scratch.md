# Getting started from scratch
## Quorum with Raft consensus
1. Build Quorum as described in the [getting set up](../Setup%20Overview%20%26%20Quickstart) section. Ensure that PATH contains geth and bootnode
2. Create a working directory which will be the base for the new node(s) and change into it
3. Generate one or more accounts for this node using `geth --datadir new-node-1 account new` and take down the account address. A funded account may be required depending what you are trying to accomplish
4. Create a `genesis.json` file see example [here](../genesis). The `alloc` field should be pre-populated with the account you generated at previous step
5. Generate node key `bootnode --genkey=nodekey` and copy it into datadir
6. Execute `bootnode --nodekey=new-node-1/nodekey --writeaddress` and take note of the displayed output. This is the enode id of the new node
7. Create a file called `static-nodes.json` and edit it to match this [example](../permissioned-nodes). Your file should contain a single line for your node with your enode's id and the ports you are going to use for devp2p and raft. Ensure that this file is in your nodes data directory
8. Initialize new node with `geth --datadir new-node-1 init genesis.json`
9. Start your node and send into background with `PRIVATE_CONFIG=ignore nohup geth --datadir new-node-1 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50000 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft --emitcheckpoints --port 21000 2>>node.log &`

Your node is now operational and you may attach to it with `geth attach new-node-1/geth.ipc`. This configuration starts Quorum without privacy support as could be evidenced in prefix `PRIVATE_CONFIG=ignore`, please see below sections on [how to enable privacy with privacy transaction managers](../Getting-Started-From-Scratch#adding-privacy-transaction-manager).

### Adding additional node
1. Complete steps 1, 2, 5, and 6 from the previous guide
2. Retrieve current chains `genesis.json` and `static-nodes.json`. `static-nodes.json` should be placed into new nodes data dir
3. Initialize new node with `geth --datadir new-node-2 init genesis.json`
4. Edit `static-nodes.json` and add new entry for the new node you are configuring (should be last)
5. Start your node and send into background with `PRIVATE_CONFIG=ignore nohup geth --datadir new-node-2 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50005 --rpc --rpcaddr 0.0.0.0 --rpcport 22005 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft --emitcheckpoints --port 21005 2>>node.log &`
6. Connect to an already running node of the chain and execute `raft.addPeer('enode://new-nodes-enode-address-from-step-6-of-the-above@127.0.0.1:21005?discport=0&raftport=50005')`
7. Optional: share new `static-nodes.json` with all other chain participants

Your additional node is now operational and is part of the same chain as the previously set up node.



## Quorum with Istanbul BFT consensus
1. Build Quorum as described in the [getting set up](../Setup%20Overview%20%26%20Quickstart) section. Ensure that PATH contains geth and bootnode
2. Install [istanbul-tools](https://github.com/jpmorganchase/istanbul-tools) and place `istanbul` binary into PATH
3. Create a working directory for each of the X number of initial validator nodes
4. Change into the lead (whichever one you consider first) node's working directory and generate the setup files for X initial validator nodes by executing `istanbul setup --num X --nodes --quorum --save --verbose` **only execute this instruction once, i.e. not X times**. This command will generate several items of interest: `static-nodes.json`, `genesis.json`, and nodekeys for all the initial validator nodes which will sit in numbered directories from 0 to X-1
5. Update `static-nodes.json` to include the intended IP and port numbers of all initial validator nodes. In `static-nodes.json`, you will see a different row for each node. For the rest of the installation guide, row Y refers to node Y and row 1 is assumed to correspond to the lead node
6. In each node's working directory, create a data directory called `data`, and inside `data` create the `geth` directory
7. Now we will generate initial accounts for any of the nodes by executing `geth --datadir data account new` in the required node's working directory. The resulting public account address printed in the terminal should be recorded. Repeat as many times as necessary. A set of funded accounts may be required depending what you are trying to accomplish
8. To add accounts to the initial block, edit the `genesis.json` file in the lead node's working directory and update the `alloc` field with the account(s) that were generated at previous step
9. Next we need to distribute the files created in part 4, which currently reside in the lead node's working directory, to all other nodes. To do so, place `genesis.json` in the working directory of all nodes, place `static-nodes.json` in the data folder of each node and place `X/nodekey` in node (X-1)'s `data/geth` directory
10. Switch into working directory of lead node and initialize it with `geth --datadir data init genesis.json`. Repeat for every working directory X created in step 3. *The resulting hash given by executing `geth init` must match for every node*
11. Start all nodes and send into background with `PRIVATE_CONFIG=ignore nohup geth --datadir data --permissioned --nodiscover --istanbul.blockperiod 5 --syncmode full --mine --minerthreads 1 --verbosity 5 --networkid 10 --rpc --rpcaddr 0.0.0.0 --rpcport YOUR_NODES_RPC_PORT_NUMBER --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --emitcheckpoints --port YOUR_NODES_PORT_NUMBER 2>>node.log &`, remember to replace `YOUR_NODES_RPC_PORT_NUMBER` and `YOUR_NODES_PORT_NUMBER` with your node's designated port numbers. `YOUR_NODES_PORT_NUMBER` must match the port number for this node decided on in part 5

Your node is now operational and you may attach to it with `geth attach data/geth.ipc`. This configuration starts Quorum without privacy support as could be evidenced in prefix `PRIVATE_CONFIG=ignore`, please see below sections on [[how to enable privacy with privacy transaction managers|From-Scratch#adding-privacy-transaction-manager]].

Please note that istanbul-tools may be used to generate X number of nodes, more information is available in the [docs](https://github.com/jpmorganchase/istanbul-tools).


## Adding privacy transaction manager
### Tessera
1. Build Quorum and install [Tessera](https://github.com/jpmorganchase/tessera/releases) as described in the [getting set up](../Setup%20Overview%20%26%20Quickstart) section. Ensure that PATH contains geth and bootnode. Be aware of the location of the `tessera.jar` release file
2. Generate new keys using `java -jar /path-to-tessera/tessera.jar -keygen -filename new-node-1`
3. Create new configuration file referencing samples [here](../../Privacy/Tessera/Configuration/Sample%20Configuration) with newly generated keys referenced. Note the name of the file or name it `config.json` as done in this example
4. Start your tessera node and send it into background with `java -jar /path-to-tessera/tessera.jar -configfile config.json >> tessera.log 2>&1 &`
5. Start your node and send it into background with `PRIVATE_CONFIG=tm.ipc nohup geth --datadir new-node-1 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50000 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft --emitcheckpoints --port 21000 2>>node.log &`

Your node is now operational and you may attach to it with `geth attach new-node-1/geth.ipc`. Tessera IPC bridge will be over a file name defined in your `config.json`, usually named `tm.ipc` as evidenced in prefix `PRIVATE_CONFIG=tm.ipc`. Your node is now able to send and receive private transactions, advertised public node key will be in the `new-node-1.pub` file. Tessera offers a lot of configuration flexibility, please refer [Configuration](../../Privacy/Tessera/Configuration/Configuration%20Overview) section under Tessera for complete and up to date configuration options.

### Constellation
1. Build Quorum and install [Constellation](https://github.com/jpmorganchase/constellation/releases) as described in the [getting set up](../Setup%20Overview%20%26%20Quickstart) section. Ensure that PATH contains geth, bootnode, and constellation-node binaries
2. Generate new keys with `constellation-node --generatekeys=new-node-1`
3. Start your constellation node and send it into background with `constellation-node --url=https://127.0.0.1:9001/ --port=9001 --workdir=. --socket=tm.ipc --publickeys=new-node-1.pub --privatekeys=new-node-1.key --othernodes=https://127.0.0.1:9001/ >> constellation.log 2>&1 &`
4. Start your node and send it into background with `PRIVATE_CONFIG=tm.ipc nohup geth --datadir new-node-1 --nodiscover --verbosity 5 --networkid 31337 --raft --raftport 50000 --rpc --rpcaddr 0.0.0.0 --rpcport 22000 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft --emitcheckpoints --port 21000 2>>node.log &`

Your node is now operational and you may attach to it with `geth attach new-node-1/geth.ipc`. Constellation IPC bridge will be over a file name defined in your configuration: in above step #3 see option `--socket=file-name.ipc`. Your node is now able to send and receive private transactions, advertised public node key will be in the `new-node-1.pub` file.


## Enabling permissioned configuration
Quorum ships with a permissions system based on a custom whitelist. Detailed documentation is available in [Network Permissioning](../../Permissioning/Overview.md).
