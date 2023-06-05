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
	"errors"

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

// QuorumPrecompiledContracts is the default set of pre-compiled Quorum contracts (with an extended interface).
var QuorumPrecompiledContracts = map[common.Address]QuorumPrecompiledContract{
	common.QuorumPrivacyPrecompileContractAddress(): &privacyMarker{},
}

// RunQuorumPrecompiledContract runs and evaluates the output of an extended precompiled contract.
// It returns
// - the returned bytes,
// - the _remaining_ gas,
// - any error that occurred
func RunQuorumPrecompiledContract(evm *EVM, p QuorumPrecompiledContract, input []byte, suppliedGas uint64) (ret []byte, remainingGas uint64, err error) {
	gasCost := p.RequiredGas(input)
	if suppliedGas < gasCost {
		return nil, 0, ErrOutOfGas
	}
	suppliedGas -= gasCost
	output, err := p.Run(evm, input)
	return output, suppliedGas, err
}

type privacyMarker struct{}

func (c *privacyMarker) RequiredGas(_ []byte) uint64 {
	return uint64(0)
}

// privacyMarker precompile execution
// Retrieves private transaction from Tessera and executes it.
// If we are not a participant, then just ensure public state remains in sync.
//
//	input = 20 byte address of sender, 64 byte hash for the private transaction
func (c *privacyMarker) Run(evm *EVM, _ []byte) ([]byte, error) {
	log.Debug("Running privacy marker precompile")

	// support vanilla ethereum tests where tx is not set
	if evm.currentTx == nil {
		return nil, nil
	}
	logger := log.New("pmtHash", evm.currentTx.Hash())

	if evm.depth != 0 || !evm.currentTx.IsPrivacyMarker() {
		// only supporting direct precompile calls so far
		logger.Warn("Invalid privacy marker precompile execution")
		return nil, nil
	}

	if evm.currentTx.IsPrivate() {
		//only public transactions can call the precompile
		logger.Warn("PMT is not a public transaction")
		return nil, nil
	}

	tx, _, _, err := private.FetchPrivateTransaction(evm.currentTx.Data())
	if err != nil {
		logger.Error("Failed to retrieve inner transaction from private transaction manager", "err", err)
		return nil, nil
	}

	if tx == nil {
		logger.Debug("Not a participant, skipping execution")
		return nil, nil
	}

	if !tx.IsPrivate() {
		//should only allow private txns from inside precompile, as many assumptions
		//about how a tx operates are based on its privacy (e.g. which dbs to use, PE checks etc)
		logger.Warn("Inner transaction retrieved from private transaction manager is not a private transaction, skipping execution")
		return nil, nil
	}
	//validate the private tx is signed, and that it's the same signer as the PMT
	signedBy := tx.From()
	if signedBy.Hex() == (common.Address{}).Hex() || signedBy.Hex() != evm.currentTx.From().Hex() {
		logger.Warn("PMT and inner private transaction have different signers, skipping execution")
		return nil, nil
	}

	// validate the private tx has the same nonce as the PMT
	if tx.Nonce() != evm.currentTx.Nonce() {
		logger.Warn("PMT and inner private transaction have different nonces, skipping execution")
		return nil, nil
	}

	if err := applyTransactionWithoutIncrementingNonce(evm, tx); err != nil {
		logger.Warn("Unable to apply PMT's inner transaction to EVM, skipping execution", "err", err)
		return nil, nil
	}
	logger.Debug("Inner private transaction applied")
	return nil, nil
}

// Effectively execute the internal private transaction without incrementing the nonce of the sender account.
// (1)  make a copy of the sender's starting (i.e. current) account nonce.
// (2)  decrement the sender's account nonce in the public state so that the internal private transaction (which has the same 'from' and 'nonce' as the outer PMT) can be executed.
// (3a)  execute the internal private transaction.
// (3b) if the internal private tx is successfully executed then the sender's account nonce will be incremented back to the starting nonce.
// (3c) if the execution was unsuccessful then the nonce may not be incremented.
// (4)  force reset the nonce to the starting value in any case.
func applyTransactionWithoutIncrementingNonce(evm *EVM, tx *types.Transaction) error {
	if evm.InnerApply == nil {
		return errors.New("nil inner apply function")
	}

	fromAddr := evm.currentTx.From()

	startingNonce := evm.PublicState().GetNonce(fromAddr)
	evm.publicState.SetNonce(fromAddr, startingNonce-1)
	defer evm.publicState.SetNonce(fromAddr, startingNonce)

	return evm.InnerApply(tx)
}
