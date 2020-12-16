// Copyright 2019 The go-ethereum Authors
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

package state

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
)

func BenchmarkCutOriginal(b *testing.B) {
	value := common.HexToHash("0x01")
	for i := 0; i < b.N; i++ {
		bytes.TrimLeft(value[:], "\x00")
	}
}

func BenchmarkCutsetterFn(b *testing.B) {
	value := common.HexToHash("0x01")
	cutSetFn := func(r rune) bool {
		return int32(r) == int32(0)
	}
	for i := 0; i < b.N; i++ {
		bytes.TrimLeftFunc(value[:], cutSetFn)
	}
}

func BenchmarkCutCustomTrim(b *testing.B) {
	value := common.HexToHash("0x01")
	for i := 0; i < b.N; i++ {
		common.TrimLeftZeroes(value[:])
	}
}

func xTestFuzzCutter(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for {
		v := make([]byte, 20)
		zeroes := rand.Intn(21)
		rand.Read(v[zeroes:])
		exp := bytes.TrimLeft(v[:], "\x00")
		got := common.TrimLeftZeroes(v)
		if !bytes.Equal(exp, got) {

			fmt.Printf("Input %x\n", v)
			fmt.Printf("Exp %x\n", exp)
			fmt.Printf("Got %x\n", got)
			t.Fatalf("Error")
		}
		//break
	}
}

type privacyMetadataOld struct {
	CreationTxHash common.EncryptedPayloadHash
	PrivacyFlag    engine.PrivacyFlagType
}

// privacyMetadataToBytes is the utility function under test from previous implementation
func privacyMetadataToBytes(pm *privacyMetadataOld) ([]byte, error) {
	return rlp.EncodeToBytes(pm)
}

// bytesToPrivacyMetadata is the utility function under test from previous implementation
func bytesToPrivacyMetadata(b []byte) (*privacyMetadataOld, error) {
	var data *privacyMetadataOld
	if err := rlp.DecodeBytes(b, &data); err != nil {
		return nil, fmt.Errorf("unable to decode privacy metadata. Cause: %v", err)
	}
	return data, nil
}

func TestRLP_PrivacyMetadata_DecodeBackwardCompatibility(t *testing.T) {
	existingPM := &privacyMetadataOld{
		CreationTxHash: common.BytesToEncryptedPayloadHash([]byte("arbitrary-hash")),
		PrivacyFlag:    engine.PrivacyFlagStateValidation,
	}
	existing, err := privacyMetadataToBytes(existingPM)
	assert.NoError(t, err)

	var actual PrivacyMetadata
	err = rlp.DecodeBytes(existing, &actual)

	assert.NoError(t, err, "Must decode PrivacyMetadata successfully")
	assert.Equal(t, existingPM.CreationTxHash, actual.CreationTxHash)
	assert.Equal(t, existingPM.PrivacyFlag, actual.PrivacyFlag)
}

func TestRLP_PrivacyMetadata_DecodeForwardCompatibility(t *testing.T) {
	pm := &PrivacyMetadata{
		CreationTxHash: common.BytesToEncryptedPayloadHash([]byte("arbitrary-hash")),
		PrivacyFlag:    engine.PrivacyFlagStateValidation,
	}
	existing, err := rlp.EncodeToBytes(pm)
	assert.NoError(t, err)

	var actual *privacyMetadataOld
	actual, err = bytesToPrivacyMetadata(existing)

	assert.NoError(t, err, "Must encode PrivacyMetadata successfully")
	assert.Equal(t, pm.CreationTxHash, actual.CreationTxHash)
	assert.Equal(t, pm.PrivacyFlag, actual.PrivacyFlag)
}

// From initial privacy enhancements, the privacy metadata is RLP encoded
// we now wrap PrivacyMetadata in a more generic struct. This test is to make sure
// we support backward compatibility.
func TestRLP_AccountExtraData_BackwardCompatibility(t *testing.T) {
	// prepare existing RLP bytes
	arbitraryExistingMetadata := &PrivacyMetadata{
		CreationTxHash: common.BytesToEncryptedPayloadHash([]byte("arbitrary-existing-privacy-metadata-creation-hash")),
		PrivacyFlag:    engine.PrivacyFlagPartyProtection,
	}
	existing, err := rlp.EncodeToBytes(arbitraryExistingMetadata)
	assert.NoError(t, err)

	// now try to decode with the new struct
	var actual AccountExtraData
	err = rlp.DecodeBytes(existing, &actual)

	assert.NoError(t, err, "Must decode successfully")
	assert.Equal(t, arbitraryExistingMetadata.CreationTxHash, actual.PrivacyMetadata.CreationTxHash)
	assert.Equal(t, arbitraryExistingMetadata.PrivacyFlag, actual.PrivacyMetadata.PrivacyFlag)
}

