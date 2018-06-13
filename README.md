# Quorum

<a href="https://quorumslack.azurewebsites.net" target="_blank" rel="noopener"><img title="Quorum Slack" src="https://quorumslack.azurewebsites.net/badge.svg" alt="Quorum Slack" /></a>

Quorum is an Ethereum-based distributed ledger protocol with transaction/contract privacy and new consensus mechanisms.

Quorum is a fork of [go-ethereum](https://github.com/ethereum/go-ethereum) and is updated in line with go-ethereum releases.

Key enhancements over go-ethereum:

  * __Privacy__ - Quorum supports private transactions and private contracts through public/private state separation and utilising [Constellation](https://github.com/jpmorganchase/constellation), a peer-to-peer encrypted message exchange for directed transfer of private data to network participants
  * __Alternative Consensus Mechanisms__ - with no need for POW/POS in a permissioned network, Quorum instead offers multiple consensus mechanisms that are more appropriate for consortium chains:
    * __Raft-based Consensus__ - a consensus model for faster blocktimes, transaction finality, and on-demand block creation
    * __Istanbul BFT__ - a PBFT-inspired consensus algorithm with transaction finality, by AMIS.
  * __Peer Permissioning__ - node/peer permissioning using smart contracts, ensuring only known parties can join the network
  * __Higher Performance__ - Quorum offers significantly higher performance than public geth

Note: The QuorumChain consensus algorithm is not yet supported by this release.

## Architecture

<a href="https://github.com/jpmorganchase/quorum/wiki/Transaction-Processing#private-transaction-process-flow">![Quorum privacy architecture](https://github.com/jpmorganchase/quorum-docs/raw/master/images/QuorumTransactionProcessing.JPG)</a>

The above diagram is a high-level overview of the privacy architecture used by Quorum. For more in-depth discussion of the components, refer to the [wiki](https://github.com/jpmorganchase/quorum/wiki/) pages.

## Quickstart

The quickest way to get started with Quorum is using [VirtualBox](https://www.virtualbox.org/wiki/Downloads) and [Vagrant](https://www.vagrantup.com/downloads.html):

```sh
git clone https://github.com/jpmorganchase/quorum-examples
cd quorum-examples
vagrant up
# (should take 5 or so minutes)
vagrant ssh
```

Now that you have a fully-functioning Quorum environment set up, let's run the 7-node cluster example. This will spin up several nodes with a mix of voters, block makers, and unprivileged nodes.

```sh
# (from within vagrant env, use `vagrant ssh` to enter)
ubuntu@ubuntu-xenial:~$ cd quorum-examples/7nodes

$ ./raft-init.sh
# (output condensed for clarity)
[*] Cleaning up temporary data directories
[*] Configuring node 1
[*] Configuring node 2 as block maker and voter
[*] Configuring node 3
[*] Configuring node 4 as voter
[*] Configuring node 5 as voter
[*] Configuring node 6
[*] Configuring node 7

$ ./raft-start.sh
[*] Starting Constellation nodes
[*] Starting bootnode... waiting... done
[*] Starting node 1
[*] Starting node 2
[*] Starting node 3
[*] Starting node 4
[*] Starting node 5
[*] Starting node 6
[*] Starting node 7
[*] Unlocking account and sending first transaction
Contract transaction send: TransactionHash: 0xbfb7bfb97ba9bacbf768e67ac8ef05e4ac6960fc1eeb6ab38247db91448b8ec6 waiting to be mined...
true
```

We now have a 7-node Quorum cluster with a [private smart contract](https://github.com/jpmorganchase/quorum-examples/blob/master/examples/7nodes/script1.js) (SimpleStorage) sent from `node 1` "for" `node 7` (denoted by the public key passed via `privateFor: ["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="]` in the `sendTransaction` call).

Connect to any of the nodes and inspect them using the following commands:

```sh
$ geth attach ipc:qdata/dd1/geth.ipc
$ geth attach ipc:qdata/dd2/geth.ipc
...
$ geth attach ipc:qdata/dd7/geth.ipc


# e.g.

$ geth attach ipc:qdata/dd2/geth.ipc
Welcome to the Geth JavaScript console!

instance: Geth/v1.5.0-unstable/linux/go1.7.3
coinbase: 0xca843569e3427144cead5e4d5999a3d0ccf92b8e
at block: 679 (Tue, 15 Nov 2016 00:01:05 UTC)
 datadir: /home/ubuntu/quorum-examples/7nodes/qdata/dd2
 modules: admin:1.0 debug:1.0 eth:1.0 net:1.0 personal:1.0 quorum:1.0 rpc:1.0 txpool:1.0 web3:1.0

# let's look at the private txn created earlier:
> eth.getTransaction("0xbfb7bfb97ba9bacbf768e67ac8ef05e4ac6960fc1eeb6ab38247db91448b8ec6")
{
  blockHash: "0xb6aec633ef1f79daddc071bec8a56b7099ab08ac9ff2dc2764ffb34d5a8d15f8",
  blockNumber: 1,
  from: "0xed9d02e382b34818e88b88a309c7fe71e65f419d",
  gas: 300000,
  gasPrice: 0,
  hash: "0xbfb7bfb97ba9bacbf768e67ac8ef05e4ac6960fc1eeb6ab38247db91448b8ec6",
  input: "0x9820c1a5869713757565daede6fcec57f3a6b45d659e59e72c98c531dcba9ed206fd0012c75ce72dc8b48cd079ac08536d3214b1a4043da8cea85be858b39c1d",
  nonce: 0,
  r: "0x226615349dc143a26852d91d2dff1e57b4259b576f675b06173e9972850089e7",
  s: "0x45d74765c5400c5c280dd6285a84032bdcb1de85a846e87b57e9e0cedad6c427",
  to: null,
  transactionIndex: 1,
  v: "0x25",
  value: 0
}
```

Note in particular the `v` field value of "0x25" or "0x26" (37 or 38 in decimal) which marks this transaction as having a private payload (input).

## Demonstrating Privacy
Documentation detailing steps to demonstrate the privacy features of Quorum can be found in [quorum-examples/7nodes/README](https://github.com/jpmorganchase/quorum-examples/tree/master/examples/7nodes/README.md).

## Further Reading

Further documentation can be found in the [docs](docs/) folder and on the [wiki](https://github.com/jpmorganchase/quorum/wiki/).

## See also

* [Quorum](https://github.com/jpmorganchase/quorum): this repository
* [Constellation](https://github.com/jpmorganchase/constellation): peer-to-peer encrypted message exchange for transaction privacy
* [Raft Consensus Documentation](raft/doc.md)
* [Istanbul BFT Consensus Documentation](https://github.com/ethereum/EIPs/issues/650): [RPC API](https://github.com/getamis/go-ethereum/wiki/RPC-API) and [technical article](https://medium.com/getamis/istanbul-bft-ibft-c2758b7fe6ff)
* [ZSL](https://github.com/jpmorganchase/quorum/wiki/ZSL) wiki page and [documentation](https://github.com/jpmorganchase/zsl-q/blob/master/README.md)
* [quorum-examples](https://github.com/jpmorganchase/quorum-examples): example quorum clusters
* [quorum-tools](https://github.com/jpmorganchase/quorum-tools): local cluster orchestration, and integration testing tool
* [Quorum Wiki](https://github.com/jpmorganchase/quorum/wiki)

## Third Party Tools/Libraries

The following Quorum-related libraries/applications have been created by Third Parties and as such are not specifically endorsed by J.P. Morgan.  A big thanks to the developers for improving the tooling around Quorum!

* [Quorum Blockchain Explorer](https://github.com/blk-io/blk-explorer-free) - a Blockchain Explorer for Quorum which supports viewing private transactions
* [Quorum-Genesis](https://github.com/davebryson/quorum-genesis) - A simple CL utility for Quorum to help populate the genesis file with voters and makers
* [Quorum Maker](https://github.com/synechron-finlabs/quorum-maker/) - a utility to create Quorum nodes
* [QuorumNetworkManager](https://github.com/ConsenSys/QuorumNetworkManager) - makes creating & managing Quorum networks easy
* [ERC20 REST service](https://github.com/blk-io/erc20-rest-service) - a Quorum-supported RESTful service for creating and managing ERC-20 tokens
* [Nethereum Quorum](https://github.com/Nethereum/Nethereum/tree/master/src/Nethereum.Quorum) - a .NET Quorum adapter
* [web3j-quorum](https://github.com/web3j/quorum) - an extension to the web3j Java library providing support for the Quorum API

## Contributing

Thank you for your interest in contributing to Quorum!

Quorum is built on open source and we invite you to contribute enhancements. Upon review you will be required to complete a Contributor License Agreement (CLA) before we are able to merge. If you have any questions about the contribution process, please feel free to send an email to [quorum_info@jpmorgan.com](mailto:quorum_info@jpmorgan.com).

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.
