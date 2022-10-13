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

package types

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

var (
	// IstanbulDigest represents a hash of "Istanbul practical byzantine fault tolerance"
	// to identify whether the block is from Istanbul consensus engine
	IstanbulDigest = common.HexToHash("0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365")

	IstanbulExtraVanity = 32 // Fixed number of extra-data bytes reserved for validator vanity
	IstanbulExtraSeal   = 65 // Fixed number of extra-data bytes reserved for validator seal

	QBFTAuthVote = byte(0xFF) // Magic number to vote on adding a new validator
	QBFTDropVote = byte(0x00) // Magic number to vote on removing a validator.

	// ErrInvalidIstanbulHeaderExtra is returned if the length of extra-data is less than 32 bytes
	ErrInvalidIstanbulHeaderExtra = errors.New("invalid istanbul header extra-data")
)

// IstanbulExtra represents the legacy IBFT header extradata
type IstanbulExtra struct {
	Validators    []common.Address
	Seal          []byte
	CommittedSeal [][]byte
}

// EncodeRLP serializes ist into the Ethereum RLP format.
func (ist *IstanbulExtra) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{
		ist.Validators,
		ist.Seal,
		ist.CommittedSeal,
	})
}

// DecodeRLP implements rlp.Decoder, and load the istanbul fields from a RLP stream.
func (ist *IstanbulExtra) DecodeRLP(s *rlp.Stream) error {
	var istanbulExtra struct {
		Validators    []common.Address
		Seal          []byte
		CommittedSeal [][]byte
	}
	if err := s.Decode(&istanbulExtra); err != nil {
		return err
	}
	ist.Validators, ist.Seal, ist.CommittedSeal = istanbulExtra.Validators, istanbulExtra.Seal, istanbulExtra.CommittedSeal
	return nil
}

// ExtractIstanbulExtra extracts all values of the IstanbulExtra from the header. It returns an
// error if the length of the given extra-data is less than 32 bytes or the extra-data can not
// be decoded.
func ExtractIstanbulExtra(h *Header) (*IstanbulExtra, error) {
	if len(h.Extra) < IstanbulExtraVanity {
		return nil, ErrInvalidIstanbulHeaderExtra
	}

	var istanbulExtra *IstanbulExtra
	err := rlp.DecodeBytes(h.Extra[IstanbulExtraVanity:], &istanbulExtra)
	if err != nil {
		return nil, err
	}
	return istanbulExtra, nil
}

// FilteredHeader returns a filtered header which some information (like seal, committed seals)
// are clean to fulfill the Istanbul hash rules. It first check if the extradata can be extracted into IstanbulExtra if that fails,
// it extracts extradata into QBFTExtra struct
func FilteredHeader(h *Header) *Header {
	// Check if the header extradata can be decoded in IstanbulExtra, if yes, then call IstanbulFilteredHeader()
	// if not then call QBFTFilteredHeader()
	_, err := ExtractIstanbulExtra(h)
	if err != nil {
		return QBFTFilteredHeader(h)
	}
	return IstanbulFilteredHeader(h, true)
}

// IstanbulFilteredHeader returns a filtered header which some information (like seal, committed seals)
// are clean to fulfill the Istanbul hash rules. It returns nil if the extra-data cannot be
// decoded/encoded by rlp.
func IstanbulFilteredHeader(h *Header, keepSeal bool) *Header {
	newHeader := CopyHeader(h)
	istanbulExtra, err := ExtractIstanbulExtra(newHeader)
	if err != nil {
		return nil
	}

	if !keepSeal {
		istanbulExtra.Seal = []byte{}
	}
	istanbulExtra.CommittedSeal = [][]byte{}

	payload, err := rlp.EncodeToBytes(&istanbulExtra)
	if err != nil {
		return nil
	}

	newHeader.Extra = append(newHeader.Extra[:IstanbulExtraVanity], payload...)

	return newHeader
}

