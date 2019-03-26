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

// PermImplABI is the input ABI used to generate the binding from.
const PermImplABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_acct\",\"type\":\"address\"}],\"name\":\"checkIfVoterExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getVoteCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getRoleDetails\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nwAdminOrg\",\"type\":\"string\"},{\"name\":\"_nwAdminRole\",\"type\":\"string\"},{\"name\":\"_oAdminRole\",\"type\":\"string\"}],\"name\":\"setPolicy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"}],\"name\":\"getAccountDetails\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgManager\",\"type\":\"address\"},{\"name\":\"_rolesManager\",\"type\":\"address\"},{\"name\":\"_acctManager\",\"type\":\"address\"},{\"name\":\"_voterManager\",\"type\":\"address\"},{\"name\":\"_nodeManager\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"getNodeStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateNetworkBootStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNetworkBootStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgIndex\",\"type\":\"uint256\"}],\"name\":\"getOrgInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"validateAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"}],\"name\":\"addAdminAccounts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_acct\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"assignAccountRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"approveOrgAdminAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeIndex\",\"type\":\"uint256\"}],\"name\":\"getNodeDetailsFromIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"assignOrgAdminAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getNumberOfVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"isOrgAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_access\",\"type\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"bool\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"addNewRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isNetworkAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"addAdminNodes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getPendingOp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_msg\",\"type\":\"string\"}],\"name\":\"Dummy\",\"type\":\"event\"}]"

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

// CheckIfVoterExists is a free data retrieval call binding the contract method 0x00b813df.
//
// Solidity: function checkIfVoterExists(_orgId string, _acct address) constant returns(bool)
func (_PermImpl *PermImplCaller) CheckIfVoterExists(opts *bind.CallOpts, _orgId string, _acct common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PermImpl.contract.Call(opts, out, "checkIfVoterExists", _orgId, _acct)
	return *ret0, err
}

// CheckIfVoterExists is a free data retrieval call binding the contract method 0x00b813df.
//
// Solidity: function checkIfVoterExists(_orgId string, _acct address) constant returns(bool)
func (_PermImpl *PermImplSession) CheckIfVoterExists(_orgId string, _acct common.Address) (bool, error) {
	return _PermImpl.Contract.CheckIfVoterExists(&_PermImpl.CallOpts, _orgId, _acct)
}

