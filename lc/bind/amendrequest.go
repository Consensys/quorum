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

// IAmendRequestAmendStage is an auto generated low-level Go binding around an user-defined struct.
type IAmendRequestAmendStage struct {
	Stage    *big.Int
	SubStage *big.Int
	Content  IStageContractContent
}

// IAmendRequestConfirmation is an auto generated low-level Go binding around an user-defined struct.
type IAmendRequestConfirmation struct {
	IssuingBank        string
	AdvisingBank       string
	ReimbursingBank    string
	IssuingBankSig     []byte
	AdvisingBankSig    []byte
	ReimbursingBankSig []byte
}

// IAmendRequestRequest is an auto generated low-level Go binding around an user-defined struct.
type IAmendRequestRequest struct {
	TypeOf          *big.Int
	Proposer        common.Address
	MigratingStages [][32]byte
	AmendStage      IAmendRequestAmendStage
	Confirmed       IAmendRequestConfirmation
	IsFulfilled     bool
}

// IStageContractContent is an auto generated low-level Go binding around an user-defined struct.
type IStageContractContent struct {
	RootHash       [32]byte
	SignedTime     *big.Int
	PrevHash       [32]byte
	NumOfDocuments *big.Int
	ContentHash    [][32]byte
	Url            string
	Acknowledge    []byte
	Signature      []byte
}

// AmendRequestABI is the input ABI used to generate the binding from.
const AmendRequestABI = "[{\"inputs\":[{\"internalType\":\"contractILCManagement\",\"name\":\"_management\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"documentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApprovedAmendment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"documentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"SubmittedAmendment\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_approver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"fulfilled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getAmendRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"typeOf\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"migratingStages\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subStage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signedTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numOfDocuments\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"contentHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"acknowledge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIStageContract.Content\",\"name\":\"content\",\"type\":\"tuple\"}],\"internalType\":\"structIAmendRequest.AmendStage\",\"name\":\"amendStage\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"issuingBank\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"advisingBank\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reimbursingBank\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"issuingBankSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"advisingBankSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reimbursingBankSig\",\"type\":\"bytes\"}],\"internalType\":\"structIAmendRequest.Confirmation\",\"name\":\"confirmed\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isFulfilled\",\"type\":\"bool\"}],\"internalType\":\"structIAmendRequest.Request\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"isFulfilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"isProposer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"management\",\"outputs\":[{\"internalType\":\"contractILCManagement\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILCManagement\",\"name\":\"_management\",\"type\":\"address\"}],\"name\":\"setLCManagement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"_migratingStages\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subStage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signedTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numOfDocuments\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"contentHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"acknowledge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIStageContract.Content\",\"name\":\"content\",\"type\":\"tuple\"}],\"internalType\":\"structIAmendRequest.AmendStage\",\"name\":\"_amendStage\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var AmendRequestParsedABI, _ = abi.JSON(strings.NewReader(AmendRequestABI))

// AmendRequest is an auto generated Go binding around an Ethereum contract.
type AmendRequest struct {
	AmendRequestCaller     // Read-only binding to the contract
	AmendRequestTransactor // Write-only binding to the contract
	AmendRequestFilterer   // Log filterer for contract events
}

