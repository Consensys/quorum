# account Plugins

!!! info
    `account` plugins are currently in beta

`account` plugins can be used with Quorum or `clef` to provide additional account management.  

It is recommended to first read the [Pluggable Architecture overview](../../../../PluggableArchitecture/Overview) to learn how to use plugins.

## Usage
### Using with Quorum & clef

```shell tab="Quorum"
geth --plugins file:///path/to/plugins.json ...
```

```shell tab="clef"
clef --plugins file:///path/to/plugins.json ...
```

where the [plugins settings file](../../../../PluggableArchitecture/Settings), `plugins.json`, defines an `account` provider: 

```json
{
    "providers": {
        "account": {
            "name": "quorum-account-plugin-<NAME>",
            "version": "<VERSION>",
            "config": "file:///path/to/plugin.json"
        }
    }
}
```

See [Available plugins](#available-plugins) for a list of available plugins and their corresponding `<NAME>` and `<VERSION>`.

### RPC API
A limited API allows users to interact directly with `account` plugins:

!!! info 
    Quorum must explicitly expose the API with `--rpcapi plugin@account` or `--wsapi plugin@account`

#### plugin@account_newAccount
Create a new plugin-managed account:

| Parameter | Description |
| --- | --- |
| `config` | Plugin-specific json configuration for creating a new account.  See the plugin's documentation for more info on the json config required 

##### Example
```shell tab="quorum"
curl -X POST \
     -H "Content-Type:application/json" \
     -d '
        {
            "jsonrpc":"2.0",
            "method":"plugin@account_newAccount",
            "params":[{<config>}], 
            "id":1
        }' \
     http://localhost:22000
``` 

```js tab="js console"
plugin_account.newAccount({<config>})
``` 

```shell tab="clef"
echo '
    {
        "jsonrpc":"2.0",
        "method":"plugin@account_newAccount",
        "params":[{<config>}], 
        "id":1
    }
' | nc -U /path/to/clef.ipc
```

#### plugin@account_importRawKey
Create a new plugin-managed account from an existing private key:  

!!! note 
    Although this API can be used to move plugin-managed accounts between nodes, the plugin may provide a more preferable alternative.  See the plugin's documentation for more info.

| Parameter | Description |
| --- | --- |
| `rawkey` | Hex-encoded account private key (without 0x prefix) 
| `config` | Plugin-specific json configuration for creating a new account.  See the plugin's documentation for more info on the json config required

##### Example
```shell tab="quorum"
curl -X POST \
     -H "Content-Type:application/json" \
     -d '
         {
             "jsonrpc":"2.0",
             "method":"plugin@account_importRawKey",
             "params":["<rawkey>", {<config>}], 
             "id":1
         }' \
     http://localhost:22000
```

```js tab="js console"
plugin_account.importRawKey(<rawkey>, {<config>})
``` 

```text tab="clef"
not supported - use CLI instead
```


### CLI
A limited CLI allows users to interact directly with `account` plugins:

```shell
geth account plugin --help
```
!!! info
    Use the `--verbosity` flag to hide log output, e.g. `geth --verbosity 1 account plugin new ...`

#### geth account plugin new
Create a new plugin-managed account: 

| Parameter | Description |
| --- | --- |
| <span style="white-space:nowrap">`plugins.account.config`</span> | Plugin-specific configuration for creating a new account.  Can be `file://` or inline-json. See the plugin's documentation for more info on the json config required

```shell tab="json file"
geth account plugin new \
    --plugins file:///path/to/plugin-config.json \
    --plugins.account.config file:///path/to/new-acct-config.json
```

```shell tab="inline json"
geth account plugin new \
    --plugins file:///path/to/plugin-config.json \
    --plugins.account.config '{<json>}'
```

#### geth account plugin import
Create a new plugin-managed account from an existing private key:  

| Parameter | Description |
| --- | --- |
| <span style="white-space:nowrap">`plugins.account.config`</span> | Plugin-specific configuration for creating a new account.  Can be `file://` or inline-json. See the plugin's documentation for more info on the json config required
| `rawkey` | Path to file containing hex-encoded account private key (without 0x prefix) (e.g. `/path/to/raw.key`)

```shell tab="json file"
geth account plugin import \
    --plugins file:///path/to/plugin-config.json \
    --plugins.account.config file:///path/to/new-acct-config.json \
    /path/to/raw.key
```

```shell tab="inline json"
geth account plugin import \
    --plugins file:///path/to/plugin-config.json \
    --plugins.account.config '{<json>}'
    /path/to/raw.key
```

#### geth account plugin list
List all plugin-managed accounts for a given config:

```shell
geth account plugin list \
    --plugins file:///path/to/plugin-config.json
```

## Available plugins 
| Name | Version |  | Description |
| --- | --- | --- | --- |
| <span style="white-space:nowrap">`hashicorp-vault`</span> | `0.0.1` | <span style="white-space:nowrap">[Docs](../Hashicorp-Vault/Overview) / [Source](https://www.github.com/jpmorganchase/quorum-account-plugin-hashicorp-vault)</span> | Enables storage of Quorum account keys in a Hashicorp Vault kv v2 engine.  Written in Go. 

## Developers
See [For Developers](../../../../PluggableArchitecture/Plugins/account/For-Developers). 
