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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ModeABI is the input ABI used to generate the binding from.
const ModeABI = "[{\"inputs\":[{\"internalType\":\"contractIPermissionsInterface\",\"name\":\"_permission\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"checkAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permission\",\"outputs\":[{\"internalType\":\"contractIPermissionsInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permission\",\"type\":\"address\"}],\"name\":\"setPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchedToDAO\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var ModeParsedABI, _ = abi.JSON(strings.NewReader(ModeABI))

// Mode is an auto generated Go binding around an Ethereum contract.
type Mode struct {
	ModeCaller     // Read-only binding to the contract
	ModeTransactor // Write-only binding to the contract
	ModeFilterer   // Log filterer for contract events
}

// ModeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModeSession struct {
	Contract     *Mode             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModeCallerSession struct {
	Contract *ModeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ModeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModeTransactorSession struct {
	Contract     *ModeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModeRaw struct {
	Contract *Mode // Generic contract binding to access the raw methods on
}

// ModeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModeCallerRaw struct {
	Contract *ModeCaller // Generic read-only contract binding to access the raw methods on
}

// ModeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModeTransactorRaw struct {
	Contract *ModeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMode creates a new instance of Mode, bound to a specific deployed contract.
func NewMode(address common.Address, backend bind.ContractBackend) (*Mode, error) {
	contract, err := bindMode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mode{ModeCaller: ModeCaller{contract: contract}, ModeTransactor: ModeTransactor{contract: contract}, ModeFilterer: ModeFilterer{contract: contract}}, nil
}

// NewModeCaller creates a new read-only instance of Mode, bound to a specific deployed contract.
func NewModeCaller(address common.Address, caller bind.ContractCaller) (*ModeCaller, error) {
	contract, err := bindMode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModeCaller{contract: contract}, nil
}

// NewModeTransactor creates a new write-only instance of Mode, bound to a specific deployed contract.
func NewModeTransactor(address common.Address, transactor bind.ContractTransactor) (*ModeTransactor, error) {
	contract, err := bindMode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModeTransactor{contract: contract}, nil
}

// NewModeFilterer creates a new log filterer instance of Mode, bound to a specific deployed contract.
func NewModeFilterer(address common.Address, filterer bind.ContractFilterer) (*ModeFilterer, error) {
	contract, err := bindMode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModeFilterer{contract: contract}, nil
}

// bindMode binds a generic wrapper to an already deployed contract.
func bindMode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ModeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mode *ModeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mode.Contract.ModeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mode *ModeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mode.Contract.ModeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mode *ModeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mode.Contract.ModeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mode *ModeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mode *ModeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mode *ModeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mode.Contract.contract.Transact(opts, method, params...)
}

// CheckAuthorization is a free data retrieval call binding the contract method 0x89f4dd47.
//
// Solidity: function checkAuthorization(address _operator) view returns(bool)
func (_Mode *ModeCaller) CheckAuthorization(opts *bind.CallOpts, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _Mode.contract.Call(opts, &out, "checkAuthorization", _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckAuthorization is a free data retrieval call binding the contract method 0x89f4dd47.
//
// Solidity: function checkAuthorization(address _operator) view returns(bool)
func (_Mode *ModeSession) CheckAuthorization(_operator common.Address) (bool, error) {
	return _Mode.Contract.CheckAuthorization(&_Mode.CallOpts, _operator)
}

// CheckAuthorization is a free data retrieval call binding the contract method 0x89f4dd47.
//
// Solidity: function checkAuthorization(address _operator) view returns(bool)
func (_Mode *ModeCallerSession) CheckAuthorization(_operator common.Address) (bool, error) {
	return _Mode.Contract.CheckAuthorization(&_Mode.CallOpts, _operator)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Mode *ModeCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mode.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Mode *ModeSession) Dao() (common.Address, error) {
	return _Mode.Contract.Dao(&_Mode.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Mode *ModeCallerSession) Dao() (common.Address, error) {
	return _Mode.Contract.Dao(&_Mode.CallOpts)
}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_Mode *ModeCaller) Permission(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mode.contract.Call(opts, &out, "permission")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_Mode *ModeSession) Permission() (common.Address, error) {
	return _Mode.Contract.Permission(&_Mode.CallOpts)
}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_Mode *ModeCallerSession) Permission() (common.Address, error) {
	return _Mode.Contract.Permission(&_Mode.CallOpts)
}

// SwitchedToDAO is a free data retrieval call binding the contract method 0x84e82810.
//
// Solidity: function switchedToDAO() view returns(bool)
func (_Mode *ModeCaller) SwitchedToDAO(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mode.contract.Call(opts, &out, "switchedToDAO")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwitchedToDAO is a free data retrieval call binding the contract method 0x84e82810.
//
// Solidity: function switchedToDAO() view returns(bool)
func (_Mode *ModeSession) SwitchedToDAO() (bool, error) {
	return _Mode.Contract.SwitchedToDAO(&_Mode.CallOpts)
}

// SwitchedToDAO is a free data retrieval call binding the contract method 0x84e82810.
//
// Solidity: function switchedToDAO() view returns(bool)
func (_Mode *ModeCallerSession) SwitchedToDAO() (bool, error) {
	return _Mode.Contract.SwitchedToDAO(&_Mode.CallOpts)
}

// SetDAO is a paid mutator transaction binding the contract method 0xe73a914c.
//
// Solidity: function setDAO(address _dao) returns()
func (_Mode *ModeTransactor) SetDAO(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _Mode.contract.Transact(opts, "setDAO", _dao)
}

// SetDAO is a paid mutator transaction binding the contract method 0xe73a914c.
//
// Solidity: function setDAO(address _dao) returns()
func (_Mode *ModeSession) SetDAO(_dao common.Address) (*types.Transaction, error) {
	return _Mode.Contract.SetDAO(&_Mode.TransactOpts, _dao)
}

// SetDAO is a paid mutator transaction binding the contract method 0xe73a914c.
//
// Solidity: function setDAO(address _dao) returns()
func (_Mode *ModeTransactorSession) SetDAO(_dao common.Address) (*types.Transaction, error) {
	return _Mode.Contract.SetDAO(&_Mode.TransactOpts, _dao)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_Mode *ModeTransactor) SetPermission(opts *bind.TransactOpts, _permission common.Address) (*types.Transaction, error) {
	return _Mode.contract.Transact(opts, "setPermission", _permission)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_Mode *ModeSession) SetPermission(_permission common.Address) (*types.Transaction, error) {
	return _Mode.Contract.SetPermission(&_Mode.TransactOpts, _permission)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_Mode *ModeTransactorSession) SetPermission(_permission common.Address) (*types.Transaction, error) {
	return _Mode.Contract.SetPermission(&_Mode.TransactOpts, _permission)
}

// SwitchMode is a paid mutator transaction binding the contract method 0xb5680cb5.
//
// Solidity: function switchMode() returns()
func (_Mode *ModeTransactor) SwitchMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mode.contract.Transact(opts, "switchMode")
}

// SwitchMode is a paid mutator transaction binding the contract method 0xb5680cb5.
//
// Solidity: function switchMode() returns()
func (_Mode *ModeSession) SwitchMode() (*types.Transaction, error) {
	return _Mode.Contract.SwitchMode(&_Mode.TransactOpts)
}

// SwitchMode is a paid mutator transaction binding the contract method 0xb5680cb5.
//
// Solidity: function switchMode() returns()
func (_Mode *ModeTransactorSession) SwitchMode() (*types.Transaction, error) {
	return _Mode.Contract.SwitchMode(&_Mode.TransactOpts)
}
