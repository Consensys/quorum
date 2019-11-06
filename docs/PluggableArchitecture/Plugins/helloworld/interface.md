<!-- This is auto generated file from running `go generate` in plugin/proto folder. Please do not edit -->



<a name="helloworld.proto"></a>

## helloworld.proto
This plugin interface is to demonstrate a hello world example only.

### Services


<a name="proto.PluginGreeting"></a>

#### `PluginGreeting`
Greeting service

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Greeting | [PluginHelloWorld.Request](#proto.PluginHelloWorld.Request) | [PluginHelloWorld.Response](#proto.PluginHelloWorld.Response) |  |

 <!-- end services -->

### Messsages


<a name="proto.PluginHelloWorld"></a>

#### `PluginHelloWorld`







<a name="proto.PluginHelloWorld.Request"></a>

#### `PluginHelloWorld.Request`
Initialization data for the plugin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msg | [string](#string) |  | a message to the plugin |






<a name="proto.PluginHelloWorld.Response"></a>

#### `PluginHelloWorld.Response`



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msg | [string](#string) |  | a response message from the plugin |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

