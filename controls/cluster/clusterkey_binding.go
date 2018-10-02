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
const ClusterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"deleteOrgKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"printAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"addOrgKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"OrgKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"OrgKeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_voterAccount\",\"type\":\"string\"}],\"name\":\"orgVoterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"KeyNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"OrgNotFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"PrintAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_privateKey\",\"type\":\"string\"}],\"name\":\"KeyExists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_keyExists\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"loopCnt\",\"type\":\"uint256\"}],\"name\":\"Dummy\",\"type\":\"event\"}]"

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

// PrintAll is a paid mutator transaction binding the contract method 0xb9f41ba3.
//
// Solidity: function printAll() returns()
func (_Cluster *ClusterTransactor) PrintAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "printAll")
}

// PrintAll is a paid mutator transaction binding the contract method 0xb9f41ba3.
//
// Solidity: function printAll() returns()
func (_Cluster *ClusterSession) PrintAll() (*types.Transaction, error) {
	return _Cluster.Contract.PrintAll(&_Cluster.TransactOpts)
}

// PrintAll is a paid mutator transaction binding the contract method 0xb9f41ba3.
//
// Solidity: function printAll() returns()
func (_Cluster *ClusterTransactorSession) PrintAll() (*types.Transaction, error) {
	return _Cluster.Contract.PrintAll(&_Cluster.TransactOpts)
}

// ClusterDummyIterator is returned from FilterDummy and is used to iterate over the raw logs and unpacked data for Dummy events raised by the Cluster contract.
type ClusterDummyIterator struct {
	Event *ClusterDummy // Event containing the contract specifics and raw log

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
func (it *ClusterDummyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterDummy)
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
		it.Event = new(ClusterDummy)
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
func (it *ClusterDummyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterDummyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterDummy represents a Dummy event raised by the Cluster contract.
type ClusterDummy struct {
	OrgId     *big.Int
	KeyExists bool
	LoopCnt   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDummy is a free log retrieval operation binding the contract event 0xd58c0f112df16fab0019f11d0dd6b109778672dd444bdd4ef464ff369d83eacd.
//
// Solidity: e Dummy(_orgId uint256, _keyExists bool, loopCnt uint256)
func (_Cluster *ClusterFilterer) FilterDummy(opts *bind.FilterOpts) (*ClusterDummyIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "Dummy")
	if err != nil {
		return nil, err
	}
	return &ClusterDummyIterator{contract: _Cluster.contract, event: "Dummy", logs: logs, sub: sub}, nil
}

// WatchDummy is a free log subscription operation binding the contract event 0xd58c0f112df16fab0019f11d0dd6b109778672dd444bdd4ef464ff369d83eacd.
//
// Solidity: e Dummy(_orgId uint256, _keyExists bool, loopCnt uint256)
func (_Cluster *ClusterFilterer) WatchDummy(opts *bind.WatchOpts, sink chan<- *ClusterDummy) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "Dummy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterDummy)
				if err := _Cluster.contract.UnpackLog(event, "Dummy", log); err != nil {
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
