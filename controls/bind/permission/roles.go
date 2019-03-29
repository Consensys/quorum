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

// RoleManagerABI is the input ABI used to generate the binding from.
const RoleManagerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getRoleDetails\",\"outputs\":[{\"name\":\"roleId\",\"type\":\"string\"},{\"name\":\"orgId\",\"type\":\"string\"},{\"name\":\"accessType\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"bool\"},{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"isVoterRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"isFullAccessRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_baseAccess\",\"type\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"bool\"}],\"name\":\"addRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"roleExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfRoles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"rIndex\",\"type\":\"uint256\"}],\"name\":\"getRoleDetailsFromIndex\",\"outputs\":[{\"name\":\"roleId\",\"type\":\"string\"},{\"name\":\"orgId\",\"type\":\"string\"},{\"name\":\"accessType\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"bool\"},{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_roleId\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"removeRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_baseAccess\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_isVoter\",\"type\":\"bool\"}],\"name\":\"RoleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_roleId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"}]"

// RoleManager is an auto generated Go binding around an Ethereum contract.
type RoleManager struct {
	RoleManagerCaller     // Read-only binding to the contract
	RoleManagerTransactor // Write-only binding to the contract
	RoleManagerFilterer   // Log filterer for contract events
}

// RoleManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoleManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoleManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoleManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoleManagerSession struct {
	Contract     *RoleManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoleManagerCallerSession struct {
	Contract *RoleManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RoleManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoleManagerTransactorSession struct {
	Contract     *RoleManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RoleManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoleManagerRaw struct {
	Contract *RoleManager // Generic contract binding to access the raw methods on
}

// RoleManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoleManagerCallerRaw struct {
	Contract *RoleManagerCaller // Generic read-only contract binding to access the raw methods on
}

// RoleManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoleManagerTransactorRaw struct {
	Contract *RoleManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoleManager creates a new instance of RoleManager, bound to a specific deployed contract.
func NewRoleManager(address common.Address, backend bind.ContractBackend) (*RoleManager, error) {
	contract, err := bindRoleManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoleManager{RoleManagerCaller: RoleManagerCaller{contract: contract}, RoleManagerTransactor: RoleManagerTransactor{contract: contract}, RoleManagerFilterer: RoleManagerFilterer{contract: contract}}, nil
}

// NewRoleManagerCaller creates a new read-only instance of RoleManager, bound to a specific deployed contract.
func NewRoleManagerCaller(address common.Address, caller bind.ContractCaller) (*RoleManagerCaller, error) {
	contract, err := bindRoleManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoleManagerCaller{contract: contract}, nil
}

// NewRoleManagerTransactor creates a new write-only instance of RoleManager, bound to a specific deployed contract.
func NewRoleManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*RoleManagerTransactor, error) {
	contract, err := bindRoleManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoleManagerTransactor{contract: contract}, nil
}

// NewRoleManagerFilterer creates a new log filterer instance of RoleManager, bound to a specific deployed contract.
func NewRoleManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*RoleManagerFilterer, error) {
	contract, err := bindRoleManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoleManagerFilterer{contract: contract}, nil
}

// bindRoleManager binds a generic wrapper to an already deployed contract.
func bindRoleManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoleManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleManager *RoleManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RoleManager.Contract.RoleManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleManager *RoleManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleManager.Contract.RoleManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleManager *RoleManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleManager.Contract.RoleManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleManager *RoleManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RoleManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleManager *RoleManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleManager *RoleManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleManager.Contract.contract.Transact(opts, method, params...)
}

// GetNumberOfRoles is a free data retrieval call binding the contract method 0x87f55d31.
//
// Solidity: function getNumberOfRoles() constant returns(uint256)
func (_RoleManager *RoleManagerCaller) GetNumberOfRoles(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RoleManager.contract.Call(opts, out, "getNumberOfRoles")
	return *ret0, err
}

// GetNumberOfRoles is a free data retrieval call binding the contract method 0x87f55d31.
//
// Solidity: function getNumberOfRoles() constant returns(uint256)
func (_RoleManager *RoleManagerSession) GetNumberOfRoles() (*big.Int, error) {
	return _RoleManager.Contract.GetNumberOfRoles(&_RoleManager.CallOpts)
}

