// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractWhitelistManagerMetaData contains all meta data concerning the ContractWhitelistManager contract.
var ContractWhitelistManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"ContractWhitelistModified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"addNewContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cIndex\",\"type\":\"uint256\"}],\"name\":\"getContractWhitelistDetailsFromIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfWhitelistedContracts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516108a03803806108a0833981810160405281019061003191906100d4565b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a38261007a565b9050919050565b6100b381610099565b81146100bd575f80fd5b50565b5f815190506100ce816100aa565b92915050565b5f602082840312156100e9576100e8610076565b5b5f6100f6848285016100c0565b91505092915050565b6107948061010c5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80639a7fda6614610043578063b20f4fa51461005f578063b5a93e261461008f575b5f80fd5b61005d600480360381019061005891906104fa565b6100ad565b005b61007960048036038101906100749190610558565b6103f8565b6040516100869190610592565b60405180910390f35b61009761043e565b6040516100a491906105ba565b60405180910390f35b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b8152600401602060405180830381865afa158015610115573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061013991906105e7565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019d9061066c565b60405180910390fd5b60035460025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055507fb354cac8f12f93c803369656a115296370a81eafcfed866fce91e47408be0c2a816040516102199190610592565b60405180910390a15f61022b8261044a565b90505f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054146102d25781600182815481106102865761028561068a565b5b905f5260205f20015f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506103bd565b60035f8154809291906102e4906106e4565b919050555060035460025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550600160405180602001604052808473ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003905f5260205f20015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505b7fb354cac8f12f93c803369656a115296370a81eafcfed866fce91e47408be0c2a826040516103ec9190610592565b60405180910390a15050565b5f6001828154811061040d5761040c61068a565b5b905f5260205f20015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f600180549050905090565b5f600160025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610495919061072b565b9050919050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104c9826104a0565b9050919050565b6104d9816104bf565b81146104e3575f80fd5b50565b5f813590506104f4816104d0565b92915050565b5f6020828403121561050f5761050e61049c565b5b5f61051c848285016104e6565b91505092915050565b5f819050919050565b61053781610525565b8114610541575f80fd5b50565b5f813590506105528161052e565b92915050565b5f6020828403121561056d5761056c61049c565b5b5f61057a84828501610544565b91505092915050565b61058c816104bf565b82525050565b5f6020820190506105a55f830184610583565b92915050565b6105b481610525565b82525050565b5f6020820190506105cd5f8301846105ab565b92915050565b5f815190506105e1816104d0565b92915050565b5f602082840312156105fc576105fb61049c565b5b5f610609848285016105d3565b91505092915050565b5f82825260208201905092915050565b7f696e76616c69642063616c6c65720000000000000000000000000000000000005f82015250565b5f610656600e83610612565b915061066182610622565b602082019050919050565b5f6020820190508181035f8301526106838161064a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6106ee82610525565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036107205761071f6106b7565b5b600182019050919050565b5f61073582610525565b915061074083610525565b9250828203905081811115610758576107576106b7565b5b9291505056fea2646970667358221220f7521aece7d6a4a6701f3c9e64be45388c1c175cd9f78b83ce609db5a8708c7664736f6c63430008180033",
}

// ContractWhitelistManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractWhitelistManagerMetaData.ABI instead.
var ContractWhitelistManagerABI = ContractWhitelistManagerMetaData.ABI

var ContractWhitelistManagerParsedABI, _ = abi.JSON(strings.NewReader(ContractWhitelistManagerABI))

// ContractWhitelistManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractWhitelistManagerMetaData.Bin instead.
var ContractWhitelistManagerBin = ContractWhitelistManagerMetaData.Bin

// DeployContractWhitelistManager deploys a new Ethereum contract, binding an instance of ContractWhitelistManager to it.
func DeployContractWhitelistManager(auth *bind.TransactOpts, backend bind.ContractBackend, _permUpgradable common.Address) (common.Address, *types.Transaction, *ContractWhitelistManager, error) {
	parsed, err := ContractWhitelistManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractWhitelistManagerBin), backend, _permUpgradable)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractWhitelistManager{ContractWhitelistManagerCaller: ContractWhitelistManagerCaller{contract: contract}, ContractWhitelistManagerTransactor: ContractWhitelistManagerTransactor{contract: contract}, ContractWhitelistManagerFilterer: ContractWhitelistManagerFilterer{contract: contract}}, nil
}

// ContractWhitelistManager is an auto generated Go binding around an Ethereum contract.
type ContractWhitelistManager struct {
	ContractWhitelistManagerCaller     // Read-only binding to the contract
	ContractWhitelistManagerTransactor // Write-only binding to the contract
	ContractWhitelistManagerFilterer   // Log filterer for contract events
}

// ContractWhitelistManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractWhitelistManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractWhitelistManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractWhitelistManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractWhitelistManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractWhitelistManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractWhitelistManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractWhitelistManagerSession struct {
	Contract     *ContractWhitelistManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractWhitelistManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractWhitelistManagerCallerSession struct {
	Contract *ContractWhitelistManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractWhitelistManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractWhitelistManagerTransactorSession struct {
	Contract     *ContractWhitelistManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractWhitelistManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractWhitelistManagerRaw struct {
	Contract *ContractWhitelistManager // Generic contract binding to access the raw methods on
}

// ContractWhitelistManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractWhitelistManagerCallerRaw struct {
	Contract *ContractWhitelistManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractWhitelistManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractWhitelistManagerTransactorRaw struct {
	Contract *ContractWhitelistManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractWhitelistManager creates a new instance of ContractWhitelistManager, bound to a specific deployed contract.
func NewContractWhitelistManager(address common.Address, backend bind.ContractBackend) (*ContractWhitelistManager, error) {
	contract, err := bindContractWhitelistManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManager{ContractWhitelistManagerCaller: ContractWhitelistManagerCaller{contract: contract}, ContractWhitelistManagerTransactor: ContractWhitelistManagerTransactor{contract: contract}, ContractWhitelistManagerFilterer: ContractWhitelistManagerFilterer{contract: contract}}, nil
}

// NewContractWhitelistManagerCaller creates a new read-only instance of ContractWhitelistManager, bound to a specific deployed contract.
func NewContractWhitelistManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractWhitelistManagerCaller, error) {
	contract, err := bindContractWhitelistManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerCaller{contract: contract}, nil
}

// NewContractWhitelistManagerTransactor creates a new write-only instance of ContractWhitelistManager, bound to a specific deployed contract.
func NewContractWhitelistManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractWhitelistManagerTransactor, error) {
	contract, err := bindContractWhitelistManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerTransactor{contract: contract}, nil
}

// NewContractWhitelistManagerFilterer creates a new log filterer instance of ContractWhitelistManager, bound to a specific deployed contract.
func NewContractWhitelistManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractWhitelistManagerFilterer, error) {
	contract, err := bindContractWhitelistManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerFilterer{contract: contract}, nil
}

// bindContractWhitelistManager binds a generic wrapper to an already deployed contract.
func bindContractWhitelistManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractWhitelistManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractWhitelistManager *ContractWhitelistManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractWhitelistManager.Contract.ContractWhitelistManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractWhitelistManager *ContractWhitelistManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.ContractWhitelistManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractWhitelistManager *ContractWhitelistManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.ContractWhitelistManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractWhitelistManager *ContractWhitelistManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractWhitelistManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.contract.Transact(opts, method, params...)
}

