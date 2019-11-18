# DNS for Quorum

DNS support in Quorum has two distinct areas, usage in the static nodes file and usage in the 
node discovery protocol. You are free to use one and not the other, or to mix them as the use case
requires.

## Static nodes

Static nodes are nodes we keep reference to even if the node is not alive, so that is the nodes comes alive, 
then we can connect to it. Using DNS names in the static node definitions means that if the node is offline and
it's IP changes, we will still be able to connect to it when it comes back online, unlike Ethereum. 

## Discovery

If using the DiscoveryV4 protocol with DNS resolution, then you will need to provide the hostname that you wish
for others to see to use to connect to you. This is achieved with the `--hostname <hostname>` flag. If this 
flag is set to something that does not resolve to your IP address, then other peers will try to connect to
the wrong address and fail, so it is important that this value is correct.
If you provide a value that does not resolve to any IP address, the node will fail to start up and emit an error
notifying the node operator of such; but it cannot protect against valid but incorrect hostname being provided.