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

// OrgManagerABI is the input ABI used to generate the binding from.
const OrgManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"updateOrg\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_action\",\"type\":\"uint256\"}],\"name\":\"approveOrgStatusUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getUltimateParent\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pOrgId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addSubOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgIndex\",\"type\":\"uint256\"}],\"name\":\"getOrgInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfOrgs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_orgStatus\",\"type\":\"uint256\"}],\"name\":\"checkOrgStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_breadth\",\"type\":\"uint256\"},{\"name\":\"_depth\",\"type\":\"uint256\"}],\"name\":\"setUpOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"approveOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"checkOrgExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_porgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_ultParent\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_level\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"OrgApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_porgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_ultParent\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_level\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"OrgPendingApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_porgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_ultParent\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_level\",\"type\":\"uint256\"}],\"name\":\"OrgSuspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_porgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_ultParent\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_level\",\"type\":\"uint256\"}],\"name\":\"OrgSuspensionRevoked\",\"type\":\"event\"}]"

// OrgManager is an auto generated Go binding around an Ethereum contract.
type OrgManager struct {
	OrgManagerCaller     // Read-only binding to the contract
	OrgManagerTransactor // Write-only binding to the contract
	OrgManagerFilterer   // Log filterer for contract events
}

// OrgManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrgManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrgManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrgManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrgManagerSession struct {
	Contract     *OrgManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrgManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrgManagerCallerSession struct {
	Contract *OrgManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OrgManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrgManagerTransactorSession struct {
	Contract     *OrgManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OrgManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrgManagerRaw struct {
	Contract *OrgManager // Generic contract binding to access the raw methods on
}

// OrgManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrgManagerCallerRaw struct {
	Contract *OrgManagerCaller // Generic read-only contract binding to access the raw methods on
}

// OrgManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrgManagerTransactorRaw struct {
	Contract *OrgManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrgManager creates a new instance of OrgManager, bound to a specific deployed contract.
func NewOrgManager(address common.Address, backend bind.ContractBackend) (*OrgManager, error) {
	contract, err := bindOrgManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrgManager{OrgManagerCaller: OrgManagerCaller{contract: contract}, OrgManagerTransactor: OrgManagerTransactor{contract: contract}, OrgManagerFilterer: OrgManagerFilterer{contract: contract}}, nil
}

// NewOrgManagerCaller creates a new read-only instance of OrgManager, bound to a specific deployed contract.
func NewOrgManagerCaller(address common.Address, caller bind.ContractCaller) (*OrgManagerCaller, error) {
	contract, err := bindOrgManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrgManagerCaller{contract: contract}, nil
}

// NewOrgManagerTransactor creates a new write-only instance of OrgManager, bound to a specific deployed contract.
func NewOrgManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*OrgManagerTransactor, error) {
	contract, err := bindOrgManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrgManagerTransactor{contract: contract}, nil
}

// NewOrgManagerFilterer creates a new log filterer instance of OrgManager, bound to a specific deployed contract.
func NewOrgManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*OrgManagerFilterer, error) {
	contract, err := bindOrgManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrgManagerFilterer{contract: contract}, nil
}

// bindOrgManager binds a generic wrapper to an already deployed contract.
func bindOrgManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrgManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrgManager *OrgManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrgManager.Contract.OrgManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrgManager *OrgManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrgManager.Contract.OrgManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrgManager *OrgManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrgManager.Contract.OrgManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrgManager *OrgManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrgManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrgManager *OrgManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrgManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrgManager *OrgManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrgManager.Contract.contract.Transact(opts, method, params...)
}

// CheckOrgExists is a free data retrieval call binding the contract method 0xffe40d1d.
//
// Solidity: function checkOrgExists(_orgId string) constant returns(bool)
func (_OrgManager *OrgManagerCaller) CheckOrgExists(opts *bind.CallOpts, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OrgManager.contract.Call(opts, out, "checkOrgExists", _orgId)
	return *ret0, err
}

// CheckOrgExists is a free data retrieval call binding the contract method 0xffe40d1d.
//
// Solidity: function checkOrgExists(_orgId string) constant returns(bool)
func (_OrgManager *OrgManagerSession) CheckOrgExists(_orgId string) (bool, error) {
	return _OrgManager.Contract.CheckOrgExists(&_OrgManager.CallOpts, _orgId)
}

