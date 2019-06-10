### Tessera 
[Tessera](https://github.com/jpmorganchase/tessera/wiki) is Quorum's Transaction Manager.  Quorum privacy features depends on Tessera to Encrypt/Decrypt, and broadcast the orchestrations of a private transaction payload. 
Tessera uses an enclave to perform the encryption/decryption of private transactions payload. The encryption keys should be stored in high secure environments such a hardware security module (HSM).
Tessera communication with its dependencies (Enclave, Quorum node, Payload Storage Database, Secret Storage Service) must be secured. To ensure the privacy and authentication of the communication between Tessera the network must be configured to Certificate Based Mutual Authentication (MTLS).

### Encryption Keys
Encryption keys is the most critical element of the privacy model, if the encryption key is compromised the network loses its privacy. Tessera support integration with Trusted Platform Modules (TPM) and Hardware Security Modules (HSM) to reduce surface attack and provide highly secure environment.


### Security Checklist 
    
    - Tessera should run in independent network segment in production
    - Tessera must leverage certificate based mutual authentication with its dependencies
    - Secret storage services must support key rotation.
    - Depending on the deployment model Encryption Keys must be backed-up in offline secured locations.
    - Secret storage service  must be in complete isolation of external network.
    - Tessera connection strings must not be stored in clear text in configuration files. 
    - Secret storage in cloud deployment should run under a single tenancy model.
    - Host firewall should be enabled, inbound and outbound traffic should be limited to only vault services and restricted to consumers of those services. This includes essential host services like DNS, and NTP.
    - Restrict remote access to Secret Storage instance to whitelisted IP addresses and enable MFA.
    - Disable remote root access to Tessera/Secret storage hosts.
    - Enable remote centralized logging for tessera and its dependencies.
    - Disable core dumps in tessera host.
    - Tessera upgrades should be using immutable strategy and frequent.
    
    
