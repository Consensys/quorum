// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ENSABI is the input ABI used to generate the binding from.
const ENSABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"label\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"NewResolver\",\"type\":\"event\"}]"

// ENSBin is the compiled bytecode used for deploying new contracts.
const ENSBin = `0x608060405234801561001057600080fd5b50604051602080610409833981016040525160008080526020527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb58054600160a060020a03909216600160a060020a0319909216919091179055610390806100796000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630178b8bf811461007157806302571be3146100a557806306ab5923146100bd5780631896f70a146100e65780635b0fc9c31461010a575b600080fd5b34801561007d57600080fd5b5061008960043561012e565b60408051600160a060020a039092168252519081900360200190f35b3480156100b157600080fd5b5061008960043561014c565b3480156100c957600080fd5b506100e4600435602435600160a060020a0360443516610167565b005b3480156100f257600080fd5b506100e4600435600160a060020a0360243516610221565b34801561011657600080fd5b506100e4600435600160a060020a03602435166102c4565b600090815260208190526040902060010154600160a060020a031690565b600090815260208190526040902054600160a060020a031690565b6000838152602081905260408120548490600160a060020a0316331461018c57600080fd5b60408051868152602080820187905282519182900383018220600160a060020a03871683529251929450869288927fce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e8292908290030190a3506000908152602081905260409020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555050565b6000828152602081905260409020548290600160a060020a0316331461024657600080fd5b60408051600160a060020a0384168152905184917f335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0919081900360200190a250600091825260208290526040909120600101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909216919091179055565b6000828152602081905260409020548290600160a060020a031633146102e957600080fd5b60408051600160a060020a0384168152905184917fd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266919081900360200190a250600091825260208290526040909120805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039092169190911790555600a165627a7a72305820df4aac84ae1b2457acdbef94dbcf9e82967d538dc402aa754b2cddc657b9dee40029`

// DeployENS deploys a new Ethereum contract, binding an instance of ENS to it.
func DeployENS(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address) (common.Address, *types.Transaction, *ENS, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ENSBin), backend, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ENS{ENSCaller: ENSCaller{contract: contract}, ENSTransactor: ENSTransactor{contract: contract}, ENSFilterer: ENSFilterer{contract: contract}}, nil
}

// ENS is an auto generated Go binding around an Ethereum contract.
type ENS struct {
	ENSCaller     // Read-only binding to the contract
	ENSTransactor // Write-only binding to the contract
	ENSFilterer   // Log filterer for contract events
}

// ENSCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ENSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENSSession struct {
	Contract     *ENS              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENSCallerSession struct {
	Contract *ENSCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ENSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENSTransactorSession struct {
	Contract     *ENSTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENSRaw struct {
	Contract *ENS // Generic contract binding to access the raw methods on
}

// ENSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENSCallerRaw struct {
	Contract *ENSCaller // Generic read-only contract binding to access the raw methods on
}

// ENSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENSTransactorRaw struct {
	Contract *ENSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENS creates a new instance of ENS, bound to a specific deployed contract.
func NewENS(address common.Address, backend bind.ContractBackend) (*ENS, error) {
	contract, err := bindENS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENS{ENSCaller: ENSCaller{contract: contract}, ENSTransactor: ENSTransactor{contract: contract}, ENSFilterer: ENSFilterer{contract: contract}}, nil
}

// NewENSCaller creates a new read-only instance of ENS, bound to a specific deployed contract.
func NewENSCaller(address common.Address, caller bind.ContractCaller) (*ENSCaller, error) {
	contract, err := bindENS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ENSCaller{contract: contract}, nil
}

// NewENSTransactor creates a new write-only instance of ENS, bound to a specific deployed contract.
func NewENSTransactor(address common.Address, transactor bind.ContractTransactor) (*ENSTransactor, error) {
	contract, err := bindENS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ENSTransactor{contract: contract}, nil
}

// NewENSFilterer creates a new log filterer instance of ENS, bound to a specific deployed contract.
func NewENSFilterer(address common.Address, filterer bind.ContractFilterer) (*ENSFilterer, error) {
	contract, err := bindENS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ENSFilterer{contract: contract}, nil
}

// bindENS binds a generic wrapper to an already deployed contract.
func bindENS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENS *ENSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ENS.Contract.ENSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENS *ENSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENS.Contract.ENSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENS *ENSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENS.Contract.ENSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENS *ENSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ENS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENS *ENSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENS *ENSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENS.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_ENS *ENSCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ENS.contract.Call(opts, out, "owner", node)
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_ENS *ENSSession) Owner(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Owner(&_ENS.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(node bytes32) constant returns(address)
func (_ENS *ENSCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Owner(&_ENS.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_ENS *ENSCaller) Resolver(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ENS.contract.Call(opts, out, "resolver", node)
	return *ret0, err
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_ENS *ENSSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Resolver(&_ENS.CallOpts, node)
}

// Resolver is a free data retrieval call binding the contract method 0x0178b8bf.
//
// Solidity: function resolver(node bytes32) constant returns(address)
func (_ENS *ENSCallerSession) Resolver(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Resolver(&_ENS.CallOpts, node)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_ENS *ENSTransactor) SetOwner(opts *bind.TransactOpts, node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setOwner", node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_ENS *ENSSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetOwner(&_ENS.TransactOpts, node, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x5b0fc9c3.
//
// Solidity: function setOwner(node bytes32, owner address) returns()
func (_ENS *ENSTransactorSession) SetOwner(node [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetOwner(&_ENS.TransactOpts, node, owner)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_ENS *ENSTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_ENS *ENSSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetResolver(&_ENS.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(node bytes32, resolver address) returns()
func (_ENS *ENSTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetResolver(&_ENS.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_ENS *ENSTransactor) SetSubnodeOwner(opts *bind.TransactOpts, node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.contract.Transact(opts, "setSubnodeOwner", node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_ENS *ENSSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetSubnodeOwner(&_ENS.TransactOpts, node, label, owner)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0x06ab5923.
//
// Solidity: function setSubnodeOwner(node bytes32, label bytes32, owner address) returns()
func (_ENS *ENSTransactorSession) SetSubnodeOwner(node [32]byte, label [32]byte, owner common.Address) (*types.Transaction, error) {
	return _ENS.Contract.SetSubnodeOwner(&_ENS.TransactOpts, node, label, owner)
}

// ENSNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the ENS contract.
type ENSNewOwnerIterator struct {
	Event *ENSNewOwner // Event containing the contract specifics and raw log

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
func (it *ENSNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSNewOwner)
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
		it.Event = new(ENSNewOwner)
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
func (it *ENSNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSNewOwner represents a NewOwner event raised by the ENS contract.
type ENSNewOwner struct {
	Node  [32]byte
	Label [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: e NewOwner(node indexed bytes32, label indexed bytes32, owner address)
func (_ENS *ENSFilterer) FilterNewOwner(opts *bind.FilterOpts, node [][32]byte, label [][32]byte) (*ENSNewOwnerIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return &ENSNewOwnerIterator{contract: _ENS.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0xce0457fe73731f824cc272376169235128c118b49d344817417c6d108d155e82.
//
// Solidity: e NewOwner(node indexed bytes32, label indexed bytes32, owner address)
func (_ENS *ENSFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *ENSNewOwner, node [][32]byte, label [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "NewOwner", nodeRule, labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSNewOwner)
				if err := _ENS.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ENSNewResolverIterator is returned from FilterNewResolver and is used to iterate over the raw logs and unpacked data for NewResolver events raised by the ENS contract.
type ENSNewResolverIterator struct {
	Event *ENSNewResolver // Event containing the contract specifics and raw log

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
func (it *ENSNewResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSNewResolver)
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
		it.Event = new(ENSNewResolver)
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
func (it *ENSNewResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSNewResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSNewResolver represents a NewResolver event raised by the ENS contract.
type ENSNewResolver struct {
	Node     [32]byte
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewResolver is a free log retrieval operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: e NewResolver(node indexed bytes32, resolver address)
func (_ENS *ENSFilterer) FilterNewResolver(opts *bind.FilterOpts, node [][32]byte) (*ENSNewResolverIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSNewResolverIterator{contract: _ENS.contract, event: "NewResolver", logs: logs, sub: sub}, nil
}

// WatchNewResolver is a free log subscription operation binding the contract event 0x335721b01866dc23fbee8b6b2c7b1e14d6f05c28cd35a2c934239f94095602a0.
//
// Solidity: e NewResolver(node indexed bytes32, resolver address)
func (_ENS *ENSFilterer) WatchNewResolver(opts *bind.WatchOpts, sink chan<- *ENSNewResolver, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "NewResolver", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSNewResolver)
				if err := _ENS.contract.UnpackLog(event, "NewResolver", log); err != nil {
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

// ENSTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ENS contract.
type ENSTransferIterator struct {
	Event *ENSTransfer // Event containing the contract specifics and raw log

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
func (it *ENSTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ENSTransfer)
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
		it.Event = new(ENSTransfer)
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
func (it *ENSTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ENSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ENSTransfer represents a Transfer event raised by the ENS contract.
type ENSTransfer struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: e Transfer(node indexed bytes32, owner address)
func (_ENS *ENSFilterer) FilterTransfer(opts *bind.FilterOpts, node [][32]byte) (*ENSTransferIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.FilterLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ENSTransferIterator{contract: _ENS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd4735d920b0f87494915f556dd9b54c8f309026070caea5c737245152564d266.
//
// Solidity: e Transfer(node indexed bytes32, owner address)
func (_ENS *ENSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ENSTransfer, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _ENS.contract.WatchLogs(opts, "Transfer", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ENSTransfer)
				if err := _ENS.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// FIFSRegistrarABI is the input ABI used to generate the binding from.
const FIFSRegistrarABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"subnode\",\"type\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ensAddr\",\"type\":\"address\"},{\"name\":\"node\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// FIFSRegistrarBin is the compiled bytecode used for deploying new contracts.
const FIFSRegistrarBin = `0x608060405234801561001057600080fd5b5060405160408061087283398101604052805160209091015160008054600160a060020a031916600160a060020a0384161790558161004d6100a4565b600160a060020a03909116815260405190819003602001906000f08015801561007a573d6000803e3d6000fd5b5060018054600160a060020a031916600160a060020a0392909216919091179055600255506100b4565b6040516104ac806103c683390190565b610303806100c36000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663d22057a98114610045575b600080fd5b34801561005157600080fd5b50610069600435600160a060020a036024351661006b565b005b60025460408051918252602080830185905281519283900382018320600080547f02571be300000000000000000000000000000000000000000000000000000000865260048601839052935191949093600160a060020a0316926302571be39260248084019382900301818787803b1580156100e657600080fd5b505af11580156100fa573d6000803e3d6000fd5b505050506040513d602081101561011057600080fd5b50519050600160a060020a038116158015906101355750600160a060020a0381163314155b1561013f57600080fd5b60008054600254604080517f06ab592300000000000000000000000000000000000000000000000000000000815260048101929092526024820188905230604483015251600160a060020a03909216926306ab59239260648084019382900301818387803b1580156101b057600080fd5b505af11580156101c4573d6000803e3d6000fd5b505060008054600154604080517f1896f70a00000000000000000000000000000000000000000000000000000000815260048101899052600160a060020a0392831660248201529051919092169450631896f70a93506044808301939282900301818387803b15801561023657600080fd5b505af115801561024a573d6000803e3d6000fd5b505060008054604080517f5b0fc9c300000000000000000000000000000000000000000000000000000000815260048101889052600160a060020a0389811660248301529151919092169450635b0fc9c393506044808301939282900301818387803b1580156102b957600080fd5b505af11580156102cd573d6000803e3d6000fd5b50505050505050505600a165627a7a72305820ae55b028e60774a969020d5f3463d7dc2cb5315c985ebef83412a1e0a1d386c40029608060405234801561001057600080fd5b506040516020806104ac833981016040525160008054600160a060020a03909216600160a060020a031990921691909117905561045a806100526000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632dff6941811461007e5780633b3b57de146100a857806341b9dc2b146100dc578063c3d014d61461010b578063d5fa2b0014610128575b34801561007857600080fd5b50600080fd5b34801561008a57600080fd5b5061009660043561014c565b60408051918252519081900360200190f35b3480156100b457600080fd5b506100c060043561016c565b60408051600160a060020a039092168252519081900360200190f35b3480156100e857600080fd5b506100f7600435602435610190565b604080519115158252519081900360200190f35b34801561011757600080fd5b50610126600435602435610220565b005b34801561013457600080fd5b50610126600435600160a060020a0360243516610314565b60008181526002602052604090205480151561016757600080fd5b919050565b600081815260016020526040902054600160a060020a031680151561016757600080fd5b60007f6164647200000000000000000000000000000000000000000000000000000000821480156101d75750600083815260016020526040902054600160a060020a031615155b8061021957507f636f6e74656e740000000000000000000000000000000000000000000000000082148015610219575060008381526002602052604090205415155b9392505050565b60008054604080517f02571be300000000000000000000000000000000000000000000000000000000815260048101869052905185933393600160a060020a0316926302571be39260248083019360209383900390910190829087803b15801561028957600080fd5b505af115801561029d573d6000803e3d6000fd5b505050506040513d60208110156102b357600080fd5b5051600160a060020a0316146102c857600080fd5b6000838152600260209081526040918290208490558151848152915185927f0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc92908290030190a2505050565b60008054604080517f02571be300000000000000000000000000000000000000000000000000000000815260048101869052905185933393600160a060020a0316926302571be39260248083019360209383900390910190829087803b15801561037d57600080fd5b505af1158015610391573d6000803e3d6000fd5b505050506040513d60208110156103a757600080fd5b5051600160a060020a0316146103bc57600080fd5b600083815260016020908152604091829020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386169081179091558251908152915185927f52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd292908290030190a25050505600a165627a7a72305820e828d8c34d5c1fa2d5ef031c0663cb888b98ce4d9d76f94612232dbd8499833c0029`

// DeployFIFSRegistrar deploys a new Ethereum contract, binding an instance of FIFSRegistrar to it.
func DeployFIFSRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend, ensAddr common.Address, node [32]byte) (common.Address, *types.Transaction, *FIFSRegistrar, error) {
	parsed, err := abi.JSON(strings.NewReader(FIFSRegistrarABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FIFSRegistrarBin), backend, ensAddr, node)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FIFSRegistrar{FIFSRegistrarCaller: FIFSRegistrarCaller{contract: contract}, FIFSRegistrarTransactor: FIFSRegistrarTransactor{contract: contract}, FIFSRegistrarFilterer: FIFSRegistrarFilterer{contract: contract}}, nil
}

// FIFSRegistrar is an auto generated Go binding around an Ethereum contract.
type FIFSRegistrar struct {
	FIFSRegistrarCaller     // Read-only binding to the contract
	FIFSRegistrarTransactor // Write-only binding to the contract
	FIFSRegistrarFilterer   // Log filterer for contract events
}

// FIFSRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type FIFSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FIFSRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FIFSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FIFSRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FIFSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FIFSRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FIFSRegistrarSession struct {
	Contract     *FIFSRegistrar    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FIFSRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FIFSRegistrarCallerSession struct {
	Contract *FIFSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FIFSRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FIFSRegistrarTransactorSession struct {
	Contract     *FIFSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FIFSRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type FIFSRegistrarRaw struct {
	Contract *FIFSRegistrar // Generic contract binding to access the raw methods on
}

// FIFSRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FIFSRegistrarCallerRaw struct {
	Contract *FIFSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// FIFSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FIFSRegistrarTransactorRaw struct {
	Contract *FIFSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFIFSRegistrar creates a new instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrar(address common.Address, backend bind.ContractBackend) (*FIFSRegistrar, error) {
	contract, err := bindFIFSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrar{FIFSRegistrarCaller: FIFSRegistrarCaller{contract: contract}, FIFSRegistrarTransactor: FIFSRegistrarTransactor{contract: contract}, FIFSRegistrarFilterer: FIFSRegistrarFilterer{contract: contract}}, nil
}

// NewFIFSRegistrarCaller creates a new read-only instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*FIFSRegistrarCaller, error) {
	contract, err := bindFIFSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrarCaller{contract: contract}, nil
}

// NewFIFSRegistrarTransactor creates a new write-only instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*FIFSRegistrarTransactor, error) {
	contract, err := bindFIFSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrarTransactor{contract: contract}, nil
}

// NewFIFSRegistrarFilterer creates a new log filterer instance of FIFSRegistrar, bound to a specific deployed contract.
func NewFIFSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*FIFSRegistrarFilterer, error) {
	contract, err := bindFIFSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FIFSRegistrarFilterer{contract: contract}, nil
}

// bindFIFSRegistrar binds a generic wrapper to an already deployed contract.
func bindFIFSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FIFSRegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FIFSRegistrar *FIFSRegistrarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FIFSRegistrar.Contract.FIFSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FIFSRegistrar *FIFSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.FIFSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FIFSRegistrar *FIFSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.FIFSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FIFSRegistrar *FIFSRegistrarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FIFSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FIFSRegistrar *FIFSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FIFSRegistrar *FIFSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// Register is a paid mutator transaction binding the contract method 0xd22057a9.
//
// Solidity: function register(subnode bytes32, owner address) returns()
func (_FIFSRegistrar *FIFSRegistrarTransactor) Register(opts *bind.TransactOpts, subnode [32]byte, owner common.Address) (*types.Transaction, error) {
	return _FIFSRegistrar.contract.Transact(opts, "register", subnode, owner)
}

// Register is a paid mutator transaction binding the contract method 0xd22057a9.
//
// Solidity: function register(subnode bytes32, owner address) returns()
func (_FIFSRegistrar *FIFSRegistrarSession) Register(subnode [32]byte, owner common.Address) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.Register(&_FIFSRegistrar.TransactOpts, subnode, owner)
}

// Register is a paid mutator transaction binding the contract method 0xd22057a9.
//
// Solidity: function register(subnode bytes32, owner address) returns()
func (_FIFSRegistrar *FIFSRegistrarTransactorSession) Register(subnode [32]byte, owner common.Address) (*types.Transaction, error) {
	return _FIFSRegistrar.Contract.Register(&_FIFSRegistrar.TransactOpts, subnode, owner)
}

// PublicResolverABI is the input ABI used to generate the binding from.
const PublicResolverABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"ret\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"kind\",\"type\":\"bytes32\"}],\"name\":\"has\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"setContent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ensAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ContentChanged\",\"type\":\"event\"}]"

// PublicResolverBin is the compiled bytecode used for deploying new contracts.
const PublicResolverBin = `0x608060405234801561001057600080fd5b506040516020806104ac833981016040525160008054600160a060020a03909216600160a060020a031990921691909117905561045a806100526000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632dff6941811461007e5780633b3b57de146100a857806341b9dc2b146100dc578063c3d014d61461010b578063d5fa2b0014610128575b34801561007857600080fd5b50600080fd5b34801561008a57600080fd5b5061009660043561014c565b60408051918252519081900360200190f35b3480156100b457600080fd5b506100c060043561016c565b60408051600160a060020a039092168252519081900360200190f35b3480156100e857600080fd5b506100f7600435602435610190565b604080519115158252519081900360200190f35b34801561011757600080fd5b50610126600435602435610220565b005b34801561013457600080fd5b50610126600435600160a060020a0360243516610314565b60008181526002602052604090205480151561016757600080fd5b919050565b600081815260016020526040902054600160a060020a031680151561016757600080fd5b60007f6164647200000000000000000000000000000000000000000000000000000000821480156101d75750600083815260016020526040902054600160a060020a031615155b8061021957507f636f6e74656e740000000000000000000000000000000000000000000000000082148015610219575060008381526002602052604090205415155b9392505050565b60008054604080517f02571be300000000000000000000000000000000000000000000000000000000815260048101869052905185933393600160a060020a0316926302571be39260248083019360209383900390910190829087803b15801561028957600080fd5b505af115801561029d573d6000803e3d6000fd5b505050506040513d60208110156102b357600080fd5b5051600160a060020a0316146102c857600080fd5b6000838152600260209081526040918290208490558151848152915185927f0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc92908290030190a2505050565b60008054604080517f02571be300000000000000000000000000000000000000000000000000000000815260048101869052905185933393600160a060020a0316926302571be39260248083019360209383900390910190829087803b15801561037d57600080fd5b505af1158015610391573d6000803e3d6000fd5b505050506040513d60208110156103a757600080fd5b5051600160a060020a0316146103bc57600080fd5b600083815260016020908152604091829020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386169081179091558251908152915185927f52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd292908290030190a25050505600a165627a7a72305820e828d8c34d5c1fa2d5ef031c0663cb888b98ce4d9d76f94612232dbd8499833c0029`

// DeployPublicResolver deploys a new Ethereum contract, binding an instance of PublicResolver to it.
func DeployPublicResolver(auth *bind.TransactOpts, backend bind.ContractBackend, ensAddr common.Address) (common.Address, *types.Transaction, *PublicResolver, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicResolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PublicResolverBin), backend, ensAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicResolver{PublicResolverCaller: PublicResolverCaller{contract: contract}, PublicResolverTransactor: PublicResolverTransactor{contract: contract}, PublicResolverFilterer: PublicResolverFilterer{contract: contract}}, nil
}

// PublicResolver is an auto generated Go binding around an Ethereum contract.
type PublicResolver struct {
	PublicResolverCaller     // Read-only binding to the contract
	PublicResolverTransactor // Write-only binding to the contract
	PublicResolverFilterer   // Log filterer for contract events
}

// PublicResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicResolverSession struct {
	Contract     *PublicResolver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PublicResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicResolverCallerSession struct {
	Contract *PublicResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PublicResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicResolverTransactorSession struct {
	Contract     *PublicResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PublicResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicResolverRaw struct {
	Contract *PublicResolver // Generic contract binding to access the raw methods on
}

// PublicResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicResolverCallerRaw struct {
	Contract *PublicResolverCaller // Generic read-only contract binding to access the raw methods on
}

// PublicResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicResolverTransactorRaw struct {
	Contract *PublicResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicResolver creates a new instance of PublicResolver, bound to a specific deployed contract.
func NewPublicResolver(address common.Address, backend bind.ContractBackend) (*PublicResolver, error) {
	contract, err := bindPublicResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicResolver{PublicResolverCaller: PublicResolverCaller{contract: contract}, PublicResolverTransactor: PublicResolverTransactor{contract: contract}, PublicResolverFilterer: PublicResolverFilterer{contract: contract}}, nil
}

// NewPublicResolverCaller creates a new read-only instance of PublicResolver, bound to a specific deployed contract.
func NewPublicResolverCaller(address common.Address, caller bind.ContractCaller) (*PublicResolverCaller, error) {
	contract, err := bindPublicResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicResolverCaller{contract: contract}, nil
}

// NewPublicResolverTransactor creates a new write-only instance of PublicResolver, bound to a specific deployed contract.
func NewPublicResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicResolverTransactor, error) {
	contract, err := bindPublicResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicResolverTransactor{contract: contract}, nil
}

// NewPublicResolverFilterer creates a new log filterer instance of PublicResolver, bound to a specific deployed contract.
func NewPublicResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicResolverFilterer, error) {
	contract, err := bindPublicResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicResolverFilterer{contract: contract}, nil
}

// bindPublicResolver binds a generic wrapper to an already deployed contract.
func bindPublicResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PublicResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicResolver *PublicResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PublicResolver.Contract.PublicResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicResolver *PublicResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicResolver.Contract.PublicResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicResolver *PublicResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicResolver.Contract.PublicResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicResolver *PublicResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PublicResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicResolver *PublicResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicResolver *PublicResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicResolver.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_PublicResolver *PublicResolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PublicResolver.contract.Call(opts, out, "addr", node)
	return *ret0, err
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_PublicResolver *PublicResolverSession) Addr(node [32]byte) (common.Address, error) {
	return _PublicResolver.Contract.Addr(&_PublicResolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_PublicResolver *PublicResolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _PublicResolver.Contract.Addr(&_PublicResolver.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_PublicResolver *PublicResolverCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PublicResolver.contract.Call(opts, out, "content", node)
	return *ret0, err
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_PublicResolver *PublicResolverSession) Content(node [32]byte) ([32]byte, error) {
	return _PublicResolver.Contract.Content(&_PublicResolver.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_PublicResolver *PublicResolverCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _PublicResolver.Contract.Content(&_PublicResolver.CallOpts, node)
}

// Has is a paid mutator transaction binding the contract method 0x41b9dc2b.
//
// Solidity: function has(node bytes32, kind bytes32) returns(bool)
func (_PublicResolver *PublicResolverTransactor) Has(opts *bind.TransactOpts, node [32]byte, kind [32]byte) (*types.Transaction, error) {
	return _PublicResolver.contract.Transact(opts, "has", node, kind)
}

// Has is a paid mutator transaction binding the contract method 0x41b9dc2b.
//
// Solidity: function has(node bytes32, kind bytes32) returns(bool)
func (_PublicResolver *PublicResolverSession) Has(node [32]byte, kind [32]byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.Has(&_PublicResolver.TransactOpts, node, kind)
}

// Has is a paid mutator transaction binding the contract method 0x41b9dc2b.
//
// Solidity: function has(node bytes32, kind bytes32) returns(bool)
func (_PublicResolver *PublicResolverTransactorSession) Has(node [32]byte, kind [32]byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.Has(&_PublicResolver.TransactOpts, node, kind)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_PublicResolver *PublicResolverTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _PublicResolver.contract.Transact(opts, "setAddr", node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_PublicResolver *PublicResolverSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetAddr(&_PublicResolver.TransactOpts, node, addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, addr address) returns()
func (_PublicResolver *PublicResolverTransactorSession) SetAddr(node [32]byte, addr common.Address) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetAddr(&_PublicResolver.TransactOpts, node, addr)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_PublicResolver *PublicResolverTransactor) SetContent(opts *bind.TransactOpts, node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _PublicResolver.contract.Transact(opts, "setContent", node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_PublicResolver *PublicResolverSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetContent(&_PublicResolver.TransactOpts, node, hash)
}

// SetContent is a paid mutator transaction binding the contract method 0xc3d014d6.
//
// Solidity: function setContent(node bytes32, hash bytes32) returns()
func (_PublicResolver *PublicResolverTransactorSession) SetContent(node [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _PublicResolver.Contract.SetContent(&_PublicResolver.TransactOpts, node, hash)
}

// PublicResolverAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the PublicResolver contract.
type PublicResolverAddrChangedIterator struct {
	Event *PublicResolverAddrChanged // Event containing the contract specifics and raw log

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
func (it *PublicResolverAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicResolverAddrChanged)
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
		it.Event = new(PublicResolverAddrChanged)
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
func (it *PublicResolverAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicResolverAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicResolverAddrChanged represents a AddrChanged event raised by the PublicResolver contract.
type PublicResolverAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: e AddrChanged(node indexed bytes32, a address)
func (_PublicResolver *PublicResolverFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicResolverAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicResolverAddrChangedIterator{contract: _PublicResolver.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: e AddrChanged(node indexed bytes32, a address)
func (_PublicResolver *PublicResolverFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *PublicResolverAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicResolverAddrChanged)
				if err := _PublicResolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// PublicResolverContentChangedIterator is returned from FilterContentChanged and is used to iterate over the raw logs and unpacked data for ContentChanged events raised by the PublicResolver contract.
type PublicResolverContentChangedIterator struct {
	Event *PublicResolverContentChanged // Event containing the contract specifics and raw log

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
func (it *PublicResolverContentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicResolverContentChanged)
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
		it.Event = new(PublicResolverContentChanged)
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
func (it *PublicResolverContentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicResolverContentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicResolverContentChanged represents a ContentChanged event raised by the PublicResolver contract.
type PublicResolverContentChanged struct {
	Node [32]byte
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContentChanged is a free log retrieval operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: e ContentChanged(node indexed bytes32, hash bytes32)
func (_PublicResolver *PublicResolverFilterer) FilterContentChanged(opts *bind.FilterOpts, node [][32]byte) (*PublicResolverContentChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.FilterLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &PublicResolverContentChangedIterator{contract: _PublicResolver.contract, event: "ContentChanged", logs: logs, sub: sub}, nil
}

// WatchContentChanged is a free log subscription operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: e ContentChanged(node indexed bytes32, hash bytes32)
func (_PublicResolver *PublicResolverFilterer) WatchContentChanged(opts *bind.WatchOpts, sink chan<- *PublicResolverContentChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _PublicResolver.contract.WatchLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicResolverContentChanged)
				if err := _PublicResolver.contract.UnpackLog(event, "ContentChanged", log); err != nil {
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

// ResolverABI is the input ABI used to generate the binding from.
const ResolverABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"content\",\"outputs\":[{\"name\":\"ret\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"ret\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"kind\",\"type\":\"bytes32\"}],\"name\":\"has\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ContentChanged\",\"type\":\"event\"}]"

// ResolverBin is the compiled bytecode used for deploying new contracts.
const ResolverBin = `0x`

// DeployResolver deploys a new Ethereum contract, binding an instance of Resolver to it.
func DeployResolver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Resolver, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ResolverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Resolver{ResolverCaller: ResolverCaller{contract: contract}, ResolverTransactor: ResolverTransactor{contract: contract}, ResolverFilterer: ResolverFilterer{contract: contract}}, nil
}

// Resolver is an auto generated Go binding around an Ethereum contract.
type Resolver struct {
	ResolverCaller     // Read-only binding to the contract
	ResolverTransactor // Write-only binding to the contract
	ResolverFilterer   // Log filterer for contract events
}

// ResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolverSession struct {
	Contract     *Resolver         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolverCallerSession struct {
	Contract *ResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolverTransactorSession struct {
	Contract     *ResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolverRaw struct {
	Contract *Resolver // Generic contract binding to access the raw methods on
}

// ResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolverCallerRaw struct {
	Contract *ResolverCaller // Generic read-only contract binding to access the raw methods on
}

// ResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolverTransactorRaw struct {
	Contract *ResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolver creates a new instance of Resolver, bound to a specific deployed contract.
func NewResolver(address common.Address, backend bind.ContractBackend) (*Resolver, error) {
	contract, err := bindResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Resolver{ResolverCaller: ResolverCaller{contract: contract}, ResolverTransactor: ResolverTransactor{contract: contract}, ResolverFilterer: ResolverFilterer{contract: contract}}, nil
}

// NewResolverCaller creates a new read-only instance of Resolver, bound to a specific deployed contract.
func NewResolverCaller(address common.Address, caller bind.ContractCaller) (*ResolverCaller, error) {
	contract, err := bindResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverCaller{contract: contract}, nil
}

// NewResolverTransactor creates a new write-only instance of Resolver, bound to a specific deployed contract.
func NewResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolverTransactor, error) {
	contract, err := bindResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolverTransactor{contract: contract}, nil
}

// NewResolverFilterer creates a new log filterer instance of Resolver, bound to a specific deployed contract.
func NewResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*ResolverFilterer, error) {
	contract, err := bindResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolverFilterer{contract: contract}, nil
}

// bindResolver binds a generic wrapper to an already deployed contract.
func bindResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolver *ResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Resolver.Contract.ResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolver *ResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolver.Contract.ResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolver *ResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolver.Contract.ResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Resolver *ResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Resolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Resolver *ResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Resolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Resolver *ResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Resolver.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_Resolver *ResolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Resolver.contract.Call(opts, out, "addr", node)
	return *ret0, err
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_Resolver *ResolverSession) Addr(node [32]byte) (common.Address, error) {
	return _Resolver.Contract.Addr(&_Resolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(ret address)
func (_Resolver *ResolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _Resolver.Contract.Addr(&_Resolver.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_Resolver *ResolverCaller) Content(opts *bind.CallOpts, node [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Resolver.contract.Call(opts, out, "content", node)
	return *ret0, err
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_Resolver *ResolverSession) Content(node [32]byte) ([32]byte, error) {
	return _Resolver.Contract.Content(&_Resolver.CallOpts, node)
}

// Content is a free data retrieval call binding the contract method 0x2dff6941.
//
// Solidity: function content(node bytes32) constant returns(ret bytes32)
func (_Resolver *ResolverCallerSession) Content(node [32]byte) ([32]byte, error) {
	return _Resolver.Contract.Content(&_Resolver.CallOpts, node)
}

// Has is a paid mutator transaction binding the contract method 0x41b9dc2b.
//
// Solidity: function has(node bytes32, kind bytes32) returns(bool)
func (_Resolver *ResolverTransactor) Has(opts *bind.TransactOpts, node [32]byte, kind [32]byte) (*types.Transaction, error) {
	return _Resolver.contract.Transact(opts, "has", node, kind)
}

// Has is a paid mutator transaction binding the contract method 0x41b9dc2b.
//
// Solidity: function has(node bytes32, kind bytes32) returns(bool)
func (_Resolver *ResolverSession) Has(node [32]byte, kind [32]byte) (*types.Transaction, error) {
	return _Resolver.Contract.Has(&_Resolver.TransactOpts, node, kind)
}

// Has is a paid mutator transaction binding the contract method 0x41b9dc2b.
//
// Solidity: function has(node bytes32, kind bytes32) returns(bool)
func (_Resolver *ResolverTransactorSession) Has(node [32]byte, kind [32]byte) (*types.Transaction, error) {
	return _Resolver.Contract.Has(&_Resolver.TransactOpts, node, kind)
}

// ResolverAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the Resolver contract.
type ResolverAddrChangedIterator struct {
	Event *ResolverAddrChanged // Event containing the contract specifics and raw log

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
func (it *ResolverAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverAddrChanged)
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
		it.Event = new(ResolverAddrChanged)
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
func (it *ResolverAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverAddrChanged represents a AddrChanged event raised by the Resolver contract.
type ResolverAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: e AddrChanged(node indexed bytes32, a address)
func (_Resolver *ResolverFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverAddrChangedIterator{contract: _Resolver.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: e AddrChanged(node indexed bytes32, a address)
func (_Resolver *ResolverFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *ResolverAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverAddrChanged)
				if err := _Resolver.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ResolverContentChangedIterator is returned from FilterContentChanged and is used to iterate over the raw logs and unpacked data for ContentChanged events raised by the Resolver contract.
type ResolverContentChangedIterator struct {
	Event *ResolverContentChanged // Event containing the contract specifics and raw log

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
func (it *ResolverContentChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolverContentChanged)
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
		it.Event = new(ResolverContentChanged)
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
func (it *ResolverContentChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolverContentChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolverContentChanged represents a ContentChanged event raised by the Resolver contract.
type ResolverContentChanged struct {
	Node [32]byte
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContentChanged is a free log retrieval operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: e ContentChanged(node indexed bytes32, hash bytes32)
func (_Resolver *ResolverFilterer) FilterContentChanged(opts *bind.FilterOpts, node [][32]byte) (*ResolverContentChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.FilterLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ResolverContentChangedIterator{contract: _Resolver.contract, event: "ContentChanged", logs: logs, sub: sub}, nil
}

// WatchContentChanged is a free log subscription operation binding the contract event 0x0424b6fe0d9c3bdbece0e7879dc241bb0c22e900be8b6c168b4ee08bd9bf83bc.
//
// Solidity: e ContentChanged(node indexed bytes32, hash bytes32)
func (_Resolver *ResolverFilterer) WatchContentChanged(opts *bind.WatchOpts, sink chan<- *ResolverContentChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Resolver.contract.WatchLogs(opts, "ContentChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolverContentChanged)
				if err := _Resolver.contract.UnpackLog(event, "ContentChanged", log); err != nil {
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
