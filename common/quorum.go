package common

import "math"

var maxPrecompileAddress = math.MaxInt8 // maxPrecompileAddress is the last address that can be used for a pre-compiled contract

func QuorumPrivacyPrecompileContractAddress() Address {
	address := maxPrecompileAddress - 1
	return BytesToAddress([]byte{byte(address)})
}
