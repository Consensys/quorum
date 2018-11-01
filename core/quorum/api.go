package quorum

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/params"
	pbind "github.com/ethereum/go-ethereum/controls/bind"
	"github.com/ethereum/go-ethereum/log"
)

var defaultGasLimit = uint64(470000000)
var defaultGasPrice = big.NewInt(0)

type PermAction int

const (
	ProposeNode PermAction = iota
	ApproveNode
	DeactivateNode
	ApproveDeactivateNode
	AddVoter
	RemoveVoter
)

type OrgKeyAction int

const (
	AddOrgKey OrgKeyAction = iota
	RemoveOrgKey
)

type PermissionAPI struct {
	txPool     *core.TxPool
	ethClnt    *ethclient.Client
	acntMgr    *accounts.Manager
	txOpt      *bind.TransactOpts
	permContr  *pbind.Permissions
	clustContr *pbind.Cluster
}

type txArgs struct {
	from   common.Address
	voter  common.Address
	nodeId string
	orgId  string
	keyId  string
}

func NewPermissionAPI(tp *core.TxPool, am *accounts.Manager) *PermissionAPI {
	return &PermissionAPI{tp, nil, am, nil, nil, nil}
}

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

func (s *PermissionAPI) AddVoter(from common.Address, vaddr common.Address) bool {
	return s.executePermAction(AddVoter, txArgs{voter: vaddr, from: from})
}

func (s *PermissionAPI) RemoveVoter(from common.Address, vaddr common.Address) bool {
	return s.executePermAction(RemoveVoter, txArgs{voter: vaddr, from: from})
}

func (s *PermissionAPI) ProposeNode(from common.Address, nodeId string) bool {
	return s.executePermAction(ProposeNode, txArgs{nodeId: nodeId, from: from})
}

func (s *PermissionAPI) ApproveNode(from common.Address, nodeId string) bool {
	return s.executePermAction(ApproveNode, txArgs{nodeId: nodeId, from: from})
}

func (s *PermissionAPI) DeactivateNode(from common.Address, nodeId string) bool {
	return s.executePermAction(DeactivateNode, txArgs{nodeId: nodeId, from: from})
}

func (s *PermissionAPI) ApproveDeactivateNode(from common.Address, nodeId string) bool {
	return s.executePermAction(ApproveDeactivateNode, txArgs{nodeId: nodeId, from: from})
}

func (s *PermissionAPI) RemoveOrgKey(from common.Address, orgId string, pvtKey string) bool {
	return s.executeOrgKeyAction(RemoveOrgKey, txArgs{from: from, orgId: orgId, keyId: pvtKey})
}

func (s *PermissionAPI) AddOrgKey(from common.Address, orgId string, pvtKey string) bool {
	return s.executeOrgKeyAction(AddOrgKey, txArgs{from: from, orgId: orgId, keyId: pvtKey})
}

func (s *PermissionAPI) executePermAction(action PermAction, args txArgs) bool {
	fromAcct, w, err := s.validateAccount(args.from)
	if err != nil {
		return false
	}
	ps := s.newPermSession(w, fromAcct)
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

func (s *PermissionAPI) executeOrgKeyAction(action OrgKeyAction, args txArgs) bool {
	fromAcct, w, err := s.validateAccount(args.from)
	if err != nil {
		return false
	}
	ps := s.newClusterSession(w, fromAcct)
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

func (s *PermissionAPI) validateAccount(from common.Address) (accounts.Account, accounts.Wallet, error) {
	acct := accounts.Account{Address: from}
	w, err := s.acntMgr.Find(acct)
	if err != nil {
		return acct, nil, err
	}
	return acct, w, nil
}

func (s *PermissionAPI) newPermSession(w accounts.Wallet, acct accounts.Account) *pbind.PermissionsSession {
	transactOpts := bind.NewWalletTransactor(w, acct)
	ps := &pbind.PermissionsSession{
		Contract: s.permContr,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     acct.Address,
			GasLimit: defaultGasLimit,
			GasPrice: defaultGasPrice,
			Signer:   transactOpts.Signer,
		},
	}
	nonce := s.txPool.Nonce(acct.Address)
	ps.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)
	return ps
}

func (s *PermissionAPI) newClusterSession(w accounts.Wallet, acct accounts.Account) *pbind.ClusterSession {
	transactOpts := bind.NewWalletTransactor(w, acct)
	return &pbind.ClusterSession{
		Contract: s.clustContr,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     acct.Address,
			GasLimit: defaultGasLimit,
			GasPrice: defaultGasPrice,
			Signer:   transactOpts.Signer,
		},
	}
}
