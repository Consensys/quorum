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

// AMCABI is the input ABI used to generate the binding from.
const AMCABI = "[{\"inputs\":[{\"internalType\":\"contractIPermissionsInterface\",\"name\":\"_permission\",\"type\":\"address\"},{\"internalType\":\"contractIMode\",\"name\":\"_mode\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_org\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"amendRequest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isNetworkAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managementOrg\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mode\",\"outputs\":[{\"internalType\":\"contractIMode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permission\",\"outputs\":[{\"internalType\":\"contractIPermissionsInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_amendRequest\",\"type\":\"address\"}],\"name\":\"setAmendRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newOrg\",\"type\":\"string\"}],\"name\":\"setManagementOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mode\",\"type\":\"address\"}],\"name\":\"setMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permission\",\"type\":\"address\"}],\"name\":\"setPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_standardFactory\",\"type\":\"address\"}],\"name\":\"setStandardFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_upasFactory\",\"type\":\"address\"}],\"name\":\"setUPASFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"standardFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upasFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_org\",\"type\":\"bytes32\"}],\"name\":\"verifyIdentity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var AMCParsedABI, _ = abi.JSON(strings.NewReader(AMCABI))

// AMC is an auto generated Go binding around an Ethereum contract.
type AMC struct {
	AMCCaller     // Read-only binding to the contract
	AMCTransactor // Write-only binding to the contract
	AMCFilterer   // Log filterer for contract events
}

