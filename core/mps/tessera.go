package mps

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/rpc"
)

type TesseraPrivateStateMetadataResolver struct {
	residentGroupByKey map[string]*core.PrivateStateMetadata
	privacyGroupById   map[types.PrivateStateIdentifier]*core.PrivateStateMetadata
}

func (t *TesseraPrivateStateMetadataResolver) ResolveForManagedParty(managedParty string) (*core.PrivateStateMetadata, error) {
	psm, found := t.residentGroupByKey[managedParty]
	if !found {
		return nil, fmt.Errorf("unable to find private state metadata for managed party %s", managedParty)
	}
	return psm, nil
}

func (t *TesseraPrivateStateMetadataResolver) ResolveForUserContext(ctx context.Context) (*core.PrivateStateMetadata, error) {
	psi, ok := ctx.Value(rpc.CtxPrivateStateIdentifier).(types.PrivateStateIdentifier)
	if !ok {
		psi = types.DefaultPrivateStateIdentifier
	}
	psm, found := t.privacyGroupById[psi]
	if !found {
		return nil, fmt.Errorf("unable to find private state for context psi %s", psi)
	}
	return psm, nil
}

func (t *TesseraPrivateStateMetadataResolver) PSIs() []types.PrivateStateIdentifier {
	psis := make([]types.PrivateStateIdentifier, 0, len(t.privacyGroupById))
	for psi := range t.privacyGroupById {
		psis = append(psis, psi)
	}
	return psis
}

func (t *TesseraPrivateStateMetadataResolver) NotIncludeAny(psm *core.PrivateStateMetadata, managedParties ...string) bool {
	return psm.NotIncludeAny(managedParties...)
}

func NewTesseraPrivateStateMetadataResolver() (core.PrivateStateMetadataResolver, error) {
	groups, err := private.P.Groups()
	if err != nil {
		return nil, err
	}
	residentGroupByKey := make(map[string]*core.PrivateStateMetadata)
	privacyGroupById := make(map[types.PrivateStateIdentifier]*core.PrivateStateMetadata)
	convertedGroups := make([]engine.PrivacyGroup, 0)
	for _, group := range groups {
		if group.Type == "RESIDENT" {
			// Resident group IDs come in base64 encoded, so revert to original ID
			decoded, err := base64.StdEncoding.DecodeString(group.PrivacyGroupId)
			if err != nil {
				return nil, err
			}
			group.PrivacyGroupId = string(decoded)
		}
		psi := types.ToPrivateStateIdentifier(group.PrivacyGroupId)
		existing, found := privacyGroupById[psi]
		if found {
			return nil, fmt.Errorf("privacy groups id clash id=%s existing.Name=%s duplicate.Name=%s", existing.ID, existing.Name, group.Name)
		}
		privacyGroupById[psi] = privacyGroupToPrivateStateMetadata(group)
		if group.Type == "RESIDENT" {
			for _, address := range group.Members {
				existing, found := residentGroupByKey[address]
				if found {
					return nil, fmt.Errorf("same address is part of two different groups: address=%s existing.Name=%s duplicate.Name=%s", address, existing.Name, group.Name)
				}
				residentGroupByKey[address] = privacyGroupToPrivateStateMetadata(group)
			}
		}
		convertedGroups = append(convertedGroups, group)
	}

	return &TesseraPrivateStateMetadataResolver{
		residentGroupByKey: residentGroupByKey,
		privacyGroupById:   privacyGroupById,
	}, nil
}

func privacyGroupToPrivateStateMetadata(group engine.PrivacyGroup) *core.PrivateStateMetadata {
	return core.NewPrivateStateMetadata(
		types.ToPrivateStateIdentifier(group.PrivacyGroupId),
		group.Name,
		group.Description,
		strTypeToPrivateStateType(group.Type),
		group.Members,
	)
}

func strTypeToPrivateStateType(strType string) core.PrivateStateType {
	if strType == "LEGACY" {
		return core.Legacy
	}
	if strType == "PANTHEON" {
		return core.Pantheon
	}
	return core.Resident
}
