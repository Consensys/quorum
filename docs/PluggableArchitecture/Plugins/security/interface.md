<!-- This is auto generated file from running `go generate` in plugin/proto folder. Please do not edit -->



<a name="security.proto"></a>

## security.proto
This plugin interface provides services to secure `geth` RPC servers, which includes:

- TLS configuration to enable HTTPS/WSS servers
- Authentication

### Services


<a name="proto.AuthenticationManager"></a>

#### `AuthenticationManager`
`Required`
RPC service authenticate the preauthenticated token. Response is the token containing expiry date and granted authorities

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Authenticate | [AuthenticationToken](#proto.AuthenticationToken) | [PreAuthenticatedAuthenticationToken](#proto.PreAuthenticatedAuthenticationToken) | Perform authentication of the token. Return a token that contains expiry date and granted authorities |


<a name="proto.TLSConfigurationSource"></a>

#### `TLSConfigurationSource`
`Optional`
RPC service to provide TLS configuration to enable HTTPS/WSS in `geth` RPC Servers

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [TLSConfiguration.Request](#proto.TLSConfiguration.Request) | [TLSConfiguration.Response](#proto.TLSConfiguration.Response) |  |

 <!-- end services -->

### Messsages


<a name="proto.AuthenticationToken"></a>

#### `AuthenticationToken`
Representing the access token for an authentication request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rawToken | [bytes](#bytes) |  |  |






<a name="proto.GrantedAuthority"></a>

#### `GrantedAuthority`
Representing a permission being extracted from access token by the plugin implementation.
This permission is then stored in security context of a request and
used internally to decide if the access is granted/denied


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service | [string](#string) |  | `geth` RPC API namespace. E.g.: rpc, eth, admin, debug, ... |
| method | [string](#string) |  | `geth` RPC API function. E.g.: nodeInfo, blockNumber, ... |
| raw | [string](#string) |  | raw string of the the granted authority value. This gives plugin implementation freedom to interpret the value |






<a name="proto.PreAuthenticatedAuthenticationToken"></a>

#### `PreAuthenticatedAuthenticationToken`
Representing an authenticated principal after `AuthenticationToken` has been processed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rawToken | [bytes](#bytes) |  |  |
| expiredAt | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| authorities | [GrantedAuthority](#proto.GrantedAuthority) | repeated |  |






<a name="proto.TLSConfiguration"></a>

#### `TLSConfiguration`
A wrapper message to logically group other messages






<a name="proto.TLSConfiguration.Data"></a>

#### `TLSConfiguration.Data`
TLS configuration data for `geth`


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| keyPem | [bytes](#bytes) |  | Private key in PEM format |
| certPem | [bytes](#bytes) |  | Certificate in PEM format |
| cipherSuites | [uint32](#uint32) | repeated | List of cipher suites constants being supported by the server |






<a name="proto.TLSConfiguration.Request"></a>

#### `TLSConfiguration.Request`
It's an empty Request received by RPC service






<a name="proto.TLSConfiguration.Response"></a>

#### `TLSConfiguration.Response`
Response from RPC service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [TLSConfiguration.Data](#proto.TLSConfiguration.Data) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

