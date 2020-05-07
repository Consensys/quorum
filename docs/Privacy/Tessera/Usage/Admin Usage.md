Administrators of a Tessera node can use the `admin` CLI command to make changes to the node.  These changes are made while the node is running and do not require a node restart.

The `admin` CLI makes use of the [ADMIN server API](../Interface%20%26%20API) and provides some additional features.  An ADMIN server must have been configured at startup (see [Configuration Overview](../../Configuration/Configuration%20Overview)).  

After starting a node with `tessera -configfile /path/to/node-config.json`, the admin CLI can be used.  Currently supported admin commands are:
- `addpeer`: Add a new peer to a running node

### `addpeer`
```
tessera admin -configfile /path/to/node-config.json -addpeer <new-peer-url>
```
The provided configfile is the same configfile used to start the Tessera node.

This will do two things:

1. Add `<new-peer-url>` to the node's list of peers, by using the ADMIN API
1. Update the configfile `/path/to/node-config.json` to include `<new-peer-url>` in the `peer` list.  Updating the configfile in this way means that if the node is stopped and started again, the admin changes will still be present.

If the configfile should not be updated, use the ADMIN API directly.
