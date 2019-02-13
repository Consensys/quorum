// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cluster

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

// ClusterABI is the input ABI used to generate the binding from.
const ClusterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_morgId\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"checkIfVoterExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_morgId\",\"type\":\"string\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getVoter\",\"outputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_morgId\",\"type\":\"string\"}],\"name\":\"addSubOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getOrgKeyCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getOrgPendingOp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"approvePendingOp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"checkIfKeyExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"deleteOrgKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_keyIndex\",\"type\":\"uint256\"}],\"name\":\"getOrgKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_morgId\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_morgId\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"deleteVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgIndex\",\"type\":\"uint256\"}],\"name\":\"getOrgInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfOrgs\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"checkKeyClash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_morgId\",\"type\":\"string\"}],\"name\":\"getNumberOfVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_morgId\",\"type\":\"string\"}],\"name\":\"addMasterOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"checkVotingAccountExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"addOrgKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_morgId\",\"type\":\"string\"}],\"name\":\"checkMasterOrgExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getOrgVoteCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkOrgContractExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getPendingOp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"checkOrgPendingOp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"checkOrgExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"MasterOrgAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"SubOrgAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"OrgKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"OrgKeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_pendingOp\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"ItemForApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"VoterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"VoterDeleted\",\"type\":\"event\"}]"

// Cluster is an auto generated Go binding around an Ethereum contract.
type Cluster struct {
	ClusterCaller     // Read-only binding to the contract
	ClusterTransactor // Write-only binding to the contract
	ClusterFilterer   // Log filterer for contract events
}

// ClusterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClusterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClusterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClusterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClusterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClusterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClusterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClusterSession struct {
	Contract     *Cluster          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClusterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClusterCallerSession struct {
	Contract *ClusterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ClusterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClusterTransactorSession struct {
	Contract     *ClusterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ClusterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClusterRaw struct {
	Contract *Cluster // Generic contract binding to access the raw methods on
}

// ClusterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClusterCallerRaw struct {
	Contract *ClusterCaller // Generic read-only contract binding to access the raw methods on
}

// ClusterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClusterTransactorRaw struct {
	Contract *ClusterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCluster creates a new instance of Cluster, bound to a specific deployed contract.
func NewCluster(address common.Address, backend bind.ContractBackend) (*Cluster, error) {
	contract, err := bindCluster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cluster{ClusterCaller: ClusterCaller{contract: contract}, ClusterTransactor: ClusterTransactor{contract: contract}, ClusterFilterer: ClusterFilterer{contract: contract}}, nil
}

// NewClusterCaller creates a new read-only instance of Cluster, bound to a specific deployed contract.
func NewClusterCaller(address common.Address, caller bind.ContractCaller) (*ClusterCaller, error) {
	contract, err := bindCluster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClusterCaller{contract: contract}, nil
}

// NewClusterTransactor creates a new write-only instance of Cluster, bound to a specific deployed contract.
func NewClusterTransactor(address common.Address, transactor bind.ContractTransactor) (*ClusterTransactor, error) {
	contract, err := bindCluster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClusterTransactor{contract: contract}, nil
}

// NewClusterFilterer creates a new log filterer instance of Cluster, bound to a specific deployed contract.
func NewClusterFilterer(address common.Address, filterer bind.ContractFilterer) (*ClusterFilterer, error) {
	contract, err := bindCluster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClusterFilterer{contract: contract}, nil
}

// bindCluster binds a generic wrapper to an already deployed contract.
func bindCluster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClusterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cluster *ClusterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cluster.Contract.ClusterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cluster *ClusterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cluster.Contract.ClusterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cluster *ClusterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cluster.Contract.ClusterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cluster *ClusterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cluster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cluster *ClusterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cluster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cluster *ClusterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cluster.Contract.contract.Transact(opts, method, params...)
}

// CheckIfKeyExists is a free data retrieval call binding the contract method 0x4898598e.
//
// Solidity: function checkIfKeyExists(_orgId string, _tmKey string) constant returns(bool)
func (_Cluster *ClusterCaller) CheckIfKeyExists(opts *bind.CallOpts, _orgId string, _tmKey string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "checkIfKeyExists", _orgId, _tmKey)
	return *ret0, err
}

// CheckIfKeyExists is a free data retrieval call binding the contract method 0x4898598e.
//
// Solidity: function checkIfKeyExists(_orgId string, _tmKey string) constant returns(bool)
func (_Cluster *ClusterSession) CheckIfKeyExists(_orgId string, _tmKey string) (bool, error) {
	return _Cluster.Contract.CheckIfKeyExists(&_Cluster.CallOpts, _orgId, _tmKey)
}

// CheckIfKeyExists is a free data retrieval call binding the contract method 0x4898598e.
//
// Solidity: function checkIfKeyExists(_orgId string, _tmKey string) constant returns(bool)
func (_Cluster *ClusterCallerSession) CheckIfKeyExists(_orgId string, _tmKey string) (bool, error) {
	return _Cluster.Contract.CheckIfKeyExists(&_Cluster.CallOpts, _orgId, _tmKey)
}

// CheckIfVoterExists is a free data retrieval call binding the contract method 0x00b813df.
//
// Solidity: function checkIfVoterExists(_morgId string, _address address) constant returns(bool)
func (_Cluster *ClusterCaller) CheckIfVoterExists(opts *bind.CallOpts, _morgId string, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "checkIfVoterExists", _morgId, _address)
	return *ret0, err
}

