// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package quorum

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// VotingContractABI is the input ABI used to generate the binding from.
const VotingContractABI = `[{"constant":false,"inputs":[{"name":"threshold","type":"uint256"}],"name":"setVoteThreshold","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"removeBlockMaker","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"voterCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"canCreateBlocks","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"voteThreshold","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"height","type":"uint256"}],"name":"getCanonHash","outputs":[{"name":"","type":"bytes32"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"height","type":"uint256"},{"name":"hash","type":"bytes32"}],"name":"vote","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"addBlockMaker","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"removeVoter","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"height","type":"uint256"},{"name":"n","type":"uint256"}],"name":"getEntry","outputs":[{"name":"","type":"bytes32"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"addr","type":"address"}],"name":"isVoter","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"canVote","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"blockMakerCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"getSize","outputs":[{"name":"","type":"uint256"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"addr","type":"address"}],"name":"isBlockMaker","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"addVoter","outputs":[],"payable":false,"type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"sender","type":"address"},{"indexed":false,"name":"blockNumber","type":"uint256"},{"indexed":false,"name":"blockHash","type":"bytes32"}],"name":"Vote","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"","type":"address"}],"name":"AddVoter","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"","type":"address"}],"name":"RemovedVoter","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"","type":"address"}],"name":"AddBlockMaker","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"","type":"address"}],"name":"RemovedBlockMaker","type":"event"}]`

// VotingContract is an auto generated Go binding around an Ethereum contract.
type VotingContract struct {
	VotingContractCaller     // Read-only binding to the contract
	VotingContractTransactor // Write-only binding to the contract
}

// VotingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingContractSession struct {
	Contract     *VotingContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingContractCallerSession struct {
	Contract *VotingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VotingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingContractTransactorSession struct {
	Contract     *VotingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VotingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingContractRaw struct {
	Contract *VotingContract // Generic contract binding to access the raw methods on
}

// VotingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingContractCallerRaw struct {
	Contract *VotingContractCaller // Generic read-only contract binding to access the raw methods on
}

// VotingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingContractTransactorRaw struct {
	Contract *VotingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVotingContract creates a new instance of VotingContract, bound to a specific deployed contract.
func NewVotingContract(address common.Address, backend bind.ContractBackend) (*VotingContract, error) {
	contract, err := bindVotingContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VotingContract{VotingContractCaller: VotingContractCaller{contract: contract}, VotingContractTransactor: VotingContractTransactor{contract: contract}}, nil
}

// NewVotingContractCaller creates a new read-only instance of VotingContract, bound to a specific deployed contract.
func NewVotingContractCaller(address common.Address, caller bind.ContractCaller) (*VotingContractCaller, error) {
	contract, err := bindVotingContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &VotingContractCaller{contract: contract}, nil
}

// NewVotingContractTransactor creates a new write-only instance of VotingContract, bound to a specific deployed contract.
func NewVotingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingContractTransactor, error) {
	contract, err := bindVotingContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &VotingContractTransactor{contract: contract}, nil
}

// bindVotingContract binds a generic wrapper to an already deployed contract.
func bindVotingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VotingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingContract *VotingContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VotingContract.Contract.VotingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingContract *VotingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingContract.Contract.VotingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingContract *VotingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingContract.Contract.VotingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VotingContract *VotingContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VotingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VotingContract *VotingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VotingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VotingContract *VotingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VotingContract.Contract.contract.Transact(opts, method, params...)
}

// BlockMakerCount is a free data retrieval call binding the contract method 0xcf528985.
//
// Solidity: function blockMakerCount() constant returns(uint256)
func (_VotingContract *VotingContractCaller) BlockMakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "blockMakerCount")
	return *ret0, err
}

// BlockMakerCount is a free data retrieval call binding the contract method 0xcf528985.
//
// Solidity: function blockMakerCount() constant returns(uint256)
func (_VotingContract *VotingContractSession) BlockMakerCount() (*big.Int, error) {
	return _VotingContract.Contract.BlockMakerCount(&_VotingContract.CallOpts)
}

