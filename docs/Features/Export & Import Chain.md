# Quorum Export & Import Chain

Quorum supports export and import of chain data. This is an effective backup mechanism but Quorum has a specialized 
design of it to adapt for different consensus need and private transaction.

### Quorum Export

Quorum export chain must run after geth process is stopped. It supports at most 4 arguments. 1) export file name & 2) genesis file are mandatory. 3) first block & 4) 
last block are optional but must provide together if used.

##### Sample command for Quorum export

`geth export <export file name> <genesis file> --datadir <geth data dir>`

### Quorum Import

Quorum import chain must run on a new datadir before `geth init`. It supports arbitrary number of import files (at least 1) and a mandatory genesis file.

##### Sample command for Quorum import

`PRIVATE_CONFIG=<PTM ipc endpoint> geth import <import file names...> <genesis file> --datadir <geth data dir>`

### Special Explanation Notes

##### Private Transaction Manager

If private transactions are used in the chain data, Private Transaction Manager process for the original exported node 
must be running on the PTM ipc endpoint during import chain. Otherwise, nil pointer exceptions will be raised.

##### IBFT

IBFT block data contains seal information in the header, to use a copy of exported chain data, the new node must use an 
IBFT genesis file with exact same validator set extra data encoding as original exported node's genesis.

##### Raft

Raft consensus is not reflected at the chain data level. An exported chain data from a Raft cluster can only be used by 
new nodes adding to that cluster only. 