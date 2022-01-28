// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

//go:generate mockgen -source interface.go -destination mock_interface.go -package vm

package vm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
)

// Quorum

type AccountExtraDataStateGetter interface {
	// Return nil for public contract
	GetPrivacyMetadata(addr common.Address) (*state.PrivacyMetadata, error)
	GetManagedParties(addr common.Address) ([]string, error)
}

type AccountExtraDataStateSetter interface {
	SetPrivacyMetadata(addr common.Address, pm *state.PrivacyMetadata)
	SetManagedParties(addr common.Address, managedParties []string)
}

// Quorum uses a cut-down StateDB, MinimalApiState. We leave the methods in StateDB commented out so they'll produce a
// conflict when upstream changes.
type MinimalApiState interface {
	AccountExtraDataStateGetter

	GetBalance(addr common.Address) *big.Int
	SetBalance(addr common.Address, balance *big.Int)
	GetCode(addr common.Address) []byte
	GetState(a common.Address, b common.Hash) common.Hash
	GetNonce(addr common.Address) uint64
	SetNonce(addr common.Address, nonce uint64)
	SetCode(common.Address, []byte)

	// RLP-encoded of the state object in a given address
	// Throw error if no state object is found
	GetRLPEncodedStateObject(addr common.Address) ([]byte, error)
	GetProof(common.Address) ([][]byte, error)
	GetStorageProof(common.Address, common.Hash) ([][]byte, error)
	StorageTrie(addr common.Address) state.Trie
	Error() error
	GetCodeHash(common.Address) common.Hash
	SetState(common.Address, common.Hash, common.Hash)
	SetStorage(addr common.Address, storage map[common.Hash]common.Hash)
}

// End Quorum

// StateDB is an EVM database for full state querying.
type StateDB interface {
	// Quorum

	MinimalApiState
	AccountExtraDataStateSetter
	// End Quorum

	CreateAccount(common.Address)

	SubBalance(common.Address, *big.Int)
	AddBalance(common.Address, *big.Int)
	//GetBalance(common.Address) *big.Int

	//GetNonce(common.Address) uint64
	//SetNonce(common.Address, uint64)

	//GetCodeHash(common.Address) common.Hash
	//GetCode(common.Address) []byte
	//SetCode(common.Address, []byte)

	GetCodeSize(common.Address) int

	AddRefund(uint64)
	SubRefund(uint64)
	GetRefund() uint64

	GetCommittedState(common.Address, common.Hash) common.Hash
	//GetState(common.Address, common.Hash) common.Hash
	//SetState(common.Address, common.Hash, common.Hash)

	Suicide(common.Address) bool
	HasSuicided(common.Address) bool

	// Exist reports whether the given account exists in state.
	// Notably this should also return true for suicided accounts.
	Exist(common.Address) bool
	// Empty returns whether the given account is empty. Empty
	// is defined according to EIP161 (balance = nonce = code = 0).
	Empty(common.Address) bool

	PrepareAccessList(sender common.Address, dest *common.Address, precompiles []common.Address, txAccesses types.AccessList)
	AddressInAccessList(addr common.Address) bool
	SlotInAccessList(addr common.Address, slot common.Hash) (addressOk bool, slotOk bool)
	// AddAddressToAccessList adds the given address to the access list. This operation is safe to perform
	// even if the feature/fork is not active yet
	AddAddressToAccessList(addr common.Address)
	// AddSlotToAccessList adds the given (address,slot) to the access list. This operation is safe to perform
	// even if the feature/fork is not active yet
	AddSlotToAccessList(addr common.Address, slot common.Hash)

	RevertToSnapshot(int)
	Snapshot() int

	AddLog(*types.Log)
	AddPreimage(common.Hash, []byte)

	ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) error
}

// CallContext provides a basic interface for the EVM calling conventions. The EVM
// depends on this context being implemented for doing subcalls and initialising new EVM contracts.
type CallContext interface {
	// Call another contract
	Call(env *EVM, me ContractRef, addr common.Address, data []byte, gas, value *big.Int) ([]byte, error)
	// Take another's contract code and execute within our own context
	CallCode(env *EVM, me ContractRef, addr common.Address, data []byte, gas, value *big.Int) ([]byte, error)
	// Same as CallCode except sender and value is propagated from parent to child scope
	DelegateCall(env *EVM, me ContractRef, addr common.Address, data []byte, gas *big.Int) ([]byte, error)
	// Create a new contract
	Create(env *EVM, me ContractRef, data []byte, gas, value *big.Int) ([]byte, common.Address, error)
}
