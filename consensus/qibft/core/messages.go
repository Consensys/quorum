package core

import (
	"bytes"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// QBFT message_deprecated codes
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
	View() View
	Source() common.Address
	SetSource(address common.Address)
	EncodePayload() ([]byte, error)
	Signature() []byte
	SetSignature(signature []byte)
}

// QBFT Messages
type CommonPayload struct {
	code uint64
	source common.Address
	Sequence *big.Int
	Round *big.Int
	signature []byte
}

func (m *CommonPayload) Code() uint64 {
	return m.code
}

func (m *CommonPayload) Source() common.Address {
	return m.source
}

func (m *CommonPayload) SetSource(address common.Address) {
	m.source = address
}

func (m *CommonPayload) View() View {
	return View{Sequence: m.Sequence, Round: m.Round}
}

func (m *CommonPayload) Signature() []byte {
	return m.signature
}

func (m *CommonPayload) SetSignature(signature []byte) {
	m.signature = signature
}

type CommitMsg struct {
	CommonPayload
	Digest     common.Hash
	CommitSeal []byte
}

func (m *CommitMsg) EncodePayload() ([]byte, error) {
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
			m.signature})
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
	m.signature = message.Signature
	return nil
}

// RLP
func DecodeMessage(code uint64, data []byte) (QBFTMessage, error){
	switch code {
	case preprepareMsgCode:
		var preprepare PreprepareMsg
		if err := rlp.DecodeBytes(data, &preprepare); err != nil {
			return nil, errFailedDecodePreprepare
		}
		preprepare.code = preprepareMsgCode
		return &preprepare, nil
	case prepareMsgCode:
		var prepare PrepareMsg
		if err := rlp.DecodeBytes(data, &prepare); err != nil {
			return nil, errFailedDecodeCommit
		}
		prepare.code = prepareMsgCode
		return &prepare, nil
	case commitMsgCode:
		var commit CommitMsg
		if err := rlp.DecodeBytes(data, &commit); err != nil {
			return nil, errFailedDecodeCommit
		}
		commit.code = commitMsgCode
		return &commit, nil
	case roundChangeMsgCode:
		var roundChange RoundChangeMsg
		if err := rlp.DecodeBytes(data, &roundChange); err != nil {
			return nil, errFailedDecodeRoundChange
		}
		roundChange.code = roundChangeMsgCode
		return &roundChange, nil
	}

	return nil, errInvalidMessage
}

// ROUND-CHANGE
type RoundChangeMsg struct {
	SignedRoundChangePayload
	PreparedBlock *types.Block
	Justification []*SignedPreparePayload
}

type SignedRoundChangePayload struct {
	CommonPayload
	PreparedRound  *big.Int
	PreparedDigest common.Hash
}

func (p *SignedRoundChangePayload) EncodeRLP(w io.Writer) error {
	var encodedPayload rlp.RawValue
	encodedPayload, err := p.EncodePayload()
	if err != nil {
		return err
	}

	return rlp.Encode(
		w,
		[]interface{}{encodedPayload, p.signature})

}

func (p *SignedRoundChangePayload) DecodeRLP(stream *rlp.Stream) error {
	// Signed Payload
	if _, err := stream.List(); err != nil {
		log.Error("QBFT: Error List() Signed Payload", "err", err)
		return err
	}

	// Payload
	encodedPayload, err := stream.Raw()
	if err != nil {
		log.Error("QBFT: Error Raw()", "err", err)
		return err
	}

	payloadStream := rlp.NewStream(bytes.NewReader(encodedPayload), 0)

	if _, err = payloadStream.List(); err != nil {
		log.Error("QBFT: Error List() Payload", "err", err)
		return err
	}

	if err = payloadStream.Decode(&p.Sequence); err != nil {
		log.Error("QBFT: Error Decode(&m.Sequence)", "err", err)
		return err
	}
	if err = payloadStream.Decode(&p.Round); err != nil {
		log.Error("QBFT: Error Decode(&m.Round)", "err", err)
		return err
	}

	// Prepared
	var size uint64
	if size, err = payloadStream.List(); err != nil {
		log.Error("QBFT: Error List() Prepared", "err", err)
		return err
	}
	if size > 0 {
		if err = payloadStream.Decode(&p.PreparedRound); err != nil {
			log.Error("QBFT: Error Decode(&m.PreparedRound)", "err", err)
			return err
		}
		if err = payloadStream.Decode(&p.PreparedDigest); err != nil {
			log.Error("QBFT: Error Decode(&p.PreparedDigest)", "err", err)
			return err
		}
	}
	// End Prepared
	if err = payloadStream.ListEnd(); err != nil {
		return err
	}

	// End Payload
	if err = payloadStream.ListEnd(); err != nil {
		return err
	}

	if err = stream.Decode(&p.signature); err != nil {
		return err
	}
	// End SignedPayload
	if err = stream.ListEnd(); err != nil {
		return err
	}

	p.code = roundChangeMsgCode

	log.Info("QBFT: Correctly decoded SignedRoundChangePayload", "p", p)

	return nil
}


