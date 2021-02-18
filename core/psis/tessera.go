package psis

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

type TesseraPrivacyGroupPSISImpl struct {
	groups             []engine.PrivacyGroup
	residentGroupByKey map[string]*core.PrivateStateMetadata
	privacyGroupById   map[string]*core.PrivateStateMetadata
}

func (t *TesseraPrivacyGroupPSISImpl) ResolveForManagedParty(managedParty string) (*core.PrivateStateMetadata, error) {
	psm, found := t.residentGroupByKey[managedParty]
	if !found {
		return nil, fmt.Errorf("Unable to find private state for managed party %s", managedParty)
	}

	return psm, nil
}

func (t *TesseraPrivacyGroupPSISImpl) ResolveForUserContext(ctx context.Context) (*core.PrivateStateMetadata, error) {
	psiTyped, ok := ctx.Value(rpc.CtxPrivateStateIdentifier).(types.PrivateStateIdentifier)
	var psi string
	if !ok {
		psi = "private"
	} else {
		psi = string(psiTyped)
	}
	psm, found := t.privacyGroupById[psi]
	if !found {
		return nil, fmt.Errorf("unable to find private state for context psi %s", psi)
	}
	return psm, nil
}

func (t *TesseraPrivacyGroupPSISImpl) Groups() []engine.PrivacyGroup {
	return t.groups
}

func NewTesseraPrivacyGroupPSIS() (core.PrivateStateIdentifierService, error) {
	groups, err := private.P.Groups()
	if err != nil {
		return nil, err
	}
	residentGroupByKey := make(map[string]*core.PrivateStateMetadata)
	privacyGroupById := make(map[string]*core.PrivateStateMetadata)
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

		existing, found := privacyGroupById[group.PrivacyGroupId]
		if found {
			return nil, fmt.Errorf("privacy groups id clash id=%s existing.Name=%s duplicate.Name=%s", existing.ID, existing.Name, group.Name)
		}
		privacyGroupById[group.PrivacyGroupId] = privacyGroupToPrivateStateMetadata(group)
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

	return &TesseraPrivacyGroupPSISImpl{
		groups:             convertedGroups,
		residentGroupByKey: residentGroupByKey,
		privacyGroupById:   privacyGroupById,
	}, nil
}

func privacyGroupToPrivateStateMetadata(group engine.PrivacyGroup) *core.PrivateStateMetadata {
	return &core.PrivateStateMetadata{
		ID:          group.PrivacyGroupId,
		Name:        group.Name,
		Description: group.Description,
		Type:        strTypeToPrivateStateType(group.Type),
		Addresses:   group.Members,
	}
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