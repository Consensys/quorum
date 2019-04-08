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

// PermInterfaceABI is the input ABI used to generate the binding from.
const PermInterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getPermissionsImpl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"updateNodeStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nwAdminOrg\",\"type\":\"string\"},{\"name\":\"_nwAdminRole\",\"type\":\"string\"},{\"name\":\"_oAdminRole\",\"type\":\"string\"}],\"name\":\"setPolicy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_roleId\",\"type\":\"string\"}],\"name\":\"assignAccountRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgManager\",\"type\":\"address\"},{\"name\":\"_rolesManager\",\"type\":\"address\"},{\"name\":\"_acctManager\",\"type\":\"address\"},{\"name\":\"_voterManager\",\"type\":\"address\"},{\"name\":\"_nodeManager\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateNetworkBootStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNetworkBootStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_permImplementation\",\"type\":\"address\"}],\"name\":\"setPermImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"approveOrgStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"validateAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"}],\"name\":\"addAdminAccounts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"approveOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"assignOrgAdminAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_access\",\"type\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"bool\"}],\"name\":\"addNewRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"isOrgAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"updateOrgStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isNetworkAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"approveOrgAdminAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"addAdminNodes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getPendingOp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permImplUpgradeable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PermInterface is an auto generated Go binding around an Ethereum contract.
type PermInterface struct {
	PermInterfaceCaller     // Read-only binding to the contract
	PermInterfaceTransactor // Write-only binding to the contract
	PermInterfaceFilterer   // Log filterer for contract events
}

// PermInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermInterfaceSession struct {
	Contract     *PermInterface    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermInterfaceCallerSession struct {
	Contract *PermInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PermInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermInterfaceTransactorSession struct {
	Contract     *PermInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PermInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermInterfaceRaw struct {
	Contract *PermInterface // Generic contract binding to access the raw methods on
}

// PermInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermInterfaceCallerRaw struct {
	Contract *PermInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// PermInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermInterfaceTransactorRaw struct {
	Contract *PermInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermInterface creates a new instance of PermInterface, bound to a specific deployed contract.
func NewPermInterface(address common.Address, backend bind.ContractBackend) (*PermInterface, error) {
	contract, err := bindPermInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermInterface{PermInterfaceCaller: PermInterfaceCaller{contract: contract}, PermInterfaceTransactor: PermInterfaceTransactor{contract: contract}, PermInterfaceFilterer: PermInterfaceFilterer{contract: contract}}, nil
}

// NewPermInterfaceCaller creates a new read-only instance of PermInterface, bound to a specific deployed contract.
func NewPermInterfaceCaller(address common.Address, caller bind.ContractCaller) (*PermInterfaceCaller, error) {
	contract, err := bindPermInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermInterfaceCaller{contract: contract}, nil
}

// NewPermInterfaceTransactor creates a new write-only instance of PermInterface, bound to a specific deployed contract.
func NewPermInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PermInterfaceTransactor, error) {
	contract, err := bindPermInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermInterfaceTransactor{contract: contract}, nil
}

// NewPermInterfaceFilterer creates a new log filterer instance of PermInterface, bound to a specific deployed contract.
func NewPermInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PermInterfaceFilterer, error) {
	contract, err := bindPermInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermInterfaceFilterer{contract: contract}, nil
}

// bindPermInterface binds a generic wrapper to an already deployed contract.
func bindPermInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermInterface *PermInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermInterface.Contract.PermInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermInterface *PermInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermInterface.Contract.PermInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermInterface *PermInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermInterface.Contract.PermInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermInterface *PermInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PermInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermInterface *PermInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermInterface *PermInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermInterface.Contract.contract.Transact(opts, method, params...)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_PermInterface *PermInterfaceCaller) GetNetworkBootStatus(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermInterface.contract.Call(opts, out, "getNetworkBootStatus")
	return *ret0, err
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_PermInterface *PermInterfaceSession) GetNetworkBootStatus() (bool, error) {
	return _PermInterface.Contract.GetNetworkBootStatus(&_PermInterface.CallOpts)
}

