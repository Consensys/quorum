package message

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/qibft/core"
	"math/big"
)


// Data that is common to all QBFT messages. Used for composition.
type CommonPayload struct {
	code      uint64
	source    common.Address
	Sequence  *big.Int
	Round     *big.Int
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

func (m *CommonPayload) View() core.View {
	return core.View{Sequence: m.Sequence, Round: m.Round}
}

func (m *CommonPayload) Signature() []byte {
	return m.signature
}

func (m *CommonPayload) SetSignature(signature []byte) {
	m.signature = signature
}
