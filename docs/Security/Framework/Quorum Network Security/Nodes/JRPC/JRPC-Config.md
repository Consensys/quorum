Quorum Client supports two formats for configuring JSON RPC security:

- Quorum [JSON](#json-format) configuration file (via `--rpcsec.config`
  flag)
- Standard Ethereum [TOML](#toml-format) configuration file (via
  `--config` flag)

Security configuration in TOML file is selected over JSON file if both
are provided.

Additionally, `--rpcsec` flag must be provided in order to enable the
security features for JSON RPC.

## How-to

!!! warning
    This sample is meant as a guide only. Do **not** copy and
    paste the entire content into your application. Rather, pick only the
    properties that you need.

### JSON

> `geth ... --rpcsec --rpcsec.config <path_to_json_config_file>`

Below is the complete supported configuration

```json
{
  "resourceServer": {                      
    "id": "quorum",
    "tls": {
      "auto": false,
      "certFile": "cert.pem",
      "keyFile": "key.pem"
    }
  },
  "authorizationServer": {
    "issuer": "http://goquorum.com/oauth",
    "introspect" : {
      "endpoint" : "https://localhost:5000/oauth/introspect",
      "cache": {
        "limit": 80,
        "expirationInSeconds": 3600
      },
      "authentication": {
        "method": "client_secret_basic",
        "credentials": {
          "clientId": "quorum",
          "clientSecret": "admin"
        }
      },
      "tlsConnection": {
        "insecureSkipVerify": false, 
        "certFile": "server.crt",
        "caFile": "server.ca.crt"
      }
    },
    "jws": {
      "endpoint": "https://localhost:5000/keys",
      "tlsConnection": {
        "insecureSkipVerify": false, 
        "certFile": "server.crt",
        "caFile": "server.ca.crt"
      }
    },
    "jwt": {
      "authorizationField": "scope"
    }
  }
}
```

### TOML

> `geth ... --rpcsec --config <path_to_toml_configuration_file>`

