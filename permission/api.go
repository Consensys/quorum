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

var actionSuccess = "Action completed successfully"

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
	o, err := types.OrgInfoMap.GetOrg(orgId)
	if err != nil {
		return types.OrgDetailInfo{}, err
	}

	if o == nil {
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
	orgRec, err := types.OrgInfoMap.GetOrg(orgId)
	if err != nil {
		return types.OrgDetailInfo{}, err
	}

	if orgRec == nil {
		return types.OrgDetailInfo{NodeList: nodeList, RoleList: roleList, AcctList: acctList}, nil
	}
	return types.OrgDetailInfo{NodeList: nodeList, RoleList: roleList, AcctList: acctList, SubOrgList: orgRec.SubOrgList}, nil
}

func (q *QuorumControlsAPI) initOp(txa ethapi.SendTxArgs) (*pbind.PermInterfaceSession, error) {
	var err error
	var w accounts.Wallet

	w, err = q.validateAccount(txa.From)
	if err != nil {
		return nil, types.ErrInvalidAccount
	}
	pinterf := q.newPermInterfaceSession(w, txa)

	return pinterf, nil
}

func reportExecError(action PermAction, err error) (string, error) {
	log.Error("Failed to execute permission action", "action", action, "err", err)
	msg := fmt.Sprintf("failed to execute permissions action: %v", err)
	return "", errors.New(msg)
}

