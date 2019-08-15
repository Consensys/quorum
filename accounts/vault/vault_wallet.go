package vault

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"strings"
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
	defer zero()

	if err != nil {
		return nil, err
	}

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
func (w *VaultWallet) Store(key *ecdsa.PrivateKey, config HashicorpSecretConfig) (common.Address, []string, error) {
	address := crypto.PubkeyToAddress(key.PublicKey)
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

type accountsByURL []accounts.Account

func (s accountsByURL) Len() int           { return len(s) }
func (s accountsByURL) Less(i, j int) bool { return s[i].URL.Cmp(s[j].URL) < 0 }
func (s accountsByURL) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

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
