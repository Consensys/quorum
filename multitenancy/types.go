package multitenancy

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// SchemePSI represents an URL scheme for access scope value
	SchemePSI = "psi"
	// QueryEOA query parameter captures the EOA address in the URL-based access scope
	QueryEOA = "eoa"
	// AnyEOAAddress represents wild card for EOA address
	AnyEOAAddress = "0x0"
)

// PrivateStateSecurityAttribute contains security configuration ask
// which are defined for a secure private state
type PrivateStateSecurityAttribute struct {
	psi types.PrivateStateIdentifier
	// the Externally Owned Account being used to sign transactions
	// impacting the private state
	eoa common.Address
}

func (pssa *PrivateStateSecurityAttribute) String() string {
	return fmt.Sprintf("psi=%s eoa=%s", pssa.psi, pssa.eoa.Hex())
}

func (pssa *PrivateStateSecurityAttribute) WithPSI(psi types.PrivateStateIdentifier) *PrivateStateSecurityAttribute {
	pssa.psi = psi
	return pssa
}

func (pssa *PrivateStateSecurityAttribute) WithEOA(eoa common.Address) *PrivateStateSecurityAttribute {
	pssa.eoa = eoa
	return pssa
}
