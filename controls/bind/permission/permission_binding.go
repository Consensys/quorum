// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permissions

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PermissionsABI is the input ABI used to generate the binding from.
const PermissionsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"getVoteCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getVoteStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVoterCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"activateNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"approveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"acctIndex\",\"type\":\"uint256\"}],\"name\":\"getAccountDetails\",\"outputs\":[{\"name\":\"_acct\",\"type\":\"address\"},{\"name\":\"_acctAccess\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfAccounts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"getNodeStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"enodeId\",\"type\":\"string\"}],\"name\":\"getNodeDetails\",\"outputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_ipAddrPort\",\"type\":\"string\"},{\"name\":\"_discPort\",\"type\":\"string\"},{\"name\":\"_raftPort\",\"type\":\"string\"},{\"name\":\"_nodeStatus\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"deactivateNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateNetworkBootStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"proposeDeactivation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"cancelPendingOperation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"blacklistNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initAccounts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNetworkBootStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_ipAddrPort\",\"type\":\"string\"},{\"name\":\"_discPort\",\"type\":\"string\"},{\"name\":\"_raftPort\",\"type\":\"string\"}],\"name\":\"proposeNodeBlacklisting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeIndex\",\"type\":\"uint256\"}],\"name\":\"getNodeDetailsFromIndex\",\"outputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_ipAddrPort\",\"type\":\"string\"},{\"name\":\"_discPort\",\"type\":\"string\"},{\"name\":\"_raftPort\",\"type\":\"string\"},{\"name\":\"_nodeStatus\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitAccountsCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_acctid\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_accountAccess\",\"type\":\"uint8\"}],\"name\":\"updateAccountAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_ipAddrPort\",\"type\":\"string\"},{\"name\":\"_discPort\",\"type\":\"string\"},{\"name\":\"_raftPort\",\"type\":\"string\"}],\"name\":\"proposeNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getVoter\",\"outputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_voterStatus\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfValidVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"proposeNodeActivation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFullAccessAccountCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_ipAddrPort\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_discPort\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_raftPort\",\"type\":\"string\"}],\"name\":\"NodeApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodePendingDeactivation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_ipAddrPort\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_discPort\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_raftPort\",\"type\":\"string\"}],\"name\":\"NodeDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodePendingActivation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_ipAddrPort\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_discPort\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_raftPort\",\"type\":\"string\"}],\"name\":\"NodeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodePendingBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_ipAddrPort\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_discPort\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_raftPort\",\"type\":\"string\"}],\"name\":\"NodeBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"PendingOperationCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_access\",\"type\":\"uint8\"}],\"name\":\"AccountAccessModified\",\"type\":\"event\"}]"

// Permissions is an auto generated Go binding around an Ethereum contract.
type Permissions struct {
	PermissionsCaller     // Read-only binding to the contract
	PermissionsTransactor // Write-only binding to the contract
	PermissionsFilterer   // Log filterer for contract events
}

// PermissionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionsSession struct {
	Contract     *Permissions      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermissionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionsCallerSession struct {
	Contract *PermissionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PermissionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionsTransactorSession struct {
	Contract     *PermissionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PermissionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionsRaw struct {
	Contract *Permissions // Generic contract binding to access the raw methods on
}

// PermissionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionsCallerRaw struct {
	Contract *PermissionsCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionsTransactorRaw struct {
	Contract *PermissionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissions creates a new instance of Permissions, bound to a specific deployed contract.
func NewPermissions(address common.Address, backend bind.ContractBackend) (*Permissions, error) {
	contract, err := bindPermissions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Permissions{PermissionsCaller: PermissionsCaller{contract: contract}, PermissionsTransactor: PermissionsTransactor{contract: contract}, PermissionsFilterer: PermissionsFilterer{contract: contract}}, nil
}

// NewPermissionsCaller creates a new read-only instance of Permissions, bound to a specific deployed contract.
func NewPermissionsCaller(address common.Address, caller bind.ContractCaller) (*PermissionsCaller, error) {
	contract, err := bindPermissions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionsCaller{contract: contract}, nil
}

// NewPermissionsTransactor creates a new write-only instance of Permissions, bound to a specific deployed contract.
func NewPermissionsTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionsTransactor, error) {
	contract, err := bindPermissions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionsTransactor{contract: contract}, nil
}

// NewPermissionsFilterer creates a new log filterer instance of Permissions, bound to a specific deployed contract.
func NewPermissionsFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionsFilterer, error) {
	contract, err := bindPermissions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionsFilterer{contract: contract}, nil
}

// bindPermissions binds a generic wrapper to an already deployed contract.
func bindPermissions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissions *PermissionsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissions.Contract.PermissionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissions *PermissionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.Contract.PermissionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissions *PermissionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissions.Contract.PermissionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissions *PermissionsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissions *PermissionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissions *PermissionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissions.Contract.contract.Transact(opts, method, params...)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2a0b2617.
//
// Solidity: function getAccountDetails(acctIndex uint256) constant returns(_acct address, _acctAccess uint8)
func (_Permissions *PermissionsCaller) GetAccountDetails(opts *bind.CallOpts, acctIndex *big.Int) (struct {
	Acct       common.Address
	AcctAccess uint8
}, error) {
	ret := new(struct {
		Acct       common.Address
		AcctAccess uint8
	})
	out := ret
	err := _Permissions.contract.Call(opts, out, "getAccountDetails", acctIndex)
	return *ret, err
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2a0b2617.
//
// Solidity: function getAccountDetails(acctIndex uint256) constant returns(_acct address, _acctAccess uint8)
func (_Permissions *PermissionsSession) GetAccountDetails(acctIndex *big.Int) (struct {
	Acct       common.Address
	AcctAccess uint8
}, error) {
	return _Permissions.Contract.GetAccountDetails(&_Permissions.CallOpts, acctIndex)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2a0b2617.
//
// Solidity: function getAccountDetails(acctIndex uint256) constant returns(_acct address, _acctAccess uint8)
func (_Permissions *PermissionsCallerSession) GetAccountDetails(acctIndex *big.Int) (struct {
	Acct       common.Address
	AcctAccess uint8
}, error) {
	return _Permissions.Contract.GetAccountDetails(&_Permissions.CallOpts, acctIndex)
}

// GetFullAccessAccountCount is a free data retrieval call binding the contract method 0xfc0eea99.
//
// Solidity: function getFullAccessAccountCount() constant returns(uint256)
func (_Permissions *PermissionsCaller) GetFullAccessAccountCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "getFullAccessAccountCount")
	return *ret0, err
}

// GetFullAccessAccountCount is a free data retrieval call binding the contract method 0xfc0eea99.
//
// Solidity: function getFullAccessAccountCount() constant returns(uint256)
func (_Permissions *PermissionsSession) GetFullAccessAccountCount() (*big.Int, error) {
	return _Permissions.Contract.GetFullAccessAccountCount(&_Permissions.CallOpts)
}

// GetFullAccessAccountCount is a free data retrieval call binding the contract method 0xfc0eea99.
//
// Solidity: function getFullAccessAccountCount() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) GetFullAccessAccountCount() (*big.Int, error) {
	return _Permissions.Contract.GetFullAccessAccountCount(&_Permissions.CallOpts)
}