// QBFTExtra represents header extradata for qbft protocol
type QBFTExtra struct {
	VanityData    []byte
	Validators    []common.Address
	Vote          *ValidatorVote `rlp:"nilString"`
	Round         []byte
	CommittedSeal [][]byte
}

type qbftExtraFallback struct {
	VanityData    []byte
	Validators    []common.Address
	Vote          *ValidatorVote `rlp:"nil"`
	Round         uint32
	CommittedSeal [][]byte
}

type QBFTExtraNoSealsNoRound struct {
	VanityData []byte
	Validators []common.Address
	Vote       *ValidatorVote `rlp:"nilString"`
}

type QBFTExtraNoSeals struct {
	VanityData []byte
	Validators []common.Address
	Vote       *ValidatorVote `rlp:"nilString"`
	Round      []byte
}

type ValidatorVote struct {
	RecipientAddress common.Address
	VoteType         byte
}

// EncodeRLP serializes qist into the Ethereum RLP format.
func (qst *QBFTExtra) EncodeRLP(w io.Writer) error {
	if qst.Vote == nil {
		return rlp.Encode(w, []interface{}{
			qst.VanityData,
			qst.Validators,
			make([]byte, 0),
			qst.Round,
			qst.CommittedSeal,
		})
	}
	return rlp.Encode(w, []interface{}{
		qst.VanityData,
		qst.Validators,
		qst.Vote,
		qst.Round,
		qst.CommittedSeal,
	})
}

func (qst *QBFTExtraNoSeals) EncodeRLP(w io.Writer) error {
	if qst.Vote == nil {
		return rlp.Encode(w, []interface{}{
			qst.VanityData,
			qst.Validators,
			make([]byte, 0),
			qst.Round,
		})
	}
	return rlp.Encode(w, []interface{}{
		qst.VanityData,
		qst.Validators,
		qst.Vote,
		qst.Round,
	})
}

func (qst *QBFTExtraNoSealsNoRound) EncodeRLP(w io.Writer) error {
	if qst.Vote == nil {
		return rlp.Encode(w, []interface{}{
			qst.VanityData,
			qst.Validators,
			make([]byte, 0),
		})
	}
	return rlp.Encode(w, []interface{}{
		qst.VanityData,
		qst.Validators,
		qst.Vote,
	})
}

// DecodeRLP implements rlp.Decoder, and load the QBFTExtra fields from a RLP stream.
func (qst *QBFTExtra) DecodeRLP(s *rlp.Stream) error {
	var qbftExtra struct {
		VanityData    []byte
		Validators    []common.Address
		Vote          *ValidatorVote `rlp:"nilString"`
		Round         []byte
		CommittedSeal [][]byte
	}
	if err := s.Decode(&qbftExtra); err != nil {
		var qbftExtraFallback struct {
			VanityData    []byte
			Validators    []common.Address
			Vote          *ValidatorVote `rlp:"nilString"`
			Round         uint32
			CommittedSeal [][]byte
		}
		if err := s.Decode(&qbftExtraFallback); err != nil {
			return err
		}
		round := make([]byte, 4)
		binary.BigEndian.PutUint32(round, qbftExtraFallback.Round)
		qst.VanityData, qst.Validators, qst.Vote, qst.Round, qst.CommittedSeal = qbftExtraFallback.VanityData, qbftExtraFallback.Validators, qbftExtraFallback.Vote, round, qbftExtraFallback.CommittedSeal
		return nil
	}
	qst.VanityData, qst.Validators, qst.Vote, qst.Round, qst.CommittedSeal = qbftExtra.VanityData, qbftExtra.Validators, qbftExtra.Vote, qbftExtra.Round, qbftExtra.CommittedSeal

	return nil
}

