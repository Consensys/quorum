// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind

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

// EeaPermUpgrABI is the input ABI used to generate the binding from.
const EeaPermUpgrABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getPermImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"confirmImplChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPermInterface\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permInterface\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_permImpl\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// EeaPermUpgrBin is the compiled bytecode used for deploying new contracts.
var EeaPermUpgrBin = "0x608060405234801561001057600080fd5b506040516107f73803806107f78339818101604052602081101561003357600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556002805460ff60a01b19169055610785806100726000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630e32cf901461005c57806322bcb39a14610080578063a75b87d2146100a8578063e572515c146100b0578063f09a4016146100b8575b600080fd5b6100646100e6565b604080516001600160a01b039092168252519081900360200190f35b6100a66004803603602081101561009657600080fd5b50356001600160a01b03166100f5565b005b61006461042f565b61006461043e565b6100a6600480360360408110156100ce57600080fd5b506001600160a01b038135811691602001351661044d565b6001546001600160a01b031690565b6000546001600160a01b03163314610145576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21031b0b63632b960911b604482015290519081900360640190fd5b60608060606000600160009054906101000a90046001600160a01b03166001600160a01b031663cc9ba6fa6040518163ffffffff1660e01b815260040160006040518083038186803b15801561019a57600080fd5b505afa1580156101ae573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260808110156101d757600080fd5b81019080805160405193929190846401000000008211156101f757600080fd5b90830190602082018581111561020c57600080fd5b825164010000000081118282018810171561022657600080fd5b82525081516020918201929091019080838360005b8381101561025357818101518382015260200161023b565b50505050905090810190601f1680156102805780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156102a357600080fd5b9083019060208201858111156102b857600080fd5b82516401000000008111828201881017156102d257600080fd5b82525081516020918201929091019080838360005b838110156102ff5781810151838201526020016102e7565b50505050905090810190601f16801561032c5780820380516001836020036101000a031916815260200191505b506040526020018051604051939291908464010000000082111561034f57600080fd5b90830190602082018581111561036457600080fd5b825164010000000081118282018810171561037e57600080fd5b82525081516020918201929091019080838360005b838110156103ab578181015183820152602001610393565b50505050905090810190601f1680156103d85780820380516001836020036101000a031916815260200191505b50604052602001519498509296509094509192506103fd91508690508585858561054d565b600180546001600160a01b0319166001600160a01b03878116919091179182905561042891166106ee565b5050505050565b6000546001600160a01b031690565b6002546001600160a01b031690565b6000546001600160a01b0316331461049d576040805162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21031b0b63632b960911b604482015290519081900360640190fd5b600254600160a01b900460ff16156104fc576040805162461bcd60e51b815260206004820152601960248201527f63616e206265206578656375746564206f6e6c79206f6e636500000000000000604482015290519081900360640190fd5b600180546001600160a01b038084166001600160a01b031992831617928390556002805486831693169290921790915561053691166106ee565b50506002805460ff60a01b1916600160a01b179055565b846001600160a01b031663f5ad584a858585856040518563ffffffff1660e01b81526004018080602001806020018060200185151515158152602001848103845288818151815260200191508051906020019080838360005b838110156105be5781810151838201526020016105a6565b50505050905090810190601f1680156105eb5780820380516001836020036101000a031916815260200191505b50848103835287518152875160209182019189019080838360005b8381101561061e578181015183820152602001610606565b50505050905090810190601f16801561064b5780820380516001836020036101000a031916815260200191505b50848103825286518152865160209182019188019080838360005b8381101561067e578181015183820152602001610666565b50505050905090810190601f1680156106ab5780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b1580156106cf57600080fd5b505af11580156106e3573d6000803e3d6000fd5b505050505050505050565b6002546040805163511bbd9f60e01b81526001600160a01b0384811660048301529151919092169163511bbd9f91602480830192600092919082900301818387803b15801561073c57600080fd5b505af1158015610428573d6000803e3d6000fdfea265627a7a723158208ce95bb3faa8908c89dfbcc8111dac5cdaa7c865298b2fb5effdebf8f89c548864736f6c634300050b0032"