// CheckOrgExists is a free data retrieval call binding the contract method 0xffe40d1d.
//
// Solidity: function checkOrgExists(_orgId string) constant returns(bool)
func (_OrgManager *OrgManagerCallerSession) CheckOrgExists(_orgId string) (bool, error) {
	return _OrgManager.Contract.CheckOrgExists(&_OrgManager.CallOpts, _orgId)
}

// CheckOrgStatus is a free data retrieval call binding the contract method 0x8c8642df.
//
// Solidity: function checkOrgStatus(_orgId string, _orgStatus uint256) constant returns(bool)
func (_OrgManager *OrgManagerCaller) CheckOrgStatus(opts *bind.CallOpts, _orgId string, _orgStatus *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _OrgManager.contract.Call(opts, out, "checkOrgStatus", _orgId, _orgStatus)
	return *ret0, err
}

// CheckOrgStatus is a free data retrieval call binding the contract method 0x8c8642df.
//
// Solidity: function checkOrgStatus(_orgId string, _orgStatus uint256) constant returns(bool)
func (_OrgManager *OrgManagerSession) CheckOrgStatus(_orgId string, _orgStatus *big.Int) (bool, error) {
	return _OrgManager.Contract.CheckOrgStatus(&_OrgManager.CallOpts, _orgId, _orgStatus)
}

// CheckOrgStatus is a free data retrieval call binding the contract method 0x8c8642df.
//
// Solidity: function checkOrgStatus(_orgId string, _orgStatus uint256) constant returns(bool)
func (_OrgManager *OrgManagerCallerSession) CheckOrgStatus(_orgId string, _orgStatus *big.Int) (bool, error) {
	return _OrgManager.Contract.CheckOrgStatus(&_OrgManager.CallOpts, _orgId, _orgStatus)
}

// GetNumberOfOrgs is a free data retrieval call binding the contract method 0x7755ebdd.
//
// Solidity: function getNumberOfOrgs() constant returns(uint256)
func (_OrgManager *OrgManagerCaller) GetNumberOfOrgs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OrgManager.contract.Call(opts, out, "getNumberOfOrgs")
	return *ret0, err
}

// GetNumberOfOrgs is a free data retrieval call binding the contract method 0x7755ebdd.
//
// Solidity: function getNumberOfOrgs() constant returns(uint256)
func (_OrgManager *OrgManagerSession) GetNumberOfOrgs() (*big.Int, error) {
	return _OrgManager.Contract.GetNumberOfOrgs(&_OrgManager.CallOpts)
}

// GetNumberOfOrgs is a free data retrieval call binding the contract method 0x7755ebdd.
//
// Solidity: function getNumberOfOrgs() constant returns(uint256)
func (_OrgManager *OrgManagerCallerSession) GetNumberOfOrgs() (*big.Int, error) {
	return _OrgManager.Contract.GetNumberOfOrgs(&_OrgManager.CallOpts)
}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(_orgIndex uint256) constant returns(string, string, string, uint256, uint256)
func (_OrgManager *OrgManagerCaller) GetOrgInfo(opts *bind.CallOpts, _orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _OrgManager.contract.Call(opts, out, "getOrgInfo", _orgIndex)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(_orgIndex uint256) constant returns(string, string, string, uint256, uint256)
func (_OrgManager *OrgManagerSession) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	return _OrgManager.Contract.GetOrgInfo(&_OrgManager.CallOpts, _orgIndex)
}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(_orgIndex uint256) constant returns(string, string, string, uint256, uint256)
func (_OrgManager *OrgManagerCallerSession) GetOrgInfo(_orgIndex *big.Int) (string, string, string, *big.Int, *big.Int, error) {
	return _OrgManager.Contract.GetOrgInfo(&_OrgManager.CallOpts, _orgIndex)
}

// GetUltimateParent is a free data retrieval call binding the contract method 0x177c8d8a.
//
// Solidity: function getUltimateParent(_orgId string) constant returns(string)
func (_OrgManager *OrgManagerCaller) GetUltimateParent(opts *bind.CallOpts, _orgId string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OrgManager.contract.Call(opts, out, "getUltimateParent", _orgId)
	return *ret0, err
}

