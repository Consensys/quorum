package gethRaft

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
	"github.com/ethereum/go-ethereum/rpc"
)

type PublicTransactionPoolAPI struct {
	// eventMux   *event.TypeMux
	chainDb    ethdb.Database
	blockchain *core.BlockChain
	am         *accounts.Manager
	txPool     *core.TxPool
	txMu       *sync.Mutex
}

// NewPublicTransactionPoolAPI creates a new RPC service with methods specific for the transaction pool.
func NewPublicTransactionPoolAPI(s *RaftService) *PublicTransactionPoolAPI {
	api := &PublicTransactionPoolAPI{
		// eventMux:      e.eventMux,
		chainDb:    s.chainDb,
		blockchain: s.blockchain,
		am:         s.accountManager,
		txPool:     s.txPool,
		txMu:       &s.txMu,
	}

	return api
}

// SendTransaction creates a transaction for the given argument, sign it and submit it to the
// transaction pool.
func (s *PublicTransactionPoolAPI) SendTransaction(args SendTxArgs) (common.Hash, error) {
	gas, gasPrice := mockGasAndPrice()

	s.txMu.Lock()
	defer s.txMu.Unlock()

	if args.Nonce == nil {
		args.Nonce = rpc.NewHexNumber(s.txPool.State().GetNonce(args.From))
	}

	var tx *types.Transaction
	if args.To == nil {
		tx = types.NewContractCreation(args.Nonce.Uint64(), args.Value.BigInt(), gas, gasPrice, common.FromHex(args.Data))
	} else {
		tx = types.NewTransaction(args.Nonce.Uint64(), *args.To, args.Value.BigInt(), gas, gasPrice, common.FromHex(args.Data))
	}

	signature, err := s.am.Sign(args.From, tx.SigHash().Bytes())
	if err != nil {
		return common.Hash{}, err
	}

	return submitTransaction(s.txPool, tx, signature)
}

type SendTxArgs struct {
	From  common.Address  `json:"from"`
	To    *common.Address `json:"to"`
	Value *rpc.HexNumber  `json:"value"`
	Data  string          `json:"data"`
	Nonce *rpc.HexNumber  `json:"nonce"`
}

// *** public transaction api

func mockGasAndPrice() (*big.Int, *big.Int) {
	// pretend to send 90000 gas, even though it's not actually consumed
	// TODO(joel): why not more?
	return rpc.NewHexNumber(90000).BigInt(), rpc.NewHexNumber(0).BigInt()
}

// submitTransaction is a helper function that submits tx to txPool and creates a log entry.
func submitTransaction(txPool *core.TxPool, tx *types.Transaction, signature []byte) (common.Hash, error) {
	signedTx, err := tx.WithSignature(signature)
	if err != nil {
		return common.Hash{}, err
	}

	txPool.SetLocal(signedTx)
	if err := txPool.Add(signedTx); err != nil {
		return common.Hash{}, err
	}

	if signedTx.To() == nil {
		from, _ := signedTx.From()
		addr := crypto.CreateAddress(from, signedTx.Nonce())
		glog.V(logger.Info).Infof("Tx(%s) created: %s\n", signedTx.Hash().Hex(), addr.Hex())
	} else {
		glog.V(logger.Info).Infof("Tx(%s) to: %s\n", signedTx.Hash().Hex(), tx.To().Hex())
	}

	return signedTx.Hash(), nil
}
