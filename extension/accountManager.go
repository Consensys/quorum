package extension

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/internal/ethapi"
)

var (
	//default gas limit to use if not passed in sendTxArgs
	defaultGasLimit = uint64(4712384)
	//default gas price to use if not passed in sendTxArgs
	defaultGasPrice = big.NewInt(0)

	//Private participants must be specified for contract extension related transactions
	errNotPrivate = errors.New("must specify private participants")
)

// IAccountManager is an interface for transaction and account generation
// based operations
type IAccountManager interface {
	// Exists returns whether a given address is managed by this account manager or not
	Exists(address common.Address) bool

	// GenerateTransactOptions transforms API input arguments to ethclient
	// compatible arguments
	// It will validate based on use in the Privacy Extension context, namely
	// it will check the account is managed by the this node, and that it is a
	// private transaction
	GenerateTransactOptions(apiArguments ethapi.SendTxArgs) (*bind.TransactOpts, error)
}

type AccountManager struct {
	manager *accounts.Manager
}

func NewAccountManager(manager *accounts.Manager) *AccountManager {
	return &AccountManager{manager: manager}
}

func (manager *AccountManager) Exists(address common.Address) bool {
	_, err := manager.manager.Find(accounts.Account{Address: address})
	return err == nil
}

func (manager *AccountManager) GenerateTransactOptions(txa ethapi.SendTxArgs) (*bind.TransactOpts, error) {
	if txa.PrivateFor == nil {
		return nil, errNotPrivate
	}
	if !manager.Exists(txa.From) {
		return nil, fmt.Errorf("no wallet found for account %s", txa.From.String())
	}

	//Find the account we plan to send the transaction from
	from := accounts.Account{Address: txa.From}
	wallet, _ := manager.manager.Find(from)

	txArgs := bind.NewWalletTransactor(wallet, from)
	txArgs.PrivateFrom = txa.PrivateFrom
	txArgs.PrivateFor = txa.PrivateFor
	txArgs.GasLimit = defaultGasLimit
	txArgs.GasPrice = defaultGasPrice

	if txa.GasPrice != nil {
		txArgs.GasPrice = txa.GasPrice.ToInt()
	}
	if txa.Gas != nil {
		txArgs.GasLimit = uint64(*txa.Gas)
	}
	return txArgs, nil
}
