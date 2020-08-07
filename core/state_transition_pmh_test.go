package core

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/private/engine"
	testifyassert "github.com/stretchr/testify/assert"
)

type stubPmhStateTransition struct {
	snapshot int
}

func (s *stubPmhStateTransition) SetTxPrivacyMetadata(pm *types.PrivacyMetadata) {
}

func (s *stubPmhStateTransition) IsPrivacyEnhancementsEnabled() bool {
	return true
}

func (s *stubPmhStateTransition) RevertToSnapshot(val int) {
	s.snapshot = val
}

func (s *stubPmhStateTransition) GetStatePrivacyMetadata(addr common.Address) (*state.PrivacyMetadata, error) {
	return &state.PrivacyMetadata{PrivacyFlag: engine.PrivacyFlagStateValidation, CreationTxHash: common.EncryptedPayloadHash{1}}, nil
}

func (s *stubPmhStateTransition) CalculateMerkleRoot() (common.Hash, error) {
	return common.Hash{}, fmt.Errorf("Unable to calculate MerkleRoot")
}

func (s *stubPmhStateTransition) AffectedContracts() []common.Address {
	return make([]common.Address, 0)
}

func TestPrivateMessageContextVerify_WithMerkleRootCreationError(t *testing.T) {
	assert := testifyassert.New(t)
	stateTransitionAPI := &stubPmhStateTransition{}

	pmc := newPMH(stateTransitionAPI)
	pmc.receivedPrivacyMetadata = &engine.ExtraMetadata{ACMerkleRoot: common.Hash{1}, PrivacyFlag: engine.PrivacyFlagStateValidation}
	pmc.snapshot = 10
	exitEarly, err := pmc.verify(nil)

	assert.Error(err, "verify must return an error due to the MerkleRoot calculation error")
	assert.Equal(pmc.snapshot, stateTransitionAPI.snapshot, "Revert should have been called")
	assert.True(exitEarly, "Exit early should be true")
}