// CheckIfVoterExists is a free data retrieval call binding the contract method 0x00b813df.
//
// Solidity: function checkIfVoterExists(_morgId string, _address address) constant returns(bool)
func (_Cluster *ClusterSession) CheckIfVoterExists(_morgId string, _address common.Address) (bool, error) {
	return _Cluster.Contract.CheckIfVoterExists(&_Cluster.CallOpts, _morgId, _address)
}

// CheckIfVoterExists is a free data retrieval call binding the contract method 0x00b813df.
//
// Solidity: function checkIfVoterExists(_morgId string, _address address) constant returns(bool)
func (_Cluster *ClusterCallerSession) CheckIfVoterExists(_morgId string, _address common.Address) (bool, error) {
	return _Cluster.Contract.CheckIfVoterExists(&_Cluster.CallOpts, _morgId, _address)
}

// CheckKeyClash is a free data retrieval call binding the contract method 0x8fde9c5e.
//
// Solidity: function checkKeyClash(_orgId string, _key string) constant returns(bool)
func (_Cluster *ClusterCaller) CheckKeyClash(opts *bind.CallOpts, _orgId string, _key string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "checkKeyClash", _orgId, _key)
	return *ret0, err
}

// CheckKeyClash is a free data retrieval call binding the contract method 0x8fde9c5e.
//
// Solidity: function checkKeyClash(_orgId string, _key string) constant returns(bool)
func (_Cluster *ClusterSession) CheckKeyClash(_orgId string, _key string) (bool, error) {
	return _Cluster.Contract.CheckKeyClash(&_Cluster.CallOpts, _orgId, _key)
}

// CheckKeyClash is a free data retrieval call binding the contract method 0x8fde9c5e.
//
// Solidity: function checkKeyClash(_orgId string, _key string) constant returns(bool)
func (_Cluster *ClusterCallerSession) CheckKeyClash(_orgId string, _key string) (bool, error) {
	return _Cluster.Contract.CheckKeyClash(&_Cluster.CallOpts, _orgId, _key)
}

// CheckMasterOrgExists is a free data retrieval call binding the contract method 0xd912967a.
//
// Solidity: function checkMasterOrgExists(_morgId string) constant returns(bool)
func (_Cluster *ClusterCaller) CheckMasterOrgExists(opts *bind.CallOpts, _morgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "checkMasterOrgExists", _morgId)
	return *ret0, err
}

// CheckMasterOrgExists is a free data retrieval call binding the contract method 0xd912967a.
//
// Solidity: function checkMasterOrgExists(_morgId string) constant returns(bool)
func (_Cluster *ClusterSession) CheckMasterOrgExists(_morgId string) (bool, error) {
	return _Cluster.Contract.CheckMasterOrgExists(&_Cluster.CallOpts, _morgId)
}

// CheckMasterOrgExists is a free data retrieval call binding the contract method 0xd912967a.
//
// Solidity: function checkMasterOrgExists(_morgId string) constant returns(bool)
func (_Cluster *ClusterCallerSession) CheckMasterOrgExists(_morgId string) (bool, error) {
	return _Cluster.Contract.CheckMasterOrgExists(&_Cluster.CallOpts, _morgId)
}

// CheckOrgContractExists is a free data retrieval call binding the contract method 0xee0c7dda.
//
// Solidity: function checkOrgContractExists() constant returns(bool)
func (_Cluster *ClusterCaller) CheckOrgContractExists(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "checkOrgContractExists")
	return *ret0, err
}

// CheckOrgContractExists is a free data retrieval call binding the contract method 0xee0c7dda.
//
// Solidity: function checkOrgContractExists() constant returns(bool)
func (_Cluster *ClusterSession) CheckOrgContractExists() (bool, error) {
	return _Cluster.Contract.CheckOrgContractExists(&_Cluster.CallOpts)
}

// CheckOrgContractExists is a free data retrieval call binding the contract method 0xee0c7dda.
//
// Solidity: function checkOrgContractExists() constant returns(bool)
func (_Cluster *ClusterCallerSession) CheckOrgContractExists() (bool, error) {
	return _Cluster.Contract.CheckOrgContractExists(&_Cluster.CallOpts)
}

// CheckOrgExists is a free data retrieval call binding the contract method 0xffe40d1d.
//
// Solidity: function checkOrgExists(_orgId string) constant returns(bool)
func (_Cluster *ClusterCaller) CheckOrgExists(opts *bind.CallOpts, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "checkOrgExists", _orgId)
	return *ret0, err
}

// CheckOrgExists is a free data retrieval call binding the contract method 0xffe40d1d.
//
// Solidity: function checkOrgExists(_orgId string) constant returns(bool)
func (_Cluster *ClusterSession) CheckOrgExists(_orgId string) (bool, error) {
	return _Cluster.Contract.CheckOrgExists(&_Cluster.CallOpts, _orgId)
}

// CheckOrgExists is a free data retrieval call binding the contract method 0xffe40d1d.
//
// Solidity: function checkOrgExists(_orgId string) constant returns(bool)
func (_Cluster *ClusterCallerSession) CheckOrgExists(_orgId string) (bool, error) {
	return _Cluster.Contract.CheckOrgExists(&_Cluster.CallOpts, _orgId)
}

// CheckOrgPendingOp is a free data retrieval call binding the contract method 0xfb23dedc.
//
// Solidity: function checkOrgPendingOp(_orgId string) constant returns(bool)
func (_Cluster *ClusterCaller) CheckOrgPendingOp(opts *bind.CallOpts, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "checkOrgPendingOp", _orgId)
	return *ret0, err
}

