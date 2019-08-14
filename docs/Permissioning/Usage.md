
**This section describes the usage of permission model for creation of a network, initial set up and management of network. The network management activities can be broadly categorized into:**

* [Initial network set up](#initial-network-set-up)
* [Proposing a new organization into the network](#proposing-a-new-organization-into-the-network)
* [Organization admin managing the organization level permissions](#organization-admin-managing-the-organization-level-permissions)
* [Suspending an organization temporarily](#suspending-an-organization-temporarily)
* [Revoking suspension of an organization](#revoking-suspension-of-an-organization)
* [Assigning admin privileges at organization and network level](#assigning-admin-privileges-at-organization-and-network-level)


### Initial network set up
Please refer to [set up](../setup). For an existing network running with an older version of Quorum:

* Upgrade Quorum to the latest version 
* Deploy the contracts 
* Execute the `init` method of `PermissionsUpgradable.sol` from the guardian account
* Copy the `permission-config.json`  to the data directory of each node  
* Bring `geth` up in `--permissioned` mode.

For a new network using the latest version of Quorum:

* Bring up the initial set of nodes 
* Deploy the contracts 
* Execute the `init` method of `PermissionsUpgradable.sol` from the guardian account
* Upgrade Quorum to the latest version 
* Copy the `permission-config.json`  to the data directory of each node  
* Bring `geth` up in `--permissioned` mode.

As part of network initialization:

* A network admin organization is created with the `nwAdminOrg` name specified in `permission-config.json`. All nodes which are part of `static-nodes.json` are assigned to this organization. 
* A network admin role is created with the `nwAdminRole` name specified in the config file. 
* All accounts given in the `accounts` array of the config file are assigned the network admin role. These accounts will have the ability to propose and approve new organizations into the network.

Assuming that the network was started with the `permission-config.json` given in the [set up](../setup), and assuming the network was brought up with the `static-nodes.json` file given below:
```json
[
	"enode://72c0572f7a2492cffb5efc3463ef350c68a0446402a123dacec9db5c378789205b525b3f5f623f7548379ab0e5957110bffcf43a6115e450890f97a9f65a681a@127.0.0.1:21000?discport=0",
	"enode://7a1e3b5c6ad614086a4e5fb55b6fe0a7cf7a7ac92ac3a60e6033de29df14148e7a6a7b4461eb70639df9aa379bd77487937bea0a8da862142b12d326c7285742@127.0.0.1:21001?discport=0",
	"enode://5085e86db5324ca4a55aeccfbb35befb412def36e6bc74f166102796ac3c8af3cc83a5dec9c32e6fd6d359b779dba9a911da8f3e722cb11eb4e10694c59fd4a1@127.0.0.1:21002?discport=0",
	"enode://28a4afcf56ee5e435c65b9581fc36896cc684695fa1db83c9568de4353dc6664b5cab09694d9427e9cf26a5cd2ac2fb45a63b43bb24e46ee121f21beb3a7865e@127.0.0.1:21003?discport=0"
]
```
then the network will have the following configuration once it has started up:
```
> quorumPermission.orgList
[{
    fullOrgId: "ADMINORG",
    level: 1,
    orgId: "ADMINORG",
    parentOrgId: "",
    status: 2,
    subOrgList: null,
    ultimateParent: "ADMINORG"
}]
> quorumPermission.getOrgDetails("ADMINORG")
{
  acctList: [{
      acctId: "0xed9d02e382b34818e88b88a309c7fe71e65f419d",
      isOrgAdmin: true,
      orgId: "ADMINORG",
      roleId: "ADMIN",
      status: 2
  }, {
      acctId: "0xca843569e3427144cead5e4d5999a3d0ccf92b8e",
      isOrgAdmin: true,
      orgId: "ADMINORG",
      roleId: "ADMIN",
      status: 2
  }],
  nodeList: [{
      orgId: "ADMINORG",
      status: 2,
      url: "enode://72c0572f7a2492cffb5efc3463ef350c68a0446402a123dacec9db5c378789205b525b3f5f623f7548379ab0e5957110bffcf43a6115e450890f97a9f65a681a@127.0.0.1:21000?discport=0"
  }, {
      orgId: "ADMINORG",
      status: 2,
      url: "enode://7a1e3b5c6ad614086a4e5fb55b6fe0a7cf7a7ac92ac3a60e6033de29df14148e7a6a7b4461eb70639df9aa379bd77487937bea0a8da862142b12d326c7285742@127.0.0.1:21001?discport=0"
  }, {
      orgId: "ADMINORG",
      status: 2,
      url: "enode://5085e86db5324ca4a55aeccfbb35befb412def36e6bc74f166102796ac3c8af3cc83a5dec9c32e6fd6d359b779dba9a911da8f3e722cb11eb4e10694c59fd4a1@127.0.0.1:21002?discport=0"
  }, {
      orgId: "ADMINORG",
      status: 2,
      url: "enode://28a4afcf56ee5e435c65b9581fc36896cc684695fa1db83c9568de4353dc6664b5cab09694d9427e9cf26a5cd2ac2fb45a63b43bb24e46ee121f21beb3a7865e@127.0.0.1:21003?discport=0"
  }],
  roleList: [{
      access: 3,
      active: true,
      isAdmin: true,
      isVoter: true,
      orgId: "ADMINORG",
      roleId: "ADMIN"
  }],
  subOrgList: null
}
```

### Proposing a new organization into the network
Once the network is up, the network admin accounts can then propose a new organization into the network. Majority approval from the network admin accounts is required before an organization is approved. The APIs for [proposing](../Permissioning%20apis#quorumpermissionaddorg) and [approving](../Permissioning%20apis#quorumpermissionapproveorg) an organization are documented in [permission APIs](../Permissioning%20apis)

#### Example
An example to propose and approve an organization by name `ORG1` is as shown below:

```javascript
> quorumPermission.addOrg("ORG1", "enode://de9c2d5937e599930832cecc1df8cc90b50839bdf635c1a4e68e1dab2d001cd4a11c626e155078cc65958a72e2d72c1342a28909775edd99cc39470172cce0ac@127.0.0.1:21004?discport=0", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: "0xed9d02e382b34818e88b88a309c7fe71e65f419d"})
"Action completed successfully"
```

Once the org is proposed, it will be in `Proposed` state awaiting approval from other network admin accounts. The org status is as shown below:
```javascript
> quorumPermission.orgList[1]
{
    fullOrgId: "ORG1",
    level: 1,
    orgId: "ORG1",
    parentOrgId: "",
    status: 1,
    subOrgList: null,
    ultimateParent: "ORG1"
}
```

The network admin accounts can then approve the proposed organizations and once the majority approval is achieved, the organization status is updated as `Approved`

```javascript
> quorumPermission.approveOrg("ORG1", "enode://de9c2d5937e599930832cecc1df8cc90b50839bdf635c1a4e68e1dab2d001cd4a11c626e155078cc65958a72e2d72c1342a28909775edd99cc39470172cce0ac@127.0.0.1:21004?discport=0", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: "0xca843569e3427144cead5e4d5999a3d0ccf92b8e"})
"Action completed successfully"
> quorumPermission.orgList[1]
{
    fullOrgId: "ORG1",
    level: 1,
    orgId: "ORG1",
    parentOrgId: "",
    status: 2,
    subOrgList: null,
    ultimateParent: "ORG1"
}
```

The details of the new organization approved are as below:
```javascript
> quorumPermission.getOrgDetails("ORG1")
{
    acctList: [{
        acctId: "0x0638e1574728b6d862dd5d3a3e0942c3be47d996",
        isOrgAdmin: true,
        orgId: "ORG1",
        roleId: "ORGADMIN",
        status: 2
    }],
    nodeList: [{
        orgId: "ORG1",
        status: 2,
        url: "enode://de9c2d5937e599930832cecc1df8cc90b50839bdf635c1a4e68e1dab2d001cd4a11c626e155078cc65958a72e2d72c1342a28909775edd99cc39470172cce0ac@127.0.0.1:21004?discport=0"
    }],
    roleList: [{
        access: 3,
        active: true,
        isAdmin: true,
        isVoter: true,
        orgId: "ORG1",
        roleId: "ORGADMIN"
    }],
    subOrgList: null
}
```

As can be seen from the above, as a part of approval:

* A org admin role with name as given in `orgAdminRole` in `permission-config.json` has been created and linked to the organization `ORG1`
* The account given has been linked to the organization `ORG1` and org admin role. This account acts as the organization admin account and can in turn manage further roles, nodes and accounts at organization level
* The node has been linked to organization and status has been updated as `Approved`

The new node belonging to the organization can now join the network. In case the network is running in `Raft` consensus mode, before the node joins the network, please ensure that:

*  The node has been added as a peer using `raft.addPeer(<<enodeId>>)`
*  Bring up `geth` for the new node using `--raftjoinexisting` with the peer id as obtained in the above step
 
### Organization admin managing the organization level permissions
Once the organization is approved and the node of the organization has joined the network, the organization admin can then create sub organizations, roles, add additional nodes at organization level, add accounts to the organization and change roles of existing organization level accounts. 

To add a sub org at `ORG1` level refer to [addSubOrg API](../Permissioning%20apis#quorumpermissionaddsuborg)
```javascript
> quorumPermission.addSubOrg("ORG1", "SUB1", "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0", {from: eth.accounts[0]})
"Action completed successfully"
> quorumPermission.getOrgDetails("ORG1.SUB1")
{
  acctList: null,
  nodeList: [{
      orgId: "ORG1.SUB1",
      status: 2,
      url: "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0"
  }],
  roleList: null,
  subOrgList: null
}
```

For adding a sub org the enode id is not mandatory. For the newly created sub org if the org admin desires to add an administration account, the org admin account will have to first create a role with `isAdmin` flag as `Y` and then assign this role to the account which belongs to the sub org. Once assigned the account will act as org admin at sub org level. Refer to [addNewRole API](../Permissioning%20apis#quorumpermissionaddnewrole)
```javascript
> quorumPermission.addNewRole("ORG1.SUB1", "SUBADMIN", 3, false, true,{from: eth.accounts[0]})
"Action completed successfully"
> eth.accounts[0]
"0x0638e1574728b6d862dd5d3a3e0942c3be47d996"
```

The role `SUBADMIN` can now be assigned to an account at sub org `SUB1` for making the account admin for the sub org.
```javascript
> quorumPermission.addAccountToOrg("0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0", "ORG1.SUB1", "SUBADMIN", {from: "0x0638e1574728b6d862dd5d3a3e0942c3be47d996"})
"Action completed successfully"
> quorumPermission.getOrgDetails("ORG1.SUB1")
{
  acctList: [{
      acctId: "0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0",
      isOrgAdmin: true,
      orgId: "ORG1.SUB1",
      roleId: "SUBADMIN",
      status: 2
  }],
  nodeList: [{
      orgId: "ORG1.SUB1",
      status: 2,
      url: "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0"
  }],
  roleList: [{
      access: 3,
      active: true,
      isAdmin: true,
      isVoter: false,
      orgId: "ORG1.SUB1",
      roleId: "SUBADMIN"
  }],
  subOrgList: null
}
```

The account `0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0` is now the admin for sub org `SUB1` and will be able to add roles, accounts and nodes to the sub org. It should be noted that the org admin account at master org level has the admin rights on all the sub organizations below. However the admin account at sub org level has control only in the sub org to which it is linked. 
```javascript
> quorumPermission.addNewRole("ORG1.SUB1", "TRANSACT", 1, false, true,{from: "0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0"})
"Action completed successfully"
> quorumPermission.getOrgDetails("ORG1.SUB1").roleList
[{
    access: 3,
    active: true,
    isAdmin: true,
    isVoter: false,
    orgId: "ORG1.SUB1",
    roleId: "SUBADMIN"
}, {
    access: 1,
    active: true,
    isAdmin: true,
    isVoter: false,
    orgId: "ORG1.SUB1",
    roleId: "TRANSACT"
}]
```

To add an account to an organization refer to [addAccountToOrg API](../Permissioning%20apis#quorumpermissionaddaccounttoorg).
```javascript
> quorumPermission.addAccountToOrg("0x283f3b8989ec20df621166973c93b56b0f4b5455", "ORG1.SUB1", "SUBADMIN", {from: "0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0"})
"Action completed successfully"
> quorumPermission.getOrgDetails("ORG1.SUB1").acctList

[{
    acctId: "0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0",
    isOrgAdmin: true,
    orgId: "ORG1.SUB1",
    roleId: "SUBADMIN",
    status: 2
}, {
    acctId: "0x283f3b8989ec20df621166973c93b56b0f4b5455",
    isOrgAdmin: true,
    orgId: "ORG1.SUB1",
    roleId: "TRANSACT",
    status: 2
}]
```

To suspend an account [updateAccountStatus](../Permissioning%20apis#quorumpermissionupdateaccountstatus) API can be invoked with action as 1.
```javascript
> quorumPermission.updateAccountStatus("ORG1.SUB1", "0x283f3b8989ec20df621166973c93b56b0f4b5455", 1, {from: "0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0"})
"Action completed successfully"
> quorumPermission.getOrgDetails("ORG1.SUB1").acctList
[{
    acctId: "0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0",
    isOrgAdmin: true,
    orgId: "ORG1.SUB1",
    roleId: "SUBADMIN",
    status: 2
}, {
    acctId: "0x283f3b8989ec20df621166973c93b56b0f4b5455",
    isOrgAdmin: true,
    orgId: "ORG1.SUB1",
    roleId: "TRANSACT",
    status: 1
}]
```

To revoke suspension of an account [updateAccountStatus](../Permissioning%20apis#quorumpermissionupdateaccountstatus) API can be invoked with action as 2.
```javascript
> quorumPermission.updateAccountStatus("ORG1.SUB1", "0x283f3b8989ec20df621166973c93b56b0f4b5455", 2, {from: "0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0"})
"Action completed successfully"
> quorumPermission.getOrgDetails("ORG1.SUB1").acctList

[{
    acctId: "0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0",
    isOrgAdmin: true,
    orgId: "ORG1.SUB1",
    roleId: "SUBADMIN",
    status: 2
}, {
    acctId: "0x283f3b8989ec20df621166973c93b56b0f4b5455",
    isOrgAdmin: true,
    orgId: "ORG1.SUB1",
    roleId: "TRANSACT",
    status: 2
}]
```

To [blacklist an account updateAccountStatus](../Permissioning%20apis#quorumpermissionupdateaccountstatus) API can be invoked with action as 3. Once blacklisted no further activity will be possible on the account.
```javascript
> quorumPermission.updateAccountStatus("ORG1.SUB1", "0x283f3b8989ec20df621166973c93b56b0f4b5455", 3, {from: "0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0"})
"Action completed successfully"
> quorumPermission.getOrgDetails("ORG1.SUB1").acctList

[{
    acctId: "0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0",
    isOrgAdmin: true,
    orgId: "ORG1.SUB1",
    roleId: "SUBADMIN",
    status: 2
}, {
    acctId: "0x283f3b8989ec20df621166973c93b56b0f4b5455",
    isOrgAdmin: true,
    orgId: "ORG1.SUB1",
    roleId: "TRANSACT",
    status: 5
}]
```

To [add nodes addNode ](../Permissioning%20apis#quorumpermissionaddnode) at organization and sub organization level by the org admin.
```javascript
> quorumPermission.addNode("ORG1.SUB1", "enode://eacaa74c4b0e7a9e12d2fe5fee6595eda841d6d992c35dbbcc50fcee4aa86dfbbdeff7dc7e72c2305d5a62257f82737a8cffc80474c15c611c037f52db1a3a7b@127.0.0.1:21005?discport=0", {from: "0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0"})
"Action completed successfully"
> quorumPermission.getOrgDetails("ORG1.SUB1").nodeList
[{
    orgId: "ORG1.SUB1",
    status: 2,
    url: "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0"
}, {
    orgId: "ORG1.SUB1",
    status: 2,
    url: "enode://eacaa74c4b0e7a9e12d2fe5fee6595eda841d6d992c35dbbcc50fcee4aa86dfbbdeff7dc7e72c2305d5a62257f82737a8cffc80474c15c611c037f52db1a3a7b@127.0.0.1:21005?discport=0"
}]
```

Org admin can manage the status of the nodes by using [updateNodeStatus](../Permissioning%20apis#quorumpermissionupdatenodestatus) API. To deactivate a node the API can be invoked with action 1.
```javascript
> quorumPermission.getOrgDetails("ORG1.SUB1").nodeList
[{
    orgId: "ORG1.SUB1",
    status: 2,
    url: "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0"
}, {
    orgId: "ORG1.SUB1",
    status: 3,
    url: "enode://eacaa74c4b0e7a9e12d2fe5fee6595eda841d6d992c35dbbcc50fcee4aa86dfbbdeff7dc7e72c2305d5a62257f82737a8cffc80474c15c611c037f52db1a3a7b@127.0.0.1:21005?discport=0"
}]
```

To activate the node back invoke [updateNodeStatus](../Permissioning%20apis#quorumpermissionupdatenodestatus) API with action 2.
```javascript
> quorumPermission.updateNodeStatus("ORG1.SUB1", "enode://eacaa74c4b0e7a9e12d2fe5fee6595eda841d6d992c35dbbcc50fcee4aa86dfbbdeff7dc7e72c2305d5a62257f82737a8cffc80474c15c611c037f52db1a3a7b@127.0.0.1:21005?discport=0",2, {from:"0x42ef6abedcb7ecd3e9c4816cd5f5a96df35bb9a0"})
"Action completed successfully"
> quorumPermission.getOrgDetails("ORG1.SUB1").nodeList
[{
    orgId: "ORG1.SUB1",
    status: 2,
    url: "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0"
}, {
    orgId: "ORG1.SUB1",
    status: 2,
    url: "enode://eacaa74c4b0e7a9e12d2fe5fee6595eda841d6d992c35dbbcc50fcee4aa86dfbbdeff7dc7e72c2305d5a62257f82737a8cffc80474c15c611c037f52db1a3a7b@127.0.0.1:21005?discport=0"
}]
```

To blacklist a node invoke [updateNodeStatus](../Permissioning%20apis#quorumpermissionupdatenodestatus) API with action 3. Once blacklisted the node will never be able join the network again.
```javascript
> quorumPermission.getOrgDetails("ORG1.SUB1").nodeList
[{
    orgId: "ORG1.SUB1",
    status: 2,
    url: "enode://239c1f044a2b03b6c4713109af036b775c5418fe4ca63b04b1ce00124af00ddab7cc088fc46020cdc783b6207efe624551be4c06a994993d8d70f684688fb7cf@127.0.0.1:21006?discport=0"
}, {
    orgId: "ORG1.SUB1",
    status: 4,
    url: "enode://eacaa74c4b0e7a9e12d2fe5fee6595eda841d6d992c35dbbcc50fcee4aa86dfbbdeff7dc7e72c2305d5a62257f82737a8cffc80474c15c611c037f52db1a3a7b@127.0.0.1:21005?discport=0"
}]
```

It should be noted that in the case of the `Raft` consensus mechanism, when the node is deactivated the peer id is lost and hence upon activation, the node needs to be added to Raft cluster again using `raft.addPeer` and the node should be brought up with new peer id.

Further:

* An account can transact from any of the nodes linked to org or sub org with in the same organization
* If a node is deactivated no transaction will be allowed from that node

### Suspending an organization temporarily
If there is a need to temporarily suspend all activities of an organization [updateOrgStatus](../Permissioning%20apis#quorumpermissionupdateorgstatus) API can be invoked with action 1. This can be invoked only by the network admin accounts and will reuiqre majority voting.
```javascript
> quorumPermission.updateOrgStatus("ORG1", 1, {from: "0xed9d02e382b34818e88b88a309c7fe71e65f419d"})
"Action completed successfully"
> quorumPermission.orgList[2]
{
  fullOrgId: "ORG1",
  level: 1,
  orgId: "ORG1",
  parentOrgId: "",
  status: 3,
  subOrgList: null,
  ultimateParent: "ORG1"
}
```

To approve the org, suspension majority approval from other network admin accounts is required. The api for the same is [approveOrgStatus](../Permissioning%20apis#quorumpermissionapproveorgstatus). Once approved the org status is marked as suspended.
```javascript
> quorumPermission.approveOrgStatus("ORG1", 1, {from: "0xca843569e3427144cead5e4d5999a3d0ccf92b8e"})
"Action completed successfully"
> quorumPermission.orgList[2]
{
  fullOrgId: "ORG1",
  level: 1,
  orgId: "ORG1",
  parentOrgId: "",
  status: 4,
  subOrgList: null,
  ultimateParent: "ORG1"
}
```

When the org is suspended no transaction from any of the account linked to the organization or sub organizations under it is allowed. However, the nodes linked to the organization will be active and will be syncing with the network.


### Revoking suspension of an organization
To revoke the suspension of an org [updateOrgStatus](../Permissioning%20apis#quorumpermissionupdateorgstatus) can be called with action as 2. This will require majority approval (API [approveOrgStatus](../Permissioning%20apis#quorumpermissionapproveorgstatus) with action 2). 
```javascript
> quorumPermission.updateOrgStatus("ORG1", 2, {from: "0xed9d02e382b34818e88b88a309c7fe71e65f419d"})
"Action completed successfully"
> quorumPermission.approveOrgStatus("ORG1", 2, {from: "0xed9d02e382b34818e88b88a309c7fe71e65f419d"})
"Action completed successfully"
> quorumPermission.orgList[0]
{
  fullOrgId: "ORG1.SUB1",
  level: 2,
  orgId: "SUB1",
  parentOrgId: "ORG1",
  status: 2,
  subOrgList: null,
  ultimateParent: "ORG1"
}
```

Once the revoke is approved, all accounts in the organization and sub organization will be able to transact as per role level access. 

### Assigning admin privileges at organization and network level
There may be a scenario where one of the accounts at the organization level needs to have network admin level permissions and be able to perform network admin activities. Similarly there can be a need to change the admin account at organization level. Both these activities can be performed by existing network admin accounts only, and will require majority approval from the network admin accounts. The API usage details are as below.
To assign network admin or org admin role to an account invoke [assignAdminRole](../Permissioning%20apis#quorumpermissionassignadminrole).
```javascript
> quorumPermission.assignAdminRole("ORG1", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", "ADMIN", {from: "0xed9d02e382b34818e88b88a309c7fe71e65f419d"})
"Action completed successfully"
> quorumPermission.acctList[3]
{
  acctId: "0x0638e1574728b6d862dd5d3a3e0942c3be47d996",
  isOrgAdmin: true,
  orgId: "ORG1",
  roleId: "ADMIN",
  status: 1
}
```

To approve the assignment of network admin role invoke [approveAdminRole](../Permissioning%20apis#quorumpermissionapproveadminrole) API.
```javascript
> quorumPermission.approveAdminRole("ORG1", "0x0638e1574728b6d862dd5d3a3e0942c3be47d996", {from: eth.accounts[0]})
"Action completed successfully"
> quorumPermission.acctList[4]
{
  acctId: "0x0638e1574728b6d862dd5d3a3e0942c3be47d996",
  isOrgAdmin: true,
  orgId: "ORG1",
  roleId: "ADMIN",
  status: 2
}
```

The above account can now perform all activities allowable by a network admin account and can participate in the approval process for any actions at network level.