// EncodeRLP serializes ValidatorVote into the Ethereum RLP format.
func (vv *ValidatorVote) EncodeRLP(w io.Writer) error {
	if vv.VoteType == 0 {
		// It's easier to process this edge case in such a way in order not to introduce extra logic to RLP module for only this case
		_, err := w.Write(append(append([]byte{0xd6, 0x94}, vv.RecipientAddress.Bytes()...), vv.VoteType))
		return err
	}
	return rlp.Encode(w, []interface{}{
		vv.RecipientAddress,
		vv.VoteType,
	})
}

// DecodeRLP implements rlp.Decoder, and load the ValidatorVote fields from a RLP stream.
func (vv *ValidatorVote) DecodeRLP(s *rlp.Stream) error {
	var validatorVote struct {
		RecipientAddress common.Address
		VoteType         byte
	}
	s.IgnoreCanonIntForByte = true
	if err := s.Decode(&validatorVote); err != nil {
		return err
	}
	s.IgnoreCanonIntForByte = false
	vv.RecipientAddress, vv.VoteType = validatorVote.RecipientAddress, validatorVote.VoteType
	return nil
}

// ExtractQBFTExtra extracts all values of the QBFTExtra from the header. It returns an
// error if the length of the given extra-data is less than 32 bytes or the extra-data can not
// be decoded.
func ExtractQBFTExtra(h *Header) (*QBFTExtra, error) {
	qbftExtra := new(QBFTExtra)
	err := rlp.DecodeBytes(h.Extra[:], qbftExtra)
	if err != nil {
		qbftExtraFallback := new(qbftExtraFallback)
		errFallback := rlp.DecodeBytes(h.Extra[:], qbftExtraFallback)
		if errFallback != nil {
			return nil, fmt.Errorf("got two decoding errors: 1=%w; 2=%v", err, errFallback)
		}
		log.Warn("qbft extra fallback to old version")
		round := make([]byte, 0)
		if qbftExtraFallback.Round > 0 {
			round = make([]byte, 4)
			binary.BigEndian.PutUint32(round, qbftExtraFallback.Round)
		}
		qbftExtra.VanityData, qbftExtra.Validators, qbftExtra.Vote, qbftExtra.Round, qbftExtra.CommittedSeal = qbftExtraFallback.VanityData, qbftExtraFallback.Validators, qbftExtraFallback.Vote, round, qbftExtraFallback.CommittedSeal
	}
	return qbftExtra, nil
}

// QBFTFilteredHeader returns a filtered header which some information (like committed seals, round, validator vote)
// are clean to fulfill the Istanbul hash rules. It returns nil if the extra-data cannot be
// decoded/encoded by rlp.
func QBFTFilteredHeader(h *Header) *Header {
	newHeader := CopyHeader(h)
	qbftExtra, err := ExtractQBFTExtra(newHeader)
	if err != nil {
		return nil
	}

	qbftFiltered := new(QBFTExtraNoSealsNoRound)

	qbftFiltered.Validators, qbftFiltered.VanityData, qbftFiltered.Vote = qbftExtra.Validators, qbftExtra.VanityData, qbftExtra.Vote

	payload, err := rlp.EncodeToBytes(&qbftFiltered)
	if err != nil {
		return nil
	}

	newHeader.Extra = payload

	return newHeader
}

// QBFTFilteredHeaderWithRound returns the copy of the header with round number set to the given round number
// and commit seal set to its null value
func QBFTFilteredHeaderWithRound(h *Header, round uint32) *Header {
	newHeader := CopyHeader(h)
	qbftExtra, err := ExtractQBFTExtra(newHeader)
	if err != nil {
		return nil
	}

	qbftFiltered := new(QBFTExtraNoSeals)

	qbftFiltered.Validators, qbftFiltered.VanityData, qbftFiltered.Vote = qbftExtra.Validators, qbftExtra.VanityData, qbftExtra.Vote

	a := make([]byte, 4)
	binary.BigEndian.PutUint32(a, round)

	qbftFiltered.Round = a

	payload, err := rlp.EncodeToBytes(&qbftFiltered)
	if err != nil {
		return nil
	}

	newHeader.Extra = payload

	return newHeader
}