// AmendRequestCaller is an auto generated read-only Go binding around an Ethereum contract.
type AmendRequestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmendRequestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AmendRequestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmendRequestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AmendRequestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmendRequestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AmendRequestSession struct {
	Contract     *AmendRequest     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AmendRequestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AmendRequestCallerSession struct {
	Contract *AmendRequestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AmendRequestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AmendRequestTransactorSession struct {
	Contract     *AmendRequestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AmendRequestRaw is an auto generated low-level Go binding around an Ethereum contract.
type AmendRequestRaw struct {
	Contract *AmendRequest // Generic contract binding to access the raw methods on
}

// AmendRequestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AmendRequestCallerRaw struct {
	Contract *AmendRequestCaller // Generic read-only contract binding to access the raw methods on
}

// AmendRequestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AmendRequestTransactorRaw struct {
	Contract *AmendRequestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAmendRequest creates a new instance of AmendRequest, bound to a specific deployed contract.
func NewAmendRequest(address common.Address, backend bind.ContractBackend) (*AmendRequest, error) {
	contract, err := bindAmendRequest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AmendRequest{AmendRequestCaller: AmendRequestCaller{contract: contract}, AmendRequestTransactor: AmendRequestTransactor{contract: contract}, AmendRequestFilterer: AmendRequestFilterer{contract: contract}}, nil
}

// NewAmendRequestCaller creates a new read-only instance of AmendRequest, bound to a specific deployed contract.
func NewAmendRequestCaller(address common.Address, caller bind.ContractCaller) (*AmendRequestCaller, error) {
	contract, err := bindAmendRequest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AmendRequestCaller{contract: contract}, nil
}

// NewAmendRequestTransactor creates a new write-only instance of AmendRequest, bound to a specific deployed contract.
func NewAmendRequestTransactor(address common.Address, transactor bind.ContractTransactor) (*AmendRequestTransactor, error) {
	contract, err := bindAmendRequest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AmendRequestTransactor{contract: contract}, nil
}

// NewAmendRequestFilterer creates a new log filterer instance of AmendRequest, bound to a specific deployed contract.
func NewAmendRequestFilterer(address common.Address, filterer bind.ContractFilterer) (*AmendRequestFilterer, error) {
	contract, err := bindAmendRequest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AmendRequestFilterer{contract: contract}, nil
}

// bindAmendRequest binds a generic wrapper to an already deployed contract.
func bindAmendRequest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AmendRequestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmendRequest *AmendRequestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmendRequest.Contract.AmendRequestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmendRequest *AmendRequestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmendRequest.Contract.AmendRequestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmendRequest *AmendRequestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmendRequest.Contract.AmendRequestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmendRequest *AmendRequestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmendRequest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmendRequest *AmendRequestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmendRequest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmendRequest *AmendRequestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmendRequest.Contract.contract.Transact(opts, method, params...)
}

// GetAmendRequest is a free data retrieval call binding the contract method 0xa6497af9.
//
// Solidity: function getAmendRequest(uint256 _documentId, uint256 _requestId) view returns((uint256,address,bytes32[],(uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)),(string,string,string,bytes,bytes,bytes),bool))
func (_AmendRequest *AmendRequestCaller) GetAmendRequest(opts *bind.CallOpts, _documentId *big.Int, _requestId *big.Int) (IAmendRequestRequest, error) {
	var out []interface{}
	err := _AmendRequest.contract.Call(opts, &out, "getAmendRequest", _documentId, _requestId)

	if err != nil {
		return *new(IAmendRequestRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IAmendRequestRequest)).(*IAmendRequestRequest)

	return out0, err

}

// GetAmendRequest is a free data retrieval call binding the contract method 0xa6497af9.
//
// Solidity: function getAmendRequest(uint256 _documentId, uint256 _requestId) view returns((uint256,address,bytes32[],(uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)),(string,string,string,bytes,bytes,bytes),bool))
func (_AmendRequest *AmendRequestSession) GetAmendRequest(_documentId *big.Int, _requestId *big.Int) (IAmendRequestRequest, error) {
	return _AmendRequest.Contract.GetAmendRequest(&_AmendRequest.CallOpts, _documentId, _requestId)
}

// GetAmendRequest is a free data retrieval call binding the contract method 0xa6497af9.
//
// Solidity: function getAmendRequest(uint256 _documentId, uint256 _requestId) view returns((uint256,address,bytes32[],(uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)),(string,string,string,bytes,bytes,bytes),bool))
func (_AmendRequest *AmendRequestCallerSession) GetAmendRequest(_documentId *big.Int, _requestId *big.Int) (IAmendRequestRequest, error) {
	return _AmendRequest.Contract.GetAmendRequest(&_AmendRequest.CallOpts, _documentId, _requestId)
}

// IsApproved is a free data retrieval call binding the contract method 0xbf276511.
//
// Solidity: function isApproved(uint256 _documentId, uint256 _requestId) view returns(bool)
func (_AmendRequest *AmendRequestCaller) IsApproved(opts *bind.CallOpts, _documentId *big.Int, _requestId *big.Int) (bool, error) {
	var out []interface{}
	err := _AmendRequest.contract.Call(opts, &out, "isApproved", _documentId, _requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0xbf276511.
//
// Solidity: function isApproved(uint256 _documentId, uint256 _requestId) view returns(bool)
func (_AmendRequest *AmendRequestSession) IsApproved(_documentId *big.Int, _requestId *big.Int) (bool, error) {
	return _AmendRequest.Contract.IsApproved(&_AmendRequest.CallOpts, _documentId, _requestId)
}

// IsApproved is a free data retrieval call binding the contract method 0xbf276511.
//
// Solidity: function isApproved(uint256 _documentId, uint256 _requestId) view returns(bool)
func (_AmendRequest *AmendRequestCallerSession) IsApproved(_documentId *big.Int, _requestId *big.Int) (bool, error) {
	return _AmendRequest.Contract.IsApproved(&_AmendRequest.CallOpts, _documentId, _requestId)
}

// IsFulfilled is a free data retrieval call binding the contract method 0x9f3fc6d3.
//
// Solidity: function isFulfilled(uint256 _documentId, uint256 _requestId) view returns(bool)
func (_AmendRequest *AmendRequestCaller) IsFulfilled(opts *bind.CallOpts, _documentId *big.Int, _requestId *big.Int) (bool, error) {
	var out []interface{}
	err := _AmendRequest.contract.Call(opts, &out, "isFulfilled", _documentId, _requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFulfilled is a free data retrieval call binding the contract method 0x9f3fc6d3.
//
// Solidity: function isFulfilled(uint256 _documentId, uint256 _requestId) view returns(bool)
func (_AmendRequest *AmendRequestSession) IsFulfilled(_documentId *big.Int, _requestId *big.Int) (bool, error) {
	return _AmendRequest.Contract.IsFulfilled(&_AmendRequest.CallOpts, _documentId, _requestId)
}

// IsFulfilled is a free data retrieval call binding the contract method 0x9f3fc6d3.
//
// Solidity: function isFulfilled(uint256 _documentId, uint256 _requestId) view returns(bool)
func (_AmendRequest *AmendRequestCallerSession) IsFulfilled(_documentId *big.Int, _requestId *big.Int) (bool, error) {
	return _AmendRequest.Contract.IsFulfilled(&_AmendRequest.CallOpts, _documentId, _requestId)
}

// IsProposer is a free data retrieval call binding the contract method 0x856b909b.
//
// Solidity: function isProposer(uint256 _documentId, uint256 _requestId, address _executor) view returns(bool)
func (_AmendRequest *AmendRequestCaller) IsProposer(opts *bind.CallOpts, _documentId *big.Int, _requestId *big.Int, _executor common.Address) (bool, error) {
	var out []interface{}
	err := _AmendRequest.contract.Call(opts, &out, "isProposer", _documentId, _requestId, _executor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProposer is a free data retrieval call binding the contract method 0x856b909b.
//
// Solidity: function isProposer(uint256 _documentId, uint256 _requestId, address _executor) view returns(bool)
func (_AmendRequest *AmendRequestSession) IsProposer(_documentId *big.Int, _requestId *big.Int, _executor common.Address) (bool, error) {
	return _AmendRequest.Contract.IsProposer(&_AmendRequest.CallOpts, _documentId, _requestId, _executor)
}

// IsProposer is a free data retrieval call binding the contract method 0x856b909b.
//
// Solidity: function isProposer(uint256 _documentId, uint256 _requestId, address _executor) view returns(bool)
func (_AmendRequest *AmendRequestCallerSession) IsProposer(_documentId *big.Int, _requestId *big.Int, _executor common.Address) (bool, error) {
	return _AmendRequest.Contract.IsProposer(&_AmendRequest.CallOpts, _documentId, _requestId, _executor)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_AmendRequest *AmendRequestCaller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AmendRequest.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_AmendRequest *AmendRequestSession) Management() (common.Address, error) {
	return _AmendRequest.Contract.Management(&_AmendRequest.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_AmendRequest *AmendRequestCallerSession) Management() (common.Address, error) {
	return _AmendRequest.Contract.Management(&_AmendRequest.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_AmendRequest *AmendRequestCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AmendRequest.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_AmendRequest *AmendRequestSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _AmendRequest.Contract.Nonces(&_AmendRequest.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_AmendRequest *AmendRequestCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _AmendRequest.Contract.Nonces(&_AmendRequest.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0xa687be52.
//
// Solidity: function approve(uint256 _documentId, uint256 _requestId, address _approver, bytes _signature) returns()
func (_AmendRequest *AmendRequestTransactor) Approve(opts *bind.TransactOpts, _documentId *big.Int, _requestId *big.Int, _approver common.Address, _signature []byte) (*types.Transaction, error) {
	return _AmendRequest.contract.Transact(opts, "approve", _documentId, _requestId, _approver, _signature)
}

// Approve is a paid mutator transaction binding the contract method 0xa687be52.
//
// Solidity: function approve(uint256 _documentId, uint256 _requestId, address _approver, bytes _signature) returns()
func (_AmendRequest *AmendRequestSession) Approve(_documentId *big.Int, _requestId *big.Int, _approver common.Address, _signature []byte) (*types.Transaction, error) {
	return _AmendRequest.Contract.Approve(&_AmendRequest.TransactOpts, _documentId, _requestId, _approver, _signature)
}

// Approve is a paid mutator transaction binding the contract method 0xa687be52.
//
// Solidity: function approve(uint256 _documentId, uint256 _requestId, address _approver, bytes _signature) returns()
func (_AmendRequest *AmendRequestTransactorSession) Approve(_documentId *big.Int, _requestId *big.Int, _approver common.Address, _signature []byte) (*types.Transaction, error) {
	return _AmendRequest.Contract.Approve(&_AmendRequest.TransactOpts, _documentId, _requestId, _approver, _signature)
}

// Fulfilled is a paid mutator transaction binding the contract method 0x00ebf111.
//
// Solidity: function fulfilled(uint256 _documentId, uint256 _requestId) returns()
func (_AmendRequest *AmendRequestTransactor) Fulfilled(opts *bind.TransactOpts, _documentId *big.Int, _requestId *big.Int) (*types.Transaction, error) {
	return _AmendRequest.contract.Transact(opts, "fulfilled", _documentId, _requestId)
}

// Fulfilled is a paid mutator transaction binding the contract method 0x00ebf111.
//
// Solidity: function fulfilled(uint256 _documentId, uint256 _requestId) returns()
func (_AmendRequest *AmendRequestSession) Fulfilled(_documentId *big.Int, _requestId *big.Int) (*types.Transaction, error) {
	return _AmendRequest.Contract.Fulfilled(&_AmendRequest.TransactOpts, _documentId, _requestId)
}

// Fulfilled is a paid mutator transaction binding the contract method 0x00ebf111.
//
// Solidity: function fulfilled(uint256 _documentId, uint256 _requestId) returns()
func (_AmendRequest *AmendRequestTransactorSession) Fulfilled(_documentId *big.Int, _requestId *big.Int) (*types.Transaction, error) {
	return _AmendRequest.Contract.Fulfilled(&_AmendRequest.TransactOpts, _documentId, _requestId)
}

// SetLCManagement is a paid mutator transaction binding the contract method 0xb3463971.
//
// Solidity: function setLCManagement(address _management) returns()
func (_AmendRequest *AmendRequestTransactor) SetLCManagement(opts *bind.TransactOpts, _management common.Address) (*types.Transaction, error) {
	return _AmendRequest.contract.Transact(opts, "setLCManagement", _management)
}

// SetLCManagement is a paid mutator transaction binding the contract method 0xb3463971.
//
// Solidity: function setLCManagement(address _management) returns()
func (_AmendRequest *AmendRequestSession) SetLCManagement(_management common.Address) (*types.Transaction, error) {
	return _AmendRequest.Contract.SetLCManagement(&_AmendRequest.TransactOpts, _management)
}

// SetLCManagement is a paid mutator transaction binding the contract method 0xb3463971.
//
// Solidity: function setLCManagement(address _management) returns()
func (_AmendRequest *AmendRequestTransactorSession) SetLCManagement(_management common.Address) (*types.Transaction, error) {
	return _AmendRequest.Contract.SetLCManagement(&_AmendRequest.TransactOpts, _management)
}

// Submit is a paid mutator transaction binding the contract method 0xdf879220.
//
// Solidity: function submit(uint256 _documentId, address _proposer, bytes32[] _migratingStages, (uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)) _amendStage, bytes _signature) returns()
func (_AmendRequest *AmendRequestTransactor) Submit(opts *bind.TransactOpts, _documentId *big.Int, _proposer common.Address, _migratingStages [][32]byte, _amendStage IAmendRequestAmendStage, _signature []byte) (*types.Transaction, error) {
	return _AmendRequest.contract.Transact(opts, "submit", _documentId, _proposer, _migratingStages, _amendStage, _signature)
}

// Submit is a paid mutator transaction binding the contract method 0xdf879220.
//
// Solidity: function submit(uint256 _documentId, address _proposer, bytes32[] _migratingStages, (uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)) _amendStage, bytes _signature) returns()
func (_AmendRequest *AmendRequestSession) Submit(_documentId *big.Int, _proposer common.Address, _migratingStages [][32]byte, _amendStage IAmendRequestAmendStage, _signature []byte) (*types.Transaction, error) {
	return _AmendRequest.Contract.Submit(&_AmendRequest.TransactOpts, _documentId, _proposer, _migratingStages, _amendStage, _signature)
}

// Submit is a paid mutator transaction binding the contract method 0xdf879220.
//
// Solidity: function submit(uint256 _documentId, address _proposer, bytes32[] _migratingStages, (uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)) _amendStage, bytes _signature) returns()
func (_AmendRequest *AmendRequestTransactorSession) Submit(_documentId *big.Int, _proposer common.Address, _migratingStages [][32]byte, _amendStage IAmendRequestAmendStage, _signature []byte) (*types.Transaction, error) {
	return _AmendRequest.Contract.Submit(&_AmendRequest.TransactOpts, _documentId, _proposer, _migratingStages, _amendStage, _signature)
}

// AmendRequestApprovedAmendmentIterator is returned from FilterApprovedAmendment and is used to iterate over the raw logs and unpacked data for ApprovedAmendment events raised by the AmendRequest contract.
type AmendRequestApprovedAmendmentIterator struct {
	Event *AmendRequestApprovedAmendment // Event containing the contract specifics and raw log

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
func (it *AmendRequestApprovedAmendmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmendRequestApprovedAmendment)
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
		it.Event = new(AmendRequestApprovedAmendment)
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
func (it *AmendRequestApprovedAmendmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmendRequestApprovedAmendmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmendRequestApprovedAmendment represents a ApprovedAmendment event raised by the AmendRequest contract.
type AmendRequestApprovedAmendment struct {
	DocumentId *big.Int
	RequestId  *big.Int
	Approver   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApprovedAmendment is a free log retrieval operation binding the contract event 0xe8bf39464f4f533619fa362b4e80c90fba7d4527115ffe466472fdc52f977109.
//
// Solidity: event ApprovedAmendment(uint256 indexed documentId, uint256 indexed requestId, address indexed approver)
func (_AmendRequest *AmendRequestFilterer) FilterApprovedAmendment(opts *bind.FilterOpts, documentId []*big.Int, requestId []*big.Int, approver []common.Address) (*AmendRequestApprovedAmendmentIterator, error) {

	var documentIdRule []interface{}
	for _, documentIdItem := range documentId {
		documentIdRule = append(documentIdRule, documentIdItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _AmendRequest.contract.FilterLogs(opts, "ApprovedAmendment", documentIdRule, requestIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return &AmendRequestApprovedAmendmentIterator{contract: _AmendRequest.contract, event: "ApprovedAmendment", logs: logs, sub: sub}, nil
}

var ApprovedAmendmentTopicHash = "0xe8bf39464f4f533619fa362b4e80c90fba7d4527115ffe466472fdc52f977109"

// WatchApprovedAmendment is a free log subscription operation binding the contract event 0xe8bf39464f4f533619fa362b4e80c90fba7d4527115ffe466472fdc52f977109.
//
// Solidity: event ApprovedAmendment(uint256 indexed documentId, uint256 indexed requestId, address indexed approver)
func (_AmendRequest *AmendRequestFilterer) WatchApprovedAmendment(opts *bind.WatchOpts, sink chan<- *AmendRequestApprovedAmendment, documentId []*big.Int, requestId []*big.Int, approver []common.Address) (event.Subscription, error) {

	var documentIdRule []interface{}
	for _, documentIdItem := range documentId {
		documentIdRule = append(documentIdRule, documentIdItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _AmendRequest.contract.WatchLogs(opts, "ApprovedAmendment", documentIdRule, requestIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmendRequestApprovedAmendment)
				if err := _AmendRequest.contract.UnpackLog(event, "ApprovedAmendment", log); err != nil {
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

// ParseApprovedAmendment is a log parse operation binding the contract event 0xe8bf39464f4f533619fa362b4e80c90fba7d4527115ffe466472fdc52f977109.
//
// Solidity: event ApprovedAmendment(uint256 indexed documentId, uint256 indexed requestId, address indexed approver)
func (_AmendRequest *AmendRequestFilterer) ParseApprovedAmendment(log types.Log) (*AmendRequestApprovedAmendment, error) {
	event := new(AmendRequestApprovedAmendment)
	if err := _AmendRequest.contract.UnpackLog(event, "ApprovedAmendment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmendRequestSubmittedAmendmentIterator is returned from FilterSubmittedAmendment and is used to iterate over the raw logs and unpacked data for SubmittedAmendment events raised by the AmendRequest contract.
type AmendRequestSubmittedAmendmentIterator struct {
	Event *AmendRequestSubmittedAmendment // Event containing the contract specifics and raw log

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
func (it *AmendRequestSubmittedAmendmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmendRequestSubmittedAmendment)
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
		it.Event = new(AmendRequestSubmittedAmendment)
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
func (it *AmendRequestSubmittedAmendmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmendRequestSubmittedAmendmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmendRequestSubmittedAmendment represents a SubmittedAmendment event raised by the AmendRequest contract.
type AmendRequestSubmittedAmendment struct {
	Proposer   common.Address
	DocumentId *big.Int
	Nonce      *big.Int
	RequestId  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubmittedAmendment is a free log retrieval operation binding the contract event 0x5613f7bdb7de0b9e304bd3206cfe56e5fbe6132fb78f885650b90e5ae6608810.
//
// Solidity: event SubmittedAmendment(address indexed proposer, uint256 indexed documentId, uint256 indexed nonce, uint256 requestId)
func (_AmendRequest *AmendRequestFilterer) FilterSubmittedAmendment(opts *bind.FilterOpts, proposer []common.Address, documentId []*big.Int, nonce []*big.Int) (*AmendRequestSubmittedAmendmentIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var documentIdRule []interface{}
	for _, documentIdItem := range documentId {
		documentIdRule = append(documentIdRule, documentIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _AmendRequest.contract.FilterLogs(opts, "SubmittedAmendment", proposerRule, documentIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &AmendRequestSubmittedAmendmentIterator{contract: _AmendRequest.contract, event: "SubmittedAmendment", logs: logs, sub: sub}, nil
}

var SubmittedAmendmentTopicHash = "0x5613f7bdb7de0b9e304bd3206cfe56e5fbe6132fb78f885650b90e5ae6608810"

// WatchSubmittedAmendment is a free log subscription operation binding the contract event 0x5613f7bdb7de0b9e304bd3206cfe56e5fbe6132fb78f885650b90e5ae6608810.
//
// Solidity: event SubmittedAmendment(address indexed proposer, uint256 indexed documentId, uint256 indexed nonce, uint256 requestId)
func (_AmendRequest *AmendRequestFilterer) WatchSubmittedAmendment(opts *bind.WatchOpts, sink chan<- *AmendRequestSubmittedAmendment, proposer []common.Address, documentId []*big.Int, nonce []*big.Int) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var documentIdRule []interface{}
	for _, documentIdItem := range documentId {
		documentIdRule = append(documentIdRule, documentIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _AmendRequest.contract.WatchLogs(opts, "SubmittedAmendment", proposerRule, documentIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmendRequestSubmittedAmendment)
				if err := _AmendRequest.contract.UnpackLog(event, "SubmittedAmendment", log); err != nil {
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

// ParseSubmittedAmendment is a log parse operation binding the contract event 0x5613f7bdb7de0b9e304bd3206cfe56e5fbe6132fb78f885650b90e5ae6608810.
//
// Solidity: event SubmittedAmendment(address indexed proposer, uint256 indexed documentId, uint256 indexed nonce, uint256 requestId)
func (_AmendRequest *AmendRequestFilterer) ParseSubmittedAmendment(log types.Log) (*AmendRequestSubmittedAmendment, error) {
	event := new(AmendRequestSubmittedAmendment)
	if err := _AmendRequest.contract.UnpackLog(event, "SubmittedAmendment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
