This section contains examples on how to configure Quorum Client to
enable [security features](/Security/QuorumClient/JSONRPC-Features/) for
JSON RPC server.

Security configuration can be setup with JSON or TOML format. Please
refer [here](/Security/QuorumClient/JSONRPC-Configs/) for more details.

!!! danger 
    These examples aremeant as a guide only. Do ==**not**== copy and
    paste the entire content into your application.

## Start a secured Quorum Client

```bash tab="Using JSON"
geth <your_usual_geth_args> \
    --rpc --rpcport 22000 \
    --ws  --wsport  23000 \
    --rpcsec --rpcsec.config rpc-sec-config.json
```

```bash tab="Using TOML"
geth <your_usual_geth_args> \
    --rpc --rpcport 22000 \
    --ws  --wsport  23000 \
    --rpcsec --config node1.toml
``` 

## TLS using self-signed certificate

Define TLS configuration, set `auto` to `true` and path to certificate
and primary key files being generated.

```bash tab="JSON"
cat <<EOF > rpc-sec-config.json
{
    "resourceServer": {
        "tls": {
            "auto": true,
            "certFile": "/mycerts/cert.pem",
            "keyFile": "/mycerts/key.pem"
        }
    }
}
EOF
```

```bash tab="TOML"
cat <<EOF > node1.toml
[Node.RPCSecurity]
    [Node.RPCSecurity.ResourceServer]
        [Node.RPCSecurity.ResourceServer.TLS]
            AutoGenerate = true
            CertFile = "/mycerts/cert.pem"
            KeyFile = "/mycerts/key.pem"
EOF
```

## TLS using your own certificate

In order to use your own certificate, set `auto` to `false` and provide
certificate and private key files.

```bash tab="JSON"
cat <<EOF > rpc-sec-config.json
{
    "resourceServer": {
        "tls": {
            "auto": false,
            "certFile": "/mycompany/node1_cert.pem",
            "keyFile": "/mycompany/node1_key.pem"
        }
    }
}
EOF
```

```bash tab="TOML"
cat <<EOF > node1.toml
[Node.RPCSecurity]
    [Node.RPCSecurity.ResourceServer]
        [Node.RPCSecurity.ResourceServer.TLS]
            AutoGenerate = false
            CertFile = "/mycompany/node1_cert.pem"
            KeyFile = "/mycompany/node1_key.pem"
EOF
```

## Attach to a secured-TLS `geth`

Connect to the node with `--rpcclitls.insecureskipverify` to ignore the
Server's certificate validation.

```bash
geth attach https://localhost:22000 --rpcclitls.insecureskipverify    
geth attach wss://localhost:23000   --rpcclitls.insecureskipverify    
```

Or connect to node using the above generated server certifcate.

```bash
geth attach https://localhost:22000 --rpcclitls.cert /mycerts/cert.pem    
geth attach wss://localhost:23000   --rpcclitls.cert /mycerts/cert.pem    
```

Refer [here](/Security/QuorumClient/geth-attach-monitor/) for more
details on available flags for `geth` commands.

## Protect JSON RPC

