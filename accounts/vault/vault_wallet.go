package vault

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/hashicorp/vault/api"
	"math/big"
	"os"
	"sort"
	"strconv"
	"sync"
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
	getKey(acct accounts.Account) (key *ecdsa.PrivateKey, zeroFn func(), err error)
}

func newHashicorpWallet(config HashicorpWalletConfig, updateFeed *event.Feed) (vaultWallet, error) {
	var url accounts.URL

	//to parse a string url as an accounts.URL it must first be in json format
	toParse := fmt.Sprintf("\"%v\"", config.Client.Url)

	if err := url.UnmarshalJSON([]byte(toParse)); err != nil {
		return vaultWallet{}, err
	}

	w := vaultWallet{
		url: url,
		vault: newHashicorpService(config),
		updateFeed: updateFeed,
	}

	return w, nil
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
	if !w.Contains(account) {
		return nil, accounts.ErrUnknownAccount
	}

	key, zero, err := w.vault.getKey(account)

	if err != nil {
		return nil, err
	}

	defer zero()

	return crypto.Sign(hash, key)
}

func (w vaultWallet) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	if !w.Contains(account) {
		return nil, accounts.ErrUnknownAccount
	}

	key, zero, err := w.vault.getKey(account)

	if err != nil {
		return nil, err
	}

	defer zero()

	// start quorum specific
	if tx.IsPrivate() {
		log.Info("Private transaction signing with QuorumPrivateTxSigner")
		return types.SignTx(tx, types.QuorumPrivateTxSigner{}, key)
	} // End quorum specific

	// Depending on the presence of the chain ID, sign with EIP155 or homestead
	if chainID != nil {
		return types.SignTx(tx, types.NewEIP155Signer(chainID), key)
	}
	return types.SignTx(tx, types.HomesteadSigner{}, key)
}

func (w vaultWallet) SignHashWithPassphrase(account accounts.Account, passphrase string, hash []byte) ([]byte, error) {
	return w.SignHash(account, hash)
}

func (w vaultWallet) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return w.SignTx(account, tx, chainID)
}

type hashicorpService struct {
	client      *api.Client
	config      hashicorpClientConfig
	secrets     []hashicorpSecretConfig
	mutex       sync.RWMutex
	accts       []accounts.Account
	keyHandlers map[common.Address]map[accounts.URL]hashicorpKeyHandler

}

func newHashicorpService(config HashicorpWalletConfig) *hashicorpService {
	s := &hashicorpService{
		config:      config.Client,
		secrets:     config.Secrets,
		keyHandlers: make(map[common.Address]map[accounts.URL]hashicorpKeyHandler),
	}

	return s
}

type hashicorpKeyHandler struct {
	secret hashicorpSecretConfig
	key *ecdsa.PrivateKey
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
	pollingIntervalMillis := h.config.VaultPollingIntervalMillis
	if pollingIntervalMillis == 0 {
		pollingIntervalMillis = 10000
	}
	d := time.Duration(pollingIntervalMillis) * time.Millisecond

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

			address, err := h.getAddressFromVault(s)

			if err != nil {
				log.Warn("unable to get address from Hashicorp Vault", "url", url, "err", err)
				continue
			}

			// create accounts.Account
			// to parse a string url as an accounts.URL it must first be in json format
			toParse := fmt.Sprintf("\"%v\"", url)
			var parsedUrl accounts.URL

			if err := parsedUrl.UnmarshalJSON([]byte(toParse)); err != nil {
				log.Warn("unable to parse url of account retrieved from Hashicorp Vault", "url", url, "err", err)
				continue
			}

			acct := accounts.Account{
				Address: address,
				URL: parsedUrl,
			}

			// update state
			h.mutex.Lock()

			if _, ok := h.keyHandlers[acct.Address]; !ok {
				h.keyHandlers[acct.Address] = make(map[accounts.URL]hashicorpKeyHandler)
			}

			keyHandlersByUrl := h.keyHandlers[acct.Address]

			if _, ok := keyHandlersByUrl[acct.URL]; ok {
				log.Warn("Hashicorp Vault key handler already exists.  Not updated.", "url", url)
				h.mutex.Unlock()
				continue
			}

			keyHandlersByUrl[acct.URL] = hashicorpKeyHandler{secret: s}
			h.accts = append(h.accts, acct)

			h.mutex.Unlock()
		}
	}
}

func (h *hashicorpService) getAddressFromVault(s hashicorpSecretConfig) (common.Address, error) {
	hexAddr, err := h.getSecretFromVault(s.AddressSecret, s.AddressSecretVersion, s.SecretEngine)

	if err != nil {
		return common.Address{}, err
	}

	return common.HexToAddress(hexAddr), nil
}

func countRetrievedKeys(keyHandlers map[common.Address]map[accounts.URL]hashicorpKeyHandler) int {
	var n int

	for _, khByUrl := range keyHandlers {
		for _, kh := range khByUrl {
			if kh.key != nil {
				n++
			}
		}
	}

	return n
}

