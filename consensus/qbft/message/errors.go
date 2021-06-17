package message

import "errors"

var (
	ErrFailedDecodePreprepare  = errors.New("failed to decode PRE-PREPARE message")
	ErrFailedDecodePrepare     = errors.New("failed to decode PREPARE message")
	ErrFailedDecodeCommit      = errors.New("failed to decode COMMIT message")
	ErrFailedDecodeRoundChange = errors.New("failed to decode ROUND-CHANGE message")
	ErrInvalidMessage          = errors.New("invalid message")
)
