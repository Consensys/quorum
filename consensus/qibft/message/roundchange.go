package message

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
	"io"
	"math/big"
)

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

			return errors.Wrap(ErrFailedDecodePreprepare, "digest does not match block in payload of justification PREPARE")
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
