// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permission

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

// AcctManagerABI is the input ABI used to generate the binding from.
const AcctManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"approveAdminRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"}],\"name\":\"getAccountDetails\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_roleId\",\"type\":\"string\"}],\"name\":\"assignAccountRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfAccounts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"validateAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"}],\"name\":\"getAccountRole\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"updateAccountStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"orgAdminExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"aIndex\",\"type\":\"uint256\"}],\"name\":\"getAccountDetailsFromIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nwAdminRole\",\"type\":\"string\"},{\"name\":\"_oAdminRole\",\"type\":\"string\"}],\"name\":\"setDefaults\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"revokeAccountRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"assignAdminRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_ultParent\",\"type\":\"string\"}],\"name\":\"checkOrgAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgAdmin\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"AccountAccessModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgAdmin\",\"type\":\"bool\"}],\"name\":\"AccountAccessRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"AccountStatusChanged\",\"type\":\"event\"}]"

// AcctManager is an auto generated Go binding around an Ethereum contract.
type AcctManager struct {
	AcctManagerCaller     // Read-only binding to the contract
	AcctManagerTransactor // Write-only binding to the contract
	AcctManagerFilterer   // Log filterer for contract events
}

// AcctManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AcctManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcctManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AcctManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcctManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AcctManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AcctManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AcctManagerSession struct {
	Contract     *AcctManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AcctManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AcctManagerCallerSession struct {
	Contract *AcctManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AcctManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AcctManagerTransactorSession struct {
	Contract     *AcctManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AcctManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AcctManagerRaw struct {
	Contract *AcctManager // Generic contract binding to access the raw methods on
}

// AcctManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AcctManagerCallerRaw struct {
	Contract *AcctManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AcctManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AcctManagerTransactorRaw struct {
	Contract *AcctManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAcctManager creates a new instance of AcctManager, bound to a specific deployed contract.
func NewAcctManager(address common.Address, backend bind.ContractBackend) (*AcctManager, error) {
	contract, err := bindAcctManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AcctManager{AcctManagerCaller: AcctManagerCaller{contract: contract}, AcctManagerTransactor: AcctManagerTransactor{contract: contract}, AcctManagerFilterer: AcctManagerFilterer{contract: contract}}, nil
}

// NewAcctManagerCaller creates a new read-only instance of AcctManager, bound to a specific deployed contract.
func NewAcctManagerCaller(address common.Address, caller bind.ContractCaller) (*AcctManagerCaller, error) {
	contract, err := bindAcctManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AcctManagerCaller{contract: contract}, nil
}

// NewAcctManagerTransactor creates a new write-only instance of AcctManager, bound to a specific deployed contract.
func NewAcctManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AcctManagerTransactor, error) {
	contract, err := bindAcctManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AcctManagerTransactor{contract: contract}, nil
}

// NewAcctManagerFilterer creates a new log filterer instance of AcctManager, bound to a specific deployed contract.
func NewAcctManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AcctManagerFilterer, error) {
	contract, err := bindAcctManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AcctManagerFilterer{contract: contract}, nil
}

// bindAcctManager binds a generic wrapper to an already deployed contract.
func bindAcctManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AcctManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AcctManager *AcctManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AcctManager.Contract.AcctManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AcctManager *AcctManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AcctManager.Contract.AcctManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AcctManager *AcctManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AcctManager.Contract.AcctManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AcctManager *AcctManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AcctManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AcctManager *AcctManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AcctManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AcctManager *AcctManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AcctManager.Contract.contract.Transact(opts, method, params...)
}

// CheckOrgAdmin is a free data retrieval call binding the contract method 0xe8b42bf4.
//
// Solidity: function checkOrgAdmin(_acct address, _orgId string, _ultParent string) constant returns(bool)
func (_AcctManager *AcctManagerCaller) CheckOrgAdmin(opts *bind.CallOpts, _acct common.Address, _orgId string, _ultParent string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AcctManager.contract.Call(opts, out, "checkOrgAdmin", _acct, _orgId, _ultParent)
	return *ret0, err
}

