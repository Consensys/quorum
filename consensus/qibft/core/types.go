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
	"bytes"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type Engine interface {
	Start() error
	Stop() error

	IsProposer() bool

	// verify if a hash is the same as the proposed block in the current pending request
	//
	// this is useful when the engine is currently the proposer
	//
	// pending request is populated right at the preprepare stage so this would give us the earliest verification
	// to avoid any race condition of coming propagated blocks
	IsCurrentProposal(blockHash common.Hash) bool
}

type State uint64

const (
	StateAcceptRequest State = iota
	StatePreprepared
	StatePrepared
	StateCommitted
)

func (s State) String() string {
	if s == StateAcceptRequest {
		return "Accept request"
	} else if s == StatePreprepared {
		return "Preprepared"
	} else if s == StatePrepared {
		return "Prepared"
	} else if s == StateCommitted {
		return "Committed"
	} else {
		return "Unknown"
	}
}

// Cmp compares s and y and returns:
//   -1 if s is the previous state of y
//    0 if s and y are the same state
//   +1 if s is the next state of y
func (s State) Cmp(y State) int {
	if uint64(s) < uint64(y) {
		return -1
	}
	if uint64(s) > uint64(y) {
		return 1
	}
	return 0
}


type message struct {
	Code          uint64
	Msg           []byte
	Address       common.Address
	Signature     []byte
	CommittedSeal []byte
	PiggybackMsgs []byte
}

// ==============================================
//
// define the functions that needs to be provided for rlp Encoder/Decoder.

// EncodeRLP serializes m into the Ethereum RLP format.
func (m *message) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{m.Code, m.Msg, m.Address, m.Signature, m.CommittedSeal, m.PiggybackMsgs})
}

// DecodeRLP implements rlp.Decoder, and load the consensus fields from a RLP stream.
func (m *message) DecodeRLP(s *rlp.Stream) error {
	var msg struct {
		Code          uint64
		Msg           []byte
		Address       common.Address
		Signature     []byte
		CommittedSeal []byte
		PiggybackMsgs []byte
	}

	if err := s.Decode(&msg); err != nil {
		return err
	}
	m.Code, m.Msg, m.Address, m.Signature, m.CommittedSeal, m.PiggybackMsgs = msg.Code, msg.Msg, msg.Address, msg.Signature, msg.CommittedSeal, msg.PiggybackMsgs
	return nil
}

// ==============================================
//
// define the functions that needs to be provided for core.

func (m *message) FromPayload(b []byte, validateFn func([]byte, []byte) (common.Address, error)) error {
	// Decode message
	err := rlp.DecodeBytes(b, &m)
	if err != nil {
		return err
	}

	if len(m.PiggybackMsgs) == 0 {
		m.PiggybackMsgs = nil
	}

	// Validate message (on a message without Signature)
	if validateFn != nil {
		// Verify messages signature
		if err = verifySignature(m, validateFn); err != nil {
			return err
		}
		// Verify Signature of piggyback messages
		/*if err = decodeAndVerifyPiggybackMsgs(m.PiggybackMsgs, validateFn); err != nil {
			return err
		}*/

	}
	return nil
}

/*
// decodeAndVerifyPiggybackMsgs decodes the given piggyback messages and verifies the signature of individual Prepare and Round Change messages
func decodeAndVerifyPiggybackMsgs(piggybackMsgsPayload []byte, validateFn func([]byte, []byte) (common.Address, error)) error {
	// First decode piggyback messages and then verify individual Prepare and Round Change messages
	if len(piggybackMsgsPayload) > 0 {
		var pbMsgs *PiggybackMessages
		err := rlp.DecodeBytes(piggybackMsgsPayload, &pbMsgs)
		if err != nil {
			return errFailedDecodePiggybackMsgs
		}
		if pbMsgs.PreparedMessages != nil && pbMsgs.PreparedMessages.messages != nil {
			if err = verifyPiggyBackMsgSignatures(pbMsgs.PreparedMessages.messages, validateFn); err != nil {
				return err
			}
		}
		if pbMsgs.RCMessages != nil && pbMsgs.RCMessages.messages != nil {
			if err = verifyPiggyBackMsgSignatures(pbMsgs.RCMessages.messages, validateFn); err != nil {
				return err
			}
		}
	}

	return nil
}*/

// verifySignature verifies the signature of the given message
func verifySignature(m *message, validateFn func([]byte, []byte) (common.Address, error)) error {
	var payload []byte
	payload, err := m.PayloadNoSig()
	if err != nil {
		return err
	}

	signerAdd, err := validateFn(payload, m.Signature)
	if err != nil {
		return err
	}
	if !bytes.Equal(signerAdd.Bytes(), m.Address.Bytes()) {
		return errInvalidSigner
	}
	return nil
}

// verifyPiggyBackMsgSignatures verifies signatures of piggyback messages which are sent as part of PRE-PREPARE or ROUNDCHANGE messages
func verifyPiggyBackMsgSignatures(messages map[common.Address]*message, validateFn func([]byte, []byte) (common.Address, error)) error {
	for _, msg := range messages {
		if err := verifySignature(msg, validateFn); err != nil {
			return err
		}
	}
	return nil
}

func (m *message) Payload() ([]byte, error) {
	return rlp.EncodeToBytes(m)
}