// BlockMakerCount is a free data retrieval call binding the contract method 0xcf528985.
//
// Solidity: function blockMakerCount() constant returns(uint256)
func (_VotingContract *VotingContractCallerSession) BlockMakerCount() (*big.Int, error) {
	return _VotingContract.Contract.BlockMakerCount(&_VotingContract.CallOpts)
}

// CanCreateBlocks is a free data retrieval call binding the contract method 0x488099a6.
//
// Solidity: function canCreateBlocks( address) constant returns(bool)
func (_VotingContract *VotingContractCaller) CanCreateBlocks(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "canCreateBlocks", arg0)
	return *ret0, err
}

// CanCreateBlocks is a free data retrieval call binding the contract method 0x488099a6.
//
// Solidity: function canCreateBlocks( address) constant returns(bool)
func (_VotingContract *VotingContractSession) CanCreateBlocks(arg0 common.Address) (bool, error) {
	return _VotingContract.Contract.CanCreateBlocks(&_VotingContract.CallOpts, arg0)
}

// CanCreateBlocks is a free data retrieval call binding the contract method 0x488099a6.
//
// Solidity: function canCreateBlocks( address) constant returns(bool)
func (_VotingContract *VotingContractCallerSession) CanCreateBlocks(arg0 common.Address) (bool, error) {
	return _VotingContract.Contract.CanCreateBlocks(&_VotingContract.CallOpts, arg0)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote( address) constant returns(bool)
func (_VotingContract *VotingContractCaller) CanVote(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "canVote", arg0)
	return *ret0, err
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote( address) constant returns(bool)
func (_VotingContract *VotingContractSession) CanVote(arg0 common.Address) (bool, error) {
	return _VotingContract.Contract.CanVote(&_VotingContract.CallOpts, arg0)
}

// CanVote is a free data retrieval call binding the contract method 0xadfaa72e.
//
// Solidity: function canVote( address) constant returns(bool)
func (_VotingContract *VotingContractCallerSession) CanVote(arg0 common.Address) (bool, error) {
	return _VotingContract.Contract.CanVote(&_VotingContract.CallOpts, arg0)
}

// GetCanonHash is a free data retrieval call binding the contract method 0x559c390c.
//
// Solidity: function getCanonHash(height uint256) constant returns(bytes32)
func (_VotingContract *VotingContractCaller) GetCanonHash(opts *bind.CallOpts, height *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "getCanonHash", height)
	return *ret0, err
}

// GetCanonHash is a free data retrieval call binding the contract method 0x559c390c.
//
// Solidity: function getCanonHash(height uint256) constant returns(bytes32)
func (_VotingContract *VotingContractSession) GetCanonHash(height *big.Int) ([32]byte, error) {
	return _VotingContract.Contract.GetCanonHash(&_VotingContract.CallOpts, height)
}

// GetCanonHash is a free data retrieval call binding the contract method 0x559c390c.
//
// Solidity: function getCanonHash(height uint256) constant returns(bytes32)
func (_VotingContract *VotingContractCallerSession) GetCanonHash(height *big.Int) ([32]byte, error) {
	return _VotingContract.Contract.GetCanonHash(&_VotingContract.CallOpts, height)
}

// GetEntry is a free data retrieval call binding the contract method 0x98ba676d.
//
// Solidity: function getEntry(height uint256, n uint256) constant returns(bytes32)
func (_VotingContract *VotingContractCaller) GetEntry(opts *bind.CallOpts, height *big.Int, n *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "getEntry", height, n)
	return *ret0, err
}

// GetEntry is a free data retrieval call binding the contract method 0x98ba676d.
//
// Solidity: function getEntry(height uint256, n uint256) constant returns(bytes32)
func (_VotingContract *VotingContractSession) GetEntry(height *big.Int, n *big.Int) ([32]byte, error) {
	return _VotingContract.Contract.GetEntry(&_VotingContract.CallOpts, height, n)
}

