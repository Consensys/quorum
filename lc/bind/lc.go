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

// // IStageContractContent is an auto generated low-level Go binding around an user-defined struct.
// type IStageContractContent struct {
// 	RootHash       [32]byte
// 	SignedTime     *big.Int
// 	PrevHash       [32]byte
// 	NumOfDocuments *big.Int
// 	ContentHash    [][32]byte
// 	Url            string
// 	Acknowledge    []byte
// 	Signature      []byte
// }

// IStageContractPack is an auto generated low-level Go binding around an user-defined struct.
type IStageContractPack struct {
	Sender  common.Address
	Content IStageContractContent
}

// IStageContractStage is an auto generated low-level Go binding around an user-defined struct.
type IStageContractStage struct {
	Stage    *big.Int
	SubStage *big.Int
}

// LCABI is the input ABI used to generate the binding from.
const LCABI = "[{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_orgs\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subStage\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"documentID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"approvedTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"organization\",\"type\":\"string\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"amend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_subStage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signedTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numOfDocuments\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"contentHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"acknowledge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIStageContract.Content\",\"name\":\"_content\",\"type\":\"tuple\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stage\",\"type\":\"uint256\"}],\"name\":\"checkProposer\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_org\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_subStage\",\"type\":\"uint256\"}],\"name\":\"getContent\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signedTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numOfDocuments\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"contentHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"acknowledge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIStageContract.Content\",\"name\":\"_content\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInvolvedParties\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"getMigrateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subStage\",\"type\":\"uint256\"}],\"internalType\":\"structIStageContract.Stage[]\",\"name\":\"_stages\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signedTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numOfDocuments\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"contentHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"acknowledge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIStageContract.Content\",\"name\":\"content\",\"type\":\"tuple\"}],\"internalType\":\"structIStageContract.Pack[]\",\"name\":\"_packages\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRootHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRootList\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashToStage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subStage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"contractIFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"}],\"name\":\"setCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var LCParsedABI, _ = abi.JSON(strings.NewReader(LCABI))

// LC is an auto generated Go binding around an Ethereum contract.
type LC struct {
	LCCaller     // Read-only binding to the contract
	LCTransactor // Write-only binding to the contract
	LCFilterer   // Log filterer for contract events
}