Quorum Client can be configured to protect JSON RPC APIs. Quorum Client
will then behave as a
[Resource Server](/Security/QuorumClient/JSONRPC-Features/#enterprise-authorization-protocol-integration).

To demonstrate this feature, we use `quorum-sample-oauth2-server` acting
as an authorization server which grants access tokens to JSON RPC client
and validates access tokens requested by Quorum Client.

### Start sample authorization server

Follow instruction in
[`quorum-sample-oauth2-server`](https://github.com/jpmorganchase/quorum-sample-oauth2-server)
repository to download a sample authorization server

    ./qoauth2server

    ... INF Saved client aud=<na> client_id=introspect client_secret=abc scope=<na>
    ... INF SetClient id=introspect secret=abc
    ... INF Saved client aud="node1 node2 node3" client_id=quorum client_secret=admin scope=*.*
    ... INF SetClient id=quorum secret=admin
    ... INF Saved client aud=node1 client_id=foo client_secret=bar scope="rpc.* admin.nodeInfo"
    ... INF SetClient id=foo secret=bar
    ... INF Starting server certFile=cert.pem keyFile=key.pem url=https://localhost:5000
    
By default, the server will be started locally on port `5000` with
auto-generate TLS certificate

    ... INF Starting server certFile=cert.pem keyFile=key.pem url=https://localhost:5000 

and with 3 preconfigured credentials

    ... INF Saved client aud=<na> client_id=introspect client_secret=abc scope=<na>
    ... INF Saved client aud="node1 node2 node3" client_id=quorum client_secret=admin scope=*.*
    ... INF Saved client aud=node1 client_id=foo client_secret=bar scope="rpc.* admin.nodeInfo"

- `introspect/abc` is used for resource server to authenticate when
calling introspection API
- `quorum/admin` is granted access to all APIs and the generated
  token is restricted to Quorum Client with identities: `node1`,
  `node2` and `node3`
- `foo/bar` is granted access to `rpc` API module and
  `admin_nodeInfo` API and the generated token is restricted to
  Quorum Client with identity `node1`
  
> For more available options when starting the sample authorization
> server, run `./qoauth2server --help`

### Start Quorum Client resource server

#### Define Security Configuration

If authorization server supports
[token introspection API](https://tools.ietf.org/html/rfc7662), provide
`introspect` configuration block.

!!! tip
    Quorum Client JSON RPC server can be configured to support only TLS
    or only authorization or both

```bash tab="JSON"
cat <<EOF > rpc-sec-config.json
{
    "resourceServer": {
        "id": "node1"
        "tls": {
            "auto": true,
            "certFile": "/mycerts/cert.pem",
            "keyFile": "/mycerts/key.pem"
        }
    },
    "authorizationServer": {
        "issuer": "http://goquorum.com/oauth",
        "introspect" : {
            "endpoint" : "https://localhost:5000/oauth/introspect",
            "authentication": {
                "method": "client_secret_basic",
                "credentials": {
                    "clientId": "introspect",
                    "clientSecret": "abc"
                }
            },
            "tlsConnection": {
                "insecureSkipVerify": true 
            }
        }
    }
}
EOF
```

```bash tab="TOML"
cat <<EOF > node1.toml
[Node.RPCSecurity]
    [Node.RPCSecurity.ResourceServer]
        Id = "node1"
        [Node.RPCSecurity.ResourceServer.TLS]
            AutoGenerate = true
            CertFile = "/mycerts/node1_cert.pem"
            KeyFile = "/mycerts/node1_key.pem"
    [Node.RPCSecurity.AuthorizationServer]
        Issuer = "http://goquorum.com/oauth"
        [Node.RPCSecurity.AuthorizationServer.Introspect]
            Endpoint = "https://localhost:5000/oauth/introspect"
            [Node.RPCSecurity.AuthorizationServer.Introspect.Authentication]
                Method = "client_secret_basic"
                [Node.RPCSecurity.AuthorizationServer.Introspect.Authentication.Credentials]
                    clientId = "introspect"
                    clientSecret = "abc"
            [Node.RPCSecurity.AuthorizationServer.Introspect.TLSConnection]
                InsecureSkipVerify = true
EOF
```

If not, provide `jws` configuration block.

```bash tab="JSON"
cat <<EOF > rpc-sec-config.json
{
    "resourceServer": {
        "id": "node1",
        "tls": {
            "auto": true,
            "certFile": "/mycerts/cert.pem",
            "keyFile": "/mycerts/key.pem"
        }
    },
    "authorizationServer": {
        "issuer": "http://goquorum.com/oauth",
        "jws": {
            "endpoint": "https://localhost:5000/keys",
            "tlsConnection": {
                "insecureSkipVerify": true
            }
        }
    }
}
EOF
```

```bash tab="TOML"
cat <<EOF > node1.toml
[Node.RPCSecurity]
    [Node.RPCSecurity.ResourceServer]
        Id = "node1"
        [Node.RPCSecurity.ResourceServer.TLS]
            AutoGenerate = true
            CertFile = "/mycerts/node1_cert.pem"
            KeyFile = "/mycerts/node1_key.pem"
    [Node.RPCSecurity.AuthorizationServer]
        Issuer = "http://goquorum.com/oauth"
        [Node.RPCSecurity.AuthorizationServer.JSONWebSignature]
            Endpoint = "https://localhost:5000/keys"
            [Node.RPCSecurity.AuthorizationServer.JSONWebSignature.TLSConnection]
                InsecureSkipVerify = true
EOF
```

Our sample authorization server grants access token which is issued by
predefined issuer `http://goquorum.com/oauth` hence it matches with the
value of `issuer`

#### Run Quorum Client

Refer to [Start a secured Quorum Client](#start-a-secured-quorum-client)
above

### Invoke a secured JSON RPC API

This section describes how JSON RPC client invokes a protected API. We
use `curl` to demonstrate the interaction. 

`geth attach` and `geth monitor` are
[also supported](/Security/QuorumClient/geth-attach-monitor/) by passing
`--rpcclitoken` to the command line

#### Request access token

By using `foo/bar` credentials created 
[above](#start-quorum-client-resource-server)

```bash
# request
curl -k -X POST https://localhost:5000/oauth/token --data "grant_type=client_credentials&client_id=foo&client_secret=bar"

# response 
{
  "access_token": "eyJhbGciOiJSUzI1NiIsImN0eSI6IkpXVCIsImtpZCI6IjU1OWE4NTZhLWQ3ZTctNGJjZC1iNzZjLTkyNTUyZmE5NGY4MyIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsibm9kZTEiXSwiZXhwIjoxNTU5ODU1NDkzLCJpYXQiOjE1NTk4NDgyOTMsImlzcyI6Imh0dHA6Ly9nb3F1b3J1bS5jb20vb2F1dGgiLCJzY29wZSI6InJwYy4qIGFkbWluLm5vZGVJbmZvIiwic3ViIjoicXVvcnVtIn0.OQZLPPi8ZUJRghHg4bBH0SZjmNVamGqYM-oiO5pIFP1WnDr0tL7zFAmPPIoBFPQfIvb_L5OSfQPltwi3Y6Pc0ozEd-_Ee_dPu7QCONNMpp1csJHaxgD_kvO6lQE3Ohw6cuucVrewXbq_qDtpcKAf17Z608T8vTqXBWIBelI6G5dWE9kJ2T3VDGCD6mLg_b1K68ewxWRL5j9qkvA7Se6MQLiwNDzlvHkElbwfrRPUbUbeKW3d7509U2Rybmf5--q0-ac6VEyWmMKVzfEITcAxZSxK_WRlJLeHYkFr-LIBzwWzuTR9qWHmIUcQANb7PnH2E7pIw_shwDLaZwStAquiAA",
  "expires_in": 3600,
  "scope": "rpc.* admin.nodeInfo",
  "token_type": "Bearer"
}

# set the env variable
export TOKEN=eyJhbGciOiJSUzI1NiIsImN0eSI6IkpXVCIsImtpZCI6IjU1OWE4NTZhLWQ3ZTctNGJjZC1iNzZjLTkyNTUyZmE5NGY4MyIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsibm9kZTEiXSwiZXhwIjoxNTU5ODU1NDkzLCJpYXQiOjE1NTk4NDgyOTMsImlzcyI6Imh0dHA6Ly9nb3F1b3J1bS5jb20vb2F1dGgiLCJzY29wZSI6InJwYy4qIGFkbWluLm5vZGVJbmZvIiwic3ViIjoicXVvcnVtIn0.OQZLPPi8ZUJRghHg4bBH0SZjmNVamGqYM-oiO5pIFP1WnDr0tL7zFAmPPIoBFPQfIvb_L5OSfQPltwi3Y6Pc0ozEd-_Ee_dPu7QCONNMpp1csJHaxgD_kvO6lQE3Ohw6cuucVrewXbq_qDtpcKAf17Z608T8vTqXBWIBelI6G5dWE9kJ2T3VDGCD6mLg_b1K68ewxWRL5j9qkvA7Se6MQLiwNDzlvHkElbwfrRPUbUbeKW3d7509U2Rybmf5--q0-ac6VEyWmMKVzfEITcAxZSxK_WRlJLeHYkFr-LIBzwWzuTR9qWHmIUcQANb7PnH2E7pIw_shwDLaZwStAquiAA
```

#### Invoke APIs
 
Access token generated by `foo/bar` credentials has been granted access
to `rpc.*` and `admin.nodeInfo` APIs hence response is successfully
returned

```bash
# request
curl -k -X POST https://localhost:22000/ \
    -H "Content-type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    --data '{"jsonrpc":"2.0","method":"rpc_modules","params":[],"id":1}'

# successful response
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "admin": "1.0",
    "debug": "1.0",
    "eth": "1.0",
    "miner": "1.0",
    "net": "1.0",
    "personal": "1.0",
    "raft": "1.0",
    "rpc": "1.0",
    "txpool": "1.0",
    "web3": "1.0"
  }
}
```

```bash
# request
curl -k -X POST https://localhost:22000/ \
    -H "Content-type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    --data '{"jsonrpc":"2.0","method":"admin_nodeInfo","params":[],"id":1}'

# successul response
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "id": "7577115fd9184a27da527128d8dba507e85f116b1f7e231ca8525fc9008a6966",
    "name": "Geth/node1/v1.8.18-stable-62087493(quorum-v2.2.3)/darwin-amd64/go1.11.6",
    "enode": "enode://ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608bc67c4b30db88e0a5a6c6390213f7acbe1153ff6d23ce57380104288ae19373ef@127.0.0.1:21000?discport=0",
    "enr": "0xf88fb840f1003757e7ad7a9bcdd4acec30ffeeb38490132a3c70ff92148ff88e7493b26c16e9708f2c6c58fa2bbc29ed9d9849183fba3c4c566dad61a079ab6a9970ec8a0c83636170c6c5836574683f826964827634826970847f00000189736563703235366b31a103ac6b1096ca56b9f6d004b779ae3728bf83f8e22453404cc3cef16a3d9b96608b83746370825208",
    "ip": "127.0.0.1",
    "ports": {
      "discovery": 0,
      "listener": 21000
    },
    "listenAddr": "[::]:21000",
    "protocols": {
      "eth": {
        "network": 10,
        "difficulty": 0,
        "genesis": "0x6a6605601e17bbfbc0a199104a05f222d11da37fe2320c023394ff1e516243a2",
        "config": {
          "chainId": 10,
          "homesteadBlock": 0,
          "eip150Block": 0,
          "eip150Hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
          "eip155Block": 0,
          "eip158Block": 0,
          "byzantiumBlock": 0,
          "isQuorum": true,
          "txnSizeLimit": 64
        },
        "head": "0x6a6605601e17bbfbc0a199104a05f222d11da37fe2320c023394ff1e516243a2"
      }
    }
  }
}
```

```bash
# request
curl -k -X POST https://localhost:22000/ \
    -H "Content-type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'

# error response
{
  "jsonrpc": "2.0",
  "id": 1,
  "error": {
    "code": -32001,
    "message": "eth_blockNumber - access denied"
  }
}
```

```bash
geth attach wss://localhost:23000 --rpcclitoken $TOKEN --rpcclitls.insecureskipverify
Welcome to the Geth JavaScript console!

 modules: admin:1.0 debug:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 raft:1.0 rpc:1.0 txpool:1.0 web3:1.0

> eth.blockNumber
Error: eth_blockNumber - access denied
    at web3.js:3143:20
    at web3.js:6347:15
    at get (web3.js:6247:38)
    at <unknown>
```