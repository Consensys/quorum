package vm

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/multitenancy"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/assert"
)

func TestAffectedMode_Update_whenTypical(t *testing.T) {
	testObject := ModeUnknown
	authorizedReads := []bool{true, false}
	authorizedWrites := []bool{true, false}
	for _, authorizedRead := range authorizedReads {
		for _, authorizedWrite := range authorizedWrites {
			actual := testObject.Update(authorizedRead, authorizedWrite)

			assert.True(t, actual.Has(ModeUpdated))
			assert.Equal(t, authorizedRead, actual.Has(ModeRead))
			assert.Equal(t, authorizedWrite, actual.Has(ModeWrite))
			assert.False(t, testObject.Has(ModeUpdated))
		}
	}
}

func TestCall_shouldReturnErrorNotAuthorized_whenNoCodeAndSupportsMultitenancy(t *testing.T) {
	address := common.BytesToAddress([]byte("contract"))

	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	statedb.CreateAccount(address)
	statedb.SetCode(address, []byte{})
	statedb.SetState(address, common.Hash{}, common.BytesToHash([]byte{}))
	statedb.Finalise(true) // Push the state into the "original" slot

	vmctx := Context{
		CanTransfer: func(StateDB, common.Address, *big.Int) bool { return true },
		Transfer:    func(StateDB, common.Address, common.Address, *big.Int) {},
	}
	vmenv := NewEVM(vmctx, statedb, statedb, params.QuorumTestChainConfig, Config{ExtraEips: []int{2200}})
	vmenv.SupportsMultitenancy = true
	vmenv.AuthorizeMessageCallFunc = func(contractAddress common.Address) (bool, bool, error) {
		return true, true, nil
	}

	_, _, err := vmenv.Call(AccountRef(common.Address{}), address, nil, math.MaxUint64, new(big.Int))

	assert.Error(t, err, multitenancy.ErrNotAuthorized)
}
