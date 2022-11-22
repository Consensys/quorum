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

// LCManagementABI is the input ABI used to generate the binding from.
const LCManagementABI = "[{\"inputs\":[{\"internalType\":\"contractIPermissionsInterface\",\"name\":\"_permission\",\"type\":\"address\"},{\"internalType\":\"contractIMode\",\"name\":\"_mode\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amendRequest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"isAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mode\",\"outputs\":[{\"internalType\":\"contractIMode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permission\",\"outputs\":[{\"internalType\":\"contractIPermissionsInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_amendRequest\",\"type\":\"address\"}],\"name\":\"setAmendRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mode\",\"type\":\"address\"}],\"name\":\"setMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permission\",\"type\":\"address\"}],\"name\":\"setPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_standardFactory\",\"type\":\"address\"}],\"name\":\"setStandardFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_upasFactory\",\"type\":\"address\"}],\"name\":\"setUPASFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"standardFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_orgs\",\"type\":\"bytes32[]\"}],\"name\":\"unwhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upasFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_org\",\"type\":\"bytes32\"}],\"name\":\"verifyIdentity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_orgs\",\"type\":\"bytes32[]\"}],\"name\":\"whitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"whitelistOrgs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var LCManagementParsedABI, _ = abi.JSON(strings.NewReader(LCManagementABI))

// LCManagement is an auto generated Go binding around an Ethereum contract.
type LCManagement struct {
	LCManagementCaller     // Read-only binding to the contract
	LCManagementTransactor // Write-only binding to the contract
	LCManagementFilterer   // Log filterer for contract events
}

// LCManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type LCManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LCManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LCManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LCManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LCManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LCManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LCManagementSession struct {
	Contract     *LCManagement     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LCManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LCManagementCallerSession struct {
	Contract *LCManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LCManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LCManagementTransactorSession struct {
	Contract     *LCManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LCManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type LCManagementRaw struct {
	Contract *LCManagement // Generic contract binding to access the raw methods on
}

// LCManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LCManagementCallerRaw struct {
	Contract *LCManagementCaller // Generic read-only contract binding to access the raw methods on
}

// LCManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LCManagementTransactorRaw struct {
	Contract *LCManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLCManagement creates a new instance of LCManagement, bound to a specific deployed contract.
func NewLCManagement(address common.Address, backend bind.ContractBackend) (*LCManagement, error) {
	contract, err := bindLCManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LCManagement{LCManagementCaller: LCManagementCaller{contract: contract}, LCManagementTransactor: LCManagementTransactor{contract: contract}, LCManagementFilterer: LCManagementFilterer{contract: contract}}, nil
}

// NewLCManagementCaller creates a new read-only instance of LCManagement, bound to a specific deployed contract.
func NewLCManagementCaller(address common.Address, caller bind.ContractCaller) (*LCManagementCaller, error) {
	contract, err := bindLCManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LCManagementCaller{contract: contract}, nil
}

// NewLCManagementTransactor creates a new write-only instance of LCManagement, bound to a specific deployed contract.
func NewLCManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*LCManagementTransactor, error) {
	contract, err := bindLCManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LCManagementTransactor{contract: contract}, nil
}

// NewLCManagementFilterer creates a new log filterer instance of LCManagement, bound to a specific deployed contract.
func NewLCManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*LCManagementFilterer, error) {
	contract, err := bindLCManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LCManagementFilterer{contract: contract}, nil
}

// bindLCManagement binds a generic wrapper to an already deployed contract.
func bindLCManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LCManagementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LCManagement *LCManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LCManagement.Contract.LCManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LCManagement *LCManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LCManagement.Contract.LCManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LCManagement *LCManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LCManagement.Contract.LCManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LCManagement *LCManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LCManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LCManagement *LCManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LCManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LCManagement *LCManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LCManagement.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LCManagement *LCManagementCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LCManagement *LCManagementSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LCManagement.Contract.DEFAULTADMINROLE(&_LCManagement.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LCManagement *LCManagementCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LCManagement.Contract.DEFAULTADMINROLE(&_LCManagement.CallOpts)
}

