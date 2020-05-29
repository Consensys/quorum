package extension

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

	datadir, err := ioutil.TempDir("", t.Name())
	if err != nil {
		t.Errorf("could not create temp directory for test, error: %s", err.Error())
	}

	dataHandler := NewJsonFileDataHandler(datadir)

	if err := dataHandler.Save(extensionContracts); err != nil {
		t.Errorf("error writing data to file, error: %s", err.Error())
	}

	loadedData, err := dataHandler.Load()
	if err != nil {
		t.Errorf("error reading data from file, error: %s", err.Error())
	}

	if !assert.ObjectsAreEqual(extensionContracts, loadedData) {
		expected, _ := json.Marshal(extensionContracts)
		actual, _ := json.Marshal(loadedData)
		t.Errorf("expected data from file different to data written, expected %v, got %v", string(expected), string(actual))
	}
}
