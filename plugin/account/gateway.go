package account

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/jpmorganchase/quorum-account-plugin-sdk-go/proto"
)

type service struct {
	client proto.AccountServiceClient
}

func (g *service) Status(ctx context.Context) (string, error) {
	resp, err := g.client.Status(ctx, &proto.StatusRequest{})
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", errors.New("empty response from plugin")
	}
	return resp.Status, err
}

func (g *service) Open(ctx context.Context, passphrase string) error {
	_, err := g.client.Open(ctx, &proto.OpenRequest{Passphrase: passphrase})
	return err
}

func (g *service) Close(ctx context.Context) error {
	_, err := g.client.Close(ctx, &proto.CloseRequest{})
	return err
}

func (g *service) Accounts(ctx context.Context) []accounts.Account {
	resp, err := g.client.Accounts(ctx, &proto.AccountsRequest{})
	if err != nil {
		log.Error("unable to get accounts from plugin account store", "err", err)
		return []accounts.Account{}
	}
	if resp == nil {
		log.Error("empty response from plugin")
		return []accounts.Account{}
	}

	return asAccounts(resp.Accounts)
}

func (g *service) Contains(ctx context.Context, account accounts.Account) bool {
	resp, err := g.client.Contains(ctx, &proto.ContainsRequest{Address: account.Address.Bytes()})
	if err != nil {
		log.Error("unable to check contents of plugin account store", "err", err)
	}
	if resp == nil {
		log.Error("empty response from plugin")
		return false
	}

	return resp.IsContained
}

func (g *service) Sign(ctx context.Context, account accounts.Account, toSign []byte) ([]byte, error) {
	resp, err := g.client.Sign(ctx, &proto.SignRequest{
		Address: account.Address.Bytes(),
		ToSign:  toSign,
	})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("empty response from plugin")
	}

	return resp.Sig, nil
}

func (g *service) UnlockAndSign(ctx context.Context, account accounts.Account, toSign []byte, passphrase string) ([]byte, error) {
	resp, err := g.client.UnlockAndSign(ctx, &proto.UnlockAndSignRequest{
		Address:    account.Address.Bytes(),
		ToSign:     toSign,
		Passphrase: passphrase,
	})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("empty response from plugin")
	}

	return resp.Sig, nil
}

func (g *service) TimedUnlock(ctx context.Context, account accounts.Account, password string, duration time.Duration) error {
	_, err := g.client.TimedUnlock(ctx, &proto.TimedUnlockRequest{Address: account.Address.Bytes(), Password: password, Duration: duration.Nanoseconds()})
	return err
}

func (g *service) Lock(ctx context.Context, account accounts.Account) error {
	_, err := g.client.Lock(ctx, &proto.LockRequest{Address: account.Address.Bytes()})
	return err
}

func (g *service) NewAccount(ctx context.Context, newAccountConfig interface{}) (accounts.Account, error) {
	byt, err := json.Marshal(newAccountConfig)
	if err != nil {
		return accounts.Account{}, err
	}

	req := &proto.NewAccountRequest{
		NewAccountConfig: byt,
	}
	resp, err := g.client.NewAccount(ctx, req)
	if err != nil {
		return accounts.Account{}, err
	}

	acct, err := asAccount(resp.Account)
	if err != nil {
		return accounts.Account{}, err
	}

	return acct, nil
}

func (g *service) ImportRawKey(ctx context.Context, rawKey string, newAccountConfig interface{}) (accounts.Account, error) {
	byt, err := json.Marshal(newAccountConfig)
	if err != nil {
		return accounts.Account{}, err
	}
	// validate the rawKey
	_, err = crypto.HexToECDSA(rawKey)
	if err != nil {
		return accounts.Account{}, err
	}
	req := &proto.ImportRawKeyRequest{
		RawKey:           rawKey,
		NewAccountConfig: byt,
	}
	resp, err := g.client.ImportRawKey(ctx, req)
	if err != nil {
		return accounts.Account{}, err
	}

	acct, err := asAccount(resp.Account)
	if err != nil {
		return accounts.Account{}, err
	}

	return acct, nil
}

func asAccounts(pAccts []*proto.Account) []accounts.Account {
	accts := make([]accounts.Account, 0, len(pAccts))

	for i, pAcct := range pAccts {
		acct, err := asAccount(pAcct)
		if err != nil {
			log.Error("unable to parse account from plugin account store", "index", i, "err", err)
			continue
		}

		accts = append(accts, acct)
	}

	return accts
}

func asAccount(pAcct *proto.Account) (accounts.Account, error) {
	addr := strings.TrimSpace(common.Bytes2Hex(pAcct.Address))

	if !common.IsHexAddress(addr) {
		return accounts.Account{}, fmt.Errorf("invalid hex address: %v", addr)
	}

	url, err := ToUrl(pAcct.Url)
	if err != nil {
		return accounts.Account{}, err
	}

	acct := accounts.Account{
		Address: common.HexToAddress(addr),
		URL:     url,
	}

	return acct, nil
}

func ToUrl(strUrl string) (accounts.URL, error) {
	if strUrl == "" {
		return accounts.URL{}, nil
	}

	//to parse a string url as an accounts.URL it must first be in json format
	toParse := fmt.Sprintf("\"%v\"", strUrl)

	var url accounts.URL
	if err := url.UnmarshalJSON([]byte(toParse)); err != nil {
		return accounts.URL{}, err
	}

	return url, nil
}
