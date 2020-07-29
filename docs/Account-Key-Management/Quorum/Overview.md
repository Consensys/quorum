# Quorum Account Management

As with geth, Quorum accounts can be stored in password-protected `keystore` files.  See [Keystore Files](../Keystore-Files) for more info.

Quorum v2.6.0 introduced `clef` (introduced in `geth` v1.9.x), a standalone account manager and signer that can be used to decouple account management responsibilities from the Quorum node.  `go-ethereum`'s intention is to deprecate account management within `geth` at some point in the future and replace it with `clef`.  See [Clef](../Clef) for more info.

Quorum v2.7.0 introduced the `account` plugins beta, which allows Quorum or `clef` to be extended with alternative methods of managing accounts.  See [account Plugins](../account-Plugins/Overview) for more info.
