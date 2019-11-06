title: Overview - Pluggable Architecture - Quorum

Quorum Client is a modified `geth` client. One of the unique enhancements
is the pluggable architecture which allows adding additional features as plugins to the core `geth`, 
providing extensibility, flexibility, and isolation of Quorum features.

## Why this?
 
This enhancement brings number benefits to the table:

1. Allowing selective components to decide their implementations at configuration time.
1. Supporting our community to make Quorum Client better with innovative implementations.
1. Decoupling new enterprise features from core `geth` therefore allowing us to update the code base in line with `go-ethereum` faster.

## How it works?

Each plugin exposes an implementation for a specific [plugin interface](../SupportedInterfaces).
Plugins are executed as a separate process and communicate with the main Quorum Client `geth` process
over an [gRPC](https://grpc.io/) interface.

The plugin implementation must adhere to certain gRPC services defined in a `.proto` file corresponding to the plugin interface.
Plugins can be written in different languages as gRPC provides a mechanism to generate stub code from `.proto` files. 

The network communication and RPC are handled automatically by high-level plugin library.

## Installing Plugins

Currently plugins must be manually installed into a directory (default to `plugins` directory inside `geth` data directory or via `baseDir` value in [plugins settings](../Settings)).
 
In the future, plugins distributed by Quorum are automatically installed by Quorum Client.

## Using Plugins

[Plugins settings file](../Settings) contains a JSON that describes what plugins to be used. 
Then start `geth` with `--plugins` as below:

```bash
geth ... \
     --plugins file:///<path>/<to>/plugins.json
```

## Plugin Integrity Verification
In its default settings Quorum uses its own Plugin Central Server to download and verify plugin integrity using [PGP](https://en.wikipedia.org/wiki/Pretty_Good_Privacy). 
However the architecture enables the same verification process locally via `--plugins.localverify` and `--plugins.publickey` flags or 
remotely with custom plugin central - reference the [`Settings`](../Settings/) section for more information on how to support custom plugin central. 

If the flag `--plugins.skipverify` is provided at runtime the plugin verification process will be disabled.

!!! warning
    Using `--plugins.skipverify`  is not adviced for production settings and it should be avoided as it introduces security risks.

## Example: `HelloWorld` plugin

In this example, `HelloWorld` plugin exposes a JSON RPC endpoint to return a greeting message in the configured language.
This plugin is reloadable which means we can use `admin_reloadPlugin` JSON RPC API to reload the plugin.

The `HelloWorld` plugin example is available in Quorum Git repository. In actual plugin development, plugin source code is maintained in a separate repository.

Prequisites to run this example are similar to those for building `geth` with `make`

1. Clone Quorum repository
   ```bash
   › git clone https://github.com/jpmorganchase/quorum.git
   › cd quorum
   ```
1. Build `geth` and the plugin
   ```bash
   quorum› make geth helloWorldPlugin
   ```
   Notice that there are files being created under `build/bin` directory:
    - `geth-plugin-settings.json` which is the plugin settings file for `geth`
    - `quorum-plugin-helloWorld-1.0.0.zip` which is the `HelloWorld` plugin distribution zip file
    - `helloWorld-plugin-config.json` which is the config file for the plugin
1. Run `geth` with plugin
   ```bash
   quorum› PRIVATE_CONFIG=ignore \
   build/bin/geth \
        --nodiscover \
        --verbosity 5 \
        --networkid 10 \
        --raft \
        --raftjoinexisting 1 \
        --datadir ./build/_workspace/test \
        --rpc \
        --rpcapi eth,debug,admin,net,web3,plugin@helloworld \
        --plugins file://./build/bin/geth-plugin-settings.json \
        --plugins.skipverify
   ```
   `ps -ef | grep helloWorld` would reveal the `HelloWorld` plugin process
1. Call the JSON RPC
   ```bash
   quorum› curl -X POST http://localhost:8545 \
        -H "Content-type: application/json" \
        --data '{"jsonrpc":"2.0","method":"plugin@helloworld_greeting","params":["Quorum Plugin"],"id":1}'
   {"jsonrpc":"2.0","id":1,"result":"Hello Quorum Plugin!"}
   ```
1. Update plugin config to support `es` language
   ```bash
   # update language to "es"
   quorum› vi build/bin/helloWorld-plugin-config.json
   ```
1. Reload the plugin
   ```bash
   quorum› curl -X POST http://localhost:8545 \
        -H "Content-type: application/json" \
        --data '{"jsonrpc":"2.0","method":"admin_reloadPlugin","params":["helloworld"],"id":1}'
   {"jsonrpc":"2.0","id":1,"result":true}
   ```
1. Call the JSON RPC
   ```bash
   quorum› curl -X POST http://localhost:8545 \
        -H "Content-type: application/json" \
        --data '{"jsonrpc":"2.0","method":"plugin@helloworld_greeting","params":["Quorum Plugin"],"id":1}'
   {"jsonrpc":"2.0","id":1,"result":"Hola Quorum Plugin!"}
   ```