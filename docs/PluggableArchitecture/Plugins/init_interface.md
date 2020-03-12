<!-- This is auto generated file from running `go generate` in plugin/proto folder. Please do not edit -->



<a name="init.proto"></a>

## init.proto
It is __mandatory__ that every plugin must implement this RPC service

Via this service, plugins receive a raw configuration sent by `geth`.
It's up to the plugin to interpret and parse the configuration then do the initialization
to make sure the plugin is ready to serve

### Services


<a name="proto_common.PluginInitializer"></a>

#### `PluginInitializer`
`Required`
RPC service to initialize the plugin after plugin process is started successfully

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Init | [PluginInitialization.Request](#proto_common.PluginInitialization.Request) | [PluginInitialization.Response](#proto_common.PluginInitialization.Response) |  |

 <!-- end services -->

### Messsages


<a name="proto_common.PluginInitialization"></a>

#### `PluginInitialization`
A wrapper message to logically group other messages






<a name="proto_common.PluginInitialization.Request"></a>

#### `PluginInitialization.Request`
Initialization data for the plugin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hostIdentity | [string](#string) |  | `geth` node identity |
| rawConfiguration | [bytes](#bytes) |  | Raw configuration to be processed by the plugin |






<a name="proto_common.PluginInitialization.Response"></a>

#### `PluginInitialization.Response`






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

