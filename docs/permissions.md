# Overview
Currently when Quorum geth is brought up in `--permissioned` mode, for establishing p2p connectivity the nodes check `permissioned-nodes.json` file. If the enode id of the node requesting connection is present in this file, the p2p connection is established else it is declined. The `permissioned-nodes.json` file is updated procedurally whenever a new node joins the network. Node permissioning feature will allow the existing nodes to propose a new node to join the network and once majority voting is done on the proposal, it will update the `permissioned-nodes.json` automatically. Further the existing nodes can propose any node for deactivation, blacklisting and activating a node back from a deactivated status.

Account permissioning feature introduces controls at account level. This will control the type of activities that a particular account can perform in the network.

It should be noted that both the above features will be available when Quorum geth is brought in `--permissioned` mode with the set up as described in the next section. 

## Set up
Node permissioning and account access control is managed by a smart contract [PermissionsInterface.sol](../controls/permission/PermissionsInterface.sol). The precompiled byte code of the smart contract is deployed at address `0x000000000000000000020` in network boot-up process. The binding of the precompiled byte code with the address is in `genesis.json`. The initial set of account that will have full access when the network is up, should given as a part of `genesis.json` as shown below:
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

Further, if the initial set of network nodes are brought up with `--permissioned` mode and a new approved node is joining the network without `--permissioned` flag in the `geth` start command, system will not allow the new `geth` node to come and a fatal error `Joining a permissioned network in non-permissioned mode. Bring up geth with --permissioned."` will be thrown.

## Node Permissioning 
In a permissioned network any node can be in one of the following status:
* Proposed - The node has been proposed to join the network and pending majority voting to be marked as `Approved`
* Approved - The node is approved and is part of the network
* PendingDeactivation - The node has been proposed for deactivation from network and is pending majority approval
* Deactivated - The node has been deactivated from the network. Any node in this status will be disconnected from the network and block sync with this node will not happen
* PendingActivation - A deactivated node has been proposed for activation and is pending majority approval. Once approved the node will move to `Approved` status
* PendingBlacklisting - The node has been proposed for blacklisting and is pending majority approval
* Blacklisted - The node has been blacklisted. If the node was an active node on the network, the node will be disconnected from the network and block sync will stop

It should be noted that deactivation is temporary in nature and hence the node can join back the network at a later point in time. However blacklisting is permanent in nature. A blacklisted node will never be able to join back the network. Further blacklisting can be proposed for a node which is in the network or a new node which is currently not part of the network. 

When the network is started for the first time, all the nodes present in `static-nodes.json` file are added to the permissioned nodes list maintained in the smart contract. Once the initial network is up, these nodes can then propose and approve new nodes to join the network. 

### Node Permission APIs
#### quorumNodeMgmt.permissionNodeList 
* Input: None
* Output: Returns the list of all enodes and their status 
* Example:
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
#### quorumNodeMgmt.addVoter 
Before a new node can be proposed to the network, the network should have valid voters. This api allows an account to be added as voter to the network. Only an account with full access can add another account as voter. Further the account being added as voter account should have at least transact permission. 
* Input: Account to be added as voter, transaction object
* Output: Status of the operation
* Example:
```
> quorumNodeMgmt.addVoter("0x0fBDc686b912d7722dc86510934589E0AAf3b55A", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
#### quorumNodeMgmt.removeVoter 
Allows a voter account to be removed from the network. Only an account with `FullAccess` can perform this activity.
* Input: Account to be removed, transaction object
* Output: Status of the operation
* Example:
```
> quorumNodeMgmt.removeVoter("0x0fBDc686b912d7722dc86510934589E0AAf3b55A", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
#### quorumNodeMgmt.voterList 
* Input: None
* Output: List of all voters on the network
* Example:
```
> quorumNodeMgmt.voterList
["0xed9d02e382b34818e88B88a309c7fe71E65f419d", "0x0fBDc686b912d7722dc86510934589E0AAf3b55A"]
```

#### quorumNodeMgmt.proposeNode 
Allows a new enode to be proposed to the network. This operation will be allowed only if at least one voter account is present in the network.
* Input: enode to be proposed, transaction object
* Output: Status of the operation
* Example:
```
> quorumNodeMgmt.proposeNode("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
#### quorumNodeMgmt.approveNode 
API for approving a proposed node. The node gets approved once majority votes from the voter accounts is received.
* Input: enode to be approved, transaction object
* Output: Status of the operation
* Example:
```
> quorumNodeMgmt.approveNode("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
#### quorumNodeMgmt.proposeNodeDeactivation 
API for proposing a node for deactivation. The node must be `Approved` state and there should be at least one voter account present at network.
* Input: enode to be deactivated, transaction object
* Output: Status of the operation
* Example:
```
> quorumNodeMgmt.proposeNodeDeactivation("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
#### quorumNodeMgmt.approveNodeDeactivation
API for approving node for deactivation. The node gets deactivated once majority votes from the voter accounts is received
* Input: enode to be deactivated, transaction object
* Output: Status of the operation
* Example:
```
> quorumNodeMgmt.approveNodeDeactivation("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
#### quorumNodeMgmt.proposeNodeActivation 
API for proposing activation of a deactivated node. The node must be in `Deactivated` state and there should be at least one voter account present at network.
* Input: enode to be activated, transaction object
* Output: Status of the operation
* Example:
```
> quorumNodeMgmt.proposeNodeActivation("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
#### quorumNodeMgmt.approveNodeActivation 
API for approval of activating a deactivated node. The node gets activated once majority votes from the voter accounts is received
* Input: enode to be activated, transaction object
* Output: Status of the operation
* Example:
```
> quorumNodeMgmt.approveNodeActivation("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
#### quorumNodeMgmt.proposeNodeBlacklisting 
API for blacklisting a node from the network. Any node (irrespective of node status or a node which is not part of network) can be proposed for blacklisting. Blacklisting takes precedence on any other proposal. 
* Input: enode to be blacklisted, transaction object
* Output: Status of the operation
* Example:
```
> quorumNodeMgmt.proposeNodeBlacklisting("enode://3701f007bfa4cb26512d7df18e6bbd202e8484a6e11d387af6e482b525fa25542d46ff9c99db87bd419b980c24a086117a397f6d8f88e74351b41693880ea0cb@127.0.0.1:21004?discport=0&raftport=50405", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
#### quorumNodeMgmt.approveNodeBlacklisting 
API for approving node blacklisting. The node is blacklisted once majority votes from the voter accounts. Once the node is blacklisted, it cannot rejoin the network.
* Input: enode to be blacklisted, transaction object
* Output: Status of the operation
* Example:
```
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
* Contract deploy: Accounts with this access will be able to deploy contracts and commit transactions
* Full access: Similar to "Contract deploy" access, accounts with this access will be able to deploy contracts and perform transactions. Further only an account with Full access can add voters to the network. 

### Account Access APIs
#### quorumAcctMgmt.permissionAccountList
* Input: None
* Output: Returns the list of all permissioned accounts with account access for each 
* Example:
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
#### quorumAcctMgmt.setAccountAccess
* Input: Account, access type for the account and transaction object
* Output: Status of the operation
* Example:
```
> quorumAcctMgmt.setAccountAccess("0x9186eb3d20cbd1f5f992a950d808c4495153abd5", 2, {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
### General validations for account access
The table below indicates the numeric value for each account access type.

| AccessType      |           Value |
| :-------------: | :-------------: |
| ReadOnly        |               0 |
| Transact        |               1 |
| Contract deploy |               2 |
| Full access     |               3 |

While setting the account access, system checks if the account which is setting the access has sufficient privileges to perform the activity. 
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
