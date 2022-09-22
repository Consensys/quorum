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

// ModeABI is the input ABI used to generate the binding from.
const ModeABI = "[{\"inputs\":[{\"internalType\":\"contractIPermissionsInterface\",\"name\":\"_permission\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"checkAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permission\",\"outputs\":[{\"internalType\":\"contractIPermissionsInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_permission\",\"type\":\"address\"}],\"name\":\"setPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchedToDAO\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var ModeParsedABI, _ = abi.JSON(strings.NewReader(ModeABI))

// ModeBin is the compiled bytecode used for deploying new contracts.
var ModeBin = "0x608060405234801561001057600080fd5b50604051610ac1380380610ac18339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061011a565b60008151905061008781610103565b92915050565b6000602082840312156100a3576100a26100fe565b5b60006100b184828501610078565b91505092915050565b60006100c5826100de565b9050919050565b60006100d7826100ba565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b61010c816100cc565b811461011757600080fd5b50565b610998806101296000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063b5680cb51161005b578063b5680cb5146100ee578063b85a35d2146100f8578063e73a914c14610114578063f3b0c8b7146101305761007d565b80634162169f1461008257806384e82810146100a057806389f4dd47146100be575b600080fd5b61008a61014e565b6040516100979190610747565b60405180910390f35b6100a8610174565b6040516100b59190610762565b60405180910390f35b6100d860048036038101906100d39190610634565b610187565b6040516100e59190610762565b60405180910390f35b6100f6610204565b005b610112600480360381019061010d9190610634565b610313565b005b61012e60048036038101906101299190610634565b610410565b005b610138610510565b604051610145919061077d565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160149054906101000a900460ff1681565b6000600160149054906101000a900460ff166101ab576101a682610534565b6101fd565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b9050919050565b61020d33610187565b61024c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024390610798565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600160149054906101000a900460ff166102e6576102a68173ffffffffffffffffffffffffffffffffffffffff166105e7565b6102e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102dc906107f8565b60405180910390fd5b5b600160149054906101000a900460ff1615600160146101000a81548160ff02191690831515021790555050565b61031c33610187565b61035b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035290610798565b60405180910390fd5b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156103cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c3906107b8565b60405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b61041933610187565b610458576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044f90610798565b60405180910390fd5b600160149054906101000a900460ff16156104cc5761048c8173ffffffffffffffffffffffffffffffffffffffff166105e7565b6104cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c2906107d8565b60405180910390fd5b5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d1aa0c20836040518263ffffffff1660e01b81526004016105909190610747565b60206040518083038186803b1580156105a857600080fd5b505afa1580156105bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e09190610661565b9050919050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60008135905061061981610934565b92915050565b60008151905061062e8161094b565b92915050565b60006020828403121561064a5761064961088b565b5b60006106588482850161060a565b91505092915050565b6000602082840312156106775761067661088b565b5b60006106858482850161061f565b91505092915050565b61069781610829565b82525050565b6106a68161083b565b82525050565b6106b581610867565b82525050565b60006106c8600c83610818565b91506106d382610890565b602082019050919050565b60006106eb601083610818565b91506106f6826108b9565b602082019050919050565b600061070e600f83610818565b9150610719826108e2565b602082019050919050565b6000610731601983610818565b915061073c8261090b565b602082019050919050565b600060208201905061075c600083018461068e565b92915050565b6000602082019050610777600083018461069d565b92915050565b600060208201905061079260008301846106ac565b92915050565b600060208201905081810360008301526107b1816106bb565b9050919050565b600060208201905081810360008301526107d1816106de565b9050919050565b600060208201905081810360008301526107f181610701565b9050919050565b6000602082019050818103600083015261081181610724565b9050919050565b600082825260208201905092915050565b600061083482610847565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061087282610879565b9050919050565b600061088482610847565b9050919050565b600080fd5b7f556e617574686f72697a65640000000000000000000000000000000000000000600082015250565b7f536574207a65726f206164647265737300000000000000000000000000000000600082015250565b7f496e76616c69642073657474696e670000000000000000000000000000000000600082015250565b7f556e61626c6520746f207377697463682044414f206d6f646500000000000000600082015250565b61093d81610829565b811461094857600080fd5b50565b6109548161083b565b811461095f57600080fd5b5056fea26469706673582212203bc4c71c1065ddbc365f76d2f500de4856ae5ada8f51546f38995a82684de91964736f6c63430008060033"

