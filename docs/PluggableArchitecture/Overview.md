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

**Note:** using `--plugins.skipverify`  is not advices for production settings and it should be avoided as introduces security risks.
