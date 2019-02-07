# Overview
Currently when Quorum geth is brought up in `--permissioned` mode, for establishing p2p connectivity the nodes check `permissioned-nodes.json` file. If the enode id of the node requesting connection is present in this file, the p2p connection is established else it is declined. The `permissioned-nodes.json` file is updated procedurally whenever a new node joins the network. Node permissioning feature will allow the existing nodes to propose a new node to join the network and once majority voting is done on the proposal, it will update the `permissioned-nodes.json` automatically. Further the existing nodes can propose any node for deactivation, blacklisting and activating a node back from a deactivated status.

Account permissioning feature introduces controls at account level. This will control the type of activities that a particular account can perform in the network.

It should be noted that both the above features will be available when Quorum geth is brought in `--permissioned` mode with the set up as described in the next section. 

## Set up
Node permissioning and account access control is managed by a smart contract [Permission.sol](../controls/permission/Permission.sol). The precompiled bytecode of the smart contract is deployed at address `0x000000000000000000020` in network bootup process. The binding of the precompiled byte code with the address is in `genesis.json`. The initial set of account that will have full access when the network is up, should given as a part of `genesis.json` as shown below:
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
In the above set up, accounts `"0xcA843569e3427144cEad5e4d5999a3D0cCF92B8e", "0xed9d02e382b34818e88b88a309c7fe71e65f419d", "0x0fbdc686b912d7722dc86510934589e0aaf3b55a"` will have full access when the network boot is completed. The default access for any other account in the network will be `ReadOnly`

If the network is brought up with permissions control byte code and no accounts are given as a part of storage, then geth start up will fail with error `Permissioned network being brought up with zero accounts having full access. Add permissioned full access accounts in genesis.json and bring up the network`

Further, if the initial set of network nodes are brought up with `--permissioned` mode and a new approved node is joining the network without `--permissioned` flag in the `geth` start commmand, system will not allow the new `geth` node to come and a fatal error `Joining a permissioned network in non-permissioned mode. Bring up geth with --permissioned."` will be thrown.

## Node Permissioning 
In a permissioned network any node can be in one of the following status:
* Proposed - The node has been proposed to join the network and pending majority voting to be marked as `Approved`
* Approved - The node is approved and is part of the network
* PendingDeactivation - The node has been proposed for deatvation from network and is pending majority approval
* Deactivated - The node has been deactivated from the network. Any node in this status will be disconnected from the network and block sync with this node will not happen
* PendingActivation - A deactivated node has been proposed for activation and is pending majority approval. Once approved the node will move to `Approved` status
* PendingBlacklisting - The node has been proposed for blacklisting and is pending majority approval
* Blacklisted - The node has been blacklisted. If the node was an active node on the network, the node will be disconnected from the network and block sync will stop

It should be noted that deactivation is temporary in nature and hence the node can join back the network at a later point in time. However blacklisting is permanent in nature. A blacklisted node will never be able to join back the network. Further blacklisting can be proposed for a node which is in the network or a new node which is currently not part of the network. 

When the network is started for the first time, all the nodes present in `static-nodes.json` file are added to the permissioned nodes list maintained in the smart contract. Once the initila network is up, these nodes can then propose and approve new nodes to join the network. 

