# Quorum

Quorum is a blockchain with properties that make it suitable as a consortium or private chain.
One of its core features is the ability to keep certain state private and only accessible by granted parties.

## Consensus algorithm

Quorum is a majority voting protocol where a subset of nodes within the network are given the `voting` role.
The voting role allows a node to vote which block should be the canonical head on a particular height.
The block with the most votes will win and is considered the canonical head of the chain.

Block creation is only allowed by nodes with the `block maker` role.
A node with this role can create a block, sign the block and put the signature within in `ExtraData` field of the block.
On block import nodes can verify if the block was signed by one of the nodes that have the `block maker` role.

Nodes can be given no role, one of the roles or both roles through command line arguments.
The collection of addresses with special roles is tracked within the Quorum smart contract.

Quorum is implemented in a smart contract pre-deployed on address `0x0000000000000000000000000000000000000020` and can be found [here](https://github.com/ethlab/go-ethereum-private/blob/master/core/quorum/block_voting.sol).
Voters and block makers can be added or removed and the minimum number of votes before a block is selected as winner can be configured.

## State

Quorum supports dual state:

- public state, accessible by all nodes within the network
- private state, only accessible by nodes with the correct permissions

The difference is made through the use of transactions with encrypted (private) and non-encrypted payload (public) transactions.
Nodes can determine if a transaction is private by looking at the V value of the signature.
Public transactions have a V value of 27 or 28, private transactions have a value of 37 or 38.

If the transaction is private and the node has the ability to decrypt the payload it can execute the transaction.
Nodes who are not involved in the transaction cannot decrypt the payload and process the transaction.
As a result all nodes share a common public state which is created through public transactions and have a local unique private state.

This model imposes a restriction in the ability to modify state in private transactions.
Since its a common use case that a (private) contract reads data from a public contract the virtual machine has the ability to jump into read only mode.
For each call from a private contract to a public contract the virtual machine will change to read only mode.
If the virtual machine is in read only mode and the code tries to make a stage change the virtual machine stops execution and throws an exception.

The following transactions are allowed:

S: sender, (X): private, X: public, ->: direction, []: read only mode
```
1. S -> A -> B
2. S -> (A) -> (B)
3. S -> (A) -> [B -> C]
```
The following transaction are unsupported:

```
1. (S) -> A
2. (S) -> (A)
```

### State verification

To determine if nodes are in sync the public state root hash is included in the block.
Since private transactions can only be processed by nodes that are involved its impossible to get global consensus on the private state.
To overcome this issue the RPC method `eth_storageRoot(address[, blockNumber]) -> hash` can be used.
It returns the storage root for the given address at an (optional) block number.
If the optional block number is not given the latest block number is used.
The storage root hash can be on or off chain compared by the parties involved.

## Building Quorum

Clone the repository and build the source:

```
git clone https://github.com/jpmorganchase/quorum.git
cd quorum
make all
```

Binaries are placed within `$REPO_ROOT/build/bin`.

Run the tests:

```
make test
```

## Running Quorum

Describing all command line arguments it out of the scope of this document. They can be viewed with: `geth --help`.

### Initialize chain

The first step is to generate the genesis block.

```
geth init genesis.json
```

The genesis block should include the Quorum voting contract address `0x0000000000000000000000000000000000000020`.
The code can be generated with [browser solidity](http://ethereum.github.io/browser-solidity/#version=soljson-latest.js) (note, use the runtime code) or using the solidity compiler `solc --optimize --bin-runtime block_voting.sol`.

The `7nodes` directory in the `quorum-examples` repository contains several keys (using an empty password) that are used in the example genesis file:
```
key1    vote key 1
key2    vote key 2
key3    vote key 3
key4    block maker 1
key5    block maker 2
```

Example genesis file:
```
{
  "alloc": {
    "0x0000000000000000000000000000000000000020": {
      "code": "606060405236156100c45760e060020a60003504631290948581146100c9578063284d163c146100f957806342169e4814610130578063488099a6146101395780634fe437d514610154578063559c390c1461015d57806368bb8bb61461025d57806372a571fc146102c857806386c1ff681461036957806398ba676d146103a0578063a7771ee31461040b578063adfaa72e14610433578063cf5289851461044e578063de8fa43114610457578063e814d1c71461046d578063f4ab9adf14610494575b610002565b610548600435600160a060020a03331660009081526003602052604090205460ff16156100c45760018190555b50565b610548600435600160a060020a03331660009081526005602052604090205460ff16156100c4576004546001141561055e57610002565b61045b60025481565b61054a60043560056020526000908152604090205460ff1681565b61045b60015481565b61045b60043560006000600060006000600050600186038154811015610002579080526002027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630192505b60018301548110156105d75760018301805484916000918490811015610002576000918252602080832090910154835282810193909352604091820181205485825292869052205410801561023257506001805490840180548591600091859081101561000257906000526020600020900160005054815260208101919091526040016000205410155b156102555760018301805482908110156100025760009182526020909120015491505b6001016101a8565b610548600435602435600160a060020a03331660009081526003602052604081205460ff16156100c4578054839010156105e45780548084038101808355908290829080158290116105df576002028160020283600052602060002091820191016105df919061066b565b610548600435600160a060020a03331660009081526005602052604090205460ff16156100c457600160a060020a0381166000908152604090205460ff1615156100f65760406000819020805460ff191660019081179091556004805490910190558051600160a060020a038316815290517f1a4ce6942f7aa91856332e618fc90159f13a340611a308f5d7327ba0707e56859181900360200190a16100f6565b610548600435600160a060020a03331660009081526003602052604090205460ff16156100c4576002546001141561071457610002565b61045b600435602435600060006000600050600185038154811015610002579080526002027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630181509050806001016000508381548110156100025750825250602090200154919050565b61054a600435600160a060020a03811660009081526003602052604090205460ff165b919050565b61054a60043560036020526000908152604090205460ff1681565b61045b60045481565b6000545b60408051918252519081900360200190f35b61054a600435600160a060020a03811660009081526005602052604090205460ff1661042e565b610548600435600160a060020a03331660009081526003602052604090205460ff16156100c457600160a060020a03811660009081526003602052604090205460ff1615156100f65760406000818120600160a060020a0384169182905260036020908152815460ff1916600190811790925560028054909201909155825191825291517f0ad2eca75347acd5160276fe4b5dad46987e4ff4af9e574195e3e9bc15d7e0ff929181900390910190a16100f6565b005b604080519115158252519081900360200190f35b600160a060020a03811660009081526005602052604090205460ff16156100f65760406000819020805460ff19169055600480546000190190558051600160a060020a038316815290517f8cee3054364d6799f1c8962580ad61273d9d38ca1ff26516bd1ad23c099a60229181900360200190a16100f6565b509392505050565b505050505b60008054600019850190811015610002578382526002027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563016020819052604082205490925014156106b8578060010160005080548060010182818154818355818115116106a5578183600052602060002091820191016106a5919061068d565b50506002015b808211156106a157600181018054600080835591825260208220610665918101905b808211156106a1576000815560010161068d565b5090565b5050506000928352506020909120018290555b600082815260208281526040918290208054600101905581514381529081018490528151600160a060020a033316927f3d03ba7f4b5227cdb385f2610906e5bcee147171603ec40005b30915ad20e258928290030190a2505050565b600160a060020a03811660009081526003602052604090205460ff16156100f65760406000819020805460ff19169055600280546000190190558051600160a060020a038316815290517f183393fc5cffbfc7d03d623966b85f76b9430f42d3aada2ac3f3deabc78899e89181900360200190a16100f656",
      "storage": {
        "0x0000000000000000000000000000000000000000000000000000000000000001": "0x02",

        "0x0000000000000000000000000000000000000000000000000000000000000002": "0x04",
        "0x29ecdbdf95c7f6ceec92d6150c697aa14abeb0f8595dd58d808842ea237d8494": "0x01",
        "0x6aa118c6537572d8b515a9f9154be55a3377a8de7991cd23bf6e5ceb368688e3": "0x01",
        "0x50793743212c6f01d326957d7069005b912f8215f10c7536be6b10782c6c44cd": "0x01",
        "0x38f6c908c5cc7ca668cec2f476abe61b4dbb1df20f0ad8e07ef5dbf6a2f1ffd4": "0x01",

        "0x0000000000000000000000000000000000000000000000000000000000000004": "0x02",
        "0xaca3b76ed4968740c3180dd7fa37f4aa229a2c758a848f53920e9ccb4c4bb74e": "0x01",
        "0xd188ba2dc293670542c1befaf7678b0859e5354a0727d1188b2afb6f47fe24d1": "0x01"
      }
    },
    "0xed9d02e382b34818e88b88a309c7fe71e65f419d": {
      "balance": "1000000000000000000000000000"
    },
    "0xca843569e3427144cead5e4d5999a3d0ccf92b8e": {
      "balance": "1000000000000000000000000000"
    },
    "0x0fbdc686b912d7722dc86510934589e0aaf3b55a": {
      "balance": "1000000000000000000000000000"
    },
    "0x9186eb3d20cbd1f5f992a950d808c4495153abd5": {
      "balance": "1000000000000000000000000000"
    },
    "0xed9d02e382b34818e88b88a309c7fe71e65f419d": {
      "balance": "1000000000000000000000000000"
    },
    "0x0638e1574728b6d862dd5d3a3e0942c3be47d996": {
      "balance": "1000000000000000000000000000"
    }
  },
  "coinbase": "0x0000000000000000000000000000000000000000",
  "config": {
    "homesteadBlock": 0
  },
  "difficulty": "0x0",
  "extraData": "0x",
  "gasLimit": "0x2FEFD800",
  "mixhash": "0x00000000000000000000000000000000000000647572616c65787365646c6578",
  "nonce": "0x0",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp": "0x00"
}
```

The storage key for voters and block makers is calculated with `web3.sha3(<256 bit aligned key value> + <256 bit variable index>)`.
The console can be used to calculate the storage key, in this case for vote key 1:
```
> key = "000000000000000000000000ed9d02e382b34818e88b88a309c7fe71e65f419d" + "0000000000000000000000000000000000000000000000000000000000000003"
"000000000000000000000000ed9d02e382b34818e88b88a309c7fe71e65f419d0000000000000000000000000000000000000000000000000000000000000003"
> web3.sha3(key, {"encoding": "hex"})
"0x29ecdbdf95c7f6ceec92d6150c697aa14abeb0f8595dd58d808842ea237d8494"
```

The `genesis.json` file can be found in the `7nodes` folder in the `quorum-examples` repository.

### Start node

Starting a node is as simple as `geth`. This will start the node without any of the roles and makes the node a spectator.

### Voting role

Start a node with the voting role:

```
geth --voteaccount 0xed9d02e382b34818e88b88a309c7fe71e65f419d
```

Optionally the `--votepassword` can be used to unlock the account.
If this flag is omitted the node will prompt for the password.

### Block maker role

Start a node with the block maker role:
```
geth --blockmakeraccount 0x9186eb3d20cbd1f5f992a950d808c4495153abd5
```

Created blocks will be signed with this account.

Optionally the `--blockmakerpassword` can be used to unlock the account.
If this flag is omitted the node will prompt for the password.

## Setup multi-node network

Quorum comes with several scripts to setup a private test network with 7 nodes:

* node 1, has no special roles
* node 2, has the block maker role
* node 3, has no special roles
* node 4, has the voting role
* node 5, has the voting role
* node 6, has no special roles

All scripts can be found in the `7nodes` folder in the `quorum-examples` repository.

1. Step 1, run `init.sh` and initialize data directories (change variables accordingly)
2. Step 2, start nodes with `start.sh` (change variables accordingly)
3. Step 3, stop network with `stop.sh`

### API

Quorum provides an API to inspect the current state of the voting contract.
 
$ quorum.nodeInfo returns the quorum capabilities of this node.
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

$ quorum.vote accepts a block hash and votes for this hash to be the canonical head on the current height. It returns the tx hash.
```
> quorum.vote(eth.getBlock("latest").hash)
"0x16c69b9bdf9f10c64e65dbfe50bc997d2bc1ed321c6041db602908b7f6cab2a9"
```

$ quorum.canonicalHash accepts a block height and returns the canonical hash for that height (+1 will return the hash where the current pending block will be based on top of).
```
> quorum.canonicalHash(eth.blockNumber+1)
"0xf2c8a36d0c54c7013246fddebfc29bc881f6f10f74f761d511b5ebfaa103adfa"
```

$ quorum.isVoter accepts an address and returns an indication if the given address is allowed to vote for new blocks
```
> quorum.isVoter("0xed9d02e382b34818e88b88a309c7fe71e65f419d")
true
```

$ quorum.isBlockMaker accepts an address and returns an indication if the given address is allowed to make blocks
```
> quorum.isBlockMaker("0xed9d02e382b34818e88b88a309c7fe71e65f419d")
true
```

$ quorum.makeBlock() orders the node to create a block bypassing block maker strategy.
```
> quorum.makeBlock()
"0x3a07e82a48ab3c19a3d09d247e189e3a3041d1d9eafd2e1515b4ddd5b016bfd9"
```

$ quorum.pauseBlockMaker (temporary) orders the node to stop creating blocks
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

$ quorum.resumeBlockMaker instructs the node stop begin creating blocks again when its paused.
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

## Sending Private Transactions

To send a private transaction, a PrivateTransactionManager must be configured. This is the
service which transfers private payloads to their intended recipients, performing
encryption and related operations in the process.

Currently, `constellation` is supported out of the box via the PRIVATE_CONFIG environment
variable (please note that this integration method will change in the near future.) See the
`7nodes` folder in the `quorum-examples` repository for a complete example of how to use it.
The transaction sent in `script1.js` is private for node 7's PrivateTransactionManager
public key.

Once `constellation` is launched and PRIVATE_CONFIG points to a valid configuration file,
a `SendTransaction` call can be made private by specifying the `privateFor` argument.
`privateFor` is a list of public keys of the intended recipients. (Note that in the case of
`constellation`, this public key is distinct from Ethereum account keys.) When a transaction
is private, the transaction contents will be sent to the PrivateTransactionManager and the
identifier returned will be placed in the transaction instead. When other Quorum nodes
receive a private transaction, they will query their PrivateTransactionManager for the
identifier and replace the transaction contents with the result (if any; nodes which are
not party to a transaction will not be able to retrieve the original contents.)

## Command line flags
```
QUORUM OPTIONS:
  --voteaccount value		    Address that is used to vote for blocks
  --votepassword value		    Password to unlock the voting address
  --blockmakeraccount value	    Address that is used to create blocks
  --blockmakerpassword value	Password to unlock the block maker address
  --singleblockmaker		    Indicate this node is the only node that can create blocks
  --minblocktime value		    Set minimum block time (default: 3)
  --maxblocktime value		    Set max block time (default: 10)
```

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.
