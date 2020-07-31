!!! info
    `account` plugins are currently in beta

This quickstart guide will demonstrate how to:

1. Start a `vault` dev server and configure a basic approle access control
1. Configure Quorum to use the `vault` for account management by using the Hashicorp Vault `account` plugin
1. Create new accounts in `vault`
1. Use an account to sign some data

!!! warning 
    This quickstart uses the `vault` dev server.  The dev server is quick and easy to set up but should not be used for production.
    
    The dev server does **not**:
    
    * persist data between restarts
    * encrypt HTTP communications with TLS  
    
    For more advanced Vault topics (such as configuring storage, TLS, and approle token renewal) see the [Vault docs](https://www.vaultproject.io/docs).

### Setting up a vault dev server

1. [Download `vault`](https://www.vaultproject.io/downloads)
1. Start the dev server:
    ```shell
    $ vault server -dev
    WARNING! dev mode is enabled! In this mode, Vault runs entirely in-memory
    and starts unsealed with a single unseal key. The root token is already
    authenticated to the CLI, so you can immediately begin using Vault.
    
    You may need to set the following environment variable:
    
        $ export VAULT_ADDR='http://127.0.0.1:8200'
    
    The unseal key and root token are displayed below in case you want to
    seal/unseal the Vault or re-authenticate.
    
    Unseal Key: 89TmuvV1EWRWSLCPXf7Ei7XfQ4MMfEs8DQ1pUKZz6J4=
    Root Token: s.btAGUcTLteyfQuriH840JIzG
    
    Development mode should NOT be used in production installations!
    ```
1. Setup the `vault` CLI using the values printed when starting the dev server:
    ```shell
    export VAULT_ADDR=http://127.0.0.1:8200
    export VAULT_TOKEN=s.btAGUcTLteyfQuriH840JIzG
    ```
1. The Hashicorp Vault `account` plugin requires the Vault server to be configured with a kv v2 secret engine.  The dev server has one pre-configured (with name `secret`):
    ```shell
    $ vault secrets list -detailed
    Path      Plugin   Options          Description               
    ----      ------   -------          -----------               
     ...       ...       ...                ...
    secret/   kv       map[version:2]   key/value secret storage  
    ```
    
    !!! note 
        Older versions of Vault come configured with a kv v1 secret engine by default. If the value of `Options` is not `map[version:2]`, upgrade to version 2 with `vault kv enable-versioning secret/`
        
1. Setup basic access control using approle:
  
    1. Enable approle authentication:
        ```shell
        $ vault auth enable approle
        Success! Enabled approle auth method at: approle/ 
        ```
    1. Create a basic access policy for all secrets in the `secret` kv engine:
        ```shell
        $ cat <<EOF >policy.hcl
        path "secret/*" {
            capabilities = ["create", "update", "read"]
        }
        EOF
        ```
        ```shell
        $ vault policy write basicpolicy policy.hcl
        Success! Uploaded policy: basicpolicy
        ```
    1. Create a new approle role with the access defined by `basicpolicy`:
        ```shell
        $ vault write auth/approle/role/basicrole \
            token_policies="basicpolicy" \
            secret_id_ttl=0 \
            token_ttl=0
        Success! Data written to: auth/approle/role/basicrole
        ```
        ```shell
        $ vault read auth/approle/role/basicrole/role-id
        Key        Value
        ---        -----
        role_id    dab8be4b-f17d-a9fd-f124-19fbec9bb688
        ```
        ```shell
        $ vault write -f auth/approle/role/basicrole/secret-id
        Key                   Value
        ---                   -----
        secret_id             5442577d-22e2-423f-6e11-0980af45a4fc
        secret_id_accessor    93f3fdc5-e878-4955-e557-d7a25afef6ef
        ```
        
        !!! warning
            The `secret_id` and `secret_id_accessor` should **never be shared**
    
    1. Make a note of the `role_id` and `secret_id`.  Quorum/the plugin will need these to authenticate with Vault

### Using the plugin
1. Create [Quorum's plugin config](../../../../../../PluggableArchitecture/Settings), `quorum.json`:
    ```shell
    $ cat <<EOF >quorum.json
    {
        "baseDir": "plugins-basedir",
        "providers": {
            "account": {
                "name": "quorum-account-plugin-hashicorp-vault",
                "version": "0.0.1",
                "config": "file:///path/to/plugin.json"
            }
        }
    }
    EOF
    ``` 
1. Create the [Hashicorp Vault plugin's config](../Overview#configuration), `plugin.json`:
    ```shell
    $ cat <<EOF >plugin.json 
    {
        "vault": "http://localhost:8200",
        "kvEngineName": "secret",
        "accountDirectory": "file:///path/to/accts",
        "authentication": {
            "roleId": "env://HASHICORP_ROLE_ID",
            "secretId": "env://HASHICORP_SECRET_ID",
            "approlePath": "approle"
        }
    }
    EOF
    ```
1. Use the CLI to create a new account and store it in Hashicorp Vault 
    ```shell
    $ export HASHICORP_ROLE_ID=dab8be4b-f17d-a9fd-f124-19fbec9bb688 
    $ export HASHICORP_SECRET_ID=5442577d-22e2-423f-6e11-0980af45a4fc   
    ```
    ```shell  
    $ geth account plugin new \
          --plugins file:///path/to/quorum.json \
          --plugins.skipverify \
          --plugins.account.config '{"secretName": "demoacct","overwriteProtection": {"currentVersion": 0}}'
    Your new plugin-backed account was generated
    
    Public address of the account:   0x88133AcAf18Fb9db5A79066e0dB5208cd9491Cc9
    Account URL: localhost:8200/v1/secret/data/demoacct?version=1
    
    - You can share your public address with anyone. Others need it to interact with you.
    - You must NEVER share the secret key with anyone! The key controls access to your funds!
    - Consider BACKING UP your account! The specifics of backing up will depend on the plugin backend being used.
    - The plugin backend may require you to REMEMBER part/all of the new account config to retrieve the key in the future!
      See the plugin specific documentation for more info.
    
    - See the documentation for the plugin being used for more info.
    ```
1. The `accountDirectory` will contain a new account file:
    ```shell
    $ ls accts
    UTC--2020-06-29T19-57-11.071220000Z--88133acaf18fb9db5a79066e0db5208cd9491cc9
    ```
    ```shell
    $ cat accts/UTC--2020-06-29T19-57-11.071220000Z--88133acaf18fb9db5a79066e0db5208cd9491cc9
    {"Address":"88133acaf18fb9db5a79066e0db5208cd9491cc9","VaultAccount":{"SecretName":"demoacct","SecretVersion":1},"Version":1}
    ```
1. Vault will contain the new account address and private key:
    ```shell
    $ vault kv get secret/demoacct
    ====== Metadata ======
    Key              Value
    ---              -----
    created_time     2020-06-29T20:01:02.572383Z
    deletion_time    n/a
    destroyed        false
    version          1
    
    ====================== Data ======================
    Key                                         Value
    ---                                         -----
    88133acaf18fb9db5a79066e0db5208cd9491cc9    915b18a038546d28350fb88d686f40b99e482f4264d1acadd53c734c71488643
    ```

1. Start Quorum with the Hashicorp Vault plugin:
    ```shell
    $ PRIVATE_CONFIG=ignore geth \
        --nodiscover \
        --verbosity 5 \
        --networkid 10 \
        --raft \
        --raftjoinexisting 1 \
        --datadir datadir
        --rpc \
        --rpcapi eth,debug,admin,net,web3,plugin@account \
        --plugins file:///path/to/quorum.json \
        --plugins.skipverify 
    ```
1. Attach to the node, create another account using the RPC API (available on the `geth` console), and use the account
    ```shell
    $ geth attach datadir/geth.ipc
    Welcome to the Geth JavaScript console!
    
    instance: Geth/v1.9.7-stable-f343ba05(quorum-v2.6.0)/darwin-amd64/go1.13.7
    coinbase: 0x88133acaf18fb9db5a79066e0db5208cd9491cc9
    at block: 0 (Thu, 01 Jan 1970 01:00:00 BST)
     datadir: /Users/chrishounsom/Desktop/vault-plugin-demo/datadir
     modules: admin:1.0 debug:1.0 eth:1.0 ethash:1.0 miner:1.0 net:1.0 personal:1.0 plugin@account:1.0 raft:1.0 rpc:1.0 txpool:1.0 web3:1.0
    
    > plugin_account.newAccount({"secretName": "anotherdemoacct","overwriteProtection": {"currentVersion": 0}})
    {
        address: "0x8effa64323cd1d737f068c09021dccdee1f2ce6d",
        url: "http://localhost:8200/v1/secret/data/anotherdemoacct?version=1"
    }
   
    > personal.listWallets
    [{
        accounts: [{
            address: "0x8effa64323cd1d737f068c09021dccdee1f2ce6d",
            url: "http://localhost:8200/v1/secret/data/anotherdemoacct?version=1"
        }, {
            address: "0x88133acaf18fb9db5a79066e0db5208cd9491cc9",
            url: "http://localhost:8200/v1/secret/data/demoacct?version=1"
        }],
        status: "0 unlocked account(s)",
        url: "plugin://account-plugin-hashicorp-vault"
    }]
   
    > personal.sign("0xaaaaaa", "0x88133acaf18fb9db5a79066e0db5208cd9491cc9", "")
    "0x051f8a957aa73f51dae58ad4fc387aae93b0901031114adac39caffc667547ca3746e0982744226a2ffdf763ed747f21aba36237fc4382cac72cf81d721303511b"
    ```
   