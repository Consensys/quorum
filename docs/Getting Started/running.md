# Running Quorum

## Developing Smart Contracts
Quorum uses standard [Solidity](https://solidity.readthedocs.io/en/develop/) for writing Smart Contracts, and generally, these can be designed as you would design Smart Contracts for Ethereum.  Smart Contracts can either be public (i.e. visible and executable by all participants on a given Quorum network) or private to one or more network participants.  Note, however, that Quorum does not introduce new contract Types. Instead, similar to [Transactions](../../Transaction%20Processing/Transaction%20Processing), the concept of public and private contracts is notional only.

### Creating Public Transactions/Contracts

Sending a standard Ethereum-style transaction to a given network will make it viewable and executable by all participants on the network.  As with Ethereum, leave the `to` field empty for a contract-creation transaction.

Example JSON RPC API call to send a public transaction:

``` json
{    
    "jsonrpc":"2.0",
    "method":"eth_sendTransaction",
    "params":[
        {
            "from": "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
            "to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
            "gas": "0x76c0", // 30400
            "gasPrice": "0x9184e72a000", // 10000000000000
            "value": "0x9184e72a", // 2441406250
            "data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
        }
    ],
    "id":1
}
```

See the [Quorum API](../api) page for details on the `sendTransaction` call, which includes some modifications to the standard Ethereum call.

!!! info
    See the Contract Design Considerations sections below for important points on creating Quorum contracts

### Creating Private Transactions/Contracts
In order to make a transaction/smart contract private and therefore only viewable and executable by a subset of the network, send a standard Ethereum Transaction but include the Quorum-specific `privateFor` parameter.  `privateFor` is used to provide the list of participants for the transaction/contract.  Each participant is identified by a Privacy Manager public key. 

Example JSON RPC API call to send a private transaction:

``` json
{
    "jsonrpc":"2.0",
    "method":"eth_sendTransaction",
    "params":[
        {
            "from": "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
            "to": "0xd46e8dd67c5d32be8058bb8eb970870f07244567",
            "gas": "0x76c0", // 30400
            "gasPrice": "0x9184e72a000", // 10000000000000
            "value": "0x9184e72a", // 2441406250
            "data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675",
            "privateFor": ["$PUBKEY1, $PUBKEY2"]
        }
    ],
    "id":1
}
```

See the [Quorum API](../api) page for details on the `sendTransaction` call, which includes some modifications to the standard Ethereum call.

!!! info
    See the Contract Design Considerations sections below for important points on creating Quorum contracts

### Quorum Contract Design Considerations

1. *Private contracts cannot update public contracts.*  This is because not all participants will be able to execute a private contract, and so if that contract can update a public contract, then each participant will end up with a different state for the public contract.
2. *Once a contract has been made public, it can't later be made private.*  If you do need to make a public contract private, it would need to be deleted from the blockchain and a new private contract created.

## Setting up a multi-node network

The [quorum-examples 7nodes](../Quorum-Examples) source files contain several scripts demonstrating how to set up a private test network made up of 7 nodes.  

## Permissioned Networks

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
``` json
"config": {
    "homesteadBlock": 0,
    "byzantiumBlock": 0,
    "chainId": 10,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "eip158Block": 0,
    "isQuorum": true
  },
```

Now we can initialize geth:

```
geth init genesis.json
```

### Setup Bootnode
Optionally you can set up a bootnode that all the other nodes will first connect to in order to find other peers in the network. You will first need to generate a bootnode key:

1. To generate the key for the first time:

    ``` bash
    bootnode -genkey tmp_file.txt  // this will start a bootnode with an enode address and generate a key inside a “tmp_file.txt” file`
    ```

2. To later restart the bootnode using the same key (and hence use the same enode url):

    ``` bash
    bootnode -nodekey tmp_file.txt
    ```

    or

    ``` bash
    bootnode -nodekeyhex 77bd02ffa26e3fb8f324bda24ae588066f1873d95680104de5bc2db9e7b2e510 // Key from tmp_file.txt
    ```

### Start node

Starting a node is as simple as `geth`. This will start the node without any of the roles and makes the node a spectator. If you have setup a bootnode then be sure to add the `--bootnodes` param to your startup command:

`geth --bootnodes $BOOTNODE_ENODE`

### Adding New Nodes:
Any additions to the `permissioned-nodes.json` file will be dynamically picked up by the server when subsequent incoming/outgoing requests are made. The node does not need to be restarted in order for the changes to take effect. 

### Removing existing nodes:
Removing existing connected nodes from the `permissioned-nodes.json` file will not immediately drop those existing connected nodes. However, if the connection is dropped for any reason, and a subsequent connect request is made from the dropped node ids, it will be rejected as part of that new request.

## Quorum API
Please see the [Quorum API](../../api) page for details.

## Network and Chain ID

An Ethereum network is run using a Network ID and, after [EIP-155](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md), a Chain ID.

Before EIP-155, the names "Network ID" and "Chain ID" were used interchangeably, but after this they have separate meanings.

The network ID is a property of a peer, NOT of the chain the peer is managing. A network ID can be passed in via the command line by `--networkid <id>`. It's purpose is to separate peers that are running under a different network ID. Therefore, you  cannot sync with anyone who is running a node with a different network ID. However, since it is trivial to change this, it is a less secure version of Quorum's `--permissioned` flag, and it only used for simple segregation.

The chain ID is a property of the chain managed by the node. It is used for replay protection of transactions - prior to EIP-155, a transaction run on one chain could be copied and sent to a different chain by anyone, since the transaction is already signed.

Setting the chain ID has the effect of changing one of the parameters of a transaction, namely the `V` parameter.
As the EIP explains, the `v` parameter is set to `2*ChainID + 35/36`. For the Ethereum Foundation Mainnet, which has a chain ID of `1`, this means that all transactions have a value of either `37` or `38`.

The chain ID set in the genesis configuration file, under the `config` section, and is only used when the block number is above the one set at `eip155Block`. See the [quorum-examples genesis files](../genesis) for an example. It can be changed as many times as needed whilst the chain is below the `eip155Block` number and re-rerunning `geth init` - this will not delete or modify any current sync process or saved blocks!

In Quorum, transactions are considered private if the `v` parameter is set to `37` or `38`, which clashes with networks which have a Chain ID of `1`. For this reason, Quorum will not run using chain ID `1` and will immediately quit if started with such a configuration from version 2.1.0 onwards.
If you are running a version prior to version 2.1.0, EIP-155 signing is not used, thus a chain ID of `1` was allowed; you will need to change this using `geth init` before running an updated version.


## ZSL Proof of Concept

J.P. Morgan and the Zcash team partnered to create a proof of concept (POC) implementation of ZSL for Quorum, which enables the issuance of digital assets using ZSL-enabled public smart contracts (z-contracts). We refer to such digital assets as “z-tokens”. Z-tokens can be shielded from public view and transacted privately. Proof that a shielded transaction has been executed can be presented to a private contract, thereby allowing the private contract to update its state in response to shielded transactions that are executed using public z-contracts. 

This combination of Constellation/Tessera’s private contracts with ZSL’s z-contracts, allows obligations that arise from a private contract, to be settled using shielded transfers of z-tokens, while maintaining full privacy and confidentiality.

For more information, see the [ZSL](../../ZSL) page of this wiki.

## Configurable transaction size:

Quorum allows operators of blockchains to increase maximum transaction size of accepted transactions via the genesis block. The Quorum default is currently increased to `64kb` from Ethereum's default `32kb` transaction size.  This is configurable up to `128kb` by adding `txnSizeLimit` to the config section of the genesis file:

``` json
"config": {
    "chainId": 10,
    "isQuorum":true.
    ...
    "txnSizeLimit": 128
}
```