Below is the complete supported configuration. Noted that [TOML
configuration file](https://github.com/ethereum/go-ethereum/blob/master/README.md#configuration)
can also be used to pass numerous flags to `geth` binary

```toml
[Node]
...<snip>...

[Node.RPCSecurity]
    [Node.RPCSecurity.ResourceServer]
        Id = "node1"
        [Node.RPCSecurity.ResourceServer.TLS]
            AutoGenerate = true
            CertFile = "/tmp/node1_cert.pem"
            KeyFile = "/tmp/node1_key.pem"
    [Node.RPCSecurity.AuthorizationServer]
        Issuer = "http://goquorum.com/oauth"
        [Node.RPCSecurity.AuthorizationServer.Cache]
            Limit = 80
            ExpirationInSeconds = 3600
        [Node.RPCSecurity.AuthorizationServer.Introspect]
            Endpoint = "https://localhost:5000/oauth/introspect"
            [Node.RPCSecurity.AuthorizationServer.Introspect.Authentication]
                Method = "client_secret_basic"
                [Node.RPCSecurity.AuthorizationServer.Introspect.Authentication.Credentials]
                    clientId = "introspect"
                    clientSecret = "abc"
            [Node.RPCSecurity.AuthorizationServer.Introspect.TLSConnection]
                InsecureSkipVerify = false
                CertFile = "cert.pem"
                CaFile   = "ca.pem"
        [Node.RPCSecurity.AuthorizationServer.JSONWebSignature]
            Endpoint = "https://localhost:5000/keys"
            [Node.RPCSecurity.AuthorizationServer.JSONWebSignature.TLSConnection]
                InsecureSkipVerify = false
                CertFile = "cert.pem"
                CaFile   = "ca.pem"
        [Node.RPCSecurity.AuthorizationServer.JSONWebToken]
            AuthorizationField = "scope"
```

## Structure

This section defines the structure of security configuration file. Each
listed item shows JSON object/field and TOML table/key

### `resourceServer`

TOML Table: `[Node.RPCSecurity.ResourceServer]`

Configures Quorum node as a resource server handling authenticated
requests. Below is the properties

- `id`/`Id` - Mandatory (if neither `--identity` or `UserIdent` is
  provided nor no `authorizationServer` configuration)

    Identify of the running node. Take value from `--identity` command
    line flag or `UserIdent` TOML configuration key. This is used to
    validate against the audience of a received access token to make sure
    the running node is the intended receipient of the token
  
- `tls`/`[Node.RPCSecurity.ResourceServer.TLS]` - Optional

    Configures
    [TLS](/Security/QuorumClient/JSONRPC-Features/#native-transport-layer-security)
    for JSON RPC server
  
    - `auto`/`AutoGenerate` - Optional (default to `false`)
        
        If `true`, Quorum Client will generate a self-signed
        certificate. Generated PEM certificate and key will be stored in
        files given in `certFile` and `keyFile` properties below
        
    - `certFile`/`CertFile` - Optional (default to `cert.pem`)
    
        Local file path to certificate. If the certificate is signed by
        a certificate authority, the `certFile` should be the
        concatenation of the server's certificate, any intermediates,
        and the CA's certificate.
        
    - `keyFile`/`KeyFile` - Optional (default to `key.pem`)
    
        Local file path to the above certificate-matching private key

### `authorizationServer`

TOML Table: `[Node.RPCSecurity.AuthorizationServer]`

Configuration to allow the JSON RPC server to decode and validate access
tokens issued by an external authorization server. Therefore either
`introspect` or `jws` must be provided.

- `issuer`/`Issuer` - Mandatory

    Unique identifier for the entity that issued the access token

- `cache`/`[Node.RPCSecurity.AuthorizationServer.Cache]` - Optional

    Configuration of a global LRU cache for validated tokens

    - `limit`/`Limit` - Optional (default to 80)
        
        Maximum number of items being cached
        
    - `expirationInSeconds`/`ExpirationInSeconds` - Optional (default to
      3600)

        Expiry time for each cached item
      
- `introspect`/`[Node.RPCSecurity.AuthorizationServer.Introspect]` -
  Optional
  
    Configuration of
    [introspection](https://tools.ietf.org/html/rfc7662) client
  
    - `endpoint`/`Endpoint` - Mandatory
    
        Introspection API endpoint
    
    - `authentication`/`[Node.RPCSecurity.AuthorizationServer.Introspect.Authentication]`
      \- Mandatory
      
        To prevent token scanning attacks. This must be configured.
      
        - `method`/`Method` - Mandatory
      
            Supported values are:
        
            - `client_secret_basic`: using basic authentication
            - `client_secret_form`: using form authentication
            - `private_key`: mutual TLS authentication
            
        - `credentials`/`Credentials` - Mandatory
    
            Key value pairs that support the above authentication
            method. Supported keys for each method:
            
            | Method | Keys |
            |:-------|:-----|
            | `client_secret_basic` | `clientId` and `clientSecret` |
            | `client_secret_form` | `clientId` and `clientSecret` |
            | `private_key` | `certFile` and `caFile` |
    
    - `tlsConnection`/`[Node.RPCSecurity.AuthorizationServer.Introspect.TLSConnection]`
       \- Optional
    
        - `insecureSkipVerify`/`InsecureSkipVerify` - Optional (default to `false`)
        
            If `true`, disable verification of server's TLS certificate
            on connection
        
        - `certFile`/`CertFile` - Optional
        
            Server's TLS certificate PEM file on connection
            
        - `caFile`/`CaFile` - Optional
        
            CA certificate PEM file for provided server's TLS certificate on connection
        
- `jws`/`[Node.RPCSecurity.AuthorizationServer.JSONWebSignature]` -
  Optional
  
    Configuration of how to obtain JSON Web Keysets for validating JWT
    signature
  
    - `endpoint`/`Endpoint` - Mandatory
    
        URL to retrieve JSON Web Keysets
  
    - `tlsConnection`/`[Node.RPCSecurity.AuthorizationServer.Introspect.TLSConnection]`
       \- Optional
    
        - `insecureSkipVerify`/`InsecureSkipVerify` - Optional (default to `false`)
        
            If `true`, disable verification of server's TLS certificate
            on connection
        
        - `certFile`/`CertFile` - Optional
        
            Server's TLS certificate PEM file on connection
            
        - `caFile`/`CaFile` - Optional
        
            CA certificate PEM file for provided server's TLS certificate on connection
            
- `jwt`/`[Node.RPCSecurity.AuthorizationServer.JSONWebToken]` - Optional
    - `authorizationField`/`AuthorizationField` - Optional (default to `
      scope`)
      
        Field name inside JWT token being used to contain authorization
        values

## Geth Attach
`geth attach` and `geth monitor` has additional flags allowing to
connect to secured Quorum node

    --rpcclitoken value             RPC Client access token
    --rpcclitls.insecureskipverify  Disable verification of server's TLS certificate on connection by client
    --rpcclitls.cert value          Server's TLS certificate PEM file on connection by client
    --rpcclitls.cacert value        CA certificate PEM file for provided server's TLS certificate on connection by client

!!! tip
    Run `geth attach --help` or `geth monitor --help` will display the
    same flag description as above

Check out some examples
[here](/Security/QuorumClient/JSONRPC-Examples/#attach-to-a-secured-geth)