# Introduction
The current permission model with in Quorum is limited to node level permissions only and allows a set of nodes which are part of `permissioned-nodes.json` to join the network. Considering the enterprise level needs for private and consortium Blockchains, the permissionsing model has been enhanced inline with EEA(Etnerprise Ethreum Alliance) specs. The overview of the mode is as depicted below:
![permissions mode](images/PermissionsModel.png)  
### Key Definitions
* Network - A set of organizations
* Organization - A set of Ethereum accounts, nodes having varied permissions to interact with the network
* Sub Organization - Further sub grouping with in the Organization as per business need
* Account - An Ethereum account 
* Voter - An account capable of voting for a certain action
* Role - A named job function in organization
* Node - A geth node which is part of the network and belongs to an organization or sub organization

As depicted above, in the enhanced permissions model, the network is comprises of group of organization. The voters defined at network level can add these organizations and assign an account as the organization administration account. Each organization is has roles, accounts linked to these roles and nodes. The organization can further have sub organizations and each sub organization can have its own set of roles accounts and sets. The organization administration account manage and control all activities at the organization level. Further the organization administrator can create an admin role and assign the same to a different account to manage the administration of a sub organization. The access rights of an account are derived based on the role assigned to it. The account will be able to transact via any node linked to the sub org or overall organizations level.  

### Smart Contract design for permissions
The permissions model is completely built on smart contracts. The smart contract design is as below:
![contract design](images/ContractDesign.png)

The permissions smart contract design follows the Proxy-Implementation-Storage pattern which allows the implementation logic to be changed without changing the storage or interface layer. Brief description of the smart contracts is as below:
* `PermissionsUpgradable.sol`: This contract stores the address of latest implementation contract and is owned by a custodian( an Ethereum account). Only custodian is allowed to change the implementation contract address. 
* `PermissionsInterface.sol`: This contract interface contract and holds the interfaces for permissions related actions
* `PermissionsImplementation.sol`: This contract has the business logic for the permissions actions. This contract can receive requests only from a valid interface as defined in `PermissionsUpgradable.sol` and interact with all the storage contracts for respective actions.
* `OrgManager.sol`: This contract stores data for organizations and sub organizations. This contract can receive request from valid implementation contract as defined in `PermissionsUpgrdable.sol`
* `AccountManager.sol`: This contract can receive request from valid implementation contract as defined in `PermissionsUpgrdable.sol`. This contracts stores the data of all accounts, their linkage to organization and various roles. This contracts also stores the status of an account. The account can be in any of the following states - `PendingApproval`, `Active`, `Suspended`, `Blacklisted`, `Revoked`
* `NodeManager.sol`: This contract can receive request from valid implementation contract as defined in `PermissionsUpgrdable.sol`. This contracts stores the data of a node, its linkage to a organization or sub organization and status of the node. The node status can be any one of the following states - `PendingApproval`, `Active`, `Deactivated`, `Blacklisted`
* `RoleManager.sol`: This contract can receive request from valid implementation contract as defined in `PermissionsUpgrdable.sol`. This contract stores data for various roles and the organization to which it is linked. At access at role level can be any one of the following: `Readonly` which allows only read operations, `Transact` which allows value transfer but no contract deployment access, `ContractDeploy` which allows both value transfer and contract deployment access and `FullAccess` which allows additional network level accesses in addition to value transfer and contract deployment. If a role is revoked all accounts which are linked to the role lose all access rights.
* `VoterManager.sol`: This contract can receive request from valid implementation contract as defined in `PermissionsUpgrdable.sol`. This contract stores the data of valid voters at network level which can approve identified activities e.g. adding a new organization to the network. Any account which is linked to a predefined network admin role will be marked as a voter. Whenever a network level activity which requires voting is performed, a voting item is added to this contract and each voter account can vote for the activity. The activity is marked as `Approved` upon majority voting.



## Set up


