Tessera configuration varies by version as new features are added or changed. Below is a list of sample configurations that show a possible structure. There may be more features that are not included in the sample; a full list of features can be found [here](../Configuration%20Overview).

## Samples

| Version       |
| ------------- |
| [0.10.3](../Tessera%20v0.10.3%20sample%20settings) |
| [0.10.2](../Tessera%20v0.10.2%20sample%20settings) |
| [0.10](../Tessera%20v0.10.0%20sample%20settings) |
| [0.9](../Tessera%20v0.9%20sample%20settings) |
| [0.8](../Tessera%20v0.8%20sample%20settings)      |
| [0.7.3](../Tessera%20v0.7.3%20sample%20settings)      |

## Changelist
### 0.10.3
- The `keys.azureKeyVaultConfig` and `keys.hashicorpKeyVaultConfig` fields are now deprecated.  Instead, the generic `keys.keyVaultConfigs` should be used.  See [Keys Config](../Keys) for more info.

### 0.10.2
- The `keys.keyData.passwords` field is no longer supported.  Instead, use `keys.keyData.passwordFile` or utilise the [CLI password prompt](../Keys#providing-key-passwords-at-runtime) when starting the node.

- Added configuration to choose alternative curves/symmetric ciphers. If no encryptor configuration is provided it will default to NaCl (see [Supporting alternative curves in Tessera](../Configuration Overview#supporting-alternative-curves-in-tessera) for more details).

    e.g.
    ```json
    {
        "encryptor": {
            "type":"EC",
            "properties":{
                "symmetricCipher":"AES/GCM/NoPadding",
                "ellipticCurve":"secp256r1",
                "nonceLength":"24",
                "sharedKeyLength":"32"
            }
        },
        ...
    }
    ``` 

### 0.10
- Added feature-toggle for remote key validation.  Disabled by default.
    ```json
    {
        "features": {
            "enableRemoteKeyValidation": false
        },
        ...
    }
    ```
### 0.9
- Collapsed server socket definitions into a single property `serverAddress`, e.g.
    ```json
    {
        "serverConfigs": [
            {
                "serverSocket": {
                    "type":"INET",
                    "port": 9001,
                    "hostName": "http://localhost"
                },
                ...
            }
        ],
        ...
    }
    ```
    becomes
    ```json
    {
        "serverConfigs": [
            {
                "serverAddress": "http://localhost:9001",
                ...
            }
        ],
        ...
    }
    ```

### 0.8
- Added modular server configurations
