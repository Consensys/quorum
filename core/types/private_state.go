package types

import (
	"encoding/json"
	"strings"

	"github.com/ethereum/go-ethereum/private/engine"
)

const (
	// DefaultPrivateStateIdentifier is the default privacy group name created when
	// no privacy group has been configured in Tessera
	DefaultPrivateStateIdentifier PrivateStateIdentifier = "private"
)

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

func (psi PrivateStateIdentifier) String() string {
	return string(psi)
}

func ToPrivateStateIdentifier(s string) PrivateStateIdentifier {
	return PrivateStateIdentifier(s)
}

// EncodePSI includes counter and PSI value in an JSON message ID.
// i.e.: <counter> becomes "<psi>/32"
func EncodePSI(idCounterBytes []byte, psi PrivateStateIdentifier) json.RawMessage {
	if len(psi) == 0 {
		return idCounterBytes
	}
	newID := make([]byte, len(idCounterBytes)+len(psi)+3) // including 2 double quotes and '@'
	newID[0], newID[len(newID)-1] = '"', '"'
	copy(newID[1:len(psi)+1], psi)
	copy(newID[len(psi)+1:], append([]byte("/"), idCounterBytes...))
	return newID
}

// DecodePSI extracts PSI value from an encoded JSON message ID. Return second value as false
// if no PSI is encoded
// i.e.: "<counter>/<psi>" returns <psi>
func DecodePSI(id json.RawMessage) (PrivateStateIdentifier, bool) {
	idStr := string(id)
	if !strings.HasPrefix(idStr, "\"") || !strings.HasSuffix(idStr, "\"") {
		return "", false
	}
	sepIdx := strings.Index(idStr, "/")
	if sepIdx == -1 {
		return "", false
	}
	return PrivateStateIdentifier(id[1:sepIdx]), true
}
