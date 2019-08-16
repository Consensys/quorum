# Using accounts stored in Vault

## Authenticating Quorum with Vault
Depending on the authentication method being used, different environment variables must be set:

**Approle (recommended)**

Set `VAULT_ROLE_ID` and `VAULT_SECRET_ID`.  These credentials are obtained as detailed in the  [Vault AppRole documentation](https://www.vaultproject.io/docs/auth/approle.html#configuration).

**Token**

Set `VAULT_TOKEN`. This can be the root token or a token obtained by logging in outside of Quorum (e.g. using the HTTP API).

## Configuration

Vault configuration is provided in the standard toml Quorum config file and applied at start up (i.e. with the `geth --config ...` CLI option), e.g.:

```toml
[Node.HashicorpVault]
    [Node.HashicorpVault.Client]
    Url        = "https://localhost:8200"
    Approle    = "auth"                   # optional, default: approle
    CaCert     = "/path/to/ca.pem"        # optional
    ClientCert = "/path/to/client.pem"    # optional
    ClientKey  = "/path/to/client.key"    # optional
    UnlockAll  = true                     # optional, default: false
    VaultPollingIntervalMillis = 5000     # optional, default: 10000

    [[Node.HashicorpVault.Secrets]]
    AddressSecret           = "primaryAddr"
    AddressSecretVersion    = 10
    PrivateKeySecret        = "primaryKey"
    PrivateKeySecretVersion = 10
    SecretEngine            = "kv"
    
    [[Node.HashicorpVault.Secrets]]
    AddressSecret           = "otherAddr"
    AddressSecretVersion    = 5
    PrivateKeySecret        = "otherKey"
    PrivateKeySecretVersion = 5
    SecretEngine            = "kv"
```

### Client config

The `Client` block configures the Vault API client used by Quorum.

- `Url`: Address of the Vault server expressed as a URL and port
- `Approle`: Vault path for an enabled Approle auth method (default: `"approle"`)
- `CaCert`: Path to a PEM-encoded CA certificate file. Used to verify the Vault server's SSL certificate
- `ClientCert`: Path to a PEM-encoded client certificate. Required when communicating with the Vault server using TLS
- `ClientKey`: Path to an unencrypted, PEM-encoded private key which corresponds to the matching client certificate
- `UnlockAll`: Unlock all accounts at start up (i.e. retrieve all private keys from the Vault and retain in memory indefinitely)  (default: `false`)
- `VaultPollingIntervalMillis`: The node periodically queries the Vault until the addresses (and private keys if `UnlockAll = true`) for each `Secrets` block have been retrieved  (default: `10000` i.e. 10 secs)  

!!! warning
    The Vault client library used by Quorum can also be configured by setting the [default Vault environment variables](https://www.vaultproject.io/docs/commands/#environment-variables) (e.g. setting `VAULT_CACERT` instead of `CaCert`).  Any values set in the `.toml` file will take precedence.  

### Secrets config

The `Secrets` blocks allow multiple accounts to be configured for the node and tell Quorum which addresses/keys to retrieve from the Vault.  

- `AddressSecret`: Secret name for the address component of the account
- `AddressSecretVersion`: Version of the address to retrieve
- `PrivateKeySecret`: Secret name for the private key component of the account
- `PrivateKeySecretVersion`: Version of the private key to retrieve
- `SecretEngine`: Vault path for an enabled KV v2 secret engine

If the address and key configured in a particular `Secrets` block are incorrect (i.e. the retrieved address cannot be derived from the retrieved private key) then any signing operations will be prevented.  This is to ensure requests to sign with a particular account are not signed with the wrong key due to a configuration error.

## Locked/unlocked accounts

Locked accounts do not have their private key stored in memory and so API requests have to be made to the Vault server whenever a signing operation is performed.  

Accounts can be unlocked (i.e. storing the private key in memory) temporarily or indefinitely, thereby reducing the number of API requests that need to be made to the Vault.  

This is useful in situations where network latency between Quorum and the Vault is impacting transaction throughput.  An account can be unlocked for a few minutes to perform the necessary signing operations, after which it will automatically lock and zero the key.  Any subsequent signing operations will then have to make API requests to the Vault unless another unlock is applied.

### Unlocking accounts

Accounts can be unlocked in the following ways:

1. **At node start**:  All accounts can be unlocked by adding the `UnlockAll = true` option to the `config.toml`
2. **During node operation**: An individual account can be unlocked using the new `personal_unlockVaultAccount(address, duration)` API.  This API method is also available on the `geth` JS console as `personal.unlockVaultAccount`.  

    The default duration is `300` seconds.  A duration of `0` unlocks the account indefinitely.  
   
    Existing timed unlocks can be overriden with a subsequent call to the API.  Indefinitely unlocked accounts cannot be overriden with a timed unlock and must be manually locked.

### Locking accounts

Accounts can be instantly locked using the new `personal_lockVaultAccount(address)` API.  This API method is available on the `geth` JS console as `personal.lockVaultAccount`.
  
This can be used on both indefinitely unlocked and timed unlocked accounts.
