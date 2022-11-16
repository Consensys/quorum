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

// StandardLCFactoryABI is the input ABI used to generate the binding from.
const StandardLCFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_amc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"documentID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lcContractAddr\",\"type\":\"address\"}],\"name\":\"NewStandardLC\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"amc\",\"outputs\":[{\"internalType\":\"contractIAMC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_parties\",\"type\":\"bytes32[]\"}],\"name\":\"amend\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_parties\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signedTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numOfDocuments\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"contentHash\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"acknowledge\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIStageContract.Content\",\"name\":\"_content\",\"type\":\"tuple\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_documentId\",\"type\":\"uint256\"}],\"name\":\"getLCAddress\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_amc\",\"type\":\"address\"}],\"name\":\"setAMC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var StandardLCFactoryParsedABI, _ = abi.JSON(strings.NewReader(StandardLCFactoryABI))

// StandardLCFactory is an auto generated Go binding around an Ethereum contract.
type StandardLCFactory struct {
	StandardLCFactoryCaller     // Read-only binding to the contract
	StandardLCFactoryTransactor // Write-only binding to the contract
	StandardLCFactoryFilterer   // Log filterer for contract events
}

// StandardLCFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardLCFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardLCFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardLCFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardLCFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardLCFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardLCFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardLCFactorySession struct {
	Contract     *StandardLCFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StandardLCFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardLCFactoryCallerSession struct {
	Contract *StandardLCFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StandardLCFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardLCFactoryTransactorSession struct {
	Contract     *StandardLCFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StandardLCFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardLCFactoryRaw struct {
	Contract *StandardLCFactory // Generic contract binding to access the raw methods on
}

// StandardLCFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardLCFactoryCallerRaw struct {
	Contract *StandardLCFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// StandardLCFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardLCFactoryTransactorRaw struct {
	Contract *StandardLCFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardLCFactory creates a new instance of StandardLCFactory, bound to a specific deployed contract.
func NewStandardLCFactory(address common.Address, backend bind.ContractBackend) (*StandardLCFactory, error) {
	contract, err := bindStandardLCFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardLCFactory{StandardLCFactoryCaller: StandardLCFactoryCaller{contract: contract}, StandardLCFactoryTransactor: StandardLCFactoryTransactor{contract: contract}, StandardLCFactoryFilterer: StandardLCFactoryFilterer{contract: contract}}, nil
}

// NewStandardLCFactoryCaller creates a new read-only instance of StandardLCFactory, bound to a specific deployed contract.
func NewStandardLCFactoryCaller(address common.Address, caller bind.ContractCaller) (*StandardLCFactoryCaller, error) {
	contract, err := bindStandardLCFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardLCFactoryCaller{contract: contract}, nil
}

// NewStandardLCFactoryTransactor creates a new write-only instance of StandardLCFactory, bound to a specific deployed contract.
func NewStandardLCFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardLCFactoryTransactor, error) {
	contract, err := bindStandardLCFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardLCFactoryTransactor{contract: contract}, nil
}

// NewStandardLCFactoryFilterer creates a new log filterer instance of StandardLCFactory, bound to a specific deployed contract.
func NewStandardLCFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardLCFactoryFilterer, error) {
	contract, err := bindStandardLCFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardLCFactoryFilterer{contract: contract}, nil
}

// bindStandardLCFactory binds a generic wrapper to an already deployed contract.
func bindStandardLCFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardLCFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardLCFactory *StandardLCFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardLCFactory.Contract.StandardLCFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardLCFactory *StandardLCFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardLCFactory.Contract.StandardLCFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardLCFactory *StandardLCFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardLCFactory.Contract.StandardLCFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardLCFactory *StandardLCFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardLCFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardLCFactory *StandardLCFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardLCFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardLCFactory *StandardLCFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardLCFactory.Contract.contract.Transact(opts, method, params...)
}

