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
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contractAddr\",\"type\":\"address\"}],\"name\":\"ContractWhitelistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contractAddr\",\"type\":\"address\"}],\"name\":\"ContractWhitelistRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"addWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfWhitelistedContracts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistedContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"isContractWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"revokeWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610f6f8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80637e47daa1146100645780639c7f331514610082578063b5a93e261461009e578063c057058a146100bc578063c4d66de8146100ec578063f80f5dd514610108575b5f80fd5b61006c610124565b6040516100799190610b45565b60405180910390f35b61009c60048036038101906100979190610b93565b61013e565b005b6100a66102ff565b6040516100b39190610bd6565b60405180910390f35b6100d660048036038101906100d19190610b93565b610318565b6040516100e39190610c09565b60405180910390f35b61010660048036038101906101019190610b93565b61033d565b005b610122600480360381019061011d9190610b93565b61056b565b005b606061013961013161072d565b600101610793565b905090565b61014661072d565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101b0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101d49190610c36565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610241576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023890610cbb565b60405180910390fd5b5f61024a61072d565b905061026282826001016107b290919063ffffffff16565b6102a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029890610d23565b60405180910390fd5b6102b782826001016107df90919063ffffffff16565b508173ffffffffffffffffffffffffffffffffffffffff167fe5044c30eab65bda022826bec45e211a3f9c73ed493965c7fef15d626fb065d960405160405180910390a25050565b5f61031361030b61072d565b60010161080c565b905090565b5f6103368261032561072d565b6001016107b290919063ffffffff16565b9050919050565b5f61034661081f565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff1614801561038e5750825b90505f60018367ffffffffffffffff161480156103c157505f3073ffffffffffffffffffffffffffffffffffffffff163b145b9050811580156103cf575080155b15610406576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610453576001855f0160086101000a81548160ff0219169083151502179055505b5f73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16036104c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b890610d8b565b60405180910390fd5b856104ca61072d565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508315610563575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2600160405161055a9190610dfe565b60405180910390a15b505050505050565b61057361072d565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630e32cf906040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105dd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106019190610c36565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066590610cbb565b60405180910390fd5b5f61067761072d565b905061068f82826001016107b290919063ffffffff16565b156106cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c690610e61565b60405180910390fd5b6106e5828260010161084690919063ffffffff16565b508173ffffffffffffffffffffffffffffffffffffffff167f69d296e2418833651c8d83f409c6339e9f1243ad2653b985e62a39d4307ee50b60405160405180910390a25050565b5f8060ff5f1b1960017fdc0a0fb9b8c3742858130ce0eafb2fa7793d4ff4fec8654c10918f0e0dfd8c765f1c6107639190610eac565b6040516020016107739190610bd6565b604051602081830303815290604052805190602001201690508091505090565b60605f6107a1835f01610873565b905060608190508092505050919050565b5f6107d7835f018373ffffffffffffffffffffffffffffffffffffffff165f1b6108cc565b905092915050565b5f610804835f018373ffffffffffffffffffffffffffffffffffffffff165f1b6108ec565b905092915050565b5f610818825f016109e8565b9050919050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f61086b835f018373ffffffffffffffffffffffffffffffffffffffff165f1b6109f7565b905092915050565b6060815f018054806020026020016040519081016040528092919081815260200182805480156108c057602002820191905f5260205f20905b8154815260200190600101908083116108ac575b50505050509050919050565b5f80836001015f8481526020019081526020015f20541415905092915050565b5f80836001015f8481526020019081526020015f205490505f81146109dd575f6001826109199190610eac565b90505f6001865f018054905061092f9190610eac565b9050808214610995575f865f01828154811061094e5761094d610edf565b5b905f5260205f200154905080875f01848154811061096f5761096e610edf565b5b905f5260205f20018190555083876001015f8381526020019081526020015f2081905550505b855f018054806109a8576109a7610f0c565b5b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506109e2565b5f9150505b92915050565b5f815f01805490509050919050565b5f610a0283836108cc565b610a5457825f0182908060018154018082558091505060019003905f5260205f20015f9091909190915055825f0180549050836001015f8481526020019081526020015f208190555060019050610a58565b5f90505b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ab082610a87565b9050919050565b610ac081610aa6565b82525050565b5f610ad18383610ab7565b60208301905092915050565b5f602082019050919050565b5f610af382610a5e565b610afd8185610a68565b9350610b0883610a78565b805f5b83811015610b38578151610b1f8882610ac6565b9750610b2a83610add565b925050600181019050610b0b565b5085935050505092915050565b5f6020820190508181035f830152610b5d8184610ae9565b905092915050565b5f80fd5b610b7281610aa6565b8114610b7c575f80fd5b50565b5f81359050610b8d81610b69565b92915050565b5f60208284031215610ba857610ba7610b65565b5b5f610bb584828501610b7f565b91505092915050565b5f819050919050565b610bd081610bbe565b82525050565b5f602082019050610be95f830184610bc7565b92915050565b5f8115159050919050565b610c0381610bef565b82525050565b5f602082019050610c1c5f830184610bfa565b92915050565b5f81519050610c3081610b69565b92915050565b5f60208284031215610c4b57610c4a610b65565b5b5f610c5884828501610c22565b91505092915050565b5f82825260208201905092915050565b7f696e76616c69642063616c6c65720000000000000000000000000000000000005f82015250565b5f610ca5600e83610c61565b9150610cb082610c71565b602082019050919050565b5f6020820190508181035f830152610cd281610c99565b9050919050565b7f77686974656c69737420646f6573206e6f7420657869737400000000000000005f82015250565b5f610d0d601883610c61565b9150610d1882610cd9565b602082019050919050565b5f6020820190508181035f830152610d3a81610d01565b9050919050565b7f43616e6e6f742073657420746f20656d707479206164647265737300000000005f82015250565b5f610d75601b83610c61565b9150610d8082610d41565b602082019050919050565b5f6020820190508181035f830152610da281610d69565b9050919050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f819050919050565b5f610de8610de3610dde84610da9565b610dc5565b610db2565b9050919050565b610df881610dce565b82525050565b5f602082019050610e115f830184610def565b92915050565b7f77686974656c69737420616c72656164792065786973747300000000000000005f82015250565b5f610e4b601883610c61565b9150610e5682610e17565b602082019050919050565b5f6020820190508181035f830152610e7881610e3f565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610eb682610bbe565b9150610ec183610bbe565b9250828203905081811115610ed957610ed8610e7f565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea2646970667358221220ac597c0c8eed2f244d40ea629ef71261069daac05f3a2df25a9c31fb65fb6e6264736f6c63430008180033",
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