func (h *hashicorpService) privateKeyRetrievalLoop(ticker *time.Ticker) {
	for range ticker.C {
		h.mutex.RLock()
		keyHandlers := h.keyHandlers
		h.mutex.RUnlock()

		if countRetrievedKeys(keyHandlers) == len(h.secrets) {
			ticker.Stop()
			return
		}

		for addr, byUrl := range keyHandlers {
			for u, kh := range byUrl {
				key, err := h.getKeyFromVault(kh.secret)

				if err != nil {
					path := fmt.Sprintf("%s/data/%s", kh.secret.SecretEngine, kh.secret.PrivateKeySecret)
					url := fmt.Sprintf("%v/v1/%v?version=%v", h.client.Address(), path, kh.secret.PrivateKeySecretVersion)

					log.Warn("unable to get key from Hashicorp Vault", "url", url, "err", err)
					continue
				}

				h.mutex.Lock()
				existing := h.keyHandlers[addr][u]
				updated := hashicorpKeyHandler{secret: existing.secret, key: key}
				h.keyHandlers[addr][u] = updated
				h.mutex.Unlock()
			}
		}
	}
}

func (h *hashicorpService) getKeyFromVault(s hashicorpSecretConfig) (*ecdsa.PrivateKey, error) {
	hexKey, err := h.getSecretFromVault(s.PrivateKeySecret, s.PrivateKeySecretVersion, s.SecretEngine)

	if err != nil {
		return nil, err
	}

	key, err := crypto.HexToECDSA(hexKey)

	if err != nil {
		return nil, fmt.Errorf("unable to parse data from Hashicorp Vault to *ecdsa.PrivateKey: %v", err)
	}

	return key, nil
}

func (h *hashicorpService) getSecretFromVault(name string, version int, engine string) (string, error) {
	path := fmt.Sprintf("%s/data/%s", engine, name)

	versionData := make(map[string][]string)
	versionData["version"] = []string{strconv.Itoa(version)}

	resp, err := h.client.Logical().ReadWithData(path, versionData)

	if err != nil {
		return "", fmt.Errorf("unable to get secret from Hashicorp Vault: %v", err)
	}

	respData, ok := resp.Data["data"].(map[string]interface{})

	if !ok {
		return "", errors.New("Hashicorp Vault response does not contain data")
	}

	if len(respData) != 1 {
		return "", errors.New("only one key/value pair is allowed in each Hashicorp Vault secret")
	}

	// get secret regardless of key in map
	var s interface{}
	for _, d := range respData {
		s = d
	}

	secret, ok := s.(string)

	if !ok {
		return "", errors.New("Hashicorp Vault response data is not in string format")
	}

	return secret, nil
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

type accountsByURL []accounts.Account

func (s accountsByURL) Len() int           { return len(s) }
func (s accountsByURL) Less(i, j int) bool { return s[i].URL.Cmp(s[j].URL) < 0 }
func (s accountsByURL) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (h *hashicorpService) getKey(acct accounts.Account) (*ecdsa.PrivateKey, func(), error) {
	keyHandlersByUrl, ok := h.keyHandlers[acct.Address]

	if !ok {
		return nil, nil, accounts.ErrUnknownAccount
	}

	if (acct.URL == accounts.URL{}) && len(keyHandlersByUrl) > 1 {
		ambiguousAccounts := []accounts.Account{}

		for url, _ := range keyHandlersByUrl {
			ambiguousAccounts = append(ambiguousAccounts, accounts.Account{Address: acct.Address, URL: url})
		}

		sort.Sort(accountsByURL(ambiguousAccounts))

		err := &keystore.AmbiguousAddrError{
			Addr: acct.Address,
			Matches: ambiguousAccounts,
		}

		return nil, nil, err
	}

	// return the only key for this address
	if (acct.URL == accounts.URL{}) && len(keyHandlersByUrl) == 1 {
		var keyHandler hashicorpKeyHandler

		for _, kh := range keyHandlersByUrl {
			keyHandler = kh
		}

		return h.getKeyFromHandler(keyHandler)
	}

	keyHandler, ok := keyHandlersByUrl[acct.URL]

	if !ok {
		return nil, nil, accounts.ErrUnknownAccount
	}

	return h.getKeyFromHandler(keyHandler)
}

func (h *hashicorpService) getKeyFromHandler(handler hashicorpKeyHandler) (*ecdsa.PrivateKey, func(), error) {
	if key := handler.key; key != nil {
		return key, func(){}, nil
	}

	key, err := h.getKeyFromVault(handler.secret)

	if err != nil {
		return nil, nil, err
	}

	// zeroFn zeroes the retrieved private key
	zeroFn := func () {
		b := key.D.Bits()
		for i := range b {
			b[i] = 0
		}
		key = nil
	}

	return key, zeroFn, nil
}
