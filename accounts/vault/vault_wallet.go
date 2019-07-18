package vault

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/hashicorp/vault/api"
	"math/big"
	"os"
	"strconv"
	"time"
)

type vaultWallet struct {
	url accounts.URL
	vault vaultService
	updateFeed *event.Feed
}

// vault related behaviour that will be specific to each vault type
type vaultService interface {
	status() (string, error)
	open() error
	close() error
	accounts() []accounts.Account
}

func newHashicorpWallet(config hashicorpWalletConfig, updateFeed *event.Feed) (vaultWallet, error) {
	var url accounts.URL

	//to parse a string url as an accounts.URL it must first be in json format
	toParse := fmt.Sprintf("\"%v\"", config.Client.Url)

	if err := url.UnmarshalJSON([]byte(toParse)); err != nil {
		return vaultWallet{}, err
	}

	return vaultWallet{url: url, vault: &hashicorpService{config: config.Client, secrets: config.Secrets}, updateFeed: updateFeed}, nil
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
	return w.vault.accounts()
}

func (w vaultWallet) Contains(account accounts.Account) bool {
	equal := func(a, b accounts.Account) bool {
		return a.Address == b.Address && (a.URL == b.URL || a.URL == accounts.URL{} || b.URL == accounts.URL{})
	}

	accts := w.Accounts()

	for _, a := range accts {
		if equal(a, account) {
			return true
		}
	}
	return false
}

func (w vaultWallet) Derive(path accounts.DerivationPath, pin bool) (accounts.Account, error) {
	return accounts.Account{}, accounts.ErrNotSupported
}

func (w vaultWallet) SelfDerive(base accounts.DerivationPath, chain ethereum.ChainStateReader) {}

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
	secrets []hashicorpSecretData
	accts []accounts.Account
	keys []*ecdsa.PrivateKey
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

	// 10s polling interval by default
	pollingIntervalSecs := h.config.VaultPollingIntervalSecs
	if pollingIntervalSecs == 0 {
		pollingIntervalSecs = 10
	}
	d := time.Duration(pollingIntervalSecs) * time.Second

	go h.accountRetrievalLoop(time.NewTicker(d))

	if h.config.StorePrivateKeys {
		go h.privateKeyRetrievalLoop(time.NewTicker(d))
	}

	return nil
}

func (h *hashicorpService) accountRetrievalLoop(ticker *time.Ticker) {
	for range ticker.C {
		if len(h.accts) == len(h.secrets) {
			ticker.Stop()
			return
		}

		for _, s := range h.secrets {
			path := fmt.Sprintf("%s/data/%s", s.SecretEngine, s.AddressSecret)

			url := fmt.Sprintf("%v/v1/%v?version=%v", h.client.Address(), path, s.AddressSecretVersion)

			versionData := make(map[string][]string)
			versionData["version"] = []string{strconv.Itoa(s.AddressSecretVersion)}

			// get address from vault
			resp, err := h.client.Logical().ReadWithData(path, versionData)

			if err != nil {
				log.Warn("unable to get secret from Hashicorp Vault", "url", url, "err", err)
				continue
			}

			respData, ok := resp.Data["data"].(map[string]interface{})

			if !ok {
				log.Warn("Hashicorp Vault response does not contain data", "url", url)
				continue
			}

			if len(respData) != 1 {
				log.Warn("only one key/value pair is allowed in each Hashicorp Vault secret", "url", url)
				continue
			}

			// get secret regardless of key in map
			var addr interface{}
			for _, d := range respData {
				addr = d
			}

			address, ok := addr.(string)

			if !ok {
				log.Warn("Hashicorp Vault response data is not in string format", "url", url)
				continue
			}

			// create accounts.Account
			//to parse a string url as an accounts.URL it must first be in json format
			toParse := fmt.Sprintf("\"%v\"", url)
			var parsedUrl accounts.URL

			if err := parsedUrl.UnmarshalJSON([]byte(toParse)); err != nil {
				log.Warn("unable to parse url of account retrieved from Hashicorp Vault", "url", url, "err", err)
				continue
			}

			acct := accounts.Account{
				Address: common.HexToAddress(address),
				URL: parsedUrl,
			}

			h.accts = append(h.accts, acct)
		}
	}
}

func (h *hashicorpService) privateKeyRetrievalLoop(ticker *time.Ticker) {
	for range ticker.C {
		if len(h.keys) == len(h.secrets) {
			ticker.Stop()
			return
		}

		for _, s := range h.secrets {
			path := fmt.Sprintf("%s/data/%s", s.SecretEngine, s.PrivateKeySecret)

			url := fmt.Sprintf("%v/v1/%v?version=%v", h.client.Address(), path, s.PrivateKeySecretVersion)

			versionData := make(map[string][]string)
			versionData["version"] = []string{strconv.Itoa(s.PrivateKeySecretVersion)}

			// get key from vault
			resp, err := h.client.Logical().ReadWithData(path, versionData)

			if err != nil {
				log.Warn("unable to get secret from Hashicorp Vault", "url", url, "err", err)
				continue
			}

			respData, ok := resp.Data["data"].(map[string]interface{})

			if !ok {
				log.Warn("Hashicorp Vault response does not contain data", "url", url)
				continue
			}

			if len(respData) != 1 {
				log.Warn("only one key/value pair is allowed in each Hashicorp Vault secret", "url", url)
				continue
			}

			// get secret regardless of key in map
			var k interface{}
			for _, d := range respData {
				k = d
			}

			hex, ok := k.(string)

			if !ok {
				log.Warn("Hashicorp Vault response data is not in string format", "url", url)
				continue
			}

			// create *ecdsa.PrivateKey
			key, err := crypto.HexToECDSA(hex)

			if err != nil {
				log.Warn("unable to parse data from Hashicorp Vault to *ecdsa.PrivateKey", "url", url, "err", err)
				continue
			}

			h.keys = append(h.keys, key)
		}
	}
}

func usingApproleAuth(roleID, secretID string) bool {
	return roleID != "" && secretID != ""
}

func (h *hashicorpService) close() error {
	h.client = nil

	return nil
}

func (h *hashicorpService) accounts() []accounts.Account {
	cpy := make([]accounts.Account, len(h.accts))
	copy(cpy, h.accts)

	return cpy
}
