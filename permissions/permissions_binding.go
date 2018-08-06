// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permissions

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PermissionsABI is the input ABI used to generate the binding from.
const PermissionsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"getNodeStatus\",\"outputs\":[{\"name\":\"_status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"ProposeDeactivation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_acctId\",\"type\":\"address\"},{\"name\":\"access\",\"type\":\"uint8\"}],\"name\":\"updateAcctAccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_canWrite\",\"type\":\"bool\"},{\"name\":\"_canLead\",\"type\":\"bool\"}],\"name\":\"ProposeNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"DeactivateNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"ApproveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_canWrite\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_canLead\",\"type\":\"bool\"}],\"name\":\"NewNodeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodeApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodePendingDeactivation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodeDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"acctId\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"access\",\"type\":\"uint8\"}],\"name\":\"AcctAccessModified\",\"type\":\"event\"}]"

// Permissions is an auto generated Go binding around an Ethereum contract.
type Permissions struct {
	PermissionsCaller     // Read-only binding to the contract
	PermissionsTransactor // Write-only binding to the contract
	PermissionsFilterer   // Log filterer for contract events
}

// PermissionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionsSession struct {
	Contract     *Permissions      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermissionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionsCallerSession struct {
	Contract *PermissionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PermissionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionsTransactorSession struct {
	Contract     *PermissionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PermissionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionsRaw struct {
	Contract *Permissions // Generic contract binding to access the raw methods on
}

// PermissionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionsCallerRaw struct {
	Contract *PermissionsCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionsTransactorRaw struct {
	Contract *PermissionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissions creates a new instance of Permissions, bound to a specific deployed contract.
func NewPermissions(address common.Address, backend bind.ContractBackend) (*Permissions, error) {
	contract, err := bindPermissions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Permissions{PermissionsCaller: PermissionsCaller{contract: contract}, PermissionsTransactor: PermissionsTransactor{contract: contract}, PermissionsFilterer: PermissionsFilterer{contract: contract}}, nil
}

// NewPermissionsCaller creates a new read-only instance of Permissions, bound to a specific deployed contract.
func NewPermissionsCaller(address common.Address, caller bind.ContractCaller) (*PermissionsCaller, error) {
	contract, err := bindPermissions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionsCaller{contract: contract}, nil
}

// NewPermissionsTransactor creates a new write-only instance of Permissions, bound to a specific deployed contract.
func NewPermissionsTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionsTransactor, error) {
	contract, err := bindPermissions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionsTransactor{contract: contract}, nil
}

// NewPermissionsFilterer creates a new log filterer instance of Permissions, bound to a specific deployed contract.
func NewPermissionsFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionsFilterer, error) {
	contract, err := bindPermissions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionsFilterer{contract: contract}, nil
}

// bindPermissions binds a generic wrapper to an already deployed contract.
func bindPermissions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermissionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissions *PermissionsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissions.Contract.PermissionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissions *PermissionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.Contract.PermissionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissions *PermissionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissions.Contract.PermissionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Permissions *PermissionsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Permissions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Permissions *PermissionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Permissions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Permissions *PermissionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Permissions.Contract.contract.Transact(opts, method, params...)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(_status uint8)
func (_Permissions *PermissionsCaller) GetNodeStatus(opts *bind.CallOpts, _enodeId string) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Permissions.contract.Call(opts, out, "getNodeStatus", _enodeId)
	return *ret0, err
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(_status uint8)
func (_Permissions *PermissionsSession) GetNodeStatus(_enodeId string) (uint8, error) {
	return _Permissions.Contract.GetNodeStatus(&_Permissions.CallOpts, _enodeId)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(_status uint8)
func (_Permissions *PermissionsCallerSession) GetNodeStatus(_enodeId string) (uint8, error) {
	return _Permissions.Contract.GetNodeStatus(&_Permissions.CallOpts, _enodeId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0xc1dd1d92.
//
// Solidity: function ApproveNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactor) ApproveNode(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "ApproveNode", _enodeId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0xc1dd1d92.
//
// Solidity: function ApproveNode(_enodeId string) returns()
func (_Permissions *PermissionsSession) ApproveNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ApproveNode(&_Permissions.TransactOpts, _enodeId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0xc1dd1d92.
//
// Solidity: function ApproveNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactorSession) ApproveNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ApproveNode(&_Permissions.TransactOpts, _enodeId)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0xbc67f80b.
//
// Solidity: function DeactivateNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactor) DeactivateNode(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "DeactivateNode", _enodeId)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0xbc67f80b.
//
// Solidity: function DeactivateNode(_enodeId string) returns()
func (_Permissions *PermissionsSession) DeactivateNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.DeactivateNode(&_Permissions.TransactOpts, _enodeId)
}

// DeactivateNode is a paid mutator transaction binding the contract method 0xbc67f80b.
//
// Solidity: function DeactivateNode(_enodeId string) returns()
func (_Permissions *PermissionsTransactorSession) DeactivateNode(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.DeactivateNode(&_Permissions.TransactOpts, _enodeId)
}

// ProposeDeactivation is a paid mutator transaction binding the contract method 0x6dcd03e0.
//
// Solidity: function ProposeDeactivation(_enodeId string) returns()
func (_Permissions *PermissionsTransactor) ProposeDeactivation(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "ProposeDeactivation", _enodeId)
}

// ProposeDeactivation is a paid mutator transaction binding the contract method 0x6dcd03e0.
//
// Solidity: function ProposeDeactivation(_enodeId string) returns()
func (_Permissions *PermissionsSession) ProposeDeactivation(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeDeactivation(&_Permissions.TransactOpts, _enodeId)
}

// ProposeDeactivation is a paid mutator transaction binding the contract method 0x6dcd03e0.
//
// Solidity: function ProposeDeactivation(_enodeId string) returns()
func (_Permissions *PermissionsTransactorSession) ProposeDeactivation(_enodeId string) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeDeactivation(&_Permissions.TransactOpts, _enodeId)
}

// ProposeNode is a paid mutator transaction binding the contract method 0x9e1ebc14.
//
// Solidity: function ProposeNode(_enodeId string, _canWrite bool, _canLead bool) returns()
func (_Permissions *PermissionsTransactor) ProposeNode(opts *bind.TransactOpts, _enodeId string, _canWrite bool, _canLead bool) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "ProposeNode", _enodeId, _canWrite, _canLead)
}

