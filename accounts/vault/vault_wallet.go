package vault

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
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
	"strings"
	"sync"
	"time"
)

type VaultWallet struct {
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
	timedUnlock(acct accounts.Account, timeout time.Duration) error
	lock(acct accounts.Account) error
	writeSecret(name, value, secretEngine string) (path string, version int64, err error)
}

func newHashicorpWallet(config HashicorpWalletConfig, updateFeed *event.Feed) (VaultWallet, error) {
	var url accounts.URL

	//to parse a string url as an accounts.URL it must first be in json format
	toParse := fmt.Sprintf("\"%v\"", config.Client.Url)

	if err := url.UnmarshalJSON([]byte(toParse)); err != nil {
		return VaultWallet{}, err
	}

	w := VaultWallet{
		url: url,
		vault: newHashicorpService(config),
		updateFeed: updateFeed,
	}

	return w, nil
}

func (w VaultWallet) URL() accounts.URL {
	return w.url
}

// the vault service should return open and nil error if status is good
func (w VaultWallet) Status() (string, error) {
	return w.vault.status()
}

func (w VaultWallet) Open(passphrase string) error {
	if err := w.vault.open(); err != nil {
		return err
	}

	w.updateFeed.Send(accounts.WalletEvent{Wallet: w, Kind: accounts.WalletOpened})

	return nil
}

func (w VaultWallet) Close() error {
	return w.vault.close()
}

func (w VaultWallet) Accounts() []accounts.Account {
	return w.vault.accounts()
}

