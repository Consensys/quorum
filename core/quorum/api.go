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
	"github.com/ethereum/go-ethereum/params"
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

const (
	AddMasterOrg OrgKeyAction = iota
	AddSubOrg
	AddOrgVoter
	RemoveOrgVoter
	AddOrgKey
	RemoveOrgKey
	ApprovePendingOp
)

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
	permContr   *pbind.Permissions
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
	ErrNoVoterAccount       = ExecStatus{false, "No voter account registered. Add voter first"}
	ErrInvalidNode          = ExecStatus{false, "Invalid node id"}
	ErrAccountNotAVoter     = ExecStatus{false, "Account is not a voter. Action cannot be approved"}
	ErrInvalidAccount       = ExecStatus{false, "Invalid account id"}
	ErrInvalidAccountAccess = ExecStatus{false, "Invalid account access type"}
	ErrFailedExecution      = ExecStatus{false, "Failed to execute permission action"}
	ErrNodeDetailsMismatch  = ExecStatus{false, "Node details mismatch"}
	ErrPermissionDisabled   = ExecStatus{false, "Permissions control not enabled"}
	ErrOrgDisabled          = ExecStatus{false, "Org key management not enabled for the network"}
	ErrAccountAccess        = ExecStatus{false, "Account does not have sufficient access for operation"}
	ErrVoterAccountAccess   = ExecStatus{false, "Voter account does not have sufficient access"}
	ErrMasterOrgExists      = ExecStatus{false, "Master org already exists"}
	ErrInvalidMasterOrg     = ExecStatus{false, "Master org does not exist. Add master org first"}
	ErrInvalidOrg           = ExecStatus{false, "Org does not exist. Add org first"}
	ErrOrgExists            = ExecStatus{false, "Org already exists"}
	ErrVoterExists          = ExecStatus{false, "Voter account exists"}
	ErrPendingApprovals     = ExecStatus{false, "Pending approvals for the organization. Approve first"}
	ErrKeyExists            = ExecStatus{false, "Key exists for the organization"}
	ErrKeyInUse             = ExecStatus{false, "Key already in use in another master organization"}
	ErrKeyNotFound          = ExecStatus{false, "Key not found for the organization"}
	ErrNothingToApprove     = ExecStatus{false, "Nothing to approve"}
	ErrNothingToCancel      = ExecStatus{false, "Nothing to cancel"}
	ErrNodeProposed         = ExecStatus{false, "Node already proposed for the action"}
	ErrAccountIsNotVoter    = ExecStatus{false, "Not a voter account"}
	ErrBlacklistedNode      = ExecStatus{false, "Blacklisted node. Operation not allowed"}
	ErrOpNotAllowed         = ExecStatus{false, "Operation not allowed"}
	ErrLastFullAccessAcct   = ExecStatus{false, "Last account with full access. Operation not allowed"}
	ExecSuccess             = ExecStatus{true, "Action completed successfully"}
)

var (
	nodeApproveStatus = map[uint8]string{
		0: "NotInNetwork",
		1: "PendingApproval",
		2: "Approved",
		3: "PendingDeactivation",
		4: "Deactivated",
		5: "PendingActivation",
		6: "PendingBlacklisting",
		7: "Blacklisted",
	}

	accountPermMap = map[uint8]string{
		0: "ReadOnly",
		1: "Transact",
		2: "ContractDeploy",
		3: "FullAccess",
	}

	pendingOpMap = map[uint8]string{
		0: "None",
		1: "Add",
		2: "Remove",
	}
)

// NewQuorumControlsAPI creates a new QuorumControlsAPI to access quorum services
func NewQuorumControlsAPI(tp *core.TxPool, am *accounts.Manager) *QuorumControlsAPI {
	return &QuorumControlsAPI{tp, nil, am, nil, nil, nil, nil, false, false, nil, nil}
}

