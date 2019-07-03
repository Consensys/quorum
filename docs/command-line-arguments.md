# Command Line Arguments

Since Quorum is a fork of go-ethereum, it inherits all of the same CLI arguments. 
See [here](https://geth.ethereum.org/docs/interface/command-line-options) for a list of options, or run `geth help`.

## Quorum CLI arguments

Quorum adds a number of CLI arguments for the features it has added. These are also visible by running `geth help`, and 
are laid out here.

### DNS support

`--hostname <name>` - the hostname you want to advewrtise to other peers to connect to you. e.g. `--hostname host1.com` would
set your enode ID to be `enode://<enodeId>@host1.com:21005`.

`--raft.dns-enable` - DNS support for the RAFT layer must be explicitly enabled, and is imcompatible with older 
versions. This should be enabled when all nodes have upgraded.