**Quorum clients** is a thick-client whose Private Transaction feature operation depends on a Transaction Manager Client that encrypts and decrypts 
private transactions payload. Both Quorum client and its dependencies Transaction Manager, Peers, and Enclave uses traditional TCP/UDP transport layer to communicate.  

As any asset in a network its security depends on multiple elements  (E.g the security of the Host, Data, and Accounts). In Quorum it will be the security of 
the Client and Transaction Manager host/host-runtime, encryption keys, Consensus runtime and Network Access Controls.

### Host Security
As any asset in a Quorum network (Client Host, Transaction Manager Host, Private Transaction Storage Host, ..etc ) must be hardened following industry best practices. A host IDS should be used to detect any malicious activities on the host. Direct access to the host should not be allowed, instead a jump server should be used and access allowed only limited number of administrators. 
Operation systems, software and services will have vulnerabilities Quorum network hosts must implement a robust patch management program.  

### Client Security 
Quorum client instance exposes a Json-Remote Produce Call (RPC) interface through HTTP, Web Socket, or Inter-Process communication techniques. The Json-RPC interfaces
allows the remote interaction with the ledger features, and Smart Contracts. The JRPC interface must be secured in order to preserve the integrity of the ledger runtime.

Each client in the network must be uniquely identified. In Quorum this is done by using nodes identity. Node identity is represented through a public key/private key, where
the public key identifies the node in the network. Quorum Smart Contract Permissioning models depends on nodes identity to authorize TCP level communication between nodes, as result securing 
the private key of a node is a paramount activity required to prevent unauthorized node from joining the network.

 
### Users Security
Blockchain technology uses public key cryptography to protect the integrity of transactions and blocks. The security of a user’s Private keys is dependent on the security operation elements implemented to 
preserve the Private key from compromise. In Ethereum Accounts Private keys are encrypted with user specified seed (password). Users password should never be saved across the ecosystem or stored in ledger host in any form.
 
### Security Checklist
    Host:
    - Harden Quorum Host Opertion System (e.g remove irrelevant services, root access...etc).
    - Disable direct remote network access to Quorum host management interface in production.
    - Use Host Based Intrusion Detection System (HIDS) to monitoring Quorum node host.
    - Enable Host Based Firewall Rules that enforces network access to JRPC interface to only a preidentified, trusted and required systems.
    - Implement a robust Patch Management Prgoram, and always keep the host updated to latest stable version.
    - Ensure host level Isolation of responsability between Quorum client and its dependency (e.g do not run the transaction manager and its database in the same host) 
    - Ensure Quorum network hosts run with appropiate service level agreement (SLA) that can ensure a defense against nonvulnerability based denail of service.
        
    Client:
    - Enable Secure Transport Security (TLS) to encrypt all communications from/to JRPC interface to prevent data leakage and man in the middle attacks (MITM).
    - Enable Quorum Enterprise JRPC authorization model to enforce atomic access controls to ledger modules functionalities (e.g personal.OpenWallet).
    - Implement a robust Patch Management Prgoram, and always keep the client updated to latest stable version.
    - Ensure Quorum client run configuration is not started with unlocked accounts options.
    - Ensure cross domain access of the JRPC interface is configured appropiately.  
    - Ensure peer discovery is appropiately set based on the consotrium requierments.
    - In Raft Based Consensus there is no guarantee a leader would not be acting maliciously, as result raft 
    should not be used in enviorment where network ledger is managed by third party authorities.
    - Quorum clients must run with metrics collection capability in order to perserve opertional security.
    
    Users:
    - Accounts Private Key encryption password should never be stored in the ledger host in any form.
    - In an architecture where accounts private keys are not offloaded to ledger node clients, the encrypted private keys should be backedup to secure environment regularly. 
    
         