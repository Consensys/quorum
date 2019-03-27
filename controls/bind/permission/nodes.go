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

// NodeManagerABI is the input ABI used to generate the binding from.
const NodeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"approveNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"getNodeStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"enodeId\",\"type\":\"string\"}],\"name\":\"getNodeDetails\",\"outputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_nodeStatus\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addOrgNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeIndex\",\"type\":\"uint256\"}],\"name\":\"getNodeDetailsFromIndex\",\"outputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_nodeStatus\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"addNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfNodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodeApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodePendingDeactivation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodeDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodePendingActivation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enodeId\",\"type\":\"string\"}],\"name\":\"NodePendingBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"string\"}],\"name\":\"NodeBlacklisted\",\"type\":\"event\"}]"

// NodeManager is an auto generated Go binding around an Ethereum contract.
type NodeManager struct {
	NodeManagerCaller     // Read-only binding to the contract
	NodeManagerTransactor // Write-only binding to the contract
	NodeManagerFilterer   // Log filterer for contract events
}

// NodeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeManagerSession struct {
	Contract     *NodeManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeManagerCallerSession struct {
	Contract *NodeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeManagerTransactorSession struct {
	Contract     *NodeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeManagerRaw struct {
	Contract *NodeManager // Generic contract binding to access the raw methods on
}

// NodeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeManagerCallerRaw struct {
	Contract *NodeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NodeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeManagerTransactorRaw struct {
	Contract *NodeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeManager creates a new instance of NodeManager, bound to a specific deployed contract.
func NewNodeManager(address common.Address, backend bind.ContractBackend) (*NodeManager, error) {
	contract, err := bindNodeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeManager{NodeManagerCaller: NodeManagerCaller{contract: contract}, NodeManagerTransactor: NodeManagerTransactor{contract: contract}, NodeManagerFilterer: NodeManagerFilterer{contract: contract}}, nil
}

// NewNodeManagerCaller creates a new read-only instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerCaller(address common.Address, caller bind.ContractCaller) (*NodeManagerCaller, error) {
	contract, err := bindNodeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerCaller{contract: contract}, nil
}

// NewNodeManagerTransactor creates a new write-only instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeManagerTransactor, error) {
	contract, err := bindNodeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeManagerTransactor{contract: contract}, nil
}

// NewNodeManagerFilterer creates a new log filterer instance of NodeManager, bound to a specific deployed contract.
func NewNodeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeManagerFilterer, error) {
	contract, err := bindNodeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeManagerFilterer{contract: contract}, nil
}