// GetWhitelistedContracts is a free data retrieval call binding the contract method 0x7e47daa1.
//
// Solidity: function getWhitelistedContracts() view returns(address[])
func (_ContractWhitelistManager *ContractWhitelistManagerCaller) GetWhitelistedContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractWhitelistManager.contract.Call(opts, &out, "getWhitelistedContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelistedContracts is a free data retrieval call binding the contract method 0x7e47daa1.
//
// Solidity: function getWhitelistedContracts() view returns(address[])
func (_ContractWhitelistManager *ContractWhitelistManagerSession) GetWhitelistedContracts() ([]common.Address, error) {
	return _ContractWhitelistManager.Contract.GetWhitelistedContracts(&_ContractWhitelistManager.CallOpts)
}

// GetWhitelistedContracts is a free data retrieval call binding the contract method 0x7e47daa1.
//
// Solidity: function getWhitelistedContracts() view returns(address[])
func (_ContractWhitelistManager *ContractWhitelistManagerCallerSession) GetWhitelistedContracts() ([]common.Address, error) {
	return _ContractWhitelistManager.Contract.GetWhitelistedContracts(&_ContractWhitelistManager.CallOpts)
}

// IsContractWhitelisted is a free data retrieval call binding the contract method 0xc057058a.
//
// Solidity: function isContractWhitelisted(address _contract) view returns(bool)
func (_ContractWhitelistManager *ContractWhitelistManagerCaller) IsContractWhitelisted(opts *bind.CallOpts, _contract common.Address) (bool, error) {
	var out []interface{}
	err := _ContractWhitelistManager.contract.Call(opts, &out, "isContractWhitelisted", _contract)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContractWhitelisted is a free data retrieval call binding the contract method 0xc057058a.
//
// Solidity: function isContractWhitelisted(address _contract) view returns(bool)
func (_ContractWhitelistManager *ContractWhitelistManagerSession) IsContractWhitelisted(_contract common.Address) (bool, error) {
	return _ContractWhitelistManager.Contract.IsContractWhitelisted(&_ContractWhitelistManager.CallOpts, _contract)
}

// IsContractWhitelisted is a free data retrieval call binding the contract method 0xc057058a.
//
// Solidity: function isContractWhitelisted(address _contract) view returns(bool)
func (_ContractWhitelistManager *ContractWhitelistManagerCallerSession) IsContractWhitelisted(_contract common.Address) (bool, error) {
	return _ContractWhitelistManager.Contract.IsContractWhitelisted(&_ContractWhitelistManager.CallOpts, _contract)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactor) AddWhitelist(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.contract.Transact(opts, "addWhitelist", _contract)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerSession) AddWhitelist(_contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.AddWhitelist(&_ContractWhitelistManager.TransactOpts, _contract)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0xf80f5dd5.
//
// Solidity: function addWhitelist(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorSession) AddWhitelist(_contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.AddWhitelist(&_ContractWhitelistManager.TransactOpts, _contract)
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

// RevokeWhitelist is a paid mutator transaction binding the contract method 0x9c7f3315.
//
// Solidity: function revokeWhitelist(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactor) RevokeWhitelist(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.contract.Transact(opts, "revokeWhitelist", _contract)
}

// RevokeWhitelist is a paid mutator transaction binding the contract method 0x9c7f3315.
//
// Solidity: function revokeWhitelist(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerSession) RevokeWhitelist(_contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.RevokeWhitelist(&_ContractWhitelistManager.TransactOpts, _contract)
}

// RevokeWhitelist is a paid mutator transaction binding the contract method 0x9c7f3315.
//
// Solidity: function revokeWhitelist(address _contract) returns()
func (_ContractWhitelistManager *ContractWhitelistManagerTransactorSession) RevokeWhitelist(_contract common.Address) (*types.Transaction, error) {
	return _ContractWhitelistManager.Contract.RevokeWhitelist(&_ContractWhitelistManager.TransactOpts, _contract)
}

// ContractWhitelistManagerContractWhitelistAddedIterator is returned from FilterContractWhitelistAdded and is used to iterate over the raw logs and unpacked data for ContractWhitelistAdded events raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerContractWhitelistAddedIterator struct {
	Event *ContractWhitelistManagerContractWhitelistAdded // Event containing the contract specifics and raw log

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
func (it *ContractWhitelistManagerContractWhitelistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWhitelistManagerContractWhitelistAdded)
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
		it.Event = new(ContractWhitelistManagerContractWhitelistAdded)
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
func (it *ContractWhitelistManagerContractWhitelistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWhitelistManagerContractWhitelistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWhitelistManagerContractWhitelistAdded represents a ContractWhitelistAdded event raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerContractWhitelistAdded struct {
	ContractAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractWhitelistAdded is a free log retrieval operation binding the contract event 0x69d296e2418833651c8d83f409c6339e9f1243ad2653b985e62a39d4307ee50b.
//
// Solidity: event ContractWhitelistAdded(address indexed _contractAddr)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) FilterContractWhitelistAdded(opts *bind.FilterOpts, _contractAddr []common.Address) (*ContractWhitelistManagerContractWhitelistAddedIterator, error) {

	var _contractAddrRule []interface{}
	for _, _contractAddrItem := range _contractAddr {
		_contractAddrRule = append(_contractAddrRule, _contractAddrItem)
	}

	logs, sub, err := _ContractWhitelistManager.contract.FilterLogs(opts, "ContractWhitelistAdded", _contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerContractWhitelistAddedIterator{contract: _ContractWhitelistManager.contract, event: "ContractWhitelistAdded", logs: logs, sub: sub}, nil
}

var ContractWhitelistAddedTopicHash = "0x69d296e2418833651c8d83f409c6339e9f1243ad2653b985e62a39d4307ee50b"

// WatchContractWhitelistAdded is a free log subscription operation binding the contract event 0x69d296e2418833651c8d83f409c6339e9f1243ad2653b985e62a39d4307ee50b.
//
// Solidity: event ContractWhitelistAdded(address indexed _contractAddr)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) WatchContractWhitelistAdded(opts *bind.WatchOpts, sink chan<- *ContractWhitelistManagerContractWhitelistAdded, _contractAddr []common.Address) (event.Subscription, error) {

	var _contractAddrRule []interface{}
	for _, _contractAddrItem := range _contractAddr {
		_contractAddrRule = append(_contractAddrRule, _contractAddrItem)
	}

	logs, sub, err := _ContractWhitelistManager.contract.WatchLogs(opts, "ContractWhitelistAdded", _contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWhitelistManagerContractWhitelistAdded)
				if err := _ContractWhitelistManager.contract.UnpackLog(event, "ContractWhitelistAdded", log); err != nil {
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

// ParseContractWhitelistAdded is a log parse operation binding the contract event 0x69d296e2418833651c8d83f409c6339e9f1243ad2653b985e62a39d4307ee50b.
//
// Solidity: event ContractWhitelistAdded(address indexed _contractAddr)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) ParseContractWhitelistAdded(log types.Log) (*ContractWhitelistManagerContractWhitelistAdded, error) {
	event := new(ContractWhitelistManagerContractWhitelistAdded)
	if err := _ContractWhitelistManager.contract.UnpackLog(event, "ContractWhitelistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWhitelistManagerContractWhitelistRevokedIterator is returned from FilterContractWhitelistRevoked and is used to iterate over the raw logs and unpacked data for ContractWhitelistRevoked events raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerContractWhitelistRevokedIterator struct {
	Event *ContractWhitelistManagerContractWhitelistRevoked // Event containing the contract specifics and raw log

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
func (it *ContractWhitelistManagerContractWhitelistRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWhitelistManagerContractWhitelistRevoked)
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
		it.Event = new(ContractWhitelistManagerContractWhitelistRevoked)
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
func (it *ContractWhitelistManagerContractWhitelistRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWhitelistManagerContractWhitelistRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWhitelistManagerContractWhitelistRevoked represents a ContractWhitelistRevoked event raised by the ContractWhitelistManager contract.
type ContractWhitelistManagerContractWhitelistRevoked struct {
	ContractAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractWhitelistRevoked is a free log retrieval operation binding the contract event 0xe5044c30eab65bda022826bec45e211a3f9c73ed493965c7fef15d626fb065d9.
//
// Solidity: event ContractWhitelistRevoked(address indexed _contractAddr)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) FilterContractWhitelistRevoked(opts *bind.FilterOpts, _contractAddr []common.Address) (*ContractWhitelistManagerContractWhitelistRevokedIterator, error) {

	var _contractAddrRule []interface{}
	for _, _contractAddrItem := range _contractAddr {
		_contractAddrRule = append(_contractAddrRule, _contractAddrItem)
	}

	logs, sub, err := _ContractWhitelistManager.contract.FilterLogs(opts, "ContractWhitelistRevoked", _contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &ContractWhitelistManagerContractWhitelistRevokedIterator{contract: _ContractWhitelistManager.contract, event: "ContractWhitelistRevoked", logs: logs, sub: sub}, nil
}

var ContractWhitelistRevokedTopicHash = "0xe5044c30eab65bda022826bec45e211a3f9c73ed493965c7fef15d626fb065d9"

// WatchContractWhitelistRevoked is a free log subscription operation binding the contract event 0xe5044c30eab65bda022826bec45e211a3f9c73ed493965c7fef15d626fb065d9.
//
// Solidity: event ContractWhitelistRevoked(address indexed _contractAddr)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) WatchContractWhitelistRevoked(opts *bind.WatchOpts, sink chan<- *ContractWhitelistManagerContractWhitelistRevoked, _contractAddr []common.Address) (event.Subscription, error) {

	var _contractAddrRule []interface{}
	for _, _contractAddrItem := range _contractAddr {
		_contractAddrRule = append(_contractAddrRule, _contractAddrItem)
	}

	logs, sub, err := _ContractWhitelistManager.contract.WatchLogs(opts, "ContractWhitelistRevoked", _contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWhitelistManagerContractWhitelistRevoked)
				if err := _ContractWhitelistManager.contract.UnpackLog(event, "ContractWhitelistRevoked", log); err != nil {
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

// ParseContractWhitelistRevoked is a log parse operation binding the contract event 0xe5044c30eab65bda022826bec45e211a3f9c73ed493965c7fef15d626fb065d9.
//
// Solidity: event ContractWhitelistRevoked(address indexed _contractAddr)
func (_ContractWhitelistManager *ContractWhitelistManagerFilterer) ParseContractWhitelistRevoked(log types.Log) (*ContractWhitelistManagerContractWhitelistRevoked, error) {
	event := new(ContractWhitelistManagerContractWhitelistRevoked)
	if err := _ContractWhitelistManager.contract.UnpackLog(event, "ContractWhitelistRevoked", log); err != nil {
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
