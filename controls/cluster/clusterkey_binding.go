// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cluster

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ClusterABI is the input ABI used to generate the binding from.
const ClusterABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"printAllOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"approvePendingOp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"deleteOrgKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"deleteVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"printAllVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"addOrgKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"OrgKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"OrgKeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_voterAccount\",\"type\":\"string\"}],\"name\":\"orgVoterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"KeyNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"OrgNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"KeyExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"VoterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"VoterExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"VoterNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"VoterAccountDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NoVotingAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"PendingApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_pendingOp\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"ItemForApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NothingToApprove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"PrintAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_voterAccount\",\"type\":\"address\"}],\"name\":\"PrintVoter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_pendingOp\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_pendingKey\",\"type\":\"string\"}],\"name\":\"PrintKey\",\"type\":\"event\"}]"

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

// AddOrgKey is a paid mutator transaction binding the contract method 0xd88ce6bb.
//
// Solidity: function addOrgKey(_orgId string, _privateKey string) returns()
func (_Cluster *ClusterTransactor) AddOrgKey(opts *bind.TransactOpts, _orgId string, _privateKey string) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "addOrgKey", _orgId, _privateKey)
}

// AddOrgKey is a paid mutator transaction binding the contract method 0xd88ce6bb.
//
// Solidity: function addOrgKey(_orgId string, _privateKey string) returns()
func (_Cluster *ClusterSession) AddOrgKey(_orgId string, _privateKey string) (*types.Transaction, error) {
	return _Cluster.Contract.AddOrgKey(&_Cluster.TransactOpts, _orgId, _privateKey)
}

// AddOrgKey is a paid mutator transaction binding the contract method 0xd88ce6bb.
//
// Solidity: function addOrgKey(_orgId string, _privateKey string) returns()
func (_Cluster *ClusterTransactorSession) AddOrgKey(_orgId string, _privateKey string) (*types.Transaction, error) {
	return _Cluster.Contract.AddOrgKey(&_Cluster.TransactOpts, _orgId, _privateKey)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(_orgId string, _address address) returns()
func (_Cluster *ClusterTransactor) AddVoter(opts *bind.TransactOpts, _orgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "addVoter", _orgId, _address)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(_orgId string, _address address) returns()
func (_Cluster *ClusterSession) AddVoter(_orgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.Contract.AddVoter(&_Cluster.TransactOpts, _orgId, _address)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(_orgId string, _address address) returns()
func (_Cluster *ClusterTransactorSession) AddVoter(_orgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.Contract.AddVoter(&_Cluster.TransactOpts, _orgId, _address)
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
// Solidity: function deleteOrgKey(_orgId string, _privateKey string) returns()
func (_Cluster *ClusterTransactor) DeleteOrgKey(opts *bind.TransactOpts, _orgId string, _privateKey string) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "deleteOrgKey", _orgId, _privateKey)
}

// DeleteOrgKey is a paid mutator transaction binding the contract method 0x49379c50.
//
// Solidity: function deleteOrgKey(_orgId string, _privateKey string) returns()
func (_Cluster *ClusterSession) DeleteOrgKey(_orgId string, _privateKey string) (*types.Transaction, error) {
	return _Cluster.Contract.DeleteOrgKey(&_Cluster.TransactOpts, _orgId, _privateKey)
}

// DeleteOrgKey is a paid mutator transaction binding the contract method 0x49379c50.
//
// Solidity: function deleteOrgKey(_orgId string, _privateKey string) returns()
func (_Cluster *ClusterTransactorSession) DeleteOrgKey(_orgId string, _privateKey string) (*types.Transaction, error) {
	return _Cluster.Contract.DeleteOrgKey(&_Cluster.TransactOpts, _orgId, _privateKey)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(_orgId string, _address address) returns()
func (_Cluster *ClusterTransactor) DeleteVoter(opts *bind.TransactOpts, _orgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "deleteVoter", _orgId, _address)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(_orgId string, _address address) returns()
func (_Cluster *ClusterSession) DeleteVoter(_orgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.Contract.DeleteVoter(&_Cluster.TransactOpts, _orgId, _address)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(_orgId string, _address address) returns()
func (_Cluster *ClusterTransactorSession) DeleteVoter(_orgId string, _address common.Address) (*types.Transaction, error) {
	return _Cluster.Contract.DeleteVoter(&_Cluster.TransactOpts, _orgId, _address)
}

// PrintAllOrg is a paid mutator transaction binding the contract method 0x2bbc5084.
//
// Solidity: function printAllOrg() returns()
func (_Cluster *ClusterTransactor) PrintAllOrg(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "printAllOrg")
}

// PrintAllOrg is a paid mutator transaction binding the contract method 0x2bbc5084.
//
// Solidity: function printAllOrg() returns()
func (_Cluster *ClusterSession) PrintAllOrg() (*types.Transaction, error) {
	return _Cluster.Contract.PrintAllOrg(&_Cluster.TransactOpts)
}

// PrintAllOrg is a paid mutator transaction binding the contract method 0x2bbc5084.
//
// Solidity: function printAllOrg() returns()
func (_Cluster *ClusterTransactorSession) PrintAllOrg() (*types.Transaction, error) {
	return _Cluster.Contract.PrintAllOrg(&_Cluster.TransactOpts)
}

// PrintAllVoter is a paid mutator transaction binding the contract method 0x73f9cee0.
//
// Solidity: function printAllVoter() returns()
func (_Cluster *ClusterTransactor) PrintAllVoter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "printAllVoter")
}

