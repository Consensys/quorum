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
	"testing"

	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
)

func TestRoundChangeSet(t *testing.T) {
	vset := validator.NewSet(generateValidators(4), istanbul.RoundRobin)
	rc := newRoundChangeSet(vset)
	rc.NewRound(big.NewInt(1))

	view := &View{
		Sequence: big.NewInt(1),
		Round:    big.NewInt(1),
	}
	r := &RoundChangeMessage{
		View:                view,
		PreparedRound:       big.NewInt(1),
		PreparedBlockDigest: newTestProposal().Hash(),
	}
	//m, _ := Encode(r)

	// Test Add()
	// Add message from all validators
	for i, v := range vset.List() {
		roundChange := RoundChangeMsg{
			CommonMsg:            CommonMsg{
				code:           roundChangeMsgCode,
				source:         v.Address(),
				Sequence:       big.NewInt(1),
				Round:          big.NewInt(1),
				EncodedPayload: nil,
				Signature:      nil,
			},
			PreparedRound:        big.NewInt(1),
			PreparedValue:        newTestProposal(),
			Justification:        nil,
			EncodedSignedPayload: nil,
		}
		
		/*msg := &message{
			Code:    msgRoundChange,
			Msg:     m,
			Address: v.Address(),
		}*/
		rc.Add(view.Round, &roundChange, r.PreparedRound, newTestProposal(), newMessageSet(vset), vset.Size())
		if rc.roundChanges[view.Round.Uint64()].Size() != i+1 {
			t.Errorf("the size of round change messages mismatch: have %v, want %v", rc.roundChanges[view.Round.Uint64()].Size(), i+1)
		}
	}

	// Add message again from all validators, but the size should be the same
	for _, v := range vset.List() {
		roundChange := RoundChangeMsg{
			CommonMsg:            CommonMsg{
				code:           roundChangeMsgCode,
				source:         v.Address(),
				Sequence:       big.NewInt(1),
				Round:          big.NewInt(1),
				EncodedPayload: nil,
				Signature:      nil,
			},
			PreparedRound:        big.NewInt(1),
			PreparedValue:        newTestProposal(),
			Justification:        nil,
			EncodedSignedPayload: nil,
		}
		rc.Add(view.Round, &roundChange, r.PreparedRound, newTestProposal(), newMessageSet(vset), vset.Size())
		if rc.roundChanges[view.Round.Uint64()].Size() != vset.Size() {
			t.Errorf("the size of round change messages mismatch: have %v, want %v", rc.roundChanges[view.Round.Uint64()].Size(), vset.Size())
		}
	}

	// Test MaxRound()
	for i := 0; i < 10; i++ {
		maxRound := rc.MaxRound(i)
		if i <= vset.Size() {
			if maxRound == nil || maxRound.Cmp(view.Round) != 0 {
				t.Errorf("max round mismatch: have %v, want %v", maxRound, view.Round)
			}
		} else if maxRound != nil {
			t.Errorf("max round mismatch: have %v, want nil", maxRound)
		}
	}

	// Test ClearLowerThan()
	for i := int64(0); i < 2; i++ {
		rc.ClearLowerThan(big.NewInt(i))
		if rc.roundChanges[view.Round.Uint64()].Size() != vset.Size() {
			t.Errorf("the size of round change messages mismatch: have %v, want %v", rc.roundChanges[view.Round.Uint64()].Size(), vset.Size())
		}
	}
	rc.ClearLowerThan(big.NewInt(2))
	if rc.roundChanges[view.Round.Uint64()] != nil {
		t.Errorf("the change messages mismatch: have %v, want nil", rc.roundChanges[view.Round.Uint64()])
	}
}

func TestGetMinRoundChange(t *testing.T) {
	rcs := getRoundChangeSetForPositveTests()
	minRC := rcs.getMinRoundChange(big.NewInt(1))
	if minRC.Uint64() != 2 {
		t.Errorf("min Round Change mismatch: have %v, want 2", minRC.Uint64())
	}
}

