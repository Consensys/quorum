package permission

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/params"
	pcore "github.com/ethereum/go-ethereum/permission/core"
	ptype "github.com/ethereum/go-ethereum/permission/core/types"
)

// This is to make sure all contract instances are ready and initialized
//
// Required to be call after standard service start lifecycle
func (p *PermissionCtrl) AfterStart() error {
	log.Debug("permission service: binding contracts")
	err := <-p.errorChan // capture any error happened during asyncStart. Also wait here if asyncStart is not yet finish
	if err != nil {
		return err
	}
	if err = p.contract.BindContracts(); err != nil {
		return fmt.Errorf("populateInitPermissions failed to bind contracts: %v", err)
	}

	// populate the initial list of permissioned nodes and account accesses
	if err := p.populateInitPermissions(params.DEFAULT_ORGCACHE_SIZE, params.DEFAULT_ROLECACHE_SIZE,
		params.DEFAULT_NODECACHE_SIZE, params.DEFAULT_ACCOUNTCACHE_SIZE); err != nil {
		return fmt.Errorf("populateInitPermissions failed: %v", err)
	}

	// set the function point for transaction allowed check
	pcore.PermissionTransactionAllowedFunc = p.IsTransactionAllowed
	setPermissionService(p)

	// set the default access to ReadOnly
	pcore.SetDefaults(p.permConfig.NwAdminRole, p.permConfig.OrgAdminRole, p.IsV2Permission())
	for _, f := range []func() error{
		p.monitorQIP714Block,               // monitor block number to activate new permissions controls
		p.backend.ManageOrgPermissions,     // monitor org management related events
		p.backend.ManageNodePermissions,    // monitor org  level Node management events
		p.backend.ManageRolePermissions,    // monitor org level role management events
		p.backend.ManageAccountPermissions, // monitor org level account management events
	} {
		if err := f(); err != nil {
			return err
		}
	}

	log.Info("permission service: is now ready")

	return nil
}

// start service asynchronously due to dependencies
func (p *PermissionCtrl) asyncStart() {
	var ethereum *eth.Ethereum
	// will be blocked here until Node is up
	if err := p.node.Service(&ethereum); err != nil {
		p.errorChan <- fmt.Errorf("dependent ethereum service not started")
		return
	}
	defer func() {
		p.errorChan <- nil
	}()
	// for cases where the node is joining an existing network, permission service
	// can be brought up only after block syncing is complete. This function
	// waits for block syncing before the starting permissions
	p.startWaitGroup.Add(1)
	go func(_wg *sync.WaitGroup) {
		log.Debug("permission service: waiting for downloader")
		stopChan, stopSubscription := ptype.SubscribeStopEvent()
		pollingTicker := time.NewTicker(10 * time.Millisecond)
		defer func(start time.Time) {
			log.Debug("permission service: downloader completed", "took", time.Since(start))
			stopSubscription.Unsubscribe()
			pollingTicker.Stop()
			_wg.Done()
		}(time.Now())
		for {
			select {
			case <-pollingTicker.C:
				if pcore.GetSyncStatus() && !ethereum.Downloader().Synchronising() {
					return
				}
			case <-stopChan:
				return
			}
		}
	}(p.startWaitGroup) // wait for downloader to sync if any

	log.Debug("permission service: waiting for all dependencies to be ready")
	p.startWaitGroup.Wait()
	client, err := p.node.Attach()
	if err != nil {
		p.errorChan <- fmt.Errorf("unable to create rpc client: %v", err)
		return
	}
	p.ethClnt = ethclient.NewClient(client)
	p.eth = ethereum
	p.isRaft = p.eth.BlockChain().Config().Istanbul == nil && p.eth.BlockChain().Config().Clique == nil
	p.updateBackEnd()
}

