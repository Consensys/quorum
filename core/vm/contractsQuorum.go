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
// Retrieves private transaction from Tessera and executes it.
// If we are not a participant, then just ensure public state remains in sync.
//		input = 20 byte address of sender, 64 byte hash for the private transaction
func (c *privacyMarker) Run(evm *EVM, input []byte) ([]byte, error) {
	log.Debug("Running privacy marker precompile")

	if evm.currentTx.IsPrivate() {
		//only public transactions can call the precompile
		log.Warn("Private transaction called precompile", "tx hash", evm.currentTx.Hash())
		return nil, nil
	}

	data := evm.currentTx.Data()
	txHash := common.BytesToEncryptedPayloadHash(data[20:])
	_, _, txData, _, err := private.P.Receive(txHash) //TODO: should use returned metadata...
	if err != nil {
		log.Error("Failed to retrieve transaction from private transaction manager", "err", err)
		return nil, err
	}

	//TODO (peter): sender from tx data should be removed when possible
	fromAddr := common.BytesToAddress(data[:20])

	if txData == nil {
		log.Debug("not a participant, precompile performing no action")
		// must increment the nonce to mirror the state change that is done in evm.create() for participants
		evm.publicState.SetNonce(fromAddr, evm.publicState.GetNonce(fromAddr)+1)
		return nil, nil
	}

	var tx types.Transaction
	if err := json.NewDecoder(bytes.NewReader(txData)).Decode(&tx); err != nil {
		log.Trace("failed to deserialize privacy marker transaction", "err", err)
		return nil, err
	}

	if !tx.IsPrivate() {
		//should only allow private txns from inside precompile, as many assumptions
		//about how a tx operates are based on its privacy (e.g. which dbs to use, PE checks etc)
		log.Warn("Public transaction pulled from PTM during privacy precompile execution")

		// non-participants have already incremented the public nonce, so we need to as well
		evm.publicState.SetNonce(fromAddr, evm.publicState.GetNonce(fromAddr)+1)

		return nil, nil
	}

	_, err = c.runUsingNewEVM(evm, tx)

	return nil, err
}

func (c *privacyMarker) runUsingNewEVM(evm *EVM, tx types.Transaction) ([]byte, error) {
	_, err := evm.InnerApply(&tx)
	return nil, err
}
