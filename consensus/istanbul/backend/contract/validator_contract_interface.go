// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ValidatorContractInterfaceABI is the input ABI used to generate the binding from.
const ValidatorContractInterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

var ValidatorContractInterfaceParsedABI, _ = abi.JSON(strings.NewReader(ValidatorContractInterfaceABI))

// ValidatorContractInterface is an auto generated Go binding around an Ethereum contract.
type ValidatorContractInterface struct {
	ValidatorContractInterfaceCaller     // Read-only binding to the contract
	ValidatorContractInterfaceTransactor // Write-only binding to the contract
	ValidatorContractInterfaceFilterer   // Log filterer for contract events
}

// ValidatorContractInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorContractInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorContractInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorContractInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorContractInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorContractInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorContractInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorContractInterfaceSession struct {
	Contract     *ValidatorContractInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ValidatorContractInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorContractInterfaceCallerSession struct {
	Contract *ValidatorContractInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ValidatorContractInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorContractInterfaceTransactorSession struct {
	Contract     *ValidatorContractInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ValidatorContractInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorContractInterfaceRaw struct {
	Contract *ValidatorContractInterface // Generic contract binding to access the raw methods on
}

// ValidatorContractInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorContractInterfaceCallerRaw struct {
	Contract *ValidatorContractInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorContractInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorContractInterfaceTransactorRaw struct {
	Contract *ValidatorContractInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorContractInterface creates a new instance of ValidatorContractInterface, bound to a specific deployed contract.
func NewValidatorContractInterface(address common.Address, backend bind.ContractBackend) (*ValidatorContractInterface, error) {
	contract, err := bindValidatorContractInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorContractInterface{ValidatorContractInterfaceCaller: ValidatorContractInterfaceCaller{contract: contract}, ValidatorContractInterfaceTransactor: ValidatorContractInterfaceTransactor{contract: contract}, ValidatorContractInterfaceFilterer: ValidatorContractInterfaceFilterer{contract: contract}}, nil
}

// NewValidatorContractInterfaceCaller creates a new read-only instance of ValidatorContractInterface, bound to a specific deployed contract.
func NewValidatorContractInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ValidatorContractInterfaceCaller, error) {
	contract, err := bindValidatorContractInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorContractInterfaceCaller{contract: contract}, nil
}

// NewValidatorContractInterfaceTransactor creates a new write-only instance of ValidatorContractInterface, bound to a specific deployed contract.
func NewValidatorContractInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorContractInterfaceTransactor, error) {
	contract, err := bindValidatorContractInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorContractInterfaceTransactor{contract: contract}, nil
}

// NewValidatorContractInterfaceFilterer creates a new log filterer instance of ValidatorContractInterface, bound to a specific deployed contract.
func NewValidatorContractInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorContractInterfaceFilterer, error) {
	contract, err := bindValidatorContractInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorContractInterfaceFilterer{contract: contract}, nil
}

// bindValidatorContractInterface binds a generic wrapper to an already deployed contract.
func bindValidatorContractInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorContractInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorContractInterface *ValidatorContractInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorContractInterface.Contract.ValidatorContractInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorContractInterface *ValidatorContractInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorContractInterface.Contract.ValidatorContractInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorContractInterface *ValidatorContractInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorContractInterface.Contract.ValidatorContractInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorContractInterface *ValidatorContractInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorContractInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorContractInterface *ValidatorContractInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorContractInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorContractInterface *ValidatorContractInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorContractInterface.Contract.contract.Transact(opts, method, params...)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ValidatorContractInterface *ValidatorContractInterfaceCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ValidatorContractInterface.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ValidatorContractInterface *ValidatorContractInterfaceSession) GetValidators() ([]common.Address, error) {
	return _ValidatorContractInterface.Contract.GetValidators(&_ValidatorContractInterface.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ValidatorContractInterface *ValidatorContractInterfaceCallerSession) GetValidators() ([]common.Address, error) {
	return _ValidatorContractInterface.Contract.GetValidators(&_ValidatorContractInterface.CallOpts)
}
