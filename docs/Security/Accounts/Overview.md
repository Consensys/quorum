# Managing accounts

!!! info
    External Vaults can also be used to manage Tessera key pairs.  See [Setting up Hashicorp Vault](http://localhost:8000/Privacy/Tessera/Tessera%20Services/Keys/Setting%20up%20a%20Hashicorp%20Vault/) and [Setting up Azure Key Vault]( http://localhost:8000/Privacy/Tessera/Tessera%20Services/Keys/Setting%20up%20an%20Azure%20Key%20Vault/) from the Tessera documentation for more info.

As with geth, Quorum accounts can be stored in password-protected `keystore` files (Ledger and Trezor hardware wallets are not yet supported).  See the [geth documentation](https://github.com/ethereum/go-ethereum/wiki/Managing-your-accounts) for details on using file-based accounts. 
 
In addition to this, Quorum also supports the storage of accounts in a Hashicorp Vault.  This section details how to [set up a Vault](Configuring-Hashicorp-Vault.md), [create new accounts in Vault](Creating-new-accounts-in-Vault.md), and how to [use accounts stored in Vault](Using-accounts-stored-in-Vault.md).

## Managing accounts in a Hashicorp Vault

Managing Quorum accounts in a Hashicorp Vault offers several benefits over using standard `keystore` files:

* Your account private keys are stored in a Hashicorp Vault which can be deployed on separate infrastructure to your Quorum node  

* Quorum can be configured to retrieve account private keys from the Vault to use in signing **only when needed**.  Keys are never written to disk by Quorum and are only held in memory indefinitely if configured

* Vault enables you to configure permissions on a per-secret basis to ensure account private keys can only be accessed by authorised users/applications 
