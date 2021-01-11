package extension

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

/*
	The file can have two formats:

	1.
	{
		"psiContracts": {
			"psi1": {
				"contract1address": ...,
				"contract2address": ...,
			},
			...
		}
	}

	2.
	{
		"contract1address": ...,
		"contract2address": ...,
	}

*/

const extensionContractData = "activeExtensions.json"

type DataHandler interface {
	Load() (map[string]map[common.Address]*ExtensionContract, error)

	Save(extensionContracts map[string]map[common.Address]*ExtensionContract) error
}

type JsonFileDataHandler struct {
	saveFile string
}

func NewJsonFileDataHandler(dataDirectory string) *JsonFileDataHandler {
	return &JsonFileDataHandler{
		saveFile: filepath.Join(dataDirectory, extensionContractData),
	}
}

/*
	The stratehy when loading the save file is too check if the newer "psiContracts" field is present.
	If so, then everything should exist under that key, and so we can unmarshal and return immediately.

	If not, then the save file was made from a previous version. Load up all the data as before and
	put it under the "private" PSI.

	It should never be the case the file contains both types of data at once.
*/
func (handler *JsonFileDataHandler) Load() (map[string]map[common.Address]*ExtensionContract, error) {
	if _, err := os.Stat(handler.saveFile); !(err == nil || !os.IsNotExist(err)) {
		return map[string]map[common.Address]*ExtensionContract{"private": {}}, nil
	}

	blob, err := ioutil.ReadFile(handler.saveFile)
	if err != nil {
		return nil, err
	}

	var untyped map[string]json.RawMessage
	if err := json.Unmarshal(blob, &untyped); err != nil {
		return nil, err
	}

	if psiContracts, ok := untyped["psiContracts"]; ok {
		var contracts map[string]map[common.Address]*ExtensionContract
		json.Unmarshal(psiContracts, &contracts)
		return contracts, nil
	}

	currentContracts := make(map[common.Address]*ExtensionContract)
	for key, val := range untyped {
		extAddress := common.HexToAddress(key)
		var ext ExtensionContract
		json.Unmarshal(val, &ext)
		currentContracts[extAddress] = &ext
	}
	return map[string]map[common.Address]*ExtensionContract{"private": currentContracts}, nil
}

func (handler *JsonFileDataHandler) Save(extensionContracts map[string]map[common.Address]*ExtensionContract) error {
	//we want to put the map under "psiContracts" key to distinguish from existing data
	saveData := make(map[string]interface{})
	saveData["psiContracts"] = extensionContracts

	//no unmarshallable types, so can't error
	output, _ := json.Marshal(&saveData)

	if errSaving := ioutil.WriteFile(handler.saveFile, output, 0644); errSaving != nil {
		log.Error("Couldn't save outstanding extension contract details")
		return errSaving
	}
	return nil
}