// DeployEeaPermUpgr deploys a new Ethereum contract, binding an instance of EeaPermUpgr to it.
func DeployEeaPermUpgr(auth *bind.TransactOpts, backend bind.ContractBackend, _guardian common.Address) (common.Address, *types.Transaction, *EeaPermUpgr, error) {
	parsed, err := abi.JSON(strings.NewReader(EeaPermUpgrABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EeaPermUpgrBin), backend, _guardian)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EeaPermUpgr{EeaPermUpgrCaller: EeaPermUpgrCaller{contract: contract}, EeaPermUpgrTransactor: EeaPermUpgrTransactor{contract: contract}, EeaPermUpgrFilterer: EeaPermUpgrFilterer{contract: contract}}, nil
}

// EeaPermUpgr is an auto generated Go binding around an Ethereum contract.
type EeaPermUpgr struct {
	EeaPermUpgrCaller     // Read-only binding to the contract
	EeaPermUpgrTransactor // Write-only binding to the contract
	EeaPermUpgrFilterer   // Log filterer for contract events
}

// EeaPermUpgrCaller is an auto generated read-only Go binding around an Ethereum contract.
type EeaPermUpgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EeaPermUpgrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EeaPermUpgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EeaPermUpgrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EeaPermUpgrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EeaPermUpgrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EeaPermUpgrSession struct {
	Contract     *EeaPermUpgr      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EeaPermUpgrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EeaPermUpgrCallerSession struct {
	Contract *EeaPermUpgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EeaPermUpgrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EeaPermUpgrTransactorSession struct {
	Contract     *EeaPermUpgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EeaPermUpgrRaw is an auto generated low-level Go binding around an Ethereum contract.
type EeaPermUpgrRaw struct {
	Contract *EeaPermUpgr // Generic contract binding to access the raw methods on
}

// EeaPermUpgrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EeaPermUpgrCallerRaw struct {
	Contract *EeaPermUpgrCaller // Generic read-only contract binding to access the raw methods on
}

// EeaPermUpgrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EeaPermUpgrTransactorRaw struct {
	Contract *EeaPermUpgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEeaPermUpgr creates a new instance of EeaPermUpgr, bound to a specific deployed contract.
func NewEeaPermUpgr(address common.Address, backend bind.ContractBackend) (*EeaPermUpgr, error) {
	contract, err := bindEeaPermUpgr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EeaPermUpgr{EeaPermUpgrCaller: EeaPermUpgrCaller{contract: contract}, EeaPermUpgrTransactor: EeaPermUpgrTransactor{contract: contract}, EeaPermUpgrFilterer: EeaPermUpgrFilterer{contract: contract}}, nil
}

// NewEeaPermUpgrCaller creates a new read-only instance of EeaPermUpgr, bound to a specific deployed contract.
func NewEeaPermUpgrCaller(address common.Address, caller bind.ContractCaller) (*EeaPermUpgrCaller, error) {
	contract, err := bindEeaPermUpgr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EeaPermUpgrCaller{contract: contract}, nil
}

// NewEeaPermUpgrTransactor creates a new write-only instance of EeaPermUpgr, bound to a specific deployed contract.
func NewEeaPermUpgrTransactor(address common.Address, transactor bind.ContractTransactor) (*EeaPermUpgrTransactor, error) {
	contract, err := bindEeaPermUpgr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EeaPermUpgrTransactor{contract: contract}, nil
}

// NewEeaPermUpgrFilterer creates a new log filterer instance of EeaPermUpgr, bound to a specific deployed contract.
func NewEeaPermUpgrFilterer(address common.Address, filterer bind.ContractFilterer) (*EeaPermUpgrFilterer, error) {
	contract, err := bindEeaPermUpgr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EeaPermUpgrFilterer{contract: contract}, nil
}

// bindEeaPermUpgr binds a generic wrapper to an already deployed contract.
func bindEeaPermUpgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EeaPermUpgrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EeaPermUpgr *EeaPermUpgrRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EeaPermUpgr.Contract.EeaPermUpgrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EeaPermUpgr *EeaPermUpgrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EeaPermUpgr.Contract.EeaPermUpgrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EeaPermUpgr *EeaPermUpgrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EeaPermUpgr.Contract.EeaPermUpgrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EeaPermUpgr *EeaPermUpgrCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EeaPermUpgr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EeaPermUpgr *EeaPermUpgrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EeaPermUpgr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EeaPermUpgr *EeaPermUpgrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EeaPermUpgr.Contract.contract.Transact(opts, method, params...)
}

// GetGuardian is a free data retrieval call binding the contract method 0xa75b87d2.
//
// Solidity: function getGuardian() constant returns(address)
func (_EeaPermUpgr *EeaPermUpgrCaller) GetGuardian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EeaPermUpgr.contract.Call(opts, out, "getGuardian")
	return *ret0, err
}