// AmendRequest is a free data retrieval call binding the contract method 0x66fba795.
//
// Solidity: function amendRequest() view returns(address)
func (_LCManagement *LCManagementCaller) AmendRequest(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "amendRequest")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmendRequest is a free data retrieval call binding the contract method 0x66fba795.
//
// Solidity: function amendRequest() view returns(address)
func (_LCManagement *LCManagementSession) AmendRequest() (common.Address, error) {
	return _LCManagement.Contract.AmendRequest(&_LCManagement.CallOpts)
}

// AmendRequest is a free data retrieval call binding the contract method 0x66fba795.
//
// Solidity: function amendRequest() view returns(address)
func (_LCManagement *LCManagementCallerSession) AmendRequest() (common.Address, error) {
	return _LCManagement.Contract.AmendRequest(&_LCManagement.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LCManagement *LCManagementCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LCManagement *LCManagementSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LCManagement.Contract.GetRoleAdmin(&_LCManagement.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LCManagement *LCManagementCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LCManagement.Contract.GetRoleAdmin(&_LCManagement.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_LCManagement *LCManagementCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_LCManagement *LCManagementSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _LCManagement.Contract.GetRoleMember(&_LCManagement.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_LCManagement *LCManagementCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _LCManagement.Contract.GetRoleMember(&_LCManagement.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_LCManagement *LCManagementCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_LCManagement *LCManagementSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _LCManagement.Contract.GetRoleMemberCount(&_LCManagement.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_LCManagement *LCManagementCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _LCManagement.Contract.GetRoleMemberCount(&_LCManagement.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LCManagement *LCManagementCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LCManagement *LCManagementSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LCManagement.Contract.HasRole(&_LCManagement.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LCManagement *LCManagementCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LCManagement.Contract.HasRole(&_LCManagement.CallOpts, role, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _account) view returns(bool)
func (_LCManagement *LCManagementCaller) IsAdmin(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "isAdmin", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _account) view returns(bool)
func (_LCManagement *LCManagementSession) IsAdmin(_account common.Address) (bool, error) {
	return _LCManagement.Contract.IsAdmin(&_LCManagement.CallOpts, _account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _account) view returns(bool)
func (_LCManagement *LCManagementCallerSession) IsAdmin(_account common.Address) (bool, error) {
	return _LCManagement.Contract.IsAdmin(&_LCManagement.CallOpts, _account)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _caller) view returns(bool)
func (_LCManagement *LCManagementCaller) IsAuthorized(opts *bind.CallOpts, _caller common.Address) (bool, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "isAuthorized", _caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _caller) view returns(bool)
func (_LCManagement *LCManagementSession) IsAuthorized(_caller common.Address) (bool, error) {
	return _LCManagement.Contract.IsAuthorized(&_LCManagement.CallOpts, _caller)
}

// IsAuthorized is a free data retrieval call binding the contract method 0xfe9fbb80.
//
// Solidity: function isAuthorized(address _caller) view returns(bool)
func (_LCManagement *LCManagementCallerSession) IsAuthorized(_caller common.Address) (bool, error) {
	return _LCManagement.Contract.IsAuthorized(&_LCManagement.CallOpts, _caller)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _account) view returns(bool)
func (_LCManagement *LCManagementCaller) IsOperator(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "isOperator", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _account) view returns(bool)
func (_LCManagement *LCManagementSession) IsOperator(_account common.Address) (bool, error) {
	return _LCManagement.Contract.IsOperator(&_LCManagement.CallOpts, _account)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _account) view returns(bool)
func (_LCManagement *LCManagementCallerSession) IsOperator(_account common.Address) (bool, error) {
	return _LCManagement.Contract.IsOperator(&_LCManagement.CallOpts, _account)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address _account) view returns(bool)
func (_LCManagement *LCManagementCaller) IsVerifier(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "isVerifier", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address _account) view returns(bool)
func (_LCManagement *LCManagementSession) IsVerifier(_account common.Address) (bool, error) {
	return _LCManagement.Contract.IsVerifier(&_LCManagement.CallOpts, _account)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address _account) view returns(bool)
func (_LCManagement *LCManagementCallerSession) IsVerifier(_account common.Address) (bool, error) {
	return _LCManagement.Contract.IsVerifier(&_LCManagement.CallOpts, _account)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(address)
func (_LCManagement *LCManagementCaller) Mode(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "mode")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(address)
func (_LCManagement *LCManagementSession) Mode() (common.Address, error) {
	return _LCManagement.Contract.Mode(&_LCManagement.CallOpts)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(address)
func (_LCManagement *LCManagementCallerSession) Mode() (common.Address, error) {
	return _LCManagement.Contract.Mode(&_LCManagement.CallOpts)
}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_LCManagement *LCManagementCaller) Permission(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "permission")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_LCManagement *LCManagementSession) Permission() (common.Address, error) {
	return _LCManagement.Contract.Permission(&_LCManagement.CallOpts)
}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_LCManagement *LCManagementCallerSession) Permission() (common.Address, error) {
	return _LCManagement.Contract.Permission(&_LCManagement.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_LCManagement *LCManagementCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_LCManagement *LCManagementSession) Router() (common.Address, error) {
	return _LCManagement.Contract.Router(&_LCManagement.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_LCManagement *LCManagementCallerSession) Router() (common.Address, error) {
	return _LCManagement.Contract.Router(&_LCManagement.CallOpts)
}

// StandardFactory is a free data retrieval call binding the contract method 0x317f8638.
//
// Solidity: function standardFactory() view returns(address)
func (_LCManagement *LCManagementCaller) StandardFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "standardFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StandardFactory is a free data retrieval call binding the contract method 0x317f8638.
//
// Solidity: function standardFactory() view returns(address)
func (_LCManagement *LCManagementSession) StandardFactory() (common.Address, error) {
	return _LCManagement.Contract.StandardFactory(&_LCManagement.CallOpts)
}

// StandardFactory is a free data retrieval call binding the contract method 0x317f8638.
//
// Solidity: function standardFactory() view returns(address)
func (_LCManagement *LCManagementCallerSession) StandardFactory() (common.Address, error) {
	return _LCManagement.Contract.StandardFactory(&_LCManagement.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LCManagement *LCManagementCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LCManagement *LCManagementSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LCManagement.Contract.SupportsInterface(&_LCManagement.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LCManagement *LCManagementCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LCManagement.Contract.SupportsInterface(&_LCManagement.CallOpts, interfaceId)
}

// UpasFactory is a free data retrieval call binding the contract method 0x5f8501c8.
//
// Solidity: function upasFactory() view returns(address)
func (_LCManagement *LCManagementCaller) UpasFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "upasFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpasFactory is a free data retrieval call binding the contract method 0x5f8501c8.
//
// Solidity: function upasFactory() view returns(address)
func (_LCManagement *LCManagementSession) UpasFactory() (common.Address, error) {
	return _LCManagement.Contract.UpasFactory(&_LCManagement.CallOpts)
}

// UpasFactory is a free data retrieval call binding the contract method 0x5f8501c8.
//
// Solidity: function upasFactory() view returns(address)
func (_LCManagement *LCManagementCallerSession) UpasFactory() (common.Address, error) {
	return _LCManagement.Contract.UpasFactory(&_LCManagement.CallOpts)
}

// VerifyIdentity is a free data retrieval call binding the contract method 0x5581f372.
//
// Solidity: function verifyIdentity(address _account, bytes32 _org) view returns(bool)
func (_LCManagement *LCManagementCaller) VerifyIdentity(opts *bind.CallOpts, _account common.Address, _org [32]byte) (bool, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "verifyIdentity", _account, _org)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyIdentity is a free data retrieval call binding the contract method 0x5581f372.
//
// Solidity: function verifyIdentity(address _account, bytes32 _org) view returns(bool)
func (_LCManagement *LCManagementSession) VerifyIdentity(_account common.Address, _org [32]byte) (bool, error) {
	return _LCManagement.Contract.VerifyIdentity(&_LCManagement.CallOpts, _account, _org)
}

// VerifyIdentity is a free data retrieval call binding the contract method 0x5581f372.
//
// Solidity: function verifyIdentity(address _account, bytes32 _org) view returns(bool)
func (_LCManagement *LCManagementCallerSession) VerifyIdentity(_account common.Address, _org [32]byte) (bool, error) {
	return _LCManagement.Contract.VerifyIdentity(&_LCManagement.CallOpts, _account, _org)
}

// WhitelistOrgs is a free data retrieval call binding the contract method 0xf1c28ca5.
//
// Solidity: function whitelistOrgs(bytes32 ) view returns(bool)
func (_LCManagement *LCManagementCaller) WhitelistOrgs(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _LCManagement.contract.Call(opts, &out, "whitelistOrgs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistOrgs is a free data retrieval call binding the contract method 0xf1c28ca5.
//
// Solidity: function whitelistOrgs(bytes32 ) view returns(bool)
func (_LCManagement *LCManagementSession) WhitelistOrgs(arg0 [32]byte) (bool, error) {
	return _LCManagement.Contract.WhitelistOrgs(&_LCManagement.CallOpts, arg0)
}

// WhitelistOrgs is a free data retrieval call binding the contract method 0xf1c28ca5.
//
// Solidity: function whitelistOrgs(bytes32 ) view returns(bool)
func (_LCManagement *LCManagementCallerSession) WhitelistOrgs(arg0 [32]byte) (bool, error) {
	return _LCManagement.Contract.WhitelistOrgs(&_LCManagement.CallOpts, arg0)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LCManagement *LCManagementTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LCManagement.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LCManagement *LCManagementSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.GrantRole(&_LCManagement.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LCManagement *LCManagementTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.GrantRole(&_LCManagement.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LCManagement *LCManagementTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LCManagement.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LCManagement *LCManagementSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.RenounceRole(&_LCManagement.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LCManagement *LCManagementTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.RenounceRole(&_LCManagement.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LCManagement *LCManagementTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LCManagement.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LCManagement *LCManagementSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.RevokeRole(&_LCManagement.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LCManagement *LCManagementTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.RevokeRole(&_LCManagement.TransactOpts, role, account)
}

// SetAmendRequest is a paid mutator transaction binding the contract method 0x71e771a2.
//
// Solidity: function setAmendRequest(address _amendRequest) returns()
func (_LCManagement *LCManagementTransactor) SetAmendRequest(opts *bind.TransactOpts, _amendRequest common.Address) (*types.Transaction, error) {
	return _LCManagement.contract.Transact(opts, "setAmendRequest", _amendRequest)
}

// SetAmendRequest is a paid mutator transaction binding the contract method 0x71e771a2.
//
// Solidity: function setAmendRequest(address _amendRequest) returns()
func (_LCManagement *LCManagementSession) SetAmendRequest(_amendRequest common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetAmendRequest(&_LCManagement.TransactOpts, _amendRequest)
}

// SetAmendRequest is a paid mutator transaction binding the contract method 0x71e771a2.
//
// Solidity: function setAmendRequest(address _amendRequest) returns()
func (_LCManagement *LCManagementTransactorSession) SetAmendRequest(_amendRequest common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetAmendRequest(&_LCManagement.TransactOpts, _amendRequest)
}

// SetMode is a paid mutator transaction binding the contract method 0x9e694cea.
//
// Solidity: function setMode(address _mode) returns()
func (_LCManagement *LCManagementTransactor) SetMode(opts *bind.TransactOpts, _mode common.Address) (*types.Transaction, error) {
	return _LCManagement.contract.Transact(opts, "setMode", _mode)
}

// SetMode is a paid mutator transaction binding the contract method 0x9e694cea.
//
// Solidity: function setMode(address _mode) returns()
func (_LCManagement *LCManagementSession) SetMode(_mode common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetMode(&_LCManagement.TransactOpts, _mode)
}

// SetMode is a paid mutator transaction binding the contract method 0x9e694cea.
//
// Solidity: function setMode(address _mode) returns()
func (_LCManagement *LCManagementTransactorSession) SetMode(_mode common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetMode(&_LCManagement.TransactOpts, _mode)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_LCManagement *LCManagementTransactor) SetPermission(opts *bind.TransactOpts, _permission common.Address) (*types.Transaction, error) {
	return _LCManagement.contract.Transact(opts, "setPermission", _permission)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_LCManagement *LCManagementSession) SetPermission(_permission common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetPermission(&_LCManagement.TransactOpts, _permission)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_LCManagement *LCManagementTransactorSession) SetPermission(_permission common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetPermission(&_LCManagement.TransactOpts, _permission)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_LCManagement *LCManagementTransactor) SetRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _LCManagement.contract.Transact(opts, "setRouter", _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_LCManagement *LCManagementSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetRouter(&_LCManagement.TransactOpts, _router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _router) returns()
func (_LCManagement *LCManagementTransactorSession) SetRouter(_router common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetRouter(&_LCManagement.TransactOpts, _router)
}

// SetStandardFactory is a paid mutator transaction binding the contract method 0x005fa939.
//
// Solidity: function setStandardFactory(address _standardFactory) returns()
func (_LCManagement *LCManagementTransactor) SetStandardFactory(opts *bind.TransactOpts, _standardFactory common.Address) (*types.Transaction, error) {
	return _LCManagement.contract.Transact(opts, "setStandardFactory", _standardFactory)
}

// SetStandardFactory is a paid mutator transaction binding the contract method 0x005fa939.
//
// Solidity: function setStandardFactory(address _standardFactory) returns()
func (_LCManagement *LCManagementSession) SetStandardFactory(_standardFactory common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetStandardFactory(&_LCManagement.TransactOpts, _standardFactory)
}

// SetStandardFactory is a paid mutator transaction binding the contract method 0x005fa939.
//
// Solidity: function setStandardFactory(address _standardFactory) returns()
func (_LCManagement *LCManagementTransactorSession) SetStandardFactory(_standardFactory common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetStandardFactory(&_LCManagement.TransactOpts, _standardFactory)
}

// SetUPASFactory is a paid mutator transaction binding the contract method 0x39cd8e96.
//
// Solidity: function setUPASFactory(address _upasFactory) returns()
func (_LCManagement *LCManagementTransactor) SetUPASFactory(opts *bind.TransactOpts, _upasFactory common.Address) (*types.Transaction, error) {
	return _LCManagement.contract.Transact(opts, "setUPASFactory", _upasFactory)
}

// SetUPASFactory is a paid mutator transaction binding the contract method 0x39cd8e96.
//
// Solidity: function setUPASFactory(address _upasFactory) returns()
func (_LCManagement *LCManagementSession) SetUPASFactory(_upasFactory common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetUPASFactory(&_LCManagement.TransactOpts, _upasFactory)
}

// SetUPASFactory is a paid mutator transaction binding the contract method 0x39cd8e96.
//
// Solidity: function setUPASFactory(address _upasFactory) returns()
func (_LCManagement *LCManagementTransactorSession) SetUPASFactory(_upasFactory common.Address) (*types.Transaction, error) {
	return _LCManagement.Contract.SetUPASFactory(&_LCManagement.TransactOpts, _upasFactory)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x3908d9a0.
//
// Solidity: function unwhitelist(bytes32[] _orgs) returns()
func (_LCManagement *LCManagementTransactor) Unwhitelist(opts *bind.TransactOpts, _orgs [][32]byte) (*types.Transaction, error) {
	return _LCManagement.contract.Transact(opts, "unwhitelist", _orgs)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x3908d9a0.
//
// Solidity: function unwhitelist(bytes32[] _orgs) returns()
func (_LCManagement *LCManagementSession) Unwhitelist(_orgs [][32]byte) (*types.Transaction, error) {
	return _LCManagement.Contract.Unwhitelist(&_LCManagement.TransactOpts, _orgs)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x3908d9a0.
//
// Solidity: function unwhitelist(bytes32[] _orgs) returns()
func (_LCManagement *LCManagementTransactorSession) Unwhitelist(_orgs [][32]byte) (*types.Transaction, error) {
	return _LCManagement.Contract.Unwhitelist(&_LCManagement.TransactOpts, _orgs)
}

// Whitelist is a paid mutator transaction binding the contract method 0x3b9f8383.
//
// Solidity: function whitelist(bytes32[] _orgs) returns()
func (_LCManagement *LCManagementTransactor) Whitelist(opts *bind.TransactOpts, _orgs [][32]byte) (*types.Transaction, error) {
	return _LCManagement.contract.Transact(opts, "whitelist", _orgs)
}

// Whitelist is a paid mutator transaction binding the contract method 0x3b9f8383.
//
// Solidity: function whitelist(bytes32[] _orgs) returns()
func (_LCManagement *LCManagementSession) Whitelist(_orgs [][32]byte) (*types.Transaction, error) {
	return _LCManagement.Contract.Whitelist(&_LCManagement.TransactOpts, _orgs)
}

// Whitelist is a paid mutator transaction binding the contract method 0x3b9f8383.
//
// Solidity: function whitelist(bytes32[] _orgs) returns()
func (_LCManagement *LCManagementTransactorSession) Whitelist(_orgs [][32]byte) (*types.Transaction, error) {
	return _LCManagement.Contract.Whitelist(&_LCManagement.TransactOpts, _orgs)
}

// LCManagementRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LCManagement contract.
type LCManagementRoleAdminChangedIterator struct {
	Event *LCManagementRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LCManagementRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LCManagementRoleAdminChanged)
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
		it.Event = new(LCManagementRoleAdminChanged)
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
func (it *LCManagementRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LCManagementRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LCManagementRoleAdminChanged represents a RoleAdminChanged event raised by the LCManagement contract.
type LCManagementRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LCManagement *LCManagementFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LCManagementRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LCManagement.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LCManagementRoleAdminChangedIterator{contract: _LCManagement.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// var RoleAdminChangedTopicHash = "0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff"

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LCManagement *LCManagementFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LCManagementRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LCManagement.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LCManagementRoleAdminChanged)
				if err := _LCManagement.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LCManagement *LCManagementFilterer) ParseRoleAdminChanged(log types.Log) (*LCManagementRoleAdminChanged, error) {
	event := new(LCManagementRoleAdminChanged)
	if err := _LCManagement.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LCManagementRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LCManagement contract.
type LCManagementRoleGrantedIterator struct {
	Event *LCManagementRoleGranted // Event containing the contract specifics and raw log

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
func (it *LCManagementRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LCManagementRoleGranted)
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
		it.Event = new(LCManagementRoleGranted)
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
func (it *LCManagementRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LCManagementRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LCManagementRoleGranted represents a RoleGranted event raised by the LCManagement contract.
type LCManagementRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LCManagement *LCManagementFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LCManagementRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LCManagement.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LCManagementRoleGrantedIterator{contract: _LCManagement.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

var RoleGrantedTopicHash = "0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d"

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LCManagement *LCManagementFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LCManagementRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LCManagement.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LCManagementRoleGranted)
				if err := _LCManagement.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LCManagement *LCManagementFilterer) ParseRoleGranted(log types.Log) (*LCManagementRoleGranted, error) {
	event := new(LCManagementRoleGranted)
	if err := _LCManagement.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LCManagementRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LCManagement contract.
type LCManagementRoleRevokedIterator struct {
	Event *LCManagementRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LCManagementRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LCManagementRoleRevoked)
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
		it.Event = new(LCManagementRoleRevoked)
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
func (it *LCManagementRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LCManagementRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LCManagementRoleRevoked represents a RoleRevoked event raised by the LCManagement contract.
type LCManagementRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LCManagement *LCManagementFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LCManagementRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LCManagement.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LCManagementRoleRevokedIterator{contract: _LCManagement.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

var RoleRevokedTopicHash = "0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b"

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LCManagement *LCManagementFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LCManagementRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LCManagement.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LCManagementRoleRevoked)
				if err := _LCManagement.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LCManagement *LCManagementFilterer) ParseRoleRevoked(log types.Log) (*LCManagementRoleRevoked, error) {
	event := new(LCManagementRoleRevoked)
	if err := _LCManagement.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
