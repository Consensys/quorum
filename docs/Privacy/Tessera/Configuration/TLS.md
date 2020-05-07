### Usage
Communications via TLS/SSL can be enabled by setting `"tls": "STRICT"`. 

If the value is set to `"OFF"`, the rest of the SSL configuration will not be considered.

!!! warning
    If using TLS make sure to update the hostname of the node to use `https` instead of `http`

```json
{
  "sslConfig": {
    "tls": "[Authentication mode : OFF,STRICT]",
           
    "serverTrustMode": "[Possible values: CA, TOFU, WHITELIST, CA_OR_TOFU, NONE]",
    "serverKeyStore": "[Path to server keystore]",
    "serverKeyStorePassword": "[Password required for server KeyStore]",
    "serverTrustStore": "[Server trust store path]",
    "serverTrustStorePassword": "[Password required for server trust store]",
    "serverTlsKeyPath": "[Path to server TLS key path]",
    "serverTlsCertificatePath": "[Path to server TLS cert path]",
    "serverTrustCertificates": [
      "[Array of truststore certificates if no truststore is defined.]"
    ],
    
    "clientTrustMode": "[Possible values: CA, TOFU, WHITELIST, CA_OR_TOFU, NONE]",
    "clientKeyStore": "[Path to client keystore. The keystore that is used when communicating to other nodes.]",
    "clientKeyStorePassword": "[Password required for client KeyStore]",
    "clientTrustStore": "[Path to client TrustStore]",
    "clientTrustStorePassword": "[Password required for client trust store]",
    "clientTlsKeyPath": "[Path to client TLS Key]",
    "clientTlsCertificatePath": "[Path to client TLS cert]",
    "clientTrustCertificates": [
      "[Array of truststore certificates if no truststore is defined.]"
    ],
    
    "knownClientsFile": "[TLS known clients file for the server. This contains the fingerprints of public keys of other nodes that are allowed to connect to this one.]",
    "knownServersFile": "[TLS known servers file for the client. This contains the fingerprints of public keys of other nodes that this node has encountered.]",
    "generateKeyStoreIfNotExisted": "[boolean]",
    "environmentVariablePrefix": "[Prefix to uniquely identify environment variables for this particular server ssl config]"
  }
}
```

When SSL is enabled, each node will need to have certificates and keys defined for both client-side and server-side. These can be defined in multiple ways:

1. Secured & unsecured `.jks` (Java keystore) format files 
    * `serverKeyStore`, `serverKeyStorePassword`, `serverTrustStore`, `serverTrustStorePassword` 
    * `clientKeyStore`, `clientKeyStorePassword`, `clientTrustStore`, `clientTrustStorePassword`
2. `.pem` format certificate and key files
    * `serverTlsKeyPath`, `serverTlsCertificatePath`, `serverTrustCertificates`
    * `clientTlsKeyPath`, `clientTlsCertificatePath`, `clientTrustCertificates`

`.jks` files take precedence over `.pem` files if both are provided for client-side or server-side.

#### Keystores
##### Passwords
Passwords for secured `.jks` keystores can be provided in multiple ways, and in the following order of precedence:
1. *Prefixed* environment variables
    * `<PREFIX>_TESSERA_SERVER_KEYSTORE_PWD`, `<PREFIX>_TESSERA_SERVER_TRUSTSTORE_PWD`
    * `<PREFIX>_TESSERA_CLIENT_KEYSTORE_PWD`, `<PREFIX>_TESSERA_CLIENT_TRUSTSTORE_PWD`
    
2. Config file 
    * `serverKeyStorePassword`, `serverTrustStorePassword`
    * `clientKeyStorePassword`, `clientTrustStorePassword`
    
3. *Global* environment variables
    * `TESSERA_SERVER_KEYSTORE_PWD`, `TESSERA_SERVER_TRUSTSTORE_PWD`
    * `TESSERA_CLIENT_KEYSTORE_PWD`, `TESSERA_CLIENT_TRUSTSTORE_PWD`

The *global* environment variables, if set, are applied to all server configs defined in the configfile (i.e. if a P2P and ADMIN server are both configured with TLS then the values set for the global environment variables will be used for both).  These values are ignored if the passwords are also provided in the configfile or as prefixed environment variables.

The *prefixed* environment variables are only applied to the servers with that `environmentVariablePrefix` value defined in their config.  This allows, for example, a P2P and ADMIN server to be configured with different prefixes, `P2P` and `ADMIN`.  Different keystores can then be used for each server and the individual passwords provided with `P2P_<...>` and `ADMIN_<...>`.

