package privacyExtension

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	extension "github.com/ethereum/go-ethereum/extension/extensionContracts"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/private/engine/notinuse"
	"github.com/stretchr/testify/assert"
)

var input = `{"0x2222222222222222222222222222222222222222":
				{"state":
					{"balance":"22",
						"nonce":5,
						"root":"56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
						"codeHash":"87874902497a5bb968da31a2998d8f22e949d1ef6214bcdedd8bae24cca4b9e3",
						"code":"03030303030303",
						"storage":{
							"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421": "2a"
						}
				}}}`

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

func createStateDb(t *testing.T, metadata *state.PrivacyMetadata) *state.StateDB {
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)

	var accounts map[string]extension.AccountWithMetadata
	if err := json.Unmarshal([]byte(input), &accounts); err != nil {
		t.Errorf("error when unmarshalling static data: %s", err.Error())
	}

	success := setState(statedb, accounts, metadata, nil)
	if !success {
		t.Errorf("unexpected error when setting state")
	}

	return statedb
}

func TestStateSetWithListedAccounts(t *testing.T) {
	statedb := createStateDb(t, &state.PrivacyMetadata{})

	address := common.HexToAddress("0x2222222222222222222222222222222222222222")
	balance := statedb.GetBalance(address)
	code := statedb.GetCode(address)
	nonce := statedb.GetNonce(address)
	storage, _ := statedb.GetStorageRoot(address)

	// we don't save PrivacyMetadata if it's standardprivate
	privacyMetaData, err := statedb.GetPrivacyMetadata(address)
	assert.Error(t, err, common.ErrNoAccountExtraData)
	assert.Nil(t, privacyMetaData)

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

	stateVal := statedb.GetState(address, common.BytesToHash(expectedStorageHash))
	assert.Equal(t, common.HexToHash("0x2a"), stateVal)
}

func TestStateSetWithListedAccountsFailsOnInvalidBalance(t *testing.T) {
	input := `{"0x2222222222222222222222222222222222222222":{"state":{"balance":"invalid","nonce":5,"root":"56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","codeHash":"87874902497a5bb968da31a2998d8f22e949d1ef6214bcdedd8bae24cca4b9e3","code":"03030303030303","storage":{}}}}`
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)

	var accounts map[string]extension.AccountWithMetadata
	if err := json.Unmarshal([]byte(input), &accounts); err != nil {
		t.Errorf("error when unmarshalling static data: %s", err.Error())
	}

	success := setState(statedb, accounts, &state.PrivacyMetadata{}, nil)
	if success {
		t.Errorf("error expected when setting state")
	}
}

func Test_setPrivacyMetadata(t *testing.T) {
	privacyMetaData := &state.PrivacyMetadata{}

	statedb := createStateDb(t, privacyMetaData)
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	// call setPrivacyMetaData
	arbitraryBytes1 := []byte{10}
	hash := common.BytesToEncryptedPayloadHash(arbitraryBytes1)
	setPrivacyMetadata(statedb, address, base64.StdEncoding.EncodeToString(arbitraryBytes1))

	// we don't save PrivacyMetadata if it's standardprivate
	_, err := statedb.GetPrivacyMetadata(address)
	assert.Error(t, err, common.ErrNoAccountExtraData)

	privacyMetaData = &state.PrivacyMetadata{CreationTxHash: hash, PrivacyFlag: engine.PrivacyFlagPartyProtection}
	statedb.SetPrivacyMetadata(address, privacyMetaData)

	privacyMetaData, err = statedb.GetPrivacyMetadata(address)
	if err != nil {
		t.Errorf("expected error to be nil, got err %s", err)
	}
	assert.Equal(t, engine.PrivacyFlagPartyProtection, privacyMetaData.PrivacyFlag)
	assert.Equal(t, hash, privacyMetaData.CreationTxHash)

	arbitraryBytes2 := []byte{20}
	newHash := common.BytesToEncryptedPayloadHash(arbitraryBytes2)
	setPrivacyMetadata(statedb, address, base64.StdEncoding.EncodeToString(arbitraryBytes2))

	privacyMetaData, err = statedb.GetPrivacyMetadata(address)
	if err != nil {
		t.Errorf("expected error to be nil, got err %s", err)
	}
	assert.Equal(t, engine.PrivacyFlagPartyProtection, privacyMetaData.PrivacyFlag)
	assert.Equal(t, newHash, privacyMetaData.CreationTxHash)
}

func Test_setState_WithManagedParties(t *testing.T) {
	statedb := createStateDb(t, &state.PrivacyMetadata{})
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	presetManagedParties := []string{"mp1", "mp2"}
	statedb.SetManagedParties(address, presetManagedParties)

	mp, err := statedb.GetManagedParties(address)
	assert.Nil(t, err)
	assert.EqualValues(t, presetManagedParties, mp)

	extraManagedParties := []string{"mp1", "mp2", "mp3"}
	var accounts map[string]extension.AccountWithMetadata
	json.Unmarshal([]byte(input), &accounts)
	success := setState(statedb, accounts, &state.PrivacyMetadata{}, extraManagedParties)
	assert.True(t, success)

	mp, err = statedb.GetManagedParties(address)
	assert.Nil(t, err)
	assert.EqualValues(t, []string{"mp1", "mp2", "mp3"}, mp)
}