// ProposeNode is a paid mutator transaction binding the contract method 0x9e1ebc14.
//
// Solidity: function ProposeNode(_enodeId string, _canWrite bool, _canLead bool) returns()
func (_Permissions *PermissionsSession) ProposeNode(_enodeId string, _canWrite bool, _canLead bool) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeNode(&_Permissions.TransactOpts, _enodeId, _canWrite, _canLead)
}

// ProposeNode is a paid mutator transaction binding the contract method 0x9e1ebc14.
//
// Solidity: function ProposeNode(_enodeId string, _canWrite bool, _canLead bool) returns()
func (_Permissions *PermissionsTransactorSession) ProposeNode(_enodeId string, _canWrite bool, _canLead bool) (*types.Transaction, error) {
	return _Permissions.Contract.ProposeNode(&_Permissions.TransactOpts, _enodeId, _canWrite, _canLead)
}

// UpdateAcctAccess is a paid mutator transaction binding the contract method 0x913fc7d7.
//
// Solidity: function updateAcctAccess(_acctId address, access uint8) returns()
func (_Permissions *PermissionsTransactor) UpdateAcctAccess(opts *bind.TransactOpts, _acctId common.Address, access uint8) (*types.Transaction, error) {
	return _Permissions.contract.Transact(opts, "updateAcctAccess", _acctId, access)
}

// UpdateAcctAccess is a paid mutator transaction binding the contract method 0x913fc7d7.
//
// Solidity: function updateAcctAccess(_acctId address, access uint8) returns()
func (_Permissions *PermissionsSession) UpdateAcctAccess(_acctId common.Address, access uint8) (*types.Transaction, error) {
	return _Permissions.Contract.UpdateAcctAccess(&_Permissions.TransactOpts, _acctId, access)
}

// UpdateAcctAccess is a paid mutator transaction binding the contract method 0x913fc7d7.
//
// Solidity: function updateAcctAccess(_acctId address, access uint8) returns()
func (_Permissions *PermissionsTransactorSession) UpdateAcctAccess(_acctId common.Address, access uint8) (*types.Transaction, error) {
	return _Permissions.Contract.UpdateAcctAccess(&_Permissions.TransactOpts, _acctId, access)
}

