package vault

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/hashicorp/vault/api"
	"math/big"
	"os"
)

type vaultWallet struct {
	url accounts.URL
	vault vaultService
	updateFeed *event.Feed
	accounts []accounts.Account
}

// vault related behaviour that will be specific to each vault type
type vaultService interface {
	status() (string, error)
	open() error
	close() error
}

func newHashicorpWallet(config hashicorpWalletConfig, updateFeed *event.Feed) (vaultWallet, error) {
	var url accounts.URL

	//to parse a string url as an accounts.URL it must first be in json format
	toParse := fmt.Sprintf("\"%v\"", config.Client.Url)

	if err := url.UnmarshalJSON([]byte(toParse)); err != nil {
		return vaultWallet{}, err
	}

	return vaultWallet{url: url, vault: &hashicorpService{config: config.Client}, updateFeed: updateFeed}, nil
}

func (w vaultWallet) URL() accounts.URL {
	return w.url
}

// the vault service should return open and nil error if status is good
func (w vaultWallet) Status() (string, error) {
	return w.vault.status()
}

func (w vaultWallet) Open(passphrase string) error {
	if err := w.vault.open(); err != nil {
		return err
	}

	w.updateFeed.Send(accounts.WalletEvent{Wallet: w, Kind: accounts.WalletOpened})

	return nil
}

func (w vaultWallet) Close() error {
	return w.vault.close()
}

func (w vaultWallet) Accounts() []accounts.Account {
	cpy := make([]accounts.Account, len(w.accounts))
	copy(cpy, w.accounts)

	return cpy
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
	config hashicorpClientConfig
}

const (
	open = "open"
	closed = "closed"
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

func (h *hashicorpService) status() (string, error) {
	if h.client == nil {
		return closed, nil
	}

	health, err := h.client.Sys().Health()

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

const (
	roleIDEnv = "VAULT_ROLE_ID"
	secretIDEnv = "VAULT_SECRET_ID"
)

var (
	noHashicorpEnvSetErr = fmt.Errorf("environment variables must be set when creating the Hashicorp client.  Set %v and %v if the Vault is configured to use Approle authentication.  Else set %v", roleIDEnv, secretIDEnv, api.EnvVaultToken)
	invalidApproleAuthErr = fmt.Errorf("both %v and %v must be set if using Approle authentication", roleIDEnv, secretIDEnv)
)

func (h *hashicorpService) open() error {
	if h.client != nil {
		return accounts.ErrWalletAlreadyOpen
	}

	conf := api.DefaultConfig()
	conf.Address = h.config.Url

	tlsConfig := &api.TLSConfig{
		CACert: h.config.CaCert,
		ClientCert: h.config.ClientCert,
		ClientKey: h.config.ClientKey,
	}

	if err := conf.ConfigureTLS(tlsConfig); err != nil {
		return fmt.Errorf("error creating Hashicorp client: %v", err)
	}

	c, err := api.NewClient(conf)

	if err != nil {
		return fmt.Errorf("error creating Hashicorp client: %v", err)
	}

	roleID := os.Getenv(roleIDEnv)
	secretID := os.Getenv(secretIDEnv)

	if roleID == "" && secretID == "" && os.Getenv(api.EnvVaultToken) == "" {
		return noHashicorpEnvSetErr
	}

	if roleID == "" && secretID != "" || roleID != "" && secretID == "" {
		return invalidApproleAuthErr
	}

	if usingApproleAuth(roleID, secretID) {
		//authenticate the client using approle
		body := map[string]interface{} {"role_id": roleID, "secret_id": secretID}

		approle := h.config.Approle

		if approle == "" {
			approle = "approle"
		}

		resp, err := c.Logical().Write(fmt.Sprintf("auth/%s/login", approle), body)

		if err != nil {
			return err
		}

		token, err := resp.TokenID()

		c.SetToken(token)
	}

	// api.Client uses the token at VAULT_TOKEN by default so nothing extra needs to be done when not using approle
	h.client = c

	return nil
}

func usingApproleAuth(roleID, secretID string) bool {
	return roleID != "" && secretID != ""
}

func (h *hashicorpService) close() error {
	h.client = nil

	return nil
}
