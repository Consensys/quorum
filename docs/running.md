
# Running Quorum

A `--permissioned` CLI argument was introduced with Quorum.

```
QUORUM OPTIONS:
  --permissioned              If enabled, the node will allow only a defined list of nodes to connect
```

The full list of arguments can be viewed by running `geth --help`.

### Initialize chain

The first step is to generate the genesis block.

The `7nodes` directory in the `quorum-examples` repository contains several keys (using an empty password) that are used in the example genesis file:

```
key1    vote key 1
key2    vote key 2
key3    vote key 3
key4    block maker 1
key5    block maker 2
```

Example genesis file (copy to `genesis.json`):
```json
{
  "alloc": {},
  "coinbase": "0x0000000000000000000000000000000000000000",
  "config": {
    "homesteadBlock": 0
  },
  "difficulty": "0x0",
  "extraData": "0x",
  "gasLimit": "0x2FEFD800",
  "mixhash": "0x00000000000000000000000000000000000000647572616c65787365646c6578",
  "nonce": "0x0",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp": "0x00"
}
```

Now we can initialize geth:

```
geth init genesis.json
```

### Setup Bootnode
Optionally you can set up a bootnode that all the other nodes will first connect to in order to find other peers in the network. You will first need to generate a bootnode key:

1. To generate the key for the first time:

`bootnode -genkey tmp_file.txt  // this will start a bootnode with an enode address and generate a key inside a “tmp_file.txt” file`

2. To later restart the bootnode using the same key (and hence use the same enode url):

`bootnode -nodekey tmp_file.txt`

or

`bootnode -nodekeyhex 77bd02ffa26e3fb8f324bda24ae588066f1873d95680104de5bc2db9e7b2e510 // Key from tmp_file.txt`


### Start node

Starting a node is as simple as `geth`. This will start the node without any of the roles and makes the node a spectator. If you have setup a bootnode then be sure to add the `--bootnodes` param to your startup command:

`geth --bootnodes $BOOTNODE_ENODE`

## Setup multi-node network

Quorum comes with several scripts to setup a private test network with 7 nodes in the `7nodes` folder in the `quorum-examples` repository.

1. Step 1, run `raft-init.sh` and initialize data directories (change variables accordingly)
2. Step 2, start nodes with `raft-start.sh` (change variables accordingly)
3. Step 3, stop network with `stop.sh`

## Permissioned Network

Node Permissioning is a feature that controls which nodes can connect to a given node and also to which nodes this node can dial out to. Currently, it is managed at individual node level by the command line flag `--permissioned` while starting the node.

If the `--permissioned` node is present, the node looks for a file named `<data-dir>/permissioned-nodes.json`. This file contains the list of enodes that this node can connect to and also accepts connections only from those nodes. In other words, if permissioning is enabled, only the nodes that are listed in this file become part of the network. It is an error to enable `--permissioned` but not have the `permissioned-nodes.json` file. If the flag is given, but no nodes are present in this file, then this node can neither connect to any node or accept any incoming connections.

The `permissioned-nodes.json` follows following pattern (similar to `static-nodes.json`):

```json
[
  "enode://enodehash1@ip1:port1",
  "enode://enodehash2@ip2:port2",
  "enode://enodehash3@ip3:port3",
]
```

Sample file:

```json
[
  "enode://6598638ac5b15ee386210156a43f565fa8c48592489d3e66ac774eac759db9eb52866898cf0c5e597a1595d9e60e1a19c84f77df489324e2f3a967207c047470@127.0.0.1:30300",
]
```

In the current release, every node has its own copy of `permissioned-nodes.json`. In a future release, the permissioned nodes list will be moved to a smart contract, thereby keeping the list on chain and one global list of nodes that connect to the network.
