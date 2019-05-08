The private/public key pairs used by Tessera can be [stored](../Keys) in and [retrieved](../../../Configuration/Keys) from a key vault, preventing the need to store the keys locally.

This page details how to set up and configure a Hashicorp Vault for use with Tessera.

 The [Hashicorp Vault Getting Started documentation](https://learn.hashicorp.com/vault/) provides much of the information needed to get started.  The following section goes over some additional considerations when running Tessera with Vault.

## Configuring the vault

### TLS
When running in production situations it is advised to configure the Vault server for 2-way (mutual) TLS communication.  Tessera also supports 1-way TLS and unsecured (no TLS) communications with a Vault server.

An example configuration for the Vault listener to use 2-way TLS is shown below.  This can be included as part of the `.hcl` used when starting the Vault server:

```
listener "tcp" {
        tls_min_version = "tls12"
        tls_cert_file = "/path/to/server.crt"
        tls_key_file = "/path/to/server.key"
        tls_require_and_verify_client_cert = "true"
        tls_client_ca_file = "/path/to/client-ca.crt"
}
```

### Auth methods
Tessera directly supports the [AppRole](https://www.vaultproject.io/docs/auth/approle.html) auth method.  If required, other auth methods can be used by logging in outside of Tessera (e.g. using the HTTP API) and providing the resulting vault token to Tessera.  See the *Enabling Tessera to use the vault* section below for more information.

When using AppRole, Tessera assumes the default auth path to be `approle`, however this value can be overwritten.  See [Keys](../../../Configuration/Keys) for more information.

### Policies
To be able to carry out all possible interactions with a Vault, Tessera requires the following policy capabilities: `["create", "update", "read"]`.  A subset of these capabilities can be configured if not all functionality is required.

### Secret engines
Tessera can read and write keys to the following secret engine type:

- [K/V Version 2](https://www.vaultproject.io/docs/secrets/kv/kv-v2.html)

The K/V Version 2 secret engine supports versioning of secrets, however only a limited number of versions are retained.  This number can be changed as part of the Vault configuration process.

## Enabling Tessera to use the vault
### Environment Variables
If using a Hashicorp Vault, Tessera requires certain environment variables to be set depending on the auth method being used.

- If using the AppRole auth method, set:
    1. `HASHICORP_ROLE_ID`
    2. `HASHICORP_SECRET_ID`

  These credentials are obtained as outlined in the [AppRole documentation](https://www.vaultproject.io/docs/auth/approle.html).  Tessera will use these credentials to authenticate with Vault.

- If using the root token or you already have a token due to authorising with an alternative method, set:
  1. `HASHICORP_TOKEN`

!!! note
    If using TLS additional environment variables must be set.  See [Keys](../../../Configuration/Keys) for more information as well as details of the Tessera configuration required to retrieve keys from a Vault.

### Dependencies
The Hashicorp dependencies are included in the `tessera-app-<version>-app.jar`.  If using the `tessera-simple-<version>-app.jar` then `hashicorp-key-vault-<version>-all.jar` must be added to the classpath.
