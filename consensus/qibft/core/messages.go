package core

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/log"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/common"
)

// QBFT message codes
const (
	preprepareMsgCode  = 0x12
	prepareMsgCode     = 0x13
	commitMsgCode      = 0x14
	roundChangeMsgCode = 0x15
)

func MessageCodes() map[uint64]struct{} {
	return map[uint64]struct{}{
		preprepareMsgCode:  {},
		prepareMsgCode:     {},
		commitMsgCode:      {},
		roundChangeMsgCode: {},
	}
}

// Common interface for QBFT Messages
type QBFTMessage interface {
	Code() uint64
	Source() common.Address
	View() View
	SignedPayload() SignedPayload
}

// The RLP-encoded payload of a message and its respective signature.
type SignedPayload struct {
	EncodedPayload []byte
	Signature      []byte
}

func SignPayload(payload []byte, privateKey *ecdsa.PrivateKey) (*SignedPayload, error) {
	signedPayload := &SignedPayload{
		EncodedPayload: payload,
		Signature:      nil,
	}
	return signedPayload, nil
}

// QBFT Messages
type CommonMsg struct {
	code uint64
	source common.Address
	Sequence *big.Int
	Round *big.Int
	EncodedPayload []byte
	Signature []byte
}

func (m *CommonMsg) Code() uint64 {
	return m.code
}

func (m *CommonMsg) Source() common.Address {
	return m.source
}

func (m *CommonMsg) View() View {
	return View{Sequence: m.Sequence, Round: m.Round}
}

func (m *CommonMsg) SignedPayload() SignedPayload {
	return SignedPayload{EncodedPayload: m.EncodedPayload, Signature: m.Signature}
}

type CommitMsg struct {
	CommonMsg
	//View
	Digest     common.Hash
	CommitSeal []byte
//	source     common.Address
//	SignedPayload
}

func (m *CommitMsg) EncodedPayload() ([]byte, error) {
	return rlp.EncodeToBytes([]interface{}{m.Sequence, m.Round, m.Digest, m.CommitSeal})
}

func (m *CommitMsg) decodePayload(stream *rlp.Stream) error {
	var payload struct {
		Sequence   *big.Int
		Round      *big.Int
		Digest     common.Hash
		CommitSeal []byte
	}
	if err := stream.Decode(&payload); err != nil {
		return err
	}
	m.Sequence = payload.Sequence
	m.Round = payload.Round
	m.Digest = payload.Digest
	m.CommitSeal = payload.CommitSeal
	return nil
}

func (m *CommitMsg) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{
			[]interface{}{m.Sequence, m.Round, m.Digest, m.CommitSeal},
			m.Signature})
}

func (m *CommitMsg) DecodeRLP(stream *rlp.Stream) error {
	var message struct {
		Payload struct {
			Sequence   *big.Int
			Round      *big.Int
			Digest     common.Hash
			CommitSeal []byte
		}
		Signature []byte
	}
	if err := stream.Decode(&message); err != nil {
		return err
	}
	m.Sequence = message.Payload.Sequence
	m.Round = message.Payload.Round
	m.Digest = message.Payload.Digest
	m.CommitSeal = message.Payload.CommitSeal
	m.Signature = message.Signature
	return nil
}

// RLP
func DecodeQBFTMessage(code uint64, data []byte) (QBFTMessage, error){
	switch code {
	case commitMsgCode:
		var commit CommitMsg
		if err := rlp.DecodeBytes(data, &commit); err != nil {
			log.Error("QBFT: Error decoding message", "code", code)
			return nil, errFailedDecodeCommit
		}
		commit.code = commitMsgCode
		return &commit, nil
	}
	return nil, errInvalidMessage
}