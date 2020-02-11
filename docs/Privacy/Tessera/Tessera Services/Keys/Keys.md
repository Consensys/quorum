## Generating keys

### File-stored keys
Generate a key pair and save in new files `new.pub` and `new.key` (will start an interactive prompt to provide passwords):   
```
tessera -keygen -filename new
```
Multiple key pairs can be generated at the same time by providing a comma-separated list of values:
```
tessera -keygen -filename /path/to/key1,/path/to/key2
```

To generate an unlocked key, the following can be used to tell Tessera to not expect any input:

```bash tab="v0.8.x onwards"
tessera -keygen < /dev/null
```

```bash tab="v0.7.x and earlier"
printf "\n\n" | tessera -keygen
```

### Azure Key Vault-stored keys
Generate a key pair as secrets with IDs `Pub` and `Key` and save to an Azure Key Vault with DNS name `<url>`:
```
tessera -keygen -keygenvaulttype AZURE -keygenvaulturl <url>
```

The `-filename` option can be used to specify alternate IDs.  Multiple key pairs can be generated at the same time by providing a comma-separated list of values:
```
tessera -keygen -keygenvaulttype AZURE -keygenvaulturl <url> -filename id1,id2
```

!!! warning
    If saving new keys with the same ID as keys that already exist in the vault, the existing keys will be replaced by the newer version.  When doing this, make sure to [specify the correct secret version in your Tessera configuration](../../../Configuration/Keys/#azure-key-vault-key-pairs) 

!!! note
    Environment variables must be set if using an Azure Key Vault, for more information see [Setting up an Azure key vault](../Setting%20up%20an%20Azure%20Key%20Vault)

### Hashicorp Vault-stored keys
Generate a key pair and save to a Hashicorp Vault at the secret path `secretEngine/secretName` with IDs `publicKey` and `privateKey`:
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

!!! warning 
    Saving a new key pair to an existing secret will overwrite the values stored at that secret.  Previous versions of secrets may be retained and be retrievable by Tessera depending on how the K/V secrets engine is configured.  When doing this, make sure to [specify the correct secret version in your Tessera configuration](../../../Configuration/Keys/#hashicorp-vault-key-pairs)

!!! note
    Environment variables must be set if using a Hashicorp Vault, and a version 2 K/V secret engine must be enabled.  For more information see [Setting up a Hashicorp Vault](../Setting%20up%20a%20Hashicorp%20Vault)
    
    
### AWS Secrets Manager-stored keys
Generate a key pair and save to an AWS Secrets Manager, with endpoint `<url>`, as secrets with IDs `Pub` and `Key`:

```bash
tessera -keygen -keygenvaulttype AWS -keygenvaulturl <url>
```

The `-filename` option can be used to specify alternate IDs.  Multiple key pairs can be generated at the same time by providing a comma-separated list of values:

```bash
tessera -keygen -keygenvaulttype AWS -keygenvaulturl <url> -filename id1,id2
```

!!! note
    Environment variables must be set if using an AWS Secrets Manager, for more information see [Setting up an AWS Secrets Manager](../Setting%20up%20an%20AWS%20Secrets%20Manager)

### Updating a configfile with newly generated keys 
Any newly generated keys must be added to a Tessera `.json` configfile.  Often it is easiest to do this manually.  

However, the `tessera keygen` `-configfile` option can be used to automatically update a configfile after key generation.  This is particularly useful when scripting.

```
tessera -keygen -filename key1 -configfile /path/to/config.json --configout /path/to/new.json --pwdout /path/to/new.pwds
```

The above command will prompt for a password and generate the `key1` pair as usual.  The Tessera configuration from `/path/to/config.json` will be read, updated and saved to `/path/to/new.json`.  New passwords will be appended to the existing password file as defined in `/path/to/config.json` and written to `/path/to/new.pwds`.

If the `--configout` and `--pwdout` options are not provided, the updated `.json` config will be printed to the terminal.

!!! note "Note: Differences between v0.10.3 and earlier versions"
    Before Tessera version 0.10.3 the node would start after updating the configfile.  
    
    In v0.10.3, this behaviour was removed to ensure clearer distinction of responsibilities between each Tessera command.  The same behaviour can be achieved in v0.10.3 onwards by running:
     ```
     tessera keygen ... -output /path/to/new.json
     tessera -configfile /path/to/new.json
     ```  

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

## Using alternative curve key types
By default the `-keygen` and `-updatepassword` commands generate and update [NaCl](https://nacl.cr.yp.to/) compatible keys.  

As of Tessera v0.10.2, the `--encryptor.type=EC` CLI option can be provided to generate/update keys of different types.  See [encryptor config](../../../Configuration/Configuration Overview/#encryptor-supporting-alternative-curves-in-tessera) for more details.
