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

const (
	msgPreprepare uint64 = iota
	msgPrepare
	msgCommit
	msgRoundChange
)

type message struct {
	Code          uint64
	Msg           []byte
	Address       common.Address
	Signature     []byte
	CommittedSeal []byte
}

// ==============================================
//
// define the functions that needs to be provided for rlp Encoder/Decoder.

// EncodeRLP serializes m into the Ethereum RLP format.
func (m *message) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{m.Code, m.Msg, m.Address, m.Signature, m.CommittedSeal})
}

// DecodeRLP implements rlp.Decoder, and load the consensus fields from a RLP stream.
func (m *message) DecodeRLP(s *rlp.Stream) error {
	var msg struct {
		Code          uint64
		Msg           []byte
		Address       common.Address
		Signature     []byte
		CommittedSeal []byte
	}

	if err := s.Decode(&msg); err != nil {
		return err
	}
	m.Code, m.Msg, m.Address, m.Signature, m.CommittedSeal = msg.Code, msg.Msg, msg.Address, msg.Signature, msg.CommittedSeal
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

	// Validate message (on a message without Signature)
	if validateFn != nil {
		var payload []byte
		payload, err = m.PayloadNoSig()
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
	RCMessages      *messageSet
	PrepareMessages *messageSet
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
	View             *View
	Proposal         istanbul.Proposal
	RCMessages       *messageSet
	PreparedMessages *messageSet
}

// EncodeRLP serializes b into the Ethereum RLP format.
func (b *Preprepare) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{b.View, b.Proposal, b.RCMessages, b.PreparedMessages})
}

// DecodeRLP implements rlp.Decoder, and load the consensus fields from a RLP stream.
func (b *Preprepare) DecodeRLP(s *rlp.Stream) error {
	var preprepare struct {
		View             *View
		Proposal         *types.Block
		RCMessages       *messageSet
		PreparedMessages *messageSet
	}

	if err := s.Decode(&preprepare); err != nil {
		return err
	}
	b.View, b.Proposal, b.RCMessages, b.PreparedMessages = preprepare.View, preprepare.Proposal, preprepare.RCMessages, preprepare.PreparedMessages

	return nil
}

// Subject represents the message sent when msgPrepare and msgCommit is broadcasted
type Subject struct {
	View   *View
	Digest istanbul.Proposal
}

// EncodeRLP serializes b into the Ethereum RLP format.
func (b *Subject) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{b.View, b.Digest})
}

// DecodeRLP implements rlp.Decoder, and load the consensus fields from a RLP stream.
func (b *Subject) DecodeRLP(s *rlp.Stream) error {
	var subject struct {
		View   *View
		Digest *types.Block
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
	View             *View
	PreparedRound    *big.Int
	PreparedBlock    istanbul.Proposal
	PreparedMessages *messageSet
}

func (r *RoundChangeMessage) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{r.View, r.PreparedRound, r.PreparedBlock, r.PreparedMessages})
}

func (r *RoundChangeMessage) DecodeRLP(s *rlp.Stream) error {
	var rcMessage struct {
		View             *View
		PreparedRound    *big.Int
		PreparedBlock    *types.Block
		PreparedMessages *messageSet
	}

	if err := s.Decode(&rcMessage); err != nil {
		return err
	}
	r.View, r.PreparedRound, r.PreparedBlock, r.PreparedMessages = rcMessage.View, rcMessage.PreparedRound, rcMessage.PreparedBlock, rcMessage.PreparedMessages
	return nil
}