// AMCCaller is an auto generated read-only Go binding around an Ethereum contract.
type AMCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AMCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AMCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AMCSession struct {
	Contract     *AMC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AMCCallerSession struct {
	Contract *AMCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AMCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AMCTransactorSession struct {
	Contract     *AMCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMCRaw is an auto generated low-level Go binding around an Ethereum contract.
type AMCRaw struct {
	Contract *AMC // Generic contract binding to access the raw methods on
}

// AMCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AMCCallerRaw struct {
	Contract *AMCCaller // Generic read-only contract binding to access the raw methods on
}

// AMCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AMCTransactorRaw struct {
	Contract *AMCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAMC creates a new instance of AMC, bound to a specific deployed contract.
func NewAMC(address common.Address, backend bind.ContractBackend) (*AMC, error) {
	contract, err := bindAMC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AMC{AMCCaller: AMCCaller{contract: contract}, AMCTransactor: AMCTransactor{contract: contract}, AMCFilterer: AMCFilterer{contract: contract}}, nil
}

// NewAMCCaller creates a new read-only instance of AMC, bound to a specific deployed contract.
func NewAMCCaller(address common.Address, caller bind.ContractCaller) (*AMCCaller, error) {
	contract, err := bindAMC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AMCCaller{contract: contract}, nil
}

// NewAMCTransactor creates a new write-only instance of AMC, bound to a specific deployed contract.
func NewAMCTransactor(address common.Address, transactor bind.ContractTransactor) (*AMCTransactor, error) {
	contract, err := bindAMC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AMCTransactor{contract: contract}, nil
}

// NewAMCFilterer creates a new log filterer instance of AMC, bound to a specific deployed contract.
func NewAMCFilterer(address common.Address, filterer bind.ContractFilterer) (*AMCFilterer, error) {
	contract, err := bindAMC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AMCFilterer{contract: contract}, nil
}

// bindAMC binds a generic wrapper to an already deployed contract.
func bindAMC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AMCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMC *AMCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AMC.Contract.AMCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMC *AMCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMC.Contract.AMCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMC *AMCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMC.Contract.AMCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMC *AMCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AMC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMC *AMCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMC *AMCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMC.Contract.contract.Transact(opts, method, params...)
}

// AmendRequest is a free data retrieval call binding the contract method 0x66fba795.
//
// Solidity: function amendRequest() view returns(address)
func (_AMC *AMCCaller) AmendRequest(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "amendRequest")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmendRequest is a free data retrieval call binding the contract method 0x66fba795.
//
// Solidity: function amendRequest() view returns(address)
func (_AMC *AMCSession) AmendRequest() (common.Address, error) {
	return _AMC.Contract.AmendRequest(&_AMC.CallOpts)
}

// AmendRequest is a free data retrieval call binding the contract method 0x66fba795.
//
// Solidity: function amendRequest() view returns(address)
func (_AMC *AMCCallerSession) AmendRequest() (common.Address, error) {
	return _AMC.Contract.AmendRequest(&_AMC.CallOpts)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _account) view returns(bool)
func (_AMC *AMCCaller) IsAuthorized(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "isAuthorized", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _account) view returns(bool)
func (_AMC *AMCSession) IsAuthorized(_account common.Address) (bool, error) {
	return _AMC.Contract.IsAuthorized(&_AMC.CallOpts, _account)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _account) view returns(bool)
func (_AMC *AMCCallerSession) IsAuthorized(_account common.Address) (bool, error) {
	return _AMC.Contract.IsAuthorized(&_AMC.CallOpts, _account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_AMC *AMCCaller) IsManager(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "isManager", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_AMC *AMCSession) IsManager(_account common.Address) (bool, error) {
	return _AMC.Contract.IsManager(&_AMC.CallOpts, _account)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address _account) view returns(bool)
func (_AMC *AMCCallerSession) IsManager(_account common.Address) (bool, error) {
	return _AMC.Contract.IsManager(&_AMC.CallOpts, _account)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_AMC *AMCCaller) IsNetworkAdmin(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "isNetworkAdmin", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_AMC *AMCSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _AMC.Contract.IsNetworkAdmin(&_AMC.CallOpts, _account)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(address _account) view returns(bool)
func (_AMC *AMCCallerSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _AMC.Contract.IsNetworkAdmin(&_AMC.CallOpts, _account)
}

// ManagementOrg is a free data retrieval call binding the contract method 0x86db5f8f.
//
// Solidity: function managementOrg() view returns(string)
func (_AMC *AMCCaller) ManagementOrg(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "managementOrg")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ManagementOrg is a free data retrieval call binding the contract method 0x86db5f8f.
//
// Solidity: function managementOrg() view returns(string)
func (_AMC *AMCSession) ManagementOrg() (string, error) {
	return _AMC.Contract.ManagementOrg(&_AMC.CallOpts)
}

// ManagementOrg is a free data retrieval call binding the contract method 0x86db5f8f.
//
// Solidity: function managementOrg() view returns(string)
func (_AMC *AMCCallerSession) ManagementOrg() (string, error) {
	return _AMC.Contract.ManagementOrg(&_AMC.CallOpts)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(address)
func (_AMC *AMCCaller) Mode(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "mode")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(address)
func (_AMC *AMCSession) Mode() (common.Address, error) {
	return _AMC.Contract.Mode(&_AMC.CallOpts)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(address)
func (_AMC *AMCCallerSession) Mode() (common.Address, error) {
	return _AMC.Contract.Mode(&_AMC.CallOpts)
}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_AMC *AMCCaller) Permission(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "permission")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_AMC *AMCSession) Permission() (common.Address, error) {
	return _AMC.Contract.Permission(&_AMC.CallOpts)
}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_AMC *AMCCallerSession) Permission() (common.Address, error) {
	return _AMC.Contract.Permission(&_AMC.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_AMC *AMCCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_AMC *AMCSession) Router() (common.Address, error) {
	return _AMC.Contract.Router(&_AMC.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_AMC *AMCCallerSession) Router() (common.Address, error) {
	return _AMC.Contract.Router(&_AMC.CallOpts)
}

// StandardFactory is a free data retrieval call binding the contract method 0x317f8638.
//
// Solidity: function standardFactory() view returns(address)
func (_AMC *AMCCaller) StandardFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "standardFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StandardFactory is a free data retrieval call binding the contract method 0x317f8638.
//
// Solidity: function standardFactory() view returns(address)
func (_AMC *AMCSession) StandardFactory() (common.Address, error) {
	return _AMC.Contract.StandardFactory(&_AMC.CallOpts)
}

// StandardFactory is a free data retrieval call binding the contract method 0x317f8638.
//
// Solidity: function standardFactory() view returns(address)
func (_AMC *AMCCallerSession) StandardFactory() (common.Address, error) {
	return _AMC.Contract.StandardFactory(&_AMC.CallOpts)
}

// UpasFactory is a free data retrieval call binding the contract method 0x5f8501c8.
//
// Solidity: function upasFactory() view returns(address)
func (_AMC *AMCCaller) UpasFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "upasFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpasFactory is a free data retrieval call binding the contract method 0x5f8501c8.
//
// Solidity: function upasFactory() view returns(address)
func (_AMC *AMCSession) UpasFactory() (common.Address, error) {
	return _AMC.Contract.UpasFactory(&_AMC.CallOpts)
}

// UpasFactory is a free data retrieval call binding the contract method 0x5f8501c8.
//
// Solidity: function upasFactory() view returns(address)
func (_AMC *AMCCallerSession) UpasFactory() (common.Address, error) {
	return _AMC.Contract.UpasFactory(&_AMC.CallOpts)
}

// VerifyIdentity is a free data retrieval call binding the contract method 0x5581f372.
//
// Solidity: function verifyIdentity(address _account, bytes32 _org) view returns(bool)
func (_AMC *AMCCaller) VerifyIdentity(opts *bind.CallOpts, _account common.Address, _org [32]byte) (bool, error) {
	var out []interface{}
	err := _AMC.contract.Call(opts, &out, "verifyIdentity", _account, _org)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyIdentity is a free data retrieval call binding the contract method 0x5581f372.
//
// Solidity: function verifyIdentity(address _account, bytes32 _org) view returns(bool)
func (_AMC *AMCSession) VerifyIdentity(_account common.Address, _org [32]byte) (bool, error) {
	return _AMC.Contract.VerifyIdentity(&_AMC.CallOpts, _account, _org)
}

// VerifyIdentity is a free data retrieval call binding the contract method 0x5581f372.
//
// Solidity: function verifyIdentity(address _account, bytes32 _org) view returns(bool)
func (_AMC *AMCCallerSession) VerifyIdentity(_account common.Address, _org [32]byte) (bool, error) {
	return _AMC.Contract.VerifyIdentity(&_AMC.CallOpts, _account, _org)
}

// SetAmendRequest is a paid mutator transaction binding the contract method 0x71e771a2.
//
// Solidity: function setAmendRequest(address _amendRequest) returns()
func (_AMC *AMCTransactor) SetAmendRequest(opts *bind.TransactOpts, _amendRequest common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setAmendRequest", _amendRequest)
}

// SetAmendRequest is a paid mutator transaction binding the contract method 0x71e771a2.
//
// Solidity: function setAmendRequest(address _amendRequest) returns()
func (_AMC *AMCSession) SetAmendRequest(_amendRequest common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetAmendRequest(&_AMC.TransactOpts, _amendRequest)
}

// SetAmendRequest is a paid mutator transaction binding the contract method 0x71e771a2.
//
// Solidity: function setAmendRequest(address _amendRequest) returns()
func (_AMC *AMCTransactorSession) SetAmendRequest(_amendRequest common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetAmendRequest(&_AMC.TransactOpts, _amendRequest)
}

// SetManagementOrg is a paid mutator transaction binding the contract method 0x82929792.
//
// Solidity: function setManagementOrg(string _newOrg) returns()
func (_AMC *AMCTransactor) SetManagementOrg(opts *bind.TransactOpts, _newOrg string) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setManagementOrg", _newOrg)
}

// SetManagementOrg is a paid mutator transaction binding the contract method 0x82929792.
//
// Solidity: function setManagementOrg(string _newOrg) returns()
func (_AMC *AMCSession) SetManagementOrg(_newOrg string) (*types.Transaction, error) {
	return _AMC.Contract.SetManagementOrg(&_AMC.TransactOpts, _newOrg)
}

// SetManagementOrg is a paid mutator transaction binding the contract method 0x82929792.
//
// Solidity: function setManagementOrg(string _newOrg) returns()
func (_AMC *AMCTransactorSession) SetManagementOrg(_newOrg string) (*types.Transaction, error) {
	return _AMC.Contract.SetManagementOrg(&_AMC.TransactOpts, _newOrg)
}

// SetMode is a paid mutator transaction binding the contract method 0x9e694cea.
//
// Solidity: function setMode(address _mode) returns()
func (_AMC *AMCTransactor) SetMode(opts *bind.TransactOpts, _mode common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setMode", _mode)
}

// SetMode is a paid mutator transaction binding the contract method 0x9e694cea.
//
// Solidity: function setMode(address _mode) returns()
func (_AMC *AMCSession) SetMode(_mode common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetMode(&_AMC.TransactOpts, _mode)
}

// SetMode is a paid mutator transaction binding the contract method 0x9e694cea.
//
// Solidity: function setMode(address _mode) returns()
func (_AMC *AMCTransactorSession) SetMode(_mode common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetMode(&_AMC.TransactOpts, _mode)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_AMC *AMCTransactor) SetPermission(opts *bind.TransactOpts, _permission common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setPermission", _permission)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_AMC *AMCSession) SetPermission(_permission common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetPermission(&_AMC.TransactOpts, _permission)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_AMC *AMCTransactorSession) SetPermission(_permission common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetPermission(&_AMC.TransactOpts, _permission)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_AMC *AMCTransactor) SetRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setRouter", _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_AMC *AMCSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetRouter(&_AMC.TransactOpts, _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_AMC *AMCTransactorSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetRouter(&_AMC.TransactOpts, _router)
}

// SetStandardFactory is a paid mutator transaction binding the contract method 0x005fa939.
//
// Solidity: function setStandardFactory(address _standardFactory) returns()
func (_AMC *AMCTransactor) SetStandardFactory(opts *bind.TransactOpts, _standardFactory common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setStandardFactory", _standardFactory)
}

// SetStandardFactory is a paid mutator transaction binding the contract method 0x005fa939.
//
// Solidity: function setStandardFactory(address _standardFactory) returns()
func (_AMC *AMCSession) SetStandardFactory(_standardFactory common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetStandardFactory(&_AMC.TransactOpts, _standardFactory)
}

// SetStandardFactory is a paid mutator transaction binding the contract method 0x005fa939.
//
// Solidity: function setStandardFactory(address _standardFactory) returns()
func (_AMC *AMCTransactorSession) SetStandardFactory(_standardFactory common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetStandardFactory(&_AMC.TransactOpts, _standardFactory)
}

// SetUPASFactory is a paid mutator transaction binding the contract method 0x39cd8e96.
//
// Solidity: function setUPASFactory(address _upasFactory) returns()
func (_AMC *AMCTransactor) SetUPASFactory(opts *bind.TransactOpts, _upasFactory common.Address) (*types.Transaction, error) {
	return _AMC.contract.Transact(opts, "setUPASFactory", _upasFactory)
}

// SetUPASFactory is a paid mutator transaction binding the contract method 0x39cd8e96.
//
// Solidity: function setUPASFactory(address _upasFactory) returns()
func (_AMC *AMCSession) SetUPASFactory(_upasFactory common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetUPASFactory(&_AMC.TransactOpts, _upasFactory)
}

// SetUPASFactory is a paid mutator transaction binding the contract method 0x39cd8e96.
//
// Solidity: function setUPASFactory(address _upasFactory) returns()
func (_AMC *AMCTransactorSession) SetUPASFactory(_upasFactory common.Address) (*types.Transaction, error) {
	return _AMC.Contract.SetUPASFactory(&_AMC.TransactOpts, _upasFactory)
}
