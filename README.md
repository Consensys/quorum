# <img src="https://raw.githubusercontent.com/jpmorganchase/quorum/master/logo.png" width="200" height="200"/>

<a href="https://www.goquorum.com/slack-inviter" target="_blank" rel="noopener"><img title="Quorum Slack" src="https://93ecjxb0d3.execute-api.us-east-1.amazonaws.com/Express/badge.svg" alt="Quorum Slack" /></a>
![Build Check](https://github.com/jpmorganchase/quorum/workflows/Build%20Check/badge.svg?branch=master)
[![Download](https://api.bintray.com/packages/quorumengineering/quorum/geth/images/download.svg)](https://bintray.com/quorumengineering/quorum/geth/_latestVersion)
[![Docker Pulls](https://img.shields.io/docker/pulls/quorumengineering/quorum)](https://hub.docker.com/r/quorumengineering/quorum)

Quorum is an Ethereum-based distributed ledger protocol with transaction/contract privacy and new consensus mechanisms.

Quorum is a fork of [go-ethereum](https://github.com/ethereum/go-ethereum) and is updated in line with go-ethereum releases.

Key enhancements over go-ethereum:

* [__Privacy__](http://docs.goquorum.com/en/latest/Privacy/Overview/) - Quorum supports private transactions and private contracts through public/private state separation, and utilises peer-to-peer encrypted message exchanges (see [Constellation](https://github.com/consensys/constellation) and [Tessera](https://github.com/consensys/tessera)) for directed transfer of private data to network participants
* [__Alternative Consensus Mechanisms__](http://docs.goquorum.com/en/latest/Consensus/Consensus/) - with no need for POW/POS in a permissioned network, Quorum instead offers multiple consensus mechanisms that are more appropriate for consortium chains:
    * [__Raft-based Consensus__](http://docs.goquorum.com/en/latest/Consensus/raft/raft/) - a consensus model for faster blocktimes, transaction finality, and on-demand block creation
    * [__Istanbul BFT__](http://docs.goquorum.com/en/latest/Consensus/ibft/ibft/) - a PBFT-inspired consensus algorithm with transaction finality, by AMIS.
    * [__Clique POA Consensus__](https://github.com/ethereum/EIPs/issues/225) - a default POA consensus algorithm bundled with Go Ethereum.
* [__Peer Permissioning__](http://docs.goquorum.com/en/latest/Permissioning/Permissions%20Overview/) - node/peer permissioning, ensuring only known parties can join the network
* [__Account Management__](http://docs.goquorum.com/en/latest/Account-Key-Management/Overview/) - Quorum introduced account plugins, which allows Quorum or clef to be extended with alternative methods of managing accounts including external vaults.
* [__Pluggable Architecture__](http://docs.goquorum.com/en/latest/PluggableArchitecture/Overview/) -  allows adding additional features as plugins to the core `geth`, providing extensibility, flexibility, and distinct isolation of Quorum features.
* __Higher Performance__ - Quorum offers significantly higher performance throughput than public geth

## Architecture

![Quorum Tessera Privacy Flow](https://github.com/consensys/quorum/blob/master/docs/Quorum%20Design.png)

The above diagram is very high-level overview of component architecture used by Quorum. For more in-depth discussion of the components and how they interact, please refer to [lifecycle of a private transaction](http://docs.goquorum.com/en/latest/Privacy/Lifecycle-of-a-private-transaction/).

## Quickstart
There are [several ways](https://docs.goquorum.com/en/latest/Getting%20Started/Getting%20Started%20Overview/) to quickly get up and running with Quorum.  One of the easiest is to use [Quorum Wizard](https://docs.goquorum.com/en/latest/Getting%20Started/Getting%20Started%20Overview/#quickstart-with-quorum-wizard) - a command line tool that allows users to set up a development Quorum network on their local machine in less than *2 minutes*.

## Quorum Projects

Check out some of the interesting projects we are actively working on:

* [quorum-wizard](http://docs.goquorum.com/en/latest/Wizard/GettingStarted/): Setup a Quorum network in 2 minutes!
* [quorum-remix-plugin](http://docs.goquorum.com/en/latest/RemixPlugin/Overview/): The Quorum plugin for Ethereum's Remix IDE adds support for creating and interacting with private contracts on a Quorum network.
* [Cakeshop](http://docs.goquorum.com/en/latest/Cakeshop/Overview/): An integrated development environment and SDK for Quorum
* [Quorum-Profiling](http://docs.goquorum.com/en/latest/Quorum%20Profiling/Overview/): Toolset for stress testing & benchmarking Quorum networks. 
* [quorum-examples](http://docs.goquorum.com/en/latest/Getting%20Started/Quorum-Examples/): Quorum demonstration examples
* <img src="docs/images/qubernetes/k8s-logo.png" width="15"/> [qubernetes](http://docs.goquorum.com/en/latest/Getting%20Started/Getting%20Started%20Overview/#quorum-on-kubernetes): Deploy Quorum on Kubernetes
* [quorum-cloud](http://docs.goquorum.com/en/latest/Getting%20Started/Getting%20Started%20Overview/#creating-a-network-deployed-in-the-cloud): Tools to help deploy Quorum network in a cloud provider of choice
* [quorum.js](http://docs.goquorum.com/en/latest/quorum.js/Overview/): Extends web3.js to support Quorum-specific APIs
* Zero Knowledge on Quorum
   * [ZSL](https://docs.goquorum.consensys.net/en/latest/Reference/GoQuorum-Projects/#zsl-proof-of-concept) POC and [ZSL on Quorum](https://github.com/ConsenSys/zsl-q/blob/master/README.md)
   * [Anonymous Zether](https://github.com/ConsenSys/anonymous-zether) implementation



## Official Docker Containers
The official docker containers can be found under https://hub.docker.com/u/quorumengineering/

## Third Party Tools/Libraries

The following Quorum-related libraries/applications have been created by Third Parties and as such are not specifically endorsed by J.P. Morgan.  A big thanks to the developers for improving the tooling around Quorum!

* [Quorum Blockchain Explorer](https://github.com/web3labs/epirus-free) - a Blockchain Explorer for Quorum which supports viewing private transactions
* [Quorum-Genesis](https://github.com/davebryson/quorum-genesis) - A simple CL utility for Quorum to help populate the genesis file with voters and makers
* [Quorum Maker](https://github.com/synechron-finlabs/quorum-maker/) - a utility to create Quorum nodes
* [QuorumNetworkManager](https://github.com/ConsenSys/QuorumNetworkManager) - makes creating & managing Quorum networks easy
* [ERC20 REST service](https://github.com/web3labs/erc20-rest-service) - a Quorum-supported RESTful service for creating and managing ERC-20 tokens
* [Nethereum Quorum](https://github.com/Nethereum/Nethereum/tree/master/src/Nethereum.Quorum) - a .NET Quorum adapter
* [web3j-quorum](https://github.com/web3j/web3j-quorum) - an extension to the web3j Java library providing support for the Quorum API
* [Apache Camel](http://github.com/apache/camel) - an Apache Camel component providing support for the Quorum API using web3j library. Here is the artcile describing how to use Apache Camel with Ethereum and Quorum https://medium.com/@bibryam/enterprise-integration-for-ethereum-fa67a1577d43

## Contributing
Quorum is built on open source and we invite you to contribute enhancements. Upon review you will be required to complete a Contributor License Agreement (CLA) before we are able to merge. If you have any questions about the contribution process, please feel free to send an email to [info@goquorum.com](mailto:info@goquorum.com). Please see the [Contributors guide](.github/CONTRIBUTING.md) for more information about the process.

## Reporting Security Bugs
Security is part of our commitment to our users. At Quorum we have a close relationship with the security community, we understand the realm, and encourage security researchers to become part of our mission of building secure reliable software. This section explains how to submit security bugs, and what to expect in return.

All security bugs in [Quorum](https://github.com/consensys/quorum) and its ecosystem ([Tessera](https://github.com/consensys/tessera), [Constellation](https://github.com/consensys/constellation), [Cakeshop](https://github.com/consensys/cakeshop), ..etc)  should be reported by email to [security-quorum@consensys.net](mailto:security-quorum@consensys.net). Please use the prefix **[security]** in your subject. This email is delivered to Quorum security team. Your email will be acknowledged, and you'll receive a more detailed response to your email as soon as possible indicating the next steps in handling your report. After the initial reply to your report, the security team will endeavor to keep you informed of the progress being made towards a fix and full announcement.

If you have not received a reply to your email or you have not heard from the security team please contact any team member through quorum slack security channel. **Please note that Quorum slack channels are public discussion forum**. When escalating to this medium, please do not disclose the details of the issue. Simply state that you're trying to reach a member of the security team.

#### Responsible Disclosure Process
Quorum project uses the following responsible disclosure process:

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