// PrintAllVoter is a paid mutator transaction binding the contract method 0x73f9cee0.
//
// Solidity: function printAllVoter() returns()
func (_Cluster *ClusterSession) PrintAllVoter() (*types.Transaction, error) {
	return _Cluster.Contract.PrintAllVoter(&_Cluster.TransactOpts)
}

// PrintAllVoter is a paid mutator transaction binding the contract method 0x73f9cee0.
//
// Solidity: function printAllVoter() returns()
func (_Cluster *ClusterTransactorSession) PrintAllVoter() (*types.Transaction, error) {
	return _Cluster.Contract.PrintAllVoter(&_Cluster.TransactOpts)
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
	OrgId      string
	PendingOp  uint8
	PrivateKey string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterItemForApproval is a free log retrieval operation binding the contract event 0x4475befcee492797e02530076fd7e138aa058eb3bcd028a9df5c0f2815ba9f4a.
//
// Solidity: e ItemForApproval(_orgId string, _pendingOp uint8, _privateKey string)
func (_Cluster *ClusterFilterer) FilterItemForApproval(opts *bind.FilterOpts) (*ClusterItemForApprovalIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "ItemForApproval")
	if err != nil {
		return nil, err
	}
	return &ClusterItemForApprovalIterator{contract: _Cluster.contract, event: "ItemForApproval", logs: logs, sub: sub}, nil
}

// WatchItemForApproval is a free log subscription operation binding the contract event 0x4475befcee492797e02530076fd7e138aa058eb3bcd028a9df5c0f2815ba9f4a.
//
// Solidity: e ItemForApproval(_orgId string, _pendingOp uint8, _privateKey string)
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

