package message

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/qibft/core"
)

// QBFT message codes
const (
	preprepareMsgCode  = 0x12
	prepareMsgCode     = 0x13
	commitMsgCode      = 0x14
	roundChangeMsgCode = 0x15
)

// A set containing the messages codes for all QBFT messages.
func MessageCodes() map[uint64]struct{} {
	return map[uint64]struct{}{
		preprepareMsgCode:  {},
		prepareMsgCode:     {},
		commitMsgCode:      {},
		roundChangeMsgCode: {},
	}
}

// Common interface for all QBFT messages
type QBFTMessage interface {
	Code() uint64
	View() core.View
	Source() common.Address
	SetSource(address common.Address)
	EncodePayload() ([]byte, error)
	Signature() []byte
	SetSignature(signature []byte)
}
