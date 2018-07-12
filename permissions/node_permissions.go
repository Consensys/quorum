package permissions

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"gopkg.in/urfave/cli.v1"
)
//populates the nodes list from permissioned-nodes.json into the permissions
//smart contract
func PopulateNodes(ctx *cli.Context, stack *node.Node ){

	  log.Trace("Quorum permissioning v2 started")

	  var e *eth.Ethereum
	  if err := stack.Service(&e); err != nil {
		  utils.Fatalf("Ethereum service not running: %v", err)
	  }

	  rpcClient, err := stack.Attach()
	  if err != nil {
		  utils.Fatalf("Failed to attach to self: %v", err)
	  }
	  stateReader := ethclient.NewClient(rpcClient)
	  log.Trace("rpc connection to permissions contract established")

	  datadir := ctx.GlobalString(utils.DataDirFlag.Name)

	  files, err := ioutil.ReadDir(filepath.Join(datadir, "keystore"))
	  if err != nil {
		  utils.Fatalf("Failed to read keystore directory: %v", err)
	  }
	  log.Trace("reading account keys...")

	  // (zekun) HACK: here we always use the first key as transactor
	  var keyPath string
	  for _, f := range files {
		  keyPath = filepath.Join(datadir, "keystore", f.Name())
		  break
	  }
	  keyBlob, err := ioutil.ReadFile(keyPath)
	  if err != nil {
		  utils.Fatalf("Failed to read key file: %v", err)
	  }
	  log.Debug("Finished reading key file", "keyPath", keyPath, "keyBlob", keyBlob)
	  // n := bytes.IndexByte(keyBlob, 0)
	  n := len(keyBlob)
	  log.Debug("Decoding keyBlob", "length", n)
	  key := string(keyBlob[:n])
	  log.Debug("Decoded key", "key", key)

	  contractAddr := common.HexToAddress("0x0000000000000000000000000000000000000020") // hard coded in genesis
	  permissionsContract, err := NewPermissions(contractAddr, stateReader)
	  if err != nil {
		  utils.Fatalf("Failed to instantiate a Permissions contract: %v", err)
	  }
	  log.Debug("Permissions contract instantiated")
	  auth, err := bind.NewTransactor(strings.NewReader(key), "")
	  if err != nil {
		  utils.Fatalf("Failed to create authorized transactor: %v", err)
	  }
	  log.Debug("Transactor created")
	  permissionsSession := &PermissionsSession{
		  Contract: permissionsContract,
		  CallOpts: bind.CallOpts{
			  Pending: true,
		  },
		  TransactOpts: bind.TransactOpts{
			  From:     auth.From,
			  Signer:   auth.Signer,
			  GasLimit: 3558096384,
			  GasPrice: big.NewInt(0),
		  },
	  }

	  nodes := p2p.ParsePermissionedNodes(datadir)
	  for _, node := range nodes {
		  enodeID := fmt.Sprintf("%x", node.ID[:])
		  log.Trace("Adding node to permissions contract", "enodeID", enodeID)

		  nonce := e.TxPool().Nonce(permissionsSession.TransactOpts.From)
		  permissionsSession.TransactOpts.Nonce = new(big.Int).SetUint64(nonce)
		  log.Trace("Current Nonce", "nonce", nonce)

		  tx, err := permissionsSession.ProposeNode(enodeID, true, true)
		  if err != nil {
			  log.Warn("Failed to propose node", "err", err)
		  }
		  log.Debug("Transaction pending", "tx hash", tx.Hash())
	  }
}
