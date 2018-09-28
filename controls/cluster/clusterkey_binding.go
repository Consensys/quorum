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
const ClusterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_privateKeys\",\"type\":\"string\"}],\"name\":\"updatedOrgKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_privateKeys\",\"type\":\"string\"}],\"name\":\"OrgKeyUpdated\",\"type\":\"event\"}]"

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

// UpdatedOrgKeys is a paid mutator transaction binding the contract method 0x57b006b1.
//
// Solidity: function updatedOrgKeys(_orgId string, _privateKeys string) returns()
func (_Cluster *ClusterTransactor) UpdatedOrgKeys(opts *bind.TransactOpts, _orgId string, _privateKeys string) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "updatedOrgKeys", _orgId, _privateKeys)
}

// UpdatedOrgKeys is a paid mutator transaction binding the contract method 0x57b006b1.
//
// Solidity: function updatedOrgKeys(_orgId string, _privateKeys string) returns()
func (_Cluster *ClusterSession) UpdatedOrgKeys(_orgId string, _privateKeys string) (*types.Transaction, error) {
	return _Cluster.Contract.UpdatedOrgKeys(&_Cluster.TransactOpts, _orgId, _privateKeys)
}

// UpdatedOrgKeys is a paid mutator transaction binding the contract method 0x57b006b1.
//
// Solidity: function updatedOrgKeys(_orgId string, _privateKeys string) returns()
func (_Cluster *ClusterTransactorSession) UpdatedOrgKeys(_orgId string, _privateKeys string) (*types.Transaction, error) {
	return _Cluster.Contract.UpdatedOrgKeys(&_Cluster.TransactOpts, _orgId, _privateKeys)
}

// ClusterOrgKeyUpdatedIterator is returned from FilterOrgKeyUpdated and is used to iterate over the raw logs and unpacked data for OrgKeyUpdated events raised by the Cluster contract.
type ClusterOrgKeyUpdatedIterator struct {
	Event *ClusterOrgKeyUpdated // Event containing the contract specifics and raw log

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
func (it *ClusterOrgKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClusterOrgKeyUpdated)
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
		it.Event = new(ClusterOrgKeyUpdated)
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
func (it *ClusterOrgKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClusterOrgKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClusterOrgKeyUpdated represents a OrgKeyUpdated event raised by the Cluster contract.
type ClusterOrgKeyUpdated struct {
	OrgId       string
	PrivateKeys string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrgKeyUpdated is a free log retrieval operation binding the contract event 0x824be03dd319521ddd2d26c548a12bf867ac723abc5f3c1d9f5a1eb8b5420bde.
//
// Solidity: e OrgKeyUpdated(_orgId string, _privateKeys string)
func (_Cluster *ClusterFilterer) FilterOrgKeyUpdated(opts *bind.FilterOpts) (*ClusterOrgKeyUpdatedIterator, error) {

	logs, sub, err := _Cluster.contract.FilterLogs(opts, "OrgKeyUpdated")
	if err != nil {
		return nil, err
	}
	return &ClusterOrgKeyUpdatedIterator{contract: _Cluster.contract, event: "OrgKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchOrgKeyUpdated is a free log subscription operation binding the contract event 0x824be03dd319521ddd2d26c548a12bf867ac723abc5f3c1d9f5a1eb8b5420bde.
//
// Solidity: e OrgKeyUpdated(_orgId string, _privateKeys string)
func (_Cluster *ClusterFilterer) WatchOrgKeyUpdated(opts *bind.WatchOpts, sink chan<- *ClusterOrgKeyUpdated) (event.Subscription, error) {

	logs, sub, err := _Cluster.contract.WatchLogs(opts, "OrgKeyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClusterOrgKeyUpdated)
				if err := _Cluster.contract.UnpackLog(event, "OrgKeyUpdated", log); err != nil {
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
