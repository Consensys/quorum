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
	"github.com/ethereum/go-ethereum/rlp"
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
		View:          cv,
		PreparedRound: c.current.preparedRound,
		PreparedBlock: preparedBlock,
	}

	payload, err := Encode(rc)
	if err != nil {
		logger.Error("Failed to encode ROUND CHANGE", "rc", rc, "err", err)
		return
	}

	prepareMsgs := c.PreparedRoundPrepares
	if prepareMsgs == nil {
		prepareMsgs = newMessageSet(c.valSet)
	}

	var piggybackMsgPayload []byte
	piggybackMsg := &PiggybackMessages{PreparedMessages: prepareMsgs, RCMessages: newMessageSet(c.valSet)}
	piggybackMsgPayload, err = Encode(piggybackMsg)
	if err != nil {
		logger.Error("Failed to encode Piggyback messages accompanying ROUND CHANGE", "err", err)
		return
	}

	c.broadcast(&message{
		Code:          msgRoundChange,
		Msg:           payload,
		PiggybackMsgs: piggybackMsgPayload,
	})

}

func (c *core) handleRoundChange(msg *message, src istanbul.Validator) error {
	logger := c.logger.New("state", c.state, "from", src.Address().Hex())

	// Decode ROUND CHANGE message
	var rc *RoundChangeMessage
	if err := msg.Decode(&rc); err != nil {
		logger.Error("Failed to decode ROUND CHANGE", "err", err)
		return errFailedDecodeRoundChange
	}

	// Decode Prepare messages that piggyback Round Change message
	var piggybackMsgs *PiggybackMessages
	if msg.PiggybackMsgs != nil && len(msg.PiggybackMsgs) > 0 {
		if err := rlp.DecodeBytes(msg.PiggybackMsgs, &piggybackMsgs); err != nil {
			logger.Error("Failed to decode ROUND CHANGE Piggyback messages", "err", err)
			return errFailedDecodeRoundChange
		}

	}

	if err := c.checkMessage(msgRoundChange, rc.View); err != nil {
		return err
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
		err := c.roundChangeSet.Add(roundView.Round, msg, pr, pb, piggybackMsgs.PreparedMessages)
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

	} else if currentRoundMessages >= c.QuorumSize() && c.IsProposer() && c.justifyRoundChange(cv.Round) && c.current.preprepareSent.Cmp(cv.Round) < 0 {
		preparedRound, proposal := c.highestPrepared(cv.Round)
		if proposal == nil {
			proposal = c.current.pendingRequest.Proposal
		}

		var preparedMessages *messageSet
		if preparedRound == nil {
			preparedMessages = newMessageSet(c.valSet)
		} else {
			preparedMessages = c.roundChangeSet.preparedMessages[preparedRound.Uint64()]
		}

		r := &Request{
			Proposal:        proposal,
			RCMessages:      c.roundChangeSet.roundChanges[cv.Round.Uint64()],
			PrepareMessages: preparedMessages,
		}

		c.sendPreprepare(r)
	}
	return nil
}

// justifyRoundChange validates if the round change is valid or not
func (c *core) justifyRoundChange(round *big.Int) bool {
	if pr := c.roundChangeSet.preparedRounds[round.Uint64()]; pr == nil {
		return true
	}

	pr, pv := c.highestPrepared(round)
	// Check if the block in each prepared message is the one that is being proposed
	// To handle the case where a byzantine node can send an empty prepared block, check atleast Quorum of prepared blocks match the condition and not all
	i := 0
	for addr, msg := range c.roundChangeSet.preparedMessages[round.Uint64()].messages {
		var prepare *Subject
		if err := msg.Decode(&prepare); err != nil {
			c.logger.Error("Failed to decode Prepared Message", "err", err)
			continue
		}
		if prepare.Digest.Hash() != pv.Hash() {
			c.logger.Error("Highest Prepared Block does not match the Proposal", "Address", addr)
			continue
		}
		if prepare.View.Round.Uint64() != pr.Uint64() {
			c.logger.Error("Round in Prepared Block does not match the Highest Prepared Round", "Address", addr)
			continue
		}
		i++
		if i == c.QuorumSize() {
			// validated Quorum of prepared messages
			return true
		}
	}
	return false
}

// highestPrepared returns the highest Prepared Round and the corresponding Prepared Block
func (c *core) highestPrepared(round *big.Int) (*big.Int, istanbul.Proposal) {
	return c.roundChangeSet.preparedRounds[round.Uint64()], c.roundChangeSet.preparedBlocks[round.Uint64()]
}

// ----------------------------------------------------------------------------

func newRoundChangeSet(valSet istanbul.ValidatorSet) *roundChangeSet {
	return &roundChangeSet{
		validatorSet:     valSet,
		roundChanges:     make(map[uint64]*messageSet),
		preparedMessages: make(map[uint64]*messageSet),
		mu:               new(sync.Mutex),
	}
}

type roundChangeSet struct {
	validatorSet     istanbul.ValidatorSet
	roundChanges     map[uint64]*messageSet
	preparedMessages map[uint64]*messageSet
	preparedRounds   map[uint64]*big.Int
	preparedBlocks   map[uint64]istanbul.Proposal
	mu               *sync.Mutex
}

// Add adds the round and message into round change set
func (rcs *roundChangeSet) Add(r *big.Int, msg *message, preparedRound *big.Int, preparedBlock istanbul.Proposal, preparedMessages *messageSet) error {
	rcs.mu.Lock()
	defer rcs.mu.Unlock()

	round := r.Uint64()
	if rcs.roundChanges[round] == nil {
		rcs.roundChanges[round] = newMessageSet(rcs.validatorSet)
	}
	if err := rcs.roundChanges[round].Add(msg); err != nil {
		return err
	}

	if rcs.preparedRounds == nil {
		rcs.preparedRounds = make(map[uint64]*big.Int)
	}
	if rcs.preparedBlocks == nil {
		rcs.preparedBlocks = make(map[uint64]istanbul.Proposal)
	}

	if rcs.preparedRounds[round] == nil || preparedRound.Cmp(rcs.preparedRounds[round]) > 0 {
		rcs.preparedRounds[round] = preparedRound
		rcs.preparedBlocks[round] = preparedBlock
		rcs.preparedMessages[round] = preparedMessages
	}

	return nil
}

// higherRoundMessages returns the number of Round Change messages received for the round greater than the given round
// and from different validators
func (rcs *roundChangeSet) higherRoundMessages(round *big.Int) int {
	rcs.mu.Lock()
	defer rcs.mu.Unlock()

	addresses := make(map[common.Address]struct{})
	for k, rms := range rcs.roundChanges {
		if k > round.Uint64() {
			for addr := range rms.messages {
				addresses[addr] = struct{}{}
			}
		}
	}
	return len(addresses)
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
			delete(rcs.preparedRounds, k)
			delete(rcs.preparedBlocks, k)
			delete(rcs.preparedMessages, k)
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