### Permission APIs
#### quorumPermission.orgList 
* Input: None
* Output: Returns the list of all organizations and their status 
* Example:
```
> quorumPermission.orgList
[{
    fullOrgId: "INITORG",
    level: 1,
    orgId: "INITORG",
    parentOrgId: "",
    status: 2,
    subOrgList: null,
    ultimateParent: "INITORG"
}]
```
#### quorumPermission.acctList 
* Input: None
* Output: Returns the list of all accounts across organizations 
* Example:
```
> quorumPermission.acctList
[{
    acctId: "0xed9d02e382b34818e88b88a309c7fe71e65f419d",
    isOrgAdmin: true,
    orgId: "INITORG",
    roleId: "NWADMIN",
    status: 2
}, {
    acctId: "0xca843569e3427144cead5e4d5999a3d0ccf92b8e",
    isOrgAdmin: true,
    orgId: "INITORG",
    roleId: "NWADMIN",
    status: 2
}]
```
#### quorumPermission.nodeList 
* Input: None
* Output: Returns the list of all nodes across organizations 
* Example:
```
> quorumPermission.nodeList
[{
    orgId: "INITORG",
    status: 2,
    url: "enode://72c0572f7a2492cffb5efc3463ef350c68a0446402a123dacec9db5c378789205b525b3f5f623f7548379ab0e5957110bffcf43a6115e450890f97a9f65a681a@127.0.0.1:21000?discport=0"
}, {
    orgId: "INITORG",
    status: 2,
    url: "enode://7a1e3b5c6ad614086a4e5fb55b6fe0a7cf7a7ac92ac3a60e6033de29df14148e7a6a7b4461eb70639df9aa379bd77487937bea0a8da862142b12d326c7285742@127.0.0.1:21001?discport=0"
}, {
    orgId: "INITORG",
    status: 2,
    url: "enode://5085e86db5324ca4a55aeccfbb35befb412def36e6bc74f166102796ac3c8af3cc83a5dec9c32e6fd6d359b779dba9a911da8f3e722cb11eb4e10694c59fd4a1@127.0.0.1:21002?discport=0"
}, {
    orgId: "INITORG",
    status: 2,
    url: "enode://28a4afcf56ee5e435c65b9581fc36896cc684695fa1db83c9568de4353dc6664b5cab09694d9427e9cf26a5cd2ac2fb45a63b43bb24e46ee121f21beb3a7865e@127.0.0.1:21003?discport=0"
}]
```
#### quorumPermission.roleList 
* Input: None
* Output: Returns the list of all roles across organizations and their details
* Example:
```
> quorumPermission.roleList
[{
    access: 3,
    active: true,
    isAdmin: true,
    isVoter: true,
    orgId: "INITORG",
    roleId: "NWADMIN"
}]
```
#### quorumPermission.getOrgDetails 
This returns the list of accounts, nodes, roles, sub organizations linked to an organization
* Input: idrganization or sub organization id
* Output: list of all accounts, roles, nodes and sub orgs
* Example:
```
> quorumPermission.getOrgDetails("INITORG")
{
  acctList: [{
      acctId: "0xed9d02e382b34818e88b88a309c7fe71e65f419d",
      isOrgAdmin: true,
      orgId: "INITORG",
      roleId: "NWADMIN",
      status: 2
  }, {
      acctId: "0xca843569e3427144cead5e4d5999a3d0ccf92b8e",
      isOrgAdmin: true,
      orgId: "INITORG",
      roleId: "NWADMIN",
      status: 2
  }],
  nodeList: [{
      orgId: "INITORG",
      status: 2,
      url: "enode://72c0572f7a2492cffb5efc3463ef350c68a0446402a123dacec9db5c378789205b525b3f5f623f7548379ab0e5957110bffcf43a6115e450890f97a9f65a681a@127.0.0.1:21000?discport=0"
  }, {
      orgId: "INITORG",
      status: 2,
      url: "enode://7a1e3b5c6ad614086a4e5fb55b6fe0a7cf7a7ac92ac3a60e6033de29df14148e7a6a7b4461eb70639df9aa379bd77487937bea0a8da862142b12d326c7285742@127.0.0.1:21001?discport=0"
  }, {
      orgId: "INITORG",
      status: 2,
      url: "enode://5085e86db5324ca4a55aeccfbb35befb412def36e6bc74f166102796ac3c8af3cc83a5dec9c32e6fd6d359b779dba9a911da8f3e722cb11eb4e10694c59fd4a1@127.0.0.1:21002?discport=0"
  }, {
      orgId: "INITORG",
      status: 2,
      url: "enode://28a4afcf56ee5e435c65b9581fc36896cc684695fa1db83c9568de4353dc6664b5cab09694d9427e9cf26a5cd2ac2fb45a63b43bb24e46ee121f21beb3a7865e@127.0.0.1:21003?discport=0"
  }],
  roleList: [{
      access: 3,
      active: true,
      isAdmin: true,
      isVoter: true,
      orgId: "INITORG",
      roleId: "NWADMIN"
  }],
  subOrgList: null
}
```
#### quorumPermission.addOrg 
This api can be executed by a network admin account only for proposing a new organization into the network
* Input: Unique organization id, enode id, account id
* Output: Status of the operation
* Example:
```
> quorumPermission.addOrg("ABC", "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
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

