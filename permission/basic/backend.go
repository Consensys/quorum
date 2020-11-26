package basic

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	bb "github.com/ethereum/go-ethereum/permission/basic/bind"
	"github.com/ethereum/go-ethereum/permission/cache"
	ptype "github.com/ethereum/go-ethereum/permission/types"
)

type Backend struct {
	Ib    ptype.InterfaceBackend
	Contr *Init
}

func (b *Backend) ManageAccountPermissions() error {
	chAccessModified := make(chan *bb.AcctManagerAccountAccessModified)
	chAccessRevoked := make(chan *bb.AcctManagerAccountAccessRevoked)
	chStatusChanged := make(chan *bb.AcctManagerAccountStatusChanged)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := b.Contr.PermAcct.AcctManagerFilterer.WatchAccountAccessModified(opts, chAccessModified); err != nil {
		return fmt.Errorf("failed AccountAccessModified: %v", err)
	}

	if _, err := b.Contr.PermAcct.AcctManagerFilterer.WatchAccountAccessRevoked(opts, chAccessRevoked); err != nil {
		return fmt.Errorf("failed AccountAccessRevoked: %v", err)
	}

	if _, err := b.Contr.PermAcct.AcctManagerFilterer.WatchAccountStatusChanged(opts, chStatusChanged); err != nil {
		return fmt.Errorf("failed AccountStatusChanged: %v", err)
	}

	go func() {
		stopChan, stopSubscription := ptype.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtAccessModified := <-chAccessModified:
				cache.AcctInfoMap.UpsertAccount(evtAccessModified.OrgId, evtAccessModified.RoleId, evtAccessModified.Account, evtAccessModified.OrgAdmin, cache.AcctStatus(int(evtAccessModified.Status.Uint64())))

			case evtAccessRevoked := <-chAccessRevoked:
				cache.AcctInfoMap.UpsertAccount(evtAccessRevoked.OrgId, evtAccessRevoked.RoleId, evtAccessRevoked.Account, evtAccessRevoked.OrgAdmin, cache.AcctActive)

			case evtStatusChanged := <-chStatusChanged:
				if ac, err := cache.AcctInfoMap.GetAccount(evtStatusChanged.Account); ac != nil {
					cache.AcctInfoMap.UpsertAccount(evtStatusChanged.OrgId, ac.RoleId, evtStatusChanged.Account, ac.IsOrgAdmin, cache.AcctStatus(int(evtStatusChanged.Status.Uint64())))
				} else {
					log.Info("error fetching account information", "err", err)
				}
			case <-stopChan:
				log.Info("quit account Contr watch")
				return
			}
		}
	}()
	return nil
}

func (b *Backend) ManageRolePermissions() error {
	chRoleCreated := make(chan *bb.RoleManagerRoleCreated, 1)
	chRoleRevoked := make(chan *bb.RoleManagerRoleRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	contract := b.Contr

	if _, err := contract.PermRole.RoleManagerFilterer.WatchRoleCreated(opts, chRoleCreated); err != nil {
		return fmt.Errorf("failed WatchRoleCreated: %v", err)
	}

	if _, err := contract.PermRole.RoleManagerFilterer.WatchRoleRevoked(opts, chRoleRevoked); err != nil {
		return fmt.Errorf("failed WatchRoleRevoked: %v", err)
	}

	go func() {
		stopChan, stopSubscription := ptype.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtRoleCreated := <-chRoleCreated:
				cache.RoleInfoMap.UpsertRole(evtRoleCreated.OrgId, evtRoleCreated.RoleId, evtRoleCreated.IsVoter, evtRoleCreated.IsAdmin, cache.AccessType(int(evtRoleCreated.BaseAccess.Uint64())), true)

			case evtRoleRevoked := <-chRoleRevoked:
				if r, _ := cache.RoleInfoMap.GetRole(evtRoleRevoked.OrgId, evtRoleRevoked.RoleId); r != nil {
					cache.RoleInfoMap.UpsertRole(evtRoleRevoked.OrgId, evtRoleRevoked.RoleId, r.IsVoter, r.IsAdmin, r.Access, false)
				} else {
					log.Error("Revoke role - cache is missing role", "org", evtRoleRevoked.OrgId, "role", evtRoleRevoked.RoleId)
				}
			case <-stopChan:
				log.Info("quit role Contr watch")
				return
			}
		}
	}()
	return nil
}

