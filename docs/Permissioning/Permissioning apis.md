# Permission APIs
### quorumPermission.orgList 
* Input: None
* Output: Returns the list of all organizations and their status 
* Example:
```javascript
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
Please click [here](#organization-status-types) for the complete list of organization statuses.
### quorumPermission.acctList 
* Input: None
* Output: Returns the list of all accounts across organizations 
* Example:
```javascript
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
Please click [here](#account-status-types) for the complete list of account statuses.
### quorumPermission.nodeList 
* Input: None
* Output: Returns the list of all nodes across organizations 
* Example:
```javascript
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
Please click [here](#node-status-types) for the complete list of node statuses.

### quorumPermission.roleList 
* Input: None
* Output: Returns the list of all roles across organizations and their details
* Example:
```javascript
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
Please click [here](#account-access-types) for the complete list of different values of account access.

### quorumPermission.getOrgDetails 
This returns the list of accounts, nodes, roles, and sub organizations linked to an organization

* Input: organization id or sub organization id
* Output: list of all accounts, nodes, roles, and sub orgs
* Example:
```javascript
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
### quorumPermission.addOrg 
This api can be executed by a network admin account (`from:` in transactions args) only for proposing a new organization into the network

* Input: Unique alphanumeric organization id, enode id, account id (org admin account)
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.addOrg("ABC", "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
If there are any pending items for approval, proposal of any new organization will fail. Also the enode id and accounts can be linked to one organization only. 
```javascript
> quorumPermission.addOrg("ABC", "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
{
  msg: "Pending approvals for the organization. Approve first",
  status: false
}
> quorumPermission.addOrg("XYZ", "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
{
  msg: "EnodeId already part of network.",
  status: false
}
> quorumPermission.addOrg("XYZ", "enode://de9c2d5937e599930832cecc1df8cc90b50839bdf635c1a4e68e1dab2d001cd4a11c626e155078cc65958a72e2d72c1342a28909775edd99cc39470172cce0ac@127.0.0.1:21004?discport=0", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
{
  msg: "Account already in use in another organization",
  status: false
}

```

### quorumPermission.approveOrg 
This api can be executed by a network admin account (`from:` in transactions args) only for approving a proposed organization into the network.

* Input: Unique organization id, enode id, account id (org admin account)
* Output: Status of the operation
* Example:
```javascript
quorumPermission.approveOrg("ABC", "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
### quorumPermission.updateOrgStatus
This api can only be executed by a network admin account and is used for temporarily suspending an organization or re-enabling a suspended organization. This activity can be performed for master organization only and requires majority approval from network admins.

* Input: organization id, action (1 for suspending the organization and 2 for activating a suspended organization)
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.updateOrgStatus("ABC", 1, {from:eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
### quorumPermission.approveOrgStatus
This api can only be executed by a network admin account and is used for approving the org status change proposal.  Once majority approval is received from network admins, the org status is updated.

* Input: organization id, action (1 for suspending the organization and 2 for activating a suspended organization)
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.approveOrgStatus("ABC", 1, {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
When an organization is in suspended status, no transactions or contract deploy activities are allowed from any nodes linked to the org and sub organizations under it. Similarly no transactions will be allowed from any accounts linked to the organization

### quorumPermission.addSubOrg 
This api can be executed by a organization admin account to create a sub organization under the master org. 

* Input: parent org id, alphanumeric sub organization id,  enode id (not mandatory and can be null), account id (not mandatory and can be 0x0)
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.addSubOrg("ABC", "SUB1", "", "0x0000000000000000000000000000000000000000", {from: eth.accounts[0]})

{
  msg: "Action completed successfully",
  status: true
}
```
It should be noted that, parent org id should contain the complete org hierarchy from master org id to the immediate parent. The org hierarchy is separated by `.`. For example, if master org `ABC` has a sub organization `SUB1`, then while creating the sub organization at `SUB1` level, the parent org should be given as `ABC.SUB1`. Please see the examples below: 
```javascript
> quorumPermission.addSubOrg("ABC.SUB1", "SUB2","", "0x0000000000000000000000000000000000000000", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumPermission.addSubOrg("ABC.SUB1.SUB2", "SUB3","", "0x0000000000000000000000000000000000000000", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
### quorumPermission.addNewRole
This api can be executed by an organization admin account to create a new role for the organization.

* Input: organization id or sub organization id, alphanumeric role id, account access ([access values](#account-access-types))(, isVoter, isAdminRole
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.addNewRole("ABC", "TRANSACT", 1, false, false,{from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumPermission.addNewRole("ABC.SUB1.SUB2.SUB3", "TRANSACT", 1, false, false,{from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
### quorumPermission.removeRole
This api can be executed by an organization admin account to create a new role for the organization.

* Input: organization id or sub organization id, role id
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.removeRole("ABC.SUB1.SUB2.SUB3", "TRANSACT", {from: eth.accounts[1]})
{
  msg: "Action completed successfully",
  status: true
}
```
### quorumPermission.addAccountToOrg
This api can be executed by an organization admin to add an account to an organization and assign a role to the account

* Input: Account id, organization id or sub organization id, role to be assigned
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.addAccountToOrg("0xf017976fdf1521de2e108e63b423380307f501f8", "ABC", "TRANSACT", {from: eth.accounts[1]})
{
  msg: "Action completed successfully",
  status: true
}
```
The account can at best be linked to a single organization or sub organization and cannot belong to multiple organizations or sub organizations
```javascript
> quorumPermission.assignAccountRole("0xf017976fdf1521de2e108e63b423380307f501f8", "ABC.SUB1", "TRANSACT", {from: eth.accounts[1]})
{
  msg: "Account already in use in another organization",
  status: false
}
```
### quorumPermission.changeAccountRole
This api can be executed by an organization admin account to assign a role to an account.

* Input: Account id, organization id or sub organization id, role to be assigned
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.changeAccountRole("0xf017976fdf1521de2e108e63b423380307f501f8", "ABC", "TRANSACT", {from: eth.accounts[1]})
{
  msg: "Action completed successfully",
  status: true
}
```

### quorumPermission.updateAccountStatus
This api can be executed by an organization admin account to update the account status.

* Input:  organization id or sub organization id, Account id, action (1 for suspending the account, 2 for activating a suspended account, 3 for blacklisting the account)
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.updateAccountStatus("ABC", "0xf017976fdf1521de2e108e63b423380307f501f8", 1, {from: eth.accounts[1]})
{
  msg: "Action completed successfully",
  status: true
}
```
Once a account is blacklisted no further action is allowed on it.

### quorumPermission.assignAdminRole
This api can be executed by the network admin to add a new account as network admin or change the org admin account for an organization.

* Input: organization id to which the account belongs, account id, role id (it can be either org admin role or network admin role)
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.assignAdminRole("ABC", "0xf017976fdf1521de2e108e63b423380307f501f8", "NWADMIN", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```

### quorumPermission.approveAdminRole 
This api can be executed by the network admin to approve the organization admin or network admin role assignment to an account. The role is approved once majority approval is received.

* Input: organization id to which the account belongs, account id
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.approveAdminRole("ABC", "0xf017976fdf1521de2e108e63b423380307f501f8",  {from: eth.accounts[0]})

{
  msg: "Action completed successfully",
  status: true
}
```

### quorumPermission.addNode
This api can be executed by the organization admin account to add a node to the organization or sub organization.

* Input:  organization id or sub organization id, enode id
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.addNode("ABC.SUB1.SUB2.SUB3", "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0&raftport=50407", {from: eth.accounts[1]})
{
  msg: "Action completed successfully",
  status: true
}
```
A node cannot be part of multiple organizations. 

### quorumPermission.updateNodeStatus
This api can be executed by the organization admin account to update the status of a node.

* Input:  organization id or sub organization id, enode id, action (1 for deactivating the node, 2 for activating a deactivated node and 3 for blacklisting a node)
* Output: Status of the operation
* Example:
```javascript
> quorumPermission.updateNodeStatus("ABC.SUB1.SUB2.SUB3", "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0&raftport=50407",3, {from: eth.accounts[1]})
{
  msg: "Action completed successfully",
  status: true
}
```
Once a node is blacklisted no further action is possible on the same.

### Roles

#### Organization status types
The table below indicates the numeric value for various organization statuses.

| OrgStatus                 |           Value |
| :-----------------------: | :-------------: |
| NotInList                 |               0 |
| Proposed                  |               1 |
| Approved                  |               2 |
| PendingSuspension         |               3 |
| Suspended                 |               4 |
| AwaitingSuspensionRevoke  |               5 |

#### Account status types
The table below indicates the numeric value for various account statuses.

| AccountStatus   |           Value |
| :-------------: | :-------------: |
| NotInList       |               0 |
| PendingApproval |               1 |
| Active          |               2 |
| Inactive        |               3 |
| Suspended       |               4 |
| Blacklisted     |               5 |
| Revoked         |               6 |

#### Account access types
The table below indicates the numeric value for each account access type.

| AccessType      |           Value |
| :-------------: | :-------------: |
| ReadOnly        |               0 |
| Transact        |               1 |
| Contract deploy |               2 |
| Full access     |               3 |

When setting the account access, the system checks if the account setting the access has sufficient privileges to perform the activity. 

* Accounts with `FullAccess` can grant any access type ( FullAccess, Transact, ContractDeploy or ReadOnly) to any other account
* Accounts with `ContractDeploy` can grant only `Transact`, `ContractDeploy` or `ReadOnly` access to other accounts
* Accounts with `Transact` access can grant only `Transact` or `ReadOnly` access to other accounts
* Accounts with `ReadOnly` access cannot grant any access

#### Node Status types
The table below indicates the numeric value for various node statuses.

| NodeStatus                |           Value |
| :-----------------------: | :-------------: |
| NotInList                 |               0 |
| PendingApproval           |               1 |
| Approved                  |               2 |
| Deactivated               |               3 |
| Blacklisted               |               4 |