// GetInitAccountsCount is a free data retrieval call binding the contract method 0xa430db5f.
//
// Solidity: function getInitAccountsCount() constant returns(uint256)
func (_Permissions *PermissionsCaller) GetInitAccountsCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "getInitAccountsCount")
	return *ret0, err
}

// GetInitAccountsCount is a free data retrieval call binding the contract method 0xa430db5f.
//
// Solidity: function getInitAccountsCount() constant returns(uint256)
func (_Permissions *PermissionsSession) GetInitAccountsCount() (*big.Int, error) {
	return _Permissions.Contract.GetInitAccountsCount(&_Permissions.CallOpts)
}

// GetInitAccountsCount is a free data retrieval call binding the contract method 0xa430db5f.
//
// Solidity: function getInitAccountsCount() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) GetInitAccountsCount() (*big.Int, error) {
	return _Permissions.Contract.GetInitAccountsCount(&_Permissions.CallOpts)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_Permissions *PermissionsCaller) GetNetworkBootStatus(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "getNetworkBootStatus")
	return *ret0, err
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_Permissions *PermissionsSession) GetNetworkBootStatus() (bool, error) {
	return _Permissions.Contract.GetNetworkBootStatus(&_Permissions.CallOpts)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_Permissions *PermissionsCallerSession) GetNetworkBootStatus() (bool, error) {
	return _Permissions.Contract.GetNetworkBootStatus(&_Permissions.CallOpts)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(enodeId string) constant returns(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string, _nodeStatus uint8)
func (_Permissions *PermissionsCaller) GetNodeDetails(opts *bind.CallOpts, enodeId string) (struct {
	EnodeId    string
	IpAddrPort string
	DiscPort   string
	RaftPort   string
	NodeStatus uint8
}, error) {
	ret := new(struct {
		EnodeId    string
		IpAddrPort string
		DiscPort   string
		RaftPort   string
		NodeStatus uint8
	})
	out := ret
	err := _Permissions.contract.Call(opts, out, "getNodeDetails", enodeId)
	return *ret, err
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(enodeId string) constant returns(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string, _nodeStatus uint8)
func (_Permissions *PermissionsSession) GetNodeDetails(enodeId string) (struct {
	EnodeId    string
	IpAddrPort string
	DiscPort   string
	RaftPort   string
	NodeStatus uint8
}, error) {
	return _Permissions.Contract.GetNodeDetails(&_Permissions.CallOpts, enodeId)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(enodeId string) constant returns(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string, _nodeStatus uint8)
func (_Permissions *PermissionsCallerSession) GetNodeDetails(enodeId string) (struct {
	EnodeId    string
	IpAddrPort string
	DiscPort   string
	RaftPort   string
	NodeStatus uint8
}, error) {
	return _Permissions.Contract.GetNodeDetails(&_Permissions.CallOpts, enodeId)
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(nodeIndex uint256) constant returns(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string, _nodeStatus uint8)
func (_Permissions *PermissionsCaller) GetNodeDetailsFromIndex(opts *bind.CallOpts, nodeIndex *big.Int) (struct {
	EnodeId    string
	IpAddrPort string
	DiscPort   string
	RaftPort   string
	NodeStatus uint8
}, error) {
	ret := new(struct {
		EnodeId    string
		IpAddrPort string
		DiscPort   string
		RaftPort   string
		NodeStatus uint8
	})
	out := ret
	err := _Permissions.contract.Call(opts, out, "getNodeDetailsFromIndex", nodeIndex)
	return *ret, err
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(nodeIndex uint256) constant returns(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string, _nodeStatus uint8)
func (_Permissions *PermissionsSession) GetNodeDetailsFromIndex(nodeIndex *big.Int) (struct {
	EnodeId    string
	IpAddrPort string
	DiscPort   string
	RaftPort   string
	NodeStatus uint8
}, error) {
	return _Permissions.Contract.GetNodeDetailsFromIndex(&_Permissions.CallOpts, nodeIndex)
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(nodeIndex uint256) constant returns(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string, _nodeStatus uint8)
func (_Permissions *PermissionsCallerSession) GetNodeDetailsFromIndex(nodeIndex *big.Int) (struct {
	EnodeId    string
	IpAddrPort string
	DiscPort   string
	RaftPort   string
	NodeStatus uint8
}, error) {
	return _Permissions.Contract.GetNodeDetailsFromIndex(&_Permissions.CallOpts, nodeIndex)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(uint8)
func (_Permissions *PermissionsCaller) GetNodeStatus(opts *bind.CallOpts, _enodeId string) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "getNodeStatus", _enodeId)
	return *ret0, err
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(uint8)
func (_Permissions *PermissionsSession) GetNodeStatus(_enodeId string) (uint8, error) {
	return _Permissions.Contract.GetNodeStatus(&_Permissions.CallOpts, _enodeId)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(uint8)
func (_Permissions *PermissionsCallerSession) GetNodeStatus(_enodeId string) (uint8, error) {
	return _Permissions.Contract.GetNodeStatus(&_Permissions.CallOpts, _enodeId)
}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() constant returns(uint256)
func (_Permissions *PermissionsCaller) GetNumberOfAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "getNumberOfAccounts")
	return *ret0, err
}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() constant returns(uint256)
func (_Permissions *PermissionsSession) GetNumberOfAccounts() (*big.Int, error) {
	return _Permissions.Contract.GetNumberOfAccounts(&_Permissions.CallOpts)
}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) GetNumberOfAccounts() (*big.Int, error) {
	return _Permissions.Contract.GetNumberOfAccounts(&_Permissions.CallOpts)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_Permissions *PermissionsCaller) GetNumberOfNodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "getNumberOfNodes")
	return *ret0, err
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_Permissions *PermissionsSession) GetNumberOfNodes() (*big.Int, error) {
	return _Permissions.Contract.GetNumberOfNodes(&_Permissions.CallOpts)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) GetNumberOfNodes() (*big.Int, error) {
	return _Permissions.Contract.GetNumberOfNodes(&_Permissions.CallOpts)
}

// GetNumberOfValidVoters is a free data retrieval call binding the contract method 0xd12d44dc.
//
// Solidity: function getNumberOfValidVoters() constant returns(uint256)
func (_Permissions *PermissionsCaller) GetNumberOfValidVoters(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "getNumberOfValidVoters")
	return *ret0, err
}

// GetNumberOfValidVoters is a free data retrieval call binding the contract method 0xd12d44dc.
//
// Solidity: function getNumberOfValidVoters() constant returns(uint256)
func (_Permissions *PermissionsSession) GetNumberOfValidVoters() (*big.Int, error) {
	return _Permissions.Contract.GetNumberOfValidVoters(&_Permissions.CallOpts)
}

// GetNumberOfValidVoters is a free data retrieval call binding the contract method 0xd12d44dc.
//
// Solidity: function getNumberOfValidVoters() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) GetNumberOfValidVoters() (*big.Int, error) {
	return _Permissions.Contract.GetNumberOfValidVoters(&_Permissions.CallOpts)
}

