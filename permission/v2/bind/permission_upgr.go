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

// PermUpgrMetaData contains all meta data concerning the PermUpgr contract.
var PermUpgrMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"confirmImplChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermInterface\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permInterface\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_permImpl\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610e0e8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80630e32cf901461006457806322bcb39a14610082578063a75b87d21461009e578063c4d66de8146100bc578063e572515c146100d8578063f09a4016146100f6575b5f80fd5b61006c610112565b6040516100799190610879565b60405180910390f35b61009c600480360381019061009791906108cd565b61013a565b005b6100a66102e2565b6040516100b39190610879565b60405180910390f35b6100d660048036038101906100d191906108cd565b610309565b005b6100e0610548565b6040516100ed9190610879565b60405180910390f35b610110600480360381019061010b91906108f8565b610570565b005b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101be90610990565b60405180910390fd5b5f805f8060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cc9ba6fa6040518163ffffffff1660e01b81526004015f60405180830381865afa158015610234573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061025c9190610b39565b93509350935093506102718585858585610716565b8460015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506102db60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610789565b5050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f610312610813565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff1614801561035a5750825b90505f60018367ffffffffffffffff1614801561038d57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b90508115801561039b575080155b156103d2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550831561041f576001855f0160086101000a81548160ff0219169083151502179055505b5f73ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff160361048d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048490610c3b565b60405180910390fd5b855f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f600260146101000a81548160ff0219169083151502179055508315610540575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516105379190610cae565b60405180910390a15b505050505050565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f490610990565b60405180910390fd5b600260149054906101000a900460ff161561064d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064490610d11565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506106f760015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610789565b6001600260146101000a81548160ff0219169083151502179055505050565b8473ffffffffffffffffffffffffffffffffffffffff1663f5ad584a858585856040518563ffffffff1660e01b81526004016107559493929190610d80565b5f604051808303815f87803b15801561076c575f80fd5b505af115801561077e573d5f803e3d5ffd5b505050505050505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663511bbd9f826040518263ffffffff1660e01b81526004016107e39190610879565b5f604051808303815f87803b1580156107fa575f80fd5b505af115801561080c573d5f803e3d5ffd5b5050505050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108638261083a565b9050919050565b61087381610859565b82525050565b5f60208201905061088c5f83018461086a565b92915050565b5f604051905090565b5f80fd5b5f80fd5b6108ac81610859565b81146108b6575f80fd5b50565b5f813590506108c7816108a3565b92915050565b5f602082840312156108e2576108e161089b565b5b5f6108ef848285016108b9565b91505092915050565b5f806040838503121561090e5761090d61089b565b5b5f61091b858286016108b9565b925050602061092c858286016108b9565b9150509250929050565b5f82825260208201905092915050565b7f696e76616c69642063616c6c65720000000000000000000000000000000000005f82015250565b5f61097a600e83610936565b915061098582610946565b602082019050919050565b5f6020820190508181035f8301526109a78161096e565b9050919050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6109fc826109b6565b810181811067ffffffffffffffff82111715610a1b57610a1a6109c6565b5b80604052505050565b5f610a2d610892565b9050610a3982826109f3565b919050565b5f67ffffffffffffffff821115610a5857610a576109c6565b5b610a61826109b6565b9050602081019050919050565b5f5b83811015610a8b578082015181840152602081019050610a70565b5f8484015250505050565b5f610aa8610aa384610a3e565b610a24565b905082815260208101848484011115610ac457610ac36109b2565b5b610acf848285610a6e565b509392505050565b5f82601f830112610aeb57610aea6109ae565b5b8151610afb848260208601610a96565b91505092915050565b5f8115159050919050565b610b1881610b04565b8114610b22575f80fd5b50565b5f81519050610b3381610b0f565b92915050565b5f805f8060808587031215610b5157610b5061089b565b5b5f85015167ffffffffffffffff811115610b6e57610b6d61089f565b5b610b7a87828801610ad7565b945050602085015167ffffffffffffffff811115610b9b57610b9a61089f565b5b610ba787828801610ad7565b935050604085015167ffffffffffffffff811115610bc857610bc761089f565b5b610bd487828801610ad7565b9250506060610be587828801610b25565b91505092959194509250565b7f43616e6e6f742073657420746f20656d707479206164647265737300000000005f82015250565b5f610c25601b83610936565b9150610c3082610bf1565b602082019050919050565b5f6020820190508181035f830152610c5281610c19565b9050919050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f819050919050565b5f610c98610c93610c8e84610c59565b610c75565b610c62565b9050919050565b610ca881610c7e565b82525050565b5f602082019050610cc15f830184610c9f565b92915050565b7f63616e206265206578656375746564206f6e6c79206f6e6365000000000000005f82015250565b5f610cfb601983610936565b9150610d0682610cc7565b602082019050919050565b5f6020820190508181035f830152610d2881610cef565b9050919050565b5f81519050919050565b5f610d4382610d2f565b610d4d8185610936565b9350610d5d818560208601610a6e565b610d66816109b6565b840191505092915050565b610d7a81610b04565b82525050565b5f6080820190508181035f830152610d988187610d39565b90508181036020830152610dac8186610d39565b90508181036040830152610dc08185610d39565b9050610dcf6060830184610d71565b9594505050505056fea26469706673582212208ff7673a9b6b4a84ebbf6423a23a39767839736f4ca9c31d52a6bd30a6ffc80164736f6c63430008180033",
}