func Test_validateAccountsExist_AllPresent(t *testing.T) {
	expected := []common.Address{
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
	}
	actual := map[string]extension.AccountWithMetadata{
		"0x2222222222222222222222222222222222222222": {},
		"0x3333333333333333333333333333333333333333": {},
	}

	equal := validateAccountsExist(expected, actual)

	assert.True(t, equal)
}

func Test_validateAccountsExist_NotAllPresent(t *testing.T) {
	expected := []common.Address{
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
	}
	actual := map[string]extension.AccountWithMetadata{
		"0x2222222222222222222222222222222222222222": {},
		"0x4444444444444444444444444444444444444444": {},
	}

	equal := validateAccountsExist(expected, actual)

	assert.False(t, equal)
}

func Test_setManagedParties(t *testing.T) {
	statedb := createStateDb(t, &state.PrivacyMetadata{})
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	presetManagedParties := []string{"mp1", "mp2"}
	statedb.SetManagedParties(address, presetManagedParties)

	mp, err := statedb.GetManagedParties(address)
	assert.Nil(t, err)
	assert.EqualValues(t, presetManagedParties, mp)

	extraManagedParties := []string{"mp1", "mp3"}
	mpm := &mockPrivateTransactionManager{
		returns: map[string][]interface{}{"Receive": {"", extraManagedParties, nil, nil, nil}},
	}

	ptmHash := common.EncryptedPayloadHash{86}.ToBase64()
	setManagedParties(mpm, statedb, address, ptmHash)

	mp, err = statedb.GetManagedParties(address)
	assert.Nil(t, err)
	assert.Len(t, mp, 3)
	assert.Contains(t, mp, "mp1")
	assert.Contains(t, mp, "mp2")
	assert.Contains(t, mp, "mp3")
}

func Test_setManagedPartiesInvalidHash(t *testing.T) {
	statedb := createStateDb(t, &state.PrivacyMetadata{})
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	presetManagedParties := []string{"mp1", "mp2"}
	statedb.SetManagedParties(address, presetManagedParties)

	mp, err := statedb.GetManagedParties(address)
	assert.Nil(t, err)
	assert.EqualValues(t, presetManagedParties, mp)

	extraManagedParties := []string{"mp1", "mp3"}
	mpm := &mockPrivateTransactionManager{
		returns: map[string][]interface{}{"Receive": {"", extraManagedParties, nil, nil, nil}},
	}

	ptmHash := common.EncryptedPayloadHash{86}.Hex() //should be base64, so hex will fail
	setManagedParties(mpm, statedb, address, ptmHash)

	mp, err = statedb.GetManagedParties(address)
	assert.Nil(t, err)
	assert.EqualValues(t, presetManagedParties, mp)
}

type mockPSMR struct {
	core.DefaultPrivateStateManager
	returns map[string][]interface{}
}

type mockPrivateTransactionManager struct {
	notinuse.PrivateTransactionManager
	returns map[string][]interface{}
}

func (mpsmr *mockPSMR) ResolveForManagedParty(managedParty string) (*types.PrivateStateMetadata, error) {
	values := mpsmr.returns["ResolveForManagedParty"]
	var (
		r1 *types.PrivateStateMetadata
		r2 error
	)
	if values[0] != nil {
		r1 = values[0].(*types.PrivateStateMetadata)
	}
	if values[1] != nil {
		r2 = values[1].(error)
	}
	return r1, r2
}

func (mpm *mockPrivateTransactionManager) Receive(data common.EncryptedPayloadHash) (string, []string, []byte, *engine.ExtraMetadata, error) {
	values := mpm.returns["Receive"]
	var (
		r1 string
		r2 []string
		r3 []byte
		r4 *engine.ExtraMetadata
		r5 error
	)
	if values[0] != nil {
		r1 = values[0].(string)
	}
	if values[1] != nil {
		r2 = values[1].([]string)
	}
	if values[2] != nil {
		r3 = values[2].([]byte)
	}
	if values[3] != nil {
		r4 = values[3].(*engine.ExtraMetadata)
	}
	if values[4] != nil {
		r5 = values[4].(error)
	}
	return r1, r2, r3, r4, r5
}

func (mpm *mockPrivateTransactionManager) IsSender(txHash common.EncryptedPayloadHash) (bool, error) {
	values := mpm.returns["IsSender"]
	var (
		r1 bool
		r2 error
	)
	if values[0] != nil {
		r1 = values[0].(bool)
	}
	if values[1] != nil {
		r2 = values[1].(error)
	}
	return r1, r2
}

func (mpm *mockPrivateTransactionManager) DecryptPayload(payload common.DecryptRequest) ([]byte, *engine.ExtraMetadata, error) {
	values := mpm.returns["DecryptPayload"]
	var (
		r3 []byte
		r4 *engine.ExtraMetadata
		r5 error
	)
	if values[0] != nil {
		r3 = values[0].([]byte)
	}
	if values[1] != nil {
		r4 = values[1].(*engine.ExtraMetadata)
	}
	if values[2] != nil {
		r5 = values[2].(error)
	}
	return r3, r4, r5
}

func (mpm *mockPrivateTransactionManager) GetCache() state.Database {
	return nil
}
