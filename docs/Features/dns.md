# DNS for Quorum

DNS support in Quorum has two distinct areas, usage in the static nodes file and usage in the 
node discovery protocol. You are free to use one and not the other, or to mix them as the use case
requires.

## Static nodes

Static nodes are nodes we keep reference to even if the node is not alive, so that is the nodes comes alive, 
then we can connect to it. Hostnames are permitted here, and are resolved once at startup. If a static peer goes offline
and its IP address changes, then it is expected that that peer would re-establish the connection in a fully static 
network, or have discovery enabled.

## Discovery

DNS is not supported for the discovery protocol. Use a bootnode instead, which can use a DNS name that is repeatedly
resolved.

## Compatibility
For Raft, the whole network must be on version 2.4.0 of Quorum for DNS to function properly.  DNS must 
be explicitly enabled using the `--raftdnsenable` flag for each node once the node has migrated to version 2.4.0 of Quorum
The network can still run in fine when some nodes are in 2.4.0 version and some in older version as long as this feature is not enabled. For safe migration the recommended approach is as below:
* migrate the nodes to `geth` 2.4.0 version without using `--raftdnsenable` flag
* once the network is fully migrated, restart the nodes with `--raftdnsenable` to enable the feature

Please note that in a partially migrated network  (where some nodes are on version 2.4.0 and others on lower version) **with DNS feature enabled** for migrated nodes, `raft.addPeer` should not be invoked with Hostname till entire network migrates to 2.4.0 version. `raft.addPeer` can still be invoked with IP address and network will work fine. 