// GetUltimateParent is a free data retrieval call binding the contract method 0x177c8d8a.
//
// Solidity: function getUltimateParent(_orgId string) constant returns(string)
func (_OrgManager *OrgManagerSession) GetUltimateParent(_orgId string) (string, error) {
	return _OrgManager.Contract.GetUltimateParent(&_OrgManager.CallOpts, _orgId)
}

// GetUltimateParent is a free data retrieval call binding the contract method 0x177c8d8a.
//
// Solidity: function getUltimateParent(_orgId string) constant returns(string)
func (_OrgManager *OrgManagerCallerSession) GetUltimateParent(_orgId string) (string, error) {
	return _OrgManager.Contract.GetUltimateParent(&_OrgManager.CallOpts, _orgId)
}

// AddOrg is a paid mutator transaction binding the contract method 0xf9953de5.
//
// Solidity: function addOrg(_orgId string) returns()
func (_OrgManager *OrgManagerTransactor) AddOrg(opts *bind.TransactOpts, _orgId string) (*types.Transaction, error) {
	return _OrgManager.contract.Transact(opts, "addOrg", _orgId)
}

// AddOrg is a paid mutator transaction binding the contract method 0xf9953de5.
//
// Solidity: function addOrg(_orgId string) returns()
func (_OrgManager *OrgManagerSession) AddOrg(_orgId string) (*types.Transaction, error) {
	return _OrgManager.Contract.AddOrg(&_OrgManager.TransactOpts, _orgId)
}

