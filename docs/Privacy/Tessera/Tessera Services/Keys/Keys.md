## Generating keys
 
Key generation can be used in multiple ways:

1. Generate a key pair and save in new files `.pub` and `.key`:  
    ```
    tessera -keygen
    ```
    This command will require interactive input for passwords. 
If you wish to generate an unlocked key, `/dev/null` can be used for stdin to tell the application not to expect any input (version 0.8 only):
    ```
    # Version 0.8+
    tessera -keygen < /dev/null

    # Version 0.7.x or before
    printf "\n\n" | tessera -keygen
    ```

    The `-filename` option can be used to specify alternate filepaths.  Multiple key pairs can be generated at the same time by providing a comma-separated list of values:
    ```
    tessera -keygen -filename /path/to/key1,/path/to/key2
    ```

1. Generate a key pair and save to an Azure Key Vault, with DNS name `<url>`, as secrets with IDs `Pub` and `Key`:
    ```
    tessera -keygen -keygenvaulttype AZURE -keygenvaulturl <url>
    ```
    
    The `-filename` option can be used to specify alternate IDs.  Multiple key pairs can be generated at the same time by providing a comma-separated list of values:
    ```
    tessera -keygen -keygenvaulttype AZURE -keygenvaulturl <url> -filename id1,id2
    ```
    
    **Note: If saving new keys with the same ID as keys that already exist in the vault, the existing keys will be replaced by the newer version.**
    
    > Environment variables must be set if using an Azure Key Vault, for more information see [Setting up an Azure key vault](../Setting%20up%20an%20Azure%20Key%20Vault)
    
1. Generate a key pair and save to a Hashicorp Vault at the secret path `secretEngine/secretName` with IDs `publicKey` and `privateKey`:
    ```bash
    tessera -keygen -keygenvaulttype HASHICORP -keygenvaulturl <url> \
       -keygenvaultsecretengine secretEngine -filename secretName 
    ```
    Options exist for configuring TLS and AppRole authentication (by default the AppRole path is set to `approle`):
    ```bash
    tessera -keygen -keygenvaulttype HASHICORP -keygenvaulturl <url> \
       -keygenvaultsecretengine <secretEngineName> -filename <secretName> \
       -keygenvaultkeystore <JKS file> -keygenvaulttruststore <JKS file> \
       -keygenvaultapprole <authpath>
    ```
    The `-filename` option can be used to generate and store multiple key pairs at the same time:
    ```bash
    tessera -keygen -keygenvaulttype HASHICORP -keygenvaulturl <url> \
       -keygenvaultsecretengine secretEngine -filename myNode/keypairA,myNode/keypairB 
    ```
    **Saving a new key pair to an existing secret will overwrite the values stored at that secret.  Previous versions of secrets may be retained and be retrievable by Tessera depending on how the K/V secrets engine is configured.  See [Keys](../../../Configuration/Keys) for more information on configuring Tessera for use with Vault.**
    
    > Environment variables must be set if using a Hashicorp Vault, and a version 2 K/V secret engine must be enabled.  For more information see [Setting up a Hashicorp Vault](../Setting%20up%20a%20Hashicorp%20Vault).

1. Generate a key pair, save to files and then start Tessera using a provided config
    ```
    tessera -keygen -configfile /path/to/config.json
    ```
    ```
    tessera -keygen -filename key1 -configfile /path/to/config.json 
    ```
    Tessera loads `config.json` as usual and includes the newly generated key data before starting.  
    
    An updated `.json` configfile is printed to the terminal (or to a file if using the `-output` CLI option).  No changes are made to the `config.json` file itself.

## Securing private keys
Generated private keys can be encrypted with a password.  This is prompted for on the console during key generation.  After generating password-protected keys, the password must be added to your configuration to ensure Tessera can read the keys.  The password is not saved anywhere but must be added to the configuration else the key will not be able to be decrypted.  

Passwords can be added to the json config either inline using `"passwords":[]`, or stored in an external file that is referenced by `"passwordFile": "Path"`.  Note that the number of arguments/file-lines provided must equal the total number of private keys.  For example, if there are 3 total keys and the second is not password secured, the 2nd argument/line must be blank or contain dummy data.

Tessera uses Argon2 in the process of encrypting private keys.  By default, Argon2 is configured as follows:
```
{
    "variant": "id",
    "memory": 1048576,
    "iterations": 10,
    "parallelism": 4
}
```
The Argon2 configuration can be altered by using the `-keygenconfig` option.  Any override file must have the same format as the default configuration above and all options must be provided.
```
tessera -keygen -filename /path/to/key1 -keygenconfig /path/to/argonoptions.json
```

For more information on Argon2 see the [Argon2 Github page](https://github.com/P-H-C/phc-winner-argon2).

### Updating password protected private keys
The password of a private key stored in a file can be updated.  Password update uses the `--keys.keyData.privateKeyPath` CLI option to get the path to the file. 

Password update can be used in multiple ways.  Running any of these commands will start a CLI prompt to allow you to set a new password.

1. Add a password to an unlocked key
    ```
    tessera -updatepassword --keys.keyData.privateKeyPath /path/to/.key
    ```
    
1. Change the password of a locked key.  This requires providing the current password for the key (either inline or as a file)
    ```
    tessera -updatepassword --keys.keyData.privateKeyPath /path/to/.key --keys.passwords <password>
    ```
    or
    ```
    tessera -updatepassword --keys.keyData.privateKeyPath /path/to/.key --keys.passwordFile /path/to/pwds
    ```

1. Use different Argon2 options from the defaults when updating the password
    ```
    tessera --keys.keyData.privateKeyPath <path to keyfile> --keys.keyData.config.data.aopts.algorithm <algorithm> --keys.keyData.config.data.aopts.iterations <iterations> --keys.keyData.config.data.aopts.memory <memory> --keys.keyData.config.data.aopts.parallelism <parallelism>
    ```
    All options have been overriden here but only the options you wish to alter from their defaults need to be provided.
