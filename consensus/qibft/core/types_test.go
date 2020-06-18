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
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/rlp"
)

func testPreprepare(t *testing.T) {
	valSet := newTestValidatorSet(4)
	pp := &PreprepareWithPiggybackMsgs{
		Preprepare: &Preprepare{
			View: &View{
				Round:    big.NewInt(1),
				Sequence: big.NewInt(2),
			},
			Proposal: makeBlock(1),
		},
		PiggybackMessages: &PiggybackMessages{
			RCMessages: newMessageSet(valSet),
		},
	}
	prepreparePayload, _ := Encode(pp)

	m := &message{
		Code:    msgPreprepare,
		Msg:     prepreparePayload,
		Address: common.HexToAddress("0x1234567890"),
	}

	msgPayload, err := m.Payload()
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	decodedMsg := new(message)
	err = decodedMsg.FromPayload(msgPayload, nil)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	var decodedPP *PreprepareWithPiggybackMsgs
	err = decodedMsg.Decode(&decodedPP)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	// if block is encoded/decoded by rlp, we cannot to compare interface data type using reflect.DeepEqual. (like istanbul.Proposal)
	// so individual comparison here.
	if !reflect.DeepEqual(pp.Preprepare.Proposal.Hash(), decodedPP.Preprepare.Proposal.Hash()) {
		t.Errorf("proposal hash mismatch: have %v, want %v", decodedPP.Preprepare.Proposal.Hash(), pp.Preprepare.Proposal.Hash())
	}

	if !reflect.DeepEqual(pp.Preprepare.View, decodedPP.Preprepare.View) {
		t.Errorf("view mismatch: have %v, want %v", decodedPP.Preprepare.View, pp.Preprepare.View)
	}

	if !reflect.DeepEqual(pp.Preprepare.Proposal.Number(), decodedPP.Preprepare.Proposal.Number()) {
		t.Errorf("proposal number mismatch: have %v, want %v", decodedPP.Preprepare.Proposal.Number(), pp.Preprepare.Proposal.Number())
	}
}

func testSubject(t *testing.T) {
	s := &Subject{
		View: &View{
			Round:    big.NewInt(1),
			Sequence: big.NewInt(2),
		},
		Digest: makeBlock(5),
	}

	subjectPayload, _ := Encode(s)

	m := &message{
		Code:    msgPreprepare,
		Msg:     subjectPayload,
		Address: common.HexToAddress("0x1234567890"),
	}

	msgPayload, err := m.Payload()
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	decodedMsg := new(message)
	err = decodedMsg.FromPayload(msgPayload, nil)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	var decodedSub *Subject
	err = decodedMsg.Decode(&decodedSub)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	if !reflect.DeepEqual(s.View, decodedSub.View) || s.Digest.Hash().Hex() != decodedSub.Digest.Hash().Hex() {
		t.Errorf("subject mismatch: have %v, want %v", decodedSub, s)
	}
}

