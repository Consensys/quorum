# Istanbul RPC API
This is an up to date copy of original wiki entry located here https://github.com/getamis/go-ethereum/wiki/RPC-API


## Getting Started
1. Run Istanbul geth with `--rpcapi "istanbul"`
2. `geth attach`

## API Reference

### istanbul.candidates
Candidates returns the current candidates which the node tries to vote in or out.
```
istanbul.candidates
```

#### Returns
`map[string] boolean` - returns the current candidates map.

### istanbul.discard
Discard drops a currently running candidate, stopping the validator from casting further votes (either for or against).
```
istanbul.discard(address)
```

#### Parameters
`string` - the address of the candidate

### istanbul.getSnapshot
GetSnapshot retrieves the state snapshot at a given block.
```
istanbul.getSnapshot(blockHashOrBlockNumber)
```

#### Parameters
`String|Number` - The block number, the string "latest" or nil. nil is the same with string "latest" and means the latest block

#### Returns
`Object` - The snapshot object

### istanbul.getSnapshotAtHash
GetSnapshotAtHash retrieves the state snapshot at a given block.
```
istanbul.getSnapshotAtHash(blockHash)
```

#### Parameters
`String` - The block hash

#### Returns
`Object` - The snapshot object

### istanbul.getValidators
GetValidators retrieves the list of authorized validators at the specified block.
```
istanbul.getValidators(blockHashOrBlockNumber)
```

#### Parameters
`String|Number` - The block number, the string "latest" or nil. nil is the same with string "latest" and means the latest block

#### Returns
`[]string` - The validator address array

### istanbul.getValidatorsAtHash
GetValidatorsAtHash retrieves the list of authorized validators at the specified block.
```
istanbul.getValidatorsAtHash(blockHash)
```

#### Parameters
`String` - The block hash

#### Returns
`[]string` - The validator address array

### istanbul.propose
Propose injects a new authorization candidate that the validator will attempt to push through. If the number of vote is larger than 1/2 of validators to vote in/out, the candidate will be added/removed in validator set.

```
istanbul.propose(address, auth)
```

#### Parameters
`String` - The address of candidate
`bool` - `true` votes in and `false` votes out

### istanbul.nodeAddress
Retrieves the public address that is used to sign proposals, which is derived from the nodes `nodekey`.
```
istanbul.nodeAddress()
```

#### Returns
`string` - The nodes public signing address

### istanbul.getSignersFromBlock
Retrieves the public addresses for whose seals are included in the block. This means that they participated in the
consensus for this block and attested to its validity.
A block number may be optionally given, or else the current block is assumed.
```
istanbul.getSignersFromBlock(blockNumber)
```

#### Parameters
`Number` - The block number to retrieve

#### Returns
`Object` -
    - `number`: `Number` - The retrieved block's number
    - `hash`: `String` - The retrieved block's hash
    - `author`: `String` - The address of the block proposer
    - `committers`: `[]String` - The list of all addresses whose seal appears in this block

### istanbul.getSignersFromBlockByHash
Retrieves the public addresses for whose seals are included in the block. This means that they participated in the
consensus for this block and attested to its validity. A block hash must be given, and does NOT default to the current
latest block.
```
istanbul.getSignersFromBlockByHash(blockHash)
```

#### Parameters
`String` - The hash of the block to retrieve

#### Returns
`Object` -
    - `number`: `Number` - The retrieved block's number
    - `hash`: `String` - The retrieved block's hash
    - `author`: `String` - The address of the block proposer
    - `committers`: `[]String` - The list of all addresses whose seal appears in this block