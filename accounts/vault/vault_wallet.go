package vault

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hashicorp/vault/api"
	"math/big"
)

type vaultWallet struct {
	url accounts.URL
	vault vaultService
}

// vault related behaviour that will be specific to each vault type
type vaultService interface {
	status() (string, error)
}

func newHashicorpWallet(config hashicorpWalletConfig) (vaultWallet, error) {
	var url accounts.URL

	//to parse a string url as an accounts.URL it must first be in json format
	toParse := fmt.Sprintf("\"%v\"", config.Client.Url)

	if err := url.UnmarshalJSON([]byte(toParse)); err != nil {
		return vaultWallet{}, err
	}

	return vaultWallet{url: url, vault: hashicorpService{}}, nil
}

func (w vaultWallet) URL() accounts.URL {
	return w.url
}

const (
	open = "open"
	closed = "closed"
)

// the vault service should return open and nil error if status is good
func (w vaultWallet) Status() (string, error) {
	if w.vault == nil {
		return closed, nil
	}

	return w.vault.status()
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

type hashicorpService struct {
	client *api.Client
}

const (
	hashicorpHealthcheckFailed = "Hashicorp Vault healthcheck failed"
	hashicorpUninitialized = "Hashicorp Vault uninitialized"
	hashicorpSealed = "Hashicorp Vault sealed"
)

var (
	hashicorpSealedErr = errors.New(hashicorpSealed)
	hashicorpUninitializedErr = errors.New(hashicorpUninitialized)
)

type hashicorpHealthcheckErr struct {
	err error
}

func (e hashicorpHealthcheckErr) Error() string {
	return fmt.Sprintf("%v: %v", hashicorpHealthcheckFailed, e.err)
}

func (c hashicorpService) status() (string, error) {
	health, err := c.client.Sys().Health()

	if err != nil {
		return hashicorpHealthcheckFailed, hashicorpHealthcheckErr{err: err}
	}

	if !health.Initialized {
		return hashicorpUninitialized, hashicorpUninitializedErr
	}

	if health.Sealed {
		return hashicorpSealed, hashicorpSealedErr
	}

	return open, nil
}