func TestClearLowerThan(t *testing.T) {
	rcs := getRoundChangeSetForPositveTests()
	rcs.ClearLowerThan(big.NewInt(3))
	if len(rcs.roundChanges) > 0 {
		t.Errorf("Number of Round Change messages mismatch: have %v, want 0", len(rcs.roundChanges))
	}

	rcs = getRoundChangeSetForPositveTests()
	rcs.ClearLowerThan(big.NewInt(2))
	rcMsgs := rcs.roundChanges[2]
	if len(rcMsgs.messages) != 3 {
		t.Errorf("Number of Round Change messages mismatch: have %v, want 3", len(rcs.roundChanges))
	}
}

func TestGetRCMessagesForGivenRound(t *testing.T) {
	rcs := getRoundChangeSetForPositveTests()
	numOfRCMessages := rcs.getRCMessagesForGivenRound(big.NewInt(2))
	if numOfRCMessages != 3 {
		t.Errorf("Number of RoundChange messages for the given round do not match : have %v, want 3", numOfRCMessages)
	}

	// Check for messages for round 1, should return 0
	numOfRCMessages = rcs.getRCMessagesForGivenRound(big.NewInt(0))
	if numOfRCMessages != 0 {
		t.Errorf("Number of RoundChange messages for the given round do not match : have %v, want 0", numOfRCMessages)
	}
}

func TestHigherRoundMessages(t *testing.T) {
	rcs := getRoundChangeSetForPositveTests()
	// The above rcs messages are for round 2, so messages higher than that should be 0
	higherRCMsgs := rcs.higherRoundMessages(big.NewInt(2))
	if higherRCMsgs != 0 {
		t.Errorf("Number of RoundChange messages for the given round do not match : have %v, want 0", higherRCMsgs)
	}
	// The rcs messages are for round 2, so messages higher than round 1 should be 3
	higherRCMsgs = rcs.higherRoundMessages(big.NewInt(1))
	if higherRCMsgs != 3 {
		t.Errorf("Number of RoundChange messages for the given round do not match : have %v, want 3", higherRCMsgs)
	}
}

func getRoundChangeSetForPositveTests() *roundChangeSet {
	vset := validator.NewSet(generateValidators(4), istanbul.RoundRobin)

	view := &View{
		Sequence: big.NewInt(1),
		Round:    big.NewInt(2),
	}
	proposal := makeBlock(1)

	rcs := newRoundChangeSet(vset)
	rcs.NewRound(big.NewInt(2))

	/*encodedRCMsg1, _ := Encode(&RoundChangeMessage{
		View:                view,
		PreparedRound:       big.NewInt(1),
		PreparedBlockDigest: proposal.Hash(),
	})*/

	msg1 := &RoundChangeMsg{
		CommonMsg:            CommonMsg{
			code:           roundChangeMsgCode,
			source:         vset.GetByIndex(0).Address(),
			Sequence:       big.NewInt(1),
			Round:          big.NewInt(1),
			EncodedPayload: nil,
			Signature:      nil,
		},
		PreparedRound:        big.NewInt(1),
		PreparedValue:        newTestProposal(),
		Justification:        nil,
		EncodedSignedPayload: nil,
	}

	msg2 := &RoundChangeMsg{
		CommonMsg:            CommonMsg{
			code:           roundChangeMsgCode,
			source:         vset.GetByIndex(1).Address(),
			Sequence:       big.NewInt(1),
			Round:          big.NewInt(1),
			EncodedPayload: nil,
			Signature:      nil,
		},
		PreparedRound:        big.NewInt(1),
		PreparedValue:        newTestProposal(),
		Justification:        nil,
		EncodedSignedPayload: nil,
	}

	msg3 := &RoundChangeMsg{
		CommonMsg:            CommonMsg{
			code:           roundChangeMsgCode,
			source:         vset.GetByIndex(2).Address(),
			Sequence:       big.NewInt(1),
			Round:          big.NewInt(1),
			EncodedPayload: nil,
			Signature:      nil,
		},
		PreparedRound:        big.NewInt(1),
		PreparedValue:        newTestProposal(),
		Justification:        nil,
		EncodedSignedPayload: nil,
	}

	rcs.Add(view.Round, msg1, big.NewInt(1), proposal, newMessageSet(rcs.validatorSet), vset.Size())
	rcs.Add(view.Round, msg2, big.NewInt(1), proposal, newMessageSet(rcs.validatorSet), vset.Size())
	rcs.Add(view.Round, msg3, big.NewInt(1), proposal, newMessageSet(rcs.validatorSet), vset.Size())

	return rcs
}
