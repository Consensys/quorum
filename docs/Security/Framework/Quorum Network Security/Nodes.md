**Quorum client** is a thick-client whose Private Transaction feature operation depends on a Transaction Manager Client that encrypts and decrypts 
private transactions payload. Both Quorum client and its dependencies i.e, Transaction Manager, Peers, and Enclave use traditional TCP/UDP transport layer to communicate.  

As any asset in a network its security depends on multiple elements  (E.g the security of the Host, Data, and Accounts). In Quorum it will be the security of 
the Client and Transaction Manager host/host-runtime, encryption keys, Consensus runtime and Network Access Controls.

### Host Security
Any asset in a Quorum network (Client Host, Transaction Manager Host, Private Transaction Storage Host, ..etc ) must be hardened following industry best practices. A host IDS should be used to detect any malicious activities on the host. Direct access to the host should not be allowed, instead a jump server should be used and access limited to small number of administrators. 
Operating systems, software and services will have vulnerabilities. Quorum network hosts must implement a robust patch management program.  

### Client Security 
Quorum client instance exposes a JSON-Remote Procedure Call (RPC) interface through HTTP, Web Socket, or Inter-Process communication techniques. The JSON-RPC interfaces
allows the remote interaction with the ledger features, and Smart Contracts. The JRPC interface must be secured in order to preserve the integrity of the ledger runtime.

Each client in the network must be uniquely identified. In Quorum this is done by using nodes identity. Node identity is represented through a public key/private key, where
the public key identifies the node in the network. Quorum Smart Contract Permissioning models depends on nodes identity to authorize TCP level communication between nodes, as such securing 
the private key of a node is a paramount activity required to prevent unauthorized node from joining the network.

 
### Users Security
Blockchain technology uses public key cryptography to protect the integrity of transactions and blocks. The security of a user’s Private keys is dependent on the security operation elements implemented to 
preserve the Private key from compromise. In Ethereum Accounts Private keys are encrypted with user specified seed (password). Users password should never be saved across the ecosystem or stored in ledger host in any form.
 
### Security Checklist

#### Host

!!! success "Harden Quorum Host Operating System (e.g remove irrelevant services, root access...etc)."

!!! success "Disable direct remote network access to Quorum host management interface in production."

!!! success "Use Host Based Intrusion Detection System (HIDS) to monitoring Quorum node host."

!!! success "Enable Host Based Firewall Rules that enforces network access to JRPC interface to only a preidentified, trusted and required systems."

!!! success "Implement a robust Patch Management Program, and always keep the host updated to latest stable version."

!!! success "Ensure host level Isolation of responsability between Quorum client and its dependency (e.g do not run the transaction manager and its database in the same host) "

!!! success "Ensure Quorum network hosts run with appropiate service level agreement (SLA) that can ensure a defense against non-vulnerability based denial of service."

#### Client

!!! success "Enable Secure Transport Security (TLS) to encrypt all communications from/to JRPC interface to prevent data leakage and man in the middle attacks (MITM)."

!!! success "Enable Quorum Enterprise JRPC authorization model to enforce atomic access controls to ledger modules functionalities (e.g personal.OpenWallet)."

!!! success "Implement a robust Patch Management Prgoram, and always keep the client updated to latest stable version."

!!! success "Ensure Quorum client run configuration is not started with unlocked accounts options."

!!! success "Ensure cross domain access of the JRPC interface is configured appropriately.  "

!!! success "Ensure peer discovery is appropriately set based on the consortium requirements."

!!! success "In Raft Based Consensus there is no guarantee a leader would not be acting maliciously, hence raft should not be used in environment where network ledger is managed by third party authorities."

!!! success "Quorum clients must run with metrics collection capability in order to preserve operational security."

#### Users

!!! success "Accounts Private Key encryption password should never be stored in the ledger host in any form."

!!! success "In an architecture where accounts private keys are not offloaded to ledger node clients, the encrypted private keys should be backed-up to secure environment regularly. "