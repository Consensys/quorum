# Transaction Manager key management at Organization level
Currently when private transactions are done Quorum, the individual transaction manager (Tessera or Constellation) public keys have to be mentioned in the `privateFor` attribute. This feature allows multiple transaction manager keys to be grouped under a single organization name and at transaction level, the organization id can be passed in `privateFor` attribute instead of the keys.

Further, it will be possible to define a hierarchy of master organization and multiple sub organizations under the master org. e.g. There can be a master org "ABC" having 10 nodes and hence 10 keys. However there may be subset of nodes which are participating in various private transactions. These subsets can be set up as suborgs with in the master org with each suborg having a distincy identifier. While sending the private transaction, the suborg identifier can be give as a part of `privateFor` attribute.


## Set up
Organization level key management is  managed by a smart contract [Clusterkeys.sol](../controls/cluster/Clusterkeys.sol). The precompiled smart contract is deployed at address `0x000000000000000000034` in network bootup process. The binding of the precompiled byte code with the address is in `genesis.json`. 

## Organization level key management
The following apis are available for managing the organization level keys:
* For adding a new master organization, `quorumOrgMgmt.addMasterOrg` api to be used. The usage is as shown below. The input to this api are the name of the master org and transaction object.
```
> quorumOrgMgmt.addMasterOrg("ABC", {from:eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
```
* For adding a new sub org under a master org, `quorumOrgMgmt.addSubOrg` api to be used. The usage is as shown below. The input to this api are - unique sub org name, masgter org to which it belongs and transaction object. The call will fail if the master org is not existing. It should be noted that the sub org identifiers are unique across master organizations. 
```
> quorumOrgMgmt.addSubOrg("ENTITY1", "ABC", {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumOrgMgmt.addSubOrg("ENTITY1", "DEF", {from: eth.accounts[0]})
{
  msg: "Master org does not exist. Add master org first.",
  status: false
}
```
* Voters for approving the key management activity can be added at master org level using api `quorumOrgMgmt.addVoter`. This api takes the master org name, account to be added voter and transaction object as input. To check the list of voters for any master organizations, `quorumOrgMgmt.getOrgVoterList` api can be used. This api takes the master org identifier as input. To remove an account from the voter list for a master org `quorumOrgMgmt.getOrgVoterList` can be used. This takes master org name, account to be added voter and transaction object as input. Further adding a key or removing a key from an sub org can be performed only when there are voter accounts for the organization. If there are no voter accounts, system will allow the key add or delete operation. 
```
> quorumOrgMgmt.addVoter("ABC", eth.accounts[0], {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumOrgMgmt.getOrgVoterList("ABC")
["0xed9d02e382b34818e88B88a309c7fe71E65f419d"]
> quorumOrgMgmt.deleteVoter("ABC", eth.accounts[0], {from: eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumOrgMgmt.getOrgVoterList("ABC")
[]
> quorumOrgMgmt.addOrgKey("ENTITY1", "1iTZde/ndBHvzhcl7V68x44Vx7pl8nwx9LqnM/AfJUg=", {from: eth.accounts[0]})
{
  msg: "No voter account registered. Add voter first",
  status: false
}
```
* For adding a tansaction manager key to a sub org, the api `quorumOrgMgmt.addOrgKey` can be used. This api takes, the sub org name, transaction manager public key and transaction object as input. It should be noted that, the same key can be used across multiple sub org with in a master org. However the same key cannot be used across master organizations. If the key is already used with in another master org, the key add operation will fail. Once the key add operation is successful, it goes into a pending approval state and awaits approval from the voters. To check the pending operations at an sub org level, `quorumOrgMgmt.getPendingOpDetails`. System will not allow new key add or delete of an existing key if there are any pending approval activities at the sub org level. 
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
> quorumOrgMgmt.getPendingOpDetails("ENTITY1")
{
  pendingKey: "BULeR8JyUWhiuuCMU/HLA0Q5pzkYT+cHII3ZKBey3Bo=",
  pendingOp: "Add"
}
> quorumOrgMgmt.addOrgKey("ENTITY1", "QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc=", {from: eth.accounts[0]})
{
  msg: "Pending approvals for the organization. Approve first",
  status: false
}
```
* The voter accounts can approve any pending approval activities by using `quorumOrgMgmt.approvePendingOp`. This api takes the sub org identifier and the transaction object. 
```
> quorumOrgMgmt.approvePendingOp("ENTITY1", {from:eth.accounts[0]})
{
  msg: "Action completed successfully",
  status: true
}
> quorumOrgMgmt.approvePendingOp("ENTITY1", {from:eth.accounts[0]})
{
  msg: "Nothing to approve",
  status: false
}
>
```
* To get the list of all master orgs, sub orgs and the keys at each sub org level, `quorumOrgMgmt.orgKeyInfo` api can be used. This api outputs the list of all master orgs, sub orgs and keys at each sub org level. 
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