func (q *QuorumControlsAPI) AddOrg(orgId string, url string, acct common.Address, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, url: url, acctId: acct, txa: txa}

	if err := q.valAddOrg(args, pinterf); err != nil {
		return "", err
	}
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.AddOrg(args.orgId, enodeId, ip, port, raftPort, args.acctId)
	if err != nil {
		return reportExecError(AddOrg, err)
	}
	log.Debug("executed permission action", "action", AddOrg, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) AddSubOrg(porgId, orgId string, url string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{porgId: porgId, orgId: orgId, url: url, txa: txa}

	if err := q.valAddSubOrg(args, pinterf); err != nil {
		return "", err
	}
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.AddSubOrg(args.porgId, args.orgId, enodeId, ip, port, raftPort)
	if err != nil {
		return reportExecError(AddSubOrg, err)
	}
	log.Debug("executed permission action", "action", AddSubOrg, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) ApproveOrg(orgId string, url string, acct common.Address, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, url: url, acctId: acct, txa: txa}
	if err := q.valApproveOrg(args, pinterf); err != nil {
		return "", err
	}
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.ApproveOrg(args.orgId, enodeId, ip, port, raftPort, args.acctId)
	if err != nil {
		return reportExecError(ApproveOrg, err)
	}
	log.Debug("executed permission action", "action", ApproveOrg, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) UpdateOrgStatus(orgId string, status uint8, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, action: status, txa: txa}
	if err := q.valUpdateOrgStatus(args, pinterf); err != nil {
		return "", err
	}
	// and in suspended state for suspension revoke
	tx, err := pinterf.UpdateOrgStatus(args.orgId, big.NewInt(int64(args.action)))
	if err != nil {
		return reportExecError(UpdateOrgStatus, err)
	}
	log.Debug("executed permission action", "action", UpdateOrgStatus, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) AddNode(orgId string, url string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, url: url, txa: txa}
	if err := q.valAddNode(args, pinterf); err != nil {
		return "", err
	}
	// check if node is already there
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.AddNode(args.orgId, enodeId, ip, port, raftPort)
	if err != nil {
		return reportExecError(AddNode, err)
	}
	log.Debug("executed permission action", "action", AddNode, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) UpdateNodeStatus(orgId string, url string, action uint8, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, url: url, action: action, txa: txa}
	if err := q.valUpdateNodeStatus(args, UpdateNodeStatus, pinterf); err != nil {
		return "", err
	}
	// check node status for operation
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.UpdateNodeStatus(args.orgId, enodeId, ip, port, raftPort, big.NewInt(int64(args.action)))
	if err != nil {
		return reportExecError(UpdateNodeStatus, err)
	}
	log.Debug("executed permission action", "action", UpdateNodeStatus, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) ApproveOrgStatus(orgId string, status uint8, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, action: status, txa: txa}
	if err := q.valApproveOrgStatus(args, pinterf); err != nil {
		return "", err
	}
	// validate that status change is pending approval
	tx, err := pinterf.ApproveOrgStatus(args.orgId, big.NewInt(int64(args.action)))
	if err != nil {
		return reportExecError(ApproveOrgStatus, err)
	}
	log.Debug("executed permission action", "action", ApproveOrgStatus, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) AssignAdminRole(orgId string, acct common.Address, roleId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, acctId: acct, roleId: roleId, txa: txa}
	if err := q.valAssignAdminRole(args, pinterf); err != nil {
		return "", err
	}
	// check if account is already in use in another org
	tx, err := pinterf.AssignAdminRole(args.orgId, args.acctId, args.roleId)
	if err != nil {
		return reportExecError(AssignAdminRole, err)
	}
	log.Debug("executed permission action", "action", AssignAdminRole, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) ApproveAdminRole(orgId string, acct common.Address, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, acctId: acct, txa: txa}
	if err := q.valApproveAdminRole(args, pinterf); err != nil {
		return "", err
	}
	// check if anything is pending approval
	tx, err := pinterf.ApproveAdminRole(args.orgId, args.acctId)
	if err != nil {
		return reportExecError(ApproveAdminRole, err)
	}
	log.Debug("executed permission action", "action", ApproveAdminRole, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) AddNewRole(orgId string, roleId string, access uint8, isVoter bool, isAdmin bool, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, roleId: roleId, accessType: access, isVoter: isVoter, isAdmin: isAdmin, txa: txa}
	if err := q.valAddNewRole(args, pinterf); err != nil {
		return "", err
	}
	// check if role is already there in the org
	tx, err := pinterf.AddNewRole(args.roleId, args.orgId, big.NewInt(int64(args.accessType)), args.isVoter, args.isAdmin)
	if err != nil {
		return reportExecError(AddNewRole, err)
	}
	log.Debug("executed permission action", "action", AddNewRole, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) RemoveRole(orgId string, roleId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, roleId: roleId, txa: txa}

	if err := q.valRemoveRole(args, pinterf); err != nil {
		return "", err
	}
	tx, err := pinterf.RemoveRole(args.roleId, args.orgId)
	if err != nil {
		return reportExecError(RemoveRole, err)
	}
	log.Debug("executed permission action", "action", RemoveRole, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) TransactionAllowed(srcacct common.Address, tgtacct common.Address, txa ethapi.SendTxArgs) (bool, error) {
	pinterf, execStatus := q.initOp(txa)
	_, err := execStatus.OpStatus()
	if err != nil {
		return false, err
	}
	return pinterf.TransactionAllowed(srcacct, tgtacct)
}

func (q *QuorumControlsAPI) ConnectionAllowed(url string, txa ethapi.SendTxArgs) (bool, error) {
	pinterf, execStatus := q.initOp(txa)
	_, err := execStatus.OpStatus()
	if err != nil {
		return false, err
	}
	var enodeId string
	var port uint16
	var raftport uint16
	var ip [32]byte
	enodeId, ip, port, raftport, err = q.getNodeDetails(url)
	return pinterf.ConnectionAllowedImpl(enodeId, ip, port, raftport)
}

func (q *QuorumControlsAPI) AddAccountToOrg(acct common.Address, orgId string, roleId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, roleId: roleId, acctId: acct, txa: txa}

	if err := q.valAssignRole(args, pinterf); err != nil {
		return "", err
	}
	tx, err := pinterf.AssignAccountRole(args.acctId, args.orgId, args.roleId)
	if err != nil {
		return reportExecError(AddAccountToOrg, err)
	}
	log.Debug("executed permission action", "action", AddAccountToOrg, "tx", tx)
	return actionSuccess, nil
}
func (q *QuorumControlsAPI) ChangeAccountRole(acct common.Address, orgId string, roleId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, roleId: roleId, acctId: acct, txa: txa}

	if err := q.valAssignRole(args, pinterf); err != nil {
		return "", err
	}
	tx, err := pinterf.AssignAccountRole(args.acctId, args.orgId, args.roleId)
	if err != nil {
		return reportExecError(ChangeAccountRole, err)
	}
	log.Debug("executed permission action", "action", ChangeAccountRole, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) UpdateAccountStatus(orgId string, acct common.Address, status uint8, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, acctId: acct, action: status, txa: txa}

	if err := q.valUpdateAccountStatus(args, UpdateAccountStatus, pinterf); err != nil {
		return "", err
	}
	tx, err := pinterf.UpdateAccountStatus(args.orgId, args.acctId, big.NewInt(int64(args.action)))
	if err != nil {
		return reportExecError(UpdateAccountStatus, err)
	}
	log.Debug("executed permission action", "action", UpdateAccountStatus, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) RecoverBlackListedNode(orgId string, enodeId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, url: enodeId, txa: txa}

	if err := q.valRecoverNode(args, pinterf, InitiateNodeRecovery); err != nil {
		return "", err
	}
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.StartBlacklistedNodeRecovery(args.orgId, enodeId, ip, port, raftPort)
	if err != nil {
		return reportExecError(InitiateNodeRecovery, err)
	}
	log.Debug("executed permission action", "action", InitiateNodeRecovery, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) ApproveBlackListedNodeRecovery(orgId string, enodeId string, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, url: enodeId, txa: txa}

	if err := q.valRecoverNode(args, pinterf, ApproveNodeRecovery); err != nil {
		return "", err
	}
	enodeId, ip, port, raftPort, _ := q.getNodeDetails(args.url)
	tx, err := pinterf.ApproveBlacklistedNodeRecovery(args.orgId, enodeId, ip, port, raftPort)
	if err != nil {
		return reportExecError(ApproveNodeRecovery, err)
	}
	log.Debug("executed permission action", "action", ApproveNodeRecovery, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) RecoverBlackListedAccount(orgId string, acctId common.Address, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, acctId: acctId, txa: txa}

	if err := q.valRecoverAccount(args, pinterf, InitiateAccountRecovery); err != nil {
		return "", err
	}
	tx, err := pinterf.StartBlacklistedAccountRecovery(args.orgId, args.acctId)
	if err != nil {
		return reportExecError(InitiateAccountRecovery, err)
	}
	log.Debug("executed permission action", "action", InitiateAccountRecovery, "tx", tx)
	return actionSuccess, nil
}