// helper function decodes the node status to string
func decodeNodeStatus(nodeStatus uint8) string {
	if status, ok := nodeApproveStatus[nodeStatus]; ok {
		return status
	}
	return "Unknown"
}

// helper function decodes the node status to string
func decodePendingOp(pendingOp uint8) string {
	if desc, ok := pendingOpMap[pendingOp]; ok {
		return desc
	}
	return "Unknown"
}

//Init initializes QuorumControlsAPI with eth client, permission contract and org key management control
func (p *QuorumControlsAPI) Init(ethClnt *ethclient.Client, key *ecdsa.PrivateKey, apiName string, pconfig *types.PermissionConfig, pc *pbind.PermInterface) error {
	p.ethClnt = ethClnt
	if apiName == "quorumPermission" {
		var contractAddress common.Address
		//TODO: to be updated with new contract API
		//if pconfig.IsEmpty() {
		contractAddress = params.QuorumPermissionsContract
		//} else {
		//	contractAddress = common.HexToAddress(pconfig.InterfAddress)
		//}
		permContr, err := pbind.NewPermissions(contractAddress, p.ethClnt)
		if err != nil {
			return err
		}
		p.permConfig = pconfig
		p.permContr = permContr
		p.permEnabled = true
	} else {
		clustContr, err := obind.NewCluster(params.QuorumPrivateKeyManagementContract, p.ethClnt)
		if err != nil {
			return err
		}
		p.clustContr = clustContr
		p.orgEnabled = true
	}
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

func (s *QuorumControlsAPI) newOrgKeySessionWithNodeKeySigner() *obind.ClusterSession {
	auth := bind.NewKeyedTransactor(s.key)
	cs := &obind.ClusterSession{
		Contract: s.clustContr,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 4700000,
			GasPrice: big.NewInt(0),
		},
	}
	return cs
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

// AddMasterOrg adds an new master organization to the contract
func (s *QuorumControlsAPI) AddMasterOrg(morgId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(AddMasterOrg, txArgs{txa: txa, morgId: morgId})
}

// AddSubOrg ass a sub org to the master org
func (s *QuorumControlsAPI) AddSubOrg(orgId string, morgId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(AddSubOrg, txArgs{txa: txa, orgId: orgId, morgId: morgId})
}

// AddOrgVoter adds voter account to a master org
func (s *QuorumControlsAPI) AddOrgVoter(morgId string, acctId common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(AddOrgVoter, txArgs{txa: txa, morgId: morgId, acctId: acctId})
}

// RemoveOrgVoter removes voter account to a master org
func (s *QuorumControlsAPI) RemoveOrgVoter(morgId string, acctId common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(RemoveOrgVoter, txArgs{txa: txa, morgId: morgId, acctId: acctId})
}

// AddOrgKey adds an org key to the org id
func (s *QuorumControlsAPI) AddOrgKey(orgId string, tmKey string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(AddOrgKey, txArgs{txa: txa, orgId: orgId, tmKey: tmKey})
}

// RemoveOrgKey removes an org key combination from the org key map
func (s *QuorumControlsAPI) RemoveOrgKey(orgId string, tmKey string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(RemoveOrgKey, txArgs{txa: txa, orgId: orgId, tmKey: tmKey})
}

// ApprovePendingOp approves any key add or delete activity
func (s *QuorumControlsAPI) ApprovePendingOp(orgId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(ApprovePendingOp, txArgs{txa: txa, orgId: orgId})
}

// executePermAction helps to execute an action in permission contract
func (s *QuorumControlsAPI) executePermAction(action PermAction, args txArgs) ExecStatus {

	if !s.permEnabled {
		return ErrPermissionDisabled
	}
	log.Info("AJ-exec perm action", "action", action, "txargs", args)
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
		tx, err = pinterf.AddOrg(args.orgId, args.url)

	case ApproveOrg:
		tx, err = pinterf.ApproveOrg(args.orgId, args.url)

	case UpdateOrgStatus:
		tx, err = pinterf.UpdateOrgStatus(args.orgId, big.NewInt(int64(args.status)))

	case ApproveOrgStatus:
		tx, err = pinterf.ApproveOrgStatus(args.orgId, big.NewInt(int64(args.status)))

	case AddNode:
		tx, err = pinterf.AddNode(args.orgId, args.url)

	case UpdateNodeStatus:
		tx, err = pinterf.UpdateNodeStatus(args.orgId, args.url, big.NewInt(int64(args.status)))

	case AssignOrgAdminAccount:
		tx, err = pinterf.AssignOrgAdminAccount(args.orgId, args.acctId)

	case ApproveOrgAdminAccount:
		tx, err = pinterf.ApproveOrgAdminAccount(args.acctId)

	case AddNewRole:
		tx, err = pinterf.AddNewRole(args.roleId, args.orgId, big.NewInt(int64(args.accessType)), args.isVoter)

	case RemoveRole:
		tx, err = pinterf.RemoveRole(args.roleId, args.orgId)

	case AssignAccountRole:
		tx, err = pinterf.AssignAccountRole(args.acctId, args.orgId, args.roleId)
	}

	if err != nil {
		log.Error("Failed to execute permission action", "action", action, "err", err)
		return ErrFailedExecution
	}
	log.Debug("executed permission action", "action", action, "tx", tx)
	return ExecSuccess
}

