
With no need for POW/POS in a permissioned network, Quorum instead offers multiple consensus mechanisms that are more appropriate for consortium chains:  


* __Raft-based Consensus__: A consensus model for faster blocktimes, transaction finality, and on-demand block creation.  See [Raft-based consensus for Ethereum/Quorum](../raft) for more information 


* __Istanbul BFT (Byzantine Fault Tolerance) Consensus__: A PBFT-inspired consensus algorithm with transaction finality, by AMIS.  See [Istanbul BFT Consensus documentation](https://github.com/ethereum/EIPs/issues/650), the [RPC API](../istanbul-rpc-api), and this [technical web article](https://medium.com/getamis/istanbul-bft-ibft-c2758b7fe6ff) for more information


* __Clique POA Consensus__: a default POA consensus algorithm bundled with Go Ethereum.  See [Clique POA Consensus Documentation](https://github.com/ethereum/EIPs/issues/225) and a [guide to setup clique json](https://hackernoon.com/hands-on-creating-your-own-local-private-geth-node-beginner-friendly-3d45902cc612) with [puppeth](https://blog.ethereum.org/2017/04/14/geth-1-6-puppeth-master/)