func (w VaultWallet) Contains(account accounts.Account) bool {
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

func (w VaultWallet) Derive(path accounts.DerivationPath, pin bool) (accounts.Account, error) {
	return accounts.Account{}, accounts.ErrNotSupported
}

func (w VaultWallet) SelfDerive(base accounts.DerivationPath, chain ethereum.ChainStateReader) {}

func (w VaultWallet) SignHash(account accounts.Account, hash []byte) ([]byte, error) {
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

func (w VaultWallet) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
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

func (w VaultWallet) SignHashWithPassphrase(account accounts.Account, passphrase string, hash []byte) ([]byte, error) {
	return w.SignHash(account, hash)
}

func (w VaultWallet) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return w.SignTx(account, tx, chainID)
}

func (w VaultWallet) TimedUnlock(account accounts.Account, timeout time.Duration) error {
	if !w.Contains(account) {
		return accounts.ErrUnknownAccount
	}

	return w.vault.timedUnlock(account, timeout)
}

func (w VaultWallet) Lock(account accounts.Account) error {
	if !w.Contains(account) {
		return accounts.ErrUnknownAccount
	}

	return w.vault.lock(account)
}

// Store writes the provided private key to the vault.  The hex string values of the key and address are stored in the locations specified by config.
// TODO write tests
func (w *VaultWallet) Store(key *ecdsa.PrivateKey, config HashicorpSecretConfig) (common.Address, []string, error) {
	address := crypto.PubkeyToAddress(key.PublicKey)
	// TODO check if this trim behaviour is in filesystem account creation
	addrHex := strings.TrimPrefix(address.Hex(), "0x")

	addrPath, addrVersion, err := w.vault.writeSecret(config.AddressSecret, addrHex, config.SecretEngine)

	if err != nil {
		return common.Address{}, nil, fmt.Errorf("unable to store address: %v", err.Error())
	}

	addrSecretUrl := fmt.Sprintf("%v/v1/%v?version=%v", w.url, addrPath, addrVersion)

	keyBytes := crypto.FromECDSA(key)
	keyHex := hex.EncodeToString(keyBytes)

	keyPath, keyVersion, err := w.vault.writeSecret(config.PrivateKeySecret, keyHex, config.SecretEngine)

	if err != nil {
		return common.Address{}, nil, fmt.Errorf("unable to store key: %v", err.Error())
	}

	keySecretUrl := fmt.Sprintf("%v/v1/%v?version=%v", w.url, keyPath, keyVersion)

	return address, []string{addrSecretUrl, keySecretUrl}, nil
}

type hashicorpService struct {
	client      *api.Client
	config      HashicorpClientConfig
	secrets     []HashicorpSecretConfig
	mutex       sync.RWMutex
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
	RoleIDEnv   = "VAULT_ROLE_ID"
	SecretIDEnv = "VAULT_SECRET_ID"
)

var (
	noHashicorpEnvSetErr = fmt.Errorf("environment variables must be set when creating the Hashicorp client.  Set %v and %v if the Vault is configured to use Approle authentication.  Else set %v", RoleIDEnv, SecretIDEnv, api.EnvVaultToken)
	invalidApproleAuthErr = fmt.Errorf("both %v and %v must be set if using Approle authentication", RoleIDEnv, SecretIDEnv)
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
				h.keyHandlers[acct.Address] = make(map[accounts.URL]*hashicorpKeyHandler)
			}

			keyHandlersByUrl := h.keyHandlers[acct.Address]

			if _, ok := keyHandlersByUrl[acct.URL]; ok {
				log.Warn("Hashicorp Vault key handler already exists.  Not updated.", "url", url)
				h.mutex.Unlock()
				continue
			}

			keyHandlersByUrl[acct.URL] = &hashicorpKeyHandler{secret: s}
			h.accts = append(h.accts, acct)

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
				key, err := h.getKeyFromVault(kh.secret)

				if err != nil {
					path := fmt.Sprintf("%s/data/%s", kh.secret.SecretEngine, kh.secret.PrivateKeySecret)
					url := fmt.Sprintf("%v/v1/%v?version=%v", h.client.Address(), path, kh.secret.PrivateKeySecretVersion)

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
	keyHandler, err := h.getKeyHandler(acct)

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

func (h *hashicorpService) timedUnlock(acct accounts.Account, duration time.Duration) error {
	keyHandler, err := h.getKeyHandler(acct)

	if err != nil {
		return err
	}

	alreadyExisted, err := h.updateKeyHandler(keyHandler)

	if err != nil {
		return err
	}

	if alreadyExisted {
		if keyHandler.cancel == nil {
			return nil
		}

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

func (h *hashicorpService) updateKeyHandler(handler *hashicorpKeyHandler) (bool, error) {
	var (
		key *ecdsa.PrivateKey
		alreadyExisted bool
	)

	if k := handler.key; k != nil && k.D.Int64() != 0 {
		key = k
		alreadyExisted = true
	} else {
		var err error
		key, err = h.getKeyFromVault(handler.secret)

		if err != nil {
			return false, err
		}

		handler.key = key
	}

	return alreadyExisted, nil
}

func (h *hashicorpService) lock(acct accounts.Account) error {
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

	resp, err := h.client.Logical().Write(path, data)

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

// zeroKey zeroes a private key in memory
// TODO use where appropriate
func zeroKey(k *ecdsa.PrivateKey) {
	b := k.D.Bits()
	for i := range b {
		b[i] = 0
	}
}

// CreateAccount generates a secp256k1 key and corresponding Geth address and stored both in the Vault defined in the provided config.
// The key and address are stored in hex string format.
//
// The generated key and address will be saved to only the first HashicorpSecretConfig provided.  Any other secret configs are ignored.
func CreateAccount(config HashicorpWalletConfig) (common.Address, []string, error) {
	w, err := newHashicorpWallet(config, &event.Feed{})

	if err != nil {
		return common.Address{}, nil, err
	}

	err = w.Open("")

	if err != nil {
		return common.Address{}, nil, err
	}

	if status, err := w.Status(); err != nil {
		return common.Address{}, nil, err
	} else if status != open {
		return common.Address{}, nil, fmt.Errorf("error creating Vault client, %v", status)
	}

	key, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return common.Address{}, nil, err
	}
	defer zeroKey(key)

	// This gets tricky as an error while storing the key would occur after the addr has already been stored.  The user should be made aware of this as data has been stored in the vault, so even if an error is returned address and secretInfo may still be populated.  We also need to close the wallet so  do not return straight away in the case of an error.
	var errMsgs []string

	address, secretInfo, err := w.Store(key, config.Secrets[0])
	if err != nil {
		errMsgs = append(errMsgs, err.Error())
	}

	if err := w.Close(); err != nil {
		errMsgs = append(errMsgs, fmt.Sprintf("unable to close Hashicorp Vault wallet: %v", err))
	}

	if len(errMsgs) > 0 {
		return address, secretInfo, fmt.Errorf(strings.Join(errMsgs, "\n"))
	}

	return address, secretInfo, nil
}
