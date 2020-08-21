package permission

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	eea "github.com/ethereum/go-ethereum/permission/eea/bind"
)

type backendEea struct {
	pc *PermissionCtrl
}

func (b *backendEea) manageAccountPermissions() error {
	p := b.pc
	chAccessModified := make(chan *eea.EeaAcctManagerAccountAccessModified)
	chAccessRevoked := make(chan *eea.EeaAcctManagerAccountAccessRevoked)
	chStatusChanged := make(chan *eea.EeaAcctManagerAccountStatusChanged)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var contract *PermissionContractEea
	var ok bool
	if contract, ok = p.contract.(*PermissionContractEea); !ok {
		return fmt.Errorf("error casting permission contract service to EEA contract")
	}
	if _, err := contract.permAcct.EeaAcctManagerFilterer.WatchAccountAccessModified(opts, chAccessModified); err != nil {
		return fmt.Errorf("failed AccountAccessModified: %v", err)
	}

	if _, err := contract.permAcct.EeaAcctManagerFilterer.WatchAccountAccessRevoked(opts, chAccessRevoked); err != nil {
		return fmt.Errorf("failed AccountAccessRevoked: %v", err)
	}

	if _, err := contract.permAcct.EeaAcctManagerFilterer.WatchAccountStatusChanged(opts, chStatusChanged); err != nil {
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

func (b *backendEea) manageRolePermissions() error {
	p := b.pc
	chRoleCreated := make(chan *eea.EeaRoleManagerRoleCreated, 1)
	chRoleRevoked := make(chan *eea.EeaRoleManagerRoleRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var contract *PermissionContractEea
	var ok bool
	if contract, ok = p.contract.(*PermissionContractEea); !ok {
		return fmt.Errorf("error casting permission contract service to basic contract")
	}

	if _, err := contract.permRole.EeaRoleManagerFilterer.WatchRoleCreated(opts, chRoleCreated); err != nil {
		return fmt.Errorf("failed WatchRoleCreated: %v", err)
	}

	if _, err := contract.permRole.EeaRoleManagerFilterer.WatchRoleRevoked(opts, chRoleRevoked); err != nil {
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

func (b *backendEea) manageOrgPermissions() error {
	p := b.pc
	chPendingApproval := make(chan *eea.EeaOrgManagerOrgPendingApproval, 1)
	chOrgApproved := make(chan *eea.EeaOrgManagerOrgApproved, 1)
	chOrgSuspended := make(chan *eea.EeaOrgManagerOrgSuspended, 1)
	chOrgReactivated := make(chan *eea.EeaOrgManagerOrgSuspensionRevoked, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var contract *PermissionContractEea
	var ok bool
	if contract, ok = p.contract.(*PermissionContractEea); !ok {
		return fmt.Errorf("error casting permission contract service to basic contract")
	}
	if _, err := contract.permOrg.EeaOrgManagerFilterer.WatchOrgPendingApproval(opts, chPendingApproval); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := contract.permOrg.EeaOrgManagerFilterer.WatchOrgApproved(opts, chOrgApproved); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := contract.permOrg.EeaOrgManagerFilterer.WatchOrgSuspended(opts, chOrgSuspended); err != nil {
		return fmt.Errorf("failed WatchNodePendingApproval: %v", err)
	}

	if _, err := contract.permOrg.EeaOrgManagerFilterer.WatchOrgSuspensionRevoked(opts, chOrgReactivated); err != nil {
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

func (b *backendEea) manageNodePermissions() error {
	p := b.pc
	chNodeApproved := make(chan *eea.EeaNodeManagerNodeApproved, 1)
	chNodeProposed := make(chan *eea.EeaNodeManagerNodeProposed, 1)
	chNodeDeactivated := make(chan *eea.EeaNodeManagerNodeDeactivated, 1)
	chNodeActivated := make(chan *eea.EeaNodeManagerNodeActivated, 1)
	chNodeBlacklisted := make(chan *eea.EeaNodeManagerNodeBlacklisted)
	chNodeRecoveryInit := make(chan *eea.EeaNodeManagerNodeRecoveryInitiated, 1)
	chNodeRecoveryDone := make(chan *eea.EeaNodeManagerNodeRecoveryCompleted, 1)

	opts := &bind.WatchOpts{}
	var blockNumber uint64 = 1
	opts.Start = &blockNumber
	var contract *PermissionContractEea
	var ok bool
	if contract, ok = p.contract.(*PermissionContractEea); !ok {
		return fmt.Errorf("error casting permission contract service to basic contract")
	}

	if _, err := contract.permNode.EeaNodeManagerFilterer.WatchNodeApproved(opts, chNodeApproved); err != nil {
		return fmt.Errorf("failed WatchNodeApproved: %v", err)
	}

	if _, err := contract.permNode.EeaNodeManagerFilterer.WatchNodeProposed(opts, chNodeProposed); err != nil {
		return fmt.Errorf("failed WatchNodeProposed: %v", err)
	}

	if _, err := contract.permNode.EeaNodeManagerFilterer.WatchNodeDeactivated(opts, chNodeDeactivated); err != nil {
		return fmt.Errorf("failed NodeDeactivated: %v", err)
	}
	if _, err := contract.permNode.EeaNodeManagerFilterer.WatchNodeActivated(opts, chNodeActivated); err != nil {
		return fmt.Errorf("failed WatchNodeActivated: %v", err)
	}

	if _, err := contract.permNode.EeaNodeManagerFilterer.WatchNodeBlacklisted(opts, chNodeBlacklisted); err != nil {
		return fmt.Errorf("failed NodeBlacklisting: %v", err)
	}

	if _, err := contract.permNode.EeaNodeManagerFilterer.WatchNodeRecoveryInitiated(opts, chNodeRecoveryInit); err != nil {
		return fmt.Errorf("failed NodeRecoveryInitiated: %v", err)
	}

	if _, err := contract.permNode.EeaNodeManagerFilterer.WatchNodeRecoveryCompleted(opts, chNodeRecoveryDone); err != nil {
		return fmt.Errorf("failed NodeRecoveryCompleted: %v", err)
	}

	go func() {
		stopChan, stopSubscription := p.subscribeStopEvent()
		defer stopSubscription.Unsubscribe()
		for {
			select {
			case evtNodeApproved := <-chNodeApproved:
				p.updatePermissionedNodes(types.GetNodeUrl(evtNodeApproved.EnodeId, evtNodeApproved.Ip[:], evtNodeApproved.Port, evtNodeApproved.Raftport), NodeAdd)
				types.NodeInfoMap.UpsertNode(evtNodeApproved.OrgId, types.GetNodeUrl(evtNodeApproved.EnodeId, evtNodeApproved.Ip[:], evtNodeApproved.Port, evtNodeApproved.Raftport), types.NodeApproved)

			case evtNodeProposed := <-chNodeProposed:
				types.NodeInfoMap.UpsertNode(evtNodeProposed.OrgId, types.GetNodeUrl(evtNodeProposed.EnodeId, evtNodeProposed.Ip[:], evtNodeProposed.Port, evtNodeProposed.Raftport), types.NodePendingApproval)

			case evtNodeDeactivated := <-chNodeDeactivated:
				p.updatePermissionedNodes(types.GetNodeUrl(evtNodeDeactivated.EnodeId, evtNodeDeactivated.Ip[:], evtNodeDeactivated.Port, evtNodeDeactivated.Raftport), NodeDelete)
				types.NodeInfoMap.UpsertNode(evtNodeDeactivated.OrgId, types.GetNodeUrl(evtNodeDeactivated.EnodeId, evtNodeDeactivated.Ip[:], evtNodeDeactivated.Port, evtNodeDeactivated.Raftport), types.NodeDeactivated)

			case evtNodeActivated := <-chNodeActivated:
				p.updatePermissionedNodes(types.GetNodeUrl(evtNodeActivated.EnodeId, evtNodeActivated.Ip[:], evtNodeActivated.Port, evtNodeActivated.Raftport), NodeAdd)
				types.NodeInfoMap.UpsertNode(evtNodeActivated.OrgId, types.GetNodeUrl(evtNodeActivated.EnodeId, evtNodeActivated.Ip[:], evtNodeActivated.Port, evtNodeActivated.Raftport), types.NodeApproved)

			case evtNodeBlacklisted := <-chNodeBlacklisted:
				types.NodeInfoMap.UpsertNode(evtNodeBlacklisted.OrgId, types.GetNodeUrl(evtNodeBlacklisted.EnodeId, evtNodeBlacklisted.Ip[:], evtNodeBlacklisted.Port, evtNodeBlacklisted.Raftport), types.NodeBlackListed)
				p.updateDisallowedNodes(types.GetNodeUrl(evtNodeBlacklisted.EnodeId, evtNodeBlacklisted.Ip[:], evtNodeBlacklisted.Port, evtNodeBlacklisted.Raftport), NodeAdd)
				p.updatePermissionedNodes(types.GetNodeUrl(evtNodeBlacklisted.EnodeId, evtNodeBlacklisted.Ip[:], evtNodeBlacklisted.Port, evtNodeBlacklisted.Raftport), NodeDelete)

			case evtNodeRecoveryInit := <-chNodeRecoveryInit:
				types.NodeInfoMap.UpsertNode(evtNodeRecoveryInit.OrgId, types.GetNodeUrl(evtNodeRecoveryInit.EnodeId, evtNodeRecoveryInit.Ip[:], evtNodeRecoveryInit.Port, evtNodeRecoveryInit.Raftport), types.NodeRecoveryInitiated)

			case evtNodeRecoveryDone := <-chNodeRecoveryDone:
				types.NodeInfoMap.UpsertNode(evtNodeRecoveryDone.OrgId, types.GetNodeUrl(evtNodeRecoveryDone.EnodeId, evtNodeRecoveryDone.Ip[:], evtNodeRecoveryDone.Port, evtNodeRecoveryDone.Raftport), types.NodeApproved)
				p.updateDisallowedNodes(types.GetNodeUrl(evtNodeRecoveryDone.EnodeId, evtNodeRecoveryDone.Ip[:], evtNodeRecoveryDone.Port, evtNodeRecoveryDone.Raftport), NodeDelete)
				p.updatePermissionedNodes(types.GetNodeUrl(evtNodeRecoveryDone.EnodeId, evtNodeRecoveryDone.Ip[:], evtNodeRecoveryDone.Port, evtNodeRecoveryDone.Raftport), NodeAdd)

			case <-stopChan:
				log.Info("quit Node contract watch")
				return
			}
		}
	}()
	return nil
}
