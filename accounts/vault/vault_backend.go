package vault

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"reflect"
	"sort"
)

// BackendType is the reflect type of a vault backend.
var BackendType = reflect.TypeOf(&VaultBackend{})

// VaultBackend implements accounts.Backend to manage all wallets for a particular vendor's vault
type VaultBackend struct {
	wallets     []accounts.Wallet
	updateScope event.SubscriptionScope
	updateFeed  *event.Feed
	// Other backend impls require mutexes for safety as their wallets can change at any time (e.g. if a file/usb is added/removed).  vaultWallets can only be created at startup so there is no danger of concurrent reads and writes.
}

// NewHashicorpBackend creates a VaultBackend containing Hashicorp Vault compatible vaultWallets for each of the provided walletConfigs
func NewHashicorpBackend(walletConfigs []HashicorpWalletConfig) VaultBackend {
	wallets := []accounts.Wallet{}

	var updateFeed event.Feed

	for _, conf := range walletConfigs {
		w, err := newHashicorpWallet(conf, &updateFeed)
		if err != nil {
			log.Error("unable to create Hashicorp wallet from config", "err", err)
			continue
		}
		wallets = append(wallets, w)
	}

	sort.Sort(walletsByUrl(wallets))

	return VaultBackend{
		wallets:    wallets,
		updateFeed: &updateFeed,
	}
}

// Wallets implements accounts.Backend returning a copy of the list of wallets managed by the VaultBackend
func (b *VaultBackend) Wallets() []accounts.Wallet {
	cpy := make([]accounts.Wallet, len(b.wallets))
	copy(cpy, b.wallets)
	return cpy
}

// Subscribe implements accounts.Backend, creating an async subscription to receive notifications on the additional of vaultWallets
func (b *VaultBackend) Subscribe(sink chan<- accounts.WalletEvent) event.Subscription {
	return b.updateScope.Track(b.updateFeed.Subscribe(sink))
}

// walletsByUrl implements the sort interface to enable the sorting of a slice of wallets alphanumerically by their urls
type walletsByUrl []accounts.Wallet

func (w walletsByUrl) Len() int {
	return len(w)
}

func (w walletsByUrl) Less(i, j int) bool {
	return (w[i].URL()).Cmp(w[j].URL()) < 0
}

func (w walletsByUrl) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
