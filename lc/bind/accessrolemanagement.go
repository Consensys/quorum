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

// AccessRoleManagementABI is the input ABI used to generate the binding from.
const AccessRoleManagementABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accessRoleList\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accessRole\",\"type\":\"string\"}],\"name\":\"addAccessRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accessRole\",\"type\":\"string\"}],\"name\":\"getIndexAccessRole\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accessRole\",\"type\":\"string\"}],\"name\":\"isAccessRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_accessRole\",\"type\":\"string\"}],\"name\":\"removeAccessRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAccessRoles\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var AccessRoleManagementParsedABI, _ = abi.JSON(strings.NewReader(AccessRoleManagementABI))

// AccessRoleManagement is an auto generated Go binding around an Ethereum contract.
type AccessRoleManagement struct {
	AccessRoleManagementCaller     // Read-only binding to the contract
	AccessRoleManagementTransactor // Write-only binding to the contract
	AccessRoleManagementFilterer   // Log filterer for contract events
}

// AccessRoleManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessRoleManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessRoleManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessRoleManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessRoleManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessRoleManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessRoleManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessRoleManagementSession struct {
	Contract     *AccessRoleManagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AccessRoleManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessRoleManagementCallerSession struct {
	Contract *AccessRoleManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AccessRoleManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessRoleManagementTransactorSession struct {
	Contract     *AccessRoleManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AccessRoleManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessRoleManagementRaw struct {
	Contract *AccessRoleManagement // Generic contract binding to access the raw methods on
}

// AccessRoleManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessRoleManagementCallerRaw struct {
	Contract *AccessRoleManagementCaller // Generic read-only contract binding to access the raw methods on
}

// AccessRoleManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessRoleManagementTransactorRaw struct {
	Contract *AccessRoleManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessRoleManagement creates a new instance of AccessRoleManagement, bound to a specific deployed contract.
func NewAccessRoleManagement(address common.Address, backend bind.ContractBackend) (*AccessRoleManagement, error) {
	contract, err := bindAccessRoleManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessRoleManagement{AccessRoleManagementCaller: AccessRoleManagementCaller{contract: contract}, AccessRoleManagementTransactor: AccessRoleManagementTransactor{contract: contract}, AccessRoleManagementFilterer: AccessRoleManagementFilterer{contract: contract}}, nil
}

// NewAccessRoleManagementCaller creates a new read-only instance of AccessRoleManagement, bound to a specific deployed contract.
func NewAccessRoleManagementCaller(address common.Address, caller bind.ContractCaller) (*AccessRoleManagementCaller, error) {
	contract, err := bindAccessRoleManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessRoleManagementCaller{contract: contract}, nil
}

// NewAccessRoleManagementTransactor creates a new write-only instance of AccessRoleManagement, bound to a specific deployed contract.
func NewAccessRoleManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessRoleManagementTransactor, error) {
	contract, err := bindAccessRoleManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessRoleManagementTransactor{contract: contract}, nil
}

// NewAccessRoleManagementFilterer creates a new log filterer instance of AccessRoleManagement, bound to a specific deployed contract.
func NewAccessRoleManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessRoleManagementFilterer, error) {
	contract, err := bindAccessRoleManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessRoleManagementFilterer{contract: contract}, nil
}

// bindAccessRoleManagement binds a generic wrapper to an already deployed contract.
func bindAccessRoleManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessRoleManagementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessRoleManagement *AccessRoleManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessRoleManagement.Contract.AccessRoleManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessRoleManagement *AccessRoleManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.AccessRoleManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessRoleManagement *AccessRoleManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.AccessRoleManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessRoleManagement *AccessRoleManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessRoleManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessRoleManagement *AccessRoleManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessRoleManagement *AccessRoleManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessRoleManagement *AccessRoleManagementCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessRoleManagement.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessRoleManagement *AccessRoleManagementSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessRoleManagement.Contract.DEFAULTADMINROLE(&_AccessRoleManagement.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessRoleManagement *AccessRoleManagementCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessRoleManagement.Contract.DEFAULTADMINROLE(&_AccessRoleManagement.CallOpts)
}

