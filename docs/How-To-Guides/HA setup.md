
Quorum architecture allows for true HA setup end to end for heightened availability for various input output operations. Although this increase the footprint of each node, the cost offset is compensate by low to zero downtime & horizontal scalability. In this section we will go through the configuration and setup in detail:

Quorum HA Setup:

For non-validator Quorum node (or “transaction node”), we will start another Quorum node (let’s call them Q 1-1 and Q1-2).
The inbound RPC requests from clients will be load balanced to either of those Quorum nodes (primary/backup mode.)
These 2 Quorum nodes will have their own enodes and addresses (act as 2 separate nodes in the network)
They will need to share the same private state, meaning talk to the same Tessera database
These 2 Quorum nodes will need to share the same account (the public/private keypair that is used to sign transactions)
A proxy will need to be installed on each Quorum node to listen on the local ipc file and direct the requests to Tessera Q2T HTTP
Run shell script to monitor the geth keystore directory then sftp it to the other host every time there is a new file and also in periodic intervals if the new file was created at the time of other node being down. Alternative is to create a shared NFS mount for the keystore directory across both VM’s.
