package privacyExtension

import (
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestExtensionHandler_CheckExtensionAndSetPrivateState_NoLogs(t *testing.T) {
	ptm := &mockPrivateTransactionManager{}
	handler := NewExtensionHandler(ptm)
	statedb := createStateDb(t, &state.PrivacyMetadata{})

	dbBefore := statedb.Copy()
	rootBeforeExtension, _ := statedb.Commit(true)

	handler.CheckExtensionAndSetPrivateState(nil, statedb, types.DefaultPrivateStateIdentifier)

	rootAfterExtension, _ := statedb.Commit(true)
	assert.Equal(t, rootBeforeExtension, rootAfterExtension)

	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	beforeManagedParties, _ := dbBefore.GetManagedParties(address)
	afterManagedParties, _ := statedb.GetManagedParties(address)
	assert.Equal(t, beforeManagedParties, afterManagedParties)

	beforePrivacyMetadata, _ := dbBefore.GetPrivacyMetadata(address)
	afterPrivacyMetadata, _ := statedb.GetPrivacyMetadata(address)
	assert.Equal(t, beforePrivacyMetadata, afterPrivacyMetadata)
}

func TestExtensionHandler_CheckExtensionAndSetPrivateState_LogsAreNotExtensionLogs(t *testing.T) {
	ptm := &mockPrivateTransactionManager{}
	handler := NewExtensionHandler(ptm)
	statedb := createStateDb(t, &state.PrivacyMetadata{})

	dbBefore := statedb.Copy()
	rootBeforeExtension, _ := statedb.Commit(true)

	notExtensionLogs := []*types.Log{
		{
			Address:     common.HexToAddress("0x9ccd1e1089c79fe1cca81601fc9ccfa24f77eb58"),
			Topics:      []common.Hash{common.HexToHash("0x24ec1d3ff24c2f6ff210738839dbc339cd45a5294d85c79361016243157aae7b")},
			Data:        []byte{},
			BlockNumber: 6,
			TxHash:      common.HexToHash("0x5faf9ffe6fedc1139bdc1af20b26a2e113d16d736be872571458f8d4bcc048c7"),
			TxIndex:     0,
			BlockHash:   common.HexToHash("0x7e7fb6985ff7e1c7293b8e3202a2b101458acd0b93b5fbed18aab40e8cbeb587"),
			Index:       0,
			Removed:     false,
			PSI:         types.DefaultPrivateStateIdentifier,
		},
	}
	handler.CheckExtensionAndSetPrivateState(notExtensionLogs, statedb, types.DefaultPrivateStateIdentifier)

	rootAfterExtension, _ := statedb.Commit(true)
	assert.Equal(t, rootBeforeExtension, rootAfterExtension)

	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	beforeManagedParties, _ := dbBefore.GetManagedParties(address)
	afterManagedParties, _ := statedb.GetManagedParties(address)
	assert.Equal(t, beforeManagedParties, afterManagedParties)

	beforePrivacyMetadata, _ := dbBefore.GetPrivacyMetadata(address)
	afterPrivacyMetadata, _ := statedb.GetPrivacyMetadata(address)
	assert.Equal(t, beforePrivacyMetadata, afterPrivacyMetadata)
}

func TestExtensionHandler_UuidIsOwn_EmptyUUID(t *testing.T) {
	ptm := &mockPrivateTransactionManager{}
	handler := NewExtensionHandler(ptm)

	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	isOwn := handler.UuidIsOwn(address, "", types.DefaultPrivateStateIdentifier)

	assert.False(t, isOwn)
}

func TestExtensionHandler_UuidIsOwn_IsSenderIsFalse(t *testing.T) {
	ptm := &mockPrivateTransactionManager{
		returns: map[string][]interface{}{"IsSender": {false, nil}},
	}
	handler := NewExtensionHandler(ptm)

	const uuid = "0xabcd"
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	isOwn := handler.UuidIsOwn(address, uuid, types.DefaultPrivateStateIdentifier)

	assert.False(t, isOwn)
}

func TestExtensionHandler_UuidIsOwn_WrongPSIFails(t *testing.T) {
	ptm := &mockPrivateTransactionManager{
		returns: map[string][]interface{}{
			"IsSender": {true, nil},
			"Receive":  {"psi1", nil, []byte{}, nil, nil},
		},
	}
	handler := NewExtensionHandler(ptm)
	psmr := &mockPSMR{
		returns: map[string][]interface{}{
			"ResolveForManagedParty": {&types.PrivateStateMetadata{ID: "psi1", Type: types.Resident}, nil},
		},
	}
	handler.SetPSMR(psmr)

	uuid := "0xabcd"
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	isOwn := handler.UuidIsOwn(address, uuid, "other")

	assert.False(t, isOwn)
}

func TestExtensionHandler_UuidIsOwn_DecryptPayloadFails(t *testing.T) {
	ptm := &mockPrivateTransactionManager{
		returns: map[string][]interface{}{
			"IsSender":       {true, nil},
			"Receive":        {"psi1", nil, []byte(`{"somedata": "val"}`), nil, nil},
			"DecryptPayload": {nil, nil, errors.New("test error")},
		},
	}
	handler := NewExtensionHandler(ptm)
	psmr := &mockPSMR{
		returns: map[string][]interface{}{
			"ResolveForManagedParty": {&types.PrivateStateMetadata{ID: "psi1", Type: types.Resident}, nil},
		},
	}
	handler.SetPSMR(psmr)

	uuid := "0xabcd"
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	isOwn := handler.UuidIsOwn(address, uuid, "psi1")

	assert.False(t, isOwn)
}

func TestExtensionHandler_UuidIsOwn_AddressDoesntMatch(t *testing.T) {
	ptm := &mockPrivateTransactionManager{
		returns: map[string][]interface{}{
			"IsSender":       {true, nil},
			"Receive":        {"psi1", nil, []byte(`{"somedata": "val"}`), nil, nil},
			"DecryptPayload": {[]byte(`unknown`), nil, nil},
		},
	}
	handler := NewExtensionHandler(ptm)
	psmr := &mockPSMR{
		returns: map[string][]interface{}{
			"ResolveForManagedParty": {&types.PrivateStateMetadata{ID: "psi1", Type: types.Resident}, nil},
		},
	}
	handler.SetPSMR(psmr)

	uuid := "0xabcd"
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	isOwn := handler.UuidIsOwn(address, uuid, "psi1")

	assert.False(t, isOwn)
}

func TestExtensionHandler_UuidIsOwn_AddressMatches(t *testing.T) {
	uuid := "0xabcd"
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	ptm := &mockPrivateTransactionManager{
		returns: map[string][]interface{}{
			"IsSender":       {true, nil},
			"Receive":        {"psi1", nil, []byte(`{"somedata": "val"}`), nil, nil},
			"DecryptPayload": {address.Bytes(), nil, nil},
		},
	}
	handler := NewExtensionHandler(ptm)
	psmr := &mockPSMR{
		returns: map[string][]interface{}{
			"ResolveForManagedParty": {&types.PrivateStateMetadata{ID: "psi1", Type: types.Resident}, nil},
		},
	}
	handler.SetPSMR(psmr)

	isOwn := handler.UuidIsOwn(address, uuid, "psi1")

	assert.True(t, isOwn)
}

func TestExtensionHandler_UuidIsOwn_PrivatePSMRSucceeds(t *testing.T) {
	uuid := "0xabcd"
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	ptm := &mockPrivateTransactionManager{
		returns: map[string][]interface{}{
			"IsSender":       {true, nil},
			"Receive":        {"psi1, private", nil, []byte(`{"somedata": "val"}`), nil, nil},
			"DecryptPayload": {address.Bytes(), nil, nil},
		},
	}
	handler := NewExtensionHandler(ptm)
	psmr := &mockPSMR{
		returns: map[string][]interface{}{
			"ResolveForManagedParty": {&types.PrivateStateMetadata{ID: "private", Type: types.Resident}, nil},
		},
	}
	handler.SetPSMR(psmr)

	isOwn := handler.UuidIsOwn(address, uuid, types.DefaultPrivateStateIdentifier)

	assert.True(t, isOwn)
}
