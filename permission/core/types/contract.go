package types

import (
	"crypto/ecdsa"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/permission/core"
)

// TxArgs holds arguments required for execute functions
type TxArgs struct {
	OrgId      string
	POrgId     string
	Url        string
	RoleId     string
	IsVoter    bool
	IsAdmin    bool
	AcctId     common.Address
	AccessType uint8
	Action     uint8
	Txa        ethapi.SendTxArgs
}

type ContractBackend struct {
	EthClnt    bind.ContractBackend
	Key        *ecdsa.PrivateKey
	PermConfig *PermissionConfig
	IsRaft     bool
	UseDns     bool
}

type RoleService interface {
	AddNewRole(_args TxArgs) (*types.Transaction, error)
	RemoveRole(_args TxArgs) (*types.Transaction, error)
}

// Org services
type OrgService interface {
	AddOrg(_args TxArgs) (*types.Transaction, error)
	AddSubOrg(_args TxArgs) (*types.Transaction, error)
	ApproveOrg(_args TxArgs) (*types.Transaction, error)
	UpdateOrgStatus(_args TxArgs) (*types.Transaction, error)
	ApproveOrgStatus(_args TxArgs) (*types.Transaction, error)
}

// Node services
type NodeService interface {
	AddNode(_args TxArgs) (*types.Transaction, error)
	UpdateNodeStatus(_args TxArgs) (*types.Transaction, error)
	StartBlacklistedNodeRecovery(_args TxArgs) (*types.Transaction, error)
	ApproveBlacklistedNodeRecovery(_args TxArgs) (*types.Transaction, error)
}

// Account services
type AccountService interface {
	AssignAccountRole(_args TxArgs) (*types.Transaction, error)
	AssignAdminRole(_args TxArgs) (*types.Transaction, error)
	ApproveAdminRole(_args TxArgs) (*types.Transaction, error)
	UpdateAccountStatus(_args TxArgs) (*types.Transaction, error)
	StartBlacklistedAccountRecovery(_args TxArgs) (*types.Transaction, error)
	ApproveBlacklistedAccountRecovery(_args TxArgs) (*types.Transaction, error)
}

// Control services
type ControlService interface {
	ConnectionAllowed(_enodeId, _ip string, _port, _raftPort uint16) (bool, error)
	TransactionAllowed(_sender common.Address, _target common.Address, _value *big.Int, _gasPrice *big.Int, _gasLimit *big.Int, _payload []byte, _transactionType core.TransactionType) error
}

// Audit services
type AuditService interface {
	ValidatePendingOp(authOrg, orgId, url string, account common.Address, pendingOp int64) bool
	CheckPendingOp(_orgId string) bool
}

type InitService interface {
	BindContracts() error
	Init(_breadth *big.Int, _depth *big.Int) (*types.Transaction, error)
	UpdateNetworkBootStatus() (*types.Transaction, error)
	SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error)
	GetNetworkBootStatus() (bool, error)

	AddAdminAccount(_acct common.Address) (*types.Transaction, error)
	AddAdminNode(url string) (*types.Transaction, error)
	GetAccountDetailsFromIndex(_aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error)
	GetNumberOfAccounts() (*big.Int, error)
	GetAccountDetails(_account common.Address) (common.Address, string, string, *big.Int, bool, error)

	GetRoleDetailsFromIndex(_rIndex *big.Int) (struct {
		RoleId     string
		OrgId      string
		AccessType *big.Int
		Voter      bool
		Admin      bool
		Active     bool
	}, error)
	GetNumberOfRoles() (*big.Int, error)
	GetRoleDetails(_roleId string, _orgId string) (struct {
		RoleId     string
		OrgId      string
		AccessType *big.Int
		Voter      bool
		Admin      bool
		Active     bool
	}, error)

	GetNumberOfOrgs() (*big.Int, error)
	GetSubOrgIndexes(_orgId string) ([]*big.Int, error)
	GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error)
	GetOrgDetails(_orgId string) (string, string, string, *big.Int, *big.Int, error)

	GetNodeDetailsFromIndex(_nodeIndex *big.Int) (string, string, *big.Int, error)
	GetNumberOfNodes() (*big.Int, error)
	GetNodeDetails(enodeId string) (string, string, *big.Int, error)
}

func BindContract(contractInstance interface{}, bindFunc func() (interface{}, error)) error {
	element := reflect.ValueOf(contractInstance).Elem()
	instance, err := bindFunc()
	if err != nil {
		return err
	}
	element.Set(reflect.ValueOf(instance))
	return nil
}