// AccessRoleList is a free data retrieval call binding the contract method 0x41442299.
//
// Solidity: function accessRoleList(uint256 ) view returns(string)
func (_AccessRoleManagement *AccessRoleManagementCaller) AccessRoleList(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _AccessRoleManagement.contract.Call(opts, &out, "accessRoleList", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AccessRoleList is a free data retrieval call binding the contract method 0x41442299.
//
// Solidity: function accessRoleList(uint256 ) view returns(string)
func (_AccessRoleManagement *AccessRoleManagementSession) AccessRoleList(arg0 *big.Int) (string, error) {
	return _AccessRoleManagement.Contract.AccessRoleList(&_AccessRoleManagement.CallOpts, arg0)
}

// AccessRoleList is a free data retrieval call binding the contract method 0x41442299.
//
// Solidity: function accessRoleList(uint256 ) view returns(string)
func (_AccessRoleManagement *AccessRoleManagementCallerSession) AccessRoleList(arg0 *big.Int) (string, error) {
	return _AccessRoleManagement.Contract.AccessRoleList(&_AccessRoleManagement.CallOpts, arg0)
}

// GetIndexAccessRole is a free data retrieval call binding the contract method 0xe1ee9d1c.
//
// Solidity: function getIndexAccessRole(string _accessRole) view returns(uint256)
func (_AccessRoleManagement *AccessRoleManagementCaller) GetIndexAccessRole(opts *bind.CallOpts, _accessRole string) (*big.Int, error) {
	var out []interface{}
	err := _AccessRoleManagement.contract.Call(opts, &out, "getIndexAccessRole", _accessRole)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexAccessRole is a free data retrieval call binding the contract method 0xe1ee9d1c.
//
// Solidity: function getIndexAccessRole(string _accessRole) view returns(uint256)
func (_AccessRoleManagement *AccessRoleManagementSession) GetIndexAccessRole(_accessRole string) (*big.Int, error) {
	return _AccessRoleManagement.Contract.GetIndexAccessRole(&_AccessRoleManagement.CallOpts, _accessRole)
}

// GetIndexAccessRole is a free data retrieval call binding the contract method 0xe1ee9d1c.
//
// Solidity: function getIndexAccessRole(string _accessRole) view returns(uint256)
func (_AccessRoleManagement *AccessRoleManagementCallerSession) GetIndexAccessRole(_accessRole string) (*big.Int, error) {
	return _AccessRoleManagement.Contract.GetIndexAccessRole(&_AccessRoleManagement.CallOpts, _accessRole)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessRoleManagement *AccessRoleManagementCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessRoleManagement.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessRoleManagement *AccessRoleManagementSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessRoleManagement.Contract.GetRoleAdmin(&_AccessRoleManagement.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessRoleManagement *AccessRoleManagementCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessRoleManagement.Contract.GetRoleAdmin(&_AccessRoleManagement.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessRoleManagement *AccessRoleManagementCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessRoleManagement.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessRoleManagement *AccessRoleManagementSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessRoleManagement.Contract.GetRoleMember(&_AccessRoleManagement.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessRoleManagement *AccessRoleManagementCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessRoleManagement.Contract.GetRoleMember(&_AccessRoleManagement.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessRoleManagement *AccessRoleManagementCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AccessRoleManagement.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessRoleManagement *AccessRoleManagementSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessRoleManagement.Contract.GetRoleMemberCount(&_AccessRoleManagement.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessRoleManagement *AccessRoleManagementCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessRoleManagement.Contract.GetRoleMemberCount(&_AccessRoleManagement.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessRoleManagement *AccessRoleManagementCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessRoleManagement.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessRoleManagement *AccessRoleManagementSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessRoleManagement.Contract.HasRole(&_AccessRoleManagement.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessRoleManagement *AccessRoleManagementCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessRoleManagement.Contract.HasRole(&_AccessRoleManagement.CallOpts, role, account)
}

// IsAccessRole is a free data retrieval call binding the contract method 0xcf8af71f.
//
// Solidity: function isAccessRole(string _accessRole) view returns(bool)
func (_AccessRoleManagement *AccessRoleManagementCaller) IsAccessRole(opts *bind.CallOpts, _accessRole string) (bool, error) {
	var out []interface{}
	err := _AccessRoleManagement.contract.Call(opts, &out, "isAccessRole", _accessRole)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccessRole is a free data retrieval call binding the contract method 0xcf8af71f.
//
// Solidity: function isAccessRole(string _accessRole) view returns(bool)
func (_AccessRoleManagement *AccessRoleManagementSession) IsAccessRole(_accessRole string) (bool, error) {
	return _AccessRoleManagement.Contract.IsAccessRole(&_AccessRoleManagement.CallOpts, _accessRole)
}

// IsAccessRole is a free data retrieval call binding the contract method 0xcf8af71f.
//
// Solidity: function isAccessRole(string _accessRole) view returns(bool)
func (_AccessRoleManagement *AccessRoleManagementCallerSession) IsAccessRole(_accessRole string) (bool, error) {
	return _AccessRoleManagement.Contract.IsAccessRole(&_AccessRoleManagement.CallOpts, _accessRole)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessRoleManagement *AccessRoleManagementCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessRoleManagement.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessRoleManagement *AccessRoleManagementSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessRoleManagement.Contract.SupportsInterface(&_AccessRoleManagement.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessRoleManagement *AccessRoleManagementCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessRoleManagement.Contract.SupportsInterface(&_AccessRoleManagement.CallOpts, interfaceId)
}

// TotalAccessRoles is a free data retrieval call binding the contract method 0x62837df1.
//
// Solidity: function totalAccessRoles() view returns(uint256)
func (_AccessRoleManagement *AccessRoleManagementCaller) TotalAccessRoles(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessRoleManagement.contract.Call(opts, &out, "totalAccessRoles")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAccessRoles is a free data retrieval call binding the contract method 0x62837df1.
//
// Solidity: function totalAccessRoles() view returns(uint256)
func (_AccessRoleManagement *AccessRoleManagementSession) TotalAccessRoles() (*big.Int, error) {
	return _AccessRoleManagement.Contract.TotalAccessRoles(&_AccessRoleManagement.CallOpts)
}

// TotalAccessRoles is a free data retrieval call binding the contract method 0x62837df1.
//
// Solidity: function totalAccessRoles() view returns(uint256)
func (_AccessRoleManagement *AccessRoleManagementCallerSession) TotalAccessRoles() (*big.Int, error) {
	return _AccessRoleManagement.Contract.TotalAccessRoles(&_AccessRoleManagement.CallOpts)
}

// AddAccessRole is a paid mutator transaction binding the contract method 0x1e3b2af9.
//
// Solidity: function addAccessRole(string _accessRole) returns()
func (_AccessRoleManagement *AccessRoleManagementTransactor) AddAccessRole(opts *bind.TransactOpts, _accessRole string) (*types.Transaction, error) {
	return _AccessRoleManagement.contract.Transact(opts, "addAccessRole", _accessRole)
}

// AddAccessRole is a paid mutator transaction binding the contract method 0x1e3b2af9.
//
// Solidity: function addAccessRole(string _accessRole) returns()
func (_AccessRoleManagement *AccessRoleManagementSession) AddAccessRole(_accessRole string) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.AddAccessRole(&_AccessRoleManagement.TransactOpts, _accessRole)
}

// AddAccessRole is a paid mutator transaction binding the contract method 0x1e3b2af9.
//
// Solidity: function addAccessRole(string _accessRole) returns()
func (_AccessRoleManagement *AccessRoleManagementTransactorSession) AddAccessRole(_accessRole string) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.AddAccessRole(&_AccessRoleManagement.TransactOpts, _accessRole)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_AccessRoleManagement *AccessRoleManagementTransactor) GrantRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _AccessRoleManagement.contract.Transact(opts, "grantRole", _role, _account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_AccessRoleManagement *AccessRoleManagementSession) GrantRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.GrantRole(&_AccessRoleManagement.TransactOpts, _role, _account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 _role, address _account) returns()
func (_AccessRoleManagement *AccessRoleManagementTransactorSession) GrantRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.GrantRole(&_AccessRoleManagement.TransactOpts, _role, _account)
}

// RemoveAccessRole is a paid mutator transaction binding the contract method 0xd407b00c.
//
// Solidity: function removeAccessRole(string _accessRole) returns()
func (_AccessRoleManagement *AccessRoleManagementTransactor) RemoveAccessRole(opts *bind.TransactOpts, _accessRole string) (*types.Transaction, error) {
	return _AccessRoleManagement.contract.Transact(opts, "removeAccessRole", _accessRole)
}

// RemoveAccessRole is a paid mutator transaction binding the contract method 0xd407b00c.
//
// Solidity: function removeAccessRole(string _accessRole) returns()
func (_AccessRoleManagement *AccessRoleManagementSession) RemoveAccessRole(_accessRole string) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.RemoveAccessRole(&_AccessRoleManagement.TransactOpts, _accessRole)
}

// RemoveAccessRole is a paid mutator transaction binding the contract method 0xd407b00c.
//
// Solidity: function removeAccessRole(string _accessRole) returns()
func (_AccessRoleManagement *AccessRoleManagementTransactorSession) RemoveAccessRole(_accessRole string) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.RemoveAccessRole(&_AccessRoleManagement.TransactOpts, _accessRole)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _account) returns()
func (_AccessRoleManagement *AccessRoleManagementTransactor) RenounceRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _AccessRoleManagement.contract.Transact(opts, "renounceRole", _role, _account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _account) returns()
func (_AccessRoleManagement *AccessRoleManagementSession) RenounceRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.RenounceRole(&_AccessRoleManagement.TransactOpts, _role, _account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 _role, address _account) returns()
func (_AccessRoleManagement *AccessRoleManagementTransactorSession) RenounceRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.RenounceRole(&_AccessRoleManagement.TransactOpts, _role, _account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_AccessRoleManagement *AccessRoleManagementTransactor) RevokeRole(opts *bind.TransactOpts, _role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _AccessRoleManagement.contract.Transact(opts, "revokeRole", _role, _account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_AccessRoleManagement *AccessRoleManagementSession) RevokeRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.RevokeRole(&_AccessRoleManagement.TransactOpts, _role, _account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 _role, address _account) returns()
func (_AccessRoleManagement *AccessRoleManagementTransactorSession) RevokeRole(_role [32]byte, _account common.Address) (*types.Transaction, error) {
	return _AccessRoleManagement.Contract.RevokeRole(&_AccessRoleManagement.TransactOpts, _role, _account)
}

// AccessRoleManagementRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessRoleManagement contract.
type AccessRoleManagementRoleAdminChangedIterator struct {
	Event *AccessRoleManagementRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessRoleManagementRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessRoleManagementRoleAdminChanged)
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
		it.Event = new(AccessRoleManagementRoleAdminChanged)
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
func (it *AccessRoleManagementRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessRoleManagementRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessRoleManagementRoleAdminChanged represents a RoleAdminChanged event raised by the AccessRoleManagement contract.
type AccessRoleManagementRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessRoleManagement *AccessRoleManagementFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessRoleManagementRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AccessRoleManagement.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessRoleManagementRoleAdminChangedIterator{contract: _AccessRoleManagement.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

var RoleAdminChangedTopicHash = "0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff"

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessRoleManagement *AccessRoleManagementFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessRoleManagementRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AccessRoleManagement.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessRoleManagementRoleAdminChanged)
				if err := _AccessRoleManagement.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AccessRoleManagement *AccessRoleManagementFilterer) ParseRoleAdminChanged(log types.Log) (*AccessRoleManagementRoleAdminChanged, error) {
	event := new(AccessRoleManagementRoleAdminChanged)
	if err := _AccessRoleManagement.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessRoleManagementRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessRoleManagement contract.
type AccessRoleManagementRoleGrantedIterator struct {
	Event *AccessRoleManagementRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessRoleManagementRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessRoleManagementRoleGranted)
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
		it.Event = new(AccessRoleManagementRoleGranted)
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
func (it *AccessRoleManagementRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessRoleManagementRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessRoleManagementRoleGranted represents a RoleGranted event raised by the AccessRoleManagement contract.
type AccessRoleManagementRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessRoleManagement *AccessRoleManagementFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessRoleManagementRoleGrantedIterator, error) {

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

	logs, sub, err := _AccessRoleManagement.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessRoleManagementRoleGrantedIterator{contract: _AccessRoleManagement.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

var RoleGrantedTopicHash = "0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d"

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessRoleManagement *AccessRoleManagementFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessRoleManagementRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessRoleManagement.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessRoleManagementRoleGranted)
				if err := _AccessRoleManagement.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AccessRoleManagement *AccessRoleManagementFilterer) ParseRoleGranted(log types.Log) (*AccessRoleManagementRoleGranted, error) {
	event := new(AccessRoleManagementRoleGranted)
	if err := _AccessRoleManagement.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessRoleManagementRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessRoleManagement contract.
type AccessRoleManagementRoleRevokedIterator struct {
	Event *AccessRoleManagementRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessRoleManagementRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessRoleManagementRoleRevoked)
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
		it.Event = new(AccessRoleManagementRoleRevoked)
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
func (it *AccessRoleManagementRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessRoleManagementRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessRoleManagementRoleRevoked represents a RoleRevoked event raised by the AccessRoleManagement contract.
type AccessRoleManagementRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessRoleManagement *AccessRoleManagementFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessRoleManagementRoleRevokedIterator, error) {

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

	logs, sub, err := _AccessRoleManagement.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessRoleManagementRoleRevokedIterator{contract: _AccessRoleManagement.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

var RoleRevokedTopicHash = "0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b"

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessRoleManagement *AccessRoleManagementFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessRoleManagementRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessRoleManagement.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessRoleManagementRoleRevoked)
				if err := _AccessRoleManagement.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AccessRoleManagement *AccessRoleManagementFilterer) ParseRoleRevoked(log types.Log) (*AccessRoleManagementRoleRevoked, error) {
	event := new(AccessRoleManagementRoleRevoked)
	if err := _AccessRoleManagement.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
