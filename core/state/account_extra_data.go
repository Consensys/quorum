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

package state

import (
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/private/engine"
	"github.com/ethereum/go-ethereum/rlp"
)

// Quorum
// AccountExtraData is to contain extra data that supplements existing Account data.
// It is also maintained in a trie to support rollback.
// Note:
// - update copy() method
// - update DecodeRLP and EncodeRLP when adding new field
type AccountExtraData struct {
	// for privacy enhancements
	PrivacyMetadata *PrivacyMetadata
	// list of public keys managed by the corresponding Tessera.
	// This is for multitenancy
	ManagedParties []string
}

func (qmd *AccountExtraData) DecodeRLP(stream *rlp.Stream) error {
	var dataRLP struct {
		// from state.PrivacyMetadata, this is required to support
		// backward compatibility with RLP-encoded state.PrivacyMetadata.
		// Refer to rlp/doc.go for decoding rules.
		CreationTxHash *common.EncryptedPayloadHash `rlp:"nil"`
		// from state.PrivacyMetadata, this is required to support
		// backward compatibility with RLP-encoded state.PrivacyMetadata.
		// Refer to rlp/doc.go for decoding rules.
		PrivacyFlag *engine.PrivacyFlagType `rlp:"nil"`

		Rest []rlp.RawValue `rlp:"tail"` // to maintain forward compatibility
	}
	if err := stream.Decode(&dataRLP); err != nil {
		return err
	}
	if dataRLP.CreationTxHash != nil && dataRLP.PrivacyFlag != nil {
		qmd.PrivacyMetadata = &PrivacyMetadata{
			CreationTxHash: *dataRLP.CreationTxHash,
			PrivacyFlag:    *dataRLP.PrivacyFlag,
		}
	}
	if len(dataRLP.Rest) > 0 {
		var managedParties []string
		if err := rlp.DecodeBytes(dataRLP.Rest[0], &managedParties); err != nil {
			return fmt.Errorf("fail to decode managedParties with error %v", err)
		}
		// As RLP encodes empty slice or nil slice as an empty string (192)
		// we won't be able to determine when decoding. So we use pragmatic approach
		// to default to nil value. Downstream usage would deal with it easier.
		if len(managedParties) == 0 {
			qmd.ManagedParties = nil
		} else {
			qmd.ManagedParties = managedParties
		}
	}
	return nil
}

func (qmd *AccountExtraData) EncodeRLP(writer io.Writer) error {
	var (
		hash *common.EncryptedPayloadHash
		flag *engine.PrivacyFlagType
	)
	if qmd.PrivacyMetadata != nil {
		hash = &qmd.PrivacyMetadata.CreationTxHash
		flag = &qmd.PrivacyMetadata.PrivacyFlag
	}
	return rlp.Encode(writer, struct {
		CreationTxHash *common.EncryptedPayloadHash `rlp:"nil"`
		PrivacyFlag    *engine.PrivacyFlagType      `rlp:"nil"`
		ManagedParties []string
	}{
		CreationTxHash: hash,
		PrivacyFlag:    flag,
		ManagedParties: qmd.ManagedParties,
	})
}

func (qmd *AccountExtraData) copy() *AccountExtraData {
	if qmd == nil {
		return nil
	}
	var copyPM *PrivacyMetadata
	if qmd.PrivacyMetadata != nil {
		copyPM = &PrivacyMetadata{
			CreationTxHash: qmd.PrivacyMetadata.CreationTxHash,
			PrivacyFlag:    qmd.PrivacyMetadata.PrivacyFlag,
		}
	}
	copyManagedParties := make([]string, len(qmd.ManagedParties))
	copy(copyManagedParties, qmd.ManagedParties)
	return &AccountExtraData{
		PrivacyMetadata: copyPM,
		ManagedParties:  copyManagedParties,
	}
}

// attached to every private contract account
type PrivacyMetadata struct {
	CreationTxHash common.EncryptedPayloadHash `json:"creationTxHash"`
	PrivacyFlag    engine.PrivacyFlagType      `json:"privacyFlag"`
}

// Quorum
// privacyMetadataRLP struct is to make sure
// field order is preserved regardless changes in the PrivacyMetadata and its internal
//
// Edit this struct with care to make sure forward and backward compatibility
type privacyMetadataRLP struct {
	CreationTxHash common.EncryptedPayloadHash
	PrivacyFlag    engine.PrivacyFlagType

	Rest []rlp.RawValue `rlp:"tail"` // to maintain forward compatibility
}

func (p *PrivacyMetadata) DecodeRLP(stream *rlp.Stream) error {
	var dataRLP privacyMetadataRLP
	if err := stream.Decode(&dataRLP); err != nil {
		return err
	}
	p.CreationTxHash = dataRLP.CreationTxHash
	p.PrivacyFlag = dataRLP.PrivacyFlag
	return nil
}

func (p *PrivacyMetadata) EncodeRLP(writer io.Writer) error {
	return rlp.Encode(writer, privacyMetadataRLP{
		CreationTxHash: p.CreationTxHash,
		PrivacyFlag:    p.PrivacyFlag,
	})
}

func NewStatePrivacyMetadata(creationTxHash common.EncryptedPayloadHash, privacyFlag engine.PrivacyFlagType) *PrivacyMetadata {
	return &PrivacyMetadata{
		CreationTxHash: creationTxHash,
		PrivacyFlag:    privacyFlag,
	}
}