// GetNumberOfVoters is a free data retrieval call binding the contract method 0x84865b66.
//
// Solidity: function getNumberOfVoters() constant returns(uint256)
func (_Permissions *PermissionsCaller) GetNumberOfVoters(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "getNumberOfVoters")
	return *ret0, err
}

// GetNumberOfVoters is a free data retrieval call binding the contract method 0x84865b66.
//
// Solidity: function getNumberOfVoters() constant returns(uint256)
func (_Permissions *PermissionsSession) GetNumberOfVoters() (*big.Int, error) {
	return _Permissions.Contract.GetNumberOfVoters(&_Permissions.CallOpts)
}

// GetNumberOfVoters is a free data retrieval call binding the contract method 0x84865b66.
//
// Solidity: function getNumberOfVoters() constant returns(uint256)
func (_Permissions *PermissionsCallerSession) GetNumberOfVoters() (*big.Int, error) {
	return _Permissions.Contract.GetNumberOfVoters(&_Permissions.CallOpts)
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(_enodeId string) constant returns(uint256)
func (_Permissions *PermissionsCaller) GetVoteCount(opts *bind.CallOpts, _enodeId string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "getVoteCount", _enodeId)
	return *ret0, err
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(_enodeId string) constant returns(uint256)
func (_Permissions *PermissionsSession) GetVoteCount(_enodeId string) (*big.Int, error) {
	return _Permissions.Contract.GetVoteCount(&_Permissions.CallOpts, _enodeId)
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(_enodeId string) constant returns(uint256)
func (_Permissions *PermissionsCallerSession) GetVoteCount(_enodeId string) (*big.Int, error) {
	return _Permissions.Contract.GetVoteCount(&_Permissions.CallOpts, _enodeId)
}

// GetVoteStatus is a free data retrieval call binding the contract method 0x0fdc2150.
//
// Solidity: function getVoteStatus(_enodeId string, _address address) constant returns(bool)
func (_Permissions *PermissionsCaller) GetVoteStatus(opts *bind.CallOpts, _enodeId string, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "getVoteStatus", _enodeId, _address)
	return *ret0, err
}

// GetVoteStatus is a free data retrieval call binding the contract method 0x0fdc2150.
//
// Solidity: function getVoteStatus(_enodeId string, _address address) constant returns(bool)
func (_Permissions *PermissionsSession) GetVoteStatus(_enodeId string, _address common.Address) (bool, error) {
	return _Permissions.Contract.GetVoteStatus(&_Permissions.CallOpts, _enodeId, _address)
}

// GetVoteStatus is a free data retrieval call binding the contract method 0x0fdc2150.
//
// Solidity: function getVoteStatus(_enodeId string, _address address) constant returns(bool)
func (_Permissions *PermissionsCallerSession) GetVoteStatus(_enodeId string, _address common.Address) (bool, error) {
	return _Permissions.Contract.GetVoteStatus(&_Permissions.CallOpts, _enodeId, _address)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(i uint256) constant returns(_addr address, _voterStatus uint8)
func (_Permissions *PermissionsCaller) GetVoter(opts *bind.CallOpts, i *big.Int) (struct {
	Addr        common.Address
	VoterStatus uint8
}, error) {
	ret := new(struct {
		Addr        common.Address
		VoterStatus uint8
	})
	out := ret
	err := _Permissions.contract.Call(opts, out, "getVoter", i)
	return *ret, err
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(i uint256) constant returns(_addr address, _voterStatus uint8)
func (_Permissions *PermissionsSession) GetVoter(i *big.Int) (struct {
	Addr        common.Address
	VoterStatus uint8
}, error) {
	return _Permissions.Contract.GetVoter(&_Permissions.CallOpts, i)
}

// GetVoter is a free data retrieval call binding the contract method 0xd07bff0c.
//
// Solidity: function getVoter(i uint256) constant returns(_addr address, _voterStatus uint8)
func (_Permissions *PermissionsCallerSession) GetVoter(i *big.Int) (struct {
	Addr        common.Address
	VoterStatus uint8
}, error) {
	return _Permissions.Contract.GetVoter(&_Permissions.CallOpts, i)
}

// GetVoterCount is a free data retrieval call binding the contract method 0x11174a29.
//
// Solidity: function getVoterCount() constant returns(uint256, uint256)
func (_Permissions *PermissionsCaller) GetVoterCount(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Permissions.contract.Call(opts, out, "getVoterCount")
	return *ret0, *ret1, err
}

// GetVoterCount is a free data retrieval call binding the contract method 0x11174a29.
//
// Solidity: function getVoterCount() constant returns(uint256, uint256)
func (_Permissions *PermissionsSession) GetVoterCount() (*big.Int, *big.Int, error) {
	return _Permissions.Contract.GetVoterCount(&_Permissions.CallOpts)
}

// GetVoterCount is a free data retrieval call binding the contract method 0x11174a29.
//
// Solidity: function getVoterCount() constant returns(uint256, uint256)
func (_Permissions *PermissionsCallerSession) GetVoterCount() (*big.Int, *big.Int, error) {
	return _Permissions.Contract.GetVoterCount(&_Permissions.CallOpts)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(_acctid address) constant returns(bool)
func (_Permissions *PermissionsCaller) IsVoter(opts *bind.CallOpts, _acctid common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "isVoter", _acctid)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(_acctid address) constant returns(bool)
func (_Permissions *PermissionsSession) IsVoter(_acctid common.Address) (bool, error) {
	return _Permissions.Contract.IsVoter(&_Permissions.CallOpts, _acctid)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(_acctid address) constant returns(bool)
func (_Permissions *PermissionsCallerSession) IsVoter(_acctid common.Address) (bool, error) {
	return _Permissions.Contract.IsVoter(&_Permissions.CallOpts, _acctid)
}

// ActivateNode is a paid mutator transaction binding the contract method 0x14a945e5.
//
// Solidity: function activateNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactor) ActivateNode(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "activateNode", _enodeId)
}

// ActivateNode is a paid mutator transaction binding the contract method 0x14a945e5.
//
// Solidity: function activateNode(_enodeId string) returns()
func (_Permissions *PermissionsSession) ActivateNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ActivateNode(&_Permissions.TransactOpts, _enodeId)
}

// ActivateNode is a paid mutator transaction binding the contract method 0x14a945e5.
//
// Solidity: function activateNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactorSession) ActivateNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ActivateNode(&_Permissions.TransactOpts, _enodeId)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(_address address) returns()
func (_Permissions *PermissionsTransactor) AddVoter(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "addVoter", _address)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(_address address) returns()
func (_Permissions *PermissionsSession) AddVoter(_address common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.AddVoter(&_Permissions.TransactOpts, _address)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(_address address) returns()
func (_Permissions *PermissionsTransactorSession) AddVoter(_address common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.AddVoter(&_Permissions.TransactOpts, _address)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x21c67088.
//
// Solidity: function approveNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactor) ApproveNode(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "approveNode", _enodeId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x21c67088.
//
// Solidity: function approveNode(_enodeId string) returns()
func (_Permissions *PermissionsSession) ApproveNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ApproveNode(&_Permissions.TransactOpts, _enodeId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x21c67088.
//
// Solidity: function approveNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactorSession) ApproveNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ApproveNode(&_Permissions.TransactOpts, _enodeId)
}

// BlacklistNode is a paid mutator transaction binding the contract method 0x487363f9.
//
// Solidity: function blacklistNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactor) BlacklistNode(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "blacklistNode", _enodeId)
}

// BlacklistNode is a paid mutator transaction binding the contract method 0x487363f9.
//
// Solidity: function blacklistNode(_enodeId string) returns()
func (_Permissions *PermissionsSession) BlacklistNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.BlacklistNode(&_Permissions.TransactOpts, _enodeId)
}

// BlacklistNode is a paid mutator transaction binding the contract method 0x487363f9.
//
// Solidity: function blacklistNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactorSession) BlacklistNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.BlacklistNode(&_Permissions.TransactOpts, _enodeId)
}

// CancelPendingOperation is a paid mutator transaction binding the contract method 0x47dd4953.
//
// Solidity: function cancelPendingOperation(_enodeId string) returns()
func (_Permissions *PermissionsTransactor) CancelPendingOperation(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "cancelPendingOperation", _enodeId)
}

// CancelPendingOperation is a paid mutator transaction binding the contract method 0x47dd4953.
//
// Solidity: function cancelPendingOperation(_enodeId string) returns()
func (_Permissions *PermissionsSession) CancelPendingOperation(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.CancelPendingOperation(&_Permissions.TransactOpts, _enodeId)
}

// CancelPendingOperation is a paid mutator transaction binding the contract method 0x47dd4953.
//
// Solidity: function cancelPendingOperation(_enodeId string) returns()
func (_Permissions *PermissionsTransactorSession) CancelPendingOperation(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.CancelPendingOperation(&_Permissions.TransactOpts, _enodeId)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0x420c26de.
//
// Solidity: function deactivateNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactor) DeactivateNode(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "deactivateNode", _enodeId)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0x420c26de.
//
// Solidity: function deactivateNode(_enodeId string) returns()
func (_Permissions *PermissionsSession) DeactivateNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.DeactivateNode(&_Permissions.TransactOpts, _enodeId)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0x420c26de.
//
// Solidity: function deactivateNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactorSession) DeactivateNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.DeactivateNode(&_Permissions.TransactOpts, _enodeId)
}

// InitAccounts is a paid mutator transaction binding the contract method 0x49f2545b.
//
// Solidity: function initAccounts() returns()
func (_Permissions *PermissionsTransactor) InitAccounts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "initAccounts")
}

