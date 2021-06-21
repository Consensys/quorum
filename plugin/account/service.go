package account

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
)

type Service interface {
	Status(ctx context.Context) (string, error)
	Open(ctx context.Context, passphrase string) error
	Close(ctx context.Context) error
	Accounts(ctx context.Context) []accounts.Account
	Contains(ctx context.Context, account accounts.Account) bool
	Sign(ctx context.Context, account accounts.Account, toSign []byte) ([]byte, error)
	UnlockAndSign(ctx context.Context, account accounts.Account, toSign []byte, passphrase string) ([]byte, error)
	TimedUnlock(ctx context.Context, account accounts.Account, password string, duration time.Duration) error
	Lock(ctx context.Context, account accounts.Account) error
	CreatorService
}

type CreatorService interface {
	NewAccount(ctx context.Context, newAccountConfig interface{}) (accounts.Account, error)
	ImportRawKey(ctx context.Context, rawKey string, newAccountConfig interface{}) (accounts.Account, error)
}