// CheckOrgPendingOp is a free data retrieval call binding the contract method 0xfb23dedc.
//
// Solidity: function checkOrgPendingOp(_orgId string) constant returns(bool)
func (_Cluster *ClusterSession) CheckOrgPendingOp(_orgId string) (bool, error) {
	return _Cluster.Contract.CheckOrgPendingOp(&_Cluster.CallOpts, _orgId)
}

// CheckOrgPendingOp is a free data retrieval call binding the contract method 0xfb23dedc.
//
// Solidity: function checkOrgPendingOp(_orgId string) constant returns(bool)
func (_Cluster *ClusterCallerSession) CheckOrgPendingOp(_orgId string) (bool, error) {
	return _Cluster.Contract.CheckOrgPendingOp(&_Cluster.CallOpts, _orgId)
}

// CheckVotingAccountExists is a free data retrieval call binding the contract method 0xcb2c45dc.
//
// Solidity: function checkVotingAccountExists(_orgId string) constant returns(bool)
func (_Cluster *ClusterCaller) CheckVotingAccountExists(opts *bind.CallOpts, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "checkVotingAccountExists", _orgId)
	return *ret0, err
}

// CheckVotingAccountExists is a free data retrieval call binding the contract method 0xcb2c45dc.
//
// Solidity: function checkVotingAccountExists(_orgId string) constant returns(bool)
func (_Cluster *ClusterSession) CheckVotingAccountExists(_orgId string) (bool, error) {
	return _Cluster.Contract.CheckVotingAccountExists(&_Cluster.CallOpts, _orgId)
}

// CheckVotingAccountExists is a free data retrieval call binding the contract method 0xcb2c45dc.
//
// Solidity: function checkVotingAccountExists(_orgId string) constant returns(bool)
func (_Cluster *ClusterCallerSession) CheckVotingAccountExists(_orgId string) (bool, error) {
	return _Cluster.Contract.CheckVotingAccountExists(&_Cluster.CallOpts, _orgId)
}

// GetNumberOfOrgs is a free data retrieval call binding the contract method 0x7755ebdd.
//
// Solidity: function getNumberOfOrgs() constant returns(uint256)
func (_Cluster *ClusterCaller) GetNumberOfOrgs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "getNumberOfOrgs")
	return *ret0, err
}

// GetNumberOfOrgs is a free data retrieval call binding the contract method 0x7755ebdd.
//
// Solidity: function getNumberOfOrgs() constant returns(uint256)
func (_Cluster *ClusterSession) GetNumberOfOrgs() (*big.Int, error) {
	return _Cluster.Contract.GetNumberOfOrgs(&_Cluster.CallOpts)
}

// GetNumberOfOrgs is a free data retrieval call binding the contract method 0x7755ebdd.
//
// Solidity: function getNumberOfOrgs() constant returns(uint256)
func (_Cluster *ClusterCallerSession) GetNumberOfOrgs() (*big.Int, error) {
	return _Cluster.Contract.GetNumberOfOrgs(&_Cluster.CallOpts)
}

// GetNumberOfVoters is a free data retrieval call binding the contract method 0x9b904f0a.
//
// Solidity: function getNumberOfVoters(_morgId string) constant returns(uint256)
func (_Cluster *ClusterCaller) GetNumberOfVoters(opts *bind.CallOpts, _morgId string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "getNumberOfVoters", _morgId)
	return *ret0, err
}

// GetNumberOfVoters is a free data retrieval call binding the contract method 0x9b904f0a.
//
// Solidity: function getNumberOfVoters(_morgId string) constant returns(uint256)
func (_Cluster *ClusterSession) GetNumberOfVoters(_morgId string) (*big.Int, error) {
	return _Cluster.Contract.GetNumberOfVoters(&_Cluster.CallOpts, _morgId)
}

// GetNumberOfVoters is a free data retrieval call binding the contract method 0x9b904f0a.
//
// Solidity: function getNumberOfVoters(_morgId string) constant returns(uint256)
func (_Cluster *ClusterCallerSession) GetNumberOfVoters(_morgId string) (*big.Int, error) {
	return _Cluster.Contract.GetNumberOfVoters(&_Cluster.CallOpts, _morgId)
}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(_orgIndex uint256) constant returns(string, string)
func (_Cluster *ClusterCaller) GetOrgInfo(opts *bind.CallOpts, _orgIndex *big.Int) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Cluster.contract.Call(opts, out, "getOrgInfo", _orgIndex)
	return *ret0, *ret1, err
}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(_orgIndex uint256) constant returns(string, string)
func (_Cluster *ClusterSession) GetOrgInfo(_orgIndex *big.Int) (string, string, error) {
	return _Cluster.Contract.GetOrgInfo(&_Cluster.CallOpts, _orgIndex)
}

// GetOrgInfo is a free data retrieval call binding the contract method 0x5c4f32ee.
//
// Solidity: function getOrgInfo(_orgIndex uint256) constant returns(string, string)
func (_Cluster *ClusterCallerSession) GetOrgInfo(_orgIndex *big.Int) (string, string, error) {
	return _Cluster.Contract.GetOrgInfo(&_Cluster.CallOpts, _orgIndex)
}

