package quorum

import (
	"crypto/ecdsa"
	"fmt"
	"errors"
	"math/big"
	"strings"

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
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/params"
)

//default gas limit to use if not passed in sendTxArgs
var defaultGasLimit = uint64(470000000)

//default gas price to use if not passed in sendTxArgs
var defaultGasPrice = big.NewInt(0)

// PermAction represents actions in permission contract
type PermAction int

const (
	ProposeNode PermAction = iota
	ApproveNode
	ProposeNodeDeactivation
	ApproveNodeDeactivation
	ProposeNodeActivation
	ApproveNodeActivation
	ProposeNodeBlacklisting
	ApproveNodeBlacklisting
	AddVoter
	RemoveVoter
	SetAccountAccess
	CancelPendingOperation
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
	DetailsMismatch
	NothingToApprove
	NothingToCancel
)

// Voter access type
type VoterAccessType uint8

const (
	Active VoterAccessType = iota
	Inactive
)
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
}

// txArgs holds arguments required for execute functions
type txArgs struct {
	voter      common.Address
	nodeId     string
	orgId      string
	morgId     string
	tmKey      string
	txa        ethapi.SendTxArgs
	acctId     common.Address
	accessType uint8
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
	ErrNothingToCancel		= ExecStatus{false, "Nothing to cancel"}
	ErrNodeProposed		    = ExecStatus{false, "Node already proposed for the action"}
	ErrAccountIsNotVoter    = ExecStatus{false, "Not a voter account"}
	ErrBlacklistedNode      = ExecStatus{false, "Blacklisted node. Operation not allowed"}
	ErrOpNotAllowed         = ExecStatus{false, "Operation not allowed"}
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
	return &QuorumControlsAPI{tp, nil, am, nil, nil, nil, nil, false, false}
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
func (p *QuorumControlsAPI) Init(ethClnt *ethclient.Client, key *ecdsa.PrivateKey, apiName string) error {
	p.ethClnt = ethClnt
	if apiName == "quorumNodeMgmt" || apiName == "quorumAcctMgmt" {
		permContr, err := pbind.NewPermissions(params.QuorumPermissionsContract, p.ethClnt)
		if err != nil {
			return err
		}
		p.permContr = permContr
		p.permEnabled = true
	} else {
		clustContr, err := obind.NewCluster(params.QuorumPrivateKeyManagementContract, p.ethClnt)
		if err != nil {
			return err
		}
		if clustContr == nil {
		}
		p.clustContr = clustContr
		p.orgEnabled = true
	}
	p.key = key
	return nil
}

// Returns the list of Nodes and status of each
func (s *QuorumControlsAPI) PermissionNodeList() []nodeStatus {
	if !s.permEnabled {
		nodeStatArr := make([]nodeStatus, 1)
		nodeStatArr[0].EnodeId = "Permisssions control not enabled for network"
		return nodeStatArr
	}
	ps := s.newPermSessionWithNodeKeySigner()
	// get the total number of nodes on the contract
	nodeCnt, err := ps.GetNumberOfNodes()
	if err != nil {
		return nil
	}
	nodeCntI := nodeCnt.Int64()
	nodeStatArr := make([]nodeStatus, nodeCntI)
	// loop for each index and get the node details from the contract
	i := int64(0)
	for i < nodeCntI {
		nodeDtls, err := ps.GetNodeDetailsFromIndex(big.NewInt(i))
		if err != nil {
			log.Error("error getting node details", "err", err)
		} else {
			nodeStatArr[i].EnodeId = "enode://" + nodeDtls.EnodeId + "@" + nodeDtls.IpAddrPort
			nodeStatArr[i].EnodeId += "?discport=" + nodeDtls.DiscPort
			if len(nodeDtls.RaftPort) > 0 {
				nodeStatArr[i].EnodeId += "&raftport=" + nodeDtls.RaftPort
			}
			nodeStatArr[i].Status = decodeNodeStatus(nodeDtls.NodeStatus)
		}
		i++
	}
	return nodeStatArr
}

func (s *QuorumControlsAPI) PermissionAccountList() []accountInfo {
	if !s.permEnabled {
		acctInfoArr := make([]accountInfo, 1)
		acctInfoArr[0].Address = "Account access control not enable for the network"
		return acctInfoArr
	}
	ps := s.newPermSessionWithNodeKeySigner()
	// get the total number of accounts with permissions
	acctCnt, err := ps.GetNumberOfAccounts()
	if err != nil {
		return nil
	}
	acctCntI := acctCnt.Int64()
	log.Debug("total permission accounts", "count", acctCntI)
	acctInfoArr := make([]accountInfo, acctCntI)
	// loop for each index and get the node details from the contract
	i := int64(0)
	for i < acctCntI {
		a, err := ps.GetAccountDetails(big.NewInt(i))
		if err != nil {
			log.Error("error getting account info", "err", err)
		} else {
			acctInfoArr[i].Address = a.Acct.String()
			acctInfoArr[i].Access = decodeAccountPermission(a.AcctAccess)
		}
		i++
	}
	return acctInfoArr
}

func (s *QuorumControlsAPI) VoterList() []string {
	if !s.permEnabled {
		voterArr := make([]string, 1)
		voterArr[0] = "Permissions control not enabled for the network"
		return voterArr
	}
	ps := s.newPermSessionWithNodeKeySigner()
	// get the total number of accounts with permissions
	voterCnt, err := ps.GetNumberOfVoters()
	if err != nil {
		return nil
	}
	voterCntI := voterCnt.Int64()
	log.Debug("total voters", "count", voterCntI)
	var voterArr []string
	// loop for each index and get the node details from the contract
	i := int64(0)
	for i < voterCntI {
		voter, err := ps.GetVoter(big.NewInt(i))
		if err != nil {
			log.Error("error getting voter info", "err", err)
		} else {
			if voter.VoterStatus == uint8(Active){
				voterArr = append(voterArr, voter.Addr.String())
			}
		}
		i++
	}
	return voterArr
}

func (s *QuorumControlsAPI) newPermSessionWithNodeKeySigner() *pbind.PermissionsSession {
	auth := bind.NewKeyedTransactor(s.key)
	ps := &pbind.PermissionsSession{
		Contract: s.permContr,
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
	return ps
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
func decodeAccountPermission(access uint8) string {
	if status, ok := accountPermMap[access]; ok {
		return status
	}
	return "Unknown"
}

// AddVoter adds an account to the list of accounts that can approve nodes proposed or deactivated
func (s *QuorumControlsAPI) AddVoter(vaddr common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(AddVoter, txArgs{voter: vaddr, txa: txa})
}

// RemoveVoter removes an account from the list of accounts that can approve nodes proposed or deactivated
func (s *QuorumControlsAPI) RemoveVoter(vaddr common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(RemoveVoter, txArgs{voter: vaddr, txa: txa})
}

// ProposeNode proposes a node to join the network
func (s *QuorumControlsAPI) ProposeNode(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ProposeNode, txArgs{nodeId: nodeId, txa: txa})
}

// ApproveNode approves a proposed node to join the network
func (s *QuorumControlsAPI) ApproveNode(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveNode, txArgs{nodeId: nodeId, txa: txa})
}

// DeactivateNode requests a node to get deactivated
func (s *QuorumControlsAPI) ProposeNodeDeactivation(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ProposeNodeDeactivation, txArgs{nodeId: nodeId, txa: txa})
}

