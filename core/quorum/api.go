package quorum

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	pbind "github.com/ethereum/go-ethereum/controls/bind"
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
)

// OrgKeyAction represents an action in cluster contract
type OrgKeyAction int

const (
	AddOrgKey OrgKeyAction = iota
	RemoveOrgKey
)

// PermissionAPI provides an API to access Quorum's node permission and org key management related services
type PermissionAPI struct {
	txPool     *core.TxPool
	ethClnt    *ethclient.Client
	acntMgr    *accounts.Manager
	txOpt      *bind.TransactOpts
	permContr  *pbind.Permissions
	clustContr *pbind.Cluster
	key        *ecdsa.PrivateKey
}

// txArgs holds arguments required for execute functions
type txArgs struct {
	voter      common.Address
	nodeId     string
	orgId      string
	keyId      string
	txa        ethapi.SendTxArgs
	acctId     common.Address
	accessType string
}

type nodeStatus struct {
	Name   string
	Status string
}

type ExecStatus struct {
	Status bool
	Msg    string
}

var nodeApproveStatus = map[uint8]string{
	0: "Unknown",
	1: "PendingApproval",
	2: "Approved",
	3: "PendingDeactivation",
	4: "Deactivated",
	5: "PendingActivation",
	6: "PendingBlacklisting",
	7: "Blacklisted",
}

// NewPermissionAPI creates a new PermissionAPI to access quorum services
func NewPermissionAPI(tp *core.TxPool, am *accounts.Manager) *PermissionAPI {
	return &PermissionAPI{tp, nil, am, nil, nil, nil, nil}
}

// helper function decodes the node status to string
func decodeNodeStatus(nodeStatus uint8) string {
	if status, ok := nodeApproveStatus[nodeStatus]; ok {
		return status
	}
	return "Unknown"
}

//Init initializes PermissionAPI with eth client, permission contract and org key management control
func (p *PermissionAPI) Init(ethClnt *ethclient.Client, key *ecdsa.PrivateKey) error {
	p.ethClnt = ethClnt
	permContr, err := pbind.NewPermissions(params.QuorumPermissionsContract, p.ethClnt)
	if err != nil {
		return err
	}
	p.permContr = permContr
	clustContr, err := pbind.NewCluster(params.QuorumPrivateKeyManagementContract, p.ethClnt)
	if err != nil {
		return err
	}
	p.clustContr = clustContr
	p.key = key
	return nil
}

// Returns the list of Nodes and status of each
func (s *PermissionAPI) PermissionNodeList() []nodeStatus {
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
	// get the total number of nodes on the contract
	nodeCnt, err := ps.GetNumberOfNodes()
	nodeCntI := nodeCnt.Int64()
	nodeStatArr := make([]nodeStatus, nodeCntI)
	// loop for each index and get the node details from the contract
	if err == nil {
		i := int64(0)
		for i < nodeCntI {
			nodeDtls, _ := ps.GetNodeDetails(big.NewInt(i))
			nodeStatArr[i].Name = "enode://" + nodeDtls.EnodeId + "@" + nodeDtls.IpAddrPort
			nodeStatArr[i].Name += "?discport=" + nodeDtls.DiscPort
			if len(nodeDtls.RaftPort) > 0 {
				nodeStatArr[i].Name += "&raftport" + nodeDtls.RaftPort
			}
			nodeStatArr[i].Status = decodeNodeStatus(nodeDtls.NodeStatus)
			i++
		}
	}
	return nodeStatArr
}

// AddVoter adds an account to the list of accounts that can approve nodes proposed or deactivated
func (s *PermissionAPI) AddVoter(vaddr common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(AddVoter, txArgs{voter: vaddr, txa: txa})
}

// RemoveVoter removes an account from the list of accounts that can approve nodes proposed or deactivated
func (s *PermissionAPI) RemoveVoter(vaddr common.Address, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(RemoveVoter, txArgs{voter: vaddr, txa: txa})
}

// ProposeNode proposes a node to join the network
func (s *PermissionAPI) ProposeNode(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ProposeNode, txArgs{nodeId: nodeId, txa: txa})
}

