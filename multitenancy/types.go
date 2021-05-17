package multitenancy

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	// SchemePSI represents an URL scheme for access scope value
	SchemePSI = "psi"
	// QueryNodeEOA query parameter captures the node-manged EOA address in the URL-based access scope
	QueryNodeEOA = "node.eoa"
	// QuerySelfEOA query parameter captures the self-manged EOA address in the URL-based access scope
	QuerySelfEOA = "self.eoa"
	// AnyEOAAddress represents wild card for EOA address
	AnyEOAAddress = "0x0"
)

// PrivateStateSecurityAttribute contains security configuration ask
// which are defined for a secure private state
type PrivateStateSecurityAttribute struct {
	psi types.PrivateStateIdentifier
	// the node-managed Externally Owned Account being used to sign transactions
	// impacting the private state
	nodeEOA *common.Address
	// the self-managed Externally Owned Account being used to sign transactions
	// impacting the private state
	selfEOA *common.Address
}

func (pssa *PrivateStateSecurityAttribute) String() string {
	return fmt.Sprintf("psi=%s node.eoa=%s self.eoa=%s", pssa.psi, toHexAddress(pssa.nodeEOA), toHexAddress(pssa.selfEOA))
}

func (pssa *PrivateStateSecurityAttribute) WithPSI(psi types.PrivateStateIdentifier) *PrivateStateSecurityAttribute {
	pssa.psi = psi
	return pssa
}

// WithSelfEOAIf calls WithSelfEOA if b is true, otherwise calls WithNodeEOA
func (pssa *PrivateStateSecurityAttribute) WithSelfEOAIf(b bool, eoa common.Address) *PrivateStateSecurityAttribute {
	if b {
		return pssa.WithSelfEOA(eoa)
	} else {
		return pssa.WithNodeEOA(eoa)
	}
}

// WithNodeEOA set node-mannaged EOA value and unset self-managed EOA value
func (pssa *PrivateStateSecurityAttribute) WithNodeEOA(eoa common.Address) *PrivateStateSecurityAttribute {
	pssa.nodeEOA, pssa.selfEOA = &eoa, nil
	return pssa
}

// WithSelfEOA set self-mannaged EOA value and unset node-managed EOA value
func (pssa *PrivateStateSecurityAttribute) WithSelfEOA(eoa common.Address) *PrivateStateSecurityAttribute {
	pssa.selfEOA, pssa.nodeEOA = &eoa, nil
	return pssa
}