##### Generating keystores
If keystores do not already exist, Tessera can generate `.jks` (Java keystore) files for use with non-CA Trust Modes (see Trust Modes).

By setting `"generateKeyStoreIfNotExisted": "true"`, Tessera will check whether files already exist at the paths provided in the `serverKeyStore` and `clientKeyStore` config values.  If the files do not exist: 

1. New keystores will be generated and saved at the `serverKeyStore` and `clientKeyStore` paths
2. The keystores will be secured using the corresponding passwords if they are provided (see Passwords)

#### PEM files
Below is a config sample for using the `.pem` file format:
```json
"sslConfig" : {
    "tls" : "STRICT",
    "generateKeyStoreIfNotExisted" : "false",
    "serverTlsKeyPath" : "server-key.pem",  
    "serverTlsCertificatePath" : "server-cert.pem",
    "serverTrustCertificates" : ["server-trust.pem"]
    "serverTrustMode" : "CA",
    "clientTlsKeyPath" : "client-key.pem",
    "clientTlsCertificatePath" : "client-cert.pem",
    "clientTrustCertificates" : ["client-trust.pem"]
    "clientTrustMode" : "TOFU",
    "knownClientsFile" : "knownClients",
    "knownServersFile" : "knownServers"
}
```

#### Trust Modes
The Trust Mode for both client and server must also be specified. Multiple trust modes are supported: `TOFU`, `WHITELIST`, `CA`, `CA_OR_TOFU`, and `NONE`.

* `TOFU` (Trust-on-first-use)  
    Only the first node that connects identifying as a certain host will be allowed to connect as the same host in the future. When connecting for the first time, the host and its certificate will be added to `knownClientsFile` (for server), or `knownServersFile` (for client). These files will be generated if not already existed, using the values specified in `knownClientsFile` and `knownServersFile`.

    A config sample for `TOFU` trust mode is:

    ```json
    "sslConfig" : {
        "tls" : "STRICT",
        "generateKeyStoreIfNotExisted" : "true",
        "serverKeyStore" : "server-keystore",
        "serverKeyStorePassword" : "tessera",
        "serverTrustMode" : "TOFU",
        "clientKeyStore" : "client-keystore",
        "clientKeyStorePassword" : "tessera",
        "clientTrustMode" : "TOFU",
        "knownClientsFile" : "knownClients",
        "knownServersFile" : "knownServers"
    }
    ```

* `WHITELIST`   
    Only nodes that have previously connected to this node and have been added to the `knownClients` file will be allowed to connect. Similarly, this node will only be allowed to make connections to nodes that have been added to the `knownServers` file. This trust mode will not add new entries to the `knownClients` or `knownServers` files.
    
    With this trust mode, the whitelist files (`knownClientsFile` and `knownServersFile`) must be provided. 
    
    A config sample for `WHITELIST` trust mode is:
    
    ```json
    "sslConfig" : {
        "tls" : "STRICT",
        "generateKeyStoreIfNotExisted" : "true",
        "serverKeyStore" : "server-keystore",
        "serverKeyStorePassword" : "tessera",
        "serverTrustMode" : "WHITELIST",
        "clientKeyStore" : "client-keystore",
        "clientKeyStorePassword" : "tessera",
        "clientTrustMode" : "WHITELIST",
        "knownClientsFile" : "knownClients",
        "knownServersFile" : "knownServers"
    }
    ```

* `CA` 
    Only nodes with a valid certificate and chain of trust are allowed to connect. For this trust mode, trust stores must be provided and must contain a list of trust certificates.
    
    A config sample for `CA` trust mode is:
    
    ```json
    "sslConfig" : {
        "tls" : "STRICT",
        "generateKeyStoreIfNotExisted" : "false", //You can't generate trust stores when using CA
        "serverKeyStore" : "server-keystore",
        "serverKeyStorePassword" : "tessera",
        "serverTrustStore" : "server-truststore",
        "serverTrustStorePassword" : "tessera",
        "serverTrustMode" : "CA",
        "clientKeyStore" : "client-keystore",
        "clientKeyStorePassword" : "tessera",
        "clientTrustStore" : "client-truststore",
        "clientTrustStorePassword" : "tessera",
        "clientTrustMode" : "CA",
        "knownClientsFile" : "knownClients",
        "knownServersFile" : "knownServers"
    }
    ```
