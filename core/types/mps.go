package types

import "fmt"

var (
	// DefaultPrivateStateIdentifier is the default privacy group name created when
	// no privacy group has been configured in Tessera
	DefaultPrivateStateIdentifier PrivateStateIdentifier = ToPrivateStateIdentifier("private")
	// EmptyPrivateStateIdentifier is the identifier for the empty private state
	// which is to hold state of transactions "as if" to which the node is not party
	EmptyPrivateStateIdentifier PrivateStateIdentifier = ToPrivateStateIdentifier("empty")
	// DefaultPrivateStateMetadata is the metadata for the single private state being used
	// when MPS is disabled
	DefaultPrivateStateMetadata = NewPrivateStateMetadata(
		DefaultPrivateStateIdentifier,
		"private",
		"legacy private state",
		Resident,
		nil,
	)
	// EmptyPrivateStateMetadata is the metadata for the empty private state being used
	// when MPS is enabled
	EmptyPrivateStateMetadata = NewPrivateStateMetadata(
		EmptyPrivateStateIdentifier,
		"empty",
		"empty state metadata",
		Resident,
		nil,
	)
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

type PrivateStateType uint64

const (
	Resident PrivateStateType = iota                          // 0
	Legacy   PrivateStateType = 1 << PrivateStateType(iota-1) // 1
	Pantheon PrivateStateType = 1 << PrivateStateType(iota-1) // 2
)

// PrivateStateMetadata is the domain model in Quorum which maps with
// PrivacyGroup domain in Tessera
type PrivateStateMetadata struct {
	ID          PrivateStateIdentifier
	Name        string
	Description string
	Type        PrivateStateType
	// Addresses stores the public keys in the SAME ORDER being configured
	// in Tessera
	Addresses []string
	// addressIndex is to facilitate fast searching
	addressIndex map[string]struct{}
}

func (psm *PrivateStateMetadata) NotIncludeAny(addresses ...string) bool {
	for _, addr := range addresses {
		if _, found := psm.addressIndex[addr]; found {
			return false
		}
	}
	return true
}

func (psm *PrivateStateMetadata) String() string {
	return fmt.Sprintf("ID=%s,Name=%s,Desc=%s,Type=%d,Addresses=%v", psm.ID, psm.Name, psm.Description, psm.Type, psm.Addresses)
}

func NewPrivateStateMetadata(id PrivateStateIdentifier, name, desc string, t PrivateStateType, addresses []string) *PrivateStateMetadata {
	index := make(map[string]struct{}, len(addresses))
	for _, a := range addresses {
		index[a] = struct{}{}
	}
	return &PrivateStateMetadata{
		ID:           id,
		Name:         name,
		Description:  desc,
		Type:         t,
		Addresses:    addresses[:],
		addressIndex: index,
	}
}
