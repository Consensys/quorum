package extension

import (
	"encoding/json"
	"io/ioutil"
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

func TestWriteContentsToFileWritesOkay(t *testing.T) {
	extensionContracts := make(map[common.Address]*ExtensionContract)
	extensionContracts[common.HexToAddress("0x2222222222222222222222222222222222222222")] = &ExtensionContract{
		Address:                   common.HexToAddress("0x1111111111111111111111111111111111111111"),
		AllHaveVoted:              false,
		Initiator:                 common.HexToAddress("0x3333333333333333333333333333333333333333"),
		ManagementContractAddress: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		CreationData:              []byte("Sample Transaction Data"),
	}

	datadir, err := ioutil.TempDir("", t.Name())
	if err != nil {
		t.Errorf("could not create temp directory for test, error: %s", err.Error())
	}

	if err := writeContentsToFile(extensionContracts, datadir+"/"); err != nil {
		t.Errorf("error writing data to file, error: %s", err.Error())
	}

	data, err := ioutil.ReadFile(datadir + "/" + extensionContractData)
	if err != nil {
		t.Errorf("error reading data from file, error: %s", err.Error())
	}

	output, _ := json.Marshal(&extensionContracts)

	if string(data) != string(output) {
		t.Errorf("expected data from file different to data written, got %s, expected %s", string(data), string(output))
	}
}

//func TestWriteContentsToFileErrCantWriteToFile(t *testing.T) {}

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

func (wallet *MockWallet) SelfDerive(base accounts.DerivationPath, chain ethereum.ChainStateReader) {
	panic("not implemented")
}

func (wallet *MockWallet) SignHash(account accounts.Account, hash []byte) ([]byte, error) {
	panic("not implemented")
}

func (wallet *MockWallet) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	panic("not implemented")
}

func (wallet *MockWallet) SignHashWithPassphrase(account accounts.Account, passphrase string, hash []byte) ([]byte, error) {
	panic("not implemented")
}

func (wallet *MockWallet) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	panic("not implemented")
}

func TestGenerateTransactionOptionsErrorsWhenAccountNotFound(t *testing.T) {
	sendTxArgs := ethapi.SendTxArgs{
		From: common.Address{},
	}

	mockBackend := MockBackend{}
	manager := accounts.NewManager(&mockBackend)

	_, err := generateTransactOpts(manager, sendTxArgs)
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

	sendTxArgs := ethapi.SendTxArgs{
		From:        from,
		PrivateFor:  []string{"privateFor1", "privateFor2"},
		PrivateFrom: "privateFrom",
	}

	mockWallet := &MockWallet{isContained: true}
	mockBackend := MockBackend{wallets: []accounts.Wallet{mockWallet}}
	manager := accounts.NewManager(&mockBackend)

	generatedOptions, err := generateTransactOpts(manager, sendTxArgs)
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

	sendTxArgs := ethapi.SendTxArgs{
		From:     from,
		Gas:      &gasLimit,
		GasPrice: &gasPrice,
	}

	mockWallet := &MockWallet{isContained: true}
	mockBackend := MockBackend{wallets: []accounts.Wallet{mockWallet}}
	manager := accounts.NewManager(&mockBackend)

	generatedOptions, err := generateTransactOpts(manager, sendTxArgs)
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

