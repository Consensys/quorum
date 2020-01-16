package extension

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

const extensionContractData = "activeExtensions.json"

type DataHandler interface {
	Load() (map[common.Address]*ExtensionContract, error)

	Save(extensionContracts map[common.Address]*ExtensionContract) error
}

type JsonFileDataHandler struct {
	saveFile string
}

func NewJsonFileDataHandler(dataDirectory string) *JsonFileDataHandler {
	return &JsonFileDataHandler{
		saveFile: filepath.Join(dataDirectory, extensionContractData),
	}
}

func (handler *JsonFileDataHandler) Load() (map[common.Address]*ExtensionContract, error) {
	currentContracts := make(map[common.Address]*ExtensionContract)
	if _, err := os.Stat(handler.saveFile); err == nil || !os.IsNotExist(err) {
		blob, err := ioutil.ReadFile(handler.saveFile)
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(blob, &currentContracts); err != nil {
			return nil, err
		}
	}
	return currentContracts, nil
}

func (handler *JsonFileDataHandler) Save(extensionContracts map[common.Address]*ExtensionContract) error {
	//no unmarshallable types, so can't error
	output, _ := json.Marshal(&extensionContracts)

	if errSaving := ioutil.WriteFile(handler.saveFile, output, 0644); errSaving != nil {
		log.Error("Couldn't save outstanding extension contract details")
		return errSaving
	}
	return nil
}
