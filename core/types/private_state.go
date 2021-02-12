package types

import "github.com/ethereum/go-ethereum/private/engine"

// PrivacyMetadata encapsulates privacy information to be attached
// to a transaction being processed
type PrivacyMetadata struct {
	PrivacyFlag engine.PrivacyFlagType
}

// PrivateStateIdentifier is an unique identifier of a private state.
// The value comes from Tessera privacy group detail,
// it could be a privacy group name or ID
type PrivateStateIdentifier string
