package extensionContracts

import "github.com/ethereum/go-ethereum/common"

func UnpackStateSharedLog(logData []byte) (common.Address, string, string, error) {
	decodedLog := new(ContractExtenderStateShared)
	if err := ContractExtensionABI.Unpack(decodedLog, "StateShared", logData); err != nil {
		return common.Address{}, "", "", err
	}
	return decodedLog.ToExtend, decodedLog.Tesserahash, decodedLog.Uuid, nil
}
