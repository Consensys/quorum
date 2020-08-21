package permission

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	basic "github.com/ethereum/go-ethereum/permission/basic/bind"
)

type backendBasic struct {
	pc *PermissionCtrl
}

func (b *backendBasic) manageAccountPermissions() error {
	p := b.pc
	chAccessModified := make(chan *basic.AcctManagerAccountAccessModified)
	chAccessRevoked := make(chan *basic.AcctManagerAccountAccessRevoked)
	chStatusChanged := make(chan *basic.AcctManagerAccountStatusChanged)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var contract *BasicContract
	var ok bool
	if contract, ok = p.contract.(*BasicContract); !ok {
		return fmt.Errorf("error casting permission contract service to basic contract")
	}

	if _, err := contract.permAcct.AcctManagerFilterer.WatchAccountAccessModified(opts, chAccessModified); err != nil {
		return fmt.Errorf("failed AccountAccessModified: %v", err)
	}

	if _, err := contract.permAcct.AcctManagerFilterer.WatchAccountAccessRevoked(opts, chAccessRevoked); err != nil {
		return fmt.Errorf("failed AccountAccessRevoked: %v", err)
	}

	if _, err := contract.permAcct.AcctManagerFilterer.WatchAccountStatusChanged(opts, chStatusChanged); err != nil {
		return fmt.Errorf("failed AccountStatusChanged: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.subscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtAccessModified := <-chAccessModified:
				types.AcctInfoMap.UpsertAccount(evtAccessModified.OrgId, evtAccessModified.RoleId, evtAccessModified.Account, evtAccessModified.OrgAdmin, types.AcctStatus(int(evtAccessModified.Status.Uint64())))

			case evtAccessRevoked := <-chAccessRevoked:
				types.AcctInfoMap.UpsertAccount(evtAccessRevoked.OrgId, evtAccessRevoked.RoleId, evtAccessRevoked.Account, evtAccessRevoked.OrgAdmin, types.AcctActive)

			case evtStatusChanged := <-chStatusChanged:
				if ac, err := types.AcctInfoMap.GetAccount(evtStatusChanged.Account); ac != nil {
					types.AcctInfoMap.UpsertAccount(evtStatusChanged.OrgId, ac.RoleId, evtStatusChanged.Account, ac.IsOrgAdmin, types.AcctStatus(int(evtStatusChanged.Status.Uint64())))
				} else {
					log.Info("error fetching account information", "err", err)
				}
			case <-stopChan:
				log.Info("quit account contract watch")
				return
			}
		}
	}()
	return nil
}

func (b *backendBasic) manageRolePermissions() error {
	p := b.pc
	chRoleCreated := make(chan *basic.RoleManagerRoleCreated, 1)
	chRoleRevoked := make(chan *basic.RoleManagerRoleRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var contract *BasicContract
	var ok bool
	if contract, ok = p.contract.(*BasicContract); !ok {
		return fmt.Errorf("error casting permission contract service to basic contract")
	}
	if _, err := contract.permRole.RoleManagerFilterer.WatchRoleCreated(opts, chRoleCreated); err != nil {
		return fmt.Errorf("failed WatchRoleCreated: %v", err)
	}

	if _, err := contract.permRole.RoleManagerFilterer.WatchRoleRevoked(opts, chRoleRevoked); err != nil {
		return fmt.Errorf("failed WatchRoleRemoved: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.subscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtRoleCreated := <-chRoleCreated:
				types.RoleInfoMap.UpsertRole(evtRoleCreated.OrgId, evtRoleCreated.RoleId, evtRoleCreated.IsVoter, evtRoleCreated.IsAdmin, types.AccessType(int(evtRoleCreated.BaseAccess.Uint64())), true)

			case evtRoleRevoked := <-chRoleRevoked:
				if r, _ := types.RoleInfoMap.GetRole(evtRoleRevoked.OrgId, evtRoleRevoked.RoleId); r != nil {
					types.RoleInfoMap.UpsertRole(evtRoleRevoked.OrgId, evtRoleRevoked.RoleId, r.IsVoter, r.IsAdmin, r.Access, false)
				} else {
					log.Error("Revoke role - cache is missing role", "org", evtRoleRevoked.OrgId, "role", evtRoleRevoked.RoleId)
				}
			case <-stopChan:
				log.Info("quit role contract watch")
				return
			}
		}
	}()
	return nil
}

func (b *backendBasic) manageOrgPermissions() error {
	p := b.pc
	chPendingApproval := make(chan *basic.OrgManagerOrgPendingApproval, 1)
	chOrgApproved := make(chan *basic.OrgManagerOrgApproved, 1)
	chOrgSuspended := make(chan *basic.OrgManagerOrgSuspended, 1)
	chOrgReactivated := make(chan *basic.OrgManagerOrgSuspensionRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var contract *BasicContract
	var ok bool
	if contract, ok = p.contract.(*BasicContract); !ok {
		return fmt.Errorf("error casting permission contract service to basic contract")
	}
	if _, err := contract.permOrg.OrgManagerFilterer.WatchOrgPendingApproval(opts, chPendingApproval); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := contract.permOrg.OrgManagerFilterer.WatchOrgApproved(opts, chOrgApproved); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := contract.permOrg.OrgManagerFilterer.WatchOrgSuspended(opts, chOrgSuspended); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := contract.permOrg.OrgManagerFilterer.WatchOrgSuspensionRevoked(opts, chOrgReactivated); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.subscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtPendingApproval := <-chPendingApproval:
				types.OrgInfoMap.UpsertOrg(evtPendingApproval.OrgId, evtPendingApproval.PorgId, evtPendingApproval.UltParent, evtPendingApproval.Level, types.OrgStatus(evtPendingApproval.Status.Uint64()))

			case evtOrgApproved := <-chOrgApproved:
				types.OrgInfoMap.UpsertOrg(evtOrgApproved.OrgId, evtOrgApproved.PorgId, evtOrgApproved.UltParent, evtOrgApproved.Level, types.OrgApproved)

			case evtOrgSuspended := <-chOrgSuspended:
				types.OrgInfoMap.UpsertOrg(evtOrgSuspended.OrgId, evtOrgSuspended.PorgId, evtOrgSuspended.UltParent, evtOrgSuspended.Level, types.OrgSuspended)

			case evtOrgReactivated := <-chOrgReactivated:
				types.OrgInfoMap.UpsertOrg(evtOrgReactivated.OrgId, evtOrgReactivated.PorgId, evtOrgReactivated.UltParent, evtOrgReactivated.Level, types.OrgApproved)
			case <-stopChan:
				log.Info("quit org contract watch")
				return
			}
		}
	}()
	return nil
}

