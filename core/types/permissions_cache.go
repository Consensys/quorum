package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/hashicorp/golang-lru"
	"sync"
)

type AccessType uint8

const (
	ReadOnly AccessType = iota
	Transact
	ContractDeploy
	FullAccess
)

type PermStruct struct {
	AcctId     common.Address
	AcctAccess AccessType
}
type OrgStruct struct {
	OrgId string
	Keys  []string
}

var DefaultAccess = FullAccess

const acctMapLimit = 100
const orgKeyMapLimit = 100

var AcctMap, _ = lru.New(acctMapLimit)

var OrgKeyMap, _ = lru.New(orgKeyMapLimit)

var orgKeyLock sync.Mutex

func SetDefaultAccess() {
	DefaultAccess = ReadOnly
}
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
	return DefaultAccess
}

func AddOrgKey(orgId string, key string) {
	if OrgKeyMap.Len() != 0 {
		if val, ok := OrgKeyMap.Get(orgId); ok {
			orgKeyLock.Lock()
			// Org record exists. Append the key only
			vo := val.(*OrgStruct)
			vo.Keys = append(vo.Keys, key)
			orgKeyLock.Unlock()
			return
		}
	}
	OrgKeyMap.Add(orgId, &OrgStruct{OrgId: orgId, Keys: []string{key}})
}

func DeleteOrgKey(orgId string, key string) {
	defer orgKeyLock.Unlock()
	if val, ok := OrgKeyMap.Get(orgId); ok {
		orgKeyLock.Lock()
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