// InitAccounts is a paid mutator transaction binding the contract method 0x49f2545b.
//
// Solidity: function initAccounts() returns()
func (_Permissions *PermissionsSession) InitAccounts() (*types.Transaction, error) {
	return _Permissions.Contract.InitAccounts(&_Permissions.TransactOpts)
}

// InitAccounts is a paid mutator transaction binding the contract method 0x49f2545b.
//
// Solidity: function initAccounts() returns()
func (_Permissions *PermissionsTransactorSession) InitAccounts() (*types.Transaction, error) {
	return _Permissions.Contract.InitAccounts(&_Permissions.TransactOpts)
}

// ProposeDeactivation is a paid mutator transaction binding the contract method 0x47b8fe57.
//
// Solidity: function proposeDeactivation(_enodeId string) returns()
func (_Permissions *PermissionsTransactor) ProposeDeactivation(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "proposeDeactivation", _enodeId)
}

// ProposeDeactivation is a paid mutator transaction binding the contract method 0x47b8fe57.
//
// Solidity: function proposeDeactivation(_enodeId string) returns()
func (_Permissions *PermissionsSession) ProposeDeactivation(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeDeactivation(&_Permissions.TransactOpts, _enodeId)
}

// ProposeDeactivation is a paid mutator transaction binding the contract method 0x47b8fe57.
//
// Solidity: function proposeDeactivation(_enodeId string) returns()
func (_Permissions *PermissionsTransactorSession) ProposeDeactivation(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeDeactivation(&_Permissions.TransactOpts, _enodeId)
}

