# Introduction
The [current permission model](../Old%20Permissioning.md) with in Quorum is limited to node level permissions only and allows a set of nodes which are part of `permissioned-nodes.json` to join the network. The model has been enhanced to cater for enterprise level needs to have **smart contract based permissions model** which allows flexibility to manage nodes, accounts and account level access controls. The overview of the model is as depicted below:
![permissions mode](images/PermissionsModel.png)  
### Key Definitions
* Network - A set of interconnected nodes representing an enterprise blockchain which contains organizations
* Organization - A set of Ethereum accounts, nodes having varied permissions to interact with the network
* Sub Organization - Further sub grouping with in the Organization as per business need
* Account - An Ethereum account which is an EOA (Externally Owned Account)
* Voter - An account capable of voting for a certain action
* Role - A named job function in organization
* Node - A `geth` node which is part of the network and belongs to an organization or sub organization

As depicted above, in the enhanced permissions model, the network comprises of group of organization. The network admin accounts defined at network level can propose and approve new organizations to join the network and assign an account as the organization administration account. The organization admin account can in turn creates roles, sub organizations, assign roles to its accounts and add any other node which is part of the organization. The sub organizations in turn can have its own set of roles accounts and sub organizations. The organization administration account manages and controls all activities at the organization level. The organization administrator can create an admin role and assign the same to a different account to manage the administration of a sub organization. The access rights of an account are derived based on the role assigned to it. The account will be able to transact via any node linked to the sub org or at overall organizations level.  

A sample network view is as depicted below:
![sample mode](images/sampleNetwork.png)
