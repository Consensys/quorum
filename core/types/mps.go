package types

var (
	// DefaultPrivateStateIdentifier is the default privacy group name created when
	// no privacy group has been configured in Tessera
	DefaultPrivateStateIdentifier = ToPrivateStateIdentifier("private")
	// EmptyPrivateStateIdentifier is the identifier for the empty private state
	// which is to hold state of transactions "as if" to which the node is not party
	EmptyPrivateStateIdentifier = ToPrivateStateIdentifier("empty")
)

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