// monitors QIP714Block and set default access
func (p *PermissionCtrl) monitorQIP714Block() error {
	// if QIP714block is not given, set the default access
	// to readonly
	if p.eth.BlockChain().Config().QIP714Block == nil || p.eth.BlockChain().Config().IsQIP714(p.eth.BlockChain().CurrentBlock().Number()) {
		pcore.SetQIP714BlockReached()
		return nil
	}
	//QIP714block is given, monitor block count
	go func() {
		chainHeadCh := make(chan core.ChainHeadEvent, 1)
		headSub := p.eth.BlockChain().SubscribeChainHeadEvent(chainHeadCh)
		defer headSub.Unsubscribe()
		stopChan, stopSubscription := ptype.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case head := <-chainHeadCh:
				if p.eth.BlockChain().Config().IsQIP714(head.Block.Number()) {
					pcore.SetQIP714BlockReached()
					return
				}
			case <-stopChan:
				return
			}
		}
	}()
	return nil
}

func (p *PermissionCtrl) instantiateCache(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int) {
	// instantiate the cache objects for permissions
	pcore.OrgInfoMap = pcore.NewOrgCache(orgCacheSize)
	pcore.OrgInfoMap.PopulateCacheFunc(p.populateOrgToCache)

	pcore.RoleInfoMap = pcore.NewRoleCache(roleCacheSize)
	pcore.RoleInfoMap.PopulateCacheFunc(p.populateRoleToCache)

	pcore.NodeInfoMap = pcore.NewNodeCache(nodeCacheSize)
	pcore.NodeInfoMap.PopulateCacheFunc(p.populateNodeCache)
	pcore.NodeInfoMap.PopulateValidateFunc(p.populateNodeCacheAndValidate)

	pcore.AcctInfoMap = pcore.NewAcctCache(accountCacheSize)
	pcore.AcctInfoMap.PopulateCacheFunc(p.populateAccountToCache)
}

// Thus function checks if the initial network boot up status and if no
// populates permissions model with details from permission-config.json
func (p *PermissionCtrl) populateInitPermissions(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize int) error {
	p.instantiateCache(orgCacheSize, roleCacheSize, nodeCacheSize, accountCacheSize)
	networkInitialized, err := p.contract.GetNetworkBootStatus()
	if err != nil {
		// handle the scenario of no contract code.
		log.Warn("Failed to retrieve network boot status ", "err", err)
		return err
	}

	if !networkInitialized {
		p.backend.MonitorNetworkBootUp()
		if err := p.bootupNetwork(); err != nil {
			return err
		}
	} else {
		//populate orgs, nodes, roles and accounts from contract
		for _, f := range []func() error{
			p.populateOrgsFromContract,
			p.populateNodesFromContract,
			p.populateRolesFromContract,
			p.populateAccountsFromContract,
		} {
			if err := f(); err != nil {
				return err
			}
		}
		pcore.SetNetworkBootUpCompleted()
	}
	return nil
}

// initialize the permissions model and populate initial values
func (p *PermissionCtrl) bootupNetwork() error {
	if _, err := p.contract.SetPolicy(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, p.permConfig.OrgAdminRole); err != nil {
		log.Error("bootupNetwork SetPolicy failed", "err", err)
		return err
	}
	if _, err := p.contract.Init(p.permConfig.SubOrgBreadth, p.permConfig.SubOrgDepth); err != nil {
		log.Error("bootupNetwork init failed", "err", err)
		return err
	}

	pcore.OrgInfoMap.UpsertOrg(p.permConfig.NwAdminOrg, "", p.permConfig.NwAdminOrg, big.NewInt(1), pcore.OrgApproved)
	pcore.RoleInfoMap.UpsertRole(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, true, true, pcore.FullAccess, true)
	// populate the initial Node list from static-nodes.json
	if err := p.populateStaticNodesToContract(); err != nil {
		return err
	}
	// populate initial account access to full access
	if err := p.populateInitAccountAccess(); err != nil {
		return err
	}

	// update network status to boot completed
	if err := p.updateNetworkStatus(); err != nil {
		log.Error("failed to updated network boot status", "error", err)
		return err
	}
	return nil
}

