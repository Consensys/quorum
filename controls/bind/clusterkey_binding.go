// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind

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
const ClusterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_morgId\",\"type\":\"string\"}],\"name\":\"addSubOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"approvePendingOp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"deleteOrgKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_morgId\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_morgId\",\"type\":\"string\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"deleteVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_morgId\",\"type\":\"string\"}],\"name\":\"addMasterOrg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"addOrgKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_morgId\",\"type\":\"string\"}],\"name\":\"checkMasterOrgExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"MasterOrgAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"MasterOrgExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"MasterOrgNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"SubOrgAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"SubOrgExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"SubOrgNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"OrgKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"OrgKeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"KeyNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"KeyExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"OrgNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"PendingApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_pendingOp\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"ItemForApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NothingToApprove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"NoVotingAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"VoterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"VoterNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"VoterDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"VoterExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_tmKey\",\"type\":\"string\"}],\"name\":\"PrintAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_voterAccount\",\"type\":\"address\"}],\"name\":\"PrintVoter\",\"type\":\"event\"}]"

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
	OrgId string
	TmKey string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterKeyExists is a free log retrieval operation binding the contract event 0xfd2bb3c1cfc78c051cb1f0ed88147fb9348eba128a594dc66fbf35dc63fe692d.
//
// Solidity: e KeyExists(_orgId string, _tmKey string)
func (_Cluster *ClusterFilterer) FilterKeyExists(opts *bind.FilterOpts) (*ClusterKeyExistsIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "KeyExists")
	if err != nil {
		return nil, err
	}
	return &ClusterKeyExistsIterator{contract: _Cluster.contract, event: "KeyExists", logs: logs, sub: sub}, nil
}

// WatchKeyExists is a free log subscription operation binding the contract event 0xfd2bb3c1cfc78c051cb1f0ed88147fb9348eba128a594dc66fbf35dc63fe692d.
//
// Solidity: e KeyExists(_orgId string, _tmKey string)
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
	TmKey string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterKeyNotFound is a free log retrieval operation binding the contract event 0x1625cf45f71f82c8ccf66926c15856f85b1e08dbe285065512100db776fdeb28.
//
// Solidity: e KeyNotFound(_tmKey string)
func (_Cluster *ClusterFilterer) FilterKeyNotFound(opts *bind.FilterOpts) (*ClusterKeyNotFoundIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "KeyNotFound")
	if err != nil {
		return nil, err
	}
	return &ClusterKeyNotFoundIterator{contract: _Cluster.contract, event: "KeyNotFound", logs: logs, sub: sub}, nil
}

// WatchKeyNotFound is a free log subscription operation binding the contract event 0x1625cf45f71f82c8ccf66926c15856f85b1e08dbe285065512100db776fdeb28.
//
// Solidity: e KeyNotFound(_tmKey string)
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

