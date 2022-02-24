package core

import (
	"encoding/base64"
	"fmt"

	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/privatecache"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/private/engine"
)

// newPrivateStateManager instantiates an instance of mps.PrivateStateManager based on
// the given isMPS flag.
//
// If isMPS is true, it also does the validation to make sure
// the target private.PrivateTransactionManager supports MPS
func newPrivateStateManager(db ethdb.Database, privateCacheProvider privatecache.Provider, isMPS bool) (mps.PrivateStateManager, error) {
	if isMPS {
		// validation
		if !private.P.HasFeature(engine.MultiplePrivateStates) {
			return nil, fmt.Errorf("cannot instantiate MultiplePrivateStateManager while the transaction manager does not support multiple private states")
		}
		groups, err := private.P.Groups()
		if err != nil {
			return nil, err
		}
		residentGroupByKey := make(map[string]*mps.PrivateStateMetadata)
		privacyGroupById := make(map[types.PrivateStateIdentifier]*mps.PrivateStateMetadata)
		for _, group := range groups {
			if group.Type == engine.PrivacyGroupResident {
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
			if group.Type == engine.PrivacyGroupResident {
				for _, address := range group.Members {
					existing, found := residentGroupByKey[address]
					if found {
						return nil, fmt.Errorf("same address is part of two different groups: address=%s existing.Name=%s duplicate.Name=%s", address, existing.Name, group.Name)
					}
					residentGroupByKey[address] = privacyGroupToPrivateStateMetadata(group)
				}
			}
		}
		return newMultiplePrivateStateManager(db, privateCacheProvider, residentGroupByKey, privacyGroupById)
	} else {
		return newDefaultPrivateStateManager(db, privateCacheProvider), nil
	}
}

func privacyGroupToPrivateStateMetadata(group engine.PrivacyGroup) *mps.PrivateStateMetadata {
	return mps.NewPrivateStateMetadata(
		types.ToPrivateStateIdentifier(group.PrivacyGroupId),
		group.Name,
		group.Description,
		strTypeToPrivateStateType(group.Type),
		group.Members,
	)
}

func strTypeToPrivateStateType(strType string) mps.PrivateStateType {
	switch strType {
	case engine.PrivacyGroupLegacy:
		return mps.Legacy
	case engine.PrivacyGroupPantheon:
		return mps.Pantheon
	default:
		return mps.Resident
	}
}
