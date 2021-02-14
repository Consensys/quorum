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
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// broadcastNextRoundChange sends the ROUND CHANGE message with current round + 1
func (c *core) broadcastNextRoundChange() {
	cv := c.currentView()
	c.broadcastRoundChange(new(big.Int).Add(cv.Round, common.Big1))
}

// sendRoundChange sends the ROUND CHANGE message with the given round
func (c *core) broadcastRoundChange(round *big.Int) {
	logger := c.logger.New("state", c.state)

	cv := c.currentView()
	if cv.Round.Cmp(round) > 0 {
		logger.Error("Cannot send out the round change", "current round", cv.Round, "target round", round)
		return
	}

	rcMsg := &RoundChangeMsg{
		SignedRoundChangePayload: SignedRoundChangePayload{
			CommonPayload: CommonPayload{
				code:      roundChangeMsgCode,
				source:    c.Address(),
				Sequence:  c.current.Sequence(),
				Round:     round,
				signature: nil,
			},
			PreparedRound: nil,
			PreparedValue: nil,
		},
	}

	// Fill in prepared round and prepared value
	if c.current.preparedRound != nil && c.current.preparedBlock != nil {
		rcMsg.PreparedRound = c.current.preparedRound
		rcMsg.PreparedValue = c.current.preparedBlock.(*types.Block)
	}

	// Sign message
	encodedPayload, err := rcMsg.EncodePayload()
	if err != nil {
		logger.Error("QBFT: Failed to encode round-change message", "msg", rcMsg, "err", err)
		return
	}
	rcMsg.signature, err = c.backend.Sign(encodedPayload)
	if err != nil {
		logger.Error("QBFT: Failed to sign round-change message", "msg", rcMsg, "err", err)
		return
	}

	// Add justification
	if c.QBFTPreparedPrepares != nil {
		rcMsg.Justification = c.QBFTPreparedPrepares
		logger.Info("QBFT: On RoundChange", "justification", rcMsg.Justification)
	}

	// RLP-encode message
	data, err := rlp.EncodeToBytes(rcMsg)
	if err != nil {
		logger.Error("QBFT: Failed to encode round-change message", "msg", rcMsg, "err", err)
		return
	}

	logger.Info("QBFT: broadcast round-change message", "msg", rcMsg)
	// Broadcast RLP-encoded message
	if err = c.backend.Broadcast(c.valSet, roundChangeMsgCode, data); err != nil {
		logger.Error("QBFT: Failed to broadcast message", "msg", rcMsg, "err", err)
		return
	}
}

func (c *core) handleRoundChange(roundChange *RoundChangeMsg) error {
	logger := c.logger.New("state", c.state)

	logger.Info("QBFT: handleRoundChange", "m", roundChange)

	view := roundChange.View()
	currentRound := c.currentView().Round

	// Add the ROUND CHANGE message to its message set and return how many
	// messages we've got with the same round number and sequence number.
	if view.Round.Cmp(currentRound) >= 0 {
		var prepareMessages []*SignedPreparePayload = nil
		var pr *big.Int = nil
		var pb *types.Block = nil
		if roundChange.PreparedRound != nil && roundChange.PreparedValue != nil && roundChange.Justification != nil && len(roundChange.Justification) > 0 {
			prepareMessages = roundChange.Justification
			pr = roundChange.PreparedRound
			pb = roundChange.PreparedValue
		}
		err := c.roundChangeSet.Add(view.Round, roundChange, pr, pb, prepareMessages, c.QuorumSize())
		if err != nil {
			logger.Warn("Failed to add round change message", "msg", roundChange, "err", err)
			return err
		}
	}

	num := c.roundChangeSet.higherRoundMessages(currentRound)
	currentRoundMessages := c.roundChangeSet.getRCMessagesForGivenRound(currentRound)
	log.Info("QBFT: handleRoundChange count", "higherRoundMsgs", num, "currentRoundMsgs", currentRoundMessages)
	if num == c.valSet.F()+1 {
		newRound := c.roundChangeSet.getMinRoundChange(currentRound)
		logger.Trace("Starting new Round", "round", newRound)
		c.startNewRound(newRound)
		c.broadcastRoundChange(newRound)
	} else if currentRoundMessages >= c.QuorumSize() && c.IsProposer() && c.current.preprepareSent.Cmp(currentRound) < 0 {
		_, proposal := c.highestPrepared(currentRound)
		if proposal == nil {
			proposal = c.current.pendingRequest.Proposal
		}

		roundChangeMessages := c.roundChangeSet.roundChanges[currentRound.Uint64()]
		prepareMessages := c.roundChangeSet.prepareMessages[currentRound.Uint64()]

		// Justification
		rcSignedPayloads := make([]*SignedRoundChangePayload, 0)
		for _, m := range roundChangeMessages.Values() {
			rcMsg := m.(*RoundChangeMsg)
			rcSignedPayloads = append(rcSignedPayloads, &rcMsg.SignedRoundChangePayload)
		}

		if !justify(proposal, rcSignedPayloads, prepareMessages, c.QuorumSize()) {
			return nil
		}

		log.Info("QBFT: handleRoundChange - broadcasting pre-prepare")
		r := &Request{
			Proposal:        proposal,
			RCMessages:      roundChangeMessages,
			PrepareMessages: prepareMessages,
		}
		c.sendPreprepareMsg(r)
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
		roundChanges:         make(map[uint64]*qbftMsgSet),
		prepareMessages:      make(map[uint64][]*SignedPreparePayload),
		highestPreparedRound: make(map[uint64]*big.Int),
		highestPreparedBlock: make(map[uint64]istanbul.Proposal),
		mu:                   new(sync.Mutex),
	}
}

type roundChangeSet struct {
	validatorSet         istanbul.ValidatorSet
	roundChanges         map[uint64]*qbftMsgSet
	prepareMessages      map[uint64][]*SignedPreparePayload
	highestPreparedRound map[uint64]*big.Int
	highestPreparedBlock map[uint64]istanbul.Proposal
	mu                   *sync.Mutex
}

func (rcs *roundChangeSet) NewRound(r *big.Int) {
	rcs.mu.Lock()
	defer rcs.mu.Unlock()
	round := r.Uint64()
	rcs.roundChanges[round] = newQBFTMsgSet(rcs.validatorSet)
	rcs.prepareMessages[round] = make([]*SignedPreparePayload, 0)
}

// Add adds the round and message into round change set
func (rcs *roundChangeSet) Add(r *big.Int, msg QBFTMessage, preparedRound *big.Int, preparedBlock istanbul.Proposal, prepareMessages []*SignedPreparePayload, quorumSize int) error {
	rcs.mu.Lock()
	defer rcs.mu.Unlock()

	round := r.Uint64()
	if rcs.roundChanges[round] == nil {
		rcs.roundChanges[round] = newQBFTMsgSet(rcs.validatorSet)
	}
	if err := rcs.roundChanges[round].Add(msg); err != nil {
		return err
	}

	if preparedRound != nil && (rcs.highestPreparedRound[round] == nil || preparedRound.Cmp(rcs.highestPreparedRound[round]) > 0) {
		roundChange := msg.(*RoundChangeMsg)
		if hasMatchingRoundChangeAndPreparesFORQBFT(roundChange, prepareMessages, quorumSize) {
			rcs.highestPreparedRound[round] = preparedRound
			rcs.highestPreparedBlock[round] = preparedBlock
			rcs.prepareMessages[round] = prepareMessages
		}
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

// NilBlock represents a nil block and is sent in RoundChangeMessage if PreparedBlockDigest is nil
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