func (q *QuorumControlsAPI) ApproveBlackListedAccountRecovery(orgId string, acctId common.Address, txa ethapi.SendTxArgs) (string, error) {
	pinterf, err := q.initOp(txa)
	if err != nil {
		return "", err
	}
	args := txArgs{orgId: orgId, acctId: acctId, txa: txa}

	if err := q.valRecoverAccount(args, pinterf, ApproveAccountRecovery); err != nil {
		return "", err
	}
	tx, err := pinterf.ApproveBlacklistedAccountRecovery(args.orgId, args.acctId)
	if err != nil {
		return reportExecError(ApproveAccountRecovery, err)
	}
	log.Debug("executed permission action", "action", ApproveAccountRecovery, "tx", tx)
	return actionSuccess, nil
}

// check if the account is network admin
func (q *QuorumControlsAPI) isNetworkAdmin(account common.Address) bool {
	ac, _ := types.AcctInfoMap.GetAccount(account)
	return ac != nil && ac.RoleId == q.permCtrl.permConfig.NwAdminRole
}

func (q *QuorumControlsAPI) isOrgAdmin(account common.Address, orgId string) error {
	org, err := types.OrgInfoMap.GetOrg(orgId)
	if err != nil {
		return err
	}
	if org == nil {
		return types.ErrOrgDoesNotExists
	}
	ac, _ := types.AcctInfoMap.GetAccount(account)
	if ac == nil {
		return types.ErrNotOrgAdmin
	}
	// check if the account is network admin
	if !(ac.IsOrgAdmin && (ac.OrgId == orgId || ac.OrgId == org.UltimateParent)) {
		return types.ErrNotOrgAdmin
	}
	return nil
}

func (q *QuorumControlsAPI) validateOrg(orgId, pOrgId string) error {
	// validate Parent org id
	if pOrgId != "" {
		if _, err := types.OrgInfoMap.GetOrg(pOrgId); err != nil {
			return types.ErrInvalidParentOrg
		}
		locOrgId := pOrgId + "." + orgId
		if lorgRec, _ := types.OrgInfoMap.GetOrg(locOrgId); lorgRec != nil {
			return types.ErrOrgExists
		}
	} else if orgRec, _ := types.OrgInfoMap.GetOrg(orgId); orgRec != nil {
		return types.ErrOrgExists
	}
	return nil
}