// PermissionsAcctAccessModifiedIterator is returned from FilterAcctAccessModified and is used to iterate over the raw logs and unpacked data for AcctAccessModified events raised by the Permissions contract.
type PermissionsAcctAccessModifiedIterator struct {
	Event *PermissionsAcctAccessModified // Event containing the contract specifics and raw log

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
func (it *PermissionsAcctAccessModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsAcctAccessModified)
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
		it.Event = new(PermissionsAcctAccessModified)
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
func (it *PermissionsAcctAccessModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsAcctAccessModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsAcctAccessModified represents a AcctAccessModified event raised by the Permissions contract.
type PermissionsAcctAccessModified struct {
	AcctId common.Address
	Access uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAcctAccessModified is a free log retrieval operation binding the contract event 0xb9f8d88def106a1a167b617a197bee4323d3997cb1fa46b785c1a7b9af6e82a8.
//
// Solidity: e AcctAccessModified(acctId address, access uint8)
func (_Permissions *PermissionsFilterer) FilterAcctAccessModified(opts *bind.FilterOpts) (*PermissionsAcctAccessModifiedIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "AcctAccessModified")
	if err != nil {
		return nil, err
	}
	return &PermissionsAcctAccessModifiedIterator{contract: _Permissions.contract, event: "AcctAccessModified", logs: logs, sub: sub}, nil
}

// WatchAcctAccessModified is a free log subscription operation binding the contract event 0xb9f8d88def106a1a167b617a197bee4323d3997cb1fa46b785c1a7b9af6e82a8.
//
// Solidity: e AcctAccessModified(acctId address, access uint8)
func (_Permissions *PermissionsFilterer) WatchAcctAccessModified(opts *bind.WatchOpts, sink chan<- *PermissionsAcctAccessModified) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "AcctAccessModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsAcctAccessModified)
				if err := _Permissions.contract.UnpackLog(event, "AcctAccessModified", log); err != nil {
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

// PermissionsNewNodeProposedIterator is returned from FilterNewNodeProposed and is used to iterate over the raw logs and unpacked data for NewNodeProposed events raised by the Permissions contract.
type PermissionsNewNodeProposedIterator struct {
	Event *PermissionsNewNodeProposed // Event containing the contract specifics and raw log

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
func (it *PermissionsNewNodeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNewNodeProposed)
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
		it.Event = new(PermissionsNewNodeProposed)
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
func (it *PermissionsNewNodeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNewNodeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNewNodeProposed represents a NewNodeProposed event raised by the Permissions contract.
type PermissionsNewNodeProposed struct {
	EnodeId  string
	CanWrite bool
	CanLead  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewNodeProposed is a free log retrieval operation binding the contract event 0xd630d8689250dadadf783ef2b5c310617046e2f16f9cddbb44e8d96801e3a347.
//
// Solidity: e NewNodeProposed(_enodeId string, _canWrite bool, _canLead bool)
func (_Permissions *PermissionsFilterer) FilterNewNodeProposed(opts *bind.FilterOpts) (*PermissionsNewNodeProposedIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NewNodeProposed")
	if err != nil {
		return nil, err
	}
	return &PermissionsNewNodeProposedIterator{contract: _Permissions.contract, event: "NewNodeProposed", logs: logs, sub: sub}, nil
}

// WatchNewNodeProposed is a free log subscription operation binding the contract event 0xd630d8689250dadadf783ef2b5c310617046e2f16f9cddbb44e8d96801e3a347.
//
// Solidity: e NewNodeProposed(_enodeId string, _canWrite bool, _canLead bool)
func (_Permissions *PermissionsFilterer) WatchNewNodeProposed(opts *bind.WatchOpts, sink chan<- *PermissionsNewNodeProposed) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NewNodeProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNewNodeProposed)
				if err := _Permissions.contract.UnpackLog(event, "NewNodeProposed", log); err != nil {
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

// PermissionsNodeApprovedIterator is returned from FilterNodeApproved and is used to iterate over the raw logs and unpacked data for NodeApproved events raised by the Permissions contract.
type PermissionsNodeApprovedIterator struct {
	Event *PermissionsNodeApproved // Event containing the contract specifics and raw log

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
func (it *PermissionsNodeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNodeApproved)
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
		it.Event = new(PermissionsNodeApproved)
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
func (it *PermissionsNodeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNodeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNodeApproved represents a NodeApproved event raised by the Permissions contract.
type PermissionsNodeApproved struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeApproved is a free log retrieval operation binding the contract event 0xc8f0c6e7f31c7ba4e6e29615ae2ab658fdda704c49912bb6118db07a4c36d478.
//
// Solidity: e NodeApproved(_enodeId string)
func (_Permissions *PermissionsFilterer) FilterNodeApproved(opts *bind.FilterOpts) (*PermissionsNodeApprovedIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return &PermissionsNodeApprovedIterator{contract: _Permissions.contract, event: "NodeApproved", logs: logs, sub: sub}, nil
}

// WatchNodeApproved is a free log subscription operation binding the contract event 0xc8f0c6e7f31c7ba4e6e29615ae2ab658fdda704c49912bb6118db07a4c36d478.
//
// Solidity: e NodeApproved(_enodeId string)
func (_Permissions *PermissionsFilterer) WatchNodeApproved(opts *bind.WatchOpts, sink chan<- *PermissionsNodeApproved) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNodeApproved)
				if err := _Permissions.contract.UnpackLog(event, "NodeApproved", log); err != nil {
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

// PermissionsNodeDeactivatedIterator is returned from FilterNodeDeactivated and is used to iterate over the raw logs and unpacked data for NodeDeactivated events raised by the Permissions contract.
type PermissionsNodeDeactivatedIterator struct {
	Event *PermissionsNodeDeactivated // Event containing the contract specifics and raw log

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
func (it *PermissionsNodeDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNodeDeactivated)
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
		it.Event = new(PermissionsNodeDeactivated)
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
func (it *PermissionsNodeDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNodeDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNodeDeactivated represents a NodeDeactivated event raised by the Permissions contract.
type PermissionsNodeDeactivated struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeDeactivated is a free log retrieval operation binding the contract event 0xb4551525dafbacbcbad53f3a1ad477e2de2428dcd5832ae46d8edacf8c2959d5.
//
// Solidity: e NodeDeactivated(_enodeId string)
func (_Permissions *PermissionsFilterer) FilterNodeDeactivated(opts *bind.FilterOpts) (*PermissionsNodeDeactivatedIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return &PermissionsNodeDeactivatedIterator{contract: _Permissions.contract, event: "NodeDeactivated", logs: logs, sub: sub}, nil
}

// WatchNodeDeactivated is a free log subscription operation binding the contract event 0xb4551525dafbacbcbad53f3a1ad477e2de2428dcd5832ae46d8edacf8c2959d5.
//
// Solidity: e NodeDeactivated(_enodeId string)
func (_Permissions *PermissionsFilterer) WatchNodeDeactivated(opts *bind.WatchOpts, sink chan<- *PermissionsNodeDeactivated) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNodeDeactivated)
				if err := _Permissions.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
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

// PermissionsNodePendingDeactivationIterator is returned from FilterNodePendingDeactivation and is used to iterate over the raw logs and unpacked data for NodePendingDeactivation events raised by the Permissions contract.
type PermissionsNodePendingDeactivationIterator struct {
	Event *PermissionsNodePendingDeactivation // Event containing the contract specifics and raw log

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
func (it *PermissionsNodePendingDeactivationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionsNodePendingDeactivation)
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
		it.Event = new(PermissionsNodePendingDeactivation)
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
func (it *PermissionsNodePendingDeactivationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionsNodePendingDeactivationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionsNodePendingDeactivation represents a NodePendingDeactivation event raised by the Permissions contract.
type PermissionsNodePendingDeactivation struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodePendingDeactivation is a free log retrieval operation binding the contract event 0x2b5689b33f48f1dcbda2084e130a9bee7b3bf14dc767ea74cbdf3e5fffb118e4.
//
// Solidity: e NodePendingDeactivation(_enodeId string)
func (_Permissions *PermissionsFilterer) FilterNodePendingDeactivation(opts *bind.FilterOpts) (*PermissionsNodePendingDeactivationIterator, error) {

	logs, sub, err := _Permissions.contract.FilterLogs(opts, "NodePendingDeactivation")
	if err != nil {
		return nil, err
	}
	return &PermissionsNodePendingDeactivationIterator{contract: _Permissions.contract, event: "NodePendingDeactivation", logs: logs, sub: sub}, nil
}

// WatchNodePendingDeactivation is a free log subscription operation binding the contract event 0x2b5689b33f48f1dcbda2084e130a9bee7b3bf14dc767ea74cbdf3e5fffb118e4.
//
// Solidity: e NodePendingDeactivation(_enodeId string)
func (_Permissions *PermissionsFilterer) WatchNodePendingDeactivation(opts *bind.WatchOpts, sink chan<- *PermissionsNodePendingDeactivation) (event.Subscription, error) {

	logs, sub, err := _Permissions.contract.WatchLogs(opts, "NodePendingDeactivation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionsNodePendingDeactivation)
				if err := _Permissions.contract.UnpackLog(event, "NodePendingDeactivation", log); err != nil {
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
