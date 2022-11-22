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
	IssuingBank        [32]byte
	AdvisingBank       [32]byte
	ReimbursingBank    [32]byte
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

// RouterServiceABI is the input ABI used to generate the binding from.
const RouterServiceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_management\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_subStage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signedTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numOfDocuments\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"contentHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"acknowledge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIStageContract.Content\",\"name\":\"_content\",\"type\":\"tuple\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"approveAmendment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"}],\"name\":\"closeLC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"fulfillAmendment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_typeOf\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getAmendmentRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"typeOf\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"migratingStages\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subStage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signedTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numOfDocuments\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"contentHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"acknowledge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIStageContract.Content\",\"name\":\"content\",\"type\":\"tuple\"}],\"internalType\":\"structIAmendRequest.AmendStage\",\"name\":\"amendStage\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"issuingBank\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"advisingBank\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"reimbursingBank\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"issuingBankSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"advisingBankSig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reimbursingBankSig\",\"type\":\"bytes\"}],\"internalType\":\"structIAmendRequest.Confirmation\",\"name\":\"confirmed\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isFulfilled\",\"type\":\"bool\"}],\"internalType\":\"structIAmendRequest.Request\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"}],\"name\":\"getRootHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_subStage\",\"type\":\"uint256\"}],\"name\":\"getStageContent\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signedTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numOfDocuments\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"contentHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"acknowledge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIStageContract.Content\",\"name\":\"_content\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"isAmendApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"management\",\"outputs\":[{\"internalType\":\"contractILCManagement\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_management\",\"type\":\"address\"}],\"name\":\"setLCManagement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_migratingStages\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subStage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signedTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numOfDocuments\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"contentHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"acknowledge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIStageContract.Content\",\"name\":\"content\",\"type\":\"tuple\"}],\"internalType\":\"structIAmendRequest.AmendStage\",\"name\":\"_amendStage\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"submitAmendment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var RouterServiceParsedABI, _ = abi.JSON(strings.NewReader(RouterServiceABI))

// RouterService is an auto generated Go binding around an Ethereum contract.
type RouterService struct {
	RouterServiceCaller     // Read-only binding to the contract
	RouterServiceTransactor // Write-only binding to the contract
	RouterServiceFilterer   // Log filterer for contract events
}

// RouterServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterServiceSession struct {
	Contract     *RouterService    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterServiceCallerSession struct {
	Contract *RouterServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RouterServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterServiceTransactorSession struct {
	Contract     *RouterServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RouterServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterServiceRaw struct {
	Contract *RouterService // Generic contract binding to access the raw methods on
}

// RouterServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterServiceCallerRaw struct {
	Contract *RouterServiceCaller // Generic read-only contract binding to access the raw methods on
}

// RouterServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterServiceTransactorRaw struct {
	Contract *RouterServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouterService creates a new instance of RouterService, bound to a specific deployed contract.
func NewRouterService(address common.Address, backend bind.ContractBackend) (*RouterService, error) {
	contract, err := bindRouterService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RouterService{RouterServiceCaller: RouterServiceCaller{contract: contract}, RouterServiceTransactor: RouterServiceTransactor{contract: contract}, RouterServiceFilterer: RouterServiceFilterer{contract: contract}}, nil
}

// NewRouterServiceCaller creates a new read-only instance of RouterService, bound to a specific deployed contract.
func NewRouterServiceCaller(address common.Address, caller bind.ContractCaller) (*RouterServiceCaller, error) {
	contract, err := bindRouterService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterServiceCaller{contract: contract}, nil
}

// NewRouterServiceTransactor creates a new write-only instance of RouterService, bound to a specific deployed contract.
func NewRouterServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterServiceTransactor, error) {
	contract, err := bindRouterService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterServiceTransactor{contract: contract}, nil
}

// NewRouterServiceFilterer creates a new log filterer instance of RouterService, bound to a specific deployed contract.
func NewRouterServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterServiceFilterer, error) {
	contract, err := bindRouterService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterServiceFilterer{contract: contract}, nil
}

