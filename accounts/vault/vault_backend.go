package vault

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/event"
)

type vaultBackend struct {
	wallets []accounts.Wallet
	updateScope event.SubscriptionScope
	updateFeed event.Feed
	// Other backend impls require mutexes for safety as their wallets can change at any time (e.g. if a file/usb is added/removed).  vaultWallets can only be created at startup so there is no chance for concurrent access.
}

func (b *vaultBackend) Wallets() []accounts.Wallet {
	cpy := make([]accounts.Wallet, len(b.wallets))
	copy(cpy, b.wallets)
	return cpy
}

func (b *vaultBackend) Subscribe(sink chan<- accounts.WalletEvent) event.Subscription {
	return b.updateScope.Track(b.updateFeed.Subscribe(sink))
}