// GetNumberOfRoles is a free data retrieval call binding the contract method 0x87f55d31.
//
// Solidity: function getNumberOfRoles() constant returns(uint256)
func (_RoleManager *RoleManagerCallerSession) GetNumberOfRoles() (*big.Int, error) {
	return _RoleManager.Contract.GetNumberOfRoles(&_RoleManager.CallOpts)
}

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(_roleId string, _orgId string) constant returns(roleId string, orgId string, accessType uint256, voter bool, active bool)
func (_RoleManager *RoleManagerCaller) GetRoleDetails(opts *bind.CallOpts, _roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Active     bool
}, error) {
	ret := new(struct {
		RoleId     string
		OrgId      string
		AccessType *big.Int
		Voter      bool
		Active     bool
	})
	out := ret
	err := _RoleManager.contract.Call(opts, out, "getRoleDetails", _roleId, _orgId)
	return *ret, err
}

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(_roleId string, _orgId string) constant returns(roleId string, orgId string, accessType uint256, voter bool, active bool)
func (_RoleManager *RoleManagerSession) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Active     bool
}, error) {
	return _RoleManager.Contract.GetRoleDetails(&_RoleManager.CallOpts, _roleId, _orgId)
}

// GetRoleDetails is a free data retrieval call binding the contract method 0x1870aba3.
//
// Solidity: function getRoleDetails(_roleId string, _orgId string) constant returns(roleId string, orgId string, accessType uint256, voter bool, active bool)
func (_RoleManager *RoleManagerCallerSession) GetRoleDetails(_roleId string, _orgId string) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Active     bool
}, error) {
	return _RoleManager.Contract.GetRoleDetails(&_RoleManager.CallOpts, _roleId, _orgId)
}

// GetRoleDetailsFromIndex is a free data retrieval call binding the contract method 0xa451d4a8.
//
// Solidity: function getRoleDetailsFromIndex(rIndex uint256) constant returns(roleId string, orgId string, accessType uint256, voter bool, active bool)
func (_RoleManager *RoleManagerCaller) GetRoleDetailsFromIndex(opts *bind.CallOpts, rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Active     bool
}, error) {
	ret := new(struct {
		RoleId     string
		OrgId      string
		AccessType *big.Int
		Voter      bool
		Active     bool
	})
	out := ret
	err := _RoleManager.contract.Call(opts, out, "getRoleDetailsFromIndex", rIndex)
	return *ret, err
}

// GetRoleDetailsFromIndex is a free data retrieval call binding the contract method 0xa451d4a8.
//
// Solidity: function getRoleDetailsFromIndex(rIndex uint256) constant returns(roleId string, orgId string, accessType uint256, voter bool, active bool)
func (_RoleManager *RoleManagerSession) GetRoleDetailsFromIndex(rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Active     bool
}, error) {
	return _RoleManager.Contract.GetRoleDetailsFromIndex(&_RoleManager.CallOpts, rIndex)
}

// GetRoleDetailsFromIndex is a free data retrieval call binding the contract method 0xa451d4a8.
//
// Solidity: function getRoleDetailsFromIndex(rIndex uint256) constant returns(roleId string, orgId string, accessType uint256, voter bool, active bool)
func (_RoleManager *RoleManagerCallerSession) GetRoleDetailsFromIndex(rIndex *big.Int) (struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Active     bool
}, error) {
	return _RoleManager.Contract.GetRoleDetailsFromIndex(&_RoleManager.CallOpts, rIndex)
}

// IsFullAccessRole is a free data retrieval call binding the contract method 0x476ff5cc.
//
// Solidity: function isFullAccessRole(_roleId string, _orgId string) constant returns(bool)
func (_RoleManager *RoleManagerCaller) IsFullAccessRole(opts *bind.CallOpts, _roleId string, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RoleManager.contract.Call(opts, out, "isFullAccessRole", _roleId, _orgId)
	return *ret0, err
}

// IsFullAccessRole is a free data retrieval call binding the contract method 0x476ff5cc.
//
// Solidity: function isFullAccessRole(_roleId string, _orgId string) constant returns(bool)
func (_RoleManager *RoleManagerSession) IsFullAccessRole(_roleId string, _orgId string) (bool, error) {
	return _RoleManager.Contract.IsFullAccessRole(&_RoleManager.CallOpts, _roleId, _orgId)
}