// ApproveDeactivateNode approves a node to get deactivated
func (s *QuorumControlsAPI) ApproveNodeDeactivation(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveNodeDeactivation, txArgs{nodeId: nodeId, txa: txa})
}

// DeactivateNode requests a node to get deactivated
func (s *QuorumControlsAPI) ProposeNodeActivation(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ProposeNodeActivation, txArgs{nodeId: nodeId, txa: txa})
}

// ApproveDeactivateNode approves a node to get deactivated
func (s *QuorumControlsAPI) ApproveNodeActivation(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveNodeActivation, txArgs{nodeId: nodeId, txa: txa})
}

// DeactivateNode requests a node to get deactivated
func (s *QuorumControlsAPI) ProposeNodeBlacklisting(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ProposeNodeBlacklisting, txArgs{nodeId: nodeId, txa: txa})
}

// ApproveDeactivateNode approves a node to get deactivated
func (s *QuorumControlsAPI) ApproveNodeBlacklisting(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveNodeBlacklisting, txArgs{nodeId: nodeId, txa: txa})
}

// CancelPendingOperation cancels a pending operation
func (s *QuorumControlsAPI) CancelPendingOperation(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(CancelPendingOperation, txArgs{nodeId: nodeId, txa: txa})
}

// AddMasterOrg adds an new master organization to the contract
func (s *QuorumControlsAPI) AddMasterOrg(morgId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(AddMasterOrg, txArgs{txa: txa, morgId: morgId})
}

