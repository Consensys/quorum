# Set up
The steps to enable new permissions model are as described below:
* For a new network, bring up the initial set of nodes which will be part of the network
* Deploy the `PermissionsUpgradable.sol` in the network. The deployment of this contract will require a custodian account to be given as a part of deployment. 
* Deploy the rest of the contracts. All the other contracts will require the address of `PermissionsUpgradable.sol` contract as a part of deployment.
* Once all the contracts are deployed create a file `permission-config.json` which will have the following construct:
```$xslt
{
           "upgrdableAddress": "0x1932c48b2bf8102ba33b4a6b545c32236e342f34",
           "interfaceAddress": "0x4d3bfd7821e237ffe84209d8e638f9f309865b87",
           "implAddress": "0xfe0602d820f42800e3ef3f89e1c39cd15f78d283",
           "nodeMgrAddress": "0x8a5e2a6343108babed07899510fb42297938d41f",
           "accountMgrAddress": "0x9d13c6d3afe1721beef56b55d303b09e021e27ab",
           "roleMgrAddress": "0x1349f3e1b8d71effb47b840594ff27da7e603d17",
           "voterMgrAddress": "0xd9d64b7dc034fafdba5dc2902875a67b5d586420",
           "orgMgrAddress" : "0x938781b9796aea6376e40ca158f67fa89d5d8a18",
           "nwAdminOrg": "ADMINORG",
           "nwAdminRole" : "ADMIN",
           "orgAdminRole" : "ORGADMIN",
           "accounts":["0xed9d02e382b34818e88b88a309c7fe71e65f419d", "0xca843569e3427144cead5e4d5999a3d0ccf92b8e"],
           "subOrgBreadth" : "3",
           "subOrgDepth" : "4"
}
```
> * `upgrdableAddress` is the address of deployed contract `PermissionsUpgradable.sol`
> * `interfaceAddress` is the address of deployed contract `PermissionsInterface.sol`
> * `implAddress` is the address of deployed contract `PermissionsImplementation.sol`
> * `nodeMgrAddress` is the address of deployed contract `NodeManager.sol`
> * `accountMgrAddress` is the address of deployed contract `AccountManager.sol`
> * `roleMgrAddress` is the address of deployed contract `RoleManager.sol`
> * `voterMgrAddress` is the address of deployed contract `VoterManager.sol`
> * `orgMgrAddress` is the address of deployed contract `OrgManager.sol`
> * `nwAdminOrg` is the name of initial organization that will be created as a part of network boot up with new permissions model. This organization will own all the initial nodes which come at the time of network boot up and accounts which will be the network admin account
> * `nwAdminRole` is role id for which will have full access and will be network admin
> * `accounts` holds the initial list of accounts which will be linked to the network admin organization and will be assigned the network admin role. These accounts will have complete control on the network and can propose and approve new organizations into the network
> * `subOrgBreadth` indicates the number of sub organizations that any org can have
> * `subOrgDepth` indicates the maximum depth sub org hierarchy allowed in the network

* Bring down the all `geth` nodes in network and copy `permission-config.json` into the data directory of each of the node
* Bring up all `geth` nodes in `--permissioned` mode for new permissions model to take effect