// GetOrgKey is a free data retrieval call binding the contract method 0x5002dadf.
//
// Solidity: function getOrgKey(_orgId string, _keyIndex uint256) constant returns(string, bool)
func (_Cluster *ClusterCaller) GetOrgKey(opts *bind.CallOpts, _orgId string, _keyIndex *big.Int) (string, bool, error) {
	var (
		ret0 = new(string)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Cluster.contract.Call(opts, out, "getOrgKey", _orgId, _keyIndex)
	return *ret0, *ret1, err
}

// GetOrgKey is a free data retrieval call binding the contract method 0x5002dadf.
//
// Solidity: function getOrgKey(_orgId string, _keyIndex uint256) constant returns(string, bool)
func (_Cluster *ClusterSession) GetOrgKey(_orgId string, _keyIndex *big.Int) (string, bool, error) {
	return _Cluster.Contract.GetOrgKey(&_Cluster.CallOpts, _orgId, _keyIndex)
}

// GetOrgKey is a free data retrieval call binding the contract method 0x5002dadf.
//
// Solidity: function getOrgKey(_orgId string, _keyIndex uint256) constant returns(string, bool)
func (_Cluster *ClusterCallerSession) GetOrgKey(_orgId string, _keyIndex *big.Int) (string, bool, error) {
	return _Cluster.Contract.GetOrgKey(&_Cluster.CallOpts, _orgId, _keyIndex)
}

// GetOrgKeyCount is a free data retrieval call binding the contract method 0x243cc506.
//
// Solidity: function getOrgKeyCount(_orgId string) constant returns(uint256)
func (_Cluster *ClusterCaller) GetOrgKeyCount(opts *bind.CallOpts, _orgId string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "getOrgKeyCount", _orgId)
	return *ret0, err
}

// GetOrgKeyCount is a free data retrieval call binding the contract method 0x243cc506.
//
// Solidity: function getOrgKeyCount(_orgId string) constant returns(uint256)
func (_Cluster *ClusterSession) GetOrgKeyCount(_orgId string) (*big.Int, error) {
	return _Cluster.Contract.GetOrgKeyCount(&_Cluster.CallOpts, _orgId)
}

// GetOrgKeyCount is a free data retrieval call binding the contract method 0x243cc506.
//
// Solidity: function getOrgKeyCount(_orgId string) constant returns(uint256)
func (_Cluster *ClusterCallerSession) GetOrgKeyCount(_orgId string) (*big.Int, error) {
	return _Cluster.Contract.GetOrgKeyCount(&_Cluster.CallOpts, _orgId)
}

// GetOrgPendingOp is a free data retrieval call binding the contract method 0x33680eb7.
//
// Solidity: function getOrgPendingOp(_orgId string) constant returns(string, uint8)
func (_Cluster *ClusterCaller) GetOrgPendingOp(opts *bind.CallOpts, _orgId string) (string, uint8, error) {
	var (
		ret0 = new(string)
		ret1 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Cluster.contract.Call(opts, out, "getOrgPendingOp", _orgId)
	return *ret0, *ret1, err
}

// GetOrgPendingOp is a free data retrieval call binding the contract method 0x33680eb7.
//
// Solidity: function getOrgPendingOp(_orgId string) constant returns(string, uint8)
func (_Cluster *ClusterSession) GetOrgPendingOp(_orgId string) (string, uint8, error) {
	return _Cluster.Contract.GetOrgPendingOp(&_Cluster.CallOpts, _orgId)
}

// GetOrgPendingOp is a free data retrieval call binding the contract method 0x33680eb7.
//
// Solidity: function getOrgPendingOp(_orgId string) constant returns(string, uint8)
func (_Cluster *ClusterCallerSession) GetOrgPendingOp(_orgId string) (string, uint8, error) {
	return _Cluster.Contract.GetOrgPendingOp(&_Cluster.CallOpts, _orgId)
}

// GetOrgVoteCount is a free data retrieval call binding the contract method 0xe7089a0c.
//
// Solidity: function getOrgVoteCount(_orgId string) constant returns(uint256)
func (_Cluster *ClusterCaller) GetOrgVoteCount(opts *bind.CallOpts, _orgId string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "getOrgVoteCount", _orgId)
	return *ret0, err
}

// GetOrgVoteCount is a free data retrieval call binding the contract method 0xe7089a0c.
//
// Solidity: function getOrgVoteCount(_orgId string) constant returns(uint256)
func (_Cluster *ClusterSession) GetOrgVoteCount(_orgId string) (*big.Int, error) {
	return _Cluster.Contract.GetOrgVoteCount(&_Cluster.CallOpts, _orgId)
}

