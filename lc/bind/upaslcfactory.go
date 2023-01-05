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

// UPASLCFactoryABI is the input ABI used to generate the binding from.
const UPASLCFactoryABI = "[{\"inputs\":[{\"internalType\":\"contractILCManagement\",\"name\":\"_management\",\"type\":\"address\"},{\"internalType\":\"contractIWrapper\",\"name\":\"_wrapper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"documentID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lcContractAddr\",\"type\":\"address\"}],\"name\":\"NewUPASLC\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPAS_WRAPPER\",\"outputs\":[{\"internalType\":\"contractIWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_parties\",\"type\":\"string[]\"}],\"name\":\"amend\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_parties\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signedTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numOfDocuments\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"contentHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"acknowledge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIStageContract.Content\",\"name\":\"_content\",\"type\":\"tuple\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"}],\"name\":\"getLCAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"management\",\"outputs\":[{\"internalType\":\"contractILCManagement\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_management\",\"type\":\"address\"}],\"name\":\"setLCManagement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var UPASLCFactoryParsedABI, _ = abi.JSON(strings.NewReader(UPASLCFactoryABI))

// UPASLCFactory is an auto generated Go binding around an Ethereum contract.
type UPASLCFactory struct {
	UPASLCFactoryCaller     // Read-only binding to the contract
	UPASLCFactoryTransactor // Write-only binding to the contract
	UPASLCFactoryFilterer   // Log filterer for contract events
}

// UPASLCFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UPASLCFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UPASLCFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UPASLCFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UPASLCFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UPASLCFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UPASLCFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UPASLCFactorySession struct {
	Contract     *UPASLCFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UPASLCFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UPASLCFactoryCallerSession struct {
	Contract *UPASLCFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UPASLCFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UPASLCFactoryTransactorSession struct {
	Contract     *UPASLCFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UPASLCFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UPASLCFactoryRaw struct {
	Contract *UPASLCFactory // Generic contract binding to access the raw methods on
}

// UPASLCFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UPASLCFactoryCallerRaw struct {
	Contract *UPASLCFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UPASLCFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UPASLCFactoryTransactorRaw struct {
	Contract *UPASLCFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUPASLCFactory creates a new instance of UPASLCFactory, bound to a specific deployed contract.
func NewUPASLCFactory(address common.Address, backend bind.ContractBackend) (*UPASLCFactory, error) {
	contract, err := bindUPASLCFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UPASLCFactory{UPASLCFactoryCaller: UPASLCFactoryCaller{contract: contract}, UPASLCFactoryTransactor: UPASLCFactoryTransactor{contract: contract}, UPASLCFactoryFilterer: UPASLCFactoryFilterer{contract: contract}}, nil
}

// NewUPASLCFactoryCaller creates a new read-only instance of UPASLCFactory, bound to a specific deployed contract.
func NewUPASLCFactoryCaller(address common.Address, caller bind.ContractCaller) (*UPASLCFactoryCaller, error) {
	contract, err := bindUPASLCFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UPASLCFactoryCaller{contract: contract}, nil
}

// NewUPASLCFactoryTransactor creates a new write-only instance of UPASLCFactory, bound to a specific deployed contract.
func NewUPASLCFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UPASLCFactoryTransactor, error) {
	contract, err := bindUPASLCFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UPASLCFactoryTransactor{contract: contract}, nil
}

// NewUPASLCFactoryFilterer creates a new log filterer instance of UPASLCFactory, bound to a specific deployed contract.
func NewUPASLCFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UPASLCFactoryFilterer, error) {
	contract, err := bindUPASLCFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UPASLCFactoryFilterer{contract: contract}, nil
}

// bindUPASLCFactory binds a generic wrapper to an already deployed contract.
func bindUPASLCFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UPASLCFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UPASLCFactory *UPASLCFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UPASLCFactory.Contract.UPASLCFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UPASLCFactory *UPASLCFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UPASLCFactory.Contract.UPASLCFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UPASLCFactory *UPASLCFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UPASLCFactory.Contract.UPASLCFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UPASLCFactory *UPASLCFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UPASLCFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UPASLCFactory *UPASLCFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UPASLCFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UPASLCFactory *UPASLCFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UPASLCFactory.Contract.contract.Transact(opts, method, params...)
}