// ProposeNode is a paid mutator transaction binding the contract method 0xc7ab7ccf.
//
// Solidity: function proposeNode(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string) returns()
func (_Permissions *PermissionsTransactor) ProposeNode(opts *bind.TransactOpts, _enodeId string, _ipAddrPort string, _discPort string, _raftPort string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "proposeNode", _enodeId, _ipAddrPort, _discPort, _raftPort)
}

// ProposeNode is a paid mutator transaction binding the contract method 0xc7ab7ccf.
//
// Solidity: function proposeNode(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string) returns()
func (_Permissions *PermissionsSession) ProposeNode(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeNode(&_Permissions.TransactOpts, _enodeId, _ipAddrPort, _discPort, _raftPort)
}

// ProposeNode is a paid mutator transaction binding the contract method 0xc7ab7ccf.
//
// Solidity: function proposeNode(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string) returns()
func (_Permissions *PermissionsTransactorSession) ProposeNode(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeNode(&_Permissions.TransactOpts, _enodeId, _ipAddrPort, _discPort, _raftPort)
}

// ProposeNodeActivation is a paid mutator transaction binding the contract method 0xe51008e1.
//
// Solidity: function proposeNodeActivation(_enodeId string) returns()
func (_Permissions *PermissionsTransactor) ProposeNodeActivation(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "proposeNodeActivation", _enodeId)
}

// ProposeNodeActivation is a paid mutator transaction binding the contract method 0xe51008e1.
//
// Solidity: function proposeNodeActivation(_enodeId string) returns()
func (_Permissions *PermissionsSession) ProposeNodeActivation(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeNodeActivation(&_Permissions.TransactOpts, _enodeId)
}

// ProposeNodeActivation is a paid mutator transaction binding the contract method 0xe51008e1.
//
// Solidity: function proposeNodeActivation(_enodeId string) returns()
func (_Permissions *PermissionsTransactorSession) ProposeNodeActivation(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeNodeActivation(&_Permissions.TransactOpts, _enodeId)
}

// ProposeNodeBlacklisting is a paid mutator transaction binding the contract method 0x60514a5a.
//
// Solidity: function proposeNodeBlacklisting(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string) returns()
func (_Permissions *PermissionsTransactor) ProposeNodeBlacklisting(opts *bind.TransactOpts, _enodeId string, _ipAddrPort string, _discPort string, _raftPort string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "proposeNodeBlacklisting", _enodeId, _ipAddrPort, _discPort, _raftPort)
}

// ProposeNodeBlacklisting is a paid mutator transaction binding the contract method 0x60514a5a.
//
// Solidity: function proposeNodeBlacklisting(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string) returns()
func (_Permissions *PermissionsSession) ProposeNodeBlacklisting(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeNodeBlacklisting(&_Permissions.TransactOpts, _enodeId, _ipAddrPort, _discPort, _raftPort)
}

// ProposeNodeBlacklisting is a paid mutator transaction binding the contract method 0x60514a5a.
//
// Solidity: function proposeNodeBlacklisting(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string) returns()
func (_Permissions *PermissionsTransactorSession) ProposeNodeBlacklisting(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeNodeBlacklisting(&_Permissions.TransactOpts, _enodeId, _ipAddrPort, _discPort, _raftPort)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(_address address) returns()
func (_Permissions *PermissionsTransactor) RemoveVoter(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "removeVoter", _address)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(_address address) returns()
func (_Permissions *PermissionsSession) RemoveVoter(_address common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.RemoveVoter(&_Permissions.TransactOpts, _address)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(_address address) returns()
func (_Permissions *PermissionsTransactorSession) RemoveVoter(_address common.Address) (*types.Transaction, error) {
	return _Permissions.Contract.RemoveVoter(&_Permissions.TransactOpts, _address)
}

// UpdateAccountAccess is a paid mutator transaction binding the contract method 0xc6962b99.
//
// Solidity: function updateAccountAccess(_address address, _accountAccess uint8) returns()
func (_Permissions *PermissionsTransactor) UpdateAccountAccess(opts *bind.TransactOpts, _address common.Address, _accountAccess uint8) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "updateAccountAccess", _address, _accountAccess)
}

// UpdateAccountAccess is a paid mutator transaction binding the contract method 0xc6962b99.
//
// Solidity: function updateAccountAccess(_address address, _accountAccess uint8) returns()
func (_Permissions *PermissionsSession) UpdateAccountAccess(_address common.Address, _accountAccess uint8) (*types.Transaction, error) {
	return _Permissions.Contract.UpdateAccountAccess(&_Permissions.TransactOpts, _address, _accountAccess)
}

