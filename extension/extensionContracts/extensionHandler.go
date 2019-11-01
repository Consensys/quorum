package extensionContracts

func UnpackStateSharedLog(logData []byte) (string, string, error) {
	decodedLog := new(ContractExtenderStateShared)
	if err := ContractExtensionABI.Unpack(decodedLog, "StateShared", logData); err != nil {
		return "", "", err
	}
	return decodedLog.Hash, decodedLog.Uuid, nil
}