// GetGuardian is a free data retrieval call binding the contract method 0xa75b87d2.
//
// Solidity: function getGuardian() constant returns(address)
func (_EeaPermUpgr *EeaPermUpgrSession) GetGuardian() (common.Address, error) {
	return _EeaPermUpgr.Contract.GetGuardian(&_EeaPermUpgr.CallOpts)
}

// GetGuardian is a free data retrieval call binding the contract method 0xa75b87d2.
//
// Solidity: function getGuardian() constant returns(address)
func (_EeaPermUpgr *EeaPermUpgrCallerSession) GetGuardian() (common.Address, error) {
	return _EeaPermUpgr.Contract.GetGuardian(&_EeaPermUpgr.CallOpts)
}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() constant returns(address)
func (_EeaPermUpgr *EeaPermUpgrCaller) GetPermImpl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EeaPermUpgr.contract.Call(opts, out, "getPermImpl")
	return *ret0, err
}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() constant returns(address)
func (_EeaPermUpgr *EeaPermUpgrSession) GetPermImpl() (common.Address, error) {
	return _EeaPermUpgr.Contract.GetPermImpl(&_EeaPermUpgr.CallOpts)
}

// GetPermImpl is a free data retrieval call binding the contract method 0x0e32cf90.
//
// Solidity: function getPermImpl() constant returns(address)
func (_EeaPermUpgr *EeaPermUpgrCallerSession) GetPermImpl() (common.Address, error) {
	return _EeaPermUpgr.Contract.GetPermImpl(&_EeaPermUpgr.CallOpts)
}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() constant returns(address)
func (_EeaPermUpgr *EeaPermUpgrCaller) GetPermInterface(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EeaPermUpgr.contract.Call(opts, out, "getPermInterface")
	return *ret0, err
}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() constant returns(address)
func (_EeaPermUpgr *EeaPermUpgrSession) GetPermInterface() (common.Address, error) {
	return _EeaPermUpgr.Contract.GetPermInterface(&_EeaPermUpgr.CallOpts)
}

// GetPermInterface is a free data retrieval call binding the contract method 0xe572515c.
//
// Solidity: function getPermInterface() constant returns(address)
func (_EeaPermUpgr *EeaPermUpgrCallerSession) GetPermInterface() (common.Address, error) {
	return _EeaPermUpgr.Contract.GetPermInterface(&_EeaPermUpgr.CallOpts)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x22bcb39a.
//
// Solidity: function confirmImplChange(address _proposedImpl) returns()
func (_EeaPermUpgr *EeaPermUpgrTransactor) ConfirmImplChange(opts *bind.TransactOpts, _proposedImpl common.Address) (*types.Transaction, error) {
	return _EeaPermUpgr.contract.Transact(opts, "confirmImplChange", _proposedImpl)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x22bcb39a.
//
// Solidity: function confirmImplChange(address _proposedImpl) returns()
func (_EeaPermUpgr *EeaPermUpgrSession) ConfirmImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _EeaPermUpgr.Contract.ConfirmImplChange(&_EeaPermUpgr.TransactOpts, _proposedImpl)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x22bcb39a.
//
// Solidity: function confirmImplChange(address _proposedImpl) returns()
func (_EeaPermUpgr *EeaPermUpgrTransactorSession) ConfirmImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _EeaPermUpgr.Contract.ConfirmImplChange(&_EeaPermUpgr.TransactOpts, _proposedImpl)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _permInterface, address _permImpl) returns()
func (_EeaPermUpgr *EeaPermUpgrTransactor) Init(opts *bind.TransactOpts, _permInterface common.Address, _permImpl common.Address) (*types.Transaction, error) {
	return _EeaPermUpgr.contract.Transact(opts, "init", _permInterface, _permImpl)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _permInterface, address _permImpl) returns()
func (_EeaPermUpgr *EeaPermUpgrSession) Init(_permInterface common.Address, _permImpl common.Address) (*types.Transaction, error) {
	return _EeaPermUpgr.Contract.Init(&_EeaPermUpgr.TransactOpts, _permInterface, _permImpl)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _permInterface, address _permImpl) returns()
func (_EeaPermUpgr *EeaPermUpgrTransactorSession) Init(_permInterface common.Address, _permImpl common.Address) (*types.Transaction, error) {
	return _EeaPermUpgr.Contract.Init(&_EeaPermUpgr.TransactOpts, _permInterface, _permImpl)
}