// GetContractWhitelistDetailsFromIndex is a free data retrieval call binding the contract method 0xb20f4fa5.
//
// Solidity: function getContractWhitelistDetailsFromIndex(uint256 _cIndex) view returns(address)
func (_ContractWhitelistManager *ContractWhitelistManagerCaller) GetContractWhitelistDetailsFromIndex(opts *bind.CallOpts, _cIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractWhitelistManager.contract.Call(opts, &out, "getContractWhitelistDetailsFromIndex", _cIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContractWhitelistDetailsFromIndex is a free data retrieval call binding the contract method 0xb20f4fa5.
//
// Solidity: function getContractWhitelistDetailsFromIndex(uint256 _cIndex) view returns(address)
func (_ContractWhitelistManager *ContractWhitelistManagerSession) GetContractWhitelistDetailsFromIndex(_cIndex *big.Int) (common.Address, error) {
	return _ContractWhitelistManager.Contract.GetContractWhitelistDetailsFromIndex(&_ContractWhitelistManager.CallOpts, _cIndex)
}

// GetContractWhitelistDetailsFromIndex is a free data retrieval call binding the contract method 0xb20f4fa5.
//
// Solidity: function getContractWhitelistDetailsFromIndex(uint256 _cIndex) view returns(address)
func (_ContractWhitelistManager *ContractWhitelistManagerCallerSession) GetContractWhitelistDetailsFromIndex(_cIndex *big.Int) (common.Address, error) {
	return _ContractWhitelistManager.Contract.GetContractWhitelistDetailsFromIndex(&_ContractWhitelistManager.CallOpts, _cIndex)
}

// GetNumberOfWhitelistedContracts is a free data retrieval call binding the contract method 0xb5a93e26.
//
// Solidity: function getNumberOfWhitelistedContracts() view returns(uint256)
func (_ContractWhitelistManager *ContractWhitelistManagerCaller) GetNumberOfWhitelistedContracts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractWhitelistManager.contract.Call(opts, &out, "getNumberOfWhitelistedContracts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfWhitelistedContracts is a free data retrieval call binding the contract method 0xb5a93e26.
//
// Solidity: function getNumberOfWhitelistedContracts() view returns(uint256)
func (_ContractWhitelistManager *ContractWhitelistManagerSession) GetNumberOfWhitelistedContracts() (*big.Int, error) {
	return _ContractWhitelistManager.Contract.GetNumberOfWhitelistedContracts(&_ContractWhitelistManager.CallOpts)
}

// GetNumberOfWhitelistedContracts is a free data retrieval call binding the contract method 0xb5a93e26.
//
// Solidity: function getNumberOfWhitelistedContracts() view returns(uint256)
func (_ContractWhitelistManager *ContractWhitelistManagerCallerSession) GetNumberOfWhitelistedContracts() (*big.Int, error) {
	return _ContractWhitelistManager.Contract.GetNumberOfWhitelistedContracts(&_ContractWhitelistManager.CallOpts)
}

// AddNewContract is a paid mutator transaction binding the contract method 0x9a7fda66.
//
// Solidity: function addNewContract(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactor) AddNewContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.contract.Transact(opts, "addNewContract", _contract)
}

// AddNewContract is a paid mutator transaction binding the contract method 0x9a7fda66.
//
// Solidity: function addNewContract(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerSession) AddNewContract(_contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.AddNewContract(&_ContractWhitelistManager.TransactOpts, _contract)
}

// AddNewContract is a paid mutator transaction binding the contract method 0x9a7fda66.
//
// Solidity: function addNewContract(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorSession) AddNewContract(_contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.AddNewContract(&_ContractWhitelistManager.TransactOpts, _contract)
}

// ContractWhitelistManagerContractWhitelistModifiedIterator is returned from FilterContractWhitelistModified and is used to iterate over the raw logs and unpacked data for ContractWhitelistModified events raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerContractWhitelistModifiedIterator struct {
	Event *ContractWhitelistManagerContractWhitelistModified // Event containing the contract specifics and raw log

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
func (it *ContractWhitelistManagerContractWhitelistModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWhitelistManagerContractWhitelistModified)
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
		it.Event = new(ContractWhitelistManagerContractWhitelistModified)
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
func (it *ContractWhitelistManagerContractWhitelistModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWhitelistManagerContractWhitelistModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWhitelistManagerContractWhitelistModified represents a ContractWhitelistModified event raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerContractWhitelistModified struct {
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractWhitelistModified is a free log retrieval operation binding the contract event 0xb354cac8f12f93c803369656a115296370a81eafcfed866fce91e47408be0c2a.
//
// Solidity: event ContractWhitelistModified(address _contract)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) FilterContractWhitelistModified(opts *bind.FilterOpts) (*ContractWhitelistManagerContractWhitelistModifiedIterator, error) {

	logs, sub, err := _ContractWhitelistManager.contract.FilterLogs(opts, "ContractWhitelistModified")
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerContractWhitelistModifiedIterator{contract: _ContractWhitelistManager.contract, event: "ContractWhitelistModified", logs: logs, sub: sub}, nil
}

var ContractWhitelistModifiedTopicHash = "0xb354cac8f12f93c803369656a115296370a81eafcfed866fce91e47408be0c2a"

// WatchContractWhitelistModified is a free log subscription operation binding the contract event 0xb354cac8f12f93c803369656a115296370a81eafcfed866fce91e47408be0c2a.
//
// Solidity: event ContractWhitelistModified(address _contract)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) WatchContractWhitelistModified(opts *bind.WatchOpts, sink chan<- *ContractWhitelistManagerContractWhitelistModified) (event.Subscription, error) {

	logs, sub, err := _ContractWhitelistManager.contract.WatchLogs(opts, "ContractWhitelistModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWhitelistManagerContractWhitelistModified)
				if err := _ContractWhitelistManager.contract.UnpackLog(event, "ContractWhitelistModified", log); err != nil {
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

// ParseContractWhitelistModified is a log parse operation binding the contract event 0xb354cac8f12f93c803369656a115296370a81eafcfed866fce91e47408be0c2a.
//
// Solidity: event ContractWhitelistModified(address _contract)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) ParseContractWhitelistModified(log types.Log) (*ContractWhitelistManagerContractWhitelistModified, error) {
	event := new(ContractWhitelistManagerContractWhitelistModified)
	if err := _ContractWhitelistManager.contract.UnpackLog(event, "ContractWhitelistModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
