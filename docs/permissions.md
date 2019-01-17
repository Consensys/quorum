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
Node permissioning and Account access control is managed by a smart contract [Permission.sol](../controls/permission/Permission.sol). This is deployed as precompiled contract at the time of initial network bootup. The precompiled contract is deployed at address `0x000000000000000000032`. The binding of the precompiled byte code with the address is in `genesis.json`. The initial set of account which will have full access when the network is up, should given as a part of `genesis.json` as shown below:
```
{
  "alloc": {
    "0x0000000000000000000000000000000000000020": {
    "code": "<<compiled contract bytecode>>"
    "storage": {
      "0x0000000000000000000000000000000000000000000000000000000000000000" : "0x0000000000000000000000000000000000000003",
      "0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563" : "0xcA843569e3427144cEad5e4d5999a3D0cCF92B8e",
      "0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e564" : "0xed9d02e382b34818e88b88a309c7fe71e65f419d",
      "0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e565" : "0x0fbdc686b912d7722dc86510934589e0aaf3b55a"
    },
    "balance": "1000000000000000000000000000"
  },
```
In the above set up, accounts "0xcA843569e3427144cEad5e4d5999a3D0cCF92B8e", "0xcA843569e3427144cEad5e4d5999a3D0cCF92B8e", "0xcA843569e3427144cEad5e4d5999a3D0cCF92B8e" will have full access when the network boot is completed. 

Further, if the initial set of network nodes are brought up with `--permissioned` mode and a new approved node is joining the networl without `--permissioned` flag in the `geth` start commmand, system will not allow the new `geth` node to come and a fatal error `Joining a permissioned network in non-permissioned mode. Bring up geth with --permissioned."`.
## Node Permissioning 

## Account Access Control
