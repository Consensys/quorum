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
var PermUpgrBin = "0x608060405234801561001057600080fd5b506040516020806106e78339810180604052602081101561003057600080fd5b5051600080546001600160a01b039092166001600160a01b031990921691909117905560028054600160a01b60ff0219169055610675806100726000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630e32cf901461005c57806322bcb39a14610080578063a75b87d2146100a8578063e572515c146100b0578063f09a4016146100b8575b600080fd5b6100646100e6565b604080516001600160a01b039092168252519081900360200190f35b6100a66004803603602081101561009657600080fd5b50356001600160a01b03166100f5565b005b61006461030b565b61006461031a565b6100a6600480360360408110156100ce57600080fd5b506001600160a01b0381358116916020013516610329565b6001546001600160a01b031690565b6000546001600160a01b0316331461014b5760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b60608060606000600160009054906101000a90046001600160a01b03166001600160a01b031663cc9ba6fa6040518163ffffffff1660e01b815260040160006040518083038186803b1580156101a057600080fd5b505afa1580156101b4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260808110156101dd57600080fd5b8101908080516401000000008111156101f557600080fd5b8201602081018481111561020857600080fd5b815164010000000081118282018710171561022257600080fd5b5050929190602001805164010000000081111561023e57600080fd5b8201602081018481111561025157600080fd5b815164010000000081118282018710171561026b57600080fd5b5050929190602001805164010000000081111561028757600080fd5b8201602081018481111561029a57600080fd5b81516401000000008111828201871017156102b457600080fd5b50506020909101519498509296509194509192506102d9915086905085858585610443565b600180546001600160a01b0319166001600160a01b03878116919091179182905561030491166105e4565b5050505050565b6000546001600160a01b031690565b6002546001600160a01b031690565b6000546001600160a01b0316331461037f5760408051600160e51b62461bcd02815260206004820152600e6024820152600160911b6d34b73b30b634b21031b0b63632b902604482015290519081900360640190fd5b600254600160a01b900460ff16156103e15760408051600160e51b62461bcd02815260206004820152601960248201527f63616e206265206578656375746564206f6e6c79206f6e636500000000000000604482015290519081900360640190fd5b600180546001600160a01b038084166001600160a01b031992831617928390556002805486831693169290921790915561041b91166105e4565b50506002805474ff00000000000000000000000000000000000000001916600160a01b179055565b846001600160a01b031663f5ad584a858585856040518563ffffffff1660e01b81526004018080602001806020018060200185151515158152602001848103845288818151815260200191508051906020019080838360005b838110156104b457818101518382015260200161049c565b50505050905090810190601f1680156104e15780820380516001836020036101000a031916815260200191505b50848103835287518152875160209182019189019080838360005b838110156105145781810151838201526020016104fc565b50505050905090810190601f1680156105415780820380516001836020036101000a031916815260200191505b50848103825286518152865160209182019188019080838360005b8381101561057457818101518382015260200161055c565b50505050905090810190601f1680156105a15780820380516001836020036101000a031916815260200191505b50975050505050505050600060405180830381600087803b1580156105c557600080fd5b505af11580156105d9573d6000803e3d6000fd5b505050505050505050565b60025460408051600160e01b63511bbd9f0281526001600160a01b0384811660048301529151919092169163511bbd9f91602480830192600092919082900301818387803b15801561063557600080fd5b505af1158015610304573d6000803e3d6000fdfea165627a7a7230582055489d1e43ffd1f6646b629ccf78d3fb7551dd246e111ec3ccbf9ae12f8b900a0029"

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