// RemoveOrgKey removes an org key combination from the org key map
func (s *QuorumControlsAPI) AddSubOrg(orgId string, morgId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(AddSubOrg, txArgs{txa: txa, orgId: orgId, morgId: morgId})
}

// AddOrgKey adds an org key combination to the org key map
func (s *QuorumControlsAPI) AddOrgVoter(morgId string, acctId common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(AddOrgVoter, txArgs{txa: txa, morgId: morgId, acctId: acctId})
}

// RemoveOrgKey removes an org key combination from the org key map
func (s *QuorumControlsAPI) RemoveOrgVoter(morgId string, acctId common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(RemoveOrgVoter, txArgs{txa: txa, morgId: morgId, acctId: acctId})
}

func (s *QuorumControlsAPI) AddOrgKey(orgId string, tmKey string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(AddOrgKey, txArgs{txa: txa, orgId: orgId, tmKey: tmKey})
}

// RemoveOrgKey removes an org key combination from the org key map
func (s *QuorumControlsAPI) RemoveOrgKey(orgId string, tmKey string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(RemoveOrgKey, txArgs{txa: txa, orgId: orgId, tmKey: tmKey})
}

// RemoveOrgKey removes an org key combination from the org key map
func (s *QuorumControlsAPI) ApprovePendingOp(orgId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(ApprovePendingOp, txArgs{txa: txa, orgId: orgId})
}

func (s *QuorumControlsAPI) SetAccountAccess(acct common.Address, access uint8, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(SetAccountAccess, txArgs{acctId: acct, accessType: access, txa: txa})
}


func getNodeDetailsFromEnode(nodeId string) (string, string, string, string, error){
	node, err := discover.ParseNode(nodeId)
	if err != nil {
		log.Error("invalid node id: %v", err)
		return "", "", "", "", err
	} 
	enodeID := node.ID.String()
	ipAddr := node.IP.String()
	port := fmt.Sprintf("%v", node.TCP)
	discPort := fmt.Sprintf("%v", node.UDP)
	raftPort := fmt.Sprintf("%v", node.RaftPort)
	ipAddrPort := ipAddr + ":" + port

	return enodeID, discPort, raftPort, ipAddrPort, nil
}

