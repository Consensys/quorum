# Permission APIs
## APIs
### `quorumPermission_orgList` 
Returns the list of all organizations with the status of each organization in the network
#### Parameters
None
#### Returns
* `fullOrgId`: complete org id including the all parent org ids separated by ".". 
* `level`: level of the org in org hierarchy
* `orgId`: organization identifier
* `parentOrgId`: immediate parent org id
* `status`: org status. [refer](#organization-status-types) for complete list of statuses
* `subOrgList`: list of sub orgs linked to the org
* `ultimateParent`: Master org under which the org falls 
#### Examples
```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_orgList","id":10}' --header "Content-Type: application/json"

// Response
{
    fullOrgId: "INITORG",
    level: 1,
    orgId: "INITORG",
    parentOrgId: "",
    status: 2,
    subOrgList: null,
    ultimateParent: "INITORG"
}
```

```javascript tab="geth console"
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
### `quorumPermission_acctList` 
Returns the list of accounts permissioned in the network

#### Parameters
None

#### Returns
* `acctId`: account id 
* `isOrgAdmin`: indicates if the account is admin account for the organization
* `orgId`: org identifier
* `roleId`: role assigned to the account
* `status`: account status. [refer](#account-status-types) for the complete list of account status.

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_acctList","id":10}' --header "Content-Type: application/json"

// Response
{
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
}
```

```javascript tab="geth console"
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
### `quorumPermission_nodeList` 
Returms the list of nodes part of the network
#### Parameters
None
#### Returns
* `orgId`: org id to which the node belongs
* `status`: status of the node. [refer](#node-status-types) for the complete list of node statuses
* `url`: complete enode id
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_nodeList","id":10}' --header "Content-Type: application/json"

// Response
{
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
}
```

```javascript tab="geth console"
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

### `quorumPermission_roleList` 
Returns the list of roles in the network
#### Parameters
None
#### Returns
* `access`: account access. [refer](#account-access-types) for the complete list of different values of account access.
* `active`: indicates if the role is active or not
* `isAdmin`: indicates if the role is org admin role
* `isVoter`: indicates if the role is enabled for voting. Applicable only for network admin role
* `orgId`: org id to which the role is linked
* `roleId`: unique role id
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_roleList","id":10}' --header "Content-Type: application/json"

// Response
{
    access: 3,
    active: true,
    isAdmin: true,
    isVoter: true,
    orgId: "INITORG",
    roleId: "NWADMIN"
}
```

```javascript tab="geth console"
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

### `quorumPermission_getOrgDetails` 
This returns the list of accounts, nodes, roles, and sub organizations linked to an organization
#### Parameters
* org or sub org id
#### Returns
* `acctList`
* `nodeList`
* `roleList`
* `subOrgList`: array of sub orgs linked to the org
* Output: list of all accounts, nodes, roles, and sub orgs
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_getOrgDetails","params":["INITORG"],"id":10}' --header "Content-Type: application/json"

// Response
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

```javascript tab="geth console"
> quorumPermission_getOrgDetails("INITORG")
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
### `quorumPermission_addOrg` 
This api can be executed by a network admin account (`from:` in transactions args) only for proposing a new organization into the network
#### Parameter
* `orgId`: unique org identfiier
* `enodeId`: complete enode id
* `accountId`: account which will be the org admin account

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_addOrg","params":["ABC", "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.addOrg("ABC", "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
"Action completed successfully"
```
If there are any pending items for approval, proposal of any new organization will fail. Also the enode id and accounts can be linked to one organization only. 
```javascript tab="geth console"
> quorumPermission.addOrg("ABC", "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
Error: Pending approvals for the organization. Approve first
    at web3.js:3143:20
    at web3.js:6347:15
    at web3.js:5081:36
    at <anonymous>:1:1

> quorumPermission.addOrg("XYZ", "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
Error: EnodeId already part of network.
    at web3.js:3143:20
    at web3.js:6347:15
    at web3.js:5081:36
    at <anonymous>:1:1
> quorumPermission.addOrg("XYZ", "enode://de9c2d5937e599930832cecc1df8cc90b50839bdf635c1a4e68e1dab2d001cd4a11c626e155078cc65958a72e2d72c1342a28909775edd99cc39470172cce0ac@127.0.0.1:21004?discport=0", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
Error: Account already in use in another organization
    at web3.js:3143:20
    at web3.js:6347:15
    at web3.js:5081:36
    at <anonymous>:1:1

```
### `quorumPermission_approveOrg` 
This api can be executed by a network admin account (`from:` in transactions args) only for approving a proposed organization into the network.
#### Parameters
* `orgId`: unique org identfiier
* `enodeId`: complete enode id
* `accountId`: account which will be the org admin account
#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_approveOrg","params":["ABC", "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
quorumPermission.approveOrg("ABC", "enode://3d9ca5956b38557aba991e31cf510d4df641dce9cc26bfeb7de082f0c07abb6ede3a58410c8f249dabeecee4ad3979929ac4c7c496ad20b8cfdd061b7401b4f5@127.0.0.1:21003?discport=0&raftport=50404", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
"Action completed successfully"
```
### `quorumPermission_updateOrgStatus`
This api can only be executed by a network admin account and is used for temporarily suspending an organization or re-enabling a suspended organization. This activity can be performed for master organization only and requires majority approval from network admins.
#### Parameters
* `orgId`: org id 
* `action`: 
    * 1 - for suspending a org
    * 2 - for activating a suspended organization
#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_updateOrgStatus","params":["ABC", 1, {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"
//Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.updateOrgStatus("ABC", 1, {from:eth.accounts[0]})
"Action completed successfully"
```

### `quorumPermission_approveOrgStatus`
This api can only be executed by a network admin account and is used for approving the org status change proposal.  Once majority approval is received from network admins, the org status is updated.

#### Parameters
* `orgId`: org id 
* `action`: 
    * 1 - for approving org suspension
    * 2 - for approving activation of suspended org

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_approveOrgStatus","params":["ABC", 1, {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

//Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
quorumPermission.approveOrgStatus("ABC", 1, {from: eth.accounts[0]})
"Action completed successfully"

```

When an organization is in suspended status, no transactions or contract deploy activities are allowed from any nodes linked to the org and sub organizations under it. Similarly no transactions will be allowed from any accounts linked to the organization

### `quorumPermission_addSubOrg` 
This api can be executed by a organization admin account to create a sub organization under the master org.
#### Parameters
* `parentOrgId`: parent org id under which the sub org is being added. parent org id should contain the complete org hierarchy from master org id to the immediate parent. The org hierarchy is separated by `.`. For example, if master org `ABC` has a sub organization `SUB1`, then while creating the sub organization at `SUB1` level, the parent org should be given as `ABC.SUB1`
* `subOrgId`: sub org identifier
* `enodeId`: complete enode id of the node linked to the sub org id
#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_addSubOrg","params":["ABC", "SUB1","", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.addSubOrg("ABC", "SUB1", "", {from: eth.accounts[0]})
"Action completed successfully"
```

Few examples of adding sub org in nested hierarchy:
```javascript
> quorumPermission.addSubOrg("ABC.SUB1", "SUB2","",  {from: eth.accounts[0]})
"Action completed successfully"

> quorumPermission.addSubOrg("ABC.SUB1.SUB2", "SUB3","",  {from: eth.accounts[0]})
"Action completed successfully"
```

### `quorumPermission_addNewRole`
This api can be executed by an organization admin account to create a new role for the organization.

#### Parameters
* `orgId`: org id for which the role is being created
* `roleId`: unique role identifier
* `accountAccess`: account level access. [Refer](#account-access-types) for complete list
* `isVoter`: `bool` indicates if its a voting role
* `isAdminRole`: `bool` indicates if its an admin role

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_addNewRole","params":["ABC", "TRANSACT",1,false,false, {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.addNewRole("ABC", "TRANSACT", 1, false, false,{from: eth.accounts[0]})
"Action completed successfully"
> quorumPermission.addNewRole("ABC.SUB1.SUB2.SUB3", "TRANSACT", 1, false, false,{from: eth.accounts[0]})
"Action completed successfully"
```

### `quorumPermission_removeRole`
This api can be executed by an organization admin account to create a new role for the organization.

#### Parameters
* `orgId`: org or sub org id to which the role belongs
* `roleId`: role id

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_removeRole","params":["ABC", "TRANSACT", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.removeRole("ABC.SUB1.SUB2.SUB3", "TRANSACT", {from: eth.accounts[1]})
"Action completed successfully"
```

### `quorumPermission_addAccountToOrg`
This api can be executed by an organization admin to add an account to an organization and assign a role to the account

#### Parameters
* `acctId`: org or sub org id to which the role belongs
* `orgId`: org id
* `roleId`: role id

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_addAccountToOrg","params":["0xf017976fdf1521de2e108e63b423380307f501f8", "ABC", "TRANSACT", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.addAccountToOrg("0xf017976fdf1521de2e108e63b423380307f501f8", "ABC", "TRANSACT", {from: eth.accounts[1]})
"Action completed successfully"
```

The account can at best be linked to a single organization or sub organization and cannot belong to multiple organizations or sub organizations
```javascript
> quorumPermission.addAccountToOrg("0xf017976fdf1521de2e108e63b423380307f501f8", "ABC.SUB1", "TRANSACT", {from: eth.accounts[1]})
Error: Account already in use in another organization
    at web3.js:3143:20
    at web3.js:6347:15
    at web3.js:5081:36
    at <anonymous>:1:1
```
### `quorumPermission_changeAccountRole`
This api can be executed by an organization admin account to assign a role to an account.
#### Parameters
* `acctId`: account id
* `orgId`: org id
* `roleId`: new role id to be assigned to the account
#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure
#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_changeAccountRole","params":["0xf017976fdf1521de2e108e63b423380307f501f8", "ABC", "TRANSACT", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.changeAccountRole("0xf017976fdf1521de2e108e63b423380307f501f8", "ABC", "TRANSACT", {from: eth.accounts[1]})
"Action completed successfully"
```

### `quorumPermission_updateAccountStatus`
This api can be executed by an organization admin account to update the account status.

#### Parameters
* `orgId`: org id
* `acctId`: org or sub org id to which the role belongs
* `action`: 
    * 1 - for suspending the account
    * 2 - for activating a suspended account
    * 3 - for blacklisting an account
    
#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_updateAccountStatus","params":["ABC", "0xf017976fdf1521de2e108e63b423380307f501f8", 1, {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.updateAccountStatus("ABC", "0xf017976fdf1521de2e108e63b423380307f501f8", 1, {from: eth.accounts[1]})
"Action completed successfully"
```

Once a account is blacklisted it can only be recovered by network admins. Refer to [quorumPermission_recoverBlackListedAccount](#quorumpermission_recoverblacklistedaccount) and [quorumPermission_approveBlackListedAccountRecovery](#quorumpermission_approveblacklistedaccountrecovery) for further details.

### `quorumPermission_recoverBlackListedAccount`
This api can be executed by the network admin account to initiate the recovery of a blacklisted account. Post majority approval from network admin accounts, the blacklisted account will be marked as active.  

#### Parameters
* `orgId`: org or sub org id to which the node belongs
* `acctId`: blacklisted account id

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure


```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_recoverBlackListedAccount","params":["ABC.SUB1.SUB2.SUB3", "0xf017976fdf1521de2e108e63b423380307f501f8", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.recoverBlackListedAccount("ABC.SUB1.SUB2.SUB3", "0xf017976fdf1521de2e108e63b423380307f501f8", {from: eth.accounts[1]})
"Action completed successfully"
```

### `quorumPermission_approveBlackListedAccountRecovery`
This api can be executed by the network admin approve the recovery of a blacklisted account. Once majority approvals from network admin accounts is received, the account is marked as active. 

#### Parameters
* `orgId`: org or sub org id to which the node belongs
* `acctId`: blacklisted account id

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_approveBlackListedNodeRecovery","params":["ABC.SUB1.SUB2.SUB3", "0xf017976fdf1521de2e108e63b423380307f501f8", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.approveBlackListedNodeRecovery("ABC.SUB1.SUB2.SUB3", "0xf017976fdf1521de2e108e63b423380307f501f8", {from: eth.accounts[1]})
"Action completed successfully"
```

### `quorumPermission_assignAdminRole`
This api can be executed by the network admin to add a new account as network admin or change the org admin account for an organization.

#### Parameters
* `orgId`: org id to which the account belongs
* `acctId`: account id
* `roleId`: new role id to be assigned to the account. This can be the network admin role or org admin role only

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_assignAdminRole","params":["ABC", "0xf017976fdf1521de2e108e63b423380307f501f8", "NWADMIN", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"
// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.assignAdminRole("ABC", "0xf017976fdf1521de2e108e63b423380307f501f8", "NWADMIN", {from: eth.accounts[0]})
"Action completed successfully"
```

### `quorumPermission_approveAdminRole` 
This api can be executed by the network admin to approve the organization admin or network admin role assignment to an account. The role is approved once majority approval is received.

#### Parameters
* `orgId`: org id to which the account belongs
* `acctId`: account id

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_approveAdminRole","params":["ABC", "0xf017976fdf1521de2e108e63b423380307f501f8", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.approveAdminRole("ABC", "0xf017976fdf1521de2e108e63b423380307f501f8",  {from: eth.accounts[0]})
"Action completed successfully"
```

### `quorumPermission_addNode`
This api can be executed by the organization admin account to add a node to the organization or sub organization. A node cannot be part of multiple organizations.

#### Parameters
* `orgId`: org or sub org id to which the node belongs
* `enodeId`: complete enode id

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_addNode","params":["ABC.SUB1.SUB2.SUB3", "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0&raftport=50407", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.addNode("ABC.SUB1.SUB2.SUB3", "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0&raftport=50407", {from: eth.accounts[1]})
"Action completed successfully"
```

### `quorumPermission_updateNodeStatus`
This api can be executed by the organization admin account to update the status of a node.

#### Parameters
* `orgId`: org or sub org id to which the node belongs
* `enodeId`: complete enode id
* `action`: 
    * 1 - for deactivating the node
    * 2 - for activating a deactivated node
    * 3 - for blacklisting a node

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_updateNodeStatus","params":["ABC.SUB1.SUB2.SUB3", "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0&raftport=50407",1, {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.updateNodeStatus("ABC.SUB1.SUB2.SUB3", "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0&raftport=50407",3, {from: eth.accounts[1]})
"Action completed successfully"
```

Once a node is blacklisted it can only be recovered by network admins. Refer to [quorumPermission_recoverBlackListedNode](#quorumpermission_recoverblacklistednode) and [quorumPermission_approveBlackListedNodeRecovery](#quorumpermission_approveblacklistednoderecovery) for further details.

### `quorumPermission_recoverBlackListedNode`
This api can be executed by the network admin account to initiate the recovery of a blacklisted node. Post majority approval from network admin accounts, the blacklisted node will be marked as active.  

#### Parameters
* `orgId`: org or sub org id to which the node belongs
* `enodeId`: complete enode id

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_recoverBlackListedNode","params":["ABC.SUB1.SUB2.SUB3", "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0&raftport=50407", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.recoverBlackListedNode("ABC.SUB1.SUB2.SUB3", "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0&raftport=50407", {from: eth.accounts[1]})
"Action completed successfully"
```

### `quorumPermission_approveBlackListedNodeRecovery`
This api can be executed by the network admin approve the recovery of a blacklisted node. Once majority approvals from network admin accounts is received, the node is marked as active. 

#### Parameters
* `orgId`: org or sub org id to which the node belongs
* `enodeId`: complete enode id

#### Returns
* `msg`: response message
* `status`: `bool` indicating if the operation was success or failure

#### Examples

```jshelllanguage tab="JSON RPC"
// Request
curl -X POST http://127.0.0.1:22000 --data '{"jsonrpc":"2.0","method":"quorumPermission_approveBlackListedNodeRecovery","params":["ABC.SUB1.SUB2.SUB3", "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0&raftport=50407", {"from":"0xed9d02e382b34818e88b88a309c7fe71e65f419d"}],"id":10}' --header "Content-Type: application/json"

// Response
{"jsonrpc":"2.0","id":10,"result":"Action completed successfully"}
```

```javascript tab="geth console"
> quorumPermission.approveBlackListedNodeRecovery("ABC.SUB1.SUB2.SUB3", "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0&raftport=50407", {from: eth.accounts[1]})
"Action completed successfully"
```

## Roles
The table below indicates the numeric value for each account access type.

|   AccessType   | Value |
|:--------------:|:-----:|
|    ReadOnly    |   0   |
|    Transact    |   1   |
| ContractDeploy |   2   |
|   FullAccess   |   3   |

When setting the account access, the system checks if the account setting the access has sufficient privileges to perform the activity. 

* Accounts with `FullAccess` can grant any access type (`FullAccess`, `Transact`, `ContractDeploy` or `ReadOnly`) to any other account
* Accounts with `ContractDeploy` can grant only `Transact`, `ContractDeploy` or `ReadOnly` access to other accounts
* Accounts with `Transact` access can grant only `Transact` or `ReadOnly` access to other accounts
* Accounts with `ReadOnly` access cannot grant any access

## Status Mapping
### Organization status types
The table below indicates the numeric value for various organization status.

| OrgStatus                 |           Value |
| :-----------------------: | :-------------: |
| NotInList                 |               0 |
| Proposed                  |               1 |
| Approved                  |               2 |
| PendingSuspension         |               3 |
| Suspended                 |               4 |
| AwaitingSuspensionRevoke  |               5 |

### Account status types
The table below indicates the numeric value for various account status.

| AccountStatus                                         |             Value |
| :-------------:                                       |   :-------------: |
| Not In List                                           |                 0 |
| Pending Approval                                      |                 1 |
| Active                                                |                 2 |
| Inactive                                              |                 3 |
| Suspended                                             |                 4 |
| Blacklisted                                           |                 5 |
| Revoked                                               |                 6 |
| Recovery initiated for Blacklisted accounts           |                 7 |

### Node Status types
The table below indicates the numeric value for various node status.

| NodeStatus                                        |           Value |
| :-----------------------:                         | :-------------: |
| NotInList                                         |               0 |
| PendingApproval                                   |               1 |
| Approved                                          |               2 |
| Deactivated                                       |               3 |
| Blacklisted                                       |               4 |
| Recovery initiated for Blacklisted Node           |               5 |