// populates the account access details from contract into cache
func (p *PermissionCtrl) populateAccountsFromContract() error {
	if numberOfRoles, err := p.contract.GetNumberOfAccounts(); err == nil {
		iOrgNum := numberOfRoles.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if addr, org, role, status, orgAdmin, err := p.contract.GetAccountDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				pcore.AcctInfoMap.UpsertAccount(org, role, addr, orgAdmin, pcore.AcctStatus(int(status.Int64())))
			}
		}
	} else {
		return err
	}
	return nil
}

// populates the role details from contract into cache
func (p *PermissionCtrl) populateRolesFromContract() error {
	if numberOfRoles, err := p.contract.GetNumberOfRoles(); err == nil {
		iOrgNum := numberOfRoles.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if roleStruct, err := p.contract.GetRoleDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				pcore.RoleInfoMap.UpsertRole(roleStruct.OrgId, roleStruct.RoleId, roleStruct.Voter, roleStruct.Admin, pcore.AccessType(int(roleStruct.AccessType.Int64())), roleStruct.Active)
			}
		}

	} else {
		return err
	}
	return nil
}

// populates the Node details from contract into cache
func (p *PermissionCtrl) populateNodesFromContract() error {
	if numberOfNodes, err := p.contract.GetNumberOfNodes(); err == nil {
		iOrgNum := numberOfNodes.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if orgId, url, status, err := p.contract.GetNodeDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				pcore.NodeInfoMap.UpsertNode(orgId, url, pcore.NodeStatus(int(status.Int64())))
			}
		}
	} else {
		return err
	}
	return nil
}

// populates the org details from contract into cache
func (p *PermissionCtrl) populateOrgsFromContract() error {

	if numberOfOrgs, err := p.contract.GetNumberOfOrgs(); err == nil {
		iOrgNum := numberOfOrgs.Uint64()
		for k := uint64(0); k < iOrgNum; k++ {
			if orgId, porgId, ultParent, level, status, err := p.contract.GetOrgInfo(big.NewInt(int64(k))); err == nil {
				pcore.OrgInfoMap.UpsertOrg(orgId, porgId, ultParent, level, pcore.OrgStatus(int(status.Int64())))
			}
		}
	} else {
		return err
	}
	return nil
}

// Reads the node list from static-nodes.json and populates into the contract
func (p *PermissionCtrl) populateStaticNodesToContract() error {
	nodes := p.node.Server().Config.StaticNodes
	for _, node := range nodes {
		url := pcore.GetNodeUrl(node.EnodeID(), node.IP().String(), uint16(node.TCP()), uint16(node.RaftPort()), p.isRaft)
		_, err := p.contract.AddAdminNode(url)
		if err != nil {
			log.Warn("Failed to propose node", "err", err, "enode", node.EnodeID())
			return err
		}
		pcore.NodeInfoMap.UpsertNode(p.permConfig.NwAdminOrg, url, 2)
	}
	return nil
}

// Invokes the initAccounts function of smart contract to set the initial
// set of accounts access to full access
func (p *PermissionCtrl) populateInitAccountAccess() error {
	for _, a := range p.permConfig.Accounts {
		_, er := p.contract.AddAdminAccount(a)
		if er != nil {
			log.Warn("Error adding permission initial account list", "err", er, "account", a)
			return er
		}
		pcore.AcctInfoMap.UpsertAccount(p.permConfig.NwAdminOrg, p.permConfig.NwAdminRole, a, true, 2)
	}
	return nil
}

// updates network boot status to true
func (p *PermissionCtrl) updateNetworkStatus() error {
	_, err := p.contract.UpdateNetworkBootStatus()
	if err != nil {
		log.Warn("Failed to udpate network boot status ", "err", err)
		return err
	}
	return nil
}