// Amc is a free data retrieval call binding the contract method 0xf3737c04.
//
// Solidity: function amc() view returns(address)
func (_StandardLCFactory *StandardLCFactoryCaller) Amc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StandardLCFactory.contract.Call(opts, &out, "amc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Amc is a free data retrieval call binding the contract method 0xf3737c04.
//
// Solidity: function amc() view returns(address)
func (_StandardLCFactory *StandardLCFactorySession) Amc() (common.Address, error) {
	return _StandardLCFactory.Contract.Amc(&_StandardLCFactory.CallOpts)
}

// Amc is a free data retrieval call binding the contract method 0xf3737c04.
//
// Solidity: function amc() view returns(address)
func (_StandardLCFactory *StandardLCFactoryCallerSession) Amc() (common.Address, error) {
	return _StandardLCFactory.Contract.Amc(&_StandardLCFactory.CallOpts)
}

// GetLCAddress is a free data retrieval call binding the contract method 0x793e97c6.
//
// Solidity: function getLCAddress(uint256 _documentId) view returns(address[])
func (_StandardLCFactory *StandardLCFactoryCaller) GetLCAddress(opts *bind.CallOpts, _documentId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _StandardLCFactory.contract.Call(opts, &out, "getLCAddress", _documentId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLCAddress is a free data retrieval call binding the contract method 0x793e97c6.
//
// Solidity: function getLCAddress(uint256 _documentId) view returns(address[])
func (_StandardLCFactory *StandardLCFactorySession) GetLCAddress(_documentId *big.Int) ([]common.Address, error) {
	return _StandardLCFactory.Contract.GetLCAddress(&_StandardLCFactory.CallOpts, _documentId)
}

// GetLCAddress is a free data retrieval call binding the contract method 0x793e97c6.
//
// Solidity: function getLCAddress(uint256 _documentId) view returns(address[])
func (_StandardLCFactory *StandardLCFactoryCallerSession) GetLCAddress(_documentId *big.Int) ([]common.Address, error) {
	return _StandardLCFactory.Contract.GetLCAddress(&_StandardLCFactory.CallOpts, _documentId)
}

// Amend is a paid mutator transaction binding the contract method 0xa9464b75.
//
// Solidity: function amend(address _executor, uint256 _documentId, bytes32[] _parties) returns(address _contract)
func (_StandardLCFactory *StandardLCFactoryTransactor) Amend(opts *bind.TransactOpts, _executor common.Address, _documentId *big.Int, _parties [][32]byte) (*types.Transaction, error) {
	return _StandardLCFactory.contract.Transact(opts, "amend", _executor, _documentId, _parties)
}

// Amend is a paid mutator transaction binding the contract method 0xa9464b75.
//
// Solidity: function amend(address _executor, uint256 _documentId, bytes32[] _parties) returns(address _contract)
func (_StandardLCFactory *StandardLCFactorySession) Amend(_executor common.Address, _documentId *big.Int, _parties [][32]byte) (*types.Transaction, error) {
	return _StandardLCFactory.Contract.Amend(&_StandardLCFactory.TransactOpts, _executor, _documentId, _parties)
}

// Amend is a paid mutator transaction binding the contract method 0xa9464b75.
//
// Solidity: function amend(address _executor, uint256 _documentId, bytes32[] _parties) returns(address _contract)
func (_StandardLCFactory *StandardLCFactoryTransactorSession) Amend(_executor common.Address, _documentId *big.Int, _parties [][32]byte) (*types.Transaction, error) {
	return _StandardLCFactory.Contract.Amend(&_StandardLCFactory.TransactOpts, _executor, _documentId, _parties)
}

// Create is a paid mutator transaction binding the contract method 0x9e7b9dfe.
//
// Solidity: function create(bytes32[] _parties, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns(address _contract)
func (_StandardLCFactory *StandardLCFactoryTransactor) Create(opts *bind.TransactOpts, _parties [][32]byte, _content IStageContractContent) (*types.Transaction, error) {
	return _StandardLCFactory.contract.Transact(opts, "create", _parties, _content)
}

// Create is a paid mutator transaction binding the contract method 0x9e7b9dfe.
//
// Solidity: function create(bytes32[] _parties, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns(address _contract)
func (_StandardLCFactory *StandardLCFactorySession) Create(_parties [][32]byte, _content IStageContractContent) (*types.Transaction, error) {
	return _StandardLCFactory.Contract.Create(&_StandardLCFactory.TransactOpts, _parties, _content)
}

// Create is a paid mutator transaction binding the contract method 0x9e7b9dfe.
//
// Solidity: function create(bytes32[] _parties, (bytes32,uint256,bytes32,uint256,bytes32[],string,bytes,bytes) _content) returns(address _contract)
func (_StandardLCFactory *StandardLCFactoryTransactorSession) Create(_parties [][32]byte, _content IStageContractContent) (*types.Transaction, error) {
	return _StandardLCFactory.Contract.Create(&_StandardLCFactory.TransactOpts, _parties, _content)
}

// SetAMC is a paid mutator transaction binding the contract method 0x171cba35.
//
// Solidity: function setAMC(address _amc) returns()
func (_StandardLCFactory *StandardLCFactoryTransactor) SetAMC(opts *bind.TransactOpts, _amc common.Address) (*types.Transaction, error) {
	return _StandardLCFactory.contract.Transact(opts, "setAMC", _amc)
}

// SetAMC is a paid mutator transaction binding the contract method 0x171cba35.
//
// Solidity: function setAMC(address _amc) returns()
func (_StandardLCFactory *StandardLCFactorySession) SetAMC(_amc common.Address) (*types.Transaction, error) {
	return _StandardLCFactory.Contract.SetAMC(&_StandardLCFactory.TransactOpts, _amc)
}

// SetAMC is a paid mutator transaction binding the contract method 0x171cba35.
//
// Solidity: function setAMC(address _amc) returns()
func (_StandardLCFactory *StandardLCFactoryTransactorSession) SetAMC(_amc common.Address) (*types.Transaction, error) {
	return _StandardLCFactory.Contract.SetAMC(&_StandardLCFactory.TransactOpts, _amc)
}

// StandardLCFactoryNewStandardLCIterator is returned from FilterNewStandardLC and is used to iterate over the raw logs and unpacked data for NewStandardLC events raised by the StandardLCFactory contract.
type StandardLCFactoryNewStandardLCIterator struct {
	Event *StandardLCFactoryNewStandardLC // Event containing the contract specifics and raw log

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
func (it *StandardLCFactoryNewStandardLCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardLCFactoryNewStandardLC)
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
		it.Event = new(StandardLCFactoryNewStandardLC)
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
func (it *StandardLCFactoryNewStandardLCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardLCFactoryNewStandardLCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardLCFactoryNewStandardLC represents a NewStandardLC event raised by the StandardLCFactory contract.
type StandardLCFactoryNewStandardLC struct {
	DocumentID     *big.Int
	Creator        common.Address
	LcContractAddr common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewStandardLC is a free log retrieval operation binding the contract event 0x19412006b3c99a5b101aa57f5fe2d0393d999e864ee328c4e662211b7070c050.
//
// Solidity: event NewStandardLC(uint256 indexed documentID, address indexed creator, address indexed lcContractAddr)
func (_StandardLCFactory *StandardLCFactoryFilterer) FilterNewStandardLC(opts *bind.FilterOpts, documentID []*big.Int, creator []common.Address, lcContractAddr []common.Address) (*StandardLCFactoryNewStandardLCIterator, error) {

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

	logs, sub, err := _StandardLCFactory.contract.FilterLogs(opts, "NewStandardLC", documentIDRule, creatorRule, lcContractAddrRule)
	if err != nil {
		return nil, err
	}
	return &StandardLCFactoryNewStandardLCIterator{contract: _StandardLCFactory.contract, event: "NewStandardLC", logs: logs, sub: sub}, nil
}

var NewStandardLCTopicHash = "0x19412006b3c99a5b101aa57f5fe2d0393d999e864ee328c4e662211b7070c050"

// WatchNewStandardLC is a free log subscription operation binding the contract event 0x19412006b3c99a5b101aa57f5fe2d0393d999e864ee328c4e662211b7070c050.
//
// Solidity: event NewStandardLC(uint256 indexed documentID, address indexed creator, address indexed lcContractAddr)
func (_StandardLCFactory *StandardLCFactoryFilterer) WatchNewStandardLC(opts *bind.WatchOpts, sink chan<- *StandardLCFactoryNewStandardLC, documentID []*big.Int, creator []common.Address, lcContractAddr []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StandardLCFactory.contract.WatchLogs(opts, "NewStandardLC", documentIDRule, creatorRule, lcContractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardLCFactoryNewStandardLC)
				if err := _StandardLCFactory.contract.UnpackLog(event, "NewStandardLC", log); err != nil {
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

// ParseNewStandardLC is a log parse operation binding the contract event 0x19412006b3c99a5b101aa57f5fe2d0393d999e864ee328c4e662211b7070c050.
//
// Solidity: event NewStandardLC(uint256 indexed documentID, address indexed creator, address indexed lcContractAddr)
func (_StandardLCFactory *StandardLCFactoryFilterer) ParseNewStandardLC(log types.Log) (*StandardLCFactoryNewStandardLC, error) {
	event := new(StandardLCFactoryNewStandardLC)
	if err := _StandardLCFactory.contract.UnpackLog(event, "NewStandardLC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
