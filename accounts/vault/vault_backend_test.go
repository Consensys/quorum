package vault

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/event"
	"reflect"
	"strings"
	"testing"
)

func TestNewHashicorpBackend_CreatesWalletsWithUrlsFromConfig(t *testing.T) {
	makeConfs := func (url string, urls... string) []HashicorpWalletConfig {
		var confs []HashicorpWalletConfig

		confs = append(confs, HashicorpWalletConfig{Client: hashicorpClientConfig{Url: url}})

		for _, u := range urls {
			confs = append(confs, HashicorpWalletConfig{Client: hashicorpClientConfig{Url: u}})
		}

		return confs
	}

	makeUrls := func(strUrl string, strUrls... string) []accounts.URL {
		var urls []accounts.URL

		s := strings.Split(strUrl, "://")
		scheme, path := s[0], s[1]

		urls = append(urls, accounts.URL{Scheme: scheme, Path: path})

		for _, u := range strUrls {
			s := strings.Split(u, "://")
			scheme, path := s[0], s[1]

			urls = append(urls, accounts.URL{Scheme: scheme, Path: path})
		}

		return urls
	}

	tests := map[string]struct{
		in []HashicorpWalletConfig
		wantUrls []accounts.URL
	}{
		"no config": {in: []HashicorpWalletConfig{}, wantUrls: []accounts.URL(nil)},
		"single": {in: makeConfs("http://url:1"), wantUrls: makeUrls("http://url:1")},
		"multiple": {in: makeConfs("http://url:1", "http://url:2"), wantUrls: makeUrls("http://url:1", "http://url:2")},
		"orders by url":  {
			in: makeConfs("https://url:1", "https://a:9", "http://url:2", "http://url:1"),
			wantUrls: makeUrls("http://url:1", "http://url:2", "https://a:9", "https://url:1")},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			b := NewHashicorpBackend(tt.in)

			if len(tt.wantUrls) != len(b.wallets) {
				t.Fatalf("wallets created with incorrect urls or incorrectly ordered by url: want: %v, got: %v", len(tt.wantUrls), len(b.wallets))
			}

			var gotUrls []accounts.URL

			for _, wlt := range b.wallets {
				gotUrls = append(gotUrls, wlt.URL())
			}

			if !reflect.DeepEqual(tt.wantUrls, gotUrls) {
				t.Fatalf("incorrect wallets created/wallets incorrectly ordered\nwant: %v\ngot : %v", tt.wantUrls, gotUrls)
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
	b := VaultBackend{updateFeed: &event.Feed{}}

	subscriber := make(chan accounts.WalletEvent, 1)
	b.Subscribe(subscriber)

	if b.updateScope.Count() != 1 {
		t.Fatalf("incorrect number of subscribers for backend: want: %v, got: %v", 1, b.updateScope.Count())
	}

	// mock an event
	e := accounts.WalletEvent{Wallet: vaultWallet{}, Kind: accounts.WalletOpened}
	b.updateFeed.Send(e)

	if len(subscriber) != 1 {
		t.Fatal("event not added to subscriber")
	}

	got := <-subscriber

	if !reflect.DeepEqual(e, got) {
		t.Fatalf("want: %v, got: %v", e, got)
	}
}

func TestVaultBackend_Subscribe_SubscriberReceivesEventsAddedToFeedByHashicorpWallet(t *testing.T) {
	conf := HashicorpWalletConfig{Client: hashicorpClientConfig{Url: "http://url:1"}}
	b := NewHashicorpBackend([]HashicorpWalletConfig{conf})

	if len(b.wallets) != 1 {
		t.Fatalf("incorrect number of wallets: want: %v, got: %v", 1, len(b.wallets))
	}

	w := b.wallets[0].(vaultWallet)

	subscriber := make(chan accounts.WalletEvent, 1)
	b.Subscribe(subscriber)

	if b.updateScope.Count() != 1 {
		t.Fatalf("incorrect number of subscribers for backend: want: %v, got: %v", 1, b.updateScope.Count())
	}

	// mock an event
	e := accounts.WalletEvent{Wallet: vaultWallet{}, Kind: accounts.WalletOpened}
	w.updateFeed.Send(e)

	if len(subscriber) != 1 {
		t.Fatal("event not added to subscriber")
	}

	got := <-subscriber

	if !reflect.DeepEqual(e, got) {
		t.Fatalf("want: %v, got: %v", e, got)
	}

}
