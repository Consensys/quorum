title: security - Plugin Implementation - Quorum

# `security` Enterprise Plugin

| Version | Language |
|:--------|:---------|
| 1.0.0   | Go       |

This plugin implementation provides the following enterprise features to `geth` JSON RPC server:

- Providing TLS configuration to HTTP and WS transports
- Enabling `geth` JSON RPC (HTTP/WS) server to be an OAuth2-compliant resource server

## Configuration

<pre>
{
    "tls": object(<a href="#tlsconfiguration">TLSConfiguration</a>),
    "tokenValidation": object(<a href="#tokenvalidationconfiguration">TokenValidationConfiguration</a>)
}
</pre>

| Fields                    | Description                                                                                                                                                                           |
|:--------------------------|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `tls`                     | **(Optional)** If provided, serve the TLS configuration. See [TLSConfiguration](#tlsconfiguration) for more details                                                                   |
| `tokenValidation`         | **(Required)** Configuration to verify access token and extract granted authorities from the token. See [TokenValidationConfiguration](#tokenvalidationconfiguration) for more details |

### `TLSConfiguration`

<pre>
{
    "auto": bool,
    "certFile": <a href="#environmentawaredvalue">EnvironmentAwaredValue</a>,
    "keyFile": <a href="#environmentawaredvalue">EnvironmentAwaredValue</a>,
    "advanced": object(<a href="#tlsadvancedconfiguration">TLSAdvancedConfiguration</a>)
}
</pre>

| Fields     | Description                                                                                                                                                                                                          |
|:-----------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `auto`     | If true, generate a self-signed TLS certificate. Then save the generated certificate and private key in PEM format in `certFile` and `keyFile` respectively <br/> If false, use values from `certFile` and `keyFile` |
| `certFile` | Location to a file storing certificate in PEM format. Default is `cert.pem`                                                                                                                                          |
| `keyFile`  | Location to a file storing private key in PEM format. Default is `key.pem`                                                                                                                                           |
| `advanced` | Additional TLS configuration                                                                                                                                                                                         |

### `TLSAdvancedConfiguration`

<pre>
{
    "cipherSuites": array,
}
</pre>

| Fields         | Description                                                                                                                                                                                                                                              |
|:---------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `cipherSuites` | List of cipher suites to be enforced. Default to <ul><li>`TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384`</li><li>`TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384`</li><li>`TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA`</li><li>`TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA`</li></ul> Go [here](#supported-cipher-suites) to view all supported cipher suites|

### `TokenValidationConfiguration`

<pre>
{
    "issuers": array,
    "cache": object(<a href="#cacheconfiguration">CacheConfiguration</a>),
    "introspect": object(<a href="#introspectionconfiguration">IntrospectionConfiguration</a>),
    "jws": object(<a href="#jwsconfiguration">JWSConfiguration</a>),
    "jwt": object(<a href="#jwtconfiguration">JWTConfiguration</a>),
}
</pre>
| Fields       | Description                                                                            |
|:-------------|:---------------------------------------------------------------------------------------|
| `issuers`    | Array of strings specifying approved entities who issue tokens                         |
| `cache`      | Configuration of a token cache                                                         |
| `introspect` | Configuration of how to connect to introspection API                                   |
| `jws`        | Configuration of how to obtain JSON Web Keyset in order to validate JSON Web Signature |
| `jwt`        | Configuration of how to handle JSON Web Token                                          |

### `CacheConfiguration`

An LRU cache which also checks for expiration before returning the value.
Below is the default configuration if not specified

<pre>
{
    "limit": 80,
    "expirationInSeconds": 3600
}
</pre>

| Fields                | Description                        |
|:----------------------|:-----------------------------------|
| `limit`               | Max number of items in the cache   |
| `expirationInSeconds` | Expiry time for a cache item       |

### `IntrospectionConfiguration`

<pre>
{
    "endpoint": string,
    "authentication": object(<a href="#authenticationconfiguration">AuthenticationConfiguration</a>),
    "tlsConnection": object(<a href="#tlsconnectionconfiguration">TLSConnectionConfiguration</a>)
}
</pre>

| Fields           | Description                                                   |
|:-----------------|:--------------------------------------------------------------|
| `endpoint`       | Introspection API endpoint                                    |
| `authentication` | Configuration of how to authenticate when invoking `endpoint` |
| `tlsConnection`  | Configuration of TLS when connecting to `endpoint`            |

### `AuthenticationConfiguration`

<pre>
{
    "method": string,
    "credentials": map(string-><a href="#environmentawaredvalue">EnvironmentAwaredValue</a>)
}
</pre>

| Fields        | Description                                                                                                                                                                                                             |
|:--------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `method`      | Defines authentication mechanism. Supported values are <ul><li>`client_secret_basic`: basic authentication</li><li>`client_secret_form`: form authentication</li><li>`private_key`: mutual TLS authentication</li></ul> |
| `credentials` | Defines key value pair used for the given authentication mechanism above. See below for the supported keys                                                                                                              |

| Method                | Keys                       |
|:----------------------|:---------------------------|
| `client_secret_basic` | `clientId`, `clientSecret` |
| `client_secret_form`  | `clientId`, `clientSecret` |
| `private_key`         | `certFile`, `keyFile`      |

### `TLSConnectionConfiguration`

<pre>
{
    "insecureSkipVerify": bool,
    "certFile": <a href="#environmentawaredvalue">EnvironmentAwaredValue</a>,
    "caFile": <a href="#environmentawaredvalue">EnvironmentAwaredValue</a>
}
</pre>

| Fields               | Description                                                                                 |
|:---------------------|:--------------------------------------------------------------------------------------------|
| `insecureSkipVerify` | If true, do not verify server TLS certificate                                               |
| `certFile`           | Location to a file storing server certificate in PEM format. Default is `server.crt`        |
| `caFile`             | Location to a file storing server CA certificate in PEM format. Default is `server.ca.cert` |

### `JWSConfiguration`

<pre>
{
    "endpoint": string,
    "tlsConnection": object(<a href="#tlsconnectionconfiguration">TLSConnectionConfiguration</a>)
}
</pre>

| Fields          | Description                                        |
|:----------------|:---------------------------------------------------|
| `endpoint`      | API endpoint to obtain JSON Web Keyset             |
| `tlsConnection` | Configuration of TLS when connecting to `endpoint` |

### `JWTConfiguration`

<pre>
{
    "authorizationField": string,
    "preferIntrospection": bool
}
</pre>

| Fields                | Description                                                                           |
|:----------------------|:--------------------------------------------------------------------------------------|
| `authorizationField`  | Claim field name that is used to extract scopes for authorization. Default to `scope` |
| `preferIntrospection` | If true, introspection (if defined) result is used                                    |

### `EnvironmentAwaredValue`

A regular string which allows value being read from an environment variable 
by specifying an URI with `env` scheme. For example: `env://MY_VAR` will return
value from `MY_VAR` environment variable.

### Supported Cipher Suites

- `TLS_RSA_WITH_RC4_128_SHA`
- `TLS_RSA_WITH_3DES_EDE_CBC_SHA`
- `TLS_RSA_WITH_AES_128_CBC_SHA`
- `TLS_RSA_WITH_AES_256_CBC_SHA`
- `TLS_RSA_WITH_AES_128_CBC_SHA256`
- `TLS_RSA_WITH_AES_128_GCM_SHA256`
- `TLS_RSA_WITH_AES_256_GCM_SHA384`
- `TLS_ECDHE_ECDSA_WITH_RC4_128_SHA`
- `TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA`
- `TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA`
- `TLS_ECDHE_RSA_WITH_RC4_128_SHA`
- `TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA`
- `TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA`
- `TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA`
- `TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256`
- `TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256`
- `TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256`
- `TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256`
- `TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384`
- `TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384`
- `TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305`
- `TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305`
- `TLS_AES_128_GCM_SHA256`
- `TLS_AES_256_GCM_SHA384`
- `TLS_CHACHA20_POLY1305_SHA256`

## OAuth2 Authz Server Integration

Examples on how to integrate Quorum Security Plugin with an OAuth2 Authorization Server are [here](https://github.com/jpmorganchase/quorum-security-plugin-enterprise/examples).

## OAuth2 Scopes

Scope is a mechanism to limit a client's access to protected resources
in Quorum Client RPC server. A client can request one ore more scopes 
from a token endpoint of an OAuth2 Provider. The access token issued to 
the client will be limited to the scopes granted.

The scope syntax is as follow:
```text
    scope := "rpc://"rpc-string

    rpc-string := service-name delimiter method-name
   
    service-name := string
   
    delimiter := "." or "_"
   
    method-name := string
```

### Examples

#### Protecting APIs

| Scope                          | Description                                                                                 |
|:-------------------------------|:--------------------------------------------------------------------------------------------|
| `rpc://web3.clientVersion`     | Allow access to `web3_clientVersion` API                                                    |
| `rpc://eth_*` <br/>or `rpc://eth_`   | Allow access to all APIs under `eth` namespace                                              |
| `rpc://*_version` <br/>or `rpc://_version` | Allow access to `version` method of all namespaces. <br/>E.g.: `net_version`, `ssh_version` |

## Change Log

### v1.0.0

Initial release
