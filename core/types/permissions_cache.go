package types
import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type AccessType uint8

const (
	FullAccess AccessType = iota
	ReadOnly
	Transact
	ContractDeploy
)

type PermStruct struct {
	AcctId common.Address
	AcctAccess AccessType
}
type OrgStruct struct {
	OrgId string
	Keys []string
}

type PermAccountsMap map[common.Address][] *PermStruct

type PermOrgKeyMap map[string][] *OrgStruct

var AcctMap = make(map[common.Address] *PermStruct)

var OrgKeyMap = make(map[string] *OrgStruct)

func AddAccountAccess(acctId common.Address, access uint8)  {
	log.Info("SMK-AddAccountAccess @36 ", "acct", acctId, "access", access)
	mu := sync.RWMutex{}

	mu.Lock()
    AcctMap[acctId] = &PermStruct {AcctId : acctId, AcctAccess : AccessType(access)}
	mu.Unlock()
}

func GetAcctAccess(acctId common.Address) AccessType {
	log.Info("SMK-GetAcctAccess @44 ")
	mu := sync.RWMutex{}
	if len(AcctMap) != 0 {
		log.Info("SMK-GetAcctAccess @47 in if")
		if _, ok := AcctMap[acctId]; ok {
			log.Info("SMK-GetAcctAccess @49 in if")
			mu.RLock()
			acctAccess := AcctMap[acctId].AcctAccess
			mu.RUnlock()
			log.Info("SMK-GetAcctAccess @53 in if", "acctAccess", acctAccess)
			return acctAccess
		}
	}
	if len(AcctMap) == 0 {
		log.Info("SMK-GetAcctAccess @58 in if full access")
		return FullAccess
	} else {
		log.Info("SMK-GetAcctAccess @58 in if else readonly")
		return ReadOnly
	}
}

func AddOrgKey(orgId string, keys string){

	if len(OrgKeyMap) != 0 {
		if _, ok := OrgKeyMap[orgId]; ok {
			// Org record exists. Append the key only
			OrgKeyMap[orgId].Keys = append (OrgKeyMap[orgId].Keys, keys)
			return
		}
	}
	// first record into the map or firts record for the org
	var locKeys []string
	locKeys = append(locKeys, keys);
	OrgKeyMap[orgId] = &OrgStruct {OrgId : orgId, Keys : locKeys}
}

func DeleteOrgKey(orgId string, keys string){

	if len(OrgKeyMap) != 0 {
		if _, ok := OrgKeyMap[orgId]; ok {
			for i, keyVal := range OrgKeyMap[orgId].Keys{
				if keyVal == keys {
					OrgKeyMap[orgId].Keys = append(OrgKeyMap[orgId].Keys[:i], OrgKeyMap[orgId].Keys[i+1:]...)
					break
				}
			}
		}
	}
}

func ResolvePrivateForKeys(orgId string ) []string {
	var keys []string
	mu := sync.RWMutex{}

	if len(OrgKeyMap) != 0 {
		if _, ok := OrgKeyMap[orgId]; ok {
			if len(OrgKeyMap[orgId].Keys) > 0{
				mu.RLock()
				keys = OrgKeyMap[orgId].Keys
				mu.RUnlock()
			} else {
				keys = append(keys, orgId)
			}
			return keys
		}
	}
	keys = append(keys, orgId)
	return keys
}
