package permission

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
	pbind "github.com/ethereum/go-ethereum/permission/bind"
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

// QuorumControlsAPI provides an API to access Quorum's node permission and org key management related services
type QuorumControlsAPI struct {
	permCtrl *PermissionCtrl
}

// txArgs holds arguments required for execute functions
type txArgs struct {
	orgId      string
	porgId     string
	url        string
	roleId     string
	isVoter    bool
	isAdmin    bool
	acctId     common.Address
	accessType uint8
	action     uint8
	voter      common.Address
	morgId     string
	tmKey      string
	txa        ethapi.SendTxArgs
}

type PendingOpInfo struct {
	PendingKey string `json:"pendingKey"`
	PendingOp  string `json:"pendingOp"`
}

type ExecStatus struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

func (e ExecStatus) OpStatus() (string, error) {
	if e.Status {
		return e.Msg, nil
	}
	return "", fmt.Errorf("%s", e.Msg)
}

var (
	ErrNotNetworkAdmin    = ExecStatus{false, "Operation can be performed by network admin only. Account not a network admin."}
	ErrNotOrgAdmin        = ExecStatus{false, "Operation can be performed by org admin only. Account not a org admin."}
	ErrNodePresent        = ExecStatus{false, "EnodeId already part of network."}
	ErrInvalidNode        = ExecStatus{false, "Invalid enode id"}
	ErrInvalidAccount     = ExecStatus{false, "Invalid account id"}
	ErrOrgExists          = ExecStatus{false, "Org already exists"}
	ErrPendingApprovals   = ExecStatus{false, "Pending approvals for the organization. Approve first"}
	ErrNothingToApprove   = ExecStatus{false, "Nothing to approve"}
	ErrOpNotAllowed       = ExecStatus{false, "Operation not allowed"}
	ErrNodeOrgMismatch    = ExecStatus{false, "Enode id passed does not belong to the organization."}
	ErrBlacklistedNode    = ExecStatus{false, "Blacklisted node. Operation not allowed"}
	ErrBlacklistedAccount = ExecStatus{false, "Blacklisted account. Operation not allowed"}
	ErrAccountOrgAdmin    = ExecStatus{false, "Account already org admin for the org"}
	ErrOrgAdminExists     = ExecStatus{false, "Org admin exists for the org"}
	ErrAccountInUse       = ExecStatus{false, "Account already in use in another organization"}
	ErrRoleExists         = ExecStatus{false, "Role exists for the org"}
	ErrRoleActive         = ExecStatus{false, "Accounts linked to the role. Cannot be removed"}
	ErrAdminRoles         = ExecStatus{false, "Admin role cannot be removed"}
	ErrInvalidOrgName     = ExecStatus{false, "Org id cannot contain special characters"}
	ErrInvalidParentOrg   = ExecStatus{false, "Invalid parent org id"}
	ErrAccountNotThere    = ExecStatus{false, "Account does not exists"}
	ErrOrgNotOwner        = ExecStatus{false, "Account does not belong to this org"}
	ErrMaxDepth           = ExecStatus{false, "Max depth for sub orgs reached"}
	ErrMaxBreadth         = ExecStatus{false, "Max breadth for sub orgs reached"}
	ErrNodeDoesNotExists  = ExecStatus{false, "Node does not exists"}
	ErrOrgDoesNotExists   = ExecStatus{false, "Org does not exists"}
	ErrInactiveRole       = ExecStatus{false, "Role is already inactive"}
	ErrInvalidRole        = ExecStatus{false, "Invalid role"}
	ErrInvalidInput       = ExecStatus{false, "Invalid input"}
	ErrNotMasterOrg       = ExecStatus{false, "Org is not a master org"}

	ExecSuccess = ExecStatus{true, "Action completed successfully"}
)

// NewQuorumControlsAPI creates a new QuorumControlsAPI to access quorum services
func NewQuorumControlsAPI(p *PermissionCtrl) *QuorumControlsAPI {
	return &QuorumControlsAPI{p}
}

func (q *QuorumControlsAPI) OrgList() []types.OrgInfo {
	return types.OrgInfoMap.GetOrgList()
}

func (q *QuorumControlsAPI) NodeList() []types.NodeInfo {
	return types.NodeInfoMap.GetNodeList()
}

func (q *QuorumControlsAPI) RoleList() []types.RoleInfo {
	return types.RoleInfoMap.GetRoleList()
}

func (q *QuorumControlsAPI) AcctList() []types.AccountInfo {
	return types.AcctInfoMap.GetAcctList()
}