func (b *Backend) ManageOrgPermissions() error {
	chPendingApproval := make(chan *bb.OrgManagerOrgPendingApproval, 1)
	chOrgApproved := make(chan *bb.OrgManagerOrgApproved, 1)
	chOrgSuspended := make(chan *bb.OrgManagerOrgSuspended, 1)
	chOrgReactivated := make(chan *bb.OrgManagerOrgSuspensionRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	contract := b.Contr

	if _, err := contract.PermOrg.OrgManagerFilterer.WatchOrgPendingApproval(opts, chPendingApproval); err != nil {
		return fmt.Errorf("failed WatchOrgPendingApproval: %v", err)
	}

	if _, err := contract.PermOrg.OrgManagerFilterer.WatchOrgApproved(opts, chOrgApproved); err != nil {
		return fmt.Errorf("failed WatchOrgApproved: %v", err)
	}

	if _, err := contract.PermOrg.OrgManagerFilterer.WatchOrgSuspended(opts, chOrgSuspended); err != nil {
		return fmt.Errorf("failed WatchOrgSuspended: %v", err)
	}

	if _, err := contract.PermOrg.OrgManagerFilterer.WatchOrgSuspensionRevoked(opts, chOrgReactivated); err != nil {
		return fmt.Errorf("failed WatchOrgSuspensionRevoked: %v", err)
	}

	go func() {
		stopChan, stopSubscription := ptype.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtPendingApproval := <-chPendingApproval:
				cache.OrgInfoMap.UpsertOrg(evtPendingApproval.OrgId, evtPendingApproval.PorgId, evtPendingApproval.UltParent, evtPendingApproval.Level, cache.OrgStatus(evtPendingApproval.Status.Uint64()))

			case evtOrgApproved := <-chOrgApproved:
				cache.OrgInfoMap.UpsertOrg(evtOrgApproved.OrgId, evtOrgApproved.PorgId, evtOrgApproved.UltParent, evtOrgApproved.Level, cache.OrgApproved)

			case evtOrgSuspended := <-chOrgSuspended:
				cache.OrgInfoMap.UpsertOrg(evtOrgSuspended.OrgId, evtOrgSuspended.PorgId, evtOrgSuspended.UltParent, evtOrgSuspended.Level, cache.OrgSuspended)

			case evtOrgReactivated := <-chOrgReactivated:
				cache.OrgInfoMap.UpsertOrg(evtOrgReactivated.OrgId, evtOrgReactivated.PorgId, evtOrgReactivated.UltParent, evtOrgReactivated.Level, cache.OrgApproved)
			case <-stopChan:
				log.Info("quit org Contr watch")
				return
			}
		}
	}()
	return nil
}