// getter to get an account record from the contract
func (p *PermissionCtrl) populateAccountToCache(acctId common.Address) (*pcore.AccountInfo, error) {
	account, orgId, roleId, status, isAdmin, err := p.contract.GetAccountDetails(acctId)
	if err != nil {
		return nil, err
	}

	if status.Int64() == 0 {
		return nil, ptype.ErrAccountNotThere
	}
	return &pcore.AccountInfo{AcctId: account, OrgId: orgId, RoleId: roleId, Status: pcore.AcctStatus(status.Int64()), IsOrgAdmin: isAdmin}, nil
}

// getter to get a org record from the contract
func (p *PermissionCtrl) populateOrgToCache(orgId string) (*pcore.OrgInfo, error) {
	org, parentOrgId, ultimateParentId, orgLevel, orgStatus, err := p.contract.GetOrgDetails(orgId)
	if err != nil {
		return nil, err
	}
	if orgStatus.Int64() == 0 {
		return nil, ptype.ErrOrgDoesNotExists
	}
	orgInfo := pcore.OrgInfo{OrgId: org, ParentOrgId: parentOrgId, UltimateParent: ultimateParentId, Status: pcore.OrgStatus(orgStatus.Int64()), Level: orgLevel}
	// now need to build the list of sub orgs for this org
	subOrgIndexes, err := p.contract.GetSubOrgIndexes(orgId)
	if err != nil {
		return nil, err
	}

	if len(subOrgIndexes) == 0 {
		return &orgInfo, nil
	}

	// range through the sub org indexes and get the org ids to populate the suborg list
	for _, s := range subOrgIndexes {
		subOrgId, _, _, _, _, err := p.contract.GetOrgInfo(s)

		if err != nil {
			return nil, err
		}
		orgInfo.SubOrgList = append(orgInfo.SubOrgList, orgId+"."+subOrgId)

	}
	return &orgInfo, nil
}

// getter to get a role record from the contract
func (p *PermissionCtrl) populateRoleToCache(roleKey *pcore.RoleKey) (*pcore.RoleInfo, error) {
	roleDetails, err := p.contract.GetRoleDetails(roleKey.RoleId, roleKey.OrgId)

	if err != nil {
		return nil, err
	}

	if roleDetails.OrgId == "" {
		return nil, ptype.ErrInvalidRole
	}
	return &pcore.RoleInfo{OrgId: roleDetails.OrgId, RoleId: roleDetails.RoleId, IsVoter: roleDetails.Voter, IsAdmin: roleDetails.Admin, Access: pcore.AccessType(roleDetails.AccessType.Int64()), Active: roleDetails.Active}, nil
}

// getter to get a role record from the contract
func (p *PermissionCtrl) populateNodeCache(url string) (*pcore.NodeInfo, error) {
	orgId, url, status, err := p.contract.GetNodeDetails(url)
	if err != nil {
		return nil, err
	}

	if status.Int64() == 0 {
		return nil, ptype.ErrNodeDoesNotExists
	}
	return &pcore.NodeInfo{OrgId: orgId, Url: url, Status: pcore.NodeStatus(status.Int64())}, nil
}

// getter to get a Node record from the contract
func (p *PermissionCtrl) populateNodeCacheAndValidate(hexNodeId, ultimateParentId string) bool {
	txnAllowed := false
	passedEnode, _ := enode.ParseV4(hexNodeId)
	if numberOfNodes, err := p.contract.GetNumberOfNodes(); err == nil {
		numNodes := numberOfNodes.Uint64()
		for k := uint64(0); k < numNodes; k++ {
			if orgId, url, status, err := p.contract.GetNodeDetailsFromIndex(big.NewInt(int64(k))); err == nil {
				if orgRec, err := pcore.OrgInfoMap.GetOrg(orgId); err != nil {
					if orgRec.UltimateParent == ultimateParentId {
						recEnode, _ := enode.ParseV4(url)
						if recEnode.ID() == passedEnode.ID() {
							txnAllowed = true
							pcore.NodeInfoMap.UpsertNode(orgId, url, pcore.NodeStatus(int(status.Int64())))
						}
					}
				}
			}
		}
	}
	return txnAllowed
}
