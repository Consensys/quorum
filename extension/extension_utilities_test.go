package extension

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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
