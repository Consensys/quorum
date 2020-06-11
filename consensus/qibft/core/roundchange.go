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
	"math/big"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core/types"
)

// sendNextRoundChange sends the ROUND CHANGE message with current round + 1
func (c *core) sendNextRoundChange() {
	cv := c.currentView()
	c.sendRoundChange(new(big.Int).Add(cv.Round, common.Big1))
}

// sendRoundChange sends the ROUND CHANGE message with the given round
func (c *core) sendRoundChange(round *big.Int) {
	logger := c.logger.New("state", c.state)

	cv := c.currentView()
	if cv.Round.Cmp(round) > 0 {
		logger.Error("Cannot send out the round change", "current round", cv.Round, "target round", round)
		return
	}

	// Now we have the new round number and sequence number
	prepares := c.PreparedRoundPrepares
	if prepares == nil {
		prepares = newMessageSet(c.valSet)
	}
	// If a block has not been prepared and a round change message occurs the preparedBlock is nil, setting it to NilBlock, so that decoding works fine
	preparedBlock := c.current.preparedBlock
	if preparedBlock == nil {
		preparedBlock = NilBlock()
	}
	cv = c.currentView()
	rc := &RoundChangeMessage{
		View:             cv,
		PreparedRound:    c.current.preparedRound,
		PreparedBlock:    preparedBlock,
		PreparedMessages: prepares,
	}

	payload, err := Encode(rc)
	if err != nil {
		logger.Error("Failed to encode ROUND CHANGE", "rc", rc, "err", err)
		return
	}

	c.broadcast(&message{
		Code: msgRoundChange,
		Msg:  payload,
	})

}

func (c *core) handleRoundChange(msg *message, src istanbul.Validator) error {
	logger := c.logger.New("state", c.state, "from", src.Address().Hex())

	// Decode ROUND CHANGE message
	var rc *RoundChangeMessage
	if err := msg.Decode(&rc); err != nil {
		logger.Error("Failed to decode ROUND CHANGE", "err", err)
		return errInvalidMessage
	}

	if err := c.checkMessage(msgRoundChange, rc.View); err != nil {
		return err
	}

	if rc.PreparedMessages != nil && rc.PreparedMessages.messages != nil {
		if c.validateFn != nil {
			if err := validateMsgSignature(rc.PreparedMessages.messages, c.validateFn); err != nil {
				logger.Error("Unable to validate round change message signature", "err", err)
				return err
			}
		}
	}

	cv := c.currentView()
	roundView := rc.View

	// Add the ROUND CHANGE message to its message set and return how many
	// messages we've got with the same round number and sequence number.
	if roundView.Round.Cmp(cv.Round) >= 0 {
		pb := rc.PreparedBlock
		pr := rc.PreparedRound
		// Checking if NilBlock was sent as prepared block
		if NilBlock().Hash() == pb.Hash() {
			pb = nil
			pr = nil
		}
		err := c.roundChangeSet.Add(roundView.Round, msg, pr, pb, rc.PreparedMessages, c.QuorumSize())
		if err != nil {
			logger.Warn("Failed to add round change message", "from", src, "msg", msg, "err", err)
			return err
		}
	}

	num := c.roundChangeSet.higherRoundMessages(cv.Round)
	currentRoundMessages := c.roundChangeSet.getRCMessagesForGivenRound(cv.Round)
	if num == c.valSet.F()+1 {
		newRound := c.roundChangeSet.getMinRoundChange(cv.Round)
		logger.Trace("Starting new Round", "round", newRound)
		c.startNewRound(newRound)
		c.sendRoundChange(newRound)
	} else if currentRoundMessages >= c.QuorumSize() && c.IsProposer() && c.current.preprepareSent.Cmp(cv.Round) < 0 {
		_, proposal := c.highestPrepared(cv.Round)
		if proposal == nil {
			proposal = c.current.pendingRequest.Proposal
		}

		roundChangeMessages := c.roundChangeSet.roundChanges[cv.Round.Uint64()]
		prepareMessages := c.roundChangeSet.prepareMessages[cv.Round.Uint64()]

		if !justify(proposal, roundChangeMessages, prepareMessages, c.QuorumSize()) {
			return nil
		}

		r := &Request{
			Proposal:        proposal,
			RCMessages:      roundChangeMessages,
			PrepareMessages: prepareMessages,
		}

		c.sendPreprepare(r)
	}
	return nil
}

// highestPrepared returns the highest Prepared Round and the corresponding Prepared Block
func (c *core) highestPrepared(round *big.Int) (*big.Int, istanbul.Proposal) {
	return c.roundChangeSet.highestPreparedRound[round.Uint64()], c.roundChangeSet.highestPreparedBlock[round.Uint64()]
}

// ----------------------------------------------------------------------------

