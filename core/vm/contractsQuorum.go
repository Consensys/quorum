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
	"github.com/ethereum/go-ethereum/common"
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
func QuorumRunPrecompiledContract(evm *EVM, p QuorumPrecompiledContract, input []byte, contract *Contract) (ret []byte, err error) {
	gas := p.RequiredGas(input)
	if contract.UseGas(gas) {
		return p.Run(evm, input)
	}
	return nil, ErrOutOfGas
}

type privacyMarker struct{}

const privacyMarkerGas uint64 = 3000 //TODO: needs to match Besu gas usage

func PrivacyMarkerAddress() common.Address {
	return common.BytesToAddress([]byte{0x7f, 0xff, 0xff, 0xff}) //using Address = MaxInt32
}

func (c *privacyMarker) RequiredGas(input []byte) uint64 {
	return privacyMarkerGas
}

// privacyMarker precompile execution
// retrieves transaction data from Tessera and executes it (if we are a participant)
//		input = 64 byte private hash for the private transaction
func (c *privacyMarker) Run(evm *EVM, input []byte) ([]byte, error) {
	encryptedHash := common.BytesToEncryptedPayloadHash(input)
	_, _, data, _, err := private.P.Receive(encryptedHash)
	if err != nil {
		log.Error("Failed to retrieve transaction from private transaction manager", "err", err)
		return nil, err
	}

	//if this node is not a participant then no action
	if data == nil {
		log.Trace("not a participant, precompile performing no action")
		return nil, nil
	}

	return c.runUsingSandboxEVM(evm, data)
}

func (c *privacyMarker) runUsingSandboxEVM(evm *EVM, data []byte) ([]byte, error) {
	log.Warn("privacy marker transaction not yet supported - no action performed")
	return nil, nil
}
