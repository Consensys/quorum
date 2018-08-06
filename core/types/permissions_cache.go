package types
import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type PermStruct struct {
	AcctId common.Address
	Access uint8
}

type PermAccountsMap map[common.Address][] *PermStruct

var AcctMap = make(map[common.Address] *PermStruct)

func AddAccountAccess(acctId common.Address, access uint8)  {
	log.Info("Inside PutAcctmap adding :", "acctId", acctId, "access", access)

	mu := sync.RWMutex{}

	mu.Lock()
    AcctMap[acctId] = &PermStruct {AcctId : acctId, Access : access}
	mu.Unlock()
}

func GetAcctAccess(acctId common.Address) uint8 {
	mu := sync.RWMutex{}

	if len(AcctMap) != 0 {
		if _, ok := AcctMap[acctId]; ok {
			log.Info("Inside GetAcct sending :", "acctId", AcctMap[acctId].AcctId, "access", AcctMap[acctId].Access)

			mu.RLock()
			access := AcctMap[acctId].Access
			mu.RUnlock()

			return access
		}
	}
	return 99
}
