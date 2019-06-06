# <img src="https://raw.githubusercontent.com/jpmorganchase/quorum/master/logo.png" width="100" height="100"/>

<a href="https://clh7rniov2.execute-api.us-east-1.amazonaws.com/Express/" target="_blank" rel="noopener"><img title="Quorum Slack" src="https://clh7rniov2.execute-api.us-east-1.amazonaws.com/Express/badge.svg" alt="Quorum Slack" /></a>
[![Build Status](https://travis-ci.org/jpmorganchase/quorum.svg?branch=master)](https://travis-ci.org/jpmorganchase/quorum)
[![Download](https://api.bintray.com/packages/quorumengineering/quorum/geth/images/download.svg)](https://bintray.com/quorumengineering/quorum/geth/_latestVersion)

Quorum is an Ethereum-based distributed ledger protocol with transaction/contract privacy and new consensus mechanisms.

Quorum is a fork of [go-ethereum](https://github.com/ethereum/go-ethereum) and is updated in line with go-ethereum releases.

Key enhancements over go-ethereum:

* __Privacy__ - Quorum supports private transactions and private contracts through public/private state separation, and utilises peer-to-peer encrypted message exchanges (see [Constellation](https://github.com/jpmorganchase/constellation) and [Tessera](https://github.com/jpmorganchase/tessera)) for directed transfer of private data to network participants
* __Alternative Consensus Mechanisms__ - with no need for POW/POS in a permissioned network, Quorum instead offers multiple consensus mechanisms that are more appropriate for consortium chains:
    * __Raft-based Consensus__ - a consensus model for faster blocktimes, transaction finality, and on-demand block creation
    * __Istanbul BFT__ - a PBFT-inspired consensus algorithm with transaction finality, by AMIS.
* __Peer Permissioning__ - node/peer permissioning using smart contracts, ensuring only known parties can join the network
* __Higher Performance__ - Quorum offers significantly higher performance than public geth

## Architecture

![Quorum Tessera Privacy Flow](https://raw.githubusercontent.com/jpmorganchase/quorum-docs/master/images/QuorumTransactionProcessing.JPG)

The above diagram is a high-level overview of the privacy architecture used by Quorum. For more in-depth discussion of the components, refer to the [wiki](https://github.com/jpmorganchase/quorum/wiki) pages.

## Quickstart
The quickest way to get started with Quorum is by following instructions in the [Quorum Examples](https://github.com/jpmorganchase/quorum-examples) repository.  This allows you to quickly create a network of Quorum nodes, and includes a step-by-step demonstration of the privacy features of Quorum.

## Further Reading
Further documentation can be found in the [docs](docs/) folder and on the [wiki](https://github.com/jpmorganchase/quorum/wiki).

## Official Docker Containers
The official docker containers can be found under https://hub.docker.com/u/quorumengineering/ 


## See also

* [Quorum](https://github.com/jpmorganchase/quorum): this repository
* [Quorum Wiki](https://github.com/jpmorganchase/quorum/wiki)
* [quorum-examples](https://github.com/jpmorganchase/quorum-examples): Quorum demonstration examples
* [Quorum Community Slack Inviter](https://clh7rniov2.execute-api.us-east-1.amazonaws.com/Express/): Quorum Slack community entry point
* Quorum Transaction Managers
   * [Constellation](https://github.com/jpmorganchase/constellation): Haskell implementation of peer-to-peer encrypted message exchange for transaction privacy
   * [Tessera](https://github.com/jpmorganchase/tessera): Java implementation of peer-to-peer encrypted message exchange for transaction privacy
* Quorum supported consensuses
   * [Raft Consensus Documentation](docs/raft.md)
   * [Istanbul BFT Consensus Documentation](https://github.com/ethereum/EIPs/issues/650): [RPC API](https://github.com/jpmorganchase/quorum/blob/master/docs/istanbul-rpc-api.md) and [technical article](https://medium.com/getamis/istanbul-bft-ibft-c2758b7fe6ff). <span style="background-color: #ffffbf">Please note</span> that updated istanbul-tools is now hosted in [this](https://github.com/jpmorganchase/istanbul-tools/) repository
   * [Clique POA Consensus Documentation](https://github.com/ethereum/EIPs/issues/225) and a [guide to setup clique json](https://modalduality.org/posts/puppeth/) with [puppeth](https://blog.ethereum.org/2017/04/14/geth-1-6-puppeth-master/)
* [ZSL](https://github.com/jpmorganchase/quorum/wiki/ZSL) wiki page and [documentation](https://github.com/jpmorganchase/zsl-q/blob/master/README.md)
* [quorum-tools](https://github.com/jpmorganchase/quorum-tools): local cluster orchestration, and integration testing tool

## Third Party Tools/Libraries

The following Quorum-related libraries/applications have been created by Third Parties and as such are not specifically endorsed by J.P. Morgan.  A big thanks to the developers for improving the tooling around Quorum!

* [Quorum Blockchain Explorer](https://github.com/blk-io/epirus-free) - a Blockchain Explorer for Quorum which supports viewing private transactions
* [Quorum-Genesis](https://github.com/davebryson/quorum-genesis) - A simple CL utility for Quorum to help populate the genesis file with voters and makers
* [Quorum Maker](https://github.com/synechron-finlabs/quorum-maker/) - a utility to create Quorum nodes
* [QuorumNetworkManager](https://github.com/ConsenSys/QuorumNetworkManager) - makes creating & managing Quorum networks easy
* [ERC20 REST service](https://github.com/blk-io/erc20-rest-service) - a Quorum-supported RESTful service for creating and managing ERC-20 tokens
* [Nethereum Quorum](https://github.com/Nethereum/Nethereum/tree/master/src/Nethereum.Quorum) - a .NET Quorum adapter
* [web3j-quorum](https://github.com/web3j/quorum) - an extension to the web3j Java library providing support for the Quorum API
* [Apache Camel](http://github.com/apache/camel) - an Apache Camel component providing support for the Quorum API using web3j library. Here is the artcile describing how to use Apache Camel with Ethereum and Quorum https://medium.com/@bibryam/enterprise-integration-for-ethereum-fa67a1577d43


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
