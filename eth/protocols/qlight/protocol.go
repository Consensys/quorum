package qlight

import (
	"errors"

	"github.com/ethereum/go-ethereum/qlight"
)

const (
	QLightStatusMsg              = 0x11
	QLightTokenUpdateMsg         = 0x12
	QLightNewBlockPrivateDataMsg = 0x13
)

const QLightProtocolLength = 20

const QLIGHT65 = 65
const ProtocolName = "qlight"

// maxMessageSize is the maximum cap on the size of a protocol message.
const maxMessageSize = 10 * 1024 * 1024

var (
	errNoStatusMsg    = errors.New("no status message")
	errMsgTooLarge    = errors.New("message too long")
	errDecode         = errors.New("invalid message")
	errInvalidMsgCode = errors.New("invalid message code")
)

type qLightStatusData struct {
	ProtocolVersion uint32
	Server          bool
	PSI             string
	Token           string
}

type qLightTokenUpdateData struct {
	Token string
}

type BlockPrivateDataPacket []qlight.BlockPrivateData

func (*BlockPrivateDataPacket) Name() string { return "BlockPrivateData" }
func (*BlockPrivateDataPacket) Kind() byte   { return QLightNewBlockPrivateDataMsg }
