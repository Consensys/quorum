package vault

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"reflect"
	"testing"
)

func TestVaultBackend_Wallets(t *testing.T) {
	tests := map[string]struct {
			in []accounts.Wallet
			want []accounts.Wallet
	}{
		"empty": {in: []accounts.Wallet{}, want: []accounts.Wallet{}},
		"single": {in: []accounts.Wallet{VaultWallet{}}, want: []accounts.Wallet{VaultWallet{}}},
		"multiple": {in: []accounts.Wallet{VaultWallet{}, VaultWallet{}}, want: []accounts.Wallet{VaultWallet{}, VaultWallet{}}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			b := vaultBackend{wallets: tt.in}

			got := b.Wallets()

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

func TestVaultBackend_Wallets_ReturnsCopy(t *testing.T) {
	b := vaultBackend{wallets: []accounts.Wallet{VaultWallet{}}}

	got := b.Wallets()

	got[0] = OtherVaultWallet{}

	if reflect.DeepEqual(b.wallets, got) {
		t.Fatal("changes to returned slice should not affect slice in backend")
	}
}

func TestVaultBackend_Subscribe_SubscriberReceivesEventsAddedToFeed(t *testing.T) {
	subscriber := make(chan accounts.WalletEvent, 1)
	b := vaultBackend{}

	b.Subscribe(subscriber)

	if b.updateScope.Count() != 1 {
		t.Fatalf("incorrect number of subscribers for backend: want: %v, got: %v", 1, b.updateScope.Count())
	}

	// mock an event
	event := accounts.WalletEvent{Wallet: VaultWallet{}, Kind: accounts.WalletOpened}
	b.updateFeed.Send(event)

	if len(subscriber) != 1 {
		t.Fatal("event not added to subscriber")
	}

	got := <-subscriber

	if !reflect.DeepEqual(event, got) {
		t.Fatalf("want: %v, got: %v", event, got)
	}
}

type OtherVaultWallet struct {}

func (OtherVaultWallet) URL() accounts.URL {
	panic("implement me")
}

func (OtherVaultWallet) Status() (string, error) {
	panic("implement me")
}

func (OtherVaultWallet) Open(passphrase string) error {
	panic("implement me")
}

func (OtherVaultWallet) Close() error {
	panic("implement me")
}

func (OtherVaultWallet) Accounts() []accounts.Account {
	panic("implement me")
}

func (OtherVaultWallet) Contains(account accounts.Account) bool {
	panic("implement me")
}

func (OtherVaultWallet) Derive(path accounts.DerivationPath, pin bool) (accounts.Account, error) {
	panic("implement me")
}

func (OtherVaultWallet) SelfDerive(base accounts.DerivationPath, chain ethereum.ChainStateReader) {
	panic("implement me")
}

func (OtherVaultWallet) SignHash(account accounts.Account, hash []byte) ([]byte, error) {
	panic("implement me")
}

func (OtherVaultWallet) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int, isQuorum bool) (*types.Transaction, error) {
	panic("implement me")
}

func (OtherVaultWallet) SignHashWithPassphrase(account accounts.Account, passphrase string, hash []byte) ([]byte, error) {
	panic("implement me")
}

func (OtherVaultWallet) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	panic("implement me")
}

type VaultWallet struct {}

func (VaultWallet) URL() accounts.URL {
	panic("implement me")
}

func (VaultWallet) Status() (string, error) {
	panic("implement me")
}

func (VaultWallet) Open(passphrase string) error {
	panic("implement me")
}

func (VaultWallet) Close() error {
	panic("implement me")
}

func (VaultWallet) Accounts() []accounts.Account {
	panic("implement me")
}

func (VaultWallet) Contains(account accounts.Account) bool {
	panic("implement me")
}

func (VaultWallet) Derive(path accounts.DerivationPath, pin bool) (accounts.Account, error) {
	panic("implement me")
}

func (VaultWallet) SelfDerive(base accounts.DerivationPath, chain ethereum.ChainStateReader) {
	panic("implement me")
}

func (VaultWallet) SignHash(account accounts.Account, hash []byte) ([]byte, error) {
	panic("implement me")
}

func (VaultWallet) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int, isQuorum bool) (*types.Transaction, error) {
	panic("implement me")
}

func (VaultWallet) SignHashWithPassphrase(account accounts.Account, passphrase string, hash []byte) ([]byte, error) {
	panic("implement me")
}

func (VaultWallet) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	panic("implement me")
}