// AddOrg is a paid mutator transaction binding the contract method 0xf9953de5.
//
// Solidity: function addOrg(_orgId string) returns()
func (_OrgManager *OrgManagerTransactorSession) AddOrg(_orgId string) (*types.Transaction, error) {
	return _OrgManager.Contract.AddOrg(&_OrgManager.TransactOpts, _orgId)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x1f953480.
//
// Solidity: function addSubOrg(_pOrgId string, _orgId string) returns()
func (_OrgManager *OrgManagerTransactor) AddSubOrg(opts *bind.TransactOpts, _pOrgId string, _orgId string) (*types.Transaction, error) {
	return _OrgManager.contract.Transact(opts, "addSubOrg", _pOrgId, _orgId)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x1f953480.
//
// Solidity: function addSubOrg(_pOrgId string, _orgId string) returns()
func (_OrgManager *OrgManagerSession) AddSubOrg(_pOrgId string, _orgId string) (*types.Transaction, error) {
	return _OrgManager.Contract.AddSubOrg(&_OrgManager.TransactOpts, _pOrgId, _orgId)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x1f953480.
//
// Solidity: function addSubOrg(_pOrgId string, _orgId string) returns()
func (_OrgManager *OrgManagerTransactorSession) AddSubOrg(_pOrgId string, _orgId string) (*types.Transaction, error) {
	return _OrgManager.Contract.AddSubOrg(&_OrgManager.TransactOpts, _pOrgId, _orgId)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xe3028316.
//
// Solidity: function approveOrg(_orgId string) returns()
func (_OrgManager *OrgManagerTransactor) ApproveOrg(opts *bind.TransactOpts, _orgId string) (*types.Transaction, error) {
	return _OrgManager.contract.Transact(opts, "approveOrg", _orgId)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xe3028316.
//
// Solidity: function approveOrg(_orgId string) returns()
func (_OrgManager *OrgManagerSession) ApproveOrg(_orgId string) (*types.Transaction, error) {
	return _OrgManager.Contract.ApproveOrg(&_OrgManager.TransactOpts, _orgId)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0xe3028316.
//
// Solidity: function approveOrg(_orgId string) returns()
func (_OrgManager *OrgManagerTransactorSession) ApproveOrg(_orgId string) (*types.Transaction, error) {
	return _OrgManager.Contract.ApproveOrg(&_OrgManager.TransactOpts, _orgId)
}

// ApproveOrgStatusUpdate is a paid mutator transaction binding the contract method 0x14f775f9.
//
// Solidity: function approveOrgStatusUpdate(_orgId string, _action uint256) returns()
func (_OrgManager *OrgManagerTransactor) ApproveOrgStatusUpdate(opts *bind.TransactOpts, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManager.contract.Transact(opts, "approveOrgStatusUpdate", _orgId, _action)
}

// ApproveOrgStatusUpdate is a paid mutator transaction binding the contract method 0x14f775f9.
//
// Solidity: function approveOrgStatusUpdate(_orgId string, _action uint256) returns()
func (_OrgManager *OrgManagerSession) ApproveOrgStatusUpdate(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManager.Contract.ApproveOrgStatusUpdate(&_OrgManager.TransactOpts, _orgId, _action)
}

// ApproveOrgStatusUpdate is a paid mutator transaction binding the contract method 0x14f775f9.
//
// Solidity: function approveOrgStatusUpdate(_orgId string, _action uint256) returns()
func (_OrgManager *OrgManagerTransactorSession) ApproveOrgStatusUpdate(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManager.Contract.ApproveOrgStatusUpdate(&_OrgManager.TransactOpts, _orgId, _action)
}

// SetUpOrg is a paid mutator transaction binding the contract method 0x9e58eb9f.
//
// Solidity: function setUpOrg(_orgId string, _breadth uint256, _depth uint256) returns()
func (_OrgManager *OrgManagerTransactor) SetUpOrg(opts *bind.TransactOpts, _orgId string, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _OrgManager.contract.Transact(opts, "setUpOrg", _orgId, _breadth, _depth)
}

// SetUpOrg is a paid mutator transaction binding the contract method 0x9e58eb9f.
//
// Solidity: function setUpOrg(_orgId string, _breadth uint256, _depth uint256) returns()
func (_OrgManager *OrgManagerSession) SetUpOrg(_orgId string, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _OrgManager.Contract.SetUpOrg(&_OrgManager.TransactOpts, _orgId, _breadth, _depth)
}

// SetUpOrg is a paid mutator transaction binding the contract method 0x9e58eb9f.
//
// Solidity: function setUpOrg(_orgId string, _breadth uint256, _depth uint256) returns()
func (_OrgManager *OrgManagerTransactorSession) SetUpOrg(_orgId string, _breadth *big.Int, _depth *big.Int) (*types.Transaction, error) {
	return _OrgManager.Contract.SetUpOrg(&_OrgManager.TransactOpts, _orgId, _breadth, _depth)
}

// UpdateOrg is a paid mutator transaction binding the contract method 0x0cc27493.
//
// Solidity: function updateOrg(_orgId string, _action uint256) returns(uint256)
func (_OrgManager *OrgManagerTransactor) UpdateOrg(opts *bind.TransactOpts, _orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManager.contract.Transact(opts, "updateOrg", _orgId, _action)
}

// UpdateOrg is a paid mutator transaction binding the contract method 0x0cc27493.
//
// Solidity: function updateOrg(_orgId string, _action uint256) returns(uint256)
func (_OrgManager *OrgManagerSession) UpdateOrg(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManager.Contract.UpdateOrg(&_OrgManager.TransactOpts, _orgId, _action)
}

// UpdateOrg is a paid mutator transaction binding the contract method 0x0cc27493.
//
// Solidity: function updateOrg(_orgId string, _action uint256) returns(uint256)
func (_OrgManager *OrgManagerTransactorSession) UpdateOrg(_orgId string, _action *big.Int) (*types.Transaction, error) {
	return _OrgManager.Contract.UpdateOrg(&_OrgManager.TransactOpts, _orgId, _action)
}

// OrgManagerOrgApprovedIterator is returned from FilterOrgApproved and is used to iterate over the raw logs and unpacked data for OrgApproved events raised by the OrgManager contract.
type OrgManagerOrgApprovedIterator struct {
	Event *OrgManagerOrgApproved // Event containing the contract specifics and raw log

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
func (it *OrgManagerOrgApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgManagerOrgApproved)
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
		it.Event = new(OrgManagerOrgApproved)
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
func (it *OrgManagerOrgApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgManagerOrgApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgManagerOrgApproved represents a OrgApproved event raised by the OrgManager contract.
type OrgManagerOrgApproved struct {
	OrgId     string
	PorgId    string
	UltParent string
	Level     *big.Int
	Status    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrgApproved is a free log retrieval operation binding the contract event 0xd705723a50859c9cc1d3953e10b8b9478820e7a62927ad3215897ed87b20591c.
//
// Solidity: e OrgApproved(_orgId string, _porgId string, _ultParent string, _level uint256, _status uint256)
func (_OrgManager *OrgManagerFilterer) FilterOrgApproved(opts *bind.FilterOpts) (*OrgManagerOrgApprovedIterator, error) {

	logs, sub, err := _OrgManager.contract.FilterLogs(opts, "OrgApproved")
	if err != nil {
		return nil, err
	}
	return &OrgManagerOrgApprovedIterator{contract: _OrgManager.contract, event: "OrgApproved", logs: logs, sub: sub}, nil
}

// WatchOrgApproved is a free log subscription operation binding the contract event 0xd705723a50859c9cc1d3953e10b8b9478820e7a62927ad3215897ed87b20591c.
//
// Solidity: e OrgApproved(_orgId string, _porgId string, _ultParent string, _level uint256, _status uint256)
func (_OrgManager *OrgManagerFilterer) WatchOrgApproved(opts *bind.WatchOpts, sink chan<- *OrgManagerOrgApproved) (event.Subscription, error) {

	logs, sub, err := _OrgManager.contract.WatchLogs(opts, "OrgApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgManagerOrgApproved)
				if err := _OrgManager.contract.UnpackLog(event, "OrgApproved", log); err != nil {
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

// OrgManagerOrgPendingApprovalIterator is returned from FilterOrgPendingApproval and is used to iterate over the raw logs and unpacked data for OrgPendingApproval events raised by the OrgManager contract.
type OrgManagerOrgPendingApprovalIterator struct {
	Event *OrgManagerOrgPendingApproval // Event containing the contract specifics and raw log

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
func (it *OrgManagerOrgPendingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgManagerOrgPendingApproval)
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
		it.Event = new(OrgManagerOrgPendingApproval)
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
func (it *OrgManagerOrgPendingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgManagerOrgPendingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgManagerOrgPendingApproval represents a OrgPendingApproval event raised by the OrgManager contract.
type OrgManagerOrgPendingApproval struct {
	OrgId     string
	PorgId    string
	UltParent string
	Level     *big.Int
	Status    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrgPendingApproval is a free log retrieval operation binding the contract event 0x0e8b7be64e0c730234ba2cd252b227fb481d7a247ba806d1941144c535bf054b.
//
// Solidity: e OrgPendingApproval(_orgId string, _porgId string, _ultParent string, _level uint256, _status uint256)
func (_OrgManager *OrgManagerFilterer) FilterOrgPendingApproval(opts *bind.FilterOpts) (*OrgManagerOrgPendingApprovalIterator, error) {

	logs, sub, err := _OrgManager.contract.FilterLogs(opts, "OrgPendingApproval")
	if err != nil {
		return nil, err
	}
	return &OrgManagerOrgPendingApprovalIterator{contract: _OrgManager.contract, event: "OrgPendingApproval", logs: logs, sub: sub}, nil
}

// WatchOrgPendingApproval is a free log subscription operation binding the contract event 0x0e8b7be64e0c730234ba2cd252b227fb481d7a247ba806d1941144c535bf054b.
//
// Solidity: e OrgPendingApproval(_orgId string, _porgId string, _ultParent string, _level uint256, _status uint256)
func (_OrgManager *OrgManagerFilterer) WatchOrgPendingApproval(opts *bind.WatchOpts, sink chan<- *OrgManagerOrgPendingApproval) (event.Subscription, error) {

	logs, sub, err := _OrgManager.contract.WatchLogs(opts, "OrgPendingApproval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgManagerOrgPendingApproval)
				if err := _OrgManager.contract.UnpackLog(event, "OrgPendingApproval", log); err != nil {
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

// OrgManagerOrgSuspendedIterator is returned from FilterOrgSuspended and is used to iterate over the raw logs and unpacked data for OrgSuspended events raised by the OrgManager contract.
type OrgManagerOrgSuspendedIterator struct {
	Event *OrgManagerOrgSuspended // Event containing the contract specifics and raw log

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
func (it *OrgManagerOrgSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgManagerOrgSuspended)
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
		it.Event = new(OrgManagerOrgSuspended)
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
func (it *OrgManagerOrgSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgManagerOrgSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgManagerOrgSuspended represents a OrgSuspended event raised by the OrgManager contract.
type OrgManagerOrgSuspended struct {
	OrgId     string
	PorgId    string
	UltParent string
	Level     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrgSuspended is a free log retrieval operation binding the contract event 0x73ccf8d6c8385bf5347269bd59712da33183c1a5e1702494bcdb87d0f4674d96.
//
// Solidity: e OrgSuspended(_orgId string, _porgId string, _ultParent string, _level uint256)
func (_OrgManager *OrgManagerFilterer) FilterOrgSuspended(opts *bind.FilterOpts) (*OrgManagerOrgSuspendedIterator, error) {

	logs, sub, err := _OrgManager.contract.FilterLogs(opts, "OrgSuspended")
	if err != nil {
		return nil, err
	}
	return &OrgManagerOrgSuspendedIterator{contract: _OrgManager.contract, event: "OrgSuspended", logs: logs, sub: sub}, nil
}

// WatchOrgSuspended is a free log subscription operation binding the contract event 0x73ccf8d6c8385bf5347269bd59712da33183c1a5e1702494bcdb87d0f4674d96.
//
// Solidity: e OrgSuspended(_orgId string, _porgId string, _ultParent string, _level uint256)
func (_OrgManager *OrgManagerFilterer) WatchOrgSuspended(opts *bind.WatchOpts, sink chan<- *OrgManagerOrgSuspended) (event.Subscription, error) {

	logs, sub, err := _OrgManager.contract.WatchLogs(opts, "OrgSuspended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgManagerOrgSuspended)
				if err := _OrgManager.contract.UnpackLog(event, "OrgSuspended", log); err != nil {
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

// OrgManagerOrgSuspensionRevokedIterator is returned from FilterOrgSuspensionRevoked and is used to iterate over the raw logs and unpacked data for OrgSuspensionRevoked events raised by the OrgManager contract.
type OrgManagerOrgSuspensionRevokedIterator struct {
	Event *OrgManagerOrgSuspensionRevoked // Event containing the contract specifics and raw log

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
func (it *OrgManagerOrgSuspensionRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgManagerOrgSuspensionRevoked)
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
		it.Event = new(OrgManagerOrgSuspensionRevoked)
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
func (it *OrgManagerOrgSuspensionRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgManagerOrgSuspensionRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgManagerOrgSuspensionRevoked represents a OrgSuspensionRevoked event raised by the OrgManager contract.
type OrgManagerOrgSuspensionRevoked struct {
	OrgId     string
	PorgId    string
	UltParent string
	Level     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrgSuspensionRevoked is a free log retrieval operation binding the contract event 0x882f030c609566cd82918a97d457fd48f9cfcefd11282e2654cde3f94579c15f.
//
// Solidity: e OrgSuspensionRevoked(_orgId string, _porgId string, _ultParent string, _level uint256)
func (_OrgManager *OrgManagerFilterer) FilterOrgSuspensionRevoked(opts *bind.FilterOpts) (*OrgManagerOrgSuspensionRevokedIterator, error) {

	logs, sub, err := _OrgManager.contract.FilterLogs(opts, "OrgSuspensionRevoked")
	if err != nil {
		return nil, err
	}
	return &OrgManagerOrgSuspensionRevokedIterator{contract: _OrgManager.contract, event: "OrgSuspensionRevoked", logs: logs, sub: sub}, nil
}

// WatchOrgSuspensionRevoked is a free log subscription operation binding the contract event 0x882f030c609566cd82918a97d457fd48f9cfcefd11282e2654cde3f94579c15f.
//
// Solidity: e OrgSuspensionRevoked(_orgId string, _porgId string, _ultParent string, _level uint256)
func (_OrgManager *OrgManagerFilterer) WatchOrgSuspensionRevoked(opts *bind.WatchOpts, sink chan<- *OrgManagerOrgSuspensionRevoked) (event.Subscription, error) {

	logs, sub, err := _OrgManager.contract.WatchLogs(opts, "OrgSuspensionRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgManagerOrgSuspensionRevoked)
				if err := _OrgManager.contract.UnpackLog(event, "OrgSuspensionRevoked", log); err != nil {
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
