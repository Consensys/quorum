# Hashicorp Vault account Plugin

!!! info
    `account` plugins are currently in beta

The Hashicorp Vault plugin for Quorum enables the storage of Quorum account private keys in a [Hashicorp Vault KV v2 secret engine](https://www.vaultproject.io/docs/secrets/kv/kv-v2/).

Using the Hashicorp Vault plugin offers several benefits:

* Account private keys are stored in a Hashicorp Vault which can be deployed on separate infrastructure to the node  

* Vault allows for fine-grained access control to secrets 

## Building

Quorum will automatically download the plugin from bintray at startup.

Alternatively, the plugin can be downloaded or built manually and added to the [`baseDir`](https://docs.goquorum.com/en/latest/PluggableArchitecture/Settings/):

```shell
make
cp build/dist/quorum-account-plugin-hashicorp-vault-<version>.zip /path/to/baseDir
```

## Quickstart

See [Quickstart](../Quickstart) for a step-by-step guide on using Vault for Quorum account management.

## Configuration

Add the following `providers` config to the [`--plugins` file](../../../../../PluggableArchitecture/Settings):
```json
{
    "providers": {
        "account": {
            "name": "quorum-account-plugin-hashicorp-vault",
            "version": "0.0.1",
            "config": "<config>"
        }
    }
}
```

Based on this config, Quorum will look for [`quorum-account-plugin-hashicorp-vault-0.0.1.zip` in the default `baseDir`](../../../../../PluggableArchitecture/Internals#discovery).

`<config>` is the Hashicorp Vault plugin configuration:

!!! info   
    This config can be provided in [several ways](../../../../../PluggableArchitecture/Settings#plugindefinition)

```json
{
    "vault": "https://localhost:8200",
    "kvEngineName": "my-kv-engine",
    "accountDirectory": "file:///path/to/accts",
    "unlock": ["1a31744b4a6ee9f3c3d1550beb56d53d2a4fa454"],
    "authentication": {
        "roleId": "env://HASHICORP_ROLE_ID",
        "secretId": "env://HASHICORP_SECRET_ID",
        "approlePath": "approle"
    },
    "tls": {
        "caCert": "file:///path/to/ca.pem",
        "clientCert": "file:///path/to/client.pem",
        "clientKey": "file:///path/to/client.key"
    }
}
```

| Field | Description |
| --- | --- |
| `vault` | Vault server URL |
| `kvEngineName` | Name of an enabled Vault KV v2 secret engine to use for account storage |
| `accountDirectory` | Absolute `file://` URL of the account directory.  See [accountDirectory](#accountdirectory) |
| `unlock` | (Optional) List of accounts to retrieve from Vault at startup and store in memory |
| `authentication` | See [authentication](#authentication) |
| `tls` | (Optional) See [tls](#tls) |

### accountDirectory
The `accountDirectory` contains config files for each account managed by the plugin.  These files are similar to [keystore files](../../../Keystore-Files), except they do not contain any private data.

Typically these files do not have to be created or edited manually.  See [Creating accounts](#creating-accounts).

#### Example account file contents
```json
{
   "Address" : "1a31744b4a6ee9f3c3d1550beb56d53d2a4fa454",
   "VaultAccount" : {
      "SecretName" : "myacct",
      "SecretVersion" : 4
   },
   "Version" : 1
}
```


### authentication

The plugin can authenticate with Vault using [approle](https://www.vaultproject.io/docs/auth/approle) or [token](https://www.vaultproject.io/docs/auth/token) Vault authentication methods.


#### approle
!!! warning 
    approle is recommended in production
    
| Field | Description |
| --- | --- |
| `roleId` | approle role ID env URL (e.g. `env://VAR` will use the value of the `VAR` env variable) |
| `secretId` | approle secret ID env URL (e.g. `env://VAR` will use the value of the `VAR` env variable) |
| <span style="white-space:nowrap">`approlePath`</span> | name/path of the approle engine to login to |

#### token
| Field | Description |
| --- | --- |
| `token` | Vault token env URL (e.g. `env://VAR` will use the value of the `VAR` env variable) |

### tls

!!! warning 
    TLS is recommended in production

| Field | Description |
| --- | --- |
| `caCert` | Absolute `file://` URL of PEM-encoded CA certificate |
| `clientCert` | Absolute `file://` URL of PEM-encoded client certificate |
| `clientKey` | Absolute `file://` URL of PEM-encoded client key |

## Creating accounts

New accounts can be created and stored directly into the Vault by using the `account` plugin [RPC API](../../Overview#rpc-api) or [CLI](../../Overview#cli).  

!!! info 
    The plugin creates the account in memory, writes it to Vault, and zeros the private key.  The plugin never writes the private key to the node's disk.

A json config must be provided to the API/CLI when creating accounts.  Example:

```json
{
    "secretName": "myacct",
    "overwriteProtection": {
      "currentVersion": 4
    }
}
```

| Field | Description |
| --- | --- |
| `secretName` | Secret name/path the plugin will store the new account at |
| <span style="white-space:nowrap">`overwriteProtection.currentVersion`</span><br/>*or*<br/><span style="white-space:nowrap">`overwriteProtection.insecureDisable`</span> | Current integer version of this secret in Vault (`0` if no previous version exists)<br/>*or*<br/>Disable overwrite protection |

#### overwriteProtection

Typical usage will be to create separate Vault secrets for each account.  However, KV v2 secret engines also support secret versioning. 

The plugin uses [KV v2's Check-And-Set (CAS) feature](https://www.vaultproject.io/api-docs/secret/kv/kv-v2#create-update-secret) to protect against accidentally creating a new version of an existing secret.

If a secret with the same name already exists, `currentVersion` must be provided and must equal the current version number of the secret.

The CAS check can be skipped by setting `"insecureDisable": "true"`.  

!!! warning "Warning: Prevent accidental loss of account data"
    The K/V Version 2 secret engine supports versioning of secrets, however only a limited number of versions are retained (10 by default).  The `max-versions` number for a secret engine can be  set during creation of the secret engine or changed at a later date by using the Vault CLI or the Vault [HTTP API](https://www.vaultproject.io/api/secret/kv/kv-v2.html).
        
    To change `max-versions` using the CLI:
    ``` bash
    vault kv metadata put -max-versions <num> <kvEngineName>/<secretName>
    ``` 

## FAQ

### What data is stored in Vault?
The string hex representations of the account address and private key, e.g.:

```shell
$ vault kv get kv/myacct
====== Metadata ======
Key              Value
---              -----
created_time     2020-06-29T13:23:00.234716Z
deletion_time    n/a
destroyed        false
version          4

====================== Data ======================
Key                                         Value
---                                         -----
bcc328f4679fcc781d983da1c8be3d3baa6e5ae5    dfe8b73d2771380d3f36bd78ce537715e812d7797c0b055fe944cd42cc750853
```

### What are locked/unlocked accounts?
Accounts can be:

* locked: The plugin does not have the private key (it exists only in Vault)
* unlocked: The plugin has the private key

As with keystore accounts, accounts must be unlocked to sign data.  Accounts can be unlocked by:

* *Recommended*: Using geth's [`personal` API](https://geth.ethereum.org/docs/rpc/ns-personal)
* Setting `unlock` in the [config](#configuration)

The `personal` API minimises the time an account is unlocked.  The `unlock` config is useful if you need to unlock accounts in bulk for an indefinite amount of time (e.g. testing / Vault requests are impacting performance).

Any unlocked account can be locked with `personal_lockAccount`. 

The `personal_listWallets` API shows account status:
```js
> personal.listWallets
[{
    accounts: [{
        address: "0xda71f07446ed1eca304485dd00c4827ed0984998",
        url: "https://localhost:8200/v1/kv/data/myacct?version=1"
    }, {
        address: "0x0c069eb20e97f18e89e2151d312e9810e80fe089",
        url: "https://localhost:8200/v1/kv/data/myacct?version=2"
    }],
    status: "1 unlocked account(s): [0xda71f07446ed1eca304485dd00c4827ed0984998]",
    url: "plugin://account-plugin-hashicorp-vault"
}]
```

### Removing accounts/moving between nodes 

The files in the `accountDirectory` can be moved as required.  Afterwards, reload the plugin to apply any changes:

```shell tab="HTTP API"
curl -X POST http://localhost:<quorum-rpc-port> \
     -H "Content-type: application/json" \
     --data '{"jsonrpc":"2.0","method":"admin_reloadPlugin","params":["account"],"id":1}'
``` 

```js tab="js console"
admin.reloadPlugin("account")
```

!!! info
    If the account defined by the file is not available in the target node's Vault then use the `account` plugin [RPC API](../../Overview#rpc-api) or [CLI](../../Overview#cli) to import the account.  This will create the necessary file in the target node's account directory.  

### What password do I use for the personal API?
The `personal` APIs take a `passphrase` argument.  The Hashicorp Vault plugin does not use passwords as the Vault handles encryption of the account data.  

The plugin does not use the `passphrase` so any value can be used, e.g.:

```js
> personal.listWallets
[{
    accounts: [{
        address: "0xda71f07446ed1eca304485dd00c4827ed0984998",
        url: "https://localhost:8200/v1/kv/data/myacct?version=1"
    }],
    status: "0 unlocked account(s): []",
    url: "plugin://account-plugin-hashicorp-vault"
}]

// any value password can be used 
> personal.sign("0xaaaaaa", "0xda71f07446ed1eca304485dd00c4827ed0984998", "")
"0xc432436161788558a1e6387f83b703fecb90cf0507b39afdcd0d54769adc6fe71976bfac421076d54e31d3f45ddf76dcb47ad1a7035a3495d0b40bacfc258df41b"
> personal.sign("0xaaaaaa", "0xda71f07446ed1eca304485dd00c4827ed0984998", "pwd")
"0xc432436161788558a1e6387f83b703fecb90cf0507b39afdcd0d54769adc6fe71976bfac421076d54e31d3f45ddf76dcb47ad1a7035a3495d0b40bacfc258df41b"
``` 

### Approle token renewal
The plugin will automatically renew approle tokens where possible.  If the token is no longer renewable (e.g. because the max TTL has been reached) then the plugin will attempt to reauthenticate and retrieve a new token.  If the token obtained from an approle login is not renewable, then the plugin will not attempt renewal.

For more information about Hashicorp Vault TTL, leases and renewal see the [Vault documentation](https://www.vaultproject.io/docs/concepts/lease.html). 

### Approle policy requirements
To carry out all possible interactions with a Vault, a role must have the following policy capabilities: `["create", "update", "read"]`.  A subset of these capabilities can be configured if not all functionality is required.  