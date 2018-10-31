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
	"github.com/ethereum/go-ethereum/controls/permbind"
	)

type PermissionAPI struct {
	txPool      *core.TxPool
	ethClnt     *ethclient.Client
	permContr   *permbind.Permissions
	transOpts   *bind.TransactOpts
}

func NewPermissionAPI(e *core.TxPool) *PermissionAPI {
	pa := &PermissionAPI{e, nil, nil, nil}
	return pa
}

func (p *PermissionAPI) Init(ethClnt *ethclient.Client, datadir string) error {
	p.ethClnt = ethClnt
	key, kerr := getKeyFromKeyStore(datadir)
	if kerr != nil {
		log.Error("error reading key file", "err", kerr)
		return kerr
	}

	permContr, err := permbind.NewPermissions(params.QuorumPermissionsContract, p.ethClnt)
	if err != nil {
		return err
	}
	p.permContr = permContr
	auth, err := bind.NewTransactor(strings.NewReader(key), "")
	if err != nil {
		return err
	}
	p.transOpts = auth

	return nil
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

	nonce := s.txPool.Nonce(s.transOpts.From)
	s.transOpts.Nonce = new(big.Int).SetUint64(nonce)

	permissionsSession := &permbind.PermissionsSession{
		Contract: s.permContr,
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
