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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"ContractWhitelistModified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"addNewContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfWhitelistedContracts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516107ca3803806107ca833981810160405281019061003191906100d4565b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506100ff565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a38261007a565b9050919050565b6100b381610099565b81146100bd575f80fd5b50565b5f815190506100ce816100aa565b92915050565b5f602082840312156100e9576100e8610076565b5b5f6100f6848285016100c0565b91505092915050565b6106be8061010c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80639a7fda6614610038578063b5a93e2614610054575b5f80fd5b610052600480360381019061004d9190610479565b610072565b005b61005c6103bd565b60405161006991906104bc565b60405180910390f35b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100da573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100fe91906104e9565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461016b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101629061056e565b60405180910390fd5b60035460025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055507fb354cac8f12f93c803369656a115296370a81eafcfed866fce91e47408be0c2a816040516101de919061059b565b60405180910390a15f6101f0826103c9565b90505f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541461029757816001828154811061024b5761024a6105b4565b5b905f5260205f20015f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610382565b60035f8154809291906102a99061060e565b919050555060035460025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550600160405180602001604052808473ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003905f5260205f20015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505b7fb354cac8f12f93c803369656a115296370a81eafcfed866fce91e47408be0c2a826040516103b1919061059b565b60405180910390a15050565b5f600180549050905090565b5f600160025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546104149190610655565b9050919050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104488261041f565b9050919050565b6104588161043e565b8114610462575f80fd5b50565b5f813590506104738161044f565b92915050565b5f6020828403121561048e5761048d61041b565b5b5f61049b84828501610465565b91505092915050565b5f819050919050565b6104b6816104a4565b82525050565b5f6020820190506104cf5f8301846104ad565b92915050565b5f815190506104e38161044f565b92915050565b5f602082840312156104fe576104fd61041b565b5b5f61050b848285016104d5565b91505092915050565b5f82825260208201905092915050565b7f696e76616c69642063616c6c65720000000000000000000000000000000000005f82015250565b5f610558600e83610514565b915061056382610524565b602082019050919050565b5f6020820190508181035f8301526105858161054c565b9050919050565b6105958161043e565b82525050565b5f6020820190506105ae5f83018461058c565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610618826104a4565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361064a576106496105e1565b5b600182019050919050565b5f61065f826104a4565b915061066a836104a4565b9250828203905081811115610682576106816105e1565b5b9291505056fea2646970667358221220d6afde719a10e7c20b6a0061700632a85ea81dc14d1caf8d0a08fe972910b6bc64736f6c63430008180033",
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
