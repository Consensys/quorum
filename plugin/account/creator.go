package account

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts"
)

type creator struct {
	service Service
}

// NewCreator creates an implementation of CreatorService that simply acts as a delegate to service.  This method
// exists to allow for only the CreatorService methods to be exposed as APIs with the plugin delegate API framework.
func NewCreator(service Service) CreatorService {
	return &creator{service: service}
}

func (a *creator) NewAccount(ctx context.Context, newAccountConfig interface{}) (accounts.Account, error) {
	return a.service.NewAccount(ctx, newAccountConfig)
}

func (a *creator) ImportRawKey(ctx context.Context, rawKey string, newAccountConfig interface{}) (accounts.Account, error) {
	return a.service.ImportRawKey(ctx, rawKey, newAccountConfig)
}
