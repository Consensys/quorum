
# Running Quorum

Quorum introduces the `--permissioned` CLI argument: 

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

The [quorum-examples 7nodes](https://github.com/jpmorganchase/quorum-examples) source files contain several scripts demonstrating how to set up a private test network made up of 7 nodes.  

## Permissioned Network

Node Permissioning is a feature of Quorum that is used to define:
1. The nodes that a particular Quorum node is able to connect to
2. The nodes that a particular Quorum node is able to receive connections from

Permissioning is managed at the individual node level by using the `--permissioned` command line flag when starting the node.

If a node is started with `--permissioned` set, the node will look for a `<data-dir>/permissioned-nodes.json` file.  This file contains the list of enodes that this node can connect to and accept connections from.  In other words, if permissioning is enabled, only the nodes that are listed in the `permissioned-nodes.json` file become part of the network. 

If `--permissioned` is set, a `permissioned-nodes.json` file must be provided. If the flag is set but no nodes are present in this file, then the node will be unable to make any outward connections or accept any incoming connections.

The format of `permissioned-nodes.json` is similar to `static-nodes.json`:

```json
[
  "enode://enodehash1@ip1:port1",
  "enode://enodehash2@ip2:port2",
  "enode://enodehash3@ip3:port3"
]
```

For example, including the hash, a sample file might look like:

```json
[
  "enode://6598638ac5b15ee386210156a43f565fa8c48592489d3e66ac774eac759db9eb52866898cf0c5e597a1595d9e60e1a19c84f77df489324e2f3a967207c047470@127.0.0.1:30300"
]
```

In the current release, every node has its own copy of `permissioned-nodes.json`. In a future release, the permissioned nodes list will be moved to a smart contract, thereby keeping the list on-chain and requiring just one global list of nodes that connect to the network.