func (q *QuorumControlsAPI) validatePendingOp(authOrg, orgId, url string, account common.Address, pendingOp int64, pinterf *pbind.PermInterfaceSession) bool {
	pOrg, pUrl, pAcct, op, err := pinterf.GetPendingOp(authOrg)
	return err == nil && (op.Int64() == pendingOp && pOrg == orgId && pUrl == url && pAcct == account)
}

func (q *QuorumControlsAPI) checkPendingOp(orgId string, pinterf *pbind.PermInterfaceSession) bool {
	_, _, _, op, err := pinterf.GetPendingOp(orgId)
	return err == nil && op.Int64() != 0
}

func (q *QuorumControlsAPI) checkOrgStatus(orgId string, op uint8) error {
	org, _ := types.OrgInfoMap.GetOrg(orgId)

	if org == nil {
		return types.ErrOrgDoesNotExists
	}
	// check if its a master org. operation is allowed only if its a master org
	if org.Level.Cmp(big.NewInt(1)) != 0 {
		return types.ErrNotMasterOrg
	}

	if !((op == 1 && org.Status == types.OrgApproved) || (op == 2 && org.Status == types.OrgSuspended)) {
		return types.ErrOpNotAllowed
	}
	return nil
}

func (q *QuorumControlsAPI) valNodeStatusChange(orgId, url string, op NodeUpdateAction, permAction PermAction) error {
	// validates if the enode is linked the passed organization
	// validate node id and
	if len(url) == 0 {
		return types.ErrInvalidNode
	}
	if err := q.valNodeDetails(url); err != nil && err.Error() != types.ErrNodePresent.Error() {
		return err
	}

	node, err := types.NodeInfoMap.GetNodeByUrl(url)
	if err != nil {
		return err
	}

	if node.OrgId != orgId {
		return types.ErrNodeOrgMismatch
	}

	if node.Status == types.NodeBlackListed && op != RecoverBlacklistedNode {
		return types.ErrBlacklistedNode
	}

	// validate the op and node status and check if the op can be performed
	if (permAction == UpdateNodeStatus && (op != SuspendNode && op != ActivateSuspendedNode && op != BlacklistNode)) ||
		(permAction == InitiateNodeRecovery && op != RecoverBlacklistedNode) ||
		(permAction == ApproveNodeRecovery && op != ApproveBlacklistedNodeRecovery) {
		return types.ErrOpNotAllowed
	}

	if (op == SuspendNode && node.Status != types.NodeApproved) ||
		(op == ActivateSuspendedNode && node.Status != types.NodeDeactivated) ||
		(op == BlacklistNode && node.Status == types.NodeRecoveryInitiated) ||
		(op == RecoverBlacklistedNode && node.Status != types.NodeBlackListed) ||
		(op == ApproveBlacklistedNodeRecovery && node.Status != types.NodeRecoveryInitiated) {
		return types.ErrOpNotAllowed
	}

	return nil
}

func (q *QuorumControlsAPI) validateRole(orgId, roleId string) bool {
	var r *types.RoleInfo
	r, err := types.RoleInfoMap.GetRole(orgId, roleId)
	if err != nil {
		return false
	}

	orgRec, err := types.OrgInfoMap.GetOrg(orgId)
	if err != nil {
		return false
	}
	r, err = types.RoleInfoMap.GetRole(orgRec.UltimateParent, roleId)
	if err != nil {
		return false
	}

	return r != nil && r.Active
}

func (q *QuorumControlsAPI) valAccountStatusChange(orgId string, account common.Address, permAction PermAction, op AccountUpdateAction) error {
	// validates if the enode is linked the passed organization
	ac, err := types.AcctInfoMap.GetAccount(account)
	if err != nil {
		return err
	}

	if ac.IsOrgAdmin && (ac.RoleId == q.permCtrl.permConfig.NwAdminRole || ac.RoleId == q.permCtrl.permConfig.OrgAdminRole) && (op == 1 || op == 3) {
		return types.ErrOpNotAllowed
	}

	if ac.OrgId != orgId {
		return types.ErrOrgNotOwner
	}
	if (permAction == UpdateAccountStatus && (op != SuspendAccount && op != ActivateSuspendedAccount && op != BlacklistAccount)) ||
		(permAction == InitiateAccountRecovery && op != RecoverBlacklistedAccount) ||
		(permAction == ApproveAccountRecovery && op != ApproveBlacklistedAccountRecovery) {
		return types.ErrOpNotAllowed
	}

	if ac.Status == types.AcctBlacklisted && op != RecoverBlacklistedAccount {
		return types.ErrBlacklistedAccount
	}

	if (op == SuspendAccount && ac.Status != types.AcctActive) ||
		(op == ActivateSuspendedAccount && ac.Status != types.AcctSuspended) ||
		(op == BlacklistAccount && ac.Status == types.AcctRecoveryInitiated) ||
		(op == RecoverBlacklistedAccount && ac.Status != types.AcctBlacklisted) ||
		(op == ApproveBlacklistedAccountRecovery && ac.Status != types.AcctRecoveryInitiated) {
		return types.ErrOpNotAllowed
	}
	return nil
}

