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
)

//default gas limit to use if not passed in sendTxArgs
var defaultGasLimit = uint64(470000000)

//default gas price to use if not passed in sendTxArgs
var defaultGasPrice = big.NewInt(0)

// PermAction represents actions in permission contract
type PermAction int

const (
	AddOrg PermAction = iota
	ApproveOrg
	UpdateOrgStatus
	ApproveOrgStatus
	AddNode
	UpdateNodeStatus
	AssignOrgAdminAccount
	ApproveOrgAdminAccount
	AddNewRole
	RemoveRole
	AssignAccountRole
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

type nodeStatus struct {
	EnodeId string `json:"enodeId"`
	Status  string `json:"status"`
}

type accountInfo struct {
	Address string `json:"address"`
	Access  string `json:"access"`
}

type orgInfo struct {
	MasterOrgId   string   `json:"masterOrgId"`
	SubOrgId      string   `json:"subOrgId"`
	SubOrgKeyList []string `json:"subOrgKeyList"`
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
	ErrInvalidNode        = ExecStatus{false, "Invalid node id"}
	ErrInvalidAccount     = ExecStatus{false, "Invalid account id"}
	ErrFailedExecution    = ExecStatus{false, "Failed to execute permission action"}
	ErrPermissionDisabled = ExecStatus{false, "Permissions control not enabled"}
	ErrAccountAccess      = ExecStatus{false, "Account does not have sufficient access for operation"}
	ErrVoterAccountAccess = ExecStatus{false, "Voter account does not have sufficient access"}
	ErrOrgExists          = ExecStatus{false, "Org already exists"}
	ErrPendingApprovals   = ExecStatus{false, "Pending approvals for the organization. Approve first"}
	ErrNothingToApprove   = ExecStatus{false, "Nothing to approve"}
	ErrOpNotAllowed       = ExecStatus{false, "Operation not allowed"}
	ExecSuccess           = ExecStatus{true, "Action completed successfully"}

	//ErrNoVoterAccount       = ExecStatus{false, "No voter account registered. Add voter first"}
	//ErrAccountNotAVoter     = ExecStatus{false, "Account is not a voter. Action cannot be approved"}
	//ErrInvalidAccountAccess = ExecStatus{false, "Invalid account access type"}
	//ErrNodeDetailsMismatch  = ExecStatus{false, "Node details mismatch"}
	//ErrOrgDisabled          = ExecStatus{false, "Org key management not enabled for the network"}
	//ErrMasterOrgExists      = ExecStatus{false, "Master org already exists"}
	//ErrInvalidMasterOrg     = ExecStatus{false, "Master org does not exist. Add master org first"}
	//ErrInvalidOrg           = ExecStatus{false, "Org does not exist. Add org first"}
	//ErrVoterExists          = ExecStatus{false, "Voter account exists"}
	//ErrKeyExists            = ExecStatus{false, "Key exists for the organization"}
	//ErrKeyInUse             = ExecStatus{false, "Key already in use in another master organization"}
	//ErrKeyNotFound          = ExecStatus{false, "Key not found for the organization"}
	//ErrNothingToCancel      = ExecStatus{false, "Nothing to cancel"}
	//ErrNodeProposed         = ExecStatus{false, "Node already proposed for the action"}
	//ErrAccountIsNotVoter    = ExecStatus{false, "Not a voter account"}
	//ErrBlacklistedNode      = ExecStatus{false, "Blacklisted node. Operation not allowed"}
	//ErrLastFullAccessAcct   = ExecStatus{false, "Last account with full access. Operation not allowed"}
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

func (s *QuorumControlsAPI) AddOrg(orgId string, url string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(AddOrg, txArgs{orgId: orgId, url: url, txa: txa})
}

func (s *QuorumControlsAPI) ApproveOrg(orgId string, url string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveOrg, txArgs{orgId: orgId, url: url, txa: txa})
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

func (s *QuorumControlsAPI) AddNewRole(roleId string, orgId string, access uint8, isVoter bool, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(AddNewRole, txArgs{orgId: orgId, roleId: roleId, accessType: access, isVoter: isVoter, txa: txa})
}

func (s *QuorumControlsAPI) RemoveRole(roleId string, orgId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(RemoveRole, txArgs{orgId: orgId, roleId: roleId, txa: txa})
}

func (s *QuorumControlsAPI) AssignAccountRole(acct common.Address, orgId string, roleId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(AssignAccountRole, txArgs{orgId: orgId, roleId: roleId, acctId: acct, txa: txa})
}

// check if the account is network admin
func (s *QuorumControlsAPI) isNetworkAdmin(account common.Address) bool {
	ac := types.AcctInfoMap.GetAccount(account)
	return ac != nil && ac.RoleId == s.permConfig.NwAdminRole
}

func (s *QuorumControlsAPI) isOrgAdmin(account common.Address, orgId string) bool {
	ac := types.AcctInfoMap.GetAccount(account)
	return ac != nil && (ac.RoleId == s.permConfig.OrgAdminRole && ac.OrgId == orgId)
}

func (s *QuorumControlsAPI) checkOrgExists(orgId string) bool {
	org := types.OrgInfoMap.GetOrg(orgId)
	return org != nil
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
		// check if caller is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}

		// check if any previous op is pending approval for network admin
		if s.checkPendingOp(s.permConfig.NwAdminOrg, pinterf) {
			return ErrPendingApprovals
		}
		// check if org already exists
		if s.checkOrgExists(args.orgId) {
			return ErrOrgExists
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

		tx, err = pinterf.AddOrg(args.orgId, args.url)

	case ApproveOrg:
		// check caller is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}

		if !s.validatePendingOp(s.permConfig.NwAdminOrg, args.orgId, args.url, common.Address{}, 1, pinterf) {
			return ErrNothingToApprove
		}

		// check if anything pending approval
		tx, err = pinterf.ApproveOrg(args.orgId, args.url)

	case UpdateOrgStatus:
		// check if called is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}

		if !s.checkOrgStatus(args.orgId, args.status) {
			return ErrOpNotAllowed
		}

		// check if status update can be performed. Org should be approved for suspension

		// and in suspended state for suspension revoke

		tx, err = pinterf.UpdateOrgStatus(args.orgId, big.NewInt(int64(args.status)))

	case ApproveOrgStatus:
		// check if called is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}

		// check if anything is pending approval

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

		// check node status for operation
		tx, err = pinterf.UpdateNodeStatus(args.orgId, args.url, big.NewInt(int64(args.status)))

	case AssignOrgAdminAccount:
		// check if caller is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}

		// check if account is already in use in another org
		tx, err = pinterf.AssignOrgAdminAccount(args.orgId, args.acctId)

	case ApproveOrgAdminAccount:
		// check if caller is network admin
		if !s.isNetworkAdmin(args.txa.From) {
			return ErrNotNetworkAdmin
		}

		// check if anything is pending approval
		tx, err = pinterf.ApproveOrgAdminAccount(args.acctId)

	case AddNewRole:
		// check if org admin
		if !s.isOrgAdmin(args.txa.From, args.orgId) {
			return ErrNotOrgAdmin
		}
		// validate

		// check if role is already there in the org
		tx, err = pinterf.AddNewRole(args.roleId, args.orgId, big.NewInt(int64(args.accessType)), args.isVoter)

	case RemoveRole:
		// check if org admin
		if !s.isOrgAdmin(args.txa.From, args.orgId) {
			return ErrNotOrgAdmin
		}

		tx, err = pinterf.RemoveRole(args.roleId, args.orgId)

	case AssignAccountRole:
		// check if org admin
		if !s.isOrgAdmin(args.txa.From, args.orgId) {
			return ErrNotOrgAdmin
		}

		// check if account is used in another org
		tx, err = pinterf.AssignAccountRole(args.acctId, args.orgId, args.roleId)
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
