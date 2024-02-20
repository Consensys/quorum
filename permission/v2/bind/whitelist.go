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
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"ContractWhitelistModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"addNewContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cIndex\",\"type\":\"uint256\"}],\"name\":\"getContractWhitelistDetailsFromIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfWhitelistedContracts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610add8061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c80639a7fda661461004e578063b20f4fa51461006a578063b5a93e261461009a578063c4d66de8146100b8575b5f80fd5b6100686004803603810190610063919061076d565b6100d4565b005b610084600480360381019061007f91906107cb565b61041f565b6040516100919190610805565b60405180910390f35b6100a2610465565b6040516100af919061082d565b60405180910390f35b6100d260048036038101906100cd919061076d565b610471565b005b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b8152600401602060405180830381865afa15801561013c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610160919061085a565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101c4906108df565b60405180910390fd5b60035460025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055507fb354cac8f12f93c803369656a115296370a81eafcfed866fce91e47408be0c2a816040516102409190610805565b60405180910390a15f61025282610696565b90505f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054146102f95781600182815481106102ad576102ac6108fd565b5b905f5260205f20015f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506103e4565b60035f81548092919061030b90610957565b919050555060035460025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550600160405180602001604052808473ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003905f5260205f20015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505b7fb354cac8f12f93c803369656a115296370a81eafcfed866fce91e47408be0c2a826040516104139190610805565b60405180910390a15050565b5f60018281548110610434576104336108fd565b5b905f5260205f20015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f600180549050905090565b5f61047a6106e8565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff161480156104c25750825b90505f60018367ffffffffffffffff161480156104f557505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610503575080155b1561053a576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610587576001855f0160086101000a81548160ff0219169083151502179055505b5f73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16036105f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ec906109e8565b60405180910390fd5b855f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550831561068e575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516106859190610a5b565b60405180910390a15b505050505050565b5f600160025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546106e19190610a74565b9050919050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61073c82610713565b9050919050565b61074c81610732565b8114610756575f80fd5b50565b5f8135905061076781610743565b92915050565b5f602082840312156107825761078161070f565b5b5f61078f84828501610759565b91505092915050565b5f819050919050565b6107aa81610798565b81146107b4575f80fd5b50565b5f813590506107c5816107a1565b92915050565b5f602082840312156107e0576107df61070f565b5b5f6107ed848285016107b7565b91505092915050565b6107ff81610732565b82525050565b5f6020820190506108185f8301846107f6565b92915050565b61082781610798565b82525050565b5f6020820190506108405f83018461081e565b92915050565b5f8151905061085481610743565b92915050565b5f6020828403121561086f5761086e61070f565b5b5f61087c84828501610846565b91505092915050565b5f82825260208201905092915050565b7f696e76616c69642063616c6c65720000000000000000000000000000000000005f82015250565b5f6108c9600e83610885565b91506108d482610895565b602082019050919050565b5f6020820190508181035f8301526108f6816108bd565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61096182610798565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036109935761099261092a565b5b600182019050919050565b7f43616e6e6f742073657420746f20656d707479206164647265737300000000005f82015250565b5f6109d2601b83610885565b91506109dd8261099e565b602082019050919050565b5f6020820190508181035f8301526109ff816109c6565b9050919050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f819050919050565b5f610a45610a40610a3b84610a06565b610a22565b610a0f565b9050919050565b610a5581610a2b565b82525050565b5f602082019050610a6e5f830184610a4c565b92915050565b5f610a7e82610798565b9150610a8983610798565b9250828203905081811115610aa157610aa061092a565b5b9291505056fea26469706673582212200b1ab579455996a96e7a9d16df0fbc2a4dc7f88afda3873aa93ae69e96ec8a2864736f6c63430008180033",
}

// ContractWhitelistManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractWhitelistManagerMetaData.ABI instead.
var ContractWhitelistManagerABI = ContractWhitelistManagerMetaData.ABI

var ContractWhitelistManagerParsedABI, _ = abi.JSON(strings.NewReader(ContractWhitelistManagerABI))

// ContractWhitelistManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractWhitelistManagerMetaData.Bin instead.
var ContractWhitelistManagerBin = ContractWhitelistManagerMetaData.Bin

// DeployContractWhitelistManager deploys a new Ethereum contract, binding an instance of ContractWhitelistManager to it.
func DeployContractWhitelistManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractWhitelistManager, error) {
	parsed, err := ContractWhitelistManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractWhitelistManagerBin), backend)
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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permUpgradable) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactor) Initialize(opts *bind.TransactOpts, _permUpgradable common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.contract.Transact(opts, "initialize", _permUpgradable)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permUpgradable) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerSession) Initialize(_permUpgradable common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.Initialize(&_ContractWhitelistManager.TransactOpts, _permUpgradable)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _permUpgradable) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorSession) Initialize(_permUpgradable common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.Initialize(&_ContractWhitelistManager.TransactOpts, _permUpgradable)
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

// ContractWhitelistManagerInitializedWhitelistIterator is returned from FilterInitializedWhitelist and is used to iterate over the raw logs and unpacked data for InitializedWhitelist events raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerInitializedWhitelistIterator struct {
	Event *ContractWhitelistManagerInitializedWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractWhitelistManagerInitializedWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWhitelistManagerInitializedWhitelist)
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
		it.Event = new(ContractWhitelistManagerInitializedWhitelist)
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
func (it *ContractWhitelistManagerInitializedWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWhitelistManagerInitializedWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWhitelistManagerInitializedWhitelist represents a InitializedWhitelist event raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerInitializedWhitelist struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitializedWhitelist is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) FilterInitializedWhitelist(opts *bind.FilterOpts) (*ContractWhitelistManagerInitializedWhitelistIterator, error) {

	logs, sub, err := _ContractWhitelistManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerInitializedWhitelistIterator{contract: _ContractWhitelistManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

var InitializedWhitelistTopicHash = "0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2"

// WatchInitializedWhitelist is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) WatchInitializedWhitelist(opts *bind.WatchOpts, sink chan<- *ContractWhitelistManagerInitializedWhitelist) (event.Subscription, error) {

	logs, sub, err := _ContractWhitelistManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWhitelistManagerInitializedWhitelist)
				if err := _ContractWhitelistManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitializedWhitelist is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) ParseInitializedWhitelist(log types.Log) (*ContractWhitelistManagerInitializedWhitelist, error) {
	event := new(ContractWhitelistManagerInitializedWhitelist)
	if err := _ContractWhitelistManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
