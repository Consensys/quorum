package types
import (
	"sync"

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

type PermAccountsMap map[common.Address][] *PermStruct

var AcctMap = make(map[common.Address] *PermStruct)

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
