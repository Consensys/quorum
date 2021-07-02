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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PermUpgrABI is the input ABI used to generate the binding from.
const PermUpgrABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getPermImpl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"confirmImplChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGuardian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPermInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_permInterface\",\"type\":\"address\"},{\"name\":\"_permImpl\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_guardian\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

var PermUpgrParsedABI, _ = abi.JSON(strings.NewReader(PermUpgrABI))

// PermUpgrBin is the compiled bytecode used for deploying new contracts.
var PermUpgrBin = "0x608060405234801561001057600080fd5b50604051610bfa380380610bfa8339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260146101000a81548160ff02191690831515021790555050610b4b806100af6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630e32cf901461005c57806322bcb39a146100a6578063a75b87d2146100ea578063e572515c14610134578063f09a40161461017e575b600080fd5b6100646101e2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6100e8600480360360208110156100bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061020c565b005b6100f2610639565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61013c610662565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101e06004803603604081101561019457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061068c565b005b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c69642063616c6c657200000000000000000000000000000000000081525060200191505060405180910390fd5b60608060606000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cc9ba6fa6040518163ffffffff1660e01b815260040160006040518083038186803b15801561033d57600080fd5b505afa158015610351573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250608081101561037b57600080fd5b810190808051604051939291908464010000000082111561039b57600080fd5b838201915060208201858111156103b157600080fd5b82518660018202830111640100000000821117156103ce57600080fd5b8083526020830192505050908051906020019080838360005b838110156104025780820151818401526020810190506103e7565b50505050905090810190601f16801561042f5780820380516001836020036101000a031916815260200191505b506040526020018051604051939291908464010000000082111561045257600080fd5b8382019150602082018581111561046857600080fd5b825186600182028301116401000000008211171561048557600080fd5b8083526020830192505050908051906020019080838360005b838110156104b957808201518184015260208101905061049e565b50505050905090810190601f1680156104e65780820380516001836020036101000a031916815260200191505b506040526020018051604051939291908464010000000082111561050957600080fd5b8382019150602082018581111561051f57600080fd5b825186600182028301116401000000008211171561053c57600080fd5b8083526020830192505050908051906020019080838360005b83811015610570578082015181840152602081019050610555565b50505050905090810190601f16801561059d5780820380516001836020036101000a031916815260200191505b506040526020018051906020019092919050505093509350935093506105c6858585858561089d565b84600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610632600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610a5a565b5050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461074e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f696e76616c69642063616c6c657200000000000000000000000000000000000081525060200191505060405180910390fd5b600260149054906101000a900460ff16156107d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f63616e206265206578656375746564206f6e6c79206f6e63650000000000000081525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061087e600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610a5a565b6001600260146101000a81548160ff0219169083151502179055505050565b8473ffffffffffffffffffffffffffffffffffffffff1663f5ad584a858585856040518563ffffffff1660e01b81526004018080602001806020018060200185151515158152602001848103845288818151815260200191508051906020019080838360005b8381101561091e578082015181840152602081019050610903565b50505050905090810190601f16801561094b5780820380516001836020036101000a031916815260200191505b50848103835287818151815260200191508051906020019080838360005b83811015610984578082015181840152602081019050610969565b50505050905090810190601f1680156109b15780820380516001836020036101000a031916815260200191505b50848103825286818151815260200191508051906020019080838360005b838110156109ea5780820151818401526020810190506109cf565b50505050905090810190601f168015610a175780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b158015610a3b57600080fd5b505af1158015610a4f573d6000803e3d6000fd5b505050505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663511bbd9f826040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b158015610afb57600080fd5b505af1158015610b0f573d6000803e3d6000fd5b505050505056fea265627a7a72315820316b667b04f5e53cd6b4085849e4dd3299ffca85b1af531372c992cf4c25dde464736f6c63430005110032"

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
