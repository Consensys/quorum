package quorum

import (
	"fmt"
	"math/big"

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
	DeactivateNode
	ApproveDeactivateNode
	AddVoter
	RemoveVoter
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
}

// txArgs holds arguments required for execute functions
type txArgs struct {
	voter  common.Address
	nodeId string
	orgId  string
	keyId  string
	txa    ethapi.SendTxArgs
}

// NewPermissionAPI creates a new PermissionAPI to access quorum services
func NewPermissionAPI(tp *core.TxPool, am *accounts.Manager) *PermissionAPI {
	return &PermissionAPI{tp, nil, am, nil, nil, nil}
}

//Init initializes PermissionAPI with eth client, permission contract and org key management control
func (p *PermissionAPI) Init(ethClnt *ethclient.Client) error {
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
	return nil
}

// AddVoter adds an account to the list of accounts that can approve nodes proposed or deactivated
func (s *PermissionAPI) AddVoter(vaddr common.Address, txa ethapi.SendTxArgs) bool {
	return s.executePermAction(AddVoter, txArgs{voter: vaddr, txa: txa})
}

// RemoveVoter removes an account from the list of accounts that can approve nodes proposed or deactivated
func (s *PermissionAPI) RemoveVoter(vaddr common.Address, txa ethapi.SendTxArgs) bool {
	return s.executePermAction(RemoveVoter, txArgs{voter: vaddr, txa: txa})
}

// ProposeNode proposes a node to join the network
func (s *PermissionAPI) ProposeNode(nodeId string, txa ethapi.SendTxArgs) bool {
	return s.executePermAction(ProposeNode, txArgs{nodeId: nodeId, txa: txa})
}

// ApproveNode approves a proposed node to join the network
func (s *PermissionAPI) ApproveNode(nodeId string, txa ethapi.SendTxArgs) bool {
	return s.executePermAction(ApproveNode, txArgs{nodeId: nodeId, txa: txa})
}

// DeactivateNode requests a node to get deactivated
func (s *PermissionAPI) DeactivateNode(nodeId string, txa ethapi.SendTxArgs) bool {
	return s.executePermAction(DeactivateNode, txArgs{nodeId: nodeId, txa: txa})
}

// ApproveDeactivateNode approves a node to get deactivated
func (s *PermissionAPI) ApproveDeactivateNode(nodeId string, txa ethapi.SendTxArgs) bool {
	return s.executePermAction(ApproveDeactivateNode, txArgs{nodeId: nodeId, txa: txa})
}

// RemoveOrgKey removes an org key combination from the org key map
func (s *PermissionAPI) RemoveOrgKey(orgId string, pvtKey string, txa ethapi.SendTxArgs) bool {
	return s.executeOrgKeyAction(RemoveOrgKey, txArgs{txa: txa, orgId: orgId, keyId: pvtKey})
}

// AddOrgKey adds an org key combination to the org key map
func (s *PermissionAPI) AddOrgKey(orgId string, pvtKey string, txa ethapi.SendTxArgs) bool {
	return s.executeOrgKeyAction(AddOrgKey, txArgs{txa: txa, orgId: orgId, keyId: pvtKey})
}

// executePermAction helps to execute an action in permission contract
func (s *PermissionAPI) executePermAction(action PermAction, args txArgs) bool {
	w, err := s.validateAccount(args.txa.From)
	if err != nil {
		return false
	}
	ps := s.newPermSession(w, args.txa)
	var tx *types.Transaction

	switch action {
	case AddVoter:
		tx, err = ps.AddVoter(args.voter)
	case RemoveVoter:
		tx, err = ps.RemoveVoter(args.voter)
	case ProposeNode:
		node, err := discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return false
		}
		enodeID := node.ID.String()
		ipAddr := node.IP.String()
		port := fmt.Sprintf("%v", node.TCP)
		discPort := fmt.Sprintf("%v", node.UDP)
		raftPort := fmt.Sprintf("%v", node.RaftPort)
		ipAddrPort := ipAddr + ":" + port

		tx, err = ps.ProposeNode(enodeID, ipAddrPort, discPort, raftPort)
	case ApproveNode:
		node, err := discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return false
		}
		enodeID := node.ID.String()
		tx, err = ps.ApproveNode(enodeID)
	case DeactivateNode:
		node, err := discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return false
		}
		enodeID := node.ID.String()
		tx, err = ps.DeactivateNode(enodeID)
	case ApproveDeactivateNode:
		node, err := discover.ParseNode(args.nodeId)
		if err != nil {
			log.Error("invalid node id: %v", err)
			return false
		}
		enodeID := node.ID.String()
		//TODO change to approve deactivate node
		tx, err = ps.DeactivateNode(enodeID)

	}
	if err != nil {
		log.Error("Failed to execute permission action", "action", action, "err", err)
		return false
	}
	log.Debug("executed permission action", "action", action, "tx", tx)
	return true
}

// executeOrgKeyAction helps to execute an action in cluster contract
func (s *PermissionAPI) executeOrgKeyAction(action OrgKeyAction, args txArgs) bool {
	w, err := s.validateAccount(args.txa.From)
	if err != nil {
		return false
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
		return false
	}
	log.Debug("executed orgKey action", "action", action, "tx", tx)
	return true
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
