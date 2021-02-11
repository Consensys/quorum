package pluggable

import (
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/event"
	plugin "github.com/ethereum/go-ethereum/plugin/account"
)

var BackendType = reflect.TypeOf(&Backend{})

type Backend struct {
	wallets []accounts.Wallet
}

func NewBackend() *Backend {
	return &Backend{
		wallets: []accounts.Wallet{
			&wallet{
				url: accounts.URL{
					Scheme: "plugin",
					Path:   "account",
				},
			},
		},
	}
}

func (b *Backend) Wallets() []accounts.Wallet {
	cpy := make([]accounts.Wallet, len(b.wallets))
	copy(cpy, b.wallets)
	return cpy
}

// Subscribe implements accounts.Backend, creating a new subscription that is a no-op and simply exits when the Unsubscribe is called
func (b *Backend) Subscribe(_ chan<- accounts.WalletEvent) event.Subscription {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		<-quit
		return nil
	})
}

func (b *Backend) SetPluginService(s plugin.Service) error {
	return b.wallet().setPluginService(s)
}

func (b *Backend) TimedUnlock(account accounts.Account, password string, duration time.Duration) error {
	return b.wallet().timedUnlock(account, password, duration)
}

func (b *Backend) Lock(account accounts.Account) error {
	return b.wallet().lock(account)
}

// AccountCreator is the interface that wraps the plugin account creation methods.
// This interface is used to simplify the pluggable.Backend API available to the account plugin CLI and enables easier testing.
type AccountCreator interface {
	NewAccount(newAccountConfig interface{}) (accounts.Account, error)
	ImportRawKey(rawKey string, newAccountConfig interface{}) (accounts.Account, error)
}

func (b *Backend) NewAccount(newAccountConfig interface{}) (accounts.Account, error) {
	return b.wallet().newAccount(newAccountConfig)
}

func (b *Backend) ImportRawKey(rawKey string, newAccountConfig interface{}) (accounts.Account, error) {
	return b.wallet().importRawKey(rawKey, newAccountConfig)
}

func (b *Backend) wallet() *wallet {
	return b.wallets[0].(*wallet)
}