// ClusterMasterOrgExistsIterator is returned from FilterMasterOrgExists and is used to iterate over the raw logs and unpacked data for MasterOrgExists events raised by the Cluster contract.
type ClusterMasterOrgExistsIterator struct {
	Event *ClusterMasterOrgExists // Event containing the contract specifics and raw log

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
func (it *ClusterMasterOrgExistsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMasterOrgExists)
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
		it.Event = new(ClusterMasterOrgExists)
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
func (it *ClusterMasterOrgExistsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMasterOrgExistsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMasterOrgExists represents a MasterOrgExists event raised by the Cluster contract.
type ClusterMasterOrgExists struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMasterOrgExists is a free log retrieval operation binding the contract event 0x5c3bfabea6adb09ab9fcc026934927d15573bb92c42edcf1a4052b2342089f5e.
//
// Solidity: e MasterOrgExists(_orgId string)
func (_Cluster *ClusterFilterer) FilterMasterOrgExists(opts *bind.FilterOpts) (*ClusterMasterOrgExistsIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "MasterOrgExists")
	if err != nil {
		return nil, err
	}
	return &ClusterMasterOrgExistsIterator{contract: _Cluster.contract, event: "MasterOrgExists", logs: logs, sub: sub}, nil
}

// WatchMasterOrgExists is a free log subscription operation binding the contract event 0x5c3bfabea6adb09ab9fcc026934927d15573bb92c42edcf1a4052b2342089f5e.
//
// Solidity: e MasterOrgExists(_orgId string)
func (_Cluster *ClusterFilterer) WatchMasterOrgExists(opts *bind.WatchOpts, sink chan<- *ClusterMasterOrgExists) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "MasterOrgExists")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMasterOrgExists)
				if err := _Cluster.contract.UnpackLog(event, "MasterOrgExists", log); err != nil {
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

// ClusterMasterOrgNotFoundIterator is returned from FilterMasterOrgNotFound and is used to iterate over the raw logs and unpacked data for MasterOrgNotFound events raised by the Cluster contract.
type ClusterMasterOrgNotFoundIterator struct {
	Event *ClusterMasterOrgNotFound // Event containing the contract specifics and raw log

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
func (it *ClusterMasterOrgNotFoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterMasterOrgNotFound)
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
		it.Event = new(ClusterMasterOrgNotFound)
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
func (it *ClusterMasterOrgNotFoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterMasterOrgNotFoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterMasterOrgNotFound represents a MasterOrgNotFound event raised by the Cluster contract.
type ClusterMasterOrgNotFound struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMasterOrgNotFound is a free log retrieval operation binding the contract event 0xa4f69f0a4296104a861bac36ea23551d6d71a4b2d9f788ea28468eef956b4e57.
//
// Solidity: e MasterOrgNotFound(_orgId string)
func (_Cluster *ClusterFilterer) FilterMasterOrgNotFound(opts *bind.FilterOpts) (*ClusterMasterOrgNotFoundIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "MasterOrgNotFound")
	if err != nil {
		return nil, err
	}
	return &ClusterMasterOrgNotFoundIterator{contract: _Cluster.contract, event: "MasterOrgNotFound", logs: logs, sub: sub}, nil
}

// WatchMasterOrgNotFound is a free log subscription operation binding the contract event 0xa4f69f0a4296104a861bac36ea23551d6d71a4b2d9f788ea28468eef956b4e57.
//
// Solidity: e MasterOrgNotFound(_orgId string)
func (_Cluster *ClusterFilterer) WatchMasterOrgNotFound(opts *bind.WatchOpts, sink chan<- *ClusterMasterOrgNotFound) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "MasterOrgNotFound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterMasterOrgNotFound)
				if err := _Cluster.contract.UnpackLog(event, "MasterOrgNotFound", log); err != nil {
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
	OrgId string
	TmKey string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPrintAll is a free log retrieval operation binding the contract event 0x3d030f7cce2619e90f621cb560eb4327f74d9a412c2daa8bed5a892d759187ec.
//
// Solidity: e PrintAll(_orgId string, _tmKey string)
func (_Cluster *ClusterFilterer) FilterPrintAll(opts *bind.FilterOpts) (*ClusterPrintAllIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "PrintAll")
	if err != nil {
		return nil, err
	}
	return &ClusterPrintAllIterator{contract: _Cluster.contract, event: "PrintAll", logs: logs, sub: sub}, nil
}

// WatchPrintAll is a free log subscription operation binding the contract event 0x3d030f7cce2619e90f621cb560eb4327f74d9a412c2daa8bed5a892d759187ec.
//
// Solidity: e PrintAll(_orgId string, _tmKey string)
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

// ClusterSubOrgExistsIterator is returned from FilterSubOrgExists and is used to iterate over the raw logs and unpacked data for SubOrgExists events raised by the Cluster contract.
type ClusterSubOrgExistsIterator struct {
	Event *ClusterSubOrgExists // Event containing the contract specifics and raw log

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
func (it *ClusterSubOrgExistsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterSubOrgExists)
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
		it.Event = new(ClusterSubOrgExists)
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
func (it *ClusterSubOrgExistsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterSubOrgExistsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterSubOrgExists represents a SubOrgExists event raised by the Cluster contract.
type ClusterSubOrgExists struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubOrgExists is a free log retrieval operation binding the contract event 0x2b24431229489a8557abb42bd24fcf95defc95e5c1277f2b08c4860e7f62ee35.
//
// Solidity: e SubOrgExists(_orgId string)
func (_Cluster *ClusterFilterer) FilterSubOrgExists(opts *bind.FilterOpts) (*ClusterSubOrgExistsIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "SubOrgExists")
	if err != nil {
		return nil, err
	}
	return &ClusterSubOrgExistsIterator{contract: _Cluster.contract, event: "SubOrgExists", logs: logs, sub: sub}, nil
}

// WatchSubOrgExists is a free log subscription operation binding the contract event 0x2b24431229489a8557abb42bd24fcf95defc95e5c1277f2b08c4860e7f62ee35.
//
// Solidity: e SubOrgExists(_orgId string)
func (_Cluster *ClusterFilterer) WatchSubOrgExists(opts *bind.WatchOpts, sink chan<- *ClusterSubOrgExists) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "SubOrgExists")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterSubOrgExists)
				if err := _Cluster.contract.UnpackLog(event, "SubOrgExists", log); err != nil {
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

// ClusterSubOrgNotFoundIterator is returned from FilterSubOrgNotFound and is used to iterate over the raw logs and unpacked data for SubOrgNotFound events raised by the Cluster contract.
type ClusterSubOrgNotFoundIterator struct {
	Event *ClusterSubOrgNotFound // Event containing the contract specifics and raw log

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
func (it *ClusterSubOrgNotFoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterSubOrgNotFound)
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
		it.Event = new(ClusterSubOrgNotFound)
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
func (it *ClusterSubOrgNotFoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterSubOrgNotFoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterSubOrgNotFound represents a SubOrgNotFound event raised by the Cluster contract.
type ClusterSubOrgNotFound struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubOrgNotFound is a free log retrieval operation binding the contract event 0xbdc2f111e14dcb54d9904c6194c792ff3d70aec9d1c642165aa04e89a28536cd.
//
// Solidity: e SubOrgNotFound(_orgId string)
func (_Cluster *ClusterFilterer) FilterSubOrgNotFound(opts *bind.FilterOpts) (*ClusterSubOrgNotFoundIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "SubOrgNotFound")
	if err != nil {
		return nil, err
	}
	return &ClusterSubOrgNotFoundIterator{contract: _Cluster.contract, event: "SubOrgNotFound", logs: logs, sub: sub}, nil
}

// WatchSubOrgNotFound is a free log subscription operation binding the contract event 0xbdc2f111e14dcb54d9904c6194c792ff3d70aec9d1c642165aa04e89a28536cd.
//
// Solidity: e SubOrgNotFound(_orgId string)
func (_Cluster *ClusterFilterer) WatchSubOrgNotFound(opts *bind.WatchOpts, sink chan<- *ClusterSubOrgNotFound) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "SubOrgNotFound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterSubOrgNotFound)
				if err := _Cluster.contract.UnpackLog(event, "SubOrgNotFound", log); err != nil {
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