// GetEntry is a free data retrieval call binding the contract method 0x98ba676d.
//
// Solidity: function getEntry(height uint256, n uint256) constant returns(bytes32)
func (_VotingContract *VotingContractCallerSession) GetEntry(height *big.Int, n *big.Int) ([32]byte, error) {
	return _VotingContract.Contract.GetEntry(&_VotingContract.CallOpts, height, n)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_VotingContract *VotingContractCaller) GetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "getSize")
	return *ret0, err
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_VotingContract *VotingContractSession) GetSize() (*big.Int, error) {
	return _VotingContract.Contract.GetSize(&_VotingContract.CallOpts)
}

// GetSize is a free data retrieval call binding the contract method 0xde8fa431.
//
// Solidity: function getSize() constant returns(uint256)
func (_VotingContract *VotingContractCallerSession) GetSize() (*big.Int, error) {
	return _VotingContract.Contract.GetSize(&_VotingContract.CallOpts)
}

// IsBlockMaker is a free data retrieval call binding the contract method 0xe814d1c7.
//
// Solidity: function isBlockMaker(addr address) constant returns(bool)
func (_VotingContract *VotingContractCaller) IsBlockMaker(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "isBlockMaker", addr)
	return *ret0, err
}

// IsBlockMaker is a free data retrieval call binding the contract method 0xe814d1c7.
//
// Solidity: function isBlockMaker(addr address) constant returns(bool)
func (_VotingContract *VotingContractSession) IsBlockMaker(addr common.Address) (bool, error) {
	return _VotingContract.Contract.IsBlockMaker(&_VotingContract.CallOpts, addr)
}

