package psmr

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type PrivateStateType uint64

const (
	Resident PrivateStateType = iota                          // 0
	Legacy   PrivateStateType = 1 << PrivateStateType(iota-1) // 1
	Pantheon PrivateStateType = 1 << PrivateStateType(iota-1) // 2
)

// PrivateStateMetadata is the domain model in Quorum which maps with
// PrivacyGroup domain in Tessera
type PrivateStateMetadata struct {
	ID          types.PrivateStateIdentifier
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

func NewPrivateStateMetadata(id types.PrivateStateIdentifier, name, desc string, t PrivateStateType, addresses []string) *PrivateStateMetadata {
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

var EmptyPrivateStateMetadata = NewPrivateStateMetadata(
	types.ToPrivateStateIdentifier("empty"),
	"empty",
	"empty state metadata",
	Resident,
	nil,
)

var DefaultPrivateStateMetadata = NewPrivateStateMetadata(
	types.DefaultPrivateStateIdentifier,
	"private",
	"legacy private state",
	Resident,
	nil,
)

type PrivateStateMetadataResolver interface {
	ResolveForManagedParty(managedParty string) (*PrivateStateMetadata, error)
	ResolveForUserContext(ctx context.Context) (*PrivateStateMetadata, error)
	PSIs() []types.PrivateStateIdentifier
	// NotIncludeAny returns true if NONE of the managedParties is a member
	// of the given psm, otherwise returns false
	NotIncludeAny(psm *PrivateStateMetadata, managedParties ...string) bool
}

type DefaultPrivateStateMetadataResolver struct {
}

func (dpsmr *DefaultPrivateStateMetadataResolver) ResolveForManagedParty(_ string) (*PrivateStateMetadata, error) {
	return DefaultPrivateStateMetadata, nil
}

func (dpsmr *DefaultPrivateStateMetadataResolver) ResolveForUserContext(ctx context.Context) (*PrivateStateMetadata, error) {
	psi, ok := ctx.Value(rpc.CtxPrivateStateIdentifier).(types.PrivateStateIdentifier)
	if !ok {
		psi = types.DefaultPrivateStateIdentifier
	}
	return &PrivateStateMetadata{ID: psi, Type: Resident}, nil
}

func (dpsmr *DefaultPrivateStateMetadataResolver) PSIs() []types.PrivateStateIdentifier {
	return []types.PrivateStateIdentifier{
		types.DefaultPrivateStateIdentifier,
	}
}

func (dpsmr *DefaultPrivateStateMetadataResolver) NotIncludeAny(_ *PrivateStateMetadata, _ ...string) bool {
	// with default implementation, all managedParties are members of the psm
	return false
}
