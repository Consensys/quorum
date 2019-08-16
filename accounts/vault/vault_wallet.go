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

// VaultWallet implements accounts.Wallet and represents the common functionality shared by all wallets that manage accounts stored in vaults
type VaultWallet struct {
	url        accounts.URL
	vault      vaultService
	updateFeed *event.Feed
}

// vaultService defines the vendor-specific functionality that vault wallets must implement
type vaultService interface {
	// Status returns a textual status to aid the user in the current state of the
	// wallet. It also returns an error indicating any failure the wallet might have
	// encountered.
	status() (string, error)

	// open initializes access to a wallet.  It establishes a connection to the vault but does not retrieve private keys from the vault by default.
	open() error

	// close releases any resources held by an open wallet instance.
	close() error

	// accounts returns a copy of the list of signing accounts the wallet is currently aware of.
	accounts() []accounts.Account

	// getKey returns the key for the given account, making a request to the vault if the account is locked.  zeroFn is the corresponding zero function for the returned key and should be called to clean up once the key has been used.
	getKey(acct accounts.Account) (key *ecdsa.PrivateKey, zeroFn func(), err error)

	// timedUnlock unlocks the given account for the duration of timeout. A timeout of 0 unlocks the account until the program exits.
	//
	// If the account address is already unlocked for a duration, TimedUnlock extends or
	// shortens the active unlock timeout. If the address was previously unlocked
	// indefinitely the timeout is not altered.
	timedUnlock(acct accounts.Account, timeout time.Duration) error

	// lock removes the private key for the given account from memory.
	lock(acct accounts.Account) error

	// writeSecret writes a new secret with name and value to the vault.
	//
	// The returned path is the location path for the secret within the vault.  version is the version of the new secret.
	//
	// TODO make vendor agnostic
	writeSecret(name, value, secretEngine string) (path string, version int64, err error)
}

// newHashicorpWallet creates a Hashicorp Vault compatible VaultWallet using the provided config.  Wallet events will be applied to updateFeed.
func newHashicorpWallet(config HashicorpWalletConfig, updateFeed *event.Feed) (VaultWallet, error) {
	var url accounts.URL

	//to parse a string url as an accounts.URL it must first be in json format
	toParse := fmt.Sprintf("\"%v\"", config.Client.Url)

	if err := url.UnmarshalJSON([]byte(toParse)); err != nil {
		return VaultWallet{}, err
	}

	w := VaultWallet{
		url:        url,
		vault:      newHashicorpService(config),
		updateFeed: updateFeed,
	}

	return w, nil
}

// URL implements accounts.Wallet, returning the URL of the configured vault.
func (w VaultWallet) URL() accounts.URL {
	return w.url
}

// Status implements accounts.Wallet, returning a custom status message from the
// underlying vendor-specific vault service implementation.
func (w VaultWallet) Status() (string, error) {
	return w.vault.status()
}

// Open implements accounts.Wallet, attempting to open a connection to the
// vault.
func (w VaultWallet) Open(passphrase string) error {
	if err := w.vault.open(); err != nil {
		return err
	}

	w.updateFeed.Send(accounts.WalletEvent{Wallet: w, Kind: accounts.WalletOpened})

	return nil
}

// Close implements accounts.Wallet, closing the connection to the vault.
func (w VaultWallet) Close() error {
	return w.vault.close()
}

// Accounts implements accounts.Wallet, returning the list of accounts the wallet is
// currently aware of.
func (w VaultWallet) Accounts() []accounts.Account {
	return w.vault.accounts()
}

// Contains implements accounts.Wallet, returning whether a particular account is
// or is not managed by this wallet. An account with no url only needs to match
// on the address to return true.
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

// Derive implements accounts.Wallet, but is a noop for vault wallets since there
// is no notion of hierarchical account derivation for vault-stored accounts.
func (w VaultWallet) Derive(path accounts.DerivationPath, pin bool) (accounts.Account, error) {
	return accounts.Account{}, accounts.ErrNotSupported
}

// SelfDerive implements accounts.Wallet, but is a noop for vault wallets since
// there is no notion of hierarchical account derivation for vault-stored accounts.
func (w VaultWallet) SelfDerive(base accounts.DerivationPath, chain ethereum.ChainStateReader) {}

// SignHash implements accounts.Wallet, attempting to sign the given hash with
// the given account. If the wallet does not manage this particular account, an
// error is returned.
//
// If the account is locked, the wallet will unlock it but only for the duration
// of the signing.
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

// SignTx implements accounts.Wallet, attempting to sign the given transaction
// with the given account. If the wallet does not manage this particular account,
// an error is returned.
//
// If the account is locked, the wallet will unlock it but only for the duration
// of the signing.
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

// SignHashWithPassphrase implements accounts.Wallet, attempting to sign the given
// hash with the given account using passphrase as extra authentication.
// For Vault wallets the passphrase is ignored and so SignHash is called.
func (w VaultWallet) SignHashWithPassphrase(account accounts.Account, passphrase string, hash []byte) ([]byte, error) {
	return w.SignHash(account, hash)
}

// SignTxWithPassphrase implements accounts.Wallet, attempting to sign the given
// transaction with the given account using passphrase as extra authentication.
// For Vault wallets the passphrase is ignored and so SignTx is called.
func (w VaultWallet) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return w.SignTx(account, tx, chainID)
}

// TimedUnlock unlocks the given account for the duration of timeout. A timeout of 0 unlocks the account until the program exits.
//
// If the account address is already unlocked for a duration, TimedUnlock extends or
// shortens the active unlock timeout. If the address was previously unlocked
// indefinitely the timeout is not altered.
//
// If the wallet does not manage this particular account, an error is returned.
func (w VaultWallet) TimedUnlock(account accounts.Account, timeout time.Duration) error {
	if !w.Contains(account) {
		return accounts.ErrUnknownAccount
	}

	return w.vault.timedUnlock(account, timeout)
}

// Lock locks the given account thereby removing the corresponding private key from memory. If the
// wallet does not manage this particular account, an error is returned.
func (w VaultWallet) Lock(account accounts.Account) error {
	if !w.Contains(account) {
		return accounts.ErrUnknownAccount
	}

	return w.vault.lock(account)
}

// Store writes the provided private key to the vault.  The hex string values of the key and address are stored in the locations specified by config.
// TODO make vendor agnostic
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

// accountsByURL implements the sort interface to enable the sorting of a slice of accounts alphanumerically by their urls
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

// CreateAccount generates a secp256k1 key and corresponding Geth address and stores both in the Vault defined in the provided config.
// The key and address are stored in hex string format.
//
// The generated key and address will be saved to only the first HashicorpSecretConfig provided in config.  Any other secret configs are ignored.
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