// returns the master org, org and linked key details
func (s *QuorumControlsAPI) OrgKeyInfo() []orgInfo {
	if !s.orgEnabled {
		orgInfoArr := make([]orgInfo, 1)
		orgInfoArr[0].MasterOrgId = "Org key management not enabled for the network"
		return orgInfoArr
	}
	ps := s.newOrgKeySessionWithNodeKeySigner()
	// get the total number of accounts with permissions
	orgCnt, err := ps.GetNumberOfOrgs()
	if err != nil {
		return nil
	}
	orgCntI := orgCnt.Int64()
	log.Debug("total orgs", "count", orgCntI)
	orgArr := make([]orgInfo, orgCntI)
	// loop for each index and get the node details from the contract
	i := int64(0)
	for i < orgCntI {
		orgId, morgId, err := ps.GetOrgInfo(big.NewInt(i))
		if err != nil {
			log.Error("error getting org info", "err", err)
		} else {
			orgArr[i].SubOrgId = orgId
			orgArr[i].MasterOrgId = morgId
			// get the list of keys for the organization
			keyCnt, err := ps.GetOrgKeyCount(orgId)
			if err != nil {
				return nil
			}
			keyCntI := keyCnt.Int64()
			log.Debug("total keys", "count", keyCntI)
			var keyArr []string
			// loop for each index and get the node details from the contract
			j := int64(0)
			for j < keyCntI {
				key, status, err := ps.GetOrgKey(orgId, big.NewInt(j))
				if err == nil && status {
					keyArr = append(keyArr, key)
				}
				j++
			}
			orgArr[i].SubOrgKeyList = keyArr
		}
		i++
	}
	return orgArr
}

