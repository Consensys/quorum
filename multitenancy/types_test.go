package multitenancy

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	arbitraryAddress = common.StringToAddress("0xArbitraryContractAddress")
	arbitraryEOA     = common.StringToAddress("0xArbitraryEOA")
)

func TestToContractSecurityAttribute_forPublic(t *testing.T) {
	arbitraryContractParties := &ContractIndexItem{
		CreatorAddress: arbitraryEOA,
		IsPrivate:      false,
		Parties:        nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIndexReader := NewMockContractIndexReader(ctrl)
	mockIndexReader.EXPECT().
		ReadIndex(gomock.Eq(arbitraryAddress)).
		Return(arbitraryContractParties, nil)

	attr, err := ToContractSecurityAttribute(mockIndexReader, arbitraryAddress)

	assert.NoError(t, err)
	assert.Equal(t, VisibilityPublic, attr.Visibility)
	assert.Equal(t, ActionRead, attr.Action)
	assert.Equal(t, arbitraryEOA, attr.From)
	assert.Equal(t, arbitraryEOA, attr.To)
	assert.Empty(t, attr.Parties)
	assert.Empty(t, attr.PrivateFrom)
}

func TestToContractSecurityAttribute_forPrivate(t *testing.T) {
	arbitraryContractParties := &ContractIndexItem{
		CreatorAddress: arbitraryEOA,
		IsPrivate:      true,
		Parties:        []string{"arbitrary participant"},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIndexReader := NewMockContractIndexReader(ctrl)
	mockIndexReader.EXPECT().
		ReadIndex(gomock.Eq(arbitraryAddress)).
		Return(arbitraryContractParties, nil)

	attr, err := ToContractSecurityAttribute(mockIndexReader, arbitraryAddress)

	assert.NoError(t, err)
	assert.Equal(t, VisibilityPrivate, attr.Visibility)
	assert.Equal(t, ActionRead, attr.Action)
	assert.Equal(t, arbitraryEOA, attr.From)
	assert.Equal(t, arbitraryEOA, attr.To)
	assert.Equal(t, arbitraryContractParties.Parties, attr.Parties)
	assert.Empty(t, attr.PrivateFrom)
}

func TestToContractSecurityAttribute_forPrivate_nonPartyNode(t *testing.T) {
	arbitraryContractParties := &ContractIndexItem{
		CreatorAddress: arbitraryEOA,
		IsPrivate:      true,
		Parties:        []string{},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIndexReader := NewMockContractIndexReader(ctrl)
	mockIndexReader.EXPECT().
		ReadIndex(gomock.Eq(arbitraryAddress)).
		Return(arbitraryContractParties, nil)

	attr, err := ToContractSecurityAttribute(mockIndexReader, arbitraryAddress)

	assert.NoError(t, err)
	assert.Equal(t, VisibilityPrivate, attr.Visibility)
	assert.Equal(t, ActionRead, attr.Action)
	assert.Equal(t, arbitraryEOA, attr.From)
	assert.Equal(t, arbitraryEOA, attr.To)
	assert.Equal(t, arbitraryContractParties.Parties, attr.Parties)
	assert.Empty(t, attr.PrivateFrom)
}