func (p *SignedRoundChangePayload) EncodePayload() ([]byte, error) {
	var prepared = []interface{}{}
	if p.PreparedRound != nil && !p.PreparedDigest.IsEmpty() {
		prepared = []interface{}{p.PreparedRound, p.PreparedDigest}
	}
	return rlp.EncodeToBytes(
			[]interface{}{
				p.Sequence,
				p.Round,
				prepared})
}

func (m *RoundChangeMsg) EncodeRLP(w io.Writer) error {
	var prepared = []interface{}{}
	if m.PreparedRound != nil && !m.PreparedDigest.IsEmpty() {
		prepared = []interface{}{m.PreparedRound, m.PreparedDigest}
	}

	return rlp.Encode(
		w,
		[]interface{}{
			[]interface{}{
				[]interface{}{m.Sequence, m.Round, prepared},
				m.signature,
			},
			m.PreparedBlock, m.Justification,
		})
}

func (m *RoundChangeMsg) DecodeRLP(stream *rlp.Stream) error {
	var err error

	// RoundChange Message
	if _, err = stream.List(); err != nil {
		return err
	}

	// Signed Payload
	if _, err = stream.List(); err != nil {
		log.Error("QBFT: Error List() Signed Payload", "err", err)
		return err
	}

	// Payload
	encodedPayload, err := stream.Raw()
	if err != nil {
		log.Error("QBFT: Error Raw()", "err", err)
		return err
	}

	payloadStream := rlp.NewStream(bytes.NewReader(encodedPayload), 0)

	if _, err = payloadStream.List(); err != nil {
		log.Error("QBFT: Error List() Payload", "err", err)
		return err
	}

	if err = payloadStream.Decode(&m.Sequence); err != nil {
		log.Error("QBFT: Error Decode(&m.Sequence)", "err", err)
		return err
	}
	if err = payloadStream.Decode(&m.Round); err != nil {
		log.Error("QBFT: Error Decode(&m.Round)", "err", err)
		return err
	}

	// Prepared
	var size uint64
	if size, err = payloadStream.List(); err != nil {
		log.Error("QBFT: Error List() Prepared", "err", err)
		return err
	}
	if size > 0 {
		if err = payloadStream.Decode(&m.PreparedRound); err != nil {
			log.Error("QBFT: Error Decode(&m.PreparedRound)", "err", err)
			return err
		}
		if err = payloadStream.Decode(&m.PreparedDigest); err != nil {
			log.Error("QBFT: Error Decode(&m.PreparedDigest)", "err", err)
			return err
		}
	}
	// End Prepared
	if err = payloadStream.ListEnd(); err != nil {
		return err
	}

	// End Payload
	if err = payloadStream.ListEnd(); err != nil {
		return err
	}

	if err = stream.Decode(&m.signature); err != nil {
		return err
	}
	// End SignedPayload
	if err = stream.ListEnd(); err != nil {
		return err
	}

	if _, size, err = stream.Kind(); err != nil {
		log.Error("QBFT: Error Kind()", "err", err)
		return err
	}
	if size == 0 {
		if _, err = stream.Raw(); err != nil {
			log.Error("QBFT: Error Raw()", "err", err)
			return err
		}
	} else {
		if err = stream.Decode(&m.PreparedBlock); err != nil {
			log.Error("QBFT: Error Decode(&m.PreparedDigest)", "err", err)
			return err
		}
		if m.PreparedBlock.Hash() != m.PreparedDigest {
			log.Error("QBFT: Error m.PreparedDigest.Hash() != digest")
			return errFailedDecodeRoundChange
		}
	}

	if _, size, err = stream.Kind(); err != nil {
		log.Error("QBFT: Error Kind()", "err", err)
		return err
	}
	if size == 0 {
		if _, err = stream.Raw(); err != nil {
			log.Error("QBFT: Error Raw()", "err", err)
			return err
		}
	} else {
		if err = stream.Decode(&m.Justification); err != nil {
			log.Error("QBFT: Error Decode(&m.Justification)", "err", err)
			return err
		}
	}

	// End RoundChange Message
	if err = stream.ListEnd(); err != nil {
		return err
	}

	return nil
}


type PreprepareMsg struct {
	CommonPayload
	Proposal istanbul.Proposal
	JustificationRoundChanges []*SignedRoundChangePayload
	JustificationPrepares []*SignedPreparePayload
}

func (m *PreprepareMsg) EncodePayload() ([]byte, error) {
	return rlp.EncodeToBytes(
		[]interface{}{m.Sequence, m.Round, m.Proposal})
}

func (m *PreprepareMsg) EncodeRLP(w io.Writer) error {
	return rlp.Encode(
		w,
		[]interface{}{
			[]interface{}{
				[]interface{}{m.Sequence, m.Round, m.Proposal},
				m.signature,
			},
			[]interface{}{
				m.JustificationRoundChanges,
				m.JustificationPrepares,
			},
		})
}

