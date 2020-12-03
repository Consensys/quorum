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

type ContractSecurityAttributeBuilder struct {
	secAttr ContractSecurityAttribute
}

func NewContractSecurityAttributeBuilder() *ContractSecurityAttributeBuilder {
	return &ContractSecurityAttributeBuilder{
		secAttr: ContractSecurityAttribute{
			AccountStateSecurityAttribute: &AccountStateSecurityAttribute{},
			Parties:                       make([]string, 0),
		},
	}
}

func (csab *ContractSecurityAttributeBuilder) FromEOA(eoa common.Address) *ContractSecurityAttributeBuilder {
	csab.secAttr.AccountStateSecurityAttribute.From = eoa
	return csab
}

// ethereum account destination
func (csab *ContractSecurityAttributeBuilder) ToEOA(eoa common.Address) *ContractSecurityAttributeBuilder {
	csab.secAttr.AccountStateSecurityAttribute.To = eoa
	return csab
}

func (csab *ContractSecurityAttributeBuilder) PrivateFrom(tmPubKey string) *ContractSecurityAttributeBuilder {
	csab.secAttr.PrivateFrom = tmPubKey
	return csab
}

// set privateFrom only if b is true, ignore otherwise
func (csab *ContractSecurityAttributeBuilder) PrivateFromOnlyIf(b bool, tmPubKey string) *ContractSecurityAttributeBuilder {
	if b {
		csab.secAttr.PrivateFrom = tmPubKey
	}
	return csab
}

func (csab *ContractSecurityAttributeBuilder) Visibility(v ContractVisibility) *ContractSecurityAttributeBuilder {
	csab.secAttr.Visibility = v
	return csab
}

func (csab *ContractSecurityAttributeBuilder) Private() *ContractSecurityAttributeBuilder {
	return csab.Visibility(VisibilityPrivate)
}

// set VisibilityPrivate if b is true, VisibilityPublic otherwise
func (csab *ContractSecurityAttributeBuilder) PrivateIf(b bool) *ContractSecurityAttributeBuilder {
	if b {
		return csab.Visibility(VisibilityPrivate)
	} else {
		return csab.Visibility(VisibilityPublic)
	}
}

func (csab *ContractSecurityAttributeBuilder) Public() *ContractSecurityAttributeBuilder {
	return csab.Visibility(VisibilityPublic)
}

func (csab *ContractSecurityAttributeBuilder) Action(a ContractAction) *ContractSecurityAttributeBuilder {
	csab.secAttr.Action = a
	return csab
}

func (csab *ContractSecurityAttributeBuilder) Create() *ContractSecurityAttributeBuilder {
	return csab.Action(ActionCreate)
}

func (csab *ContractSecurityAttributeBuilder) Read() *ContractSecurityAttributeBuilder {
	return csab.Action(ActionRead)
}

func (csab *ContractSecurityAttributeBuilder) Write() *ContractSecurityAttributeBuilder {
	return csab.Action(ActionWrite)
}

// set ActionRead only if b is true, ignore otherwise
func (csab *ContractSecurityAttributeBuilder) ReadOnlyIf(b bool) *ContractSecurityAttributeBuilder {
	if b {
		return csab.Action(ActionRead)
	} else {
		return csab
	}
}

// set ActionWrite only if b is true, ignore otherwise
func (csab *ContractSecurityAttributeBuilder) WriteOnlyIf(b bool) *ContractSecurityAttributeBuilder {
	if b {
		return csab.Action(ActionWrite)
	} else {
		return csab
	}
}

func (csab *ContractSecurityAttributeBuilder) Parties(tmPubKeys []string) *ContractSecurityAttributeBuilder {
	parties := make([]string, len(tmPubKeys))
	copy(parties, tmPubKeys)
	csab.secAttr.Parties = parties
	return csab
}

func (csab *ContractSecurityAttributeBuilder) Party(tmPubKey string) *ContractSecurityAttributeBuilder {
	csab.secAttr.Parties = append(csab.secAttr.Parties, tmPubKey)
	return csab
}

func (csab *ContractSecurityAttributeBuilder) Build() *ContractSecurityAttribute {
	return &csab.secAttr
}
