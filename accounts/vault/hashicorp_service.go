package vault

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/hashicorp/vault/api"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type hashicorpService struct {
	config      HashicorpClientConfig
	secrets     []HashicorpSecretConfig
	mutex       sync.RWMutex
	client      *api.Client
	accts       []accounts.Account
	keyHandlers map[common.Address]map[accounts.URL]*hashicorpKeyHandler
}

func newHashicorpService(config HashicorpWalletConfig) *hashicorpService {
	s := &hashicorpService{
		config:      config.Client,
		secrets:     config.Secrets,
		keyHandlers: make(map[common.Address]map[accounts.URL]*hashicorpKeyHandler),
	}

	return s
}

type hashicorpKeyHandler struct {
	secret HashicorpSecretConfig
	mutex  sync.RWMutex
	key    *ecdsa.PrivateKey
	cancel chan struct{}
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
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	client := h.client

	if client == nil {
		return h.withAcctStatuses(closed), nil
	}

	health, err := client.Sys().Health()

	if err != nil {
		return h.withAcctStatuses(hashicorpHealthcheckFailed), hashicorpHealthcheckErr{err: err}
	}

	if !health.Initialized {
		return h.withAcctStatuses(hashicorpUninitialized), hashicorpUninitializedErr
	}

	if health.Sealed {
		return h.withAcctStatuses(hashicorpSealed), hashicorpSealedErr
	}

	return h.withAcctStatuses(open), nil
}

func (h *hashicorpService) withAcctStatuses(walletStatus string) string {
	status := []string{walletStatus}

	for addr, h := range h.keyHandlers {
		for _, hh := range h {
			var acctStatus string
			if hh.key == nil {
				acctStatus = "locked"
			} else {
				acctStatus = "unlocked"
			}
			status = append(status, fmt.Sprintf("%v: %v", addr.Hex(), acctStatus))
		}
	}

	return strings.Join(status, " | ")
}

const (
	RoleIDEnv   = "VAULT_ROLE_ID"
	SecretIDEnv = "VAULT_SECRET_ID"
)

var (
	noHashicorpEnvSetErr = fmt.Errorf("environment variables must be set when creating the Hashicorp client.  Set %v and %v if the Vault is configured to use Approle authentication.  Else set %v", RoleIDEnv, SecretIDEnv, api.EnvVaultToken)
	invalidApproleAuthErr = fmt.Errorf("both %v and %v must be set if using Approle authentication", RoleIDEnv, SecretIDEnv)
)

func (h *hashicorpService) open() error {
	if h.getClient() != nil {
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

	roleID := os.Getenv(RoleIDEnv)
	secretID := os.Getenv(SecretIDEnv)

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
	h.setClient(c)

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
		if len(h.getAccts()) == len(h.secrets) {
			ticker.Stop()
			return
		}

		for _, s := range h.secrets {
			path := fmt.Sprintf("%s/data/%s", s.SecretEngine, s.AddressSecret)

			url := fmt.Sprintf("%v/v1/%v?version=%v", h.getClient().Address(), path, s.AddressSecretVersion)

			h.mutex.RLock()
			address, err := h.getAddressFromVault(s)
			h.mutex.RUnlock()

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
				h.keyHandlers[acct.Address] = make(map[accounts.URL]*hashicorpKeyHandler)
			}

			keyHandlersByUrl := h.keyHandlers[acct.Address]

			if _, ok := keyHandlersByUrl[acct.URL]; ok {
				log.Warn("Hashicorp Vault key handler already exists.  Not updated.", "url", url)
			} else {
				keyHandlersByUrl[acct.URL] = &hashicorpKeyHandler{secret: s}
				h.accts = append(h.accts, acct)
			}

			h.mutex.Unlock()
		}
	}
}

func (h *hashicorpService) getAddressFromVault(s HashicorpSecretConfig) (common.Address, error) {
	hexAddr, err := h.getSecretFromVault(s.AddressSecret, s.AddressSecretVersion, s.SecretEngine)

	if err != nil {
		return common.Address{}, err
	}

	return common.HexToAddress(hexAddr), nil
}

func countRetrievedKeys(keyHandlers map[common.Address]map[accounts.URL]*hashicorpKeyHandler) int {
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
				h.mutex.RLock()
				key, err := h.getKeyFromVault(kh.secret)
				h.mutex.RUnlock()

				if err != nil {
					path := fmt.Sprintf("%s/data/%s", kh.secret.SecretEngine, kh.secret.PrivateKeySecret)
					url := fmt.Sprintf("%v/v1/%v?version=%v", h.getClient().Address(), path, kh.secret.PrivateKeySecretVersion)

					log.Warn("unable to get key from Hashicorp Vault", "url", url, "err", err)
					continue
				}

				h.mutex.Lock()
				h.keyHandlers[addr][u].key = key
				h.mutex.Unlock()
			}
		}
	}
}