// ApproveNode approves a proposed node to join the network
func (s *PermissionAPI) ApproveNode(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveNode, txArgs{nodeId: nodeId, txa: txa})
}

// DeactivateNode requests a node to get deactivated
func (s *PermissionAPI) ProposeNodeDeactivation(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ProposeNodeDeactivation, txArgs{nodeId: nodeId, txa: txa})
}

// ApproveDeactivateNode approves a node to get deactivated
func (s *PermissionAPI) ApproveNodeDeactivation(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveNodeDeactivation, txArgs{nodeId: nodeId, txa: txa})
}

// DeactivateNode requests a node to get deactivated
func (s *PermissionAPI) ProposeNodeActivation(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ProposeNodeActivation, txArgs{nodeId: nodeId, txa: txa})
}

// ApproveDeactivateNode approves a node to get deactivated
func (s *PermissionAPI) ApproveNodeActivation(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveNodeActivation, txArgs{nodeId: nodeId, txa: txa})
}

// DeactivateNode requests a node to get deactivated
func (s *PermissionAPI) ProposeNodeBlacklisting(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ProposeNodeBlacklisting, txArgs{nodeId: nodeId, txa: txa})
}

// ApproveDeactivateNode approves a node to get deactivated
func (s *PermissionAPI) ApproveNodeBlacklisting(nodeId string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(ApproveNodeBlacklisting, txArgs{nodeId: nodeId, txa: txa})
}

// RemoveOrgKey removes an org key combination from the org key map
func (s *PermissionAPI) RemoveOrgKey(orgId string, pvtKey string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(RemoveOrgKey, txArgs{txa: txa, orgId: orgId, keyId: pvtKey})
}

// AddOrgKey adds an org key combination to the org key map
func (s *PermissionAPI) AddOrgKey(orgId string, pvtKey string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executeOrgKeyAction(AddOrgKey, txArgs{txa: txa, orgId: orgId, keyId: pvtKey})
}

func (s *PermissionAPI) SetAccountAccess(acct common.Address, access string, txa ethapi.SendTxArgs) ExecStatus {
	return s.executePermAction(SetAccountAccess, txArgs{acctId: acct, accessType: access, txa: txa})
}