// IsBlockMaker is a free data retrieval call binding the contract method 0xe814d1c7.
//
// Solidity: function isBlockMaker(addr address) constant returns(bool)
func (_VotingContract *VotingContractCallerSession) IsBlockMaker(addr common.Address) (bool, error) {
	return _VotingContract.Contract.IsBlockMaker(&_VotingContract.CallOpts, addr)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(bool)
func (_VotingContract *VotingContractCaller) IsVoter(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "isVoter", addr)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(bool)
func (_VotingContract *VotingContractSession) IsVoter(addr common.Address) (bool, error) {
	return _VotingContract.Contract.IsVoter(&_VotingContract.CallOpts, addr)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(addr address) constant returns(bool)
func (_VotingContract *VotingContractCallerSession) IsVoter(addr common.Address) (bool, error) {
	return _VotingContract.Contract.IsVoter(&_VotingContract.CallOpts, addr)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() constant returns(uint256)
func (_VotingContract *VotingContractCaller) VoteThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "voteThreshold")
	return *ret0, err
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() constant returns(uint256)
func (_VotingContract *VotingContractSession) VoteThreshold() (*big.Int, error) {
	return _VotingContract.Contract.VoteThreshold(&_VotingContract.CallOpts)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() constant returns(uint256)
func (_VotingContract *VotingContractCallerSession) VoteThreshold() (*big.Int, error) {
	return _VotingContract.Contract.VoteThreshold(&_VotingContract.CallOpts)
}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() constant returns(uint256)
func (_VotingContract *VotingContractCaller) VoterCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VotingContract.contract.Call(opts, out, "voterCount")
	return *ret0, err
}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() constant returns(uint256)
func (_VotingContract *VotingContractSession) VoterCount() (*big.Int, error) {
	return _VotingContract.Contract.VoterCount(&_VotingContract.CallOpts)
}

// VoterCount is a free data retrieval call binding the contract method 0x42169e48.
//
// Solidity: function voterCount() constant returns(uint256)
func (_VotingContract *VotingContractCallerSession) VoterCount() (*big.Int, error) {
	return _VotingContract.Contract.VoterCount(&_VotingContract.CallOpts)
}

// AddBlockMaker is a paid mutator transaction binding the contract method 0x72a571fc.
//
// Solidity: function addBlockMaker(addr address) returns()
func (_VotingContract *VotingContractTransactor) AddBlockMaker(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "addBlockMaker", addr)
}

// AddBlockMaker is a paid mutator transaction binding the contract method 0x72a571fc.
//
// Solidity: function addBlockMaker(addr address) returns()
func (_VotingContract *VotingContractSession) AddBlockMaker(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.AddBlockMaker(&_VotingContract.TransactOpts, addr)
}

// AddBlockMaker is a paid mutator transaction binding the contract method 0x72a571fc.
//
// Solidity: function addBlockMaker(addr address) returns()
func (_VotingContract *VotingContractTransactorSession) AddBlockMaker(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.AddBlockMaker(&_VotingContract.TransactOpts, addr)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(addr address) returns()
func (_VotingContract *VotingContractTransactor) AddVoter(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "addVoter", addr)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(addr address) returns()
func (_VotingContract *VotingContractSession) AddVoter(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.AddVoter(&_VotingContract.TransactOpts, addr)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(addr address) returns()
func (_VotingContract *VotingContractTransactorSession) AddVoter(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.AddVoter(&_VotingContract.TransactOpts, addr)
}

// RemoveBlockMaker is a paid mutator transaction binding the contract method 0x284d163c.
//
// Solidity: function removeBlockMaker(addr address) returns()
func (_VotingContract *VotingContractTransactor) RemoveBlockMaker(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "removeBlockMaker", addr)
}

// RemoveBlockMaker is a paid mutator transaction binding the contract method 0x284d163c.
//
// Solidity: function removeBlockMaker(addr address) returns()
func (_VotingContract *VotingContractSession) RemoveBlockMaker(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.RemoveBlockMaker(&_VotingContract.TransactOpts, addr)
}

// RemoveBlockMaker is a paid mutator transaction binding the contract method 0x284d163c.
//
// Solidity: function removeBlockMaker(addr address) returns()
func (_VotingContract *VotingContractTransactorSession) RemoveBlockMaker(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.RemoveBlockMaker(&_VotingContract.TransactOpts, addr)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(addr address) returns()
func (_VotingContract *VotingContractTransactor) RemoveVoter(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "removeVoter", addr)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(addr address) returns()
func (_VotingContract *VotingContractSession) RemoveVoter(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.RemoveVoter(&_VotingContract.TransactOpts, addr)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(addr address) returns()
func (_VotingContract *VotingContractTransactorSession) RemoveVoter(addr common.Address) (*types.Transaction, error) {
	return _VotingContract.Contract.RemoveVoter(&_VotingContract.TransactOpts, addr)
}

// SetVoteThreshold is a paid mutator transaction binding the contract method 0x12909485.
//
// Solidity: function setVoteThreshold(threshold uint256) returns()
func (_VotingContract *VotingContractTransactor) SetVoteThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "setVoteThreshold", threshold)
}

// SetVoteThreshold is a paid mutator transaction binding the contract method 0x12909485.
//
// Solidity: function setVoteThreshold(threshold uint256) returns()
func (_VotingContract *VotingContractSession) SetVoteThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _VotingContract.Contract.SetVoteThreshold(&_VotingContract.TransactOpts, threshold)
}

// SetVoteThreshold is a paid mutator transaction binding the contract method 0x12909485.
//
// Solidity: function setVoteThreshold(threshold uint256) returns()
func (_VotingContract *VotingContractTransactorSession) SetVoteThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _VotingContract.Contract.SetVoteThreshold(&_VotingContract.TransactOpts, threshold)
}

// Vote is a paid mutator transaction binding the contract method 0x68bb8bb6.
//
// Solidity: function vote(height uint256, hash bytes32) returns()
func (_VotingContract *VotingContractTransactor) Vote(opts *bind.TransactOpts, height *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _VotingContract.contract.Transact(opts, "vote", height, hash)
}

// Vote is a paid mutator transaction binding the contract method 0x68bb8bb6.
//
// Solidity: function vote(height uint256, hash bytes32) returns()
func (_VotingContract *VotingContractSession) Vote(height *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _VotingContract.Contract.Vote(&_VotingContract.TransactOpts, height, hash)
}

// Vote is a paid mutator transaction binding the contract method 0x68bb8bb6.
//
// Solidity: function vote(height uint256, hash bytes32) returns()
func (_VotingContract *VotingContractTransactorSession) Vote(height *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _VotingContract.Contract.Vote(&_VotingContract.TransactOpts, height, hash)
}
