**Changes:**
- collapsed server socket definitions into a single property `serverAddress`

e.g.
```json
"serverSocket": {
    "type":"INET",
    "port": 9001,
    "hostName": "http://localhost"
},
```
becomes
```
"serverAddress": "http://localhost:9001",
```


---

**Sample**

```json
{
    "useWhiteList": "boolean",

    "jdbc": {
        "url": "String",
        "username": "String",
        "password": "String"
    },

    "serverConfigs": [
        {
            "app": "ENCLAVE", // Defines us using a remote enclave, leave out if using built-in enclave
            "enabled": true,
            "serverAddress": "http://localhost:9081", //Where to find the remote enclave
            "communicationType": "REST"
        },
        {
            "app": "ThirdParty",
            "enabled": true,
            "serverAddress": "http://localhost:9081",
            "bindingAddress": "String - url with port e.g. http://127.0.0.1:9081",
            "communicationType": "REST"
        },

        {
            "app": "Q2T",
            "enabled": true,
            "serverAddress": "unix:/tmp/tm.ipc",
            "communicationType": "REST"
        },

        {
            "app": "P2P",
            "enabled": true,
            "serverAddress": "http://localhost:9001",
            "bindingAddress": "String - url with port e.g. http://127.0.0.1:9001",
            "sslConfig": {
                "tls": "enum STRICT,OFF",
                "generateKeyStoreIfNotExisted": "boolean",
                "serverKeyStore": "Path",
                "serverTlsKeyPath": "Path",
                "serverTlsCertificatePath": "Path",
                "serverKeyStorePassword": "String",
                "serverTrustStore": "Path",
                "serverTrustCertificates": [
                    "Path..."
                ],
                "serverTrustStorePassword": "String",
                "serverTrustMode": "Enumeration: CA, TOFU, WHITELIST, CA_OR_TOFU, NONE",
                "clientKeyStore": "Path",
                "clientTlsKeyPath": "Path",
                "clientTlsCertificatePath": "Path",
                "clientKeyStorePassword": "String",
                "clientTrustStore": "Path",
                "clientTrustCertificates": [
                    "Path..."
                ],
                "clientTrustStorePassword": "String",
                "clientTrustMode": "Enumeration: CA, TOFU, WHITELIST, CA_OR_TOFU, NONE",
                "knownClientsFile": "Path",
                "knownServersFile": "Path"
            },
            "communicationType": "REST"
        }
    ],

    "peer": [
        {
            "url": "url e.g. http://127.0.0.1:9000/"
        }
    ],

    "keys": {
        "passwords": [
            "String..."
        ],
        "passwordFile": "Path",
        "azureKeyVaultConfig": {
            "url": "Azure Key Vault url"
        },
        "hashicorpKeyVaultConfig": {
            "url": "Hashicorp Vault url",
            "approlePath": "String (defaults to 'approle' if not set)",
            "tlsKeyStorePath": "Path to jks key store",
            "tlsTrustStorePath": "Path to jks trust store"
        },

        "keyData": [
            {
                "config": {
                    "data": {
                        "aopts": {
                            "variant": "Enum : id,d or i",
                            "memory": "int",
                            "iterations": "int",
                            "parallelism": "int"
                        },
                        "bytes": "String",
                        "snonce": "String",
                        "asalt": "String",
                        "sbox": "String",
                        "password": "String"
                    },
                    "type": "Enum: argon2sbox or unlocked. If unlocked is defined then config data is required. "
                },
                "privateKey": "String",
                "privateKeyPath": "Path",
                "azureVaultPrivateKeyId": "String",
                "azureVaultPrivateKeyVersion": "String",
                "publicKey": "String",
                "publicKeyPath": "Path",
                "azureVaultPublicKeyId": "String",
                "azureVaultPublicKeyVersion": "String",
                "hashicorpVaultSecretEngineName": "String",
                "hashicorpVaultSecretName": "String",
                "hashicorpVaultSecretVersion": "Integer (defaults to 0 (latest) if not set)",
                "hashicorpVaultPrivateKeyId": "String",
                "hashicorpVaultPublicKeyId": "String"
            }
        ]
    },

    "alwaysSendTo": [
        "String..."
    ],

    "unixSocketFile": "Path"
}
```

---

**Sample enclave settings**

```json
{
    "serverConfigs": [
        {
            "app": "ENCLAVE",
            "enabled": true,
            "serverAddress": "http://localhost:9001",
            "bindingAddress": "String - url with port e.g. http://127.0.0.1:9001",
            "sslConfig": {
                "tls": "enum STRICT,OFF",
                "generateKeyStoreIfNotExisted": "boolean",
                "serverKeyStore": "Path",
                "serverTlsKeyPath": "Path",
                "serverTlsCertificatePath": "Path",
                "serverKeyStorePassword": "String",
                "serverTrustStore": "Path",
                "serverTrustCertificates": [
                    "Path..."
                ],
                "serverTrustStorePassword": "String",
                "serverTrustMode": "Enumeration: CA, TOFU, WHITELIST, CA_OR_TOFU, NONE",
                "clientKeyStore": "Path",
                "clientTlsKeyPath": "Path",
                "clientTlsCertificatePath": "Path",
                "clientKeyStorePassword": "String",
                "clientTrustStore": "Path",
                "clientTrustCertificates": [
                    "Path..."
                ],
                "clientTrustStorePassword": "String",
                "clientTrustMode": "Enumeration: CA, TOFU, WHITELIST, CA_OR_TOFU, NONE",
                "knownClientsFile": "Path",
                "knownServersFile": "Path"
            },
            "communicationType": "REST"
        }
    ],

    "keys": {
        "passwords": [
            "String..."
        ],
        "passwordFile": "Path",
        "azureKeyVaultConfig": {
            "url": "Azure Key Vault url"
        },
        "hashicorpKeyVaultConfig": {
            "url": "Hashicorp Vault url",
            "approlePath": "String (defaults to 'approle' if not set)",
            "tlsKeyStorePath": "Path to jks key store",
            "tlsTrustStorePath": "Path to jks trust store"
        },

        "keyData": [
            {
                "config": {
                    "data": {
                        "aopts": {
                            "variant": "Enum : id,d or i",
                            "memory": "int",
                            "iterations": "int",
                            "parallelism": "int"
                        },
                        "bytes": "String",
                        "snonce": "String",
                        "asalt": "String",
                        "sbox": "String",
                        "password": "String"
                    },
                    "type": "Enum: argon2sbox or unlocked. If unlocked is defined then config data is required. "
                },
                "privateKey": "String",
                "privateKeyPath": "Path",
                "azureVaultPrivateKeyId": "String",
                "azureVaultPrivateKeyVersion": "String",
                "publicKey": "String",
                "publicKeyPath": "Path",
                "azureVaultPublicKeyId": "String",
                "azureVaultPublicKeyVersion": "String",
                "hashicorpVaultSecretEngineName": "String",
                "hashicorpVaultSecretName": "String",
                "hashicorpVaultSecretVersion": "Integer (defaults to 0 (latest) if not set)",
                "hashicorpVaultPrivateKeyId": "String",
                "hashicorpVaultPublicKeyId": "String"
            }
        ]
    },

    "alwaysSendTo": [
        "String..."
    ]
}
```
