# Configuration file

A `.json` file including required configuration details must be provided using the `-configfile` command-line property when starting Tessera.

Many configuration options can be overridden using the command-line.  See the [Using CLI to override config](../Using%20CLI%20to%20override%20config) page for more information.

## Configuration options
The configuration options are explained in more detail in this section.  Configuration options that require more than a brief explanation are covered in separate pages.

### Cryptographic Keys
See [Keys page](../Keys).

### Whitelist
If set to true, the `peers` list will be used as the whitelisted urls for the Tessera node:
```
"useWhiteList": true,
```

---

### Database
Tessera's database uses JDBC to connect to an external database. Any valid JDBC URL may be specified, refer to your providers details to construct a valid JDBC URL.

```
"jdbc": {
  "url": "[JDBC URL]",
  "username": "[JDBC Username]",
  "password": "[JDBC Password]"
}
```

---

### Server
> **For Tessera versions prior to 0.8:** See [Legacy Server Settings](../Legacy%20server%20settings).

To allow for a greater level of control, Tessera's API has been separated into distinct groups.  Each group is only accessible over a specific server type.  Tessera can be started with different combinations of these servers depending on the functionality required.  This is defined in the configuration and determines the APIs that are available and how they are accessed.
 
The possible server types are:

- `P2P` - Tessera uses this server to communicate with other Transaction Managers (the URI for this server can be shared with other nodes to be used in their `peer` list - see below)
- `Q2T` - This server is used for communications between Tessera and its corresponding Quorum node
- `ENCLAVE` - If using a remote enclave, this defines the connection details for the remote enclave server (see the [Enclave docs](../../Tessera%20Services/Enclave#types-of-enclave) for more info) 
- `ThirdParty` - This server is used to expose certain Transaction Manager functionality to external services such as Quorum.js
- `ADMIN` - This server is used for configuration management. It is intended for use by the administrator of the Tessera node and is not recommended to be advertised publicly

The servers to be started are provided as a list:
```
"serverConfigs": [
   ...<server settings>...
]
```

Each server is individually configurable and can advertise over HTTP, HTTPS or a Unix Socket.  The format of an individual server config is slightly different between Tessera v0.9 and v0.8:

#### Server configuration (v0.9)
HTTP:
```
{
    "app": "<app type>",
    "enabled": <boolean>,
    "serverAddress":"http://[host]:[port]/[path]
    "communicationType" : <enum>, // "REST" or "GRPC"
}
```
HTTPS:
```
{
    "app": "<app type>",
    "enabled": <boolean>,
    "serverAddress":"https://[host]:[port]/[path]
    "communicationType" : <enum>, // "REST" or "GRPC"
    "sslConfig": {
        ...<SSL settings, see below>...
    }
}
```
Unix Socket:
```
{
    "app": "<app type>",
    "enabled": <boolean>,
    "serverAddress":"unix://[path],
    "communicationType" : "REST"
}
```

#### Server configuration (v0.8)
HTTP:
```
{
    "app": "<app type>",
    "enabled": <boolean>,
    "serverSocket":{
        "type": "INET",
        "port": <int>, //The port to advertise and bind on (if binding address not set)
        "hostName": <string> // The hostname to advertise and bind on (if binding address not set)
    },
    "communicationType" : <enum>, // "REST" or "GRPC"
    "bindingAddress": <string> //An address to bind the server to that overrides the one defined above
}
```

HTTPS:
```
{
    "app": "<app type>",
    "enabled": <boolean>,
    "serverSocket":{
        "type": "INET",
        "port": <int>, //The port to advertise and bind on (if binding address not set)
        "hostName": <string> // The hostname to advertise and bind on (if binding address not set)
    },
    "communicationType" : <enum>, // "REST" or "GRPC"
    "bindingAddress": <string>, //An address to bind the server to that overrides the one defined above
    "sslConfig": {
       ...<SSL settings, see below>...
    }
}
```

Unix Socket: 
```
{
    "app": "<app type>",  
    "enabled": <boolean>, 
    "serverSocket":{
        "type":"UNIX",
        "path": <string> //the path of the unix socket to create
    },
    "communicationType" : "UNIX_SOCKET"
}
```

### TLS/SSL: server sub-config
See [TLS/SSL](../TLS) page.

### CORS: server sub-config
For the ThirdParty server type it may be relevant to configure CORS.
```
{
    "app":"ThirdParty",
    "enabled": true,
    "serverAddress": "http://localhost:9081",
    "communicationType" : "REST",
    "cors" : {
        "allowedMethods" : ["GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"],
        "allowedOrigins" : ["http://localhost:63342"],
        "allowedHeaders" : ["content-type"],
        "allowCredentials" : true
    }
},
```
The configurale fields are:
* `allowedMethods` - the list of allowed HTTP methods. If omitted the default list containing `"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"` is used.
* `allowedOrigins` - the list of domains from which to accept cross origin requests (browser enforced). Each entry in the list can contain the "*" (wildcard) character which matches any sequence of characters. Ex: "*locahost" would match "http://localhost" or "https://localhost". There is no default for this field.
* `allowedHeaders` - the list of allowed headers. If omitted the request `Access-Control-Request-Headers` are copied into the response as `Access-Control-Allow-Headers`.
* `allowCredentials` - the value for the `Access-Control-Allow-Credentials` response header. If omitted the default `true` value would be used.  

### InfluxDB Config: server sub-config
Configuration details to allow Tessera to record monitoring data to a running InfluxDB instance.
```
"influxConfig": {
  "hostName": "[Hostname of Influx instance]",
  "port": "[Port of Influx instance]",
  "pushIntervalInSecs": "[How often to push data to InfluxDB]",
  "dbName": "[Name of InfluxDB]"
}
```

---

### Peers
A list of URLs used by Tessera to communicate with other nodes.  Peer info is shared between nodes during runtime (however, please note the section on `Peer Discovery` below).
```
"peer": [
  {
    "url": "http://myhost.com:9000"
  },
  {
    "url": "http://myhost.com:9001"
  },
  {
    "url": "http://myhost.com:9002"
  }
]
```

### Disabling peer discovery
If peer discovery is disabled, then **only** peers defined in the configuration file will be communicated with; any peers notified by other nodes will be ignored. This allows nodes to be 'locked down' if desired.
```
"disablePeerDiscovery": true
```

---

### Always-send-to
It is possible to configure a node that will be sent a copy of every transaction, even if it is not specified as a party to the transaction. This could be used, for example, to send a copy of every transaction to a node for audit purposes. Specify the public keys to forward transactions onto, and these will be included as if you had specified them on the `privateFor` field to start with.

```
"alwaysSendTo":["<public key 1>", "<public key 2>"]
```

---

### Remote-Key-Validation
Tessera provides an API `/partyinfo` on Tessera P2P server to discover all the peers in the network. In order to prevent attackers trying to inject malicious addresses against public keys, where they will try to assign the address to direct private transactions to them instead of the real owner of the key, we have added a feature to enable node level validation on the remote key that checks the remote node does in fact own the keys that were advertised. Only after the keys are validated with the remote node to ensure they own them, the keys are added to the local network info (partyinfo) store.

Default configuration for this is `false` as this is BREAKABLE change to lower versions to Tessera 0.10.0. To enable this, simple set below parameter to true in the configuration:

```
 "features": {
    "enableRemoteKeyValidation": true
  }
```

---


