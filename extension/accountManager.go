package extension

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"math/big"
)

var (
	//default gas limit to use if not passed in sendTxArgs
	defaultGasLimit = uint64(4712384)
	//default gas price to use if not passed in sendTxArgs
	defaultGasPrice = big.NewInt(0)

	//Private participants must be specified for contract extension related transactions
	errNotPrivate = errors.New("must specify private participants")
)

type AccountManager struct {
	manager *accounts.Manager
}

func NewAccountManager(manager *accounts.Manager) *AccountManager {
	return &AccountManager{manager: manager}
}

func (manager *AccountManager) Exists(address common.Address) bool {
	from := accounts.Account{Address: address}
	_, err := manager.manager.Find(from)
	return err == nil
}

func (manager *AccountManager) generateTransactOpts(txa ethapi.SendTxArgs) (*bind.TransactOpts, error) {
	if txa.PrivateFor == nil {
		return nil, errNotPrivate
	}
	if !manager.Exists(txa.From) {
		return nil, fmt.Errorf("no wallet found for account %s", txa.From.String())
	}

	//Find the account we plan to send the transaction from
	frmAcct := accounts.Account{Address: txa.From}
	wallet, _ := manager.manager.Find(frmAcct)

	txArgs := bind.NewWalletTransactor(wallet, frmAcct)
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