// PermUpgrABI is the input ABI used to generate the binding from.
// Deprecated: Use PermUpgrMetaData.ABI instead.
var PermUpgrABI = PermUpgrMetaData.ABI

var PermUpgrParsedABI, _ = abi.JSON(strings.NewReader(PermUpgrABI))

// PermUpgrBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PermUpgrMetaData.Bin instead.
var PermUpgrBin = PermUpgrMetaData.Bin

// DeployPermUpgr deploys a new Ethereum contract, binding an instance of PermUpgr to it.
func DeployPermUpgr(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PermUpgr, error) {
	parsed, err := PermUpgrMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PermUpgrBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PermUpgr{PermUpgrCaller: PermUpgrCaller{contract: contract}, PermUpgrTransactor: PermUpgrTransactor{contract: contract}, PermUpgrFilterer: PermUpgrFilterer{contract: contract}}, nil
}

// PermUpgr is an auto generated Go binding around an Ethereum contract.
type PermUpgr struct {
	PermUpgrCaller     // Read-only binding to the contract
	PermUpgrTransactor // Write-only binding to the contract
	PermUpgrFilterer   // Log filterer for contract events
}

// PermUpgrCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermUpgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermUpgrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermUpgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermUpgrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermUpgrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermUpgrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermUpgrSession struct {
	Contract     *PermUpgr         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermUpgrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermUpgrCallerSession struct {
	Contract *PermUpgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PermUpgrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermUpgrTransactorSession struct {
	Contract     *PermUpgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PermUpgrRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermUpgrRaw struct {
	Contract *PermUpgr // Generic contract binding to access the raw methods on
}

// PermUpgrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermUpgrCallerRaw struct {
	Contract *PermUpgrCaller // Generic read-only contract binding to access the raw methods on
}

// PermUpgrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermUpgrTransactorRaw struct {
	Contract *PermUpgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermUpgr creates a new instance of PermUpgr, bound to a specific deployed contract.
func NewPermUpgr(address common.Address, backend bind.ContractBackend) (*PermUpgr, error) {
	contract, err := bindPermUpgr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermUpgr{PermUpgrCaller: PermUpgrCaller{contract: contract}, PermUpgrTransactor: PermUpgrTransactor{contract: contract}, PermUpgrFilterer: PermUpgrFilterer{contract: contract}}, nil
}

// NewPermUpgrCaller creates a new read-only instance of PermUpgr, bound to a specific deployed contract.
func NewPermUpgrCaller(address common.Address, caller bind.ContractCaller) (*PermUpgrCaller, error) {
	contract, err := bindPermUpgr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermUpgrCaller{contract: contract}, nil
}

// NewPermUpgrTransactor creates a new write-only instance of PermUpgr, bound to a specific deployed contract.
func NewPermUpgrTransactor(address common.Address, transactor bind.ContractTransactor) (*PermUpgrTransactor, error) {
	contract, err := bindPermUpgr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermUpgrTransactor{contract: contract}, nil
}

// NewPermUpgrFilterer creates a new log filterer instance of PermUpgr, bound to a specific deployed contract.
func NewPermUpgrFilterer(address common.Address, filterer bind.ContractFilterer) (*PermUpgrFilterer, error) {
	contract, err := bindPermUpgr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermUpgrFilterer{contract: contract}, nil
}

// bindPermUpgr binds a generic wrapper to an already deployed contract.
func bindPermUpgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermUpgrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermUpgr *PermUpgrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermUpgr.Contract.PermUpgrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermUpgr *PermUpgrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermUpgr.Contract.PermUpgrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermUpgr *PermUpgrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermUpgr.Contract.PermUpgrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermUpgr *PermUpgrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermUpgr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermUpgr *PermUpgrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermUpgr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermUpgr *PermUpgrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermUpgr.Contract.contract.Transact(opts, method, params...)
}