func (q *QuorumControlsAPI) checkOrgAdminExists(orgId, roleId string, account common.Address) error {
	if ac, _ := types.AcctInfoMap.GetAccount(account); ac != nil {
		if ac.OrgId != orgId {
			return types.ErrAccountInUse
		}
		if roleId != "" && roleId == q.permCtrl.permConfig.OrgAdminRole && ac.IsOrgAdmin {
			return types.ErrAccountOrgAdmin
		}
	}
	return nil
}

func (q *QuorumControlsAPI) valSubOrgBreadthDepth(porgId string) error {
	org, err := types.OrgInfoMap.GetOrg(porgId)
	if err != nil {
		return types.ErrOpNotAllowed
	}

	if q.permCtrl.permConfig.SubOrgDepth.Cmp(org.Level) == 0 {
		return types.ErrMaxDepth
	}

	if q.permCtrl.permConfig.SubOrgBreadth.Cmp(big.NewInt(int64(len(org.SubOrgList)))) == 0 {
		return types.ErrMaxBreadth
	}

	return nil
}

func (q *QuorumControlsAPI) checkNodeExists(url, enodeId string) bool {
	node, _ := types.NodeInfoMap.GetNodeByUrl(url)
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

func (q *QuorumControlsAPI) valNodeDetails(url string) error {
	// validate node id and
	if len(url) != 0 {
		enodeDet, err := enode.ParseV4(url)
		if err != nil {
			return types.ErrInvalidNode
		}
		// check if node already there
		if q.checkNodeExists(url, enodeDet.EnodeID()) {
			return types.ErrNodePresent
		}
	}
	return nil
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
func (q *QuorumControlsAPI) valAddOrg(args txArgs, pinterf *pbind.PermInterfaceSession) error {
	// check if the org id contains "."
	if args.orgId == "" || args.url == "" || args.acctId == (common.Address{0}) {
		return types.ErrInvalidInput
	}
	if !isStringAlphaNumeric(args.orgId) {
		return types.ErrInvalidOrgName
	}

	// check if caller is network admin
	if !q.isNetworkAdmin(args.txa.From) {
		return types.ErrNotNetworkAdmin
	}

	// check if any previous op is pending approval for network admin
	if q.checkPendingOp(q.permCtrl.permConfig.NwAdminOrg, pinterf) {
		return types.ErrPendingApprovals
	}
	// check if org already exists
	if er := q.validateOrg(args.orgId, ""); er != nil {
		return er
	}

	// validate node id and
	if er := q.valNodeDetails(args.url); er != nil {
		return er
	}

	// check if account is already part of another org
	if er := q.checkOrgAdminExists(args.orgId, "", args.acctId); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valApproveOrg(args txArgs, pinterf *pbind.PermInterfaceSession) error {
	// check caller is network admin
	if !q.isNetworkAdmin(args.txa.From) {
		return types.ErrNotNetworkAdmin
	}
	enodeId, _, _, _, _ := q.getNodeDetails(args.url)
	// check if anything pending approval
	if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.orgId, enodeId, args.acctId, 1, pinterf) {
		return ErrNothingToApprove
	}
	return nil
}

func (q *QuorumControlsAPI) valAddSubOrg(args txArgs, pinterf *pbind.PermInterfaceSession) error {
	// check if the org id contains "."
	if args.orgId == "" {
		return types.ErrInvalidInput
	}
	if !isStringAlphaNumeric(args.orgId) {
		return types.ErrInvalidOrgName
	}

	// check if caller is network admin
	if er := q.isOrgAdmin(args.txa.From, args.porgId); er != nil {
		return er
	}

	// check if org already exists
	if er := q.validateOrg(args.orgId, args.porgId); er != nil {
		return er
	}

	if er := q.valSubOrgBreadthDepth(args.porgId); er != nil {
		return er
	}

	if er := q.valNodeDetails(args.url); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valUpdateOrgStatus(args txArgs, pinterf *pbind.PermInterfaceSession) error {
	// check if called is network admin
	if !q.isNetworkAdmin(args.txa.From) {
		return types.ErrNotNetworkAdmin
	}
	if OrgUpdateAction(args.action) != SuspendOrg &&
		OrgUpdateAction(args.action) != ActivateSuspendedOrg {
		return types.ErrOpNotAllowed
	}

	//check if passed org id is network admin org. update should not be allowed
	if args.orgId == q.permCtrl.permConfig.NwAdminOrg {
		return types.ErrOpNotAllowed
	}
	// check if status update can be performed. Org should be approved for suspension
	if er := q.checkOrgStatus(args.orgId, args.action); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valApproveOrgStatus(args txArgs, pinterf *pbind.PermInterfaceSession) error {
	// check if called is network admin
	if !q.isNetworkAdmin(args.txa.From) {
		return types.ErrNotNetworkAdmin
	}
	// check if anything is pending approval
	var pendingOp int64
	if args.action == 1 {
		pendingOp = 2
	} else if args.action == 2 {
		pendingOp = 3
	} else {
		return types.ErrOpNotAllowed
	}
	if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.orgId, "", common.Address{}, pendingOp, pinterf) {
		return types.ErrNothingToApprove
	}
	return nil
}

func (q *QuorumControlsAPI) valAddNode(args txArgs, pinterf *pbind.PermInterfaceSession) error {
	if args.url == "" {
		return types.ErrInvalidInput
	}
	// check if caller is network admin
	if er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return er
	}

	if er := q.valNodeDetails(args.url); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valUpdateNodeStatus(args txArgs, permAction PermAction, pinterf *pbind.PermInterfaceSession) error {
	// check if org admin
	// check if caller is network admin
	if er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return er
	}

	// validation status change is with in allowed set
	if er := q.valNodeStatusChange(args.orgId, args.url, NodeUpdateAction(args.action), permAction); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valAssignAdminRole(args txArgs, pinterf *pbind.PermInterfaceSession) error {
	if args.acctId == (common.Address{0}) {
		return types.ErrInvalidInput
	}
	// check if caller is network admin
	if args.roleId != q.permCtrl.permConfig.OrgAdminRole && args.roleId != q.permCtrl.permConfig.NwAdminRole {
		return types.ErrOpNotAllowed
	}

	if !q.isNetworkAdmin(args.txa.From) {
		return types.ErrNotNetworkAdmin
	}

	if err := q.validateOrg(args.orgId, ""); err == nil {
		return types.ErrOrgDoesNotExists
	}

	// check if account is already part of another org
	if er := q.checkOrgAdminExists(args.orgId, args.roleId, args.acctId); er != nil && er.Error() != types.ErrOrgAdminExists.Error() {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valApproveAdminRole(args txArgs, pinterf *pbind.PermInterfaceSession) error {
	// check if caller is network admin
	if !q.isNetworkAdmin(args.txa.From) {
		return types.ErrNotNetworkAdmin
	}
	// check if the org exists

	// check if account is valid
	ac, _ := types.AcctInfoMap.GetAccount(args.acctId)
	if ac == nil {
		return types.ErrInvalidAccount
	}
	// validate pending op
	if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, ac.OrgId, "", args.acctId, 4, pinterf) {
		return types.ErrNothingToApprove
	}
	return nil
}

func (q *QuorumControlsAPI) valAddNewRole(args txArgs, pinterf *pbind.PermInterfaceSession) error {
	if args.roleId == "" {
		return types.ErrInvalidInput
	}
	// check if caller is network admin
	if er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return er
	}
	// validate if role is already present
	if roleRec, _ := types.RoleInfoMap.GetRole(args.orgId, args.roleId); roleRec != nil {
		return types.ErrRoleExists
	}
	return nil
}

func (q *QuorumControlsAPI) valRemoveRole(args txArgs, pinterf *pbind.PermInterfaceSession) error {
	// check if caller is network admin
	if er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return er
	}

	// admin roles cannot be removed
	if args.roleId == q.permCtrl.permConfig.OrgAdminRole || args.roleId == q.permCtrl.permConfig.NwAdminRole {
		return types.ErrAdminRoles
	}

	// check if role is alraedy inactive
	r, _ := types.RoleInfoMap.GetRole(args.orgId, args.roleId)
	if r == nil {
		return types.ErrInvalidRole
	} else if !r.Active {
		return types.ErrInactiveRole
	}

	// check if the role has active accounts. if yes operations should not be allowed
	if len(types.AcctInfoMap.GetAcctListRole(args.orgId, args.roleId)) != 0 {
		return types.ErrRoleActive
	}
	return nil
}

func (q *QuorumControlsAPI) valAssignRole(args txArgs, pinterf *pbind.PermInterfaceSession) error {
	if args.acctId == (common.Address{0}) {
		return types.ErrInvalidInput
	}
	if args.roleId == q.permCtrl.permConfig.OrgAdminRole || args.roleId == q.permCtrl.permConfig.NwAdminRole {
		return types.ErrInvalidRole
	}
	// check if caller is network admin
	if er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return er
	}

	// check if the role is valid
	if !q.validateRole(args.orgId, args.roleId) {
		return types.ErrInvalidRole
	}

	// check if the account is part of another org
	if ac, _ := types.AcctInfoMap.GetAccount(args.acctId); ac != nil {
		if ac != nil && ac.OrgId != args.orgId {
			return types.ErrAccountInUse
		}
	}
	return nil
}

func (q *QuorumControlsAPI) valUpdateAccountStatus(args txArgs, permAction PermAction, pinterf *pbind.PermInterfaceSession) error {
	// check if the caller is org admin
	if er := q.isOrgAdmin(args.txa.From, args.orgId); er != nil {
		return er
	}
	// validation status change is with in allowed set
	if er := q.valAccountStatusChange(args.orgId, args.acctId, permAction, AccountUpdateAction(args.action)); er != nil {
		return er
	}
	return nil
}

func (q *QuorumControlsAPI) valRecoverNode(args txArgs, pinterf *pbind.PermInterfaceSession, action PermAction) error {
	// check if the caller is org admin
	if !q.isNetworkAdmin(args.txa.From) {
		return types.ErrNotNetworkAdmin
	}
	// validate inputs - org id is valid, node is valid and in blacklisted state
	if err := q.validateOrg(args.orgId, ""); err.Error() != types.ErrOrgExists.Error() {
		return types.ErrInvalidOrgName
	}

	if action == InitiateNodeRecovery {
		if err := q.valNodeStatusChange(args.orgId, args.url, 4, InitiateAccountRecovery); err != nil {
			return err
		}
		// check no pending approval items
		if q.checkPendingOp(q.permCtrl.permConfig.NwAdminOrg, pinterf) {
			return types.ErrPendingApprovals
		}
	} else {
		// validate inputs - org id is valid, node is valid pending recovery state
		if err := q.valNodeStatusChange(args.orgId, args.url, 5, ApproveNodeRecovery); err != nil {
			return err
		}
		enodeId, _, _, _, _ := q.getNodeDetails(args.url)
		// check that there is a pending approval item for node recovery
		if !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.orgId, enodeId, common.Address{}, 5, pinterf) {
			return ErrNothingToApprove
		}
	}

	// if it is approval ensure that

	return nil
}

func (q *QuorumControlsAPI) valRecoverAccount(args txArgs, pinterf *pbind.PermInterfaceSession, action PermAction) error {
	// check if the caller is org admin
	if !q.isNetworkAdmin(args.txa.From) {
		return types.ErrNotNetworkAdmin
	}

	var opAction AccountUpdateAction
	if action == InitiateAccountRecovery {
		opAction = RecoverBlacklistedAccount
	} else {
		opAction = ApproveBlacklistedAccountRecovery
	}

	if err := q.valAccountStatusChange(args.orgId, args.acctId, action, opAction); err != nil {
		return err
	}

	if action == InitiateAccountRecovery && q.checkPendingOp(q.permCtrl.permConfig.NwAdminOrg, pinterf) {
		return types.ErrPendingApprovals
	}

	if action == ApproveAccountRecovery && !q.validatePendingOp(q.permCtrl.permConfig.NwAdminOrg, args.orgId, "", args.acctId, 6, pinterf) {
		return types.ErrNothingToApprove
	}
	return nil
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
