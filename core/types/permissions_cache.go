package types
import (
	"sync"
	"strings"

	"github.com/ethereum/go-ethereum/common"
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
	mu := sync.RWMutex{}

	mu.Lock()
    AcctMap[acctId] = &PermStruct {AcctId : acctId, AcctAccess : AccessType(access)}
	mu.Unlock()
}

func GetAcctAccess(acctId common.Address) AccessType {
	mu := sync.RWMutex{}
	if len(AcctMap) != 0 {
		if _, ok := AcctMap[acctId]; ok {
			mu.RLock()
			acctAccess := AcctMap[acctId].AcctAccess
			mu.RUnlock()
			return acctAccess
		}
	}
	return FullAccess
}

func AddOrgKey(orgId string, keys string){
	mu := sync.RWMutex{}

	orgKeys := strings.Fields(keys)
	mu.Lock()
	OrgKeyMap[orgId] = &OrgStruct {OrgId : orgId, Keys : orgKeys}
	mu.Unlock()
}

func ResolvePrivateForKeys(orgId string ) []string {
	var keys []string
	mu := sync.RWMutex{}

	AddOrgKey("JPMORG", "QfeDAys9MPDs2XHExtc84jKGHxZg/aj52DTh0vtA3Xc= 1iTZde/ndBHvzhcl7V68x44Vx7pl8nwx9LqnM/AfJUg=")
	if len(OrgKeyMap) != 0 {
		if _, ok := OrgKeyMap[orgId]; ok {
			mu.RLock()
			keys := OrgKeyMap[orgId].Keys
			mu.RUnlock()
			return keys
		}
	}
	keys = append(keys, orgId)
	return keys
}
