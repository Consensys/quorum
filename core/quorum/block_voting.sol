pragma solidity ^0.4.2;

// Implements a block voting algorithm to reach consensus.
//
// To vote for a block the sender must be allowed to vote. When deployed the
// deployer is the only party that is allowed to vote and can add new voters.
// Note that voters can add new voters and thus have the abbility to add multiple
// voter accounts that they control. This gives them the possibility to vote
// multiple times for a particular block. Therefore voters must be trusted.
contract BlockVoting {
    // Raised when a vote is made
    event Vote(address indexed sender, uint blockNumber, bytes32 blockHash);
    // Raised when a new address is allowed to vote.
    event AddVoter(address);
    // Raised when an address is not alloed to make votes anymore.
	event RemovedVoter(address);
	// Raised when a new address is allowed to create new blocks.
	event AddBlockMaker(address);
	// Raised when an address is not allowed to make blocks anymore.
	event RemovedBlockMaker(address);

    // The period in which voters can vote for a block that is selected
    // as the new head of the chain.
	struct Period {
	    // number of times a block is voted for
		mapping(bytes32 => uint) entries;

		// blocks up for voting
		bytes32[] indices;
	}

    // Collection of vote rounds.
	Period[] periods;

    // canonical hash must have as least voteThreshold votes before its considered valid
	uint public voteThreshold;

    // Number of addresses that are allowed to make votes.
    uint public voterCount;

    // Collection of addresses that are allowed to vote.
    mapping(address => bool) public canVote;

    // Number of addresses that are alloed to create blocks.
    uint public blockMakerCount;

    // Collection of addresses that are allowed to create blocks.
    mapping(address => bool) public canCreateBlocks;

    // Only allow addresses that are allowed to make votes.
	modifier mustBeVoter() {
		if (canVote[msg.sender]) {
		    _;
		} else {
		    throw;
		}
	}

	// Only allow addresses that are allowed to create blocks.
    modifier mustBeBlockMaker() {
        if (canCreateBlocks[msg.sender]) {
            _;
        } else {
            throw;
        }
    }

    // Set a new vote threshold. The canonical hash must have at least the given
    // threshold number of votes before it's considered valid.
	function setVoteThreshold(uint threshold) mustBeVoter {
	    voteThreshold = threshold;
	}

    // Make a vote to select a particular block as head for the previous head.
    // Only senders that are added through the addVoter are allowed to make a vote.
    // TODO: discuss if we only allow 1 vote per voter
    // (this can deadlock the system if all voters votes for something different
    // (nVotes < threshold) and cannot vote anymore, or gas limit is reached).
	function vote(uint height, bytes32 hash) mustBeVoter {
	    // start new period if this is the first transaction in the new block.
		if (periods.length < height) {
		    periods.length += height-periods.length;
		}

		// select the voting round.
		Period period = periods[height-1];

		// new block hash entry
		if(period.entries[hash] == 0) period.indices.push(hash);

		// vote
		period.entries[hash]++;

		// log vote
		Vote(msg.sender, block.number, hash);
	}

    // Get canonical head for a given block number.
    // E.g. [block 124] - [block 125] - [block 126 (pending)]
    // getCanonHash(126) will return the hash of block 125
    // (if there are enough votes for it).
	function getCanonHash(uint height) constant returns(bytes32) {
		Period period = periods[height-1];

		bytes32 best;
		for(uint i = 0; i < period.indices.length; i++) {
			if(period.entries[best] < period.entries[period.indices[i]]
			&& period.entries[period.indices[i]] >= voteThreshold) {
				best = period.indices[i];
			}
		}
		return best;
	}

	// Add an party that is allowed to make a vote.
	// Only current voters are allowed to add a new voter.
	function addVoter(address addr) mustBeVoter {
		if (!canVote[addr]) {
		    canVote[addr] = true;
		    voterCount++;
		    AddVoter(addr);
		}
	}

	// Remove a party that is allowed to vote.
	// Note, a voter can remove it self as a voter!
	function removeVoter(address addr) mustBeVoter {
	    // don't let the last voter remove it self
	    // which can cause the algorithm to stall.
	    if (voterCount == 1) throw;

        if (canVote[addr]) {
	        delete canVote[addr];
	        voterCount--;
	        RemovedVoter(addr);
        }
	}

	// isVoter returns an indication if the given address is allowed to vote.
	function isVoter(address addr) constant returns (bool) {
	    return canVote[addr];
	}

    // addBlockMaker adds the given list to the collection of addresses that
    // are allowed to create blocks.
	function addBlockMaker(address addr) mustBeBlockMaker {
        if (!canCreateBlocks[addr]) {
            canCreateBlocks[addr] = true;
            blockMakerCount++;
            AddBlockMaker(addr);
        }
	}

	// removeBlocksMaker deletes the given address of the collection of
	// addresses that are allowed to create blocks.
	function removeBlockMaker(address addr) mustBeBlockMaker {
	    if (blockMakerCount == 1) throw;

	    if (canCreateBlocks[addr]) {
	        delete canCreateBlocks[addr];
	        blockMakerCount--;
	        RemovedBlockMaker(addr);
	    }
	}

    // isBlockMaker returns an indication if the given address can create blocks.
	function isBlockMaker(address addr) constant returns (bool) {
	    return canCreateBlocks[addr];
	}

    // Number of voting rounds.
	function getSize() constant returns(uint) {
		return periods.length;
	}

    // Return a blockhash by period and index.
	function getEntry(uint height, uint n) constant returns(bytes32) {
		Period period = periods[height-1];
		return period.indices[n];
	}
}