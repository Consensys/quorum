
# API

## Privacy APIs

### `web3.eth.sendTransaction(object)` was modified to support private transactions

```js
web3.eth.sendTransaction(transactionObject [, callback])
```

Sends a transaction to the network.

##### Parameters

1. `Object` - The transaction object to send:
  - `from`: `String` - The address for the sending account. Uses the [web3.eth.defaultAccount](#web3ethdefaultaccount) property, if not specified.
  - `to`: `String` - (optional) The destination address of the message, left undefined for a contract-creation transaction.
  - `value`: `Number|String|BigNumber` - (optional) The value transferred for the transaction in Wei, also the endowment if it's a contract-creation transaction.
  - `gas`: `Number|String|BigNumber` - (optional, default: To-Be-Determined) The amount of gas to use for the transaction (unused gas is refunded).
  - <strike>`gasPrice`: `Number|String|BigNumber` - (optional, default: To-Be-Determined) The price of gas for this transaction in wei, defaults to the mean network gas price.</strike>
  - `data`: `String` - (optional) Either a [byte string](https://github.com/ethereum/wiki/wiki/Solidity,-Docs-and-ABI) containing the associated data of the message, or in the case of a contract-creation transaction, the initialisation code.
  - `nonce`: `Number`  - (optional) Integer of a nonce. This allows to overwrite your own pending transactions that use the same nonce.
  - `privateFrom`: `String`  - (optional) When sending a private transaction, the sending party's base64-encoded public key to use. If not present *and* passing `privateFor`, use the default key as configured in the `TransactionManager`.
  - `privateFor`: `List<String>`  - (optional) When sending a private transaction, an array of the recipients' base64-encoded public keys.
2. `Function` - (optional) If you pass a callback the HTTP request is made asynchronous. See [this note](#using-callbacks) for details.

##### Returns

`String` - The 32 Bytes transaction hash as HEX string.

If the transaction was a contract creation use [web3.eth.getTransactionReceipt()](#web3gettransactionreceipt) to get the contract address, after the transaction was mined.

##### Example

```js
// compiled solidity source code using https://chriseth.github.io/cpp-ethereum/
var code = "603d80600c6000396000f3007c01000000000000000000000000000000000000000000000000000000006000350463c6888fa18114602d57005b6007600435028060005260206000f3";

web3.eth.sendTransaction({
    data: code,
    privateFor: ["ROAZBWtSacxXQrOe3FGAqJDyJjFePR5ce4TSIzmJ0Bc="]
  },
  function(err, address) {
    if (!err) {
      console.log(address); // "0x7f9fade1c0d57a7af66ab4ead7c2eb7b11a91385"
    }
  }
});
```


## QuorumChain APIs

Quorum provides an API to inspect the current state of the voting contract.

### `quorum.nodeInfo` returns the quorum capabilities of this node

Example output for a node that is configured as block maker and voter:

```
> quorum.nodeInfo
{
  blockMakerAccount: "0xed9d02e382b34818e88b88a309c7fe71e65f419d",
  blockmakestrategy: {
    maxblocktime: 6,
    minblocktime: 3,
    status: "active",
    type: "deadline"
  },
  canCreateBlocks: true,
  canVote: true,
  voteAccount: "0xed9d02e382b34818e88b88a309c7fe71e65f419d"
}
```

### `quorum.vote` votes for the given hash to be the canonical head on the current height and returns the tx hash

```
> quorum.vote(eth.getBlock("latest").hash)
"0x16c69b9bdf9f10c64e65dbfe50bc997d2bc1ed321c6041db602908b7f6cab2a9"
```

### `quorum.canonicalHash` returns the canonical hash for the given block height (add 1 for the hash that the current pending block will be based on top of)

```
> quorum.canonicalHash(eth.blockNumber+1)
"0xf2c8a36d0c54c7013246fddebfc29bc881f6f10f74f761d511b5ebfaa103adfa"
```

### `quorum.isVoter` returns whether the given address is allowed to vote for new blocks

```
> quorum.isVoter("0xed9d02e382b34818e88b88a309c7fe71e65f419d")
true
```

### `quorum.isBlockMaker` returns whether the given address is allowed to make blocks

```
> quorum.isBlockMaker("0xed9d02e382b34818e88b88a309c7fe71e65f419d")
true
```

### `quorum.makeBlock` orders the node to create a block bypassing block maker strategy

```
> quorum.makeBlock()
"0x3a07e82a48ab3c19a3d09d247e189e3a3041d1d9eafd2e1515b4ddd5b016bfd9"
```

### `quorum.pauseBlockMaker` (temporary) orders the node to stop creating blocks

```
> quorum.pauseBlockMaker()
null
> quorum.nodeInfo
{
  blockMakerAccount: "0xed9d02e382b34818e88b88a309c7fe71e65f419d",
  blockmakestrategy: {
    maxblocktime: 6,
    minblocktime: 3,
    status: "paused",
    type: "deadline"
  },
  canCreateBlocks: true,
  canVote: true,
  voteAccount: "0xed9d02e382b34818e88b88a309c7fe71e65f419d"
}
```

### `quorum.resumeBlockMaker` instructs a paused node to begin creating blocks again

```
> quorum.resumeBlockMaker()
null
> quorum.nodeInfo
{
  blockMakerAccount: "0xed9d02e382b34818e88b88a309c7fe71e65f419d",
  blockmakestrategy: {
    maxblocktime: 6,
    minblocktime: 3,
    status: "active",
    type: "deadline"
  },
  canCreateBlocks: true,
  canVote: true,
  voteAccount: "0xed9d02e382b34818e88b88a309c7fe71e65f419d"
}
```