func (m *message) PayloadNoSig() ([]byte, error) {
	return rlp.EncodeToBytes(&message{
		Code:          m.Code,
		Msg:           m.Msg,
		Address:       m.Address,
		Signature:     []byte{},
		CommittedSeal: m.CommittedSeal,
	})
}

func (m *message) Decode(val interface{}) error {
	return rlp.DecodeBytes(m.Msg, val)
}

func (m *message) String() string {
	return fmt.Sprintf("{Code: %v, Address: %v}", m.Code, m.Address.String())
}

// ==============================================
//
// helper functions

func Encode(val interface{}) ([]byte, error) {
	return rlp.EncodeToBytes(val)
}

// Request is used to construct a Preprepare message
type Request struct {
	Proposal        istanbul.Proposal
	RCMessages      *qbftMsgSet
	PrepareMessages []*SignedPreparePayload
}

// View includes a round number and a sequence number.
// Sequence is the block number we'd like to commit.
// Each round has a number and is composed by 3 steps: preprepare, prepare and commit.
//
// If the given block is not accepted by validators, a round change will occur
// and the validators start a new round with round+1.
type View struct {
	Round    *big.Int
	Sequence *big.Int
}

// EncodeRLP serializes b into the Ethereum RLP format.
func (v *View) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{v.Round, v.Sequence})
}

// DecodeRLP implements rlp.Decoder, and load the consensus fields from a RLP stream.
func (v *View) DecodeRLP(s *rlp.Stream) error {
	var view struct {
		Round    *big.Int
		Sequence *big.Int
	}

	if err := s.Decode(&view); err != nil {
		return err
	}
	v.Round, v.Sequence = view.Round, view.Sequence
	return nil
}

func (v *View) String() string {
	return fmt.Sprintf("{Round: %d, Sequence: %d}", v.Round.Uint64(), v.Sequence.Uint64())
}

// Cmp compares v and y and returns:
//   -1 if v <  y
//    0 if v == y
//   +1 if v >  y
func (v *View) Cmp(y *View) int {
	if v.Sequence.Cmp(y.Sequence) != 0 {
		return v.Sequence.Cmp(y.Sequence)
	}
	if v.Round.Cmp(y.Round) != 0 {
		return v.Round.Cmp(y.Round)
	}
	return 0
}

// Preprepare represents the message sent, when msgPreprepare is broadcasted
type Preprepare struct {
	View     *View
	Proposal istanbul.Proposal
}

// EncodeRLP serializes b into the Ethereum RLP format.
func (b *Preprepare) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{b.View, b.Proposal})
}

// DecodeRLP implements rlp.Decoder, and load the consensus fields from a RLP stream.
func (b *Preprepare) DecodeRLP(s *rlp.Stream) error {
	var preprepare struct {
		View     *View
		Proposal *types.Block
	}

	if err := s.Decode(&preprepare); err != nil {
		return err
	}
	b.View, b.Proposal = preprepare.View, preprepare.Proposal

	return nil
}

// Subject represents the message sent when msgPrepare and msgCommit is broadcasted
type Subject struct {
	View   *View
	Digest common.Hash
}

// EncodeRLP serializes b into the Ethereum RLP format.
func (b *Subject) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{b.View, b.Digest})
}

// DecodeRLP implements rlp.Decoder, and load the consensus fields from a RLP stream.
func (b *Subject) DecodeRLP(s *rlp.Stream) error {
	var subject struct {
		View   *View
		Digest common.Hash
	}

	if err := s.Decode(&subject); err != nil {
		return err
	}
	b.View, b.Digest = subject.View, subject.Digest
	return nil
}

func (b *Subject) String() string {
	return fmt.Sprintf("{View: %v, Proposal: %v}", b.View, b.Digest.String())
}

// RoundChangeMessage represents the message sent when msgRoundChange is broadcasted
type RoundChangeMessage struct {
	View                *View
	PreparedRound       *big.Int
	PreparedBlockDigest common.Hash
}

func (r *RoundChangeMessage) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{r.View, r.PreparedRound, r.PreparedBlockDigest})
}

func (r *RoundChangeMessage) DecodeRLP(s *rlp.Stream) error {
	var rcMessage struct {
		View                *View
		PreparedRound       *big.Int
		PreparedBlockDigest common.Hash
	}

	if err := s.Decode(&rcMessage); err != nil {
		return err
	}
	r.View, r.PreparedRound, r.PreparedBlockDigest = rcMessage.View, rcMessage.PreparedRound, rcMessage.PreparedBlockDigest
	return nil
}

type PiggybackMessages struct {
	RCMessages       *qbftMsgSet
	PreparedMessages *messageSet
	Proposal         istanbul.Proposal
}

func (p *PiggybackMessages) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{p.RCMessages, p.PreparedMessages, p.Proposal})
}

func (p *PiggybackMessages) DecodeRLP(s *rlp.Stream) error {
	var piggybackMsgs struct {
		RCMessages       *qbftMsgSet
		PreparedMessages *messageSet
		Proposal         *types.Block
	}

	if err := s.Decode(&piggybackMsgs); err != nil {
		return err
	}
	p.RCMessages, p.PreparedMessages, p.Proposal = piggybackMsgs.RCMessages, piggybackMsgs.PreparedMessages, piggybackMsgs.Proposal

	return nil
}
