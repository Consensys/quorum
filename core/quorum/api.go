package quorum

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	obind "github.com/ethereum/go-ethereum/controls/bind/cluster"
	pbind "github.com/ethereum/go-ethereum/controls/bind/permission"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"math/big"
	"regexp"
	"strings"
)

var isStringAlphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9_-]*$`).MatchString

//default gas limit to use if not passed in sendTxArgs
var defaultGasLimit = uint64(470000000)

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
	AssignOrgAdminAccount
	ApproveOrgAdminAccount
	AddNewRole
	RemoveRole
	AssignAccountRole
	UpdateAccountStatus
)

// OrgKeyAction represents an action in cluster contract
type OrgKeyAction int

// return values for checkNodeDetails function
type NodeCheckRetVal int

const (
	Success NodeCheckRetVal = iota
)

// Voter access type
type VoterAccessType uint8

const (
	Active VoterAccessType = iota
	Inactive
)

type PermissionContracts struct {
	PermInterf *pbind.PermInterface
}

// QuorumControlsAPI provides an API to access Quorum's node permission and org key management related services
type QuorumControlsAPI struct {
	txPool      *core.TxPool
	ethClnt     *ethclient.Client
	acntMgr     *accounts.Manager
	txOpt       *bind.TransactOpts
	clustContr  *obind.Cluster
	key         *ecdsa.PrivateKey
	permEnabled bool
	orgEnabled  bool
	permConfig  *types.PermissionConfig
	permInterf  *pbind.PermInterface
}

// txArgs holds arguments required for execute functions
type txArgs struct {
	orgId      string
	porgId     string
	url        string
	roleId     string
	isVoter    bool
	acctId     common.Address
	accessType uint8
	status     uint8
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

var (
	ErrNotNetworkAdmin    = ExecStatus{false, "Operation can be performed by network admin only. Account not a network admin."}
	ErrNotOrgAdmin        = ExecStatus{false, "Operation can be performed by org admin only. Account not a org admin."}
	ErrNodePresent        = ExecStatus{false, "EnodeId already part of network."}
	ErrInvalidNode        = ExecStatus{false, "Invalid enode id"}
	ErrInvalidAccount     = ExecStatus{false, "Invalid account id"}
	ErrFailedExecution    = ExecStatus{false, "Failed to execute permission action"}
	ErrPermissionDisabled = ExecStatus{false, "Permissions control not enabled"}
	ErrAccountAccess      = ExecStatus{false, "Account does not have sufficient access for operation"}
	ErrVoterAccountAccess = ExecStatus{false, "Voter account does not have sufficient access"}
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
	ErrRoleDoesNotExist   = ExecStatus{false, "Role not found for org. Add role first"}
	ErrRoleActive         = ExecStatus{false, "Accounts linked to the role. Cannot be removed"}
	ErrAdminRoles         = ExecStatus{false, "Admin role cannot be removed"}
	ErrInvalidOrgName     = ExecStatus{false, "Org id cannot contain special characters"}
	ErrInvalidParentOrg   = ExecStatus{false, "Invalid parent org id"}
	ErrAccountNotThere    = ExecStatus{false, "Account does not exists"}
	ErrOrgNotOwner        = ExecStatus{false, "Account does not belong to this org"}
	ExecSuccess           = ExecStatus{true, "Action completed successfully"}
)

// NewQuorumControlsAPI creates a new QuorumControlsAPI to access quorum services
func NewQuorumControlsAPI(tp *core.TxPool, am *accounts.Manager) *QuorumControlsAPI {
	return &QuorumControlsAPI{tp, nil, am, nil, nil, nil, false, false, nil, nil}
}

//Init initializes QuorumControlsAPI with eth client, permission contract and org key management control
func (p *QuorumControlsAPI) Init(ethClnt *ethclient.Client, key *ecdsa.PrivateKey, apiName string, pconfig *types.PermissionConfig, pc *pbind.PermInterface) error {
	// check if the interface contract is deployed or not. if not
	// permissions apis will not work. return error
	p.ethClnt = ethClnt
	p.permConfig = pconfig

	if _, err := pbind.NewPermInterface(p.permConfig.InterfAddress, p.ethClnt); err != nil {
		return err
	}
	p.permEnabled = true
	p.key = key
	p.permInterf = pc

	return nil
}

func (s *QuorumControlsAPI) OrgList() []types.OrgInfo {
	return types.OrgInfoMap.GetOrgList()
}

func (s *QuorumControlsAPI) NodeList() []types.NodeInfo {
	return types.NodeInfoMap.GetNodeList()
}

func (s *QuorumControlsAPI) RoleList() []types.RoleInfo {
	return types.RoleInfoMap.GetRoleList()
}

func (s *QuorumControlsAPI) AcctList() []types.AccountInfo {
	return types.AcctInfoMap.GetAcctList()
}

func (s *QuorumControlsAPI) GetOrgDetails(orgId string) types.OrgDetailInfo {
	var acctList []types.AccountInfo
	var roleList []types.RoleInfo
	var nodeList []types.NodeInfo
	for _, a := range s.AcctList() {
		if a.OrgId == orgId {
			acctList = append(acctList, a)
		}
	}
	for _, a := range s.RoleList() {
		if a.OrgId == orgId {
			roleList = append(roleList, a)
		}
	}
	for _, a := range s.NodeList() {
		if a.OrgId == orgId {
			nodeList = append(nodeList, a)
		}
	}
	return types.OrgDetailInfo{NodeList: nodeList, RoleList: roleList, AcctList: acctList, SubOrgList: types.OrgInfoMap.GetOrg(orgId).SubOrgList}
}

func (s *QuorumControlsAPI) AddOrg(orgId string, url string, acct common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(AddOrg, txArgs{orgId: orgId, url: url, acctId: acct, txa: txa})
}

func (s *QuorumControlsAPI) AddSubOrg(porgId, orgId string, url string, acct common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(AddSubOrg, txArgs{porgId: porgId, orgId: orgId, url: url, acctId: acct, txa: txa})
}

func (s *QuorumControlsAPI) ApproveOrg(orgId string, url string, acct common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveOrg, txArgs{orgId: orgId, url: url, acctId: acct, txa: txa})
}

func (s *QuorumControlsAPI) UpdateOrgStatus(orgId string, status uint8, txa ethapi.SendTxArgs) ExecStatus {
	log.Info("AJ-update org status", "org", orgId, "status", status)
	return s.executePermAction(UpdateOrgStatus, txArgs{orgId: orgId, status: status, txa: txa})
}

func (s *QuorumControlsAPI) AddNode(orgId string, url string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(AddNode, txArgs{orgId: orgId, url: url, txa: txa})
}

func (s *QuorumControlsAPI) UpdateNodeStatus(orgId string, url string, status uint8, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(UpdateNodeStatus, txArgs{orgId: orgId, url: url, status: status, txa: txa})
}

func (s *QuorumControlsAPI) ApproveOrgStatus(orgId string, status uint8, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveOrgStatus, txArgs{orgId: orgId, status: status, txa: txa})
}

func (s *QuorumControlsAPI) AssignOrgAdminAccount(orgId string, acct common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(AssignOrgAdminAccount, txArgs{orgId: orgId, acctId: acct, txa: txa})
}

func (s *QuorumControlsAPI) ApproveOrgAdminAccount(acct common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveOrgAdminAccount, txArgs{acctId: acct, txa: txa})
}

func (s *QuorumControlsAPI) AddNewRole(orgId string, roleId string, access uint8, isVoter bool, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(AddNewRole, txArgs{orgId: orgId, roleId: roleId, accessType: access, isVoter: isVoter, txa: txa})
}

func (s *QuorumControlsAPI) RemoveRole(orgId string, roleId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(RemoveRole, txArgs{orgId: orgId, roleId: roleId, txa: txa})
}

func (s *QuorumControlsAPI) AssignAccountRole(acct common.Address, orgId string, roleId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(AssignAccountRole, txArgs{orgId: orgId, roleId: roleId, acctId: acct, txa: txa})
}

func (s *QuorumControlsAPI) UpdateAccountStatus(orgId string, acct common.Address, status uint8, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(UpdateAccountStatus, txArgs{orgId: orgId, acctId: acct, status: status, txa: txa})
}

// check if the account is network admin
func (s *QuorumControlsAPI) isNetworkAdmin(account common.Address) bool {
	ac := types.AcctInfoMap.GetAccount(account)
	return ac != nil && ac.RoleId == s.permConfig.NwAdminRole
}

//TODO (Amal) get it reviewed by Sai
func (s *QuorumControlsAPI) isOrgAdmin(account common.Address, orgId string) bool {
	ac := types.AcctInfoMap.GetAccount(account)
	return ac != nil && ((ac.OrgId == s.permConfig.NwAdminOrg && ac.RoleId == s.permConfig.NwAdminRole) ||
		(ac.RoleId == s.permConfig.OrgAdminRole && strings.Contains(orgId, ac.OrgId)))
}

func (s *QuorumControlsAPI) validateOrg(orgId, pOrgId string) (ExecStatus, error) {
	// validate Parent org id
	if pOrgId != "" && types.OrgInfoMap.GetOrg(pOrgId) == nil {
		return ErrInvalidParentOrg, errors.New("invalid parent org")
	} else {
		locOrgId := pOrgId + "." + orgId
		if types.OrgInfoMap.GetOrg(locOrgId) != nil {
			return ErrOrgExists, errors.New("org exists")
		}
	}
	return ExecSuccess, nil
}

func (s *QuorumControlsAPI) checkNodeExists(enodeId string) bool {
	node := types.NodeInfoMap.GetNodeByUrl(enodeId)
	return node != nil
}

func (s *QuorumControlsAPI) validatePendingOp(authOrg, orgId, url string, account common.Address, pendingOp int64, pinterf *pbind.PermInterfaceSession) bool {
	pOrg, pUrl, pAcct, op, err := pinterf.GetPendingOp(authOrg)
	return err == nil && (op.Int64() == pendingOp && pOrg == orgId && pUrl == url && pAcct == account)
}

func (s *QuorumControlsAPI) checkPendingOp(orgId string, pinterf *pbind.PermInterfaceSession) bool {
	_, _, _, op, err := pinterf.GetPendingOp(orgId)
	return err == nil && op.Int64() != 0
}

func (s *QuorumControlsAPI) checkOrgStatus(orgId string, op uint8) bool {
	org := types.OrgInfoMap.GetOrg(orgId)
	return (op == 3 && org.Status == types.OrgApproved) || (op == 5 && org.Status == types.OrgSuspended)
}

func (s *QuorumControlsAPI) valNodeStatusChange(orgId, url string, op int64) (ExecStatus, error) {
	// validates if the enode is linked the passed organization
	node := types.NodeInfoMap.GetNodeByUrl(url)

	if node.OrgId != orgId {
		return ErrNodeOrgMismatch, errors.New("node does not belong to the organization passed")
	}

	if node.Status == types.NodeBlackListed {
		return ErrBlacklistedNode, errors.New("blacklisted node. operation not allowed")
	}

	// validate the op and node status and check if the op can be performed
	if op != 3 && op != 4 && op != 5 {
		return ErrOpNotAllowed, errors.New("invalid node status change operation")
	}

	if (op == 3 && node.Status != types.NodeApproved) || (op == 4 && node.Status != types.NodeDeactivated) {
		return ErrOpNotAllowed, errors.New("node status change cannot be performed")
	}
	return ExecSuccess, nil
}

func (s *QuorumControlsAPI) valAccountStatusChange(orgId string, account common.Address, op int64) (ExecStatus, error) {
	// validates if the enode is linked the passed organization
	ac := types.AcctInfoMap.GetAccount(account)

	if ac == nil {
		return ErrAccountNotThere, errors.New("account not there")
	}

	if ac.IsOrgAdmin && (op == 1 || op == 3) {
		return ErrOpNotAllowed, errors.New("operation not allowed on org admin account")
	}

	if ac.OrgId != orgId {
		return ErrOrgNotOwner, errors.New("account does not belong to the organization passed")
	}

	if ac.Status == types.AcctBlacklisted {
		return ErrBlacklistedAccount, errors.New("blacklisted account. operation not allowed")
	}

	// validate the op and node status and check if the op can be performed
	if op != 1 && op != 2 && op != 3 {
		return ErrOpNotAllowed, errors.New("invalid account status change operation")
	}

	if (op == 1 && ac.Status != types.AcctActive) || (op == 2 && ac.Status != types.AcctSuspended) {
		return ErrOpNotAllowed, errors.New("account status change cannot be performed")
	}
	return ExecSuccess, nil
}

func (s *QuorumControlsAPI) checkOrgAdminExists(orgId string, account common.Address) (ExecStatus, error) {
	ac := types.AcctInfoMap.GetAccount(account)

	if ac == nil {
		orgAcctList := types.AcctInfoMap.GetAcctListOrg(orgId)
		if len(orgAcctList) > 0 {
			for _, a := range orgAcctList {
				if a.IsOrgAdmin == true {
					return ErrOrgAdminExists, errors.New("org admin exists for the org")
				}
			}
		}
	} else {
		if ac.OrgId != orgId {
			return ErrAccountInUse, errors.New("account part of another org")
		}
		if ac.IsOrgAdmin == true {
			return ErrAccountOrgAdmin, errors.New("account already org admin for the org")
		}
	}
	return ExecSuccess, nil
}

// executePermAction helps to execute an action in permission contract
func (s *QuorumControlsAPI) executePermAction(action PermAction, args txArgs) ExecStatus {

	if !s.permEnabled {
		return ErrPermissionDisabled
	}
	var err error
	var w accounts.Wallet

	w, err = s.validateAccount(args.txa.From)
	if err != nil {
		return ErrInvalidAccount
	}

	pinterf := s.newPermInterfaceSession(w, args.txa)
	var tx *types.Transaction

	switch action {

	case AddOrg:
		// check if the org id contains "."
		if !isStringAlphaNumeric(args.orgId) {
			return ErrInvalidOrgName
		}

		// check if caller is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}

		// check if any previous op is pending approval for network admin
		if s.checkPendingOp(s.permConfig.NwAdminOrg, pinterf) {
			return ErrPendingApprovals
		}
		// check if org already exists
		if execStatus, er := s.validateOrg(args.orgId, ""); er != nil {
			return execStatus
		}

		// validate node id and
		_, err := enode.ParseV4(args.url)
		if err != nil {
			return ErrInvalidNode
		}

		// check if node already there
		if s.checkNodeExists(args.url) {
			return ErrNodePresent
		}

		// check if account is already part of another org
		if execStatus, er := s.checkOrgAdminExists(args.orgId, args.acctId); er != nil {
			return execStatus
		}

		tx, err = pinterf.AddOrg(args.orgId, args.url, args.acctId)

	case ApproveOrg:
		// check caller is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}

		if !s.validatePendingOp(s.permConfig.NwAdminOrg, args.orgId, args.url, args.acctId, 1, pinterf) {
			return ErrNothingToApprove
		}

		// check if anything pending approval
		tx, err = pinterf.ApproveOrg(args.orgId, args.url, args.acctId)

	case AddSubOrg:
		// check if the org id contains "."
		if !isStringAlphaNumeric(args.orgId) {
			return ErrInvalidOrgName
		}

		// check if caller is network admin
		if !s.isOrgAdmin(args.txa.From, args.porgId) {
			return ErrNotOrgAdmin
		}

		// check if org already exists
		if execStatus, er := s.validateOrg(args.orgId, args.porgId); er != nil {
			return execStatus
		}

		// validate node id and
		if len(args.url) != 0 {
			_, err := enode.ParseV4(args.url)
			if err != nil {
				return ErrInvalidNode
			}
			// check if node already there
			if s.checkNodeExists(args.url) {
				return ErrNodePresent
			}
		}

		// check if account is already part of another org
		if (args.acctId != common.Address{}) {
			if execStatus, er := s.checkOrgAdminExists(args.orgId, args.acctId); er != nil {
				return execStatus
			}
		}

		tx, err = pinterf.AddSubOrg(args.porgId, args.orgId, args.url, args.acctId)

	case UpdateOrgStatus:
		// check if called is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}

		// check if status update can be performed. Org should be approved for suspension
		if !s.checkOrgStatus(args.orgId, args.status) {
			return ErrOpNotAllowed
		}

		if args.status != 3 && args.status != 5 {
			return ErrOpNotAllowed
		}

		// and in suspended state for suspension revoke
		tx, err = pinterf.UpdateOrgStatus(args.orgId, big.NewInt(int64(args.status)))

	case ApproveOrgStatus:
		// check if called is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}

		// check if anything is pending approval
		var pendingOp int64
		if args.status == 3 {
			pendingOp = 2
		} else if args.status == 5 {
			pendingOp = 3
		} else {
			return ErrOpNotAllowed
		}
		if !s.validatePendingOp(s.permConfig.NwAdminOrg, args.orgId, "", common.Address{}, pendingOp, pinterf) {
			return ErrNothingToApprove
		}

		// validate that status change is pending approval
		tx, err = pinterf.ApproveOrgStatus(args.orgId, big.NewInt(int64(args.status)))

	case AddNode:
		// check if org admin
		if !s.isOrgAdmin(args.txa.From, args.orgId) {
			return ErrNotOrgAdmin
		}

		// validate node id and
		_, err := enode.ParseV4(args.url)
		if err != nil {
			return ErrInvalidNode
		}

		// check if node is already there
		tx, err = pinterf.AddNode(args.orgId, args.url)

	case UpdateNodeStatus:
		// check if org admin
		if !s.isOrgAdmin(args.txa.From, args.orgId) {
			return ErrNotOrgAdmin
		}

		// validate node id and
		_, err := enode.ParseV4(args.url)
		if err != nil {
			return ErrInvalidNode
		}

		// validation status change is with in allowed set
		if execStatus, er := s.valNodeStatusChange(args.orgId, args.url, int64(args.status)); er != nil {
			return execStatus
		}

		// check node status for operation
		tx, err = pinterf.UpdateNodeStatus(args.orgId, args.url, big.NewInt(int64(args.status)))

	case AssignOrgAdminAccount:
		// check if caller is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}
		// check if account is already part of another org
		if execStatus, er := s.checkOrgAdminExists(args.orgId, args.acctId); er != nil {
			return execStatus
		}
		// check if account is already in use in another org
		tx, err = pinterf.AssignOrgAdminAccount(args.orgId, args.acctId)

	case ApproveOrgAdminAccount:
		// check if caller is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}

		// validate pending op
		if !s.validatePendingOp(s.permConfig.NwAdminOrg, types.AcctInfoMap.GetAccount(args.acctId).OrgId, "", args.acctId, 4, pinterf) {
			return ErrNothingToApprove
		}

		// check if anything is pending approval
		tx, err = pinterf.ApproveOrgAdminAccount(args.acctId)

	case AddNewRole:
		// check if org admin
		if !s.isOrgAdmin(args.txa.From, args.orgId) {
			return ErrNotOrgAdmin
		}
		// validate if role is already present
		if types.RoleInfoMap.GetRole(args.orgId, args.roleId) != nil {
			return ErrRoleExists
		}

		// check if role is already there in the org
		tx, err = pinterf.AddNewRole(args.roleId, args.orgId, big.NewInt(int64(args.accessType)), args.isVoter)

	case RemoveRole:
		// check if org admin
		if !s.isOrgAdmin(args.txa.From, args.orgId) {
			return ErrNotOrgAdmin
		}

		// admin roles cannot be removed
		if args.roleId == s.permConfig.OrgAdminRole || args.roleId == s.permConfig.NwAdminRole {
			return ErrAdminRoles
		}

		// check if the role has active accounts. if yes operations should not be allowed
		if len(types.AcctInfoMap.GetAcctListRole(args.orgId, args.roleId)) != 0 {
			return ErrRoleActive
		}

		tx, err = pinterf.RemoveRole(args.roleId, args.orgId)

	case AssignAccountRole:
		// check if org admin
		if !s.isOrgAdmin(args.txa.From, args.orgId) {
			return ErrNotOrgAdmin
		}

		// check if the role is part of the org
		if types.RoleInfoMap.GetRole(args.orgId, args.roleId) == nil {
			// check if the role is existing at master org level
			if types.RoleInfoMap.GetRole(types.OrgInfoMap.GetOrg(args.orgId).UltimateParent, args.roleId) == nil {
				return ErrRoleDoesNotExist
			}
		}

		// check if the account is part of another org
		if ac := types.AcctInfoMap.GetAccount(args.acctId); ac != nil {
			if ac.OrgId != args.orgId {
				return ErrAccountInUse
			}
		}

		tx, err = pinterf.AssignAccountRole(args.acctId, args.orgId, args.roleId)

	case UpdateAccountStatus:
		// validation status change is with in allowed set
		if execStatus, er := s.valAccountStatusChange(args.orgId, args.acctId, int64(args.status)); er != nil {
			return execStatus
		}

		tx, err = pinterf.UpdateAccountStatus(args.orgId, args.acctId, big.NewInt(int64(args.status)))
	}

	if err != nil {
		log.Error("Failed to execute permission action", "action", action, "err", err)
		return ErrFailedExecution
	}
	log.Debug("executed permission action", "action", action, "tx", tx)
	return ExecSuccess
}

// validateAccount validates the account and returns the wallet associated with that for signing the transaction
func (s *QuorumControlsAPI) validateAccount(from common.Address) (accounts.Wallet, error) {
	acct := accounts.Account{Address: from}
	w, err := s.acntMgr.Find(acct)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (s *QuorumControlsAPI) newPermInterfaceSession(w accounts.Wallet, txa ethapi.SendTxArgs) *pbind.PermInterfaceSession {
	frmAcct, transactOpts, gasLimit, gasPrice, nonce := s.getTxParams(txa, w)
	ps := &pbind.PermInterfaceSession{
		Contract: s.permInterf,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     frmAcct.Address,
			GasLimit: gasLimit,
			GasPrice: gasPrice,
			Signer:   transactOpts.Signer,
			Nonce:    nonce,
		},
	}
	return ps
}

// getTxParams extracts the transaction related parameters
func (s *QuorumControlsAPI) getTxParams(txa ethapi.SendTxArgs, w accounts.Wallet) (accounts.Account, *bind.TransactOpts, uint64, *big.Int, *big.Int) {
	frmAcct := accounts.Account{Address: txa.From}
	transactOpts := bind.NewWalletTransactor(w, frmAcct)
	gasLimit := defaultGasLimit
	gasPrice := defaultGasPrice
	if txa.GasPrice != nil {
		gasPrice = txa.GasPrice.ToInt()
	}
	if txa.Gas != nil {
		gasLimit = uint64(*txa.Gas)
	}
	var nonce *big.Int
	if txa.Nonce != nil {
		nonce = new(big.Int).SetUint64(uint64(*txa.Nonce))
	} else {
		nonce = new(big.Int).SetUint64(s.txPool.Nonce(frmAcct.Address))
	}
	return frmAcct, transactOpts, gasLimit, gasPrice, nonce
}

// checks if the account performing the operation has sufficient access privileges
func valAccountAccessVoter(fromAcct, targetAcct common.Address) (error, ExecStatus) {
	acctAccess := types.GetAcctAccess(fromAcct)
	// only accounts with full access will be allowed to manage voters
	if acctAccess != types.FullAccess {
		return errors.New("Account performing the operation does not have sufficient access"), ErrAccountAccess
	}

	// An account with minimum of transact access can be a voter
	if targetAcct != (common.Address{}) {
		acctAccess = types.GetAcctAccess(targetAcct)
		if acctAccess == types.ReadOnly {
			return errors.New("Voter account does not have sufficient access"), ErrVoterAccountAccess
		}
	}
	return nil, ExecSuccess
}