// IsFullAccessRole is a free data retrieval call binding the contract method 0x476ff5cc.
//
// Solidity: function isFullAccessRole(_roleId string, _orgId string) constant returns(bool)
func (_RoleManager *RoleManagerCallerSession) IsFullAccessRole(_roleId string, _orgId string) (bool, error) {
	return _RoleManager.Contract.IsFullAccessRole(&_RoleManager.CallOpts, _roleId, _orgId)
}

// IsVoterRole is a free data retrieval call binding the contract method 0x2b113705.
//
// Solidity: function isVoterRole(_roleId string, _orgId string) constant returns(bool)
func (_RoleManager *RoleManagerCaller) IsVoterRole(opts *bind.CallOpts, _roleId string, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RoleManager.contract.Call(opts, out, "isVoterRole", _roleId, _orgId)
	return *ret0, err
}

// IsVoterRole is a free data retrieval call binding the contract method 0x2b113705.
//
// Solidity: function isVoterRole(_roleId string, _orgId string) constant returns(bool)
func (_RoleManager *RoleManagerSession) IsVoterRole(_roleId string, _orgId string) (bool, error) {
	return _RoleManager.Contract.IsVoterRole(&_RoleManager.CallOpts, _roleId, _orgId)
}

// IsVoterRole is a free data retrieval call binding the contract method 0x2b113705.
//
// Solidity: function isVoterRole(_roleId string, _orgId string) constant returns(bool)
func (_RoleManager *RoleManagerCallerSession) IsVoterRole(_roleId string, _orgId string) (bool, error) {
	return _RoleManager.Contract.IsVoterRole(&_RoleManager.CallOpts, _roleId, _orgId)
}

// RoleExists is a free data retrieval call binding the contract method 0x67950aab.
//
// Solidity: function roleExists(_roleId string, _orgId string) constant returns(bool)
func (_RoleManager *RoleManagerCaller) RoleExists(opts *bind.CallOpts, _roleId string, _orgId string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RoleManager.contract.Call(opts, out, "roleExists", _roleId, _orgId)
	return *ret0, err
}

// RoleExists is a free data retrieval call binding the contract method 0x67950aab.
//
// Solidity: function roleExists(_roleId string, _orgId string) constant returns(bool)
func (_RoleManager *RoleManagerSession) RoleExists(_roleId string, _orgId string) (bool, error) {
	return _RoleManager.Contract.RoleExists(&_RoleManager.CallOpts, _roleId, _orgId)
}

// RoleExists is a free data retrieval call binding the contract method 0x67950aab.
//
// Solidity: function roleExists(_roleId string, _orgId string) constant returns(bool)
func (_RoleManager *RoleManagerCallerSession) RoleExists(_roleId string, _orgId string) (bool, error) {
	return _RoleManager.Contract.RoleExists(&_RoleManager.CallOpts, _roleId, _orgId)
}

// AddRole is a paid mutator transaction binding the contract method 0x5ba4d7c5.
//
// Solidity: function addRole(_roleId string, _orgId string, _baseAccess uint256, _voter bool) returns()
func (_RoleManager *RoleManagerTransactor) AddRole(opts *bind.TransactOpts, _roleId string, _orgId string, _baseAccess *big.Int, _voter bool) (*types.Transaction, error) {
	return _RoleManager.contract.Transact(opts, "addRole", _roleId, _orgId, _baseAccess, _voter)
}

// AddRole is a paid mutator transaction binding the contract method 0x5ba4d7c5.
//
// Solidity: function addRole(_roleId string, _orgId string, _baseAccess uint256, _voter bool) returns()
func (_RoleManager *RoleManagerSession) AddRole(_roleId string, _orgId string, _baseAccess *big.Int, _voter bool) (*types.Transaction, error) {
	return _RoleManager.Contract.AddRole(&_RoleManager.TransactOpts, _roleId, _orgId, _baseAccess, _voter)
}