### Node Permission APIs
#### quorumNodeMgmt.permissionNodeList 
Input: None
Output: Returns the list of all enodes and their status 
Example:
```
> quorumNodeMgmt.permissionNodeList
[{
    enodeId: "enode://ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef@127.0.0.1:21000?discport=0&raftport=50401",
    status: "Approved"
}, {
    enodeId: "enode://0ba6b9f606a43a95edc6247cdb1c1e105145817be7bcafd6b2c0ba15d58145f0dc1a194f70ba73cd6f4cdd6864edc7687f311254c7555cc32e4d45aeb1b80416@127.0.0.1:21001?discport=0&raftport=50402",
    status: "Approved"
}, {
    enodeId: "enode://579f786d4e2830bbcc02815a27e8a9bacccc9605df4dc6f20bcc1a6eb391e7225fff7cb83e5b4ecd1f3a94d8b733803f2f66b7e871961e7b029e22c155c3a778@127.0.0.1:21002?discport=0&raftport=50403",
    status: "Approved"
}, {
    enodeId: "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404",
    status: "Approved"
}]
```
* Before any new nodes can be proposed to the network, network level voters have to be added. To manage voters at the network level, the following apis are there. `quorumNodeMgmt.addVoter` allows an account to be added as a voter. `quorumNodeMgmt.removeVoter` allows an account to be removed from the voter list. `quorumNodeMgmt.voterList` displays the list of all voters at network level
```
> quorumNodeMgmt.addVoter("0x0fBDc686b912d7722dc86510934589E0AAf3b55A", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumNodeMgmt.voterList
["0xed9d02e382b34818e88B88a309c7fe71E65f419d", "0x0fBDc686b912d7722dc86510934589E0AAf3b55A"]
> quorumNodeMgmt.removeVoter("0x0fBDc686b912d7722dc86510934589E0AAf3b55A", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumNodeMgmt.voterList
["0xed9d02e382b34818e88B88a309c7fe71E65f419d"]
```
* `quorumNodeMgmt.proposeNode` allows a new node to be propsoed into the network.
```
> quorumNodeMgmt.proposeNode("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
* `quorumNodeMgmt.approveNode` allows approval of a new node proposed to be part of the network. This api will be execyed by the accounts marked as voters and once majority voter accounts invoke the api, the node will be marked and move to `Approved` status
```
> quorumNodeMgmt.approveNode("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
* For proposing a node for deactivation and approving the node deactivation the following two apis can be used - `quorumNodeMgmt.proposeNodeDeactivation` and `quorumNodeMgmt.approveNodeDeactivation` 
```
> quorumNodeMgmt.proposeNodeDeactivation("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumNodeMgmt.approveNodeDeactivation("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
* To propose activation of a deactivated node `quorumNodeMgmt.proposeNodeActivation` can be used. The same can be approved using `quorumNodeMgmt.approveNodeActivation` api.
```
> quorumNodeMgmt.proposeNodeActivation("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumNodeMgmt.approveNodeActivation("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
* To propose a node for blacklisting `quorumNodeMgmt.proposeNodeBlacklisting` can be used. Node blacklisting can be approved using the api `quorumNodeMgmt.approveNodeBlacklisting`
```
> quorumNodeMgmt.proposeNodeBlacklisting("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumNodeMgmt.approveNodeBlacklisting("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```

## Account Access Control
The following account access types are being introduced as a part of this feature:

* ReadOnly: Accounts with this access will be able to perform only read activities and will not be able to deploy contracts or transactions. By default any account which is not permissioned will have a read only access. 
* Transact: Accounts with transact access will be able to commit transactions but will not be able to deploy contracts
* Contract Deploy: Accounts with this access will be able to deploy contracts and commit transactions
* FullAccess: Similar to "Contract Deploy" access, accounts with this access will be able to deploy contracts and perform transactions

Currently there is not any differences in the access types "Full Access" and "Contract Deploy". 

There are two apis available for managing the account permissions in the network:
* `quorumAcctMgmt.permissionAccountList` list the permissioned account list and the access that each account is having
```
> quorumAcctMgmt.permissionAccountList
[{
    access: "FullAccess",
    address: "0xcA843569e3427144cEad5e4d5999a3D0cCF92B8e"
}, {
    access: "Transact",
    address: "0xed9d02e382b34818e88B88a309c7fe71E65f419d"
}, {
    access: "ContractDeploy",
    address: "0x0fBDc686b912d7722dc86510934589E0AAf3b55A"
}, {
    access: "ReadOnly",
    address: "0x9186eb3d20Cbd1F5f992a950d808C4495153ABd5"
}
```
* To set the account access to a particula value `quorumAcctMgmt.setAccountAccess` can be used.
```
> quorumAcctMgmt.setAccountAccess("0x9186eb3d20cbd1f5f992a950d808c4495153abd5", "0", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
The table below indicates the numeric value for each account access type.

| AccessType      |           Value |
| :-------------: | :-------------: |
| FullAccess      |               0 |
| ReadOnly        |               1 |
| Transact        |               2 |
| ContractDeploy  |               3 |

Further while setting the account access, system checks if the account which is setting the access has sufficient privileges to perform the activity. For example a read only account cannot grant full access accounyt access to another account. The allowed access set ups are as below:
* Accounts with `FullAccess` can grant any access type ( FullAccess, Transact, ContractDeploy or ReadOnly) to any other account
* Accounts with `ContractDeploy` can grant only `Transact`, `ContractDeploy` or `ReadOnly` access to other accounts
* Accounts with `Transact` access grant only `Transact` or `ReadOnly` access to other accounts
* Accounts with `ReadOnly` access cannot grant any access

If an account having lower privileges tries to assign a higher privilege to an account, system will not allow the operation and will throw the an error as shown below:
```
> quorumAcctMgmt.setAccountAccess("0xAE9bc6cD5145e67FbD1887A5145271fd182F0eE7", "0", {from: eth.accounts[0]})
{
  msg: "Account does not have sufficient access for operation",
  status: false
}
```
