## Network Permissioning

Network Permissioning is a feature that controls which nodes can connect to a given node and also to which nodes the given node can dial out to. Currently, it is managed at the individual node level by the `--permissioned` command line flag when starting the node.

If the `--permissioned` flag is set, the node looks for a file named `<data-dir>/permissioned-nodes.json` . This file contains the whitelist of enodes that this node can connect to and accept connections from. Therefore, with permissioning enabled, only the nodes that are listed in the `permissioned-nodes.json` file become part of the network. If the `--permissioned` flag is specified but no nodes are added to the `permissioned-nodes.json` file then this node can neither connect to any node nor accept any incoming connections.

The `permissioned-nodes.json` file follows the below pattern, which is similar to the `<data-dir>/static-nodes.json` file that is used to specify the list of static nodes a given node always connects to:
   ``` json
    [ 
        "enode://remoteky1@ip1:port1",
        "enode://remoteky1@ip2:port2",
        "enode://remoteky1@ip3:port3", 
    ]
   ```
    
Sample file: (node id truncated for clarity)
   ``` json
    [ 
      "enode://6598638ac5b15ee386210156a43f565fa8c485924894e2f3a967207c047470@127.0.0.1:30300",
    ]
   ```

!!! Note
    In the current implementation, every node has its own copy of the `permissioned-nodes.json` file. In this case, if different nodes have a different list of remote keys then each node may have a different list of permissioned nodes - which may have an adverse effect. In a future release, the permissioned nodes list will be moved from the `permissioned-nodes.json` file to a Smart Contract, thereby ensuring that all nodes will use one global on-chain list to verify network connections. 

