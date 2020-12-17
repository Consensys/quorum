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

	"github.com/ethereum/go-ethereum/rlp"
)

func TestMessageSetWithPreprepare(t *testing.T) {
	valSet := newTestValidatorSet(4)

	ms := newMessageSet(valSet)

	view := &View{
		Round:    new(big.Int),
		Sequence: new(big.Int),
	}
	pp := &Preprepare{
		View:     view,
		Proposal: makeBlock(1),
	}

	rawPP, err := rlp.EncodeToBytes(pp)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}
	msg := &message{
		Code:    msgPreprepare,
		Msg:     rawPP,
		Address: valSet.GetProposer().Address(),
	}

	err = ms.Add(msg)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	err = ms.Add(msg)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	if ms.Size() != 1 {
		t.Errorf("the size of message set mismatch: have %v, want 1", ms.Size())
	}
}

func TestMessageSetWithSubject(t *testing.T) {
	valSet := newTestValidatorSet(4)

	ms := newMessageSet(valSet)

	view := &View{
		Round:    new(big.Int),
		Sequence: new(big.Int),
	}

	sub := &Subject{
		View:   view,
		Digest: makeBlock(5).Hash(),
	}

	rawSub, err := rlp.EncodeToBytes(sub)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	msg := &message{
		Code:    msgPrepare,
		Msg:     rawSub,
		Address: valSet.GetProposer().Address(),
	}

	err = ms.Add(msg)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	err = ms.Add(msg)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	if ms.Size() != 1 {
		t.Errorf("the size of message set mismatch: have %v, want 1", ms.Size())
	}
}

// TestMessageSetEncodeDecode tests RLP encoding and decoding of messageSet, it does so by
// first encoding a RoundChangeMessage and then encoding a messageSet.
// It verifies encoding by decoding these messages and asserting the decoded values
func TestMessageSetEncodeDecode(t *testing.T) {
	valSet := newTestValidatorSet(4)

	ms := newMessageSet(valSet)

	proposal := makeBlock(5)

	view := &View{
		Round:    big.NewInt(0),
		Sequence: big.NewInt(5),
	}

	ms.view = view

	rc := &RoundChangeMessage{
		View:                view,
		PreparedRound:       big.NewInt(0),
		PreparedBlockDigest: proposal.Hash(),
	}

	encodedRC, err := rlp.EncodeToBytes(rc)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	firstVal := valSet.List()[0]

	msg := &message{
		Code:          0,
		Msg:           encodedRC,
		Address:       firstVal.Address(),
		Signature:     []byte{},
		CommittedSeal: []byte{},
	}

	ms.messages[firstVal.Address()] = msg

	encodedMS, err := Encode(ms)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	encodedMessages := &message{
		Code: msgRoundChange,
		Msg:  encodedMS,
	}

	var decodedMsgSet *messageSet
	err = encodedMessages.Decode(&decodedMsgSet)
	if err != nil {
		t.Errorf("failed to decode messageSet: %v", err)
	}
	decodedMsg := decodedMsgSet.messages[firstVal.Address()]

	if decodedMsg.Address != firstVal.Address() {
		t.Errorf("messageset mismatch: have %v, want %v", decodedMsg.Address, firstVal.Address())
	}

	encodedRCMsg := &message{
		Code:    msgRoundChange,
		Msg:     decodedMsg.Msg,
		Address: firstVal.Address(),
	}
	var rcMsg *RoundChangeMessage
	err = encodedRCMsg.Decode(&rcMsg)

	if rcMsg.PreparedBlockDigest != rc.PreparedBlockDigest {
		t.Errorf("rc message mismatch: have %v, want %v", rcMsg.PreparedBlockDigest, rc.PreparedBlockDigest)
	}

}