func (b *backendBasic) manageNodePermissions() error {
	p := b.pc
	chNodeApproved := make(chan *basic.NodeManagerNodeApproved, 1)
	chNodeProposed := make(chan *basic.NodeManagerNodeProposed, 1)
	chNodeDeactivated := make(chan *basic.NodeManagerNodeDeactivated, 1)
	chNodeActivated := make(chan *basic.NodeManagerNodeActivated, 1)
	chNodeBlacklisted := make(chan *basic.NodeManagerNodeBlacklisted)
	chNodeRecoveryInit := make(chan *basic.NodeManagerNodeRecoveryInitiated, 1)
	chNodeRecoveryDone := make(chan *basic.NodeManagerNodeRecoveryCompleted, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var contract *BasicContract
	var ok bool
	if contract, ok = p.contract.(*BasicContract); !ok {
		return fmt.Errorf("error casting permission contract service to basic contract")
	}

	if _, err := contract.permNode.NodeManagerFilterer.WatchNodeApproved(opts, chNodeApproved); err != nil {
		return fmt.Errorf("failed WatchNodeApproved: %v", err)
	}

	if _, err := contract.permNode.NodeManagerFilterer.WatchNodeProposed(opts, chNodeProposed); err != nil {
		return fmt.Errorf("failed WatchNodeProposed: %v", err)
	}

	if _, err := contract.permNode.NodeManagerFilterer.WatchNodeDeactivated(opts, chNodeDeactivated); err != nil {
		return fmt.Errorf("failed NodeDeactivated: %v", err)
	}
	if _, err := contract.permNode.NodeManagerFilterer.WatchNodeActivated(opts, chNodeActivated); err != nil {
		return fmt.Errorf("failed WatchNodeActivated: %v", err)
	}

	if _, err := contract.permNode.NodeManagerFilterer.WatchNodeBlacklisted(opts, chNodeBlacklisted); err != nil {
		return fmt.Errorf("failed NodeBlacklisting: %v", err)
	}

	if _, err := contract.permNode.NodeManagerFilterer.WatchNodeRecoveryInitiated(opts, chNodeRecoveryInit); err != nil {
		return fmt.Errorf("failed NodeRecoveryInitiated: %v", err)
	}

	if _, err := contract.permNode.NodeManagerFilterer.WatchNodeRecoveryCompleted(opts, chNodeRecoveryDone); err != nil {
		return fmt.Errorf("failed NodeRecoveryCompleted: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.subscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtNodeApproved := <-chNodeApproved:
				p.updatePermissionedNodes(evtNodeApproved.EnodeId, NodeAdd)
				types.NodeInfoMap.UpsertNode(evtNodeApproved.OrgId, evtNodeApproved.EnodeId, types.NodeApproved)

			case evtNodeProposed := <-chNodeProposed:
				types.NodeInfoMap.UpsertNode(evtNodeProposed.OrgId, evtNodeProposed.EnodeId, types.NodePendingApproval)

			case evtNodeDeactivated := <-chNodeDeactivated:
				p.updatePermissionedNodes(evtNodeDeactivated.EnodeId, NodeDelete)
				types.NodeInfoMap.UpsertNode(evtNodeDeactivated.OrgId, evtNodeDeactivated.EnodeId, types.NodeDeactivated)

			case evtNodeActivated := <-chNodeActivated:
				p.updatePermissionedNodes(evtNodeActivated.EnodeId, NodeAdd)
				types.NodeInfoMap.UpsertNode(evtNodeActivated.OrgId, evtNodeActivated.EnodeId, types.NodeApproved)

			case evtNodeBlacklisted := <-chNodeBlacklisted:
				types.NodeInfoMap.UpsertNode(evtNodeBlacklisted.OrgId, evtNodeBlacklisted.EnodeId, types.NodeBlackListed)
				p.updateDisallowedNodes(evtNodeBlacklisted.EnodeId, NodeAdd)
				p.updatePermissionedNodes(evtNodeBlacklisted.EnodeId, NodeDelete)

			case evtNodeRecoveryInit := <-chNodeRecoveryInit:
				types.NodeInfoMap.UpsertNode(evtNodeRecoveryInit.OrgId, evtNodeRecoveryInit.EnodeId, types.NodeRecoveryInitiated)

			case evtNodeRecoveryDone := <-chNodeRecoveryDone:
				types.NodeInfoMap.UpsertNode(evtNodeRecoveryDone.OrgId, evtNodeRecoveryDone.EnodeId, types.NodeApproved)
				p.updateDisallowedNodes(evtNodeRecoveryDone.EnodeId, NodeDelete)
				p.updatePermissionedNodes(evtNodeRecoveryDone.EnodeId, NodeAdd)

			case <-stopChan:
				log.Info("quit Node contract watch")
				return
			}
		}
	}()
	return nil
}
