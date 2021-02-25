package core

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/psmr"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/private"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//Tests GetDefaultState, GetPrivateState, CommitAndWrite
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

	blocks, blockmap, blockchain := buildTestChain(2, params.QuorumTestChainConfig, &psmr.DefaultPrivateStateMetadataResolver{})

	for _, block := range blocks {
		parent := blockmap[block.ParentHash()]
		statedb, _ := state.New(parent.Root(), blockchain.stateCache, nil)
		psManager, _ := NewPrivateStateManager(blockchain, parent.Root())

		_, privateReceipts, _, _, _ := blockchain.processor.Process(block, statedb, psManager, vm.Config{})

		for _, privateReceipt := range privateReceipts {
			expectedContractAddress := privateReceipt.ContractAddress

			assert.False(t, psManager.IsMPS())
			privateState, _ := psManager.GetDefaultState()
			assert.True(t, privateState.Exist(expectedContractAddress))
			assert.NotEqual(t, privateState.GetCodeSize(expectedContractAddress), 0)
			defaultPrivateState, _ := psManager.GetPrivateState(types.DefaultPrivateStateIdentifier)
			assert.True(t, defaultPrivateState.Exist(expectedContractAddress))
			assert.NotEqual(t, defaultPrivateState.GetCodeSize(expectedContractAddress), 0)
			_, err := psManager.GetPrivateState(types.PrivateStateIdentifier("empty"))
			assert.Error(t, err, "only the 'private' psi is supported by the legacy private state manager")

		}
		//CommitAndWrite to db
		psManager.CommitAndWrite(block)

		for _, privateReceipt := range privateReceipts {
			expectedContractAddress := privateReceipt.ContractAddress
			latestBlockRoot := block.Root()
			//contract exists on default state
			_, privDb, _ := blockchain.StateAtPSI(latestBlockRoot, types.DefaultPrivateStateIdentifier)
			assert.True(t, privDb.Exist(expectedContractAddress))
			assert.NotEqual(t, privDb.GetCodeSize(expectedContractAddress), 0)
			//legacy psm doesnt have concept of emptystate
			_, _, err := blockchain.StateAtPSI(latestBlockRoot, types.ToPrivateStateIdentifier("empty"))
			assert.Error(t, err, "only the 'private' psi is supported by the legacy private state manager")
			//legacy psm doesnt support other private states
			_, _, err = blockchain.StateAtPSI(latestBlockRoot, types.ToPrivateStateIdentifier("other"))
			assert.Error(t, err, "only the 'private' psi is supported by the legacy private state manager")
		}
	}
}