// GetOrgVoteCount is a free data retrieval call binding the contract method 0xe7089a0c.
//
// Solidity: function getOrgVoteCount(_orgId string) constant returns(uint256)
func (_Cluster *ClusterCallerSession) GetOrgVoteCount(_orgId string) (*big.Int, error) {
	return _Cluster.Contract.GetOrgVoteCount(&_Cluster.CallOpts, _orgId)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(_orgId string) constant returns(string, uint8)
func (_Cluster *ClusterCaller) GetPendingOp(opts *bind.CallOpts, _orgId string) (string, uint8, error) {
	var (
		ret0 = new(string)
		ret1 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Cluster.contract.Call(opts, out, "getPendingOp", _orgId)
	return *ret0, *ret1, err
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(_orgId string) constant returns(string, uint8)
func (_Cluster *ClusterSession) GetPendingOp(_orgId string) (string, uint8, error) {
	return _Cluster.Contract.GetPendingOp(&_Cluster.CallOpts, _orgId)
}

// GetPendingOp is a free data retrieval call binding the contract method 0xf346a3a7.
//
// Solidity: function getPendingOp(_orgId string) constant returns(string, uint8)
func (_Cluster *ClusterCallerSession) GetPendingOp(_orgId string) (string, uint8, error) {
	return _Cluster.Contract.GetPendingOp(&_Cluster.CallOpts, _orgId)
}

// GetVoter is a free data retrieval call binding the contract method 0x17a2fb72.
//
// Solidity: function getVoter(_morgId string, i uint256) constant returns(_addr address, _active bool)
func (_Cluster *ClusterCaller) GetVoter(opts *bind.CallOpts, _morgId string, i *big.Int) (struct {
	Addr   common.Address
	Active bool
}, error) {
	ret := new(struct {
		Addr   common.Address
		Active bool
	})
	out := ret
	err := _Cluster.contract.Call(opts, out, "getVoter", _morgId, i)
	return *ret, err
}

// GetVoter is a free data retrieval call binding the contract method 0x17a2fb72.
//
// Solidity: function getVoter(_morgId string, i uint256) constant returns(_addr address, _active bool)
func (_Cluster *ClusterSession) GetVoter(_morgId string, i *big.Int) (struct {
	Addr   common.Address
	Active bool
}, error) {
	return _Cluster.Contract.GetVoter(&_Cluster.CallOpts, _morgId, i)
}

// GetVoter is a free data retrieval call binding the contract method 0x17a2fb72.
//
// Solidity: function getVoter(_morgId string, i uint256) constant returns(_addr address, _active bool)
func (_Cluster *ClusterCallerSession) GetVoter(_morgId string, i *big.Int) (struct {
	Addr   common.Address
	Active bool
}, error) {
	return _Cluster.Contract.GetVoter(&_Cluster.CallOpts, _morgId, i)
}

// IsVoter is a free data retrieval call binding the contract method 0xbd9e887a.
//
// Solidity: function isVoter(_orgId string, account address) constant returns(bool)
func (_Cluster *ClusterCaller) IsVoter(opts *bind.CallOpts, _orgId string, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "isVoter", _orgId, account)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xbd9e887a.
//
// Solidity: function isVoter(_orgId string, account address) constant returns(bool)
func (_Cluster *ClusterSession) IsVoter(_orgId string, account common.Address) (bool, error) {
	return _Cluster.Contract.IsVoter(&_Cluster.CallOpts, _orgId, account)
}

// IsVoter is a free data retrieval call binding the contract method 0xbd9e887a.
//
// Solidity: function isVoter(_orgId string, account address) constant returns(bool)
func (_Cluster *ClusterCallerSession) IsVoter(_orgId string, account common.Address) (bool, error) {
	return _Cluster.Contract.IsVoter(&_Cluster.CallOpts, _orgId, account)
}

// AddMasterOrg is a paid mutator transaction binding the contract method 0xc7304f3f.
//
// Solidity: function addMasterOrg(_morgId string) returns()
func (_Cluster *ClusterTransactor) AddMasterOrg(opts *bind.TransactOpts, _morgId string) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "addMasterOrg", _morgId)
}

// AddMasterOrg is a paid mutator transaction binding the contract method 0xc7304f3f.
//
// Solidity: function addMasterOrg(_morgId string) returns()
func (_Cluster *ClusterSession) AddMasterOrg(_morgId string) (*types.Transaction, error) {
	return _Cluster.Contract.AddMasterOrg(&_Cluster.TransactOpts, _morgId)
}

// AddMasterOrg is a paid mutator transaction binding the contract method 0xc7304f3f.
//
// Solidity: function addMasterOrg(_morgId string) returns()
func (_Cluster *ClusterTransactorSession) AddMasterOrg(_morgId string) (*types.Transaction, error) {
	return _Cluster.Contract.AddMasterOrg(&_Cluster.TransactOpts, _morgId)
}

// AddOrgKey is a paid mutator transaction binding the contract method 0xd88ce6bb.
//
// Solidity: function addOrgKey(_orgId string, _tmKey string) returns()
func (_Cluster *ClusterTransactor) AddOrgKey(opts *bind.TransactOpts, _orgId string, _tmKey string) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "addOrgKey", _orgId, _tmKey)
}

// AddOrgKey is a paid mutator transaction binding the contract method 0xd88ce6bb.
//
// Solidity: function addOrgKey(_orgId string, _tmKey string) returns()
func (_Cluster *ClusterSession) AddOrgKey(_orgId string, _tmKey string) (*types.Transaction, error) {
	return _Cluster.Contract.AddOrgKey(&_Cluster.TransactOpts, _orgId, _tmKey)
}

// AddOrgKey is a paid mutator transaction binding the contract method 0xd88ce6bb.
//
// Solidity: function addOrgKey(_orgId string, _tmKey string) returns()
func (_Cluster *ClusterTransactorSession) AddOrgKey(_orgId string, _tmKey string) (*types.Transaction, error) {
	return _Cluster.Contract.AddOrgKey(&_Cluster.TransactOpts, _orgId, _tmKey)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x1f953480.
//
// Solidity: function addSubOrg(_orgId string, _morgId string) returns()
func (_Cluster *ClusterTransactor) AddSubOrg(opts *bind.TransactOpts, _orgId string, _morgId string) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "addSubOrg", _orgId, _morgId)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x1f953480.
//
// Solidity: function addSubOrg(_orgId string, _morgId string) returns()
func (_Cluster *ClusterSession) AddSubOrg(_orgId string, _morgId string) (*types.Transaction, error) {
	return _Cluster.Contract.AddSubOrg(&_Cluster.TransactOpts, _orgId, _morgId)
}