func (b *Backend) ManageNodePermissions() error {
	chNodeApproved := make(chan *bb.NodeManagerNodeApproved, 1)
	chNodeProposed := make(chan *bb.NodeManagerNodeProposed, 1)
	chNodeDeactivated := make(chan *bb.NodeManagerNodeDeactivated, 1)
	chNodeActivated := make(chan *bb.NodeManagerNodeActivated, 1)
	chNodeBlacklisted := make(chan *bb.NodeManagerNodeBlacklisted)
	chNodeRecoveryInit := make(chan *bb.NodeManagerNodeRecoveryInitiated, 1)
	chNodeRecoveryDone := make(chan *bb.NodeManagerNodeRecoveryCompleted, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	contract := b.Contr

	if _, err := contract.PermNode.NodeManagerFilterer.WatchNodeApproved(opts, chNodeApproved); err != nil {
		return fmt.Errorf("failed WatchNodeApproved: %v", err)
	}

	if _, err := contract.PermNode.NodeManagerFilterer.WatchNodeProposed(opts, chNodeProposed); err != nil {
		return fmt.Errorf("failed WatchNodeProposed: %v", err)
	}

	if _, err := contract.PermNode.NodeManagerFilterer.WatchNodeDeactivated(opts, chNodeDeactivated); err != nil {
		return fmt.Errorf("failed NodeDeactivated: %v", err)
	}
	if _, err := contract.PermNode.NodeManagerFilterer.WatchNodeActivated(opts, chNodeActivated); err != nil {
		return fmt.Errorf("failed WatchNodeActivated: %v", err)
	}

	if _, err := contract.PermNode.NodeManagerFilterer.WatchNodeBlacklisted(opts, chNodeBlacklisted); err != nil {
		return fmt.Errorf("failed NodeBlacklisting: %v", err)
	}

	if _, err := contract.PermNode.NodeManagerFilterer.WatchNodeRecoveryInitiated(opts, chNodeRecoveryInit); err != nil {
		return fmt.Errorf("failed NodeRecoveryInitiated: %v", err)
	}

	if _, err := contract.PermNode.NodeManagerFilterer.WatchNodeRecoveryCompleted(opts, chNodeRecoveryDone); err != nil {
		return fmt.Errorf("failed NodeRecoveryCompleted: %v", err)
	}

	go func() {
		stopChan, stopSubscription := ptype.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtNodeApproved := <-chNodeApproved:
				err := ptype.UpdatePermissionedNodes(b.Ib.Node(), b.Ib.DataDir(), evtNodeApproved.EnodeId, ptype.NodeAdd, b.Ib.IsRaft())
				if err != nil {
					log.Error("error updating permissioned-nodes.json", "err", err)
				}
				cache.NodeInfoMap.UpsertNode(evtNodeApproved.OrgId, evtNodeApproved.EnodeId, cache.NodeApproved)

			case evtNodeProposed := <-chNodeProposed:
				cache.NodeInfoMap.UpsertNode(evtNodeProposed.OrgId, evtNodeProposed.EnodeId, cache.NodePendingApproval)

			case evtNodeDeactivated := <-chNodeDeactivated:
				err := ptype.UpdatePermissionedNodes(b.Ib.Node(), b.Ib.DataDir(), evtNodeDeactivated.EnodeId, ptype.NodeDelete, b.Ib.IsRaft())
				if err != nil {
					log.Error("error updating permissioned-nodes.json", "err", err)
				}
				cache.NodeInfoMap.UpsertNode(evtNodeDeactivated.OrgId, evtNodeDeactivated.EnodeId, cache.NodeDeactivated)

			case evtNodeActivated := <-chNodeActivated:
				err := ptype.UpdatePermissionedNodes(b.Ib.Node(), b.Ib.DataDir(), evtNodeActivated.EnodeId, ptype.NodeAdd, b.Ib.IsRaft())
				if err != nil {
					log.Error("error updating permissioned-nodes.json", "err", err)
				}
				cache.NodeInfoMap.UpsertNode(evtNodeActivated.OrgId, evtNodeActivated.EnodeId, cache.NodeApproved)

			case evtNodeBlacklisted := <-chNodeBlacklisted:
				cache.NodeInfoMap.UpsertNode(evtNodeBlacklisted.OrgId, evtNodeBlacklisted.EnodeId, cache.NodeBlackListed)
				err := ptype.UpdateDisallowedNodes(b.Ib.DataDir(), evtNodeBlacklisted.EnodeId, ptype.NodeAdd)
				log.Error("error updating disallowed-nodes.json", "err", err)
				err = ptype.UpdatePermissionedNodes(b.Ib.Node(), b.Ib.DataDir(), evtNodeBlacklisted.EnodeId, ptype.NodeDelete, b.Ib.IsRaft())
				if err != nil {
					log.Error("error updating permissioned-nodes.json", "err", err)
				}

			case evtNodeRecoveryInit := <-chNodeRecoveryInit:
				cache.NodeInfoMap.UpsertNode(evtNodeRecoveryInit.OrgId, evtNodeRecoveryInit.EnodeId, cache.NodeRecoveryInitiated)

			case evtNodeRecoveryDone := <-chNodeRecoveryDone:
				cache.NodeInfoMap.UpsertNode(evtNodeRecoveryDone.OrgId, evtNodeRecoveryDone.EnodeId, cache.NodeApproved)
				err := ptype.UpdateDisallowedNodes(b.Ib.DataDir(), evtNodeRecoveryDone.EnodeId, ptype.NodeDelete)
				log.Error("error updating disallowed-nodes.json", "err", err)
				err = ptype.UpdatePermissionedNodes(b.Ib.Node(), b.Ib.DataDir(), evtNodeRecoveryDone.EnodeId, ptype.NodeAdd, b.Ib.IsRaft())
				if err != nil {
					log.Error("error updating permissioned-nodes.json", "err", err)
				}

			case <-stopChan:
				log.Info("quit Node Contr watch")
				return
			}
		}
	}()
	return nil
}