// UpdateAccountAccess is a paid mutator transaction binding the contract method 0xc6962b99.
//
// Solidity: function updateAccountAccess(_address address, _accountAccess uint8) returns()
func (_Permissions *PermissionsTransactorSession) UpdateAccountAccess(_address common.Address, _accountAccess uint8) (*types.Transaction, error) {
	return _Permissions.Contract.UpdateAccountAccess(&_Permissions.TransactOpts, _address, _accountAccess)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_Permissions *PermissionsTransactor) UpdateNetworkBootStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "updateNetworkBootStatus")
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_Permissions *PermissionsSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _Permissions.Contract.UpdateNetworkBootStatus(&_Permissions.TransactOpts)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_Permissions *PermissionsTransactorSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _Permissions.Contract.UpdateNetworkBootStatus(&_Permissions.TransactOpts)
}

// PermissionsAccountAccessModifiedIterator is returned from FilterAccountAccessModified and is used to iterate over the raw logs and unpacked data for AccountAccessModified events raised by the Permissions contract.
type PermissionsAccountAccessModifiedIterator struct {
	Event *PermissionsAccountAccessModified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PermissionsAccountAccessModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsAccountAccessModified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PermissionsAccountAccessModified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PermissionsAccountAccessModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsAccountAccessModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsAccountAccessModified represents a AccountAccessModified event raised by the Permissions contract.
type PermissionsAccountAccessModified struct {
	Address common.Address
	Access  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountAccessModified is a free log retrieval operation binding the contract event 0x5c7c83802ef5601aed89f3f4e4ab42298ecf8ac3fe099adad5712fc65ba9676d.
//
// Solidity: e AccountAccessModified(_address address, _access uint8)
func (_Permissions *PermissionsFilterer) FilterAccountAccessModified(opts *bind.FilterOpts) (*PermissionsAccountAccessModifiedIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "AccountAccessModified")
	if err != nil {
		return nil, err
	}
	return &PermissionsAccountAccessModifiedIterator{contract: _Permissions.contract, event: "AccountAccessModified", logs: logs, sub: sub}, nil
}

// WatchAccountAccessModified is a free log subscription operation binding the contract event 0x5c7c83802ef5601aed89f3f4e4ab42298ecf8ac3fe099adad5712fc65ba9676d.
//
// Solidity: e AccountAccessModified(_address address, _access uint8)
func (_Permissions *PermissionsFilterer) WatchAccountAccessModified(opts *bind.WatchOpts, sink chan<- *PermissionsAccountAccessModified) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "AccountAccessModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsAccountAccessModified)
				if err := _Permissions.contract.UnpackLog(event, "AccountAccessModified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PermissionsNodeActivatedIterator is returned from FilterNodeActivated and is used to iterate over the raw logs and unpacked data for NodeActivated events raised by the Permissions contract.
type PermissionsNodeActivatedIterator struct {
	Event *PermissionsNodeActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PermissionsNodeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNodeActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PermissionsNodeActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PermissionsNodeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNodeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNodeActivated represents a NodeActivated event raised by the Permissions contract.
type PermissionsNodeActivated struct {
	EnodeId    string
	IpAddrPort string
	DiscPort   string
	RaftPort   string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeActivated is a free log retrieval operation binding the contract event 0xd277bd13c43f8ddd20884d02df780044b5faaa1d9e2d4db2d0416fdfcb65d6bf.
//
// Solidity: e NodeActivated(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string)
func (_Permissions *PermissionsFilterer) FilterNodeActivated(opts *bind.FilterOpts) (*PermissionsNodeActivatedIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NodeActivated")
	if err != nil {
		return nil, err
	}
	return &PermissionsNodeActivatedIterator{contract: _Permissions.contract, event: "NodeActivated", logs: logs, sub: sub}, nil
}

// WatchNodeActivated is a free log subscription operation binding the contract event 0xd277bd13c43f8ddd20884d02df780044b5faaa1d9e2d4db2d0416fdfcb65d6bf.
//
// Solidity: e NodeActivated(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string)
func (_Permissions *PermissionsFilterer) WatchNodeActivated(opts *bind.WatchOpts, sink chan<- *PermissionsNodeActivated) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NodeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNodeActivated)
				if err := _Permissions.contract.UnpackLog(event, "NodeActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PermissionsNodeApprovedIterator is returned from FilterNodeApproved and is used to iterate over the raw logs and unpacked data for NodeApproved events raised by the Permissions contract.
type PermissionsNodeApprovedIterator struct {
	Event *PermissionsNodeApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PermissionsNodeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNodeApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PermissionsNodeApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PermissionsNodeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNodeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNodeApproved represents a NodeApproved event raised by the Permissions contract.
type PermissionsNodeApproved struct {
	EnodeId    string
	IpAddrPort string
	DiscPort   string
	RaftPort   string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeApproved is a free log retrieval operation binding the contract event 0xc6d86deaa3b3cf7c374cfd405aae9f08571fce2bf6ccfe8f98a399cda8960a98.
//
// Solidity: e NodeApproved(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string)
func (_Permissions *PermissionsFilterer) FilterNodeApproved(opts *bind.FilterOpts) (*PermissionsNodeApprovedIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return &PermissionsNodeApprovedIterator{contract: _Permissions.contract, event: "NodeApproved", logs: logs, sub: sub}, nil
}

// WatchNodeApproved is a free log subscription operation binding the contract event 0xc6d86deaa3b3cf7c374cfd405aae9f08571fce2bf6ccfe8f98a399cda8960a98.
//
// Solidity: e NodeApproved(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string)
func (_Permissions *PermissionsFilterer) WatchNodeApproved(opts *bind.WatchOpts, sink chan<- *PermissionsNodeApproved) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNodeApproved)
				if err := _Permissions.contract.UnpackLog(event, "NodeApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PermissionsNodeBlacklistedIterator is returned from FilterNodeBlacklisted and is used to iterate over the raw logs and unpacked data for NodeBlacklisted events raised by the Permissions contract.
type PermissionsNodeBlacklistedIterator struct {
	Event *PermissionsNodeBlacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PermissionsNodeBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNodeBlacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PermissionsNodeBlacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PermissionsNodeBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNodeBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNodeBlacklisted represents a NodeBlacklisted event raised by the Permissions contract.
type PermissionsNodeBlacklisted struct {
	EnodeId    string
	IpAddrPort string
	DiscPort   string
	RaftPort   string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeBlacklisted is a free log retrieval operation binding the contract event 0xe1b239bf9d5854aeca74dfeac25d6ce470230bdb5f0eec48713c4375becfe97e.
//
// Solidity: e NodeBlacklisted(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string)
func (_Permissions *PermissionsFilterer) FilterNodeBlacklisted(opts *bind.FilterOpts) (*PermissionsNodeBlacklistedIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NodeBlacklisted")
	if err != nil {
		return nil, err
	}
	return &PermissionsNodeBlacklistedIterator{contract: _Permissions.contract, event: "NodeBlacklisted", logs: logs, sub: sub}, nil
}

// WatchNodeBlacklisted is a free log subscription operation binding the contract event 0xe1b239bf9d5854aeca74dfeac25d6ce470230bdb5f0eec48713c4375becfe97e.
//
// Solidity: e NodeBlacklisted(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string)
func (_Permissions *PermissionsFilterer) WatchNodeBlacklisted(opts *bind.WatchOpts, sink chan<- *PermissionsNodeBlacklisted) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NodeBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNodeBlacklisted)
				if err := _Permissions.contract.UnpackLog(event, "NodeBlacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PermissionsNodeDeactivatedIterator is returned from FilterNodeDeactivated and is used to iterate over the raw logs and unpacked data for NodeDeactivated events raised by the Permissions contract.
type PermissionsNodeDeactivatedIterator struct {
	Event *PermissionsNodeDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PermissionsNodeDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNodeDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PermissionsNodeDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PermissionsNodeDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNodeDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNodeDeactivated represents a NodeDeactivated event raised by the Permissions contract.
type PermissionsNodeDeactivated struct {
	EnodeId    string
	IpAddrPort string
	DiscPort   string
	RaftPort   string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeDeactivated is a free log retrieval operation binding the contract event 0xd5fa0ecdea15b332dd0a270c65234bc4aee212edf2ed62eb2fd182ef55ca98a1.
//
// Solidity: e NodeDeactivated(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string)
func (_Permissions *PermissionsFilterer) FilterNodeDeactivated(opts *bind.FilterOpts) (*PermissionsNodeDeactivatedIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return &PermissionsNodeDeactivatedIterator{contract: _Permissions.contract, event: "NodeDeactivated", logs: logs, sub: sub}, nil
}

// WatchNodeDeactivated is a free log subscription operation binding the contract event 0xd5fa0ecdea15b332dd0a270c65234bc4aee212edf2ed62eb2fd182ef55ca98a1.
//
// Solidity: e NodeDeactivated(_enodeId string, _ipAddrPort string, _discPort string, _raftPort string)
func (_Permissions *PermissionsFilterer) WatchNodeDeactivated(opts *bind.WatchOpts, sink chan<- *PermissionsNodeDeactivated) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNodeDeactivated)
				if err := _Permissions.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PermissionsNodePendingActivationIterator is returned from FilterNodePendingActivation and is used to iterate over the raw logs and unpacked data for NodePendingActivation events raised by the Permissions contract.
type PermissionsNodePendingActivationIterator struct {
	Event *PermissionsNodePendingActivation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PermissionsNodePendingActivationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNodePendingActivation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PermissionsNodePendingActivation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PermissionsNodePendingActivationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNodePendingActivationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNodePendingActivation represents a NodePendingActivation event raised by the Permissions contract.
type PermissionsNodePendingActivation struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodePendingActivation is a free log retrieval operation binding the contract event 0x7b961104d9e9db7d30803aff3fa117bc41799d2faa2d2e339cf1a1f3513b0eef.
//
// Solidity: e NodePendingActivation(_enodeId string)
func (_Permissions *PermissionsFilterer) FilterNodePendingActivation(opts *bind.FilterOpts) (*PermissionsNodePendingActivationIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NodePendingActivation")
	if err != nil {
		return nil, err
	}
	return &PermissionsNodePendingActivationIterator{contract: _Permissions.contract, event: "NodePendingActivation", logs: logs, sub: sub}, nil
}

// WatchNodePendingActivation is a free log subscription operation binding the contract event 0x7b961104d9e9db7d30803aff3fa117bc41799d2faa2d2e339cf1a1f3513b0eef.
//
// Solidity: e NodePendingActivation(_enodeId string)
func (_Permissions *PermissionsFilterer) WatchNodePendingActivation(opts *bind.WatchOpts, sink chan<- *PermissionsNodePendingActivation) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NodePendingActivation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNodePendingActivation)
				if err := _Permissions.contract.UnpackLog(event, "NodePendingActivation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PermissionsNodePendingBlacklistIterator is returned from FilterNodePendingBlacklist and is used to iterate over the raw logs and unpacked data for NodePendingBlacklist events raised by the Permissions contract.
type PermissionsNodePendingBlacklistIterator struct {
	Event *PermissionsNodePendingBlacklist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PermissionsNodePendingBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNodePendingBlacklist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PermissionsNodePendingBlacklist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PermissionsNodePendingBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNodePendingBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNodePendingBlacklist represents a NodePendingBlacklist event raised by the Permissions contract.
type PermissionsNodePendingBlacklist struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodePendingBlacklist is a free log retrieval operation binding the contract event 0xb249ebebf429f1c79f3c9663998b3e22d45f242de6527c4a95e41d4d28115d74.
//
// Solidity: e NodePendingBlacklist(_enodeId string)
func (_Permissions *PermissionsFilterer) FilterNodePendingBlacklist(opts *bind.FilterOpts) (*PermissionsNodePendingBlacklistIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NodePendingBlacklist")
	if err != nil {
		return nil, err
	}
	return &PermissionsNodePendingBlacklistIterator{contract: _Permissions.contract, event: "NodePendingBlacklist", logs: logs, sub: sub}, nil
}

// WatchNodePendingBlacklist is a free log subscription operation binding the contract event 0xb249ebebf429f1c79f3c9663998b3e22d45f242de6527c4a95e41d4d28115d74.
//
// Solidity: e NodePendingBlacklist(_enodeId string)
func (_Permissions *PermissionsFilterer) WatchNodePendingBlacklist(opts *bind.WatchOpts, sink chan<- *PermissionsNodePendingBlacklist) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NodePendingBlacklist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNodePendingBlacklist)
				if err := _Permissions.contract.UnpackLog(event, "NodePendingBlacklist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PermissionsNodePendingDeactivationIterator is returned from FilterNodePendingDeactivation and is used to iterate over the raw logs and unpacked data for NodePendingDeactivation events raised by the Permissions contract.
type PermissionsNodePendingDeactivationIterator struct {
	Event *PermissionsNodePendingDeactivation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PermissionsNodePendingDeactivationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNodePendingDeactivation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PermissionsNodePendingDeactivation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PermissionsNodePendingDeactivationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNodePendingDeactivationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNodePendingDeactivation represents a NodePendingDeactivation event raised by the Permissions contract.
type PermissionsNodePendingDeactivation struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodePendingDeactivation is a free log retrieval operation binding the contract event 0x2b5689b33f48f1dcbda2084e130a9bee7b3bf14dc767ea74cbdf3e5fffb118e4.
//
// Solidity: e NodePendingDeactivation(_enodeId string)
func (_Permissions *PermissionsFilterer) FilterNodePendingDeactivation(opts *bind.FilterOpts) (*PermissionsNodePendingDeactivationIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NodePendingDeactivation")
	if err != nil {
		return nil, err
	}
	return &PermissionsNodePendingDeactivationIterator{contract: _Permissions.contract, event: "NodePendingDeactivation", logs: logs, sub: sub}, nil
}

// WatchNodePendingDeactivation is a free log subscription operation binding the contract event 0x2b5689b33f48f1dcbda2084e130a9bee7b3bf14dc767ea74cbdf3e5fffb118e4.
//
// Solidity: e NodePendingDeactivation(_enodeId string)
func (_Permissions *PermissionsFilterer) WatchNodePendingDeactivation(opts *bind.WatchOpts, sink chan<- *PermissionsNodePendingDeactivation) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NodePendingDeactivation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNodePendingDeactivation)
				if err := _Permissions.contract.UnpackLog(event, "NodePendingDeactivation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PermissionsNodeProposedIterator is returned from FilterNodeProposed and is used to iterate over the raw logs and unpacked data for NodeProposed events raised by the Permissions contract.
type PermissionsNodeProposedIterator struct {
	Event *PermissionsNodeProposed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PermissionsNodeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNodeProposed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PermissionsNodeProposed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PermissionsNodeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNodeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNodeProposed represents a NodeProposed event raised by the Permissions contract.
type PermissionsNodeProposed struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeProposed is a free log retrieval operation binding the contract event 0xaddedf3fcf588a85e0b0c3210c30da3f5597ae35221859f7e19427397a2ba80a.
//
// Solidity: e NodeProposed(_enodeId string)
func (_Permissions *PermissionsFilterer) FilterNodeProposed(opts *bind.FilterOpts) (*PermissionsNodeProposedIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NodeProposed")
	if err != nil {
		return nil, err
	}
	return &PermissionsNodeProposedIterator{contract: _Permissions.contract, event: "NodeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeProposed is a free log subscription operation binding the contract event 0xaddedf3fcf588a85e0b0c3210c30da3f5597ae35221859f7e19427397a2ba80a.
//
// Solidity: e NodeProposed(_enodeId string)
func (_Permissions *PermissionsFilterer) WatchNodeProposed(opts *bind.WatchOpts, sink chan<- *PermissionsNodeProposed) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NodeProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNodeProposed)
				if err := _Permissions.contract.UnpackLog(event, "NodeProposed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PermissionsPendingOperationCancelledIterator is returned from FilterPendingOperationCancelled and is used to iterate over the raw logs and unpacked data for PendingOperationCancelled events raised by the Permissions contract.
type PermissionsPendingOperationCancelledIterator struct {
	Event *PermissionsPendingOperationCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PermissionsPendingOperationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsPendingOperationCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PermissionsPendingOperationCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PermissionsPendingOperationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsPendingOperationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsPendingOperationCancelled represents a PendingOperationCancelled event raised by the Permissions contract.
type PermissionsPendingOperationCancelled struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPendingOperationCancelled is a free log retrieval operation binding the contract event 0x94e103729ee0a5bdbfe00e23d13d932dca50dd85b3ad4d85d1299a81583dda93.
//
// Solidity: e PendingOperationCancelled(_enodeId string)
func (_Permissions *PermissionsFilterer) FilterPendingOperationCancelled(opts *bind.FilterOpts) (*PermissionsPendingOperationCancelledIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "PendingOperationCancelled")
	if err != nil {
		return nil, err
	}
	return &PermissionsPendingOperationCancelledIterator{contract: _Permissions.contract, event: "PendingOperationCancelled", logs: logs, sub: sub}, nil
}

// WatchPendingOperationCancelled is a free log subscription operation binding the contract event 0x94e103729ee0a5bdbfe00e23d13d932dca50dd85b3ad4d85d1299a81583dda93.
//
// Solidity: e PendingOperationCancelled(_enodeId string)
func (_Permissions *PermissionsFilterer) WatchPendingOperationCancelled(opts *bind.WatchOpts, sink chan<- *PermissionsPendingOperationCancelled) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "PendingOperationCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsPendingOperationCancelled)
				if err := _Permissions.contract.UnpackLog(event, "PendingOperationCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
