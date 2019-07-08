package vault

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"sort"
)

type VaultBackend struct {
	wallets []accounts.Wallet
	updateScope event.SubscriptionScope
	updateFeed *event.Feed
	// Other backend impls require mutexes for safety as their wallets can change at any time (e.g. if a file/usb is added/removed).  vaultWallets can only be created at startup so there is no danger of concurrent reads and writes.
}

func NewHashicorpBackend(walletConfigs []hashicorpWalletConfig) VaultBackend {
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
		wallets: wallets,
		updateFeed: &updateFeed,
	}
}

func (b *VaultBackend) Wallets() []accounts.Wallet {
	cpy := make([]accounts.Wallet, len(b.wallets))
	copy(cpy, b.wallets)
	return cpy
}

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