// bindNodeManager binds a generic wrapper to an already deployed contract.
func bindNodeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManager *NodeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeManager.Contract.NodeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManager *NodeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManager.Contract.NodeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManager *NodeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManager.Contract.NodeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeManager *NodeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NodeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeManager *NodeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeManager *NodeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeManager.Contract.contract.Transact(opts, method, params...)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(enodeId string) constant returns(_enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerCaller) GetNodeDetails(opts *bind.CallOpts, enodeId string) (struct {
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	ret := new(struct {
		EnodeId    string
		NodeStatus *big.Int
	})
	out := ret
	err := _NodeManager.contract.Call(opts, out, "getNodeDetails", enodeId)
	return *ret, err
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(enodeId string) constant returns(_enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerSession) GetNodeDetails(enodeId string) (struct {
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetails(&_NodeManager.CallOpts, enodeId)
}

// GetNodeDetails is a free data retrieval call binding the contract method 0x3f0e0e47.
//
// Solidity: function getNodeDetails(enodeId string) constant returns(_enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerCallerSession) GetNodeDetails(enodeId string) (struct {
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetails(&_NodeManager.CallOpts, enodeId)
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(nodeIndex uint256) constant returns(_orgId string, _enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerCaller) GetNodeDetailsFromIndex(opts *bind.CallOpts, nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	ret := new(struct {
		OrgId      string
		EnodeId    string
		NodeStatus *big.Int
	})
	out := ret
	err := _NodeManager.contract.Call(opts, out, "getNodeDetailsFromIndex", nodeIndex)
	return *ret, err
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(nodeIndex uint256) constant returns(_orgId string, _enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerSession) GetNodeDetailsFromIndex(nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetailsFromIndex(&_NodeManager.CallOpts, nodeIndex)
}

// GetNodeDetailsFromIndex is a free data retrieval call binding the contract method 0x97c07a9b.
//
// Solidity: function getNodeDetailsFromIndex(nodeIndex uint256) constant returns(_orgId string, _enodeId string, _nodeStatus uint256)
func (_NodeManager *NodeManagerCallerSession) GetNodeDetailsFromIndex(nodeIndex *big.Int) (struct {
	OrgId      string
	EnodeId    string
	NodeStatus *big.Int
}, error) {
	return _NodeManager.Contract.GetNodeDetailsFromIndex(&_NodeManager.CallOpts, nodeIndex)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(uint256)
func (_NodeManager *NodeManagerCaller) GetNodeStatus(opts *bind.CallOpts, _enodeId string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeManager.contract.Call(opts, out, "getNodeStatus", _enodeId)
	return *ret0, err
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(uint256)
func (_NodeManager *NodeManagerSession) GetNodeStatus(_enodeId string) (*big.Int, error) {
	return _NodeManager.Contract.GetNodeStatus(&_NodeManager.CallOpts, _enodeId)
}

// GetNodeStatus is a free data retrieval call binding the contract method 0x397eeccb.
//
// Solidity: function getNodeStatus(_enodeId string) constant returns(uint256)
func (_NodeManager *NodeManagerCallerSession) GetNodeStatus(_enodeId string) (*big.Int, error) {
	return _NodeManager.Contract.GetNodeStatus(&_NodeManager.CallOpts, _enodeId)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_NodeManager *NodeManagerCaller) GetNumberOfNodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NodeManager.contract.Call(opts, out, "getNumberOfNodes")
	return *ret0, err
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_NodeManager *NodeManagerSession) GetNumberOfNodes() (*big.Int, error) {
	return _NodeManager.Contract.GetNumberOfNodes(&_NodeManager.CallOpts)
}

// GetNumberOfNodes is a free data retrieval call binding the contract method 0xb81c806a.
//
// Solidity: function getNumberOfNodes() constant returns(uint256)
func (_NodeManager *NodeManagerCallerSession) GetNumberOfNodes() (*big.Int, error) {
	return _NodeManager.Contract.GetNumberOfNodes(&_NodeManager.CallOpts)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactor) AddNode(opts *bind.TransactOpts, _enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "addNode", _enodeId, _orgId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerSession) AddNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddNode is a paid mutator transaction binding the contract method 0xa97a4406.
//
// Solidity: function addNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactorSession) AddNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x3f5e1a45.
//
// Solidity: function addOrgNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactor) AddOrgNode(opts *bind.TransactOpts, _enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "addOrgNode", _enodeId, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x3f5e1a45.
//
// Solidity: function addOrgNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerSession) AddOrgNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddOrgNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// AddOrgNode is a paid mutator transaction binding the contract method 0x3f5e1a45.
//
// Solidity: function addOrgNode(_enodeId string, _orgId string) returns()
func (_NodeManager *NodeManagerTransactorSession) AddOrgNode(_enodeId string, _orgId string) (*types.Transaction, error) {
	return _NodeManager.Contract.AddOrgNode(&_NodeManager.TransactOpts, _enodeId, _orgId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x21c67088.
//
// Solidity: function approveNode(_enodeId string) returns()
func (_NodeManager *NodeManagerTransactor) ApproveNode(opts *bind.TransactOpts, _enodeId string) (*types.Transaction, error) {
	return _NodeManager.contract.Transact(opts, "approveNode", _enodeId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x21c67088.
//
// Solidity: function approveNode(_enodeId string) returns()
func (_NodeManager *NodeManagerSession) ApproveNode(_enodeId string) (*types.Transaction, error) {
	return _NodeManager.Contract.ApproveNode(&_NodeManager.TransactOpts, _enodeId)
}

// ApproveNode is a paid mutator transaction binding the contract method 0x21c67088.
//
// Solidity: function approveNode(_enodeId string) returns()
func (_NodeManager *NodeManagerTransactorSession) ApproveNode(_enodeId string) (*types.Transaction, error) {
	return _NodeManager.Contract.ApproveNode(&_NodeManager.TransactOpts, _enodeId)
}

// NodeManagerNodeActivatedIterator is returned from FilterNodeActivated and is used to iterate over the raw logs and unpacked data for NodeActivated events raised by the NodeManager contract.
type NodeManagerNodeActivatedIterator struct {
	Event *NodeManagerNodeActivated // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeActivated)
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
		it.Event = new(NodeManagerNodeActivated)
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
func (it *NodeManagerNodeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeActivated represents a NodeActivated event raised by the NodeManager contract.
type NodeManagerNodeActivated struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeActivated is a free log retrieval operation binding the contract event 0xeee4ff0a0bd593fd1a0d7504eca37b4027b72bc4cc3df6bf3d1c7e27848d279e.
//
// Solidity: e NodeActivated(_enodeId string)
func (_NodeManager *NodeManagerFilterer) FilterNodeActivated(opts *bind.FilterOpts) (*NodeManagerNodeActivatedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeActivated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeActivatedIterator{contract: _NodeManager.contract, event: "NodeActivated", logs: logs, sub: sub}, nil
}

// WatchNodeActivated is a free log subscription operation binding the contract event 0xeee4ff0a0bd593fd1a0d7504eca37b4027b72bc4cc3df6bf3d1c7e27848d279e.
//
// Solidity: e NodeActivated(_enodeId string)
func (_NodeManager *NodeManagerFilterer) WatchNodeActivated(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeActivated) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeActivated)
				if err := _NodeManager.contract.UnpackLog(event, "NodeActivated", log); err != nil {
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

// NodeManagerNodeApprovedIterator is returned from FilterNodeApproved and is used to iterate over the raw logs and unpacked data for NodeApproved events raised by the NodeManager contract.
type NodeManagerNodeApprovedIterator struct {
	Event *NodeManagerNodeApproved // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeApproved)
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
		it.Event = new(NodeManagerNodeApproved)
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
func (it *NodeManagerNodeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeApproved represents a NodeApproved event raised by the NodeManager contract.
type NodeManagerNodeApproved struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeApproved is a free log retrieval operation binding the contract event 0xc8f0c6e7f31c7ba4e6e29615ae2ab658fdda704c49912bb6118db07a4c36d478.
//
// Solidity: e NodeApproved(_enodeId string)
func (_NodeManager *NodeManagerFilterer) FilterNodeApproved(opts *bind.FilterOpts) (*NodeManagerNodeApprovedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeApprovedIterator{contract: _NodeManager.contract, event: "NodeApproved", logs: logs, sub: sub}, nil
}

// WatchNodeApproved is a free log subscription operation binding the contract event 0xc8f0c6e7f31c7ba4e6e29615ae2ab658fdda704c49912bb6118db07a4c36d478.
//
// Solidity: e NodeApproved(_enodeId string)
func (_NodeManager *NodeManagerFilterer) WatchNodeApproved(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeApproved) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeApproved)
				if err := _NodeManager.contract.UnpackLog(event, "NodeApproved", log); err != nil {
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

// NodeManagerNodeBlacklistedIterator is returned from FilterNodeBlacklisted and is used to iterate over the raw logs and unpacked data for NodeBlacklisted events raised by the NodeManager contract.
type NodeManagerNodeBlacklistedIterator struct {
	Event *NodeManagerNodeBlacklisted // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeBlacklisted)
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
		it.Event = new(NodeManagerNodeBlacklisted)
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
func (it *NodeManagerNodeBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeBlacklisted represents a NodeBlacklisted event raised by the NodeManager contract.
type NodeManagerNodeBlacklisted struct {
	string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNodeBlacklisted is a free log retrieval operation binding the contract event 0xf97ae29ad6493ecc1ebba750b96015715fca52d1f0160b48f39788cbc6204a8e.
//
// Solidity: e NodeBlacklisted( string)
func (_NodeManager *NodeManagerFilterer) FilterNodeBlacklisted(opts *bind.FilterOpts) (*NodeManagerNodeBlacklistedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeBlacklisted")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeBlacklistedIterator{contract: _NodeManager.contract, event: "NodeBlacklisted", logs: logs, sub: sub}, nil
}

// WatchNodeBlacklisted is a free log subscription operation binding the contract event 0xf97ae29ad6493ecc1ebba750b96015715fca52d1f0160b48f39788cbc6204a8e.
//
// Solidity: e NodeBlacklisted( string)
func (_NodeManager *NodeManagerFilterer) WatchNodeBlacklisted(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeBlacklisted) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeBlacklisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeBlacklisted)
				if err := _NodeManager.contract.UnpackLog(event, "NodeBlacklisted", log); err != nil {
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

// NodeManagerNodeDeactivatedIterator is returned from FilterNodeDeactivated and is used to iterate over the raw logs and unpacked data for NodeDeactivated events raised by the NodeManager contract.
type NodeManagerNodeDeactivatedIterator struct {
	Event *NodeManagerNodeDeactivated // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeDeactivated)
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
		it.Event = new(NodeManagerNodeDeactivated)
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
func (it *NodeManagerNodeDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeDeactivated represents a NodeDeactivated event raised by the NodeManager contract.
type NodeManagerNodeDeactivated struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeDeactivated is a free log retrieval operation binding the contract event 0xb4551525dafbacbcbad53f3a1ad477e2de2428dcd5832ae46d8edacf8c2959d5.
//
// Solidity: e NodeDeactivated(_enodeId string)
func (_NodeManager *NodeManagerFilterer) FilterNodeDeactivated(opts *bind.FilterOpts) (*NodeManagerNodeDeactivatedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeDeactivatedIterator{contract: _NodeManager.contract, event: "NodeDeactivated", logs: logs, sub: sub}, nil
}

// WatchNodeDeactivated is a free log subscription operation binding the contract event 0xb4551525dafbacbcbad53f3a1ad477e2de2428dcd5832ae46d8edacf8c2959d5.
//
// Solidity: e NodeDeactivated(_enodeId string)
func (_NodeManager *NodeManagerFilterer) WatchNodeDeactivated(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeDeactivated) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeDeactivated)
				if err := _NodeManager.contract.UnpackLog(event, "NodeDeactivated", log); err != nil {
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

// NodeManagerNodePendingActivationIterator is returned from FilterNodePendingActivation and is used to iterate over the raw logs and unpacked data for NodePendingActivation events raised by the NodeManager contract.
type NodeManagerNodePendingActivationIterator struct {
	Event *NodeManagerNodePendingActivation // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodePendingActivationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodePendingActivation)
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
		it.Event = new(NodeManagerNodePendingActivation)
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
func (it *NodeManagerNodePendingActivationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodePendingActivationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodePendingActivation represents a NodePendingActivation event raised by the NodeManager contract.
type NodeManagerNodePendingActivation struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodePendingActivation is a free log retrieval operation binding the contract event 0x7b961104d9e9db7d30803aff3fa117bc41799d2faa2d2e339cf1a1f3513b0eef.
//
// Solidity: e NodePendingActivation(_enodeId string)
func (_NodeManager *NodeManagerFilterer) FilterNodePendingActivation(opts *bind.FilterOpts) (*NodeManagerNodePendingActivationIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodePendingActivation")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodePendingActivationIterator{contract: _NodeManager.contract, event: "NodePendingActivation", logs: logs, sub: sub}, nil
}

// WatchNodePendingActivation is a free log subscription operation binding the contract event 0x7b961104d9e9db7d30803aff3fa117bc41799d2faa2d2e339cf1a1f3513b0eef.
//
// Solidity: e NodePendingActivation(_enodeId string)
func (_NodeManager *NodeManagerFilterer) WatchNodePendingActivation(opts *bind.WatchOpts, sink chan<- *NodeManagerNodePendingActivation) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodePendingActivation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodePendingActivation)
				if err := _NodeManager.contract.UnpackLog(event, "NodePendingActivation", log); err != nil {
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

// NodeManagerNodePendingBlacklistIterator is returned from FilterNodePendingBlacklist and is used to iterate over the raw logs and unpacked data for NodePendingBlacklist events raised by the NodeManager contract.
type NodeManagerNodePendingBlacklistIterator struct {
	Event *NodeManagerNodePendingBlacklist // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodePendingBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodePendingBlacklist)
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
		it.Event = new(NodeManagerNodePendingBlacklist)
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
func (it *NodeManagerNodePendingBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodePendingBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodePendingBlacklist represents a NodePendingBlacklist event raised by the NodeManager contract.
type NodeManagerNodePendingBlacklist struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodePendingBlacklist is a free log retrieval operation binding the contract event 0xb249ebebf429f1c79f3c9663998b3e22d45f242de6527c4a95e41d4d28115d74.
//
// Solidity: e NodePendingBlacklist(_enodeId string)
func (_NodeManager *NodeManagerFilterer) FilterNodePendingBlacklist(opts *bind.FilterOpts) (*NodeManagerNodePendingBlacklistIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodePendingBlacklist")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodePendingBlacklistIterator{contract: _NodeManager.contract, event: "NodePendingBlacklist", logs: logs, sub: sub}, nil
}

// WatchNodePendingBlacklist is a free log subscription operation binding the contract event 0xb249ebebf429f1c79f3c9663998b3e22d45f242de6527c4a95e41d4d28115d74.
//
// Solidity: e NodePendingBlacklist(_enodeId string)
func (_NodeManager *NodeManagerFilterer) WatchNodePendingBlacklist(opts *bind.WatchOpts, sink chan<- *NodeManagerNodePendingBlacklist) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodePendingBlacklist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodePendingBlacklist)
				if err := _NodeManager.contract.UnpackLog(event, "NodePendingBlacklist", log); err != nil {
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

// NodeManagerNodePendingDeactivationIterator is returned from FilterNodePendingDeactivation and is used to iterate over the raw logs and unpacked data for NodePendingDeactivation events raised by the NodeManager contract.
type NodeManagerNodePendingDeactivationIterator struct {
	Event *NodeManagerNodePendingDeactivation // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodePendingDeactivationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodePendingDeactivation)
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
		it.Event = new(NodeManagerNodePendingDeactivation)
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
func (it *NodeManagerNodePendingDeactivationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodePendingDeactivationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodePendingDeactivation represents a NodePendingDeactivation event raised by the NodeManager contract.
type NodeManagerNodePendingDeactivation struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodePendingDeactivation is a free log retrieval operation binding the contract event 0x2b5689b33f48f1dcbda2084e130a9bee7b3bf14dc767ea74cbdf3e5fffb118e4.
//
// Solidity: e NodePendingDeactivation(_enodeId string)
func (_NodeManager *NodeManagerFilterer) FilterNodePendingDeactivation(opts *bind.FilterOpts) (*NodeManagerNodePendingDeactivationIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodePendingDeactivation")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodePendingDeactivationIterator{contract: _NodeManager.contract, event: "NodePendingDeactivation", logs: logs, sub: sub}, nil
}

// WatchNodePendingDeactivation is a free log subscription operation binding the contract event 0x2b5689b33f48f1dcbda2084e130a9bee7b3bf14dc767ea74cbdf3e5fffb118e4.
//
// Solidity: e NodePendingDeactivation(_enodeId string)
func (_NodeManager *NodeManagerFilterer) WatchNodePendingDeactivation(opts *bind.WatchOpts, sink chan<- *NodeManagerNodePendingDeactivation) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodePendingDeactivation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodePendingDeactivation)
				if err := _NodeManager.contract.UnpackLog(event, "NodePendingDeactivation", log); err != nil {
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

// NodeManagerNodeProposedIterator is returned from FilterNodeProposed and is used to iterate over the raw logs and unpacked data for NodeProposed events raised by the NodeManager contract.
type NodeManagerNodeProposedIterator struct {
	Event *NodeManagerNodeProposed // Event containing the contract specifics and raw log

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
func (it *NodeManagerNodeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeManagerNodeProposed)
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
		it.Event = new(NodeManagerNodeProposed)
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
func (it *NodeManagerNodeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeManagerNodeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeManagerNodeProposed represents a NodeProposed event raised by the NodeManager contract.
type NodeManagerNodeProposed struct {
	EnodeId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeProposed is a free log retrieval operation binding the contract event 0xaddedf3fcf588a85e0b0c3210c30da3f5597ae35221859f7e19427397a2ba80a.
//
// Solidity: e NodeProposed(_enodeId string)
func (_NodeManager *NodeManagerFilterer) FilterNodeProposed(opts *bind.FilterOpts) (*NodeManagerNodeProposedIterator, error) {

	logs, sub, err := _NodeManager.contract.FilterLogs(opts, "NodeProposed")
	if err != nil {
		return nil, err
	}
	return &NodeManagerNodeProposedIterator{contract: _NodeManager.contract, event: "NodeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeProposed is a free log subscription operation binding the contract event 0xaddedf3fcf588a85e0b0c3210c30da3f5597ae35221859f7e19427397a2ba80a.
//
// Solidity: e NodeProposed(_enodeId string)
func (_NodeManager *NodeManagerFilterer) WatchNodeProposed(opts *bind.WatchOpts, sink chan<- *NodeManagerNodeProposed) (event.Subscription, error) {

	logs, sub, err := _NodeManager.contract.WatchLogs(opts, "NodeProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeManagerNodeProposed)
				if err := _NodeManager.contract.UnpackLog(event, "NodeProposed", log); err != nil {
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