// AddSubOrg is a paid mutator transaction binding the contract method 0x1f953480.
//
// Solidity: function addSubOrg(_orgId string, _morgId string) returns()
func (_Cluster *ClusterTransactorSession) AddSubOrg(_orgId string, _morgId string) (*types.Transaction, error) {
	return _Cluster.Contract.AddSubOrg(&_Cluster.TransactOpts, _orgId, _morgId)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(_morgId string, _address address) returns()
func (_Cluster *ClusterTransactor) AddVoter(opts *bind.TransactOpts, _morgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "addVoter", _morgId, _address)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(_morgId string, _address address) returns()
func (_Cluster *ClusterSession) AddVoter(_morgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.Contract.AddVoter(&_Cluster.TransactOpts, _morgId, _address)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(_morgId string, _address address) returns()
func (_Cluster *ClusterTransactorSession) AddVoter(_morgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.Contract.AddVoter(&_Cluster.TransactOpts, _morgId, _address)
}

// ApprovePendingOp is a paid mutator transaction binding the contract method 0x35dc4772.
//
// Solidity: function approvePendingOp(_orgId string) returns()
func (_Cluster *ClusterTransactor) ApprovePendingOp(opts *bind.TransactOpts, _orgId string) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "approvePendingOp", _orgId)
}

// ApprovePendingOp is a paid mutator transaction binding the contract method 0x35dc4772.
//
// Solidity: function approvePendingOp(_orgId string) returns()
func (_Cluster *ClusterSession) ApprovePendingOp(_orgId string) (*types.Transaction, error) {
	return _Cluster.Contract.ApprovePendingOp(&_Cluster.TransactOpts, _orgId)
}

// ApprovePendingOp is a paid mutator transaction binding the contract method 0x35dc4772.
//
// Solidity: function approvePendingOp(_orgId string) returns()
func (_Cluster *ClusterTransactorSession) ApprovePendingOp(_orgId string) (*types.Transaction, error) {
	return _Cluster.Contract.ApprovePendingOp(&_Cluster.TransactOpts, _orgId)
}

// DeleteOrgKey is a paid mutator transaction binding the contract method 0x49379c50.
//
// Solidity: function deleteOrgKey(_orgId string, _tmKey string) returns()
func (_Cluster *ClusterTransactor) DeleteOrgKey(opts *bind.TransactOpts, _orgId string, _tmKey string) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "deleteOrgKey", _orgId, _tmKey)
}

// DeleteOrgKey is a paid mutator transaction binding the contract method 0x49379c50.
//
// Solidity: function deleteOrgKey(_orgId string, _tmKey string) returns()
func (_Cluster *ClusterSession) DeleteOrgKey(_orgId string, _tmKey string) (*types.Transaction, error) {
	return _Cluster.Contract.DeleteOrgKey(&_Cluster.TransactOpts, _orgId, _tmKey)
}