// executePermAction helps to execute an action in permission contract
func (s *PermissionAPI) executePermAction(action PermAction, args txArgs) ExecStatus {
	var err error
	var w accounts.Wallet
	voterErr := ExecStatus{false, "add voter first"}
	invalidNodeErr := ExecStatus{false, "invalid node id"}

	w, err = s.validateAccount(args.txa.From)
	if err != nil {
		return ExecStatus{false, err.Error()}
	}
	ps := s.newPermSession(w, args.txa)
	var tx *types.Transaction
	var node *discover.Node

	switch action {
	case AddVoter:
		tx, err = ps.AddVoter(args.voter)

	case RemoveVoter:
		tx, err = ps.RemoveVoter(args.voter)

	case ProposeNode:
		if !checkVoterExists(ps) {
			return voterErr
		}
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return invalidNodeErr
		}
		enodeID := node.ID.String()
		ipAddr := node.IP.String()
		port := fmt.Sprintf("%v", node.TCP)
		discPort := fmt.Sprintf("%v", node.UDP)
		raftPort := fmt.Sprintf("%v", node.RaftPort)
		ipAddrPort := ipAddr + ":" + port

		tx, err = ps.ProposeNode(enodeID, ipAddrPort, discPort, raftPort)

	case ApproveNode:
		if !checkIsVoter(ps, args.txa.From) {
			return voterErr
		}
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return invalidNodeErr
		}
		enodeID := node.ID.String()
		tx, err = ps.ApproveNode(enodeID)

	case ProposeNodeDeactivation:
		if !checkVoterExists(ps) {
			return voterErr
		}
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return invalidNodeErr
		}
		enodeID := node.ID.String()
		tx, err = ps.ProposeDeactivation(enodeID)

	case ApproveNodeDeactivation:
		if !checkIsVoter(ps, args.txa.From) {
			return voterErr
		}
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return invalidNodeErr
		}
		enodeID := node.ID.String()
		tx, err = ps.DeactivateNode(enodeID)

	case ProposeNodeActivation:
		if !checkVoterExists(ps) {
			return voterErr
		}
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return invalidNodeErr
		}
		enodeID := node.ID.String()
		tx, err = ps.ProposeNodeActivation(enodeID)

	case ApproveNodeActivation:
		if !checkIsVoter(ps, args.txa.From) {
			return voterErr
		}
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return invalidNodeErr
		}
		enodeID := node.ID.String()
		tx, err = ps.ActivateNode(enodeID)

	case ProposeNodeBlacklisting:
		if !checkVoterExists(ps) {
			return voterErr
		}
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return invalidNodeErr
		}
		enodeID := node.ID.String()
		ipAddr := node.IP.String()
		port := fmt.Sprintf("%v", node.TCP)
		discPort := fmt.Sprintf("%v", node.UDP)
		raftPort := fmt.Sprintf("%v", node.RaftPort)
		ipAddrPort := ipAddr + ":" + port

		tx, err = ps.ProposeNodeBlacklisting(enodeID, ipAddrPort, discPort, raftPort)
	case ApproveNodeBlacklisting:
		if !checkIsVoter(ps, args.txa.From) {
			return voterErr
		}
		node, err = discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return invalidNodeErr
		}
		enodeID := node.ID.String()
		tx, err = ps.BlacklistNode(enodeID)

	case SetAccountAccess:
		var access uint64
		access, err = strconv.ParseUint(args.accessType, 10, 8)
		if err != nil {
			return ExecStatus{false, "invalid access type"}
		}
		tx, err = ps.UpdateAccountAccess(args.acctId, uint8(access))
	}

	if err != nil {
		log.Error("Failed to execute permission action", "action", action, "err", err)
		return ExecStatus{false, err.Error()}
	}
	log.Debug("executed permission action", "action", action, "tx", tx)
	return ExecStatus{true, ""}
}

// executeOrgKeyAction helps to execute an action in cluster contract
func (s *PermissionAPI) executeOrgKeyAction(action OrgKeyAction, args txArgs) ExecStatus {
	w, err := s.validateAccount(args.txa.From)
	if err != nil {
		return ExecStatus{false, err.Error()}
	}
	ps := s.newClusterSession(w, args.txa)
	var tx *types.Transaction

	switch action {
	case AddOrgKey:
		tx, err = ps.AddOrgKey(args.orgId, args.keyId)
	case RemoveOrgKey:
		tx, err = ps.DeleteOrgKey(args.orgId, args.keyId)
	}
	if err != nil {
		log.Error("Failed to execute orgKey action", "action", action, "err", err)
		return ExecStatus{false, err.Error()}
	}
	log.Debug("executed orgKey action", "action", action, "tx", tx)
	return ExecStatus{true, ""}
}

// validateAccount validates the account and returns the wallet associated with that for signing the transaction
func (s *PermissionAPI) validateAccount(from common.Address) (accounts.Wallet, error) {
	acct := accounts.Account{Address: from}
	w, err := s.acntMgr.Find(acct)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// checkVoterExists checks if any vote accounts are there. If yes returns true, else false
func checkVoterExists(ps *pbind.PermissionsSession) bool {
	tx, err := ps.GetNumberOfVoters()
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
func (s *PermissionAPI) newPermSession(w accounts.Wallet, txa ethapi.SendTxArgs) *pbind.PermissionsSession {
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
func (s *PermissionAPI) newClusterSession(w accounts.Wallet, txa ethapi.SendTxArgs) *pbind.ClusterSession {
	frmAcct, transactOpts, gasLimit, gasPrice, nonce := s.getTxParams(txa, w)
	cs := &pbind.ClusterSession{
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
func (s *PermissionAPI) getTxParams(txa ethapi.SendTxArgs, w accounts.Wallet) (accounts.Account, *bind.TransactOpts, uint64, *big.Int, *big.Int) {
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
