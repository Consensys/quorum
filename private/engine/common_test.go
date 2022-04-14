package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivacyFlag_whenTypical(t *testing.T) {
	assert := assert.New(t)

	flag := PrivacyFlagPartyProtection | PrivacyFlagStateValidation

	assert.True(flag.Has(PrivacyFlagStateValidation))
}

func TestPrivacyFlag_whenCheckingMultipleFlags(t *testing.T) {
	assert := assert.New(t)

	flag := PrivacyFlagPartyProtection

	assert.False(flag.Has(PrivacyFlagStateValidation | PrivacyFlagPartyProtection))
}

func TestPrivacyFlag_whenCheckingMultipleFlagsArray(t *testing.T) {
	assert := assert.New(t)

	flag := PrivacyFlagStateValidation | PrivacyFlagPartyProtection

	assert.True(flag.HasAll(PrivacyFlagStateValidation, PrivacyFlagPartyProtection))
}

func TestPrivacyFlag_whenCheckingStandardPrivateFlag(t *testing.T) {
	assert := assert.New(t)

	flag := PrivacyFlagStandardPrivate

	assert.True(flag.IsStandardPrivate())
}

func TestPrivacyFlag_whenCheckingNotStandardPrivateFlag(t *testing.T) {
	assert := assert.New(t)

	flag := PrivacyFlagPartyProtection

	assert.True(flag.IsNotStandardPrivate())
}

func TestPrivacyFlag_whenPrivateStateValidation(t *testing.T) {
	assert := assert.New(t)

	t.Logf("PrivateFlagStateValidation: %d", PrivacyFlagStateValidation)

	assert.True(PrivacyFlagStateValidation.Has(PrivacyFlagPartyProtection), "State Validation must have party protection by default")
}

func TestPrivacyFlag_whenMandatoryRecipients(t *testing.T) {
	assert := assert.New(t)

	flag := PrivacyFlagMandatoryRecipients

	assert.NoError(flag.Validate())
	assert.True(flag.Has(PrivacyFlagMandatoryRecipients))
	assert.True(PrivacyFlagStateValidation.Has(flag))

}

func TestPrivacyFlagType_Validate_whenSuccess(t *testing.T) {
	assert := assert.New(t)

	flag := PrivacyFlagStateValidation

	assert.NoError(flag.Validate())
}

func TestPrivacyFlagType_Validate_whenFailure(t *testing.T) {
	assert := assert.New(t)

	flag := PrivacyFlagType(4)

	assert.Error(flag.Validate())
}

func TestFeatureSet_HasFeature(t *testing.T) {
	assert := assert.New(t)

	featureSet := NewFeatureSet(PrivacyEnhancements, MultiTenancy, MultiplePrivateStates, MandatoryRecipients)
	assert.True(featureSet.HasFeature(MandatoryRecipients))
	assert.True(featureSet.HasFeature(MultiplePrivateStates))
}
