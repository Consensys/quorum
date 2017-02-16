package quorum

import (
	"time"

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
	publicState, privateState *state.StateDB
	tcount                    int // tx count in cycle
	gp                        *core.GasPool
	ownedAccounts             *set.Set
	txs                       types.Transactions // set of transactions
	lowGasTxs                 types.Transactions
	failedTxs                 types.Transactions
	parent                    *types.Block

	header   *types.Header
	receipts types.Receipts
	logs     vm.Logs

	createdAt time.Time
}

func (ps *pendingState) applyTransaction(tx *types.Transaction, bc *core.BlockChain, cc *core.ChainConfig) (error, vm.Logs) {
	publicSnaphot, privateSnapshot := ps.publicState.Snapshot(), ps.privateState.Snapshot()

	// this is a bit of a hack to force jit for the miners
	config := cc.VmConfig
	if !(config.EnableJit && config.ForceJit) {
		config.EnableJit = false
	}
	config.ForceJit = false // disable forcing jit

	publicReceipt, _, _, err := core.ApplyTransaction(cc, bc, ps.gp, ps.publicState, ps.privateState, ps.header, tx, ps.header.GasUsed, config)
	if err != nil {
		ps.publicState.RevertToSnapshot(publicSnaphot)
		ps.privateState.RevertToSnapshot(privateSnapshot)

		return err, nil
	}
	ps.txs = append(ps.txs, tx)
	ps.receipts = append(ps.receipts, publicReceipt)

	return nil, publicReceipt.Logs
}

func (ps *pendingState) applyTransactions(txs *types.TransactionsByPriorityAndNonce, mux *event.TypeMux, bc *core.BlockChain, cc *core.ChainConfig) (types.Transactions, types.Transactions) {
	var (
		lowGasTxs types.Transactions
		failedTxs types.Transactions
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

		// Start executing the transaction
		ps.publicState.StartRecord(tx.Hash(), common.Hash{}, 0)

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
