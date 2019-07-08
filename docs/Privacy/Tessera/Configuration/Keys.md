Tessera uses cryptographic keys to provide transaction privacy.  

You can use existing private/public key pairs as well as use Tessera to generate new key pairs for you.  See [Generating & securing keys](../../Tessera%20Services/Keys/Keys) for more info.
```
"keys": {
    "passwords": [],
    "passwordFile": "Path",
    "azureKeyVaultConfig": {
        "url": "Url"
    },
    "hashicorpKeyVaultConfig": {
        "url": "Url",
        "approlePath": "String",
        "tlsKeyStorePath": "Path",
        "tlsTrustStorePath": "Path" 
    },
    "keyData": [
        {
            //The data for a private/public key pair
        }
    ]
}
```

## KeyData
Key pairs can be provided in several ways:

#### 1. Direct key pairs
Direct key pairs are convenient but are the least secure configuration option available, as you expose your private key in the configuration file. More secure options are available and preferable for production environments.

The key pair data is provided in plain text in the configfile:
```
"keys": {
    "keyData": [
        {
        "privateKey": "yAWAJjwPqUtNVlqGjSrBmr1/iIkghuOh1803Yzx9jLM=",
        "publicKey": "/+UuD63zItL1EbjxkKUljMgG8Z1w0AJ8pNOR4iq2yQc="
        }
    ]
}
```  

#### 2. Inline key pairs
The public key is provided in plain text.  The private key is provided through additional config:
```
"keys": {
    "keyData": [
        {
            "config": {
                "data": {
                    "bytes": "yAWAJjwPqUtNVlqGjSrBmr1/iIkghuOh1803Yzx9jLM="
                },
                "type": "unlocked"
            },
            "publicKey": "/+UuD63zItL1EbjxkKUljMgG8Z1w0AJ8pNOR4iq2yQc="
        }
    ]
}     
```

This allows for the use of Argon2 password-secured private keys by including the corresponding Argon2 settings in the additional config:

```
"keys": {
    "passwords": ["password"],
    "keyData": [
        {
            "config": {
                "data": {
                    "aopts": {
                        "variant": "id",
                        "memory": 1048576,
                        "iterations": 10,
                        "parallelism": 4,
                    },
                    "snonce": "x3HUNXH6LQldKtEv3q0h0hR4S12Ur9pC",
                    "asalt": "7Sem2tc6fjEfW3yYUDN/kSslKEW0e1zqKnBCWbZu2Zw=",
                    "sbox": "d0CmRus0rP0bdc7P7d/wnOyEW14pwFJmcLbdu2W3HmDNRWVJtoNpHrauA/Sr5Vxc"
                },
                "type": "argon2sbox"
            },
            "publicKey": "/+UuD63zItL1EbjxkKUljMgG8Z1w0AJ8pNOR4iq2yQc="
        }
    ]
}
```

#### 3. Azure Key Vault key pairs
The keys in the pair are stored as secrets in an Azure Key Vault.  This requires providing the vault url and the secret IDs for both keys:
```
"keys": {
    "azureKeyVaultConfig": {
        "url": "https://my-vault.vault.azure.net"
    },
    "keyData": [
        {
            "azureVaultPrivateKeyId": "Key",
            "azureVaultPublicKeyId": "Pub",
            "azureVaultPublicKeyVersion": "bvfw05z4cbu11ra2g94e43v9xxewqdq7",
            "azureVaultPrivateKeyVersion": "0my1ora2dciijx5jq9gv07sauzs5wjo2"
        }
    ]
}
```

This example configuration will retrieve the specified versions of the secrets `Key` and `Pub` from the key vault with DNS name `https://my-vault.vault.azure.net`.    If no version is specified then the latest version of the secret is retrieved.

> Environment variables must be set if using an Azure Key Vault, for more information see [Setting up an Azure Key Vault](../../Tessera%20Services/Keys/Setting%20up%20an%20Azure%20Key%20Vault)

#### 4. Hashicorp Vault key pairs
The keys in the pair are stored as a secret in a Hashicorp Vault.  Additional configuration can also be provided if the Vault is configured to use TLS and if the AppRole auth method is being used at a different path to the default (`approle`):
```
"hashicorpKeyVaultConfig": {
    "url": "https://localhost:8200",
    "tlsKeyStorePath": "/path/to/keystore.jks",
    "tlsTrustStorePath": "/path/to/truststore.jks",
    "approlePath": "not-default",
},
"keyData": [
    {
        "hashicorpVaultSecretEngineName": "engine",
        "hashicorpVaultSecretName": "secret",
        "hashicorpVaultSecretVersion": 1,
        "hashicorpVaultPrivateKeyId": "privateKey",
        "hashicorpVaultPublicKeyId": "publicKey",
    }
]
```

This example configuration will retrieve version 1 of the secret `engine/secret` from Vault and its corresponding values for `privateKey` and `publicKey`.  

If no `hashicorpVaultSecretVersion` is provided then the latest version for the secret will be retrieved by default.

Tessera requires TLS certificates and keys to be stored in `.jks` Java keystore format.  If the `.jks` files are password protected then the following environment variables must be set: 
* `HASHICORP_CLIENT_KEYSTORE_PWD`
* `HASHICORP_CLIENT_TRUSTSTORE_PWD` 

> If using a Hashicorp Vault additional environment variables must be set and a version 2 K/V secret engine must be enabled.  For more information see [Setting up a Hashicorp Vault](../../Tessera%20Services/Keys/Setting%20up%20a%20Hashicorp%20Vault).

#### 5. Filesystem key pairs   
The keys in the pair are stored in files:
```
"keys": {
    "passwordFile": "/path/to/passwords",
    "keyData": [
        {
            "privateKeyPath": "/path/to/privateKey.key",
            "publicKeyPath": "/path/to/publicKey.pub"
        }
    ]
}
```
The contents of the public key file must contain the public key only, e.g.: 
```
/+UuD63zItL1EbjxkKUljMgG8Z1w0AJ8pNOR4iq2yQc=
```

The contents of the private key file must contain the private key in the config format, e.g.:
```
{
    "type" : "unlocked",
    "data" : {
        "bytes" : "DK0HDgMWJKtZVaP31mPhk6TJNACfVzz7VZv2PsQZeKM="
    }
}
```
## Multiple Keys
If wished, multiple key pairs can be specified for a Tessera node. In this case, any one of the public keys can be used to address a private transaction to that node. Tessera will sequentially try each key to find one that can decrypt the payload. This can be used, for example, to simplify key rotation.

Note that multiple key pairs can only be set up within the configuration file, not via separate filesystem key files.

## Viewing the keys registered for a node
An ADMIN API endpoint `/config/keypairs` exists to allow you to view the public keys of the key pairs currently in use by your Tessera node.  This requires configuring an ADMIN server in the node's configuration file, as described in [Configuration Overview](../Configuration%20Overview).

A sample response for the request `adminhost:port/config/keypairs` is:
```json
[
   {
      "publicKey" : "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="
   },
   {
      "publicKey" : "ABn6zhBth2qpdrJXp98IvjExV212ALl3j4U//nj4FAI="
   }
]
```
