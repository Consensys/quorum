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

// PermUpgrABI is the input ABI used to generate the binding from.
const PermUpgrABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getPermImpl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"confirmImplChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGuardian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPermInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_permInterface\",\"type\":\"address\"},{\"name\":\"_permImpl\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_guardian\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

var PermUpgrParsedABI, _ = abi.JSON(strings.NewReader(PermUpgrABI))

// PermUpgrBin is the compiled bytecode used for deploying new contracts.
var PermUpgrBin = "0x608060405234801561001057600080fd5b50604051602080610abc8339810180604052602081101561003057600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260146101000a81548160ff02191690831515021790555050610a10806100ac6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630e32cf901461005c57806322bcb39a146100a6578063a75b87d2146100ea578063e572515c14610134578063f09a40161461017e575b600080fd5b6100646101e2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100e8600480360360208110156100bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061020c565b005b6100f2610503565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61013c61052c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101e06004803603604081101561019457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610556565b005b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156102d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c69642063616c6c657200000000000000000000000000000000000081525060200191505060405180910390fd5b60608060606000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cc9ba6fa6040518163ffffffff1660e01b815260040160006040518083038186803b15801561033f57600080fd5b505afa158015610353573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250608081101561037d57600080fd5b81019080805164010000000081111561039557600080fd5b828101905060208101848111156103ab57600080fd5b81518560018202830111640100000000821117156103c857600080fd5b505092919060200180516401000000008111156103e457600080fd5b828101905060208101848111156103fa57600080fd5b815185600182028301116401000000008211171561041757600080fd5b5050929190602001805164010000000081111561043357600080fd5b8281019050602081018481111561044957600080fd5b815185600182028301116401000000008211171561046657600080fd5b5050929190602001805190602001909291905050509350935093509350610490858585858561076b565b84600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506104fc600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610928565b5050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561061a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c69642063616c6c657200000000000000000000000000000000000081525060200191505060405180910390fd5b600260149054906101000a900460ff1615151561069f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616e206265206578656375746564206f6e6c79206f6e63650000000000000081525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061074c600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610928565b6001600260146101000a81548160ff0219169083151502179055505050565b8473ffffffffffffffffffffffffffffffffffffffff1663f5ad584a858585856040518563ffffffff1660e01b81526004018080602001806020018060200185151515158152602001848103845288818151815260200191508051906020019080838360005b838110156107ec5780820151818401526020810190506107d1565b50505050905090810190601f1680156108195780820380516001836020036101000a031916815260200191505b50848103835287818151815260200191508051906020019080838360005b83811015610852578082015181840152602081019050610837565b50505050905090810190601f16801561087f5780820380516001836020036101000a031916815260200191505b50848103825286818151815260200191508051906020019080838360005b838110156108b857808201518184015260208101905061089d565b50505050905090810190601f1680156108e55780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b15801561090957600080fd5b505af115801561091d573d6000803e3d6000fd5b505050505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663511bbd9f826040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b1580156109c957600080fd5b505af11580156109dd573d6000803e3d6000fd5b505050505056fea165627a7a72305820f908a2c3502835d6af50d8b34c9530ace90aae30e51acebeac5d74b7615fe2580029"

// DeployPermUpgr deploys a new Ethereum contract, binding an instance of PermUpgr to it.
func DeployPermUpgr(auth *bind.TransactOpts, backend bind.ContractBackend, _guardian common.Address) (common.Address, *types.Transaction, *PermUpgr, error) {
	parsed, err := abi.JSON(strings.NewReader(PermUpgrABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PermUpgrBin), backend, _guardian)
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
func (_PermUpgr *PermUpgrRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_PermUpgr *PermUpgrCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function getGuardian() constant returns(address)
func (_PermUpgr *PermUpgrCaller) GetGuardian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PermUpgr.contract.Call(opts, out, "getGuardian")
	return *ret0, err
}

// GetGuardian is a free data retrieval call binding the contract method 0xa75b87d2.
//
// Solidity: function getGuardian() constant returns(address)
func (_PermUpgr *PermUpgrSession) GetGuardian() (common.Address, error) {
	return _PermUpgr.Contract.GetGuardian(&_PermUpgr.CallOpts)
}

// GetGuardian is a free data retrieval call binding the contract method 0xa75b87d2.
//
// Solidity: function getGuardian() constant returns(address)
func (_PermUpgr *PermUpgrCallerSession) GetGuardian() (common.Address, error) {
	return _PermUpgr.Contract.GetGuardian(&_PermUpgr.CallOpts)
}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() constant returns(address)
func (_PermUpgr *PermUpgrCaller) GetPermImpl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PermUpgr.contract.Call(opts, out, "getPermImpl")
	return *ret0, err
}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() constant returns(address)
func (_PermUpgr *PermUpgrSession) GetPermImpl() (common.Address, error) {
	return _PermUpgr.Contract.GetPermImpl(&_PermUpgr.CallOpts)
}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() constant returns(address)
func (_PermUpgr *PermUpgrCallerSession) GetPermImpl() (common.Address, error) {
	return _PermUpgr.Contract.GetPermImpl(&_PermUpgr.CallOpts)
}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() constant returns(address)
func (_PermUpgr *PermUpgrCaller) GetPermInterface(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PermUpgr.contract.Call(opts, out, "getPermInterface")
	return *ret0, err
}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() constant returns(address)
func (_PermUpgr *PermUpgrSession) GetPermInterface() (common.Address, error) {
	return _PermUpgr.Contract.GetPermInterface(&_PermUpgr.CallOpts)
}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() constant returns(address)
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
