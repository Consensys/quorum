
# Running Quorum

The following new CLI arguments were introduced as part of Quorum:

```
QUORUM OPTIONS:
  --voteaccount value         Address that is used to vote for blocks
  --votepassword value        Password to unlock the voting address
  --blockmakeraccount value   Address that is used to create blocks
  --blockmakerpassword value  Password to unlock the block maker address
  --singleblockmaker          Indicate this node is the only node that can create blocks
  --minblocktime value        Set minimum block time (default: 3)
  --maxblocktime value        Set max block time (default: 10)
  --permissioned              If enabled, the node will allow only a defined list of nodes to connect
```

The full list of arguments can be viewed by running `geth --help`.

### Initialize chain

The first step is to generate the genesis block.

The genesis block should include the Quorum voting contract address `0x0000000000000000000000000000000000000020`.
The code can be generated with [browser solidity](http://ethereum.github.io/browser-solidity/#version=soljson-latest.js) (note, use the runtime code) or using the solidity compiler: `solc --optimize --bin-runtime block_voting.sol`.

The `7nodes` directory in the `quorum-examples` repository contains several keys (using an empty password) that are used in the example genesis file:

```
key1    vote key 1
key2    vote key 2
key3    vote key 3
key4    block maker 1
key5    block maker 2
```

Example genesis file (copy to `genesis.json`):
```json
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

Now we can initialize geth:

```
geth init genesis.json
```

The storage key for voters and block makers is calculated with `web3.sha3(<256 bit aligned key value> + <256 bit variable index>)`.
The console can be used to calculate the storage key, in this case for vote key 1:
```
> key = "000000000000000000000000ed9d02e382b34818e88b88a309c7fe71e65f419d" + "0000000000000000000000000000000000000000000000000000000000000003"
"000000000000000000000000ed9d02e382b34818e88b88a309c7fe71e65f419d0000000000000000000000000000000000000000000000000000000000000003"
> web3.sha3(key, {"encoding": "hex"})
"0x29ecdbdf95c7f6ceec92d6150c697aa14abeb0f8595dd58d808842ea237d8494"
```
From the above example, the `<256 bit aligned key value>` is the ethereum account address that should be added to the voting map, ed9d02e382b34818e88b88a309c7fe71e65f419d, padded to 256bits. The `<256 bit variable index>` is the index(3) of the canVote mapping in the solidity [voting smart contract](https://github.com/jpmorganchase/quorum/blob/master/core/quorum/block_voting.sol#L42) padded to 256bits. The index is calculated based on the location of canVote:

* Period[] periods --> index 0
* uint public voteThreshold --> index 1
* uint public voterCount --> index 2
* mapping(address => bool) public canVote --> index 3

The `genesis.json` file can be found in the `7nodes` folder in the `quorum-examples` repository.

### Setup Bootnode
Optionally you can set up a bootnode that all the other nodes will first connect to in order to find other peers in the network. You will first need to generate a bootnode key: 

1-	To generate the key for the first time:

`bootnode –genkey tmp_file.txt  //this will start a bootnode with an enode address and generate a key inside a “tmp_file.txt” file`

2-	To later restart the bootnode using the same key (and hence use the same enode url):

`bootnode –nodekey tmp_file.txt`
              
or
                
`bootnode –nodekeyhex 77bd02ffa26e3fb8f324bda24ae588066f1873d95680104de5bc2db9e7b2e510 // Key from tmp_file.txt`


### Start node

Starting a node is as simple as `geth`. This will start the node without any of the roles and makes the node a spectator. If you have setup a bootnode then be sure to add the `--bootnodes` param to your startup command:

`geth --bootnodes $BOOTNODE_ENODE`

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

## Permissioned Network

Node Permissioning is a feature that controls which nodes can connect to a given node and also to which nodes this node can dial out to. Currently, it is managed at individual node level by the command line flag `--permissioned` while starting the node.

If the `--permissioned` node is present, the node looks for a file named `<data-dir>/permissioned-nodes.json`. This file contains the list of enodes that this node can connect to and also accepts connections only from those nodes. In other words, if permissioning is enabled, only the nodes that are listed in this file become part of the network. It is an error to enable `--permissioned` but not have the `permissioned-nodes.json` file. If the flag is given, but no nodes are present in this file, then this node can neither connect to any node or accept any incoming connections.

The `permissioned-nodes.json` follows following pattern (similar to `static-nodes.json`):

```json
[
  "enode://enodehash1@ip1:port1",
  "enode://enodehash2@ip2:port2",
  "enode://enodehash3@ip3:port3",
]
```

Sample file:

```json
[
  "enode://6598638ac5b15ee386210156a43f565fa8c48592489d3e66ac774eac759db9eb52866898cf0c5e597a1595d9e60e1a19c84f77df489324e2f3a967207c047470@127.0.0.1:30300",
]
```

In the current release, every node has its own copy of `permissioned-nodes.json`. In a future release, the permissioned nodes list will be moved to a smart contract, thereby keeping the list on chain and one global list of nodes that connect to the network.