// AddRole is a paid mutator transaction binding the contract method 0x5ba4d7c5.
//
// Solidity: function addRole(_roleId string, _orgId string, _baseAccess uint256, _voter bool) returns()
func (_RoleManager *RoleManagerTransactorSession) AddRole(_roleId string, _orgId string, _baseAccess *big.Int, _voter bool) (*types.Transaction, error) {
	return _RoleManager.Contract.AddRole(&_RoleManager.TransactOpts, _roleId, _orgId, _baseAccess, _voter)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(_roleId string, _orgId string) returns()
func (_RoleManager *RoleManagerTransactor) RemoveRole(opts *bind.TransactOpts, _roleId string, _orgId string) (*types.Transaction, error) {
	return _RoleManager.contract.Transact(opts, "removeRole", _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(_roleId string, _orgId string) returns()
func (_RoleManager *RoleManagerSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _RoleManager.Contract.RemoveRole(&_RoleManager.TransactOpts, _roleId, _orgId)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xa6343012.
//
// Solidity: function removeRole(_roleId string, _orgId string) returns()
func (_RoleManager *RoleManagerTransactorSession) RemoveRole(_roleId string, _orgId string) (*types.Transaction, error) {
	return _RoleManager.Contract.RemoveRole(&_RoleManager.TransactOpts, _roleId, _orgId)
}

// RoleManagerRoleCreatedIterator is returned from FilterRoleCreated and is used to iterate over the raw logs and unpacked data for RoleCreated events raised by the RoleManager contract.
type RoleManagerRoleCreatedIterator struct {
	Event *RoleManagerRoleCreated // Event containing the contract specifics and raw log

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
func (it *RoleManagerRoleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleManagerRoleCreated)
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
		it.Event = new(RoleManagerRoleCreated)
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
func (it *RoleManagerRoleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleManagerRoleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleManagerRoleCreated represents a RoleCreated event raised by the RoleManager contract.
type RoleManagerRoleCreated struct {
	RoleId     string
	OrgId      string
	BaseAccess *big.Int
	IsVoter    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoleCreated is a free log retrieval operation binding the contract event 0x386ac6109c3e45c782fc5c1ad923957645d668ed4197e3173966eb66413e07c6.
//
// Solidity: e RoleCreated(_roleId string, _orgId string, _baseAccess uint256, _isVoter bool)
func (_RoleManager *RoleManagerFilterer) FilterRoleCreated(opts *bind.FilterOpts) (*RoleManagerRoleCreatedIterator, error) {

	logs, sub, err := _RoleManager.contract.FilterLogs(opts, "RoleCreated")
	if err != nil {
		return nil, err
	}
	return &RoleManagerRoleCreatedIterator{contract: _RoleManager.contract, event: "RoleCreated", logs: logs, sub: sub}, nil
}

// WatchRoleCreated is a free log subscription operation binding the contract event 0x386ac6109c3e45c782fc5c1ad923957645d668ed4197e3173966eb66413e07c6.
//
// Solidity: e RoleCreated(_roleId string, _orgId string, _baseAccess uint256, _isVoter bool)
func (_RoleManager *RoleManagerFilterer) WatchRoleCreated(opts *bind.WatchOpts, sink chan<- *RoleManagerRoleCreated) (event.Subscription, error) {

	logs, sub, err := _RoleManager.contract.WatchLogs(opts, "RoleCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleManagerRoleCreated)
				if err := _RoleManager.contract.UnpackLog(event, "RoleCreated", log); err != nil {
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

// RoleManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RoleManager contract.
type RoleManagerRoleRevokedIterator struct {
	Event *RoleManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RoleManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleManagerRoleRevoked)
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
		it.Event = new(RoleManagerRoleRevoked)
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
func (it *RoleManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleManagerRoleRevoked represents a RoleRevoked event raised by the RoleManager contract.
type RoleManagerRoleRevoked struct {
	RoleId string
	OrgId  string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0x1196059dd83524bf989fd94bb65808c09dbea2ab791fb6bfa87a0e0aa64b2ea6.
//
// Solidity: e RoleRevoked(_roleId string, _orgId string)
func (_RoleManager *RoleManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts) (*RoleManagerRoleRevokedIterator, error) {

	logs, sub, err := _RoleManager.contract.FilterLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return &RoleManagerRoleRevokedIterator{contract: _RoleManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0x1196059dd83524bf989fd94bb65808c09dbea2ab791fb6bfa87a0e0aa64b2ea6.
//
// Solidity: e RoleRevoked(_roleId string, _orgId string)
func (_RoleManager *RoleManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RoleManagerRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _RoleManager.contract.WatchLogs(opts, "RoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleManagerRoleRevoked)
				if err := _RoleManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
