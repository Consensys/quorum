package message

import (
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/qbft"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type Preprepare struct {
	CommonPayload
	Proposal                  qbft.Proposal
	JustificationRoundChanges []*SignedRoundChangePayload
	JustificationPrepares     []*SignedPreparePayload
}

func NewPreprepare(sequence *big.Int, round *big.Int, proposal qbft.Proposal) *Preprepare {
	return &Preprepare{
		CommonPayload: CommonPayload{
			code:     PreprepareCode,
			Sequence: sequence,
			Round:    round,
		},
		Proposal: proposal,
	}
}

func (m *Preprepare) EncodePayloadForSigning() ([]byte, error) {
	return rlp.EncodeToBytes(
		[]interface{}{
			m.Code(),
			[]interface{}{m.Sequence, m.Round, m.Proposal},
		})
}

func (m *Preprepare) EncodeRLP(w io.Writer) error {
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

func (m *Preprepare) DecodeRLP(stream *rlp.Stream) error {
	var message struct {
		SignedPayload struct {
			Payload struct {
				Sequence *big.Int
				Round    *big.Int
				Proposal *types.Block
			}
			Signature []byte
		}
		Justification struct {
			RoundChanges []*SignedRoundChangePayload
			Prepares     []*SignedPreparePayload
		}
	}
	if err := stream.Decode(&message); err != nil {
		return err
	}
	m.code = PreprepareCode
	m.Sequence = message.SignedPayload.Payload.Sequence
	m.Round = message.SignedPayload.Payload.Round
	m.Proposal = message.SignedPayload.Payload.Proposal
	m.signature = message.SignedPayload.Signature
	m.JustificationPrepares = message.Justification.Prepares
	m.JustificationRoundChanges = message.Justification.RoundChanges
	return nil
}

func (m *Preprepare) String() string {
	return fmt.Sprintf("code: %d, sequence: %d, round: %d, proposal: %v", m.code, m.Sequence, m.Round, m.Proposal.Hash().Hex())
}