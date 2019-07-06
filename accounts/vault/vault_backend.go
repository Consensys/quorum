package vault

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"sort"
)

type VaultBackend struct {
	wallets []accounts.Wallet
	updateScope event.SubscriptionScope
	updateFeed event.Feed
	// Other backend impls require mutexes for safety as their wallets can change at any time (e.g. if a file/usb is added/removed).  vaultWallets can only be created at startup so there is no danger of concurrent reads and writes.
}

func (b *VaultBackend) Wallets() []accounts.Wallet {
	cpy := make([]accounts.Wallet, len(b.wallets))
	copy(cpy, b.wallets)
	return cpy
}

func (b *VaultBackend) Subscribe(sink chan<- accounts.WalletEvent) event.Subscription {
	return b.updateScope.Track(b.updateFeed.Subscribe(sink))
}

func NewHashicorpBackend(walletConfigs []hashicorpWalletConfig) VaultBackend {
	wallets := []accounts.Wallet{}

	for _, conf := range walletConfigs {
		w, err := newHashicorpWallet(conf)
		if err != nil {
			// do something with error and do not append returned w to wallets
			log.Error("unable to create Hashicorp wallet from config", "err", err)
			continue
		}
		wallets = append(wallets, w)
	}

	sort.Sort(walletsByUrl(wallets))

	return VaultBackend{
		wallets: wallets,
	}
}

func newHashicorpWallet(config hashicorpWalletConfig) (vaultWallet, error) {
	var url accounts.URL

	//to parse a string url as an accounts.URL it must first be in json format
	toParse := fmt.Sprintf("\"%v\"", config.Client.Url)

	if err := url.UnmarshalJSON([]byte(toParse)); err != nil {
		return vaultWallet{}, err
	}

	return vaultWallet{Url: url}, nil
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