func (q *QuorumControlsAPI) GetOrgDetails(orgId string) (types.OrgDetailInfo, error) {
	if o := types.OrgInfoMap.GetOrg(orgId); o == nil {
		return types.OrgDetailInfo{}, errors.New("org does not exist")
	}
	var acctList []types.AccountInfo
	var roleList []types.RoleInfo
	var nodeList []types.NodeInfo
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
	return types.OrgDetailInfo{NodeList: nodeList, RoleList: roleList, AcctList: acctList, SubOrgList: types.OrgInfoMap.GetOrg(orgId).SubOrgList}, nil
}

func (q *QuorumControlsAPI) initOp(txa ethapi.SendTxArgs) (*pbind.PermInterfaceSession, ExecStatus) {
	var err error
	var w accounts.Wallet

	w, err = q.validateAccount(txa.From)
	if err != nil {
		return nil, ErrInvalidAccount
	}
	pinterf := q.newPermInterfaceSession(w, txa)

	return pinterf, ExecSuccess
}

func reportExecError(action PermAction, err error) (string, error) {
	log.Error("Failed to execute permission action", "action", action, "err", err)
	msg := fmt.Sprintf("failed to execute permissions action: %v", err)
	return ExecStatus{false, msg}.OpStatus()
}

func (q *QuorumControlsAPI) AddOrg(orgId string, url string, acct common.Address, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, url: url, acctId: acct, txa: txa}

	if execStatus := q.valAddOrg(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.AddOrg(args.orgId, enodeId, ip, port, raftPort, args.acctId)
	if err != nil {
		return reportExecError(AddOrg, err)
	}
	log.Debug("executed permission action", "action", AddOrg, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) AddSubOrg(porgId, orgId string, url string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{porgId: porgId, orgId: orgId, url: url, txa: txa}

	if execStatus := q.valAddSubOrg(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.AddSubOrg(args.porgId, args.orgId, enodeId, ip, port, raftPort)
	if err != nil {
		return reportExecError(AddSubOrg, err)
	}
	log.Debug("executed permission action", "action", AddSubOrg, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) ApproveOrg(orgId string, url string, acct common.Address, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, url: url, acctId: acct, txa: txa}
	if execStatus := q.valApproveOrg(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.ApproveOrg(args.orgId, enodeId, ip, port, raftPort, args.acctId)
	if err != nil {
		return reportExecError(ApproveOrg, err)
	}
	log.Debug("executed permission action", "action", ApproveOrg, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) UpdateOrgStatus(orgId string, status uint8, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, action: status, txa: txa}
	if execStatus := q.valUpdateOrgStatus(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	// and in suspended state for suspension revoke
	tx, err := pinterf.UpdateOrgStatus(args.orgId, big.NewInt(int64(args.action)))
	if err != nil {
		return reportExecError(UpdateOrgStatus, err)
	}
	log.Debug("executed permission action", "action", UpdateOrgStatus, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) AddNode(orgId string, url string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, url: url, txa: txa}
	if execStatus := q.valAddNode(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	// check if node is already there
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.AddNode(args.orgId, enodeId, ip, port, raftPort)
	if err != nil {
		return reportExecError(AddNode, err)
	}
	log.Debug("executed permission action", "action", AddNode, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) UpdateNodeStatus(orgId string, url string, action uint8, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, url: url, action: action, txa: txa}
	if execStatus := q.valUpdateNodeStatus(args, UpdateNodeStatus, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	// check node status for operation
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.UpdateNodeStatus(args.orgId, enodeId, ip, port, raftPort, big.NewInt(int64(args.action)))
	if err != nil {
		return reportExecError(UpdateNodeStatus, err)
	}
	log.Debug("executed permission action", "action", UpdateNodeStatus, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) ApproveOrgStatus(orgId string, status uint8, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, action: status, txa: txa}
	if execStatus := q.valApproveOrgStatus(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	// validate that status change is pending approval
	tx, err := pinterf.ApproveOrgStatus(args.orgId, big.NewInt(int64(args.action)))
	if err != nil {
		return reportExecError(ApproveOrgStatus, err)
	}
	log.Debug("executed permission action", "action", ApproveOrgStatus, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) AssignAdminRole(orgId string, acct common.Address, roleId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, acctId: acct, roleId: roleId, txa: txa}
	if execStatus := q.valAssignAdminRole(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	// check if account is already in use in another org
	tx, err := pinterf.AssignAdminRole(args.orgId, args.acctId, args.roleId)
	if err != nil {
		return reportExecError(AssignAdminRole, err)
	}
	log.Debug("executed permission action", "action", AssignAdminRole, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) ApproveAdminRole(orgId string, acct common.Address, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, acctId: acct, txa: txa}
	if execStatus := q.valApproveAdminRole(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	// check if anything is pending approval
	tx, err := pinterf.ApproveAdminRole(args.orgId, args.acctId)
	if err != nil {
		return reportExecError(ApproveAdminRole, err)
	}
	log.Debug("executed permission action", "action", ApproveAdminRole, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) AddNewRole(orgId string, roleId string, access uint8, isVoter bool, isAdmin bool, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, roleId: roleId, accessType: access, isVoter: isVoter, isAdmin: isAdmin, txa: txa}
	if execStatus := q.valAddNewRole(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	// check if role is already there in the org
	tx, err := pinterf.AddNewRole(args.roleId, args.orgId, big.NewInt(int64(args.accessType)), args.isVoter, args.isAdmin)
	if err != nil {
		return reportExecError(AddNewRole, err)
	}
	log.Debug("executed permission action", "action", AddNewRole, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) RemoveRole(orgId string, roleId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, roleId: roleId, txa: txa}

	if execStatus := q.valRemoveRole(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	tx, err := pinterf.RemoveRole(args.roleId, args.orgId)
	if err != nil {
		return reportExecError(RemoveRole, err)
	}
	log.Debug("executed permission action", "action", RemoveRole, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) AddAccountToOrg(acct common.Address, orgId string, roleId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, roleId: roleId, acctId: acct, txa: txa}

	if execStatus := q.valAssignRole(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	tx, err := pinterf.AssignAccountRole(args.acctId, args.orgId, args.roleId)
	if err != nil {
		return reportExecError(AddAccountToOrg, err)
	}
	log.Debug("executed permission action", "action", AddAccountToOrg, "tx", tx)
	return ExecSuccess.OpStatus()
}
func (q *QuorumControlsAPI) ChangeAccountRole(acct common.Address, orgId string, roleId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, roleId: roleId, acctId: acct, txa: txa}

	if execStatus := q.valAssignRole(args, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	tx, err := pinterf.AssignAccountRole(args.acctId, args.orgId, args.roleId)
	if err != nil {
		return reportExecError(ChangeAccountRole, err)
	}
	log.Debug("executed permission action", "action", ChangeAccountRole, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) UpdateAccountStatus(orgId string, acct common.Address, status uint8, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, acctId: acct, action: status, txa: txa}

	if execStatus := q.valUpdateAccountStatus(args, UpdateAccountStatus, pinterf); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	tx, err := pinterf.UpdateAccountStatus(args.orgId, args.acctId, big.NewInt(int64(args.action)))
	if err != nil {
		return reportExecError(UpdateAccountStatus, err)
	}
	log.Debug("executed permission action", "action", UpdateAccountStatus, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) RecoverBlackListedNode(orgId string, enodeId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, url: enodeId, txa: txa}

	if execStatus := q.valRecoverNode(args, pinterf, InitiateNodeRecovery); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.StartBlacklistedNodeRecovery(args.orgId, enodeId, ip, port, raftPort)
	if err != nil {
		return reportExecError(InitiateNodeRecovery, err)
	}
	log.Debug("executed permission action", "action", InitiateNodeRecovery, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) ApproveBlackListedNodeRecovery(orgId string, enodeId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, url: enodeId, txa: txa}

	if execStatus := q.valRecoverNode(args, pinterf, ApproveNodeRecovery); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.ApproveBlacklistedNodeRecovery(args.orgId, enodeId, ip, port, raftPort)
	if err != nil {
		return reportExecError(ApproveNodeRecovery, err)
	}
	log.Debug("executed permission action", "action", ApproveNodeRecovery, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) RecoverBlackListedAccount(orgId string, acctId common.Address, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, acctId: acctId, txa: txa}

	if execStatus := q.valRecoverAccount(args, pinterf, InitiateAccountRecovery); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	tx, err := pinterf.StartBlacklistedAccountRecovery(args.orgId, args.acctId)
	if err != nil {
		return reportExecError(InitiateAccountRecovery, err)
	}
	log.Debug("executed permission action", "action", InitiateAccountRecovery, "tx", tx)
	return ExecSuccess.OpStatus()
}

func (q *QuorumControlsAPI) ApproveBlackListedAccountRecovery(orgId string, acctId common.Address, txa ethapi.SendTxArgs) (string, error) {
	pinterf, execStatus := q.initOp(txa)
	if execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	args := txArgs{orgId: orgId, acctId: acctId, txa: txa}

	if execStatus := q.valRecoverAccount(args, pinterf, ApproveAccountRecovery); execStatus != ExecSuccess {
		return execStatus.OpStatus()
	}
	tx, err := pinterf.ApproveBlacklistedAccountRecovery(args.orgId, args.acctId)
	if err != nil {
		return reportExecError(ApproveAccountRecovery, err)
	}
	log.Debug("executed permission action", "action", ApproveAccountRecovery, "tx", tx)
	return ExecSuccess.OpStatus()
}

// check if the account is network admin
func (q *QuorumControlsAPI) isNetworkAdmin(account common.Address) bool {
	ac := types.AcctInfoMap.GetAccount(account)
	return ac != nil && ac.RoleId == q.permCtrl.permConfig.NwAdminRole
}

func (q *QuorumControlsAPI) isOrgAdmin(account common.Address, orgId string) (ExecStatus, error) {
	org := types.OrgInfoMap.GetOrg(orgId)
	if org == nil {
		return ErrOrgDoesNotExists, errors.New("invalid org")
	}
	ac := types.AcctInfoMap.GetAccount(account)
	if ac == nil {
		return ErrNotOrgAdmin, errors.New("not org admin")
	}
	// check if the account is network admin
	if !(ac.IsOrgAdmin && (ac.OrgId == orgId || ac.OrgId == org.UltimateParent)) {
		return ErrNotOrgAdmin, errors.New("not org admin")
	}
	return ExecSuccess, nil
}

func (q *QuorumControlsAPI) validateOrg(orgId, pOrgId string) (ExecStatus, error) {
	// validate Parent org id
	if pOrgId != "" {
		if types.OrgInfoMap.GetOrg(pOrgId) == nil {
			return ErrInvalidParentOrg, errors.New("invalid parent org")
		}
		locOrgId := pOrgId + "." + orgId
		if types.OrgInfoMap.GetOrg(locOrgId) != nil {
			return ErrOrgExists, errors.New("org exists")
		}
	} else if types.OrgInfoMap.GetOrg(orgId) != nil {
		return ErrOrgExists, errors.New("org exists")
	}
	return ExecSuccess, nil
}

func (q *QuorumControlsAPI) validatePendingOp(authOrg, orgId, url string, account common.Address, pendingOp int64, pinterf *pbind.PermInterfaceSession) bool {
	pOrg, pUrl, pAcct, op, err := pinterf.GetPendingOp(authOrg)
	return err == nil && (op.Int64() == pendingOp && pOrg == orgId && pUrl == url && pAcct == account)
}

func (q *QuorumControlsAPI) checkPendingOp(orgId string, pinterf *pbind.PermInterfaceSession) bool {
	_, _, _, op, err := pinterf.GetPendingOp(orgId)
	return err == nil && op.Int64() != 0
}

func (q *QuorumControlsAPI) checkOrgStatus(orgId string, op uint8) (ExecStatus, error) {
	org := types.OrgInfoMap.GetOrg(orgId)

	if org == nil {
		return ErrOrgDoesNotExists, errors.New("org does not exist")
	}
	// check if its a master org. operation is allowed only if its a master org
	if org.Level.Cmp(big.NewInt(1)) != 0 {
		return ErrNotMasterOrg, errors.New("org not a master org")
	}

	if !((op == 1 && org.Status == types.OrgApproved) || (op == 2 && org.Status == types.OrgSuspended)) {
		return ErrOpNotAllowed, errors.New("operation not allowed for current status")
	}
	return ExecSuccess, nil
}

func (q *QuorumControlsAPI) valNodeStatusChange(orgId, url string, op NodeUpdateAction, permAction PermAction) (ExecStatus, error) {
	// validates if the enode is linked the passed organization
	// validate node id and
	if len(url) == 0 {
		return ErrInvalidNode, errors.New("invalid node id")
	}
	if execStatus, err := q.valNodeDetails(url); err != nil && execStatus != ErrNodePresent {
		return execStatus, errors.New("node not found")
	}

	node := types.NodeInfoMap.GetNodeByUrl(url)
	if node != nil {
		if node.OrgId != orgId {
			return ErrNodeOrgMismatch, errors.New("node does not belong to the organization passed")
		}

		if node.Status == types.NodeBlackListed && op != RecoverBlacklistedNode {
			return ErrBlacklistedNode, errors.New("blacklisted node. operation not allowed")
		}

		// validate the op and node status and check if the op can be performed
		if (permAction == UpdateNodeStatus && (op != SuspendNode && op != ActivateSuspendedNode && op != BlacklistNode)) ||
			(permAction == InitiateNodeRecovery && op != RecoverBlacklistedNode) ||
			(permAction == ApproveNodeRecovery && op != ApproveBlacklistedNodeRecovery) {
			return ErrOpNotAllowed, errors.New("invalid node status change operation")
		}

		if (op == SuspendNode && node.Status != types.NodeApproved) ||
			(op == ActivateSuspendedNode && node.Status != types.NodeDeactivated) ||
			(op == BlacklistNode && node.Status == types.NodeRecoveryInitiated) ||
			(op == RecoverBlacklistedNode && node.Status != types.NodeBlackListed) ||
			(op == ApproveBlacklistedNodeRecovery && node.Status != types.NodeRecoveryInitiated) {
			return ErrOpNotAllowed, errors.New("node status change cannot be performed")
		}
	} else {
		return ErrNodeDoesNotExists, errors.New("node does not exist")
	}

	return ExecSuccess, nil
}

func (q *QuorumControlsAPI) validateRole(orgId, roleId string) bool {
	r := types.RoleInfoMap.GetRole(orgId, roleId)
	if r == nil {
		r = types.RoleInfoMap.GetRole(types.OrgInfoMap.GetOrg(orgId).UltimateParent, roleId)
	}

	return r != nil && r.Active
}

func (q *QuorumControlsAPI) valAccountStatusChange(orgId string, account common.Address, permAction PermAction, op AccountUpdateAction) (ExecStatus, error) {
	// validates if the enode is linked the passed organization
	ac := types.AcctInfoMap.GetAccount(account)

	if ac == nil {
		return ErrAccountNotThere, errors.New("account not there")
	}

	if ac.IsOrgAdmin && (ac.RoleId == q.permCtrl.permConfig.NwAdminRole || ac.RoleId == q.permCtrl.permConfig.OrgAdminRole) && (op == 1 || op == 3) {
		return ErrOpNotAllowed, errors.New("operation not allowed on org admin account")
	}

	if ac.OrgId != orgId {
		return ErrOrgNotOwner, errors.New("account does not belong to the organization passed")
	}
	if (permAction == UpdateAccountStatus && (op != SuspendAccount && op != ActivateSuspendedAccount && op != BlacklistAccount)) ||
		(permAction == InitiateAccountRecovery && op != RecoverBlacklistedAccount) ||
		(permAction == ApproveAccountRecovery && op != ApproveBlacklistedAccountRecovery) {
		return ErrOpNotAllowed, errors.New("invalid account status change operation")
	}

	if ac.Status == types.AcctBlacklisted && op != RecoverBlacklistedAccount {
		return ErrBlacklistedAccount, errors.New("blacklisted account. operation not allowed")
	}

	if (op == SuspendAccount && ac.Status != types.AcctActive) ||
		(op == ActivateSuspendedAccount && ac.Status != types.AcctSuspended) ||
		(op == BlacklistAccount && ac.Status == types.AcctRecoveryInitiated) ||
		(op == RecoverBlacklistedAccount && ac.Status != types.AcctBlacklisted) ||
		(op == ApproveBlacklistedAccountRecovery && ac.Status != types.AcctRecoveryInitiated) {
		return ErrOpNotAllowed, errors.New("account status change cannot be performed")
	}
	return ExecSuccess, nil
}

func (q *QuorumControlsAPI) checkOrgAdminExists(orgId, roleId string, account common.Address) (ExecStatus, error) {
	ac := types.AcctInfoMap.GetAccount(account)

	if ac != nil {
		if ac.OrgId != orgId {
			return ErrAccountInUse, errors.New("account part of another org")
		}
		if roleId != "" && roleId == q.permCtrl.permConfig.OrgAdminRole && ac.IsOrgAdmin {
			return ErrAccountOrgAdmin, errors.New("account already org admin for the org")
		}
	}
	return ExecSuccess, nil
}

func (q *QuorumControlsAPI) valSubOrgBreadthDepth(porgId string) (ExecStatus, error) {
	org := types.OrgInfoMap.GetOrg(porgId)

	if q.permCtrl.permConfig.SubOrgDepth.Cmp(org.Level) == 0 {
		return ErrMaxDepth, errors.New("max depth for sub orgs reached")
	}

	if q.permCtrl.permConfig.SubOrgBreadth.Cmp(big.NewInt(int64(len(org.SubOrgList)))) == 0 {
		return ErrMaxBreadth, errors.New("max breadth for sub orgs reached")
	}

	return ExecSuccess, nil
}

func (q *QuorumControlsAPI) checkNodeExists(url, enodeId string) bool {
	node := types.NodeInfoMap.GetNodeByUrl(url)
	if node != nil {
		return true
	}
	// check if the same nodeid is in use with different port numbers
	nodeList := types.NodeInfoMap.GetNodeList()
	for _, n := range nodeList {
		if enodeDet, er := enode.ParseV4(n.Url); er == nil {
			if enodeDet.EnodeID() == enodeId {
				return true
			}
		}
	}
	return false
}

func (q *QuorumControlsAPI) valNodeDetails(url string) (ExecStatus, error) {
	// validate node id and
	if len(url) != 0 {
		enodeDet, err := enode.ParseV4(url)
		if err != nil {
			return ErrInvalidNode, errors.New("invalid node id")
		}
		// check if node already there
		if q.checkNodeExists(url, enodeDet.EnodeID()) {
			return ErrNodePresent, errors.New("duplicate node")
		}
	}
	return ExecSuccess, nil
}

func (q *QuorumControlsAPI) getNodeDetails(url string) (string, [32]byte, uint16, uint16, error) {
	// validate node id and
	var ip [32]byte
	if len(url) == 0 {
		return "", ip, 0, 0, errors.New("invalid node id")
	}
	enodeDet, err := enode.ParseV4(url)
	if err != nil {
		return "", ip, 0, 0, errors.New("invalid node id")
	}
	enodeId, ip, port, raftport := enodeDet.NodeDetails()
	return enodeId, ip, port, raftport, err
}

// all validations for add org operation
func (q *QuorumControlsAPI) valAddOrg(args txArgs, pinterf *pbind.PermInterfaceSession) ExecStatus {
	// check if the org id contains "."
	if args.orgId == "" || args.url == "" || args.acctId == (common.Address{0}) {
		return ErrInvalidInput
	}
	if !isStringAlphaNumeric(args.orgId) {
		return ErrInvalidOrgName
	}

	// check if caller is network admin
	if !q.isNetworkAdmin(args.txa.From) {
		return ErrNotNetworkAdmin
	}

	// check if any previous op is pending approval for network admin
	if q.checkPendingOp(q.permCtrl.permConfig.NwAdminOrg, pinterf) {
		return ErrPendingApprovals
	}
	// check if org already exists
	if execStatus, er := q.validateOrg(args.orgId, ""); er != nil {
		return execStatus
	}

	// validate node id and
	if execStatus, er := q.valNodeDetails(args.url); er != nil {
		return execStatus
	}

	// check if account is already part of another org
	if execStatus, er := q.checkOrgAdminExists(args.orgId, "", args.acctId); er != nil {
		return execStatus
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valApproveOrg(args txArgs, pinterf *pbind.PermInterfaceSession) ExecStatus {
	// check caller is network admin
	if !q.isNetworkAdmin(args.txa.From) {
		return ErrNotNetworkAdmin
	}
	enodeId, _, _, _, _ := q.getNodeDetails(args.url)
	// check if anything pending approval
	if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.orgId, enodeId, args.acctId, 1, pinterf) {
		return ErrNothingToApprove
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valAddSubOrg(args txArgs, pinterf *pbind.PermInterfaceSession) ExecStatus {
	// check if the org id contains "."
	if args.orgId == "" {
		return ErrInvalidInput
	}
	if !isStringAlphaNumeric(args.orgId) {
		return ErrInvalidOrgName
	}

	// check if caller is network admin
	if execStatus, er := q.isOrgAdmin(args.txa.From, args.porgId); er != nil {
		return execStatus
	}

	// check if org already exists
	if execStatus, er := q.validateOrg(args.orgId, args.porgId); er != nil {
		return execStatus
	}

	if execStatus, er := q.valSubOrgBreadthDepth(args.porgId); er != nil {
		return execStatus
	}

	if execStatus, er := q.valNodeDetails(args.url); er != nil {
		return execStatus
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valUpdateOrgStatus(args txArgs, pinterf *pbind.PermInterfaceSession) ExecStatus {
	// check if called is network admin
	if !q.isNetworkAdmin(args.txa.From) {
		return ErrNotNetworkAdmin
	}
	if OrgUpdateAction(args.action) != SuspendOrg &&
		OrgUpdateAction(args.action) != ActivateSuspendedOrg {
		return ErrOpNotAllowed
	}

	//check if passed org id is network admin org. update should not be allowed
	if args.orgId == q.permCtrl.permConfig.NwAdminOrg {
		return ErrOpNotAllowed
	}
	// check if status update can be performed. Org should be approved for suspension
	if execStatus, er := q.checkOrgStatus(args.orgId, args.action); er != nil {
		return execStatus
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valApproveOrgStatus(args txArgs, pinterf *pbind.PermInterfaceSession) ExecStatus {
	// check if called is network admin
	if !q.isNetworkAdmin(args.txa.From) {
		return ErrNotNetworkAdmin
	}
	// check if anything is pending approval
	var pendingOp int64
	if args.action == 1 {
		pendingOp = 2
	} else if args.action == 2 {
		pendingOp = 3
	} else {
		return ErrOpNotAllowed
	}
	if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.orgId, "", common.Address{}, pendingOp, pinterf) {
		return ErrNothingToApprove
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valAddNode(args txArgs, pinterf *pbind.PermInterfaceSession) ExecStatus {
	if args.url == "" {
		return ErrInvalidInput
	}
	// check if caller is network admin
	if execStatus, er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return execStatus
	}

	if execStatus, er := q.valNodeDetails(args.url); er != nil {
		return execStatus
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valUpdateNodeStatus(args txArgs, permAction PermAction, pinterf *pbind.PermInterfaceSession) ExecStatus {
	// check if org admin
	// check if caller is network admin
	if execStatus, er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return execStatus
	}

	// validation status change is with in allowed set
	if execStatus, er := q.valNodeStatusChange(args.orgId, args.url, NodeUpdateAction(args.action), permAction); er != nil {
		return execStatus
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valAssignAdminRole(args txArgs, pinterf *pbind.PermInterfaceSession) ExecStatus {
	if args.acctId == (common.Address{0}) {
		return ErrInvalidInput
	}
	// check if caller is network admin
	if args.roleId != q.permCtrl.permConfig.OrgAdminRole && args.roleId != q.permCtrl.permConfig.NwAdminRole {
		return ErrOpNotAllowed
	}

	if !q.isNetworkAdmin(args.txa.From) {
		return ErrNotNetworkAdmin
	}

	if _, err := q.validateOrg(args.orgId, ""); err == nil {
		return ErrOrgDoesNotExists
	}

	// check if account is already part of another org
	if execStatus, er := q.checkOrgAdminExists(args.orgId, args.roleId, args.acctId); er != nil && execStatus != ErrOrgAdminExists {
		return execStatus
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valApproveAdminRole(args txArgs, pinterf *pbind.PermInterfaceSession) ExecStatus {
	// check if caller is network admin
	if !q.isNetworkAdmin(args.txa.From) {
		return ErrNotNetworkAdmin
	}
	// check if the org exists

	// check if account is valid
	ac := types.AcctInfoMap.GetAccount(args.acctId)
	if ac == nil {
		return ErrInvalidAccount
	}
	// validate pending op
	if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, ac.OrgId, "", args.acctId, 4, pinterf) {
		return ErrNothingToApprove
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valAddNewRole(args txArgs, pinterf *pbind.PermInterfaceSession) ExecStatus {
	if args.roleId == "" {
		return ErrInvalidInput
	}
	// check if caller is network admin
	if execStatus, er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return execStatus
	}
	// validate if role is already present
	if types.RoleInfoMap.GetRole(args.orgId, args.roleId) != nil {
		return ErrRoleExists
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valRemoveRole(args txArgs, pinterf *pbind.PermInterfaceSession) ExecStatus {
	// check if caller is network admin
	if execStatus, er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return execStatus
	}

	// admin roles cannot be removed
	if args.roleId == q.permCtrl.permConfig.OrgAdminRole || args.roleId == q.permCtrl.permConfig.NwAdminRole {
		return ErrAdminRoles
	}

	// check if role is alraedy inactive
	r := types.RoleInfoMap.GetRole(args.orgId, args.roleId)
	if r == nil {
		return ErrInvalidRole
	} else if !r.Active {
		return ErrInactiveRole
	}

	// check if the role has active accounts. if yes operations should not be allowed
	if len(types.AcctInfoMap.GetAcctListRole(args.orgId, args.roleId)) != 0 {
		return ErrRoleActive
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valAssignRole(args txArgs, pinterf *pbind.PermInterfaceSession) ExecStatus {
	if args.acctId == (common.Address{0}) {
		return ErrInvalidInput
	}
	if args.roleId == q.permCtrl.permConfig.OrgAdminRole || args.roleId == q.permCtrl.permConfig.NwAdminRole {
		return ErrInvalidRole
	}
	// check if caller is network admin
	if execStatus, er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return execStatus
	}

	// check if the role is valid
	if !q.validateRole(args.orgId, args.roleId) {
		return ErrInvalidRole
	}

	// check if the account is part of another org
	if ac := types.AcctInfoMap.GetAccount(args.acctId); ac != nil {
		if ac.OrgId != args.orgId {
			return ErrAccountInUse
		}
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valUpdateAccountStatus(args txArgs, permAction PermAction, pinterf *pbind.PermInterfaceSession) ExecStatus {
	// check if the caller is org admin
	if execStatus, er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return execStatus
	}
	// validation status change is with in allowed set
	if execStatus, er := q.valAccountStatusChange(args.orgId, args.acctId, permAction, AccountUpdateAction(args.action)); er != nil {
		return execStatus
	}
	return ExecSuccess
}

func (q *QuorumControlsAPI) valRecoverNode(args txArgs, pinterf *pbind.PermInterfaceSession, action PermAction) ExecStatus {
	// check if the caller is org admin
	if !q.isNetworkAdmin(args.txa.From) {
		return ErrNotNetworkAdmin
	}
	// validate inputs - org id is valid, node is valid and in blacklisted state
	if execStatus, _ := q.validateOrg(args.orgId, ""); execStatus != ErrOrgExists {
		return ErrInvalidOrgName
	}

	if action == InitiateNodeRecovery {
		if execStatus, _ := q.valNodeStatusChange(args.orgId, args.url, 4, InitiateAccountRecovery); execStatus != ExecSuccess {
			return execStatus
		}
		// check no pending approval items
		if q.checkPendingOp(q.permCtrl.permConfig.NwAdminOrg, pinterf) {
			return ErrPendingApprovals
		}
	} else {
		// validate inputs - org id is valid, node is valid pending recovery state
		if execStatus, _ := q.valNodeStatusChange(args.orgId, args.url, 5, ApproveNodeRecovery); execStatus != ExecSuccess {
			return execStatus
		}
		enodeId, _, _, _, _ := q.getNodeDetails(args.url)
		// check that there is a pending approval item for node recovery
		if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.orgId, enodeId, common.Address{}, 5, pinterf) {
			return ErrNothingToApprove
		}
	}

	// if it is approval ensure that

	return ExecSuccess
}

func (q *QuorumControlsAPI) valRecoverAccount(args txArgs, pinterf *pbind.PermInterfaceSession, action PermAction) ExecStatus {
	// check if the caller is org admin
	if !q.isNetworkAdmin(args.txa.From) {
		return ErrNotNetworkAdmin
	}

	var opAction AccountUpdateAction
	if action == InitiateAccountRecovery {
		opAction = RecoverBlacklistedAccount
	} else {
		opAction = ApproveBlacklistedAccountRecovery
	}

	if execStatus, err := q.valAccountStatusChange(args.orgId, args.acctId, action, opAction); err != nil {
		return execStatus
	}

	if action == InitiateAccountRecovery && q.checkPendingOp(q.permCtrl.permConfig.NwAdminOrg, pinterf) {
		return ErrPendingApprovals
	}

	if action == ApproveAccountRecovery && !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.orgId, "", args.acctId, 6, pinterf) {
		return ErrNothingToApprove
	}
	return ExecSuccess
}

// validateAccount validates the account and returns the wallet associated with that for signing the transaction
func (q *QuorumControlsAPI) validateAccount(from common.Address) (accounts.Wallet, error) {
	acct := accounts.Account{Address: from}
	w, err := q.permCtrl.eth.AccountManager().Find(acct)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (q *QuorumControlsAPI) newPermInterfaceSession(w accounts.Wallet, txa ethapi.SendTxArgs) *pbind.PermInterfaceSession {
	frmAcct, transactOpts, gasLimit, gasPrice := q.getTxParams(txa, w)
	ps := &pbind.PermInterfaceSession{
		Contract: q.permCtrl.permInterf,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     frmAcct.Address,
			GasLimit: gasLimit,
			GasPrice: gasPrice,
			Signer:   transactOpts.Signer,
		},
	}
	return ps
}

// getTxParams extracts the transaction related parameters
func (q *QuorumControlsAPI) getTxParams(txa ethapi.SendTxArgs, w accounts.Wallet) (accounts.Account, *bind.TransactOpts, uint64, *big.Int) {
	fromAcct := accounts.Account{Address: txa.From}
	transactOpts := bind.NewWalletTransactor(w, fromAcct)
	gasLimit := defaultGasLimit
	gasPrice := defaultGasPrice
	if txa.GasPrice != nil {
		gasPrice = txa.GasPrice.ToInt()
	}
	if txa.Gas != nil {
		gasLimit = uint64(*txa.Gas)
	}
	return fromAcct, transactOpts, gasLimit, gasPrice
}
