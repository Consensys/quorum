package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	istanbul "github.com/ethereum/go-ethereum/consensus/qibft"
)

// Returns true if the `proposal` is justified by the set `roundChangeMessages` of ROUND-CHANGE messages
// and by the set `prepareMessages` of PREPARE messages.
// For this we must either have:
//     - a quorum of ROUND-CHANGE messages with preparedRound and preparedBlockDigest equal to nil; or
//     - a ROUND-CHANGE message (1) whose preparedRound is not nil and is equal or higher than the
//           preparedRound of `quorumSize` ROUND-CHANGE messages and (2) whose preparedRound and
//           preparedBlockDigest match the round and block of `quorumSize` PREPARE messages.
func justify(proposal istanbul.Proposal, roundChangeMessages []*SignedRoundChangePayload, prepareMessages []*SignedPreparePayload, quorumSize int) bool {
	// Check the size of the set of ROUND-CHANGE messages
	if len(roundChangeMessages) < quorumSize {
		return false
	}

	// Check the size of the set of PREPARE messages
	if len(prepareMessages) != 0 && len(prepareMessages) < quorumSize {
		return false
	}

	// If there are PREPARE messages, they all need to have the same round and match `proposal`
	var preparedRound *big.Int
	iteration := 0
	for _, spp := range prepareMessages {
		if iteration == 0 { // Get the round of the first message
			preparedRound = spp.Round
		}
		if preparedRound.Cmp(spp.Round) != 0 || proposal.Hash() != spp.Digest {
			return false
		}
		iteration++
	}

	if preparedRound == nil {
		return hasQuorumOfRoundChangeMessagesForNil(roundChangeMessages, quorumSize)
	} else {
		return hasQuorumOfRoundChangeMessagesForPreparedRoundAndBlock(roundChangeMessages, preparedRound, proposal, quorumSize)
	}
}

// Checks whether a set of ROUND-CHANGE messages has `quorumSize` messages with nil prepared round and
// prepared block.
func hasQuorumOfRoundChangeMessagesForNil(roundChangeMessages []*SignedRoundChangePayload, quorumSize int) bool {
	nilCount := 0
	for _, m := range roundChangeMessages {
		if (m.PreparedRound == nil || m.PreparedRound.Cmp(common.Big0) == 0) && m.PreparedValue == nil {
			nilCount++
			if nilCount == quorumSize {
				return true
			}
		}
	}
	return false
}

// Checks whether a set of ROUND-CHANGE messages has some message with `preparedRound` and `preparedBlockDigest`,
// and has `quorumSize` messages with prepared round equal to nil or equal or lower than `preparedRound`.
func hasQuorumOfRoundChangeMessagesForPreparedRoundAndBlock(roundChangeMessages []*SignedRoundChangePayload, preparedRound *big.Int, preparedBlock istanbul.Proposal, quorumSize int) bool {
	lowerOrEqualRoundCount := 0
	hasMatchingMessage := false
	for _, m := range roundChangeMessages {
		if m.PreparedRound == nil || m.PreparedRound.Cmp(preparedRound) <= 0 {
			lowerOrEqualRoundCount++
			if m.PreparedRound != nil && m.PreparedRound.Cmp(preparedRound) == 0 && m.PreparedValue.Hash() == preparedBlock.Hash() {
				hasMatchingMessage = true
			}
			if lowerOrEqualRoundCount >= quorumSize && hasMatchingMessage {
				return true
			}
		}
	}

	return false
}

// Checks whether the round and block of a set of PREPARE messages of at least quorumSize match the
// preparedRound and preparedBlockDigest of a ROUND-CHANGE message.
func hasMatchingRoundChangeAndPrepares(roundChangeMessage *RoundChangeMessage, prepareMessages *messageSet, quorumSize int) bool {
	if prepareMessages.Size() < quorumSize {
		return false
	}

	for _, msg := range prepareMessages.messages {
		var prepare *Subject
		if err := msg.Decode(&prepare); err != nil {
			return false
		}
		if prepare.Digest != roundChangeMessage.PreparedBlockDigest {
			return false
		}
		if prepare.View.Round.Cmp(roundChangeMessage.PreparedRound) != 0 {
			return false
		}
	}
	return true
}

func hasMatchingRoundChangeAndPreparesFORQBFT(roundChange *RoundChangeMsg, prepareMessages []*SignedPreparePayload, quorumSize int) bool {
	if len(prepareMessages) < quorumSize {
		return false
	}

	for _, spp := range prepareMessages {
		if spp.Digest != roundChange.PreparedValue.Hash() {
			return false
		}
		if spp.Round.Cmp(roundChange.PreparedRound) != 0 {
			return false
		}
	}
	return true
}
