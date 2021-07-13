package common

import "math"

// Set contract address, using MaxInt8 as the last address that can be used for a quorum pre-compiled contract
func QuorumPrivacyPrecompileContractAddress() Address {

	return BytesToAddress([]byte{byte(math.MaxInt8 - 1)})

}
