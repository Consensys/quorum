**Smart Contracts Security** must be considered as any other application security, it might contain logical vulnerabilities, insecure design and it might run on vulnerable components (ledgers). However Smart Contracts
are the core element of Ethereum Blockchain, unlike other software concepts it is constrained by several Blockchain Technology primitives:

- Smart Contracts runtime is sandboxed, this means obtaining a secure randomness source is hard. 
- Smart Contracts can hold, transfer or destroy funds making them an economical risk component.
- Smart Contracts cannot be directly upgraded unless an upgrade mechanism is introduced from the design phase.
- Smart Contracts are immutable, and have an irrevocable self destruction feature.  
- Smart Contracts can have one or multiple owners.

### Ownership
Unlike traditional software management process Smart Contracts support the following technologically enforced ownership model:

Single Ownership:
The contract has one owner who is responsible for the contract administration process. 

Shared Custody Ownership:
Suitable for agreement between two or more parties in a network of N parties, where any party can unilaterally perform administrative action over the contract.

Consortium Based Ownership:
Is a form of expanded Shared Custody Ownership that requires consensus over the administrative actions. 


### Security Patterns:

**Checks-Effects-Interaction Pattern** Interacting with other contracts should always be the last step in contract function. It’s crucial that the current contract has finished its functionality before handling control to other contract and does not depend on the execution of the other contract. 

**Circuit Breaker** is logical emergency stop execution logic. Implementing emergency stops in logic of smart contract is a good security practice. A Circuit breaker can be triggered manually by trusted parties included in the contract like the contract owner or by using programmatic consensus rules that automatically trigger the circuit breaker when the defined conditions are met.

**Rate Limit** smart contract function within a period of time allows better control of resources that can be abused.

**Speed Bumps** introduces a delay in the action execution allowing a time to act if action is considered malicious. 


### Common Contract Vulnerabilities 

**Reentrancy**: Reentrancy occurs when external contract calls are allowed to make new calls to the calling contract before the initial execution is complete. For a function, this means that the contract state may change in the middle of its execution as a result of a call to an untrusted contract or the use of a low level function with an external address.

**Access Control**: While insecure visibility settings give attackers straightforward ways to access a contract's private values or logic, access control bypasses are sometimes more subtle. These vulnerabilities can occur when contracts use the deprecated tx.origin to validate callers, handle large authorization logic with lengthy require and make reckless use of delegatecall in proxy libraries or proxy contracts.

**Arithmetic**: Integer overflows and underflows are not a new class of vulnerability, but they are especially dangerous in smart contracts, where unsigned integers are prevalent and most developers are used to simple int types (which are often just signed integers). If overflows occur, many benign-seeming codepaths become vectors for theft or denial of service.

**Unchecked Low Level Calls**: One of the deeper features of Solidity are the low level functions such as call(), callcode(), delegatecall() and send(). Their behavior in accounting for errors is quite different from other Solidity functions, as they will not propagate (or bubble up) and will not lead to a total reversion of the current execution. Instead, they will return a boolean value set to false, and the code will continue to run. This could surprise developers and, if the return value of such low-level calls are not checked, it could lead to fail-opens and other unwanted outcomes
 
**Bad Randomness**: Is hard to get right in Ethereum. While Solidity offers functions and variables that can access apparently hard-to-predict values, they are generally either more public than they seem. Because randomness sources are to an extent predictable in ethereum, malicious users can generally replicate it and attack the function relying on its unpredictablility and this applies to dApps built on top of Quorum too.

**Front Running**: In public ethereum client miners always get rewarded via gas fees for running code on behalf of externally owned addresses (EOA), users can specify higher fees to have their transactions mined more quickly. Since the Ethereum blockchain is public, everyone can see the contents of others' pending transactions. This means if a given user is revealing the solution to a puzzle or other valuable secret, a malicious user can steal the solution and copy their transaction with higher fees to preempt the original solution. If developers of smart contracts are not careful, this situation can lead to practical and devastating front-running attacks. On the other hand since Quorum does not uses Proof Of Work (PoW) in its consensus algorithm and the gas cost is zero, PoW mining related vulnerabilities are not applicable when building on top of Quorum, however fron-running remains a risk and will dependent on the consensus algorithm in use.

**Time Manipulation**: From locking a token sale to unlocking funds at a specific time, contracts sometimes need to rely on the current time. This is usually done via block.timestamp or its alias now in Solidity. In public ethereum this value comes from the miners, however in Quorum it comes from the minter, or validators, as result smart contracts should avoid relying strongly on the block time for critical decision making. Note that block.timestamp should not be used for the generation of random numbers.

**Short Addresses**: attacks are a side-effect of the EVM itself accepting incorrectly padded arguments. Attackers can exploit this by using specially-crafted addresses to make poorly coded clients encode arguments incorrectly before including them in transactions. 


### Security Checklist

#### Ownership

!!! success "No ownership contacts must be prevented in Enteprise Blockchain."

!!! success "Contracts must include initilization phase where all owners are clearly identified and set `init(owners_list)`."

!!! success "Identify contract ownership model before starting the design of the smart contract logic."

!!! success "Define the consensus model for Constortium Based Ownership."

!!! success "Contract upgradability and ownership functionalies must verify new addresses are valid."

!!! success "Ownership relate events must be broadcasted to all the network participants."

!!! success "In a Constortium based ownership structure changing activities that are bound to approval from Constortium members before they are commited (E.g editing Constortium structure) must have an aproval pending expiration date."

!!! success "Constorium based voting must involve realtime notification through EVM event emition. "

#### Contract Implementation

!!! success "Contract should use a locked compiler version"

!!! success "The compiler version should be consistent across all contracts"

!!! success "Contract should not shadow or overwrite built-in functions"

!!! success "Contract should never use tx.origin as authorization mechanism "

!!! success "Contract should never use timestamp as source of randomness"

!!! success "Contract should never use block number or hash as a source for randomness"

!!! success "Contract should never use block number or timestamp as critical decision making conditions"

!!! success "Contract should never misuse multiple inheritance "

!!! success "Modifiers must perserve the contract state or performing an external call"

!!! success "Contract should never contain cross function race conditions"

!!! success "Contract should never use plain arithmetic computation instead safe math should be used"

!!! success "Contract fallback functions should be free of unknown states that might introduce security implications"

!!! success "Contract should avoid shadowed variables "

!!! success "Contract public variables/functions should be reviewed to ensure visibility is appropriate"

!!! success "Contract private variables should not contain sensitive data "

!!! success "Contract functions should explicitly declare visibility"

!!! success "Contract public functions should perform proper authorization checks "

!!! success "Contract should validate the input of all public and external functions"

!!! success "Contract using old solidity version constructor name must match contract name"

!!! success "Contract should explicitly mark untrusted contracts as 'untrusted'"

!!! success "Contract functions logic should perform state changing actions before making external calls"

!!! success "Contract logic should use send() and transfer() over call.value when possible"

!!! success "Contract usage of delegatecall should be properly handled"

!!! success "Contract logic must correctly handle return value of any external call"

!!! success "Contract must never assume it has been created with balance of 0 "

!!! success "Contract logic should not contain loops that are vulnerable to denial of service attacks"

!!! success "Multiparty contract logic action should not be dependent on a single party"

!!! success "Prevent Toke transfers to 0x0 address"