// DeleteOrgKey is a paid mutator transaction binding the contract method 0x49379c50.
//
// Solidity: function deleteOrgKey(_orgId string, _tmKey string) returns()
func (_Cluster *ClusterTransactorSession) DeleteOrgKey(_orgId string, _tmKey string) (*types.Transaction, error) {
	return _Cluster.Contract.DeleteOrgKey(&_Cluster.TransactOpts, _orgId, _tmKey)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(_morgId string, _address address) returns()
func (_Cluster *ClusterTransactor) DeleteVoter(opts *bind.TransactOpts, _morgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "deleteVoter", _morgId, _address)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(_morgId string, _address address) returns()
func (_Cluster *ClusterSession) DeleteVoter(_morgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.Contract.DeleteVoter(&_Cluster.TransactOpts, _morgId, _address)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(_morgId string, _address address) returns()
func (_Cluster *ClusterTransactorSession) DeleteVoter(_morgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.Contract.DeleteVoter(&_Cluster.TransactOpts, _morgId, _address)
}

// ClusterItemForApprovalIterator is returned from FilterItemForApproval and is used to iterate over the raw logs and unpacked data for ItemForApproval events raised by the Cluster contract.
type ClusterItemForApprovalIterator struct {
	Event *ClusterItemForApproval // Event containing the contract specifics and raw log

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
func (it *ClusterItemForApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterItemForApproval)
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
		it.Event = new(ClusterItemForApproval)
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
func (it *ClusterItemForApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterItemForApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterItemForApproval represents a ItemForApproval event raised by the Cluster contract.
type ClusterItemForApproval struct {
	OrgId     string
	PendingOp uint8
	TmKey     string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterItemForApproval is a free log retrieval operation binding the contract event 0x4475befcee492797e02530076fd7e138aa058eb3bcd028a9df5c0f2815ba9f4a.
//
// Solidity: e ItemForApproval(_orgId string, _pendingOp uint8, _tmKey string)
func (_Cluster *ClusterFilterer) FilterItemForApproval(opts *bind.FilterOpts) (*ClusterItemForApprovalIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "ItemForApproval")
	if err != nil {
		return nil, err
	}
	return &ClusterItemForApprovalIterator{contract: _Cluster.contract, event: "ItemForApproval", logs: logs, sub: sub}, nil
}

// WatchItemForApproval is a free log subscription operation binding the contract event 0x4475befcee492797e02530076fd7e138aa058eb3bcd028a9df5c0f2815ba9f4a.
//
// Solidity: e ItemForApproval(_orgId string, _pendingOp uint8, _tmKey string)
func (_Cluster *ClusterFilterer) WatchItemForApproval(opts *bind.WatchOpts, sink chan<- *ClusterItemForApproval) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "ItemForApproval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterItemForApproval)
				if err := _Cluster.contract.UnpackLog(event, "ItemForApproval", log); err != nil {
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

// ClusterMasterOrgAddedIterator is returned from FilterMasterOrgAdded and is used to iterate over the raw logs and unpacked data for MasterOrgAdded events raised by the Cluster contract.
type ClusterMasterOrgAddedIterator struct {
	Event *ClusterMasterOrgAdded // Event containing the contract specifics and raw log

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
func (it *ClusterMasterOrgAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMasterOrgAdded)
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
		it.Event = new(ClusterMasterOrgAdded)
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
func (it *ClusterMasterOrgAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMasterOrgAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMasterOrgAdded represents a MasterOrgAdded event raised by the Cluster contract.
type ClusterMasterOrgAdded struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMasterOrgAdded is a free log retrieval operation binding the contract event 0xfe62f8d1508aa8ddbb57fd8a6d631f4418cfcbafa90c6ce6d4b8105da5609729.
//
// Solidity: e MasterOrgAdded(_orgId string)
func (_Cluster *ClusterFilterer) FilterMasterOrgAdded(opts *bind.FilterOpts) (*ClusterMasterOrgAddedIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "MasterOrgAdded")
	if err != nil {
		return nil, err
	}
	return &ClusterMasterOrgAddedIterator{contract: _Cluster.contract, event: "MasterOrgAdded", logs: logs, sub: sub}, nil
}

// WatchMasterOrgAdded is a free log subscription operation binding the contract event 0xfe62f8d1508aa8ddbb57fd8a6d631f4418cfcbafa90c6ce6d4b8105da5609729.
//
// Solidity: e MasterOrgAdded(_orgId string)
func (_Cluster *ClusterFilterer) WatchMasterOrgAdded(opts *bind.WatchOpts, sink chan<- *ClusterMasterOrgAdded) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "MasterOrgAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMasterOrgAdded)
				if err := _Cluster.contract.UnpackLog(event, "MasterOrgAdded", log); err != nil {
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

// ClusterOrgKeyAddedIterator is returned from FilterOrgKeyAdded and is used to iterate over the raw logs and unpacked data for OrgKeyAdded events raised by the Cluster contract.
type ClusterOrgKeyAddedIterator struct {
	Event *ClusterOrgKeyAdded // Event containing the contract specifics and raw log

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
func (it *ClusterOrgKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterOrgKeyAdded)
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
		it.Event = new(ClusterOrgKeyAdded)
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
func (it *ClusterOrgKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterOrgKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterOrgKeyAdded represents a OrgKeyAdded event raised by the Cluster contract.
type ClusterOrgKeyAdded struct {
	OrgId string
	TmKey string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrgKeyAdded is a free log retrieval operation binding the contract event 0x6f4d370520782587dabc737a258f46de81ad45d733a42cd5a0045cff1e46deb4.
//
// Solidity: e OrgKeyAdded(_orgId string, _tmKey string)
func (_Cluster *ClusterFilterer) FilterOrgKeyAdded(opts *bind.FilterOpts) (*ClusterOrgKeyAddedIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "OrgKeyAdded")
	if err != nil {
		return nil, err
	}
	return &ClusterOrgKeyAddedIterator{contract: _Cluster.contract, event: "OrgKeyAdded", logs: logs, sub: sub}, nil
}

// WatchOrgKeyAdded is a free log subscription operation binding the contract event 0x6f4d370520782587dabc737a258f46de81ad45d733a42cd5a0045cff1e46deb4.
//
// Solidity: e OrgKeyAdded(_orgId string, _tmKey string)
func (_Cluster *ClusterFilterer) WatchOrgKeyAdded(opts *bind.WatchOpts, sink chan<- *ClusterOrgKeyAdded) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "OrgKeyAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterOrgKeyAdded)
				if err := _Cluster.contract.UnpackLog(event, "OrgKeyAdded", log); err != nil {
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

// ClusterOrgKeyDeletedIterator is returned from FilterOrgKeyDeleted and is used to iterate over the raw logs and unpacked data for OrgKeyDeleted events raised by the Cluster contract.
type ClusterOrgKeyDeletedIterator struct {
	Event *ClusterOrgKeyDeleted // Event containing the contract specifics and raw log

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
func (it *ClusterOrgKeyDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterOrgKeyDeleted)
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
		it.Event = new(ClusterOrgKeyDeleted)
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
func (it *ClusterOrgKeyDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterOrgKeyDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterOrgKeyDeleted represents a OrgKeyDeleted event raised by the Cluster contract.
type ClusterOrgKeyDeleted struct {
	OrgId string
	TmKey string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrgKeyDeleted is a free log retrieval operation binding the contract event 0x2e0a2dc845dce9ef7206b8fe38f3dacaad17ba74d7be9fba469c9858ae16a5d6.
//
// Solidity: e OrgKeyDeleted(_orgId string, _tmKey string)
func (_Cluster *ClusterFilterer) FilterOrgKeyDeleted(opts *bind.FilterOpts) (*ClusterOrgKeyDeletedIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "OrgKeyDeleted")
	if err != nil {
		return nil, err
	}
	return &ClusterOrgKeyDeletedIterator{contract: _Cluster.contract, event: "OrgKeyDeleted", logs: logs, sub: sub}, nil
}

// WatchOrgKeyDeleted is a free log subscription operation binding the contract event 0x2e0a2dc845dce9ef7206b8fe38f3dacaad17ba74d7be9fba469c9858ae16a5d6.
//
// Solidity: e OrgKeyDeleted(_orgId string, _tmKey string)
func (_Cluster *ClusterFilterer) WatchOrgKeyDeleted(opts *bind.WatchOpts, sink chan<- *ClusterOrgKeyDeleted) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "OrgKeyDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterOrgKeyDeleted)
				if err := _Cluster.contract.UnpackLog(event, "OrgKeyDeleted", log); err != nil {
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

// ClusterSubOrgAddedIterator is returned from FilterSubOrgAdded and is used to iterate over the raw logs and unpacked data for SubOrgAdded events raised by the Cluster contract.
type ClusterSubOrgAddedIterator struct {
	Event *ClusterSubOrgAdded // Event containing the contract specifics and raw log

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
func (it *ClusterSubOrgAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterSubOrgAdded)
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
		it.Event = new(ClusterSubOrgAdded)
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
func (it *ClusterSubOrgAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterSubOrgAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterSubOrgAdded represents a SubOrgAdded event raised by the Cluster contract.
type ClusterSubOrgAdded struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubOrgAdded is a free log retrieval operation binding the contract event 0xd734c07873f32f0735016e51dc718e21a48a3bec999d5be38cf3af363fbfedab.
//
// Solidity: e SubOrgAdded(_orgId string)
func (_Cluster *ClusterFilterer) FilterSubOrgAdded(opts *bind.FilterOpts) (*ClusterSubOrgAddedIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "SubOrgAdded")
	if err != nil {
		return nil, err
	}
	return &ClusterSubOrgAddedIterator{contract: _Cluster.contract, event: "SubOrgAdded", logs: logs, sub: sub}, nil
}

// WatchSubOrgAdded is a free log subscription operation binding the contract event 0xd734c07873f32f0735016e51dc718e21a48a3bec999d5be38cf3af363fbfedab.
//
// Solidity: e SubOrgAdded(_orgId string)
func (_Cluster *ClusterFilterer) WatchSubOrgAdded(opts *bind.WatchOpts, sink chan<- *ClusterSubOrgAdded) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "SubOrgAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterSubOrgAdded)
				if err := _Cluster.contract.UnpackLog(event, "SubOrgAdded", log); err != nil {
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

// ClusterVoterAddedIterator is returned from FilterVoterAdded and is used to iterate over the raw logs and unpacked data for VoterAdded events raised by the Cluster contract.
type ClusterVoterAddedIterator struct {
	Event *ClusterVoterAdded // Event containing the contract specifics and raw log

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
func (it *ClusterVoterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterVoterAdded)
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
		it.Event = new(ClusterVoterAdded)
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
func (it *ClusterVoterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterVoterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterVoterAdded represents a VoterAdded event raised by the Cluster contract.
type ClusterVoterAdded struct {
	OrgId   string
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVoterAdded is a free log retrieval operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: e VoterAdded(_orgId string, _address address)
func (_Cluster *ClusterFilterer) FilterVoterAdded(opts *bind.FilterOpts) (*ClusterVoterAddedIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "VoterAdded")
	if err != nil {
		return nil, err
	}
	return &ClusterVoterAddedIterator{contract: _Cluster.contract, event: "VoterAdded", logs: logs, sub: sub}, nil
}

// WatchVoterAdded is a free log subscription operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: e VoterAdded(_orgId string, _address address)
func (_Cluster *ClusterFilterer) WatchVoterAdded(opts *bind.WatchOpts, sink chan<- *ClusterVoterAdded) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "VoterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterVoterAdded)
				if err := _Cluster.contract.UnpackLog(event, "VoterAdded", log); err != nil {
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

// ClusterVoterDeletedIterator is returned from FilterVoterDeleted and is used to iterate over the raw logs and unpacked data for VoterDeleted events raised by the Cluster contract.
type ClusterVoterDeletedIterator struct {
	Event *ClusterVoterDeleted // Event containing the contract specifics and raw log

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
func (it *ClusterVoterDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterVoterDeleted)
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
		it.Event = new(ClusterVoterDeleted)
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
func (it *ClusterVoterDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterVoterDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterVoterDeleted represents a VoterDeleted event raised by the Cluster contract.
type ClusterVoterDeleted struct {
	OrgId   string
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVoterDeleted is a free log retrieval operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: e VoterDeleted(_orgId string, _address address)
func (_Cluster *ClusterFilterer) FilterVoterDeleted(opts *bind.FilterOpts) (*ClusterVoterDeletedIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "VoterDeleted")
	if err != nil {
		return nil, err
	}
	return &ClusterVoterDeletedIterator{contract: _Cluster.contract, event: "VoterDeleted", logs: logs, sub: sub}, nil
}

// WatchVoterDeleted is a free log subscription operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: e VoterDeleted(_orgId string, _address address)
func (_Cluster *ClusterFilterer) WatchVoterDeleted(opts *bind.WatchOpts, sink chan<- *ClusterVoterDeleted) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "VoterDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterVoterDeleted)
				if err := _Cluster.contract.UnpackLog(event, "VoterDeleted", log); err != nil {
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
