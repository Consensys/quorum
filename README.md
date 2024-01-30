# <img src="https://raw.githubusercontent.com/consensys/quorum/master/logo.png" width="200" height="35"/>


![Build Check](https://github.com/jpmorganchase/quorum/workflows/Build%20Check/badge.svg?branch=master)
[![Docker Pulls](https://img.shields.io/docker/pulls/quorumengineering/quorum)](https://hub.docker.com/r/quorumengineering/quorum)
[![Discord](https://img.shields.io/discord/697535391594446898)](https://discord.com/channels/697535391594446898/747810572937986240)


GoQuorum is an Ethereum-based distributed ledger protocol with transaction/contract privacy and new consensus mechanisms.

GoQuorum is a fork of [go-ethereum](https://github.com/ethereum/go-ethereum) and is updated in line with go-ethereum releases.

Key enhancements over go-ethereum:

* [__Privacy__](https://consensys.net/docs/goquorum//en/latest/concepts/privacy/) - GoQuorum supports private transactions and private contracts through public/private state separation, and utilises peer-to-peer encrypted message exchanges (see [Tessera](https://github.com/consensys/tessera)) for directed transfer of private data to network participants
* [__Alternative Consensus Mechanisms__](https://consensys.net/docs/goquorum//en/latest/concepts/consensus/overview/) - with no need for POW/POS in a permissioned network, GoQuorum instead offers multiple consensus mechanisms that are more appropriate for consortium chains:
    * [__QBFT__](https://consensys.net/docs/goquorum/en/latest/configure-and-manage/configure/consensus-protocols/qbft/) - Improved version of IBFT that is interoperable with Hyperledger Besu
    * [__Istanbul BFT__](https://consensys.net/docs/goquorum/en/latest/configure-and-manage/configure/consensus-protocols/ibft/) - a PBFT-inspired consensus algorithm with transaction finality, by AMIS.
    * [__Clique POA Consensus__](https://github.com/ethereum/EIPs/issues/225) - a default POA consensus algorithm bundled with Go Ethereum.
    * [__Raft-based Consensus__](https://consensys.net/docs/goquorum/en/latest/configure-and-manage/configure/consensus-protocols/raft/) - a consensus model for faster blocktimes, transaction finality, and on-demand block creation
* [__Peer Permissioning__](https://consensys.net/docs/goquorum/en/latest/concepts/permissions-overview/) - node/peer permissioning, ensuring only known parties can join the network
* [__Account Management__](https://consensys.net/docs/goquorum/en/latest/concepts/account-management/) - GoQuorum introduced account plugins, which allows GoQuorum or clef to be extended with alternative methods of managing accounts including external vaults.
* [__Pluggable Architecture__](https://consensys.net/docs/goquorum/en/latest/concepts/plugins/) -  allows adding additional features as plugins to the core `geth`, providing extensibility, flexibility, and distinct isolation of GoQuorum features.
* __Higher Performance__ - GoQuorum offers significantly higher performance throughput than public geth

## Architecture

![GoQuorum Tessera Privacy Flow](https://github.com/consensys/quorum/blob/master/docs/Quorum%20Design.png)

The above diagram is very high-level overview of component architecture used by GoQuorum. For more in-depth discussion of the components and how they interact, please refer to [lifecycle of a private transaction](https://consensys.net/docs/goquorum/en/latest/concepts/privacy/private-transaction-lifecycle/).

## Quickstart
The easiest way to get started is to use * [quorum-dev-quickstart](https://consensys.net/docs/goquorum/en/latest/tutorials/quorum-dev-quickstart/using-the-quickstart/) - a command line tool that allows users to set up a development GoQuorum network on their local machine in less than *2 minutes*.

## GoQuorum Projects

Check out some of the interesting projects we are actively working on:

* [quorum-remix-plugin](https://consensys.net/docs/goquorum/en/latest/tutorials/quorum-dev-quickstart/remix/): The GoQuorum plugin for Ethereum's Remix IDE adds support for creating and interacting with private contracts on a GoQuorum network.
* [Cakeshop](https://consensys.net/docs/goquorum/en/latest/configure-and-manage/monitor/cakeshop/): An integrated development environment and SDK for GoQuorum
* [quorum-examples](https://github.com/ConsenSys/quorum-examples): GoQuorum demonstration examples
* <img src="docs/images/qubernetes/k8s-logo.png" width="15"/> [Quorum-Kubernetes](https://consensys.net/docs/goquorum/en/latest/deploy/install/kubernetes/): Deploy GoQuorum on Kubernetes
* [we3js-quorum](https://consensys.net/docs/goquorum/en/latest/reference/web3js-quorum/): Extends web3.js to support GoQuorum and Hyperledger Besu specific APIs
* Zero Knowledge on GoQuorum
   * [ZSL on GoQuorum](https://github.com/ConsenSys/zsl-q/)
   * [Anonymous Zether](https://github.com/ConsenSys/anonymous-zether)



## Official Docker Containers
The official docker containers can be found under https://hub.docker.com/u/quorumengineering/

## Third Party Tools/Libraries

The following GoQuorum-related libraries/applications have been created by Third Parties and as such are not specifically endorsed by J.P. Morgan.  A big thanks to the developers for improving the tooling around GoQuorum!

* [Chainlens Blockchain Explorer](https://github.com/web3labs/chainlens-free) - a Blockchain Explorer for GoQuorum which supports viewing private transactions
* [Quorum-Genesis](https://github.com/davebryson/quorum-genesis) - A simple CL utility for GoQuorum to help populate the genesis file with voters and makers
* [Quorum Maker](https://github.com/synechron-finlabs/quorum-maker/) - a utility to create GoQuorum nodes
* [ERC20 REST service](https://github.com/web3labs/erc20-rest-service) - a GoQuorum-supported RESTful service for creating and managing ERC-20 tokens
* [Nethereum Quorum](https://github.com/Nethereum/Nethereum/tree/master/src/Nethereum.Quorum) - a .NET GoQuorum adapter
* [web3j-quorum](https://github.com/web3j/web3j-quorum) - an extension to the web3j Java library providing support for the GoQuorum API
* [Apache Camel](http://github.com/apache/camel) - an Apache Camel component providing support for the GoQuorum API using web3j library. Here is the article describing how to use Apache Camel with Ethereum and GoQuorum https://medium.com/@bibryam/enterprise-integration-for-ethereum-fa67a1577d43

## Contributing
GoQuorum is built on open source and we invite you to contribute enhancements. Upon review you will be required to complete a Contributor License Agreement (CLA) before we are able to merge. If you have any questions about the contribution process, please feel free to send an email to [info@goquorum.com](mailto:info@goquorum.com). Please see the [Contributors guide](.github/CONTRIBUTING.md) for more information about the process.

## Reporting Security Bugs
Security is part of our commitment to our users. At GoQuorum we have a close relationship with the security community, we understand the realm, and encourage security researchers to become part of our mission of building secure reliable software. This section explains how to submit security bugs, and what to expect in return.

All security bugs in [GoQuorum](https://github.com/consensys/quorum) and its ecosystem ([Tessera](https://github.com/consensys/tessera), [Cakeshop](https://github.com/consensys/cakeshop), ..etc)  should be reported by email to [security-quorum@consensys.net](mailto:security-quorum@consensys.net). Please use the prefix **[security]** in your subject. This email is delivered to GoQuorum security team. Your email will be acknowledged, and you'll receive a more detailed response to your email as soon as possible indicating the next steps in handling your report. After the initial reply to your report, the security team will endeavor to keep you informed of the progress being made towards a fix and full announcement.

If you have not received a reply to your email or you have not heard from the security team please contact any team member through GoQuorum slack security channel. **Please note that GoQuorum discord channels are public discussion forum**. When escalating to this medium, please do not disclose the details of the issue. Simply state that you're trying to reach a member of the security team.

#### Responsible Disclosure Process
GoQuorum project uses the following responsible disclosure process:

- Once the security report is received it is assigned a primary handler. This person coordinates the fix and release process.
- The issue is confirmed and a list of affected software is determined.
- Code is audited to find any potential similar problems.
- If it is determined, in consultation with the submitter, that a CVE-ID is required, the primary handler will trigger the process.
- Fixes are applied to the public repository and a new release is issued.
- On the date that the fixes are applied, announcements are sent to Quorum-announce.
- At this point you would be able to disclose publicly your finding.

**Note:** This process can take some time. Every effort will be made to handle the security bug in as timely a manner as possible, however it's important that we follow the process described above to ensure that disclosures are handled consistently.

#### Receiving Security Updates
The best way to receive security announcements is to subscribe to the Quorum-announce mailing list/channel. Any messages pertaining to a security issue will be prefixed with **[security]**.

Comments on This Policy
If you have any suggestions to improve this policy, please send an email to info@goquorum.com for discussion.

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.

Any project planning to use the `crypto/secp256k1` sub-module must use the specific [secp256k1 standalone library](https://github.com/ConsenSys/goquorum-crypto-secp256k1) licensed under 3-clause BSD.
