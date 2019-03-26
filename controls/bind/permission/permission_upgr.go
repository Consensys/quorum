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

// PermUpgrABI is the input ABI used to generate the binding from.
const PermUpgrABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getPermImpl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"confirmImplChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCustodian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPermInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_permInterface\",\"type\":\"address\"},{\"name\":\"_permImpl\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_custodian\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PermUpgr is an auto generated Go binding around an Ethereum contract.
type PermUpgr struct {
	PermUpgrCaller     // Read-only binding to the contract
	PermUpgrTransactor // Write-only binding to the contract
	PermUpgrFilterer   // Log filterer for contract events
}

// PermUpgrCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermUpgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermUpgrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermUpgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermUpgrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermUpgrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermUpgrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermUpgrSession struct {
	Contract     *PermUpgr         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermUpgrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermUpgrCallerSession struct {
	Contract *PermUpgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PermUpgrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermUpgrTransactorSession struct {
	Contract     *PermUpgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PermUpgrRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermUpgrRaw struct {
	Contract *PermUpgr // Generic contract binding to access the raw methods on
}

// PermUpgrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermUpgrCallerRaw struct {
	Contract *PermUpgrCaller // Generic read-only contract binding to access the raw methods on
}

// PermUpgrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermUpgrTransactorRaw struct {
	Contract *PermUpgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermUpgr creates a new instance of PermUpgr, bound to a specific deployed contract.
func NewPermUpgr(address common.Address, backend bind.ContractBackend) (*PermUpgr, error) {
	contract, err := bindPermUpgr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermUpgr{PermUpgrCaller: PermUpgrCaller{contract: contract}, PermUpgrTransactor: PermUpgrTransactor{contract: contract}, PermUpgrFilterer: PermUpgrFilterer{contract: contract}}, nil
}

// NewPermUpgrCaller creates a new read-only instance of PermUpgr, bound to a specific deployed contract.
func NewPermUpgrCaller(address common.Address, caller bind.ContractCaller) (*PermUpgrCaller, error) {
	contract, err := bindPermUpgr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermUpgrCaller{contract: contract}, nil
}

// NewPermUpgrTransactor creates a new write-only instance of PermUpgr, bound to a specific deployed contract.
func NewPermUpgrTransactor(address common.Address, transactor bind.ContractTransactor) (*PermUpgrTransactor, error) {
	contract, err := bindPermUpgr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermUpgrTransactor{contract: contract}, nil
}

// NewPermUpgrFilterer creates a new log filterer instance of PermUpgr, bound to a specific deployed contract.
func NewPermUpgrFilterer(address common.Address, filterer bind.ContractFilterer) (*PermUpgrFilterer, error) {
	contract, err := bindPermUpgr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermUpgrFilterer{contract: contract}, nil
}

// bindPermUpgr binds a generic wrapper to an already deployed contract.
func bindPermUpgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermUpgrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermUpgr *PermUpgrRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermUpgr.Contract.PermUpgrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermUpgr *PermUpgrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermUpgr.Contract.PermUpgrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermUpgr *PermUpgrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermUpgr.Contract.PermUpgrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermUpgr *PermUpgrCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermUpgr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermUpgr *PermUpgrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermUpgr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermUpgr *PermUpgrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermUpgr.Contract.contract.Transact(opts, method, params...)
}

// GetCustodian is a free data retrieval call binding the contract method 0xc561d4b7.
//
// Solidity: function getCustodian() constant returns(address)
func (_PermUpgr *PermUpgrCaller) GetCustodian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PermUpgr.contract.Call(opts, out, "getCustodian")
	return *ret0, err
}

// GetCustodian is a free data retrieval call binding the contract method 0xc561d4b7.
//
// Solidity: function getCustodian() constant returns(address)
func (_PermUpgr *PermUpgrSession) GetCustodian() (common.Address, error) {
	return _PermUpgr.Contract.GetCustodian(&_PermUpgr.CallOpts)
}

// GetCustodian is a free data retrieval call binding the contract method 0xc561d4b7.
//
// Solidity: function getCustodian() constant returns(address)
func (_PermUpgr *PermUpgrCallerSession) GetCustodian() (common.Address, error) {
	return _PermUpgr.Contract.GetCustodian(&_PermUpgr.CallOpts)
}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() constant returns(address)
func (_PermUpgr *PermUpgrCaller) GetPermImpl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PermUpgr.contract.Call(opts, out, "getPermImpl")
	return *ret0, err
}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() constant returns(address)
func (_PermUpgr *PermUpgrSession) GetPermImpl() (common.Address, error) {
	return _PermUpgr.Contract.GetPermImpl(&_PermUpgr.CallOpts)
}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() constant returns(address)
func (_PermUpgr *PermUpgrCallerSession) GetPermImpl() (common.Address, error) {
	return _PermUpgr.Contract.GetPermImpl(&_PermUpgr.CallOpts)
}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() constant returns(address)
func (_PermUpgr *PermUpgrCaller) GetPermInterface(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PermUpgr.contract.Call(opts, out, "getPermInterface")
	return *ret0, err
}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() constant returns(address)
func (_PermUpgr *PermUpgrSession) GetPermInterface() (common.Address, error) {
	return _PermUpgr.Contract.GetPermInterface(&_PermUpgr.CallOpts)
}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() constant returns(address)
func (_PermUpgr *PermUpgrCallerSession) GetPermInterface() (common.Address, error) {
	return _PermUpgr.Contract.GetPermInterface(&_PermUpgr.CallOpts)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x22bcb39a.
//
// Solidity: function confirmImplChange(_proposedImpl address) returns()
func (_PermUpgr *PermUpgrTransactor) ConfirmImplChange(opts *bind.TransactOpts, _proposedImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.contract.Transact(opts, "confirmImplChange", _proposedImpl)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x22bcb39a.
//
// Solidity: function confirmImplChange(_proposedImpl address) returns()
func (_PermUpgr *PermUpgrSession) ConfirmImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.Contract.ConfirmImplChange(&_PermUpgr.TransactOpts, _proposedImpl)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x22bcb39a.
//
// Solidity: function confirmImplChange(_proposedImpl address) returns()
func (_PermUpgr *PermUpgrTransactorSession) ConfirmImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.Contract.ConfirmImplChange(&_PermUpgr.TransactOpts, _proposedImpl)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(_permInterface address, _permImpl address) returns()
func (_PermUpgr *PermUpgrTransactor) Init(opts *bind.TransactOpts, _permInterface common.Address, _permImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.contract.Transact(opts, "init", _permInterface, _permImpl)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(_permInterface address, _permImpl address) returns()
func (_PermUpgr *PermUpgrSession) Init(_permInterface common.Address, _permImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.Contract.Init(&_PermUpgr.TransactOpts, _permInterface, _permImpl)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(_permInterface address, _permImpl address) returns()
func (_PermUpgr *PermUpgrTransactorSession) Init(_permInterface common.Address, _permImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.Contract.Init(&_PermUpgr.TransactOpts, _permInterface, _permImpl)
}
