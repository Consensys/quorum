### Monitoring
Monitoring a network for security events start from log collection activity. Quorum network as any other network should produce logs, that must be analyzed for anomalies.
The following parameters are of interest to be collected and analyzed:

 - Hosts access and events
 - Ethereum accounts on the network
 - Active ledger, transaction manager nodes in the network
 - Public and Private transaction rates per account in the network. 
 - Number of public Smart contracts in the network.
 - Network connections to ledger nodes and metadata.
 - Consensus protocol metadata (E.g Block creation rate, and source ...etc)

### Security Checklist

!!! success "Ensure all activities of Quorum hosts are being logged to centralized log system"

!!! success "Centralized log system most be able to provide query capabilites over the following parameters:"
    - Ethereum accounts on the network
    - Active ledger, transaction manager nodes in the network
    - Public and Private transaction rates per account in the network.
    - Number of public Smart contracts in the network.
    - Network connections to ledger nodes and metadata.
    - Consensus protocol metadata (E.g Block creation rate, and source ...etc)

!!! success "Logs must be backed-up and integrity verified. "

!!! success "An alerting system should be put in place in order to monitor consensus protocol anomalies "

