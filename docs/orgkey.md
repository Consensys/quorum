# Transaction Manager key management at Organization level
For sending private transactions in Quorum, the individual transaction manager (Tessera or Constellation) public keys have to be mentioned in the `privateFor` attribute. If the private transaction is intended for multiple nodes, this sometimes becomes challenging to manage. This feature allows multiple transaction manager keys to be grouped under a single organization name. The organization name can then be used in `privateFor` attribute instead of individual transaction manager keys. 

Further this feature allows to define a hierarchy of master organization and multiple sub organizations under the master org. e.g. There can be a master org "ABC" having 10 nodes and hence 10 keys. However there may be subset of nodes which are participating in various private transactions. These subsets can be set up as suborgs with in the master org with each suborg having a distincy identifier. While sending the private transaction, the suborg identifier can be give as a part of `privateFor` attribute.

## Set up
Organization level key management is  managed by a smart contract [Clusterkeys.sol](../controls/cluster/Clusterkeys.sol). The precompiled smart contract is deployed at address `0x000000000000000000022` in network bootup process. The binding of the precompiled byte code with the address is in `genesis.json`. 

## APIs for organization level key management
### quorumOrgMgmt.addMasterOrg
* Input: saster org id and transaction object. The master org name has to be unique
* Output: status of operation
* Example:
```
> quorumOrgMgmt.addMasterOrg("ABC", {from:eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
### quorumOrgMgmt.addSubOrg
* Input: Sub org id, master org id for the sub org and transaction object. The sub org name is unique across master organizations
* Output: status of operation
* Example:
```
> quorumOrgMgmt.addSubOrg("ENTITY1", "ABC", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
### quorumOrgMgmt.addVoter
Before any keys can be added to a sub org id, voters need to be added at master org level to which the sub org is linked. This API is used for adding a voter to the master org. Only an account with full access can add an account as voter. Further the account being added as voter account should have at least transact permission.
* Input: master org id, voter account id, transaction object 
* Output: status of operation
* Example:
```
> quorumOrgMgmt.addVoter("ABC", eth.accounts[0], {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
### quorumOrgMgmt.removeVoter
This API is used for removing a voter to the master org. Only an account with full access can perform this activity.
* Input: master org id, voter account id, transaction object 
* Output: status of operation
* Example:
```
> quorumOrgMgmt.removeVoter("ABC", eth.accounts[0], {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
### quorumOrgMgmt.getOrgVoterList
* Input: master org id
* Output: list of voters accounts for the master org
* Example:
```
> quorumOrgMgmt.getOrgVoterList("ABC")
["0xed9d02e382b34818e88B88a309c7fe71E65f419d"]
```
### quorumOrgMgmt.addOrgKey
For adding a key to a sub org, there should be valid voters at master org level to which the sub org belongs. Further the key should not be in use in any of the other master orgs. Onec the key is added successfully, it goes into pending approval status and awaits approval from voters at master org level.
* Input: sub org id, transaction manager public key, transaction object
* Output: status of the operation
* Example:
```
> quorumOrgMgmt.addOrgKey("ENTITY1", "BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=", {from:eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumOrgMgmt.addOrgKey("ENTITY1", "1iTZde/ndBHvzhcl7V68x44Vx7pl8nwx9LqnM/AfJUg=", {from: eth.accounts[0]})
{
  msg: "Key already in use in another master organization",
  status: false
}
```
### quorumOrgMgmt.getPendingOpDetails
* Input: sub org id
* Output: pending operation(add/delete), key 
* Example:
```
> quorumOrgMgmt.getPendingOpDetails("ENTITY1")
{
  pendingKey: "BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=",
  pendingOp: "Add"
}
```
### quorumOrgMgmt.approvePendingOp
Any valid voter account at master org level can invoke this API to approve the pending key add/delete operation. 
* Input: sub org id
* Output: status of the activity
* Example:
```
> quorumOrgMgmt.approvePendingOp("ENTITY1", {from:eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
### quorumOrgMgmt.orgKeyInfo
* Input: none
* Output: list of all master org ids, sub org ids and keys at each sub org id level
* Example:
```
> quorumOrgMgmt.orgKeyInfo
[{
    masterOrgId: "FFF",
    subOrgId: "FFF1",
    subOrgKeyList: []
}, {
    masterOrgId: "DEF",
    subOrgId: "ENTITY3",
    subOrgKeyList: ["1iTZde/ndBHvzhcl7V68x44Vx7pl8nwx9LqnM/AfJUg="]
}, {
    masterOrgId: "ABC",
    subOrgId: "ENTITY1",
    subOrgKeyList: ["BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo="]
}, {
    masterOrgId: "ABC",
    subOrgId: "ENTITY2",
    subOrgKeyList: ["QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", "BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo"]
}]
```

## Sending private transaction to sub orgs
Using [simplestore.sol](https://github.com/jpmorganchase/quorum-examples/blob/master/examples/7nodes/simplestorage.sol) as example, if one has to deploy this as a private contract between node1 and node2 - the deployment command will be as below:
```
a = eth.accounts[0]
web3.eth.defaultAccount = a;

var abi = [{"constant":true,"inputs":[],"name":"storedData","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"uint256"}],"name":"set","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[{"name":"initVal","type":"uint256"}],"payable":false,"type":"constructor"}];

var bytecode = "0x6060604052341561000f57600080fd5b604051602080610149833981016040528080519060200190919050505b806000819055505b505b610104806100456000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a1afcd914605157806360fe47b11460775780636d4ce63c146097575b600080fd5b3415605b57600080fd5b606160bd565b6040518082815260200191505060405180910390f35b3415608157600080fd5b6095600480803590602001909190505060c3565b005b341560a157600080fd5b60a760ce565b6040518082815260200191505060405180910390f35b60005481565b806000819055505b50565b6000805490505b905600a165627a7a72305820d5851baab720bba574474de3d09dbeaabc674a15f4dd93b974908476542c23f00029";

var simpleContract = web3.eth.contract(abi);
var simple = simpleContract.new(42, {from:web3.eth.accounts[0], data: bytecode, gas: 0x47b760, privateFor: ["QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc="]}, function(e, contract) {
  if (e) {
    console.log("err creating contract", e);
  } 
  else {
    if (!contract.address) {
            console.log("Contract transaction send: TransactionHash: " + contract.transactionHash + " waiting to be mined...");
    } else {
            console.log("Contract mined! Address: " + contract.address);
            console.log(contract);
    }
  }
});
```
In the above deployment call, the transaction manager key of node2 is passed as a part of the `privateFor` argument. Now the privateFor attribute will accepts the distinct sub org identifir. The deployment script with `privateFor` value as sub org is as showb below:
```
a = eth.accounts[0]
web3.eth.defaultAccount = a;

var abi = [{"constant":true,"inputs":[],"name":"storedData","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"x","type":"uint256"}],"name":"set","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"get","outputs":[{"name":"retVal","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[{"name":"initVal","type":"uint256"}],"payable":false,"type":"constructor"}];

var bytecode = "0x6060604052341561000f57600080fd5b604051602080610149833981016040528080519060200190919050505b806000819055505b505b610104806100456000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632a1afcd914605157806360fe47b11460775780636d4ce63c146097575b600080fd5b3415605b57600080fd5b606160bd565b6040518082815260200191505060405180910390f35b3415608157600080fd5b6095600480803590602001909190505060c3565b005b341560a157600080fd5b60a760ce565b6040518082815260200191505060405180910390f35b60005481565b806000819055505b50565b6000805490505b905600a165627a7a72305820d5851baab720bba574474de3d09dbeaabc674a15f4dd93b974908476542c23f00029";

var simpleContract = web3.eth.contract(abi);
var simple = simpleContract.new(42, {from:web3.eth.accounts[0], data: bytecode, gas: 0x47b760, privateFor: ["ENTITY1"]}, function(e, contract) {
    if (e) {
        console.log("err creating contract", e);
    } else {
        if (!contract.address) {
            console.log("Contract transaction send: TransactionHash: " + contract.transactionHash + " waiting to be mined...");
        } else {
            console.log("Contract mined! Address: " + contract.address);
            console.log(contract);
        }
    }
});
```