func newRoundChangeSet(valSet istanbul.ValidatorSet) *roundChangeSet {
	return &roundChangeSet{
		validatorSet:         valSet,
		roundChanges:         make(map[uint64]*messageSet),
		prepareMessages:      make(map[uint64]*messageSet),
		highestPreparedRound: make(map[uint64]*big.Int),
		highestPreparedBlock: make(map[uint64]istanbul.Proposal),
		mu:                   new(sync.Mutex),
	}
}

type roundChangeSet struct {
	validatorSet         istanbul.ValidatorSet
	roundChanges         map[uint64]*messageSet
	prepareMessages      map[uint64]*messageSet
	highestPreparedRound map[uint64]*big.Int
	highestPreparedBlock map[uint64]istanbul.Proposal
	mu                   *sync.Mutex
}

func (rcs *roundChangeSet) NewRound(r *big.Int) {
	rcs.mu.Lock()
	defer rcs.mu.Unlock()
	round := r.Uint64()
	rcs.roundChanges[round] = newMessageSet(rcs.validatorSet)
	rcs.prepareMessages[round] = newMessageSet(rcs.validatorSet)
}

// Add adds the round and message into round change set
func (rcs *roundChangeSet) Add(r *big.Int, msg *message, preparedRound *big.Int, preparedBlock istanbul.Proposal, prepareMessages *messageSet, quorumSize int) error {
	rcs.mu.Lock()
	defer rcs.mu.Unlock()

	round := r.Uint64()
	if err := rcs.roundChanges[round].Add(msg); err != nil {
		return err
	}

	if preparedRound != nil && (rcs.highestPreparedRound[round] == nil || preparedRound.Cmp(rcs.highestPreparedRound[round]) > 0) {
		var roundChangeMessage *RoundChangeMessage
		if err := msg.Decode(&roundChangeMessage); err != nil {
			return err
		}
		if hasMatchingRoundChangeAndPrepares(roundChangeMessage, prepareMessages, quorumSize) {
			rcs.highestPreparedRound[round] = preparedRound
			rcs.highestPreparedBlock[round] = preparedBlock
			rcs.prepareMessages[round] = prepareMessages
		}
	}

	return nil
}

// higherRoundMessages returns the number of Round Change messages received for the round greater than the given round
func (rcs *roundChangeSet) higherRoundMessages(round *big.Int) int {
	rcs.mu.Lock()
	defer rcs.mu.Unlock()

	num := 0
	for k, rms := range rcs.roundChanges {
		if k > round.Uint64() {
			num = num + len(rms.messages)
		}
	}
	return num
}

func (rcs *roundChangeSet) getRCMessagesForGivenRound(round *big.Int) int {
	rcs.mu.Lock()
	defer rcs.mu.Unlock()

	if rms := rcs.roundChanges[round.Uint64()]; rms != nil {
		return len(rms.messages)
	}
	return 0
}

// getMinRoundChange returns the minimum round greater than the given round
func (rcs *roundChangeSet) getMinRoundChange(round *big.Int) *big.Int {
	rcs.mu.Lock()
	defer rcs.mu.Unlock()

	var keys []int
	for k := range rcs.roundChanges {
		if k > round.Uint64() {
			keys = append(keys, int(k))
		}
	}
	sort.Ints(keys)
	if len(keys) == 0 {
		return round
	}
	return big.NewInt(int64(keys[0]))
}

// ClearLowerThan deletes the messages for round earlier than the given round
func (rcs *roundChangeSet) ClearLowerThan(round *big.Int) {
	rcs.mu.Lock()
	defer rcs.mu.Unlock()

	for k, rms := range rcs.roundChanges {
		if len(rms.Values()) == 0 || k < round.Uint64() {
			delete(rcs.roundChanges, k)
			delete(rcs.highestPreparedRound, k)
			delete(rcs.highestPreparedBlock, k)
			delete(rcs.prepareMessages, k)
		}
	}
}

// MaxRound returns the max round which the number of messages is equal or larger than num
func (rcs *roundChangeSet) MaxRound(num int) *big.Int {
	rcs.mu.Lock()
	defer rcs.mu.Unlock()

	var maxRound *big.Int
	for k, rms := range rcs.roundChanges {
		if rms.Size() < num {
			continue
		}
		r := big.NewInt(int64(k))
		if maxRound == nil || maxRound.Cmp(r) < 0 {
			maxRound = r
		}
	}
	return maxRound
}

// NilBlock represents a nil block and is sent in RoundChangeMessage if PreparedBlock is nil
func NilBlock() *types.Block {
	header := &types.Header{
		Difficulty: big.NewInt(0),
		Number:     big.NewInt(0),
		GasLimit:   0,
		GasUsed:    0,
		Time:       0,
	}
	block := &types.Block{}
	return block.WithSeal(header)
}