// GetNetworkBootStatus is a free data retrieval call binding the contract method 0x4cbfa82e.
//
// Solidity: function getNetworkBootStatus() constant returns(bool)
func (_PermInterface *PermInterfaceCallerSession) GetNetworkBootStatus() (bool, error) {
	return _PermInterface.Contract.GetNetworkBootStatus(&_PermInterface.CallOpts)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(_orgId string) constant returns(string, string, address, uint256)
func (_PermInterface *PermInterfaceCaller) GetPendingOp(opts *bind.CallOpts, _orgId string) (string, string, common.Address, *big.Int, error) {
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
	err := _PermInterface.contract.Call(opts, out, "getPendingOp", _orgId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(_orgId string) constant returns(string, string, address, uint256)
func (_PermInterface *PermInterfaceSession) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _PermInterface.Contract.GetPendingOp(&_PermInterface.CallOpts, _orgId)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(_orgId string) constant returns(string, string, address, uint256)
func (_PermInterface *PermInterfaceCallerSession) GetPendingOp(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _PermInterface.Contract.GetPendingOp(&_PermInterface.CallOpts, _orgId)
}

// GetPermissionsImpl is a free data retrieval call binding the contract method 0x03ed6933.
//
// Solidity: function getPermissionsImpl() constant returns(address)
func (_PermInterface *PermInterfaceCaller) GetPermissionsImpl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PermInterface.contract.Call(opts, out, "getPermissionsImpl")
	return *ret0, err
}

// GetPermissionsImpl is a free data retrieval call binding the contract method 0x03ed6933.
//
// Solidity: function getPermissionsImpl() constant returns(address)
func (_PermInterface *PermInterfaceSession) GetPermissionsImpl() (common.Address, error) {
	return _PermInterface.Contract.GetPermissionsImpl(&_PermInterface.CallOpts)
}

// GetPermissionsImpl is a free data retrieval call binding the contract method 0x03ed6933.
//
// Solidity: function getPermissionsImpl() constant returns(address)
func (_PermInterface *PermInterfaceCallerSession) GetPermissionsImpl() (common.Address, error) {
	return _PermInterface.Contract.GetPermissionsImpl(&_PermInterface.CallOpts)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(_account address) constant returns(bool)
func (_PermInterface *PermInterfaceCaller) IsNetworkAdmin(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermInterface.contract.Call(opts, out, "isNetworkAdmin", _account)
	return *ret0, err
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(_account address) constant returns(bool)
func (_PermInterface *PermInterfaceSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _PermInterface.Contract.IsNetworkAdmin(&_PermInterface.CallOpts, _account)
}

// IsNetworkAdmin is a free data retrieval call binding the contract method 0xd1aa0c20.
//
// Solidity: function isNetworkAdmin(_account address) constant returns(bool)
func (_PermInterface *PermInterfaceCallerSession) IsNetworkAdmin(_account common.Address) (bool, error) {
	return _PermInterface.Contract.IsNetworkAdmin(&_PermInterface.CallOpts, _account)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(_account address, _orgId string) constant returns(bool)
func (_PermInterface *PermInterfaceCaller) IsOrgAdmin(opts *bind.CallOpts, _account common.Address, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermInterface.contract.Call(opts, out, "isOrgAdmin", _account, _orgId)
	return *ret0, err
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(_account address, _orgId string) constant returns(bool)
func (_PermInterface *PermInterfaceSession) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return _PermInterface.Contract.IsOrgAdmin(&_PermInterface.CallOpts, _account, _orgId)
}

// IsOrgAdmin is a free data retrieval call binding the contract method 0x9bd38101.
//
// Solidity: function isOrgAdmin(_account address, _orgId string) constant returns(bool)
func (_PermInterface *PermInterfaceCallerSession) IsOrgAdmin(_account common.Address, _orgId string) (bool, error) {
	return _PermInterface.Contract.IsOrgAdmin(&_PermInterface.CallOpts, _account, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(_account address, _orgId string) constant returns(bool)
func (_PermInterface *PermInterfaceCaller) ValidateAccount(opts *bind.CallOpts, _account common.Address, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermInterface.contract.Call(opts, out, "validateAccount", _account, _orgId)
	return *ret0, err
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(_account address, _orgId string) constant returns(bool)
func (_PermInterface *PermInterfaceSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _PermInterface.Contract.ValidateAccount(&_PermInterface.CallOpts, _account, _orgId)
}

// ValidateAccount is a free data retrieval call binding the contract method 0x6b568d76.
//
// Solidity: function validateAccount(_account address, _orgId string) constant returns(bool)
func (_PermInterface *PermInterfaceCallerSession) ValidateAccount(_account common.Address, _orgId string) (bool, error) {
	return _PermInterface.Contract.ValidateAccount(&_PermInterface.CallOpts, _account, _orgId)
}

// AddAdminAccounts is a paid mutator transaction binding the contract method 0x71f57931.
//
// Solidity: function addAdminAccounts(_acct address) returns()
func (_PermInterface *PermInterfaceTransactor) AddAdminAccounts(opts *bind.TransactOpts, _acct common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "addAdminAccounts", _acct)
}

// AddAdminAccounts is a paid mutator transaction binding the contract method 0x71f57931.
//
// Solidity: function addAdminAccounts(_acct address) returns()
func (_PermInterface *PermInterfaceSession) AddAdminAccounts(_acct common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.AddAdminAccounts(&_PermInterface.TransactOpts, _acct)
}

// AddAdminAccounts is a paid mutator transaction binding the contract method 0x71f57931.
//
// Solidity: function addAdminAccounts(_acct address) returns()
func (_PermInterface *PermInterfaceTransactorSession) AddAdminAccounts(_acct common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.AddAdminAccounts(&_PermInterface.TransactOpts, _acct)
}

// AddAdminNodes is a paid mutator transaction binding the contract method 0xe5e5b85d.
//
// Solidity: function addAdminNodes(_enodeId string) returns()
func (_PermInterface *PermInterfaceTransactor) AddAdminNodes(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "addAdminNodes", _enodeId)
}

// AddAdminNodes is a paid mutator transaction binding the contract method 0xe5e5b85d.
//
// Solidity: function addAdminNodes(_enodeId string) returns()
func (_PermInterface *PermInterfaceSession) AddAdminNodes(_enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AddAdminNodes(&_PermInterface.TransactOpts, _enodeId)
}

// AddAdminNodes is a paid mutator transaction binding the contract method 0xe5e5b85d.
//
// Solidity: function addAdminNodes(_enodeId string) returns()
func (_PermInterface *PermInterfaceTransactorSession) AddAdminNodes(_enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AddAdminNodes(&_PermInterface.TransactOpts, _enodeId)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x9485d7a6.
//
// Solidity: function addNewRole(_roleId string, _orgId string, _access uint256, _voter bool) returns()
func (_PermInterface *PermInterfaceTransactor) AddNewRole(opts *bind.TransactOpts, _roleId string, _orgId string, _access *big.Int, _voter bool) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "addNewRole", _roleId, _orgId, _access, _voter)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x9485d7a6.
//
// Solidity: function addNewRole(_roleId string, _orgId string, _access uint256, _voter bool) returns()
func (_PermInterface *PermInterfaceSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool) (*types.Transaction, error) {
	return _PermInterface.Contract.AddNewRole(&_PermInterface.TransactOpts, _roleId, _orgId, _access, _voter)
}

// AddNewRole is a paid mutator transaction binding the contract method 0x9485d7a6.
//
// Solidity: function addNewRole(_roleId string, _orgId string, _access uint256, _voter bool) returns()
func (_PermInterface *PermInterfaceTransactorSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool) (*types.Transaction, error) {
	return _PermInterface.Contract.AddNewRole(&_PermInterface.TransactOpts, _roleId, _orgId, _access, _voter)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(_orgId string, _enodeId string) returns()
func (_PermInterface *PermInterfaceTransactor) AddNode(opts *bind.TransactOpts, _orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "addNode", _orgId, _enodeId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(_orgId string, _enodeId string) returns()
func (_PermInterface *PermInterfaceSession) AddNode(_orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AddNode(&_PermInterface.TransactOpts, _orgId, _enodeId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(_orgId string, _enodeId string) returns()
func (_PermInterface *PermInterfaceTransactorSession) AddNode(_orgId string, _enodeId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AddNode(&_PermInterface.TransactOpts, _orgId, _enodeId)
}

// AddOrg is a paid mutator transaction binding the contract method 0x8f362a3e.
//
// Solidity: function addOrg(_orgId string, _enodeId string, _account address) returns()
func (_PermInterface *PermInterfaceTransactor) AddOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "addOrg", _orgId, _enodeId, _account)
}

// AddOrg is a paid mutator transaction binding the contract method 0x8f362a3e.
//
// Solidity: function addOrg(_orgId string, _enodeId string, _account address) returns()
func (_PermInterface *PermInterfaceSession) AddOrg(_orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.AddOrg(&_PermInterface.TransactOpts, _orgId, _enodeId, _account)
}

// AddOrg is a paid mutator transaction binding the contract method 0x8f362a3e.
//
// Solidity: function addOrg(_orgId string, _enodeId string, _account address) returns()
func (_PermInterface *PermInterfaceTransactorSession) AddOrg(_orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.AddOrg(&_PermInterface.TransactOpts, _orgId, _enodeId, _account)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x7e461258.
//
// Solidity: function approveOrg(_orgId string, _enodeId string, _account address) returns()
func (_PermInterface *PermInterfaceTransactor) ApproveOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "approveOrg", _orgId, _enodeId, _account)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x7e461258.
//
// Solidity: function approveOrg(_orgId string, _enodeId string, _account address) returns()
func (_PermInterface *PermInterfaceSession) ApproveOrg(_orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveOrg(&_PermInterface.TransactOpts, _orgId, _enodeId, _account)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x7e461258.
//
// Solidity: function approveOrg(_orgId string, _enodeId string, _account address) returns()
func (_PermInterface *PermInterfaceTransactorSession) ApproveOrg(_orgId string, _enodeId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveOrg(&_PermInterface.TransactOpts, _orgId, _enodeId, _account)
}

// ApproveOrgAdminAccount is a paid mutator transaction binding the contract method 0xd5b6b443.
//
// Solidity: function approveOrgAdminAccount(_account address) returns()
func (_PermInterface *PermInterfaceTransactor) ApproveOrgAdminAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "approveOrgAdminAccount", _account)
}

// ApproveOrgAdminAccount is a paid mutator transaction binding the contract method 0xd5b6b443.
//
// Solidity: function approveOrgAdminAccount(_account address) returns()
func (_PermInterface *PermInterfaceSession) ApproveOrgAdminAccount(_account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveOrgAdminAccount(&_PermInterface.TransactOpts, _account)
}

// ApproveOrgAdminAccount is a paid mutator transaction binding the contract method 0xd5b6b443.
//
// Solidity: function approveOrgAdminAccount(_account address) returns()
func (_PermInterface *PermInterfaceTransactorSession) ApproveOrgAdminAccount(_account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveOrgAdminAccount(&_PermInterface.TransactOpts, _account)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0x5be9672c.
//
// Solidity: function approveOrgStatus(_orgId string, _status uint256) returns()
func (_PermInterface *PermInterfaceTransactor) ApproveOrgStatus(opts *bind.TransactOpts, _orgId string, _status *big.Int) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "approveOrgStatus", _orgId, _status)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0x5be9672c.
//
// Solidity: function approveOrgStatus(_orgId string, _status uint256) returns()
func (_PermInterface *PermInterfaceSession) ApproveOrgStatus(_orgId string, _status *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveOrgStatus(&_PermInterface.TransactOpts, _orgId, _status)
}

// ApproveOrgStatus is a paid mutator transaction binding the contract method 0x5be9672c.
//
// Solidity: function approveOrgStatus(_orgId string, _status uint256) returns()
func (_PermInterface *PermInterfaceTransactorSession) ApproveOrgStatus(_orgId string, _status *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.ApproveOrgStatus(&_PermInterface.TransactOpts, _orgId, _status)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x2f7f0a12.
//
// Solidity: function assignAccountRole(_acct address, _orgId string, _roleId string) returns()
func (_PermInterface *PermInterfaceTransactor) AssignAccountRole(opts *bind.TransactOpts, _acct common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "assignAccountRole", _acct, _orgId, _roleId)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x2f7f0a12.
//
// Solidity: function assignAccountRole(_acct address, _orgId string, _roleId string) returns()
func (_PermInterface *PermInterfaceSession) AssignAccountRole(_acct common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AssignAccountRole(&_PermInterface.TransactOpts, _acct, _orgId, _roleId)
}

// AssignAccountRole is a paid mutator transaction binding the contract method 0x2f7f0a12.
//
// Solidity: function assignAccountRole(_acct address, _orgId string, _roleId string) returns()
func (_PermInterface *PermInterfaceTransactorSession) AssignAccountRole(_acct common.Address, _orgId string, _roleId string) (*types.Transaction, error) {
	return _PermInterface.Contract.AssignAccountRole(&_PermInterface.TransactOpts, _acct, _orgId, _roleId)
}

// AssignOrgAdminAccount is a paid mutator transaction binding the contract method 0x8baa2fc8.
//
// Solidity: function assignOrgAdminAccount(_orgId string, _account address) returns()
func (_PermInterface *PermInterfaceTransactor) AssignOrgAdminAccount(opts *bind.TransactOpts, _orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "assignOrgAdminAccount", _orgId, _account)
}

// AssignOrgAdminAccount is a paid mutator transaction binding the contract method 0x8baa2fc8.
//
// Solidity: function assignOrgAdminAccount(_orgId string, _account address) returns()
func (_PermInterface *PermInterfaceSession) AssignOrgAdminAccount(_orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.AssignOrgAdminAccount(&_PermInterface.TransactOpts, _orgId, _account)
}

// AssignOrgAdminAccount is a paid mutator transaction binding the contract method 0x8baa2fc8.
//
// Solidity: function assignOrgAdminAccount(_orgId string, _account address) returns()
func (_PermInterface *PermInterfaceTransactorSession) AssignOrgAdminAccount(_orgId string, _account common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.AssignOrgAdminAccount(&_PermInterface.TransactOpts, _orgId, _account)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(_orgManager address, _rolesManager address, _acctManager address, _voterManager address, _nodeManager address) returns()
func (_PermInterface *PermInterfaceTransactor) Init(opts *bind.TransactOpts, _orgManager common.Address, _rolesManager common.Address, _acctManager common.Address, _voterManager common.Address, _nodeManager common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "init", _orgManager, _rolesManager, _acctManager, _voterManager, _nodeManager)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(_orgManager address, _rolesManager address, _acctManager address, _voterManager address, _nodeManager address) returns()
func (_PermInterface *PermInterfaceSession) Init(_orgManager common.Address, _rolesManager common.Address, _acctManager common.Address, _voterManager common.Address, _nodeManager common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.Init(&_PermInterface.TransactOpts, _orgManager, _rolesManager, _acctManager, _voterManager, _nodeManager)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(_orgManager address, _rolesManager address, _acctManager address, _voterManager address, _nodeManager address) returns()
func (_PermInterface *PermInterfaceTransactorSession) Init(_orgManager common.Address, _rolesManager common.Address, _acctManager common.Address, _voterManager common.Address, _nodeManager common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.Init(&_PermInterface.TransactOpts, _orgManager, _rolesManager, _acctManager, _voterManager, _nodeManager)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(_roleId string, _orgId string) returns()
func (_PermInterface *PermInterfaceTransactor) RemoveRole(opts *bind.TransactOpts, _roleId string, _orgId string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "removeRole", _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(_roleId string, _orgId string) returns()
func (_PermInterface *PermInterfaceSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _PermInterface.Contract.RemoveRole(&_PermInterface.TransactOpts, _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(_roleId string, _orgId string) returns()
func (_PermInterface *PermInterfaceTransactorSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _PermInterface.Contract.RemoveRole(&_PermInterface.TransactOpts, _roleId, _orgId)
}

// SetPermImplementation is a paid mutator transaction binding the contract method 0x511bbd9f.
//
// Solidity: function setPermImplementation(_permImplementation address) returns()
func (_PermInterface *PermInterfaceTransactor) SetPermImplementation(opts *bind.TransactOpts, _permImplementation common.Address) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "setPermImplementation", _permImplementation)
}

// SetPermImplementation is a paid mutator transaction binding the contract method 0x511bbd9f.
//
// Solidity: function setPermImplementation(_permImplementation address) returns()
func (_PermInterface *PermInterfaceSession) SetPermImplementation(_permImplementation common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.SetPermImplementation(&_PermInterface.TransactOpts, _permImplementation)
}

// SetPermImplementation is a paid mutator transaction binding the contract method 0x511bbd9f.
//
// Solidity: function setPermImplementation(_permImplementation address) returns()
func (_PermInterface *PermInterfaceTransactorSession) SetPermImplementation(_permImplementation common.Address) (*types.Transaction, error) {
	return _PermInterface.Contract.SetPermImplementation(&_PermInterface.TransactOpts, _permImplementation)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) returns()
func (_PermInterface *PermInterfaceTransactor) SetPolicy(opts *bind.TransactOpts, _nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "setPolicy", _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) returns()
func (_PermInterface *PermInterfaceSession) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermInterface.Contract.SetPolicy(&_PermInterface.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x1b610220.
//
// Solidity: function setPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) returns()
func (_PermInterface *PermInterfaceTransactorSession) SetPolicy(_nwAdminOrg string, _nwAdminRole string, _oAdminRole string) (*types.Transaction, error) {
	return _PermInterface.Contract.SetPolicy(&_PermInterface.TransactOpts, _nwAdminOrg, _nwAdminRole, _oAdminRole)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermInterface *PermInterfaceTransactor) UpdateNetworkBootStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "updateNetworkBootStatus")
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermInterface *PermInterfaceSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateNetworkBootStatus(&_PermInterface.TransactOpts)
}

// UpdateNetworkBootStatus is a paid mutator transaction binding the contract method 0x44478e79.
//
// Solidity: function updateNetworkBootStatus() returns(bool)
func (_PermInterface *PermInterfaceTransactorSession) UpdateNetworkBootStatus() (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateNetworkBootStatus(&_PermInterface.TransactOpts)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(_orgId string, _enodeId string, _status uint256) returns()
func (_PermInterface *PermInterfaceTransactor) UpdateNodeStatus(opts *bind.TransactOpts, _orgId string, _enodeId string, _status *big.Int) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "updateNodeStatus", _orgId, _enodeId, _status)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(_orgId string, _enodeId string, _status uint256) returns()
func (_PermInterface *PermInterfaceSession) UpdateNodeStatus(_orgId string, _enodeId string, _status *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateNodeStatus(&_PermInterface.TransactOpts, _orgId, _enodeId, _status)
}

// UpdateNodeStatus is a paid mutator transaction binding the contract method 0x0cc50146.
//
// Solidity: function updateNodeStatus(_orgId string, _enodeId string, _status uint256) returns()
func (_PermInterface *PermInterfaceTransactorSession) UpdateNodeStatus(_orgId string, _enodeId string, _status *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateNodeStatus(&_PermInterface.TransactOpts, _orgId, _enodeId, _status)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0xbb3b6e80.
//
// Solidity: function updateOrgStatus(_orgId string, _status uint256) returns()
func (_PermInterface *PermInterfaceTransactor) UpdateOrgStatus(opts *bind.TransactOpts, _orgId string, _status *big.Int) (*types.Transaction, error) {
	return _PermInterface.contract.Transact(opts, "updateOrgStatus", _orgId, _status)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0xbb3b6e80.
//
// Solidity: function updateOrgStatus(_orgId string, _status uint256) returns()
func (_PermInterface *PermInterfaceSession) UpdateOrgStatus(_orgId string, _status *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateOrgStatus(&_PermInterface.TransactOpts, _orgId, _status)
}

// UpdateOrgStatus is a paid mutator transaction binding the contract method 0xbb3b6e80.
//
// Solidity: function updateOrgStatus(_orgId string, _status uint256) returns()
func (_PermInterface *PermInterfaceTransactorSession) UpdateOrgStatus(_orgId string, _status *big.Int) (*types.Transaction, error) {
	return _PermInterface.Contract.UpdateOrgStatus(&_PermInterface.TransactOpts, _orgId, _status)
}
