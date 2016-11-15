
# API

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
