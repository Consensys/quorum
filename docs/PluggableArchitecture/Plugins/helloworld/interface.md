<!-- This is auto generated file from running `go generate` in plugin/proto folder. Please do not edit -->



<a name="helloworld.proto"></a>

## helloworld.proto
This plugin interface is to demonstrate a hello world plugin example

### Services


<a name="proto.PluginGreeting"></a>

#### `PluginGreeting`
Greeting remote service saying Hello in English and Spanish

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Greeting | [PluginHelloWorld.Request](#proto.PluginHelloWorld.Request) | [PluginHelloWorld.Response](#proto.PluginHelloWorld.Response) |  |

 <!-- end services -->

### Messsages


<a name="proto.PluginHelloWorld"></a>

#### `PluginHelloWorld`
A wrapper logically groups other messages






<a name="proto.PluginHelloWorld.Request"></a>

#### `PluginHelloWorld.Request`



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

