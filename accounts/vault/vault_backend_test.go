package vault

import (
	"github.com/ethereum/go-ethereum/accounts"
	"reflect"
	"strings"
	"testing"
)

func TestNewHashicorpBackend_CreatesWalletsFromConfig(t *testing.T) {
	makeConfs := func (url string, urls... string) []hashicorpWalletConfig {
		var confs []hashicorpWalletConfig

		confs = append(confs, hashicorpWalletConfig{Client: hashicorpClientConfig{Url: url}})

		for _, u := range urls {
			confs = append(confs, hashicorpWalletConfig{Client: hashicorpClientConfig{Url: u}})
		}

		return confs
	}

	// makeWlts crudely splits the urls to get them as accounts.URLs so as to not use the same parsing method as in the production code.  This is fine for simple urls but may not be suitable for tests that require more complex urls.
	makeWlts := func(url string, urls... string) []accounts.Wallet {
		var wlts []accounts.Wallet

		s := strings.Split(url, "://")
		scheme, path := s[0], s[1]

		wlts = append(wlts, vaultWallet{url: accounts.URL{Scheme: scheme, Path: path}})

		for _, u := range urls {
			s := strings.Split(u, "://")
			scheme, path := s[0], s[1]

			wlts = append(wlts, vaultWallet{url: accounts.URL{Scheme: scheme, Path: path}})
		}

		return wlts
	}

	tests := map[string]struct{
		in []hashicorpWalletConfig
		want []accounts.Wallet
	}{
		"no config": {in: []hashicorpWalletConfig{}, want: []accounts.Wallet{}},
		"single": {in: makeConfs("http://url:1"), want: makeWlts("http://url:1")},
		"multiple": {in: makeConfs("http://url:1", "http://url:2"), want: makeWlts("http://url:1", "http://url:2")},
		"orders by url":  {
			in: makeConfs("https://url:1", "https://a:9", "http://url:2", "http://url:1"),
			want: makeWlts("http://url:1", "http://url:2", "https://a:9", "https://url:1")},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			b := NewHashicorpBackend(tt.in)

			if !reflect.DeepEqual(tt.want, b.wallets) {
				t.Fatalf("\nwant: %v, \ngot : %v", tt.want, b.wallets)
			}
		})
	}

}

func TestVaultBackend_Wallets_ReturnsWallets(t *testing.T) {
	tests := map[string]struct {
			in []accounts.Wallet
			want []accounts.Wallet
	}{
		"empty": {in: []accounts.Wallet{}, want: []accounts.Wallet{}},
		"single": {in: []accounts.Wallet{vaultWallet{}}, want: []accounts.Wallet{vaultWallet{}}},
		"multiple": {in: []accounts.Wallet{vaultWallet{}, vaultWallet{}}, want: []accounts.Wallet{vaultWallet{}, vaultWallet{}}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			b := VaultBackend{wallets: tt.in}

			got := b.Wallets()

			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

func TestVaultBackend_Wallets_ReturnsCopy(t *testing.T) {
	b := VaultBackend{
		wallets: []accounts.Wallet{
			vaultWallet{url: accounts.URL{Scheme: "http", Path: "url"}},
		},
	}

	got := b.Wallets()

	got[0] = vaultWallet{url: accounts.URL{Scheme: "http", Path: "otherurl"}}

	if reflect.DeepEqual(b.wallets, got) {
		t.Fatal("changes to returned slice should not affect slice in backend")
	}
}

func TestVaultBackend_Subscribe_SubscriberReceivesEventsAddedToFeed(t *testing.T) {
	subscriber := make(chan accounts.WalletEvent, 1)
	b := VaultBackend{}

	b.Subscribe(subscriber)

	if b.updateScope.Count() != 1 {
		t.Fatalf("incorrect number of subscribers for backend: want: %v, got: %v", 1, b.updateScope.Count())
	}

	// mock an event
	event := accounts.WalletEvent{Wallet: vaultWallet{}, Kind: accounts.WalletOpened}
	b.updateFeed.Send(event)

	if len(subscriber) != 1 {
		t.Fatal("event not added to subscriber")
	}

	got := <-subscriber

	if !reflect.DeepEqual(event, got) {
		t.Fatalf("want: %v, got: %v", event, got)
	}
}
