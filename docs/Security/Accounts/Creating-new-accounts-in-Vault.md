# Creating new accounts in Vault

## Authenticating Quorum with Vault
Depending on the authentication method being used, different environment variables must be set:

**Approle (recommended)**

Set `VAULT_ROLE_ID` and `VAULT_SECRET_ID`.  These credentials are obtained as detailed in the  [Vault AppRole documentation](https://www.vaultproject.io/docs/auth/approle.html#configuration).

**Token**

Set `VAULT_TOKEN`. This can be the root token or a token obtained by logging in outside of Quorum (e.g. using the HTTP API).

## Creating new accounts

Quorum can create new accounts and immediately store them in a Vault, never storing them on the local filesystem.  

Accounts are stored in two components:

1. The private key used for signing
2. The 20 byte Ethereum/Quorum address derived from the private key

The `geth account new` CLI has been extended to support Vault by adding the following CLI options:

```bash
$ geth account new --help
# truncated output
HASHICORP VAULT OPTIONS:
  --hashicorp                   Store the newly created account in a Hashicorp Vault
  --hashicorp.url value         Address of the Vault server expressed as a URL and port, for example: https://127.0.0.1:8200/
  --hashicorp.approle value     Vault path for an enabled Approle auth method (requires VAULT_ROLE_ID and VAULT_SECRET_ID env vars to be set) (default: "approle")
  --hashicorp.clientcert value  Path to a PEM-encoded client certificate. Required when communicating with the Vault server using TLS
  --hashicorp.clientkey value   Path to an unencrypted, PEM-encoded private key which corresponds to the matching client certificate
  --hashicorp.cacert value      Path to a PEM-encoded CA certificate file. Used to verify the Vault server's SSL certificate
  --hashicorp.engine value      Vault path for an enabled KV v2 secret engine
  --hashicorp.nameprefix value  The new address and key will be created with name <prefix>Addr and <prefix>Key respectively.   Secrets with the same name in the Vault will be versioned and overwritten.
```

These options allow the user to specify the Vault location to store the address and private key and, if necessary, provide TLS config and Approle authentication config.  

`--hashicorp`, `--hashicorp.url` and `--hashicorp.engine` are required options.

The private key and address are stored with the same `nameprefix` and with suffixes `Key` and `Addr` respectively.  The Vault path of the created secrets and the new account address are returned on successful creation, e.g.:  

```bash
$ geth account new --hashicorp --hashicorp.url http://localhost:8200 \
                   --hashicorp.engine kv --hashicorp.nameprefix primary

Written to Vault: http://localhost:8200/v1/kv/data/primaryAddr?version=1
Written to Vault: http://localhost:8200/v1/kv/data/primaryKey?version=1
Address: {bcc328f4679fcc781d983da1c8be3d3baa6e5ae5}
```

Addresses and keys are stored in the Vault in their string hex representation, e.g.:

```json
// example Vault data for created address
{
  ...
  "data" : {
      "secret" : "bcc328f4679fCC781d983DA1c8bE3D3bAA6E5AE5"
  },
  ...
}


// example Vault data for created private key
{
  ...
  "data" : {
      "secret" : "dfe8b73d2771380d3f36bd78ce537715e812d7797c0b055fe944cd42cc750853"
  },
  ...
}
```

Saving a new account to an existing secret will overwrite the values stored at that secret. Previous versions may be retained and be retrievable depending on how the K/V secrets engine is configured.