// CheckIfVoterExists is a free data retrieval call binding the contract method 0x00b813df.
//
// Solidity: function checkIfVoterExists(_orgId string, _acct address) constant returns(bool)
func (_PermImpl *PermImplCallerSession) CheckIfVoterExists(_orgId string, _acct common.Address) (bool, error) {
	return _PermImpl.Contract.CheckIfVoterExists(&_PermImpl.CallOpts, _orgId, _acct)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(_acct address) constant returns(address, string, string, uint256, bool)
func (_PermImpl *PermImplCaller) GetAccountDetails(opts *bind.CallOpts, _acct common.Address) (common.Address, string, string, *big.Int, bool, error) {
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
	err := _PermImpl.contract.Call(opts, out, "getAccountDetails", _acct)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(_acct address) constant returns(address, string, string, uint256, bool)
func (_PermImpl *PermImplSession) GetAccountDetails(_acct common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return _PermImpl.Contract.GetAccountDetails(&_PermImpl.CallOpts, _acct)
}

// GetAccountDetails is a free data retrieval call binding the contract method 0x2aceb534.
//
// Solidity: function getAccountDetails(_acct address) constant returns(address, string, string, uint256, bool)
func (_PermImpl *PermImplCallerSession) GetAccountDetails(_acct common.Address) (common.Address, string, string, *big.Int, bool, error) {
	return _PermImpl.Contract.GetAccountDetails(&_PermImpl.CallOpts, _acct)
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

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(nodeIndex uint256) constant returns(string, uint256)
func (_PermImpl *PermImplCaller) GetNodeDetailsFromIndex(opts *bind.CallOpts, nodeIndex *big.Int) (string, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _PermImpl.contract.Call(opts, out, "getNodeDetailsFromIndex", nodeIndex)
	return *ret0, *ret1, err
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(nodeIndex uint256) constant returns(string, uint256)
func (_PermImpl *PermImplSession) GetNodeDetailsFromIndex(nodeIndex *big.Int) (string, *big.Int, error) {
	return _PermImpl.Contract.GetNodeDetailsFromIndex(&_PermImpl.CallOpts, nodeIndex)
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(nodeIndex uint256) constant returns(string, uint256)
func (_PermImpl *PermImplCallerSession) GetNodeDetailsFromIndex(nodeIndex *big.Int) (string, *big.Int, error) {
	return _PermImpl.Contract.GetNodeDetailsFromIndex(&_PermImpl.CallOpts, nodeIndex)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(uint256)
func (_PermImpl *PermImplCaller) GetNodeStatus(opts *bind.CallOpts, _enodeId string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermImpl.contract.Call(opts, out, "getNodeStatus", _enodeId)
	return *ret0, err
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(uint256)
func (_PermImpl *PermImplSession) GetNodeStatus(_enodeId string) (*big.Int, error) {
	return _PermImpl.Contract.GetNodeStatus(&_PermImpl.CallOpts, _enodeId)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(uint256)
func (_PermImpl *PermImplCallerSession) GetNodeStatus(_enodeId string) (*big.Int, error) {
	return _PermImpl.Contract.GetNodeStatus(&_PermImpl.CallOpts, _enodeId)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_PermImpl *PermImplCaller) GetNumberOfNodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermImpl.contract.Call(opts, out, "getNumberOfNodes")
	return *ret0, err
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_PermImpl *PermImplSession) GetNumberOfNodes() (*big.Int, error) {
	return _PermImpl.Contract.GetNumberOfNodes(&_PermImpl.CallOpts)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_PermImpl *PermImplCallerSession) GetNumberOfNodes() (*big.Int, error) {
	return _PermImpl.Contract.GetNumberOfNodes(&_PermImpl.CallOpts)
}

// GetNumberOfVoters is a free data retrieval call binding the contract method 0x9b904f0a.
//
// Solidity: function getNumberOfVoters(_orgId string) constant returns(uint256)
func (_PermImpl *PermImplCaller) GetNumberOfVoters(opts *bind.CallOpts, _orgId string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PermImpl.contract.Call(opts, out, "getNumberOfVoters", _orgId)
	return *ret0, err
}

// GetNumberOfVoters is a free data retrieval call binding the contract method 0x9b904f0a.
//
// Solidity: function getNumberOfVoters(_orgId string) constant returns(uint256)
func (_PermImpl *PermImplSession) GetNumberOfVoters(_orgId string) (*big.Int, error) {
	return _PermImpl.Contract.GetNumberOfVoters(&_PermImpl.CallOpts, _orgId)
}

// GetNumberOfVoters is a free data retrieval call binding the contract method 0x9b904f0a.
//
// Solidity: function getNumberOfVoters(_orgId string) constant returns(uint256)
func (_PermImpl *PermImplCallerSession) GetNumberOfVoters(_orgId string) (*big.Int, error) {
	return _PermImpl.Contract.GetNumberOfVoters(&_PermImpl.CallOpts, _orgId)
}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(_orgIndex uint256) constant returns(string, uint256)
func (_PermImpl *PermImplCaller) GetOrgInfo(opts *bind.CallOpts, _orgIndex *big.Int) (string, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _PermImpl.contract.Call(opts, out, "getOrgInfo", _orgIndex)
	return *ret0, *ret1, err
}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(_orgIndex uint256) constant returns(string, uint256)
func (_PermImpl *PermImplSession) GetOrgInfo(_orgIndex *big.Int) (string, *big.Int, error) {
	return _PermImpl.Contract.GetOrgInfo(&_PermImpl.CallOpts, _orgIndex)
}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(_orgIndex uint256) constant returns(string, uint256)
func (_PermImpl *PermImplCallerSession) GetOrgInfo(_orgIndex *big.Int) (string, *big.Int, error) {
	return _PermImpl.Contract.GetOrgInfo(&_PermImpl.CallOpts, _orgIndex)
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

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(_roleId string, _orgId string) constant returns(string, string, uint256, bool, bool)
func (_PermImpl *PermImplCaller) GetRoleDetails(opts *bind.CallOpts, _roleId string, _orgId string) (string, string, *big.Int, bool, bool, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _PermImpl.contract.Call(opts, out, "getRoleDetails", _roleId, _orgId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(_roleId string, _orgId string) constant returns(string, string, uint256, bool, bool)
func (_PermImpl *PermImplSession) GetRoleDetails(_roleId string, _orgId string) (string, string, *big.Int, bool, bool, error) {
	return _PermImpl.Contract.GetRoleDetails(&_PermImpl.CallOpts, _roleId, _orgId)
}

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(_roleId string, _orgId string) constant returns(string, string, uint256, bool, bool)
func (_PermImpl *PermImplCallerSession) GetRoleDetails(_roleId string, _orgId string) (string, string, *big.Int, bool, bool, error) {
	return _PermImpl.Contract.GetRoleDetails(&_PermImpl.CallOpts, _roleId, _orgId)
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(_orgId string) constant returns(uint256, uint256)
func (_PermImpl *PermImplCaller) GetVoteCount(opts *bind.CallOpts, _orgId string) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _PermImpl.contract.Call(opts, out, "getVoteCount", _orgId)
	return *ret0, *ret1, err
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(_orgId string) constant returns(uint256, uint256)
func (_PermImpl *PermImplSession) GetVoteCount(_orgId string) (*big.Int, *big.Int, error) {
	return _PermImpl.Contract.GetVoteCount(&_PermImpl.CallOpts, _orgId)
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(_orgId string) constant returns(uint256, uint256)
func (_PermImpl *PermImplCallerSession) GetVoteCount(_orgId string) (*big.Int, *big.Int, error) {
	return _PermImpl.Contract.GetVoteCount(&_PermImpl.CallOpts, _orgId)
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

// AddNewRole is a paid mutator transaction binding the contract method 0xa2ca82fc.
//
// Solidity: function addNewRole(_roleId string, _orgId string, _access uint256, _voter bool, _caller address) returns()
func (_PermImpl *PermImplTransactor) AddNewRole(opts *bind.TransactOpts, _roleId string, _orgId string, _access *big.Int, _voter bool, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addNewRole", _roleId, _orgId, _access, _voter, _caller)
}

// AddNewRole is a paid mutator transaction binding the contract method 0xa2ca82fc.
//
// Solidity: function addNewRole(_roleId string, _orgId string, _access uint256, _voter bool, _caller address) returns()
func (_PermImpl *PermImplSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddNewRole(&_PermImpl.TransactOpts, _roleId, _orgId, _access, _voter, _caller)
}

// AddNewRole is a paid mutator transaction binding the contract method 0xa2ca82fc.
//
// Solidity: function addNewRole(_roleId string, _orgId string, _access uint256, _voter bool, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) AddNewRole(_roleId string, _orgId string, _access *big.Int, _voter bool, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddNewRole(&_PermImpl.TransactOpts, _roleId, _orgId, _access, _voter, _caller)
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

// AddOrg is a paid mutator transaction binding the contract method 0x8f362a3e.
//
// Solidity: function addOrg(_orgId string, _enodeId string, _caller address) returns()
func (_PermImpl *PermImplTransactor) AddOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "addOrg", _orgId, _enodeId, _caller)
}

// AddOrg is a paid mutator transaction binding the contract method 0x8f362a3e.
//
// Solidity: function addOrg(_orgId string, _enodeId string, _caller address) returns()
func (_PermImpl *PermImplSession) AddOrg(_orgId string, _enodeId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _caller)
}

// AddOrg is a paid mutator transaction binding the contract method 0x8f362a3e.
//
// Solidity: function addOrg(_orgId string, _enodeId string, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) AddOrg(_orgId string, _enodeId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AddOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x7e461258.
//
// Solidity: function approveOrg(_orgId string, _enodeId string, _caller address) returns()
func (_PermImpl *PermImplTransactor) ApproveOrg(opts *bind.TransactOpts, _orgId string, _enodeId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "approveOrg", _orgId, _enodeId, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x7e461258.
//
// Solidity: function approveOrg(_orgId string, _enodeId string, _caller address) returns()
func (_PermImpl *PermImplSession) ApproveOrg(_orgId string, _enodeId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _caller)
}

// ApproveOrg is a paid mutator transaction binding the contract method 0x7e461258.
//
// Solidity: function approveOrg(_orgId string, _enodeId string, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) ApproveOrg(_orgId string, _enodeId string, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrg(&_PermImpl.TransactOpts, _orgId, _enodeId, _caller)
}

// ApproveOrgAdminAccount is a paid mutator transaction binding the contract method 0x940132d6.
//
// Solidity: function approveOrgAdminAccount(_account address, _caller address) returns()
func (_PermImpl *PermImplTransactor) ApproveOrgAdminAccount(opts *bind.TransactOpts, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "approveOrgAdminAccount", _account, _caller)
}

// ApproveOrgAdminAccount is a paid mutator transaction binding the contract method 0x940132d6.
//
// Solidity: function approveOrgAdminAccount(_account address, _caller address) returns()
func (_PermImpl *PermImplSession) ApproveOrgAdminAccount(_account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrgAdminAccount(&_PermImpl.TransactOpts, _account, _caller)
}

// ApproveOrgAdminAccount is a paid mutator transaction binding the contract method 0x940132d6.
//
// Solidity: function approveOrgAdminAccount(_account address, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) ApproveOrgAdminAccount(_account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.ApproveOrgAdminAccount(&_PermImpl.TransactOpts, _account, _caller)
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

// AssignOrgAdminAccount is a paid mutator transaction binding the contract method 0x98ea3fa1.
//
// Solidity: function assignOrgAdminAccount(_orgId string, _account address, _caller address) returns()
func (_PermImpl *PermImplTransactor) AssignOrgAdminAccount(opts *bind.TransactOpts, _orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "assignOrgAdminAccount", _orgId, _account, _caller)
}

// AssignOrgAdminAccount is a paid mutator transaction binding the contract method 0x98ea3fa1.
//
// Solidity: function assignOrgAdminAccount(_orgId string, _account address, _caller address) returns()
func (_PermImpl *PermImplSession) AssignOrgAdminAccount(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AssignOrgAdminAccount(&_PermImpl.TransactOpts, _orgId, _account, _caller)
}

// AssignOrgAdminAccount is a paid mutator transaction binding the contract method 0x98ea3fa1.
//
// Solidity: function assignOrgAdminAccount(_orgId string, _account address, _caller address) returns()
func (_PermImpl *PermImplTransactorSession) AssignOrgAdminAccount(_orgId string, _account common.Address, _caller common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.AssignOrgAdminAccount(&_PermImpl.TransactOpts, _orgId, _account, _caller)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(_orgManager address, _rolesManager address, _acctManager address, _voterManager address, _nodeManager address) returns()
func (_PermImpl *PermImplTransactor) Init(opts *bind.TransactOpts, _orgManager common.Address, _rolesManager common.Address, _acctManager common.Address, _voterManager common.Address, _nodeManager common.Address) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "init", _orgManager, _rolesManager, _acctManager, _voterManager, _nodeManager)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(_orgManager address, _rolesManager address, _acctManager address, _voterManager address, _nodeManager address) returns()
func (_PermImpl *PermImplSession) Init(_orgManager common.Address, _rolesManager common.Address, _acctManager common.Address, _voterManager common.Address, _nodeManager common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.Init(&_PermImpl.TransactOpts, _orgManager, _rolesManager, _acctManager, _voterManager, _nodeManager)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(_orgManager address, _rolesManager address, _acctManager address, _voterManager address, _nodeManager address) returns()
func (_PermImpl *PermImplTransactorSession) Init(_orgManager common.Address, _rolesManager common.Address, _acctManager common.Address, _voterManager common.Address, _nodeManager common.Address) (*types.Transaction, error) {
	return _PermImpl.Contract.Init(&_PermImpl.TransactOpts, _orgManager, _rolesManager, _acctManager, _voterManager, _nodeManager)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(_roleId string, _orgId string) returns()
func (_PermImpl *PermImplTransactor) RemoveRole(opts *bind.TransactOpts, _roleId string, _orgId string) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "removeRole", _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(_roleId string, _orgId string) returns()
func (_PermImpl *PermImplSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _PermImpl.Contract.RemoveRole(&_PermImpl.TransactOpts, _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(_roleId string, _orgId string) returns()
func (_PermImpl *PermImplTransactorSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _PermImpl.Contract.RemoveRole(&_PermImpl.TransactOpts, _roleId, _orgId)
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

// PermImplDummyIterator is returned from FilterDummy and is used to iterate over the raw logs and unpacked data for Dummy events raised by the PermImpl contract.
type PermImplDummyIterator struct {
	Event *PermImplDummy // Event containing the contract specifics and raw log

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
func (it *PermImplDummyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermImplDummy)
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
		it.Event = new(PermImplDummy)
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
func (it *PermImplDummyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermImplDummyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermImplDummy represents a Dummy event raised by the PermImpl contract.
type PermImplDummy struct {
	Msg string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDummy is a free log retrieval operation binding the contract event 0xe4909ae09a5f09db1c974cfab835cf594054bde73d77a5bd128f2d5842036a66.
//
// Solidity: e Dummy(_msg string)
func (_PermImpl *PermImplFilterer) FilterDummy(opts *bind.FilterOpts) (*PermImplDummyIterator, error) {

	logs, sub, err := _PermImpl.contract.FilterLogs(opts, "Dummy")
	if err != nil {
		return nil, err
	}
	return &PermImplDummyIterator{contract: _PermImpl.contract, event: "Dummy", logs: logs, sub: sub}, nil
}

// WatchDummy is a free log subscription operation binding the contract event 0xe4909ae09a5f09db1c974cfab835cf594054bde73d77a5bd128f2d5842036a66.
//
// Solidity: e Dummy(_msg string)
func (_PermImpl *PermImplFilterer) WatchDummy(opts *bind.WatchOpts, sink chan<- *PermImplDummy) (event.Subscription, error) {

	logs, sub, err := _PermImpl.contract.WatchLogs(opts, "Dummy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermImplDummy)
				if err := _PermImpl.contract.UnpackLog(event, "Dummy", log); err != nil {
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