func TestRLP_AccountExtraData_withField_ManagedParties(t *testing.T) {
	// prepare existing RLP bytes
	arbitraryExtraData := &AccountExtraData{
		PrivacyMetadata: &PrivacyMetadata{
			CreationTxHash: common.BytesToEncryptedPayloadHash([]byte("arbitrary-existing-privacy-metadata-creation-hash")),
			PrivacyFlag:    engine.PrivacyFlagPartyProtection,
		},
		ManagedParties: []string{"Arbitrary PK1", "Arbitrary PK2"},
	}
	existing, err := rlp.EncodeToBytes(arbitraryExtraData)
	assert.NoError(t, err)

	// now try to decode with the new struct
	var actual AccountExtraData
	err = rlp.DecodeBytes(existing, &actual)

	assert.NoError(t, err, "Must decode successfully")
	assert.Equal(t, arbitraryExtraData.PrivacyMetadata.CreationTxHash, actual.PrivacyMetadata.CreationTxHash)
	assert.Equal(t, arbitraryExtraData.PrivacyMetadata.PrivacyFlag, actual.PrivacyMetadata.PrivacyFlag)
	assert.Equal(t, arbitraryExtraData.ManagedParties, actual.ManagedParties)
}

func TestRLP_AccountExtraData_whenTypical(t *testing.T) {
	expected := AccountExtraData{
		PrivacyMetadata: &PrivacyMetadata{
			CreationTxHash: common.BytesToEncryptedPayloadHash([]byte("arbitrary-payload-hash")),
			PrivacyFlag:    engine.PrivacyFlagPartyProtection,
		},
		ManagedParties: []string{"XYZ"},
	}

	data, err := rlp.EncodeToBytes(&expected)
	assert.NoError(t, err)

	var actual AccountExtraData
	assert.NoError(t, rlp.DecodeBytes(data, &actual))
	assert.Equal(t, expected.PrivacyMetadata.CreationTxHash, actual.PrivacyMetadata.CreationTxHash)
	assert.Equal(t, expected.PrivacyMetadata.PrivacyFlag, actual.PrivacyMetadata.PrivacyFlag)
	assert.Equal(t, expected.ManagedParties, actual.ManagedParties)
}

func TestRLP_AccountExtraData_whenHavingPrivacyMetadataOnly(t *testing.T) {
	expected := AccountExtraData{
		PrivacyMetadata: &PrivacyMetadata{
			CreationTxHash: common.BytesToEncryptedPayloadHash([]byte("arbitrary-payload-hash")),
			PrivacyFlag:    engine.PrivacyFlagPartyProtection,
		},
	}

	data, err := rlp.EncodeToBytes(&expected)
	assert.NoError(t, err)

	var actual AccountExtraData
	assert.NoError(t, rlp.DecodeBytes(data, &actual))
	assert.Equal(t, expected.PrivacyMetadata.CreationTxHash, actual.PrivacyMetadata.CreationTxHash)
	assert.Equal(t, expected.PrivacyMetadata.PrivacyFlag, actual.PrivacyMetadata.PrivacyFlag)
}

func TestRLP_AccountExtraData_whenHavingNilManagedParties(t *testing.T) {
	expected := AccountExtraData{
		PrivacyMetadata: nil,
		ManagedParties:  nil,
	}

	data, err := rlp.EncodeToBytes(&expected)
	assert.NoError(t, err)

	var actual AccountExtraData
	assert.NoError(t, rlp.DecodeBytes(data, &actual))
	assert.Nil(t, actual.ManagedParties)
	assert.Nil(t, actual.PrivacyMetadata)
}

func TestRLP_AccountExtraData_whenHavingEmptyManagedParties(t *testing.T) {
	expected := AccountExtraData{
		PrivacyMetadata: nil,
		ManagedParties:  []string{},
	}

	data, err := rlp.EncodeToBytes(&expected)
	assert.NoError(t, err)

	var actual AccountExtraData
	assert.NoError(t, rlp.DecodeBytes(data, &actual))
	assert.Nil(t, actual.ManagedParties)
	assert.Nil(t, actual.PrivacyMetadata)
}

func TestCopy_whenNil(t *testing.T) {
	var testObj *AccountExtraData = nil

	assert.Nil(t, testObj.copy())
}
