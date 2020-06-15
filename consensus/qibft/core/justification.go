package core

import (
	"math/big"

	istanbul "github.com/ethereum/go-ethereum/consensus/qibft"
)

// Returns true if the `proposal` is justified by the set `roundChangeMessages` of ROUND-CHANGE messages
// and by the set `prepareMessages` of PREPARE messages.
// For this we must either have:
//     - a quorum of ROUND-CHANGE messages with preparedRound and preparedBlock equal to nil; or
//     - a ROUND-CHANGE message (1) whose preparedRound is not nil and is equal or higher than the
//           preparedRound of `quorumSize` ROUND-CHANGE messages and (2) whose preparedRound and
//           preparedBlock match the round and block of `quorumSize` PREPARE messages.
func justify(proposal istanbul.Proposal, roundChangeMessages *messageSet, prepareMessages *messageSet, quorumSize int) bool {
	// Check the size of the set of ROUND-CHANGE messages
	if roundChangeMessages.Size() < quorumSize {
		return false
	}

	// Check the size of the set of PREPARE messages
	if prepareMessages.Size() != 0 && prepareMessages.Size() < quorumSize {
		return false
	}

	// If there are PREPARE messages, they all need to have the same round and match `proposal`
	var preparedRound *big.Int
	for _, msg := range prepareMessages.messages {
		var prepareMessage *Subject
		if err := msg.Decode(&prepareMessage); err != nil {
			return false
		}
		if preparedRound != nil { // Get the round of the first message
			preparedRound = prepareMessage.View.Round
		}
		if preparedRound != prepareMessage.View.Round || proposal.Hash() != prepareMessage.Digest.Hash() {
			return false
		}
	}

	if preparedRound == nil {
		return hasQuorumOfRoundChangeMessagesForNil(roundChangeMessages, quorumSize)
	} else {
		return hasQuorumOfRoundChangeMessagesForPreparedRoundAndBlock(roundChangeMessages, preparedRound, proposal, quorumSize)
	}
}

// Checks whether a set of ROUND-CHANGE messages has `quorumSize` messages with nil prepared round and
// prepared block.
func hasQuorumOfRoundChangeMessagesForNil(roundChangeMessages *messageSet, quorumSize int) bool {
	nilCount := 0
	for _, msg := range roundChangeMessages.messages {
		var roundChangeMessage *RoundChangeMessage
		if err := msg.Decode(&roundChangeMessage); err != nil {
			continue
		}
		if roundChangeMessage.PreparedRound == nil && roundChangeMessage.PreparedBlock == NilBlock() {
			nilCount++
			if nilCount == quorumSize {
				return true
			}
		}
	}
	return false
}

// Checks whether a set of ROUND-CHANGE messages has some message with `preparedRound` and `preparedBlock`,
// and has `quorumSize` messages with prepared round equal to nil or equal or lower than `preparedRound`.
func hasQuorumOfRoundChangeMessagesForPreparedRoundAndBlock(roundChangeMessages *messageSet, preparedRound *big.Int, preparedBlock istanbul.Proposal, quorumSize int) bool {
	lowerOrEqualRoundCount := 0
	hasMatchingMessage := false
	for _, msg := range roundChangeMessages.messages {
		var roundChangeMessage *RoundChangeMessage
		if err := msg.Decode(&roundChangeMessage); err != nil {
			continue
		}

		if roundChangeMessage.PreparedRound == nil || roundChangeMessage.PreparedRound.Cmp(preparedRound) <= 0 {
			lowerOrEqualRoundCount++
			if roundChangeMessage.PreparedRound != nil && roundChangeMessage.PreparedRound.Cmp(preparedRound) == 0 && roundChangeMessage.PreparedBlock.Hash() == preparedBlock.Hash() {
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
// preparedRound and preparedBlock of a ROUND-CHANGE message.
func hasMatchingRoundChangeAndPrepares(roundChangeMessage *RoundChangeMessage, prepareMessages *messageSet, quorumSize int) bool {
	if prepareMessages.Size() < quorumSize {
		return false
	}

	for _, msg := range prepareMessages.messages {
		var prepare *Subject
		if err := msg.Decode(&prepare); err != nil {
			return false
		}
		if prepare.Digest.Hash() != roundChangeMessage.PreparedBlock.Hash() {
			return false
		}
		if prepare.View.Round.Uint64() != roundChangeMessage.PreparedRound.Uint64() {
			return false
		}
	}
	return true
}
