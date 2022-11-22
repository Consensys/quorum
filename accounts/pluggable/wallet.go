package pluggable

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	plugin "github.com/ethereum/go-ethereum/plugin/account"
)

type wallet struct {
	url           accounts.URL
	mu            sync.Mutex
	pluginService plugin.Service
}

func (w *wallet) setPluginService(s plugin.Service) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.pluginService = s

	return nil
}

func (w *wallet) URL() accounts.URL {
	return w.url
}

func (w *wallet) Status() (string, error) {
	return w.pluginService.Status(context.Background())
}

func (w *wallet) Open(passphrase string) error {
	return w.pluginService.Open(context.Background(), passphrase)
}

func (w *wallet) Close() error {
	return w.pluginService.Close(context.Background())
}

func (w *wallet) Accounts() []accounts.Account {
	if w.pluginService == nil {
		return []accounts.Account{}
	}
	return w.pluginService.Accounts(context.Background())
}

func (w *wallet) Contains(account accounts.Account) bool {
	return w.pluginService.Contains(context.Background(), account)
}

func (w *wallet) Derive(_ accounts.DerivationPath, _ bool) (accounts.Account, error) {
	return accounts.Account{}, accounts.ErrNotSupported
}

func (w *wallet) SelfDerive(_ []accounts.DerivationPath, _ ethereum.ChainStateReader) {}

func (w *wallet) SignData(account accounts.Account, _ string, data []byte) ([]byte, error) {
	return w.pluginService.Sign(context.Background(), account, crypto.Keccak256(data))
}

func (w *wallet) SignDataWithPassphrase(account accounts.Account, passphrase, _ string, data []byte) ([]byte, error) {
	return w.pluginService.UnlockAndSign(context.Background(), account, crypto.Keccak256(data), passphrase)
}

func (w *wallet) SignText(account accounts.Account, text []byte) ([]byte, error) {
	return w.pluginService.Sign(context.Background(), account, accounts.TextHash(text))
}

func (w *wallet) SignTextWithPassphrase(account accounts.Account, passphrase string, text []byte) ([]byte, error) {
	return w.pluginService.UnlockAndSign(context.Background(), account, accounts.TextHash(text), passphrase)
}

func (w *wallet) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	toSign, signer := prepareTxForSign(tx, chainID)

	sig, err := w.pluginService.Sign(context.Background(), account, toSign.Bytes())
	if err != nil {
		return nil, err
	}

	return tx.WithSignature(signer, sig)
}

func (w *wallet) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	toSign, signer := prepareTxForSign(tx, chainID)

	sig, err := w.pluginService.UnlockAndSign(context.Background(), account, toSign.Bytes(), passphrase)
	if err != nil {
		return nil, err
	}

	return tx.WithSignature(signer, sig)
}

func (w *wallet) timedUnlock(account accounts.Account, password string, duration time.Duration) error {
	return w.pluginService.TimedUnlock(context.Background(), account, password, duration)
}

func (w *wallet) lock(account accounts.Account) error {
	return w.pluginService.Lock(context.Background(), account)
}

func (w *wallet) newAccount(newAccountConfig interface{}) (accounts.Account, error) {
	return w.pluginService.NewAccount(context.Background(), newAccountConfig)
}

func (w *wallet) importRawKey(rawKey string, newAccountConfig interface{}) (accounts.Account, error) {
	return w.pluginService.ImportRawKey(context.Background(), rawKey, newAccountConfig)
}

// prepareTxForSign determines which Signer to use for the given tx and chainID, and returns the Signer's hash of the tx and the Signer itself
func prepareTxForSign(tx *types.Transaction, chainID *big.Int) (common.Hash, types.Signer) {
	var s types.Signer

	if tx.IsPrivate() {
		s = types.QuorumPrivateTxSigner{}
	} else {
		s = types.LatestSignerForChainID(chainID)
	}

	return s.Hash(tx), s
}
