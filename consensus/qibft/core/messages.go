package core

import (
	"crypto/ecdsa"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/common"
)

// QBFT message codes
const (
	preprepareMsgCode  = 0x81
	prepareMsgCode     = 0x82
	commitMsgCode      = 0x83
	roundChangeMsgCode = 0x84
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
	GetView() *View
	Source() common.Address
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
type CommitMsg struct {
	View
	Digest     common.Hash
	CommitSeal []byte
	source     common.Address
	SignedPayload
}

func (m *CommitMsg) Source() common.Address {
	return m.source
}

func (m *CommitMsg) GetView() *View {
	return &m.View
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
		[]interface{}{m.Sequence, m.Round, m.Digest},
		m.Signature})
}

func (m *CommitMsg) DecodeRLP(stream *rlp.Stream) error {
	var err error

	if _, err = stream.List(); err != nil {
		return err
	}

	m.decodePayload(stream)

	if m.Signature, err = stream.Bytes(); err != nil {
		return err
	}

	if err = stream.ListEnd(); err != nil {
		return err
	}

	return nil
}