// LCCaller is an auto generated read-only Go binding around an Ethereum contract.
type LCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LCSession struct {
	Contract     *LC               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LCCallerSession struct {
	Contract *LCCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LCTransactorSession struct {
	Contract     *LCTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LCRaw is an auto generated low-level Go binding around an Ethereum contract.
type LCRaw struct {
	Contract *LC // Generic contract binding to access the raw methods on
}

// LCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LCCallerRaw struct {
	Contract *LCCaller // Generic read-only contract binding to access the raw methods on
}

// LCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LCTransactorRaw struct {
	Contract *LCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLC creates a new instance of LC, bound to a specific deployed contract.
func NewLC(address common.Address, backend bind.ContractBackend) (*LC, error) {
	contract, err := bindLC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LC{LCCaller: LCCaller{contract: contract}, LCTransactor: LCTransactor{contract: contract}, LCFilterer: LCFilterer{contract: contract}}, nil
}

// NewLCCaller creates a new read-only instance of LC, bound to a specific deployed contract.
func NewLCCaller(address common.Address, caller bind.ContractCaller) (*LCCaller, error) {
	contract, err := bindLC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LCCaller{contract: contract}, nil
}

// NewLCTransactor creates a new write-only instance of LC, bound to a specific deployed contract.
func NewLCTransactor(address common.Address, transactor bind.ContractTransactor) (*LCTransactor, error) {
	contract, err := bindLC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LCTransactor{contract: contract}, nil
}

// NewLCFilterer creates a new log filterer instance of LC, bound to a specific deployed contract.
func NewLCFilterer(address common.Address, filterer bind.ContractFilterer) (*LCFilterer, error) {
	contract, err := bindLC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LCFilterer{contract: contract}, nil
}

// bindLC binds a generic wrapper to an already deployed contract.
func bindLC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LC *LCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LC.Contract.LCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LC *LCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LC.Contract.LCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LC *LCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LC.Contract.LCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LC *LCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LC *LCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LC *LCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LC.Contract.contract.Transact(opts, method, params...)
}

// Amended is a free data retrieval call binding the contract method 0xbaa156fc.
//
// Solidity: function amended() view returns(bool)
func (_LC *LCCaller) Amended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "amended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Amended is a free data retrieval call binding the contract method 0xbaa156fc.
//
// Solidity: function amended() view returns(bool)
func (_LC *LCSession) Amended() (bool, error) {
	return _LC.Contract.Amended(&_LC.CallOpts)
}

// Amended is a free data retrieval call binding the contract method 0xbaa156fc.
//
// Solidity: function amended() view returns(bool)
func (_LC *LCCallerSession) Amended() (bool, error) {
	return _LC.Contract.Amended(&_LC.CallOpts)
}

// CheckProposer is a free data retrieval call binding the contract method 0x7902f176.
//
// Solidity: function checkProposer(address _proposer, uint256 _stage) view returns(string _org)
func (_LC *LCCaller) CheckProposer(opts *bind.CallOpts, _proposer common.Address, _stage *big.Int) (string, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "checkProposer", _proposer, _stage)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CheckProposer is a free data retrieval call binding the contract method 0x7902f176.
//
// Solidity: function checkProposer(address _proposer, uint256 _stage) view returns(string _org)
func (_LC *LCSession) CheckProposer(_proposer common.Address, _stage *big.Int) (string, error) {
	return _LC.Contract.CheckProposer(&_LC.CallOpts, _proposer, _stage)
}

// CheckProposer is a free data retrieval call binding the contract method 0x7902f176.
//
// Solidity: function checkProposer(address _proposer, uint256 _stage) view returns(string _org)
func (_LC *LCCallerSession) CheckProposer(_proposer common.Address, _stage *big.Int) (string, error) {
	return _LC.Contract.CheckProposer(&_LC.CallOpts, _proposer, _stage)
}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_LC *LCCaller) Closed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "closed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_LC *LCSession) Closed() (bool, error) {
	return _LC.Contract.Closed(&_LC.CallOpts)
}

// Closed is a free data retrieval call binding the contract method 0x597e1fb5.
//
// Solidity: function closed() view returns(bool)
func (_LC *LCCallerSession) Closed() (bool, error) {
	return _LC.Contract.Closed(&_LC.CallOpts)
}

// GetContent is a free data retrieval call binding the contract method 0x31730a1d.
//
// Solidity: function getContent(uint256 _stage, uint256 _subStage) view returns((bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content)
func (_LC *LCCaller) GetContent(opts *bind.CallOpts, _stage *big.Int, _subStage *big.Int) (IStageContractContent, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "getContent", _stage, _subStage)

	if err != nil {
		return *new(IStageContractContent), err
	}

	out0 := *abi.ConvertType(out[0], new(IStageContractContent)).(*IStageContractContent)

	return out0, err

}

// GetContent is a free data retrieval call binding the contract method 0x31730a1d.
//
// Solidity: function getContent(uint256 _stage, uint256 _subStage) view returns((bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content)
func (_LC *LCSession) GetContent(_stage *big.Int, _subStage *big.Int) (IStageContractContent, error) {
	return _LC.Contract.GetContent(&_LC.CallOpts, _stage, _subStage)
}

// GetContent is a free data retrieval call binding the contract method 0x31730a1d.
//
// Solidity: function getContent(uint256 _stage, uint256 _subStage) view returns((bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content)
func (_LC *LCCallerSession) GetContent(_stage *big.Int, _subStage *big.Int) (IStageContractContent, error) {
	return _LC.Contract.GetContent(&_LC.CallOpts, _stage, _subStage)
}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() view returns(uint256)
func (_LC *LCCaller) GetCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "getCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() view returns(uint256)
func (_LC *LCSession) GetCounter() (*big.Int, error) {
	return _LC.Contract.GetCounter(&_LC.CallOpts)
}