// UPASWRAPPER is a free data retrieval call binding the contract method 0x1134edd3.
//
// Solidity: function UPAS_WRAPPER() view returns(address)
func (_UPASLCFactory *UPASLCFactoryCaller) UPASWRAPPER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UPASLCFactory.contract.Call(opts, &out, "UPAS_WRAPPER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UPASWRAPPER is a free data retrieval call binding the contract method 0x1134edd3.
//
// Solidity: function UPAS_WRAPPER() view returns(address)
func (_UPASLCFactory *UPASLCFactorySession) UPASWRAPPER() (common.Address, error) {
	return _UPASLCFactory.Contract.UPASWRAPPER(&_UPASLCFactory.CallOpts)
}

// UPASWRAPPER is a free data retrieval call binding the contract method 0x1134edd3.
//
// Solidity: function UPAS_WRAPPER() view returns(address)
func (_UPASLCFactory *UPASLCFactoryCallerSession) UPASWRAPPER() (common.Address, error) {
	return _UPASLCFactory.Contract.UPASWRAPPER(&_UPASLCFactory.CallOpts)
}

// GetLCAddress is a free data retrieval call binding the contract method 0x793e97c6.
//
// Solidity: function getLCAddress(uint256 _documentId) view returns(address[])
func (_UPASLCFactory *UPASLCFactoryCaller) GetLCAddress(opts *bind.CallOpts, _documentId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _UPASLCFactory.contract.Call(opts, &out, "getLCAddress", _documentId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLCAddress is a free data retrieval call binding the contract method 0x793e97c6.
//
// Solidity: function getLCAddress(uint256 _documentId) view returns(address[])
func (_UPASLCFactory *UPASLCFactorySession) GetLCAddress(_documentId *big.Int) ([]common.Address, error) {
	return _UPASLCFactory.Contract.GetLCAddress(&_UPASLCFactory.CallOpts, _documentId)
}

// GetLCAddress is a free data retrieval call binding the contract method 0x793e97c6.
//
// Solidity: function getLCAddress(uint256 _documentId) view returns(address[])
func (_UPASLCFactory *UPASLCFactoryCallerSession) GetLCAddress(_documentId *big.Int) ([]common.Address, error) {
	return _UPASLCFactory.Contract.GetLCAddress(&_UPASLCFactory.CallOpts, _documentId)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_UPASLCFactory *UPASLCFactoryCaller) Management(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UPASLCFactory.contract.Call(opts, &out, "management")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_UPASLCFactory *UPASLCFactorySession) Management() (common.Address, error) {
	return _UPASLCFactory.Contract.Management(&_UPASLCFactory.CallOpts)
}

// Management is a free data retrieval call binding the contract method 0x88a8d602.
//
// Solidity: function management() view returns(address)
func (_UPASLCFactory *UPASLCFactoryCallerSession) Management() (common.Address, error) {
	return _UPASLCFactory.Contract.Management(&_UPASLCFactory.CallOpts)
}

// Amend is a paid mutator transaction binding the contract method 0x6a654994.
//
// Solidity: function amend(address _executor, uint256 _documentId, string[] _parties) returns(address _contract)
func (_UPASLCFactory *UPASLCFactoryTransactor) Amend(opts *bind.TransactOpts, _executor common.Address, _documentId *big.Int, _parties []string) (*types.Transaction, error) {
	return _UPASLCFactory.contract.Transact(opts, "amend", _executor, _documentId, _parties)
}

// Amend is a paid mutator transaction binding the contract method 0x6a654994.
//
// Solidity: function amend(address _executor, uint256 _documentId, string[] _parties) returns(address _contract)
func (_UPASLCFactory *UPASLCFactorySession) Amend(_executor common.Address, _documentId *big.Int, _parties []string) (*types.Transaction, error) {
	return _UPASLCFactory.Contract.Amend(&_UPASLCFactory.TransactOpts, _executor, _documentId, _parties)
}

// Amend is a paid mutator transaction binding the contract method 0x6a654994.
//
// Solidity: function amend(address _executor, uint256 _documentId, string[] _parties) returns(address _contract)
func (_UPASLCFactory *UPASLCFactoryTransactorSession) Amend(_executor common.Address, _documentId *big.Int, _parties []string) (*types.Transaction, error) {
	return _UPASLCFactory.Contract.Amend(&_UPASLCFactory.TransactOpts, _executor, _documentId, _parties)
}

// Create is a paid mutator transaction binding the contract method 0xebdcefdd.
//
// Solidity: function create(string[] _parties, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns(address _contract)
func (_UPASLCFactory *UPASLCFactoryTransactor) Create(opts *bind.TransactOpts, _parties []string, _content IStageContractContent) (*types.Transaction, error) {
	return _UPASLCFactory.contract.Transact(opts, "create", _parties, _content)
}

// Create is a paid mutator transaction binding the contract method 0xebdcefdd.
//
// Solidity: function create(string[] _parties, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns(address _contract)
func (_UPASLCFactory *UPASLCFactorySession) Create(_parties []string, _content IStageContractContent) (*types.Transaction, error) {
	return _UPASLCFactory.Contract.Create(&_UPASLCFactory.TransactOpts, _parties, _content)
}

// Create is a paid mutator transaction binding the contract method 0xebdcefdd.
//
// Solidity: function create(string[] _parties, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns(address _contract)
func (_UPASLCFactory *UPASLCFactoryTransactorSession) Create(_parties []string, _content IStageContractContent) (*types.Transaction, error) {
	return _UPASLCFactory.Contract.Create(&_UPASLCFactory.TransactOpts, _parties, _content)
}

// SetLCManagement is a paid mutator transaction binding the contract method 0xb3463971.
//
// Solidity: function setLCManagement(address _management) returns()
func (_UPASLCFactory *UPASLCFactoryTransactor) SetLCManagement(opts *bind.TransactOpts, _management common.Address) (*types.Transaction, error) {
	return _UPASLCFactory.contract.Transact(opts, "setLCManagement", _management)
}

// SetLCManagement is a paid mutator transaction binding the contract method 0xb3463971.
//
// Solidity: function setLCManagement(address _management) returns()
func (_UPASLCFactory *UPASLCFactorySession) SetLCManagement(_management common.Address) (*types.Transaction, error) {
	return _UPASLCFactory.Contract.SetLCManagement(&_UPASLCFactory.TransactOpts, _management)
}

// SetLCManagement is a paid mutator transaction binding the contract method 0xb3463971.
//
// Solidity: function setLCManagement(address _management) returns()
func (_UPASLCFactory *UPASLCFactoryTransactorSession) SetLCManagement(_management common.Address) (*types.Transaction, error) {
	return _UPASLCFactory.Contract.SetLCManagement(&_UPASLCFactory.TransactOpts, _management)
}

// UPASLCFactoryNewUPASLCIterator is returned from FilterNewUPASLC and is used to iterate over the raw logs and unpacked data for NewUPASLC events raised by the UPASLCFactory contract.
type UPASLCFactoryNewUPASLCIterator struct {
	Event *UPASLCFactoryNewUPASLC // Event containing the contract specifics and raw log

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
func (it *UPASLCFactoryNewUPASLCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UPASLCFactoryNewUPASLC)
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
		it.Event = new(UPASLCFactoryNewUPASLC)
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
func (it *UPASLCFactoryNewUPASLCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UPASLCFactoryNewUPASLCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UPASLCFactoryNewUPASLC represents a NewUPASLC event raised by the UPASLCFactory contract.
type UPASLCFactoryNewUPASLC struct {
	DocumentID     *big.Int
	Creator        common.Address
	LcContractAddr common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewUPASLC is a free log retrieval operation binding the contract event 0x7730961fea105a63d4883ad6da7b350a89f5ced379b199c51a2ca698a49ea392.
//
// Solidity: event NewUPASLC(uint256 indexed documentID, address indexed creator, address indexed lcContractAddr)
func (_UPASLCFactory *UPASLCFactoryFilterer) FilterNewUPASLC(opts *bind.FilterOpts, documentID []*big.Int, creator []common.Address, lcContractAddr []common.Address) (*UPASLCFactoryNewUPASLCIterator, error) {

	var documentIDRule []interface{}
	for _, documentIDItem := range documentID {
		documentIDRule = append(documentIDRule, documentIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var lcContractAddrRule []interface{}
	for _, lcContractAddrItem := range lcContractAddr {
		lcContractAddrRule = append(lcContractAddrRule, lcContractAddrItem)
	}

	logs, sub, err := _UPASLCFactory.contract.FilterLogs(opts, "NewUPASLC", documentIDRule, creatorRule, lcContractAddrRule)
	if err != nil {
		return nil, err
	}
	return &UPASLCFactoryNewUPASLCIterator{contract: _UPASLCFactory.contract, event: "NewUPASLC", logs: logs, sub: sub}, nil
}

var NewUPASLCTopicHash = "0x7730961fea105a63d4883ad6da7b350a89f5ced379b199c51a2ca698a49ea392"

// WatchNewUPASLC is a free log subscription operation binding the contract event 0x7730961fea105a63d4883ad6da7b350a89f5ced379b199c51a2ca698a49ea392.
//
// Solidity: event NewUPASLC(uint256 indexed documentID, address indexed creator, address indexed lcContractAddr)
func (_UPASLCFactory *UPASLCFactoryFilterer) WatchNewUPASLC(opts *bind.WatchOpts, sink chan<- *UPASLCFactoryNewUPASLC, documentID []*big.Int, creator []common.Address, lcContractAddr []common.Address) (event.Subscription, error) {

	var documentIDRule []interface{}
	for _, documentIDItem := range documentID {
		documentIDRule = append(documentIDRule, documentIDItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var lcContractAddrRule []interface{}
	for _, lcContractAddrItem := range lcContractAddr {
		lcContractAddrRule = append(lcContractAddrRule, lcContractAddrItem)
	}

	logs, sub, err := _UPASLCFactory.contract.WatchLogs(opts, "NewUPASLC", documentIDRule, creatorRule, lcContractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UPASLCFactoryNewUPASLC)
				if err := _UPASLCFactory.contract.UnpackLog(event, "NewUPASLC", log); err != nil {
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

// ParseNewUPASLC is a log parse operation binding the contract event 0x7730961fea105a63d4883ad6da7b350a89f5ced379b199c51a2ca698a49ea392.
//
// Solidity: event NewUPASLC(uint256 indexed documentID, address indexed creator, address indexed lcContractAddr)
func (_UPASLCFactory *UPASLCFactoryFilterer) ParseNewUPASLC(log types.Log) (*UPASLCFactoryNewUPASLC, error) {
	event := new(UPASLCFactoryNewUPASLC)
	if err := _UPASLCFactory.contract.UnpackLog(event, "NewUPASLC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