func (b *Backend) MonitorNetworkBootUp() error {
	netWorkBootCh := make(chan *bb.PermImplPermissionsInitialized, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber

	if _, err := b.Contr.PermImpl.PermImplFilterer.WatchPermissionsInitialized(opts, netWorkBootCh); err != nil {
		return fmt.Errorf("failed WatchPermissionsInitialized: %v", err)
	}

	go func() {
		stopChan, stopSubscription := ptype.SubscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtMetworkBootUpCompleted := <-netWorkBootCh:
				if evtMetworkBootUpCompleted.NetworkBootStatus {
					cache.SetNetworkBootUpCompleted()
				}
				return
			case <-stopChan:
				log.Info("quit implementation contract network boot watch")
				return
			}
		}
	}()
	return nil
}

func getInterfaceContractSession(permInterfaceInstance *bb.PermInterface, contractAddress common.Address, backend bind.ContractBackend) (*bb.PermInterfaceSession, error) {
	if err := ptype.BindContract(&permInterfaceInstance, func() (interface{}, error) { return bb.NewPermInterface(contractAddress, backend) }); err != nil {
		return nil, err
	}
	ps := &bb.PermInterfaceSession{
		Contract: permInterfaceInstance,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	return ps, nil
}

func getBackend(contractBackend ptype.ContractBackend) (*Basic, error) {
	basicBackend := Basic{ContractBackend: contractBackend}
	ps, err := getInterfaceContractSession(basicBackend.PermInterf, contractBackend.PermConfig.InterfAddress, contractBackend.EthClnt)
	if err != nil {
		return nil, err
	}
	basicBackend.PermInterfSession = ps
	return &basicBackend, nil
}

func getBackendWithTransactOpts(contractBackend ptype.ContractBackend, transactOpts *bind.TransactOpts) (*Basic, error) {
	basicBackend, err := getBackend(contractBackend)
	if err != nil {
		return nil, err
	}
	basicBackend.PermInterfSession.TransactOpts = *transactOpts

	return basicBackend, nil
}

func (b *Backend) GetRoleService(transactOpts *bind.TransactOpts, roleBackend ptype.ContractBackend) (ptype.RoleService, error) {
	backEnd, err := getBackendWithTransactOpts(roleBackend, transactOpts)
	if err != nil {
		return nil, err
	}
	return &Role{Backend: backEnd}, nil

}

func (b *Backend) GetOrgService(transactOpts *bind.TransactOpts, orgBackend ptype.ContractBackend) (ptype.OrgService, error) {
	backEnd, err := getBackendWithTransactOpts(orgBackend, transactOpts)
	if err != nil {
		return nil, err
	}
	return &Org{Backend: backEnd}, nil

}

func (b *Backend) GetNodeService(transactOpts *bind.TransactOpts, nodeBackend ptype.ContractBackend) (ptype.NodeService, error) {
	backEnd, err := getBackendWithTransactOpts(nodeBackend, transactOpts)
	if err != nil {
		return nil, err
	}
	return &Node{Backend: backEnd}, nil

}

func (b *Backend) GetAccountService(transactOpts *bind.TransactOpts, accountBackend ptype.ContractBackend) (ptype.AccountService, error) {
	backEnd, err := getBackendWithTransactOpts(accountBackend, transactOpts)
	if err != nil {
		return nil, err
	}
	return &Account{Backend: backEnd}, nil

}

func (b *Backend) GetAuditService(auditBackend ptype.ContractBackend) (ptype.AuditService, error) {
	backEnd, err := getBackend(auditBackend)
	if err != nil {
		return nil, err
	}
	return &Audit{Backend: backEnd}, nil

}

func (b *Backend) GetControlService(controlBackend ptype.ContractBackend) (ptype.ControlService, error) {
	return &Control{}, nil
}
