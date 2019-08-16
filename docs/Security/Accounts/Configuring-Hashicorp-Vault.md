The [Hashicorp Vault Getting Started documentation](https://learn.hashicorp.com/vault/) provides much of the information needed to get a Vault up and running.  The following section goes over some additional considerations to be made when running Quorum with Vault.

## TLS
When running in production environments it is advised to configure the Vault server for 2-way (mutual) TLS communication.  Quorum also supports 1-way TLS and unsecured (no TLS) communications with a Vault server.

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

## Auth methods
Quorum directly supports the [AppRole](https://www.vaultproject.io/docs/auth/approle.html) auth method.  If required, other auth methods can be used by logging in outside of Quorum (e.g. using the HTTP API) and providing the resulting vault token to Quorum.  See [Configuring Quorum to use Vault](Using-accounts-stored-in-Vault.md) for more information.

### Policies
To be able to carry out all possible interactions with a Vault, the Quorum role requires the following policy capabilities: `["create", "update", "read"]`.  A subset of these capabilities can be configured if not all functionality is required.

## Secret engines
Quorum can read and write accounts to the following secret engine type:

- [K/V Version 2](https://www.vaultproject.io/docs/secrets/kv/kv-v2.html)

The K/V Version 2 secret engine supports versioning of secrets, however only a limited number of versions are retained.  The `max-versions` number for a secret engine can be  set during creation of the secret engine or changed at a later date by using the Vault CLI or the Vault [HTTP API](https://www.vaultproject.io/api/secret/kv/kv-v2.html).

For example, to use the CLI to change `max-versions` for a single secret:
``` bash
vault kv metadata put -max-versions <num> <secret-engine>/<secret-name>
``` 
