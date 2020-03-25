package privacyExtension

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/core/rawdb"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	extension "github.com/ethereum/go-ethereum/extension/extensionContracts"
)

func TestLogContainsExtensionTopicWithWrongLengthReturnsFalse(t *testing.T) {
	testLog := &types.Log{
		Topics: []common.Hash{{}, {}},
	}

	contained := logContainsExtensionTopic(testLog)

	if contained {
		t.Errorf("expected value '%t', but got '%t'", false, contained)
	}
}

func TestLogContainsExtensionTopicWithWrongHashReturnsFalse(t *testing.T) {
	testLog := &types.Log{
		Topics: []common.Hash{common.HexToHash("0xc05e76a85299aba9028bd0e0c3ab6fd798db442ed25ce08eb9d2098acc5a2904")},
	}

	contained := logContainsExtensionTopic(testLog)

	if contained {
		t.Errorf("expected value '%t', but got '%t'", false, contained)
	}
}

func TestLogContainsExtensionTopicWithCorrectHashReturnsTrue(t *testing.T) {
	testLog := &types.Log{
		Topics: []common.Hash{common.HexToHash("0x67a92539f3cbd7c5a9b36c23c0e2beceb27d2e1b3cd8eda02c623689267ae71e")},
	}

	contained := logContainsExtensionTopic(testLog)

	if !contained {
		t.Errorf("expected value '%t', but got '%t'", true, contained)
	}
}

func TestStateSetWithListedAccounts(t *testing.T) {
	input := `{"0x2222222222222222222222222222222222222222":{"state":{"balance":"22","nonce":5,"root":"56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","codeHash":"87874902497a5bb968da31a2998d8f22e949d1ef6214bcdedd8bae24cca4b9e3","code":"03030303030303","storage":{}}}}`
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()))

	var accounts map[string]extension.AccountWithMetadata
	if err := json.Unmarshal([]byte(input), &accounts); err != nil {
		t.Errorf("error when unmarshalling static data: %s", err.Error())
	}

	success := setState(statedb, accounts)
	if !success {
		t.Errorf("unexpected error when setting state")
	}

	address := common.HexToAddress("0x2222222222222222222222222222222222222222")
	balance := statedb.GetBalance(address)
	code := statedb.GetCode(address)
	nonce := statedb.GetNonce(address)
	storage, _ := statedb.GetStorageRoot(address)

	if balance.Uint64() != 22 {
		t.Errorf("expect Balance value of '%d', but got '%d'", 22, balance.Uint64())
		return
	}

	expectedCode := []byte{3, 3, 3, 3, 3, 3, 3}
	if !bytes.Equal(code, expectedCode) {
		t.Errorf("expect Code value of '%d', but got '%d'", expectedCode, code)
		return
	}

	if nonce != 5 {
		t.Errorf("expect Nonce value of '%d', but got '%d'", 5, nonce)
		return
	}

	expectedStorageHash := common.FromHex("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")
	if !bytes.Equal(storage.Bytes(), expectedStorageHash) {
		t.Errorf("expect Storage value of '%d', but got '%s'", expectedStorageHash, storage)
		return
	}
}

func TestStateSetWithListedAccountsFailsOnInvalidBalance(t *testing.T) {
	input := `{"0x2222222222222222222222222222222222222222":{"state":{"balance":"invalid","nonce":5,"root":"56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","codeHash":"87874902497a5bb968da31a2998d8f22e949d1ef6214bcdedd8bae24cca4b9e3","code":"03030303030303","storage":{}}}}`
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()))

	var accounts map[string]extension.AccountWithMetadata
	if err := json.Unmarshal([]byte(input), &accounts); err != nil {
		t.Errorf("error when unmarshalling static data: %s", err.Error())
	}

	success := setState(statedb, accounts)
	if success {
		t.Errorf("error expected when setting state")
	}
}
