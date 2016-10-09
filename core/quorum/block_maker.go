package quorum

import (
	"time"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
	"gopkg.in/fatih/set.v0"
)

// BlockMaker defines the interface that block makers must provide.
type BlockMaker interface {
	// Pending returns the pending block and pending state.
	Pending() (*types.Block, *state.StateDB)
}

type pendingState struct {
	state         *state.StateDB // apply state changes here
	tcount        int            // tx count in cycle
	gp            *core.GasPool
	ownedAccounts *set.Set
	txs           types.Transactions // set of transactions
	lowGasTxs     types.Transactions
	failedTxs     types.Transactions
	parent        *types.Block

	header   *types.Header
	receipts types.Receipts
	logs     vm.Logs

	createdAt time.Time
}

func (ps *pendingState) applyTransaction(tx *types.Transaction, bc *core.BlockChain, cc *core.ChainConfig) (error, vm.Logs) {
	snap := ps.state.Snapshot()

	// this is a bit of a hack to force jit for the miners
	config := cc.VmConfig
	if !(config.EnableJit && config.ForceJit) {
		config.EnableJit = false
	}
	config.ForceJit = false // disable forcing jit

	receipt, logs, _, err := core.ApplyTransaction(cc, bc, ps.gp, ps.state, ps.header, tx, ps.header.GasUsed, config)
	if err != nil {
		ps.state.RevertToSnapshot(snap)
		return err, nil
	}
	ps.txs = append(ps.txs, tx)
	ps.receipts = append(ps.receipts, receipt)

	return nil, logs
}

func (ps *pendingState) applyTransactions(txs *types.TransactionsByPriceAndNonce, mux *event.TypeMux, bc *core.BlockChain, cc *core.ChainConfig) (types.Transactions, types.Transactions) {
	var (
		lowGasTxs types.Transactions
		failedTxs types.Transactions
		gasPrice  = new(big.Int).Mul(big.NewInt(20), common.Shannon)
	)

	var coalescedLogs vm.Logs
	for {
		// Retrieve the next transaction and abort if all done
		tx := txs.Peek()
		if tx == nil {
			break
		}
		// Error may be ignored here. The error has already been checked
		// during transaction acceptance is the transaction pool.
		from, _ := tx.From()

		// Ignore any transactions (and accounts subsequently) with low gas limits
		if tx.GasPrice().Cmp(gasPrice) < 0 && !ps.ownedAccounts.Has(from) {
			// Pop the current low-priced transaction without shifting in the next from the account
			glog.V(logger.Info).Infof("Transaction (%x) below gas price (tx=%v ask=%v). All sequential txs from this address(%x) will be ignored\n", tx.Hash().Bytes()[:4], common.CurrencyToString(tx.GasPrice()), common.CurrencyToString(gasPrice), from[:4])
			lowGasTxs = append(lowGasTxs, tx)
			txs.Pop()
			continue
		}
		// Start executing the transaction
		ps.state.StartRecord(tx.Hash(), common.Hash{}, 0)

		err, logs := ps.applyTransaction(tx, bc, cc)
		switch {
		case core.IsGasLimitErr(err):
			// Pop the current out-of-gas transaction without shifting in the next from the account
			glog.V(logger.Detail).Infof("Gas limit reached for (%x) in this block. Continue to try smaller txs\n", from[:4])
			txs.Pop()
		case err != nil:
			// Pop the current failed transaction without shifting in the next from the account
			glog.V(logger.Detail).Infof("Transaction (%x) failed, will be removed: %v\n", tx.Hash().Bytes()[:4], err)
			failedTxs = append(failedTxs, tx)
			txs.Pop()
		default:
			// Everything ok, collect the logs and shift in the next transaction from the same account
			coalescedLogs = append(coalescedLogs, logs...)
			ps.tcount++
			txs.Shift()
		}
	}
	if len(coalescedLogs) > 0 || ps.tcount > 0 {
		go func(logs vm.Logs, tcount int) {
			if len(logs) > 0 {
				mux.Post(core.PendingLogsEvent{Logs: logs})
			}
			if tcount > 0 {
				mux.Post(core.PendingStateEvent{})
			}
		}(coalescedLogs, ps.tcount)
	}

	return lowGasTxs, failedTxs
}
