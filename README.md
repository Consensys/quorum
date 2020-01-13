# <img src="https://raw.githubusercontent.com/jpmorganchase/quorum/master/logo.png" width="200" height="200"/>

<a href="https://bit.ly/quorum-slack" target="_blank" rel="noopener"><img title="Quorum Slack" src="https://clh7rniov2.execute-api.us-east-1.amazonaws.com/Express/badge.svg" alt="Quorum Slack" /></a>
[![Build Status](https://travis-ci.org/jpmorganchase/quorum.svg?branch=master)](https://travis-ci.org/jpmorganchase/quorum)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/quorumengineering/quorum)](https://hub.docker.com/r/quorumengineering/quorum/builds)
[![Documentation Status](https://readthedocs.org/projects/goquorum/badge/?version=latest)](http://docs.goquorum.com/en/latest/?badge=latest)
[![Download](https://api.bintray.com/packages/quorumengineering/quorum/geth/images/download.svg)](https://bintray.com/quorumengineering/quorum/geth/_latestVersion)
[![Docker Pulls](https://img.shields.io/docker/pulls/quorumengineering/quorum)](https://hub.docker.com/r/quorumengineering/quorum)

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

The above diagram is a high-level overview of the privacy architecture used by Quorum. For more in-depth discussion of the components, refer to the [documentation site](https://docs.goquorum.com).

## Quickstart
The quickest way to get started with Quorum is by following instructions in the [Quorum Examples](https://github.com/jpmorganchase/quorum-examples) repository.  This allows you to quickly create a network of Quorum nodes, and includes a step-by-step demonstration of the privacy features of Quorum.

## Further Reading
Further documentation can be found in the [docs](docs/) folder and on the [documentation site](https://docs.goquorum.com).

## Official Docker Containers
The official docker containers can be found under https://hub.docker.com/u/quorumengineering/ 


## See also

* [Quorum](https://github.com/jpmorganchase/quorum): this repository
* [Quorum Documentation](https://docs.goquorum.com)
* [quorum-examples](https://github.com/jpmorganchase/quorum-examples): Quorum demonstration examples
* [Quorum Community Slack Inviter](https://bit.ly/quorum-slack): Quorum Slack community entry point
* Quorum Transaction Managers
   * [Constellation](https://github.com/jpmorganchase/constellation): Haskell implementation of peer-to-peer encrypted message exchange for transaction privacy
   * [Tessera](https://github.com/jpmorganchase/tessera): Java implementation of peer-to-peer encrypted message exchange for transaction privacy
* Quorum supported consensuses
   * [Raft Consensus Documentation](https://docs.goquorum.com/en/latest/Consensus/raft/)
   * [Istanbul BFT Consensus Documentation](https://github.com/ethereum/EIPs/issues/650): [RPC API](https://docs.goquorum.com/en/latest/Consensus/ibft/istanbul-rpc-api.md) and [technical article](https://medium.com/getamis/istanbul-bft-ibft-c2758b7fe6ff). __Please note__ that updated istanbul-tools is now hosted in [this](https://github.com/jpmorganchase/istanbul-tools/) repository
   * [Clique POA Consensus Documentation](https://github.com/ethereum/EIPs/issues/225) and a [guide to setup clique json](https://modalduality.org/posts/puppeth/) with [puppeth](https://blog.ethereum.org/2017/04/14/geth-1-6-puppeth-master/)
* Zero Knowledge on Quorum
   * [ZSL](https://github.com/jpmorganchase/quorum/wiki/ZSL) wiki page and [documentation](https://github.com/jpmorganchase/zsl-q/blob/master/README.md)
   * [Anonymous Zether](https://github.com/jpmorganchase/anonymous-zether) implementation
* [quorum-cloud](https://github.com/jpmorganchase/quorum-cloud): Tools to help deploy Quorum network in a cloud provider of choice
* [Cakeshop](https://github.com/jpmorganchase/cakeshop): An integrated development environment and SDK for Quorum

## Third Party Tools/Libraries

The following Quorum-related libraries/applications have been created by Third Parties and as such are not specifically endorsed by J.P. Morgan.  A big thanks to the developers for improving the tooling around Quorum!

* [Quorum Blockchain Explorer](https://github.com/blk-io/epirus-free) - a Blockchain Explorer for Quorum which supports viewing private transactions
* [Quorum-Genesis](https://github.com/davebryson/quorum-genesis) - A simple CL utility for Quorum to help populate the genesis file with voters and makers
* [Quorum Maker](https://github.com/synechron-finlabs/quorum-maker/) - a utility to create Quorum nodes
* [QuorumNetworkManager](https://github.com/ConsenSys/QuorumNetworkManager) - makes creating & managing Quorum networks easy
* [ERC20 REST service](https://github.com/blk-io/erc20-rest-service) - a Quorum-supported RESTful service for creating and managing ERC-20 tokens
* [Nethereum Quorum](https://github.com/Nethereum/Nethereum/tree/master/src/Nethereum.Quorum) - a .NET Quorum adapter
* [web3j-quorum](https://github.com/web3j/web3j-quorum) - an extension to the web3j Java library providing support for the Quorum API
* [Apache Camel](http://github.com/apache/camel) - an Apache Camel component providing support for the Quorum API using web3j library. Here is the artcile describing how to use Apache Camel with Ethereum and Quorum https://medium.com/@bibryam/enterprise-integration-for-ethereum-fa67a1577d43

## Contributing
Quorum is built on open source and we invite you to contribute enhancements. Upon review you will be required to complete a Contributor License Agreement (CLA) before we are able to merge. If you have any questions about the contribution process, please feel free to send an email to [info@goquorum.com](mailto:info@goquorum.com).

## Reporting Security Bugs
Security is part of our commitment to our users. At Quorum we have a close relationship with the security community, we understand the realm, and encourage security researchers to become part of our mission of building secure reliable software. This section explains how to submit security bugs, and what to expect in return.

All security bugs in [Quorum](https://github.com/jpmorganchase/quorum) and its ecosystem ([Tessera](https://github.com/jpmorganchase/tessera), [Constellation](https://github.com/jpmorganchase/constellation), [Cakeshop](https://github.com/jpmorganchase/cakeshop), ..etc)  should be reported by email to [info@goquorum.com](mailto:info@goquorum.com). Please use the prefix **[security]** in your subject. This email is delivered to Quorum security team. Your email will be acknowledged, and you'll receive a more detailed response to your email as soon as possible indicating the next steps in handling your report. After the initial reply to your report, the security team will endeavor to keep you informed of the progress being made towards a fix and full announcement.

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