// checks if the input node details for approval is matching with details stored
// in contract
func checkNodeDetails(ps *pbind.PermissionsSession, nodeId string, action PermAction)  (error, ExecStatus){
	enodeID, discPort, raftPort, ipAddrPort, err := getNodeDetailsFromEnode(nodeId)

	cnode, err := ps.GetNodeDetails(enodeID)
	if err == nil {
		nodeStatus := decodeNodeStatus(cnode.NodeStatus)
		// if node status is Blacklisted no activities are allowed on the same.
		if nodeStatus == "Blacklisted" {
			return errors.New("Cannot propose blacklisted node"), ErrBlacklistedNode
		}

		if nodeStatus == "NotInNetwork" && (action == ProposeNodeDeactivation || action == ProposeNodeActivation){
			return errors.New("operation cannot be performed"), ErrOpNotAllowed
		}

		newNode := false;
		if nodeStatus == "NotInNetwork" && len(cnode.IpAddrPort) == 0{
			newNode = true
		}
		detailsMatch := false;
		if strings.Compare(ipAddrPort, cnode.IpAddrPort) == 0 && strings.Compare(discPort, cnode.DiscPort) == 0 && strings.Compare(raftPort, cnode.RaftPort) == 0 {
			detailsMatch = true
		}
		// if the node is not in network and is being proposed for blacklisting or as a new node
		// allow the operation. For anyother operation, the node will be in the network and all details
		// should match
		if action == ProposeNode || action == ProposeNodeBlacklisting {
			if !newNode && !detailsMatch {
				return errors.New("Details Mismtach"), ErrNodeDetailsMismatch
			}
		} else if !detailsMatch {
			return errors.New("Details Mismtach"), ErrNodeDetailsMismatch
		}

		// if propose action, check if node status allows the operation
		if ((action == ProposeNode && nodeStatus != "NotInNetwork") ||
			(action == ProposeNodeDeactivation && nodeStatus != "Approved") ||
			(action == ProposeNodeActivation && nodeStatus != "Deactivated")) {
			return errors.New("operation cannot be performed"), ErrOpNotAllowed
		}

		// if approval action, check if anything pendinga approval
		if (action == ApproveNode && nodeStatus != "PendingApproval") ||
			(action == ApproveNodeDeactivation && nodeStatus != "PendingDeactivation") ||
			(action == ApproveNodeActivation && nodeStatus != "PendingActivation") ||
			(action == ApproveNodeBlacklisting && nodeStatus != "PendingBlacklisting") {
			return errors.New("Nothing to approve"), ErrNothingToApprove
		}

		if (action == CancelPendingOperation && nodeStatus != "PendingApproval" &&
			nodeStatus != "PendingDeactivation" && nodeStatus != "PendingActivation" &&
			nodeStatus != "PendingBlacklisting") {
			return errors.New("Nothing to cancel"), ErrNothingToCancel
		}

		if (action == ProposeNode && nodeStatus == "PendingApproval") ||
			(action == ProposeNodeDeactivation && nodeStatus == "PendingDeactivation") ||
			(action == ProposeNodeActivation && nodeStatus == "PendingActivation") ||
			(action == ProposeNodeBlacklisting && nodeStatus == "PendingBlacklisting") {
			return errors.New("Node already proposed"), ErrNodeProposed
		}

	} 

	return nil, ExecSuccess
}

