# JSON-RPC Security Features

## Native Transport Layer Security
 
The native Transport Layer Security (TLS) introduces an encryption layer
to the JSON-RPC request/response communication channel for both HTTP,
and Web Socket listeners. By using a simple configuration flag this
feature allows the automatic generation of self signed certificate for
testing environment, or a smooth integration with certificate
authorities for enterprise deployment.
 
## Enterprise Authorization Protocol Integration
 
Enterprise authorization protocol integration introduces an access
control layer that authorizes each JSON RPC invocation to an atomic
module function level (E.g `personal_OpenWallet`) using
[OAuth 2.0](https://tools.ietf.org/html/rfc6749) industry standard
protocol. This feature allows managing distributed application (dApps),
and Quorum Clients access control in an efficient approach.