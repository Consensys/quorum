package multitenancy

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ContractVisibility string
type ContractAction string

const (
	VisibilityPublic  ContractVisibility = "public"
	VisibilityPrivate ContractVisibility = "private"
	ActionRead        ContractAction     = "read"
	ActionWrite       ContractAction     = "write"
	ActionCreate      ContractAction     = "create"

	// QueryOwnedEOA query parameter is to capture the EOA address
	// For value transfer, it represents the account owner
	// For message call, it represents the EOA that signed the contract creation transaction
	// in other words, the EOA that owns the contract
	QueryOwnedEOA = "owned.eoa"
	// QueryToEOA query parameter is to capture the EOA address which is the
	// target account in value transfer scenarios
	QueryToEOA = "to.eoa"
	// QueryFromTM query parameter is to capture the Tessera Public Key
	// which indicates the sender of a private transaction or participant of a private contract
	QueryFromTM = "from.tm"

	// AnyEOAAddress represents wild card for EOA address
	AnyEOAAddress = "0x0"
)

// AccountStateSecurityAttribute contains security configuration ask
// which are defined for a secure account state
type AccountStateSecurityAttribute struct {
	From common.Address // Account Address
	To   common.Address
}

func (assa *AccountStateSecurityAttribute) String() string {
	return fmt.Sprintf("from=%s to=%s", assa.From.Hex(), assa.To.Hex())
}

// ContractSecurityAttribute contains security configuration ask
// which are defined for a secure contract account
type ContractSecurityAttribute struct {
	*AccountStateSecurityAttribute
	Visibility  ContractVisibility // public/private
	Action      ContractAction     // create/read/write
	PrivateFrom string             // TM Key, only if Visibility is private, for write/create
	Parties     []string           // TM Keys, only if Visibility is private, for read
}

func (csa *ContractSecurityAttribute) String() string {
	return fmt.Sprintf("%v visibility=%s action=%s privateFrom=%s parties=%v", csa.AccountStateSecurityAttribute, csa.Visibility, csa.Action, csa.PrivateFrom, csa.Parties)
}

// Construct a list of READ security ask from contract event logs
func ToContractSecurityAttributes(contractIndex ContractIndexReader, logs []*types.Log) ([]*ContractSecurityAttribute, error) {
	attributes := make([]*ContractSecurityAttribute, 0)
	for _, l := range logs {
		attr, err := ToContractSecurityAttribute(contractIndex, l.Address)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, attr)
	}
	return attributes, nil
}

// Construct a READ security attribute for a contract from the index
func ToContractSecurityAttribute(contractIndex ContractIndexReader, contractAddress common.Address) (*ContractSecurityAttribute, error) {
	ci, err := contractIndex.ReadIndex(contractAddress)
	if err != nil {
		return nil, fmt.Errorf("contract %s not found in the index due to %s", contractAddress.Hex(), err.Error())
	}
	attr := &ContractSecurityAttribute{
		AccountStateSecurityAttribute: &AccountStateSecurityAttribute{
			From: ci.CreatorAddress, // TODO must figure out what this value must be when tighten access control for account
			To:   ci.CreatorAddress,
		},
		Action:  ActionRead,
		Parties: ci.Parties,
	}
	if ci.IsPrivate {
		attr.Visibility = VisibilityPrivate
	} else {
		attr.Visibility = VisibilityPublic
	}
	return attr, nil
}