// CheckOrgAdmin is a free data retrieval call binding the contract method 0xe8b42bf4.
//
// Solidity: function checkOrgAdmin(_acct address, _orgId string, _ultParent string) constant returns(bool)
func (_AcctManager *AcctManagerSession) CheckOrgAdmin(_acct common.Address, _orgId string, _ultParent string) (bool, error) {
	return _AcctManager.Contract.CheckOrgAdmin(&_AcctManager.CallOpts, _acct, _orgId, _ultParent)
}

// CheckOrgAdmin is a free data retrieval call binding the contract method 0xe8b42bf4.
//
// Solidity: function checkOrgAdmin(_acct address, _orgId string, _ultParent string) constant returns(bool)
func (_AcctManager *AcctManagerCallerSession) CheckOrgAdmin(_acct common.Address, _orgId string, _ultParent string) (bool, error) {
	return _AcctManager.Contract.CheckOrgAdmin(&_AcctManager.CallOpts, _acct, _orgId, _ultParent)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(_acct address) constant returns(address, string, string, uint256, bool)
func (_AcctManager *AcctManagerCaller) GetAccountDetails(opts *bind.CallOpts, _acct common.Address) (common.Address, string, string, *big.Int, bool, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(*big.Int)
		ret4 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _AcctManager.contract.Call(opts, out, "getAccountDetails", _acct)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(_acct address) constant returns(address, string, string, uint256, bool)
func (_AcctManager *AcctManagerSession) GetAccountDetails(_acct common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return _AcctManager.Contract.GetAccountDetails(&_AcctManager.CallOpts, _acct)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(_acct address) constant returns(address, string, string, uint256, bool)
func (_AcctManager *AcctManagerCallerSession) GetAccountDetails(_acct common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return _AcctManager.Contract.GetAccountDetails(&_AcctManager.CallOpts, _acct)
}

// GetAccountDetailsFromIndex is a free data retrieval call binding the contract method 0xb2018568.
//
// Solidity: function getAccountDetailsFromIndex(aIndex uint256) constant returns(address, string, string, uint256, bool)
func (_AcctManager *AcctManagerCaller) GetAccountDetailsFromIndex(opts *bind.CallOpts, aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(*big.Int)
		ret4 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _AcctManager.contract.Call(opts, out, "getAccountDetailsFromIndex", aIndex)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetAccountDetailsFromIndex is a free data retrieval call binding the contract method 0xb2018568.
//
// Solidity: function getAccountDetailsFromIndex(aIndex uint256) constant returns(address, string, string, uint256, bool)
func (_AcctManager *AcctManagerSession) GetAccountDetailsFromIndex(aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {
	return _AcctManager.Contract.GetAccountDetailsFromIndex(&_AcctManager.CallOpts, aIndex)
}

// GetAccountDetailsFromIndex is a free data retrieval call binding the contract method 0xb2018568.
//
// Solidity: function getAccountDetailsFromIndex(aIndex uint256) constant returns(address, string, string, uint256, bool)
func (_AcctManager *AcctManagerCallerSession) GetAccountDetailsFromIndex(aIndex *big.Int) (common.Address, string, string, *big.Int, bool, error) {
	return _AcctManager.Contract.GetAccountDetailsFromIndex(&_AcctManager.CallOpts, aIndex)
}

// GetAccountRole is a free data retrieval call binding the contract method 0x81d66b23.
//
// Solidity: function getAccountRole(_acct address) constant returns(string)
func (_AcctManager *AcctManagerCaller) GetAccountRole(opts *bind.CallOpts, _acct common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AcctManager.contract.Call(opts, out, "getAccountRole", _acct)
	return *ret0, err
}

// GetAccountRole is a free data retrieval call binding the contract method 0x81d66b23.
//
// Solidity: function getAccountRole(_acct address) constant returns(string)
func (_AcctManager *AcctManagerSession) GetAccountRole(_acct common.Address) (string, error) {
	return _AcctManager.Contract.GetAccountRole(&_AcctManager.CallOpts, _acct)
}

// GetAccountRole is a free data retrieval call binding the contract method 0x81d66b23.
//
// Solidity: function getAccountRole(_acct address) constant returns(string)
func (_AcctManager *AcctManagerCallerSession) GetAccountRole(_acct common.Address) (string, error) {
	return _AcctManager.Contract.GetAccountRole(&_AcctManager.CallOpts, _acct)
}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() constant returns(uint256)
func (_AcctManager *AcctManagerCaller) GetNumberOfAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AcctManager.contract.Call(opts, out, "getNumberOfAccounts")
	return *ret0, err
}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() constant returns(uint256)
func (_AcctManager *AcctManagerSession) GetNumberOfAccounts() (*big.Int, error) {
	return _AcctManager.Contract.GetNumberOfAccounts(&_AcctManager.CallOpts)
}

// GetNumberOfAccounts is a free data retrieval call binding the contract method 0x309e36ef.
//
// Solidity: function getNumberOfAccounts() constant returns(uint256)
func (_AcctManager *AcctManagerCallerSession) GetNumberOfAccounts() (*big.Int, error) {
	return _AcctManager.Contract.GetNumberOfAccounts(&_AcctManager.CallOpts)
}

// OrgAdminExists is a free data retrieval call binding the contract method 0x950145cf.
//
// Solidity: function orgAdminExists(_orgId string) constant returns(bool)
func (_AcctManager *AcctManagerCaller) OrgAdminExists(opts *bind.CallOpts, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AcctManager.contract.Call(opts, out, "orgAdminExists", _orgId)
	return *ret0, err
}

// OrgAdminExists is a free data retrieval call binding the contract method 0x950145cf.
//
// Solidity: function orgAdminExists(_orgId string) constant returns(bool)
func (_AcctManager *AcctManagerSession) OrgAdminExists(_orgId string) (bool, error) {
	return _AcctManager.Contract.OrgAdminExists(&_AcctManager.CallOpts, _orgId)
}

// OrgAdminExists is a free data retrieval call binding the contract method 0x950145cf.
//
// Solidity: function orgAdminExists(_orgId string) constant returns(bool)
func (_AcctManager *AcctManagerCallerSession) OrgAdminExists(_orgId string) (bool, error) {
	return _AcctManager.Contract.OrgAdminExists(&_AcctManager.CallOpts, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(_acct address, _orgId string) constant returns(bool)
func (_AcctManager *AcctManagerCaller) ValidateAccount(opts *bind.CallOpts, _acct common.Address, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AcctManager.contract.Call(opts, out, "validateAccount", _acct, _orgId)
	return *ret0, err
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(_acct address, _orgId string) constant returns(bool)
func (_AcctManager *AcctManagerSession) ValidateAccount(_acct common.Address, _orgId string) (bool, error) {
	return _AcctManager.Contract.ValidateAccount(&_AcctManager.CallOpts, _acct, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(_acct address, _orgId string) constant returns(bool)
func (_AcctManager *AcctManagerCallerSession) ValidateAccount(_acct common.Address, _orgId string) (bool, error) {
	return _AcctManager.Contract.ValidateAccount(&_AcctManager.CallOpts, _acct, _orgId)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x16724c44.
//
// Solidity: function approveAdminRole(_orgId string, _address address) returns()
func (_AcctManager *AcctManagerTransactor) ApproveAdminRole(opts *bind.TransactOpts, _orgId string, _address common.Address) (*types.Transaction, error) {
	return _AcctManager.contract.Transact(opts, "approveAdminRole", _orgId, _address)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x16724c44.
//
// Solidity: function approveAdminRole(_orgId string, _address address) returns()
func (_AcctManager *AcctManagerSession) ApproveAdminRole(_orgId string, _address common.Address) (*types.Transaction, error) {
	return _AcctManager.Contract.ApproveAdminRole(&_AcctManager.TransactOpts, _orgId, _address)
}

// ApproveAdminRole is a paid mutator transaction binding the contract method 0x16724c44.
//
// Solidity: function approveAdminRole(_orgId string, _address address) returns()
func (_AcctManager *AcctManagerTransactorSession) ApproveAdminRole(_orgId string, _address common.Address) (*types.Transaction, error) {
	return _AcctManager.Contract.ApproveAdminRole(&_AcctManager.TransactOpts, _orgId, _address)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x2f7f0a12.
//
// Solidity: function assignAccountRole(_address address, _orgId string, _roleId string) returns()
func (_AcctManager *AcctManagerTransactor) AssignAccountRole(opts *bind.TransactOpts, _address common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return _AcctManager.contract.Transact(opts, "assignAccountRole", _address, _orgId, _roleId)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x2f7f0a12.
//
// Solidity: function assignAccountRole(_address address, _orgId string, _roleId string) returns()
func (_AcctManager *AcctManagerSession) AssignAccountRole(_address common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return _AcctManager.Contract.AssignAccountRole(&_AcctManager.TransactOpts, _address, _orgId, _roleId)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x2f7f0a12.
//
// Solidity: function assignAccountRole(_address address, _orgId string, _roleId string) returns()
func (_AcctManager *AcctManagerTransactorSession) AssignAccountRole(_address common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return _AcctManager.Contract.AssignAccountRole(&_AcctManager.TransactOpts, _address, _orgId, _roleId)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0xe3483a9d.
//
// Solidity: function assignAdminRole(_address address, _orgId string, _roleId string, _status uint256) returns()
func (_AcctManager *AcctManagerTransactor) AssignAdminRole(opts *bind.TransactOpts, _address common.Address, _orgId string, _roleId string, _status *big.Int) (*types.Transaction, error) {
	return _AcctManager.contract.Transact(opts, "assignAdminRole", _address, _orgId, _roleId, _status)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0xe3483a9d.
//
// Solidity: function assignAdminRole(_address address, _orgId string, _roleId string, _status uint256) returns()
func (_AcctManager *AcctManagerSession) AssignAdminRole(_address common.Address, _orgId string, _roleId string, _status *big.Int) (*types.Transaction, error) {
	return _AcctManager.Contract.AssignAdminRole(&_AcctManager.TransactOpts, _address, _orgId, _roleId, _status)
}

// AssignAdminRole is a paid mutator transaction binding the contract method 0xe3483a9d.
//
// Solidity: function assignAdminRole(_address address, _orgId string, _roleId string, _status uint256) returns()
func (_AcctManager *AcctManagerTransactorSession) AssignAdminRole(_address common.Address, _orgId string, _roleId string, _status *big.Int) (*types.Transaction, error) {
	return _AcctManager.Contract.AssignAdminRole(&_AcctManager.TransactOpts, _address, _orgId, _roleId, _status)
}

// RevokeAccountRole is a paid mutator transaction binding the contract method 0xe163dcf5.
//
// Solidity: function revokeAccountRole(_address address) returns()
func (_AcctManager *AcctManagerTransactor) RevokeAccountRole(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _AcctManager.contract.Transact(opts, "revokeAccountRole", _address)
}

// RevokeAccountRole is a paid mutator transaction binding the contract method 0xe163dcf5.
//
// Solidity: function revokeAccountRole(_address address) returns()
func (_AcctManager *AcctManagerSession) RevokeAccountRole(_address common.Address) (*types.Transaction, error) {
	return _AcctManager.Contract.RevokeAccountRole(&_AcctManager.TransactOpts, _address)
}

// RevokeAccountRole is a paid mutator transaction binding the contract method 0xe163dcf5.
//
// Solidity: function revokeAccountRole(_address address) returns()
func (_AcctManager *AcctManagerTransactorSession) RevokeAccountRole(_address common.Address) (*types.Transaction, error) {
	return _AcctManager.Contract.RevokeAccountRole(&_AcctManager.TransactOpts, _address)
}

// SetDefaults is a paid mutator transaction binding the contract method 0xcef7f6af.
//
// Solidity: function setDefaults(_nwAdminRole string, _oAdminRole string) returns()
func (_AcctManager *AcctManagerTransactor) SetDefaults(opts *bind.TransactOpts, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _AcctManager.contract.Transact(opts, "setDefaults", _nwAdminRole, _oAdminRole)
}

// SetDefaults is a paid mutator transaction binding the contract method 0xcef7f6af.
//
// Solidity: function setDefaults(_nwAdminRole string, _oAdminRole string) returns()
func (_AcctManager *AcctManagerSession) SetDefaults(_nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _AcctManager.Contract.SetDefaults(&_AcctManager.TransactOpts, _nwAdminRole, _oAdminRole)
}

// SetDefaults is a paid mutator transaction binding the contract method 0xcef7f6af.
//
// Solidity: function setDefaults(_nwAdminRole string, _oAdminRole string) returns()
func (_AcctManager *AcctManagerTransactorSession) SetDefaults(_nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _AcctManager.Contract.SetDefaults(&_AcctManager.TransactOpts, _nwAdminRole, _oAdminRole)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x84b7a84a.
//
// Solidity: function updateAccountStatus(_orgId string, _account address, _status uint256) returns()
func (_AcctManager *AcctManagerTransactor) UpdateAccountStatus(opts *bind.TransactOpts, _orgId string, _account common.Address, _status *big.Int) (*types.Transaction, error) {
	return _AcctManager.contract.Transact(opts, "updateAccountStatus", _orgId, _account, _status)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x84b7a84a.
//
// Solidity: function updateAccountStatus(_orgId string, _account address, _status uint256) returns()
func (_AcctManager *AcctManagerSession) UpdateAccountStatus(_orgId string, _account common.Address, _status *big.Int) (*types.Transaction, error) {
	return _AcctManager.Contract.UpdateAccountStatus(&_AcctManager.TransactOpts, _orgId, _account, _status)
}

// UpdateAccountStatus is a paid mutator transaction binding the contract method 0x84b7a84a.
//
// Solidity: function updateAccountStatus(_orgId string, _account address, _status uint256) returns()
func (_AcctManager *AcctManagerTransactorSession) UpdateAccountStatus(_orgId string, _account common.Address, _status *big.Int) (*types.Transaction, error) {
	return _AcctManager.Contract.UpdateAccountStatus(&_AcctManager.TransactOpts, _orgId, _account, _status)
}

// AcctManagerAccountAccessModifiedIterator is returned from FilterAccountAccessModified and is used to iterate over the raw logs and unpacked data for AccountAccessModified events raised by the AcctManager contract.
type AcctManagerAccountAccessModifiedIterator struct {
	Event *AcctManagerAccountAccessModified // Event containing the contract specifics and raw log

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
func (it *AcctManagerAccountAccessModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcctManagerAccountAccessModified)
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
		it.Event = new(AcctManagerAccountAccessModified)
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
func (it *AcctManagerAccountAccessModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcctManagerAccountAccessModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcctManagerAccountAccessModified represents a AccountAccessModified event raised by the AcctManager contract.
type AcctManagerAccountAccessModified struct {
	Address  common.Address
	OrgId    string
	RoleId   string
	OrgAdmin bool
	Status   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccountAccessModified is a free log retrieval operation binding the contract event 0x68e62a03aeb0a125c2fc869eed72f2fca473680987bdd680c093a534e17cc776.
//
// Solidity: e AccountAccessModified(_address address, _orgId string, _roleId string, _orgAdmin bool, _status uint256)
func (_AcctManager *AcctManagerFilterer) FilterAccountAccessModified(opts *bind.FilterOpts) (*AcctManagerAccountAccessModifiedIterator, error) {

	logs, sub, err := _AcctManager.contract.FilterLogs(opts, "AccountAccessModified")
	if err != nil {
		return nil, err
	}
	return &AcctManagerAccountAccessModifiedIterator{contract: _AcctManager.contract, event: "AccountAccessModified", logs: logs, sub: sub}, nil
}

// WatchAccountAccessModified is a free log subscription operation binding the contract event 0x68e62a03aeb0a125c2fc869eed72f2fca473680987bdd680c093a534e17cc776.
//
// Solidity: e AccountAccessModified(_address address, _orgId string, _roleId string, _orgAdmin bool, _status uint256)
func (_AcctManager *AcctManagerFilterer) WatchAccountAccessModified(opts *bind.WatchOpts, sink chan<- *AcctManagerAccountAccessModified) (event.Subscription, error) {

	logs, sub, err := _AcctManager.contract.WatchLogs(opts, "AccountAccessModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcctManagerAccountAccessModified)
				if err := _AcctManager.contract.UnpackLog(event, "AccountAccessModified", log); err != nil {
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

// AcctManagerAccountAccessRevokedIterator is returned from FilterAccountAccessRevoked and is used to iterate over the raw logs and unpacked data for AccountAccessRevoked events raised by the AcctManager contract.
type AcctManagerAccountAccessRevokedIterator struct {
	Event *AcctManagerAccountAccessRevoked // Event containing the contract specifics and raw log

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
func (it *AcctManagerAccountAccessRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcctManagerAccountAccessRevoked)
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
		it.Event = new(AcctManagerAccountAccessRevoked)
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
func (it *AcctManagerAccountAccessRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcctManagerAccountAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcctManagerAccountAccessRevoked represents a AccountAccessRevoked event raised by the AcctManager contract.
type AcctManagerAccountAccessRevoked struct {
	Address  common.Address
	OrgId    string
	RoleId   string
	OrgAdmin bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccountAccessRevoked is a free log retrieval operation binding the contract event 0x6b5105396435a8a139aeed682dd573cd2a7e6279de77f8c11f95a30399212ad1.
//
// Solidity: e AccountAccessRevoked(_address address, _orgId string, _roleId string, _orgAdmin bool)
func (_AcctManager *AcctManagerFilterer) FilterAccountAccessRevoked(opts *bind.FilterOpts) (*AcctManagerAccountAccessRevokedIterator, error) {

	logs, sub, err := _AcctManager.contract.FilterLogs(opts, "AccountAccessRevoked")
	if err != nil {
		return nil, err
	}
	return &AcctManagerAccountAccessRevokedIterator{contract: _AcctManager.contract, event: "AccountAccessRevoked", logs: logs, sub: sub}, nil
}

// WatchAccountAccessRevoked is a free log subscription operation binding the contract event 0x6b5105396435a8a139aeed682dd573cd2a7e6279de77f8c11f95a30399212ad1.
//
// Solidity: e AccountAccessRevoked(_address address, _orgId string, _roleId string, _orgAdmin bool)
func (_AcctManager *AcctManagerFilterer) WatchAccountAccessRevoked(opts *bind.WatchOpts, sink chan<- *AcctManagerAccountAccessRevoked) (event.Subscription, error) {

	logs, sub, err := _AcctManager.contract.WatchLogs(opts, "AccountAccessRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcctManagerAccountAccessRevoked)
				if err := _AcctManager.contract.UnpackLog(event, "AccountAccessRevoked", log); err != nil {
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

// AcctManagerAccountStatusChangedIterator is returned from FilterAccountStatusChanged and is used to iterate over the raw logs and unpacked data for AccountStatusChanged events raised by the AcctManager contract.
type AcctManagerAccountStatusChangedIterator struct {
	Event *AcctManagerAccountStatusChanged // Event containing the contract specifics and raw log

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
func (it *AcctManagerAccountStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AcctManagerAccountStatusChanged)
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
		it.Event = new(AcctManagerAccountStatusChanged)
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
func (it *AcctManagerAccountStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AcctManagerAccountStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AcctManagerAccountStatusChanged represents a AccountStatusChanged event raised by the AcctManager contract.
type AcctManagerAccountStatusChanged struct {
	Address common.Address
	OrgId   string
	Status  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountStatusChanged is a free log retrieval operation binding the contract event 0x36b0ea38154dec5e98b6bf928b971a9db5e8cd4b6946350e9e43fb9848c70b25.
//
// Solidity: e AccountStatusChanged(_address address, _orgId string, _status uint256)
func (_AcctManager *AcctManagerFilterer) FilterAccountStatusChanged(opts *bind.FilterOpts) (*AcctManagerAccountStatusChangedIterator, error) {

	logs, sub, err := _AcctManager.contract.FilterLogs(opts, "AccountStatusChanged")
	if err != nil {
		return nil, err
	}
	return &AcctManagerAccountStatusChangedIterator{contract: _AcctManager.contract, event: "AccountStatusChanged", logs: logs, sub: sub}, nil
}

// WatchAccountStatusChanged is a free log subscription operation binding the contract event 0x36b0ea38154dec5e98b6bf928b971a9db5e8cd4b6946350e9e43fb9848c70b25.
//
// Solidity: e AccountStatusChanged(_address address, _orgId string, _status uint256)
func (_AcctManager *AcctManagerFilterer) WatchAccountStatusChanged(opts *bind.WatchOpts, sink chan<- *AcctManagerAccountStatusChanged) (event.Subscription, error) {

	logs, sub, err := _AcctManager.contract.WatchLogs(opts, "AccountStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AcctManagerAccountStatusChanged)
				if err := _AcctManager.contract.UnpackLog(event, "AccountStatusChanged", log); err != nil {
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
