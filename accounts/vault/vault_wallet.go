package vault

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type vaultWallet struct {
	Url accounts.URL
}

func (w vaultWallet) URL() accounts.URL {
	return w.Url
}

func (w vaultWallet) Status() (string, error) {
	panic("implement me")
}

func (w vaultWallet) Open(passphrase string) error {
	panic("implement me")
}

func (w vaultWallet) Close() error {
	panic("implement me")
}

func (w vaultWallet) Accounts() []accounts.Account {
	panic("implement me")
}

func (w vaultWallet) Contains(account accounts.Account) bool {
	panic("implement me")
}

func (w vaultWallet) Derive(path accounts.DerivationPath, pin bool) (accounts.Account, error) {
	panic("implement me")
}

func (w vaultWallet) SelfDerive(base accounts.DerivationPath, chain ethereum.ChainStateReader) {
	panic("implement me")
}

func (w vaultWallet) SignHash(account accounts.Account, hash []byte) ([]byte, error) {
	panic("implement me")
}

func (w vaultWallet) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int, isQuorum bool) (*types.Transaction, error) {
	panic("implement me")
}

func (w vaultWallet) SignHashWithPassphrase(account accounts.Account, passphrase string, hash []byte) ([]byte, error) {
	panic("implement me")
}

func (w vaultWallet) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	panic("implement me")
}