// DeployMode deploys a new Ethereum contract, binding an instance of Mode to it.
func DeployMode(auth *bind.TransactOpts, backend bind.ContractBackend, _permission common.Address) (common.Address, *types.Transaction, *Mode, error) {
	parsed, err := abi.JSON(strings.NewReader(ModeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ModeBin), backend, _permission)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mode{ModeCaller: ModeCaller{contract: contract}, ModeTransactor: ModeTransactor{contract: contract}, ModeFilterer: ModeFilterer{contract: contract}}, nil
}

// Mode is an auto generated Go binding around an Ethereum contract.
type Mode struct {
	ModeCaller     // Read-only binding to the contract
	ModeTransactor // Write-only binding to the contract
	ModeFilterer   // Log filterer for contract events
}

// ModeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModeSession struct {
	Contract     *Mode             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModeCallerSession struct {
	Contract *ModeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ModeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModeTransactorSession struct {
	Contract     *ModeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModeRaw struct {
	Contract *Mode // Generic contract binding to access the raw methods on
}

// ModeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModeCallerRaw struct {
	Contract *ModeCaller // Generic read-only contract binding to access the raw methods on
}

// ModeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModeTransactorRaw struct {
	Contract *ModeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMode creates a new instance of Mode, bound to a specific deployed contract.
func NewMode(address common.Address, backend bind.ContractBackend) (*Mode, error) {
	contract, err := bindMode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mode{ModeCaller: ModeCaller{contract: contract}, ModeTransactor: ModeTransactor{contract: contract}, ModeFilterer: ModeFilterer{contract: contract}}, nil
}

// NewModeCaller creates a new read-only instance of Mode, bound to a specific deployed contract.
func NewModeCaller(address common.Address, caller bind.ContractCaller) (*ModeCaller, error) {
	contract, err := bindMode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModeCaller{contract: contract}, nil
}

// NewModeTransactor creates a new write-only instance of Mode, bound to a specific deployed contract.
func NewModeTransactor(address common.Address, transactor bind.ContractTransactor) (*ModeTransactor, error) {
	contract, err := bindMode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModeTransactor{contract: contract}, nil
}

// NewModeFilterer creates a new log filterer instance of Mode, bound to a specific deployed contract.
func NewModeFilterer(address common.Address, filterer bind.ContractFilterer) (*ModeFilterer, error) {
	contract, err := bindMode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModeFilterer{contract: contract}, nil
}

// bindMode binds a generic wrapper to an already deployed contract.
func bindMode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ModeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mode *ModeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mode.Contract.ModeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mode *ModeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mode.Contract.ModeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mode *ModeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mode.Contract.ModeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mode *ModeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mode *ModeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mode *ModeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mode.Contract.contract.Transact(opts, method, params...)
}

