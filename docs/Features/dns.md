# DNS for Quorum

DNS support in Quorum has two distinct areas, usage in the static nodes file and usage in the 
node discovery protocol. You are free to use one and not the other, or to mix them as the use case
requires.

### Static nodes

Static nodes are nodes we keep reference to even if the node is not alive, so that is the nodes comes alive, 
then we can connect to it. Hostnames are permitted here, and are resolved once at startup. If a static peer goes offline
and its IP address changes, then it is expected that that peer would re-establish the connection in a fully static 
network, or have discovery enabled.

### Discovery

DNS is not supported for the discovery protocol. Use a bootnode instead, which can use a DNS name that is repeatedly
resolved.

## DNS in Raft

On version 2.4.1+ in Raft mode, if the enode URL is provided with `raftport`. The URL hostname will be stored in ENR. 
If `raftport` is not provided, default behavior as per Ethereum `ParseV4` which only store IP in ENR.

### Compatibility
The whole network must be on version 2.4.1+ of Quorum for DNS to function properly; because of this, DNS must 
be explicitly enabled using the `--raftdnsenable` flag. The network will support older nodes mixed with newer nodes 
if DNS is not enabled via this flag, and it is safe to enable DNS only on some nodes if all nodes are on at least 
version 2.4.1. This allows for a clear upgrade path.