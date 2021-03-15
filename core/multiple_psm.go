package core

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/rpc"
)

type MultiplePrivateStateManager struct {
	bc                     *BlockChain
	privateStatesTrieCache state.Database

	residentGroupByKey map[string]*types.PrivateStateMetadata
	privacyGroupById   map[types.PrivateStateIdentifier]*types.PrivateStateMetadata
}

func NewMultiplePrivateStateManager(bc *BlockChain) (*MultiplePrivateStateManager, error) {
	groups, err := private.P.Groups()
	if err != nil {
		return nil, err
	}
	residentGroupByKey := make(map[string]*types.PrivateStateMetadata)
	privacyGroupById := make(map[types.PrivateStateIdentifier]*types.PrivateStateMetadata)
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
	return &MultiplePrivateStateManager{
		bc:                     bc,
		privateStatesTrieCache: state.NewDatabase(bc.db),
		residentGroupByKey:     residentGroupByKey,
		privacyGroupById:       privacyGroupById,
	}, nil
}

func (m *MultiplePrivateStateManager) GetPrivateStateRepository(blockHash common.Hash) (mps.PrivateStateRepository, error) {
	return mps.NewMultiplePrivateStateRepository(m.bc.chainConfig, m.bc.db, m.privateStatesTrieCache, blockHash)
}

func (m *MultiplePrivateStateManager) ResolveForManagedParty(managedParty string) (*types.PrivateStateMetadata, error) {
	psm, found := m.residentGroupByKey[managedParty]
	if !found {
		return nil, fmt.Errorf("unable to find private state metadata for managed party %s", managedParty)
	}
	return psm, nil
}

func (m *MultiplePrivateStateManager) ResolveForUserContext(ctx context.Context) (*types.PrivateStateMetadata, error) {
	psi, ok := ctx.Value(rpc.CtxPrivateStateIdentifier).(types.PrivateStateIdentifier)
	if !ok {
		psi = types.DefaultPrivateStateIdentifier
	}
	psm, found := m.privacyGroupById[psi]
	if !found {
		return nil, fmt.Errorf("unable to find private state for context psi %s", psi)
	}
	return psm, nil
}

func (m *MultiplePrivateStateManager) PSIs() []types.PrivateStateIdentifier {
	psis := make([]types.PrivateStateIdentifier, 0, len(m.privacyGroupById))
	for psi := range m.privacyGroupById {
		psis = append(psis, psi)
	}
	return psis
}

func (m *MultiplePrivateStateManager) NotIncludeAny(psm *types.PrivateStateMetadata, managedParties ...string) bool {
	return psm.NotIncludeAny(managedParties...)
}

func (m *MultiplePrivateStateManager) GetCache() state.Database {
	return m.privateStatesTrieCache
}

func privacyGroupToPrivateStateMetadata(group engine.PrivacyGroup) *types.PrivateStateMetadata {
	return types.NewPrivateStateMetadata(
		types.ToPrivateStateIdentifier(group.PrivacyGroupId),
		group.Name,
		group.Description,
		strTypeToPrivateStateType(group.Type),
		group.Members,
	)
}

func strTypeToPrivateStateType(strType string) types.PrivateStateType {
	if strType == "LEGACY" {
		return types.Legacy
	}
	if strType == "PANTHEON" {
		return types.Pantheon
	}
	return types.Resident
}