// CheckAuthorization is a free data retrieval call binding the contract method 0x89f4dd47.
//
// Solidity: function checkAuthorization(address _operator) view returns(bool)
func (_Mode *ModeCaller) CheckAuthorization(opts *bind.CallOpts, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _Mode.contract.Call(opts, &out, "checkAuthorization", _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckAuthorization is a free data retrieval call binding the contract method 0x89f4dd47.
//
// Solidity: function checkAuthorization(address _operator) view returns(bool)
func (_Mode *ModeSession) CheckAuthorization(_operator common.Address) (bool, error) {
	return _Mode.Contract.CheckAuthorization(&_Mode.CallOpts, _operator)
}

// CheckAuthorization is a free data retrieval call binding the contract method 0x89f4dd47.
//
// Solidity: function checkAuthorization(address _operator) view returns(bool)
func (_Mode *ModeCallerSession) CheckAuthorization(_operator common.Address) (bool, error) {
	return _Mode.Contract.CheckAuthorization(&_Mode.CallOpts, _operator)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Mode *ModeCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mode.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Mode *ModeSession) Dao() (common.Address, error) {
	return _Mode.Contract.Dao(&_Mode.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Mode *ModeCallerSession) Dao() (common.Address, error) {
	return _Mode.Contract.Dao(&_Mode.CallOpts)
}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_Mode *ModeCaller) Permission(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mode.contract.Call(opts, &out, "permission")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_Mode *ModeSession) Permission() (common.Address, error) {
	return _Mode.Contract.Permission(&_Mode.CallOpts)
}

// Permission is a free data retrieval call binding the contract method 0xf3b0c8b7.
//
// Solidity: function permission() view returns(address)
func (_Mode *ModeCallerSession) Permission() (common.Address, error) {
	return _Mode.Contract.Permission(&_Mode.CallOpts)
}

// SwitchedToDAO is a free data retrieval call binding the contract method 0x84e82810.
//
// Solidity: function switchedToDAO() view returns(bool)
func (_Mode *ModeCaller) SwitchedToDAO(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mode.contract.Call(opts, &out, "switchedToDAO")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwitchedToDAO is a free data retrieval call binding the contract method 0x84e82810.
//
// Solidity: function switchedToDAO() view returns(bool)
func (_Mode *ModeSession) SwitchedToDAO() (bool, error) {
	return _Mode.Contract.SwitchedToDAO(&_Mode.CallOpts)
}

// SwitchedToDAO is a free data retrieval call binding the contract method 0x84e82810.
//
// Solidity: function switchedToDAO() view returns(bool)
func (_Mode *ModeCallerSession) SwitchedToDAO() (bool, error) {
	return _Mode.Contract.SwitchedToDAO(&_Mode.CallOpts)
}

// SetDAO is a paid mutator transaction binding the contract method 0xe73a914c.
//
// Solidity: function setDAO(address _dao) returns()
func (_Mode *ModeTransactor) SetDAO(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _Mode.contract.Transact(opts, "setDAO", _dao)
}

// SetDAO is a paid mutator transaction binding the contract method 0xe73a914c.
//
// Solidity: function setDAO(address _dao) returns()
func (_Mode *ModeSession) SetDAO(_dao common.Address) (*types.Transaction, error) {
	return _Mode.Contract.SetDAO(&_Mode.TransactOpts, _dao)
}

// SetDAO is a paid mutator transaction binding the contract method 0xe73a914c.
//
// Solidity: function setDAO(address _dao) returns()
func (_Mode *ModeTransactorSession) SetDAO(_dao common.Address) (*types.Transaction, error) {
	return _Mode.Contract.SetDAO(&_Mode.TransactOpts, _dao)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_Mode *ModeTransactor) SetPermission(opts *bind.TransactOpts, _permission common.Address) (*types.Transaction, error) {
	return _Mode.contract.Transact(opts, "setPermission", _permission)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_Mode *ModeSession) SetPermission(_permission common.Address) (*types.Transaction, error) {
	return _Mode.Contract.SetPermission(&_Mode.TransactOpts, _permission)
}

// SetPermission is a paid mutator transaction binding the contract method 0xb85a35d2.
//
// Solidity: function setPermission(address _permission) returns()
func (_Mode *ModeTransactorSession) SetPermission(_permission common.Address) (*types.Transaction, error) {
	return _Mode.Contract.SetPermission(&_Mode.TransactOpts, _permission)
}

// SwitchMode is a paid mutator transaction binding the contract method 0xb5680cb5.
//
// Solidity: function switchMode() returns()
func (_Mode *ModeTransactor) SwitchMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mode.contract.Transact(opts, "switchMode")
}

// SwitchMode is a paid mutator transaction binding the contract method 0xb5680cb5.
//
// Solidity: function switchMode() returns()
func (_Mode *ModeSession) SwitchMode() (*types.Transaction, error) {
	return _Mode.Contract.SwitchMode(&_Mode.TransactOpts)
}

// SwitchMode is a paid mutator transaction binding the contract method 0xb5680cb5.
//
// Solidity: function switchMode() returns()
func (_Mode *ModeTransactorSession) SwitchMode() (*types.Transaction, error) {
	return _Mode.Contract.SwitchMode(&_Mode.TransactOpts)
}
