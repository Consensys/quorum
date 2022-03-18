package permission

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/permission/core"
	ptype "github.com/ethereum/go-ethereum/permission/core/types"
)

var isStringAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`).MatchString

//default gas limit to use if not passed in sendTxArgs
var defaultGasLimit = uint64(4712384)

//default gas price to use if not passed in sendTxArgs
var defaultGasPrice = big.NewInt(0)

// PermAction represents actions in permission contract
type PermAction int

const (
	AddOrg PermAction = iota
	ApproveOrg
	AddSubOrg
	UpdateOrgStatus
	ApproveOrgStatus
	AddNode
	UpdateNodeStatus
	AssignAdminRole
	ApproveAdminRole
	AddNewRole
	RemoveRole
	AddAccountToOrg
	ChangeAccountRole
	UpdateAccountStatus
	InitiateNodeRecovery
	InitiateAccountRecovery
	ApproveNodeRecovery
	ApproveAccountRecovery
)

type AccountUpdateAction int

const (
	SuspendAccount AccountUpdateAction = iota + 1
	ActivateSuspendedAccount
	BlacklistAccount
	RecoverBlacklistedAccount
	ApproveBlacklistedAccountRecovery
)

type NodeUpdateAction int

const (
	SuspendNode NodeUpdateAction = iota + 1
	ActivateSuspendedNode
	BlacklistNode
	RecoverBlacklistedNode
	ApproveBlacklistedNodeRecovery
)

type OrgUpdateAction int

const (
	SuspendOrg OrgUpdateAction = iota + 1
	ActivateSuspendedOrg
)

// QuorumControlsAPI provides an API to access Quorum's enterprise permissions related services
type QuorumControlsAPI struct {
	permCtrl *PermissionCtrl
}

type PendingOpInfo struct {
	PendingKey string `json:"pendingKey"`
	PendingOp  string `json:"pendingOp"`
}

var actionSuccess = "Action completed successfully"

// NewQuorumControlsAPI creates a new QuorumControlsAPI to access quorum services
func NewQuorumControlsAPI(p *PermissionCtrl) *QuorumControlsAPI {
	return &QuorumControlsAPI{p}
}

func (q *QuorumControlsAPI) OrgList() []core.OrgInfo {
	return core.OrgInfoMap.GetOrgList()
}

func (q *QuorumControlsAPI) NodeList() []core.NodeInfo {
	return core.NodeInfoMap.GetNodeList()
}

func (q *QuorumControlsAPI) RoleList() []core.RoleInfo {
	return core.RoleInfoMap.GetRoleList()
}

func (q *QuorumControlsAPI) AcctList() []core.AccountInfo {
	return core.AcctInfoMap.GetAcctList()
}

func (q *QuorumControlsAPI) GetOrgDetails(orgId string) (core.OrgDetailInfo, error) {
	o, err := core.OrgInfoMap.GetOrg(orgId)
	if err != nil {
		return core.OrgDetailInfo{}, err
	}

	if o == nil {
		return core.OrgDetailInfo{}, errors.New("org does not exist")
	}
	var acctList []core.AccountInfo
	var roleList []core.RoleInfo
	var nodeList []core.NodeInfo
	for _, a := range q.AcctList() {
		if a.OrgId == orgId {
			acctList = append(acctList, a)
		}
	}
	for _, a := range q.RoleList() {
		if a.OrgId == orgId {
			roleList = append(roleList, a)
		}
	}
	for _, a := range q.NodeList() {
		if a.OrgId == orgId {
			nodeList = append(nodeList, a)
		}
	}
	orgRec, err := core.OrgInfoMap.GetOrg(orgId)
	if err != nil {
		return core.OrgDetailInfo{}, err
	}

	if orgRec == nil {
		return core.OrgDetailInfo{NodeList: nodeList, RoleList: roleList, AcctList: acctList}, nil
	}
	return core.OrgDetailInfo{NodeList: nodeList, RoleList: roleList, AcctList: acctList, SubOrgList: orgRec.SubOrgList}, nil
}

func reportExecError(action PermAction, err error) (string, error) {
	log.Error("Failed to execute permission action", "action", action, "err", err)
	msg := fmt.Sprintf("failed to execute permissions action: %v", err)
	return "", errors.New(msg)
}

func (q *QuorumControlsAPI) AddOrg(orgId string, url string, acct common.Address, txa ethapi.SendTxArgs) (string, error) {
	orgService, err := q.permCtrl.NewPermissionOrgService(txa)

	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, Url: url, AcctId: acct, Txa: txa}

	if err := q.valAddOrg(args); err != nil {
		return "", err
	}
	tx, err := orgService.AddOrg(args)
	if err != nil {
		return reportExecError(AddOrg, err)
	}
	log.Debug("executed permission action", "action", AddOrg, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) AddSubOrg(porgId, orgId string, url string, txa ethapi.SendTxArgs) (string, error) {
	orgService, err := q.permCtrl.NewPermissionOrgService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{POrgId: porgId, OrgId: orgId, Url: url, Txa: txa}

	if err := q.valAddSubOrg(args); err != nil {
		return "", err
	}
	tx, err := orgService.AddSubOrg(args)
	if err != nil {
		return reportExecError(AddSubOrg, err)
	}
	log.Debug("executed permission action", "action", AddSubOrg, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) ApproveOrg(orgId string, url string, acct common.Address, txa ethapi.SendTxArgs) (string, error) {
	orgService, err := q.permCtrl.NewPermissionOrgService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, Url: url, AcctId: acct, Txa: txa}
	if err := q.valApproveOrg(args); err != nil {
		return "", err
	}
	tx, err := orgService.ApproveOrg(args)
	if err != nil {
		return reportExecError(ApproveOrg, err)
	}
	log.Debug("executed permission action", "action", ApproveOrg, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) UpdateOrgStatus(orgId string, status uint8, txa ethapi.SendTxArgs) (string, error) {
	orgService, err := q.permCtrl.NewPermissionOrgService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, Action: status, Txa: txa}
	if err := q.valUpdateOrgStatus(args); err != nil {
		return "", err
	}
	// and in suspended state for suspension revoke
	tx, err := orgService.UpdateOrgStatus(args)
	if err != nil {
		return reportExecError(UpdateOrgStatus, err)
	}
	log.Debug("executed permission action", "action", UpdateOrgStatus, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) AddNode(orgId string, url string, txa ethapi.SendTxArgs) (string, error) {
	nodeService, err := q.permCtrl.NewPermissionNodeService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, Url: url, Txa: txa}
	if err := q.valAddNode(args); err != nil {
		return "", err
	}
	// check if node is already there
	tx, err := nodeService.AddNode(args)
	if err != nil {
		return reportExecError(AddNode, err)
	}
	log.Debug("executed permission action", "action", AddNode, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) UpdateNodeStatus(orgId string, url string, action uint8, txa ethapi.SendTxArgs) (string, error) {
	nodeService, err := q.permCtrl.NewPermissionNodeService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, Url: url, Action: action, Txa: txa}
	if err := q.valUpdateNodeStatus(args, UpdateNodeStatus); err != nil {
		return "", err
	}
	// check node status for operation
	tx, err := nodeService.UpdateNodeStatus(args)
	if err != nil {
		return reportExecError(UpdateNodeStatus, err)
	}
	log.Debug("executed permission action", "action", UpdateNodeStatus, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) ApproveOrgStatus(orgId string, status uint8, txa ethapi.SendTxArgs) (string, error) {
	orgService, err := q.permCtrl.NewPermissionOrgService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, Action: status, Txa: txa}
	if err := q.valApproveOrgStatus(args); err != nil {
		return "", err
	}
	// validate that status change is pending approval
	tx, err := orgService.ApproveOrgStatus(args)
	if err != nil {
		return reportExecError(ApproveOrgStatus, err)
	}
	log.Debug("executed permission action", "action", ApproveOrgStatus, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) AssignAdminRole(orgId string, acct common.Address, roleId string, txa ethapi.SendTxArgs) (string, error) {
	accountService, err := q.permCtrl.NewPermissionAccountService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, AcctId: acct, RoleId: roleId, Txa: txa}
	if err := q.valAssignAdminRole(args); err != nil {
		return "", err
	}
	// check if account is already in use in another org
	tx, err := accountService.AssignAdminRole(args)
	if err != nil {
		return reportExecError(AssignAdminRole, err)
	}
	log.Debug("executed permission action", "action", AssignAdminRole, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) ApproveAdminRole(orgId string, acct common.Address, txa ethapi.SendTxArgs) (string, error) {
	accountService, err := q.permCtrl.NewPermissionAccountService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, AcctId: acct, Txa: txa}
	if err := q.valApproveAdminRole(args); err != nil {
		return "", err
	}
	// check if anything is pending approval
	tx, err := accountService.ApproveAdminRole(args)
	if err != nil {
		return reportExecError(ApproveAdminRole, err)
	}
	log.Debug("executed permission action", "action", ApproveAdminRole, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) AddNewRole(orgId string, roleId string, access uint8, isVoter bool, isAdmin bool, txa ethapi.SendTxArgs) (string, error) {
	roleService, err := q.permCtrl.NewPermissionRoleService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, RoleId: roleId, AccessType: access, IsVoter: isVoter, IsAdmin: isAdmin, Txa: txa}
	if err := q.valAddNewRole(args); err != nil {
		return "", err
	}
	// check if role is already there in the org
	tx, err := roleService.AddNewRole(args)
	if err != nil {
		return reportExecError(AddNewRole, err)
	}
	log.Debug("executed permission action", "action", AddNewRole, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) RemoveRole(orgId string, roleId string, txa ethapi.SendTxArgs) (string, error) {
	roleService, err := q.permCtrl.NewPermissionRoleService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, RoleId: roleId, Txa: txa}

	if err := q.valRemoveRole(args); err != nil {
		return "", err
	}
	tx, err := roleService.RemoveRole(args)
	if err != nil {
		return reportExecError(RemoveRole, err)
	}
	log.Debug("executed permission action", "action", RemoveRole, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) AddAccountToOrg(acct common.Address, orgId string, roleId string, txa ethapi.SendTxArgs) (string, error) {
	accountService, err := q.permCtrl.NewPermissionAccountService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, RoleId: roleId, AcctId: acct, Txa: txa}

	if err := q.valAssignRole(args); err != nil {
		return "", err
	}
	tx, err := accountService.AssignAccountRole(args)
	if err != nil {
		return reportExecError(AddAccountToOrg, err)
	}
	log.Debug("executed permission action", "action", AddAccountToOrg, "tx", tx)
	return actionSuccess, nil
}
func (q *QuorumControlsAPI) ChangeAccountRole(acct common.Address, orgId string, roleId string, txa ethapi.SendTxArgs) (string, error) {
	accountService, err := q.permCtrl.NewPermissionAccountService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, RoleId: roleId, AcctId: acct, Txa: txa}

	if err := q.valAssignRole(args); err != nil {
		return "", err
	}
	tx, err := accountService.AssignAccountRole(args)
	if err != nil {
		return reportExecError(ChangeAccountRole, err)
	}
	log.Debug("executed permission action", "action", ChangeAccountRole, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) UpdateAccountStatus(orgId string, acct common.Address, status uint8, txa ethapi.SendTxArgs) (string, error) {
	accountService, err := q.permCtrl.NewPermissionAccountService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, AcctId: acct, Action: status, Txa: txa}

	if err := q.valUpdateAccountStatus(args, UpdateAccountStatus); err != nil {
		return "", err
	}
	tx, err := accountService.UpdateAccountStatus(args)
	if err != nil {
		return reportExecError(UpdateAccountStatus, err)
	}
	log.Debug("executed permission action", "action", UpdateAccountStatus, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) RecoverBlackListedNode(orgId string, enodeId string, txa ethapi.SendTxArgs) (string, error) {
	nodeService, err := q.permCtrl.NewPermissionNodeService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, Url: enodeId, Txa: txa}

	if err := q.valRecoverNode(args, InitiateNodeRecovery); err != nil {
		return "", err
	}
	tx, err := nodeService.StartBlacklistedNodeRecovery(args)
	if err != nil {
		return reportExecError(InitiateNodeRecovery, err)
	}
	log.Debug("executed permission action", "action", InitiateNodeRecovery, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) ApproveBlackListedNodeRecovery(orgId string, enodeId string, txa ethapi.SendTxArgs) (string, error) {
	nodeService, err := q.permCtrl.NewPermissionNodeService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, Url: enodeId, Txa: txa}

	if err := q.valRecoverNode(args, ApproveNodeRecovery); err != nil {
		return "", err
	}
	tx, err := nodeService.ApproveBlacklistedNodeRecovery(args)
	if err != nil {
		return reportExecError(ApproveNodeRecovery, err)
	}
	log.Debug("executed permission action", "action", ApproveNodeRecovery, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) RecoverBlackListedAccount(orgId string, acctId common.Address, txa ethapi.SendTxArgs) (string, error) {
	accountService, err := q.permCtrl.NewPermissionAccountService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, AcctId: acctId, Txa: txa}

	if err := q.valRecoverAccount(args, InitiateAccountRecovery); err != nil {
		return "", err
	}
	tx, err := accountService.StartBlacklistedAccountRecovery(args)
	if err != nil {
		return reportExecError(InitiateAccountRecovery, err)
	}
	log.Debug("executed permission action", "action", InitiateAccountRecovery, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) ApproveBlackListedAccountRecovery(orgId string, acctId common.Address, txa ethapi.SendTxArgs) (string, error) {
	accountService, err := q.permCtrl.NewPermissionAccountService(txa)
	if err != nil {
		return "", err
	}
	args := ptype.TxArgs{OrgId: orgId, AcctId: acctId, Txa: txa}

	if err := q.valRecoverAccount(args, ApproveAccountRecovery); err != nil {
		return "", err
	}
	tx, err := accountService.ApproveBlacklistedAccountRecovery(args)
	if err != nil {
		return reportExecError(ApproveAccountRecovery, err)
	}
	log.Debug("executed permission action", "action", ApproveAccountRecovery, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) TransactionAllowed(txa ethapi.SendTxArgs) bool {
	var value, gasPrice, gasLimit *big.Int
	var payload []byte
	var to, from common.Address
	if txa.Value != nil {
		value = txa.Value.ToInt()
	} else {
		value = big.NewInt(0)
	}
	from = txa.From
	if txa.To == nil {
		to = common.Address{}
	} else {
		to = *txa.To
	}

	if txa.GasPrice != nil {
		gasPrice = txa.GasPrice.ToInt()
	} else {
		gasPrice = big.NewInt(0)
	}

	if txa.Gas != nil {
		gasLimit = big.NewInt(int64(*txa.Gas))
	} else {
		gasLimit = big.NewInt(0)
	}

	if txa.Data != nil {
		payload = *txa.Data
	}

	transactionType := core.ValueTransferTxn

	if txa.To == nil {
		transactionType = core.ContractDeployTxn
	} else if txa.Data != nil {
		transactionType = core.ContractCallTxn
	}

	if err := core.IsTransactionAllowed(from, to, value, gasPrice, gasLimit, payload, transactionType); err != nil {
		return false
	} else {
		return true
	}
}

func (q *QuorumControlsAPI) ConnectionAllowed(enodeId, ip string, port, raftPort uint16) bool {
	controlService, err := q.permCtrl.NewPermissionControlService()
	if err != nil {
		return false
	}

	if allowed, err := controlService.ConnectionAllowed(enodeId, ip, port, raftPort); err != nil {
		return false
	} else {
		return allowed
	}
}

// check if the account is network admin
func (q *QuorumControlsAPI) isNetworkAdmin(account common.Address) bool {
	ac, _ := core.AcctInfoMap.GetAccount(account)
	return ac != nil && ac.RoleId == q.permCtrl.permConfig.NwAdminRole
}

func (q *QuorumControlsAPI) isOrgAdmin(account common.Address, orgId string) error {
	org, err := core.OrgInfoMap.GetOrg(orgId)
	if err != nil {
		return err
	}
	if org == nil {
		return ptype.ErrOrgDoesNotExists
	}
	ac, _ := core.AcctInfoMap.GetAccount(account)
	if ac == nil {
		return ptype.ErrNotOrgAdmin
	}
	// check if the account is network admin
	if !(ac.IsOrgAdmin && (ac.OrgId == orgId || ac.OrgId == org.UltimateParent)) {
		return ptype.ErrNotOrgAdmin
	}
	return nil
}

func (q *QuorumControlsAPI) validateOrg(orgId, pOrgId string) error {
	// validate Parent org id
	if pOrgId != "" {
		if _, err := core.OrgInfoMap.GetOrg(pOrgId); err != nil {
			return ptype.ErrInvalidParentOrg
		}
		locOrgId := pOrgId + "." + orgId
		if lorgRec, _ := core.OrgInfoMap.GetOrg(locOrgId); lorgRec != nil {
			return ptype.ErrOrgExists
		}
	} else if orgRec, _ := core.OrgInfoMap.GetOrg(orgId); orgRec != nil {
		return ptype.ErrOrgExists
	}
	return nil
}

func (q *QuorumControlsAPI) validatePendingOp(authOrg, orgId, url string, account common.Address, pendingOp int64) bool {
	auditService, err := q.permCtrl.NewPermissionAuditService()
	if err != nil {
		return false
	}
	return auditService.ValidatePendingOp(authOrg, orgId, url, account, pendingOp)
}

func (q *QuorumControlsAPI) checkPendingOp(orgId string) bool {
	auditService, err := q.permCtrl.NewPermissionAuditService()
	if err != nil {
		return false
	}
	return auditService.CheckPendingOp(orgId)
}

func (q *QuorumControlsAPI) checkOrgStatus(orgId string, op uint8) error {
	org, _ := core.OrgInfoMap.GetOrg(orgId)

	if org == nil {
		return ptype.ErrOrgDoesNotExists
	}
	// check if its a master org. operation is allowed only if its a master org
	if org.Level.Cmp(big.NewInt(1)) != 0 {
		return ptype.ErrNotMasterOrg
	}

	if !((op == 1 && org.Status == core.OrgApproved) || (op == 2 && org.Status == core.OrgSuspended)) {
		return ptype.ErrOpNotAllowed
	}
	return nil
}

func (q *QuorumControlsAPI) valNodeStatusChange(orgId, url string, op NodeUpdateAction, permAction PermAction) error {
	// validates if the enode is linked the passed organization
	// validate node id and if the status change can happen
	if len(url) == 0 {
		return ptype.ErrInvalidNode
	}
	if err := q.valNodeDetails(url); err != nil && err.Error() != ptype.ErrNodePresent.Error() {
		return err
	}

	node, err := core.NodeInfoMap.GetNodeByUrl(url)
	if err != nil {
		return err
	}

	if node.OrgId != orgId {
		return ptype.ErrNodeOrgMismatch
	}

	if node.Status == core.NodeBlackListed && op != RecoverBlacklistedNode {
		return ptype.ErrBlacklistedNode
	}

	// validate the op and node status and check if the op can be performed
	if (permAction == UpdateNodeStatus && (op != SuspendNode && op != ActivateSuspendedNode && op != BlacklistNode)) ||
		(permAction == InitiateNodeRecovery && op != RecoverBlacklistedNode) ||
		(permAction == ApproveNodeRecovery && op != ApproveBlacklistedNodeRecovery) {
		return ptype.ErrOpNotAllowed
	}

	if (op == SuspendNode && node.Status != core.NodeApproved) ||
		(op == ActivateSuspendedNode && node.Status != core.NodeDeactivated) ||
		(op == BlacklistNode && node.Status == core.NodeRecoveryInitiated) ||
		(op == RecoverBlacklistedNode && node.Status != core.NodeBlackListed) ||
		(op == ApproveBlacklistedNodeRecovery && node.Status != core.NodeRecoveryInitiated) {
		return ptype.ErrOpNotAllowed
	}

	return nil
}

func (q *QuorumControlsAPI) validateRole(orgId, roleId string) bool {
	var r *core.RoleInfo
	r, err := core.RoleInfoMap.GetRole(orgId, roleId)
	if err != nil {
		return false
	}

	orgRec, err := core.OrgInfoMap.GetOrg(orgId)
	if err != nil {
		return false
	}
	r, err = core.RoleInfoMap.GetRole(orgRec.UltimateParent, roleId)
	if err != nil {
		return false
	}

	return r != nil && r.Active
}

func (q *QuorumControlsAPI) valAccountStatusChange(orgId string, account common.Address, permAction PermAction, op AccountUpdateAction) error {
	// validates if the enode is linked the passed organization
	ac, err := core.AcctInfoMap.GetAccount(account)
	if err != nil {
		return err
	}

	if ac.IsOrgAdmin && (ac.RoleId == q.permCtrl.permConfig.NwAdminRole || ac.RoleId == q.permCtrl.permConfig.OrgAdminRole) && (op == 1 || op == 3) {
		return ptype.ErrOpNotAllowed
	}

	if ac.OrgId != orgId {
		return ptype.ErrOrgNotOwner
	}
	if (permAction == UpdateAccountStatus && (op != SuspendAccount && op != ActivateSuspendedAccount && op != BlacklistAccount)) ||
		(permAction == InitiateAccountRecovery && op != RecoverBlacklistedAccount) ||
		(permAction == ApproveAccountRecovery && op != ApproveBlacklistedAccountRecovery) {
		return ptype.ErrOpNotAllowed
	}

	if ac.Status == core.AcctBlacklisted && op != RecoverBlacklistedAccount {
		return ptype.ErrBlacklistedAccount
	}

	if (op == SuspendAccount && ac.Status != core.AcctActive) ||
		(op == ActivateSuspendedAccount && ac.Status != core.AcctSuspended) ||
		(op == BlacklistAccount && ac.Status == core.AcctRecoveryInitiated) ||
		(op == RecoverBlacklistedAccount && ac.Status != core.AcctBlacklisted) ||
		(op == ApproveBlacklistedAccountRecovery && ac.Status != core.AcctRecoveryInitiated) {
		return ptype.ErrOpNotAllowed
	}
	return nil
}

func (q *QuorumControlsAPI) checkOrgAdminExists(orgId, roleId string, account common.Address) error {
	if ac, _ := core.AcctInfoMap.GetAccount(account); ac != nil {
		if ac.OrgId != orgId {
			return ptype.ErrAccountInUse
		}
		if roleId != "" && roleId == q.permCtrl.permConfig.OrgAdminRole && ac.IsOrgAdmin {
			return ptype.ErrAccountOrgAdmin
		}
	}
	return nil
}

func (q *QuorumControlsAPI) valSubOrgBreadthDepth(porgId string) error {
	org, err := core.OrgInfoMap.GetOrg(porgId)
	if err != nil {
		return ptype.ErrOpNotAllowed
	}

	if q.permCtrl.permConfig.SubOrgDepth.Cmp(org.Level) == 0 {
		return ptype.ErrMaxDepth
	}

	if q.permCtrl.permConfig.SubOrgBreadth.Cmp(big.NewInt(int64(len(org.SubOrgList)))) == 0 {
		return ptype.ErrMaxBreadth
	}

	return nil
}

func (q *QuorumControlsAPI) checkNodeExists(url, enodeId string) bool {
	node, _ := core.NodeInfoMap.GetNodeByUrl(url)
	if node != nil {
		return true
	}
	// check if the same nodeid is in use with different port numbers
	nodeList := core.NodeInfoMap.GetNodeList()
	for _, n := range nodeList {
		if enodeDet, er := enode.ParseV4(n.Url); er == nil {
			if enodeDet.EnodeID() == enodeId {
				return true
			}
		}
	}
	return false
}

func (q *QuorumControlsAPI) valNodeDetails(url string) error {
	// validate node id and
	if len(url) != 0 {
		enodeDet, err := enode.ParseV4(url)
		if err != nil {
			return ptype.ErrInvalidNode
		}
		if q.permCtrl.isRaft && !q.permCtrl.useDns && enodeDet.Host() != "" {
			return ptype.ErrHostNameNotSupported
		}
		// check if node already there
		if q.checkNodeExists(url, enodeDet.EnodeID()) {
			return ptype.ErrNodePresent
		}
	}
	return nil
}

// all validations for add org operation
func (q *QuorumControlsAPI) valAddOrg(args ptype.TxArgs) error {
	// check if the org id contains "."
	if args.OrgId == "" || args.Url == "" || args.AcctId == (common.Address{0}) {
		return ptype.ErrInvalidInput
	}
	if !isStringAlphaNumeric(args.OrgId) {
		return ptype.ErrInvalidOrgName
	}

	// check if caller is network admin
	if !q.isNetworkAdmin(args.Txa.From) {
		return ptype.ErrNotNetworkAdmin
	}

	// check if any previous op is pending approval for network admin
	if q.checkPendingOp(q.permCtrl.permConfig.NwAdminOrg) {
		return ptype.ErrPendingApprovals
	}
	// check if org already exists
	if er := q.validateOrg(args.OrgId, ""); er != nil {
		return er
	}

	// validate node id and
	if er := q.valNodeDetails(args.Url); er != nil {
		return er
	}

	// check if account is already part of another org
	if er := q.checkOrgAdminExists(args.OrgId, "", args.AcctId); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valApproveOrg(args ptype.TxArgs) error {
	// check caller is network admin
	if !q.isNetworkAdmin(args.Txa.From) {
		return ptype.ErrNotNetworkAdmin
	}
	// check if anything pending approval
	if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.OrgId, args.Url, args.AcctId, 1) {
		return ptype.ErrNothingToApprove
	}
	return nil
}

func (q *QuorumControlsAPI) valAddSubOrg(args ptype.TxArgs) error {
	// check if the org id contains "."
	if args.OrgId == "" {
		return ptype.ErrInvalidInput
	}
	if !isStringAlphaNumeric(args.OrgId) {
		return ptype.ErrInvalidOrgName
	}

	// check if caller is network admin
	if er := q.isOrgAdmin(args.Txa.From, args.POrgId); er != nil {
		return er
	}

	// check if org already exists
	if er := q.validateOrg(args.OrgId, args.POrgId); er != nil {
		return er
	}

	if er := q.valSubOrgBreadthDepth(args.POrgId); er != nil {
		return er
	}

	if er := q.valNodeDetails(args.Url); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valUpdateOrgStatus(args ptype.TxArgs) error {
	// check if called is network admin
	if !q.isNetworkAdmin(args.Txa.From) {
		return ptype.ErrNotNetworkAdmin
	}
	if OrgUpdateAction(args.Action) != SuspendOrg &&
		OrgUpdateAction(args.Action) != ActivateSuspendedOrg {
		return ptype.ErrOpNotAllowed
	}

	//check if passed org id is network admin org. update should not be allowed
	if args.OrgId == q.permCtrl.permConfig.NwAdminOrg {
		return ptype.ErrOpNotAllowed
	}
	// check if status update can be performed. Org should be approved for suspension
	if er := q.checkOrgStatus(args.OrgId, args.Action); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valApproveOrgStatus(args ptype.TxArgs) error {
	// check if called is network admin
	if !q.isNetworkAdmin(args.Txa.From) {
		return ptype.ErrNotNetworkAdmin
	}
	// check if anything is pending approval
	var pendingOp int64
	if args.Action == 1 {
		pendingOp = 2
	} else if args.Action == 2 {
		pendingOp = 3
	} else {
		return ptype.ErrOpNotAllowed
	}
	if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.OrgId, "", common.Address{}, pendingOp) {
		return ptype.ErrNothingToApprove
	}
	return nil
}

func (q *QuorumControlsAPI) valAddNode(args ptype.TxArgs) error {
	if args.Url == "" {
		return ptype.ErrInvalidInput
	}
	// check if caller is network admin
	if er := q.isOrgAdmin(args.Txa.From, args.OrgId); er != nil {
		return er
	}

	if er := q.valNodeDetails(args.Url); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valUpdateNodeStatus(args ptype.TxArgs, permAction PermAction) error {
	// check if org admin
	// check if caller is network admin
	if er := q.isOrgAdmin(args.Txa.From, args.OrgId); er != nil {
		return er
	}

	// validation status change is with in allowed set
	if er := q.valNodeStatusChange(args.OrgId, args.Url, NodeUpdateAction(args.Action), permAction); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valAssignAdminRole(args ptype.TxArgs) error {
	if args.AcctId == (common.Address{0}) {
		return ptype.ErrInvalidInput
	}
	// check if caller is network admin
	if args.RoleId != q.permCtrl.permConfig.OrgAdminRole && args.RoleId != q.permCtrl.permConfig.NwAdminRole {
		return ptype.ErrOpNotAllowed
	}

	if !q.isNetworkAdmin(args.Txa.From) {
		return ptype.ErrNotNetworkAdmin
	}

	if err := q.validateOrg(args.OrgId, ""); err == nil {
		return ptype.ErrOrgDoesNotExists
	}

	// check if account is already part of another org
	if er := q.checkOrgAdminExists(args.OrgId, args.RoleId, args.AcctId); er != nil && er.Error() != ptype.ErrOrgAdminExists.Error() {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valApproveAdminRole(args ptype.TxArgs) error {
	// check if caller is network admin
	if !q.isNetworkAdmin(args.Txa.From) {
		return ptype.ErrNotNetworkAdmin
	}
	// check if the org exists

	// check if account is valid
	ac, _ := core.AcctInfoMap.GetAccount(args.AcctId)
	if ac == nil {
		return ptype.ErrInvalidAccount
	}
	// validate pending op
	if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, ac.OrgId, "", args.AcctId, 4) {
		return ptype.ErrNothingToApprove
	}
	return nil
}

func (q *QuorumControlsAPI) valAddNewRole(args ptype.TxArgs) error {
	if args.RoleId == "" {
		return ptype.ErrInvalidInput
	}
	// check if caller is network admin
	if er := q.isOrgAdmin(args.Txa.From, args.OrgId); er != nil {
		return er
	}
	// validate if role is already present
	if roleRec, _ := core.RoleInfoMap.GetRole(args.OrgId, args.RoleId); roleRec != nil {
		return ptype.ErrRoleExists
	}
	return nil
}

func (q *QuorumControlsAPI) valRemoveRole(args ptype.TxArgs) error {
	// check if caller is network admin
	if er := q.isOrgAdmin(args.Txa.From, args.OrgId); er != nil {
		return er
	}

	// admin roles cannot be removed
	if args.RoleId == q.permCtrl.permConfig.OrgAdminRole || args.RoleId == q.permCtrl.permConfig.NwAdminRole {
		return ptype.ErrAdminRoles
	}

	// check if role is alraedy inactive
	r, _ := core.RoleInfoMap.GetRole(args.OrgId, args.RoleId)
	if r == nil {
		return ptype.ErrInvalidRole
	} else if !r.Active {
		return ptype.ErrInactiveRole
	}

	// check if the role has active accounts. if yes operations should not be allowed
	if len(core.AcctInfoMap.GetAcctListRole(args.OrgId, args.RoleId)) != 0 {
		return ptype.ErrRoleActive
	}
	return nil
}

func (q *QuorumControlsAPI) valAssignRole(args ptype.TxArgs) error {
	if args.AcctId == (common.Address{0}) {
		return ptype.ErrInvalidInput
	}
	if args.RoleId == q.permCtrl.permConfig.OrgAdminRole || args.RoleId == q.permCtrl.permConfig.NwAdminRole {
		return ptype.ErrInvalidRole
	}
	// check if caller is network admin
	if er := q.isOrgAdmin(args.Txa.From, args.OrgId); er != nil {
		return er
	}

	// check if the role is valid
	if !q.validateRole(args.OrgId, args.RoleId) {
		return ptype.ErrInvalidRole
	}

	// check if the account is part of another org
	if ac, _ := core.AcctInfoMap.GetAccount(args.AcctId); ac != nil {
		if ac != nil && ac.OrgId != args.OrgId {
			return ptype.ErrAccountInUse
		}
	}
	return nil
}

func (q *QuorumControlsAPI) valUpdateAccountStatus(args ptype.TxArgs, permAction PermAction) error {
	// check if the caller is org admin
	if er := q.isOrgAdmin(args.Txa.From, args.OrgId); er != nil {
		return er
	}
	// validation status change is with in allowed set
	if er := q.valAccountStatusChange(args.OrgId, args.AcctId, permAction, AccountUpdateAction(args.Action)); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valRecoverNode(args ptype.TxArgs, action PermAction) error {
	// check if the caller is org admin
	if !q.isNetworkAdmin(args.Txa.From) {
		return ptype.ErrNotNetworkAdmin
	}
	// validate inputs - org id is valid, node is valid and in blacklisted state
	if err := q.validateOrg(args.OrgId, ""); err.Error() != ptype.ErrOrgExists.Error() {
		return ptype.ErrInvalidOrgName
	}

	if action == InitiateNodeRecovery {
		if err := q.valNodeStatusChange(args.OrgId, args.Url, 4, InitiateAccountRecovery); err != nil {
			return err
		}
		// check no pending approval items
		if q.checkPendingOp(q.permCtrl.permConfig.NwAdminOrg) {
			return ptype.ErrPendingApprovals
		}
	} else {
		// validate inputs - org id is valid, Node is valid pending recovery state
		if err := q.valNodeStatusChange(args.OrgId, args.Url, 5, ApproveNodeRecovery); err != nil {
			return err
		}
		if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.OrgId, args.Url, common.Address{}, 5) {
			return ptype.ErrNothingToApprove
		}
	}

	// if it is approval ensure that

	return nil
}

func (q *QuorumControlsAPI) valRecoverAccount(args ptype.TxArgs, action PermAction) error {
	// check if the caller is org admin
	if !q.isNetworkAdmin(args.Txa.From) {
		return ptype.ErrNotNetworkAdmin
	}

	var opAction AccountUpdateAction
	if action == InitiateAccountRecovery {
		opAction = RecoverBlacklistedAccount
	} else {
		opAction = ApproveBlacklistedAccountRecovery
	}

	if err := q.valAccountStatusChange(args.OrgId, args.AcctId, action, opAction); err != nil {
		return err
	}

	if action == InitiateAccountRecovery && q.checkPendingOp(q.permCtrl.permConfig.NwAdminOrg) {
		return ptype.ErrPendingApprovals
	}

	if action == ApproveAccountRecovery && !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.OrgId, "", args.AcctId, 6) {
		return ptype.ErrNothingToApprove
	}
	return nil
}
