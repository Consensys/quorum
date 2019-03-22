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
	AcctId common.Address
	roleId string
}
type OrgStruct struct {
	OrgId string
	Keys  []string
}

// permission config for bootstrapping
type PermissionConfig struct {
	UpgrdAddress   string
	InterfAddress  string
	ImplAddress    string
	NodeAddress    string
	AccountAddress string
	NwAdminOrg     string
	NwAdminRole    string
	OrgAdminRole   string

	Accounts []string //initial list of account that need full access
}

var DefaultAccess = FullAccess

const acctMapLimit = 100
const orgKeyMapLimit = 100

var AcctMap, _ = lru.New(acctMapLimit)

var OrgKeyMap, _ = lru.New(orgKeyMapLimit)

var orgKeyLock sync.Mutex

func (pc *PermissionConfig) IsEmpty() bool {
	return pc.InterfAddress == "" || pc.NodeAddress == "" || pc.AccountAddress == ""
}

// sets default access to ReadOnly
func SetDefaultAccess() {
	DefaultAccess = FullAccess
}

// Adds account access to the cache
func AddAccountAccess(acctId common.Address, roleId string) {
	AcctMap.Add(acctId, &PermStruct{AcctId: acctId, roleId: roleId})
}

// Returns the access type for an account. If not found returns
// default access
func GetAcctAccess(acctId common.Address) AccessType {
	if AcctMap.Len() != 0 {
		if _, ok := AcctMap.Get(acctId); ok {
			// val.(*PermStruct)
			return DefaultAccess
		}
	}
	return DefaultAccess
}

// Adds org key details to cache
func AddOrgKey(orgId string, key string) {
	if OrgKeyMap.Len() != 0 {
		if val, ok := OrgKeyMap.Get(orgId); ok {
			orgKeyLock.Lock()
			defer orgKeyLock.Unlock()
			// Org record exists. Append the key only
			vo := val.(*OrgStruct)
			vo.Keys = append(vo.Keys, key)
			return
		}
	}
	OrgKeyMap.Add(orgId, &OrgStruct{OrgId: orgId, Keys: []string{key}})
}

// deletes org key details from cache
func DeleteOrgKey(orgId string, key string) {
	if val, ok := OrgKeyMap.Get(orgId); ok {
		orgKeyLock.Lock()
		defer orgKeyLock.Unlock()
		vo := val.(*OrgStruct)
		for i, keyVal := range vo.Keys {
			if keyVal == key {
				vo.Keys = append(vo.Keys[:i], vo.Keys[i+1:]...)
				break
			}
		}
	}
}

// Givens a orgid returns the linked keys for the org
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