func testSubjectWithSignature(t *testing.T) {
	s := &Subject{
		View: &View{
			Round:    big.NewInt(1),
			Sequence: big.NewInt(2),
		},
		Digest: makeBlock(5),
	}
	expectedSig := []byte{0x01}

	subjectPayload, _ := Encode(s)
	// 1. Encode test
	address := common.HexToAddress("0x1234567890")
	m := &message{
		Code:          msgPrepare,
		Msg:           subjectPayload,
		Address:       address,
		Signature:     expectedSig,
		CommittedSeal: []byte{},
	}

	msgPayload, err := m.Payload()
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	// 2. Decode test
	// 2.1 Test normal validate func
	decodedMsg := new(message)
	err = decodedMsg.FromPayload(msgPayload, func(data []byte, sig []byte) (common.Address, error) {
		return address, nil
	})
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	if !reflect.DeepEqual(decodedMsg, m) {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	// 2.2 Test nil validate func
	decodedMsg = new(message)
	err = decodedMsg.FromPayload(msgPayload, nil)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(decodedMsg, m) {
		t.Errorf("message mismatch: have %v, want %v", decodedMsg, m)
	}

	// 2.3 Test failed validate func
	decodedMsg = new(message)
	err = decodedMsg.FromPayload(msgPayload, func(data []byte, sig []byte) (common.Address, error) {
		return common.Address{}, istanbul.ErrUnauthorizedAddress
	})
	if err != istanbul.ErrUnauthorizedAddress {
		t.Errorf("error mismatch: have %v, want %v", err, istanbul.ErrUnauthorizedAddress)
	}
}

func TestMessageEncodeDecode(t *testing.T) {
	testPreprepare(t)
	testSubject(t)
	testSubjectWithSignature(t)
}

func TestViewCompare(t *testing.T) {
	// test equality
	srvView := &View{
		Sequence: big.NewInt(2),
		Round:    big.NewInt(1),
	}
	tarView := &View{
		Sequence: big.NewInt(2),
		Round:    big.NewInt(1),
	}
	if r := srvView.Cmp(tarView); r != 0 {
		t.Errorf("source(%v) should be equal to target(%v): have %v, want %v", srvView, tarView, r, 0)
	}

	// test larger Sequence
	tarView = &View{
		Sequence: big.NewInt(1),
		Round:    big.NewInt(1),
	}
	if r := srvView.Cmp(tarView); r != 1 {
		t.Errorf("source(%v) should be larger than target(%v): have %v, want %v", srvView, tarView, r, 1)
	}

	// test larger Round
	tarView = &View{
		Sequence: big.NewInt(2),
		Round:    big.NewInt(0),
	}
	if r := srvView.Cmp(tarView); r != 1 {
		t.Errorf("source(%v) should be larger than target(%v): have %v, want %v", srvView, tarView, r, 1)
	}

	// test smaller Sequence
	tarView = &View{
		Sequence: big.NewInt(3),
		Round:    big.NewInt(1),
	}
	if r := srvView.Cmp(tarView); r != -1 {
		t.Errorf("source(%v) should be smaller than target(%v): have %v, want %v", srvView, tarView, r, -1)
	}
	tarView = &View{
		Sequence: big.NewInt(2),
		Round:    big.NewInt(2),
	}
	if r := srvView.Cmp(tarView); r != -1 {
		t.Errorf("source(%v) should be smaller than target(%v): have %v, want %v", srvView, tarView, r, -1)
	}
}

func TestPreprepareEncodeDecode(t *testing.T) {
	valSet := newTestValidatorSet(4)
	view := &View{
		Round:    big.NewInt(1),
		Sequence: big.NewInt(5),
	}

	proposal := makeBlock(5)
	preprepareWithPB := &PreprepareWithPiggybackMsgs{
		Preprepare: &Preprepare{
			View:     view,
			Proposal: proposal,
		},
		PiggybackMessages: &PiggybackMessages{
			RCMessages: newMessageSet(valSet),
		},
	}

	rawPreprepare, err := rlp.EncodeToBytes(preprepareWithPB)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	// decode preprepare message
	msg := &message{
		Code:    msgPreprepare,
		Msg:     rawPreprepare,
		Address: common.Address{},
	}
	var decPreprepare *PreprepareWithPiggybackMsgs
	err = msg.Decode(&decPreprepare)
	if err != nil {
		t.Errorf("error decoding preprepare message: %v", err)
	}

	if decPreprepare.Preprepare.Proposal.Hash() != proposal.Hash() {
		t.Errorf("error mismatch proposal hash: have %v, want %v", decPreprepare.Preprepare.Proposal.Hash(), proposal.Hash())
	}

}

func TestRCEncodeDeocdeRLP(t *testing.T) {
	view := &View{
		Round:    big.NewInt(1),
		Sequence: big.NewInt(5),
	}
	rc := &RoundChangePiggybackMsgs{
		RoundChangeMessage: &RoundChangeMessage{
			View:          view,
			PreparedRound: big.NewInt(0),
			PreparedBlock: makeBlock(5),
		},
	}
	rawRC, err := rlp.EncodeToBytes(rc)
	if err != nil {
		t.Errorf("error mismatch: have %v, want nil", err)
	}

	// decode roundchange message
	msg := &message{
		Code:    msgRoundChange,
		Msg:     rawRC,
		Address: common.Address{},
	}
	var decRC *RoundChangePiggybackMsgs
	err = msg.Decode(&decRC)
	if err != nil {
		t.Errorf("error decoding roundchange message: %v", err)
	}
	if decRC.RoundChangeMessage.View.Round.Uint64() != view.Round.Uint64() {
		t.Errorf("error mismatch view: have %v, want %v", decRC.RoundChangeMessage.View.Round.Uint64(), view.Round.Uint64())
	}
}
