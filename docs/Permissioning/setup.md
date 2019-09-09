# Set up
The steps to enable new permissions model are as described below:

## New network

* Bring up the initial set of nodes which will be part of the network
* Deploy the `PermissionsUpgradable.sol` in the network. The deployment of this contract will require a guardian account to be given as a part of deployment. 
* Deploy the rest of the contracts. All the other contracts will require the address of `PermissionsUpgradable.sol` contract as a part of deployment.
* Once all the contracts are deployed create a file `permission-config.json` which will have the following construct:
```json
{
    "upgradableAddress": "0x1932c48b2bf8102ba33b4a6b545c32236e342f34",
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
    "subOrgBreadth" : 3,
    "subOrgDepth" : 4
}
```
> * `upgradableAddress` is the address of deployed contract `PermissionsUpgradable.sol`
> * `interfaceAddress` is the address of deployed contract `PermissionsInterface.sol`
> * `implAddress` is the address of deployed contract `PermissionsImplementation.sol`
> * `nodeMgrAddress` is the address of deployed contract `NodeManager.sol`
> * `accountMgrAddress` is the address of deployed contract `AccountManager.sol`
> * `roleMgrAddress` is the address of deployed contract `RoleManager.sol`
> * `voterMgrAddress` is the address of deployed contract `VoterManager.sol`
> * `orgMgrAddress` is the address of deployed contract `OrgManager.sol`
> * `nwAdminOrg` is the name of initial organization that will be created as a part of network boot up with new permissions model. This organization will own all the initial nodes which come at the time of network boot up and accounts which will be the network admin account
> * `nwAdminRole` is role id which will have full access and will be network admin. This role will be assigned to the network admin accounts
> * `orgAdminRole` is role id which will have full access and will manage organization level administration activities. This role will be assigned to the org admin account
> * `accounts` holds the initial list of accounts which will be linked to the network admin organization and will be assigned the network admin role. These accounts will have complete control on the network and can propose and approve new organizations into the network
> * `subOrgBreadth` indicates the number of sub organizations that any org can have
> * `subOrgDepth` indicates the maximum depth of sub org hierarchy allowed in the network

* Once the contracts are deployed, `init` in `PermissionsUpgradable.sol` need to be executed by the guardian account. This will link the interface and implementation contracts. A sample script for loading the upgradable contract at `geth` prompt is as given below
```javascript
ac = eth.accounts[0];
web3.eth.defaultAccount = ac;
var abi = [{"constant":true,"inputs":[],"name":"getPermImpl","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_proposedImpl","type":"address"}],"name":"confirmImplChange","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getGuardian","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getPermInterface","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_permInterface","type":"address"},{"name":"_permImpl","type":"address"}],"name":"init","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[{"name":"_guardian","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}];
var upgr = web3.eth.contract(abi).at("0x1932c48b2bf8102ba33b4a6b545c32236e342f34"); // address of the upgradable contracts
var impl = "0xfe0602d820f42800e3ef3f89e1c39cd15f78d283" // address of the implementation contracts
var intr = "0x4d3bfd7821e237ffe84209d8e638f9f309865b87" // address of the interface contracts
```
* At `geth` prompt load the above script after replacing the contract addresses appropriately and execute `upgr.init(intr, impl, {from: <guardian account>, gas: 4500000})`
* Bring down the all `geth` nodes in the network and copy `permission-config.json` into the data directory of each node

## Migrating from an earlier version
The following steps needs to be followed when migrating from a earlier version for enabling permissions feature

* Bring down the running network in the earlier version. 
* The `maxCodeSize` attribute in `genesis.json` need to be set to 35. Update `genesis.json` to reflect the same
```javascript
  "config": {
    "homesteadBlock": 0,
    "byzantiumBlock": 0,
    "chainId": 10,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
    "eip158Block": 0,
    "maxCodeSize" : 35,
    "isQuorum":
```
* Execute `geth --datadir <<data dir path>> init genesis.json`
* Bring up the network with latest geth and deploy the contracts as explained earlier in the set up. The rest of the steps will be similar to bringing up a new network

!!! Note
    * It should be noted that the new permission model will be in force only when `permission-config.json` is present in data directory. If this file is not there and the node is brought up with `--permissioned` flag, node level permissions as per the earlier model will be effective.
    * Please ensure that `maxCodeSize` in `genesis.json` is set to 35 
