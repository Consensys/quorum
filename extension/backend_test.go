package extension

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/internal/ethapi"
)

type MockBackend struct {
	wallets []accounts.Wallet
}

func (backend *MockBackend) Wallets() []accounts.Wallet {
	return backend.wallets
}

func (backend *MockBackend) Subscribe(sink chan<- accounts.WalletEvent) event.Subscription {
	return nil
}

type MockWallet struct {
	isContained bool
}

func (wallet *MockWallet) URL() accounts.URL { panic("not implemented") }

func (wallet *MockWallet) Status() (string, error) { panic("not implemented") }

func (wallet *MockWallet) Open(passphrase string) error { panic("not implemented") }

func (wallet *MockWallet) Close() error { panic("not implemented") }

func (wallet *MockWallet) Accounts() []accounts.Account { panic("not implemented") }

func (wallet *MockWallet) Contains(account accounts.Account) bool { return wallet.isContained }

func (wallet *MockWallet) Derive(path accounts.DerivationPath, pin bool) (accounts.Account, error) {
	panic("not implemented")
}

func (wallet *MockWallet) SelfDerive(bases []accounts.DerivationPath, chain ethereum.ChainStateReader) {
	panic("not implemented")
}

func (wallet *MockWallet) SignData(account accounts.Account, mimeType string, data []byte) ([]byte, error) {
	panic("not implemented")
}

func (wallet *MockWallet) SignDataWithPassphrase(account accounts.Account, passphrase, mimeType string, data []byte) ([]byte, error) {
	panic("not implemented")
}

func (wallet *MockWallet) SignText(account accounts.Account, text []byte) ([]byte, error) {
	panic("not implemented")
}

func (wallet *MockWallet) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	panic("not implemented")
}

func (wallet *MockWallet) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	panic("not implemented")
}

func (wallet *MockWallet) SignTextWithPassphrase(account accounts.Account, passphrase string, hash []byte) ([]byte, error) {
	panic("not implemented")
}

func TestGenerateTransactionOptionsErrorsWhenNoPrivateParticipants(t *testing.T) {
	sendTxArgs := ethapi.SendTxArgs{
		From: common.Address{},
	}

	mockBackend := MockBackend{}
	accountManager := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: true}, &mockBackend)

	service := &PrivacyService{
		accountManager: accountManager,
	}

	_, err := service.GenerateTransactOptions(sendTxArgs)
	if err == nil {
		t.Errorf("expected err to not be nil")
		return
	}

	expectedErr := "must specify private participants"
	if err.Error() != expectedErr {
		t.Errorf("expected err to be '%s', but was '%s'", expectedErr, err.Error())
	}
}

func TestGenerateTransactionOptionsErrorsWhenAccountNotFound(t *testing.T) {
	privateTxArgs := ethapi.PrivateTxArgs{PrivateFor: []string{}}
	sendTxArgs := ethapi.SendTxArgs{
		From:          common.Address{},
		PrivateTxArgs: privateTxArgs,
	}

	mockBackend := MockBackend{}
	accountManager := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: true}, &mockBackend)

	service := &PrivacyService{
		accountManager: accountManager,
	}

	_, err := service.GenerateTransactOptions(sendTxArgs)
	if err == nil {
		t.Errorf("expected err to not be nil")
		return
	}

	expectedErr := "no wallet found for account 0x0000000000000000000000000000000000000000"
	if err.Error() != expectedErr {
		t.Errorf("expected err to be '%s', but was '%s'", expectedErr, err.Error())
	}
}

func TestGenerateTransactionOptionsGivesDefaults(t *testing.T) {
	from := common.HexToAddress("0x2222222222222222222222222222222222222222")

	privateTxArgs := ethapi.PrivateTxArgs{PrivateFor: []string{"privateFor1", "privateFor2"}, PrivateFrom: "privateFrom"}

	sendTxArgs := ethapi.SendTxArgs{
		From:          from,
		PrivateTxArgs: privateTxArgs,
	}

	mockWallet := &MockWallet{isContained: true}
	mockBackend := MockBackend{wallets: []accounts.Wallet{mockWallet}}
	accountManager := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: true}, &mockBackend)
	service := &PrivacyService{
		accountManager: accountManager,
	}

	generatedOptions, err := service.GenerateTransactOptions(sendTxArgs)
	if err != nil {
		t.Errorf("expected err to be '%s', but was '%s'", "nil", err.Error())
		return
	}

	if generatedOptions.PrivateFrom != sendTxArgs.PrivateFrom {
		t.Errorf("expected PrivateFrom to be '%s', but was '%s'", sendTxArgs.PrivateFrom, generatedOptions.PrivateFrom)
		return
	}

	if len(generatedOptions.PrivateFor) != 2 || generatedOptions.PrivateFor[0] != sendTxArgs.PrivateFor[0] || generatedOptions.PrivateFor[1] != sendTxArgs.PrivateFor[1] {
		t.Errorf("expected PrivateFor to be '%s', but was '%s'", sendTxArgs.PrivateFor, generatedOptions.PrivateFor)
		return
	}

	if generatedOptions.GasLimit != 4712384 {
		t.Errorf("expected GasLimit to be '%d', but was '%d'", 4712384, generatedOptions.GasLimit)
		return
	}

	if generatedOptions.GasPrice == nil || generatedOptions.GasPrice.Cmp(new(big.Int)) != 0 {
		t.Errorf("expected GasLimit to be '%d', but was '%d'", new(big.Int), generatedOptions.GasPrice)
		return
	}

	if generatedOptions.From != from {
		t.Errorf("expected From to be '%d', but was '%d'", from, generatedOptions.From)
		return
	}
}

func TestGenerateTransactionOptionsGivesNonDefaultsWhenSpecified(t *testing.T) {
	from := common.HexToAddress("0x2222222222222222222222222222222222222222")
	gasLimit := hexutil.Uint64(5000)
	gasPrice := hexutil.Big(*big.NewInt(50))

	privateTxArgs := ethapi.PrivateTxArgs{PrivateFor: []string{}}

	sendTxArgs := ethapi.SendTxArgs{
		From:          from,
		Gas:           &gasLimit,
		GasPrice:      &gasPrice,
		PrivateTxArgs: privateTxArgs,
	}

	mockWallet := &MockWallet{isContained: true}
	mockBackend := MockBackend{wallets: []accounts.Wallet{mockWallet}}
	accountManager := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: true}, &mockBackend)
	service := &PrivacyService{
		accountManager: accountManager,
	}

	generatedOptions, err := service.GenerateTransactOptions(sendTxArgs)
	if err != nil {
		t.Errorf("expected err to be '%s', but was '%s'", "nil", err.Error())
		return
	}

	if generatedOptions.GasLimit != 5000 {
		t.Errorf("expected GasLimit to be '%d', but was '%d'", 5000, generatedOptions.GasLimit)
		return
	}

	if generatedOptions.GasPrice == nil || generatedOptions.GasPrice.Cmp(big.NewInt(50)) != 0 {
		t.Errorf("expected GasLimit to be '%d', but was '%d'", big.NewInt(50), generatedOptions.GasPrice)
		return
	}
}
