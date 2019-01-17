# Overview
At present when Quorum is brought `--permissionoed` mode, for establishing p2p connectivity, the nodes checked `permissioned-nodes.json` file. If the enode id of the node requesting connection is present in this file, the connection is established else it is declined. The `permissioned-nodes.json` file is updated procedurally whenever a new node joins the network. Node permissioning feature will allow the existing nodes to propose a new node to join the network and once majority voting is done on the proposal, it will update the `permissioned-nodes.json` automatically. Further the existing nodes can propose any node for deactivation, blacklisting and activating back from a deactivated status.

Account permissioning feature will introduce the following access controls at account level:
* Read Only: Accounts with this access will be able to perform only read activities and will not be able to deploy contracts or transactions
* Transact: Accounts with transact access will be able to commit transactions but will not be able to deploy contracts
* Contract Deploy: Accounts with this access will be able to deploy contracts and commit transactions
* Full Access: Similar to "Contract Deploy" access, accounts with this access will be able to deploy contracts and perform transactions
Currently there is not any differences in the access types "Full Access" and "Contract Deploy". 

It should be noted that both the above features will be available when Quorum geth is brought in `--permissioned` mode.

## Set up
Node permissioning and Account access control is managed by a smart contract [Permission.sol](../controls/permission/Permission.sol). This is deployed as precompiled contract at the time of initial network bootup. The precompiled contract is deployed at address `0x000000000000000000032`. 
## Node Permissioning 

## Account Access Control
