// Copyright 2014 The go-ethereum Authors
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

package vm

import (
	"bytes"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/private"
)

// QuorumPrecompiledContract is an extended interface for native Quorum Go contracts. The implementation
// requires a deterministic gas count based on the input size of the Run method of the
// contract.
type QuorumPrecompiledContract interface {
	RequiredGas(input []byte) uint64            // RequiredPrice calculates the contract gas use
	Run(evm *EVM, input []byte) ([]byte, error) // Run runs the precompiled contract
}

// Ths contains the default set of pre-compiled Quorum contracts (with an extended interface).
var QuorumPrecompiledContracts = map[common.Address]QuorumPrecompiledContract{
	PrivacyMarkerAddress(): &privacyMarker{},
}

// QuorumRunPrecompiledContract runs and evaluates the output of an extended precompiled contract.
// It returns
// - the returned bytes,
// - the _remaining_ gas,
// - any error that occurred
func QuorumRunPrecompiledContract(evm *EVM, p QuorumPrecompiledContract, input []byte, suppliedGas uint64) (ret []byte, remainingGas uint64, err error) {
	gasCost := p.RequiredGas(input)
	if suppliedGas < gasCost {
		return nil, 0, ErrOutOfGas
	}
	suppliedGas -= gasCost
	output, err := p.Run(evm, input)
	return output, suppliedGas, err
}

type privacyMarker struct{}

func PrivacyMarkerAddress() common.Address {
	return common.BytesToAddress([]byte{0x7f, 0xff, 0xff, 0xff}) //using Address = MaxInt32
}

func (c *privacyMarker) RequiredGas(input []byte) uint64 {
	return uint64(0)
}

// privacyMarker precompile execution
// retrieves transaction data from Tessera and executes it (if we are a participant)
//		input = 64 byte hash for the private transaction
func (c *privacyMarker) Run(evm *EVM, input []byte) ([]byte, error) {
	log.Debug("Running privacy marker precompile")

	txHash := common.BytesToEncryptedPayloadHash(evm.currentTx.Data())
	_, _, txData, _, err := private.P.Receive(txHash) //TODO: should use returned metadata...
	if err != nil {
		log.Error("Failed to retrieve transaction from private transaction manager", "err", err)
		return nil, err
	}
	if txData == nil {
		log.Debug("not a participant, precompile performing no action")
		return nil, nil
	}

	var tx types.Transaction
	err = json.NewDecoder(bytes.NewReader(txData)).Decode(&tx)
	if err != nil {
		log.Trace("failed to deserialize privacy marker transaction", "err", err)
		return nil, err
	}

	return c.runUsingSandboxEVM(evm, tx, tx.Data())
}

func (c *privacyMarker) runUsingSandboxEVM(evm *EVM, tx types.Transaction, data []byte) ([]byte, error) {

	/* TODO: should use something like this context for the EVM:
	   // Setup context with timeout as gas un-metered
	   var cancel context.CancelFunc
	   ctx, cancel = context.WithTimeout(ctx, time.Second*5)
	   // Make sure the context is cancelled when the call has completed
	   // this makes sure resources are cleaned up.
	   defer func() { cancel() }()
	*/

	//TODO: may need to create a new vm.Context (because of multitenancy, as per GetEVM())
	privateState := evm.publicState
	if tx.To() != nil && !privateState.Exist(*tx.To()) {
		privateState = evm.SavedPrivateState
	}

	sandboxEVM := NewEVM(evm.Context, evm.publicState, privateState, evm.chainConfig, evm.vmConfig)
	sandboxEVM.SetCurrentTX(&tx)

	/* TODO:
	// Wait for the context to be done and cancel the evm. Even if the
	// EVM has finished, cancelling may be done (repeatedly)
	go func() {
	   <-ctx.Done()
	   evm.Cancel()
	}()
	*/

	ret, _, err := sandboxEVM.Call(AccountRef(tx.From()), *tx.To(), data, tx.Gas(), tx.Value())

	return ret, err
}
