package quorum

import (
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"path/filepath"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"strings"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/log"
	pbind "github.com/ethereum/go-ethereum/controls/bind"
	"github.com/ethereum/go-ethereum/common"
)

var defaultGasLimit = uint64(47000)
var defaultGasPrice = big.NewInt(0)

type PermissionAPI struct {
	txPool    *core.TxPool
	ethClnt   *ethclient.Client
	transOpts *bind.TransactOpts
	permContr *pbind.Permissions
	clustContr *pbind.Cluster
}

func NewPermissionAPI(tp *core.TxPool) *PermissionAPI {
	pa := &PermissionAPI{tp, nil, nil, nil, nil}
	return pa
}

func (p *PermissionAPI) Init(ethClnt *ethclient.Client, datadir string) error {
	p.ethClnt = ethClnt
	key, kerr := getKeyFromKeyStore(datadir)
	if kerr != nil {
		log.Error("error reading key file", "err", kerr)
		return kerr
	}
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
	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	if err != nil {
		return err
	}
	p.transOpts = auth
	return nil
}


func (s *PermissionAPI) AddVoter(addr common.Address) bool {
	ps := s.newPermSession()
	tx, err := ps.AddVoter(addr)
	if err != nil {
		log.Warn("Failed to add voter", "err", err)
		return false
	}
	txHash := tx.Hash()
	log.Info("Transaction pending", "tx hash", string(txHash[:]))
	return true
}

func (s *PermissionAPI) RemoveVoter(addr common.Address) bool {
	ps := s.newPermSession()
	tx, err := ps.RemoveVoter(addr)
	if err != nil {
		log.Warn("Failed to remove voter", "err", err)
		return false
	}
	txHash := tx.Hash()
	log.Info("Transaction pending", "tx hash", string(txHash[:]))
	return true
}


func (s *PermissionAPI) ProposeNode(nodeId string) bool {
	node, err := discover.ParseNode(nodeId)
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
	nonce := s.txPool.Nonce(s.transOpts.From)
	s.transOpts.Nonce = new(big.Int).SetUint64(nonce)

	ps := s.newPermSession()

	tx, err := ps.ProposeNode(enodeID, ipAddrPort, discPort, raftPort)
	if err != nil {
		log.Warn("Failed to propose node", "err", err)
		log.Error("Failed to propose node: %v", err)
		return false
	}
	txHash := tx.Hash()
	statusMsg := fmt.Sprintf("Transaction pending tx hash %s", string(txHash[:]))
	log.Debug(statusMsg)
	return true
}

func (s *PermissionAPI) ApproveNode(nodeId string) bool {
	node, err := discover.ParseNode(nodeId)
	if err != nil {
		log.Error("invalid node id: %v", err)
		return false
	}
	enodeID := node.ID.String()
	nonce := s.txPool.Nonce(s.transOpts.From)
	s.transOpts.Nonce = new(big.Int).SetUint64(nonce)

	ps := s.newPermSession()
	tx, err := ps.ApproveNode(enodeID)
	if err != nil {
		log.Warn("Failed to propose node", "err", err)
		return false
	}
	txHash := tx.Hash()
	log.Debug("Transaction pending", "tx hash", string(txHash[:]))
	return true
}

func (s *PermissionAPI) DeactivateNode(nodeId string) bool {
	node, err := discover.ParseNode(nodeId)
	if err != nil {
		log.Error("invalid node id: %v", err)
		return false
	}
	enodeID := node.ID.String()
	nonce := s.txPool.Nonce(s.transOpts.From)
	s.transOpts.Nonce = new(big.Int).SetUint64(nonce)

	ps := s.newPermSession()
	tx, err := ps.DeactivateNode(enodeID)
	if err != nil {
		log.Warn("Failed to propose node", "err", err)
		return false
	}
	txHash := tx.Hash()
	log.Debug("Transaction pending", "tx hash", string(txHash[:]))
	return true
}

func (s *PermissionAPI) ApproveDeactivateNode(nodeId string) bool {
	node, err := discover.ParseNode(nodeId)
	if err != nil {
		log.Error("invalid node id: %v", err)
		return false
	}
	enodeID := node.ID.String()
	nonce := s.txPool.Nonce(s.transOpts.From)
	s.transOpts.Nonce = new(big.Int).SetUint64(nonce)

	ps := s.newPermSession()
	//TODO change it to approveDeactivate node once contract is updated
	tx, err := ps.DeactivateNode(enodeID)
	if err != nil {
		log.Warn("Failed to propose node", "err", err)
		return false
	}
	txHash := tx.Hash()
	log.Debug("Transaction pending", "tx hash", string(txHash[:]))
	return true
}


func (s *PermissionAPI) newPermSession() *pbind.PermissionsSession {
	return &pbind.PermissionsSession{
		Contract: s.permContr,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     s.transOpts.From,
			Signer:   s.transOpts.Signer,
			GasLimit: defaultGasLimit,
			GasPrice: defaultGasPrice,
		},
	}
}

func (s *PermissionAPI) newClusterSession() *pbind.ClusterSession {
	return &pbind.ClusterSession{
		Contract: s.clustContr,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     s.transOpts.From,
			Signer:   s.transOpts.Signer,
			GasLimit: defaultGasLimit,
			GasPrice: defaultGasPrice,
		},
	}
}

func (s *PermissionAPI) AddOrgKey(orgId string, pvtKey string) bool {
	cs := s.newClusterSession()
	tx, err := cs.AddOrgKey(orgId, pvtKey)
	if err != nil {
		log.Warn("Failed to add org key", "err", err)
		return false
	}
	txHash := tx.Hash()
	log.Info("Transaction pending", "tx hash", string(txHash[:]))
	return true
}

func (s *PermissionAPI) RemoveOrgKey(orgId string, pvtKey string) bool {
	cs := s.newClusterSession()
	tx, err := cs.DeleteOrgKey(orgId, pvtKey)
	if err != nil {
		log.Warn("Failed to remove org key", "err", err)
		return false
	}
	txHash := tx.Hash()
	log.Info("Transaction pending", "tx hash", string(txHash[:]))
	return true
}



func getKeyFromKeyStore(datadir string) (string, error) {

	files, err := ioutil.ReadDir(filepath.Join(datadir, "keystore"))
	if err != nil {
		return "", err
	}

	// HACK: here we always use the first key as transactor
	var keyPath string
	for _, f := range files {
		keyPath = filepath.Join(datadir, "keystore", f.Name())
		break
	}
	keyBlob, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return "", err
	}
	n := len(keyBlob)

	return string(keyBlob[:n]), nil
}
