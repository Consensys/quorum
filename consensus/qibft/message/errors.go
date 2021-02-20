package message

import "errors"

var (
	ErrFailedDecodePreprepare = errors.New("failed to decode PRE-PREPARE message")
	ErrFailedDecodePrepare = errors.New("failed to decode PREPARE message")
	ErrFailedDecodeCommit = errors.New("failed to decode COMMIT message")
	ErrFailedDecodeRoundChange = errors.New("failed to decode ROUND-CHANGE message")
	ErrInvalidMessage = errors.New("invalid message")
	/*
	// errInconsistentSubject is returned when received subject is different from
	// current subject.
	errInconsistentSubject = errors.New("inconsistent subjects")
	// errNotFromProposer is returned when received message_deprecated is supposed to be from
	// proposer.
	errNotFromProposer = errors.New("message_deprecated does not come from proposer")
	// errFutureMessage is returned when current view is earlier than the
	// view of the received message_deprecated.
	errFutureMessage = errors.New("future message_deprecated")
	// errOldMessage is returned when the received message_deprecated's view is earlier
	// than current view.
	errOldMessage = errors.New("old message_deprecated")
	// errInvalidMessage is returned when the message_deprecated is malformed.
	errInvalidMessage = errors.New("invalid message_deprecated")
	// errFailedDecodePreprepare is returned when the PRE-PREPARE message_deprecated is malformed.
	errFailedDecodePreprepare = errors.New("failed to decode PRE-PREPARE")
	// errFailedDecodeRoundChange is returned when the ROUNDCHANGE message_deprecated is malformed.
	errFailedDecodeRoundChange = errors.New("failed to decode ROUNDCHANGE")
	// errFailedDecodePrepare is returned when the PREPARE message_deprecated is malformed.
	errFailedDecodePrepare = errors.New("failed to decode PREPARE")
	// errFailedDecodeCommit is returned when the COMMIT message_deprecated is malformed.
	errFailedDecodeCommit = errors.New("failed to decode COMMIT")
	// errFailedDecodePiggybackMsgs is returned when the Piggyback messages are malformed
	errFailedDecodePiggybackMsgs = errors.New("failed to decode Piggyback Messages")
	// errInvalidSigner is returned when the message_deprecated is signed by a validator different than message_deprecated sender
	errInvalidSigner = errors.New("message_deprecated not signed by the sender")
	// errInvalidPreparedBlock is returned when prepared block is not validated in round change messages
	errInvalidPreparedBlock = errors.New("invalid prepared block in round change messages")*/
)
