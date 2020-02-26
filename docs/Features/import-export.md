# Backup & Restore of Quorum Nodes

Quorum supports export and import of chain data with built in tooling. This is an effective node backup mechanism adapted for the specific needs of Quorum such 
as private transactions, permissioning, and supported consensus algorithms.


!!! note
    Quorum chain data import and export must run after `geth` process is stopped.

### Node Backup (Export)

Backup functionality mimics original `geth export` command but requires a stricter set of arguments. Quorum export accepts 3 arguments:

1. Export file name **required**
3. First block
4. Last block *are optional but must be provided together when used*

##### Sample command

`geth export <export file name> --datadir <geth data dir>`

### Node Restore (Import)

Restore functionality mimics original `geth import` command but requires a stricter set of arguments, as well as, corresponding transaction manager.
Quorum import must run on a new node with a new `--datadir` before `geth init` has been executed. Restore supports arbitrary number of import files (at least 1) and a mandatory genesis file.

!!! warning
    If private transactions are used in the chain data, Private Transaction Manager process for the original exported node 
    must be running on the PTM ipc endpoint during import chain. Otherwise, nil pointer exceptions will be raised.

##### Sample command

`PRIVATE_CONFIG=<PTM ipc endpoint> geth import <import file names...> --datadir <geth data dir>`

### Special Consensus Considerations

##### IBFT

IBFT block data contains sealer information in the header, to restore a copy of exported chain data, the new node must use an 
IBFT genesis file with exact same validator set encoded in extra data field as original exported node's genesis.

##### Raft

Raft backup do not account for current Raft state. An exported chain data from a Raft cluster can only be used by 
new nodes being added to that cluster only.