func(s *QuorumControlsAPI) validateOpDetails(ps *pbind.PermissionsSession, enodeID string, from common.Address, action PermAction) (error, ExecStatus) {

	// check if  the input node is fine
	err, execStatus := checkNodeDetails(ps, enodeID, action)
	if err != nil {
		return errors.New("Node details mismatch"), execStatus
	}

	// if action is propose type then check if voter nodes are there in the network
	if (action == ProposeNode || action == ProposeNodeDeactivation || action == ProposeNodeActivation || action == ProposeNodeBlacklisting){
		if 	!checkVoterExists(ps){
			return errors.New("No voter account"), ErrNoVoterAccount
		}
	}

	// if approval process, check if the account is a voter account
	if (action == ApproveNode || action == ApproveNodeDeactivation || action == ApproveNodeActivation || action == ApproveNodeBlacklisting || action == CancelPendingOperation){
		if !checkIsVoter(ps, from) {
			return errors.New("Not a voter account"), ErrAccountNotAVoter
		}
	}
	return nil, ExecSuccess
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
	ps := s.newPermSession(w, args.txa)
	var tx *types.Transaction
	var node *discover.Node
	var execStatus ExecStatus

	if action != SetAccountAccess {
		err, execStatus = s.validateOpDetails(ps, args.nodeId, args.txa.From, action)
		if err != nil {
			return execStatus
		}
	}

	switch action {
	case AddVoter:
		if !checkVoterAccountAccess(args.voter) {
			return ErrVoterAccountAccess
		}
		if checkIsVoter(ps, args.voter){
			return ErrVoterExists
		}
		tx, err = ps.AddVoter(args.voter)

	case RemoveVoter:
		if !checkVoterAccountAccess(args.voter) {
			return ErrVoterAccountAccess
		}
		if !checkIsVoter(ps, args.voter){
			return ErrAccountIsNotVoter
		}
		tx, err = ps.RemoveVoter(args.voter)

	case ProposeNode:
		enodeID, discPort, raftPort, ipAddrPort, locerr := getNodeDetailsFromEnode(args.nodeId)
		if locerr != nil {
			log.Error("invalid node id: %v", err)
			return ErrInvalidNode
		} 

		tx, err = ps.ProposeNode(enodeID, ipAddrPort, discPort, raftPort)

	case ApproveNode:
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return ErrInvalidNode
		}
		enodeID := node.ID.String()

		tx, err = ps.ApproveNode(enodeID)

	case ProposeNodeDeactivation:
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return ErrInvalidNode
		}
		enodeID := node.ID.String()
		tx, err = ps.ProposeDeactivation(enodeID)

	case ApproveNodeDeactivation:
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return ErrInvalidNode
		}
		enodeID := node.ID.String()

		tx, err = ps.DeactivateNode(enodeID)

	case ProposeNodeActivation:
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return ErrInvalidNode
		}
		enodeID := node.ID.String()
		tx, err = ps.ProposeNodeActivation(enodeID)

	case ApproveNodeActivation:
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return ErrInvalidNode
		}
		enodeID := node.ID.String()

		tx, err = ps.ActivateNode(enodeID)

	case ProposeNodeBlacklisting:
		enodeID, discPort, raftPort, ipAddrPort, locerr := getNodeDetailsFromEnode(args.nodeId)
		if locerr != nil {
			log.Error("invalid node id: %v", err)
			return ErrInvalidNode
		}
		tx, err = ps.ProposeNodeBlacklisting(enodeID, ipAddrPort, discPort, raftPort)
	case ApproveNodeBlacklisting:
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return ErrInvalidNode
		}
		enodeID := node.ID.String()

		tx, err = ps.BlacklistNode(enodeID)

	case SetAccountAccess:
		if (args.accessType > 3){
			return ErrInvalidAccountAccess
		}
		if !checkAccountAccess(args.txa.From, args.acctId, args.accessType) {
			return ErrAccountAccess
		}
		tx, err = ps.UpdateAccountAccess(args.acctId, args.accessType)

	case CancelPendingOperation:
		if !checkIsVoter(ps, args.txa.From) {
			return ErrAccountNotAVoter
		}
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return ErrInvalidNode
		}
		enodeID := node.ID.String()

		tx, err = ps.CancelPendingOperation(enodeID)

	}

	if err != nil {
		log.Error("Failed to execute permission action", "action", action, "err", err)
		return ErrFailedExecution
	}
	log.Debug("executed permission action", "action", action, "tx", tx)
	return ExecSuccess
}

func (s *QuorumControlsAPI) GetOrgVoterList(morgId string) []string {
	if !s.orgEnabled {
		voterArr := make([]string, 1)
		voterArr[0] = "Permissions control not enabled for the network"
		return voterArr
	}
	ps := s.newOrgKeySessionWithNodeKeySigner()
	// get the total number of accounts with permissions
	voterCnt, err := ps.GetNumberOfVoters(morgId)
	if err != nil {
		return nil
	}
	voterCntI := voterCnt.Int64()
	log.Debug("total voters", "count", voterCntI)
	voterArr := make([]string, voterCntI)
	// loop for each index and get the node details from the contract
	i := int64(0)
	for i < voterCntI {
		a, err := ps.GetVoter(morgId, big.NewInt(i))
		if err != nil {
			log.Error("error getting voter info", "err", err)
		} else {
			voterArr[i] = a.String()
		}
		i++
	}
	return voterArr
}

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
			keyArr := make([]string, keyCntI)
			// loop for each index and get the node details from the contract
			j := int64(0)
			for j < keyCntI {
				key, err := ps.GetOrgKey(orgId, big.NewInt(j))
				if err != nil {
					log.Error("error key info", "err", err)
				} else {
					keyArr[j] = key
				}
				j++
			}
			orgArr[i].SubOrgKeyList = keyArr
		}
		i++
	}
	return orgArr
}

