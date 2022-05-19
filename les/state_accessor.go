// Copyright 2021 The go-ethereum Authors
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

package les

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/light"
)

// stateAtBlock retrieves the state database associated with a certain block.
func (leth *LightEthereum) stateAtBlock(ctx context.Context, block *types.Block, reexec uint64) (*state.StateDB, mps.PrivateStateRepository, error) {
	return light.NewState(ctx, block.Header(), leth.odr), nil, nil
}

// stateAtTransaction returns the execution environment of a certain transaction.
func (leth *LightEthereum) stateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (core.Message, vm.BlockContext, *state.StateDB, *state.StateDB, mps.PrivateStateRepository, error) {
	// Short circuit if it's genesis block.
	if block.NumberU64() == 0 {
		return nil, vm.BlockContext{}, nil, nil, nil, errors.New("no transaction in genesis")
	}
	// Create the parent state database
	parent, err := leth.blockchain.GetBlock(ctx, block.ParentHash(), block.NumberU64()-1)
	if err != nil {
		return nil, vm.BlockContext{}, nil, nil, nil, err
	}
	statedb, private, err := leth.stateAtBlock(ctx, parent, reexec)
	if err != nil {
		return nil, vm.BlockContext{}, nil, nil, nil, err
	}
	privateStateDb, err := private.DefaultState()
	if err != nil {
		return nil, vm.BlockContext{}, nil, nil, nil, err
	}
	if txIndex == 0 && len(block.Transactions()) == 0 {
		return nil, vm.BlockContext{}, statedb, privateStateDb, private, nil
	}
	// Recompute transactions up to the target index.
	signer := types.MakeSigner(leth.blockchain.Config(), block.Number())
	for idx, tx := range block.Transactions() {
		// Assemble the transaction call message and return if the requested offset
		msg, _ := tx.AsMessage(signer)
		txContext := core.NewEVMTxContext(msg)
		context := core.NewEVMBlockContext(block.Header(), leth.blockchain, nil)
		statedb.Prepare(tx.Hash(), block.Hash(), idx)
		if idx == txIndex {
			return msg, context, statedb, statedb, nil, nil
		}
		// Not yet the searched for transaction, execute on top of the current state
		vmenv := vm.NewEVM(context, txContext, statedb, privateStateDb, leth.blockchain.Config(), vm.Config{})
		if _, err := core.ApplyMessage(vmenv, msg, new(core.GasPool).AddGas(tx.Gas())); err != nil {
			return nil, vm.BlockContext{}, nil, nil, nil, fmt.Errorf("transaction %#x failed: %v", tx.Hash(), err)
		}
		// Ensure any modifications are committed to the state
		// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect
		statedb.Finalise(vmenv.ChainConfig().IsEIP158(block.Number()))
	}
	return nil, vm.BlockContext{}, nil, nil, nil, fmt.Errorf("transaction index %d out of range for block %#x", txIndex, block.Hash())
}
