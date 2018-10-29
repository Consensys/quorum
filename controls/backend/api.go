package backend

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"math/big"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/ethereum/go-ethereum/controls/permbind"
)


func APIs(ec *ethclient.Client, e *eth.Ethereum, datadir string) []rpc.API {
	return []rpc.API{
		{
			Namespace: "permnode",
			Version:   "1.0",
			Service:   NewPermissionAPI(ec, e, datadir),
			Public:    true,
		},
	}
}

type PermissionAPI struct {
	ethClient *ethclient.Client
	eth *eth.Ethereum
	permissionsContr *permbind.Permissions
	transOpts *bind.TransactOpts

}




func getKeyFromKeyStore(datadir string) string {

	files, err := ioutil.ReadDir(filepath.Join(datadir, "keystore"))
	if err != nil {
		utils.Fatalf("Failed to read keystore directory: %v", err)
	}

	// HACK: here we always use the first key as transactor
	var keyPath string
	for _, f := range files {
		keyPath = filepath.Join(datadir, "keystore", f.Name())
		break
	}
	keyBlob, err := ioutil.ReadFile(keyPath)
	if err != nil {
		utils.Fatalf("Failed to read key file: %v", err)
	}
	// n := bytes.IndexByte(keyBlob, 0)
	n := len(keyBlob)

	return string(keyBlob[:n])
}

func NewPermissionAPI(ec *ethclient.Client, e *eth.Ethereum, datadir string) *PermissionAPI {
	permissionsContract, err := permbind.NewPermissions(params.QuorumPermissionsContract, ec)
	if err != nil {
		utils.Fatalf("Failed to instantiate a Permissions contract: %v", err)
	}
	//TODO check if reading from keystore is correct approach
	key := getKeyFromKeyStore(datadir)
	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	if err != nil {
		utils.Fatalf("Failed to create authorized transactor: %v", err)
	}
	return &PermissionAPI{ec, e, permissionsContract, auth}
}

func (s *PermissionAPI) AddVoter(addr string) string {
	log.Info("AJ-called1")
	return "added voter " + addr
}

func (s *PermissionAPI) ProposeNode(enodeId string) string {
	node, err := discover.ParseNode(enodeId)
	if err != nil {
		return fmt.Sprintf("invalid node id: %v", err)
	}
	enodeID := node.ID.String()
	ipAddr := node.IP.String()
	port := fmt.Sprintf("%v", node.TCP)
	discPort := fmt.Sprintf("%v", node.UDP)
	raftPort := fmt.Sprintf("%v", node.RaftPort)
	ipAddrPort := ipAddr + ":" + port

	log.Trace("AJ-Adding node to permissions contract", "enodeID", enodeID)

	nonce := s.eth.TxPool().Nonce(s.transOpts.From)
	s.transOpts.Nonce = new(big.Int).SetUint64(nonce)

	permissionsSession := &permbind.PermissionsSession{
		Contract: s.permissionsContr,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     s.transOpts.From,
			Signer:   s.transOpts.Signer,
			GasLimit: 4700000,
			GasPrice: big.NewInt(0),
		},
	}

	tx, err := permissionsSession.ProposeNode(enodeID, ipAddrPort, discPort, raftPort)
	if err != nil {
		log.Warn("AJ-Failed to propose node", "err", err)
	}
	statusMsg := fmt.Sprintf("Transaction pending tx hash %s", tx.Hash())
	log.Debug(statusMsg)
	return statusMsg
}

func (s *PermissionAPI) BlacklistNode(enodeId string) string {
	log.Info("AJ-called3")
	return "blacklisted node " + enodeId
}

func (s *PermissionAPI) RemoveNode(enodeId string) string {
	log.Info("AJ-called4")
	return "removed node " + enodeId
}

func (s *PermissionAPI) ApproveNode(enodeId string) string {
	log.Info("AJ-called5")
	return "approved node " + enodeId
}

func (s *PermissionAPI) ValidNodes() []string {
	log.Info("AJ-called6")
	return []string{"n1", "n2"}
}
