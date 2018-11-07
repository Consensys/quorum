package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/golang-lru"
)

type AccessType uint8

const (
	FullAccess AccessType = iota
	ReadOnly
	Transact
	ContractDeploy
)

type PermStruct struct {
	AcctId     common.Address
	AcctAccess AccessType
}
type OrgStruct struct {
	OrgId string
	Keys  []string
}

var AcctMap, AcctMapErr = lru.NewARC(100)

var OrgKeyMap, OrgKeyMapErr = lru.NewARC(100)

func AddAccountAccess(acctId common.Address, access uint8) {
	AcctMap.Add(acctId, &PermStruct{AcctId: acctId, AcctAccess: AccessType(access)})
}

func GetAcctAccess(acctId common.Address) AccessType {
	if AcctMap.Len() != 0 {
		if val, ok := AcctMap.Get(acctId); ok {
			vo := val.(*PermStruct)
			return vo.AcctAccess
		}
	}
	if len(AcctMap) == 0 {
		return FullAccess
	} else {
		return ReadOnly
	}
}

func AddOrgKey(orgId string, key string) {
	if OrgKeyMap.Len() != 0 {
		if val, ok := OrgKeyMap.Get(orgId); ok {
			// Org record exists. Append the key only
			vo := val.(*OrgStruct)
			vo.Keys = append(vo.Keys, key)
			return
		}
	}
	OrgKeyMap.Add(orgId, &OrgStruct{OrgId: orgId, Keys: []string{key}})
}

func DeleteOrgKey(orgId string, key string) {
	if val, ok := OrgKeyMap.Get(orgId); ok {
		vo := val.(*OrgStruct)
		for i, keyVal := range vo.Keys {
			if keyVal == key {
				vo.Keys = append(vo.Keys[:i], vo.Keys[i+1:]...)
				break
			}
		}
	}
}

func ResolvePrivateForKeys(orgId string) []string {
	var keys []string
	if val, ok := OrgKeyMap.Get(orgId); ok {
		vo := val.(*OrgStruct)
		if len(vo.Keys) > 0 {
			keys = vo.Keys
		} else {
			keys = append(keys, orgId)
		}
		return keys
	}
	keys = append(keys, orgId)
	return keys
}