// GetCounter is a free data retrieval call binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() view returns(uint256)
func (_LC *LCCallerSession) GetCounter() (*big.Int, error) {
	return _LC.Contract.GetCounter(&_LC.CallOpts)
}

// GetInvolvedParties is a free data retrieval call binding the contract method 0xf65ebf3f.
//
// Solidity: function getInvolvedParties() view returns(string[])
func (_LC *LCCaller) GetInvolvedParties(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "getInvolvedParties")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetInvolvedParties is a free data retrieval call binding the contract method 0xf65ebf3f.
//
// Solidity: function getInvolvedParties() view returns(string[])
func (_LC *LCSession) GetInvolvedParties() ([]string, error) {
	return _LC.Contract.GetInvolvedParties(&_LC.CallOpts)
}

// GetInvolvedParties is a free data retrieval call binding the contract method 0xf65ebf3f.
//
// Solidity: function getInvolvedParties() view returns(string[])
func (_LC *LCCallerSession) GetInvolvedParties() ([]string, error) {
	return _LC.Contract.GetInvolvedParties(&_LC.CallOpts)
}

// GetMigrateInfo is a free data retrieval call binding the contract method 0xb2c14ace.
//
// Solidity: function getMigrateInfo(bytes32[] _hashes) view returns((uint256,uint256)[] _stages, (address,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes))[] _packages)
func (_LC *LCCaller) GetMigrateInfo(opts *bind.CallOpts, _hashes [][32]byte) (struct {
	Stages   []IStageContractStage
	Packages []IStageContractPack
}, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "getMigrateInfo", _hashes)

	outstruct := new(struct {
		Stages   []IStageContractStage
		Packages []IStageContractPack
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stages = *abi.ConvertType(out[0], new([]IStageContractStage)).(*[]IStageContractStage)
	outstruct.Packages = *abi.ConvertType(out[1], new([]IStageContractPack)).(*[]IStageContractPack)

	return *outstruct, err

}

// GetMigrateInfo is a free data retrieval call binding the contract method 0xb2c14ace.
//
// Solidity: function getMigrateInfo(bytes32[] _hashes) view returns((uint256,uint256)[] _stages, (address,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes))[] _packages)
func (_LC *LCSession) GetMigrateInfo(_hashes [][32]byte) (struct {
	Stages   []IStageContractStage
	Packages []IStageContractPack
}, error) {
	return _LC.Contract.GetMigrateInfo(&_LC.CallOpts, _hashes)
}

// GetMigrateInfo is a free data retrieval call binding the contract method 0xb2c14ace.
//
// Solidity: function getMigrateInfo(bytes32[] _hashes) view returns((uint256,uint256)[] _stages, (address,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes))[] _packages)
func (_LC *LCCallerSession) GetMigrateInfo(_hashes [][32]byte) (struct {
	Stages   []IStageContractStage
	Packages []IStageContractPack
}, error) {
	return _LC.Contract.GetMigrateInfo(&_LC.CallOpts, _hashes)
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() view returns(bytes32)
func (_LC *LCCaller) GetRootHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "getRootHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() view returns(bytes32)
func (_LC *LCSession) GetRootHash() ([32]byte, error) {
	return _LC.Contract.GetRootHash(&_LC.CallOpts)
}

// GetRootHash is a free data retrieval call binding the contract method 0x80759f1f.
//
// Solidity: function getRootHash() view returns(bytes32)
func (_LC *LCCallerSession) GetRootHash() ([32]byte, error) {
	return _LC.Contract.GetRootHash(&_LC.CallOpts)
}

// GetRootList is a free data retrieval call binding the contract method 0x19b5c615.
//
// Solidity: function getRootList() view returns(bytes32[])
func (_LC *LCCaller) GetRootList(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "getRootList")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetRootList is a free data retrieval call binding the contract method 0x19b5c615.
//
// Solidity: function getRootList() view returns(bytes32[])
func (_LC *LCSession) GetRootList() ([][32]byte, error) {
	return _LC.Contract.GetRootList(&_LC.CallOpts)
}

// GetRootList is a free data retrieval call binding the contract method 0x19b5c615.
//
// Solidity: function getRootList() view returns(bytes32[])
func (_LC *LCCallerSession) GetRootList() ([][32]byte, error) {
	return _LC.Contract.GetRootList(&_LC.CallOpts)
}

// HashToStage is a free data retrieval call binding the contract method 0x1c207386.
//
// Solidity: function hashToStage(bytes32 ) view returns(uint256 stage, uint256 subStage)
func (_LC *LCCaller) HashToStage(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Stage    *big.Int
	SubStage *big.Int
}, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "hashToStage", arg0)

	outstruct := new(struct {
		Stage    *big.Int
		SubStage *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stage = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SubStage = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// HashToStage is a free data retrieval call binding the contract method 0x1c207386.
//
// Solidity: function hashToStage(bytes32 ) view returns(uint256 stage, uint256 subStage)
func (_LC *LCSession) HashToStage(arg0 [32]byte) (struct {
	Stage    *big.Int
	SubStage *big.Int
}, error) {
	return _LC.Contract.HashToStage(&_LC.CallOpts, arg0)
}

// HashToStage is a free data retrieval call binding the contract method 0x1c207386.
//
// Solidity: function hashToStage(bytes32 ) view returns(uint256 stage, uint256 subStage)
func (_LC *LCCallerSession) HashToStage(arg0 [32]byte) (struct {
	Stage    *big.Int
	SubStage *big.Int
}, error) {
	return _LC.Contract.HashToStage(&_LC.CallOpts, arg0)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool)
func (_LC *LCCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "isClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool)
func (_LC *LCSession) IsClosed() (bool, error) {
	return _LC.Contract.IsClosed(&_LC.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool)
func (_LC *LCCallerSession) IsClosed() (bool, error) {
	return _LC.Contract.IsClosed(&_LC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LC *LCCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LC.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LC *LCSession) Owner() (common.Address, error) {
	return _LC.Contract.Owner(&_LC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LC *LCCallerSession) Owner() (common.Address, error) {
	return _LC.Contract.Owner(&_LC.CallOpts)
}

// Amend is a paid mutator transaction binding the contract method 0x48a63d2e.
//
// Solidity: function amend() returns()
func (_LC *LCTransactor) Amend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LC.contract.Transact(opts, "amend")
}

// Amend is a paid mutator transaction binding the contract method 0x48a63d2e.
//
// Solidity: function amend() returns()
func (_LC *LCSession) Amend() (*types.Transaction, error) {
	return _LC.Contract.Amend(&_LC.TransactOpts)
}

// Amend is a paid mutator transaction binding the contract method 0x48a63d2e.
//
// Solidity: function amend() returns()
func (_LC *LCTransactorSession) Amend() (*types.Transaction, error) {
	return _LC.Contract.Amend(&_LC.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x9b1e29ba.
//
// Solidity: function approve(address _caller, uint256 _stage, uint256 _subStage, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns()
func (_LC *LCTransactor) Approve(opts *bind.TransactOpts, _caller common.Address, _stage *big.Int, _subStage *big.Int, _content IStageContractContent) (*types.Transaction, error) {
	return _LC.contract.Transact(opts, "approve", _caller, _stage, _subStage, _content)
}

// Approve is a paid mutator transaction binding the contract method 0x9b1e29ba.
//
// Solidity: function approve(address _caller, uint256 _stage, uint256 _subStage, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns()
func (_LC *LCSession) Approve(_caller common.Address, _stage *big.Int, _subStage *big.Int, _content IStageContractContent) (*types.Transaction, error) {
	return _LC.Contract.Approve(&_LC.TransactOpts, _caller, _stage, _subStage, _content)
}

// Approve is a paid mutator transaction binding the contract method 0x9b1e29ba.
//
// Solidity: function approve(address _caller, uint256 _stage, uint256 _subStage, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns()
func (_LC *LCTransactorSession) Approve(_caller common.Address, _stage *big.Int, _subStage *big.Int, _content IStageContractContent) (*types.Transaction, error) {
	return _LC.Contract.Approve(&_LC.TransactOpts, _caller, _stage, _subStage, _content)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_LC *LCTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LC.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_LC *LCSession) Close() (*types.Transaction, error) {
	return _LC.Contract.Close(&_LC.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_LC *LCTransactorSession) Close() (*types.Transaction, error) {
	return _LC.Contract.Close(&_LC.TransactOpts)
}

// SetCounter is a paid mutator transaction binding the contract method 0x8bb5d9c3.
//
// Solidity: function setCounter(uint256 _newValue) returns()
func (_LC *LCTransactor) SetCounter(opts *bind.TransactOpts, _newValue *big.Int) (*types.Transaction, error) {
	return _LC.contract.Transact(opts, "setCounter", _newValue)
}

// SetCounter is a paid mutator transaction binding the contract method 0x8bb5d9c3.
//
// Solidity: function setCounter(uint256 _newValue) returns()
func (_LC *LCSession) SetCounter(_newValue *big.Int) (*types.Transaction, error) {
	return _LC.Contract.SetCounter(&_LC.TransactOpts, _newValue)
}

// SetCounter is a paid mutator transaction binding the contract method 0x8bb5d9c3.
//
// Solidity: function setCounter(uint256 _newValue) returns()
func (_LC *LCTransactorSession) SetCounter(_newValue *big.Int) (*types.Transaction, error) {
	return _LC.Contract.SetCounter(&_LC.TransactOpts, _newValue)
}

// LCApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the LC contract.
type LCApprovedIterator struct {
	Event *LCApproved // Event containing the contract specifics and raw log

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
func (it *LCApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LCApproved)
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
		it.Event = new(LCApproved)
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
func (it *LCApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LCApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LCApproved represents a Approved event raised by the LC contract.
type LCApproved struct {
	Caller       common.Address
	Stage        *big.Int
	SubStage     *big.Int
	DocumentID   *big.Int
	ApprovedTime *big.Int
	Organization string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0x27292c739cb43063ac162d1616d542cd4a6bc6d578db97774cbb234aafc14826.
//
// Solidity: event Approved(address indexed caller, uint256 stage, uint256 subStage, uint256 indexed documentID, uint256 approvedTime, string organization)
func (_LC *LCFilterer) FilterApproved(opts *bind.FilterOpts, caller []common.Address, documentID []*big.Int) (*LCApprovedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var documentIDRule []interface{}
	for _, documentIDItem := range documentID {
		documentIDRule = append(documentIDRule, documentIDItem)
	}

	logs, sub, err := _LC.contract.FilterLogs(opts, "Approved", callerRule, documentIDRule)
	if err != nil {
		return nil, err
	}
	return &LCApprovedIterator{contract: _LC.contract, event: "Approved", logs: logs, sub: sub}, nil
}

var ApprovedTopicHash = "0x27292c739cb43063ac162d1616d542cd4a6bc6d578db97774cbb234aafc14826"

// WatchApproved is a free log subscription operation binding the contract event 0x27292c739cb43063ac162d1616d542cd4a6bc6d578db97774cbb234aafc14826.
//
// Solidity: event Approved(address indexed caller, uint256 stage, uint256 subStage, uint256 indexed documentID, uint256 approvedTime, string organization)
func (_LC *LCFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *LCApproved, caller []common.Address, documentID []*big.Int) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var documentIDRule []interface{}
	for _, documentIDItem := range documentID {
		documentIDRule = append(documentIDRule, documentIDItem)
	}

	logs, sub, err := _LC.contract.WatchLogs(opts, "Approved", callerRule, documentIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LCApproved)
				if err := _LC.contract.UnpackLog(event, "Approved", log); err != nil {
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

// ParseApproved is a log parse operation binding the contract event 0x27292c739cb43063ac162d1616d542cd4a6bc6d578db97774cbb234aafc14826.
//
// Solidity: event Approved(address indexed caller, uint256 stage, uint256 subStage, uint256 indexed documentID, uint256 approvedTime, string organization)
func (_LC *LCFilterer) ParseApproved(log types.Log) (*LCApproved, error) {
	event := new(LCApproved)
	if err := _LC.contract.UnpackLog(event, "Approved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