// GetGuardian is a free data retrieval call binding the contract method 0xa75b87d2.
//
// Solidity: function getGuardian() view returns(address)
func (_PermUpgr *PermUpgrCaller) GetGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermUpgr.contract.Call(opts, &out, "getGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGuardian is a free data retrieval call binding the contract method 0xa75b87d2.
//
// Solidity: function getGuardian() view returns(address)
func (_PermUpgr *PermUpgrSession) GetGuardian() (common.Address, error) {
	return _PermUpgr.Contract.GetGuardian(&_PermUpgr.CallOpts)
}

// GetGuardian is a free data retrieval call binding the contract method 0xa75b87d2.
//
// Solidity: function getGuardian() view returns(address)
func (_PermUpgr *PermUpgrCallerSession) GetGuardian() (common.Address, error) {
	return _PermUpgr.Contract.GetGuardian(&_PermUpgr.CallOpts)
}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() view returns(address)
func (_PermUpgr *PermUpgrCaller) GetPermImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermUpgr.contract.Call(opts, &out, "getPermImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() view returns(address)
func (_PermUpgr *PermUpgrSession) GetPermImpl() (common.Address, error) {
	return _PermUpgr.Contract.GetPermImpl(&_PermUpgr.CallOpts)
}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() view returns(address)
func (_PermUpgr *PermUpgrCallerSession) GetPermImpl() (common.Address, error) {
	return _PermUpgr.Contract.GetPermImpl(&_PermUpgr.CallOpts)
}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() view returns(address)
func (_PermUpgr *PermUpgrCaller) GetPermInterface(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermUpgr.contract.Call(opts, &out, "getPermInterface")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() view returns(address)
func (_PermUpgr *PermUpgrSession) GetPermInterface() (common.Address, error) {
	return _PermUpgr.Contract.GetPermInterface(&_PermUpgr.CallOpts)
}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() view returns(address)
func (_PermUpgr *PermUpgrCallerSession) GetPermInterface() (common.Address, error) {
	return _PermUpgr.Contract.GetPermInterface(&_PermUpgr.CallOpts)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x22bcb39a.
//
// Solidity: function confirmImplChange(address _proposedImpl) returns()
func (_PermUpgr *PermUpgrTransactor) ConfirmImplChange(opts *bind.TransactOpts, _proposedImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.contract.Transact(opts, "confirmImplChange", _proposedImpl)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x22bcb39a.
//
// Solidity: function confirmImplChange(address _proposedImpl) returns()
func (_PermUpgr *PermUpgrSession) ConfirmImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.Contract.ConfirmImplChange(&_PermUpgr.TransactOpts, _proposedImpl)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x22bcb39a.
//
// Solidity: function confirmImplChange(address _proposedImpl) returns()
func (_PermUpgr *PermUpgrTransactorSession) ConfirmImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.Contract.ConfirmImplChange(&_PermUpgr.TransactOpts, _proposedImpl)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _permInterface, address _permImpl) returns()
func (_PermUpgr *PermUpgrTransactor) Init(opts *bind.TransactOpts, _permInterface common.Address, _permImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.contract.Transact(opts, "init", _permInterface, _permImpl)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _permInterface, address _permImpl) returns()
func (_PermUpgr *PermUpgrSession) Init(_permInterface common.Address, _permImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.Contract.Init(&_PermUpgr.TransactOpts, _permInterface, _permImpl)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _permInterface, address _permImpl) returns()
func (_PermUpgr *PermUpgrTransactorSession) Init(_permInterface common.Address, _permImpl common.Address) (*types.Transaction, error) {
	return _PermUpgr.Contract.Init(&_PermUpgr.TransactOpts, _permInterface, _permImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _guardian) returns()
func (_PermUpgr *PermUpgrTransactor) Initialize(opts *bind.TransactOpts, _guardian common.Address) (*types.Transaction, error) {
	return _PermUpgr.contract.Transact(opts, "initialize", _guardian)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _guardian) returns()
func (_PermUpgr *PermUpgrSession) Initialize(_guardian common.Address) (*types.Transaction, error) {
	return _PermUpgr.Contract.Initialize(&_PermUpgr.TransactOpts, _guardian)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _guardian) returns()
func (_PermUpgr *PermUpgrTransactorSession) Initialize(_guardian common.Address) (*types.Transaction, error) {
	return _PermUpgr.Contract.Initialize(&_PermUpgr.TransactOpts, _guardian)
}

// PermUpgrInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PermUpgr contract.
type PermUpgrInitializedIterator struct {
	Event *PermUpgrInitialized // Event containing the contract specifics and raw log

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
func (it *PermUpgrInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermUpgrInitialized)
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
		it.Event = new(PermUpgrInitialized)
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
func (it *PermUpgrInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermUpgrInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermUpgrInitialized represents a Initialized event raised by the PermUpgr contract.
type PermUpgrInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PermUpgr *PermUpgrFilterer) FilterInitialized(opts *bind.FilterOpts) (*PermUpgrInitializedIterator, error) {

	logs, sub, err := _PermUpgr.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PermUpgrInitializedIterator{contract: _PermUpgr.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

var InitializedTopicHash = "0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2"

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PermUpgr *PermUpgrFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PermUpgrInitialized) (event.Subscription, error) {

	logs, sub, err := _PermUpgr.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermUpgrInitialized)
				if err := _PermUpgr.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PermUpgr *PermUpgrFilterer) ParseInitialized(log types.Log) (*PermUpgrInitialized, error) {
	event := new(PermUpgrInitialized)
	if err := _PermUpgr.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