// this function returns the approval pending action at sub org level
func (s *QuorumControlsAPI) GetPendingOpDetails(orgId string) PendingOpInfo {
	if !s.orgEnabled {
		return PendingOpInfo{"Org key management not enabled for the network", "None"}
	}
	ps := s.newOrgKeySessionWithNodeKeySigner()
	ret, _ := ps.CheckOrgExists(orgId)
	if ret {
		// get the total number of accounts with permissions
		pendingKey, pendingOp, err := ps.GetOrgPendingOp(orgId)
		if err == nil {
			pendOpInfo := PendingOpInfo{pendingKey, decodePendingOp(pendingOp)}
			return pendOpInfo
		} else {
			return PendingOpInfo{"Info not found", "None"}
		}
	} else {
		return PendingOpInfo{"Org not found", "None"}
	}
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
		ret, _ := ps.CheckMasterOrgExists(args.morgId)
		if !ret {
			return ErrInvalidMasterOrg
		}
		ret, _, _ = ps.CheckIfVoterExists(args.morgId, args.acctId)
		if ret {
			return ErrVoterExists
		}
		tx, err = ps.AddVoter(args.morgId, args.acctId)

	case RemoveOrgVoter:
		ret, _ := ps.CheckMasterOrgExists(args.morgId)
		if !ret {
			return ErrInvalidMasterOrg
		}
		ret, _, _ = ps.CheckIfVoterExists(args.morgId, args.acctId)
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
		ret, _, _ = ps.CheckIfKeyExists(args.orgId, args.tmKey)
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
		ret, _, _ = ps.CheckIfKeyExists(args.orgId, args.tmKey)
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

// checkVoterExists checks if any vote accounts are there. If yes returns true, else false
func checkVoterExists(ps *pbind.PermissionsSession) bool {
	tx, err := ps.GetNumberOfValidVoters()
	log.Debug("number of voters", "count", tx)
	if err == nil && tx.Cmp(big.NewInt(0)) > 0 {
		return true
	}
	return false
}

// checks if any accounts is a valid voter to approve the action
func checkIsVoter(ps *pbind.PermissionsSession, acctId common.Address) bool {
	tx, err := ps.IsVoter(acctId)
	if err == nil && tx {
		return true
	}
	return false
}

// newPermSession creates a new permission contract session
func (s *QuorumControlsAPI) newPermSession(w accounts.Wallet, txa ethapi.SendTxArgs) *pbind.PermissionsSession {
	frmAcct, transactOpts, gasLimit, gasPrice, nonce := s.getTxParams(txa, w)
	ps := &pbind.PermissionsSession{
		Contract: s.permContr,
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
func checkAccountAccess(from, targetAcct common.Address, accessType uint8) bool {
	fromAcctAccess := types.GetAcctAccess(from)
	targetAcctAccess := types.GetAcctAccess(targetAcct)

	retVal := false

	// Check if the from account is sufficient access rights to set the access
	if fromAcctAccess == types.FullAccess {
		retVal = true
	} else if fromAcctAccess == types.ContractDeploy && accessType != uint8(types.FullAccess) {
		retVal = true
	} else if fromAcctAccess == types.Transact && (accessType == uint8(types.Transact) || accessType == uint8(types.ReadOnly)) {
		retVal = true
	} 

	if retVal && fromAcctAccess != types.FullAccess {
		if ((fromAcctAccess == types.ContractDeploy && targetAcctAccess == types.FullAccess) ||
			(fromAcctAccess == types.Transact &&
			(targetAcctAccess == types.ContractDeploy || targetAcctAccess == types.FullAccess))){
			retVal = false
		}

	}
	return retVal
}

// checks if the account performing the operation has sufficient access privileges
func checkVoterAccountAccess(account common.Address) bool {
	acctAccess := types.GetAcctAccess(account)
	// if acctAccess == types.ReadOnly {
	// only accounts with full access will be allowed to manage voters
	if acctAccess != types.FullAccess {
		return false
	}
	return true
}
