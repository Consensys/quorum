package core

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/private"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"

	testifyassert "github.com/stretchr/testify/assert"
)

func verifyGasPoolCalculation(t *testing.T, pm private.PrivateTransactionManager) {
	assert := testifyassert.New(t)
	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = pm

	txGasLimit := uint64(100000)
	gasPool := new(GasPool).AddGas(200000)
	// this payload would give us 25288 intrinsic gas
	arbitraryEncryptedPayload := "4ab80888354582b92ab442a317828386e4bf21ea4a38d1a9183fbb715f199475269d7686939017f4a6b28310d5003ebd8e012eade530b79e157657ce8dd9692a"
	expectedGasPool := new(GasPool).AddGas(174712) // only intrinsic gas is deducted

	db := ethdb.NewMemDatabase()
	privateState, _ := state.New(common.Hash{}, state.NewDatabase(db))
	publicState, _ := state.New(common.Hash{}, state.NewDatabase(db))
	msg := privateCallMsg{
		callmsg: callmsg{
			addr:     common.Address{2},
			to:       &common.Address{},
			value:    new(big.Int),
			gas:      txGasLimit,
			gasPrice: big.NewInt(0),
			data:     common.Hex2Bytes(arbitraryEncryptedPayload),
		},
	}
	ctx := NewEVMContext(msg, &dualStateTestHeader, nil, &common.Address{})
	evm := vm.NewEVM(ctx, publicState, privateState, params.QuorumTestChainConfig, vm.Config{})
	arbitraryBalance := big.NewInt(100000000)
	publicState.SetBalance(evm.Coinbase, arbitraryBalance)
	publicState.SetBalance(msg.From(), arbitraryBalance)

	testObject := NewStateTransition(evm, msg, gasPool)

	_, _, failed, err := testObject.TransitionDb()

	assert.NoError(err)
	assert.False(failed)

	assert.Equal(new(big.Int).SetUint64(expectedGasPool.Gas()), new(big.Int).SetUint64(gasPool.Gas()), "gas pool must be calculated correctly")
	assert.Equal(arbitraryBalance, publicState.GetBalance(evm.Coinbase), "balance must not be changed")
	assert.Equal(arbitraryBalance, publicState.GetBalance(msg.From()), "balance must not be changed")
}

func TestStateTransition_TransitionDb_GasPoolCalculation_whenNonPartyNodeProcessingPrivateTransactions(t *testing.T) {
	stubPTM := &StubPrivateTransactionManager{
		responses: map[string][]interface{}{
			"Receive": {
				[]byte{},
				nil,
			},
		},
	}
	verifyGasPoolCalculation(t, stubPTM)
}

func TestStateTransition_TransitionDb_GasPoolCalculation_whenPartyNodeProcessingPrivateTransactions(t *testing.T) {
	stubPTM := &StubPrivateTransactionManager{
		responses: map[string][]interface{}{
			"Receive": {
				common.Hex2Bytes("600a6000526001601ff300"),
				nil,
			},
		},
	}
	verifyGasPoolCalculation(t, stubPTM)
}

type privateCallMsg struct {
	callmsg
}

func (pm privateCallMsg) IsPrivate() bool { return true }

type StubPrivateTransactionManager struct {
	responses map[string][]interface{}
}

func (spm *StubPrivateTransactionManager) Send(data []byte, from string, to []string) ([]byte, error) {
	return nil, fmt.Errorf("to be implemented")
}

func (spm *StubPrivateTransactionManager) SendSignedTx(data []byte, to []string) ([]byte, error) {
	return nil, fmt.Errorf("to be implemented")
}

func (spm *StubPrivateTransactionManager) Receive(data []byte) ([]byte, error) {
	res := spm.responses["Receive"]
	if err, ok := res[1].(error); ok {
		return nil, err
	}
	if ret, ok := res[0].([]byte); ok {
		return ret, nil
	}
	return nil, nil
}