// bindRouterService binds a generic wrapper to an already deployed contract.
func bindRouterService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RouterService *RouterServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RouterService.Contract.RouterServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RouterService *RouterServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterService.Contract.RouterServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RouterService *RouterServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RouterService.Contract.RouterServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RouterService *RouterServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RouterService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RouterService *RouterServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RouterService *RouterServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RouterService.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0xb93f9b0a.
//
// Solidity: function getAddress(uint256 _documentId) view returns(address _contract, uint256 _typeOf)
func (_RouterService *RouterServiceCaller) GetAddress(opts *bind.CallOpts, _documentId *big.Int) (struct {
	Contract common.Address
	TypeOf   *big.Int
}, error) {
	var out []interface{}
	err := _RouterService.contract.Call(opts, &out, "getAddress", _documentId)

	outstruct := new(struct {
		Contract common.Address
		TypeOf   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Contract = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TypeOf = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAddress is a free data retrieval call binding the contract method 0xb93f9b0a.
//
// Solidity: function getAddress(uint256 _documentId) view returns(address _contract, uint256 _typeOf)
func (_RouterService *RouterServiceSession) GetAddress(_documentId *big.Int) (struct {
	Contract common.Address
	TypeOf   *big.Int
}, error) {
	return _RouterService.Contract.GetAddress(&_RouterService.CallOpts, _documentId)
}

// GetAddress is a free data retrieval call binding the contract method 0xb93f9b0a.
//
// Solidity: function getAddress(uint256 _documentId) view returns(address _contract, uint256 _typeOf)
func (_RouterService *RouterServiceCallerSession) GetAddress(_documentId *big.Int) (struct {
	Contract common.Address
	TypeOf   *big.Int
}, error) {
	return _RouterService.Contract.GetAddress(&_RouterService.CallOpts, _documentId)
}

// GetAmendmentRequest is a free data retrieval call binding the contract method 0x75b21ea8.
//
// Solidity: function getAmendmentRequest(uint256 _documentId, uint256 _requestId) view returns((uint256,address,bytes32[],(uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)),(bytes32,bytes32,bytes32,bytes,bytes,bytes),bool))
func (_RouterService *RouterServiceCaller) GetAmendmentRequest(opts *bind.CallOpts, _documentId *big.Int, _requestId *big.Int) (IAmendRequestRequest, error) {
	var out []interface{}
	err := _RouterService.contract.Call(opts, &out, "getAmendmentRequest", _documentId, _requestId)

	if err != nil {
		return *new(IAmendRequestRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IAmendRequestRequest)).(*IAmendRequestRequest)

	return out0, err

}

// GetAmendmentRequest is a free data retrieval call binding the contract method 0x75b21ea8.
//
// Solidity: function getAmendmentRequest(uint256 _documentId, uint256 _requestId) view returns((uint256,address,bytes32[],(uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)),(bytes32,bytes32,bytes32,bytes,bytes,bytes),bool))
func (_RouterService *RouterServiceSession) GetAmendmentRequest(_documentId *big.Int, _requestId *big.Int) (IAmendRequestRequest, error) {
	return _RouterService.Contract.GetAmendmentRequest(&_RouterService.CallOpts, _documentId, _requestId)
}

// GetAmendmentRequest is a free data retrieval call binding the contract method 0x75b21ea8.
//
// Solidity: function getAmendmentRequest(uint256 _documentId, uint256 _requestId) view returns((uint256,address,bytes32[],(uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)),(bytes32,bytes32,bytes32,bytes,bytes,bytes),bool))
func (_RouterService *RouterServiceCallerSession) GetAmendmentRequest(_documentId *big.Int, _requestId *big.Int) (IAmendRequestRequest, error) {
	return _RouterService.Contract.GetAmendmentRequest(&_RouterService.CallOpts, _documentId, _requestId)
}

// GetRootHash is a free data retrieval call binding the contract method 0x093abc86.
//
// Solidity: function getRootHash(uint256 _documentId) view returns(bytes32)
func (_RouterService *RouterServiceCaller) GetRootHash(opts *bind.CallOpts, _documentId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RouterService.contract.Call(opts, &out, "getRootHash", _documentId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRootHash is a free data retrieval call binding the contract method 0x093abc86.
//
// Solidity: function getRootHash(uint256 _documentId) view returns(bytes32)
func (_RouterService *RouterServiceSession) GetRootHash(_documentId *big.Int) ([32]byte, error) {
	return _RouterService.Contract.GetRootHash(&_RouterService.CallOpts, _documentId)
}

// GetRootHash is a free data retrieval call binding the contract method 0x093abc86.
//
// Solidity: function getRootHash(uint256 _documentId) view returns(bytes32)
func (_RouterService *RouterServiceCallerSession) GetRootHash(_documentId *big.Int) ([32]byte, error) {
	return _RouterService.Contract.GetRootHash(&_RouterService.CallOpts, _documentId)
}

// GetStageContent is a free data retrieval call binding the contract method 0x850caeb0.
//
// Solidity: function getStageContent(uint256 _documentId, uint256 _stage, uint256 _subStage) view returns((bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content)
func (_RouterService *RouterServiceCaller) GetStageContent(opts *bind.CallOpts, _documentId *big.Int, _stage *big.Int, _subStage *big.Int) (IStageContractContent, error) {
	var out []interface{}
	err := _RouterService.contract.Call(opts, &out, "getStageContent", _documentId, _stage, _subStage)

	if err != nil {
		return *new(IStageContractContent), err
	}

	out0 := *abi.ConvertType(out[0], new(IStageContractContent)).(*IStageContractContent)

	return out0, err

}

// GetStageContent is a free data retrieval call binding the contract method 0x850caeb0.
//
// Solidity: function getStageContent(uint256 _documentId, uint256 _stage, uint256 _subStage) view returns((bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content)
func (_RouterService *RouterServiceSession) GetStageContent(_documentId *big.Int, _stage *big.Int, _subStage *big.Int) (IStageContractContent, error) {
	return _RouterService.Contract.GetStageContent(&_RouterService.CallOpts, _documentId, _stage, _subStage)
}

// GetStageContent is a free data retrieval call binding the contract method 0x850caeb0.
//
// Solidity: function getStageContent(uint256 _documentId, uint256 _stage, uint256 _subStage) view returns((bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content)
func (_RouterService *RouterServiceCallerSession) GetStageContent(_documentId *big.Int, _stage *big.Int, _subStage *big.Int) (IStageContractContent, error) {
	return _RouterService.Contract.GetStageContent(&_RouterService.CallOpts, _documentId, _stage, _subStage)
}

// IsAmendApproved is a free data retrieval call binding the contract method 0x8f81311a.
//
// Solidity: function isAmendApproved(uint256 _documentId, uint256 _requestId) view returns(bool)
func (_RouterService *RouterServiceCaller) IsAmendApproved(opts *bind.CallOpts, _documentId *big.Int, _requestId *big.Int) (bool, error) {
	var out []interface{}
	err := _RouterService.contract.Call(opts, &out, "isAmendApproved", _documentId, _requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAmendApproved is a free data retrieval call binding the contract method 0x8f81311a.
//
// Solidity: function isAmendApproved(uint256 _documentId, uint256 _requestId) view returns(bool)
func (_RouterService *RouterServiceSession) IsAmendApproved(_documentId *big.Int, _requestId *big.Int) (bool, error) {
	return _RouterService.Contract.IsAmendApproved(&_RouterService.CallOpts, _documentId, _requestId)
}

// IsAmendApproved is a free data retrieval call binding the contract method 0x8f81311a.
//
// Solidity: function isAmendApproved(uint256 _documentId, uint256 _requestId) view returns(bool)
func (_RouterService *RouterServiceCallerSession) IsAmendApproved(_documentId *big.Int, _requestId *big.Int) (bool, error) {
	return _RouterService.Contract.IsAmendApproved(&_RouterService.CallOpts, _documentId, _requestId)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_RouterService *RouterServiceCaller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterService.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_RouterService *RouterServiceSession) Management() (common.Address, error) {
	return _RouterService.Contract.Management(&_RouterService.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_RouterService *RouterServiceCallerSession) Management() (common.Address, error) {
	return _RouterService.Contract.Management(&_RouterService.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x142cf46a.
//
// Solidity: function approve(uint256 _documentId, uint256 _stage, uint256 _subStage, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns()
func (_RouterService *RouterServiceTransactor) Approve(opts *bind.TransactOpts, _documentId *big.Int, _stage *big.Int, _subStage *big.Int, _content IStageContractContent) (*types.Transaction, error) {
	return _RouterService.contract.Transact(opts, "approve", _documentId, _stage, _subStage, _content)
}

// Approve is a paid mutator transaction binding the contract method 0x142cf46a.
//
// Solidity: function approve(uint256 _documentId, uint256 _stage, uint256 _subStage, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns()
func (_RouterService *RouterServiceSession) Approve(_documentId *big.Int, _stage *big.Int, _subStage *big.Int, _content IStageContractContent) (*types.Transaction, error) {
	return _RouterService.Contract.Approve(&_RouterService.TransactOpts, _documentId, _stage, _subStage, _content)
}

// Approve is a paid mutator transaction binding the contract method 0x142cf46a.
//
// Solidity: function approve(uint256 _documentId, uint256 _stage, uint256 _subStage, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns()
func (_RouterService *RouterServiceTransactorSession) Approve(_documentId *big.Int, _stage *big.Int, _subStage *big.Int, _content IStageContractContent) (*types.Transaction, error) {
	return _RouterService.Contract.Approve(&_RouterService.TransactOpts, _documentId, _stage, _subStage, _content)
}

// ApproveAmendment is a paid mutator transaction binding the contract method 0xd06c26ca.
//
// Solidity: function approveAmendment(uint256 _documentId, uint256 _requestId, bytes _signature) returns()
func (_RouterService *RouterServiceTransactor) ApproveAmendment(opts *bind.TransactOpts, _documentId *big.Int, _requestId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _RouterService.contract.Transact(opts, "approveAmendment", _documentId, _requestId, _signature)
}

// ApproveAmendment is a paid mutator transaction binding the contract method 0xd06c26ca.
//
// Solidity: function approveAmendment(uint256 _documentId, uint256 _requestId, bytes _signature) returns()
func (_RouterService *RouterServiceSession) ApproveAmendment(_documentId *big.Int, _requestId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _RouterService.Contract.ApproveAmendment(&_RouterService.TransactOpts, _documentId, _requestId, _signature)
}

// ApproveAmendment is a paid mutator transaction binding the contract method 0xd06c26ca.
//
// Solidity: function approveAmendment(uint256 _documentId, uint256 _requestId, bytes _signature) returns()
func (_RouterService *RouterServiceTransactorSession) ApproveAmendment(_documentId *big.Int, _requestId *big.Int, _signature []byte) (*types.Transaction, error) {
	return _RouterService.Contract.ApproveAmendment(&_RouterService.TransactOpts, _documentId, _requestId, _signature)
}

// CloseLC is a paid mutator transaction binding the contract method 0x4d662357.
//
// Solidity: function closeLC(uint256 _documentId) returns()
func (_RouterService *RouterServiceTransactor) CloseLC(opts *bind.TransactOpts, _documentId *big.Int) (*types.Transaction, error) {
	return _RouterService.contract.Transact(opts, "closeLC", _documentId)
}

// CloseLC is a paid mutator transaction binding the contract method 0x4d662357.
//
// Solidity: function closeLC(uint256 _documentId) returns()
func (_RouterService *RouterServiceSession) CloseLC(_documentId *big.Int) (*types.Transaction, error) {
	return _RouterService.Contract.CloseLC(&_RouterService.TransactOpts, _documentId)
}

// CloseLC is a paid mutator transaction binding the contract method 0x4d662357.
//
// Solidity: function closeLC(uint256 _documentId) returns()
func (_RouterService *RouterServiceTransactorSession) CloseLC(_documentId *big.Int) (*types.Transaction, error) {
	return _RouterService.Contract.CloseLC(&_RouterService.TransactOpts, _documentId)
}

// FulfillAmendment is a paid mutator transaction binding the contract method 0x16d4eb3f.
//
// Solidity: function fulfillAmendment(uint256 _documentId, uint256 _requestId) returns()
func (_RouterService *RouterServiceTransactor) FulfillAmendment(opts *bind.TransactOpts, _documentId *big.Int, _requestId *big.Int) (*types.Transaction, error) {
	return _RouterService.contract.Transact(opts, "fulfillAmendment", _documentId, _requestId)
}

// FulfillAmendment is a paid mutator transaction binding the contract method 0x16d4eb3f.
//
// Solidity: function fulfillAmendment(uint256 _documentId, uint256 _requestId) returns()
func (_RouterService *RouterServiceSession) FulfillAmendment(_documentId *big.Int, _requestId *big.Int) (*types.Transaction, error) {
	return _RouterService.Contract.FulfillAmendment(&_RouterService.TransactOpts, _documentId, _requestId)
}

// FulfillAmendment is a paid mutator transaction binding the contract method 0x16d4eb3f.
//
// Solidity: function fulfillAmendment(uint256 _documentId, uint256 _requestId) returns()
func (_RouterService *RouterServiceTransactorSession) FulfillAmendment(_documentId *big.Int, _requestId *big.Int) (*types.Transaction, error) {
	return _RouterService.Contract.FulfillAmendment(&_RouterService.TransactOpts, _documentId, _requestId)
}

// SetLCManagement is a paid mutator transaction binding the contract method 0xb3463971.
//
// Solidity: function setLCManagement(address _management) returns()
func (_RouterService *RouterServiceTransactor) SetLCManagement(opts *bind.TransactOpts, _management common.Address) (*types.Transaction, error) {
	return _RouterService.contract.Transact(opts, "setLCManagement", _management)
}

// SetLCManagement is a paid mutator transaction binding the contract method 0xb3463971.
//
// Solidity: function setLCManagement(address _management) returns()
func (_RouterService *RouterServiceSession) SetLCManagement(_management common.Address) (*types.Transaction, error) {
	return _RouterService.Contract.SetLCManagement(&_RouterService.TransactOpts, _management)
}

// SetLCManagement is a paid mutator transaction binding the contract method 0xb3463971.
//
// Solidity: function setLCManagement(address _management) returns()
func (_RouterService *RouterServiceTransactorSession) SetLCManagement(_management common.Address) (*types.Transaction, error) {
	return _RouterService.Contract.SetLCManagement(&_RouterService.TransactOpts, _management)
}

// SubmitAmendment is a paid mutator transaction binding the contract method 0xbbf62c7e.
//
// Solidity: function submitAmendment(uint256 _documentId, bytes32[] _migratingStages, (uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)) _amendStage, bytes _signature) returns()
func (_RouterService *RouterServiceTransactor) SubmitAmendment(opts *bind.TransactOpts, _documentId *big.Int, _migratingStages [][32]byte, _amendStage IAmendRequestAmendStage, _signature []byte) (*types.Transaction, error) {
	return _RouterService.contract.Transact(opts, "submitAmendment", _documentId, _migratingStages, _amendStage, _signature)
}

// SubmitAmendment is a paid mutator transaction binding the contract method 0xbbf62c7e.
//
// Solidity: function submitAmendment(uint256 _documentId, bytes32[] _migratingStages, (uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)) _amendStage, bytes _signature) returns()
func (_RouterService *RouterServiceSession) SubmitAmendment(_documentId *big.Int, _migratingStages [][32]byte, _amendStage IAmendRequestAmendStage, _signature []byte) (*types.Transaction, error) {
	return _RouterService.Contract.SubmitAmendment(&_RouterService.TransactOpts, _documentId, _migratingStages, _amendStage, _signature)
}

// SubmitAmendment is a paid mutator transaction binding the contract method 0xbbf62c7e.
//
// Solidity: function submitAmendment(uint256 _documentId, bytes32[] _migratingStages, (uint256,uint256,(bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes)) _amendStage, bytes _signature) returns()
func (_RouterService *RouterServiceTransactorSession) SubmitAmendment(_documentId *big.Int, _migratingStages [][32]byte, _amendStage IAmendRequestAmendStage, _signature []byte) (*types.Transaction, error) {
	return _RouterService.Contract.SubmitAmendment(&_RouterService.TransactOpts, _documentId, _migratingStages, _amendStage, _signature)
}
