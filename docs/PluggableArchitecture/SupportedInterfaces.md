

<a name="init.proto"></a>

## init.proto
This plugin interface specifies how a plugin can be initialized.

It is __mandatory__ that every plugin must implement this RPC service

### Services

<a name="proto.PluginInitializer"></a>

#### PluginInitializer
`Required`
Plugin Manager to initialize the plugin after plugin process is started successfully

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Init | [PluginInitialization.Request](#proto.PluginInitialization.Request) | [PluginInitialization.Response](#proto.PluginInitialization.Response) |  |

 <!-- end services -->

### Messages

<a name="proto.PluginInitialization"></a>

#### PluginInitialization







<a name="proto.PluginInitialization.Request"></a>

#### PluginInitialization.Request
Initialization data for the plugin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hostIdentity | [string](#string) |  | `geth` node name |
| rawConfiguration | [bytes](#bytes) |  | raw configuration to be processed by the plugin |






<a name="proto.PluginInitialization.Response"></a>

#### PluginInitialization.Response






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->



<a name="helloWorld.proto"></a>

## helloWorld.proto
This plugin interface is to demonstrate a hello world example only.

### Services

<a name="proto.PluginGreeting"></a>

#### PluginGreeting
`Optional`
Ping Pong a message

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Greeting | [PluginHelloWorld.Request](#proto.PluginHelloWorld.Request) | [PluginHelloWorld.Response](#proto.PluginHelloWorld.Response) |  |

 <!-- end services -->

### Messages

<a name="proto.PluginHelloWorld"></a>

#### PluginHelloWorld







<a name="proto.PluginHelloWorld.Request"></a>

#### PluginHelloWorld.Request
Initialization data for the plugin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msg | [string](#string) |  | a message to the plugin |






<a name="proto.PluginHelloWorld.Response"></a>

#### PluginHelloWorld.Response



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msg | [string](#string) |  | a response message from the plugin |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |
