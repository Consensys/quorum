!!! warning "Change from Tessera v0.10.2+"
    The `keys.keyData.passwords` field is no longer supported as of Tessera v0.10.2.  
    
    Instead, use `keys.keyData.passwordFile` or utilise the [CLI password prompt](#providing-key-passwords-at-runtime) when starting the node.

Tessera uses cryptographic keys to provide transaction privacy.  

You can use existing private/public key pairs as well as use Tessera to generate new key pairs for you.  See [Generating & securing keys](../../Tessera%20Services/Keys/Keys) for more info.

```json tab="v0.10.3 onwards"
"keys": {
    "passwordFile": "Path",
    "keyVaultConfigs": [
        {
            "keyVaultType": "Enumeration: AZURE, HASHICORP, AWS",
            "properties": "Map[string]string"
        }
    ],
    "keyData": [
        {
            // The data for a private/public key pair
        }
    ]
}
```

```json tab="v0.10.2"
"keys": {
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
            // The data for a private/public key pair
        }
    ]
}
```

```json tab="v0.10.1 and earlier"
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
            // The data for a private/public key pair
        }
    ]
}
```

## KeyData
Key pairs can be provided in several ways:

### Direct key pairs

!!! warning
    Direct key pairs and unprotected inline key pairs are convenient but are the least secure configuration options available as the private key is exposed in the configuration file. The other options available are more secure and recommended for production environments.

The key pair data is provided in plain text in the configfile.

```json
"keys": {
    "keyData": [
        {
        "privateKey": "yAWAJjwPqUtNVlqGjSrBmr1/iIkghuOh1803Yzx9jLM=",
        "publicKey": "/+UuD63zItL1EbjxkKUljMgG8Z1w0AJ8pNOR4iq2yQc="
        }
    ]
}
```  

### Inline key pairs
#### Unprotected

!!! warning
    Direct key pairs and unprotected inline key pairs are convenient but are the least secure configuration options available as the private key is exposed in the configuration file. The other options available are more secure and recommended for production environments.
    
The key pair data is provided in plain text in the configfile.  The plain text private key is provided in a `config` json object:
```json
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

#### Protected
The public key is provided in plain text.  The private key must be password-protected using Argon2.  The corresponding encrypted data is provided in the `config` json object.

```json tab="v0.10.2 onwards"
"keys": {
    "passwordFile": "/path/to/pwds.txt",
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

```json tab="v0.10.1 and earlier"
"keys": {
    "passwords": ["password"],
    "passwordFile": "/path/to/pwds.txt",
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

Passwords must be provided so that Tessera can decrypt and use the private keys.  Passwords can be provided in multiple ways:

|        | Description                                                                                                                                                                  |
|--------|--------------|
| File   | `"passwordFile": "/path/to/pwds.txt"`<br/>Must contain only one password per line.  Empty lines should be used for unlocked keys.  Passwords must be provided in the order that key pairs are defined in the config. |
| Direct | `"passwords": ["pwd1", "pwd2", ...]`<br/>Empty strings should be used for unlocked keys.  Passwords must be provided in the order that key pairs are defined in the config.  Not recommended for production use.     |
| CLI    | Tessera will prompt on the CLI for the passwords of any encrypted keys that have not had passwords provided in the config.  This process only needs to be performed once, when starting the node.              |

### Filesystem key pairs   
The keys in the pair are stored in files:

```json tab="v0.10.2 onwards"
"keys": {
    "passwordFile": "/path/to/pwds.txt",
    "keyData": [
        {
            "privateKeyPath": "/path/to/privateKey.key",
            "publicKeyPath": "/path/to/publicKey.pub"
        }
    ]
}
```

```json tab="v0.10.1 and earlier"
"keys": {
    "passwords": ["password"],
    "passwordFile": "/path/to/pwds.txt",
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

The contents of the private key file must contain the private key in the Inline key pair format, e.g.:
```json
{
    "type" : "unlocked",
    "data" : {
        "bytes" : "DK0HDgMWJKtZVaP31mPhk6TJNACfVzz7VZv2PsQZeKM="
    }
}
```

or

```json
{
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
}
```

Passwords must be provided so that Tessera can decrypt and use the private keys.  Passwords can be provided in multiple ways:

|        | Description                                                                                                                                                                  |
|--------|--------------|
| File   | `"passwordFile": "/path/to/pwds.txt"`<br/>Must contain only one password per line.  Empty lines should be used for unlocked keys.  Passwords must be provided in the order that key pairs are defined in the config. |
| Direct | `"passwords": ["pwd1", "pwd2", ...]`<br/>Empty strings should be used for unlocked keys.  Passwords must be provided in the order that key pairs are defined in the config.  Not recommended for production use.     |
| CLI    | Tessera will prompt on the CLI for the passwords of any encrypted keys that have not had passwords provided in the config.  This process only needs to be performed once, when starting the node.              |

### Azure Key Vault key pairs
The keys in the pair are stored as secrets in an Azure Key Vault.  This requires providing the vault url and the secret IDs for both keys:

```json tab="v0.10.3 onwards"
"keys": {
    "keyVaultConfigs": [
        {
            "keyVaultType": "AZURE",
            "properties": {
                "url": "https://my-vault.vault.azure.net"
            } 
        }
    ],
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

```json tab="v0.10.2 and earlier"
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

!!! info
    Environment variables must be set if using an Azure Key Vault, for more information see [Setting up an Azure Key Vault](../../Tessera%20Services/Keys/Setting%20up%20an%20Azure%20Key%20Vault)

### Hashicorp Vault key pairs
The keys in the pair are stored as a secret in a Hashicorp Vault.  Additional configuration can also be provided if the Vault is configured to use TLS and if the AppRole auth method is being used at a different path to the default (`approle`):

```json tab="v0.10.3 onwards"
"keys": {
    "keyVaultConfigs": [
        {
            "keyVaultType": "HASHICORP",
            "properties": {
                "url": "https://localhost:8200",
                "tlsKeyStorePath": "/path/to/keystore.jks",
                "tlsTrustStorePath": "/path/to/truststore.jks",
                "approlePath": "not-default"
            }
        }
    ],
    "keyData": [
        {
            "hashicorpVaultSecretEngineName": "engine",
            "hashicorpVaultSecretName": "secret",
            "hashicorpVaultSecretVersion": 1,
            "hashicorpVaultPrivateKeyId": "privateKey",
            "hashicorpVaultPublicKeyId": "publicKey",
        }
    ]
}
```

```json tab="v0.10.2 and earlier"
"keys": {
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
}
```

This example configuration will retrieve version 1 of the secret `engine/secret` from Vault and its corresponding values for `privateKey` and `publicKey`.  

If no `hashicorpVaultSecretVersion` is provided then the latest version for the secret will be retrieved by default.

Tessera requires TLS certificates and keys to be stored in `.jks` Java keystore format.  If the `.jks` files are password protected then the following environment variables must be set: 

* `HASHICORP_CLIENT_KEYSTORE_PWD`
* `HASHICORP_CLIENT_TRUSTSTORE_PWD` 

!!! info
    If using a Hashicorp Vault additional environment variables must be set and a version 2 K/V secret engine must be enabled.  For more information see [Setting up a Hashicorp Vault](../../Tessera%20Services/Keys/Setting%20up%20a%20Hashicorp%20Vault).

### AWS Secrets Manager key pairs
The keys in the pair are stored as secrets in the _AWS Secrets Manager_.  This requires providing the secret IDs for both keys.  The endpoint is optional as the _AWS SDK_ can fallback to its inbuilt property retrieval chain (e.g. using the environment variable `AWS_REGION` or `~/.aws/config` file - see [the AWS docs](https://docs.aws.amazon.com/sdk-for-java/v2/developer-guide/credentials.html) for similar behaviour explained in the context of credentials):

```json tab="v0.10.3 onwards"
"keys": {
    "keyVaultConfigs": [
        {
            "keyVaultConfigType": "AWS",
            "properties": {
                "endpoint": "https://secretsmanager.us-west-2.amazonaws.com"
            }
        }
    ],
    "keyData": [
        {
            "awsSecretsManagerPublicKeyId": "secretIdPub",
            "awsSecretsManagerPrivateKeyId": "secretIdKey"
        }
    ]
}
```

This example configuration will retrieve the secrets `secretIdPub` and `secretIdKey` from the _AWS Secrets Manager_ using the endpoint `https://secretsmanager.us-west-2.amazonaws.com`.

!!! info
    A `Credential should be scoped to a valid region` error when starting means that the region specified in the `endpoint` differs from the region the AWS SDK has retrieved from its [property retrieval chain](https://docs.aws.amazon.com/sdk-for-java/v2/developer-guide/credentials.html).  This can be resolved by setting the `AWS_REGION` environment variable to the same region as defined in the `endpoint`.

!!! info
    Environment variables must be set if using an _AWS Secrets Manager_, for more information see [Setting up an AWS Secrets Manager](../../Tessera%20Services/Keys/Setting%20up%20an%20AWS%20Secrets%20Manager)

## Providing key passwords at runtime
Tessera will start a CLI password prompt if it has incomplete password data for its locked keys.  This prompt can be used to provide the required passwords for each key without having to provide them in the configfile itself.  

For example:   

```bash
tessera -configfile path/to/config.json
Password for key[0] missing or invalid.
Attempt 1 of 2. Enter a password for the key

2019-12-09 13:48:16.159 [main] INFO  c.q.t.config.keys.KeyEncryptorImpl - Decrypting private key
2019-12-09 13:48:19.364 [main] INFO  c.q.t.config.keys.KeyEncryptorImpl - Decrypted private key
# Tessera startup continues as normal
``` 

## Multiple Keys
If wished, multiple key pairs can be specified for a Tessera node. In this case, any one of the public keys can be used to address a private transaction to that node. Tessera will sequentially try each key to find one that can decrypt the payload. This can be used, for example, to simplify key rotation.

Note that multiple key pairs can only be set up within the configuration file, not via separate filesystem key files.

## Viewing the keys registered for a node
For Tessera v0.10.2 onwards the ThirdParty API `/keys` endpoint can be used to view the public keys of the key pairs currently in use by your Tessera node. 

For Tessera v0.10.1 and earlier, an ADMIN API endpoint `/config/keypairs` exists to allow you to view the public keys of the key pairs currently in use by your Tessera node.  

A sample response for the request `adminhost:port/config/keypairs` is:

```json tab="v0.10.2 onwards"
request: thirdpartyhost:port/keys
{
   "keys" : [
      {
         "key" : "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="
      },
      {
         "key" : "ABn6zhBth2qpdrJXp98IvjExV212ALl3j4U//nj4FAI="
      }
   ]
}
```

```json tab="v0.10.1 and earlier"
request: adminhost:port/config/keypairs
[
   {
      "publicKey" : "oNspPPgszVUFw0qmGFfWwh1uxVUXgvBxleXORHj07g8="
   },
   {
      "publicKey" : "ABn6zhBth2qpdrJXp98IvjExV212ALl3j4U//nj4FAI="
   }
]
```

The corresponding server must be configured in the node's configuration file, as described in [Configuration Overview](../Configuration%20Overview).
