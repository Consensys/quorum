title: Internals - Pluggable Architecture - Quorum

## Background

### Go Plugin
`geth` is written in the Go programming language. [Go 1.8 introduced](https://golang.org/doc/go1.8#plugin) a new plugin architecture 
which allows for the creation of plugins (via `plugin` build mode) and to use these plugins at runtime (via `plugin` package). 
In order to utilize this architecture, there are strict requirements in developing plugins. 

By using the network RPC interface, the plugin is independently built and distributed without having to rebuild `geth`. 
Especially with gRPC interfaces, plugins can be written in different languages (see our [examples](../PluginDevelopment/#examples)).
This makes it easy for you to build a prototype feature or even a proprietary plugin for your organization's internal use.

We use HashiCorp's [`go-plugin`](https://github.com/hashicorp/go-plugin) library as it fits our asks 
and it has been proven in many plugin-based production systems.

### Why we decided to use plugins

There are number of benefits:

- Dynamically-linked binaries (which you get when using plugins) are much smaller than statically compiled binaries.
- We value the ability to isolate failures. E.g.: Quorum client would continue mining/validating even if security plugin has crashed.
- Easily enables support for open source plugins written in languages other than Go.

## Design

```plantuml
skinparam componentStyle uml2
skinparam shadowing false
skinparam backgroundColor transparent
skinparam rectangle {
    roundCorner<<component>> 25
}

file "JSON File" as json
file "TOML File" as toml
note left of toml : Standard Ethereum Config
note right of json : Quorum Plugin Settings

node "geth" <<process>> {
    rectangle "CLI Flags" as flags
    frame "plugin.Settings" as settings {
        storage "Plugin1\nDefinition" as pd1
        storage "Plugin2\nDefinition" as pd2
        storage "Plugin Central\nConnectivity" as pcc
    }

    json <-down- flags : "via\n""--plugins"""
    toml <-down- flags : "via\n""--config"""
    flags -down-> settings : populate

    interface """node.Service""" as service
    rectangle """plugin.PluginManager""" <<geth service>> as pm
    note right of pm
    registered and managed
    as standard ""geth""
    service life cycle
    end note

    pm -up- service
    pm -up- settings

    card "arbitrary" <<component>> as arbitrary
    interface "internal1" as i1
    interface "internal2" as i2
    interface "internal3" as i3

    package "Plugin Interface 1" {
        rectangle "Plugin1" <<template>> as p1
        rectangle "Gateway1" <<adapter>> as p1gw1
        rectangle "Gateway2" <<adapter>> as p1gw2

        interface "grpc service interface1A" as grpcI1A
        interface "grpc service interface1B" as grpcI1B

        rectangle "GRPC Stub Client1" <<grpc client>> as grpcC1
    }
    
    package "Plugin Interface 2" {
        rectangle "Plugin2" <<template>> as p2    
        rectangle "Gateway" <<adapter>> as p2gw

        interface "grpc service interface2" as grpcI2

        rectangle "GRPC Stub Client2" <<grpc client>> as grpcC2
    }

    pm -- p1
    pm -- p2

    arbitrary --( i1
    arbitrary --( i2
    arbitrary --( i3

    p1gw1 -- i1
    p1gw2 -- i2
    p2gw -- i3

    p1 -- p1gw1
    p1 -- p1gw2
    p2 -- p2gw

    grpcC1 --( grpcI1A
    grpcC1 --( grpcI1B
    grpcC2 --( grpcI2

    p1gw1 --> grpcC1 : use
    p1gw2 --> grpcC1 : use
    p2gw --> grpcC2 : use
}

node "Plugin1" <<process>> {
    rectangle "Implementation" <<grpc server>> as impl1
}

node "Plugin2" <<process>> {
    rectangle "Implementation" <<grpc server>> as impl2
}

impl1 -up- grpcI1A
impl1 -up- grpcI1B
impl2 -up- grpcI2

```

### Discovery

The Quorum client reads the plugin [settings](../Settings) file to determine which plugins are going to be loaded and searches for installed plugins
(`<name>-<version>.zip` files) in the plugin `baseDir` (defaults to `<datadir>/plugins`). If the required plugin doesnt exist in the path, Quorum will attempt to use the configured `plugin central` to download the plugin.

### PluginManager

The `PluginManager` manages the plugins being used inside `geth`. It reads the [configuration](../Settings) and builds a registry of plugins.
`PluginManager` implements the standard `Service` interface in `geth`, hence being embedded into the `geth` service life cycle, i.e.: expose service APIs, start and stop.
The `PluginManager` service is registered as early as possible in the node lifecycle. This is to ensure the node fails fast if an issue is encountered when registering the `PluginManager`, so as not to impact other services.

### Plugin Reloading

The `PluginManager` exposes an API (`admin_reloadPlugin`) that allows reloading a plugin. This attempts to restart the current plugin process.   

Any changes to the plugin config after initial node start will be applied when reloading the plugin.  
This is demonstrated in the [HelloWorld plugin example](http://localhost:8000/PluggableArchitecture/Overview/#example-helloworld-plugin).
