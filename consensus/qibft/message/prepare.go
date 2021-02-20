package message

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
	"math/big"
)

// A QBFT PREPARE message.
type Prepare struct {
	SignedPreparePayload
}

func (m *Prepare) EncodePayload() ([]byte, error) {
	return rlp.EncodeToBytes([]interface{}{m.Sequence, m.Round, m.Digest})
}

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