// ClusterKeyExistsIterator is returned from FilterKeyExists and is used to iterate over the raw logs and unpacked data for KeyExists events raised by the Cluster contract.
type ClusterKeyExistsIterator struct {
	Event *ClusterKeyExists // Event containing the contract specifics and raw log

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
func (it *ClusterKeyExistsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterKeyExists)
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
		it.Event = new(ClusterKeyExists)
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
func (it *ClusterKeyExistsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterKeyExistsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterKeyExists represents a KeyExists event raised by the Cluster contract.
type ClusterKeyExists struct {
	OrgId      string
	PrivateKey string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterKeyExists is a free log retrieval operation binding the contract event 0xfd2bb3c1cfc78c051cb1f0ed88147fb9348eba128a594dc66fbf35dc63fe692d.
//
// Solidity: e KeyExists(_orgId string, _privateKey string)
func (_Cluster *ClusterFilterer) FilterKeyExists(opts *bind.FilterOpts) (*ClusterKeyExistsIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "KeyExists")
	if err != nil {
		return nil, err
	}
	return &ClusterKeyExistsIterator{contract: _Cluster.contract, event: "KeyExists", logs: logs, sub: sub}, nil
}

// WatchKeyExists is a free log subscription operation binding the contract event 0xfd2bb3c1cfc78c051cb1f0ed88147fb9348eba128a594dc66fbf35dc63fe692d.
//
// Solidity: e KeyExists(_orgId string, _privateKey string)
func (_Cluster *ClusterFilterer) WatchKeyExists(opts *bind.WatchOpts, sink chan<- *ClusterKeyExists) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "KeyExists")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterKeyExists)
				if err := _Cluster.contract.UnpackLog(event, "KeyExists", log); err != nil {
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

// ClusterKeyNotFoundIterator is returned from FilterKeyNotFound and is used to iterate over the raw logs and unpacked data for KeyNotFound events raised by the Cluster contract.
type ClusterKeyNotFoundIterator struct {
	Event *ClusterKeyNotFound // Event containing the contract specifics and raw log

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
func (it *ClusterKeyNotFoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterKeyNotFound)
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
		it.Event = new(ClusterKeyNotFound)
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
func (it *ClusterKeyNotFoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterKeyNotFoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterKeyNotFound represents a KeyNotFound event raised by the Cluster contract.
type ClusterKeyNotFound struct {
	PrivateKey string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterKeyNotFound is a free log retrieval operation binding the contract event 0x1625cf45f71f82c8ccf66926c15856f85b1e08dbe285065512100db776fdeb28.
//
// Solidity: e KeyNotFound(_privateKey string)
func (_Cluster *ClusterFilterer) FilterKeyNotFound(opts *bind.FilterOpts) (*ClusterKeyNotFoundIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "KeyNotFound")
	if err != nil {
		return nil, err
	}
	return &ClusterKeyNotFoundIterator{contract: _Cluster.contract, event: "KeyNotFound", logs: logs, sub: sub}, nil
}

// WatchKeyNotFound is a free log subscription operation binding the contract event 0x1625cf45f71f82c8ccf66926c15856f85b1e08dbe285065512100db776fdeb28.
//
// Solidity: e KeyNotFound(_privateKey string)
func (_Cluster *ClusterFilterer) WatchKeyNotFound(opts *bind.WatchOpts, sink chan<- *ClusterKeyNotFound) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "KeyNotFound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterKeyNotFound)
				if err := _Cluster.contract.UnpackLog(event, "KeyNotFound", log); err != nil {
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

// ClusterNoVotingAccountIterator is returned from FilterNoVotingAccount and is used to iterate over the raw logs and unpacked data for NoVotingAccount events raised by the Cluster contract.
type ClusterNoVotingAccountIterator struct {
	Event *ClusterNoVotingAccount // Event containing the contract specifics and raw log

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
func (it *ClusterNoVotingAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterNoVotingAccount)
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
		it.Event = new(ClusterNoVotingAccount)
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
func (it *ClusterNoVotingAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterNoVotingAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterNoVotingAccount represents a NoVotingAccount event raised by the Cluster contract.
type ClusterNoVotingAccount struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNoVotingAccount is a free log retrieval operation binding the contract event 0xaca1d1ef8876e1135c8c76871025becb2eefcbdb13c62fcd55c51dc174abf7af.
//
// Solidity: e NoVotingAccount(_orgId string)
func (_Cluster *ClusterFilterer) FilterNoVotingAccount(opts *bind.FilterOpts) (*ClusterNoVotingAccountIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "NoVotingAccount")
	if err != nil {
		return nil, err
	}
	return &ClusterNoVotingAccountIterator{contract: _Cluster.contract, event: "NoVotingAccount", logs: logs, sub: sub}, nil
}

// WatchNoVotingAccount is a free log subscription operation binding the contract event 0xaca1d1ef8876e1135c8c76871025becb2eefcbdb13c62fcd55c51dc174abf7af.
//
// Solidity: e NoVotingAccount(_orgId string)
func (_Cluster *ClusterFilterer) WatchNoVotingAccount(opts *bind.WatchOpts, sink chan<- *ClusterNoVotingAccount) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "NoVotingAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterNoVotingAccount)
				if err := _Cluster.contract.UnpackLog(event, "NoVotingAccount", log); err != nil {
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

// ClusterNothingToApproveIterator is returned from FilterNothingToApprove and is used to iterate over the raw logs and unpacked data for NothingToApprove events raised by the Cluster contract.
type ClusterNothingToApproveIterator struct {
	Event *ClusterNothingToApprove // Event containing the contract specifics and raw log

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
func (it *ClusterNothingToApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterNothingToApprove)
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
		it.Event = new(ClusterNothingToApprove)
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
func (it *ClusterNothingToApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterNothingToApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterNothingToApprove represents a NothingToApprove event raised by the Cluster contract.
type ClusterNothingToApprove struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNothingToApprove is a free log retrieval operation binding the contract event 0xe820171ad1d64f6ca44bb0943acbac4c6085812bf91e8b791646295639806228.
//
// Solidity: e NothingToApprove(_orgId string)
func (_Cluster *ClusterFilterer) FilterNothingToApprove(opts *bind.FilterOpts) (*ClusterNothingToApproveIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "NothingToApprove")
	if err != nil {
		return nil, err
	}
	return &ClusterNothingToApproveIterator{contract: _Cluster.contract, event: "NothingToApprove", logs: logs, sub: sub}, nil
}

// WatchNothingToApprove is a free log subscription operation binding the contract event 0xe820171ad1d64f6ca44bb0943acbac4c6085812bf91e8b791646295639806228.
//
// Solidity: e NothingToApprove(_orgId string)
func (_Cluster *ClusterFilterer) WatchNothingToApprove(opts *bind.WatchOpts, sink chan<- *ClusterNothingToApprove) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "NothingToApprove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterNothingToApprove)
				if err := _Cluster.contract.UnpackLog(event, "NothingToApprove", log); err != nil {
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
	OrgId      string
	PrivateKey string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrgKeyAdded is a free log retrieval operation binding the contract event 0x6f4d370520782587dabc737a258f46de81ad45d733a42cd5a0045cff1e46deb4.
//
// Solidity: e OrgKeyAdded(_orgId string, _privateKey string)
func (_Cluster *ClusterFilterer) FilterOrgKeyAdded(opts *bind.FilterOpts) (*ClusterOrgKeyAddedIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "OrgKeyAdded")
	if err != nil {
		return nil, err
	}
	return &ClusterOrgKeyAddedIterator{contract: _Cluster.contract, event: "OrgKeyAdded", logs: logs, sub: sub}, nil
}

// WatchOrgKeyAdded is a free log subscription operation binding the contract event 0x6f4d370520782587dabc737a258f46de81ad45d733a42cd5a0045cff1e46deb4.
//
// Solidity: e OrgKeyAdded(_orgId string, _privateKey string)
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
	OrgId      string
	PrivateKey string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrgKeyDeleted is a free log retrieval operation binding the contract event 0x2e0a2dc845dce9ef7206b8fe38f3dacaad17ba74d7be9fba469c9858ae16a5d6.
//
// Solidity: e OrgKeyDeleted(_orgId string, _privateKey string)
func (_Cluster *ClusterFilterer) FilterOrgKeyDeleted(opts *bind.FilterOpts) (*ClusterOrgKeyDeletedIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "OrgKeyDeleted")
	if err != nil {
		return nil, err
	}
	return &ClusterOrgKeyDeletedIterator{contract: _Cluster.contract, event: "OrgKeyDeleted", logs: logs, sub: sub}, nil
}

// WatchOrgKeyDeleted is a free log subscription operation binding the contract event 0x2e0a2dc845dce9ef7206b8fe38f3dacaad17ba74d7be9fba469c9858ae16a5d6.
//
// Solidity: e OrgKeyDeleted(_orgId string, _privateKey string)
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

// ClusterOrgNotFoundIterator is returned from FilterOrgNotFound and is used to iterate over the raw logs and unpacked data for OrgNotFound events raised by the Cluster contract.
type ClusterOrgNotFoundIterator struct {
	Event *ClusterOrgNotFound // Event containing the contract specifics and raw log

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
func (it *ClusterOrgNotFoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterOrgNotFound)
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
		it.Event = new(ClusterOrgNotFound)
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
func (it *ClusterOrgNotFoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterOrgNotFoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterOrgNotFound represents a OrgNotFound event raised by the Cluster contract.
type ClusterOrgNotFound struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrgNotFound is a free log retrieval operation binding the contract event 0x0d426160118ead0b6900081fd1f08b0d9b626bd033ddd50cd7d24be253e11a83.
//
// Solidity: e OrgNotFound(_orgId string)
func (_Cluster *ClusterFilterer) FilterOrgNotFound(opts *bind.FilterOpts) (*ClusterOrgNotFoundIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "OrgNotFound")
	if err != nil {
		return nil, err
	}
	return &ClusterOrgNotFoundIterator{contract: _Cluster.contract, event: "OrgNotFound", logs: logs, sub: sub}, nil
}

// WatchOrgNotFound is a free log subscription operation binding the contract event 0x0d426160118ead0b6900081fd1f08b0d9b626bd033ddd50cd7d24be253e11a83.
//
// Solidity: e OrgNotFound(_orgId string)
func (_Cluster *ClusterFilterer) WatchOrgNotFound(opts *bind.WatchOpts, sink chan<- *ClusterOrgNotFound) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "OrgNotFound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterOrgNotFound)
				if err := _Cluster.contract.UnpackLog(event, "OrgNotFound", log); err != nil {
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

// ClusterPendingApprovalIterator is returned from FilterPendingApproval and is used to iterate over the raw logs and unpacked data for PendingApproval events raised by the Cluster contract.
type ClusterPendingApprovalIterator struct {
	Event *ClusterPendingApproval // Event containing the contract specifics and raw log

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
func (it *ClusterPendingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterPendingApproval)
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
		it.Event = new(ClusterPendingApproval)
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
func (it *ClusterPendingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterPendingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterPendingApproval represents a PendingApproval event raised by the Cluster contract.
type ClusterPendingApproval struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPendingApproval is a free log retrieval operation binding the contract event 0x2de31d28953221328a1c7e30a93fa15e0d8573128a8f6fa92cf66408a0403c99.
//
// Solidity: e PendingApproval(_orgId string)
func (_Cluster *ClusterFilterer) FilterPendingApproval(opts *bind.FilterOpts) (*ClusterPendingApprovalIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "PendingApproval")
	if err != nil {
		return nil, err
	}
	return &ClusterPendingApprovalIterator{contract: _Cluster.contract, event: "PendingApproval", logs: logs, sub: sub}, nil
}

// WatchPendingApproval is a free log subscription operation binding the contract event 0x2de31d28953221328a1c7e30a93fa15e0d8573128a8f6fa92cf66408a0403c99.
//
// Solidity: e PendingApproval(_orgId string)
func (_Cluster *ClusterFilterer) WatchPendingApproval(opts *bind.WatchOpts, sink chan<- *ClusterPendingApproval) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "PendingApproval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterPendingApproval)
				if err := _Cluster.contract.UnpackLog(event, "PendingApproval", log); err != nil {
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

// ClusterPrintAllIterator is returned from FilterPrintAll and is used to iterate over the raw logs and unpacked data for PrintAll events raised by the Cluster contract.
type ClusterPrintAllIterator struct {
	Event *ClusterPrintAll // Event containing the contract specifics and raw log

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
func (it *ClusterPrintAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterPrintAll)
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
		it.Event = new(ClusterPrintAll)
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
func (it *ClusterPrintAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterPrintAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterPrintAll represents a PrintAll event raised by the Cluster contract.
type ClusterPrintAll struct {
	OrgId      string
	PrivateKey string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPrintAll is a free log retrieval operation binding the contract event 0x3d030f7cce2619e90f621cb560eb4327f74d9a412c2daa8bed5a892d759187ec.
//
// Solidity: e PrintAll(_orgId string, _privateKey string)
func (_Cluster *ClusterFilterer) FilterPrintAll(opts *bind.FilterOpts) (*ClusterPrintAllIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "PrintAll")
	if err != nil {
		return nil, err
	}
	return &ClusterPrintAllIterator{contract: _Cluster.contract, event: "PrintAll", logs: logs, sub: sub}, nil
}

// WatchPrintAll is a free log subscription operation binding the contract event 0x3d030f7cce2619e90f621cb560eb4327f74d9a412c2daa8bed5a892d759187ec.
//
// Solidity: e PrintAll(_orgId string, _privateKey string)
func (_Cluster *ClusterFilterer) WatchPrintAll(opts *bind.WatchOpts, sink chan<- *ClusterPrintAll) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "PrintAll")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterPrintAll)
				if err := _Cluster.contract.UnpackLog(event, "PrintAll", log); err != nil {
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

// ClusterPrintKeyIterator is returned from FilterPrintKey and is used to iterate over the raw logs and unpacked data for PrintKey events raised by the Cluster contract.
type ClusterPrintKeyIterator struct {
	Event *ClusterPrintKey // Event containing the contract specifics and raw log

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
func (it *ClusterPrintKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterPrintKey)
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
		it.Event = new(ClusterPrintKey)
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
func (it *ClusterPrintKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterPrintKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterPrintKey represents a PrintKey event raised by the Cluster contract.
type ClusterPrintKey struct {
	OrgId      string
	PendingOp  uint8
	PendingKey string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPrintKey is a free log retrieval operation binding the contract event 0x5ac991854acd1b6419820adc1e7485528b0ad28b55d18e070e9bcd63786e7ff8.
//
// Solidity: e PrintKey(_orgId string, _pendingOp uint8, _pendingKey string)
func (_Cluster *ClusterFilterer) FilterPrintKey(opts *bind.FilterOpts) (*ClusterPrintKeyIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "PrintKey")
	if err != nil {
		return nil, err
	}
	return &ClusterPrintKeyIterator{contract: _Cluster.contract, event: "PrintKey", logs: logs, sub: sub}, nil
}

// WatchPrintKey is a free log subscription operation binding the contract event 0x5ac991854acd1b6419820adc1e7485528b0ad28b55d18e070e9bcd63786e7ff8.
//
// Solidity: e PrintKey(_orgId string, _pendingOp uint8, _pendingKey string)
func (_Cluster *ClusterFilterer) WatchPrintKey(opts *bind.WatchOpts, sink chan<- *ClusterPrintKey) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "PrintKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterPrintKey)
				if err := _Cluster.contract.UnpackLog(event, "PrintKey", log); err != nil {
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

// ClusterPrintVoterIterator is returned from FilterPrintVoter and is used to iterate over the raw logs and unpacked data for PrintVoter events raised by the Cluster contract.
type ClusterPrintVoterIterator struct {
	Event *ClusterPrintVoter // Event containing the contract specifics and raw log

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
func (it *ClusterPrintVoterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterPrintVoter)
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
		it.Event = new(ClusterPrintVoter)
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
func (it *ClusterPrintVoterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterPrintVoterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterPrintVoter represents a PrintVoter event raised by the Cluster contract.
type ClusterPrintVoter struct {
	OrgId        string
	VoterAccount common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPrintVoter is a free log retrieval operation binding the contract event 0x0c0001a7636c2b95d29de23e25bb65060f9ad324f9f38b309f6f5659a6cb3165.
//
// Solidity: e PrintVoter(_orgId string, _voterAccount address)
func (_Cluster *ClusterFilterer) FilterPrintVoter(opts *bind.FilterOpts) (*ClusterPrintVoterIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "PrintVoter")
	if err != nil {
		return nil, err
	}
	return &ClusterPrintVoterIterator{contract: _Cluster.contract, event: "PrintVoter", logs: logs, sub: sub}, nil
}

// WatchPrintVoter is a free log subscription operation binding the contract event 0x0c0001a7636c2b95d29de23e25bb65060f9ad324f9f38b309f6f5659a6cb3165.
//
// Solidity: e PrintVoter(_orgId string, _voterAccount address)
func (_Cluster *ClusterFilterer) WatchPrintVoter(opts *bind.WatchOpts, sink chan<- *ClusterPrintVoter) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "PrintVoter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterPrintVoter)
				if err := _Cluster.contract.UnpackLog(event, "PrintVoter", log); err != nil {
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

// ClusterVoterAccountDeletedIterator is returned from FilterVoterAccountDeleted and is used to iterate over the raw logs and unpacked data for VoterAccountDeleted events raised by the Cluster contract.
type ClusterVoterAccountDeletedIterator struct {
	Event *ClusterVoterAccountDeleted // Event containing the contract specifics and raw log

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
func (it *ClusterVoterAccountDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterVoterAccountDeleted)
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
		it.Event = new(ClusterVoterAccountDeleted)
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
func (it *ClusterVoterAccountDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterVoterAccountDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterVoterAccountDeleted represents a VoterAccountDeleted event raised by the Cluster contract.
type ClusterVoterAccountDeleted struct {
	OrgId   string
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVoterAccountDeleted is a free log retrieval operation binding the contract event 0x192eeefa7720f7067bcc0a3f5bbcc2941c417a8761ec8b795d2941f3b0d2be17.
//
// Solidity: e VoterAccountDeleted(_orgId string, _address address)
func (_Cluster *ClusterFilterer) FilterVoterAccountDeleted(opts *bind.FilterOpts) (*ClusterVoterAccountDeletedIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "VoterAccountDeleted")
	if err != nil {
		return nil, err
	}
	return &ClusterVoterAccountDeletedIterator{contract: _Cluster.contract, event: "VoterAccountDeleted", logs: logs, sub: sub}, nil
}

// WatchVoterAccountDeleted is a free log subscription operation binding the contract event 0x192eeefa7720f7067bcc0a3f5bbcc2941c417a8761ec8b795d2941f3b0d2be17.
//
// Solidity: e VoterAccountDeleted(_orgId string, _address address)
func (_Cluster *ClusterFilterer) WatchVoterAccountDeleted(opts *bind.WatchOpts, sink chan<- *ClusterVoterAccountDeleted) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "VoterAccountDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterVoterAccountDeleted)
				if err := _Cluster.contract.UnpackLog(event, "VoterAccountDeleted", log); err != nil {
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

// ClusterVoterExistsIterator is returned from FilterVoterExists and is used to iterate over the raw logs and unpacked data for VoterExists events raised by the Cluster contract.
type ClusterVoterExistsIterator struct {
	Event *ClusterVoterExists // Event containing the contract specifics and raw log

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
func (it *ClusterVoterExistsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterVoterExists)
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
		it.Event = new(ClusterVoterExists)
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
func (it *ClusterVoterExistsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterVoterExistsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterVoterExists represents a VoterExists event raised by the Cluster contract.
type ClusterVoterExists struct {
	OrgId   string
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVoterExists is a free log retrieval operation binding the contract event 0x57c0436fcca42a02516ed36a118ab7196a853b19ae03db4cbe9d1f6ec5a8f30b.
//
// Solidity: e VoterExists(_orgId string, _address address)
func (_Cluster *ClusterFilterer) FilterVoterExists(opts *bind.FilterOpts) (*ClusterVoterExistsIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "VoterExists")
	if err != nil {
		return nil, err
	}
	return &ClusterVoterExistsIterator{contract: _Cluster.contract, event: "VoterExists", logs: logs, sub: sub}, nil
}

// WatchVoterExists is a free log subscription operation binding the contract event 0x57c0436fcca42a02516ed36a118ab7196a853b19ae03db4cbe9d1f6ec5a8f30b.
//
// Solidity: e VoterExists(_orgId string, _address address)
func (_Cluster *ClusterFilterer) WatchVoterExists(opts *bind.WatchOpts, sink chan<- *ClusterVoterExists) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "VoterExists")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterVoterExists)
				if err := _Cluster.contract.UnpackLog(event, "VoterExists", log); err != nil {
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

// ClusterVoterNotFoundIterator is returned from FilterVoterNotFound and is used to iterate over the raw logs and unpacked data for VoterNotFound events raised by the Cluster contract.
type ClusterVoterNotFoundIterator struct {
	Event *ClusterVoterNotFound // Event containing the contract specifics and raw log

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
func (it *ClusterVoterNotFoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterVoterNotFound)
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
		it.Event = new(ClusterVoterNotFound)
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
func (it *ClusterVoterNotFoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterVoterNotFoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterVoterNotFound represents a VoterNotFound event raised by the Cluster contract.
type ClusterVoterNotFound struct {
	OrgId   string
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVoterNotFound is a free log retrieval operation binding the contract event 0xfc25f65937d1cb43a570fb3b3c5d0bf0b69b4cdd2da44c56b9893e2d14035960.
//
// Solidity: e VoterNotFound(_orgId string, _address address)
func (_Cluster *ClusterFilterer) FilterVoterNotFound(opts *bind.FilterOpts) (*ClusterVoterNotFoundIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "VoterNotFound")
	if err != nil {
		return nil, err
	}
	return &ClusterVoterNotFoundIterator{contract: _Cluster.contract, event: "VoterNotFound", logs: logs, sub: sub}, nil
}

// WatchVoterNotFound is a free log subscription operation binding the contract event 0xfc25f65937d1cb43a570fb3b3c5d0bf0b69b4cdd2da44c56b9893e2d14035960.
//
// Solidity: e VoterNotFound(_orgId string, _address address)
func (_Cluster *ClusterFilterer) WatchVoterNotFound(opts *bind.WatchOpts, sink chan<- *ClusterVoterNotFound) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "VoterNotFound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterVoterNotFound)
				if err := _Cluster.contract.UnpackLog(event, "VoterNotFound", log); err != nil {
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

// ClusterOrgVoterAddedIterator is returned from FilterOrgVoterAdded and is used to iterate over the raw logs and unpacked data for OrgVoterAdded events raised by the Cluster contract.
type ClusterOrgVoterAddedIterator struct {
	Event *ClusterOrgVoterAdded // Event containing the contract specifics and raw log

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
func (it *ClusterOrgVoterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterOrgVoterAdded)
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
		it.Event = new(ClusterOrgVoterAdded)
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
func (it *ClusterOrgVoterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterOrgVoterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterOrgVoterAdded represents a OrgVoterAdded event raised by the Cluster contract.
type ClusterOrgVoterAdded struct {
	OrgId        string
	VoterAccount string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrgVoterAdded is a free log retrieval operation binding the contract event 0x29f608001a67850240567b3e8b7e23bcef793f113a446763600384c00899c04c.
//
// Solidity: e orgVoterAdded(_orgId string, _voterAccount string)
func (_Cluster *ClusterFilterer) FilterOrgVoterAdded(opts *bind.FilterOpts) (*ClusterOrgVoterAddedIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "orgVoterAdded")
	if err != nil {
		return nil, err
	}
	return &ClusterOrgVoterAddedIterator{contract: _Cluster.contract, event: "orgVoterAdded", logs: logs, sub: sub}, nil
}

// WatchOrgVoterAdded is a free log subscription operation binding the contract event 0x29f608001a67850240567b3e8b7e23bcef793f113a446763600384c00899c04c.
//
// Solidity: e orgVoterAdded(_orgId string, _voterAccount string)
func (_Cluster *ClusterFilterer) WatchOrgVoterAdded(opts *bind.WatchOpts, sink chan<- *ClusterOrgVoterAdded) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "orgVoterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterOrgVoterAdded)
				if err := _Cluster.contract.UnpackLog(event, "orgVoterAdded", log); err != nil {
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
