package message

import (
	"github.com/ethereum/go-ethereum/consensus/qibft"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
	"math/big"
)

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
