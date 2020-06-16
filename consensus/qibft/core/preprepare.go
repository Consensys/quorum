// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"time"

	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
)

func (c *core) sendPreprepare(request *Request) {
	logger := c.logger.New("state", c.state)
	rcMessages := request.RCMessages
	if rcMessages == nil {
		rcMessages = newMessageSet(c.valSet)
	}

	// If I'm the proposer and I have the same sequence with the proposal
	if c.current.Sequence().Cmp(request.Proposal.Number()) == 0 && c.IsProposer() {
		curView := c.currentView()
		preprepare, err := Encode(&Preprepare{
			View:             curView,
			Proposal:         request.Proposal,
			RCMessages:       rcMessages,
			PreparedMessages: request.PrepareMessages,
		})
		if err != nil {
			logger.Error("Failed to encode", "view", curView)
			return
		}
		c.broadcast(&message{
			Code: msgPreprepare,
			Msg:  preprepare,
		})
		// Set the preprepareSent to the current round
		c.current.preprepareSent = curView.Round
	}
}

func (c *core) handlePreprepare(msg *message, src istanbul.Validator) error {
	logger := c.logger.New("from", src, "state", c.state)

	// Decode PRE-PREPARE
	var preprepare *Preprepare
	err := msg.Decode(&preprepare)
	if err != nil {
		logger.Debug("Failed to decode preprepare message", "err", err)
		return errFailedDecodePreprepare
	}

	// Ensure we have the same view with the PRE-PREPARE message
	// If it is old message, see if we need to broadcast COMMIT
	if err := c.checkMessage(msgPreprepare, preprepare.View); err != nil {
		if err == errOldMessage {
			// Get validator set for the given proposal
			valSet := c.backend.ParentValidators(preprepare.Proposal).Copy()
			previousProposer := c.backend.GetProposer(preprepare.Proposal.Number().Uint64() - 1)
			valSet.CalcProposer(previousProposer, preprepare.View.Round.Uint64())
			// Broadcast COMMIT if it is an existing block
			// 1. The proposer needs to be a proposer matches the given (Sequence + Round)
			// 2. The given block must exist
			if valSet.IsProposer(src.Address()) && c.backend.HasPropsal(preprepare.Proposal.Hash(), preprepare.Proposal.Number()) {
				c.sendCommitForOldBlock(preprepare.View, preprepare.Proposal)
				return nil
			}
		}
		return err
	}

	// Check if the message comes from current proposer
	if !c.valSet.IsProposer(src.Address()) {
		logger.Warn("Ignore preprepare messages from non-proposer")
		return errNotFromProposer
	}

	if preprepare.View.Round.Uint64() > 0 && !c.validatePrepreparedMessage(preprepare, src) {
		logger.Error("Unable to verify prepared block in Round Change messages")
		return errInvalidPreparedBlock
	}

	// Verify the proposal we received
	if duration, err := c.backend.Verify(preprepare.Proposal); err != nil {
		// if it's a future block, we will handle it again after the duration
		if err == consensus.ErrFutureBlock {
			logger.Info("Proposed block will be handled in the future", "err", err, "duration", duration)
			c.stopFuturePreprepareTimer()
			c.futurePreprepareTimer = time.AfterFunc(duration, func() {
				c.sendEvent(backlogEvent{
					src: src,
					msg: msg,
				})
			})
		}
		return err
	}

	// Here is about to accept the PRE-PREPARE
	if c.state == StateAcceptRequest {
		c.newRoundChangeTimer()
		c.acceptPreprepare(preprepare)
		c.setState(StatePreprepared)
		c.sendPrepare()
	}

	return nil
}

func (c *core) acceptPreprepare(preprepare *Preprepare) {
	c.consensusTimestamp = time.Now()
	c.current.SetPreprepare(preprepare)
}

// validatePrepreparedMessage validates Preprepared message received
func (c *core) validatePrepreparedMessage(preprepare *Preprepare, src istanbul.Validator) bool {
	logger := c.logger.New("from", src, "state", c.state)
	highestPreparedRound, validRC := c.checkRoundChangeMessages(preprepare, src)
	if !validRC {
		logger.Error("Unable to verify Round Change messages in Preprepare")
		return false
	}
	if highestPreparedRound != 0 && !c.checkPreparedMessages(preprepare, highestPreparedRound, src) {
		logger.Error("Unable to verify Prepared messages in Preprepare")
		return false
	}
	return true
}

// checkRoundChangeMessages verifies if the Round Change message is signed by a valid validator and
// Also, check if the proposal was the preparedBlock corresponding to the highest preparedRound
func (c *core) checkRoundChangeMessages(preprepare *Preprepare, src istanbul.Validator) (uint64, bool) {
	logger := c.logger.New("from", src, "state", c.state)

	if preprepare.RCMessages != nil && preprepare.RCMessages.messages != nil {
		var preparedRound uint64 = 0
		var preparedBlock istanbul.Proposal
		for _, msg := range preprepare.RCMessages.messages {
			var rc *RoundChangeMessage
			if err := msg.Decode(&rc); err != nil {
				logger.Error("Failed to decode ROUND CHANGE", "err", err)
				return 0, false
			}
			if rc.PreparedRound.Uint64() > preparedRound {
				preparedRound = rc.PreparedRound.Uint64()
				preparedBlock = rc.PreparedBlock
			}
		}
		if preparedRound == 0 {
			return preparedRound, true
		}
		if preparedRound > 0 {
			return preparedRound, preparedBlock == preprepare.Proposal
		}
	}
	return 0, false
}

// checkPreparedMessages verifies if a Quorum of Prepared messages were received and
// the block in each prepared message is the same as the proposal and is prepared in the same round
func (c *core) checkPreparedMessages(preprepare *Preprepare, highestPreparedRound uint64, src istanbul.Validator) bool {
	logger := c.logger.New("from", src, "state", c.state)
	if preprepare.PreparedMessages != nil && preprepare.PreparedMessages.messages != nil {
		// Number of prepared messages should not be less than Quorum of messages
		if len(preprepare.PreparedMessages.messages) < c.QuorumSize() {
			logger.Error("Quorum of Prepared messages not found in Preprepare messages")
			return false
		}
		// Check if the block in each prepared message is the one that is being proposed
		for addr, msg := range preprepare.PreparedMessages.messages {
			var prepare *Subject
			if err := msg.Decode(&prepare); err != nil {
				logger.Error("Failed to decode Prepared Message", "err", err)
				return false
			}
			if prepare.Digest.Hash() != preprepare.Proposal.Hash() {
				logger.Error("Prepared block does not match the Proposal", "Address", addr)
				return false
			}
			if prepare.View.Round.Uint64() != highestPreparedRound {
				logger.Error("Round in Prepared Block does not match the Highest Prepared Round", "Address", addr)
				return false
			}
		}
		return true
	}
	return false
}