// executeOrgKeyAction helps to execute an action in cluster contract
func (s *QuorumControlsAPI) executeOrgKeyAction(action OrgKeyAction, args txArgs) ExecStatus {
	if !s.orgEnabled {
		return ErrOrgDisabled
	}
	w, err := s.validateAccount(args.txa.From)
	if err != nil {
		return ExecStatus{false, err.Error()}
	}
	ps := s.newClusterSession(w, args.txa)

	var tx *types.Transaction

	switch action {
	case AddMasterOrg:
		// check if the master org exists. if yes throw error
		ret, _ := ps.CheckMasterOrgExists(args.morgId)
		if ret {
			return ErrMasterOrgExists
		}
		tx, err = ps.AddMasterOrg(args.morgId)

	case AddSubOrg:
		ret, _ := ps.CheckMasterOrgExists(args.morgId)
		if !ret {
			return ErrInvalidMasterOrg
		}
		ret, err = ps.CheckOrgExists(args.orgId)
		if ret {
			return ErrOrgExists
		}
		tx, err = ps.AddSubOrg(args.orgId, args.morgId)

	case AddOrgVoter:
		if locErr, execStatus := valAccountAccessVoter(args.txa.From, args.acctId); locErr != nil {
			return execStatus
		}
		ret, _ := ps.CheckMasterOrgExists(args.morgId)
		if !ret {
			return ErrInvalidMasterOrg
		}
		ret, _ = ps.CheckIfVoterExists(args.morgId, args.acctId)
		if ret {
			return ErrVoterExists
		}
		tx, err = ps.AddVoter(args.morgId, args.acctId)

	case RemoveOrgVoter:
		if locErr, execStatus := valAccountAccessVoter(args.txa.From, common.Address{}); locErr != nil {
			return execStatus
		}
		ret, _ := ps.CheckMasterOrgExists(args.morgId)
		if !ret {
			return ErrInvalidMasterOrg
		}
		ret, _ = ps.CheckIfVoterExists(args.morgId, args.acctId)
		if !ret {
			return ErrInvalidAccount
		}
		tx, err = ps.DeleteVoter(args.morgId, args.acctId)

	case AddOrgKey:
		ret, _ := ps.CheckOrgExists(args.orgId)
		if !ret {
			return ErrInvalidOrg
		}
		ret, _ = ps.CheckVotingAccountExists(args.orgId)
		if !ret {
			return ErrNoVoterAccount
		}
		ret, _ = ps.CheckOrgPendingOp(args.orgId)
		if ret {
			return ErrPendingApprovals
		}
		ret, _ = ps.CheckIfKeyExists(args.orgId, args.tmKey)
		if ret {
			return ErrKeyExists
		}
		ret, _ = ps.CheckKeyClash(args.orgId, args.tmKey)
		if ret {
			return ErrKeyInUse
		}
		tx, err = ps.AddOrgKey(args.orgId, args.tmKey)

	case RemoveOrgKey:
		ret, _ := ps.CheckOrgExists(args.orgId)
		if !ret {
			return ErrInvalidOrg
		}
		ret, _ = ps.CheckVotingAccountExists(args.orgId)
		if !ret {
			return ErrNoVoterAccount
		}
		ret, _ = ps.CheckOrgPendingOp(args.orgId)
		if ret {
			return ErrPendingApprovals
		}
		ret, _ = ps.CheckIfKeyExists(args.orgId, args.tmKey)
		if !ret {
			return ErrKeyNotFound
		}
		tx, err = ps.DeleteOrgKey(args.orgId, args.tmKey)

	case ApprovePendingOp:
		ret, _ := ps.CheckOrgExists(args.orgId)
		if !ret {
			return ErrInvalidOrg
		}
		ret, _ = ps.IsVoter(args.orgId, args.txa.From)
		if !ret {
			return ErrAccountNotAVoter
		}
		ret, _ = ps.CheckOrgPendingOp(args.orgId)
		if !ret {
			return ErrNothingToApprove
		}
		tx, err = ps.ApprovePendingOp(args.orgId)
	}

	if err != nil {
		log.Error("Failed to execute orgKey action", "action", action, "err", err)
		return ExecStatus{false, err.Error()}
	}
	log.Debug("executed orgKey action", "action", action, "tx", tx)
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

// newClusterSession creates a new cluster contract session
func (s *QuorumControlsAPI) newClusterSession(w accounts.Wallet, txa ethapi.SendTxArgs) *obind.ClusterSession {
	frmAcct, transactOpts, gasLimit, gasPrice, nonce := s.getTxParams(txa, w)
	cs := &obind.ClusterSession{
		Contract: s.clustContr,
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
	return cs
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
