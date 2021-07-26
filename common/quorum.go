package common

// Set contract address, using value that doesn't conflict with upstream geth, or with Besu
func QuorumPrivacyPrecompileContractAddress() Address {

	return BytesToAddress([]byte{byte(0x7a)})

}
