title: Settings - Pluggable Architecture - Quorum

`geth` can load plugins from:

- JSON file which is passed via `--plugins` flag
- Ethereum TOML configuration file which is passed via `--config` flag

```json tab="JSON"
{
  "baseDir": string,
  "central": object(PluginCentralConfiguration),
  "providers": {
     <string>: object(PluginDefinition)
  }
}
```

```toml tab="TOML"
[Node.Plugins]
    BaseDir = string
    Central = object(PluginCentralConfiguration) # as inline table
    # Or as key-value table
    # [Node.Plugins.Central]
    # .. = .. from object(PluginCentralConfiguration)
    [[Node.Plugins.Providers]]
        <string> = object(PluginDefinition) # as inline table
        # Or as key-value table
        # [[Node.Plugins.Providers.<string>]]
        # .. = .. from object(PluginDefinition)
```

| Fields      | Description                                                                                                                                                                                                        |
|:------------|:-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `baseDir`   | A string indicates local directory from where plugins are read. If empty, default to `<datadir>/plugins`. <br/> To read from arbitrary enviroment variable (e.g: `MY_BASE_DIR`), provide value `env://MY_BASE_DIR` |
| `central`   | A configuration of the remote plugin central. See [PluginCentralConfiguration](#plugincentralconfiguration)                                                                                                        |
| `providers` | A map specifies supported plugin interfaces with the respected plugin provider definitions (see [PluginDefinition](#plugindefinition))                                                                             |
| `<string>`  | A string constant indicates the plugin interface. E.g: `helloworld`.                                                                                                                                               |

## `PluginCentralConfiguration`

Quorum Plugin Central Server will be used, modify this section to customize your own local plugin central

```json tab="JSON"
{
  "baseURL": string,
  "certFingerprint": string,
  "publicKeyURI": string,
  "insecureSkipVerify": bool
}
```

```toml tab="TOML"
# as inline table
{ BaseURL = string, CertFingerPrint = string, PublicKeyURI = string, InsecureSkipVerify = bool }
# as key-value table
BaseURL = string
CertFingerPrint = string
PublicKeyURI = string
InsecureSkipVerify = bool
```

| Fields               | Description                                                                                                          |
|:---------------------|:---------------------------------------------------------------------------------------------------------------------|
| `baseURL`            | A string indicating the remote plugin central URL (ex.`https://plugins.mycorp.com`)                                  |
| `certFingerprint`    | A string containing hex representation of the http server public key finger print to be used for certificate pinning |
| `publicKeyURI`       | A string that reference the location of the PGP public key to be used to perform the signature verification          |
| `insecureSkipVerify` | If true, verify the server's certificate chain and host name                                                         |

## `PluginDefinition`

Defines the plugin and its configuration

```json tab="JSON"
{
  "name": string,
  "version": string,
  "config": file/string/array/object
}
```

```toml tab="TOML"
# as inline table
{ Name = string, Version = string, Config = file/string/array/object }
# as key-value table
Name = string
Version = string
Config = file/string/array/object
```

| Fields    | Description                                                                                                                                                                                                                                                                     |
|:----------|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `name`    | A string specifies the name of the plugin                                                                                                                                                                                                                                       |
| `version` | A string specifies the version of the plugin                                                                                                                                                                                                                                    |
| `config`  | Value can be: <ul><li>uri format: support the following schemes<ul><li>`file`: location of plugin config file to be read. E.g.: `file:///opt/plugin.cfg`</li><li>`env`: value from an environment variable. E.g.: `env://MY_CONFIG_JSON`<br/>To indicate value is a file location: append `?type=file`. E.g.: `env://MY_CONFIG_FILE?type=file`</li></ul><li>string: an arbitrary JSON string</li><li>array: a valid JSON array E.g.: `["1", "2", "3"]`</li><li>object: a valid JSON object. E.g.: `{"foo" : "bar"}`</li></ul> |