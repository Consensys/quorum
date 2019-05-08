// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PermImplABI is the input ABI used to generate the binding from.
const PermImplABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_status\",\"type\":\"uint256\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateAccountStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgManager\",\"type\":\"address\"},{\"name\":\"_rolesManager\",\"type\":\"address\"},{\"name\":\"_acctManager\",\"type\":\"address\"},{\"name\":\"_voterManager\",\"type\":\"address\"},{\"name\":\"_nodeManager\",\"type\":\"address\"},{\"name\":\"_breadth\",\"type\":\"uint256\"},{\"name\":\"_depth\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_access\",\"type\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"bool\"},{\"name\":\"_admin\",\"type\":\"bool\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addNewRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nwAdminOrg\",\"type\":\"string\"},{\"name\":\"_nwAdminRole\",\"type\":\"string\"},{\"name\":\"_oAdminRole\",\"type\":\"string\"}],\"name\":\"setPolicy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_status\",\"type\":\"uint256\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateOrgStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"assignAdminRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateNetworkBootStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNetworkBootStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"validateAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"}],\"name\":\"addAdminAccounts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveAdminRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"assignAccountRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pOrg\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addSubOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"isOrgAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_status\",\"type\":\"uint256\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveOrgStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isNetworkAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_status\",\"type\":\"uint256\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"updateNodeStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"addAdminNodes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getPendingOp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PermImpl is an auto generated Go binding around an Ethereum contract.
type PermImpl struct {
	PermImplCaller     // Read-only binding to the contract
	PermImplTransactor // Write-only binding to the contract
	PermImplFilterer   // Log filterer for contract events
}

// PermImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermImplSession struct {
	Contract     *PermImpl         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermImplCallerSession struct {
	Contract *PermImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PermImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermImplTransactorSession struct {
	Contract     *PermImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PermImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermImplRaw struct {
	Contract *PermImpl // Generic contract binding to access the raw methods on
}

// PermImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermImplCallerRaw struct {
	Contract *PermImplCaller // Generic read-only contract binding to access the raw methods on
}

// PermImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermImplTransactorRaw struct {
	Contract *PermImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermImpl creates a new instance of PermImpl, bound to a specific deployed contract.
func NewPermImpl(address common.Address, backend bind.ContractBackend) (*PermImpl, error) {
	contract, err := bindPermImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermImpl{PermImplCaller: PermImplCaller{contract: contract}, PermImplTransactor: PermImplTransactor{contract: contract}, PermImplFilterer: PermImplFilterer{contract: contract}}, nil
}

// NewPermImplCaller creates a new read-only instance of PermImpl, bound to a specific deployed contract.
func NewPermImplCaller(address common.Address, caller bind.ContractCaller) (*PermImplCaller, error) {
	contract, err := bindPermImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermImplCaller{contract: contract}, nil
}

// NewPermImplTransactor creates a new write-only instance of PermImpl, bound to a specific deployed contract.
func NewPermImplTransactor(address common.Address, transactor bind.ContractTransactor) (*PermImplTransactor, error) {
	contract, err := bindPermImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermImplTransactor{contract: contract}, nil
}

// NewPermImplFilterer creates a new log filterer instance of PermImpl, bound to a specific deployed contract.
func NewPermImplFilterer(address common.Address, filterer bind.ContractFilterer) (*PermImplFilterer, error) {
	contract, err := bindPermImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermImplFilterer{contract: contract}, nil
}

// bindPermImpl binds a generic wrapper to an already deployed contract.
func bindPermImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermImpl *PermImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermImpl.Contract.PermImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermImpl *PermImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermImpl.Contract.PermImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermImpl *PermImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermImpl.Contract.PermImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermImpl *PermImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermImpl *PermImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermImpl *PermImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermImpl.Contract.contract.Transact(opts, method, params...)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_PermImpl *PermImplCaller) GetNetworkBootStatus(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermImpl.contract.Call(opts, out, "getNetworkBootStatus")
	return *ret0, err
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_PermImpl *PermImplSession) GetNetworkBootStatus() (bool, error) {
	return _PermImpl.Contract.GetNetworkBootStatus(&_PermImpl.CallOpts)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_PermImpl *PermImplCallerSession) GetNetworkBootStatus() (bool, error) {
	return _PermImpl.Contract.GetNetworkBootStatus(&_PermImpl.CallOpts)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(_orgId string) constant returns(string, string, address, uint256)
func (_PermImpl *PermImplCaller) GetPendingOp(opts *bind.CallOpts, _orgId string) (string, string, common.Address, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(common.Address)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _PermImpl.contract.Call(opts, out, "getPendingOp", _orgId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(_orgId string) constant returns(string, string, address, uint256)
func (_PermImpl *PermImplSession) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _PermImpl.Contract.GetPendingOp(&_PermImpl.CallOpts, _orgId)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(_orgId string) constant returns(string, string, address, uint256)
func (_PermImpl *PermImplCallerSession) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _PermImpl.Contract.GetPendingOp(&_PermImpl.CallOpts, _orgId)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(_account address) constant returns(bool)
func (_PermImpl *PermImplCaller) IsNetworkAdmin(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermImpl.contract.Call(opts, out, "isNetworkAdmin", _account)
	return *ret0, err
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(_account address) constant returns(bool)
func (_PermImpl *PermImplSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _PermImpl.Contract.IsNetworkAdmin(&_PermImpl.CallOpts, _account)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(_account address) constant returns(bool)
func (_PermImpl *PermImplCallerSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _PermImpl.Contract.IsNetworkAdmin(&_PermImpl.CallOpts, _account)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(_account address, _orgId string) constant returns(bool)
func (_PermImpl *PermImplCaller) IsOrgAdmin(opts *bind.CallOpts, _account common.Address, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermImpl.contract.Call(opts, out, "isOrgAdmin", _account, _orgId)
	return *ret0, err
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(_account address, _orgId string) constant returns(bool)
func (_PermImpl *PermImplSession) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return _PermImpl.Contract.IsOrgAdmin(&_PermImpl.CallOpts, _account, _orgId)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(_account address, _orgId string) constant returns(bool)
func (_PermImpl *PermImplCallerSession) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return _PermImpl.Contract.IsOrgAdmin(&_PermImpl.CallOpts, _account, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(_account address, _orgId string) constant returns(bool)
func (_PermImpl *PermImplCaller) ValidateAccount(opts *bind.CallOpts, _account common.Address, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermImpl.contract.Call(opts, out, "validateAccount", _account, _orgId)
	return *ret0, err
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(_account address, _orgId string) constant returns(bool)
func (_PermImpl *PermImplSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _PermImpl.Contract.ValidateAccount(&_PermImpl.CallOpts, _account, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(_account address, _orgId string) constant returns(bool)
func (_PermImpl *PermImplCallerSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _PermImpl.Contract.ValidateAccount(&_PermImpl.CallOpts, _account, _orgId)
}

// AddAdminAccounts is a paid mutator transaction binding the contract method 0x71f57931.
//
// Solidity: function addAdminAccounts(_acct address) returns()
func (_PermImpl *PermImplTransactor) AddAdminAccounts(opts *bind.TransactOpts, _acct common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addAdminAccounts", _acct)
}

// AddAdminAccounts is a paid mutator transaction binding the contract method 0x71f57931.
//
// Solidity: function addAdminAccounts(_acct address) returns()
func (_PermImpl *PermImplSession) AddAdminAccounts(_acct common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddAdminAccounts(&_PermImpl.TransactOpts, _acct)
}

// AddAdminAccounts is a paid mutator transaction binding the contract method 0x71f57931.
//
// Solidity: function addAdminAccounts(_acct address) returns()
func (_PermImpl *PermImplTransactorSession) AddAdminAccounts(_acct common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddAdminAccounts(&_PermImpl.TransactOpts, _acct)
}

// AddAdminNodes is a paid mutator transaction binding the contract method 0xe5e5b85d.
//
// Solidity: function addAdminNodes(_enodeId string) returns()
func (_PermImpl *PermImplTransactor) AddAdminNodes(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addAdminNodes", _enodeId)
}

// AddAdminNodes is a paid mutator transaction binding the contract method 0xe5e5b85d.
//
// Solidity: function addAdminNodes(_enodeId string) returns()
func (_PermImpl *PermImplSession) AddAdminNodes(_enodeId string) (*types.Transaction, error) {
	return _PermImpl.Contract.AddAdminNodes(&_PermImpl.TransactOpts, _enodeId)
}

// AddAdminNodes is a paid mutator transaction binding the contract method 0xe5e5b85d.
//
// Solidity: function addAdminNodes(_enodeId string) returns()
func (_PermImpl *PermImplTransactorSession) AddAdminNodes(_enodeId string) (*types.Transaction, error) {
	return _PermImpl.Contract.AddAdminNodes(&_PermImpl.TransactOpts, _enodeId)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x1b04c276.
//
// Solidity: function addNewRole(_roleId string, _orgId string, _access uint256, _voter bool, _admin bool, _caller address) returns()
func (_PermImpl *PermImplTransactor) AddNewRole(opts *bind.TransactOpts, _roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addNewRole", _roleId, _orgId, _access, _voter, _admin, _caller)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x1b04c276.
//
// Solidity: function addNewRole(_roleId string, _orgId string, _access uint256, _voter bool, _admin bool, _caller address) returns()
func (_PermImpl *PermImplSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddNewRole(&_PermImpl.TransactOpts, _roleId, _orgId, _access, _voter, _admin, _caller)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x1b04c276.
//
// Solidity: function addNewRole(_roleId string, _orgId string, _access uint256, _voter bool, _admin bool, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _admin bool, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddNewRole(&_PermImpl.TransactOpts, _roleId, _orgId, _access, _voter, _admin, _caller)
}

// AddNode is a paid mutator transaction binding the contract method 0x59a260a3.
//
// Solidity: function addNode(_orgId string, _enodeId string, _caller address) returns()
func (_PermImpl *PermImplTransactor) AddNode(opts *bind.TransactOpts, _orgId string, _enodeId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addNode", _orgId, _enodeId, _caller)
}

// AddNode is a paid mutator transaction binding the contract method 0x59a260a3.
//
// Solidity: function addNode(_orgId string, _enodeId string, _caller address) returns()
func (_PermImpl *PermImplSession) AddNode(_orgId string, _enodeId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddNode(&_PermImpl.TransactOpts, _orgId, _enodeId, _caller)
}

// AddNode is a paid mutator transaction binding the contract method 0x59a260a3.
//
// Solidity: function addNode(_orgId string, _enodeId string, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) AddNode(_orgId string, _enodeId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddNode(&_PermImpl.TransactOpts, _orgId, _enodeId, _caller)
}

// AddOrg is a paid mutator transaction binding the contract method 0xf922f802.
//
// Solidity: function addOrg(_orgId string, _enodeId string, _account address, _caller address) returns()
func (_PermImpl *PermImplTransactor) AddOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addOrg", _orgId, _enodeId, _account, _caller)
}

// AddOrg is a paid mutator transaction binding the contract method 0xf922f802.
//
// Solidity: function addOrg(_orgId string, _enodeId string, _account address, _caller address) returns()
func (_PermImpl *PermImplSession) AddOrg(_orgId string, _enodeId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _account, _caller)
}

// AddOrg is a paid mutator transaction binding the contract method 0xf922f802.
//
// Solidity: function addOrg(_orgId string, _enodeId string, _account address, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) AddOrg(_orgId string, _enodeId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _account, _caller)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x90894f0d.
//
// Solidity: function addSubOrg(_pOrg string, _orgId string, _enodeId string, _account address, _caller address) returns()
func (_PermImpl *PermImplTransactor) AddSubOrg(opts *bind.TransactOpts, _pOrg string, _orgId string, _enodeId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addSubOrg", _pOrg, _orgId, _enodeId, _account, _caller)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x90894f0d.
//
// Solidity: function addSubOrg(_pOrg string, _orgId string, _enodeId string, _account address, _caller address) returns()
func (_PermImpl *PermImplSession) AddSubOrg(_pOrg string, _orgId string, _enodeId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddSubOrg(&_PermImpl.TransactOpts, _pOrg, _orgId, _enodeId, _account, _caller)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x90894f0d.
//
// Solidity: function addSubOrg(_pOrg string, _orgId string, _enodeId string, _account address, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) AddSubOrg(_pOrg string, _orgId string, _enodeId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddSubOrg(&_PermImpl.TransactOpts, _pOrg, _orgId, _enodeId, _account, _caller)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x88843041.
//
// Solidity: function approveAdminRole(_orgId string, _account address, _caller address) returns()
func (_PermImpl *PermImplTransactor) ApproveAdminRole(opts *bind.TransactOpts, _orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "approveAdminRole", _orgId, _account, _caller)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x88843041.
//
// Solidity: function approveAdminRole(_orgId string, _account address, _caller address) returns()
func (_PermImpl *PermImplSession) ApproveAdminRole(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveAdminRole(&_PermImpl.TransactOpts, _orgId, _account, _caller)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x88843041.
//
// Solidity: function approveAdminRole(_orgId string, _account address, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) ApproveAdminRole(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveAdminRole(&_PermImpl.TransactOpts, _orgId, _account, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x3bc07dea.
//
// Solidity: function approveOrg(_orgId string, _enodeId string, _account address, _caller address) returns()
func (_PermImpl *PermImplTransactor) ApproveOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "approveOrg", _orgId, _enodeId, _account, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x3bc07dea.
//
// Solidity: function approveOrg(_orgId string, _enodeId string, _account address, _caller address) returns()
func (_PermImpl *PermImplSession) ApproveOrg(_orgId string, _enodeId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _account, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x3bc07dea.
//
// Solidity: function approveOrg(_orgId string, _enodeId string, _account address, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) ApproveOrg(_orgId string, _enodeId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _account, _caller)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0xb5546564.
//
// Solidity: function approveOrgStatus(_orgId string, _status uint256, _caller address) returns()
func (_PermImpl *PermImplTransactor) ApproveOrgStatus(opts *bind.TransactOpts, _orgId string, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "approveOrgStatus", _orgId, _status, _caller)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0xb5546564.
//
// Solidity: function approveOrgStatus(_orgId string, _status uint256, _caller address) returns()
func (_PermImpl *PermImplSession) ApproveOrgStatus(_orgId string, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrgStatus(&_PermImpl.TransactOpts, _orgId, _status, _caller)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0xb5546564.
//
// Solidity: function approveOrgStatus(_orgId string, _status uint256, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) ApproveOrgStatus(_orgId string, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrgStatus(&_PermImpl.TransactOpts, _orgId, _status, _caller)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x8baa8191.
//
// Solidity: function assignAccountRole(_acct address, _orgId string, _roleId string, _caller address) returns()
func (_PermImpl *PermImplTransactor) AssignAccountRole(opts *bind.TransactOpts, _acct common.Address, _orgId string, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "assignAccountRole", _acct, _orgId, _roleId, _caller)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x8baa8191.
//
// Solidity: function assignAccountRole(_acct address, _orgId string, _roleId string, _caller address) returns()
func (_PermImpl *PermImplSession) AssignAccountRole(_acct common.Address, _orgId string, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AssignAccountRole(&_PermImpl.TransactOpts, _acct, _orgId, _roleId, _caller)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x8baa8191.
//
// Solidity: function assignAccountRole(_acct address, _orgId string, _roleId string, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) AssignAccountRole(_acct common.Address, _orgId string, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AssignAccountRole(&_PermImpl.TransactOpts, _acct, _orgId, _roleId, _caller)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x404bf3eb.
//
// Solidity: function assignAdminRole(_orgId string, _account address, _roleId string, _caller address) returns()
func (_PermImpl *PermImplTransactor) AssignAdminRole(opts *bind.TransactOpts, _orgId string, _account common.Address, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "assignAdminRole", _orgId, _account, _roleId, _caller)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x404bf3eb.
//
// Solidity: function assignAdminRole(_orgId string, _account address, _roleId string, _caller address) returns()
func (_PermImpl *PermImplSession) AssignAdminRole(_orgId string, _account common.Address, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AssignAdminRole(&_PermImpl.TransactOpts, _orgId, _account, _roleId, _caller)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0x404bf3eb.
//
// Solidity: function assignAdminRole(_orgId string, _account address, _roleId string, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) AssignAdminRole(_orgId string, _account common.Address, _roleId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AssignAdminRole(&_PermImpl.TransactOpts, _orgId, _account, _roleId, _caller)
}

// Init is a paid mutator transaction binding the contract method 0x0fd07ea4.
//
// Solidity: function init(_orgManager address, _rolesManager address, _acctManager address, _voterManager address, _nodeManager address, _breadth uint256, _depth uint256) returns()
func (_PermImpl *PermImplTransactor) Init(opts *bind.TransactOpts, _orgManager common.Address, _rolesManager common.Address, _acctManager common.Address, _voterManager common.Address, _nodeManager common.Address, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "init", _orgManager, _rolesManager, _acctManager, _voterManager, _nodeManager, _breadth, _depth)
}

// Init is a paid mutator transaction binding the contract method 0x0fd07ea4.
//
// Solidity: function init(_orgManager address, _rolesManager address, _acctManager address, _voterManager address, _nodeManager address, _breadth uint256, _depth uint256) returns()
func (_PermImpl *PermImplSession) Init(_orgManager common.Address, _rolesManager common.Address, _acctManager common.Address, _voterManager common.Address, _nodeManager common.Address, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermImpl.Contract.Init(&_PermImpl.TransactOpts, _orgManager, _rolesManager, _acctManager, _voterManager, _nodeManager, _breadth, _depth)
}

// Init is a paid mutator transaction binding the contract method 0x0fd07ea4.
//
// Solidity: function init(_orgManager address, _rolesManager address, _acctManager address, _voterManager address, _nodeManager address, _breadth uint256, _depth uint256) returns()
func (_PermImpl *PermImplTransactorSession) Init(_orgManager common.Address, _rolesManager common.Address, _acctManager common.Address, _voterManager common.Address, _nodeManager common.Address, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _PermImpl.Contract.Init(&_PermImpl.TransactOpts, _orgManager, _rolesManager, _acctManager, _voterManager, _nodeManager, _breadth, _depth)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x5ca5adbe.
//
// Solidity: function removeRole(_roleId string, _orgId string, _caller address) returns()
func (_PermImpl *PermImplTransactor) RemoveRole(opts *bind.TransactOpts, _roleId string, _orgId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "removeRole", _roleId, _orgId, _caller)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x5ca5adbe.
//
// Solidity: function removeRole(_roleId string, _orgId string, _caller address) returns()
func (_PermImpl *PermImplSession) RemoveRole(_roleId string, _orgId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.RemoveRole(&_PermImpl.TransactOpts, _roleId, _orgId, _caller)
}

// RemoveRole is a paid mutator transaction binding the contract method 0x5ca5adbe.
//
// Solidity: function removeRole(_roleId string, _orgId string, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) RemoveRole(_roleId string, _orgId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.RemoveRole(&_PermImpl.TransactOpts, _roleId, _orgId, _caller)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) returns()
func (_PermImpl *PermImplTransactor) SetPolicy(opts *bind.TransactOpts, _nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "setPolicy", _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) returns()
func (_PermImpl *PermImplSession) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermImpl.Contract.SetPolicy(&_PermImpl.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) returns()
func (_PermImpl *PermImplTransactorSession) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermImpl.Contract.SetPolicy(&_PermImpl.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x04e81f1e.
//
// Solidity: function updateAccountStatus(_orgId string, _account address, _status uint256, _caller address) returns()
func (_PermImpl *PermImplTransactor) UpdateAccountStatus(opts *bind.TransactOpts, _orgId string, _account common.Address, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "updateAccountStatus", _orgId, _account, _status, _caller)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x04e81f1e.
//
// Solidity: function updateAccountStatus(_orgId string, _account address, _status uint256, _caller address) returns()
func (_PermImpl *PermImplSession) UpdateAccountStatus(_orgId string, _account common.Address, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateAccountStatus(&_PermImpl.TransactOpts, _orgId, _account, _status, _caller)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x04e81f1e.
//
// Solidity: function updateAccountStatus(_orgId string, _account address, _status uint256, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) UpdateAccountStatus(_orgId string, _account common.Address, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateAccountStatus(&_PermImpl.TransactOpts, _orgId, _account, _status, _caller)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermImpl *PermImplTransactor) UpdateNetworkBootStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "updateNetworkBootStatus")
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermImpl *PermImplSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateNetworkBootStatus(&_PermImpl.TransactOpts)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermImpl *PermImplTransactorSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateNetworkBootStatus(&_PermImpl.TransactOpts)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0xdbfad711.
//
// Solidity: function updateNodeStatus(_orgId string, _enodeId string, _status uint256, _caller address) returns()
func (_PermImpl *PermImplTransactor) UpdateNodeStatus(opts *bind.TransactOpts, _orgId string, _enodeId string, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "updateNodeStatus", _orgId, _enodeId, _status, _caller)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0xdbfad711.
//
// Solidity: function updateNodeStatus(_orgId string, _enodeId string, _status uint256, _caller address) returns()
func (_PermImpl *PermImplSession) UpdateNodeStatus(_orgId string, _enodeId string, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateNodeStatus(&_PermImpl.TransactOpts, _orgId, _enodeId, _status, _caller)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0xdbfad711.
//
// Solidity: function updateNodeStatus(_orgId string, _enodeId string, _status uint256, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) UpdateNodeStatus(_orgId string, _enodeId string, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateNodeStatus(&_PermImpl.TransactOpts, _orgId, _enodeId, _status, _caller)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0x3cf5f33b.
//
// Solidity: function updateOrgStatus(_orgId string, _status uint256, _caller address) returns()
func (_PermImpl *PermImplTransactor) UpdateOrgStatus(opts *bind.TransactOpts, _orgId string, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "updateOrgStatus", _orgId, _status, _caller)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0x3cf5f33b.
//
// Solidity: function updateOrgStatus(_orgId string, _status uint256, _caller address) returns()
func (_PermImpl *PermImplSession) UpdateOrgStatus(_orgId string, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateOrgStatus(&_PermImpl.TransactOpts, _orgId, _status, _caller)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0x3cf5f33b.
//
// Solidity: function updateOrgStatus(_orgId string, _status uint256, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) UpdateOrgStatus(_orgId string, _status *big.Int, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.UpdateOrgStatus(&_PermImpl.TransactOpts, _orgId, _status, _caller)
}
