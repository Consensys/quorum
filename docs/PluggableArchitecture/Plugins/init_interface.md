<!-- This is auto generated file from running `go generate` in plugin/proto folder. Please do not edit -->



<a name="init.proto"></a>

## init.proto
This plugin interface specifies how a plugin can be initialized.

It is __mandatory__ that every plugin must implement this RPC service

### Services


<a name="proto.PluginInitializer"></a>

#### `PluginInitializer`
Plugin Manager to initialize the plugin after plugin process is started successfully

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Init | [PluginInitialization.Request](#proto.PluginInitialization.Request) | [PluginInitialization.Response](#proto.PluginInitialization.Response) |  |

 <!-- end services -->

### Messsages


<a name="proto.PluginInitialization"></a>

#### `PluginInitialization`







<a name="proto.PluginInitialization.Request"></a>

#### `PluginInitialization.Request`
Initialization data for the plugin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hostIdentity | [string](#string) |  | `geth` node identity |
| rawConfiguration | [bytes](#bytes) |  | raw configuration to be processed by the plugin |






<a name="proto.PluginInitialization.Response"></a>

#### `PluginInitialization.Response`






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

