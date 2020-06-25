## What is Clef?
`clef` is a tool introduced in `geth` `v1.9.0` and is included as part of the `go-ethereum` project.  

`clef` runs as a separate process to `geth` and provides an alternative method of managing accounts and signing transactions/data.  Instead of `geth` loading and using accounts directly, `geth` delegates account management responsibilities to `clef`.

`go-ethereum`'s intention is to deprecate account management within `geth` and replace it with `clef`.

Using `clef` instead of `geth` for account management has several benefits: 

* Users and DApps no longer have a dependency on access to a synchronised local node loaded with accounts.  Transactions and DApp data can instead be signed using `clef`
* Future account-related features will likely only be available in `clef` and not found in `geth` (e.g. [EIP-191 and EIP-712 have been implemented in `clef`, but there is no intention of implementing them in `geth`](https://github.com/ethereum/go-ethereum/pull/17789/)) 
* `clef` provides several user-experience improvements to ease use and improve security
* More info can be found in the additional documentation `.md` files in the [cmd/clef directory](https://github.com/jpmorganchase/quorum/tree/master/cmd/clef)

## What is Clef for Quorum?
`clef` was introduced in Quorum `v2.6.0`.

Clef for Quorum is the standard `go-ethereum` `clef` tool, with additional support for Quorum-specific features, including:

* Support for private transactions

## Installation
`geth` and all included tools (i.e. `clef`, `bootnode`, ...) can be installed to `PATH` by [building Quorum from source with `make all`](../../Getting%20Started/Installing#quorum).

Verify the installation with:
```shell
clef help
```

## Getting Started

See [cmd/clef/tutorial.md](https://github.com/jpmorganchase/quorum/blob/master/cmd/clef/tutorial.md) for an overview and step-by-step guide on initialising and starting `clef`, as well as configuring automation rules.

## Usage

`clef` can be used in one of two ways:

1. As an external signer
1. As a `geth` signer

!!! warning
    In the long term, the preferred way of using `clef` will be as an external signer.  However, whilst waiting for tooling to support the `clef` API, the `go-ethereum` project have included the option to use `clef` as a `geth` signer.  This ensures existing tooling and user flows can remain unchanged.  The option to use `clef` as a `geth` signer **will be deprecated** in a future release of `go-ethereum` once the migration of account management from `geth` to `clef` is complete.

### As an external signer
Using `clef` as an external signer requires interacting with `clef` through its RPC API.  By default this is exposed over IPC socket.  The API can also be exposed over HTTP by using the `--rpcaddr` CLI flag.

An example workflow would be:

1. Start `clef` and make your accounts available to it
1. Sign a transaction with the account by using `clef`'s `account_signTransaction` API.  `clef` will return the signed transaction.
1. Use `eth_sendRawTransaction` or `eth_sendRawPrivateTransaction` to send the signed transaction to a Quorum node that does not have your accounts available to it
1. The Quorum node will validate the transaction and propagate it through the network for minting 

#### Example: List accounts

```shell
echo '{"id": 1, "jsonrpc": "2.0", "method": "account_list"}' | nc -U /path/to/clef.ipc
```

#### Example: Sign data

```shell
echo '{"id": 1, "jsonrpc": "2.0", "method": "account_signData", "params": ["data/plain", "0x6038dc01869425004ca0b8370f6c81cf464213b3", "0xaaaaaa"]}' | nc -U /path/to/clef.ipc
``` 

### As a geth signer
Using `clef` as a `geth` signer will not require direct interaction through the `clef` API.  Instead `geth` can be used as normal and will automatically delegate to `clef`.

To use `clef` as a `geth` signer:

1. Start `clef`
1. Start `geth` with the `--signer /path/to/clef.ipc` CLI flag 

An example workflow would be:

1. Start `clef` and make your accounts available to it
1. Start `geth` and do not make your accounts available to it
1. Use `eth_sendTransaction` to sign and submit a transaction for validation, propagation, and minting 
