package core

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/mps"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/private"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//Tests DefaultState, StatePSI, CommitAndWrite
func TestLegacyPrivateStateCreated(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockptm.EXPECT().Receive(gomock.Not(common.EncryptedPayloadHash{})).Return("", []string{"psi1", "psi2"}, common.FromHex(testCode), nil, nil).AnyTimes()

	blocks, blockmap, blockchain := buildTestChain(2, params.QuorumTestChainConfig)

	for _, block := range blocks {
		parent := blockmap[block.ParentHash()]
		statedb, _ := state.New(parent.Root(), blockchain.StateCache(), nil)
		privateStateRepo, _ := blockchain.PrivateStateManager().StateRepository(parent.Root())

		_, privateReceipts, _, _, _ := blockchain.Processor().Process(block, statedb, privateStateRepo, vm.Config{})

		for _, privateReceipt := range privateReceipts {
			expectedContractAddress := privateReceipt.ContractAddress

			assert.False(t, privateStateRepo.IsMPS())
			privateState, _ := privateStateRepo.DefaultState()
			assert.True(t, privateState.Exist(expectedContractAddress))
			assert.NotEqual(t, privateState.GetCodeSize(expectedContractAddress), 0)
			defaultPrivateState, _ := privateStateRepo.StatePSI(types.DefaultPrivateStateIdentifier)
			assert.True(t, defaultPrivateState.Exist(expectedContractAddress))
			assert.NotEqual(t, defaultPrivateState.GetCodeSize(expectedContractAddress), 0)
			_, err := privateStateRepo.StatePSI(types.PrivateStateIdentifier("empty"))
			assert.Error(t, err, "only the 'private' psi is supported by the default private state manager")

		}
		//CommitAndWrite to db
		privateStateRepo.CommitAndWrite(false, block)

		for _, privateReceipt := range privateReceipts {
			expectedContractAddress := privateReceipt.ContractAddress
			latestBlockRoot := block.Root()
			//contract exists on default state
			_, privDb, _ := blockchain.StateAtPSI(latestBlockRoot, types.DefaultPrivateStateIdentifier)
			assert.True(t, privDb.Exist(expectedContractAddress))
			assert.NotEqual(t, privDb.GetCodeSize(expectedContractAddress), 0)
			//legacy psm doesnt have concept of emptystate
			_, _, err := blockchain.StateAtPSI(latestBlockRoot, types.ToPrivateStateIdentifier("empty"))
			assert.Error(t, err, "only the 'private' psi is supported by the default private state manager")
			//legacy psm doesnt support other private states
			_, _, err = blockchain.StateAtPSI(latestBlockRoot, types.ToPrivateStateIdentifier("other"))
			assert.Error(t, err, "only the 'private' psi is supported by the default private state manager")
		}
	}
}

func TestDefaultResolver(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockptm := private.NewMockPrivateTransactionManager(mockCtrl)

	saved := private.P
	defer func() {
		private.P = saved
	}()
	private.P = mockptm

	mockptm.EXPECT().Receive(gomock.Not(common.EncryptedPayloadHash{})).Return("", []string{}, common.FromHex(testCode), nil, nil).AnyTimes()
	mockptm.EXPECT().Receive(common.EncryptedPayloadHash{}).Return("", []string{}, common.EncryptedPayloadHash{}.Bytes(), nil, nil).AnyTimes()

	_, _, blockchain := buildTestChain(1, params.QuorumTestChainConfig)

	mpsm := newDefaultPrivateStateManager(blockchain.db)

	psm1, _ := mpsm.ResolveForManagedParty("TEST")
	assert.Equal(t, psm1, mps.DefaultPrivateStateMetadata)

	ctx := rpc.WithPrivateStateIdentifier(context.Background(), types.DefaultPrivateStateIdentifier)
	psm1, _ = mpsm.ResolveForUserContext(ctx)
	assert.Equal(t, psm1, &mps.PrivateStateMetadata{ID: "private", Type: mps.Resident})
	psm1, _ = mpsm.ResolveForUserContext(context.Background())
	assert.Equal(t, psm1, &mps.PrivateStateMetadata{ID: "private", Type: mps.Resident})

	assert.Equal(t, mpsm.PSIs(), []types.PrivateStateIdentifier{types.DefaultPrivateStateIdentifier})
}
