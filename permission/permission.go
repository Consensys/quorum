package permission

import (
	"encoding/json"
	"fmt"
	eea "github.com/ethereum/go-ethereum/permission/bind"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)



type PermissionService interface {
	AfterStart() error
	AsyncStart()
	Start(srvr *p2p.Server) error
	APIs() []rpc.API
	Protocols() []p2p.Protocol
	MonitorQIP714Block() error
	ManageOrgPermissions() error
	SubscribeStopEvent() (chan types.StopEvent, event.Subscription)
	ManageNodePermissions() error
	UpdateFile(fileName, enodeId string, operation types.NodeOperation, createFile bool)
	UpdatePermissionedNodes(enodeId string, operation types.NodeOperation)
	UpdateDisallowedNodes(url string, operation types.NodeOperation)
	ManageAccountPermissions() error
	DisconnectNode(enodeId string)
	InstantiateCache(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int)
	PopulateInitPermissions(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int) error
	BootupNetwork(permInterfSession *eea.EeaPermInterfaceSession) error
	PopulateAccountsFromContract(auth *bind.TransactOpts) error
	PopulateRolesFromContract(auth *bind.TransactOpts) error
	PopulateNodesFromContract(auth *bind.TransactOpts) error
	PopulateOrgsFromContract(auth *bind.TransactOpts) error
	PopulateStaticNodesToContract(permissionsSession *eea.EeaPermInterfaceSession) error
	PopulateInitAccountAccess(permissionsSession *eea.EeaPermInterfaceSession) error
	UpdateNetworkStatus(permissionsSession *eea.EeaPermInterfaceSession) error
	ManageRolePermissions() error
	PopulateAccountToCache(acctId common.Address) (*types.AccountInfo, error)
	PopulateOrgToCache(orgId string) (*types.OrgInfo, error)
	PopulateRoleToCache(roleKey *types.RoleKey) (*types.RoleInfo, error)
	PopulateNodeCache(url string) (*types.NodeInfo, error)
	PopulateNodeCacheAndValidate(hexNodeId, ultimateParentId string) bool
	BlockChain() *core.BlockChain
	PermissionConfig() *types.PermissionConfig
	Ethereum() *eth.Ethereum
	PermissionInterface() interface{}
}



// function reads the permissions config file passed and populates the
// config structure accordingly
func ParsePermissionConfig(dir string) (types.PermissionConfig, error) {
	fullPath := filepath.Join(dir, params.PERMISSION_MODEL_CONFIG)
	f, err := os.Open(fullPath)
	if err != nil {
		log.Error("can't open file", "file", fullPath, "error", err)
		return types.PermissionConfig{}, err
	}
	defer func() {
		_ = f.Close()
	}()

	var permConfig types.PermissionConfig
	blob, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Error("error reading file", "err", err, "file", fullPath)
	}

	err = json.Unmarshal(blob, &permConfig)
	if err != nil {
		log.Error("error unmarshalling the file", "err", err, "file", fullPath)
	}

	if len(permConfig.Accounts) == 0 {
		return types.PermissionConfig{}, fmt.Errorf("no accounts given in %s. Network cannot boot up", params.PERMISSION_MODEL_CONFIG)
	}
	if permConfig.SubOrgDepth.Cmp(big.NewInt(0)) == 0 || permConfig.SubOrgBreadth.Cmp(big.NewInt(0)) == 0 {
		return types.PermissionConfig{}, fmt.Errorf("sub org breadth depth not passed in %s. Network cannot boot up", params.PERMISSION_MODEL_CONFIG)
	}
	if permConfig.IsEmpty() {
		return types.PermissionConfig{}, fmt.Errorf("missing contract addresses in %s", params.PERMISSION_MODEL_CONFIG)
	}

	return permConfig, nil
}

// Create a service instance for permissioning
//
// Permission Service depends on the following:
// 1. EthService to be ready
// 2. Downloader to sync up blocks
// 3. InProc RPC server to be ready
func NewQuorumPermissionCtrl(stack *node.Node, pconfig *types.PermissionConfig, eeaFlag bool) (PermissionService, error) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	var p PermissionService
	if eeaFlag {
		p = &EeaPermissionCtrl{
			Node:           stack,
			Key:            stack.GetNodeKey(),
			DataDir:        stack.DataDir(),
			PermConfig:     pconfig,
			StartWaitGroup: wg,
			ErrorChan:      make(chan error),
		}
	} else {
		// TODO (Amal) add basic permission contract
	}

	stopChan, stopSubscription := p.SubscribeStopEvent()
	inProcRPCServerSub := stack.EventMux().Subscribe(rpc.InProcServerReadyEvent{})
	log.Debug("permission service: waiting for InProcRPC Server")

	go func(_wg *sync.WaitGroup) {
		defer func(start time.Time) {
			log.Debug("permission service: InProcRPC server is ready", "took", time.Since(start))
			stopSubscription.Unsubscribe()
			inProcRPCServerSub.Unsubscribe()
			_wg.Done()
		}(time.Now())
		select {
		case <-inProcRPCServerSub.Chan():
		case <-stopChan:
		}
	}(wg) // wait for inproc RPC to be ready
	return p, nil
}


