package message

import (
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

// A QBFT PREPARE message.
type Prepare struct {
	SignedPreparePayload
}

func NewPrepare(sequence *big.Int, round *big.Int, digest common.Hash) *Prepare {
	return &Prepare{SignedPreparePayload{
		CommonPayload: CommonPayload{
			code:     PrepareCode,
			Sequence: sequence,
			Round:    round,
		},
		Digest: digest,
	}}
}

type SignedPreparePayload struct {
	CommonPayload
	Digest common.Hash
}

func NewSignedPreparePayload(sequence *big.Int, round *big.Int, digest common.Hash, signature []byte, source common.Address) *SignedPreparePayload {
	return &SignedPreparePayload{
		CommonPayload: CommonPayload{
			code:      PrepareCode,
			source:    source,
			Sequence:  sequence,
			Round:     round,
			signature: signature,
		},
		Digest: digest,
	}
}

func (p *SignedPreparePayload) String() string {
	return fmt.Sprintf("Prepare {seq=%v, round=%v, digest=%v}", p.Sequence, p.Round, p.Digest.Hex())
}

func (p *SignedPreparePayload) EncodePayload() ([]byte, error) {
	return rlp.EncodeToBytes([]interface{}{p.Sequence, p.Round, p.Digest})
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
			Sequence *big.Int
			Round    *big.Int
			Digest   common.Hash
		}
		Signature []byte
	}
	if err := stream.Decode(&message); err != nil {
		return err
	}
	signedPayload.code = PrepareCode
	signedPayload.Sequence = message.Payload.Sequence
	signedPayload.Round = message.Payload.Round
	signedPayload.Digest = message.Payload.Digest
	signedPayload.signature = message.Signature
	return nil
}
