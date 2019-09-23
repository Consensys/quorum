## What is it?

[Cakeshop](https://github.com/jpmorganchase/cakeshop) is a set of tools and APIs for working with [Ethereum](https://ethereum.org/)-like ledgers, packaged as a Java web application archive (WAR) that gets you up and running in under 60 seconds.

Cakeshop can either start up a geth node, which you can then interact with using the Cakeshop front-end, or it can be connected to an Ethereum-like node, such as Quorum, that you already have running. A given Cakeshop instance connects with one node on the blockchain network you connect to.

![image](console.png)

Out of the box you get:


* **Node Management** - Fully functioning Ethereum node (via geth), setting up a cluster
* **Blockchain Explorer** - view transactions, blocks and contracts, and see historical contract state at a point in time
* **Admin Console** - start & stop nodes, create a cluster and view the overall status of your network
* **Peer Management** - easily discover, add and remove peers
* **Solidity Sandbox** - develop, compile, deploy and interact with Solidity smart contracts

It provides tools for managing a local blockchain node, setting up clusters,
exploring the state of the chain, and working with contracts.

The Cakeshop package includes the [tessera](https://github.com/jpmorganchase/tessera) and [constellation](https://github.com/jpmorganchase/constellation) transaction managers, a [Solidity](https://solidity.readthedocs.org/en/latest/) compiler, and all dependencies. Cakeshop will download the latest version of [quorum](https://github.com/jpmorganchase/quorum) and bootnode from [geth](https://github.com/ethereum/go-ethereum) (to use a different version, see [here](https://github.com/jpmorganchase/cakeshop/blob/master/docs/configuration.md#custom-quorum-binaries))
