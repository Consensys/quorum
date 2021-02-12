package types

import "github.com/ethereum/go-ethereum/private/engine"

const (
	// DefaultPrivateStateIdentifier is the default privacy group name created when
	// no privacy group has been configured in Tessera
	DefaultPrivateStateIdentifier PrivateStateIdentifier = "private"
)

// PrivacyMetadata encapsulates privacy information to be attached
// to a transaction being processed
type PrivacyMetadata struct {
	PrivacyFlag engine.PrivacyFlagType
}

// PrivateStateIdentifier is an unique identifier of a private state.
// The value comes from Tessera privacy group detail,
// it could be a privacy group name or ID
type PrivateStateIdentifier string