func (h *hashicorpService) getKeyFromVault(s HashicorpSecretConfig) (*ecdsa.PrivateKey, error) {
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

	if resp == nil {
		return "", fmt.Errorf("no data for secret in Hashicorp Vault")
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
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.client = nil

	return nil
}

func (h *hashicorpService) accounts() []accounts.Account {
	accts := h.getAccts()
	cpy := make([]accounts.Account, len(accts))
	copy(cpy, accts)

	return cpy
}

func (h *hashicorpService) getKey(acct accounts.Account) (*ecdsa.PrivateKey, func(), error) {
	h.mutex.RLock()
	keyHandler, err := h.getKeyHandler(acct)
	h.mutex.RUnlock()

	if err != nil {
		return nil, nil, err
	}

	return h.getKeyFromHandler(*keyHandler)
}

func (h *hashicorpService) getKeyHandler(acct accounts.Account) (*hashicorpKeyHandler, error) {
	keyHandlersByUrl, ok := h.keyHandlers[acct.Address]

	if !ok {
		return nil, accounts.ErrUnknownAccount
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

		return nil, err
	}

	// return the only key for this address
	if (acct.URL == accounts.URL{}) && len(keyHandlersByUrl) == 1 {
		var keyHandler *hashicorpKeyHandler

		for _, kh := range keyHandlersByUrl {
			keyHandler = kh
			break
		}

		return keyHandler, nil
	}

	keyHandler, ok := keyHandlersByUrl[acct.URL]

	if !ok {
		return nil, accounts.ErrUnknownAccount
	}

	return keyHandler, nil
}

func (h *hashicorpService) getKeyFromHandler(handler hashicorpKeyHandler) (*ecdsa.PrivateKey, func(), error) {
	if key := handler.key; key != nil {
		return key, func(){}, nil
	}

	h.mutex.RLock()
	key, err := h.getKeyFromVault(handler.secret)
	h.mutex.RUnlock()

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

func (h *hashicorpService) timedUnlock(acct accounts.Account, duration time.Duration) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	keyHandler, err := h.getKeyHandler(acct)

	if err != nil {
		return err
	}

	alreadyUnlocked, err := h.unlockKeyHandler(keyHandler)

	if err != nil {
		return err
	}

	if alreadyUnlocked {
		// indefinitely unlocked, do not override
		if keyHandler.cancel == nil {
			return nil
		}

		// cancel existing timed unlock
		close(keyHandler.cancel)
		keyHandler.cancel = nil
	}

	if duration == 0 {
		keyHandler.cancel = nil
	} else if duration > 0 {
		go keyHandler.timedLock(duration)
	}

	return nil
}

func (h *hashicorpService) unlockKeyHandler(handler *hashicorpKeyHandler) (alreadyUnlocked bool, err error) {
	if k := handler.key; k != nil && k.D.Int64() != 0 {
		return true, nil
	}

	key, err := h.getKeyFromVault(handler.secret)

	if err != nil {
		return false, err
	}

	handler.key = key

	return false, nil
}

func (h *hashicorpService) lock(acct accounts.Account) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	keyHandler, err := h.getKeyHandler(acct)

	if err != nil {
		return err
	}

	// cancel any existing timed lock
	if keyHandler.cancel != nil {
		close(keyHandler.cancel)
		keyHandler.cancel = nil
	}

	// zero the key if it is present
	if keyHandler.key != nil {
		b := keyHandler.key.D.Bits()
		for i := range b {
			b[i] = 0
		}
		keyHandler.key = nil
	}

	return nil
}

// Even if error is returned, data might have been written to Vault.  path and version may contain useful information even in the case of an error. version = -1 indicates no version was retrieved from the Vault (Vault version numbers are >= 0)
func (h *hashicorpService) writeSecret(name, value, secretEngine string) (string, int64, error) {
	path := fmt.Sprintf("%s/data/%s", secretEngine, name)

	data := make(map[string]interface{})
	data["data"] = map[string]interface{}{
		"secret": value,
	}

	resp, err := h.getClient().Logical().Write(path, data)

	if err != nil {
		return "", -1, fmt.Errorf("unable to write secret to vault: %v", err)
	}

	v, ok := resp.Data["version"]

	if !ok {
		v = json.Number("-1")
	}

	vJson, ok := v.(json.Number)

	vInt, err := vJson.Int64()

	if err != nil {
		return path, -1, fmt.Errorf("unable to convert version in Vault response to int64: version number is %v", vJson.String())
	}

	return path, vInt, nil
}

func (h *hashicorpKeyHandler) timedLock(duration time.Duration) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	t := time.NewTimer(duration)
	defer t.Stop()
	h.cancel = make(chan struct{})

	select {
	case <-t.C:
		b := h.key.D.Bits()
		for i := range b {
			b[i] = 0
		}
		h.key = nil
	case <-h.cancel:
		//do nothing
	}
}

// Each of these getters takes an RLock so care should be taken not to call these within an existing Lock otherwise this will cause a deadlock.
// Should not be used if storing the returned client in a variable for later use as the fact it is a pointer means that you should be locking for the entirety of the usage of the client.
func (h *hashicorpService) getClient() *api.Client {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return h.client
}

func (h *hashicorpService) getAccts() []accounts.Account {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return h.accts
}

func (h *hashicorpService) setClient(c *api.Client) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.client = c
}
