# Adding nodes to the network

Adding new nodes to an existing network can range from a common occurence to never happening.
In public blockchains, such as the Ethereum Mainnet, new nodes continuously join and talk to the existing network.
In permissioned blockchains, this may not happen as often, but it still an important task to achieve as your network 
evolves.

When adding new nodes to the network, it is important understand that the Quorum network and Private Transaction 
Manager network are distinct and do not overlap in any way. Therefore, options applicable to one are not applicable to 
the other. In some cases, they may have their own options to achieve similar tasks, but must be specified separately.

## Prerequisites

- [Quorum installed](/Getting%20Started/Installing.md)
- [Tessera/Constellation installed](/Getting%20Started/Installing.md) if using private transactions
- A running network (see [Creating a Network From Scratch](/Getting%20Started/Creating-A-Network-From-Scratch))

## Adding Quorum nodes

Adding a new Quorum node is the most common operation, as you can choose to run a Quorum node with or without a Private 
Transaction Manager, but rarely will one do the opposite.

### Raft

1. On an *existing* node, add the new peer to the raft network
    ```
    > raft.addPeer("enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0&raftport=50407")
    7
    ```
    
    So in this example, our new node has a Raft ID of `7`.

2. If you are using permissioning, or discovery for Ethereum p2p, please refer [here](#extra-options).

3. We now need to initialise the new node with the network's genesis configuration.

    !!! note
        Where you obtain this from will be dependent on the network. You may get it from an existing peer, or a network operator, or elsewhere entirely.

    Initialising the new node is exactly the same an the original nodes.
    ```bash
    $ geth --datadir qdata/dd7 init genesis.json
    ```

4. Now we can start up the new node and let it sync with the network. The main difference now is the use of the 
`--raftjoinexisting` flag, which lets the node know that it is joining an existing network, which is handled 
differently internally. The Raft ID obtained in step 1 is passed as a parameter to this flag.

    ```bash
    $ PRIVATE_CONFIG=ignore geth --datadir qdata/dd7 ... OTHER ARGS ... --raft --raftport 50407 --rpcport 22006 --port 21006 --raftjoinexisting 7
    ```
   
   The new node is now up and running, and will start syncing the blockchain from existing peers. Once this has 
   completed, it can send new transactions just as any other peer.

### IBFT/Clique

Adding nodes to an IBFT/Clique network is a bit simpler, as it only needs to configure itself rather then be 
pre-allocated on the network (permissioning aside).

1. Initialise the new node with the network's genesis configuration.
    
    !!! note
        Where you obtain this from will be dependent on the network. You may get it from an existing peer, or a network operator, or elsewhere entirely.

    Initialising the new node is exactly the same an the original nodes.
    ```bash
    $ geth --datadir qdata/dd7 init genesis.json
    ```

2. If you are using permissioning or discovery for Ethereum peer-to-peer, please refer [here](#extra-options).

3. Start the new node, pointing either to a `bootnode` or listing an existing peer in the `static-nodes.json` file. 
Once a connection is established, the node will start syncing the blockchain, after which transactions can be sent.

### Extra options

Some options take effect regardless of the consensus mechanism used.

#### Permissioned nodes

If using the `permissioned-nodes.json` file for permissioning, then you must make sure this file is updated on all 
nodes before the new node is able to communicate with existing nodes. You do not need to restart any nodes in 
order for the changes to take effect.

#### Static node connections

If not using peer-to-peer node discovery (i.e. you have specified `--nodiscover`), then the only connections a node 
made will be to peers defined in the `static-nodes.json` file. When adding a new node, you should make sure you have 
peers defined in its `static-nodes.json` file. The more peers you have defined here, the better network connectivity 
and fault tolerance you have.

!!! note
    * You do not need to update the existing peers static nodes for the connection to be established, although it is good practise to do so.
    * You do not need to specify every peer in your static nodes file if you do not wish to connect to every peer directly.

#### Peer-to-peer discovery

If you are using discovery, then more options *in addition* to static nodes become available.

- Any nodes that are connected to your peers, which at the start will be ones defined in the static node list, will 
then be visible by you, allowing you to connect to them; this is done automatically.

- You may specify any number of bootnodes, defined by the `--bootnodes` parameter. This takes a commas separated list 
of enode URIs, similar to the `static-nodes.json` file. These act in the same way as static nodes, letting you connect 
to them and then find out about other peers, whom you then connect to.

!!! note
    If you have discovery disabled, this means you will not try to find other nodes to connect to, but others can still find and connect to you.

## Adding Private Transaction Managers

In this tutorial, there will be no focus on the advanced features of adding a new Private Transaction Manager (PTM).
This tutorial uses [Tessera](https://github.com/jpmorganchase/tessera) for any examples.

Adding a new node to the PTM is relatively straight forward, but there are a lot of extra options that can be used, 
which is what will be explained here.

### Adding a new PTM node

In a basic setting, adding a new PTM node is as simple as making sure you have one of the existing nodes listed in your 
peer list.

In Tessera, this would equate to the following in the configuration file:
```json
{
  "peers": [
    {
      "url": "http://existingpeer1.com:8080"
    }
  ]
}
```

From there, Tessera will connect to that peer and discover all the other PTM nodes in the network, connecting to each 
of them in turn.

!!! note
    You may want to include multiple peers in the peer list in case any of them are offline/unreachable.
    
### IP whitelisting

The IP Whitelist that Tessera provides allows you restrict connections much like the `permissioned-nodes.json` file 
does for Quorum. Only IP addresses/hostnames listed in your peers list will be allowed to connect to you.

See the [Tessera configuration page](/Privacy/Tessera/Configuration/Configuration%20Overview#whitelist) for details on setting it up.

In order to make sure the new node is accepted into the network:

1. You will need to add the new peer to each of the existing nodes before communication is allowed.
    Tessera provides a way to do this without needing to restart an already running node:
    ```bash
    $ java -jar tessera.jar admin -configfile /path/to/existing-node-config.json -addpeer http://newpeer.com:8080
    ```
   
2. The new peer can be started, setting the `peers` configuration to mirror the existing network.
    e.g. if there are 3 existing nodes in the network, then the new nodes configuration will look like this:
    ```json
    {
      "peers": [
        {
          "url": "http://existingpeer1.com:8080"
        },
        {
          "url": "http://existingpeer2.com:8080"
        },
        {
          "url": "http://existingpeer3.com:8080"
        }
      ]
    }
    ```

    The new node will allow incoming connections from the existing peers, and then existing peers will allow incoming 
    connections from the new peer!
    
### Discovery

Tessera discovery is very similar to the IP whitelist. The difference being that the IP whitelist blocks 
communications between nodes, whereas disabling discovery only affects which public keys we keep track of.

See the [Tessera configuration page](/Privacy/Tessera/Configuration/Configuration%20Overview#disabling-peer-discovery) for 
details on setting it up.
   
When discovery is disabled, Tessera will only allow keys that are owned by a node in its peer list to be available to 
the users. This means that if any keys are found that are owned by a node NOT in our peer list, they are discarded and 
private transactions cannot be sent to that public key.

!!! note
    This does not affect incoming transactions. Someone not in your peer list can still send transactions to your node, unless you also enable the IP Whitelist option.
    
    
    
In order to make sure the new node is accepted into the network:

1. You will need to add the new peer to each of the existing nodes before they will accept public keys that are linked 
to the new peer.
    Tessera provides a way to do this without needing to restart an already running node:
    ```bash
    $ java -jar tessera.jar admin -configfile /path/to/existing-node-config.json -addpeer http://newpeer.com:8080
    ```
    
2. The new peer can be started, setting the `peers` configuration to mirror the existing network.
    e.g. if there are 3 existing nodes in the network, then the new nodes configuration will look like this:
    ```json
    {
      "peers": [
        {
          "url": "http://existingpeer1.com:8080"
        },
        {
          "url": "http://existingpeer2.com:8080"
        },
        {
          "url": "http://existingpeer3.com:8080"
        }
      ]
    }
    ```
    
    The new node will now record public keys belonging to the existing peers, and then existing peers will record 
    public keys belonging to the new peer; this allows private transactions to be sent both directions!
    
    
## Examples

For a walkthrough of some examples that put into action the above, check out [this guide](/How-To-Guides/add_node_examples)!