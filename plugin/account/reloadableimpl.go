package account

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/log"
)

type DispenseFunc func() (Service, error)

type ReloadableService struct {
	DispenseFunc DispenseFunc
}

func (am *ReloadableService) Status(ctx context.Context) (string, error) {
	s, err := am.DispenseFunc()
	if err != nil {
		return "", err
	}
	return s.Status(ctx)
}

func (am *ReloadableService) Open(ctx context.Context, passphrase string) error {
	s, err := am.DispenseFunc()
	if err != nil {
		return err
	}
	return s.Open(ctx, passphrase)
}

func (am *ReloadableService) Close(ctx context.Context) error {
	s, err := am.DispenseFunc()
	if err != nil {
		return err
	}
	return s.Close(ctx)
}

func (am *ReloadableService) Accounts(ctx context.Context) []accounts.Account {
	s, err := am.DispenseFunc()
	if err != nil {
		log.Error("unable to dispense account plugin", "err", err)
		return []accounts.Account{}
	}
	return s.Accounts(ctx)
}

func (am *ReloadableService) Contains(ctx context.Context, account accounts.Account) bool {
	s, err := am.DispenseFunc()
	if err != nil {
		log.Error("unable to dispense account plugin", "err", err)
		return false
	}
	return s.Contains(ctx, account)
}

func (am *ReloadableService) Sign(ctx context.Context, account accounts.Account, toSign []byte) ([]byte, error) {
	s, err := am.DispenseFunc()
	if err != nil {
		return nil, err
	}
	return s.Sign(ctx, account, toSign)
}

func (am *ReloadableService) UnlockAndSign(ctx context.Context, account accounts.Account, toSign []byte, passphrase string) ([]byte, error) {
	s, err := am.DispenseFunc()
	if err != nil {
		return nil, err
	}
	return s.UnlockAndSign(ctx, account, toSign, passphrase)
}

func (am *ReloadableService) TimedUnlock(ctx context.Context, account accounts.Account, password string, duration time.Duration) error {
	s, err := am.DispenseFunc()
	if err != nil {
		return err
	}
	return s.TimedUnlock(ctx, account, password, duration)
}

func (am *ReloadableService) Lock(ctx context.Context, account accounts.Account) error {
	s, err := am.DispenseFunc()
	if err != nil {
		return err
	}
	return s.Lock(ctx, account)
}

func (am *ReloadableService) NewAccount(ctx context.Context, newAccountConfig interface{}) (accounts.Account, error) {
	s, err := am.DispenseFunc()
	if err != nil {
		return accounts.Account{}, err
	}
	return s.NewAccount(ctx, newAccountConfig)
}

func (am *ReloadableService) ImportRawKey(ctx context.Context, rawKey string, newAccountConfig interface{}) (accounts.Account, error) {
	s, err := am.DispenseFunc()
	if err != nil {
		return accounts.Account{}, err
	}
	return s.ImportRawKey(ctx, rawKey, newAccountConfig)
}
