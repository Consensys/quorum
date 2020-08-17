package privacyExtension

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	extension "github.com/ethereum/go-ethereum/extension/extensionContracts"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/stretchr/testify/assert"
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
		Topics: []common.Hash{common.HexToHash("0xf20540914db019dd7c8d05ed165316a58d1583642772ac46f3d0c29b8644bd36")},
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

func createStateDb(t *testing.T) *state.StateDB {
	input := `{"0x2222222222222222222222222222222222222222":{"state":{"balance":"22","nonce":5,"root":"56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","codeHash":"87874902497a5bb968da31a2998d8f22e949d1ef6214bcdedd8bae24cca4b9e3","code":"03030303030303","storage":{}}}}`
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()))

	var accounts map[string]extension.AccountWithMetadata
	if err := json.Unmarshal([]byte(input), &accounts); err != nil {
		t.Errorf("error when unmarshalling static data: %s", err.Error())
	}

	success := setState(statedb, accounts, &state.PrivacyMetadata{})
	if !success {
		t.Errorf("unexpected error when setting state")
	}

	return statedb
}

func TestStateSetWithListedAccounts(t *testing.T) {
	statedb := createStateDb(t)

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

	success := setState(statedb, accounts, &state.PrivacyMetadata{})
	if success {
		t.Errorf("error expected when setting state")
	}
}

func Test_setPrivacyMetadata(t *testing.T) {
	statedb := createStateDb(t)
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	// call setPrivacyMetaData
	arbitraryBytes1 := []byte{10}
	hash := common.BytesToEncryptedPayloadHash(arbitraryBytes1)
	setPrivacyMetadata(statedb, address, base64.StdEncoding.EncodeToString(arbitraryBytes1))

	privacyMetaData, err := statedb.GetStatePrivacyMetadata(address)
	if err != nil {
		t.Errorf("expected error to be nil, got err %s", err)
	}

	assert.NotEqual(t, privacyMetaData.CreationTxHash, hash)
	privacyMetaData = &state.PrivacyMetadata{hash, engine.PrivacyFlagPartyProtection}
	statedb.SetStatePrivacyMetadata(address, privacyMetaData)

	privacyMetaData, err = statedb.GetStatePrivacyMetadata(address)
	if err != nil {
		t.Errorf("expected error to be nil, got err %s", err)
	}
	assert.Equal(t, engine.PrivacyFlagPartyProtection, privacyMetaData.PrivacyFlag)
	assert.Equal(t, hash, privacyMetaData.CreationTxHash)

	arbitraryBytes2 := []byte{20}
	newHash := common.BytesToEncryptedPayloadHash(arbitraryBytes2)
	setPrivacyMetadata(statedb, address, base64.StdEncoding.EncodeToString(arbitraryBytes2))

	privacyMetaData, err = statedb.GetStatePrivacyMetadata(address)
	if err != nil {
		t.Errorf("expected error to be nil, got err %s", err)
	}
	assert.Equal(t, engine.PrivacyFlagPartyProtection, privacyMetaData.PrivacyFlag)
	assert.Equal(t, newHash, privacyMetaData.CreationTxHash)
}