func (m *PreprepareMsg) DecodeRLP(stream *rlp.Stream) error {
	var message struct {
		SignedPayload struct {
			Payload struct {
				Sequence *big.Int
				Round *big.Int
				Proposal *types.Block
			}
			Signature []byte
		}
		Justification struct {
			RoundChanges []*SignedRoundChangePayload
			Prepares []*SignedPreparePayload
		}
	}
	if err := stream.Decode(&message); err != nil {
		return err
	}
	m.Sequence = message.SignedPayload.Payload.Sequence
	m.Round = message.SignedPayload.Payload.Round
	m.Proposal = message.SignedPayload.Payload.Proposal
	m.signature = message.SignedPayload.Signature
	m.JustificationPrepares = message.Justification.Prepares
	m.JustificationRoundChanges = message.Justification.RoundChanges
	return nil
}
/*
func (m *PreprepareMsg) DecodeRLP(stream *rlp.Stream) error {
	if _, err := stream.List(); err != nil {
		log.Error("QBFT: Error List()", "err", err)
		return err
	}

	encodedPayload, err := stream.Raw()
	if err != nil {
		log.Error("QBFT: Error Raw()", "err", err)
		return err
	}

	signature, err := stream.Bytes()
	if err != nil {
		log.Error("QBFT: Error Bytes()", "err", err)
		return err
	}
	m.signature = signature


	var payload struct {
		Sequence   *big.Int
		Round      *big.Int
		Proposal   *types.Block
	}
	if err = rlp.DecodeBytes(encodedPayload, &payload); err != nil {
		log.Error("QBFT: Error DecodeBytes()", "err", err)
		return err
	}

	m.Sequence = payload.Sequence
	m.Round = payload.Round
	m.Proposal = payload.Proposal

	return stream.ListEnd()
}*/

type PrepareMsgOld struct {
	CommonPayload
	Digest     common.Hash
}

type PrepareMsg struct {
	SignedPreparePayload
}

func (m *PrepareMsg) EncodePayload() ([]byte, error) {
	return rlp.EncodeToBytes([]interface{}{m.Sequence, m.Round, m.Digest})
}

/*
func (m *PrepareMsg) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, m.SignedPreparePayload)
}

func (m *PrepareMsg) DecodeRLP(stream *rlp.Stream) error {
	m.code = prepareMsgCode
	return stream.Decode(m.SignedPreparePayload)
}
*/
type SignedPreparePayload struct {
	CommonPayload
	Digest common.Hash
}

func (signedPayload *SignedPreparePayload) EncodeRLP(w io.Writer) error {
	return rlp.Encode(
		w,
		[]interface{}{
			[]interface{}{
				signedPayload.Sequence,
				signedPayload.Round,
				signedPayload.Digest},
			signedPayload.signature,
		})
}

func (signedPayload *SignedPreparePayload) DecodeRLP(stream *rlp.Stream) error {
	var message struct {
		Payload struct {
			Sequence   *big.Int
			Round      *big.Int
			Digest     common.Hash
		}
		Signature []byte
	}
	if err := stream.Decode(&message); err != nil {
		return err
	}
	signedPayload.code = prepareMsgCode
	signedPayload.Sequence = message.Payload.Sequence
	signedPayload.Round = message.Payload.Round
	signedPayload.Digest = message.Payload.Digest
	signedPayload.signature = message.Signature
	return nil
}


func (m *PrepareMsgOld) EncodePayload() ([]byte, error) {
	return rlp.EncodeToBytes([]interface{}{m.Sequence, m.Round, m.Digest})
}

func (m *PrepareMsgOld) EncodeSignedPayload() ([]byte, error) {
	return rlp.EncodeToBytes(
		[]interface{}{
			[]interface{}{m.Sequence, m.Round, m.Digest},
			m.signature,
		})
}

func (m *PrepareMsgOld) decodePayload(stream *rlp.Stream) error {
	var payload struct {
		Sequence   *big.Int
		Round      *big.Int
		Digest     common.Hash
	}
	if err := stream.Decode(&payload); err != nil {
		return err
	}
	m.Sequence = payload.Sequence
	m.Round = payload.Round
	m.Digest = payload.Digest
	return nil
}

func (m *PrepareMsgOld) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{
		[]interface{}{m.Sequence, m.Round, m.Digest},
		m.signature})
}

func (m *PrepareMsgOld) DecodeRLP(stream *rlp.Stream) error {
	var message struct {
		Payload struct {
			Sequence   *big.Int
			Round      *big.Int
			Digest     common.Hash
		}
		Signature []byte
	}
	if err := stream.Decode(&message); err != nil {
		return err
	}
	m.Sequence = message.Payload.Sequence
	m.Round = message.Payload.Round
	m.Digest = message.Payload.Digest
	m.signature = message.Signature
	return nil
}
