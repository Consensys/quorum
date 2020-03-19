# Upgrade to Quorum 2.6.0

Quorum 2.6.0 is a fork of ethereum 1.9.7. 
See [ethereum 1.9.0](https://blog.ethereum.org/2019/07/10/geth-v1-9-0/) for the full list of features. 
Once migrated to 2.6.0 you cannot rollback to your old version. You should back up `datadir` if you need to rollback.
You can migrate node by node and have mixed node network running during migration.

### Key features to be aware of
* freezerdb - you can provide separate location for freezerdb via geth commandline arguments.

```  --datadir.ancient value             Data directory for ancient chain segments (default = inside chaindata) ```
* enable account unlocking explicitly when an account is unlocked from geth commandline arguments
``` --allow-insecure-unlock             Allow insecure account unlocking when account-related RPCs are exposed by http ```

* include `istanbulForkBlock` in `genesis.json` and update genesis

* `--exitwhensynced` geth commandline argument won't work for Raft

