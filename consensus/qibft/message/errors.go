package message

import "errors"

var (
	ErrFailedDecodePreprepare  = errors.New("failed to decode PRE-PREPARE message")
	ErrFailedDecodePrepare     = errors.New("failed to decode PREPARE message")
	ErrFailedDecodeCommit      = errors.New("failed to decode COMMIT message")
	ErrFailedDecodeRoundChange = errors.New("failed to decode ROUND-CHANGE message")
	ErrInvalidMessage          = errors.New("invalid message")
	/*
		// errInconsistentSubject is returned when received subject is different from
		// current subject.
		errInconsistentSubject = errors.New("inconsistent subjects")
		// errNotFromProposer is returned when received message is supposed to be from
		// proposer.
		errNotFromProposer = errors.New("message does not come from proposer")
		// errFutureMessage is returned when current view is earlier than the
		// view of the received message.
		errFutureMessage = errors.New("future message")
		// errOldMessage is returned when the received message's view is earlier
		// than current view.
		errOldMessage = errors.New("old message")
		// errInvalidMessage is returned when the message is malformed.
		errInvalidMessage = errors.New("invalid message")
		// errFailedDecodePreprepare is returned when the PRE-PREPARE message is malformed.
		errFailedDecodePreprepare = errors.New("failed to decode PRE-PREPARE")
		// errFailedDecodeRoundChange is returned when the ROUNDCHANGE message is malformed.
		errFailedDecodeRoundChange = errors.New("failed to decode ROUNDCHANGE")
		// errFailedDecodePrepare is returned when the PREPARE message is malformed.
		errFailedDecodePrepare = errors.New("failed to decode PREPARE")
		// errFailedDecodeCommit is returned when the COMMIT message is malformed.
		errFailedDecodeCommit = errors.New("failed to decode COMMIT")
		// errFailedDecodePiggybackMsgs is returned when the Piggyback messages are malformed
		errFailedDecodePiggybackMsgs = errors.New("failed to decode Piggyback Messages")
		// errInvalidSigner is returned when the message is signed by a validator different than message sender
		errInvalidSigner = errors.New("message not signed by the sender")
		// errInvalidPreparedBlock is returned when prepared block is not validated in round change messages
		errInvalidPreparedBlock = errors.New("invalid prepared block in round change messages")*/
)
