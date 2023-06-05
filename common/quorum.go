package common

const (
	//Hex-encoded 64 byte array of "17" values
	MaxPrivateIntrinsicDataHex = "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
)

// Set contract address, using value that doesn't conflict with upstream geth, or with Besu
func QuorumPrivacyPrecompileContractAddress() Address {
	return BytesToAddress([]byte{byte(0x7a)})
}
