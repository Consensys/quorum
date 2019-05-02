## Enclave

### What is an enclave?

An enclave is a secure processing environment that acts as a black box for processing commands and data. Enclaves come in various forms, some on hardware and others in software. In all scenarios, the purpose is to protect information that exists inside of the enclave from malicious attack.

### What does a Tessera enclave do?

The Tessera enclave is designed to handle all of the encryption/decryption operations required by the Transaction Manager, as well as all forms of key management.

This enables all sensitive operations to be handled in a single place, without any leakage into areas of program memory that don't need access. This also means that a smaller application can be run in a secure environment, where memory constraints are often more stringent, such as hardware enclaves.

The Transaction Manager, which handles peer management and database access, as well as Quorum communication does not perform **any** encryption/decryption, greatly reducing the impact an attack can have.

### What exactly does the enclave handle?

The Tessera enclave **handles** the following data:

- public/private key access
- public keys of extra recipients (** should be moved into Transaction Manager, not sensitive)
- default identity of attached nodes

The enclaves **performs** the following actions on request:

- fetching the default identity for attached nodes (default public key)
- providing forwarding keys for all transactions (** should be moved into Transaction Manager, not sensitive)
- returning all public keys managed by this enclave
- encrypting a payload for given sender and recipients
- encrypting raw payloads for given sender
- decrypting transactions for a given recipient (or sender)
- adding new recipients for existing payloads

### Where does the Enclave sit in the private transaction flow?

The Enclave is the innermost actor of the sequence of events. The below diagram demonstrates where the enclave sits:

![Quorum Tessera Privacy Flow](https://github.com/jpmorganchase/tessera/raw/master/Tessera%20Privacy%20flow.jpeg)

As the diagram shows, each enclave interacts only with it's own transaction manager and no-one else.

Tessera provides different types of Enclaves to suit different needs:

### Types of Enclave

#### Local enclave
The local enclave is the classical option that was included in versions of Tessera prior to v0.9. This includes the enclave inside the same process and the transaction manager. This is still an option, and requires all the enclave configuration to be inside the same configuration file and the Transaction Manager configuration.

##### How to use?
In order to use the local enclave, you simply need to not specify an Enclave server type in the configuration. don't forget to specify the enclave config in the Transaction Manager config file.


#### HTTP Enclave
The HTTP Enclave is a remote enclave that serves RESTful endpoints over HTTP. This allows a clear separation of concerns for between the Enclave process and Transaction Manager (TM) process. The enclave must be present and running at TM startup as it will be called upon for initialisation.

##### How to use?
The HTTP enclave can be started up by specifying an `ENCLAVE` server app type, with REST as the communication type. This same configuration should be put into the TM configuration so it knows where to find the remote enclave. Remember to set TLS settings as appropriate, with the TM being a client of the Enclave.

##### Advantage?
The HTTP enclave could be deployed in a completely secure environment away from local machine where TM process runs and it adds this additional layer of security for private keys which is only accessible from HTTP enclave.


### Setting up an Enclave

### Configuration

The configuration for the enclave is designed to the same as for the Transaction Manager.

#### Local Enclave Setup
The following should be present in the TM configuration:
```json
{
    "keys": {
        "keyData": [{
            "privateKey": "yAWAJjwPqUtNVlqGjSrBmr1/iIkghuOh1803Yzx9jLM=",
            "publicKey": "/+UuD63zItL1EbjxkKUljMgG8Z1w0AJ8pNOR4iq2yQc="
        }]
    },

    "alwaysSendTo": []
}
```
 
#### Remote Enclave Setup
The configuration required is minimal, and only requires the following from the main config (as an example):

In the remote enclave config:
```json
{
    "serverConfigs": [{
        "app": "ENCLAVE",
        "enabled": true,
        "serverAddress": "http://localhost:8080",
        "communicationType": "REST",
        "bindingAddress": "http://0.0.0.0:8080"
    }],

    "keys": {
        "keyData": [{
            "privateKey": "yAWAJjwPqUtNVlqGjSrBmr1/iIkghuOh1803Yzx9jLM=",
            "publicKey": "/+UuD63zItL1EbjxkKUljMgG8Z1w0AJ8pNOR4iq2yQc="
        }]
    },

    "alwaysSendTo": []
}
```

and in the TM configuration:
```json
"serverConfigs": [{
    "app": "ENCLAVE",
    "enabled": true,
    "serverAddress": "http://localhost:8080",
    "communicationType": "REST"
}],
```
The keys are the same as the Transaction Manager configuration, and can use all the key types including vaults.  When using a vault with the enclave, be sure to include the corresponding jar on the classpath, either:

* `/path/to/azure-key-vault-0.9-SNAPSHOT-all.jar`
* `/path/to/hashicorp-key-vault-0.9-SNAPSHOT-all.jar`

If using the all-in-one Transaction Manager jar, all the relevant files are included, and just the configuration needs to be updated for the TM.

If using the individual "make-your-own" jars, you will need the "core Transaction Manager" jar along with the "Enclave clients" jar, and add them both to the classpath as such: `java -cp /path/to/transactionmanager.jar:/path/to/enclave-client.jar com.quroum.tessera.Launcher -configfile /path/to/config.json`
