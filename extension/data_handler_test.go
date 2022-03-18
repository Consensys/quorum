package extension

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestWriteContentsToFileWritesOkay(t *testing.T) {
	extensionContracts := make(map[common.Address]*ExtensionContract)
	extensionContracts[common.HexToAddress("0x2222222222222222222222222222222222222222")] = &ExtensionContract{
		ContractExtended:          common.HexToAddress("0x1111111111111111111111111111111111111111"),
		Initiator:                 common.HexToAddress("0x3333333333333333333333333333333333333333"),
		Recipient:                 common.HexToAddress("0x4444444444444444444444444444444444444444"),
		RecipientPtmKey:           "1234567891234567891234567891234567891234567=",
		ManagementContractAddress: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		CreationData:              []byte("Sample Transaction Data"),
	}
	psiExtensions := map[types.PrivateStateIdentifier]map[common.Address]*ExtensionContract{
		types.DefaultPrivateStateIdentifier: extensionContracts,
		"somekey":                           extensionContracts,
	}

	datadir, err := ioutil.TempDir("", t.Name())
	defer os.RemoveAll(datadir)
	assert.Nil(t, err, "could not create temp directory for test")

	dataHandler := NewJsonFileDataHandler(datadir)

	err = dataHandler.Save(psiExtensions)
	assert.Nil(t, err, "error writing data from file")

	loadedData, err := dataHandler.Load()
	assert.Nil(t, err, "error reading data from file")

	if !assert.ObjectsAreEqual(psiExtensions, loadedData) {
		expected, _ := json.Marshal(extensionContracts)
		actual, _ := json.Marshal(loadedData)
		t.Errorf("expected data from file different to data written, expected %v, got %v", string(expected), string(actual))
	}
}

func TestLoadOldContents(t *testing.T) {
	extensionContracts := make(map[common.Address]*ExtensionContract)
	extensionContracts[common.HexToAddress("0x2222222222222222222222222222222222222222")] = &ExtensionContract{
		ContractExtended:          common.HexToAddress("0x1111111111111111111111111111111111111111"),
		Initiator:                 common.HexToAddress("0x3333333333333333333333333333333333333333"),
		Recipient:                 common.HexToAddress("0x4444444444444444444444444444444444444444"),
		RecipientPtmKey:           "1234567891234567891234567891234567891234567=",
		ManagementContractAddress: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		CreationData:              []byte("Sample Transaction Data"),
	}
	psiExtensions := map[types.PrivateStateIdentifier]map[common.Address]*ExtensionContract{
		types.DefaultPrivateStateIdentifier: extensionContracts,
		"somekey":                           extensionContracts,
	}

	datadir, err := ioutil.TempDir("", t.Name())
	defer os.RemoveAll(datadir)
	assert.Nil(t, err, "could not create temp directory for test")

	dataHandler := NewJsonFileDataHandler(datadir)

	err = dataHandler.Save(psiExtensions)
	assert.Nil(t, err, "error writing data from file")

	loadedData, err := dataHandler.Load()
	assert.Nil(t, err, "error reading data from file")

	if !assert.ObjectsAreEqual(psiExtensions, loadedData) {
		expected, _ := json.Marshal(extensionContracts)
		actual, _ := json.Marshal(loadedData)
		t.Errorf("expected data from file different to data written, expected %v, got %v", string(expected), string(actual))
	}
}
