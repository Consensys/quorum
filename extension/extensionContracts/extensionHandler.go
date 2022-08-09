package extensionContracts

import "github.com/ethereum/go-ethereum/common"

func UnpackStateSharedLog(logData []byte) (common.Address, string, string, error) {
	decodedLog := new(ContractExtenderStateShared)
	if err := ContractExtenderParsedABI.UnpackIntoInterface(decodedLog, "StateShared", logData); err != nil {
		return common.Address{}, "", "", err
	}
	return decodedLog.ToExtend, decodedLog.Tesserahash, decodedLog.Uuid, nil
}

func UnpackNewExtensionCreatedLog(data []byte) (*ContractExtenderNewContractExtensionContractCreated, error) {
	newExtensionEvent := new(ContractExtenderNewContractExtensionContractCreated)
	err := ContractExtenderParsedABI.UnpackIntoInterface(newExtensionEvent, "NewContractExtensionContractCreated", data)

	return newExtensionEvent, err
}
